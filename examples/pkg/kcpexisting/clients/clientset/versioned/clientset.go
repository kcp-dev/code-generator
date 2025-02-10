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

package clientset

import (
	"fmt"
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"

	client "acme.corp/pkg/generated/clientset/versioned"
	examplev1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/example/v1"
	examplev1alpha1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/example/v1alpha1"
	examplev1beta1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/example/v1beta1"
	examplev2 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/example/v2"
	example3v1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/example3/v1"
	exampledashedv1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/exampledashed/v1"
	existinginterfacesv1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/existinginterfaces/v1"
	secondexamplev1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/secondexample/v1"
)

type ClusterInterface interface {
	Cluster(logicalcluster.Path) client.Interface
	Discovery() discovery.DiscoveryInterface
	ExampleDashedV1() exampledashedv1.ExampleDashedV1ClusterInterface
	Example3V1() example3v1.Example3V1ClusterInterface
	ExampleV1() examplev1.ExampleV1ClusterInterface
	ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1ClusterInterface
	ExampleV1beta1() examplev1beta1.ExampleV1beta1ClusterInterface
	ExampleV2() examplev2.ExampleV2ClusterInterface
	ExistinginterfacesV1() existinginterfacesv1.ExistinginterfacesV1ClusterInterface
	SecondexampleV1() secondexamplev1.SecondexampleV1ClusterInterface
}

// ClusterClientset contains the clients for groups.
type ClusterClientset struct {
	*discovery.DiscoveryClient
	clientCache          kcpclient.Cache[*client.Clientset]
	exampledashedV1      *exampledashedv1.ExampleDashedV1ClusterClient
	example3V1           *example3v1.Example3V1ClusterClient
	exampleV1            *examplev1.ExampleV1ClusterClient
	exampleV1alpha1      *examplev1alpha1.ExampleV1alpha1ClusterClient
	exampleV1beta1       *examplev1beta1.ExampleV1beta1ClusterClient
	exampleV2            *examplev2.ExampleV2ClusterClient
	existinginterfacesV1 *existinginterfacesv1.ExistinginterfacesV1ClusterClient
	secondexampleV1      *secondexamplev1.SecondexampleV1ClusterClient
}

// Discovery retrieves the DiscoveryClient
func (c *ClusterClientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// ExampleDashedV1 retrieves the ExampleDashedV1ClusterClient.
func (c *ClusterClientset) ExampleDashedV1() exampledashedv1.ExampleDashedV1ClusterInterface {
	return c.exampledashedV1
}

// Example3V1 retrieves the Example3V1ClusterClient.
func (c *ClusterClientset) Example3V1() example3v1.Example3V1ClusterInterface {
	return c.example3V1
}

// ExampleV1 retrieves the ExampleV1ClusterClient.
func (c *ClusterClientset) ExampleV1() examplev1.ExampleV1ClusterInterface {
	return c.exampleV1
}

// ExampleV1alpha1 retrieves the ExampleV1alpha1ClusterClient.
func (c *ClusterClientset) ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1ClusterInterface {
	return c.exampleV1alpha1
}

// ExampleV1beta1 retrieves the ExampleV1beta1ClusterClient.
func (c *ClusterClientset) ExampleV1beta1() examplev1beta1.ExampleV1beta1ClusterInterface {
	return c.exampleV1beta1
}

// ExampleV2 retrieves the ExampleV2ClusterClient.
func (c *ClusterClientset) ExampleV2() examplev2.ExampleV2ClusterInterface {
	return c.exampleV2
}

// ExistinginterfacesV1 retrieves the ExistinginterfacesV1ClusterClient.
func (c *ClusterClientset) ExistinginterfacesV1() existinginterfacesv1.ExistinginterfacesV1ClusterInterface {
	return c.existinginterfacesV1
}

// SecondexampleV1 retrieves the SecondexampleV1ClusterClient.
func (c *ClusterClientset) SecondexampleV1() secondexamplev1.SecondexampleV1ClusterInterface {
	return c.secondexampleV1
}

// Cluster scopes this clientset to one cluster.
func (c *ClusterClientset) Cluster(clusterPath logicalcluster.Path) client.Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

// NewForConfig creates a new ClusterClientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ClusterClientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new ClusterClientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*ClusterClientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	cache := kcpclient.NewCache(c, httpClient, &kcpclient.Constructor[*client.Clientset]{
		NewForConfigAndClient: client.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}

	var cs ClusterClientset
	cs.clientCache = cache
	var err error
	cs.exampledashedV1, err = exampledashedv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.example3V1, err = example3v1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.exampleV1, err = examplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.exampleV1alpha1, err = examplev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.exampleV1beta1, err = examplev1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.exampleV2, err = examplev2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.existinginterfacesV1, err = existinginterfacesv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.secondexampleV1, err = secondexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new ClusterClientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterClientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}
