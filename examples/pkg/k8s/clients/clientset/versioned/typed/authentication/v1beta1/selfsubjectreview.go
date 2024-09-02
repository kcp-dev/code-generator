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
	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	authenticationv1beta1client "acme.corp/pkg/generated/clientset/versioned/typed/authentication/v1beta1"
)

// SelfSubjectReviewsClusterGetter has a method to return a SelfSubjectReviewClusterInterface.
// A group's cluster client should implement this interface.
type SelfSubjectReviewsClusterGetter interface {
	SelfSubjectReviews() SelfSubjectReviewClusterInterface
}

// SelfSubjectReviewClusterInterface can scope down to one cluster and return a authenticationv1beta1client.SelfSubjectReviewInterface.
type SelfSubjectReviewClusterInterface interface {
	Cluster(logicalcluster.Path) authenticationv1beta1client.SelfSubjectReviewInterface
}

type selfSubjectReviewsClusterInterface struct {
	clientCache kcpclient.Cache[*authenticationv1beta1client.AuthenticationV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *selfSubjectReviewsClusterInterface) Cluster(clusterPath logicalcluster.Path) authenticationv1beta1client.SelfSubjectReviewInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).SelfSubjectReviews()
}
