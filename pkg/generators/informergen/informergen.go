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
	"path/filepath"
	"sort"
	"strings"

	"k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/gengo/v2/namer"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/kcp-dev/code-generator/v2/pkg/internal/informergen"
	"github.com/kcp-dev/code-generator/v2/pkg/parser"
	"github.com/kcp-dev/code-generator/v2/pkg/util"
)

type Generator struct {
	// Name is the name of the clientset, e.g. "kubernetes"
	ClientsetName string `marker:",optional"`

	// ExternalOnly toggles the creation of a "externalversions" sub-directory. Set to true if you are generating
	// custom code for a project that's not using k8s.io/code-generator/generate-groups.sh for their types.
	ExternalOnly bool `marker:",optional"`

	// Standalone toggles the creation of a "cluster" sub-directory. Set to true if you are placing cluster-
	// aware code somewhere outside of the normal client tree.
	Standalone bool `marker:",optional"`

	// HeaderFile specifies the header text (e.g. license) to prepend to generated files.
	HeaderFile string `marker:",optional"`

	// Year specifies the year to substitute for " YEAR" in the header file.
	Year string `marker:",optional"`

	// OutputPackagePath is the root directory under which this tool will output files.
	// e.g. "github.com/kcp-dev/client-go/clients"
	OutputPackagePath string `marker:""`

	// APIPackagePath is the root directory under which API types exist.
	// e.g. "k8s.io/api"
	APIPackagePath string `marker:"apiPackagePath"`

	// SingleClusterClientPackagePath is the root directory under which single-cluster-aware clients exist.
	// e.g. "k8s.io/client-go/kubernetes"
	SingleClusterClientPackagePath string `marker:""`

	// SingleClusterInformerPackagePath is the package under which the cluster-unaware listers are exposed.
	// e.g. "k8s.io/client-go/informers"
	SingleClusterInformerPackagePath string `marker:",optional"`

	// SingleClusterListerPackagePath is the root directory under which single-cluster-aware listers exist,
	// for the case where we're only generating new code "on top" to enable multi-cluster use-cases.
	// e.g. "k8s.io/client-go/listers"
	SingleClusterListerPackagePath string `marker:",optional"`
}

func (g Generator) RegisterMarkers(into *markers.Registry) error {
	return markers.RegisterAll(into,
		parser.GenclientMarker(),
		parser.NonNamespacedMarker(),
		parser.GroupNameMarker(),
		parser.NoVerbsMarker(),
		parser.ReadOnlyMarker(),
		parser.SkipVerbsMarker(),
		parser.OnlyVerbsMarker(),
	)
}

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

	groupVersionKinds, err := parser.CollectKinds(ctx, "list", "watch")
	if err != nil {
		return err
	}

	groupInfo := toGroupVersionInfos(groupVersionKinds)

	logger := klog.Background()

	onlyGroups := make([]parser.Group, 0, len(groupVersionKinds))
	for group := range groupVersionKinds {
		onlyGroups = append(onlyGroups, group)
	}
	sort.Slice(onlyGroups, func(i, j int) bool {
		return onlyGroups[i].Group.PackageName() < onlyGroups[j].Group.PackageName()
	})

	if g.ClientsetName == "" {
		g.ClientsetName = "clientset"
	}

	clientsetDir := g.ClientsetName
	if !g.ExternalOnly {
		clientsetDir = filepath.Join(clientsetDir, "versioned")
	}
	if !g.Standalone {
		clientsetDir = filepath.Join(clientsetDir, "cluster")
	}
	listersDir := "listers"

	informersDir := "informers"
	if !g.ExternalOnly {
		informersDir = filepath.Join(informersDir, "externalversions")
	}
	factoryPath := filepath.Join(informersDir, "factory.go")
	logger.WithValues("path", factoryPath).Info("generating informer factory")
	if err := util.WriteGeneratedCode(ctx, headerText, &informergen.Factory{
		Groups:                           onlyGroups,
		PackagePath:                      filepath.Join(g.OutputPackagePath, informersDir),
		ClientsetPackagePath:             filepath.Join(g.OutputPackagePath, clientsetDir),
		SingleClusterClientPackagePath:   g.SingleClusterClientPackagePath,
		SingleClusterInformerPackagePath: g.SingleClusterInformerPackagePath,
	}, factoryPath); err != nil {
		return err
	}

	gvks := map[types.Group]map[parser.Version][]parser.Kind{}
	for group, versions := range groupVersionKinds {
		for version, kinds := range versions {
			info := toGroupVersionInfo(group, version)
			if _, exists := gvks[info.Group]; !exists {
				gvks[info.Group] = map[parser.Version][]parser.Kind{}
			}
			gvks[info.Group][info.Version] = kinds
		}
	}
	genericPath := filepath.Join(informersDir, "generic.go")
	logger.WithValues("path", factoryPath).Info("generating generic informers")
	if err := util.WriteGeneratedCode(ctx, headerText, &informergen.Generic{
		Groups:                           groupInfo,
		GroupVersionKinds:                gvks,
		APIPackagePath:                   g.APIPackagePath,
		SingleClusterInformerPackagePath: g.SingleClusterInformerPackagePath,
	}, genericPath); err != nil {
		return err
	}

	interfacesPath := filepath.Join(informersDir, "internalinterfaces", "factory_interfaces.go")
	logger.WithValues("path", factoryPath).Info("generating internal informer interfaces")
	if err := util.WriteGeneratedCode(ctx, headerText, &informergen.FactoryInterface{
		ClientsetPackagePath:           filepath.Join(g.OutputPackagePath, clientsetDir),
		SingleClusterClientPackagePath: g.SingleClusterClientPackagePath,
		UseUpstreamInterfaces:          g.SingleClusterInformerPackagePath != "",
	}, interfacesPath); err != nil {
		return err
	}

	for group, versions := range groupVersionKinds {
		groupDir := filepath.Join(informersDir, group.PackageName())
		outputFile := filepath.Join(groupDir, "interface.go")
		logger := logger.WithValues(
			"group", group.String(),
		)

		var onlyVersions []parser.Version
		for version := range versions {
			onlyVersions = append(onlyVersions, parser.Version(namer.IC(version.Version.String())))
		}
		sort.Slice(onlyVersions, func(i, j int) bool {
			return onlyVersions[i].PackageName() < onlyVersions[j].PackageName()
		})

		logger.WithValues("path", outputFile).Info("generating group interface")
		if err := util.WriteGeneratedCode(ctx, headerText, &informergen.GroupInterface{
			Group:                 group,
			Versions:              onlyVersions,
			PackagePath:           filepath.Join(g.OutputPackagePath, informersDir),
			UseUpstreamInterfaces: g.SingleClusterInformerPackagePath != "",
		}, outputFile); err != nil {
			return err
		}
		for version, kinds := range versions {
			versionDir := filepath.Join(groupDir, version.PackageName())
			outputFile := filepath.Join(versionDir, "interface.go")
			logger := logger.WithValues(
				"version", version.String(),
			)

			logger.WithValues("path", outputFile).Info("generating version interface")
			if err := util.WriteGeneratedCode(ctx, headerText, &informergen.VersionInterface{
				Version:               types.Version(namer.IC(version.Version.String())),
				Kinds:                 kinds,
				PackagePath:           filepath.Join(g.OutputPackagePath, informersDir),
				UseUpstreamInterfaces: g.SingleClusterInformerPackagePath != "",
			}, outputFile); err != nil {
				return err
			}

			for _, kind := range kinds {
				outputFile := filepath.Join(versionDir, strings.ToLower(kind.String())+".go")
				logger := logger.WithValues(
					"kind", kind.String(),
				)
				logger.WithValues("path", outputFile).Info("generating informer for kind")

				if err := util.WriteGeneratedCode(ctx, headerText, &informergen.Informer{
					Group:                            toGroupVersionInfo(group, version),
					Kind:                             kind,
					APIPackagePath:                   g.APIPackagePath,
					PackagePath:                      filepath.Join(g.OutputPackagePath, informersDir),
					ClientsetPackagePath:             filepath.Join(g.OutputPackagePath, clientsetDir),
					ListerPackagePath:                filepath.Join(g.OutputPackagePath, listersDir),
					SingleClusterClientPackagePath:   g.SingleClusterClientPackagePath,
					SingleClusterInformerPackagePath: g.SingleClusterInformerPackagePath,
					SingleClusterListerPackagePath:   g.SingleClusterListerPackagePath,
				}, outputFile); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// adapted from https://github.com/kubernetes/kubernetes/blob/8f269d6df2a57544b73d5ca35e04451373ef334c/staging/src/k8s.io/code-generator/cmd/client-gen/types/helpers.go#L87-L103
func toGroupVersionInfos(groupVersionKinds map[parser.Group]map[types.PackageVersion][]parser.Kind) []parser.Group {
	var info []parser.Group
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
func toGroupVersionInfo(group parser.Group, version types.PackageVersion) parser.Group {
	return parser.Group{
		Group:                group.Group,
		PackageAlias:         strings.ToLower(group.GoName + version.Version.NonEmpty()),
		Version:              parser.Version(namer.IC(version.Version.String())),
		GoName:               group.GoName,
		LowerCaseGroupGoName: namer.IL(group.GoName),
	}
}
