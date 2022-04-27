import (
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/kcp-dev/client-gen/listerpoc/listerpoc/example/listers/core/v1"
)

type ConfigMapInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ConfigMapLister
}

type configMapInformer struct {
	delegate corev1informers.ConfigMapInformer
}

func (r realclusterConfigMapInformer) Informer() cache.SharedIndexInformer {
	return r.delegate.Informer()
}

func (r realclusterConfigMapInformer) Lister() ConfigMapLister {
	return v1.NewConfigMapLister(r.delegate.Informer().GetIndexer())
}
