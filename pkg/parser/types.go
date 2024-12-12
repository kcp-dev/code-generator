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
	GoName               string `marker:",+groupGoName"`
	Version              Version
	PackageAlias         string
	LowerCaseGroupGoName string
}

func (g Group) GoPackageAlias() string {
	if g.PackageAlias == "" {
		panic("PackageAlias is empty. Programmer error.")
	}
	return strings.ReplaceAll(g.PackageAlias, "-", "")
}

func (g Group) GroupGoName() string {
	if g.GoName == "" {
		panic("GroupGoName is empty. Programmer error.")
	}
	return strings.ReplaceAll(g.GoName, "-", "") // Code base is litterate with either GoName or GroupGoName. We should unify this.
}

func (g Group) PackagePath() string {
	return strings.ToLower(g.Group.PackageName())
}

func (g Group) PackageName() string {
	return strings.ToLower(strings.ReplaceAll(g.Group.PackageName(), "-", ""))
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
	return strings.ToLower(v.NonEmpty())
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
