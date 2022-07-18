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

package parser

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"

	"k8s.io/code-generator/cmd/client-gen/args"
	"k8s.io/code-generator/cmd/client-gen/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/kcp-dev/code-generator/pkg/flag"
)

// GetGV parses the Group Versions provided in the input through flags
// and creates a list of []types.GroupVersions.
func GetGV(f flag.Flags) ([]types.GroupVersions, error) {
	dedupGVs := map[string][]types.GroupVersions{}
	groupVersions := make([]types.GroupVersions, 0)

	// Its already validated that list of group versions cannot be empty.
	inputGVs := f.GroupVersions
	for _, gv := range inputGVs {
		// arr[0] -> group, arr[1] -> versions
		arr := strings.Split(gv, ":")
		if len(arr) != 2 {
			return nil, fmt.Errorf("input to --group-version must be in <group>:<versions> format, ex: rbac:v1. Got %q", gv)
		}
		if _, ok := dedupGVs[arr[0]]; !ok {
			dedupGVs[arr[0]] = []types.GroupVersions{}
		}

		versions := strings.Split(arr[1], ",")
		for _, v := range versions {
			// input path is converted to <inputDir>/<group>/<version>.
			// example for input directory of "k8s.io/client-go/kubernetes/pkg/apis/", it would
			// be converted to "k8s.io/client-go/kubernetes/pkg/apis/rbac/v1".
			input := filepath.Join(f.InputDir, arr[0], v)
			groups := []types.GroupVersions{}
			builder := args.NewGroupVersionsBuilder(&groups)
			_ = args.NewGVPackagesValue(builder, []string{input})

			dedupGVs[arr[0]] = append(dedupGVs[arr[0]], groups...)
		}
	}
	for _, groupversions := range dedupGVs {
		finalGV := types.GroupVersions{}

		for _, groupversion := range groupversions {
			if finalGV.PackageName == "" {
				finalGV.PackageName = groupversion.PackageName
			}
			if finalGV.Group.String() == "" {
				finalGV.Group = groupversion.Group
			}
			finalGV.Versions = append(finalGV.Versions, groupversion.Versions...)

		}
		groupVersions = append(groupVersions, finalGV)
	}
	return groupVersions, nil
}

func GetGVKs(ctx *genall.GenerationContext, inputDir string, groupVersions []types.GroupVersions) (map[Group]map[types.PackageVersion][]Kind, error) {

	gvks := map[Group]map[types.PackageVersion][]Kind{}

	for _, gv := range groupVersions {
		group := Group{Name: gv.Group.String(), GoName: gv.Group.String(), FullName: gv.Group.String()}
		for _, packageVersion := range gv.Versions {

			abs, err := filepath.Abs(inputDir)
			if err != nil {
				return nil, err
			}
			path := filepath.Join(abs, group.Name, packageVersion.String())
			pkgs, err := loader.LoadRootsWithConfig(&packages.Config{
				Dir: inputDir, Mode: packages.NeedTypesInfo,
			}, path)
			if err != nil {
				return nil, err
			}
			ctx.Roots = pkgs
			for _, root := range ctx.Roots {
				packageMarkers, _ := markers.PackageMarkers(ctx.Collector, root)
				if packageMarkers != nil {
					val, ok := packageMarkers.Get(GroupNameMarker.Name).(markers.RawArguments)
					if ok {
						group.FullName = string(val)
						groupGoName := strings.Split(group.FullName, ".")[0]
						if groupGoName != "" {
							group.GoName = groupGoName
						}
					}
				}

				// Initialize the map down here so that we can use the group with the proper GoName as the key
				if _, ok := gvks[group]; !ok {
					gvks[group] = map[types.PackageVersion][]Kind{}
				}
				if _, ok := gvks[group][packageVersion]; !ok {
					gvks[group][packageVersion] = []Kind{}
				}

				if typeErr := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {

					// if not enabled for this type, skip
					if !IsEnabledForMethod(info) {
						return
					}
					namespaced := !IsClusterScoped(info)
					gvks[group][packageVersion] = append(gvks[group][packageVersion], NewKind(info.Name, namespaced))

				}); typeErr != nil {
					return nil, typeErr
				}
			}
			sort.Slice(gvks[group][packageVersion], func(i, j int) bool {
				return gvks[group][packageVersion][i].String() < gvks[group][packageVersion][j].String()
			})
			if len(gvks[group][packageVersion]) == 0 {
				klog.Warningf("No types discovered for %s:%s, will skip generation for this GroupVersion", group.Name, packageVersion.String())
				delete(gvks[group], packageVersion)
			}
		}
		if len(gvks[group]) == 0 {
			delete(gvks, group)
		}
	}

	return gvks, nil
}
