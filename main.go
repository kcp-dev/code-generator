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
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kcp-dev/client-gen/pkg/flag"
	"github.com/kcp-dev/client-gen/pkg/generators"
	"github.com/kcp-dev/client-gen/pkg/generators/clientgen"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

var (
	allGenerators = map[string]generators.Generator{
		"client": clientgen.Generator{},
	}
)

func main() {
	f := &flag.Flags{}
	cmd := &cobra.Command{
		Use:     "code-gen",
		Short:   "Generate cluster-aware kcp wrappers around clients, listers and informers.",
		Long:    "Generate cluster-aware kcp wrappers around clients, listers and informers.",
		Example: "TODO",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no arguments provided to the command. Accepted values are clients, informers and listers.")
			}

			cmdArgs := strings.Split(args[0], ",")
			enabledGenerators := []generators.Generator{}
			for _, gName := range cmdArgs {
				if gen, ok := allGenerators[gName]; ok {
					enabledGenerators = append(enabledGenerators, gen)
				}
			}

			if len(enabledGenerators) == 0 {
				return fmt.Errorf("no generator ran.")
			}

			for _, generator := range enabledGenerators {
				reg, err := generator.RegisterMarker()
				if err != nil {
					return fmt.Errorf("error registering markers in generator %s", generator.GetName())
				}

				ctx := &genall.GenerationContext{Collector: &markers.Collector{Registry: reg}}
				if err := generator.Run(ctx, *f); err != nil {
					return err
				}
			}

			return nil
		},
	}

	f.AddTo(cmd.Flags())
	pflag.Parse()

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error running all markers: %v\n", err)
		os.Exit(1)
	}
}
