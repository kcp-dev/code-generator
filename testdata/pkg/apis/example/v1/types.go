/*
Copyright 2022 The KCP Authors.

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

package v1

// +genclient
// TestType is a top-level type. A client is created for it.
type TestType struct {
	// +optional
	APIGroups []string `json:"apiGroups,omitempty" protobuf:"bytes,2,rep,name=apiGroups"`
}

// TestTypeList is a top-level list type. The client methods for lists are automatically created.
// You are not supposed to create a separated client for this one.
type TestTypeList struct {
	Items []TestType `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
type ClusterTestType struct {
	// Kind is the type of resource being referenced
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind"`
	// Name is the name of resource being referenced
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`
	// +optional
	Status ClusterTestTypeStatus `json:"status,omitempty"`
}

type ClusterTestTypeStatus struct {
	Blah string
}

type ClusterTestTypeList struct {
	Items []ClusterTestType `json:"items"`
}
