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

package parser

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/ast/astutil"

	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/gengo/v2/namer"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/kcp-dev/code-generator/v3/cmd/cluster-client-gen/types"
)

// isForbiddenGroupVersion hacks around the k8s client-set, where types have +genclient but aren't meant to have
// generated clients ... ?
func isForbiddenGroupVersion(group types.Group, version types.Version) bool {
	forbidden := map[string]sets.Set[string]{
		"abac.authorization.kubernetes.io": sets.New[string]("v0", "v1beta1"),
		"componentconfig":                  sets.New[string]("v1alpha1"),
		"imagepolicy.k8s.io":               sets.New[string]("v1alpha1"),
		"admission.k8s.io":                 sets.New[string]("v1", "v1beta1"),
	}
	versions, ok := forbidden[group.String()]
	return ok && versions.Has(version.String())
}

// CollectKinds finds all groupVersionKinds for which the k8s client-generators are run and the set of
// verbs are supported.
// When we are looking at a package, we can determine the group and version by copying the upstream
// logic:
// https://github.com/kubernetes/kubernetes/blob/f046bdf24e69ac31d3e1ed56926d9a7c715f1cc8/staging/src/k8s.io/code-generator/cmd/lister-gen/generators/lister.go#L93-L106
func CollectKinds(ctx *genall.GenerationContext, verbs ...string) (map[Group]map[types.PackageVersion][]Kind, error) {
	groupVersionKinds := map[Group]map[types.PackageVersion][]Kind{}
	for _, root := range ctx.Roots {
		logger := klog.Background()
		logger.V(4).Info("processing " + root.PkgPath)
		parts := strings.Split(root.PkgPath, "/")
		groupName := types.Group(parts[len(parts)-2])
		version := types.PackageVersion{
			Version: types.Version(parts[len(parts)-1]),
			Package: root.PkgPath,
		}

		packageMarkers, err := markers.PackageMarkers(ctx.Collector, root)
		if err != nil {
			return nil, err
		}

		// look for a reference to runtime.APIVersionInternal to skip internal APIs
		var isInternalAPI bool
		for _, file := range root.Syntax {
			astutil.Apply(file, nil, func(cursor *astutil.Cursor) bool {
				switch node := cursor.Node().(type) {
				case *ast.SelectorExpr:
					pkg, ok := node.X.(*ast.Ident)
					if !ok {
						break
					}
					if pkg.Name == "runtime" && node.Sel.Name == "APIVersionInternal" {
						isInternalAPI = true
						return false
					}
				}
				return true
			})
		}
		if isInternalAPI {
			logger.WithValues("path", root.PkgPath).V(4).Info("skipping package for internal API")
			continue
		}

		groupNameRaw, ok := packageMarkers.Get(GroupNameMarker().Name).(markers.RawArguments)
		if ok {
			// If there's a comment of the form "// +groupName=somegroup" or
			// "// +groupName=somegroup.foo.bar.io", use the first field (somegroup) as the name of the
			// group when generating. [N.B.](skuznets): even though the generators do the indexing here, the group
			// type does it for you, and handles the special case for "internal"
			logger.WithValues("original", groupName, "override", string(groupNameRaw)).V(4).Info("found a group name override")
			groupName = types.Group(groupNameRaw)
		}
		groupGoName := namer.IC(groupName.PackageName())
		// internal.apiserver.k8s.io needs to have a package name of apiserverinternal, but a Go name of internal ...
		if parts := strings.Split(groupName.NonEmpty(), "."); parts[0] == "internal" && len(parts) > 1 {
			groupGoName = namer.IC(parts[0])
		}

		groupGoNameRaw, ok := packageMarkers.Get(GroupGoNameMarker().Name).(markers.RawArguments)
		if ok {
			// If there's a comment of the form "// +groupGoName=SomeUniqueShortName", use that as
			// the Go group identifier in CamelCase.
			groupGoName = namer.IC(string(groupGoNameRaw))
		}
		group := Group{Group: groupName, GoName: groupGoName}

		logger = logger.WithValues("group", group, "version", version, "goName", groupGoName)
		if isForbiddenGroupVersion(group.Group, version.Version) {
			logger.WithValues("package", root.PkgPath).V(4).Info("skipping forbidden package")
			continue
		}
		logger.WithValues("package", root.PkgPath).V(4).Info("collecting kinds in package")

		// find types which have generated clients and support LIST + WATCH
		var kinds []Kind
		var typeErrors []error
		if err := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {
			logger = logger.WithValues("kind", info.Name)
			if !ClientsGeneratedForType(info) {
				logger.V(3).V(4).Info("skipping kind as it has no generated clients")
				return
			}

			supported, err := SupportedVerbs(info)
			if err != nil {
				typeErrors = append(typeErrors, err)
				return
			}
			extensions := ClientExtensions(info)
			if len(verbs) > 0 && !supported.HasAll(verbs...) {
				logger.V(4).Info("skipping kind as it does not support the necessary verbs")
				return
			}

			logger.V(4).Info("will generate for kind")
			kinds = append(kinds, NewKind(info.Name, IsNamespaced(info), supported, extensions))
		}); err != nil {
			return nil, err
		}
		if len(typeErrors) > 0 {
			return nil, errors.NewAggregate(typeErrors)
		}
		if len(kinds) == 0 {
			logger.V(4).Info("skipping group/version as it has no kinds that have generated clients")
			continue
		}
		if _, recorded := groupVersionKinds[group]; !recorded {
			groupVersionKinds[group] = map[types.PackageVersion][]Kind{}
		}
		groupVersionKinds[group][version] = append(groupVersionKinds[group][version], kinds...)
	}
	return groupVersionKinds, nil
}
