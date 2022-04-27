package externalversions

import (
	"reflect"
	"time"

	"github.com/kcp-dev/client-gen/listerpoc/listerpoc/example/informers/core"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

func NewSharedInformerFactory(client kubernetes.Interface, defaultResync time.Duration) *sharedInformerFactory {
	delegate := informers.NewSharedInformerFactoryWithOptions(
		client,
		defaultResync,
		informers.WithExtraClusterScopedIndexers(
			cache.Indexers{
				ClusterIndexName: ClusterIndexFunc,
			},
		),
		informers.WithExtraNamespaceScopedIndexers(
			cache.Indexers{
				ClusterIndexName:             ClusterIndexFunc,
				ClusterAndNamespaceIndexName: ClusterAndNamespaceIndexFunc,
			},
		),
		informers.WithKeyFunction(ClusterAwareKeyFunc),
	)

	return &sharedInformerFactory{
		delegate: delegate,
	}
}

type SharedInformerFactory interface {
	Start(stopChn <-chan struct{})
	InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer

	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Core() core.Interface
}

type sharedInformerFactory struct {
	delegate informers.SharedInformerFactory
}

func (s *sharedInformerFactory) Start(stopChn <-chan struct{}) {
	s.delegate.Start(stopChn)
}

func (s *sharedInformerFactory) WaitForCacheSync(stopChn <-chan struct{}) map[reflect.Type]bool {
	return s.delegate.WaitForCacheSync(stopChn)
}

func (s *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer {
	return s.delegate.InformerFor(obj, newFunc)
}

func (s *sharedInformerFactory) ExtraClusterScopedIndexers() cache.Indexers {
	return s.delegate.ExtraClusterScopedIndexers()
}

func (s *sharedInformerFactory) ExtraNamespaceScopedIndexers() cache.Indexers {
	return s.delegate.ExtraClusterScopedIndexers()
}

func (s *sharedInformerFactory) KeyFunction() cache.KeyFunc {
	return s.delegate.KeyFunction()
}

func (f *sharedInformerFactory) Core() core.Interface {
	return core.New(f.delegate.Core())
}
