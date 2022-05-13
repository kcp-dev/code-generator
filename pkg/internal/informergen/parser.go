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

package informergen

import (
	"io"
	"text/template"
)

type factory struct {
	writer      *io.Writer
	PackageName string
}

func (f *factory) WriteContent() error {
	templ, err := template.New("factory").Parse(factoryTmpl)
	if err != nil {
		return err
	}
	return templ.Execute(*f.writer, f)
}

func NewFactory(w io.Writer, packageName string) (*factory, error) {
	return &factory{writer: &w, PackageName: packageName}, nil
}

type generic struct {
	writer      *io.Writer
	PackageName string
}

func (g *generic) WriteContent() error {
	templ, err := template.New("generic").Parse(genericTmpl)
	if err != nil {
		return err
	}
	return templ.Execute(*g.writer, g)
}

func NewGeneric(w io.Writer, packageName string) (*generic, error) {
	return &generic{writer: &w, PackageName: packageName}, nil
}

type groupInterface struct {
	writer      *io.Writer
	PackageName string
}

func (g *groupInterface) WriteContent() error {
	templ, err := template.New("groupInterface").Parse(groupInterfaceTmpl)
	if err != nil {
		return err
	}
	return templ.Execute(*g.writer, g)
}

func NewGroupInterface(w io.Writer, packageName string) (*groupInterface, error) {
	return &groupInterface{writer: &w, PackageName: packageName}, nil
}

type versionInterface struct {
	writer      *io.Writer
	PackageName string
}

func (v *versionInterface) WriteContent() error {
	templ, err := template.New("versionInterface").Parse(versionInterfaceTmpl)
	if err != nil {
		return err
	}
	return templ.Execute(*v.writer, v)
}

func NewVersionInterface(w io.Writer, packageName string) (*versionInterface, error) {
	return &versionInterface{writer: &w, PackageName: packageName}, nil
}

type informer struct {
	writer      *io.Writer
	PackageName string
}

func (i *informer) WriteContent() error {
	templ, err := template.New("informer").Parse(informerTmpl)
	if err != nil {
		return err
	}
	return templ.Execute(*i.writer, i)
}

func NewInformer(w io.Writer, packageName string) (*informer, error) {
	return &informer{writer: &w, PackageName: packageName}, nil
}
