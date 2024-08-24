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

package v2

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v2 "k8s.io/api/autoscaling/v2"
	upstreamautoscalingv2client "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	rest "k8s.io/client-go/rest"
	"k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

type AutoscalingV2ClusterInterface interface {
	AutoscalingV2ClusterScoper
	HorizontalPodAutoscalersClusterGetter
}

type AutoscalingV2ClusterScoper interface {
	Cluster(logicalcluster.Path) upstreamautoscalingv2client.AutoscalingV2Interface
}

// AutoscalingV2Client is used to interact with features provided by the autoscaling group.
type AutoscalingV2ClusterClient struct {
	clientCache kcpclient.Cache[*upstreamautoscalingv2client.AutoscalingV2Client]
}

func (c *AutoscalingV2ClusterClient) Cluster(clusterPath logicalcluster.Path) upstreamautoscalingv2client.AutoscalingV2Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

func (c *AutoscalingV2ClusterClient) HorizontalPodAutoscalers() HorizontalPodAutoscalerClusterInterface {
	return &horizontalPodAutoscalersClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new AutoscalingV2Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*AutoscalingV2ClusterClient, error) {
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

// NewForConfigAndClient creates a new AutoscalingV2Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*AutoscalingV2ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*upstreamautoscalingv2client.AutoscalingV2Client]{
		NewForConfigAndClient: upstreamautoscalingv2client.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}

	return &AutoscalingV2ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new AutoscalingV2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AutoscalingV2ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func setConfigDefaults(config *rest.Config) error {
	gv := v2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}
