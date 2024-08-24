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
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	rest "k8s.io/client-go/rest"
	v1beta1 "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/node/v1beta1"
)

type cSIDriversClusterClient struct {
	*kcptesting.Fake
}

func (c *FakeNodeV1beta1) RuntimeClasses() v1beta1.RuntimeClassInterface {
	return &FakeRuntimeClasses{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNodeV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
