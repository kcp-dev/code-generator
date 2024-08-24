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
	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	upstreamauthorizationv1beta1client "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
)

// LocalSubjectAccessReviewsClusterGetter has a method to return a LocalSubjectAccessReviewClusterInterface.
// A group's client should implement this interface.
type LocalSubjectAccessReviewsClusterGetter interface {
	LocalSubjectAccessReviews() LocalSubjectAccessReviewClusterInterface
}

// LocalSubjectAccessReviewClusterInterface has methods to work with LocalSubjectAccessReview resources.
type LocalSubjectAccessReviewClusterInterface interface {
	LocalSubjectAccessReviewExpansion
}

type localSubjectAccessReviewsClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamauthorizationv1beta1client.AuthorizationV1beta1Client]
}
