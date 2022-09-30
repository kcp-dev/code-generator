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
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	genutil "k8s.io/code-generator/cmd/client-gen/generators/util"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

var (
	// In controller-tool's terms marker's are defined in the following format: <makername>:<parameter>=<values>. These
	// markers are not a part of genclient, since they do not accept any values.
	GenclientMarker     = markers.Must(markers.MakeDefinition("genclient", markers.DescribesType, GenClient{}))
	NonNamespacedMarker = markers.Must(markers.MakeDefinition("genclient:nonNamespaced", markers.DescribesType, struct{}{}))

	// These markers, are not a part of "+genclient", and are defined separately because they accept a list which is comma separated. In
	// controller-tools, comma indicates another argument, as multiple arguments need to provided with a semi-colon separator.
	SkipVerbsMarker = markers.Must(markers.MakeDefinition("genclient:skipVerbs", markers.DescribesType, markers.RawArguments("")))
	OnlyVerbsMarker = markers.Must(markers.MakeDefinition("genclient:onlyVerbs", markers.DescribesType, markers.RawArguments("")))

	GroupNameMarker   = markers.Must(markers.MakeDefinition("groupName", markers.DescribesPackage, markers.RawArguments("")))
	GroupGoNameMarker = markers.Must(markers.MakeDefinition("groupGoName", markers.DescribesPackage, markers.RawArguments("")))

	// In controller-tool's terms marker's are defined in the following format: <makername>:<parameter>=<values>. These
	// markers are not a part of genclient, since they do not accept any values.
	NoStatusMarker = markers.Must(markers.MakeDefinition("genclient:noStatus", markers.DescribesType, struct{}{}))
	NoVerbsMarker  = markers.Must(markers.MakeDefinition("genclient:noVerbs", markers.DescribesType, struct{}{}))
	ReadOnlyMarker = markers.Must(markers.MakeDefinition("genclient:readonly", markers.DescribesType, struct{}{}))
)

type GenClient struct {
	Method      *string
	Verb        *string
	Subresource *string
	Input       *string
	Result      *string
}

// ClientsGeneratedForType verifies if the genclient marker is enabled for
// this type or not.
func ClientsGeneratedForType(info *markers.TypeInfo) bool {
	return info.Markers.Get(GenclientMarker.Name) != nil
}

// IsClusterScoped verifies if the genclient marker for this
// type is namespaced or clusterscoped.
func IsClusterScoped(info *markers.TypeInfo) bool {
	return info.Markers.Get(NonNamespacedMarker.Name) != nil
}

// IsNamespaced verifies if the genclient marker for this
// type is namespaced.
func IsNamespaced(info *markers.TypeInfo) bool {
	return !IsClusterScoped(info)
}

// SupportedVerbs determines which verbs the type supports
func SupportedVerbs(info *markers.TypeInfo) (sets.String, error) {
	if info.Markers.Get(NoVerbsMarker.Name) != nil {
		return sets.NewString(), nil
	}

	if info.Markers.Get(ReadOnlyMarker.Name) != nil {
		return sets.NewString(genutil.ReadonlyVerbs...), nil
	}

	extractVerbs := func(info *markers.TypeInfo, name string) ([]string, error) {
		if items := info.Markers.Get(name); items != nil {
			val, ok := items.(markers.RawArguments)
			if !ok {
				return nil, fmt.Errorf("marker defined in wrong format %q", OnlyVerbsMarker.Name)
			}
			return strings.Split(string(val), ","), nil
		}
		return nil, nil
	}

	onlyVerbs, err := extractVerbs(info, OnlyVerbsMarker.Name)
	if err != nil {
		return sets.NewString(), err
	}
	if len(onlyVerbs) > 0 {
		return sets.NewString(onlyVerbs...), nil
	}

	skipVerbs, err := extractVerbs(info, SkipVerbsMarker.Name)
	if err != nil {
		return sets.NewString(), err
	}
	return sets.NewString(genutil.SupportedVerbs...).Difference(sets.NewString(skipVerbs...)), nil
}

// SupportsVerbs determines if the type supports all the verbs.
func SupportsVerbs(info *markers.TypeInfo, verbs ...string) (bool, error) {
	supported, err := SupportedVerbs(info)
	if err != nil {
		return false, err
	}
	return supported.HasAll(verbs...), nil
}
