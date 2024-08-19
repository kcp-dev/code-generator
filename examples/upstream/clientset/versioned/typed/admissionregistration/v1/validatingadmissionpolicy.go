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

	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	admissionregistrationv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/admissionregistration/v1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// ValidatingAdmissionPoliciesGetter has a method to return a ValidatingAdmissionPolicyInterface.
// A group's client should implement this interface.
type ValidatingAdmissionPoliciesGetter interface {
	ValidatingAdmissionPolicies() ValidatingAdmissionPolicyInterface
}

// ValidatingAdmissionPolicyInterface has methods to work with ValidatingAdmissionPolicy resources.
type ValidatingAdmissionPolicyInterface interface {
	Create(ctx context.Context, validatingAdmissionPolicy *v1.ValidatingAdmissionPolicy, opts metav1.CreateOptions) (*v1.ValidatingAdmissionPolicy, error)
	Update(ctx context.Context, validatingAdmissionPolicy *v1.ValidatingAdmissionPolicy, opts metav1.UpdateOptions) (*v1.ValidatingAdmissionPolicy, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, validatingAdmissionPolicy *v1.ValidatingAdmissionPolicy, opts metav1.UpdateOptions) (*v1.ValidatingAdmissionPolicy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ValidatingAdmissionPolicy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ValidatingAdmissionPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ValidatingAdmissionPolicy, err error)
	Apply(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1.ValidatingAdmissionPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ValidatingAdmissionPolicy, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1.ValidatingAdmissionPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ValidatingAdmissionPolicy, err error)
	ValidatingAdmissionPolicyExpansion
}

// validatingAdmissionPolicies implements ValidatingAdmissionPolicyInterface
type validatingAdmissionPolicies struct {
	*gentype.ClientWithListAndApply[*v1.ValidatingAdmissionPolicy, *v1.ValidatingAdmissionPolicyList, *admissionregistrationv1.ValidatingAdmissionPolicyApplyConfiguration]
}

// newValidatingAdmissionPolicies returns a ValidatingAdmissionPolicies
func newValidatingAdmissionPolicies(c *AdmissionregistrationV1Client) *validatingAdmissionPolicies {
	return &validatingAdmissionPolicies{
		gentype.NewClientWithListAndApply[*v1.ValidatingAdmissionPolicy, *v1.ValidatingAdmissionPolicyList, *admissionregistrationv1.ValidatingAdmissionPolicyApplyConfiguration](
			"validatingadmissionpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.ValidatingAdmissionPolicy { return &v1.ValidatingAdmissionPolicy{} },
			func() *v1.ValidatingAdmissionPolicyList { return &v1.ValidatingAdmissionPolicyList{} }),
	}
}
