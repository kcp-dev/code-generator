// This file contains things that we want to factor out into separate packages
package v1

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/informers/core"
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	ClusterIndexName             = "cluster"
	ClusterAndNamespaceIndexName = "cluster-and-namespace"
)

func ClusterIndexFunc(obj interface{}) ([]string, error) {
	meta, err := meta.Accessor(obj)
	if err != nil {
		return []string{""}, fmt.Errorf("object has no meta: %v", err)
	}
	// return []string{meta.GetZZZ_DeprecatedClusterName()}, nil
	index := []string{meta.GetZZZ_DeprecatedClusterName()}
	fmt.Printf("\t\tClusterIndexFunc -> %v\n", index)
	return index, nil
}

func ClusterAndNamespaceIndexFunc(obj interface{}) ([]string, error) {
	meta, err := meta.Accessor(obj)
	if err != nil {
		return []string{""}, fmt.Errorf("object has no meta: %v", err)
	}
	// return []string{meta.GetZZZ_DeprecatedClusterName() + "/" + meta.GetNamespace()}, nil
	index := []string{meta.GetZZZ_DeprecatedClusterName() + "/" + meta.GetNamespace()}
	fmt.Printf("\t\tClusterAndNamespaceIndexFunc -> %v\n", index)
	return index, nil

}

func ClusterAwareKeyFunc(obj interface{}) (string, error) {
	meta, err := meta.Accessor(obj)
	if err != nil {
		return "", fmt.Errorf("object has no meta: %v", err)
	}
	clusterName := meta.GetZZZ_DeprecatedClusterName()
	namespace := meta.GetNamespace()
	name := meta.GetName()

	return strings.Join([]string{clusterName, namespace, name}, "/"), nil
}

func NewSharedInformerFactory(client kubernetes.Interface) *realClusterSharedInformerFactory {
	delegate := informers.NewSharedInformerFactoryWithOptions(
		client,
		0, // defaultResync,
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

	return &realClusterSharedInformerFactory{
		delegate: delegate,
	}
}

type realClusterSharedInformerFactory struct {
	delegate informers.SharedInformerFactory
}

func (r realClusterSharedInformerFactory) Core() clustercoreInterface {
	return &realclustercoreInterface{
		delegate: r.delegate.Core(),
	}
}

type clustercoreInterface interface {
	V1() clustercorev1Interface
}

type realclustercoreInterface struct {
	delegate core.Interface
}

func (r realclustercoreInterface) V1() clustercorev1Interface {
	return &realclustercorev1Interface{
		delegate: r.delegate.V1(),
	}
}

type clustercorev1Interface interface {
	ConfigMaps() clusterConfigMapInformer
}

type realclustercorev1Interface struct {
	delegate corev1informers.Interface
}

func (r realclustercorev1Interface) ConfigMaps() clusterConfigMapInformer {
	return &realclusterConfigMapInformer{
		delegate: r.delegate.ConfigMaps(),
	}
}

type clusterConfigMapInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() ConfigMapLister
}

type realclusterConfigMapInformer struct {
	delegate corev1informers.ConfigMapInformer
}

func (r realclusterConfigMapInformer) Informer() cache.SharedIndexInformer {
	return r.delegate.Informer()
}

func (r realclusterConfigMapInformer) Lister() ConfigMapLister {
	return NewConfigMapLister(r.delegate.Informer().GetIndexer())
}
