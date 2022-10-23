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
	GenclientMarker     = markers.Must(markers.MakeDefinition("genclient", markers.DescribesType, extension{}))
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

type extension struct {
	Method      *string
	Verb        *string
	Subresource *string
	Input       *string
	Result      *string
}

type Extension struct {
	Method      string
	Verb        string
	Subresource string
	InputPath   string
	InputType   string
	ResultPath  string
	ResultType  string
}

// InputType returns the input override package path and the type.
func (e *extension) InputType() (string, string) {
	if e.Input == nil {
		return "", ""
	}
	parts := strings.Split(*e.Input, ".")
	return parts[len(parts)-1], strings.Join(parts[0:len(parts)-1], ".")
}

// ResultType returns the result override package path and the type.
func (e *extension) ResultType() (string, string) {
	if e.Result == nil {
		return "", ""
	}
	parts := strings.Split(*e.Result, ".")
	return parts[len(parts)-1], strings.Join(parts[0:len(parts)-1], ".")
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

	baseVerbs := sets.NewString(genutil.SupportedVerbs...)
	if info.Markers.Get(ReadOnlyMarker.Name) != nil {
		baseVerbs = sets.NewString(genutil.ReadonlyVerbs...)
	}

	if info.Markers.Get(NoStatusMarker.Name) != nil {
		baseVerbs = baseVerbs.Difference(sets.NewString("updateStatus", "applyStatus"))
	}

	skipVerbs, err := extractVerbs(info, SkipVerbsMarker.Name)
	if err != nil {
		return sets.NewString(), err
	}
	return baseVerbs.Difference(sets.NewString(skipVerbs...)), nil
}

func ClientExtensions(info *markers.TypeInfo) []Extension {
	values, ok := info.Markers[GenclientMarker.Name]
	if !ok || values == nil {
		return nil
	}
	var extensions []Extension
	for _, item := range values {
		extension, ok := item.(extension)
		if !ok {
			continue // should not occur
		}
		if extension.Method == nil || *extension.Method == "" {
			continue
		}
		transformed := Extension{
			Method:      deref(extension.Method),
			Verb:        deref(extension.Verb),
			Subresource: deref(extension.Subresource),
		}
		transformed.InputType, transformed.InputPath = extension.InputType()
		transformed.ResultType, transformed.ResultPath = extension.ResultType()
		extensions = append(extensions, transformed)
	}
	return extensions
}

func deref(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

// SupportsVerbs determines if the type supports all the verbs.
func SupportsVerbs(info *markers.TypeInfo, verbs ...string) (bool, error) {
	supported, err := SupportedVerbs(info)
	if err != nil {
		return false, err
	}
	return supported.HasAll(verbs...), nil
}
