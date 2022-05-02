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

package generator

import (
	"k8s.io/code-generator/cmd/client-gen/types"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

var (
	// RuleDefinition is a marker for defining rules
	RuleDefinition = markers.Must(markers.MakeDefinition("genlister", markers.DescribesType, placeholder{}))
)

type Generator struct {
	// BaseImportPath refers to the base path of the package
	inputDir string
	// Output Dir
	outputDir string
	// GroupVersions for whom the clients are to be generated
	groupVersions []types.GroupVersions

	groupVersionKinds []types.GroupVersionInfo
}

func (g *Generator) RegisterMarkers(into *markers.Registry) error {
	// TODO: implement me!
	return nil
}

func (g *Generator) Generate(gctx *GenerationContext) error {
	// TODO: implement me!
	return nil
}
