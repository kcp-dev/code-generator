/*
Copyright The KCP Authors.

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

package generators

import (
	"github.com/kcp-dev/client-gen/pkg/flag"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

// Generator knows how to register some set of markers, and then produce
// output artifacts based on loaded code containing those markers,
// sharing common loaded data.
type Generator interface {
	// Run uses the generation context, parses the flags from
	// the command line and generates templates.
	Run(ctx *genall.GenerationContext, f flag.Flags) error
	// RegisterMarkers registers all markers needed by this Generator
	// and returns a Registery.
	RegisterMarker() (*markers.Registry, error)
	// GetName returns the name of the generator.
	GetName() string
}
