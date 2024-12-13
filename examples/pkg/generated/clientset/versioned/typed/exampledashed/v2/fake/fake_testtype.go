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

// Code generated by client-gen-v0.31. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	examplev1 "acme.corp/pkg/apis/example/v1"
	v2 "acme.corp/pkg/apisv2/exampledashed/v2"
)

// FakeTestTypes implements TestTypeInterface
type FakeTestTypes struct {
	Fake *FakeExampleDashedV2
	ns   string
}

var testtypesResource = v2.SchemeGroupVersion.WithResource("testtypes")

var testtypesKind = v2.SchemeGroupVersion.WithKind("TestType")

// Get takes name of the testType, and returns the corresponding testType object, and an error if there is any.
func (c *FakeTestTypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.TestType, err error) {
	emptyResult := &v2.TestType{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(testtypesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2.TestType), err
}

// List takes label and field selectors, and returns the list of TestTypes that match those selectors.
func (c *FakeTestTypes) List(ctx context.Context, opts v1.ListOptions) (result *v2.TestTypeList, err error) {
	emptyResult := &v2.TestTypeList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(testtypesResource, testtypesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.TestTypeList{ListMeta: obj.(*v2.TestTypeList).ListMeta}
	for _, item := range obj.(*v2.TestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested testTypes.
func (c *FakeTestTypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(testtypesResource, c.ns, opts))

}

// Create takes the representation of a testType and creates it.  Returns the server's representation of the testType, and an error, if there is any.
func (c *FakeTestTypes) Create(ctx context.Context, testType *v2.TestType, opts v1.CreateOptions) (result *v2.TestType, err error) {
	emptyResult := &v2.TestType{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(testtypesResource, c.ns, testType, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2.TestType), err
}

// Update takes the representation of a testType and updates it. Returns the server's representation of the testType, and an error, if there is any.
func (c *FakeTestTypes) Update(ctx context.Context, testType *v2.TestType, opts v1.UpdateOptions) (result *v2.TestType, err error) {
	emptyResult := &v2.TestType{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(testtypesResource, c.ns, testType, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2.TestType), err
}

// Delete takes name of the testType and deletes it. Returns an error if one occurs.
func (c *FakeTestTypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(testtypesResource, c.ns, name, opts), &v2.TestType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTestTypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(testtypesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v2.TestTypeList{})
	return err
}

// Patch applies the patch and returns the patched testType.
func (c *FakeTestTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.TestType, err error) {
	emptyResult := &v2.TestType{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(testtypesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v2.TestType), err
}

// CreateField takes the representation of a field and creates it.  Returns the server's representation of the field, and an error, if there is any.
func (c *FakeTestTypes) CreateField(ctx context.Context, testTypeName string, field *examplev1.Field, opts v1.CreateOptions) (result *examplev1.Field, err error) {
	emptyResult := &examplev1.Field{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateSubresourceActionWithOptions(testtypesResource, testTypeName, "field", c.ns, field, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*examplev1.Field), err
}

// UpdateField takes the representation of a field and updates it. Returns the server's representation of the field, and an error, if there is any.
func (c *FakeTestTypes) UpdateField(ctx context.Context, testTypeName string, field *examplev1.Field, opts v1.UpdateOptions) (result *examplev1.Field, err error) {
	emptyResult := &examplev1.Field{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(testtypesResource, "field", c.ns, field, opts), &examplev1.Field{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*examplev1.Field), err
}

// GetField takes name of the testType, and returns the corresponding field object, and an error if there is any.
func (c *FakeTestTypes) GetField(ctx context.Context, testTypeName string, options v1.GetOptions) (result *examplev1.Field, err error) {
	emptyResult := &examplev1.Field{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(testtypesResource, c.ns, "field", testTypeName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*examplev1.Field), err
}
