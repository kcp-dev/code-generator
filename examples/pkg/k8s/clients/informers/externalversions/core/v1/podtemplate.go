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

// PodTemplateClusterInformer provides access to a shared informer and lister for
// PodTemplates.
type PodTemplateClusterInformer interface {
	Cluster(logicalcluster.Name) PodTemplateInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() corev1listers.PodTemplateClusterLister
}

type podTemplateClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPodTemplateClusterInformer constructs a new informer for PodTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodTemplateClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPodTemplateClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPodTemplateClusterInformer constructs a new informer for PodTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodTemplateClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().PodTemplates().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().PodTemplates().Watch(context.TODO(), options)
			},
		},
		&corev1.PodTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *podTemplateClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPodTemplateClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *podTemplateClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&corev1.PodTemplate{}, f.defaultInformer)
}

func (f *podTemplateClusterInformer) Lister() corev1listers.PodTemplateClusterLister {
	return corev1listers.NewPodTemplateClusterLister(f.Informer().GetIndexer())
}

// PodTemplateInformer provides access to a shared informer and lister for
// PodTemplates.
type PodTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1listers.PodTemplateLister
}

func (f *podTemplateClusterInformer) Cluster(clusterName logicalcluster.Name) PodTemplateInformer {
	return &podTemplateInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type podTemplateInformer struct {
	informer cache.SharedIndexInformer
	lister   corev1listers.PodTemplateLister
}

func (f *podTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *podTemplateInformer) Lister() corev1listers.PodTemplateLister {
	return f.lister
}

type podTemplateScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func (f *podTemplateScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.PodTemplate{}, f.defaultInformer)
}

func (f *podTemplateScopedInformer) Lister() corev1listers.PodTemplateLister {
	return corev1listers.NewPodTemplateLister(f.Informer().GetIndexer())
}

// NewPodTemplateInformer constructs a new informer for PodTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodTemplateInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodTemplateInformer(client, resyncPeriod, namespace, indexers, nil)
}

// NewFilteredPodTemplateInformer constructs a new informer for PodTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodTemplateInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().PodTemplates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().PodTemplates(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1.PodTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *podTemplateScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodTemplateInformer(client, resyncPeriod, f.namespace, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
	}, f.tweakListOptions)
}
