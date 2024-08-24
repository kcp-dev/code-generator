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

// nodesClusterClient implements nodeInterface
type nodesClusterClient struct {
	*kcptesting.Fake
}

var nodesResource = v1.SchemeGroupVersion.WithResource("nodes")

var nodesKind = v1.SchemeGroupVersion.WithKind("Node")

// Get takes name of the node, and returns the corresponding node object, and an error if there is any.
func (c *nodesClusterClient) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Node, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(nodesResource, c.ClusterPath, c.Namespace, name), &v1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Node), err
}

// List takes label and field selectors, and returns the list of Nodes that match those selectors.
func (c *nodesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NodeList, err error) {
	emptyResult := &v1.NodeList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(nodesResource, nodesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.NodeList{ListMeta: obj.(*v1.NodeList).ListMeta}
	for _, item := range obj.(*v1.NodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodes.
func (c *nodesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(nodesResource, opts))
}

// Create takes the representation of a node and creates it.  Returns the server's representation of the node, and an error, if there is any.
func (c *nodesClusterClient) Create(ctx context.Context, node *v1.Node, opts metav1.CreateOptions) (result *v1.Node, err error) {
	emptyResult := &v1.Node{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(nodesResource, node, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Node), err
}

// Update takes the representation of a node and updates it. Returns the server's representation of the node, and an error, if there is any.
func (c *nodesClusterClient) Update(ctx context.Context, node *v1.Node, opts metav1.UpdateOptions) (result *v1.Node, err error) {
	emptyResult := &v1.Node{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(nodesResource, node, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Node), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *nodesClusterClient) UpdateStatus(ctx context.Context, node *v1.Node, opts metav1.UpdateOptions) (result *v1.Node, err error) {
	emptyResult := &v1.Node{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(nodesResource, "status", node, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Node), err
}

// Delete takes name of the node and deletes it. Returns an error if one occurs.
func (c *nodesClusterClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(nodesResource, name, opts), &v1.Node{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *nodesClusterClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(nodesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.NodeList{})
	return err
}

// Patch applies the patch and returns the patched node.
func (c *nodesClusterClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Node, err error) {
	emptyResult := &v1.Node{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(nodesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Node), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied node.
func (c *nodesClusterClient) Apply(ctx context.Context, node *corev1.NodeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Node, err error) {
	if node == nil {
		return nil, fmt.Errorf("node provided to Apply must not be nil")
	}
	data, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	name := node.Name
	if name == nil {
		return nil, fmt.Errorf("node.Name must be provided to Apply")
	}
	emptyResult := &v1.Node{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(nodesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Node), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *nodesClusterClient) ApplyStatus(ctx context.Context, node *corev1.NodeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Node, err error) {
	if node == nil {
		return nil, fmt.Errorf("node provided to Apply must not be nil")
	}
	data, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	name := node.Name
	if name == nil {
		return nil, fmt.Errorf("node.Name must be provided to Apply")
	}
	emptyResult := &v1.Node{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(nodesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Node), err
}
