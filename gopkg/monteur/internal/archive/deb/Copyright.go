// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Copyright 2.0 (the "License");
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

// CopyrightFormat are the list of strict format values in Copyright.
type CopyrightFormat string

//nolint:lll
const (
	COPYRIGHT_FORMAT_1_0 CopyrightFormat = "https://www.debian.org/doc/packaging-manuals/copyright-format/1.0/"
)

// Copyright is the DEBIAN/copyright file contents data.
//
// More info:
//   https://www.debian.org/doc/packaging-manuals/copyright-format/1.0/
type Copyright struct {
	// Format is the .deb copyright compliant format.
	//
	// This field is **MANDATORY**.
	//
	// This field will only be printed in the compulsory header paragraph.
	Format CopyrightFormat

	// Name is the name of the software the copyright is being applied to.
	//
	// This field is **MANDATORY**.
	//
	// This field will only be printed in the compulsory header paragraph.
	Name string

	// Contact is the contactable entity for queries.
	//
	// This field is **MANDATORY**.
	//
	// This field will only be printed in the compulsory header paragraph.
	Contact *Entity

	// Source is the location of the copyright pathing.
	//
	// Good example: `https://www.example.com/software/project`
	Source string

	// Disclaimer is the statement to disclaim some legal responsibilities.
	//
	// This field will only be printed in the compulsory header paragraph.
	Disclaimer string

	// Comment is some comments.
	//
	// This field will only be printed in the compulsory header paragraph.
	//
	// This field is pending for string rendering as the specification did
	// not provide any example on where, when, and how to format the field.
	Comment string

	// License is the license short name, or otherwise, the short title.
	//
	// Multi-licenses shall be separated by comma (`,`); parenthesis (`()`);
	//  `and`; or `or` depending on the licensing mechanism. Example:
	//   License: MPL-1.1, GPL-2, LGPL-2.1
	//   License: MPL-1.1, GPL-2, and LGPL-2.1
	//   License: MPL-1.1 or GPL-2 or LGPL-2.1
	//   License: MPL-1.1 and (GPL-2 or LGPL-2.1)
	//
	// This field will only be printed in the compulsory header paragraph.
	License string

	// Copyright are the list of copyrights owners.
	//
	// This field will only be printed in the compulsory header paragraph.
	Copyright []*Entity

	// Licenses are the list of specific licnese.
	//
	// This field is pending for string rendering as the specification did
	// not provide any example on where, when, and how to format the field.
	Licenses []*License
}

// Sanitize checks all the Copyright data complying to the .deb format.
//
// It shall return error if any data is not compliant.
func (me *Copyright) Sanitize() (err error) {
	switch me.Format {
	case COPYRIGHT_FORMAT_1_0:
	default:
		return fmt.Errorf("%s: '%s'",
			ERROR_COPYRIGHT_FORMAT_BAD,
			me.Format,
		)
	}

	if me.Name == "" {
		return fmt.Errorf("%s: ''",
			ERROR_COPYRIGHT_UPSTREAM_NAME_BAD,
		)
	}

	if me.Contact == nil {
		return fmt.Errorf("%s: 'nil'",
			ERROR_COPYRIGHT_UPSTREAM_CONTACT_BAD,
		)
	}

	err = me.Contact.Sanitize()
	if err != nil {
		return fmt.Errorf("%s: '%s'",
			ERROR_COPYRIGHT_UPSTREAM_CONTACT_BAD,
			err,
		)
	}

	if me.Copyright == nil {
		me.Copyright = []*Entity{}
	}

	if me.Licenses == nil {
		me.Licenses = []*License{}
	}

	if len(me.Copyright) > 0 &&
		me.License == "" &&
		len(me.Licenses) == 0 {
		return fmt.Errorf(ERROR_COPYRIGHT_LICENSE_MISSING)
	}

	return nil
}

// Strings creates the DEBIAN/copyright file contents.
//
// It will panic if the Copyright struct was not sanitized before use and an
// error has occurred.
func (me *Copyright) String() (s string) {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	s += _FIELD_FORMAT + string(me.Format) + "\n"
	s += _FIELD_UPSTREAM_NAME + me.Name + "\n"
	s += _FIELD_UPSTREAM_CONTACT + me.Contact.Tag() + "\n"

	if me.Source != "" {
		s += _FIELD_SOURCE + me.Source + "\n"
	}

	if me.Disclaimer != "" {
		s += _FIELD_DISCLAIMER + me.Disclaimer + "\n"
	}

	if me.Comment != "" {
		s += _FIELD_COMMENT + me.Comment + "\n"
	}

	if me.License != "" {
		s += _FIELD_LICENSE + me.License + "\n"
	}

	for i, v := range me.Copyright {
		if i == 0 {
			s += _FIELD_COPYRIGHT
		} else {
			s += _FIELD_COPYRIGHT_FOLD
		}

		s += v.String() + "\n"
	}

	for _, v := range me.Licenses {
		s += "\n" + v.String()
	}

	return strings.TrimRight(s, "\n")
}

// License is the copyright's file paragraph license with strict format.
type License struct {
	// License is the license's short name or otherwise, the short title.
	//
	// This field is **MANDATORY**.
	//
	// Unlike header, this field shall **ONLY accept 1 (ONE) license**.
	License string

	// Body is the full license body without space prefix.
	//
	// Space prefix shall be done automatically.
	Body string

	// Comment is the optional comment field.
	//
	// This field is pending for string rendering as the specification did
	// not provide any example on where, when, and how to format the field.
	Comment string

	// Copyright are the list of copyright owners.
	//
	// This field is **MANDATORY**.
	//
	// This field **SHALL NOT BE EMPTY**.
	Copyright []*Entity

	// Files are the list of affected files for the license.
	//
	// This field is **MANDATORY**.
	//
	// This field **SHALL NOT BE EMPTY**. Symbols like asterisk (*) is
	// allowed to indiciate all pattern. In case of having the license
	// covers all the files, simple insert 1 asterisk entry is suffice.
	Files []string
}

// Sanitize is to check all data are compliant to DEBIAN/copyright requirements.
//
// If shall return error if any data is not compliant.
func (me *License) Sanitize() (err error) {
	if me.Files == nil {
		me.Files = []string{}
	}

	if me.Copyright == nil {
		me.Copyright = []*Entity{}
	}

	if len(me.Files) == 0 {
		return fmt.Errorf(ERROR_LICENSE_FILE_EMPTY)
	}

	if len(me.Copyright) == 0 {
		return fmt.Errorf(ERROR_LICENSE_COPYRIGHT_EMPTY)
	}

	cMap := map[string]bool{}
	cList := []*Entity{}
	for _, v := range me.Copyright {
		if v == nil {
			continue
		}

		if v.Year <= 0 {
			return fmt.Errorf("%s: Got '%d'",
				ERROR_LICENSE_COPYRIGHT_YEAR_BAD,
				v.Year,
			)
		}

		if !cMap[v.Name] {
			cList = append(cList, v)
			cMap[v.Name] = true
		}
	}
	me.Copyright = cList

	x := strings.ToLower(me.License)
	switch {
	case x == "":
		return fmt.Errorf(ERROR_LICENSE_EMPTY)
	case strings.ContainsAny(x, ";,"):
		return fmt.Errorf("%s: %s", ERROR_LICENSE_MULTI, me.License)
	case strings.Contains(x, " or "):
		return fmt.Errorf("%s: %s", ERROR_LICENSE_MULTI, me.License)
	case strings.Contains(x, " and "):
		return fmt.Errorf("%s: %s", ERROR_LICENSE_MULTI, me.License)
	default:
	}

	return nil
}

// String is to generate a license file paragraph for the DEBIAN/copyright.
//
// It shall panic if the License is not sanitized before use and has an error
// occurred.
func (me *License) String() (s string) {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	s += _FIELD_FILES + strings.Join(me.Files, ", ") + "\n"

	for i, v := range me.Copyright {
		if i == 0 {
			s += _FIELD_COPYRIGHT
		} else {
			s += _FIELD_COPYRIGHT_FOLD
		}

		s += v.String() + "\n"
	}

	if me.Comment != "" {
		s += _FIELD_COMMENT + me.Comment + "\n"
	}

	s += _FIELD_LICENSE + me.License + "\n"

	if me.Body != "" {
		s += foldSPDXText(me.Body)
	}

	return s
}
