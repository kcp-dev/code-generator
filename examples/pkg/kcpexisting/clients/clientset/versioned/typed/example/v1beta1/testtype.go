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
	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	examplev1beta1client "acme.corp/pkg/generated/clientset/versioned/typed/example/v1beta1"
)

// TestTypesClusterGetter has a method to return a TestTypeClusterInterface.
// A group's cluster client should implement this interface.
type TestTypesClusterGetter interface {
	TestTypes() TestTypeClusterInterface
}

// TestTypeClusterInterface can scope down to one cluster and return a TestTypesNamespacer.
type TestTypeClusterInterface interface {
	Cluster(logicalcluster.Name) TestTypesNamespacer
}

type testTypesClusterInterface struct {
	clientCache kcpclient.Cache[*examplev1beta1client.ExampleV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *testTypesClusterInterface) Cluster(name logicalcluster.Name) TestTypesNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &testTypesNamespacer{clientCache: c.clientCache, name: name}
}

// TestTypesNamespacer can scope to objects within a namespace, returning a examplev1beta1client.TestTypeInterface.
type TestTypesNamespacer interface {
	Namespace(string) examplev1beta1client.TestTypeInterface
}

type testTypesNamespacer struct {
	clientCache kcpclient.Cache[*examplev1beta1client.ExampleV1beta1Client]
	name        logicalcluster.Name
}

func (n *testTypesNamespacer) Namespace(namespace string) examplev1beta1client.TestTypeInterface {
	return n.clientCache.ClusterOrDie(n.name).TestTypes(namespace)
}