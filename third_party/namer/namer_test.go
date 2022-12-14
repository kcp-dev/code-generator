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

package namer

import (
	"testing"

	"github.com/kcp-dev/code-generator/v2/pkg/util"
)

func TestNamer(t *testing.T) {
	namer := &Namer{
		Exceptions: map[string]string{
			"Blah": "Blah",
		},
		Finalize: util.UpperFirst,
	}

	cases := []struct {
		typeName string
		expected string
	}{
		{
			"Pod",
			"Pods",
		},
		{
			"Entry",
			"Entries",
		},
		{
			"Fizz",
			"Fizzes",
		},
		{
			"Blah",
			"Blah",
		},
		{
			"check",
			"Checks",
		},
		{
			"Ingress",
			"Ingresses",
		},
		{
			"ray",
			"Rays",
		},
		{
			"Fox",
			"Foxes",
		},
		{
			"City",
			"Cities",
		},
		{
			"Leaf",
			"Leaves",
		},
		{
			"Rich",
			"Riches",
		},
		{
			"Life",
			"Lives",
		},
		{
			"Myth",
			"Myths",
		},
	}

	for _, c := range cases {
		output := namer.Name(c.typeName)
		if output != c.expected {
			t.Errorf("Unexpected result from namer. Expected %s, got %s", c.expected, output)
		}
	}
}
