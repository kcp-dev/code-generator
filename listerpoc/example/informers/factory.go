
import (
	"time"

	"github.com/kcp-dev/client-gen/listerpoc/listerpoc/example/informers/core"

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

type sharedInformerFactory struct {
	delegate informers.SharedInformerFactory
}

func (r sharedInformerFactory) Core() core.Interface {
	return core.New(r.delegate.Core())
}

