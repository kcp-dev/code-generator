package v1

import v1 "k8s.io/client-go/informers/apps/v1"

type Interface interface {
	ConfigMaps() clusterConfigMapInformer
}

type version struct {
	delegate v1.Interface
}

func New(delegate v1.Interface) Interface {
	return &version{delegate: delegate}
}

func (v *version) ConfigMaps() ConfigMapInformer {
	return &configMapInformer{
		delegate: r.delegate.ConfigMaps(),
	}
}
