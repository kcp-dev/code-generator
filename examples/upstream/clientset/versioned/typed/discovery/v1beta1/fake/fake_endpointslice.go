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
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/discovery/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	upstreamdiscoveryv1beta1client "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	"k8s.io/client-go/testing"
	discoveryv1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/discovery/v1beta1"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/discovery/v1beta1"
)

var endpointslicesResource = v1beta1.SchemeGroupVersion.WithResource("endpointslices")

var endpointslicesKind = v1beta1.SchemeGroupVersion.WithKind("EndpointSlice")

// endpointSlicesClusterClient implements endpointSliceInterface
type endpointSlicesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *endpointSlicesClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.EndpointSliceNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &endpointSlicesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of EndpointSlices that match those selectors.
func (c *endpointSlicesClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.EndpointSliceList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(endpointslicesResource, endpointslicesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta1.EndpointSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EndpointSliceList{ListMeta: obj.(*v1beta1.EndpointSliceList).ListMeta}
	for _, item := range obj.(*v1beta1.EndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpointSlices across all clusters.
func (c *endpointSlicesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(endpointslicesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type endpointSlicesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *endpointSlicesNamespacer) Namespace(namespace string) upstreamdiscoveryv1beta1client.EndpointSliceInterface {
	return &configMapsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type endpointSlicesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *endpointSlicesClient) Create(ctx context.Context, endpointSlice *v1beta1.EndpointSlice, opts metav1.CreateOptions) (*v1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(endpointslicesResource, c.ClusterPath, c.Namespace, endpointSlice), &v1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) Update(ctx context.Context, endpointSlice *v1beta1.EndpointSlice, opts metav1.CreateOptions) (*v1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(endpointslicesResource, c.ClusterPath, c.Namespace, endpointSlice), &v1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) UpdateStatus(ctx context.Context, endpointSlice *v1beta1.EndpointSlice, opts metav1.CreateOptions) (*v1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(endpointslicesResource, c.ClusterPath, "status", c.Namespace, endpointSlice), &v1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) Delete(ctx context.Context, name string, opts metav1.CreateOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(endpointslicesResource, c.ClusterPath, c.Namespace, name, opts), &v1beta1.EndpointSlice{})
	return err
}

func (c *endpointSlicesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(endpointslicesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.EndpointSliceList{})
	return err
}

func (c *endpointSlicesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(endpointslicesResource, c.ClusterPath, c.Namespace, name), &v1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

// List takes label and field selectors, and returns the list of v1beta1.EndpointSlice that match those selectors.
func (c *endpointSlicesClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.EndpointSliceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(endpointslicesResource, endpointslicesKind, c.ClusterPath, c.Namespace, opts), &v1beta1.EndpointSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EndpointSliceList{ListMeta: obj.(*v1beta1.EndpointSliceList).ListMeta}
	for _, item := range obj.(*v1beta1.EndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *endpointSlicesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(endpointslicesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *endpointSlicesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta1.EndpointSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(endpointslicesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}

func (c *endpointSlicesClient) Apply(ctx context.Context, applyConfiguration *discoveryv1beta1.EndpointSliceApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.EndpointSlice, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(endpointslicesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1beta1.EndpointSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EndpointSlice), err
}
