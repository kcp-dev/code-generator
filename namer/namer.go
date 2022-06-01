/*
Copyright 2015 The Kubernetes Authors.

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

// This code has been taken from https://github.com/kubernetes/gengo/blob/master/namer/plural_namer.go
// with minor modifications. The changes are:
// 1. There is no concept of plublic/private namer here. There is generic namer struct which
// parses the input and gives us the required plural form.
// 2. The input argument to function `Name` is customized to accept a string instead of `types.Type`,
// since we directly modeify the API name in our code-gen.

package namer

var consonants = "bcdfghjklmnpqrstvwxyz"

type Namer struct {
	// use this to add any exceptions to look up
	Exceptions map[string]string
	// Use this to either convert everything to lowercase/upper case
	// or anything else based on whether the parameter will be public
	// or private.
	Finalize func(string) string
}

// Name gives out the final plural names which are
// to be scaffolded
func (n *Namer) Name(input string) string {
	singular := input
	var plural string
	var ok bool
	if plural, ok = n.Exceptions[singular]; ok {
		return n.Finalize(plural)
	}
	if len(singular) < 2 {
		return n.Finalize(singular)
	}

	switch rune(singular[len(singular)-1]) {
	case 's', 'x', 'z':
		plural = esPlural(singular)
	case 'y':
		sl := rune(singular[len(singular)-2])
		if isConsonant(sl) {
			plural = iesPlural(singular)
		} else {
			plural = sPlural(singular)
		}
	case 'h':
		sl := rune(singular[len(singular)-2])
		if sl == 'c' || sl == 's' {
			plural = esPlural(singular)
		} else {
			plural = sPlural(singular)
		}
	case 'e':
		sl := rune(singular[len(singular)-2])
		if sl == 'f' {
			plural = vesPlural(singular[:len(singular)-1])
		} else {
			plural = sPlural(singular)
		}
	case 'f':
		plural = vesPlural(singular)
	default:
		plural = sPlural(singular)
	}
	return n.Finalize(plural)
}

func iesPlural(singular string) string {
	return singular[:len(singular)-1] + "ies"
}

func vesPlural(singular string) string {
	return singular[:len(singular)-1] + "ves"
}

func esPlural(singular string) string {
	return singular + "es"
}

func sPlural(singular string) string {
	return singular + "s"
}

func isConsonant(char rune) bool {
	for _, c := range consonants {
		if char == c {
			return true
		}
	}
	return false
}
