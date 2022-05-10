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

package clientgen

import (
	"testing"

	"k8s.io/code-generator/cmd/client-gen/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kcp-dev/code-generator/pkg/flag"
)

var _ = Describe("Test generator funcs", func() {
	Describe("Test setting defaults", func() {
		var (
			f flag.Flags
			g *Generator
		)
		BeforeEach(func() {
			f = flag.Flags{}
			f.InputDir = "."
			f.ClientsetAPIPath = "examples"
			f.OutputDir = "."
			f.GroupVersions = []string{"apps:v1"}

			g = &Generator{}
		})

		It("should set defaults correctly", func() {
			err := g.setDefaults(f)
			Expect(err).NotTo(HaveOccurred())
			Expect(g.inputDir).To(Equal("."))
			Expect(g.clientSetAPIPath).To(Equal("examples"))
			Expect(g.outputDir).To(Equal("."))

			expected := []types.GroupVersions{{
				PackageName: "apps",
				Group:       types.Group("apps"),
				Versions: []types.PackageVersion{
					{
						Version: types.Version("v1"),
						Package: "apps/v1",
					},
				},
			}}
			Expect(g.groupVersions).To(Equal(expected))
		})
	})

	Describe("Test gv", func() {
		var (
			f flag.Flags
		)
		BeforeEach(func() {
			f = flag.Flags{}
			f.InputDir = "."
			f.GroupVersions = []string{"apps:v1", "rbac:v2"}
		})

		It("should parse Group versions without error", func() {
			expected := []types.GroupVersions{{
				PackageName: "apps",
				Group:       types.Group("apps"),
				Versions: []types.PackageVersion{
					{
						Version: types.Version("v1"),
						Package: "apps/v1",
					},
				},
			}, {
				PackageName: "rbac",
				Group:       types.Group("rbac"),
				Versions: []types.PackageVersion{
					{
						Version: types.Version("v2"),
						Package: "rbac/v2",
					},
				},
			}}

			gvs, err := GetGV(f)
			Expect(err).NotTo(HaveOccurred())
			Expect(gvs).To(Equal(expected))
		})

		It("should parse multiple Group versions without error", func() {
			f.GroupVersions = []string{"apps:v1,v2"}
			expected := []types.GroupVersions{{
				PackageName: "apps",
				Group:       types.Group("apps"),
				Versions: []types.PackageVersion{
					{
						Version: types.Version("v1"),
						Package: "apps/v1",
					},
				},
			}, {
				PackageName: "apps",
				Group:       types.Group("apps"),
				Versions: []types.PackageVersion{
					{
						Version: types.Version("v2"),
						Package: "apps/v2",
					},
				},
			}}

			gvs, err := GetGV(f)
			Expect(err).NotTo(HaveOccurred())
			Expect(gvs).To(Equal(expected))
		})

		It("should error when wrong input is provided through flag", func() {
			f.GroupVersions = []string{"apps"}

			_, err := GetGV(f)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("input to --group-version must be in <group>:<versions> format, ex: rbac:v1"))

			f.GroupVersions = []string{"apps:v1:v2"}

			_, err = GetGV(f)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("input to --group-version must be in <group>:<versions> format, ex: rbac:v1"))

		})

	})
})

func TestMetadata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test generator suite")
}
