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
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
	"k8s.io/code-generator/cmd/client-gen/types"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/kcp-dev/code-generator/pkg/flag"
	"github.com/kcp-dev/code-generator/pkg/generators/clientgen"
	"github.com/kcp-dev/code-generator/pkg/internal/informergen"
	"github.com/kcp-dev/code-generator/pkg/util"
)

const (
	// GeneratorName is the name of the generator.
	GeneratorName = "informer"
	// packageName for typed client wrappers.
	typedPackageName = "externalversions"
)

// Assigning marker's output to a placeholder struct, to verify to
// typecast the result and make sure if it exists for the type.
type placeholder struct{}

type Generator struct {
	// inputDir is the path where types are defined.
	inputDir string
	// inputpkgPaths stores details on input directory.
	inputpkgPaths pkgPaths
	// output Dir where the wrappers are to be written.
	outputDir string
	// GroupVersions for whom the clients are to be generated.
	groupVersions []types.GroupVersions
	// headerText is the header text to be added to generated wrappers.
	// It is obtained from `--go-header-text` flag.
	headerText string
}

// TODO: Store this information in generation context, as other genrators
// may need this too.
type pkgPaths struct {
	// basePacakge path as found in go module.
	basePackage string
	// hasGoMod is a way of checking if the go.mod file is present inside
	// the input directory or not. If present the basepkg path need not be modified
	// to include the location of input directory. If not, include the location of
	// all the sub folders provided in the input directory.
	hasGoMod bool
}

func (g Generator) RegisterMarker() (*markers.Registry, error) {
	reg := &markers.Registry{}
	if err := markers.RegisterAll(reg, RuleDefinition, NonNamespacedMarker, NoStatusMarker); err != nil {
		return nil, fmt.Errorf("error registering markers")
	}
	return reg, nil
}

func (g Generator) GetName() string {
	return GeneratorName
}

// Run validates the input from the flags and sets default values, after which
// it calls the custom client genrator to create wrappers. If there are any
// errors while generating interface wrappers, it prints it out.
func (g Generator) Run(ctx *genall.GenerationContext, f flag.Flags) error {
	if err := flag.ValidateFlags(f); err != nil {
		return err
	}
	if err := g.setDefaults(f); err != nil {
		return err
	}
	if err := g.generate(ctx); err != nil {
		return err
	}

	// print all the errors consolidated from packages in the generation context.
	// skip the type errors since they occur when input path does not contain
	// go.mod files.
	hadErr := loader.PrintErrors(ctx.Roots, packages.TypeError)
	if hadErr {
		return fmt.Errorf("generator did not run successfully")
	}
	return nil
}

// setDefaults sets the default values for the generator. It also creates
// a list of group versions provided as an input.
func (g *Generator) setDefaults(f flag.Flags) (err error) {
	if f.InputDir != "" {
		g.inputDir = f.InputDir
		pkg, hasGoMod := util.CurrentPackage(f.InputDir)
		if len(pkg) == 0 {
			return fmt.Errorf("error finding the module path for this package %q", f.InputDir)
		}
		g.inputpkgPaths = pkgPaths{
			basePackage: pkg,
			hasGoMod:    hasGoMod,
		}
	}
	if f.OutputDir != "" {
		g.outputDir = f.OutputDir
	}
	g.headerText, err = util.GetHeaderText(f.GoHeaderFilePath)
	if err != nil {
		return err
	}
	gvs, err := clientgen.GetGV(f)
	if err != nil {
		return err
	}
	g.groupVersions = append(g.groupVersions, gvs...)
	return nil
}

// generate first generates the wrapper for all the interfaces provided in the input.
// Then for each type defined in the input, it recursively wraps the subsequent
// interfaces to be kcp-aware.
func (g *Generator) generate(ctx *genall.GenerationContext) error {
	err := g.writeFactory(ctx)
	if err != nil {
		return err
	}
	err = g.writeGeneric(ctx)
	if err != nil {
		return err
	}

	for _, group := range g.allGroups() {
		err = g.writeGroupInterface(ctx, group)
		if err != nil {
			return err
		}
		for _, version := range g.versionsForGroup(group) {
			err = g.writeVersionInterface(ctx, group, version)
			err = g.writeInformer(ctx, group, version)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Generator) writeHeader(out io.Writer) error {
	n, err := out.Write([]byte(g.headerText))
	if err != nil {
		return err
	}

	if n < len([]byte(g.headerText)) {
		return errors.New("header text was not written properly.")
	}
	return nil
}

func (g *Generator) allGroups() (ret []types.Group) {
	groups := map[string]types.Group{}
	for _, gv := range g.groupVersions {
		groups[gv.Group.String()] = gv.Group
	}
	for _, group := range groups {
		ret = append(ret, group)
	}
	return
}

func (g *Generator) versionsForGroup(group types.Group) (ret []types.PackageVersion) {
	versions := map[string]types.PackageVersion{}
	for _, gv := range g.groupVersions {
		v := gv.Versions[0]
		if gv.Group == group {
			versions[v.Version.String()] = v
		}
	}
	for _, version := range versions {
		ret = append(ret, version)
	}
	return
}

func (g *Generator) writeFactory(ctx *genall.GenerationContext) error {
	var out bytes.Buffer

	if err := g.writeHeader(&out); err != nil {
		return err
	}

	// TODO needs to know about each group
	t, err := informergen.NewFactory(&out, "externalversions")
	if err != nil {
		return err
	}
	t.WriteContent()

	outBytes := out.Bytes()
	formattedBytes, err := format.Source(outBytes)
	if err != nil {
		return err
	} else {
		outBytes = formattedBytes
	}

	return util.WriteContent(outBytes, "factory.go", filepath.Join(g.outputDir, "informers", typedPackageName))
}

func (g *Generator) writeGeneric(ctx *genall.GenerationContext) error {
	var out bytes.Buffer

	if err := g.writeHeader(&out); err != nil {
		return err
	}

	//TODO needs to know about each gvk

	t, err := informergen.NewGeneric(&out, "externalversions")
	if err != nil {
		return err
	}
	t.WriteContent()

	outBytes := out.Bytes()
	formattedBytes, err := format.Source(outBytes)
	if err != nil {
		return err
	} else {
		outBytes = formattedBytes
	}

	return util.WriteContent(outBytes, "generic.go", filepath.Join(g.outputDir, "informers", typedPackageName))
}

func (g *Generator) writeGroupInterface(ctx *genall.GenerationContext, group types.Group) error {
	var out bytes.Buffer

	if err := g.writeHeader(&out); err != nil {
		return err
	}

	t, err := informergen.NewGroupInterface(&out, group.String())
	if err != nil {
		return err
	}
	t.WriteContent()

	outBytes := out.Bytes()
	formattedBytes, err := format.Source(outBytes)
	if err != nil {
		return err
	} else {
		outBytes = formattedBytes
	}

	return util.WriteContent(outBytes, "interface.go", filepath.Join(g.outputDir, "informers", typedPackageName, group.String()))
}

func (g *Generator) writeVersionInterface(ctx *genall.GenerationContext, group types.Group, version types.PackageVersion) error {
	var out bytes.Buffer

	if err := g.writeHeader(&out); err != nil {
		return err
	}

	t, err := informergen.NewVersionInterface(&out, version.Version.String())
	if err != nil {
		return err
	}
	t.WriteContent()

	outBytes := out.Bytes()
	formattedBytes, err := format.Source(outBytes)
	if err != nil {
		return err
	} else {
		outBytes = formattedBytes
	}

	return util.WriteContent(outBytes, "interface.go", filepath.Join(g.outputDir, "informers", typedPackageName, group.String(), version.Version.String()))
}

func (g *Generator) writeInformer(ctx *genall.GenerationContext, group types.Group, version types.PackageVersion) error {
	basePkg := g.inputpkgPaths.basePackage
	if !g.inputpkgPaths.hasGoMod {
		cleanPkgPath := util.CleanInputDir(g.inputDir)
		if cleanPkgPath != "" {
			basePkg = filepath.Join(g.inputpkgPaths.basePackage, cleanPkgPath)
		}
	}
	path := filepath.Join(basePkg, group.String(), version.Version.String())

	pkgs, err := loader.LoadRootsWithConfig(&packages.Config{Dir: g.inputDir}, path)
	if err != nil {
		return err
	}

	// Assign the pkgs obtained from loading roots to generation context.
	// TODO: Figure out if controller-tools generation runtime can be used to
	// wire in instead.
	ctx.Roots = pkgs

	for _, root := range pkgs {
		root.NeedTypesInfo()

		// this is to accomodate multiple types defined in single group
		// byType := make(map[string][]byte)
		if eachTypeErr := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {
			var out bytes.Buffer

			// if not enabled for this type, skip
			if !IsEnabledForMethod(info) {
				return
			}
			if err := g.writeHeader(&out); err != nil {
				root.AddError(err)
				return
			}

			t, err := informergen.NewInformer(&out, version.Version.String())
			if err != nil {
				root.AddError(err)
				return
			}

			t.WriteContent()

			outBytes := out.Bytes()
			formattedBytes, err := format.Source(outBytes)
			if err != nil {
				root.AddError(err)
				return
			} else {
				outBytes = formattedBytes
			}
			err = util.WriteContent(outBytes, fmt.Sprintf("%ss.go", strings.ToLower(info.Name)), filepath.Join(g.outputDir, "informers", typedPackageName, group.String(), version.Version.String()))
			if err != nil {
				root.AddError(err)
				return
			}
		}); eachTypeErr != nil {
			return eachTypeErr
		}
	}
	return nil
}
