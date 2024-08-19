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

	v1alpha1 "k8s.io/api/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	networkingv1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/networking/v1alpha1"
)

// FakeServiceCIDRs implements ServiceCIDRInterface
type FakeServiceCIDRs struct {
	Fake *FakeNetworkingV1alpha1
}

var servicecidrsResource = v1alpha1.SchemeGroupVersion.WithResource("servicecidrs")

var servicecidrsKind = v1alpha1.SchemeGroupVersion.WithKind("ServiceCIDR")

// Get takes name of the serviceCIDR, and returns the corresponding serviceCIDR object, and an error if there is any.
func (c *FakeServiceCIDRs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceCIDR, err error) {
	emptyResult := &v1alpha1.ServiceCIDR{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(servicecidrsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ServiceCIDR), err
}

// List takes label and field selectors, and returns the list of ServiceCIDRs that match those selectors.
func (c *FakeServiceCIDRs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceCIDRList, err error) {
	emptyResult := &v1alpha1.ServiceCIDRList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(servicecidrsResource, servicecidrsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceCIDRList{ListMeta: obj.(*v1alpha1.ServiceCIDRList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceCIDRList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCIDRs.
func (c *FakeServiceCIDRs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(servicecidrsResource, opts))
}

// Create takes the representation of a serviceCIDR and creates it.  Returns the server's representation of the serviceCIDR, and an error, if there is any.
func (c *FakeServiceCIDRs) Create(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.CreateOptions) (result *v1alpha1.ServiceCIDR, err error) {
	emptyResult := &v1alpha1.ServiceCIDR{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(servicecidrsResource, serviceCIDR, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ServiceCIDR), err
}

// Update takes the representation of a serviceCIDR and updates it. Returns the server's representation of the serviceCIDR, and an error, if there is any.
func (c *FakeServiceCIDRs) Update(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.UpdateOptions) (result *v1alpha1.ServiceCIDR, err error) {
	emptyResult := &v1alpha1.ServiceCIDR{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(servicecidrsResource, serviceCIDR, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ServiceCIDR), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCIDRs) UpdateStatus(ctx context.Context, serviceCIDR *v1alpha1.ServiceCIDR, opts v1.UpdateOptions) (result *v1alpha1.ServiceCIDR, err error) {
	emptyResult := &v1alpha1.ServiceCIDR{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(servicecidrsResource, "status", serviceCIDR, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ServiceCIDR), err
}

// Delete takes name of the serviceCIDR and deletes it. Returns an error if one occurs.
func (c *FakeServiceCIDRs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(servicecidrsResource, name, opts), &v1alpha1.ServiceCIDR{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCIDRs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(servicecidrsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceCIDRList{})
	return err
}

// Patch applies the patch and returns the patched serviceCIDR.
func (c *FakeServiceCIDRs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceCIDR, err error) {
	emptyResult := &v1alpha1.ServiceCIDR{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(servicecidrsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ServiceCIDR), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied serviceCIDR.
func (c *FakeServiceCIDRs) Apply(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDRApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ServiceCIDR, err error) {
	if serviceCIDR == nil {
		return nil, fmt.Errorf("serviceCIDR provided to Apply must not be nil")
	}
	data, err := json.Marshal(serviceCIDR)
	if err != nil {
		return nil, err
	}
	name := serviceCIDR.Name
	if name == nil {
		return nil, fmt.Errorf("serviceCIDR.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.ServiceCIDR{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(servicecidrsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ServiceCIDR), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeServiceCIDRs) ApplyStatus(ctx context.Context, serviceCIDR *networkingv1alpha1.ServiceCIDRApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ServiceCIDR, err error) {
	if serviceCIDR == nil {
		return nil, fmt.Errorf("serviceCIDR provided to Apply must not be nil")
	}
	data, err := json.Marshal(serviceCIDR)
	if err != nil {
		return nil, err
	}
	name := serviceCIDR.Name
	if name == nil {
		return nil, fmt.Errorf("serviceCIDR.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.ServiceCIDR{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(servicecidrsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ServiceCIDR), err
}
