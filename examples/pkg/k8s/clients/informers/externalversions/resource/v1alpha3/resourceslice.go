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

package v1alpha3

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

	resourcev1alpha3 "acme.corp/pkg/apis/resource/v1alpha3"
	scopedclientset "acme.corp/pkg/generated/clientset/versioned"
	clientset "acme.corp/pkg/kcp/clients/clientset/versioned"
	"acme.corp/pkg/kcp/clients/informers/externalversions/internalinterfaces"
	resourcev1alpha3listers "acme.corp/pkg/kcp/clients/listers/resource/v1alpha3"
)

// ResourceSliceClusterInformer provides access to a shared informer and lister for
// ResourceSlices.
type ResourceSliceClusterInformer interface {
	Cluster(logicalcluster.Name) ResourceSliceInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() resourcev1alpha3listers.ResourceSliceClusterLister
}

type resourceSliceClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceSliceClusterInformer constructs a new informer for ResourceSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceSliceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredResourceSliceClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceSliceClusterInformer constructs a new informer for ResourceSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceSliceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceSlices().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceSlices().Watch(context.TODO(), options)
			},
		},
		&resourcev1alpha3.ResourceSlice{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceSliceClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredResourceSliceClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *resourceSliceClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha3.ResourceSlice{}, f.defaultInformer)
}

func (f *resourceSliceClusterInformer) Lister() resourcev1alpha3listers.ResourceSliceClusterLister {
	return resourcev1alpha3listers.NewResourceSliceClusterLister(f.Informer().GetIndexer())
}

// ResourceSliceInformer provides access to a shared informer and lister for
// ResourceSlices.
type ResourceSliceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() resourcev1alpha3listers.ResourceSliceLister
}

func (f *resourceSliceClusterInformer) Cluster(clusterName logicalcluster.Name) ResourceSliceInformer {
	return &resourceSliceInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type resourceSliceInformer struct {
	informer cache.SharedIndexInformer
	lister   resourcev1alpha3listers.ResourceSliceLister
}

func (f *resourceSliceInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *resourceSliceInformer) Lister() resourcev1alpha3listers.ResourceSliceLister {
	return f.lister
}

type resourceSliceScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func (f *resourceSliceScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha3.ResourceSlice{}, f.defaultInformer)
}

func (f *resourceSliceScopedInformer) Lister() resourcev1alpha3listers.ResourceSliceLister {
	return resourcev1alpha3listers.NewResourceSliceLister(f.Informer().GetIndexer())
}

// NewResourceSliceInformer constructs a new informer for ResourceSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceSliceInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceSliceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceSliceInformer constructs a new informer for ResourceSlice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceSliceInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceSlices().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha3().ResourceSlices().Watch(context.TODO(), options)
			},
		},
		&resourcev1alpha3.ResourceSlice{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceSliceScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceSliceInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}
