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
	apimachinerycache "github.com/kcp-dev/apimachinery/pkg/cache"
	existinginterfacesv1 "github.com/kcp-dev/code-generator/examples/pkg/apis/existinginterfaces/v1"
	existinginterfacesv1listers "github.com/kcp-dev/code-generator/examples/pkg/generated/listers/existinginterfaces/v1"
	"github.com/kcp-dev/logicalcluster"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

var _ existinginterfacesv1listers.TestTypeLister = &TestTypeClusterLister{}

// TestTypeClusterLister implements the existinginterfacesv1listers.TestTypeLister interface.
type TestTypeClusterLister struct {
	indexer cache.Indexer
}

// NewTestTypeClusterLister returns a new TestTypeClusterLister.
func NewTestTypeClusterLister(indexer cache.Indexer) existinginterfacesv1listers.TestTypeLister {
	return &TestTypeClusterLister{indexer: indexer}
}

// List lists all existinginterfacesv1.TestType in the indexer.
func (s TestTypeClusterLister) List(selector labels.Selector) (ret []*existinginterfacesv1.TestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*existinginterfacesv1.TestType))
	})
	return ret, err
}

// TestTypes returns an object that can list and get existinginterfacesv1.TestType.
func (s TestTypeClusterLister) TestTypes(namespace string) existinginterfacesv1listers.TestTypeNamespaceLister {
	panic("Calling 'TestTypes' is not supported before scoping lister to a workspace")
}

// Cluster returns an object that can list and get existinginterfacesv1.TestType.

func (s TestTypeClusterLister) Cluster(cluster logicalcluster.Name) existinginterfacesv1listers.TestTypeLister {
	return &TestTypeLister{indexer: s.indexer, cluster: cluster}
}

// TestTypeLister implements the existinginterfacesv1listers.TestTypeLister interface.
type TestTypeLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all existinginterfacesv1.TestType in the indexer.
func (s TestTypeLister) List(selector labels.Selector) (ret []*existinginterfacesv1.TestType, err error) {
	selectAll := selector == nil || selector.Empty()

	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), "", "")
	list, err := s.indexer.ByIndex(apimachinerycache.ClusterIndexName, key)
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*existinginterfacesv1.TestType)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// TestTypes returns an object that can list and get existinginterfacesv1.TestType.
func (s TestTypeLister) TestTypes(namespace string) existinginterfacesv1listers.TestTypeNamespaceLister {
	return &TestTypeNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// TestTypeNamespaceLister implements the existinginterfacesv1listers.TestTypeNamespaceLister interface.
type TestTypeNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all existinginterfacesv1.TestType in the indexer for a given namespace.
func (s TestTypeNamespaceLister) List(selector labels.Selector) (ret []*existinginterfacesv1.TestType, err error) {
	selectAll := selector == nil || selector.Empty()

	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), s.namespace, "")
	list, err := s.indexer.ByIndex(apimachinerycache.ClusterAndNamespaceIndexName, key)
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*existinginterfacesv1.TestType)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}
	return ret, err
}

// Get retrieves the existinginterfacesv1.TestType from the indexer for a given namespace and name.
func (s TestTypeNamespaceLister) Get(name string) (*existinginterfacesv1.TestType, error) {
	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(existinginterfacesv1.Resource("TestType"), name)
	}
	return obj.(*existinginterfacesv1.TestType), nil
}