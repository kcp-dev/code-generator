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

package v1beta1

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/authorization/v1beta1"
	upstreamauthorizationv1beta1client "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	rest "k8s.io/client-go/rest"
	"k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

type AuthorizationV1beta1ClusterInterface interface {
	AuthorizationV1beta1ClusterScoper
	LocalSubjectAccessReviewsClusterGetter
	SelfSubjectAccessReviewsClusterGetter
	SelfSubjectRulesReviewsClusterGetter
	SubjectAccessReviewsClusterGetter
}

type AuthorizationV1beta1ClusterScoper interface {
	Cluster(logicalcluster.Path) upstreamauthorizationv1beta1client.AuthorizationV1beta1Interface
}

// AuthorizationV1beta1Client is used to interact with features provided by the authorization.k8s.io group.
type AuthorizationV1beta1ClusterClient struct {
	clientCache kcpclient.Cache[*upstreamauthorizationv1beta1client.AuthorizationV1beta1Client]
}

func (c *AuthorizationV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) upstreamauthorizationv1beta1client.AuthorizationV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

func (c *AuthorizationV1beta1ClusterClient) LocalSubjectAccessReviews() LocalSubjectAccessReviewClusterInterface {
	return &localSubjectAccessReviewsClusterInterface{clientCache: c.clientCache}
}

func (c *AuthorizationV1beta1ClusterClient) SelfSubjectAccessReviews() SelfSubjectAccessReviewClusterInterface {
	return &selfSubjectAccessReviewsClusterInterface{clientCache: c.clientCache}
}

func (c *AuthorizationV1beta1ClusterClient) SelfSubjectRulesReviews() SelfSubjectRulesReviewClusterInterface {
	return &selfSubjectRulesReviewsClusterInterface{clientCache: c.clientCache}
}

func (c *AuthorizationV1beta1ClusterClient) SubjectAccessReviews() SubjectAccessReviewClusterInterface {
	return &subjectAccessReviewsClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new AuthorizationV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*AuthorizationV1beta1ClusterClient, error) {
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

// NewForConfigAndClient creates a new AuthorizationV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*AuthorizationV1beta1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*upstreamauthorizationv1beta1client.AuthorizationV1beta1Client]{
		NewForConfigAndClient: upstreamauthorizationv1beta1client.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}

	return &AuthorizationV1beta1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new AuthorizationV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AuthorizationV1beta1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}
