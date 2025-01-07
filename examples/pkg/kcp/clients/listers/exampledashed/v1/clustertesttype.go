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

	exampledashedv1 "acme.corp/pkg/apis/exampledashed/v1"
)

// ClusterTestTypeClusterLister can list ClusterTestTypes across all workspaces, or scope down to a ClusterTestTypeLister for one workspace.
// All objects returned here must be treated as read-only.
type ClusterTestTypeClusterLister interface {
	// List lists all ClusterTestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*exampledashedv1.ClusterTestType, err error)
	// Cluster returns a lister that can list and get ClusterTestTypes in one workspace.
	Cluster(clusterName logicalcluster.Name) ClusterTestTypeLister
	ClusterTestTypeClusterListerExpansion
}

type clusterTestTypeClusterLister struct {
	indexer cache.Indexer
}

// NewClusterTestTypeClusterLister returns a new ClusterTestTypeClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewClusterTestTypeClusterLister(indexer cache.Indexer) *clusterTestTypeClusterLister {
	return &clusterTestTypeClusterLister{indexer: indexer}
}

// List lists all ClusterTestTypes in the indexer across all workspaces.
func (s *clusterTestTypeClusterLister) List(selector labels.Selector) (ret []*exampledashedv1.ClusterTestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*exampledashedv1.ClusterTestType))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ClusterTestTypes.
func (s *clusterTestTypeClusterLister) Cluster(clusterName logicalcluster.Name) ClusterTestTypeLister {
	return &clusterTestTypeLister{indexer: s.indexer, clusterName: clusterName}
}

// ClusterTestTypeLister can list all ClusterTestTypes, or get one in particular.
// All objects returned here must be treated as read-only.
type ClusterTestTypeLister interface {
	// List lists all ClusterTestTypes in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*exampledashedv1.ClusterTestType, err error)
	// Get retrieves the ClusterTestType from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*exampledashedv1.ClusterTestType, error)
	ClusterTestTypeListerExpansion
}

// clusterTestTypeLister can list all ClusterTestTypes inside a workspace.
type clusterTestTypeLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all ClusterTestTypes in the indexer for a workspace.
func (s *clusterTestTypeLister) List(selector labels.Selector) (ret []*exampledashedv1.ClusterTestType, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*exampledashedv1.ClusterTestType))
	})
	return ret, err
}

// Get retrieves the ClusterTestType from the indexer for a given workspace and name.
func (s *clusterTestTypeLister) Get(name string) (*exampledashedv1.ClusterTestType, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(exampledashedv1.Resource("clustertesttypes"), name)
	}
	return obj.(*exampledashedv1.ClusterTestType), nil
}

// NewClusterTestTypeLister returns a new ClusterTestTypeLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewClusterTestTypeLister(indexer cache.Indexer) *clusterTestTypeScopedLister {
	return &clusterTestTypeScopedLister{indexer: indexer}
}

// clusterTestTypeScopedLister can list all ClusterTestTypes inside a workspace.
type clusterTestTypeScopedLister struct {
	indexer cache.Indexer
}

// List lists all ClusterTestTypes in the indexer for a workspace.
func (s *clusterTestTypeScopedLister) List(selector labels.Selector) (ret []*exampledashedv1.ClusterTestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*exampledashedv1.ClusterTestType))
	})
	return ret, err
}

// Get retrieves the ClusterTestType from the indexer for a given workspace and name.
func (s *clusterTestTypeScopedLister) Get(name string) (*exampledashedv1.ClusterTestType, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(exampledashedv1.Resource("clustertesttypes"), name)
	}
	return obj.(*exampledashedv1.ClusterTestType), nil
}
