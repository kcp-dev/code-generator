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
	v1 "k8s.io/code-generator/examples/upstream/listers/networking/v1"
	time "time"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	upstreamnetworking.k8s.iov1informers "k8s.io/client-go/informers/v1/networking.k8s.io"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	networkingv1 "k8s.io/api/networking/v1"
	cache "k8s.io/client-go/tools/cache"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	"github.com/kcp-dev/logicalcluster/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
)


// IngressClusterInformer provides access to a shared informer and lister for
// Ingresses.
type IngressClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1.IngressLister
	Cluster(logicalcluster.Name) upstreamnetworking.k8s.iov1informers.IngressInformer
}

type ingressClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace string
}

// NewIngressClusterInformer constructs a new informer for Ingress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIngressClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIngressClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIngressClusterInformer constructs a new informer for Ingress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIngressClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().Ingresses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().Ingresses(namespace).Watch(context.TODO(), options)
			},
		},
		&networkingv1.Ingress{},
		resyncPeriod,
		indexers,
	)
}

func (f *ingressClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIngressClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *ingressInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1.Ingress{}, f.defaultInformer)
}

func (f *ingressInformer) Lister() v1.IngressLister {
	return v1.NewIngressLister(f.Informer().GetIndexer())
}
