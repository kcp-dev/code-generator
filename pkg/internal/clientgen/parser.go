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

package clientgen

import (
	"fmt"
	"go/types"
	"io"
	"path/filepath"
	"strings"
	"text/template"

	codegenutil "k8s.io/code-generator/cmd/client-gen/generators/util"

	"github.com/kcp-dev/code-generator/namer"
	"github.com/kcp-dev/code-generator/pkg/util"
	gentype "k8s.io/code-generator/cmd/client-gen/types"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

// funcMap contains the list of functions which are to be registered with
// the templates.
var funcMap = template.FuncMap{
	"upperFirst": func(s string) string {
		return strings.ToUpper(string(s[0])) + s[1:]
	},
	"lowerFirst": func(s string) string {
		return strings.ToLower(string(s[0])) + s[1:]
	},
	"default": util.DefaultValue,
	"typepkg": func(a api) string {
		return fmt.Sprintf("*%sapi%s.%s", a.PkgName, a.Version, a.Name)
	},
}

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

// packages stores the info used to scaffold wrapped interfaces content
type packages struct {
	Name       string
	APIPath    string
	ClientPath string
	Version    string
	Imports    []string
	writer     io.Writer
}

// NewInterfaceWrapper returns a interfaceWrapper which can fill the templates to wrtie clientset wrappers.
func NewInterfaceWrapper(clientSetAPIPath, clientsetName, pkgPath string, gvs []gentype.GroupVersions, w io.Writer) (*interfaceWrapper, error) {
	apis := groupVersionsToApis(gvs)
	return &interfaceWrapper{
		InterfaceName:    filepath.Base(clientSetAPIPath),
		ClientsetName:    clientsetName,
		ClientsetAPIPath: clientSetAPIPath,
		TypedPkgPath:     pkgPath,
		APIs:             apis,
		writer:           &w,
	}, nil
}

// TODO: this could be converted to an interface, wherein each sub-generator has a writeContent method.
func (w *interfaceWrapper) WriteContent() error {
	templ, err := template.New("wrapper").Funcs(funcMap).Parse(wrappedInterfacesTempl)
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
		result = append(result, *a)
	}
	return result
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

// TODO: Clean this up. Its not required to convert the input to a struct
// and then to a map before passing it to template. Using of map is better,
// as it allows to add more variables dynamically.
func (p *packages) WriteContent(importList *[]string) error {
	if importList != nil {
		p.Imports = append(p.Imports, p.appendGenericImports(*importList)...)
	}
	templ, err := template.New("client").Funcs(funcMap).Parse(commonTempl)
	if err != nil {
		return err
	}
	return templ.Execute(p.writer, p)
}

// api contains info about each type
type api struct {
	Name                      string
	PkgName                   string
	Version                   string
	IsNamespaced              bool
	AdditionalMethods         []AdditionalMethod
	SkipVerbs                 []string
	NoVerbs                   bool
	HasStatus                 bool
	ApplyConfigurationPackage string
	writer                    io.Writer
	importList                *[]string

	InputType  string
	InputName  string
	ResultType string
	Method     string
}

type AdditionalMethod struct {
	Method      *string
	Verb        *string
	Subresource *string
	Input       *string
	Result      *string
}

func NewAPI(
	root *loader.Package,
	info *markers.TypeInfo,
	group, version, applyconfigurationpkg string,
	namespaceScoped bool,
	additionalMethods []AdditionalMethod,
	importList *[]string,
	skipVerbs, onlyVerbs []string,
	noVerbs, hasStatus bool,
	w io.Writer,
) (*api, error) {
	typeInfo := root.TypesInfo.TypeOf(info.RawSpec.Name)
	if typeInfo == types.Typ[types.Invalid] {
		return nil, fmt.Errorf("unknown type: %s", info.Name)
	}

	api := &api{
		Name:                      info.RawSpec.Name.Name,
		PkgName:                   group,
		Version:                   version,
		ApplyConfigurationPackage: applyconfigurationpkg,
		IsNamespaced:              namespaceScoped,
		AdditionalMethods:         additionalMethods,
		importList:                importList,
		SkipVerbs:                 skipVerbs,
		NoVerbs:                   noVerbs,
		HasStatus:                 hasStatus,
		writer:                    w,
	}
	return api, nil
}

// TODO: add templating logic for additional methods
func (a *api) WriteContent() error {
	generateApply := len(a.ApplyConfigurationPackage) > 0
	// Add common template followed by specific methods which are required.
	if err := templateExecute(wrapperMethodsTempl, *a); err != nil {
		return err
	}

	if a.NoVerbs {
		return nil
	}

	if a.hasVerb("get") {
		if err := templateExecute(getTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("list") {
		if err := templateExecute(listTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("watch") {
		if err := templateExecute(watchTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("create") {
		if err := templateExecute(createTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("update") {
		if err := templateExecute(updateTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("updateStatus") && a.HasStatus {
		if err := templateExecute(updateStatusTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("delete") {
		if err := templateExecute(deleteTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("deleteCollection") {
		if err := templateExecute(deleteCollectionTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("patch") {
		if err := templateExecute(patchTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("apply") && generateApply {
		*a.importList = append(*a.importList, util.ImportFormat(fmt.Sprintf("%sapply%s", a.PkgName, a.Version), fmt.Sprintf("%s/%s/%s", a.ApplyConfigurationPackage, a.PkgName, a.Version)))
		if err := templateExecute(applyTemplate, *a); err != nil {
			return err
		}
	}

	if a.hasVerb("applyStatus") && generateApply && a.HasStatus {
		if err := templateExecute(applyStatusTemplate, *a); err != nil {
			return err
		}
	}

	for _, extension := range a.AdditionalMethods {
		if extension.Result != nil {
			name, pkg := getPkgType(*extension.Result)
			if len(pkg) > 0 {
				*a.importList = append(*a.importList, util.ImportFormat(fmt.Sprintf("%sapi", strings.ToLower(name)), pkg))
				a.ResultType = fmt.Sprintf("*%sapi.%s", strings.ToLower(name), name)
			} else {
				a.ResultType = fmt.Sprintf("*%sapi%s.%s", a.PkgName, a.Version, name)
			}

		}

		if extension.Input != nil {
			name, pkg := getPkgType(*extension.Result)
			if len(pkg) > 0 {
				*a.importList = append(*a.importList, util.ImportFormat(fmt.Sprintf("%sapi", strings.ToLower(name)), pkg))
				a.InputType = fmt.Sprintf("*%sapi.%s", strings.ToLower(name), name)
				if *extension.Verb == "apply" {
					_, gvString := codegenutil.ParsePathGroupVersion(pkg)
					*a.importList = append(*a.importList, util.ImportFormat(fmt.Sprintf("%sapplyconfig", strings.ToLower(name)), fmt.Sprintf("%s/%s", a.ApplyConfigurationPackage, gvString)))
					a.InputType = fmt.Sprintf("*%sapplyconfig.%sApplyConfiguration", strings.ToLower(name), name)
				}
			} else {
				a.InputType = fmt.Sprintf("*%sapi%s.%s", a.PkgName, a.Version, name)
				if *extension.Verb == "apply" {
					a.InputType = fmt.Sprintf("*%sapply%s.%sApplyConfiguration", a.PkgName, a.Version, name)
				}
			}
			a.InputName = strings.ToLower(name)
		}

		if *extension.Verb == "get" {
			a.Method = *extension.Method
			if err := templateExecute(getTemplate, *a); err != nil {
				return err
			}
		}

		if *extension.Verb == "create" {
			a.Method = *extension.Method
			adjTemplate := adjustTemplate(createTemplate, "create")
			if err := templateExecute(adjTemplate, *a); err != nil {
				return err
			}
		}

		if *extension.Verb == "update" {
			a.Method = *extension.Method
			adjTemplate := adjustTemplate(updateTemplate, "update")
			if err := templateExecute(adjTemplate, *a); err != nil {
				return err
			}
		}

		if *extension.Verb == "apply" {
			a.Method = *extension.Method
			adjTemplate := adjustTemplate(applyTemplate, "apply")
			if err := templateExecute(adjTemplate, *a); err != nil {
				return err
			}
		}
	}
	return nil
}

// Execute template and append the contents to the writer.
func templateExecute(templateName string, data api) error {
	namer := namer.Namer{
		Finalize: util.UpperFirst,
	}

	// Add plural naming to the function map based on a custom namer.
	funcMap["plural"] = func(input string) string {
		return namer.Name(input)
	}
	templ, err := template.New("wrapper").Funcs(funcMap).Parse(templateName)
	if err != nil {
		return err
	}
	return templ.Execute(data.writer, data)
}

// hasVerb returns true if the verb is to be scaffolded,
// if it is a part of "skipVerbs" then it returns false.
func (a *api) hasVerb(verb string) bool {
	if len(a.SkipVerbs) == 0 {
		return true
	}
	for _, s := range a.SkipVerbs {
		if verb == s {
			return false
		}
	}
	return true
}

// getPkgType returns the result override package path and the type.
func getPkgType(input string) (string, string) {
	parts := strings.Split(input, ".")
	return parts[len(parts)-1], strings.Join(parts[0:len(parts)-1], ".")
}

// appendImports adds the regular imports needed for scaffolding.
func (p *packages) appendGenericImports(importList []string) []string {
	return append(importList, `"context"`, `"fmt"`, `"k8s.io/apimachinery/pkg/types"`, `"k8s.io/client-go/rest"`,
		`"github.com/kcp-dev/logicalcluster"`, util.ImportFormat("metav1", "k8s.io/apimachinery/pkg/apis/meta/v1"), `"k8s.io/apimachinery/pkg/watch"`, util.ImportFormat("kcp", "github.com/kcp-dev/apimachinery/pkg/client"),
		util.ImportFormat(fmt.Sprintf("%sapi%s", p.Name, p.Version), p.APIPath),
		util.ImportFormat(fmt.Sprintf("%s%s", p.Name, p.Version), fmt.Sprintf("%s/typed/%s/%s", p.ClientPath, p.Name, p.Version)))
}

func adjustTemplate(template, verb string) string {
	index := strings.Index(template, "ctx context.Context,") + len("ctx context.Context,")
	newTemplate := string(template[:index]) + " name string," + string(template[index:])
	return adjReturnValueInTemplate(newTemplate, verb)
}

func adjReturnValueInTemplate(template, verb string) string {
	s := fmt.Sprintf("w.delegate.{{default .Method \"%s\"}}(ctx,", util.UpperFirst(verb))
	index := strings.Index(template, s) + len(s)
	return string(template[:index]) + " name, " + string(template[index:])
}
