//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	corev1 "acme.corp/pkg/apis/core/v1"
	applyconfigurationscorev1 "acme.corp/pkg/generated/applyconfigurations/core/v1"
	corev1client "acme.corp/pkg/generated/clientset/versioned/typed/core/v1"
	kcpcorev1 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/core/v1"
)

var podTemplatesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "podtemplates"}
var podTemplatesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "PodTemplate"}

type podTemplatesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *podTemplatesClusterClient) Cluster(clusterPath logicalcluster.Path) kcpcorev1.PodTemplatesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &podTemplatesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PodTemplates that match those selectors across all clusters.
func (c *podTemplatesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PodTemplateList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(podTemplatesResource, podTemplatesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &corev1.PodTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.PodTemplateList{ListMeta: obj.(*corev1.PodTemplateList).ListMeta}
	for _, item := range obj.(*corev1.PodTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested PodTemplates across all clusters.
func (c *podTemplatesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(podTemplatesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type podTemplatesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *podTemplatesNamespacer) Namespace(namespace string) corev1client.PodTemplateInterface {
	return &podTemplatesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type podTemplatesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *podTemplatesClient) Create(ctx context.Context, podTemplate *corev1.PodTemplate, opts metav1.CreateOptions) (*corev1.PodTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(podTemplatesResource, c.ClusterPath, c.Namespace, podTemplate), &corev1.PodTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

func (c *podTemplatesClient) Update(ctx context.Context, podTemplate *corev1.PodTemplate, opts metav1.UpdateOptions) (*corev1.PodTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(podTemplatesResource, c.ClusterPath, c.Namespace, podTemplate), &corev1.PodTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

func (c *podTemplatesClient) UpdateStatus(ctx context.Context, podTemplate *corev1.PodTemplate, opts metav1.UpdateOptions) (*corev1.PodTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(podTemplatesResource, c.ClusterPath, "status", c.Namespace, podTemplate), &corev1.PodTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

func (c *podTemplatesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(podTemplatesResource, c.ClusterPath, c.Namespace, name, opts), &corev1.PodTemplate{})
	return err
}

func (c *podTemplatesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(podTemplatesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.PodTemplateList{})
	return err
}

func (c *podTemplatesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1.PodTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(podTemplatesResource, c.ClusterPath, c.Namespace, name), &corev1.PodTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

// List takes label and field selectors, and returns the list of PodTemplates that match those selectors.
func (c *podTemplatesClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.PodTemplateList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(podTemplatesResource, podTemplatesKind, c.ClusterPath, c.Namespace, opts), &corev1.PodTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.PodTemplateList{ListMeta: obj.(*corev1.PodTemplateList).ListMeta}
	for _, item := range obj.(*corev1.PodTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *podTemplatesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(podTemplatesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *podTemplatesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.PodTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podTemplatesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &corev1.PodTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

func (c *podTemplatesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1.PodTemplateApplyConfiguration, opts metav1.ApplyOptions) (*corev1.PodTemplate, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podTemplatesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &corev1.PodTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}

func (c *podTemplatesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1.PodTemplateApplyConfiguration, opts metav1.ApplyOptions) (*corev1.PodTemplate, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podTemplatesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &corev1.PodTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PodTemplate), err
}
