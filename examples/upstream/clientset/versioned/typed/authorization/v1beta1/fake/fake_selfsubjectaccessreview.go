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

package fake

import (
	"context"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	v1beta1 "k8s.io/api/authorization/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testing "k8s.io/client-go/testing"
)

// selfSubjectAccessReviewsClusterClient implements selfSubjectAccessReviewInterface
type selfSubjectAccessReviewsClusterClient struct {
	*kcptesting.Fake
}

var selfsubjectaccessreviewsResource = v1beta1.SchemeGroupVersion.WithResource("selfsubjectaccessreviews")

var selfsubjectaccessreviewsKind = v1beta1.SchemeGroupVersion.WithKind("SelfSubjectAccessReview")

// Create takes the representation of a selfSubjectAccessReview and creates it.  Returns the server's representation of the selfSubjectAccessReview, and an error, if there is any.
func (c *selfSubjectAccessReviewsClusterClient) Create(ctx context.Context, selfSubjectAccessReview *v1beta1.SelfSubjectAccessReview, opts v1.CreateOptions) (result *v1beta1.SelfSubjectAccessReview, err error) {
	emptyResult := &v1beta1.SelfSubjectAccessReview{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(selfsubjectaccessreviewsResource, selfSubjectAccessReview, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.SelfSubjectAccessReview), err
}
