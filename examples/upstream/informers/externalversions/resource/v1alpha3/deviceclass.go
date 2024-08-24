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

package v1alpha3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1alpha3 "k8s.io/code-generator/examples/upstream/listers/resource/v1alpha3"
	"github.com/kcp-dev/logicalcluster/v3"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	resourcev1alpha3 "k8s.io/api/resource/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	time "time"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	upstreamresource.k8s.iov1alpha3informers "k8s.io/client-go/informers/v1alpha3/resource.k8s.io"
)


// DeviceClassClusterInformer provides access to a shared informer and lister for
// DeviceClasses.
type DeviceClassClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1alpha3.DeviceClassLister
	Cluster(logicalcluster.Name) upstreamresource.k8s.iov1alpha3informers.DeviceClassInformer
}

type deviceClassClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDeviceClassClusterInformer constructs a new informer for DeviceClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeviceClassClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeviceClassClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDeviceClassClusterInformer constructs a new informer for DeviceClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeviceClassClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().DeviceClasses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().DeviceClasses().Watch(context.TODO(), options)
			},
		},
		&resourcev1alpha3.DeviceClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *deviceClassClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeviceClassClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *deviceClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha3.DeviceClass{}, f.defaultInformer)
}

func (f *deviceClassInformer) Lister() v1alpha3.DeviceClassLister {
	return v1alpha3.NewDeviceClassLister(f.Informer().GetIndexer())
}
