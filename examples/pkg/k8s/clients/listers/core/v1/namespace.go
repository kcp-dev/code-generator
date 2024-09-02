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

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	corev1 "acme.corp/pkg/apis/core/v1"
)

// NamespaceClusterLister can list Namespaces across all workspaces, or scope down to a NamespaceLister for one workspace.
// All objects returned here must be treated as read-only.
type NamespaceClusterLister interface {
	// List lists all Namespaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*corev1.Namespace, err error)
	// Cluster returns a lister that can list and get Namespaces in one workspace.
	Cluster(clusterName logicalcluster.Name) NamespaceLister
	NamespaceClusterListerExpansion
}

type namespaceClusterLister struct {
	indexer cache.Indexer
}

// NewNamespaceClusterLister returns a new NamespaceClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewNamespaceClusterLister(indexer cache.Indexer) *namespaceClusterLister {
	return &namespaceClusterLister{indexer: indexer}
}

// List lists all Namespaces in the indexer across all workspaces.
func (s *namespaceClusterLister) List(selector labels.Selector) (ret []*corev1.Namespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*corev1.Namespace))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Namespaces.
func (s *namespaceClusterLister) Cluster(clusterName logicalcluster.Name) NamespaceLister {
	return &namespaceLister{indexer: s.indexer, clusterName: clusterName}
}

// NamespaceLister can list all Namespaces, or get one in particular.
// All objects returned here must be treated as read-only.
type NamespaceLister interface {
	// List lists all Namespaces in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*corev1.Namespace, err error)
	// Get retrieves the Namespace from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*corev1.Namespace, error)
	NamespaceListerExpansion
}

// namespaceLister can list all Namespaces inside a workspace.
type namespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all Namespaces in the indexer for a workspace.
func (s *namespaceLister) List(selector labels.Selector) (ret []*corev1.Namespace, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.Namespace))
	})
	return ret, err
}

// Get retrieves the Namespace from the indexer for a given workspace and name.
func (s *namespaceLister) Get(name string) (*corev1.Namespace, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(corev1.Resource("namespaces"), name)
	}
	return obj.(*corev1.Namespace), nil
}

// NewNamespaceLister returns a new NamespaceLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewNamespaceLister(indexer cache.Indexer) *namespaceScopedLister {
	return &namespaceScopedLister{indexer: indexer}
}

// namespaceScopedLister can list all Namespaces inside a workspace.
type namespaceScopedLister struct {
	indexer cache.Indexer
}

// List lists all Namespaces in the indexer for a workspace.
func (s *namespaceScopedLister) List(selector labels.Selector) (ret []*corev1.Namespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.Namespace))
	})
	return ret, err
}

// Get retrieves the Namespace from the indexer for a given workspace and name.
func (s *namespaceScopedLister) Get(name string) (*corev1.Namespace, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(corev1.Resource("namespaces"), name)
	}
	return obj.(*corev1.Namespace), nil
}
