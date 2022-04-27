package externalversions

import (
	"fmt"

	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	case v1.SchemeGroupVersion.WithResource("configmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().ConfigMaps().Informer()}, nil
	}
	return nil, fmt.Errorf("no informer found for %v", resource)
}

type GenericLister interface {
	// List will return all objects across clusters
	List(selector labels.Selector) (ret []runtime.Object, err error)
	// Get will attempt to retrieve assuming that name==key
	Get(name string) (runtime.Object, error)
	// ByCluster will give you a GenericClusterLister for one namespace
	ByCluster(cluster logicalcluster.LogicalCluster) cache.GenericLister
}

func NewGenericLister(indexer cache.Indexer, resource schema.GroupResource) GenericLister {
	return &genericLister{indexer: indexer, resource: resource}
}

type genericLister struct {
	indexer  Indexer
	resource schema.GroupResource
}

func (s *genericLister) List(selector labels.Selector) (ret []runtime.Object, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*corev1.ConfigMap))
	})
	return ret, err
}

func (s *genericLister) Get(name string) (runtime.Object, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(s.resource, name)
	}
	return obj.(runtime.Object), nil
}

func (s *genericLister) ByCluster(cluster logicalcluster.LogicalCluster) cache.GenericLister {
	return &genericClusterLister{indexer: s.indexer, resource: s.resource, cluster: cluster}
}

type genericClusterLister struct {
	indexer  Indexer
	cluster  logicalcluster.LogicalCluster
	resource schema.GroupResource
}

func (s *genericClusterLister) List(selector labels.Selector) (ret []runtime.Object, err error) {
	list, err := c.indexer.ByIndex(ClusterIndexName, s.cluster.String())
	if err != nil {
		return nil, err
	}

	if selector == nil {
		selector = labels.Everything()
	}
	for i := range list {
		item := list[i].(*runtime.Object)
		if selector.Matches(labels.Set(item.GetLabels())) {
			ret = append(ret, item)
		}
	}

	return ret, err
}

func (s *genericClusterLister) Get(name string) (runtime.Object, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(s.resource, name)
	}
	return obj.(runtime.Object), nil
}

func (s *genericClusterLister) ByNamespace(namespace string) cache.GenericLister {
	return &genericNamespaceLister{indexer: s.indexer, namespace: namespace, resource: s.resource, cluster: s.cluster}
}

type genericNamespaceLister struct {
	indexer   Indexer
	cluster   logicalcluster.LogicalCluster
	namespace string
	resource  schema.GroupResource
}

func (s *genericNamespaceLister) List(selector labels.Selector) (ret []runtime.Object, err error) {
	list, err := s.indexer.Index(ClusterAndNamespaceIndexName, &metav1.ObjectMeta{
		ZZZ_DeprecatedClusterName: s.cluster.String(),
		Namespace:                 s.namespace,
	})
	if err != nil {
		return nil, err
	}

	if selector == nil {
		selector = labels.Everything()
	}
	for i := range list {
		item := list[i].(*runtime.Object)
		if selector.Matches(labels.Set(item.GetLabels())) {
			ret = append(ret, item)
		}
	}
	return ret, err
}

func (s *genericNamespaceLister) Get(name string) (runtime.Object, error) {
	meta := &metav1.ObjectMeta{
		ZZZ_DeprecatedClusterName: s.cluster.String(),
		Namespace:                 s.namespace,
		Name:                      name,
	}
	obj, exists, err := s.indexer.Get(meta)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(s.resource, name)
	}
	return obj.(*runtime.Object), nil
}
