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

package v2

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

	examplev2 "acme.corp/pkg/apis/example/v2"
	upstreamexamplev2informers "acme.corp/pkg/generated/informers/externalversions/example/v2"
	upstreamexamplev2listers "acme.corp/pkg/generated/listers/example/v2"
	clientset "acme.corp/pkg/kcpexisting/clients/clientset/versioned"
	"acme.corp/pkg/kcpexisting/clients/informers/externalversions/internalinterfaces"
	examplev2listers "acme.corp/pkg/kcpexisting/clients/listers/example/v2"
)

// ClusterTestTypeClusterInformer provides access to a shared informer and lister for
// ClusterTestTypes.
type ClusterTestTypeClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamexamplev2informers.ClusterTestTypeInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() examplev2listers.ClusterTestTypeClusterLister
}

type clusterTestTypeClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterTestTypeClusterInformer constructs a new informer for ClusterTestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterTestTypeClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredClusterTestTypeClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterTestTypeClusterInformer constructs a new informer for ClusterTestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterTestTypeClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV2().ClusterTestTypes().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV2().ClusterTestTypes().Watch(context.TODO(), options)
			},
		},
		&examplev2.ClusterTestType{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterTestTypeClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredClusterTestTypeClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *clusterTestTypeClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&examplev2.ClusterTestType{}, f.defaultInformer)
}

func (f *clusterTestTypeClusterInformer) Lister() examplev2listers.ClusterTestTypeClusterLister {
	return examplev2listers.NewClusterTestTypeClusterLister(f.Informer().GetIndexer())
}

func (f *clusterTestTypeClusterInformer) Cluster(clusterName logicalcluster.Name) upstreamexamplev2informers.ClusterTestTypeInformer {
	return &clusterTestTypeInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type clusterTestTypeInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamexamplev2listers.ClusterTestTypeLister
}

func (f *clusterTestTypeInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *clusterTestTypeInformer) Lister() upstreamexamplev2listers.ClusterTestTypeLister {
	return f.lister
}
