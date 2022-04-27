package core

import (
	"k8s.io/client-go/informers/core"

	v1 "github.com/kcp-dev/client-gen/listerpoc/listerpoc/example/informers/core/v1"
)

type Interface interface {
	V1() v1.Interface
}

type group struct {
	delegate core.Interface
}

func New(delegate core.Interface) Interface {
	return &group{delegate: delegate}
}

func (g *group) V1() v1.Interface {
	return v1.New(g.delegate.V1())
}
