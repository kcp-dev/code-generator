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
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	"github.com/kcp-dev/logicalcluster/v3"
	upstreamscheduling.k8s.iov1informers "k8s.io/client-go/informers/v1/scheduling.k8s.io"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	watch "k8s.io/apimachinery/pkg/watch"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/code-generator/examples/upstream/listers/scheduling/v1"
	time "time"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	schedulingv1 "k8s.io/api/scheduling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// PriorityClassClusterInformer provides access to a shared informer and lister for
// PriorityClasses.
type PriorityClassClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1.PriorityClassLister
	Cluster(logicalcluster.Name) upstreamscheduling.k8s.iov1informers.PriorityClassInformer
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
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1().PriorityClasses().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1().PriorityClasses().Watch(context.TODO(), options)
			},
		},
		&schedulingv1.PriorityClass{},
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
	return f.factory.InformerFor(&schedulingv1.PriorityClass{}, f.defaultInformer)
}

func (f *priorityClassInformer) Lister() v1.PriorityClassLister {
	return v1.NewPriorityClassLister(f.Informer().GetIndexer())
}
