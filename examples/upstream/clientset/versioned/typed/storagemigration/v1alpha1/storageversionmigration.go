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

package v1alpha1

import (
	"context"

	v1alpha1 "k8s.io/api/storagemigration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	storagemigrationv1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/storagemigration/v1alpha1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// StorageVersionMigrationsGetter has a method to return a StorageVersionMigrationInterface.
// A group's client should implement this interface.
type StorageVersionMigrationsGetter interface {
	StorageVersionMigrations() StorageVersionMigrationInterface
}

// StorageVersionMigrationInterface has methods to work with StorageVersionMigration resources.
type StorageVersionMigrationInterface interface {
	Create(ctx context.Context, storageVersionMigration *v1alpha1.StorageVersionMigration, opts v1.CreateOptions) (*v1alpha1.StorageVersionMigration, error)
	Update(ctx context.Context, storageVersionMigration *v1alpha1.StorageVersionMigration, opts v1.UpdateOptions) (*v1alpha1.StorageVersionMigration, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, storageVersionMigration *v1alpha1.StorageVersionMigration, opts v1.UpdateOptions) (*v1alpha1.StorageVersionMigration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.StorageVersionMigration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.StorageVersionMigrationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageVersionMigration, err error)
	Apply(ctx context.Context, storageVersionMigration *storagemigrationv1alpha1.StorageVersionMigrationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.StorageVersionMigration, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, storageVersionMigration *storagemigrationv1alpha1.StorageVersionMigrationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.StorageVersionMigration, err error)
	StorageVersionMigrationExpansion
}

// storageVersionMigrations implements StorageVersionMigrationInterface
type storageVersionMigrations struct {
	*gentype.ClientWithListAndApply[*v1alpha1.StorageVersionMigration, *v1alpha1.StorageVersionMigrationList, *storagemigrationv1alpha1.StorageVersionMigrationApplyConfiguration]
}

// newStorageVersionMigrations returns a StorageVersionMigrations
func newStorageVersionMigrations(c *StoragemigrationV1alpha1Client) *storageVersionMigrations {
	return &storageVersionMigrations{
		gentype.NewClientWithListAndApply[*v1alpha1.StorageVersionMigration, *v1alpha1.StorageVersionMigrationList, *storagemigrationv1alpha1.StorageVersionMigrationApplyConfiguration](
			"storageversionmigrations",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.StorageVersionMigration { return &v1alpha1.StorageVersionMigration{} },
			func() *v1alpha1.StorageVersionMigrationList { return &v1alpha1.StorageVersionMigrationList{} }),
	}
}
