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

package main

import (
	goflags "flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/kcp-dev/code-generator/pkg/flag"
	"github.com/kcp-dev/code-generator/pkg/generators"
	"github.com/kcp-dev/code-generator/pkg/generators/clientgen"
	"github.com/kcp-dev/code-generator/pkg/generators/informergen"
	"github.com/kcp-dev/code-generator/pkg/generators/listergen"
)

var (
	allGenerators = map[string]generators.Generator{
		"client":   clientgen.Generator{},
		"lister":   listergen.Generator{},
		"informer": informergen.Generator{},
	}
)

func main() {
	f := &flag.Flags{}
	cmd := &cobra.Command{
		Use:          "code-gen",
		SilenceUsage: true,
		Short:        "Generate cluster-aware kcp wrappers around clients, listers, and informers.",
		Long:         "Generate cluster-aware kcp wrappers around clients, listers, and informers.",
		Example: `Generate cluster-aware kcp clients from existing code scaffolded by k8.io/code-gen.
		For example:
		# To generate client wrappers:
		code-gen "client" --clientset-name clusterclient --go-header-file examples/header.txt 
						  --clientset-api-path=github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned 
						  --input-dir github.com/kcp-dev/code-generator/examples 
						  --output-dir examples/pkg 
						  --group-versions example:v1
		
		# To generate listers and informers (Yet to be implemented):
		code-gen "client,lister,informer" --clientset-name clusterclient --go-header-file examples/header.txt 
						  --clientset-api-path=github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned 
						  --input-dir github.com/kcp-dev/code-generator/examples 
						  --output-dir examples/pkg 
						  --group-versions example:v1
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no arguments provided to the command. Accepted values are clients, informers and listers.")
			}

			// This argument is expected to be of the form
			// "client,lister,informer". Based on the arguments,
			// enable the suitable generator.
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

	fs := goflags.NewFlagSet("klog", goflags.PanicOnError)
	klog.InitFlags(fs)
	cmd.Flags().AddGoFlagSet(fs)

	f.AddTo(cmd.Flags())

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error running all markers: %v\n", err)
		os.Exit(1)
	}
}
