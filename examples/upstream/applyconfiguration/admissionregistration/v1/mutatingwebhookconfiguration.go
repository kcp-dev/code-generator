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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// MutatingWebhookConfigurationApplyConfiguration represents a declarative configuration of the MutatingWebhookConfiguration type for use
// with apply.
type MutatingWebhookConfigurationApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Webhooks                         []MutatingWebhookApplyConfiguration `json:"webhooks,omitempty"`
}

// MutatingWebhookConfiguration constructs a declarative configuration of the MutatingWebhookConfiguration type for use with
// apply.
func MutatingWebhookConfiguration(name string) *MutatingWebhookConfigurationApplyConfiguration {
	b := &MutatingWebhookConfigurationApplyConfiguration{}
	b.WithName(name)
	b.WithKind("MutatingWebhookConfiguration")
	b.WithAPIVersion("admissionregistration.k8s.io/v1")
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithKind(value string) *MutatingWebhookConfigurationApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithAPIVersion(value string) *MutatingWebhookConfigurationApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithName(value string) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithGenerateName(value string) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithNamespace(value string) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithUID(value types.UID) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithResourceVersion(value string) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithGeneration(value int64) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithCreationTimestamp(value metav1.Time) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithLabels(entries map[string]string) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithAnnotations(entries map[string]string) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithFinalizers(values ...string) *MutatingWebhookConfigurationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *MutatingWebhookConfigurationApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithWebhooks adds the given value to the Webhooks field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Webhooks field.
func (b *MutatingWebhookConfigurationApplyConfiguration) WithWebhooks(values ...*MutatingWebhookApplyConfiguration) *MutatingWebhookConfigurationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithWebhooks")
		}
		b.Webhooks = append(b.Webhooks, *values[i])
	}
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *MutatingWebhookConfigurationApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
