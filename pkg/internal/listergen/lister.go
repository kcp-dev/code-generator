package listergen

import (
	"io"
	"text/template"

	"github.com/kcp-dev/code-generator/v2/pkg/parser"
)

type Lister struct {
	// Group is:
	// - the name of the API group, e.g. "authorization",
	// - the version and package path of the API, e.g. "v1" and "k8s.io/api/rbac/v1"
	Group parser.Group
	// Kind is the kind for which we are generating listers, e.g. "ClusterRole"
	Kind parser.Kind

	// APIPackagePath is the root directory under which API types exist.
	// e.g. "k8s.io/api"
	APIPackagePath string

	// SingleClusterListerPackagePath is the fully qualified Go package name under which the (pre-existing)
	// listers for single-cluster contexts are defined. Option. e.g. "k8s.io/client-go/listers"
	SingleClusterListerPackagePath string
}

func (l *Lister) WriteContent(w io.Writer) error {
	templ, err := template.New("lister").Funcs(templateFuncs).Parse(lister)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"group":                          l.Group,
		"kind":                           &l.Kind,
		"apiPackagePath":                 l.APIPackagePath,
		"singleClusterListerPackagePath": l.SingleClusterListerPackagePath,
		"useUpstreamInterfaces":          l.SingleClusterListerPackagePath != "",
	}
	return templ.Execute(w, m)
}

var lister = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package {{.group.Version.PackageName}}

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"	
	"github.com/kcp-dev/logicalcluster/v3"
	
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/api/errors"

	{{.group.PackageAlias}} "{{.apiPackagePath}}/{{.group.Group.PackageName}}/{{.group.Version.PackageName}}"
	{{if .useUpstreamInterfaces -}}
	{{.group.PackageAlias}}listers "{{.singleClusterListerPackagePath}}/{{.group.Group.PackageName}}/{{.group.Version.PackageName}}"
	{{end -}}
)

// {{.kind.String}}ClusterLister can list {{.kind.Plural}} across all workspaces, or scope down to a {{.kind.String}}Lister for one workspace.
// All objects returned here must be treated as read-only.
type {{.kind.String}}ClusterLister interface {
	// List lists all {{.kind.Plural}} in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error)
	// Cluster returns a lister that can list and get {{.kind.Plural}} in one workspace.
{{if not .useUpstreamInterfaces -}}
	Cluster(clusterName logicalcluster.Name) {{.kind.String}}Lister
{{else -}}
	Cluster(clusterName logicalcluster.Name){{.group.PackageAlias}}listers.{{.kind.String}}Lister
{{end -}}
	{{.kind.String}}ClusterListerExpansion
}

type {{.kind.String | lowerFirst }}ClusterLister struct {
	indexer cache.Indexer
}

// New{{.kind.String}}ClusterLister returns a new {{.kind.String}}ClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
{{ if  .kind.IsNamespaced -}}
// - has the kcpcache.ClusterAndNamespaceIndex as an index
{{end -}}
func New{{.kind.String}}ClusterLister(indexer cache.Indexer) *{{.kind.String | lowerFirst}}ClusterLister {
	return &{{.kind.String | lowerFirst}}ClusterLister{indexer: indexer}
}

// List lists all {{.kind.Plural}} in the indexer across all workspaces.
func (s *{{.kind.String | lowerFirst}}ClusterLister) List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*{{.group.PackageAlias}}.{{.kind.String}}))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get {{.kind.Plural}}.
{{if not .useUpstreamInterfaces -}}
func (s *{{.kind.String | lowerFirst}}ClusterLister) Cluster(clusterName logicalcluster.Name) {{.kind.String}}Lister {
{{else -}}
func (s *{{.kind.String | lowerFirst}}ClusterLister) Cluster(clusterName logicalcluster.Name){{.group.PackageAlias}}listers.{{.kind.String}}Lister {
{{end -}}
	return &{{.kind.String | lowerFirst}}Lister{indexer: s.indexer, clusterName: clusterName}
}

{{if not .useUpstreamInterfaces -}}
{{ if  not .kind.IsNamespaced -}}
// {{.kind.String}}Lister can list all {{.kind.Plural}}, or get one in particular.
{{else -}}
// {{.kind.String}}Lister can list {{.kind.Plural}} across all namespaces, or scope down to a {{.kind.String}}NamespaceLister for one namespace.
{{end -}}
// All objects returned here must be treated as read-only.
type {{.kind.String}}Lister interface {
	// List lists all {{.kind.Plural}} in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error)
{{ if  not .kind.IsNamespaced -}}
	// Get retrieves the {{.kind.String}} from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error)
{{else -}}
	// {{.kind.Plural}} returns a lister that can list and get {{.kind.Plural}} in one workspace and namespace.
	{{.kind.Plural}}(namespace string) {{.kind.String}}NamespaceLister
{{end -}}
	{{.kind.String}}ListerExpansion
}
{{end -}}

{{if not .useUpstreamInterfaces -}}
// {{.kind.String | lowerFirst}}Lister can list all {{.kind.Plural}} inside a workspace{{ if .kind.IsNamespaced }} or scope down to a {{.kind.String}}Lister for one namespace{{end}}.
{{else -}}
// {{.kind.String | lowerFirst}}Lister implements the {{.group.PackageAlias}}listers.{{.kind.String}}Lister interface.
{{end -}}
type {{.kind.String | lowerFirst}}Lister struct {
	indexer cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all {{.kind.Plural}} in the indexer for a workspace.
func (s *{{.kind.String | lowerFirst}}Lister) List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*{{.group.PackageAlias}}.{{.kind.String}}))
	})
	return ret, err
}

{{ if  not .kind.IsNamespaced -}}
// Get retrieves the {{.kind.String}} from the indexer for a given workspace and name.
func (s *{{.kind.String | lowerFirst}}Lister) Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound({{.group.PackageAlias}}.Resource("{{.kind.Plural | toLower}}"), name)
	}
	return obj.(*{{.group.PackageAlias}}.{{.kind.String}}), nil
}
{{ else -}}
// {{.kind.Plural}} returns an object that can list and get {{.kind.Plural}} in one namespace.
{{if not .useUpstreamInterfaces -}}
func (s *{{.kind.String | lowerFirst}}Lister) {{.kind.Plural}}(namespace string) {{.kind.String}}NamespaceLister {
{{else -}}
func (s *{{.kind.String | lowerFirst}}Lister) {{.kind.Plural}}(namespace string) {{.group.PackageAlias}}listers.{{.kind.String}}NamespaceLister {
{{end -}}
	return &{{.kind.String | lowerFirst}}NamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

{{if not .useUpstreamInterfaces -}}
// {{.kind.String | lowerFirst}}NamespaceLister helps list and get {{.kind.Plural}}.
// All objects returned here must be treated as read-only.
type {{.kind.String}}NamespaceLister interface {
	// List lists all {{.kind.Plural}} in the workspace and namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error)
	// Get retrieves the {{.kind.String}} from the indexer for a given workspace, namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error)
	{{.kind.String}}NamespaceListerExpansion
}
{{end -}}

{{ if not .useUpstreamInterfaces -}}
// {{.kind.String | lowerFirst}}NamespaceLister helps list and get {{.kind.Plural}}.
// All objects returned here must be treated as read-only.
{{ end -}}
{{ if .useUpstreamInterfaces -}}
// {{.kind.String | lowerFirst}}NamespaceLister implements the {{.group.PackageAlias}}listers.{{.kind.String}}NamespaceLister interface.
{{ end -}}
type {{.kind.String | lowerFirst}}NamespaceLister struct {
	indexer   cache.Indexer
	clusterName   logicalcluster.Name
	namespace string
}

// List lists all {{.kind.Plural}} in the indexer for a given workspace and namespace.
func (s *{{.kind.String | lowerFirst}}NamespaceLister) List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*{{.group.PackageAlias}}.{{.kind.String}}))
	})
	return ret, err
}

// Get retrieves the {{.kind.String}} from the indexer for a given workspace, namespace and name.
func (s *{{.kind.String | lowerFirst}}NamespaceLister) Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound({{.group.PackageAlias}}.Resource("{{.kind.Plural | toLower}}"), name)
	}
	return obj.(*{{.group.PackageAlias}}.{{.kind.String}}), nil
}
{{ end -}}

{{if not .useUpstreamInterfaces -}}
// New{{.kind.String}}Lister returns a new {{.kind.String}}Lister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
{{ if  .kind.IsNamespaced -}}
// - has the cache.NamespaceIndex as an index
{{end -}}
func New{{.kind.String}}Lister(indexer cache.Indexer) *{{.kind.String | lowerFirst}}ScopedLister {
	return &{{.kind.String | lowerFirst}}ScopedLister{indexer: indexer}
}

// {{.kind.String | lowerFirst}}ScopedLister can list all {{.kind.Plural}} inside a workspace{{ if .kind.IsNamespaced }} or scope down to a {{.kind.String}}Lister for one namespace{{end}}.
type {{.kind.String | lowerFirst}}ScopedLister struct {
	indexer cache.Indexer
}

// List lists all {{.kind.Plural}} in the indexer for a workspace.
func (s *{{.kind.String | lowerFirst}}ScopedLister) List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*{{.group.PackageAlias}}.{{.kind.String}}))
	})
	return ret, err
}

{{ if  not .kind.IsNamespaced -}}
// Get retrieves the {{.kind.String}} from the indexer for a given workspace and name.
func (s *{{.kind.String | lowerFirst}}ScopedLister) Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound({{.group.PackageAlias}}.Resource("{{.kind.Plural | toLower}}"), name)
	}
	return obj.(*{{.group.PackageAlias}}.{{.kind.String}}), nil
}
{{ else -}}
// {{.kind.Plural}} returns an object that can list and get {{.kind.Plural}} in one namespace.
func (s *{{.kind.String | lowerFirst}}ScopedLister) {{.kind.Plural}}(namespace string) {{.kind.String}}NamespaceLister {
	return &{{.kind.String | lowerFirst}}ScopedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// {{.kind.String | lowerFirst}}ScopedNamespaceLister helps list and get {{.kind.Plural}}.
type {{.kind.String | lowerFirst}}ScopedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all {{.kind.Plural}} in the indexer for a given workspace and namespace.
func (s *{{.kind.String | lowerFirst}}ScopedNamespaceLister) List(selector labels.Selector) (ret []*{{.group.PackageAlias}}.{{.kind.String}}, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*{{.group.PackageAlias}}.{{.kind.String}}))
	})
	return ret, err
}

// Get retrieves the {{.kind.String}} from the indexer for a given workspace, namespace and name.
func (s *{{.kind.String | lowerFirst}}ScopedNamespaceLister) Get(name string) (*{{.group.PackageAlias}}.{{.kind.String}}, error) {
	key := s.namespace + "/" + name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound({{.group.PackageAlias}}.Resource("{{.kind.Plural | toLower}}"), name)
	}
	return obj.(*{{.group.PackageAlias}}.{{.kind.String}}), nil
}
{{ end -}}
{{end -}}
`
