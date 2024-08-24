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

package v1

import (
	"context"
	time "time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamcorev1informers "k8s.io/client-go/informers/v1/core"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1 "k8s.io/code-generator/examples/upstream/listers/core/v1"
)

// EndpointsClusterInformer provides access to a shared informer and lister for
// Endpoints.
type EndpointsClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1.EndpointsLister
	Cluster(logicalcluster.Name) upstreamcorev1informers.EndpointsInformer
}

type endpointsClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEndpointsClusterInformer constructs a new informer for Endpoints type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEndpointsClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEndpointsClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEndpointsClusterInformer constructs a new informer for Endpoints type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEndpointsClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Endpoints(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Endpoints(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1.Endpoints{},
		resyncPeriod,
		indexers,
	)
}

func (f *endpointsClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEndpointsClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *endpointsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.Endpoints{}, f.defaultInformer)
}

func (f *endpointsInformer) Lister() v1.EndpointsLister {
	return v1.NewEndpointsLister(f.Informer().GetIndexer())
}
