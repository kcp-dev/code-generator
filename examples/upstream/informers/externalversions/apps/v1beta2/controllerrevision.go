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

package v1beta2

import (
	"context"
	time "time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamappsv1beta2informers "k8s.io/client-go/informers/v1beta2/apps"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1beta2 "k8s.io/code-generator/examples/upstream/listers/apps/v1beta2"
)

// ControllerRevisionClusterInformer provides access to a shared informer and lister for
// ControllerRevisions.
type ControllerRevisionClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1beta2.ControllerRevisionLister
	Cluster(logicalcluster.Name) upstreamappsv1beta2informers.ControllerRevisionInformer
}

type controllerRevisionClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewControllerRevisionClusterInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewControllerRevisionClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredControllerRevisionClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredControllerRevisionClusterInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredControllerRevisionClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1beta2().ControllerRevisions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1beta2().ControllerRevisions(namespace).Watch(context.TODO(), options)
			},
		},
		&appsv1beta2.ControllerRevision{},
		resyncPeriod,
		indexers,
	)
}

func (f *controllerRevisionClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredControllerRevisionClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *controllerRevisionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1beta2.ControllerRevision{}, f.defaultInformer)
}

func (f *controllerRevisionInformer) Lister() v1beta2.ControllerRevisionLister {
	return v1beta2.NewControllerRevisionLister(f.Informer().GetIndexer())
}
