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

package v1

import (
	"context"

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	networkingv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/networking/v1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// NetworkPoliciesGetter has a method to return a NetworkPolicyInterface.
// A group's client should implement this interface.
type NetworkPoliciesGetter interface {
	NetworkPolicies(namespace string) NetworkPolicyInterface
}

// NetworkPolicyInterface has methods to work with NetworkPolicy resources.
type NetworkPolicyInterface interface {
	Create(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.CreateOptions) (*v1.NetworkPolicy, error)
	Update(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.UpdateOptions) (*v1.NetworkPolicy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NetworkPolicy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NetworkPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkPolicy, err error)
	Apply(ctx context.Context, networkPolicy *networkingv1.NetworkPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *v1.NetworkPolicy, err error)
	NetworkPolicyExpansion
}

// networkPolicies implements NetworkPolicyInterface
type networkPolicies struct {
	*gentype.ClientWithListAndApply[*v1.NetworkPolicy, *v1.NetworkPolicyList, *networkingv1.NetworkPolicyApplyConfiguration]
}

// newNetworkPolicies returns a NetworkPolicies
func newNetworkPolicies(c *NetworkingV1Client, namespace string) *networkPolicies {
	return &networkPolicies{
		gentype.NewClientWithListAndApply[*v1.NetworkPolicy, *v1.NetworkPolicyList, *networkingv1.NetworkPolicyApplyConfiguration](
			"networkpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.NetworkPolicy { return &v1.NetworkPolicy{} },
			func() *v1.NetworkPolicyList { return &v1.NetworkPolicyList{} }),
	}
}
