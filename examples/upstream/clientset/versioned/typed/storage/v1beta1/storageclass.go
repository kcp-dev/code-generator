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
	"context"

	v1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	storagev1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/storage/v1beta1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// StorageClassesGetter has a method to return a StorageClassInterface.
// A group's client should implement this interface.
type StorageClassesGetter interface {
	StorageClasses() StorageClassInterface
}

// StorageClassInterface has methods to work with StorageClass resources.
type StorageClassInterface interface {
	Create(ctx context.Context, storageClass *v1beta1.StorageClass, opts v1.CreateOptions) (*v1beta1.StorageClass, error)
	Update(ctx context.Context, storageClass *v1beta1.StorageClass, opts v1.UpdateOptions) (*v1beta1.StorageClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.StorageClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.StorageClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StorageClass, err error)
	Apply(ctx context.Context, storageClass *storagev1beta1.StorageClassApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.StorageClass, err error)
	StorageClassExpansion
}

// storageClasses implements StorageClassInterface
type storageClasses struct {
	*gentype.ClientWithListAndApply[*v1beta1.StorageClass, *v1beta1.StorageClassList, *storagev1beta1.StorageClassApplyConfiguration]
}

// newStorageClasses returns a StorageClasses
func newStorageClasses(c *StorageV1beta1Client) *storageClasses {
	return &storageClasses{
		gentype.NewClientWithListAndApply[*v1beta1.StorageClass, *v1beta1.StorageClassList, *storagev1beta1.StorageClassApplyConfiguration](
			"storageclasses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.StorageClass { return &v1beta1.StorageClass{} },
			func() *v1beta1.StorageClassList { return &v1beta1.StorageClassList{} }),
	}
}
