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

package informergen

import (
	"io"
	"text/template"

	"k8s.io/code-generator/cmd/client-gen/types"

	"github.com/kcp-dev/code-generator/v2/pkg/parser"
)

type Generic struct {
	// Groups are the groups in this informer factory.
	Groups []parser.Group

	// GroupVersionKinds are all the kinds we need to support,indexed by group and version.
	GroupVersionKinds map[types.Group]map[parser.Version][]parser.Kind

	// APIPackagePath is the root directory under which API types exist.
	// e.g. "k8s.io/api"
	APIPackagePath string

	// SingleClusterInformerPackagePath is the package under which the cluster-unaware listers are exposed.
	// e.g. "k8s.io/client-go/informers"
	SingleClusterInformerPackagePath string
}

func (g *Generic) WriteContent(w io.Writer) error {
	templ, err := template.New("generic").Funcs(templateFuncs).Parse(genericInformer)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"groups":                           g.Groups,
		"groupVersionKinds":                g.GroupVersionKinds,
		"apiPackagePath":                   g.APIPackagePath,
		"singleClusterInformerPackagePath": g.SingleClusterInformerPackagePath,
		"useUpstreamInterfaces":            g.SingleClusterInformerPackagePath != "",
	}
	return templ.Execute(w, m)
}

var genericInformer = `
// Code generated by kcp code-generator. DO NOT EDIT.

package informers

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	{{if .useUpstreamInterfaces -}}
	upstreaminformers "{{.singleClusterInformerPackagePath}}"
	{{end -}}

{{range .groups}}	{{.GoPackageAlias}} "{{$.apiPackagePath}}/{{.PackageName}}/{{.Version.PackageName}}"
{{end -}}
)

type GenericClusterInformer interface {
	Cluster(logicalcluster.Name) {{if .useUpstreamInterfaces}}upstreaminformers.{{end}}GenericInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() kcpcache.GenericClusterLister
}

{{ if not .useUpstreamInterfaces }}
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}
{{end }}

type genericClusterInformer struct {
	informer kcpcache.ScopeableSharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.informer
}

// Lister returns the GenericClusterLister.
func (f *genericClusterInformer) Lister() kcpcache.GenericClusterLister {
	return kcpcache.NewGenericClusterLister(f.Informer().GetIndexer(), f.resource)
}

// Cluster scopes to a GenericInformer.
func (f *genericClusterInformer) Cluster(clusterName logicalcluster.Name) {{if .useUpstreamInterfaces}}upstreaminformers.{{end}}GenericInformer {
	return &genericInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().ByCluster(clusterName),
	}
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	lister   cache.GenericLister
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return f.lister
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error) {
	switch resource {
{{range $group := .groups}}	// Group={{.Group.NonEmpty}}, Version={{.Version}}
{{range $kind := index (index $.groupVersionKinds .Group) .Version}}	case {{$group.GoPackageAlias}}.SchemeGroupVersion.WithResource("{{$kind.Plural|toLower}}"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.{{$group.GroupGoName}}().{{$group.Version}}().{{$kind.Plural}}().Informer()}, nil
{{end -}}
{{end -}}
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

{{if not .useUpstreamInterfaces -}}
// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedScopedInformerFactory) ForResource(resource schema.GroupVersionResource) ({{if .useUpstreamInterfaces}}upstreaminformers.{{end}}GenericInformer, error) {
	switch resource {
{{range $group := .groups}}	// Group={{.Group.NonEmpty}}, Version={{.Version}}
{{range $kind := index (index $.groupVersionKinds .Group) .Version}}	case {{$group.GoPackageAlias}}.SchemeGroupVersion.WithResource("{{$kind.Plural|toLower}}"):
		informer := f.{{$group.GroupGoName}}().{{$group.Version}}().{{$kind.Plural}}().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
{{end -}}
{{end -}}
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
{{end}}
`
