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
	v1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	appsv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/apps/v1"
	applyconfigurationautoscalingv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/autoscaling/v1"
)

// replicaSetsClusterClient implements replicaSetInterface
type replicaSetsClusterClient struct {
	*kcptesting.Fake
}

var replicasetsResource = v1.SchemeGroupVersion.WithResource("replicasets")

var replicasetsKind = v1.SchemeGroupVersion.WithKind("ReplicaSet")

// Get takes name of the replicaSet, and returns the corresponding replicaSet object, and an error if there is any.
func (c *replicaSetsClusterClient) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ReplicaSet, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(replicasetsResource, c.ClusterPath, c.Namespace, name), &v1.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ReplicaSet), err
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors.
func (c *replicaSetsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ReplicaSetList, err error) {
	emptyResult := &v1.ReplicaSetList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(replicasetsResource, replicasetsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ReplicaSetList{ListMeta: obj.(*v1.ReplicaSetList).ListMeta}
	for _, item := range obj.(*v1.ReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicaSets.
func (c *replicaSetsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(replicasetsResource, c.ns, opts))

}

// Create takes the representation of a replicaSet and creates it.  Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *replicaSetsClusterClient) Create(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.CreateOptions) (result *v1.ReplicaSet, err error) {
	emptyResult := &v1.ReplicaSet{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(replicasetsResource, c.ns, replicaSet, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicaSet), err
}

// Update takes the representation of a replicaSet and updates it. Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *replicaSetsClusterClient) Update(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.UpdateOptions) (result *v1.ReplicaSet, err error) {
	emptyResult := &v1.ReplicaSet{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(replicasetsResource, c.ns, replicaSet, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicaSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *replicaSetsClusterClient) UpdateStatus(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.UpdateOptions) (result *v1.ReplicaSet, err error) {
	emptyResult := &v1.ReplicaSet{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(replicasetsResource, "status", c.ns, replicaSet, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicaSet), err
}

// Delete takes name of the replicaSet and deletes it. Returns an error if one occurs.
func (c *replicaSetsClusterClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(replicasetsResource, c.ns, name, opts), &v1.ReplicaSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *replicaSetsClusterClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(replicasetsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ReplicaSetList{})
	return err
}

// Patch applies the patch and returns the patched replicaSet.
func (c *replicaSetsClusterClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ReplicaSet, err error) {
	emptyResult := &v1.ReplicaSet{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(replicasetsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicaSet), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied replicaSet.
func (c *replicaSetsClusterClient) Apply(ctx context.Context, replicaSet *appsv1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicaSet, err error) {
	if replicaSet == nil {
		return nil, fmt.Errorf("replicaSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(replicaSet)
	if err != nil {
		return nil, err
	}
	name := replicaSet.Name
	if name == nil {
		return nil, fmt.Errorf("replicaSet.Name must be provided to Apply")
	}
	emptyResult := &v1.ReplicaSet{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(replicasetsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicaSet), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *replicaSetsClusterClient) ApplyStatus(ctx context.Context, replicaSet *appsv1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicaSet, err error) {
	if replicaSet == nil {
		return nil, fmt.Errorf("replicaSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(replicaSet)
	if err != nil {
		return nil, err
	}
	name := replicaSet.Name
	if name == nil {
		return nil, fmt.Errorf("replicaSet.Name must be provided to Apply")
	}
	emptyResult := &v1.ReplicaSet{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(replicasetsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicaSet), err
}

// GetScale takes name of the replicaSet, and returns the corresponding scale object, and an error if there is any.
func (c *replicaSetsClusterClient) GetScale(ctx context.Context, replicaSetName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(configMapsResource, c.ClusterPath, c.Namespace, name), &corev1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err

	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(replicasetsResource, c.ns, "scale", replicaSetName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *replicaSetsClusterClient) UpdateScale(ctx context.Context, replicaSetName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(replicasetsResource, "scale", c.ns, scale, opts), &autoscalingv1.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *replicaSetsClusterClient) ApplyScale(ctx context.Context, replicaSetName string, scale *applyconfigurationautoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (result *autoscalingv1.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(replicasetsResource, c.ns, replicaSetName, types.ApplyPatchType, data, opts.ToPatchOptions(), "scale"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}
