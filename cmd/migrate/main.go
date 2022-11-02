/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/gopackages"
	"github.com/dave/dst/dstutil"
	"github.com/sirupsen/logrus"
	"golang.org/x/tools/go/packages"
	"k8s.io/apimachinery/pkg/util/sets"
)

// change factories
// change types all the way down

// key := MetaNamespaceKeyFunc -> key := MetaClusterNamespaceKeyFunc
// ns, name := SplitMetaNamespaceKey -> clusterName, ns, name := SplitMetaClusterNamespaceKey
// lister.Get(key) -> lister.Cluster(clusterName).Get(key)
// lister.List(labelSelector) -> lister.Cluster(clusterName).List(labelSelector)

// lister.Get(clusters.ToClusterAwareKey(clusterName, name)) -> lister.Cluster(clusterName).Get(name) // kcp specific, only we have the former
// indexer.ByIndex(byWorkspace, clusterName) -> lister.Cluster(clusterName).List(labels.Everything()) // kcp specific, only we have the former

type options struct {
	packageNames []string
}

func newOptions() *options {
	o := &options{}
	return o
}

func (o *options) addFlags(fs *flag.FlagSet) {
}

func (o *options) complete(args []string) error {
	o.packageNames = args
	return nil
}

func (o *options) validate() error {
	if len(o.packageNames) == 0 {
		return errors.New("packages to scan are required arguments")
	}
	return nil
}

func main() {
	o := newOptions()
	o.addFlags(flag.CommandLine)
	flag.Parse()
	if err := o.complete(flag.Args()); err != nil {
		logrus.WithError(err).Fatal("could not complete options")
	}
	if err := o.validate(); err != nil {
		logrus.WithError(err).Fatal("invalid options")
	}
	if err := rewrite(o.packageNames); err != nil {
		logrus.WithError(err).Fatal("failed to re-write source")
	}
}

func rewrite(packageNames []string) error {
	pkgs, err := decorator.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedTypes | packages.NeedTypesSizes |
			packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedModule,
		Tests: true,
	}, packageNames...)
	if err != nil {
		return fmt.Errorf("failed to load source: %w", err)
	}
	for _, pkg := range pkgs {
		restorer := decorator.NewRestorerWithImports(pkg.PkgPath, gopackages.New(""))
		fileRestorer := restorer.FileRestorer()
		for i, file := range pkg.Syntax {
			filePath := pkg.CompiledGoFiles[i]
			dir := pkg.Dir
			if pkg.Module != nil { // even though we ask for modules, if we run on a single file we don't get them
				dir = pkg.Module.Dir
			}
			relPath, err := filepath.Rel(dir, filePath)
			if err != nil {
				return fmt.Errorf("should not happen: could not find relative path to %s from %s", filePath, pkg.Dir)
			}

			var shouldSkip bool
			for _, dec := range file.Decs.Start.All() {
				if dec == "// +kcp-code-generator:skip" {
					shouldSkip = true
					break
				}
			}
			logrus.WithFields(logrus.Fields{"file": relPath}).Info("Considering file.")
			if shouldSkip {
				logrus.WithFields(logrus.Fields{"file": relPath}).Info("Skipping file.")
				continue
			}
			var updates int
			mutatedNode := dstutil.Apply(file, nil, func(cursor *dstutil.Cursor) bool {
				for _, update := range []func(pkg *decorator.Package, cursor *dstutil.Cursor, fileRestorer *decorator.FileRestorer) int{
					rewriteClientTypes,
					rewriteClusterInterfaceCall,
					//rewriteKeyFuncs,
					//rewriteKeySplits,
					//rewriteKcpKeySplits,
					rewriteListerGet,
				} {
					updates += update(pkg, cursor, fileRestorer)
				}
				return true
			})
			if updates > 0 {
				logrus.WithFields(logrus.Fields{"file": relPath, "updates": updates}).Info("Updating file.")
				previous, err := ioutil.ReadFile(filePath)
				if err != nil {
					return fmt.Errorf("failed to read contents of %s: %w", filePath, err)
				}
				f, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0666)
				if err != nil {
					return fmt.Errorf("failed to open %s for writing: %w", relPath, err)
				}
				if err := fileRestorer.Fprint(f, mutatedNode.(*dst.File)); err != nil {
					if err := ioutil.WriteFile(filePath, previous, 0666); err != nil {
						logrus.WithFields(logrus.Fields{"file": relPath}).WithError(err).Error("Failed to restore contents of file.")
					}
					return fmt.Errorf("failed to write %s: %w", relPath, err)
				}
			}
		}
	}
	return nil
}

const (
	kcpCachePath  = "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpCacheAlias = "kcpcache"
)

type rewriteRule struct {
	from, to    string
	nameMatcher *regexp.Regexp
	formatAlias func([]string) string
}

var kcpClientTypeRules = []rewriteRule{
	{
		from:        "github.com/kcp-dev/kcp/pkg/client/clientset/versioned",
		to:          "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/cluster",
		nameMatcher: regexp.MustCompile(`.*(Interface|Getter|Clientset|Client|Config)`),
		formatAlias: func(suffix []string) string {
			// [] -> "kcpclientset"
			// ["typed", "<group>", "<version>"] -> "<group><version>"
			if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
				return "kcpclientset"
			}
			if len(suffix) == 1 {
				return "kcp" + suffix[0] + "client"
			}
			return suffix[1] + suffix[2] + "client"
		},
	},
	{
		from:        "github.com/kcp-dev/kcp/pkg/client/informers/externalversions",
		to:          "github.com/kcp-dev/kcp/pkg/client/informers/externalversions",
		nameMatcher: regexp.MustCompile(`.*(Interface|Informer|Getter|Clientset|Config)`),
		formatAlias: func(suffix []string) string {
			// [] -> "kcpinformers"
			// ["<group>", "<version>"] -> "<group><version>"informers
			if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
				return "kcpinformers"
			}
			return suffix[0] + suffix[1] + "informers"
		},
	},
	{
		from:        "github.com/kcp-dev/kcp/pkg/client/listers",
		to:          "github.com/kcp-dev/kcp/pkg/client/listers",
		nameMatcher: regexp.MustCompile(`.*(Interface|Lister|Getter|Clientset|Config)`),
		formatAlias: func(suffix []string) string {
			// ["<group>", "<version>"] -> "kcp<group><version>"listers
			return suffix[0] + suffix[1] + "listers"
		},
	},
	//{
	//	from:        "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset",
	//	to:          "github.com/kcp-dev/client-go/apiextensions/clients/clientset/versioned",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Getter|Clientset|Client|Config)`),
	//	formatAlias: func(suffix []string) string {
	//		// [] -> "kcpclientset"
	//		// ["typed", "<group>", "<version>"] -> "<group><version>"
	//		if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
	//			return "kcpapiextensionsclientset"
	//		}
	//		if len(suffix) == 1 {
	//			return "kcp" + suffix[0] + "client"
	//		}
	//		return "kcp" + suffix[1] + suffix[2] + "client"
	//	},
	//},
	//{
	//	from:        "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions",
	//	to:          "github.com/kcp-dev/client-go/apiextensions/clients/informers",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Informer|Getter|Clientset|Config)`),
	//	formatAlias: func(suffix []string) string {
	//		// [] -> "kcpapiextensionsinformers"
	//		// ["<group>", "<version>"] -> "<group><version>"informers
	//		if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
	//			return "kcpapiextensionsinformers"
	//		}
	//		return "kcp" + suffix[0] + suffix[1] + "informers"
	//	},
	//},
	//{
	//	from:        "k8s.io/apiextensions-apiserver/pkg/client/listers",
	//	to:          "github.com/kcp-dev/client-go/apiextensions/clients/listers",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Lister|Getter|Clientset|Config)`),
	//	formatAlias: func(suffix []string) string {
	//		// ["<group>", "<version>"] -> "kcp<group><version>"listers
	//		return "kcp" + suffix[0] + suffix[1] + "listers"
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/kubernetes",
	//	to:          "github.com/kcp-dev/client-go/clients/clientset/versioned",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Getter|Clientset|Client|Config)`),
	//	formatAlias: func(suffix []string) string {
	//		// [] -> "kcpclientset"
	//		// ["typed", "<group>", "<version>"] -> "<group><version>"
	//		if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
	//			return "kcpkubernetesclientset"
	//		}
	//		if len(suffix) == 1 {
	//			return "kcp" + suffix[0] + "client"
	//		}
	//		return "kcp" + suffix[1] + suffix[2] + "client"
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/informers",
	//	to:          "github.com/kcp-dev/client-go/clients/informers",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Informer|Getter|Clientset|Config)`),
	//	formatAlias: func(suffix []string) string {
	//		// [] -> "kcpkubernetesinformers"
	//		// ["<group>", "<version>"] -> "<group><version>"informers
	//		if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
	//			return "kcpkubernetesinformers"
	//		}
	//		return "kcp" + suffix[0] + suffix[1] + "informers"
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/listers",
	//	to:          "github.com/kcp-dev/client-go/clients/listers",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Lister|Getter|Clientset|Config)`),
	//	formatAlias: func(suffix []string) string {
	//		// ["<group>", "<version>"] -> "kcp<group><version>"listers
	//		return "kcp" + suffix[0] + suffix[1] + "listers"
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/dynamic",
	//	to:          "github.com/kcp-dev/client-go/clients/dynamic",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Lister|Informer|Clientset|Config)`),
	//	formatAlias: func(suffix []string) string {
	//		// [] -> "kcpdynamic"
	//		// ["dynamiclister"] -> "kcpdynamiclister"
	//		if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
	//			return "kcpdynamic"
	//		}
	//		return "kcp" + suffix[0]
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/dynamic/fake",
	//	to:          "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/dynamic/fake",
	//	nameMatcher: regexp.MustCompile(`.*`),
	//	formatAlias: func(suffix []string) string {
	//		return "kcpfakedynamic"
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/metadata",
	//	to:          "github.com/kcp-dev/client-go/clients/metadata",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Lister|Informer|Clientset|Config)`),
	//	formatAlias: func(suffix []string) string {
	//		// [] -> "kcpmetadata"
	//		// ["metadatalister"] -> "kcpmetadatalister"
	//		if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
	//			return "kcpmetadata"
	//		}
	//		return "kcp" + suffix[0]
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/metadata/fake",
	//	to:          "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/metadata/fake",
	//	nameMatcher: regexp.MustCompile(`.*`),
	//	formatAlias: func(suffix []string) string {
	//		return "kcpfakemetadata"
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/discovery",
	//	to:          "github.com/kcp-dev/client-go/clients/discovery",
	//	nameMatcher: regexp.MustCompile(`.*(Interface|Clientset)`),
	//	formatAlias: func(suffix []string) string {
	//		// [] -> "kcpdiscovery"
	//		// ["discoverylister"] -> "kcpdiscoverylister"
	//		if len(suffix) == 0 || len(suffix) == 1 && suffix[0] == "" {
	//			return "kcpdiscovery"
	//		}
	//		return "kcp" + suffix[0]
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/discovery/fake",
	//	to:          "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/discovery/fake",
	//	nameMatcher: regexp.MustCompile(`.*`),
	//	formatAlias: func(suffix []string) string {
	//		return "kcpfakediscovery"
	//	},
	//},
	//{
	//	from:        "k8s.io/client-go/testing",
	//	to:          "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing",
	//	nameMatcher: regexp.MustCompile(`.*(Action|Reactor|Reaction|Fake|ObjectTracker|Client)`),
	//	formatAlias: func(suffix []string) string {
	//		return "kcptesting"
	//	},
	//},
}

// k8s.io/client-go/kubernetes.Interface -> github.com/kcp-dev/client-go/clients/clientset/versioned.ClusterInterface
// k8s.io/client-go/kubernetes/typed/<group>/<version> -> github.com/kcp-dev/client-go/clients/clientset/versioned/typed/<group>/<version>
//  - <group><version>Interface -> <group><version>ClusterInterface
//  - <kind>Getter -> <kind>ClusterGetter
//  - <kind>Interface -> <kind>ClusterInterface
// k8s.io/client-go/informers/<group>/<version> -> github.com/kcp-dev/client-go/clients/informers/<group>/<version>
//  - Interface -> ClusterInterface
//  - <kind>Informer -> <kind>ClusterInformer
// k8s.io/client-go/listers/<group>/<version> -> github.com/kcp-dev/client-go/clients/listers/<group>/<version>
//  - Interface -> ClusterInterface
//  - <kind>Lister -> <kind>ClusterLister
// k8s.io/client-go/dynamic/dynamic<lister,informer> -> github.com/kcp-dev/client-go/clients/dynamic/dynamic<lister,informer>
//  - Interface -> ClusterInterface
//  - <kind>Lister -> <kind>ClusterLister
// k8s.io/client-go/metadata/metadata<lister,informer> -> github.com/kcp-dev/client-go/clients/metadata/metadata<lister,informer>
//  - Interface -> ClusterInterface
//  - <kind>Lister -> <kind>ClusterLister
// k8s.io/client-go/discovery -> github.com/kcp-dev/client-go/clients/discovery
//  - Interface -> ClusterInterface
//  - Clientset -> ClusterClientset
func rewriteClientTypes(pkg *decorator.Package, cursor *dstutil.Cursor, fileRestorer *decorator.FileRestorer) int {
	var updated int
	switch node := cursor.Node().(type) {
	case *dst.Ident:
		for _, rule := range kcpClientTypeRules {
			if !strings.HasPrefix(node.Path, rule.from) || strings.HasSuffix(node.Path, "scheme") {
				continue
			}
			if strings.Contains(node.Path, "cluster") {
				continue
			}
			if !rule.nameMatcher.MatchString(node.Name) {
				continue
			}
			alias := rule.formatAlias(strings.Split(strings.TrimPrefix(strings.TrimPrefix(node.Path, rule.from), "/"), "/"))
			kcpPath := strings.ReplaceAll(node.Path, rule.from, rule.to)
			fileRestorer.Alias[kcpPath] = alias
			kcpName := strings.ReplaceAll(
				strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(node.Name, "Lister", "ClusterLister"),
					"Informer", "ClusterInformer"),
					"Getter", "ClusterGetter"),
					"Interface", "ClusterInterface"),
					"Clientset", "ClusterClientset"),
					"NewClusterFor", "NewFor"),
					"SharedCluster", "Shared"),
				"ClusterCluster", "Cluster") // lol
			if node.Path == kcpPath && node.Name == kcpName {
				continue
			}
			logrus.WithFields(logrus.Fields{
				"before": fmt.Sprintf("%s.%s", node.Path, node.Name),
				"after":  fmt.Sprintf("%s.%s", kcpPath, kcpName),
				"alias":  alias,
			}).Info("updating identifier")
			cursor.Replace(&dst.Ident{
				Name: kcpName,
				Path: kcpPath,
				Obj:  node.Obj,
				Decs: node.Decs,
			})
			updated += 1
		}
	}
	return updated
}

// kcpkubernetesclientset.ClusterInterface.Cluster(logicalcluster.Wildcard) -> kcpkubernetesclientset.ClusterInterface

// kcpkubernetesclientset.ClusterInterface.(...)(logicalcluster.WithCluster(ctx, clusterName), ...) -> kcpkubernetesclientset.ClusterInterface.Cluster(clusterName).(...) // just delete wildcards here
func rewriteClusterInterfaceCall(pkg *decorator.Package, cursor *dstutil.Cursor, _ *decorator.FileRestorer) int {
	var updated int
	switch node := cursor.Node().(type) {
	case *dst.CallExpr:
		// we need a function call
		clientFunction, ok := node.Fun.(*dst.SelectorExpr)
		if !ok {
			break
		}
		// with more than one argument
		if len(node.Args) < 2 {
			break
		}
		// where the first argument itself is a function call
		clusterScopingCall, ok := node.Args[0].(*dst.CallExpr)
		if !ok {
			break
		}
		// that creates a cluster-aware context
		if id, ok := clusterScopingCall.Fun.(*dst.Ident); !ok || id.Name != "WithCluster" {
			break
		}
		// from two arguments
		if len(clusterScopingCall.Args) != 2 {
			break
		}
		ctx := clusterScopingCall.Args[0]
		clusterName := clusterScopingCall.Args[1]

		// and we get to that function call with a obj -> groupversion -> kind chain of calls
		kindCall, ok := clientFunction.X.(*dst.CallExpr)
		if !ok {
			break
		}

		kindFunction, ok := kindCall.Fun.(*dst.SelectorExpr)
		if !ok {
			break
		}

		groupVersionCall, ok := kindFunction.X.(*dst.CallExpr)
		if !ok {
			break
		}

		groupVersionFunction, ok := groupVersionCall.Fun.(*dst.SelectorExpr)
		if !ok {
			break
		}

		types := sets.NewString(
			"github.com/kcp-dev/kcp/pkg/client/clientset/versioned.ClusterInterface",
			"github.com/kcp-dev/kcp/pkg/client/clientset/versioned.Interface",
			"*github.com/kcp-dev/kcp/pkg/client/clientset/versioned.ClusterClientset",
			"*github.com/kcp-dev/kcp/pkg/client/clientset/versioned.Clientset",
			"github.com/kcp-dev/kcp/pkg/client/clientset/versioned/cluster.ClusterInterface",
			"*github.com/kcp-dev/kcp/pkg/client/clientset/versioned/cluster.ClusterClientset",
		)
		if objType := pkg.TypesInfo.TypeOf(pkg.Decorator.Ast.Nodes[groupVersionFunction.X].(ast.Expr)).String(); !types.Has(objType) {
			logrus.Infof("wrong type %s", objType)
			break
		}

		cursor.Replace(&dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X: &dst.CallExpr{
					Fun: &dst.SelectorExpr{
						X: &dst.CallExpr{
							Fun: &dst.SelectorExpr{
								X: &dst.CallExpr{
									Fun: &dst.SelectorExpr{
										X: groupVersionFunction.X,
										Sel: &dst.Ident{
											Name: "Cluster",
										},
									},
									Args: []dst.Expr{clusterName},
								},
								Sel:  groupVersionFunction.Sel,
								Decs: groupVersionFunction.Decs,
							},
							Args:     groupVersionCall.Args,
							Ellipsis: groupVersionCall.Ellipsis,
							Decs:     groupVersionCall.Decs,
						},
						Sel:  kindFunction.Sel,
						Decs: kindFunction.Decs,
					},
					Args:     kindCall.Args,
					Ellipsis: kindCall.Ellipsis,
					Decs:     kindCall.Decs,
				},
				Sel:  clientFunction.Sel,
				Decs: clientFunction.Decs,
			},
			Args:     append([]dst.Expr{ctx}, node.Args[1:]...),
			Ellipsis: node.Ellipsis,
			Decs:     node.Decs,
		})
		updated += 1
	}
	return updated
}

// <group><version>client.ClusterInterface.(...)(logicalcluster.WithCluster(ctx, clusterName), ...) -> <group><version>client.ClusterInterface.Cluster(clusterName).(...) // rewrite .Foo(ns) to .Cluster().Namespace(ns) ... wildcards here are going to require human refactor?

// key := cache.MetaNamespaceKeyFunc(obj) -> key := kcpcache.MetaClusterNamespaceKeyFunc(obj)
func rewriteKeyFuncs(pkg *decorator.Package, cursor *dstutil.Cursor, fileRestorer *decorator.FileRestorer) int {
	var updated int
	switch node := cursor.Node().(type) {
	case *dst.CallExpr:
		// we need a function call to cache.(DeletionHandling)?MetaNamespaceKeyFunc()
		expr, ok := node.Fun.(*dst.Ident)
		if !ok {
			break
		}
		if expr.Name != "MetaNamespaceKeyFunc" && expr.Name != "DeletionHandlingMetaNamespaceKeyFunc" {
			break
		}
		if expr.Path != "k8s.io/client-go/tools/cache" {
			break
		}
		name := strings.ReplaceAll(expr.Name, "MetaNamespace", "MetaClusterNamespace")
		fileRestorer.Alias[kcpCachePath] = kcpCacheAlias
		cursor.Replace(&dst.CallExpr{
			Fun: &dst.Ident{
				Name: name,
				Path: kcpCachePath,
			},
			Args: node.Args,
		})
		updated += 1
	}
	return updated
}

// ns, name, err := SplitMetaNamespaceKey -> clusterName, ns, name, err := SplitMetaClusterNamespaceKey
func rewriteKeySplits(pkg *decorator.Package, cursor *dstutil.Cursor, fileRestorer *decorator.FileRestorer) int {
	var updated int
	switch node := cursor.Node().(type) {
	case *dst.AssignStmt:
		// we need an assignment ns, name, err := cache.SplitMetaNamespaceKey(key)
		if len(node.Lhs) != 3 {
			break
		}
		if len(node.Rhs) != 1 {
			break
		}

		// we need a function call on the right hand side
		call, ok := node.Rhs[0].(*dst.CallExpr)
		if !ok {
			break
		}
		// to cache.SplitMetaNamespaceKey()
		expr, ok := call.Fun.(*dst.Ident)
		if !ok {
			break
		}
		if expr.Name != "SplitMetaNamespaceKey" {
			break
		}
		if expr.Path != "k8s.io/client-go/tools/cache" {
			break
		}
		// with one argument
		if len(call.Args) != 1 {
			break
		}

		fileRestorer.Alias[kcpCachePath] = kcpCacheAlias
		cursor.Replace(&dst.AssignStmt{
			Lhs: append([]dst.Expr{&dst.Ident{
				Name: "clusterName",
			}},
				node.Lhs...,
			),
			Tok: node.Tok,
			Rhs: []dst.Expr{
				&dst.CallExpr{
					Fun: &dst.Ident{
						Name: "SplitMetaClusterNamespaceKey",
						Path: kcpCachePath,
					},
					Args: call.Args,
				},
			},
		})
		updated += 1
	}
	return updated
}

// clusterName, name := clusters.SplitClusterAwareKey(key)
//
// into
//
// clusterName, _, name, err := kcpcache.SplitMetaClusterNamespaceKey(key)
// if err != nil {
//     runtime.HandleError(err)
//     return false
// }

// namespace, clusterAwareName, err := cache.SplitMetaNamespaceKey(key)
// if err != nil {
//     runtime.HandleError(err)
//     return false
// }
// clusterName, name := clusters.SplitClusterAwareKey(clusterAwareName)
//
// into
// clusterName, namespace, name, err := kcpcache.SplitMetaClusterNamespaceKey(key)
// if err != nil {
//     runtime.HandleError(err)
//     return false
// }
func rewriteKcpKeySplits(pkg *decorator.Package, cursor *dstutil.Cursor, fileRestorer *decorator.FileRestorer) int {
	var updated int
	switch node := cursor.Node().(type) {
	case *dst.BlockStmt:
		// in amongst the statements in this block, we're looking for calls to cache.SplitMetaNamespaceKey and clusters.SplitClusterAwareKey
		for i := range node.List {
			if i == len(node.List)-1 {
				continue // we're looking for a relationship between nodes, so we need more than one more
			}
			assignment, ok := node.List[i].(*dst.AssignStmt)
			if !ok {
				continue
			}

			// we need an assignment ns, clusterAwareName, err := cache.SplitMetaNamespaceKey(key)
			if len(assignment.Lhs) != 3 {
				continue
			}
			if len(assignment.Rhs) != 1 {
				continue
			}

			// we need a function call on the right hand side
			call, ok := assignment.Rhs[0].(*dst.CallExpr)
			if !ok {
				continue
			}
			// to cache.SplitMetaNamespaceKey()
			expr, ok := call.Fun.(*dst.Ident)
			if !ok {
				continue
			}
			if expr.Name != "SplitMetaNamespaceKey" {
				continue
			}

			for j := i + 1; j < len(node.List); j++ {
				secondAssignment, ok := node.List[j].(*dst.AssignStmt)
				if !ok {
					continue
				}
				// we need an assignment clusterName, name := clusters.SplitClusterAwareKey(clusterAwareName)
				if len(secondAssignment.Lhs) != 2 {
					continue
				}
				if len(secondAssignment.Rhs) != 1 {
					continue
				}

				// we need a function call on the right hand side
				splitClusterAwareKeyCall, ok := secondAssignment.Rhs[0].(*dst.CallExpr)
				if !ok {
					continue
				}
				// to clusters.SplitClusterAwareKey()
				splitClusterAwareKeyExpr, ok := splitClusterAwareKeyCall.Fun.(*dst.Ident)
				if !ok {
					continue
				}
				if splitClusterAwareKeyExpr.Name != "SplitClusterAwareKey" {
					continue
				}
				// we need the arguments passed to be from the first split
				keyToSplit, ok := splitClusterAwareKeyCall.Args[0].(*dst.Ident)
				if !ok {
					continue
				}
				clusterAwareKey, ok := assignment.Lhs[1].(*dst.Ident)
				if !ok {
					continue
				}
				if keyToSplit.Name != clusterAwareKey.Name {
					continue
				}
				fileRestorer.Alias[kcpCachePath] = kcpCacheAlias
				// copy all statements, but
				list := node.List[0:i]
				// rewrite the assignment to have clusterName, and
				list = append(list, &dst.AssignStmt{
					Lhs: []dst.Expr{
						secondAssignment.Lhs[0], // clusterName
						assignment.Lhs[0],       // namespace
						secondAssignment.Lhs[1], // name
						assignment.Lhs[2],       // err
					},
					Tok: assignment.Tok,
					Rhs: []dst.Expr{&dst.CallExpr{
						Fun: &dst.Ident{
							Name: "SplitMetaClusterNamespaceKey",
							Path: kcpCachePath,
						},
						Args: call.Args,
					}},
					Decs: assignment.Decs,
				})
				// remove the SplitClusterAwareKey call
				list = append(list, node.List[i+1:j]...)
				list = append(list, node.List[j+1:]...)
				cursor.Replace(&dst.BlockStmt{
					List:           list,
					RbraceHasNoPos: node.RbraceHasNoPos,
					Decs:           node.Decs,
				})
				updated += 1
			}
		}
	}
	return updated
}

// lister.Get(clusters.ToClusterAwareKey(clusterName, name)) -> lister.Cluster(clusterName).Get(name)
func rewriteListerGet(pkg *decorator.Package, cursor *dstutil.Cursor, _ *decorator.FileRestorer) int {
	var updated int
	switch node := cursor.Node().(type) {
	case *dst.CallExpr:
		// we need a method call to Get()
		expr, ok := node.Fun.(*dst.SelectorExpr)
		if !ok {
			break
		}
		if expr.Sel.Name != "Get" {
			break
		}
		// we need a method call on a k8s lister
		astNode := pkg.Decorator.Ast.Nodes[expr.X]
		if astNode == nil {
			break
		}
		astExpr, ok := astNode.(ast.Expr)
		if !ok {
			break
		}
		nodeType := pkg.TypesInfo.TypeOf(astExpr)
		if nodeType == nil {
			break
		}
		// with one argument
		if len(node.Args) != 1 {
			break
		}
		// which itself is a function call
		keyFuncCall, ok := node.Args[0].(*dst.CallExpr)
		if !ok {
			break
		}
		// that creates a cluster-aware key
		if id, ok := keyFuncCall.Fun.(*dst.Ident); !ok || id.Name != "ToClusterAwareKey" {
			break
		}
		// from two arguments
		if len(keyFuncCall.Args) != 2 {
			break
		}
		clusterName := keyFuncCall.Args[0]
		name := keyFuncCall.Args[1]
		cursor.Replace(&dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X: &dst.CallExpr{
					Fun: &dst.SelectorExpr{
						X: expr.X,
						Sel: &dst.Ident{
							Name: "Cluster",
						},
					},
					Args: []dst.Expr{clusterName},
				},
				Sel: &dst.Ident{
					Name: "Get",
				},
			},
			Args: []dst.Expr{name},
		})
		updated += 1
	}
	return updated
}
