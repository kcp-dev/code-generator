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
	"path/filepath"
	"sort"
	"strings"

	"k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/gengo/v2/namer"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/kcp-dev/code-generator/v2/pkg/internal/clientgen"
	"github.com/kcp-dev/code-generator/v2/pkg/parser"
	"github.com/kcp-dev/code-generator/v2/pkg/util"
)

type Generator struct {
	// Name is the name of this client-set, e.g. "kubernetes"
	Name string `marker:",optional"`

	// ExternalOnly toggles the creation of a "versioned" sub-directory. Set to true if you are generating
	// custom code for a project that's not using k8s.io/code-generator/generate-groups.sh for their types.
	ExternalOnly bool `marker:",optional"`

	// Standalone toggles the creation of a "cluster" sub-directory. Set to true if you are placing cluster-
	// aware code somewhere outside of the normal client tree.
	Standalone bool `marker:",optional"`

	// HeaderFile specifies the header text (e.g. license) to prepend to generated files.
	HeaderFile string `marker:",optional"`

	// Year specifies the year to substitute for " YEAR" in the header file.
	Year string `marker:",optional"`

	// SingleClusterClientPackagePath is the root directory under which single-cluster-aware clients exist.
	// e.g. "k8s.io/client-go/kubernetes"
	SingleClusterClientPackagePath string `marker:""`

	// SingleClusterApplyConfigurationsPackagePath is the root directory under which single-cluster-aware apply configurations exist.
	// e.g. "k8s.io/client-go/applyconfigurations"
	SingleClusterApplyConfigurationsPackagePath string `marker:",optional"`

	// OutputPackagePath is the root directory under which this tool will output files.
	// e.g. "github.com/kcp-dev/client-go/clients"
	OutputPackagePath string `marker:""`

	// APIPackagePath is the root directory under which API types exist.
	// e.g. "k8s.io/api"
	APIPackagePath string `marker:"apiPackagePath"`
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

// Generate will generate clients for all types that have generated clients.
func (g Generator) Generate(ctx *genall.GenerationContext) error {
	var headerText string

	if g.HeaderFile != "" {
		headerBytes, err := ctx.ReadFile(g.HeaderFile)
		if err != nil {
			return err
		}
		headerText = string(headerBytes)
	}
	var replacement string
	if g.Year != "" {
		replacement = " " + g.Year
	}
	headerText = strings.ReplaceAll(headerText, " YEAR", replacement)

	if g.Name == "" {
		g.Name = "clientset"
	}

	groupVersionKinds, err := parser.CollectKinds(ctx)
	if err != nil {
		return err
	}

	groupInfo := toGroupVersionInfos(groupVersionKinds)

	clientsetDir := g.Name
	if !g.ExternalOnly {
		clientsetDir = filepath.Join(clientsetDir, "versioned")
	}
	if !g.Standalone {
		clientsetDir = filepath.Join(clientsetDir, "cluster")
	}
	clientsetFile := filepath.Join(clientsetDir, "clientset.go")
	logger := klog.Background().WithValues("clientset", g.Name)
	logger.WithValues("path", clientsetFile).Info("generating clientset")

	if err := util.WriteGeneratedCode(ctx, headerText, &clientgen.ClientSet{
		Name:                           g.Name,
		PackagePath:                    filepath.Join(g.OutputPackagePath, clientsetDir),
		Groups:                         groupInfo,
		SingleClusterClientPackagePath: g.SingleClusterClientPackagePath,
	}, clientsetFile); err != nil {
		return err
	}

	// TODO: do we actually need the scheme? no, right? k8s client-gen will do it for you
	schemeDir := filepath.Join(clientsetDir, "scheme")
	schemeFile := filepath.Join(schemeDir, "register.go")
	logger.WithValues("path", schemeFile).Info("generating scheme")

	if err := util.WriteGeneratedCode(ctx, headerText, &clientgen.Scheme{
		Groups:         groupInfo,
		APIPackagePath: g.APIPackagePath,
	}, schemeFile); err != nil {
		return err
	}

	fakeClientsetDir := filepath.Join(clientsetDir, "fake")
	fakeClientsetFile := filepath.Join(fakeClientsetDir, "clientset.go")
	logger.WithValues("path", fakeClientsetFile).Info("generating scheme")

	if err := util.WriteGeneratedCode(ctx, headerText, &clientgen.FakeClientset{
		Name:                           g.Name,
		PackagePath:                    filepath.Join(g.OutputPackagePath, clientsetDir),
		Groups:                         groupInfo,
		SingleClusterClientPackagePath: g.SingleClusterClientPackagePath,
	}, fakeClientsetFile); err != nil {
		return err
	}

	for group, versions := range groupVersionKinds {
		for version, kinds := range versions {
			groupDir := filepath.Join(clientsetDir, "typed", group.PackageName(), version.PackageName())
			outputFile := filepath.Join(groupDir, group.PackageName()+"_client.go")
			logger := logger.WithValues(
				"group", group.String(),
				"version", version.String(),
			)
			logger.WithValues("path", outputFile).Info("generating group client")
			groupInfo := toGroupVersionInfo(group, version)
			if err := util.WriteGeneratedCode(ctx, headerText, &clientgen.Group{
				Group:                          groupInfo,
				Kinds:                          kinds,
				SingleClusterClientPackagePath: g.SingleClusterClientPackagePath,
			}, outputFile); err != nil {
				return err
			}

			fakeGroupDir := filepath.Join(groupDir, "fake")
			outputFile = filepath.Join(fakeGroupDir, group.PackageName()+"_client.go")
			logger = logger.WithValues(
				"group", group.String(),
				"version", version.String(),
			)
			logger.WithValues("path", outputFile).Info("generating fake group client")
			if err := util.WriteGeneratedCode(ctx, headerText, &clientgen.FakeGroup{
				Group:                          groupInfo,
				Kinds:                          kinds,
				PackagePath:                    filepath.Join(g.OutputPackagePath, clientsetDir),
				SingleClusterClientPackagePath: g.SingleClusterClientPackagePath,
			}, outputFile); err != nil {
				return err
			}

			for _, kind := range kinds {
				outputFile := filepath.Join(groupDir, strings.ToLower(kind.String())+".go")
				logger := logger.WithValues(
					"kind", kind.String(),
				)
				logger.WithValues("path", outputFile).Info("generating client for kind")

				if err := util.WriteGeneratedCode(ctx, headerText, &clientgen.TypedClient{
					Group:                          groupInfo,
					Kind:                           kind,
					APIPackagePath:                 g.APIPackagePath,
					SingleClusterClientPackagePath: g.SingleClusterClientPackagePath,
				}, outputFile); err != nil {
					return err
				}

				outputFile = filepath.Join(fakeGroupDir, strings.ToLower(kind.String())+".go")
				logger = logger.WithValues(
					"kind", kind.String(),
				)
				logger.WithValues("path", outputFile).Info("generating fake client for kind")

				if err := util.WriteGeneratedCode(ctx, headerText, &clientgen.FakeTypedClient{
					Group:                          groupInfo,
					Kind:                           kind,
					APIPackagePath:                 g.APIPackagePath,
					PackagePath:                    filepath.Join(g.OutputPackagePath, clientsetDir),
					SingleClusterClientPackagePath: g.SingleClusterClientPackagePath,
					SingleClusterApplyConfigurationsPackagePath: g.SingleClusterApplyConfigurationsPackagePath,
				}, outputFile); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// adapted from https://github.com/kubernetes/kubernetes/blob/8f269d6df2a57544b73d5ca35e04451373ef334c/staging/src/k8s.io/code-generator/cmd/client-gen/types/helpers.go#L87-L103
func toGroupVersionInfos(groupVersionKinds map[parser.Group]map[types.PackageVersion][]parser.Kind) []types.GroupVersionInfo {
	var info []types.GroupVersionInfo
	for group, versions := range groupVersionKinds {
		for version := range versions {
			info = append(info, toGroupVersionInfo(group, version))
		}
	}
	sort.Slice(info, func(i, j int) bool {
		return info[i].PackageAlias < info[j].PackageAlias
	})
	return info
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
