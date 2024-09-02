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

	appsv1beta1 "acme.corp/pkg/apis/apps/v1beta1"
)

// StatefulSetClusterLister can list StatefulSets across all workspaces, or scope down to a StatefulSetLister for one workspace.
// All objects returned here must be treated as read-only.
type StatefulSetClusterLister interface {
	// List lists all StatefulSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error)
	// Cluster returns a lister that can list and get StatefulSets in one workspace.
	Cluster(clusterName logicalcluster.Name) StatefulSetLister
	StatefulSetClusterListerExpansion
}

type statefulSetClusterLister struct {
	indexer cache.Indexer
}

// NewStatefulSetClusterLister returns a new StatefulSetClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewStatefulSetClusterLister(indexer cache.Indexer) *statefulSetClusterLister {
	return &statefulSetClusterLister{indexer: indexer}
}

// List lists all StatefulSets in the indexer across all workspaces.
func (s *statefulSetClusterLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*appsv1beta1.StatefulSet))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get StatefulSets.
func (s *statefulSetClusterLister) Cluster(clusterName logicalcluster.Name) StatefulSetLister {
	return &statefulSetLister{indexer: s.indexer, clusterName: clusterName}
}

// StatefulSetLister can list StatefulSets across all namespaces, or scope down to a StatefulSetNamespaceLister for one namespace.
// All objects returned here must be treated as read-only.
type StatefulSetLister interface {
	// List lists all StatefulSets in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error)
	// StatefulSets returns a lister that can list and get StatefulSets in one workspace and namespace.
	StatefulSets(namespace string) StatefulSetNamespaceLister
	StatefulSetListerExpansion
}

// statefulSetLister can list all StatefulSets inside a workspace or scope down to a StatefulSetLister for one namespace.
type statefulSetLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all StatefulSets in the indexer for a workspace.
func (s *statefulSetLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*appsv1beta1.StatefulSet))
	})
	return ret, err
}

// StatefulSets returns an object that can list and get StatefulSets in one namespace.
func (s *statefulSetLister) StatefulSets(namespace string) StatefulSetNamespaceLister {
	return &statefulSetNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// statefulSetNamespaceLister helps list and get StatefulSets.
// All objects returned here must be treated as read-only.
type StatefulSetNamespaceLister interface {
	// List lists all StatefulSets in the workspace and namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error)
	// Get retrieves the StatefulSet from the indexer for a given workspace, namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*appsv1beta1.StatefulSet, error)
	StatefulSetNamespaceListerExpansion
}

// statefulSetNamespaceLister helps list and get StatefulSets.
// All objects returned here must be treated as read-only.
type statefulSetNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all StatefulSets in the indexer for a given workspace and namespace.
func (s *statefulSetNamespaceLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*appsv1beta1.StatefulSet))
	})
	return ret, err
}

// Get retrieves the StatefulSet from the indexer for a given workspace, namespace and name.
func (s *statefulSetNamespaceLister) Get(name string) (*appsv1beta1.StatefulSet, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(appsv1beta1.Resource("statefulsets"), name)
	}
	return obj.(*appsv1beta1.StatefulSet), nil
}

// NewStatefulSetLister returns a new StatefulSetLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
// - has the cache.NamespaceIndex as an index
func NewStatefulSetLister(indexer cache.Indexer) *statefulSetScopedLister {
	return &statefulSetScopedLister{indexer: indexer}
}

// statefulSetScopedLister can list all StatefulSets inside a workspace or scope down to a StatefulSetLister for one namespace.
type statefulSetScopedLister struct {
	indexer cache.Indexer
}

// List lists all StatefulSets in the indexer for a workspace.
func (s *statefulSetScopedLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*appsv1beta1.StatefulSet))
	})
	return ret, err
}

// StatefulSets returns an object that can list and get StatefulSets in one namespace.
func (s *statefulSetScopedLister) StatefulSets(namespace string) StatefulSetNamespaceLister {
	return &statefulSetScopedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// statefulSetScopedNamespaceLister helps list and get StatefulSets.
type statefulSetScopedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StatefulSets in the indexer for a given workspace and namespace.
func (s *statefulSetScopedNamespaceLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*appsv1beta1.StatefulSet))
	})
	return ret, err
}

// Get retrieves the StatefulSet from the indexer for a given workspace, namespace and name.
func (s *statefulSetScopedNamespaceLister) Get(name string) (*appsv1beta1.StatefulSet, error) {
	key := s.namespace + "/" + name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(appsv1beta1.Resource("statefulsets"), name)
	}
	return obj.(*appsv1beta1.StatefulSet), nil
}
