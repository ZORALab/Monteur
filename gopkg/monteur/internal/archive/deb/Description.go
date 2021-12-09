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
)

//nolint:lll
// Description is the DEBIAN/control Description field with strict format.
//
// More info:
//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-description
type Description struct {
	// Synopsis is the strict 65 characters long FIRST line description.
	//
	// This field is **MANDATORY**.
	//
	// More info:
	//   https://www.debian.org/doc/manuals/developers-reference/best-pkging-practices.html#    bpp-pkg-synopsis
	Synopsis string

	// Body is the full text descriptions following up Synopsis.
	//
	// The max column length is 79 chracters maximum per line. That 1
	// character is for Description to prefix a space (`" "`) folding
	// requirement.
	//
	// If possible, avoid using tab (`\t`) as the effect is unpredictable.
	//
	// You can write your paragraph without needing to manually prefix space
	// (`" "`) for each line since the format will be folded automatically.
	Body string
}

// Sanitize checks all the Description data are complying to the .deb format.
//
// It shall return error if the data are not conforming to the .deb format or
// the Upstream value is empty.
func (me *Description) Sanitize() (err error) {
	switch {
	case me.Synopsis == "":
		return fmt.Errorf("%s: synopsis is empty",
			ERROR_CONTROL_DESCRIPTION_BAD,
		)
	case len(me.Synopsis) > 65:
		return fmt.Errorf("%s: synopsis more than 64 characters",
			ERROR_CONTROL_DESCRIPTION_BAD,
		)
	default:
	}

	return nil
}

// Strings creates the DEBIAN/control compatible Description: field.
//
// It will panic if the Description was not sanitized before and has error.
//
// An example output would be:
//   Description: single line synopsis
//    extended description over several lines
//    .
//    extended description over several lines
func (me *Description) String() (s string) {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	s += _FIELD_DESCRIPTION + me.Synopsis + "\n"
	s += foldSPDXText(me.Body)

	return s
}
