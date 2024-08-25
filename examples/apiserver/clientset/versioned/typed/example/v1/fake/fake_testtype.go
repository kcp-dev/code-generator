/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"
	v1 "k8s.io/code-generator/examples/apiserver/apis/example/v1"
	kcp "k8s.io/code-generator/examples/apiserver/clientset/versioned/typed/example/v1"
)

var testtypesResource = v1.SchemeGroupVersion.WithResource("testtypes")

var testtypesKind = v1.SchemeGroupVersion.WithKind("TestType")

// testTypesClusterClient implements testTypeInterface
type testTypesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *testTypesClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.TestTypeNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &testTypesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of TestTypes that match those selectors.
func (c *testTypesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TestTypeList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(testtypesResource, testtypesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1.TestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.TestTypeList{ListMeta: obj.(*v1.TestTypeList).ListMeta}
	for _, item := range obj.(*v1.TestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested testTypes across all clusters.
func (c *testTypesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(testtypesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type testTypesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *testTypesNamespacer) Namespace(namespace string) upstreamexamplev1client.TestTypeInterface {
	return &configMapsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type testTypesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *testTypesClient) Create(ctx context.Context, testType *v1.TestType, opts metav1.CreateOptions) (*v1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(testtypesResource, c.ClusterPath, c.Namespace, testType), &v1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.TestType), err
}

func (c *testTypesClient) Update(ctx context.Context, testType *v1.TestType, opts metav1.CreateOptions) (*v1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(testtypesResource, c.ClusterPath, c.Namespace, testType), &v1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.TestType), err
}

func (c *testTypesClient) UpdateStatus(ctx context.Context, testType *v1.TestType, opts metav1.CreateOptions) (*v1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(testtypesResource, c.ClusterPath, "status", c.Namespace, testType), &v1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.TestType), err
}

func (c *testTypesClient) Delete(ctx context.Context, name string, opts metav1.CreateOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(testtypesResource, c.ClusterPath, c.Namespace, name, opts), &v1.TestType{})
	return err
}

func (c *testTypesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(testtypesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1.TestTypeList{})
	return err
}

func (c *testTypesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(testtypesResource, c.ClusterPath, c.Namespace, name), &v1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.TestType), err
}

// List takes label and field selectors, and returns the list of v1.TestType that match those selectors.
func (c *testTypesClient) List(ctx context.Context, opts metav1.ListOptions) (*v1.TestTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(testtypesResource, testtypesKind, c.ClusterPath, c.Namespace, opts), &v1.TestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.TestTypeList{ListMeta: obj.(*v1.TestTypeList).ListMeta}
	for _, item := range obj.(*v1.TestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *testTypesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(testtypesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *testTypesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(testtypesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.TestType), err
}
