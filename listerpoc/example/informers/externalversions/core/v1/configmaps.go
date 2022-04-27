package v1

import (
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kcp-dev/client-gen/listerpoc/example/listers/core/v1"
)

type ConfigMapInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ConfigMapLister
}

type configMapInformer struct {
	delegate corev1informers.ConfigMapInformer
}

func (r *configMapInformer) Informer() cache.SharedIndexInformer {
	return r.delegate.Informer()
}

func (r *configMapInformer) Lister() v1.ConfigMapLister {
	return v1.NewConfigMapLister(r.delegate.Informer().GetIndexer())
}
