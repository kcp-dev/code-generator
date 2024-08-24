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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	v1alpha1 "k8s.io/code-generator/examples/upstream/listers/rbac/v1alpha1"
	time "time"
	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	upstreamrbac.authorization.k8s.iov1alpha1informers "k8s.io/client-go/informers/v1alpha1/rbac.authorization.k8s.io"
	"github.com/kcp-dev/logicalcluster/v3"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
)


// ClusterRoleBindingClusterInformer provides access to a shared informer and lister for
// ClusterRoleBindings.
type ClusterRoleBindingClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1alpha1.ClusterRoleBindingLister
	Cluster(logicalcluster.Name) upstreamrbac.authorization.k8s.iov1alpha1informers.ClusterRoleBindingInformer
}

type clusterRoleBindingClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	
}

// NewClusterRoleBindingClusterInformer constructs a new informer for ClusterRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterRoleBindingClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterRoleBindingClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterRoleBindingClusterInformer constructs a new informer for ClusterRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterRoleBindingClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().ClusterRoleBindings().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().ClusterRoleBindings().Watch(context.TODO(), options)
			},
		},
		&rbacv1alpha1.ClusterRoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterRoleBindingClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterRoleBindingClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *clusterRoleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbacv1alpha1.ClusterRoleBinding{}, f.defaultInformer)
}

func (f *clusterRoleBindingInformer) Lister() v1alpha1.ClusterRoleBindingLister {
	return v1alpha1.NewClusterRoleBindingLister(f.Informer().GetIndexer())
}
