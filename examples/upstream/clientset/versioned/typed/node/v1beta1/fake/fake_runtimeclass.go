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
	json "encoding/json"
	"fmt"

	v1beta1 "k8s.io/api/node/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	nodev1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/node/v1beta1"
)

// FakeRuntimeClasses implements RuntimeClassInterface
type FakeRuntimeClasses struct {
	Fake *FakeNodeV1beta1
}

var runtimeclassesResource = v1beta1.SchemeGroupVersion.WithResource("runtimeclasses")

var runtimeclassesKind = v1beta1.SchemeGroupVersion.WithKind("RuntimeClass")

// Get takes name of the runtimeClass, and returns the corresponding runtimeClass object, and an error if there is any.
func (c *FakeRuntimeClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.RuntimeClass, err error) {
	emptyResult := &v1beta1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(runtimeclassesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.RuntimeClass), err
}

// List takes label and field selectors, and returns the list of RuntimeClasses that match those selectors.
func (c *FakeRuntimeClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RuntimeClassList, err error) {
	emptyResult := &v1beta1.RuntimeClassList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(runtimeclassesResource, runtimeclassesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RuntimeClassList{ListMeta: obj.(*v1beta1.RuntimeClassList).ListMeta}
	for _, item := range obj.(*v1beta1.RuntimeClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runtimeClasses.
func (c *FakeRuntimeClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(runtimeclassesResource, opts))
}

// Create takes the representation of a runtimeClass and creates it.  Returns the server's representation of the runtimeClass, and an error, if there is any.
func (c *FakeRuntimeClasses) Create(ctx context.Context, runtimeClass *v1beta1.RuntimeClass, opts v1.CreateOptions) (result *v1beta1.RuntimeClass, err error) {
	emptyResult := &v1beta1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(runtimeclassesResource, runtimeClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.RuntimeClass), err
}

// Update takes the representation of a runtimeClass and updates it. Returns the server's representation of the runtimeClass, and an error, if there is any.
func (c *FakeRuntimeClasses) Update(ctx context.Context, runtimeClass *v1beta1.RuntimeClass, opts v1.UpdateOptions) (result *v1beta1.RuntimeClass, err error) {
	emptyResult := &v1beta1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(runtimeclassesResource, runtimeClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.RuntimeClass), err
}

// Delete takes name of the runtimeClass and deletes it. Returns an error if one occurs.
func (c *FakeRuntimeClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(runtimeclassesResource, name, opts), &v1beta1.RuntimeClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRuntimeClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(runtimeclassesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.RuntimeClassList{})
	return err
}

// Patch applies the patch and returns the patched runtimeClass.
func (c *FakeRuntimeClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.RuntimeClass, err error) {
	emptyResult := &v1beta1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(runtimeclassesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.RuntimeClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied runtimeClass.
func (c *FakeRuntimeClasses) Apply(ctx context.Context, runtimeClass *nodev1beta1.RuntimeClassApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.RuntimeClass, err error) {
	if runtimeClass == nil {
		return nil, fmt.Errorf("runtimeClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(runtimeClass)
	if err != nil {
		return nil, err
	}
	name := runtimeClass.Name
	if name == nil {
		return nil, fmt.Errorf("runtimeClass.Name must be provided to Apply")
	}
	emptyResult := &v1beta1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(runtimeclassesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.RuntimeClass), err
}
