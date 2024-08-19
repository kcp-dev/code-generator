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

	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	rbacv1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/rbac/v1alpha1"
)

// FakeClusterRoleBindings implements ClusterRoleBindingInterface
type FakeClusterRoleBindings struct {
	Fake *FakeRbacV1alpha1
}

var clusterrolebindingsResource = v1alpha1.SchemeGroupVersion.WithResource("clusterrolebindings")

var clusterrolebindingsKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterRoleBinding")

// Get takes name of the clusterRoleBinding, and returns the corresponding clusterRoleBinding object, and an error if there is any.
func (c *FakeClusterRoleBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterRoleBinding, err error) {
	emptyResult := &v1alpha1.ClusterRoleBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(clusterrolebindingsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRoleBinding), err
}

// List takes label and field selectors, and returns the list of ClusterRoleBindings that match those selectors.
func (c *FakeClusterRoleBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterRoleBindingList, err error) {
	emptyResult := &v1alpha1.ClusterRoleBindingList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(clusterrolebindingsResource, clusterrolebindingsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterRoleBindingList{ListMeta: obj.(*v1alpha1.ClusterRoleBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRoleBindings.
func (c *FakeClusterRoleBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(clusterrolebindingsResource, opts))
}

// Create takes the representation of a clusterRoleBinding and creates it.  Returns the server's representation of the clusterRoleBinding, and an error, if there is any.
func (c *FakeClusterRoleBindings) Create(ctx context.Context, clusterRoleBinding *v1alpha1.ClusterRoleBinding, opts v1.CreateOptions) (result *v1alpha1.ClusterRoleBinding, err error) {
	emptyResult := &v1alpha1.ClusterRoleBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(clusterrolebindingsResource, clusterRoleBinding, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRoleBinding), err
}

// Update takes the representation of a clusterRoleBinding and updates it. Returns the server's representation of the clusterRoleBinding, and an error, if there is any.
func (c *FakeClusterRoleBindings) Update(ctx context.Context, clusterRoleBinding *v1alpha1.ClusterRoleBinding, opts v1.UpdateOptions) (result *v1alpha1.ClusterRoleBinding, err error) {
	emptyResult := &v1alpha1.ClusterRoleBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(clusterrolebindingsResource, clusterRoleBinding, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRoleBinding), err
}

// Delete takes name of the clusterRoleBinding and deletes it. Returns an error if one occurs.
func (c *FakeClusterRoleBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterrolebindingsResource, name, opts), &v1alpha1.ClusterRoleBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRoleBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(clusterrolebindingsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterRoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched clusterRoleBinding.
func (c *FakeClusterRoleBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterRoleBinding, err error) {
	emptyResult := &v1alpha1.ClusterRoleBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusterrolebindingsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRoleBinding), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterRoleBinding.
func (c *FakeClusterRoleBindings) Apply(ctx context.Context, clusterRoleBinding *rbacv1alpha1.ClusterRoleBindingApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterRoleBinding, err error) {
	if clusterRoleBinding == nil {
		return nil, fmt.Errorf("clusterRoleBinding provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterRoleBinding)
	if err != nil {
		return nil, err
	}
	name := clusterRoleBinding.Name
	if name == nil {
		return nil, fmt.Errorf("clusterRoleBinding.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.ClusterRoleBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusterrolebindingsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRoleBinding), err
}
