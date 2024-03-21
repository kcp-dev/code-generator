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

package flag

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TODO: Rewrite into generic Go Testing format.
func TestMetadata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test flags suite")
}

var _ = Describe("Test flag inputs", func() {
	var (
		f Flags
	)

	BeforeEach(func() {
		f = Flags{}
		f.InputDir = "test"
		f.ClientsetAPIPath = "testdata/"
		f.GroupVersions = []string{"apps:v1"}
	})

	It("Should not error when input in set right", func() {
		Expect(ValidateFlags(f)).NotTo(HaveOccurred())
	})
	It("verify input path error", func() {
		f.InputDir = ""
		err := ValidateFlags(f)
		Expect(err.Error()).To(ContainSubstring("input path to API definition is required"))
	})

	It("verify clientsetAPI path", func() {
		f.ClientsetAPIPath = ""
		err := ValidateFlags(f)
		Expect(err.Error()).To(ContainSubstring("specifying client API path is required currently"))
	})

	It("verify group version list", func() {
		f.GroupVersions = []string{}
		err := ValidateFlags(f)
		Expect(err.Error()).To(ContainSubstring("list of group versions for which the clients are to be generated is required"))
	})

})
