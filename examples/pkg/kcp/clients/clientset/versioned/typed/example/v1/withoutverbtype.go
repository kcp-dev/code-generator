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

package v1

import (
	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	examplev1client "acme.corp/pkg/generated/clientset/versioned/typed/example/v1"
)

// WithoutVerbTypesClusterGetter has a method to return a WithoutVerbTypeClusterInterface.
// A group's cluster client should implement this interface.
type WithoutVerbTypesClusterGetter interface {
	WithoutVerbTypes() WithoutVerbTypeClusterInterface
}

// WithoutVerbTypeClusterInterface can scope down to one cluster and return a WithoutVerbTypesNamespacer.
type WithoutVerbTypeClusterInterface interface {
	Cluster(logicalcluster.Path) WithoutVerbTypesNamespacer
}

type withoutVerbTypesClusterInterface struct {
	clientCache kcpclient.Cache[*examplev1client.ExampleV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *withoutVerbTypesClusterInterface) Cluster(clusterPath logicalcluster.Path) WithoutVerbTypesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &withoutVerbTypesNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// WithoutVerbTypesNamespacer can scope to objects within a namespace, returning a examplev1client.WithoutVerbTypeInterface.
type WithoutVerbTypesNamespacer interface {
	Namespace(string) examplev1client.WithoutVerbTypeInterface
}

type withoutVerbTypesNamespacer struct {
	clientCache kcpclient.Cache[*examplev1client.ExampleV1Client]
	clusterPath logicalcluster.Path
}

func (n *withoutVerbTypesNamespacer) Namespace(namespace string) examplev1client.WithoutVerbTypeInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).WithoutVerbTypes(namespace)
}
