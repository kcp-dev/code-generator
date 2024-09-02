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

package resource

import (
	"acme.corp/pkg/kcp/clients/informers/externalversions/internalinterfaces"
	"acme.corp/pkg/kcp/clients/informers/externalversions/resource/v1alpha3"
)

type ClusterInterface interface {
	// V1alpha3 provides access to the shared informers in V1alpha3.
	V1alpha3() v1alpha3.ClusterInterface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &group{factory: f, tweakListOptions: tweakListOptions}
}

// V1alpha3 returns a new v1alpha3.ClusterInterface.
func (g *group) V1alpha3() v1alpha3.ClusterInterface {
	return v1alpha3.New(g.factory, g.tweakListOptions)
}

type Interface interface {
	// V1alpha3 provides access to the shared informers in V1alpha3.
	V1alpha3() v1alpha3.Interface
}

type scopedGroup struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// New returns a new Interface.
func NewScoped(f internalinterfaces.SharedScopedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &scopedGroup{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V1alpha3 returns a new v1alpha3.ClusterInterface.
func (g *scopedGroup) V1alpha3() v1alpha3.Interface {
	return v1alpha3.NewScoped(g.factory, g.namespace, g.tweakListOptions)
}
