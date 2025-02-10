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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"

	clientset "acme.corp/pkg/generated/clientset/versioned"
	examplev1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1"
	fakeexamplev1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1/fake"
	examplev1alpha1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1alpha1"
	fakeexamplev1alpha1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1alpha1/fake"
	examplev1beta1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1beta1"
	fakeexamplev1beta1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1beta1/fake"
	examplev2 "acme.corp/pkg/generated/clientset/versioned/typed/example/v2"
	fakeexamplev2 "acme.corp/pkg/generated/clientset/versioned/typed/example/v2/fake"
	example3v1 "acme.corp/pkg/generated/clientset/versioned/typed/example3/v1"
	fakeexample3v1 "acme.corp/pkg/generated/clientset/versioned/typed/example3/v1/fake"
	exampledashedv1 "acme.corp/pkg/generated/clientset/versioned/typed/exampledashed/v1"
	fakeexampledashedv1 "acme.corp/pkg/generated/clientset/versioned/typed/exampledashed/v1/fake"
	existinginterfacesv1 "acme.corp/pkg/generated/clientset/versioned/typed/existinginterfaces/v1"
	fakeexistinginterfacesv1 "acme.corp/pkg/generated/clientset/versioned/typed/existinginterfaces/v1/fake"
	secondexamplev1 "acme.corp/pkg/generated/clientset/versioned/typed/secondexample/v1"
	fakesecondexamplev1 "acme.corp/pkg/generated/clientset/versioned/typed/secondexample/v1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any field management, validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
//
// DEPRECATED: NewClientset replaces this with support for field management, which significantly improves
// server side apply testing. NewClientset is only available when apply configurations are generated (e.g.
// via --with-applyconfig).
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ExampleV1 retrieves the ExampleV1Client
func (c *Clientset) ExampleV1() examplev1.ExampleV1Interface {
	return &fakeexamplev1.FakeExampleV1{Fake: &c.Fake}
}

// ExampleV1alpha1 retrieves the ExampleV1alpha1Client
func (c *Clientset) ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1Interface {
	return &fakeexamplev1alpha1.FakeExampleV1alpha1{Fake: &c.Fake}
}

// ExampleV1beta1 retrieves the ExampleV1beta1Client
func (c *Clientset) ExampleV1beta1() examplev1beta1.ExampleV1beta1Interface {
	return &fakeexamplev1beta1.FakeExampleV1beta1{Fake: &c.Fake}
}

// ExampleV2 retrieves the ExampleV2Client
func (c *Clientset) ExampleV2() examplev2.ExampleV2Interface {
	return &fakeexamplev2.FakeExampleV2{Fake: &c.Fake}
}

// Example3V1 retrieves the Example3V1Client
func (c *Clientset) Example3V1() example3v1.Example3V1Interface {
	return &fakeexample3v1.FakeExample3V1{Fake: &c.Fake}
}

// ExampleDashedV1 retrieves the ExampleDashedV1Client
func (c *Clientset) ExampleDashedV1() exampledashedv1.ExampleDashedV1Interface {
	return &fakeexampledashedv1.FakeExampleDashedV1{Fake: &c.Fake}
}

// ExistinginterfacesV1 retrieves the ExistinginterfacesV1Client
func (c *Clientset) ExistinginterfacesV1() existinginterfacesv1.ExistinginterfacesV1Interface {
	return &fakeexistinginterfacesv1.FakeExistinginterfacesV1{Fake: &c.Fake}
}

// SecondexampleV1 retrieves the SecondexampleV1Client
func (c *Clientset) SecondexampleV1() secondexamplev1.SecondexampleV1Interface {
	return &fakesecondexamplev1.FakeSecondexampleV1{Fake: &c.Fake}
}
