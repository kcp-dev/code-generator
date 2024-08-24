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
	cache "k8s.io/client-go/tools/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	upstreaminternal.apiserver.k8s.iov1alpha1informers "k8s.io/client-go/informers/v1alpha1/internal.apiserver.k8s.io"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/code-generator/examples/upstream/listers/apiserverinternal/v1alpha1"
	time "time"
	apiserverinternalv1alpha1 "k8s.io/api/apiserverinternal/v1alpha1"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
)


// StorageVersionClusterInformer provides access to a shared informer and lister for
// StorageVersions.
type StorageVersionClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1alpha1.StorageVersionLister
	Cluster(logicalcluster.Name) upstreaminternal.apiserver.k8s.iov1alpha1informers.StorageVersionInformer
}

type storageVersionClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	
}

// NewStorageVersionClusterInformer constructs a new informer for StorageVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageVersionClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageVersionClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageVersionClusterInformer constructs a new informer for StorageVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageVersionClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InternalV1alpha1().StorageVersions().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InternalV1alpha1().StorageVersions().Watch(context.TODO(), options)
			},
		},
		&apiserverinternalv1alpha1.StorageVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageVersionClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageVersionClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *storageVersionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiserverinternalv1alpha1.StorageVersion{}, f.defaultInformer)
}

func (f *storageVersionInformer) Lister() v1alpha1.StorageVersionLister {
	return v1alpha1.NewStorageVersionLister(f.Informer().GetIndexer())
}
