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

import "sigs.k8s.io/controller-tools/pkg/markers"

var (
	// RuleDefinition is a marker for defining rules
	RuleDefinition = markers.Must(markers.MakeDefinition("genclient", markers.DescribesType, placeholder{}))
	// nonNamespacedMarker checks if resource is namespaced or clusterscoped
	NonNamespacedMarker = markers.Must(markers.MakeDefinition("genclient:nonNamespaced", markers.DescribesType, placeholder{}))
	// noStatusMarker checks if status is to scaffolded
	NoStatusMarker = markers.Must(markers.MakeDefinition("+genclient:noStatus", markers.DescribesType, placeholder{}))
)

// IsEnabledForMethod verifies if the genclient marker is enabled for
// this type or not.
func IsEnabledForMethod(info *markers.TypeInfo) bool {
	enabled := info.Markers.Get(RuleDefinition.Name)
	return enabled != nil
}

// IsClusterScoped verifies if the genclient marker for this
// type is namespaced or clusterscoped.
func IsClusterScoped(info *markers.TypeInfo) bool {
	enabled := info.Markers.Get(NonNamespacedMarker.Name)
	return enabled != nil
}
