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
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	time "time"
	"github.com/kcp-dev/logicalcluster/v3"
	upstreamscheduling.k8s.iov1beta1informers "k8s.io/client-go/informers/v1beta1/scheduling.k8s.io"
	watch "k8s.io/apimachinery/pkg/watch"
	schedulingv1beta1 "k8s.io/api/scheduling/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	v1beta1 "k8s.io/code-generator/examples/upstream/listers/scheduling/v1beta1"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
)


// PriorityClassClusterInformer provides access to a shared informer and lister for
// PriorityClasses.
type PriorityClassClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1beta1.PriorityClassLister
	Cluster(logicalcluster.Name) upstreamscheduling.k8s.iov1beta1informers.PriorityClassInformer
}

type priorityClassClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	
}

// NewPriorityClassClusterInformer constructs a new informer for PriorityClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPriorityClassClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPriorityClassClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPriorityClassClusterInformer constructs a new informer for PriorityClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPriorityClassClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1beta1().PriorityClasses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1beta1().PriorityClasses().Watch(context.TODO(), options)
			},
		},
		&schedulingv1beta1.PriorityClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *priorityClassClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPriorityClassClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *priorityClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&schedulingv1beta1.PriorityClass{}, f.defaultInformer)
}

func (f *priorityClassInformer) Lister() v1beta1.PriorityClassLister {
	return v1beta1.NewPriorityClassLister(f.Informer().GetIndexer())
}
