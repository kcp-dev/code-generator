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

package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

// CurrentPackage returns the go package of the current directory, or "" if it cannot
// be derived from the GOPATH.
// This logic is taken from k8.io/code-generator, but has a change of letting user pass the
// directory whose pacakge is to be found.
func CurrentPackage(dir string) string {
	goModPath, err := getGoModPath(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gomod, err := ioutil.ReadFile(filepath.Join(goModPath, "go.mod"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return modfile.ModulePath(gomod)
}

// getGoModPath recursively traverses up the directory path
// to find the location of go.mod file.
func getGoModPath(dir string) (string, error) {
	if dir == "/" {
		return "", fmt.Errorf("could not find go.mod")
	}
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
		return dir, nil
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
