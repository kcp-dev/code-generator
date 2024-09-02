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

	rbacv1 "acme.corp/pkg/apis/rbac/v1"
)

// RoleBindingClusterLister can list RoleBindings across all workspaces, or scope down to a RoleBindingLister for one workspace.
// All objects returned here must be treated as read-only.
type RoleBindingClusterLister interface {
	// List lists all RoleBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error)
	// Cluster returns a lister that can list and get RoleBindings in one workspace.
	Cluster(clusterName logicalcluster.Name) RoleBindingLister
	RoleBindingClusterListerExpansion
}

type roleBindingClusterLister struct {
	indexer cache.Indexer
}

// NewRoleBindingClusterLister returns a new RoleBindingClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewRoleBindingClusterLister(indexer cache.Indexer) *roleBindingClusterLister {
	return &roleBindingClusterLister{indexer: indexer}
}

// List lists all RoleBindings in the indexer across all workspaces.
func (s *roleBindingClusterLister) List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*rbacv1.RoleBinding))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get RoleBindings.
func (s *roleBindingClusterLister) Cluster(clusterName logicalcluster.Name) RoleBindingLister {
	return &roleBindingLister{indexer: s.indexer, clusterName: clusterName}
}

// RoleBindingLister can list RoleBindings across all namespaces, or scope down to a RoleBindingNamespaceLister for one namespace.
// All objects returned here must be treated as read-only.
type RoleBindingLister interface {
	// List lists all RoleBindings in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error)
	// RoleBindings returns a lister that can list and get RoleBindings in one workspace and namespace.
	RoleBindings(namespace string) RoleBindingNamespaceLister
	RoleBindingListerExpansion
}

// roleBindingLister can list all RoleBindings inside a workspace or scope down to a RoleBindingLister for one namespace.
type roleBindingLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all RoleBindings in the indexer for a workspace.
func (s *roleBindingLister) List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1.RoleBinding))
	})
	return ret, err
}

// RoleBindings returns an object that can list and get RoleBindings in one namespace.
func (s *roleBindingLister) RoleBindings(namespace string) RoleBindingNamespaceLister {
	return &roleBindingNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// roleBindingNamespaceLister helps list and get RoleBindings.
// All objects returned here must be treated as read-only.
type RoleBindingNamespaceLister interface {
	// List lists all RoleBindings in the workspace and namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error)
	// Get retrieves the RoleBinding from the indexer for a given workspace, namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*rbacv1.RoleBinding, error)
	RoleBindingNamespaceListerExpansion
}

// roleBindingNamespaceLister helps list and get RoleBindings.
// All objects returned here must be treated as read-only.
type roleBindingNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all RoleBindings in the indexer for a given workspace and namespace.
func (s *roleBindingNamespaceLister) List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1.RoleBinding))
	})
	return ret, err
}

// Get retrieves the RoleBinding from the indexer for a given workspace, namespace and name.
func (s *roleBindingNamespaceLister) Get(name string) (*rbacv1.RoleBinding, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbacv1.Resource("rolebindings"), name)
	}
	return obj.(*rbacv1.RoleBinding), nil
}

// NewRoleBindingLister returns a new RoleBindingLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
// - has the cache.NamespaceIndex as an index
func NewRoleBindingLister(indexer cache.Indexer) *roleBindingScopedLister {
	return &roleBindingScopedLister{indexer: indexer}
}

// roleBindingScopedLister can list all RoleBindings inside a workspace or scope down to a RoleBindingLister for one namespace.
type roleBindingScopedLister struct {
	indexer cache.Indexer
}

// List lists all RoleBindings in the indexer for a workspace.
func (s *roleBindingScopedLister) List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1.RoleBinding))
	})
	return ret, err
}

// RoleBindings returns an object that can list and get RoleBindings in one namespace.
func (s *roleBindingScopedLister) RoleBindings(namespace string) RoleBindingNamespaceLister {
	return &roleBindingScopedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// roleBindingScopedNamespaceLister helps list and get RoleBindings.
type roleBindingScopedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RoleBindings in the indexer for a given workspace and namespace.
func (s *roleBindingScopedNamespaceLister) List(selector labels.Selector) (ret []*rbacv1.RoleBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1.RoleBinding))
	})
	return ret, err
}

// Get retrieves the RoleBinding from the indexer for a given workspace, namespace and name.
func (s *roleBindingScopedNamespaceLister) Get(name string) (*rbacv1.RoleBinding, error) {
	key := s.namespace + "/" + name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbacv1.Resource("rolebindings"), name)
	}
	return obj.(*rbacv1.RoleBinding), nil
}
