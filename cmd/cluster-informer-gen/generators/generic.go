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
	"io"
	"sort"
	"strings"

	codegennamer "k8s.io/code-generator/pkg/namer"
	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"

	clientgentypes "github.com/kcp-dev/code-generator/v3/cmd/cluster-client-gen/types"
)

// genericGenerator generates the generic informer.
type genericGenerator struct {
	generator.GoGenerator
	outputPackage             string
	imports                   namer.ImportTracker
	groupVersions             map[string]clientgentypes.GroupVersions
	groupGoNames              map[string]string
	pluralExceptions          map[string]string
	typesForGroupVersion      map[clientgentypes.GroupVersion][]*types.Type
	singleClusterInformersPkg string
	filtered                  bool
}

var _ generator.Generator = &genericGenerator{}

func (g *genericGenerator) Filter(c *generator.Context, t *types.Type) bool {
	if !g.filtered {
		g.filtered = true
		return true
	}
	return false
}

func (g *genericGenerator) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw":                namer.NewRawNamer(g.outputPackage, g.imports),
		"allLowercasePlural": namer.NewAllLowercasePluralNamer(g.pluralExceptions),
		"publicPlural":       namer.NewPublicPluralNamer(g.pluralExceptions),
		"resource":           codegennamer.NewTagOverrideNamer("resourceName", namer.NewAllLowercasePluralNamer(g.pluralExceptions)),
	}
}

func (g *genericGenerator) Imports(c *generator.Context) (imports []string) {
	imports = append(imports, g.imports.ImportLines()...)
	imports = append(imports,
		`"github.com/kcp-dev/logicalcluster/v3"`,
		`kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"`,
	)
	return
}

type group struct {
	GroupGoName string
	Name        string
	Versions    []*version
}

type groupSort []group

func (g groupSort) Len() int { return len(g) }
func (g groupSort) Less(i, j int) bool {
	return strings.ToLower(g[i].Name) < strings.ToLower(g[j].Name)
}
func (g groupSort) Swap(i, j int) { g[i], g[j] = g[j], g[i] }

type version struct {
	Name      string
	GoName    string
	Resources []*types.Type
}

type versionSort []*version

func (v versionSort) Len() int { return len(v) }
func (v versionSort) Less(i, j int) bool {
	return strings.ToLower(v[i].Name) < strings.ToLower(v[j].Name)
}
func (v versionSort) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

func (g *genericGenerator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "{{", "}}")

	groups := []group{}
	schemeGVs := make(map[*version]*types.Type)

	orderer := namer.Orderer{Namer: namer.NewPrivateNamer(0)}
	for groupPackageName, groupVersions := range g.groupVersions {
		group := group{
			GroupGoName: g.groupGoNames[groupPackageName],
			Name:        groupVersions.Group.NonEmpty(),
			Versions:    []*version{},
		}
		for _, v := range groupVersions.Versions {
			gv := clientgentypes.GroupVersion{Group: groupVersions.Group, Version: v.Version}
			version := &version{
				Name:      v.Version.NonEmpty(),
				GoName:    namer.IC(v.Version.NonEmpty()),
				Resources: orderer.OrderTypes(g.typesForGroupVersion[gv]),
			}
			func() {
				schemeGVs[version] = c.Universe.Variable(types.Name{Package: g.typesForGroupVersion[gv][0].Name.Package, Name: "SchemeGroupVersion"})
			}()
			group.Versions = append(group.Versions, version)
		}
		sort.Sort(versionSort(group.Versions))
		groups = append(groups, group)
	}
	sort.Sort(groupSort(groups))

	genericInformerPkg := g.singleClusterInformersPkg
	generateInformerInterface := false
	if genericInformerPkg == "" {
		genericInformerPkg = g.outputPackage
		generateInformerInterface = true
	}

	m := map[string]interface{}{
		"cacheGenericLister":         c.Universe.Type(cacheGenericLister),
		"cacheNewGenericLister":      c.Universe.Function(cacheNewGenericLister),
		"cacheSharedIndexInformer":   c.Universe.Type(cacheSharedIndexInformer),
		"fmtErrorf":                  c.Universe.Type(fmtErrorfFunc),
		"groups":                     groups,
		"schemeGVs":                  schemeGVs,
		"schemaGroupResource":        c.Universe.Type(schemaGroupResource),
		"schemaGroupVersionResource": c.Universe.Type(schemaGroupVersionResource),
		"genericInformer":            c.Universe.Type(types.Name{Package: genericInformerPkg, Name: "GenericInformer"}),
		"generateInformerInterface":  generateInformerInterface,
	}

	sw.Do(genericClusterInformer, m)
	sw.Do(genericInformer, m)
	sw.Do(forResource, m)

	if generateInformerInterface {
		sw.Do(forScopedResource, m)
	}

	return sw.Error()
}

var genericClusterInformer = `
type GenericClusterInformer interface {
	Cluster(logicalcluster.Name) {{.genericInformer|raw}}
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() kcpcache.GenericClusterLister
}
{{ if .generateInformerInterface }}

type GenericInformer interface {
	Informer() {{.cacheSharedIndexInformer|raw}}
	Lister() {{.cacheGenericLister|raw}}
}
{{ end }}

type genericClusterInformer struct {
	informer kcpcache.ScopeableSharedIndexInformer
	resource {{.schemaGroupResource|raw}}
}

// Informer returns the SharedIndexInformer.
func (f *genericClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericClusterInformer) Lister() kcpcache.GenericClusterLister {
	return kcpcache.NewGenericClusterLister(f.Informer().GetIndexer(), f.resource)
}

// Cluster scopes to a GenericInformer.
func (f *genericClusterInformer) Cluster(clusterName logicalcluster.Name) {{.genericInformer|raw}} {
	return &genericInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().ByCluster(clusterName),
	}
}
`

var genericInformer = `
type genericInformer struct {
	informer {{.cacheSharedIndexInformer|raw}}
	lister   {{.cacheGenericLister|raw}}
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() {{.cacheSharedIndexInformer|raw}} {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() {{.cacheGenericLister|raw}} {
	return f.lister
}
`

var forResource = `
// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource {{.schemaGroupVersionResource|raw}}) (GenericClusterInformer, error) {
	switch resource {
		{{range $group := .groups -}}{{$GroupGoName := .GroupGoName -}}
			{{range $version := .Versions -}}
	// Group={{$group.Name}}, Version={{.Name}}
				{{range .Resources -}}
	case {{index $.schemeGVs $version|raw}}.WithResource("{{.|resource}}"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.{{$GroupGoName}}().{{$version.GoName}}().{{.|publicPlural}}().Informer()}, nil
				{{end}}
			{{end}}
		{{end -}}
	}

	return nil, {{.fmtErrorf|raw}}("no informer found for %v", resource)
}
`

var forScopedResource = `
// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedScopedInformerFactory) ForResource(resource {{.schemaGroupVersionResource|raw}}) ({{.genericInformer|raw}}, error) {
	switch resource {
		{{range $group := .groups -}}{{$GroupGoName := .GroupGoName -}}
			{{range $version := .Versions -}}
	// Group={{$group.Name}}, Version={{.Name}}
				{{range .Resources -}}
	case {{index $.schemeGVs $version|raw}}.WithResource("{{.|resource}}"):
		informer := f.{{$GroupGoName}}().{{$version.GoName}}().{{.|publicPlural}}().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
				{{end}}
			{{end}}
		{{end -}}
	}

	return nil, {{.fmtErrorf|raw}}("no informer found for %v", resource)
}
`
