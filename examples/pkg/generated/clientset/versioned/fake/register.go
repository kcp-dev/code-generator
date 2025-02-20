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

// Code generated by client-gen-v0.32. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	examplev1 "acme.corp/pkg/apis/example/v1"
	examplev1alpha1 "acme.corp/pkg/apis/example/v1alpha1"
	examplev1beta1 "acme.corp/pkg/apis/example/v1beta1"
	examplev2 "acme.corp/pkg/apis/example/v2"
	example3v1 "acme.corp/pkg/apis/example3/v1"
	exampledashedv1 "acme.corp/pkg/apis/exampledashed/v1"
	existinginterfacesv1 "acme.corp/pkg/apis/existinginterfaces/v1"
	secondexamplev1 "acme.corp/pkg/apis/secondexample/v1"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	examplev1.AddToScheme,
	examplev1alpha1.AddToScheme,
	examplev1beta1.AddToScheme,
	examplev2.AddToScheme,
	example3v1.AddToScheme,
	exampledashedv1.AddToScheme,
	existinginterfacesv1.AddToScheme,
	secondexamplev1.AddToScheme,
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
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
