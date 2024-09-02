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

package v1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	corev1 "acme.corp/pkg/apis/core/v1"
	scopedclientset "acme.corp/pkg/generated/clientset/versioned"
	clientset "acme.corp/pkg/kcp/clients/clientset/versioned"
	"acme.corp/pkg/kcp/clients/informers/externalversions/internalinterfaces"
	corev1listers "acme.corp/pkg/kcp/clients/listers/core/v1"
)

// NamespaceClusterInformer provides access to a shared informer and lister for
// Namespaces.
type NamespaceClusterInformer interface {
	Cluster(logicalcluster.Name) NamespaceInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() corev1listers.NamespaceClusterLister
}

type namespaceClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNamespaceClusterInformer constructs a new informer for Namespace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNamespaceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredNamespaceClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNamespaceClusterInformer constructs a new informer for Namespace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNamespaceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Namespaces().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Namespaces().Watch(context.TODO(), options)
			},
		},
		&corev1.Namespace{},
		resyncPeriod,
		indexers,
	)
}

func (f *namespaceClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredNamespaceClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *namespaceClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&corev1.Namespace{}, f.defaultInformer)
}

func (f *namespaceClusterInformer) Lister() corev1listers.NamespaceClusterLister {
	return corev1listers.NewNamespaceClusterLister(f.Informer().GetIndexer())
}

// NamespaceInformer provides access to a shared informer and lister for
// Namespaces.
type NamespaceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1listers.NamespaceLister
}

func (f *namespaceClusterInformer) Cluster(clusterName logicalcluster.Name) NamespaceInformer {
	return &namespaceInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type namespaceInformer struct {
	informer cache.SharedIndexInformer
	lister   corev1listers.NamespaceLister
}

func (f *namespaceInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *namespaceInformer) Lister() corev1listers.NamespaceLister {
	return f.lister
}

type namespaceScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func (f *namespaceScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.Namespace{}, f.defaultInformer)
}

func (f *namespaceScopedInformer) Lister() corev1listers.NamespaceLister {
	return corev1listers.NewNamespaceLister(f.Informer().GetIndexer())
}

// NewNamespaceInformer constructs a new informer for Namespace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNamespaceInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNamespaceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNamespaceInformer constructs a new informer for Namespace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNamespaceInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Namespaces().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Namespaces().Watch(context.TODO(), options)
			},
		},
		&corev1.Namespace{},
		resyncPeriod,
		indexers,
	)
}

func (f *namespaceScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNamespaceInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}
