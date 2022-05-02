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

package flag

import "github.com/spf13/pflag"

// Flags - Options accepted by generator
type Flags struct {
	OutputDir     string
	InputDir      string
	GroupVersions *[]string
}

func (f *Flags) AddTo(flagset *pflag.FlagSet) {
	flagset.StringVar(&f.InputDir, "input-dir", "", "Input directory where types are defined. It is assumed that 'types.go' is present inside <InputDir>/pkg/apis.")
	flagset.StringVar(&f.OutputDir, "output-dir", "output", "Output directory where wrapped clients will be generated. The wrappers will be present in '<output-dir>/generated' path.")
	gv := flagset.StringSlice("group-versions", []string{}, "specify group versions for the clients")
	f.GroupVersions = gv
}
