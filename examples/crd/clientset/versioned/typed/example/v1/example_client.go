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

package v1

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	rest "k8s.io/client-go/rest"
	v1 "k8s.io/code-generator/examples/crd/apis/example/v1"
	"k8s.io/code-generator/examples/crd/clientset/versioned/scheme"
)

type ExampleV1ClusterInterface interface {
	ExampleV1ClusterScoper
	ClusterTestTypesClusterGetter
	TestTypesClusterGetter
}

type ExampleV1ClusterScoper interface {
	Cluster(logicalcluster.Path) ExampleV1Interface
}

// ExampleV1Client is used to interact with features provided by the example.crd.code-generator.k8s.io group.
type ExampleV1ClusterClient struct {
	clientCache kcpclient.Cache[*ExampleV1Interface]
}

func (c *ExampleV1ClusterClient) ClusterTestTypes() ClusterTestTypeClusterInterface {
	return &clusterTestTypesClusterInterface{clientCache: c.clientCache}
}

func (c *ExampleV1ClusterClient) TestTypes() TestTypeClusterInterface {
	return &testTypesClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new ExampleV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ExampleV1ClusterClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ExampleV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ExampleV1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*upstreamexamplev1client.ExampleV1Client]{
		NewForConfigAndClient: upstreamexamplev1client.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}

	return &ExampleV1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new ExampleV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ExampleV1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}
