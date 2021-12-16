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
	"regexp"
	"strconv"
)

// Version is the DEBIAN/control Version field with strict format.
//
// More info:
//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#version
type Version struct {
	// Upstream is the main version string without any prefix colon (`:`).
	//
	// The prefix colon (`:`) will be added automatically when Epoch is not
	// `0`.
	//
	// Upstream value **MUST** be:
	//   1. alphanumeric (`a-zA-Z0-9`)
	//   2. starting with digits (`0-9`)
	//   3. permitted symbols (`.~+-`)
	//   4. hyphen (`-`) is not allowed if Revision is empty.
	Upstream string

	// Revision is the tailing revision without the prefix hyphen (`-`).
	//
	// The prefix hyphen will be added automatically if Revision is not
	// empty.
	//
	// Revision value **MUST** be:
	//   1. alphanumeric (`a-zA-Z0-9`)
	//   2. starting with digits (`0-9`)
	//   3. permitted symbols (`.~+`)
	Revision string

	// Epoch is the overall scheme changes.
	//
	// Default is `0` which is unset.
	Epoch uint
}

// Sanitize checks all the Version data are complying to the .deb format.
//
// It shall return error if the data are not compliant.
func (me *Version) Sanitize() (err error) {
	var filter *regexp.Regexp

	if me.Upstream == "" {
		return fmt.Errorf(ERROR_VERSION_UPSTREAM_EMPTY)
	}

	// check Upstream
	filter = regexp.MustCompile(`^[A-Za-z0-9~.+-]*$`)
	if !filter.MatchString(me.Upstream) {
		return fmt.Errorf("%s: '%v'",
			ERROR_VERSION_UPSTREAM_ILLEGAL,
			me.Upstream,
		)
	}

	switch me.Upstream[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
	default:
		return fmt.Errorf("%s: '%v'",
			ERROR_VERSION_UPSTREAM_DIGIT_FIRST,
			me.Upstream,
		)
	}

	if me.Revision == "" {
		filter = regexp.MustCompile(`^[A-Za-z0-9~.+]*$`)
		if !filter.MatchString(me.Upstream) {
			return fmt.Errorf("%s: '%v'",
				ERROR_VERSION_UPSTREAM_NO_DASH,
				me.Upstream,
			)
		}

		return nil
	}

	// check Revision
	filter = regexp.MustCompile(`^[A-Za-z0-9~.+]*$`)
	if !filter.MatchString(me.Revision) {
		return fmt.Errorf("%s: '%v'",
			ERROR_VERSION_REVISION_ILLEGAL,
			me.Revision,
		)
	}

	switch me.Revision[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
	default:
		return fmt.Errorf("%s: '%v'",
			ERROR_VERSION_REVISION_DIGIT_FIRST,
			me.Revision,
		)
	}

	return nil
}

// String creates the DEBIAN/control compatible Version field.
//
// It will panic if the Version data is not sanitized before use and an error
// has occurred.
//
// An example output would be:
//    Version: 0.9.0rc2-1-10+b1
func (me *Version) String() (s string) {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	if me.Epoch > 0 {
		s += strconv.Itoa(int(me.Epoch)) + ":"
	}

	s += me.Upstream

	if me.Revision != "" {
		s += "-"
		s += me.Revision
	}

	return s
}
