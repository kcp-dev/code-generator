/*
Copyright The KCP Authors.

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

package internal

// TODO(kcp): Add conditional branches for cluster-scoped listers; i.e. No .Namespace() method.
const listerTemplate = `
package {{.Version}}

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"

	{{.Group}}{{.Version}} "{{.SrcPackage}}"
)

{{$plural := plural ".Kind"}}
{{$fqkind := .Group.Version..Kind}}

// {{.Kind}}Lister helps list {{$plural}}.
// All objects returned here must be treated as read-only.
type {{.Kind}}Lister interface {
	// List lists all {{$plural}} in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*{{$fqkind}}, err error)

	// {{$plural}} returns an object that can list and get {{$plural}}.
	// {{$plural}}(cluster logicalcluster.LogicalCluster) {{.Kind}}ClusterLister

	// Cluster returns an object that can list and get {{$plural}} from the given logical cluster.
	Cluster(cluster logicalcluster.LogicalCluster) {{.Kind}}ClusterLister

	// Note(kcp): Workspace-capable Lister implementation doesn't support support expansions.
	// {{.Kind}}ListerExpansion
}

{{$camel := camel ".Kind"}}

// {{$camel}}Lister implements the {{.Kind}}Lister interface.
type {{$camel}}Lister struct {
	indexer cache.Indexer
}

// New{{.Kind}}Lister returns a new {{.Kind}}Lister.
func New{{.Kind}}Lister(indexer cache.Indexer) {{.Kind}}Lister {
	return &{{$camel}}Lister{indexer: indexer}
}

// List lists all {{$plural}} in the indexer.
func (s *{{$camel}}Lister) List(selector labels.Selector) (ret []*{{$fqkind}}, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*{{$fqkind}}))
	})
	return ret, err
}

// Cluster returns an object that can list and get {{$plural}}.
func (s *{{$camel}}Lister) Cluster(cluster logicalcluster.LogicalCluster) {{.Kind}}ClusterLister {
	return &{{$camel}}ClusterLister{indexer: s.indexer, cluster: cluster}
}

// {{.Kind}}Lister helps list {{$plural}}.
// All objects returned here must be treated as read-only.
type {{.Kind}}ClusterLister interface {
	// List lists all {{$plural}} in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*{{$fqkind}}, err error)
	// {{$plural}} returns an object that can list and get {{$plural}}.
	{{$plural}}(namespace string) {{.Kind}}NamespaceLister
	// Note(kcp): Workspace-capable Lister implementation doesn't support support expansions.
	// {{.Kind}}ListerExpansion
}

// {{$camel}}ClusterLister implements the {{.Kind}}Lister interface.
type {{$camel}}ClusterLister struct {
	indexer cache.Indexer
	cluster logicalcluster.LogicalCluster
}

// List lists all {{$plural}} in the indexer.
func (c *{{camel}}ClusterLister) List(selector labels.Selector) (ret []*{{fqkind}}, err error) {
	list, err := c.indexer.ByIndex(ClusterIndexName, c.cluster.String())
	if err != nil {
		return nil, err
	}

	if selector == nil {
		selector = labels.Everything()
	}
	for i := range list {
		obj := list[i].(*{{$fqkind}})
		if selector.Matches(labels.Set(obj.GetLabels())) {
			ret = append(ret, obj)
		}
	}

	return ret, err
}

// {{$plural}} returns an object that can list and get {{$plural}}.
func (c *{{$camel}}ClusterLister) {{$plural}}(namespace string) {{.Kind}}NamespaceLister {
	return {{$camel}}NamespaceLister{indexer: c.indexer, cluster: c.cluster, namespace: namespace}
}

// ConfigMapNamespaceLister helps list and get {{$plural}}.
// All objects returned here must be treated as read-only.
type ConfigMapNamespaceLister interface {
	// List lists all {{$plural}} in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*{{$fqkind}}, err error)
	// Get retrieves the ConfigMap from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*{{$fqkind}}, error)
	// Note(kcp): Workspace-capable Lister implementation doesn't support support expansions.
	// ConfigMapNamespaceListerExpansion
}

// {{$camel}}NamespaceLister implements the ConfigMapNamespaceLister
// interface.
type {{$camel}}NamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.LogicalCluster
	namespace string
}

// List lists all {{$plural}} in the indexer for a given namespace.
func (c {{$camel}}NamespaceLister) List(selector labels.Selector) (ret []*{{$fqkind}}, err error) {
	list, err := c.indexer.Index(ClusterAndNamespaceIndexName, &metav1.ObjectMeta{
		ZZZ_DeprecatedClusterName: c.cluster.String(),
		Namespace:                 c.namespace,
	})
	if err != nil {
		return nil, err
	}

	if selector == nil {
		selector = labels.Everything()
	}
	for i := range list {
		cm := list[i].(*{{$fqkind}})
		if selector.Matches(labels.Set(cm.GetLabels())) {
			ret = append(ret, cm)
		}
	}

	return ret, err
}

// Get retrieves the ConfigMap from the indexer for a given namespace and name.
func (c {{$camel}}NamespaceLister) Get(name string) (*{{$fqkind}}, error) {
	meta := &metav1.ObjectMeta{
		ZZZ_DeprecatedClusterName: c.cluster.String(),
		Namespace:                 c.namespace,
		Name:                      name,
	}
	obj, exists, err := c.indexer.Get(meta)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound({{.Resource}}.{{.Group}}, name)
	}
	return obj.(*{{$fqkind}}), nil
}
`

// Do we make this configurable?
// TODO: reformat it to be able to configure Copyright Year
const HeaderText = `
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

// Code auto-generated. DO NOT EDIT.
// +build !ignore_autogenerated
`
