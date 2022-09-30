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
	"path/filepath"
	"strings"

	"github.com/kcp-dev/code-generator/pkg/internal/listergen"
	"github.com/kcp-dev/code-generator/pkg/parser"
	"github.com/kcp-dev/code-generator/pkg/util"
	"k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/gengo/namer"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

type Generator struct {
	// HeaderFile specifies the header text (e.g. license) to prepend to generated files.
	HeaderFile string `marker:",optional"`

	// Year specifies the year to substitute for " YEAR" in the header file.
	Year string `marker:",optional"`

	// APIPackagePath is the root directory under which API types exist.
	// e.g. "k8s.io/api"
	APIPackagePath string `marker:"apiPackagePath"`

	// SingleClusterListerPackagePath is the root directory under which single-cluster-aware listers exist,
	// for the case where we're only generating new code "on top" to enable multi-cluster use-cases.
	// e.g. "k8s.io/client-go/listers"
	SingleClusterListerPackagePath string `marker:",optional"`
}

func (Generator) RegisterMarkers(into *markers.Registry) error {
	return markers.RegisterAll(into,
		parser.GenclientMarker,
		parser.NonNamespacedMarker,
		parser.GroupNameMarker,
		parser.NoVerbsMarker,
		parser.ReadOnlyMarker,
		parser.SkipVerbsMarker,
		parser.OnlyVerbsMarker,
	)
}

// Generate will generate listers for all types that have generated clients and support LIST + WATCH verbs.
func (g Generator) Generate(ctx *genall.GenerationContext) error {
	var headerText string

	if g.HeaderFile != "" {
		headerBytes, err := ctx.ReadFile(g.HeaderFile)
		if err != nil {
			return err
		}
		headerText = string(headerBytes)
	}
	headerText = strings.ReplaceAll(headerText, " YEAR", " "+g.Year)

	groupVersionKinds, err := parser.CollectKinds(ctx, "list", "watch")
	if err != nil {
		return err
	}

	for group, versions := range groupVersionKinds {
		for version, kinds := range versions {
			groupInfo := toGroupVersionInfo(group, version)
			for _, kind := range kinds {
				listerDir := filepath.Join("clients", "listers", group.PackageName(), version.PackageName())
				outputFile := filepath.Join(listerDir, strings.ToLower(kind.String())+".go")
				logger := klog.Background().WithValues(
					"group", group.String(),
					"version", version.String(),
					"kind", kind.String(),
					"path", outputFile,
				)
				logger.Info("generating lister")

				if err := util.WriteGeneratedCode(ctx, headerText, &listergen.Lister{
					Group:                          groupInfo,
					APIPackagePath:                 g.APIPackagePath,
					Kind:                           kind,
					SingleClusterListerPackagePath: g.SingleClusterListerPackagePath,
				}, outputFile); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// adapted from https://github.com/kubernetes/kubernetes/blob/8f269d6df2a57544b73d5ca35e04451373ef334c/staging/src/k8s.io/code-generator/cmd/client-gen/types/helpers.go#L87-L103
func toGroupVersionInfo(group parser.Group, version types.PackageVersion) types.GroupVersionInfo {
	return types.GroupVersionInfo{
		Group:                group.Group,
		Version:              types.Version(namer.IC(version.Version.String())),
		PackageAlias:         strings.ToLower(group.GoName + version.Version.NonEmpty()),
		GroupGoName:          group.GoName,
		LowerCaseGroupGoName: namer.IL(group.GoName),
	}
}
