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
	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	upstreamauthenticationv1client "k8s.io/client-go/kubernetes/typed/authentication/v1"
)

// TokenReviewsClusterGetter has a method to return a TokenReviewClusterInterface.
// A group's client should implement this interface.
type TokenReviewsClusterGetter interface {
	TokenReviews() TokenReviewClusterInterface
}

// TokenReviewClusterInterface has methods to work with TokenReview resources.
type TokenReviewClusterInterface interface {
	TokenReviewExpansion
}

type tokenReviewsClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamauthenticationv1client.AuthenticationV1Client]
}
