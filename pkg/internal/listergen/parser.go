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

package listergen

import (
	"fmt"
	"go/types"
	"html/template"
	"io"

	"github.com/kcp-dev/code-generator/pkg/util"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

// api contains info about each type
// TODO: This would be modified as we add more markers to client-gen.
type api struct {
	Name         string
	Version      string
	PkgName      string
	writer       io.Writer
	IsNamespaced bool

	NameLowerFirst string
}

func NewAPI(root *loader.Package, info *markers.TypeInfo, version, group string, isNamespaced bool, w io.Writer) (*api, error) {
	typeInfo := root.TypesInfo.TypeOf(info.RawSpec.Name)
	if typeInfo == types.Typ[types.Invalid] {
		return nil, fmt.Errorf("unknown type: %s", info.Name)
	}

	api := &api{
		Name:         info.RawSpec.Name.Name,
		Version:      version,
		PkgName:      group,
		writer:       w,
		IsNamespaced: isNamespaced,
	}

	api.setCased()
	return api, nil
}

func (a *api) setCased() {
	a.NameLowerFirst = util.LowerFirst(a.Name)
}

func (a *api) WriteContent() error {
	templ, err := template.New("wrapper").Parse(apiWrapper)
	if err != nil {
		return err
	}
	return templ.Execute(a.writer, a)
}
