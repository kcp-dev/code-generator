/*
Copyright 2016 The Kubernetes Authors.

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

package generators

import (
	"fmt"
	"io"
	"path"
	"strings"

	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
	"k8s.io/klog/v2"

	"github.com/kcp-dev/code-generator/v3/cmd/cluster-client-gen/generators/util"
	clientgentypes "github.com/kcp-dev/code-generator/v3/cmd/cluster-client-gen/types"
)

// informerGenerator produces a file of listers for a given GroupVersion and
// type.
type informerGenerator struct {
	generator.GoGenerator
	outputPackage                          string
	groupPkgName                           string
	groupVersion                           clientgentypes.GroupVersion
	groupGoName                            string
	typeToGenerate                         *types.Type
	imports                                namer.ImportTracker
	clientSetPackage                       string
	listersPackage                         string
	internalInterfacesPackage              string
	singleClusterVersionedClientSetPackage string
	singleClusterListersPackage            string
	singleClusterInformersPackage          string
}

var _ generator.Generator = &informerGenerator{}

func (g *informerGenerator) Filter(c *generator.Context, t *types.Type) bool {
	return t == g.typeToGenerate
}

func (g *informerGenerator) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *informerGenerator) Imports(c *generator.Context) (imports []string) {
	imports = append(imports, g.imports.ImportLines()...)
	imports = append(imports,
		`"github.com/kcp-dev/logicalcluster/v3"`,
		`kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"`,
		`kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"`,
	)
	return
}

func (g *informerGenerator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	klog.V(5).Infof("processing type %v", t)

	clusterListersPkg := fmt.Sprintf("%s/%s/%s", g.listersPackage, g.groupPkgName, strings.ToLower(g.groupVersion.Version.NonEmpty()))
	clientSetInterface := c.Universe.Type(types.Name{Package: g.singleClusterVersionedClientSetPackage, Name: "Interface"})
	clientSetClusterInterface := c.Universe.Type(types.Name{Package: g.clientSetPackage, Name: "ClusterInterface"})
	informerFor := "InformerFor"

	informerPkg := g.outputPackage
	generateScopedInformer := true
	if g.singleClusterInformersPackage != "" {
		informerPkg = path.Join(g.singleClusterInformersPackage, g.groupPkgName, strings.ToLower(g.groupVersion.Version.NonEmpty()))
		generateScopedInformer = false
	}

	listersPkg := g.listersPackage
	if g.singleClusterInformersPackage != "" {
		listersPkg = g.singleClusterListersPackage
	}
	listersPkg = path.Join(listersPkg, g.groupPkgName, strings.ToLower(g.groupVersion.Version.NonEmpty()))

	tags, err := util.ParseClientGenTags(append(t.SecondClosestCommentLines, t.CommentLines...))
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"namespaced":                            !tags.NonNamespaced,
		"apiScheme":                             c.Universe.Type(apiScheme),
		"cacheIndexers":                         c.Universe.Type(cacheIndexers),
		"cacheListWatch":                        c.Universe.Type(cacheListWatch),
		"cacheMetaNamespaceIndexFunc":           c.Universe.Function(cacheMetaNamespaceIndexFunc),
		"cacheNamespaceIndex":                   c.Universe.Variable(cacheNamespaceIndex),
		"cacheNewSharedIndexInformer":           c.Universe.Function(cacheNewSharedIndexInformer),
		"cacheSharedIndexInformer":              c.Universe.Type(cacheSharedIndexInformer),
		"scopeableCacheSharedIndexInformer":     c.Universe.Type(scopeableCacheSharedIndexInformer),
		"clientSetInterface":                    clientSetInterface,
		"clientSetClusterInterface":             clientSetClusterInterface,
		"contextContext":                        c.Universe.Type(contextContext),
		"contextBackground":                     c.Universe.Function(contextBackgroundFunc),
		"group":                                 namer.IC(g.groupGoName),
		"informerFor":                           informerFor,
		"interfacesTweakListOptionsFunc":        c.Universe.Type(types.Name{Package: g.internalInterfacesPackage, Name: "TweakListOptionsFunc"}),
		"interfacesSharedInformerFactory":       c.Universe.Type(types.Name{Package: g.internalInterfacesPackage, Name: "SharedInformerFactory"}),
		"interfacesSharedScopedInformerFactory": c.Universe.Type(types.Name{Package: g.internalInterfacesPackage, Name: "SharedScopedInformerFactory"}),
		"listOptions":                           c.Universe.Type(listOptions),
		"lister":                                c.Universe.Type(types.Name{Package: listersPkg, Name: t.Name.Name + "Lister"}),
		"clusterLister":                         c.Universe.Type(types.Name{Package: clusterListersPkg, Name: t.Name.Name + "ClusterLister"}),
		"informerInterface":                     c.Universe.Type(types.Name{Package: informerPkg, Name: t.Name.Name + "Informer"}),
		"namespaceAll":                          c.Universe.Type(metav1NamespaceAll),
		"newLister":                             c.Universe.Function(types.Name{Package: listersPkg, Name: "New" + t.Name.Name + "Lister"}),
		"listerInterface":                       c.Universe.Function(types.Name{Package: g.singleClusterListersPackage, Name: t.Name.Name + "Lister"}),
		"newClusterLister":                      c.Universe.Function(types.Name{Package: clusterListersPkg, Name: "New" + t.Name.Name + "ClusterLister"}),
		"runtimeObject":                         c.Universe.Type(runtimeObject),
		"timeDuration":                          c.Universe.Type(timeDuration),
		"type":                                  t,
		"v1ListOptions":                         c.Universe.Type(v1ListOptions),
		"version":                               namer.IC(g.groupVersion.Version.String()),
		"watchInterface":                        c.Universe.Type(watchInterface),
		"generateScopedInformer":                generateScopedInformer,
	}

	sw.Do(typeClusterInformerInterface, m)
	sw.Do(typeClusterInformerStruct, m)
	sw.Do(typeClusterInformerPublicConstructor, m)
	sw.Do(typeFilteredInformerPublicConstructor, m)
	sw.Do(typeInformerConstructor, m)
	sw.Do(typeInformerInformer, m)
	sw.Do(typeInformerLister, m)
	sw.Do(typeInformerCluster, m)
	sw.Do(typeInformer, m)

	if generateScopedInformer {
		sw.Do(typeScopedInformer, m)
	}

	return sw.Error()
}

var typeClusterInformerInterface = `
// $.type|public$ClusterInformer provides access to a shared informer and lister for
// $.type|publicPlural$.
type $.type|public$ClusterInformer interface {
	Cluster(logicalcluster.Name) $.informerInterface|raw$
	Informer() $.scopeableCacheSharedIndexInformer|raw$
	Lister() $.clusterLister|raw$
}
`

var typeClusterInformerStruct = `
type $.type|private$ClusterInformer struct {
	factory          $.interfacesSharedInformerFactory|raw$
	tweakListOptions $.interfacesTweakListOptionsFunc|raw$
}
`

var typeClusterInformerPublicConstructor = `
// New$.type|public$ClusterInformer constructs a new informer for $.type|public$ type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func New$.type|public$ClusterInformer(client $.clientSetClusterInterface|raw$, resyncPeriod $.timeDuration|raw$, indexers $.cacheIndexers|raw$) $.scopeableCacheSharedIndexInformer|raw$ {
	return NewFiltered$.type|public$ClusterInformer(client, resyncPeriod, indexers, nil)
}
`

var typeFilteredInformerPublicConstructor = `
// NewFiltered$.type|public$ClusterInformer constructs a new informer for $.type|public$ type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFiltered$.type|public$ClusterInformer(client $.clientSetClusterInterface|raw$, resyncPeriod $.timeDuration|raw$, indexers $.cacheIndexers|raw$, tweakListOptions $.interfacesTweakListOptionsFunc|raw$) $.scopeableCacheSharedIndexInformer|raw$ {
	return kcpinformers.NewSharedIndexInformer(
		&$.cacheListWatch|raw${
			ListFunc: func(options $.v1ListOptions|raw$) ($.runtimeObject|raw$, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.$.group$$.version$().$.type|publicPlural$().List($.contextBackground|raw$(), options)
			},
			WatchFunc: func(options $.v1ListOptions|raw$) ($.watchInterface|raw$, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.$.group$$.version$().$.type|publicPlural$().Watch($.contextBackground|raw$(), options)
			},
		},
		&$.type|raw${},
		resyncPeriod,
		indexers,
	)
}
`

var typeInformerConstructor = `
func (f *$.type|private$ClusterInformer) defaultInformer(client $.clientSetClusterInterface|raw$, resyncPeriod $.timeDuration|raw$) $.scopeableCacheSharedIndexInformer|raw$ {
	return NewFiltered$.type|public$ClusterInformer(client, resyncPeriod, $.cacheIndexers|raw${
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc,
	}, f.tweakListOptions)
}
`

var typeInformerInformer = `
func (f *$.type|private$ClusterInformer) Informer() $.scopeableCacheSharedIndexInformer|raw$ {
	return f.factory.$.informerFor$(&$.type|raw${}, f.defaultInformer)
}
`

var typeInformerLister = `
func (f *$.type|private$ClusterInformer) Lister() $.clusterLister|raw$ {
	return $.newClusterLister|raw$(f.Informer().GetIndexer())
}
`

var typeInformerCluster = `
func (f *$.type|private$ClusterInformer) Cluster(clusterName logicalcluster.Name) $.informerInterface|raw$ {
	return &$.type|private$Informer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}
`

var typeInformer = `
type $.type|private$Informer struct {
	informer $.cacheSharedIndexInformer|raw$
	lister   $.lister|raw$
}

func (f *$.type|private$Informer) Informer() $.cacheSharedIndexInformer|raw$ {
	return f.informer
}

func (f *$.type|private$Informer) Lister() $.lister|raw$ {
	return f.lister
}
`

var typeScopedInformer = `
// $.informerInterface|raw$ provides access to a shared informer and lister for
// $.type|publicPlural$.
type $.informerInterface|raw$ interface {
	Informer() $.cacheSharedIndexInformer|raw$
	Lister() $.lister|raw$
}

type $.type|private$ScopedInformer struct {
	factory          $.interfacesSharedScopedInformerFactory|raw$
	tweakListOptions $.interfacesTweakListOptionsFunc|raw$
	$if .namespaced$namespace        string$end -$
}

// New$.type|public$Informer constructs a new informer for $.type|public$ type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func New$.type|public$Informer(client $.clientSetInterface|raw$, resyncPeriod $.timeDuration|raw$$if .namespaced$, namespace string$end$, indexers cache.Indexers) $.cacheSharedIndexInformer|raw$ {
	return NewFiltered$.type|public$Informer(client, resyncPeriod$if .namespaced$, namespace$end$, indexers, nil)
}

// NewFiltered$.type|public$Informer constructs a new informer for $.type|public$ type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFiltered$.type|public$Informer(client $.clientSetInterface|raw$, resyncPeriod $.timeDuration|raw$$if .namespaced$, namespace string$end$, indexers $.cacheIndexers|raw$, tweakListOptions $.interfacesTweakListOptionsFunc|raw$) $.cacheSharedIndexInformer|raw$ {
	return $.cacheNewSharedIndexInformer|raw$(
		&$.cacheListWatch|raw${
			ListFunc: func(options $.v1ListOptions|raw$) ($.runtimeObject|raw$, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.$.group$$.version$().$.type|publicPlural$($if .namespaced$namespace$end$).List($.contextBackground|raw$(), options)
			},
			WatchFunc: func(options $.v1ListOptions|raw$) ($.watchInterface|raw$, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.$.group$$.version$().$.type|publicPlural$($if .namespaced$namespace$end$).Watch($.contextBackground|raw$(), options)
			},
		},
		&$.type|raw${},
		resyncPeriod,
		indexers,
	)
}

func (f *$.type|private$ScopedInformer) Informer() $.cacheSharedIndexInformer|raw$ {
	return f.factory.InformerFor(&$.type|raw${}, f.defaultInformer)
}

func (f *$.type|private$ScopedInformer) Lister() $.lister|raw$ {
	return $.newLister|raw$(f.Informer().GetIndexer())
}

func (f *$.type|private$ScopedInformer) defaultInformer(client $.clientSetInterface|raw$, resyncPeriod $.timeDuration|raw$) $.cacheSharedIndexInformer|raw$ {
$if .namespaced -$
	return NewFiltered$.type|public$Informer(client, resyncPeriod, f.namespace, $.cacheIndexers|raw${
		$.cacheNamespaceIndex|raw$: $.cacheMetaNamespaceIndexFunc|raw$,
	}, f.tweakListOptions)
$else -$
	return NewFiltered$.type|public$Informer(client, resyncPeriod, $.cacheIndexers|raw${}, f.tweakListOptions)
$end -$
}
`
