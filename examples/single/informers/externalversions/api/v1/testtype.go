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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	apiv1 "k8s.io/code-generator/examples/single/api/v1"
	versioned "k8s.io/code-generator/examples/single/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/single/informers/externalversions/internalinterfaces"
	v1 "k8s.io/code-generator/examples/single/listers/api/v1"
)

// TestTypeClusterInformer provides access to a shared informer and lister for
// TestTypes.
type TestTypeClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1.TestTypeLister
	Cluster(logicalcluster.Name) TestTypeInformer
}

type testTypeClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTestTypeClusterInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestTypeClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestTypeClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTestTypeClusterInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestTypeClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().TestTypes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().TestTypes(namespace).Watch(context.TODO(), options)
			},
		},
		&apiv1.TestType{},
		resyncPeriod,
		indexers,
	)
}

func (f *testTypeClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestTypeClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *testTypeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiv1.TestType{}, f.defaultInformer)
}

func (f *testTypeInformer) Lister() v1.TestTypeLister {
	return v1.NewTestTypeLister(f.Informer().GetIndexer())
}
