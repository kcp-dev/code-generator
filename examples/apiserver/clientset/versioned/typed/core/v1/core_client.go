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
	v1 "k8s.io/code-generator/examples/apiserver/apis/core/v1"
	"k8s.io/code-generator/examples/apiserver/clientset/versioned/scheme"
)

type CoreV1ClusterInterface interface {
	CoreV1ClusterScoper
	TestTypesClusterGetter
}

type CoreV1ClusterScoper interface {
	Cluster(logicalcluster.Path) CoreV1Interface
}

// CoreV1Client is used to interact with features provided by the  group.
type CoreV1ClusterClient struct {
	clientCache kcpclient.Cache[*CoreV1Interface]
}

func (c *CoreV1ClusterClient) TestTypes() TestTypeClusterInterface {
	return &testTypesClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new CoreV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*CoreV1ClusterClient, error) {
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

// NewForConfigAndClient creates a new CoreV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*CoreV1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*upstreamcorev1client.CoreV1Client]{
		NewForConfigAndClient: upstreamcorev1client.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}

	return &CoreV1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new CoreV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CoreV1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/api"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}
