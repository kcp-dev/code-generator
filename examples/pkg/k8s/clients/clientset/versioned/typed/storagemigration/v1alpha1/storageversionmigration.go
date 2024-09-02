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

package v1alpha1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	storagemigrationv1alpha1 "acme.corp/pkg/apis/storagemigration/v1alpha1"
	storagemigrationv1alpha1client "acme.corp/pkg/generated/clientset/versioned/typed/storagemigration/v1alpha1"
)

// StorageVersionMigrationsClusterGetter has a method to return a StorageVersionMigrationClusterInterface.
// A group's cluster client should implement this interface.
type StorageVersionMigrationsClusterGetter interface {
	StorageVersionMigrations() StorageVersionMigrationClusterInterface
}

// StorageVersionMigrationClusterInterface can operate on StorageVersionMigrations across all clusters,
// or scope down to one cluster and return a storagemigrationv1alpha1client.StorageVersionMigrationInterface.
type StorageVersionMigrationClusterInterface interface {
	Cluster(logicalcluster.Path) storagemigrationv1alpha1client.StorageVersionMigrationInterface
	List(ctx context.Context, opts metav1.ListOptions) (*storagemigrationv1alpha1.StorageVersionMigrationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type storageVersionMigrationsClusterInterface struct {
	clientCache kcpclient.Cache[*storagemigrationv1alpha1client.StoragemigrationV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *storageVersionMigrationsClusterInterface) Cluster(clusterPath logicalcluster.Path) storagemigrationv1alpha1client.StorageVersionMigrationInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).StorageVersionMigrations()
}

// List returns the entire collection of all StorageVersionMigrations across all clusters.
func (c *storageVersionMigrationsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*storagemigrationv1alpha1.StorageVersionMigrationList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).StorageVersionMigrations().List(ctx, opts)
}

// Watch begins to watch all StorageVersionMigrations across all clusters.
func (c *storageVersionMigrationsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).StorageVersionMigrations().Watch(ctx, opts)
}
