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
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	examplev1beta1 "acme.corp/pkg/apis/example/v1beta1"
	examplev1beta1client "acme.corp/pkg/generated/clientset/versioned/typed/example/v1beta1"
)

// TestTypesClusterGetter has a method to return a TestTypeClusterInterface.
// A group's cluster client should implement this interface.
type TestTypesClusterGetter interface {
	TestTypes() TestTypeClusterInterface
}

// TestTypeClusterInterface can operate on TestTypes across all clusters,
// or scope down to one cluster and return a TestTypesNamespacer.
type TestTypeClusterInterface interface {
	Cluster(logicalcluster.Path) TestTypesNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*examplev1beta1.TestTypeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type testTypesClusterInterface struct {
	clientCache kcpclient.Cache[*examplev1beta1client.ExampleV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *testTypesClusterInterface) Cluster(clusterPath logicalcluster.Path) TestTypesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &testTypesNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all TestTypes across all clusters.
func (c *testTypesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*examplev1beta1.TestTypeList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).TestTypes(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all TestTypes across all clusters.
func (c *testTypesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).TestTypes(metav1.NamespaceAll).Watch(ctx, opts)
}

// TestTypesNamespacer can scope to objects within a namespace, returning a examplev1beta1client.TestTypeInterface.
type TestTypesNamespacer interface {
	Namespace(string) examplev1beta1client.TestTypeInterface
}

type testTypesNamespacer struct {
	clientCache kcpclient.Cache[*examplev1beta1client.ExampleV1beta1Client]
	clusterPath logicalcluster.Path
}

func (n *testTypesNamespacer) Namespace(namespace string) examplev1beta1client.TestTypeInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).TestTypes(namespace)
}
