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

package v1

import (
	http "net/http"

	rest "k8s.io/client-go/rest"

	existinginterfacesv1 "acme.corp/pkg/apis/existinginterfaces/v1"
	scheme "acme.corp/pkg/generated/clientset/versioned/scheme"
)

type ExistinginterfacesV1Interface interface {
	RESTClient() rest.Interface
	ClusterTestTypesGetter
	TestTypesGetter
}

// ExistinginterfacesV1Client is used to interact with features provided by the existinginterfaces.acme.corp group.
type ExistinginterfacesV1Client struct {
	restClient rest.Interface
}

func (c *ExistinginterfacesV1Client) ClusterTestTypes() ClusterTestTypeInterface {
	return newClusterTestTypes(c)
}

func (c *ExistinginterfacesV1Client) TestTypes(namespace string) TestTypeInterface {
	return newTestTypes(c, namespace)
}

// NewForConfig creates a new ExistinginterfacesV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ExistinginterfacesV1Client, error) {
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

// NewForConfigAndClient creates a new ExistinginterfacesV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ExistinginterfacesV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ExistinginterfacesV1Client{client}, nil
}

// NewForConfigOrDie creates a new ExistinginterfacesV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ExistinginterfacesV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ExistinginterfacesV1Client for the given RESTClient.
func New(c rest.Interface) *ExistinginterfacesV1Client {
	return &ExistinginterfacesV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := existinginterfacesv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ExistinginterfacesV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
