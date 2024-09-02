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

package v1

import (
	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	authorizationv1client "acme.corp/pkg/generated/clientset/versioned/typed/authorization/v1"
)

// SelfSubjectAccessReviewsClusterGetter has a method to return a SelfSubjectAccessReviewClusterInterface.
// A group's cluster client should implement this interface.
type SelfSubjectAccessReviewsClusterGetter interface {
	SelfSubjectAccessReviews() SelfSubjectAccessReviewClusterInterface
}

// SelfSubjectAccessReviewClusterInterface can scope down to one cluster and return a authorizationv1client.SelfSubjectAccessReviewInterface.
type SelfSubjectAccessReviewClusterInterface interface {
	Cluster(logicalcluster.Path) authorizationv1client.SelfSubjectAccessReviewInterface
}

type selfSubjectAccessReviewsClusterInterface struct {
	clientCache kcpclient.Cache[*authorizationv1client.AuthorizationV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *selfSubjectAccessReviewsClusterInterface) Cluster(clusterPath logicalcluster.Path) authorizationv1client.SelfSubjectAccessReviewInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).SelfSubjectAccessReviews()
}
