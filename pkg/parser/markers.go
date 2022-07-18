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

	GroupNameMarker = markers.Must(markers.MakeDefinition("groupName", markers.DescribesPackage, markers.RawArguments("")))

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

// IsEnabledForMethod verifies if the genclient marker is enabled for
// this type or not.
func IsEnabledForMethod(info *markers.TypeInfo) bool {
	enabled := info.Markers.Get(GenclientMarker.Name)
	return enabled != nil
}

// IsClusterScoped verifies if the genclient marker for this
// type is namespaced or clusterscoped.
func IsClusterScoped(info *markers.TypeInfo) bool {
	enabled := info.Markers.Get(NonNamespacedMarker.Name)
	return enabled != nil
}

// hasStatusSubresource verifies if updateStatus verb is to be scaffolded.
// if `noStatus` marker is present is returns false. Else it checks if
// the type has Status field.
func HasStatusSubresource(info *markers.TypeInfo) bool {
	if info.Markers.Get(NoStatusMarker.Name) != nil {
		return false
	}

	hasStatusField := false
	for _, f := range info.Fields {
		if f.Name == "Status" {
			hasStatusField = true
			break
		}
	}
	return hasStatusField
}
