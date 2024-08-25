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
	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	upstreamrbacv1alpha1client "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	"k8s.io/client-go/testing"
	rbacv1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/rbac/v1alpha1"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/rbac/v1alpha1"
)

var rolebindingsResource = v1alpha1.SchemeGroupVersion.WithResource("rolebindings")

var rolebindingsKind = v1alpha1.SchemeGroupVersion.WithKind("RoleBinding")

// roleBindingsClusterClient implements roleBindingInterface
type roleBindingsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *roleBindingsClusterClient) Cluster(clusterPath logicalcluster.Path) *kcp.RoleBindingNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &roleBindingsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of RoleBindings that match those selectors.
func (c *roleBindingsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RoleBindingList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(rolebindingsResource, rolebindingsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1alpha1.RoleBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RoleBindingList{ListMeta: obj.(*v1alpha1.RoleBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.RoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleBindings across all clusters.
func (c *roleBindingsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(rolebindingsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type roleBindingsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *roleBindingsNamespacer) Namespace(namespace string) upstreamrbacv1alpha1client.RoleBindingInterface {
	return &configMapsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type roleBindingsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *roleBindingsClient) Create(ctx context.Context, roleBinding *v1alpha1.RoleBinding, opts metav1.CreateOptions) (*v1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(rolebindingsResource, c.ClusterPath, c.Namespace, roleBinding), &v1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) Update(ctx context.Context, roleBinding *v1alpha1.RoleBinding, opts metav1.CreateOptions) (*v1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(rolebindingsResource, c.ClusterPath, c.Namespace, roleBinding), &v1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) UpdateStatus(ctx context.Context, roleBinding *v1alpha1.RoleBinding, opts metav1.CreateOptions) (*v1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(rolebindingsResource, c.ClusterPath, "status", c.Namespace, roleBinding), &v1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) Delete(ctx context.Context, name string, opts metav1.CreateOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(rolebindingsResource, c.ClusterPath, c.Namespace, name, opts), &v1alpha1.RoleBinding{})
	return err
}

func (c *roleBindingsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(rolebindingsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RoleBindingList{})
	return err
}

func (c *roleBindingsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(rolebindingsResource, c.ClusterPath, c.Namespace, name), &v1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleBinding), err
}

// List takes label and field selectors, and returns the list of v1alpha1.RoleBinding that match those selectors.
func (c *roleBindingsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.RoleBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(rolebindingsResource, rolebindingsKind, c.ClusterPath, c.Namespace, opts), &v1alpha1.RoleBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RoleBindingList{ListMeta: obj.(*v1alpha1.RoleBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.RoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *roleBindingsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(rolebindingsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *roleBindingsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1alpha1.RoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(rolebindingsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleBinding), err
}

func (c *roleBindingsClient) Apply(ctx context.Context, applyConfiguration *rbacv1alpha1.RoleBindingApplyConfiguration, opts metav1.ApplyOptions) (*v1alpha1.RoleBinding, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(rolebindingsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1alpha1.RoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleBinding), err
}
