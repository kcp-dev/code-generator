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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1 "k8s.io/code-generator/examples/apiserver/apis/core/v1"
)

// testTypesClusterClient implements testTypeInterface
type testTypesClusterClient struct {
	*kcptesting.Fake
}

var testtypesResource = v1.SchemeGroupVersion.WithResource("testtypes")

var testtypesKind = v1.SchemeGroupVersion.WithKind("TestType")

// Get takes name of the testType, and returns the corresponding testType object, and an error if there is any.
func (c *testTypesClusterClient) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TestType, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(testtypesResource, c.ClusterPath, c.Namespace, name), &v1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.TestType), err
}

// List takes label and field selectors, and returns the list of TestTypes that match those selectors.
func (c *testTypesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TestTypeList, err error) {
	emptyResult := &v1.TestTypeList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(testtypesResource, testtypesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
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

// Watch returns a watch.Interface that watches the requested testTypes.
func (c *testTypesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(testtypesResource, c.ns, opts))

}

// Create takes the representation of a testType and creates it.  Returns the server's representation of the testType, and an error, if there is any.
func (c *testTypesClusterClient) Create(ctx context.Context, testType *v1.TestType, opts metav1.CreateOptions) (result *v1.TestType, err error) {
	emptyResult := &v1.TestType{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(testtypesResource, c.ns, testType, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.TestType), err
}

// Update takes the representation of a testType and updates it. Returns the server's representation of the testType, and an error, if there is any.
func (c *testTypesClusterClient) Update(ctx context.Context, testType *v1.TestType, opts metav1.UpdateOptions) (result *v1.TestType, err error) {
	emptyResult := &v1.TestType{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(testtypesResource, c.ns, testType, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.TestType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *testTypesClusterClient) UpdateStatus(ctx context.Context, testType *v1.TestType, opts metav1.UpdateOptions) (result *v1.TestType, err error) {
	emptyResult := &v1.TestType{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(testtypesResource, "status", c.ns, testType, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.TestType), err
}

// Delete takes name of the testType and deletes it. Returns an error if one occurs.
func (c *testTypesClusterClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(testtypesResource, c.ns, name, opts), &v1.TestType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *testTypesClusterClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(testtypesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.TestTypeList{})
	return err
}

// Patch applies the patch and returns the patched testType.
func (c *testTypesClusterClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TestType, err error) {
	emptyResult := &v1.TestType{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(testtypesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.TestType), err
}
