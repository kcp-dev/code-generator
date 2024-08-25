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
	v1beta1 "k8s.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	upstreampolicyv1beta1client "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	"k8s.io/client-go/testing"
	policyv1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/policy/v1beta1"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/policy/v1beta1"
)

var evictionsResource = v1beta1.SchemeGroupVersion.WithResource("evictions")

var evictionsKind = v1beta1.SchemeGroupVersion.WithKind("Eviction")

// evictionsClusterClient implements evictionInterface
type evictionsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *evictionsClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.EvictionNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &evictionsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Evictions that match those selectors.
func (c *evictionsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.EvictionList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(evictionsResource, evictionsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta1.EvictionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EvictionList{ListMeta: obj.(*v1beta1.EvictionList).ListMeta}
	for _, item := range obj.(*v1beta1.EvictionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested evictions across all clusters.
func (c *evictionsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(evictionsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type evictionsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *evictionsNamespacer) Namespace(namespace string) upstreampolicyv1beta1client.EvictionInterface {
	return &configMapsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type evictionsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *evictionsClient) Create(ctx context.Context, eviction *v1beta1.Eviction, opts metav1.CreateOptions) (*v1beta1.Eviction, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(evictionsResource, c.ClusterPath, c.Namespace, eviction), &v1beta1.Eviction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Eviction), err
}

func (c *evictionsClient) Update(ctx context.Context, eviction *v1beta1.Eviction, opts metav1.CreateOptions) (*v1beta1.Eviction, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(evictionsResource, c.ClusterPath, c.Namespace, eviction), &v1beta1.Eviction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Eviction), err
}

func (c *evictionsClient) UpdateStatus(ctx context.Context, eviction *v1beta1.Eviction, opts metav1.CreateOptions) (*v1beta1.Eviction, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(evictionsResource, c.ClusterPath, "status", c.Namespace, eviction), &v1beta1.Eviction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Eviction), err
}

func (c *evictionsClient) Delete(ctx context.Context, name string, opts metav1.CreateOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(evictionsResource, c.ClusterPath, c.Namespace, name, opts), &v1beta1.Eviction{})
	return err
}

func (c *evictionsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(evictionsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.EvictionList{})
	return err
}

func (c *evictionsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.Eviction, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(evictionsResource, c.ClusterPath, c.Namespace, name), &v1beta1.Eviction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Eviction), err
}

// List takes label and field selectors, and returns the list of v1beta1.Eviction that match those selectors.
func (c *evictionsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.EvictionList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(evictionsResource, evictionsKind, c.ClusterPath, c.Namespace, opts), &v1beta1.EvictionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EvictionList{ListMeta: obj.(*v1beta1.EvictionList).ListMeta}
	for _, item := range obj.(*v1beta1.EvictionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *evictionsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(evictionsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *evictionsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta1.Eviction, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(evictionsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1beta1.Eviction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Eviction), err
}

func (c *evictionsClient) Apply(ctx context.Context, applyConfiguration *policyv1beta1.EvictionApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.Eviction, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(evictionsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1beta1.Eviction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Eviction), err
}
