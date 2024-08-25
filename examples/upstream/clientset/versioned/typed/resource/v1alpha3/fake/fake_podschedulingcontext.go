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
	v1alpha3 "k8s.io/api/resource/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	upstreamresourcev1alpha3client "k8s.io/client-go/kubernetes/typed/resource/v1alpha3"
	"k8s.io/client-go/testing"
	resourcev1alpha3 "k8s.io/code-generator/examples/upstream/applyconfiguration/resource/v1alpha3"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/resource/v1alpha3"
)

var podschedulingcontextsResource = v1alpha3.SchemeGroupVersion.WithResource("podschedulingcontexts")

var podschedulingcontextsKind = v1alpha3.SchemeGroupVersion.WithKind("PodSchedulingContext")

// podSchedulingContextsClusterClient implements podSchedulingContextInterface
type podSchedulingContextsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *podSchedulingContextsClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.PodSchedulingContextNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &podSchedulingContextsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PodSchedulingContexts that match those selectors.
func (c *podSchedulingContextsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.PodSchedulingContextList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(podschedulingcontextsResource, podschedulingcontextsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1alpha3.PodSchedulingContextList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.PodSchedulingContextList{ListMeta: obj.(*v1alpha3.PodSchedulingContextList).ListMeta}
	for _, item := range obj.(*v1alpha3.PodSchedulingContextList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podSchedulingContexts across all clusters.
func (c *podSchedulingContextsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(podschedulingcontextsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type podSchedulingContextsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *podSchedulingContextsNamespacer) Namespace(namespace string) upstreamresourcev1alpha3client.PodSchedulingContextInterface {
	return &configMapsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type podSchedulingContextsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *podSchedulingContextsClient) Create(ctx context.Context, podSchedulingContext *v1alpha3.PodSchedulingContext, opts metav1.CreateOptions) (*v1alpha3.PodSchedulingContext, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(podschedulingcontextsResource, c.ClusterPath, c.Namespace, podSchedulingContext), &v1alpha3.PodSchedulingContext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.PodSchedulingContext), err
}

func (c *podSchedulingContextsClient) Update(ctx context.Context, podSchedulingContext *v1alpha3.PodSchedulingContext, opts metav1.CreateOptions) (*v1alpha3.PodSchedulingContext, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(podschedulingcontextsResource, c.ClusterPath, c.Namespace, podSchedulingContext), &v1alpha3.PodSchedulingContext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.PodSchedulingContext), err
}

func (c *podSchedulingContextsClient) UpdateStatus(ctx context.Context, podSchedulingContext *v1alpha3.PodSchedulingContext, opts metav1.CreateOptions) (*v1alpha3.PodSchedulingContext, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(podschedulingcontextsResource, c.ClusterPath, "status", c.Namespace, podSchedulingContext), &v1alpha3.PodSchedulingContext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.PodSchedulingContext), err
}

func (c *podSchedulingContextsClient) Delete(ctx context.Context, name string, opts metav1.CreateOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(podschedulingcontextsResource, c.ClusterPath, c.Namespace, name, opts), &v1alpha3.PodSchedulingContext{})
	return err
}

func (c *podSchedulingContextsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(podschedulingcontextsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.PodSchedulingContextList{})
	return err
}

func (c *podSchedulingContextsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1alpha3.PodSchedulingContext, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(podschedulingcontextsResource, c.ClusterPath, c.Namespace, name), &v1alpha3.PodSchedulingContext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.PodSchedulingContext), err
}

// List takes label and field selectors, and returns the list of v1alpha3.PodSchedulingContext that match those selectors.
func (c *podSchedulingContextsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1alpha3.PodSchedulingContextList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(podschedulingcontextsResource, podschedulingcontextsKind, c.ClusterPath, c.Namespace, opts), &v1alpha3.PodSchedulingContextList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.PodSchedulingContextList{ListMeta: obj.(*v1alpha3.PodSchedulingContextList).ListMeta}
	for _, item := range obj.(*v1alpha3.PodSchedulingContextList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *podSchedulingContextsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(podschedulingcontextsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *podSchedulingContextsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1alpha3.PodSchedulingContext, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podschedulingcontextsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1alpha3.PodSchedulingContext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.PodSchedulingContext), err
}

func (c *podSchedulingContextsClient) Apply(ctx context.Context, applyConfiguration *resourcev1alpha3.PodSchedulingContextApplyConfiguration, opts metav1.ApplyOptions) (*v1alpha3.PodSchedulingContext, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podschedulingcontextsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1alpha3.PodSchedulingContext{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.PodSchedulingContext), err
}
