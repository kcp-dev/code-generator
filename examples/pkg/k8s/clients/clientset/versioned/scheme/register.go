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

package scheme

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	admissionregistrationv1 "acme.corp/pkg/apis/admissionregistration/v1"
	admissionregistrationv1alpha1 "acme.corp/pkg/apis/admissionregistration/v1alpha1"
	admissionregistrationv1beta1 "acme.corp/pkg/apis/admissionregistration/v1beta1"
	internalv1alpha1 "acme.corp/pkg/apis/apiserverinternal/v1alpha1"
	appsv1 "acme.corp/pkg/apis/apps/v1"
	appsv1beta1 "acme.corp/pkg/apis/apps/v1beta1"
	appsv1beta2 "acme.corp/pkg/apis/apps/v1beta2"
	authenticationv1 "acme.corp/pkg/apis/authentication/v1"
	authenticationv1alpha1 "acme.corp/pkg/apis/authentication/v1alpha1"
	authenticationv1beta1 "acme.corp/pkg/apis/authentication/v1beta1"
	authorizationv1 "acme.corp/pkg/apis/authorization/v1"
	authorizationv1beta1 "acme.corp/pkg/apis/authorization/v1beta1"
	autoscalingv1 "acme.corp/pkg/apis/autoscaling/v1"
	autoscalingv2 "acme.corp/pkg/apis/autoscaling/v2"
	autoscalingv2beta1 "acme.corp/pkg/apis/autoscaling/v2beta1"
	autoscalingv2beta2 "acme.corp/pkg/apis/autoscaling/v2beta2"
	batchv1 "acme.corp/pkg/apis/batch/v1"
	batchv1beta1 "acme.corp/pkg/apis/batch/v1beta1"
	certificatesv1 "acme.corp/pkg/apis/certificates/v1"
	certificatesv1alpha1 "acme.corp/pkg/apis/certificates/v1alpha1"
	certificatesv1beta1 "acme.corp/pkg/apis/certificates/v1beta1"
	coordinationv1 "acme.corp/pkg/apis/coordination/v1"
	coordinationv1alpha1 "acme.corp/pkg/apis/coordination/v1alpha1"
	coordinationv1beta1 "acme.corp/pkg/apis/coordination/v1beta1"
	corev1 "acme.corp/pkg/apis/core/v1"
	discoveryv1 "acme.corp/pkg/apis/discovery/v1"
	discoveryv1beta1 "acme.corp/pkg/apis/discovery/v1beta1"
	eventsv1 "acme.corp/pkg/apis/events/v1"
	eventsv1beta1 "acme.corp/pkg/apis/events/v1beta1"
	extensionsv1beta1 "acme.corp/pkg/apis/extensions/v1beta1"
	flowcontrolv1 "acme.corp/pkg/apis/flowcontrol/v1"
	flowcontrolv1beta1 "acme.corp/pkg/apis/flowcontrol/v1beta1"
	flowcontrolv1beta2 "acme.corp/pkg/apis/flowcontrol/v1beta2"
	flowcontrolv1beta3 "acme.corp/pkg/apis/flowcontrol/v1beta3"
	networkingv1 "acme.corp/pkg/apis/networking/v1"
	networkingv1alpha1 "acme.corp/pkg/apis/networking/v1alpha1"
	networkingv1beta1 "acme.corp/pkg/apis/networking/v1beta1"
	nodev1 "acme.corp/pkg/apis/node/v1"
	nodev1alpha1 "acme.corp/pkg/apis/node/v1alpha1"
	nodev1beta1 "acme.corp/pkg/apis/node/v1beta1"
	policyv1 "acme.corp/pkg/apis/policy/v1"
	policyv1beta1 "acme.corp/pkg/apis/policy/v1beta1"
	rbacv1 "acme.corp/pkg/apis/rbac/v1"
	rbacv1alpha1 "acme.corp/pkg/apis/rbac/v1alpha1"
	rbacv1beta1 "acme.corp/pkg/apis/rbac/v1beta1"
	resourcev1alpha3 "acme.corp/pkg/apis/resource/v1alpha3"
	schedulingv1 "acme.corp/pkg/apis/scheduling/v1"
	schedulingv1alpha1 "acme.corp/pkg/apis/scheduling/v1alpha1"
	schedulingv1beta1 "acme.corp/pkg/apis/scheduling/v1beta1"
	storagev1 "acme.corp/pkg/apis/storage/v1"
	storagev1alpha1 "acme.corp/pkg/apis/storage/v1alpha1"
	storagev1beta1 "acme.corp/pkg/apis/storage/v1beta1"
	storagemigrationv1alpha1 "acme.corp/pkg/apis/storagemigration/v1alpha1"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	admissionregistrationv1.AddToScheme,
	admissionregistrationv1alpha1.AddToScheme,
	admissionregistrationv1beta1.AddToScheme,
	appsv1.AddToScheme,
	appsv1beta1.AddToScheme,
	appsv1beta2.AddToScheme,
	authenticationv1.AddToScheme,
	authenticationv1alpha1.AddToScheme,
	authenticationv1beta1.AddToScheme,
	authorizationv1.AddToScheme,
	authorizationv1beta1.AddToScheme,
	autoscalingv1.AddToScheme,
	autoscalingv2.AddToScheme,
	autoscalingv2beta1.AddToScheme,
	autoscalingv2beta2.AddToScheme,
	batchv1.AddToScheme,
	batchv1beta1.AddToScheme,
	certificatesv1.AddToScheme,
	certificatesv1alpha1.AddToScheme,
	certificatesv1beta1.AddToScheme,
	coordinationv1.AddToScheme,
	coordinationv1alpha1.AddToScheme,
	coordinationv1beta1.AddToScheme,
	corev1.AddToScheme,
	discoveryv1.AddToScheme,
	discoveryv1beta1.AddToScheme,
	eventsv1.AddToScheme,
	eventsv1beta1.AddToScheme,
	extensionsv1beta1.AddToScheme,
	flowcontrolv1.AddToScheme,
	flowcontrolv1beta1.AddToScheme,
	flowcontrolv1beta2.AddToScheme,
	flowcontrolv1beta3.AddToScheme,
	internalv1alpha1.AddToScheme,
	networkingv1.AddToScheme,
	networkingv1alpha1.AddToScheme,
	networkingv1beta1.AddToScheme,
	nodev1.AddToScheme,
	nodev1alpha1.AddToScheme,
	nodev1beta1.AddToScheme,
	policyv1.AddToScheme,
	policyv1beta1.AddToScheme,
	rbacv1.AddToScheme,
	rbacv1alpha1.AddToScheme,
	rbacv1beta1.AddToScheme,
	resourcev1alpha3.AddToScheme,
	schedulingv1.AddToScheme,
	schedulingv1alpha1.AddToScheme,
	schedulingv1beta1.AddToScheme,
	storagemigrationv1alpha1.AddToScheme,
	storagev1.AddToScheme,
	storagev1alpha1.AddToScheme,
	storagev1beta1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//	import (
//	  "k8s.io/client-go/kubernetes"
//	  clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//	  aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//	)
//
//	kclientset, _ := kubernetes.NewForConfig(c)
//	_ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
