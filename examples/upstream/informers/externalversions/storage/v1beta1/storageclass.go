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

package v1beta1

import (
	cache "k8s.io/client-go/tools/cache"
	v1beta1 "k8s.io/code-generator/examples/upstream/listers/storage/v1beta1"
	upstreamstorage.k8s.iov1beta1informers "k8s.io/client-go/informers/v1beta1/storage.k8s.io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	time "time"
	"github.com/kcp-dev/logicalcluster/v3"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
)


// StorageClassClusterInformer provides access to a shared informer and lister for
// StorageClasses.
type StorageClassClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1beta1.StorageClassLister
	Cluster(logicalcluster.Name) upstreamstorage.k8s.iov1beta1informers.StorageClassInformer
}

type storageClassClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageClassClusterInformer constructs a new informer for StorageClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageClassClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageClassClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageClassClusterInformer constructs a new informer for StorageClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageClassClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1beta1().StorageClasses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1beta1().StorageClasses().Watch(context.TODO(), options)
			},
		},
		&storagev1beta1.StorageClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageClassClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageClassClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *storageClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1beta1.StorageClass{}, f.defaultInformer)
}

func (f *storageClassInformer) Lister() v1beta1.StorageClassLister {
	return v1beta1.NewStorageClassLister(f.Informer().GetIndexer())
}
