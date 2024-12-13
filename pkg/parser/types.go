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
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/code-generator/cmd/client-gen/types"

	"github.com/kcp-dev/code-generator/v2/pkg/util"
	"github.com/kcp-dev/code-generator/v2/third_party/namer"
)

type Kind struct {
	kind           string
	namespaced     bool
	SupportedVerbs sets.Set[string]
	Extensions     []Extension
	namer          namer.Namer
}

type Group struct {
	types.Group
	Version              Version
	PackageAlias         string
	GoName               string
	LowerCaseGroupGoName string
}

func (g Group) PackageName() string {
	_g := strings.Split(g.NonEmpty(), ".")[0]
	_g = strings.ReplaceAll(_g, "-", "")
	return strings.ToLower(_g)
}

func (g Group) GroupGoName() string {
	return g.GoName
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
	return k.SupportedVerbs.HasAll("list", "watch")
}

type Version string

func (v Version) String() string {
	return string(v)
}

func (v Version) NonEmpty() string {
	if v == "" {
		return "internalVersion"
	}
	return v.String()
}

func (v Version) PackageName() string {
	_v := strings.ReplaceAll(v.NonEmpty(), "-", "")
	return strings.ToLower(_v)
}

type PackageVersion struct {
	Version
	// The fully qualified package, e.g. k8s.io/kubernetes/pkg/apis/apps, where the types.go is found.
	Package string
}

// TODO(skuznets):
// add an e2e for a kind that has no verbs, but uses an extension for something
// then ensure we add in fake_type.go entries for the extension
// changes we've already made should enable clients to exist for it

func NewKind(kind string, namespaced bool, supportedVerbs sets.Set[string], extensions []Extension) Kind {
	return Kind{
		kind:           kind,
		namespaced:     namespaced,
		SupportedVerbs: supportedVerbs,
		Extensions:     extensions,
		namer: namer.Namer{
			Finalize: util.UpperFirst,
			Exceptions: map[string]string{
				"Endpoints":               "Endpoints",
				"ResourceClaimParameters": "ResourceClaimParameters",
				"ResourceClassParameters": "ResourceClassParameters",
			},
		},
	}
}
