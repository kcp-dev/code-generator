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

	resourcev1alpha3 "acme.corp/pkg/apis/resource/v1alpha3"
	applyconfigurationsresourcev1alpha3 "acme.corp/pkg/generated/applyconfigurations/resource/v1alpha3"
	resourcev1alpha3client "acme.corp/pkg/generated/clientset/versioned/typed/resource/v1alpha3"
	kcpresourcev1alpha3 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/resource/v1alpha3"
)

var resourceClaimTemplatesResource = schema.GroupVersionResource{Group: "resource.k8s.io", Version: "v1alpha3", Resource: "resourceclaimtemplates"}
var resourceClaimTemplatesKind = schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1alpha3", Kind: "ResourceClaimTemplate"}

type resourceClaimTemplatesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *resourceClaimTemplatesClusterClient) Cluster(clusterPath logicalcluster.Path) kcpresourcev1alpha3.ResourceClaimTemplatesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &resourceClaimTemplatesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ResourceClaimTemplates that match those selectors across all clusters.
func (c *resourceClaimTemplatesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha3.ResourceClaimTemplateList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(resourceClaimTemplatesResource, resourceClaimTemplatesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &resourcev1alpha3.ResourceClaimTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha3.ResourceClaimTemplateList{ListMeta: obj.(*resourcev1alpha3.ResourceClaimTemplateList).ListMeta}
	for _, item := range obj.(*resourcev1alpha3.ResourceClaimTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ResourceClaimTemplates across all clusters.
func (c *resourceClaimTemplatesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(resourceClaimTemplatesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type resourceClaimTemplatesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *resourceClaimTemplatesNamespacer) Namespace(namespace string) resourcev1alpha3client.ResourceClaimTemplateInterface {
	return &resourceClaimTemplatesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type resourceClaimTemplatesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *resourceClaimTemplatesClient) Create(ctx context.Context, resourceClaimTemplate *resourcev1alpha3.ResourceClaimTemplate, opts metav1.CreateOptions) (*resourcev1alpha3.ResourceClaimTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, resourceClaimTemplate), &resourcev1alpha3.ResourceClaimTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha3.ResourceClaimTemplate), err
}

func (c *resourceClaimTemplatesClient) Update(ctx context.Context, resourceClaimTemplate *resourcev1alpha3.ResourceClaimTemplate, opts metav1.UpdateOptions) (*resourcev1alpha3.ResourceClaimTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, resourceClaimTemplate), &resourcev1alpha3.ResourceClaimTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha3.ResourceClaimTemplate), err
}

func (c *resourceClaimTemplatesClient) UpdateStatus(ctx context.Context, resourceClaimTemplate *resourcev1alpha3.ResourceClaimTemplate, opts metav1.UpdateOptions) (*resourcev1alpha3.ResourceClaimTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(resourceClaimTemplatesResource, c.ClusterPath, "status", c.Namespace, resourceClaimTemplate), &resourcev1alpha3.ResourceClaimTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha3.ResourceClaimTemplate), err
}

func (c *resourceClaimTemplatesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, name, opts), &resourcev1alpha3.ResourceClaimTemplate{})
	return err
}

func (c *resourceClaimTemplatesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &resourcev1alpha3.ResourceClaimTemplateList{})
	return err
}

func (c *resourceClaimTemplatesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*resourcev1alpha3.ResourceClaimTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, name), &resourcev1alpha3.ResourceClaimTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha3.ResourceClaimTemplate), err
}

// List takes label and field selectors, and returns the list of ResourceClaimTemplates that match those selectors.
func (c *resourceClaimTemplatesClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha3.ResourceClaimTemplateList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(resourceClaimTemplatesResource, resourceClaimTemplatesKind, c.ClusterPath, c.Namespace, opts), &resourcev1alpha3.ResourceClaimTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha3.ResourceClaimTemplateList{ListMeta: obj.(*resourcev1alpha3.ResourceClaimTemplateList).ListMeta}
	for _, item := range obj.(*resourcev1alpha3.ResourceClaimTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *resourceClaimTemplatesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *resourceClaimTemplatesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*resourcev1alpha3.ResourceClaimTemplate, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &resourcev1alpha3.ResourceClaimTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha3.ResourceClaimTemplate), err
}

func (c *resourceClaimTemplatesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha3.ResourceClaimTemplateApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha3.ResourceClaimTemplate, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &resourcev1alpha3.ResourceClaimTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha3.ResourceClaimTemplate), err
}

func (c *resourceClaimTemplatesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha3.ResourceClaimTemplateApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha3.ResourceClaimTemplate, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimTemplatesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &resourcev1alpha3.ResourceClaimTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha3.ResourceClaimTemplate), err
}
