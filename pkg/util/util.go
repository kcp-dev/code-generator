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

package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

// ResultTypeSupportedVerbs is a list of verb types that supports overriding the
// resulting type.
var ResultTypeSupportedVerbs = []string{
	"create",
	"update",
	"get",
	"list",
	"patch",
	"apply",
}

// UnsupportedExtensionVerbs is a list of verbs we don't support generating
// extension client functions for.
var UnsupportedExtensionVerbs = []string{
	"updateStatus",
	"deleteCollection",
	"watch",
	"delete",
}

// InputTypeSupportedVerbs is a list of verb types that supports overriding the
// input argument type.
var InputTypeSupportedVerbs = []string{
	"create",
	"update",
	"apply",
}

// DefaultValue just returns the first non-empty string among
// two inputs.
var DefaultValue = func(a, b string) string {
	if len(a) == 0 {
		return b
	}
	return a
}

// CurrentPackage returns the go package of the current directory, or "" if it cannot
// be derived from the GOPATH.
// This logic is taken from k8.io/code-generator, but has a change of letting user pass the
// directory whose pacakge is to be found.
func CurrentPackage(dir string) (string, bool) {
	goModPath, err := getGoModPath(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// hasGoMod returns true if go.mod was found in the parent dir which was
	// given as input.
	var hasGoMod bool
	if goModPath == dir {
		hasGoMod = true
	} else {
		hasGoMod = false
	}

	gomod, err := ioutil.ReadFile(filepath.Join(goModPath, "go.mod"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return modfile.ModulePath(gomod), hasGoMod
}

// getGoModPath recursively traverses up the directory path
// to find the location of go.mod file.
func getGoModPath(dir string) (string, error) {
	if _, err := os.Stat(dir); err != nil {
		return "", fmt.Errorf("error trying to find go.mod from directory %s: %w", dir, err)
	}
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
		return dir, nil
	}

	if filepath.Dir(dir) == dir {
		// Hit the root
		return "", fmt.Errorf("could not find go.mod")
	}

	return getGoModPath(filepath.Dir(dir))
}

// CleanInputDir returns a clean directory path. If
// the input is ".", it returns an empty string.
func CleanInputDir(dir string) (cleanPath string) {
	if dir == "." {
		return cleanPath
	}
	return filepath.Clean(dir)
}

// GetCleanRealtivePath checks if the outputPath already consists of
// the go's base path.If so, it returns the output path. In case it doesn't
// then it combines the base path with the output path.
// For example:
// basePath := github.com/kcp-dev/kubernetes
// outputPath := pkg/output
// It would return github.com/kcp-dev/kubernetes/pkg/output
// The other case in which:
// basePath := github.com/kcp-dev/kubernetes
// outputPath := github.com/kcp-dev/kubernetes/pkg/output
// It would return github.com/kcp-dev/kubernetes/pkg/output
func GetCleanRealtivePath(basePath, outputPath string) string {
	if strings.HasPrefix(outputPath, basePath) {
		return outputPath
	}

	return filepath.Join(basePath, filepath.Clean(outputPath))
}

// ImportFormat returns the pkgName and path formatted to be scaffolded
// for inputs.
func ImportFormat(tag, path string) string {
	return fmt.Sprintf("%s %q", tag, path)
}

// upperFirst sets the first alphabet to upperCase
func UpperFirst(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}
