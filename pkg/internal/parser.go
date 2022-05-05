/*
Copyright The KCP Authors.

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

package internal

import (
	"fmt"
	"go/types"
	"io"
	"path/filepath"
	"strings"
	"text/template"

	gentype "k8s.io/code-generator/cmd/client-gen/types"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

// interfaceWrapper is used to wrap each of the
// interfaces which are mentioned in the clientset.
type interfaceWrapper struct {
	// name of the interface provided from the input flag.
	InterfaceName string
	// clientsetname is the name of the package where the clientsets
	// are to be generated.
	ClientsetName string
	// ClientsetAPI path refers to where apis are generated.
	ClientsetAPIPath string
	// Pkgpath refers to the path where the typedClients would be written.
	TypedPkgPath string
	// APIs wrap each of the type
	APIs []api
	// writer wherein outputs are written
	writer *io.Writer
}

// api contains info about each type
type api struct {
	Name         string
	Version      string
	PkgName      string
	writer       io.Writer
	IsNamespaced bool
	HasStatus    bool

	PkgNameUpperFirst string
	VersionUpperFirst string
	NameLowerFirst    string
}

// packages stores the info used to scaffold wrapped interfaces content
type packages struct {
	Name              string
	APIPath           string
	ClientPath        string
	NameUpperFirst    string
	VersionUpperFirst string
	Version           string
	writer            io.Writer
}

// NewInterfaceWrapper returns a interfaceWrapper which can fill the templates to wrtie clientset wrappers.
func NewInterfaceWrapper(clientSetAPIPath, clientsetName, pkgPath, outputDir string, gvs []gentype.GroupVersions, w io.Writer) (*interfaceWrapper, error) {
	apis := groupVersionsToApis(gvs)
	return &interfaceWrapper{
		InterfaceName:    filepath.Base(clientSetAPIPath),
		ClientsetName:    clientsetName,
		ClientsetAPIPath: clientSetAPIPath,
		TypedPkgPath:     filepath.Join(pkgPath, filepath.Clean(outputDir), clientsetName),
		APIs:             apis,
		writer:           &w,
	}, nil
}

// TODO: this could be converted to an interface, wherein each sub-generator has a writeContent method.
func (w *interfaceWrapper) WriteContent() error {
	templ, err := template.New("wrapper").Parse(wrappedInterfacesTempl)
	if err != nil {
		return err
	}
	return templ.Execute(*w.writer, w)
}

// groupVersionToApis converts a list of types.GroupVersions to api type which can then be used for
// templating.
// Note: `Versions` in type.GroupVersions is assumed to contain only one version for now.
func groupVersionsToApis(gvs []gentype.GroupVersions) []api {
	result := make([]api, 0)

	for _, gv := range gvs {
		// this shouldn't happen, we would error out in this condition while validating flags.
		if len(gv.Versions) <= 0 {
			continue
		}
		a := &api{
			Name:    gv.Group.String(),
			Version: string(gv.Versions[0].Version),
			PkgName: gv.PackageName,
		}
		a.setCased()
		result = append(result, *a)
	}
	return result
}

func (a *api) setCased() {
	a.PkgNameUpperFirst = upperFirst(a.PkgName)
	a.VersionUpperFirst = upperFirst(a.Version)
	a.NameLowerFirst = lowerFirst(a.Name)
}

func (p *packages) setCased() {
	p.NameUpperFirst = upperFirst(p.Name)
	p.VersionUpperFirst = upperFirst(p.Version)
}

// lowerFirst sets the first alphabet to lowerCase.
func lowerFirst(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

// upperFirst sets the first alphabet to upperCase/
func upperFirst(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

// NewPackages returns a new packages instance which is used to write wrapper content.
func NewPackages(root *loader.Package, apiPath, clientPath, version, group string, w io.Writer) *packages {
	p := &packages{
		Name:       sanitize(group),
		APIPath:    apiPath,
		Version:    version,
		ClientPath: clientPath,
		writer:     w,
	}
	p.setCased()
	return p
}

// group names can have separators in them, ex: "example.com". This is a fix to
// sanitize those.
// TODO: Dig into code-gen on how they handle these cases and use the same logic
// here.
func sanitize(groupName string) string {
	if groupName != "" {
		arr := strings.Split(groupName, ".")
		groupName = arr[0]
	}
	return groupName
}

func (p *packages) WriteContent() error {
	templ, err := template.New("client").Parse(commonTempl)
	if err != nil {
		return err
	}
	return templ.Execute(p.writer, p)
}

func NewAPI(root *loader.Package, info *markers.TypeInfo, version, group string, isNamespaced bool, hasStatus bool, w io.Writer) (*api, error) {
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
		HasStatus:    hasStatus,
	}

	api.setCased()
	return api, nil
}

func (a *api) WriteContent() error {
	templ, err := template.New("wrapper").Parse(wrapperMethodsTempl)
	if err != nil {
		return err
	}
	return templ.Execute(a.writer, a)
}
