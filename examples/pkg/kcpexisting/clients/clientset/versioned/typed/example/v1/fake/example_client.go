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

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	examplev1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1"
	kcpexamplev1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/example/v1"
)

var _ kcpexamplev1.ExampleV1ClusterInterface = (*ExampleV1ClusterClient)(nil)

type ExampleV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ExampleV1ClusterClient) Cluster(clusterPath logicalcluster.Path) examplev1.ExampleV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ExampleV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *ExampleV1ClusterClient) TestTypes() kcpexamplev1.TestTypeClusterInterface {
	return &testTypesClusterClient{Fake: c.Fake}
}

func (c *ExampleV1ClusterClient) ClusterTestTypes() kcpexamplev1.ClusterTestTypeClusterInterface {
	return &clusterTestTypesClusterClient{Fake: c.Fake}
}

func (c *ExampleV1ClusterClient) WithoutVerbTypes() kcpexamplev1.WithoutVerbTypeClusterInterface {
	return &withoutVerbTypesClusterClient{Fake: c.Fake}
}

var _ examplev1.ExampleV1Interface = (*ExampleV1Client)(nil)

type ExampleV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *ExampleV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ExampleV1Client) TestTypes(namespace string) examplev1.TestTypeInterface {
	return &testTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ExampleV1Client) ClusterTestTypes() examplev1.ClusterTestTypeInterface {
	return &clusterTestTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *ExampleV1Client) WithoutVerbTypes(namespace string) examplev1.WithoutVerbTypeInterface {
	return &withoutVerbTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
