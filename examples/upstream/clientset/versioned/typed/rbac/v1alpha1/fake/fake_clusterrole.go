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
	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	rbacv1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/rbac/v1alpha1"
)

// clusterRolesClusterClient implements clusterRoleInterface
type clusterRolesClusterClient struct {
	*kcptesting.Fake
}

var clusterrolesResource = v1alpha1.SchemeGroupVersion.WithResource("clusterroles")

var clusterrolesKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterRole")

// Get takes name of the clusterRole, and returns the corresponding clusterRole object, and an error if there is any.
func (c *clusterRolesClusterClient) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterRole, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(clusterrolesResource, c.ClusterPath, c.Namespace, name), &v1alpha1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}

// List takes label and field selectors, and returns the list of ClusterRoles that match those selectors.
func (c *clusterRolesClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterRoleList, err error) {
	emptyResult := &v1alpha1.ClusterRoleList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(clusterrolesResource, clusterrolesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterRoleList{ListMeta: obj.(*v1alpha1.ClusterRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRoles.
func (c *clusterRolesClusterClient) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(clusterrolesResource, opts))
}

// Create takes the representation of a clusterRole and creates it.  Returns the server's representation of the clusterRole, and an error, if there is any.
func (c *clusterRolesClusterClient) Create(ctx context.Context, clusterRole *v1alpha1.ClusterRole, opts v1.CreateOptions) (result *v1alpha1.ClusterRole, err error) {
	emptyResult := &v1alpha1.ClusterRole{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(clusterrolesResource, clusterRole, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}

// Update takes the representation of a clusterRole and updates it. Returns the server's representation of the clusterRole, and an error, if there is any.
func (c *clusterRolesClusterClient) Update(ctx context.Context, clusterRole *v1alpha1.ClusterRole, opts v1.UpdateOptions) (result *v1alpha1.ClusterRole, err error) {
	emptyResult := &v1alpha1.ClusterRole{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(clusterrolesResource, clusterRole, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}

// Delete takes name of the clusterRole and deletes it. Returns an error if one occurs.
func (c *clusterRolesClusterClient) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterrolesResource, name, opts), &v1alpha1.ClusterRole{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *clusterRolesClusterClient) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(clusterrolesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterRoleList{})
	return err
}

// Patch applies the patch and returns the patched clusterRole.
func (c *clusterRolesClusterClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterRole, err error) {
	emptyResult := &v1alpha1.ClusterRole{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusterrolesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterRole.
func (c *clusterRolesClusterClient) Apply(ctx context.Context, clusterRole *rbacv1alpha1.ClusterRoleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterRole, err error) {
	if clusterRole == nil {
		return nil, fmt.Errorf("clusterRole provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterRole)
	if err != nil {
		return nil, err
	}
	name := clusterRole.Name
	if name == nil {
		return nil, fmt.Errorf("clusterRole.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.ClusterRole{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusterrolesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ClusterRole), err
}
