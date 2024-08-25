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
	"context"
	time "time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"
	eventsv1beta1 "k8s.io/api/events/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreameventsv1beta1informers "k8s.io/client-go/informers/v1beta1/events"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1beta1 "k8s.io/code-generator/examples/upstream/listers/events/v1beta1"
)

// EventClusterInformer provides access to a shared informer and lister for
// Events.
type EventClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1beta1.EventLister
	Cluster(logicalcluster.Name) upstreameventsv1beta1informers.EventInformer
}

type eventClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewEventClusterInformer constructs a new informer for Event type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEventClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEventClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredEventClusterInformer constructs a new informer for Event type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEventClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventsV1beta1().Events(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventsV1beta1().Events(namespace).Watch(context.TODO(), options)
			},
		},
		&eventsv1beta1.Event{},
		resyncPeriod,
		indexers,
	)
}

func (f *eventClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEventClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *eventInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventsv1beta1.Event{}, f.defaultInformer)
}

func (f *eventInformer) Lister() v1beta1.EventLister {
	return v1beta1.NewEventLister(f.Informer().GetIndexer())
}
