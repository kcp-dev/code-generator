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
	"acme.corp/pkg/kcpexisting/clients/informers/externalversions/internalinterfaces"
)

type ClusterInterface interface {
	// TestTypes returns a TestTypeClusterInformer
	TestTypes() TestTypeClusterInformer
	// ClusterTestTypes returns a ClusterTestTypeClusterInformer
	ClusterTestTypes() ClusterTestTypeClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &version{factory: f, tweakListOptions: tweakListOptions}
}

// TestTypes returns a TestTypeClusterInformer
func (v *version) TestTypes() TestTypeClusterInformer {
	return &testTypeClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterTestTypes returns a ClusterTestTypeClusterInformer
func (v *version) ClusterTestTypes() ClusterTestTypeClusterInformer {
	return &clusterTestTypeClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
