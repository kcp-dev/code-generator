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

package fake

import (
	gentype "k8s.io/client-go/gentype"

	v1beta1 "acme.corp/pkg/apis/example/v1beta1"
	examplev1beta1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1beta1"
)

// fakeTestTypes implements TestTypeInterface
type fakeTestTypes struct {
	*gentype.FakeClientWithList[*v1beta1.TestType, *v1beta1.TestTypeList]
	Fake *FakeExampleV1beta1
}

func newFakeTestTypes(fake *FakeExampleV1beta1, namespace string) examplev1beta1.TestTypeInterface {
	return &fakeTestTypes{
		gentype.NewFakeClientWithList[*v1beta1.TestType, *v1beta1.TestTypeList](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("testtypes"),
			v1beta1.SchemeGroupVersion.WithKind("TestType"),
			func() *v1beta1.TestType { return &v1beta1.TestType{} },
			func() *v1beta1.TestTypeList { return &v1beta1.TestTypeList{} },
			func(dst, src *v1beta1.TestTypeList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.TestTypeList) []*v1beta1.TestType { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.TestTypeList, items []*v1beta1.TestType) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
