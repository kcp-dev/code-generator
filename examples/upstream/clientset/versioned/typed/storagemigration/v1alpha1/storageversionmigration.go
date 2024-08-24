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

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v1alpha1 "k8s.io/api/storagemigration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamstoragemigrationv1alpha1client "k8s.io/client-go/kubernetes/typed/storagemigration/v1alpha1"
)

// StorageVersionMigrationsClusterGetter has a method to return a StorageVersionMigrationClusterInterface.
// A group's client should implement this interface.
type StorageVersionMigrationsClusterGetter interface {
	StorageVersionMigrations() StorageVersionMigrationClusterInterface
}

// StorageVersionMigrationClusterInterface has methods to work with StorageVersionMigration resources.
type StorageVersionMigrationClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.StorageVersionMigrationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) upstreamNodeMagic
	StorageVersionMigrationExpansion
}

type storageVersionMigrationsClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamstoragemigrationv1alpha1client.StoragemigrationV1alpha1Client]
}
