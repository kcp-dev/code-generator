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

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	v1 "k8s.io/api/node/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	nodev1 "k8s.io/code-generator/examples/upstream/applyconfiguration/node/v1"
)

// runtimeClassesClusterClient implements runtimeClassInterface
type runtimeClassesClusterClient struct {
	*kcptesting.Fake
}

var runtimeclassesResource = v1.SchemeGroupVersion.WithResource("runtimeclasses")

var runtimeclassesKind = v1.SchemeGroupVersion.WithKind("RuntimeClass")

// Get takes name of the runtimeClass, and returns the corresponding runtimeClass object, and an error if there is any.
func (c *runtimeClassesClusterClient) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RuntimeClass, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(runtimeclassesResource, c.ClusterPath, c.Namespace, name), &v1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RuntimeClass), err
}

// List takes label and field selectors, and returns the list of RuntimeClasses that match those selectors.
func (c *runtimeClassesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RuntimeClassList, err error) {
	emptyResult := &v1.RuntimeClassList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(runtimeclassesResource, runtimeclassesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RuntimeClassList{ListMeta: obj.(*v1.RuntimeClassList).ListMeta}
	for _, item := range obj.(*v1.RuntimeClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runtimeClasses.
func (c *runtimeClassesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(runtimeclassesResource, opts))
}

// Create takes the representation of a runtimeClass and creates it.  Returns the server's representation of the runtimeClass, and an error, if there is any.
func (c *runtimeClassesClusterClient) Create(ctx context.Context, runtimeClass *v1.RuntimeClass, opts metav1.CreateOptions) (result *v1.RuntimeClass, err error) {
	emptyResult := &v1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(runtimeclassesResource, runtimeClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RuntimeClass), err
}

// Update takes the representation of a runtimeClass and updates it. Returns the server's representation of the runtimeClass, and an error, if there is any.
func (c *runtimeClassesClusterClient) Update(ctx context.Context, runtimeClass *v1.RuntimeClass, opts metav1.UpdateOptions) (result *v1.RuntimeClass, err error) {
	emptyResult := &v1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(runtimeclassesResource, runtimeClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RuntimeClass), err
}

// Delete takes name of the runtimeClass and deletes it. Returns an error if one occurs.
func (c *runtimeClassesClusterClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(runtimeclassesResource, name, opts), &v1.RuntimeClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *runtimeClassesClusterClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(runtimeclassesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RuntimeClassList{})
	return err
}

// Patch applies the patch and returns the patched runtimeClass.
func (c *runtimeClassesClusterClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RuntimeClass, err error) {
	emptyResult := &v1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(runtimeclassesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RuntimeClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied runtimeClass.
func (c *runtimeClassesClusterClient) Apply(ctx context.Context, runtimeClass *nodev1.RuntimeClassApplyConfiguration, opts metav1.ApplyOptions) (result *v1.RuntimeClass, err error) {
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
	emptyResult := &v1.RuntimeClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(runtimeclassesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RuntimeClass), err
}
