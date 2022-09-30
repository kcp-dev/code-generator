/*
Copyright 2022 The KCP Authors.

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

package parser

import (
	"github.com/kcp-dev/code-generator/pkg/util"
	"github.com/kcp-dev/code-generator/third_party/namer"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/code-generator/cmd/client-gen/types"
)

type Kind struct {
	kind           string
	namespaced     bool
	supportedVerbs sets.String
	namer          namer.Namer
}

type Group struct {
	types.Group
	GoName string
}

func (k *Kind) Plural() string {
	return k.namer.Name(k.kind)
}

func (k *Kind) String() string {
	return k.kind
}

func (k *Kind) IsNamespaced() bool {
	return k.namespaced
}

func (k *Kind) SupportsListWatch() bool {
	return k.supportedVerbs.HasAll("list", "watch")
}

func NewKind(kind string, namespaced bool, supportedVerbs sets.String) Kind {
	return Kind{
		kind:           kind,
		namespaced:     namespaced,
		supportedVerbs: supportedVerbs,
		namer: namer.Namer{
			Finalize: util.UpperFirst,
			Exceptions: map[string]string{
				"Endpoints": "Endpoints",
			},
		},
	}
}
