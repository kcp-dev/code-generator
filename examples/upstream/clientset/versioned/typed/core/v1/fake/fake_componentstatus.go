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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	corev1 "k8s.io/code-generator/examples/upstream/applyconfiguration/core/v1"
)

// componentStatusesClusterClient implements componentStatusInterface
type componentStatusesClusterClient struct {
	*kcptesting.Fake
}

var componentstatusesResource = v1.SchemeGroupVersion.WithResource("componentstatuses")

var componentstatusesKind = v1.SchemeGroupVersion.WithKind("ComponentStatus")

// Get takes name of the componentStatus, and returns the corresponding componentStatus object, and an error if there is any.
func (c *componentStatusesClusterClient) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ComponentStatus, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(componentstatusesResource, c.ClusterPath, c.Namespace, name), &v1.ComponentStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentStatus), err
}

// List takes label and field selectors, and returns the list of ComponentStatuses that match those selectors.
func (c *componentStatusesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ComponentStatusList, err error) {
	emptyResult := &v1.ComponentStatusList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(componentstatusesResource, componentstatusesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ComponentStatusList{ListMeta: obj.(*v1.ComponentStatusList).ListMeta}
	for _, item := range obj.(*v1.ComponentStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested componentStatuses.
func (c *componentStatusesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(componentstatusesResource, opts))
}

// Create takes the representation of a componentStatus and creates it.  Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *componentStatusesClusterClient) Create(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.CreateOptions) (result *v1.ComponentStatus, err error) {
	emptyResult := &v1.ComponentStatus{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(componentstatusesResource, componentStatus, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ComponentStatus), err
}

// Update takes the representation of a componentStatus and updates it. Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *componentStatusesClusterClient) Update(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.UpdateOptions) (result *v1.ComponentStatus, err error) {
	emptyResult := &v1.ComponentStatus{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(componentstatusesResource, componentStatus, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ComponentStatus), err
}

// Delete takes name of the componentStatus and deletes it. Returns an error if one occurs.
func (c *componentStatusesClusterClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(componentstatusesResource, name, opts), &v1.ComponentStatus{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *componentStatusesClusterClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(componentstatusesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ComponentStatusList{})
	return err
}

// Patch applies the patch and returns the patched componentStatus.
func (c *componentStatusesClusterClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ComponentStatus, err error) {
	emptyResult := &v1.ComponentStatus{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(componentstatusesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ComponentStatus), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied componentStatus.
func (c *componentStatusesClusterClient) Apply(ctx context.Context, componentStatus *corev1.ComponentStatusApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ComponentStatus, err error) {
	if componentStatus == nil {
		return nil, fmt.Errorf("componentStatus provided to Apply must not be nil")
	}
	data, err := json.Marshal(componentStatus)
	if err != nil {
		return nil, err
	}
	name := componentStatus.Name
	if name == nil {
		return nil, fmt.Errorf("componentStatus.Name must be provided to Apply")
	}
	emptyResult := &v1.ComponentStatus{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(componentstatusesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ComponentStatus), err
}
