// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deb

import (
	"fmt"
	"strings"
)

// VERControl are the strict values of version conditioning a PackageMeta.
type VERControl string

const (
	VERCONTROL_STRICTLY_EARLIER VERControl = "<<"
	VERCONTROL_EARLIER_OR_EQUAL VERControl = "<="
	VERCONTROL_EXACTLY_EQUAL    VERControl = "="
	VERCONTROL_LATER_OR_EQUAL   VERControl = ">="
	VERCONTROL_STRICTLY_LATER   VERControl = ">>"
)

// PackageMeta is the metadata of a package.
//
// This is used for building the dependency list.
type PackageMeta struct {
	// Name is the name of the package.
	//
	// This field is **COMPULSORY**.
	Name string

	// Version is the version tag.
	//
	// This field is **COMPULSORY**.
	Version *Version

	// VERControl is the condition applies to the Version.
	//
	// This field is **COMPULSORY**.
	VERControl VERControl

	// Architectures is the supported architecture conditions list.
	//
	// For repeating values (as in `amd64` and `amd64`; not `amd64` and
	// `!amd64`)  and/or empty string, PackageMeta will perform filtration
	// during `Sanitize()` execution.
	Architectures []string
}

// Parse is to process a given string and convert into its data structure.
//
// It shall return error if the input is incomprehensible.
func (me *PackageMeta) Parse(in string) (err error) {
	var list []string
	var ret string

	// parse Architectures first
	list = strings.Split(in, " [")
	switch len(list) {
	case 1:
		in = list[0]
	case 2:
		ret = strings.TrimRight(list[1], "]\r\n ")
		me.Architectures = strings.Split(ret, " ")
		in = list[0]
	default:
		return fmt.Errorf("%s: %s", ERROR_PACKAGE_PARSE_ARCH_BAD, in)
	}

	// parse Name
	list = strings.Split(in, " (")
	switch len(list) {
	case 2:
		in = list[1]
		me.Name = strings.TrimLeft(list[0], " \r\n")
	default:
		return fmt.Errorf("%s: %s", ERROR_PACKAGE_PARSE_NAME_BAD, in)
	}

	// parse VERControl
	list = strings.Split(in, " ")
	switch len(list) {
	case 2:
		me.VERControl = VERControl(list[0])
		in = strings.TrimRight(list[1], ") ")
	default:
		return fmt.Errorf("%s: %s", ERROR_PACKAGE_PARSE_VERCONTROL_BAD, in)
	}

	// parse Version
	me.Version = &Version{}
	return me.Version.Parse(in)
}

// Sanitize checks all the PackageMeta data are complying to .deb format.
//
// It shall return error if the data are not compliant.
func (me *PackageMeta) Sanitize() (err error) {
	if me.Name == "" {
		return fmt.Errorf(ERROR_PACKAGE_NAME_MISSING)
	}

	switch me.VERControl {
	case VERCONTROL_STRICTLY_EARLIER:
	case VERCONTROL_EARLIER_OR_EQUAL:
	case VERCONTROL_EXACTLY_EQUAL:
	case VERCONTROL_LATER_OR_EQUAL:
	case VERCONTROL_STRICTLY_LATER:
	default:
		return fmt.Errorf("%s: '%s'",
			ERROR_PACKAGE_VER_CONTROL_UNKNOWN,
			me.VERControl,
		)
	}

	if me.Version == nil {
		return fmt.Errorf(ERROR_PACKAGE_VER_MISSING)
	}

	err = me.Version.Sanitize()
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PACKAGE_VER_BAD, err)
	}

	if me.Architectures == nil {
		me.Architectures = []string{}
	}

	checker := map[string]bool{}
	list := []string{}
	for _, v := range me.Architectures {
		if v != "" && !checker[v] {
			list = append(list, v)
			checker[v] = true
		}
	}
	me.Architectures = list

	return nil
}

// String creates the DEBIAN/control compatible package meta value.
//
// It will panic if PackageMata is not sanitized before use and an error has
// occurred.
//
// Some example outputs would be:
//   TestApp (<< 5:0.50.0~testapp-0.50.0~upstream) [amd64]
//   TestApp (= 5:0.50.0~testapp-0.50.0~upstream) [amd64]
//   TestApp (= 5:0.50.0) [amd64]
//   TestApp (= 5:0.50.0)
func (me *PackageMeta) String() (s string) {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	s += me.Name
	s += " (" + string(me.VERControl) + " "
	s += me.Version.String()
	s += ")"

	if len(me.Architectures) > 0 {
		s += " "
		s += "[" + strings.Join(me.Architectures, " ") + "]"
	}

	return s
}
