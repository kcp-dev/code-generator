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

package v1beta1

import (
	"acme.corp/pkg/kcp/clients/informers/externalversions/internalinterfaces"
)

type ClusterInterface interface {
	// ValidatingAdmissionPolicies returns a ValidatingAdmissionPolicyClusterInformer
	ValidatingAdmissionPolicies() ValidatingAdmissionPolicyClusterInformer
	// ValidatingAdmissionPolicyBindings returns a ValidatingAdmissionPolicyBindingClusterInformer
	ValidatingAdmissionPolicyBindings() ValidatingAdmissionPolicyBindingClusterInformer
	// ValidatingWebhookConfigurations returns a ValidatingWebhookConfigurationClusterInformer
	ValidatingWebhookConfigurations() ValidatingWebhookConfigurationClusterInformer
	// MutatingWebhookConfigurations returns a MutatingWebhookConfigurationClusterInformer
	MutatingWebhookConfigurations() MutatingWebhookConfigurationClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &version{factory: f, tweakListOptions: tweakListOptions}
}

// ValidatingAdmissionPolicies returns a ValidatingAdmissionPolicyClusterInformer
func (v *version) ValidatingAdmissionPolicies() ValidatingAdmissionPolicyClusterInformer {
	return &validatingAdmissionPolicyClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ValidatingAdmissionPolicyBindings returns a ValidatingAdmissionPolicyBindingClusterInformer
func (v *version) ValidatingAdmissionPolicyBindings() ValidatingAdmissionPolicyBindingClusterInformer {
	return &validatingAdmissionPolicyBindingClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ValidatingWebhookConfigurations returns a ValidatingWebhookConfigurationClusterInformer
func (v *version) ValidatingWebhookConfigurations() ValidatingWebhookConfigurationClusterInformer {
	return &validatingWebhookConfigurationClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MutatingWebhookConfigurations returns a MutatingWebhookConfigurationClusterInformer
func (v *version) MutatingWebhookConfigurations() MutatingWebhookConfigurationClusterInformer {
	return &mutatingWebhookConfigurationClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

type Interface interface {
	// ValidatingAdmissionPolicies returns a ValidatingAdmissionPolicyInformer
	ValidatingAdmissionPolicies() ValidatingAdmissionPolicyInformer
	// ValidatingAdmissionPolicyBindings returns a ValidatingAdmissionPolicyBindingInformer
	ValidatingAdmissionPolicyBindings() ValidatingAdmissionPolicyBindingInformer
	// ValidatingWebhookConfigurations returns a ValidatingWebhookConfigurationInformer
	ValidatingWebhookConfigurations() ValidatingWebhookConfigurationInformer
	// MutatingWebhookConfigurations returns a MutatingWebhookConfigurationInformer
	MutatingWebhookConfigurations() MutatingWebhookConfigurationInformer
}

type scopedVersion struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// New returns a new ClusterInterface.
func NewScoped(f internalinterfaces.SharedScopedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &scopedVersion{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ValidatingAdmissionPolicies returns a ValidatingAdmissionPolicyInformer
func (v *scopedVersion) ValidatingAdmissionPolicies() ValidatingAdmissionPolicyInformer {
	return &validatingAdmissionPolicyScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ValidatingAdmissionPolicyBindings returns a ValidatingAdmissionPolicyBindingInformer
func (v *scopedVersion) ValidatingAdmissionPolicyBindings() ValidatingAdmissionPolicyBindingInformer {
	return &validatingAdmissionPolicyBindingScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ValidatingWebhookConfigurations returns a ValidatingWebhookConfigurationInformer
func (v *scopedVersion) ValidatingWebhookConfigurations() ValidatingWebhookConfigurationInformer {
	return &validatingWebhookConfigurationScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MutatingWebhookConfigurations returns a MutatingWebhookConfigurationInformer
func (v *scopedVersion) MutatingWebhookConfigurations() MutatingWebhookConfigurationInformer {
	return &mutatingWebhookConfigurationScopedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
