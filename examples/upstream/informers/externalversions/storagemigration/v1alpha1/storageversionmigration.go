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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"
	storagemigrationv1alpha1 "k8s.io/api/storagemigration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamstoragemigrationv1alpha1informers "k8s.io/client-go/informers/v1alpha1/storagemigration"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/code-generator/examples/upstream/listers/storagemigration/v1alpha1"
)

// StorageVersionMigrationClusterInformer provides access to a shared informer and lister for
// StorageVersionMigrations.
type StorageVersionMigrationClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1alpha1.StorageVersionMigrationLister
	Cluster(logicalcluster.Name) upstreamstoragemigrationv1alpha1informers.StorageVersionMigrationInformer
}

type storageVersionMigrationClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageVersionMigrationClusterInformer constructs a new informer for StorageVersionMigration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageVersionMigrationClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageVersionMigrationClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageVersionMigrationClusterInformer constructs a new informer for StorageVersionMigration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageVersionMigrationClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StoragemigrationV1alpha1().StorageVersionMigrations().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StoragemigrationV1alpha1().StorageVersionMigrations().Watch(context.TODO(), options)
			},
		},
		&storagemigrationv1alpha1.StorageVersionMigration{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageVersionMigrationClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageVersionMigrationClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *storageVersionMigrationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagemigrationv1alpha1.StorageVersionMigration{}, f.defaultInformer)
}

func (f *storageVersionMigrationInformer) Lister() v1alpha1.StorageVersionMigrationLister {
	return v1alpha1.NewStorageVersionMigrationLister(f.Informer().GetIndexer())
}
