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
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	watch "k8s.io/apimachinery/pkg/watch"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	time "time"
	"github.com/kcp-dev/logicalcluster/v3"
	upstreamrbac.authorization.k8s.iov1informers "k8s.io/client-go/informers/v1/rbac.authorization.k8s.io"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "k8s.io/client-go/tools/cache"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1 "k8s.io/code-generator/examples/upstream/listers/rbac/v1"
)


// RoleBindingClusterInformer provides access to a shared informer and lister for
// RoleBindings.
type RoleBindingClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1.RoleBindingLister
	Cluster(logicalcluster.Name) upstreamrbac.authorization.k8s.iov1informers.RoleBindingInformer
}

type roleBindingClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace string
}

// NewRoleBindingClusterInformer constructs a new informer for RoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRoleBindingClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRoleBindingClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRoleBindingClusterInformer constructs a new informer for RoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRoleBindingClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1().RoleBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1().RoleBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&rbacv1.RoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *roleBindingClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRoleBindingClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *roleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbacv1.RoleBinding{}, f.defaultInformer)
}

func (f *roleBindingInformer) Lister() v1.RoleBindingLister {
	return v1.NewRoleBindingLister(f.Informer().GetIndexer())
}
