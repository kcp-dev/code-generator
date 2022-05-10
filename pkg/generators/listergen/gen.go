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
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"path/filepath"
	"strings"

	"github.com/kcp-dev/code-generator/pkg/flag"
	"github.com/kcp-dev/code-generator/pkg/generators/clientgen"
	genclientmarker "github.com/kcp-dev/code-generator/pkg/generators/clientgen"
	"github.com/kcp-dev/code-generator/pkg/internal/listergen"
	"github.com/kcp-dev/code-generator/pkg/util"
	"golang.org/x/tools/go/packages"
	"k8s.io/code-generator/cmd/client-gen/types"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

const (
	// GeneratorName is the name of the generator.
	GeneratorName = "lister"
)

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
	if err := markers.RegisterAll(reg, genclientmarker.RuleDefinition, genclientmarker.NonNamespacedMarker); err != nil {
		return nil, fmt.Errorf("error registering markers")
	}
	return reg, nil
}

func (g Generator) GetName() string {
	return GeneratorName
}

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

func (g *Generator) generate(ctx *genall.GenerationContext) error {
	for _, gv := range g.groupVersions {
		version := gv.Versions[0]

		// This is to accomodate the usecase wherein the apis are defined under a sub-folder inside
		// base package.
		basePkg := g.inputpkgPaths.basePackage
		if !g.inputpkgPaths.hasGoMod {
			cleanPkgPath := util.CleanInputDir(g.inputDir)
			if cleanPkgPath != "" {
				basePkg = filepath.Join(g.inputpkgPaths.basePackage, cleanPkgPath)
			}
		}

		path := filepath.Join(basePkg, gv.Group.String(), string(version.Version))

		pkgs, err := loader.LoadRootsWithConfig(&packages.Config{Dir: g.inputDir}, path)
		if err != nil {
			return err
		}

		ctx.Roots = pkgs

		for _, root := range pkgs {
			root.NeedTypesInfo()

			// this is to accomodate multiple types defined in single group
			byType := make(map[string][]byte)

			var outCommonContent bytes.Buffer
			if err := g.writeHeader(&outCommonContent); err != nil {
				root.AddError(err)
			}

			if eachTypeErr := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {
				var outContent bytes.Buffer

				// if not enabled for this type, skip
				if !clientgen.IsEnabledForMethod(info) {
					return
				}
				if err := g.writeHeader(&outContent); err != nil {
					root.AddError(err)
				}

				a, err := listergen.NewAPI(root, info, string(version.Version), gv.PackageName, path, !clientgen.IsClusterScoped(info), &outContent)
				if err != nil {
					root.AddError(err)
					return
				}

				err = a.WriteContent()
				if err != nil {
					root.AddError(err)
					return
				}

				outBytes := outContent.Bytes()
				formattedBytes, err := format.Source(outBytes)
				if err != nil {
					root.AddError(err)
				} else {
					outBytes = formattedBytes
				}
				if len(outBytes) > 0 {
					byType[info.Name] = outBytes
				}
			}); eachTypeErr != nil {
				return eachTypeErr
			}

			if len(byType) == 0 {
				return nil
			}

			for typeName, content := range byType {
				filename := strings.ToLower(typeName) + util.ExtensionGo
				err = util.WriteContent(content, filename, filepath.Join(g.outputDir, "listers", gv.Group.PackageName(), string(version.Version)))
				if err != nil {
					root.AddError(err)
				}
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
