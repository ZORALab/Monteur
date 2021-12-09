// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Testsuite 2.0 (the "License");
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
	"os"
	"path/filepath"
	"strings"
)

//nolint:lll
// Testsuite is the DEBIAN/control Testsuite: field with strict format.
//
// More info:
//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-vcs-fields
type Testsuite struct {
	// Basepath is the base directory of the package.
	//
	// If left unchecked, basepath will be set to current directory.
	Basepath string

	// Paths is the list of pathing pointing to the testsuite.
	//
	// This field is **COMPULSORY**.
	Paths []string
}

// Sanitize checks all the Testsuite is complying to the .deb format.
//
// It shall return error if the data are not compliant.
func (me *Testsuite) Sanitize() (err error) {
	if me.Paths == nil {
		me.Paths = []string{}
	}

	if len(me.Paths) == 0 {
		return fmt.Errorf("%s: empty paths", ERROR_TESTSUITE_BAD)
	}

	checker := map[string]bool{}
	list := []string{}
	for _, v := range me.Paths {
		if me.Basepath == "" {
			me.Basepath, _ = os.Getwd()
		}
		path := filepath.Join(me.Basepath, v)

		if _, err = os.Stat(path); err != nil {
			return fmt.Errorf("%s: '%s'", ERROR_TESTSUITE_BAD, err)
		}

		ok := checker[v]
		if !ok {
			list = append(list, v)
			checker[v] = true
		}
	}
	me.Paths = list

	return nil
}

// Strings creates the DEBIAN/control compatible Testsuite field's value.
//
// It will panic if the Testsuite struct was not sanitized before use and an
// error has occurred.
//
// An example output (Testsuite_GIT) would be:
//   Testsuite: test/file, test/util
func (me *Testsuite) String() string {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	return strings.Join(me.Paths, ", ")
}
