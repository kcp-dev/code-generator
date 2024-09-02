//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	extensionsv1beta1 "acme.corp/pkg/apis/extensions/v1beta1"
)

// IngressClusterLister can list Ingresses across all workspaces, or scope down to a IngressLister for one workspace.
// All objects returned here must be treated as read-only.
type IngressClusterLister interface {
	// List lists all Ingresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*extensionsv1beta1.Ingress, err error)
	// Cluster returns a lister that can list and get Ingresses in one workspace.
	Cluster(clusterName logicalcluster.Name) IngressLister
	IngressClusterListerExpansion
}

type ingressClusterLister struct {
	indexer cache.Indexer
}

// NewIngressClusterLister returns a new IngressClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewIngressClusterLister(indexer cache.Indexer) *ingressClusterLister {
	return &ingressClusterLister{indexer: indexer}
}

// List lists all Ingresses in the indexer across all workspaces.
func (s *ingressClusterLister) List(selector labels.Selector) (ret []*extensionsv1beta1.Ingress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*extensionsv1beta1.Ingress))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Ingresses.
func (s *ingressClusterLister) Cluster(clusterName logicalcluster.Name) IngressLister {
	return &ingressLister{indexer: s.indexer, clusterName: clusterName}
}

// IngressLister can list Ingresses across all namespaces, or scope down to a IngressNamespaceLister for one namespace.
// All objects returned here must be treated as read-only.
type IngressLister interface {
	// List lists all Ingresses in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*extensionsv1beta1.Ingress, err error)
	// Ingresses returns a lister that can list and get Ingresses in one workspace and namespace.
	Ingresses(namespace string) IngressNamespaceLister
	IngressListerExpansion
}

// ingressLister can list all Ingresses inside a workspace or scope down to a IngressLister for one namespace.
type ingressLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all Ingresses in the indexer for a workspace.
func (s *ingressLister) List(selector labels.Selector) (ret []*extensionsv1beta1.Ingress, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*extensionsv1beta1.Ingress))
	})
	return ret, err
}

// Ingresses returns an object that can list and get Ingresses in one namespace.
func (s *ingressLister) Ingresses(namespace string) IngressNamespaceLister {
	return &ingressNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// ingressNamespaceLister helps list and get Ingresses.
// All objects returned here must be treated as read-only.
type IngressNamespaceLister interface {
	// List lists all Ingresses in the workspace and namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*extensionsv1beta1.Ingress, err error)
	// Get retrieves the Ingress from the indexer for a given workspace, namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*extensionsv1beta1.Ingress, error)
	IngressNamespaceListerExpansion
}

// ingressNamespaceLister helps list and get Ingresses.
// All objects returned here must be treated as read-only.
type ingressNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all Ingresses in the indexer for a given workspace and namespace.
func (s *ingressNamespaceLister) List(selector labels.Selector) (ret []*extensionsv1beta1.Ingress, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*extensionsv1beta1.Ingress))
	})
	return ret, err
}

// Get retrieves the Ingress from the indexer for a given workspace, namespace and name.
func (s *ingressNamespaceLister) Get(name string) (*extensionsv1beta1.Ingress, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(extensionsv1beta1.Resource("ingresses"), name)
	}
	return obj.(*extensionsv1beta1.Ingress), nil
}

// NewIngressLister returns a new IngressLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
// - has the cache.NamespaceIndex as an index
func NewIngressLister(indexer cache.Indexer) *ingressScopedLister {
	return &ingressScopedLister{indexer: indexer}
}

// ingressScopedLister can list all Ingresses inside a workspace or scope down to a IngressLister for one namespace.
type ingressScopedLister struct {
	indexer cache.Indexer
}

// List lists all Ingresses in the indexer for a workspace.
func (s *ingressScopedLister) List(selector labels.Selector) (ret []*extensionsv1beta1.Ingress, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*extensionsv1beta1.Ingress))
	})
	return ret, err
}

// Ingresses returns an object that can list and get Ingresses in one namespace.
func (s *ingressScopedLister) Ingresses(namespace string) IngressNamespaceLister {
	return &ingressScopedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ingressScopedNamespaceLister helps list and get Ingresses.
type ingressScopedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ingresses in the indexer for a given workspace and namespace.
func (s *ingressScopedNamespaceLister) List(selector labels.Selector) (ret []*extensionsv1beta1.Ingress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*extensionsv1beta1.Ingress))
	})
	return ret, err
}

// Get retrieves the Ingress from the indexer for a given workspace, namespace and name.
func (s *ingressScopedNamespaceLister) Get(name string) (*extensionsv1beta1.Ingress, error) {
	key := s.namespace + "/" + name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(extensionsv1beta1.Resource("ingresses"), name)
	}
	return obj.(*extensionsv1beta1.Ingress), nil
}
