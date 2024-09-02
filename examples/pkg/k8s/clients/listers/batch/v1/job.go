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

	batchv1 "acme.corp/pkg/apis/batch/v1"
)

// JobClusterLister can list Jobs across all workspaces, or scope down to a JobLister for one workspace.
// All objects returned here must be treated as read-only.
type JobClusterLister interface {
	// List lists all Jobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*batchv1.Job, err error)
	// Cluster returns a lister that can list and get Jobs in one workspace.
	Cluster(clusterName logicalcluster.Name) JobLister
	JobClusterListerExpansion
}

type jobClusterLister struct {
	indexer cache.Indexer
}

// NewJobClusterLister returns a new JobClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewJobClusterLister(indexer cache.Indexer) *jobClusterLister {
	return &jobClusterLister{indexer: indexer}
}

// List lists all Jobs in the indexer across all workspaces.
func (s *jobClusterLister) List(selector labels.Selector) (ret []*batchv1.Job, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*batchv1.Job))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Jobs.
func (s *jobClusterLister) Cluster(clusterName logicalcluster.Name) JobLister {
	return &jobLister{indexer: s.indexer, clusterName: clusterName}
}

// JobLister can list Jobs across all namespaces, or scope down to a JobNamespaceLister for one namespace.
// All objects returned here must be treated as read-only.
type JobLister interface {
	// List lists all Jobs in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*batchv1.Job, err error)
	// Jobs returns a lister that can list and get Jobs in one workspace and namespace.
	Jobs(namespace string) JobNamespaceLister
	JobListerExpansion
}

// jobLister can list all Jobs inside a workspace or scope down to a JobLister for one namespace.
type jobLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all Jobs in the indexer for a workspace.
func (s *jobLister) List(selector labels.Selector) (ret []*batchv1.Job, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*batchv1.Job))
	})
	return ret, err
}

// Jobs returns an object that can list and get Jobs in one namespace.
func (s *jobLister) Jobs(namespace string) JobNamespaceLister {
	return &jobNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// jobNamespaceLister helps list and get Jobs.
// All objects returned here must be treated as read-only.
type JobNamespaceLister interface {
	// List lists all Jobs in the workspace and namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*batchv1.Job, err error)
	// Get retrieves the Job from the indexer for a given workspace, namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*batchv1.Job, error)
	JobNamespaceListerExpansion
}

// jobNamespaceLister helps list and get Jobs.
// All objects returned here must be treated as read-only.
type jobNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all Jobs in the indexer for a given workspace and namespace.
func (s *jobNamespaceLister) List(selector labels.Selector) (ret []*batchv1.Job, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*batchv1.Job))
	})
	return ret, err
}

// Get retrieves the Job from the indexer for a given workspace, namespace and name.
func (s *jobNamespaceLister) Get(name string) (*batchv1.Job, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(batchv1.Resource("jobs"), name)
	}
	return obj.(*batchv1.Job), nil
}

// NewJobLister returns a new JobLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
// - has the cache.NamespaceIndex as an index
func NewJobLister(indexer cache.Indexer) *jobScopedLister {
	return &jobScopedLister{indexer: indexer}
}

// jobScopedLister can list all Jobs inside a workspace or scope down to a JobLister for one namespace.
type jobScopedLister struct {
	indexer cache.Indexer
}

// List lists all Jobs in the indexer for a workspace.
func (s *jobScopedLister) List(selector labels.Selector) (ret []*batchv1.Job, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*batchv1.Job))
	})
	return ret, err
}

// Jobs returns an object that can list and get Jobs in one namespace.
func (s *jobScopedLister) Jobs(namespace string) JobNamespaceLister {
	return &jobScopedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// jobScopedNamespaceLister helps list and get Jobs.
type jobScopedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Jobs in the indexer for a given workspace and namespace.
func (s *jobScopedNamespaceLister) List(selector labels.Selector) (ret []*batchv1.Job, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*batchv1.Job))
	})
	return ret, err
}

// Get retrieves the Job from the indexer for a given workspace, namespace and name.
func (s *jobScopedNamespaceLister) Get(name string) (*batchv1.Job, error) {
	key := s.namespace + "/" + name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(batchv1.Resource("jobs"), name)
	}
	return obj.(*batchv1.Job), nil
}
