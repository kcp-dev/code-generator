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
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	v1beta1 "k8s.io/code-generator/examples/upstream/listers/discovery/v1beta1"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	upstreamdiscovery.k8s.iov1beta1informers "k8s.io/client-go/informers/v1beta1/discovery.k8s.io"
	discoveryv1beta1 "k8s.io/api/discovery/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cache "k8s.io/client-go/tools/cache"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	time "time"
	"github.com/kcp-dev/logicalcluster/v3"
)


// EndpointSliceClusterInformer provides access to a shared informer and lister for
// EndpointSlices.
type EndpointSliceClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1beta1.EndpointSliceLister
	Cluster(logicalcluster.Name) upstreamdiscovery.k8s.iov1beta1informers.EndpointSliceInformer
}

type endpointSliceClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewEndpointSliceClusterInformer constructs a new informer for EndpointSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEndpointSliceClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEndpointSliceClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredEndpointSliceClusterInformer constructs a new informer for EndpointSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEndpointSliceClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DiscoveryV1beta1().EndpointSlices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DiscoveryV1beta1().EndpointSlices(namespace).Watch(context.TODO(), options)
			},
		},
		&discoveryv1beta1.EndpointSlice{},
		resyncPeriod,
		indexers,
	)
}

func (f *endpointSliceClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEndpointSliceClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *endpointSliceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&discoveryv1beta1.EndpointSlice{}, f.defaultInformer)
}

func (f *endpointSliceInformer) Lister() v1beta1.EndpointSliceLister {
	return v1beta1.NewEndpointSliceLister(f.Informer().GetIndexer())
}
