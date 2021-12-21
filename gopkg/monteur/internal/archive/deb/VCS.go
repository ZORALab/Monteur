// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, VCS 2.0 (the "License");
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
	"net/url"
)

type VCSType string

const (
	VCS_ARCH       VCSType = "Vcs-Arch"
	VCS_BAZAAR     VCSType = "Vcs-Bzr"
	VCS_CVS        VCSType = "Vcs-Cvs"
	VCS_DARCS      VCSType = "Vcs-Darcs"
	VCS_GIT        VCSType = "Vcs-Git"
	VCS_MERCURIAL  VCSType = "Vcs-Hg"
	VCS_MONOTONE   VCSType = "Vcs-Mtn"
	VCS_SUBVERSION VCSType = "Vcs-Svn"
)

// VCS is the DEBIAN/control VCS field with strict format.
//
// More info:
//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-vcs-fields
type VCS struct {
	// Spcifies the VCS Type.
	//
	// This field is **MANDATORY**.
	Type VCSType

	// Specifies the URL address for web browsing the repository.
	Browser *url.URL

	// Specifies the URL source of the repository.
	//
	// This field is **MANDATORY**.
	URL *url.URL

	// Specifies the branch where the deb package is produced.
	Branch string

	// Specifies the directory pathing for supported VCS type.
	Path string
}

// Sanitize checks all the VCS is complying to the .deb format.
//
// It shall return error if the data are not compliant.
func (me *VCS) Sanitize() (err error) {
	switch me.Type {
	case VCS_ARCH:
	case VCS_BAZAAR:
	case VCS_CVS:
	case VCS_DARCS:
	case VCS_GIT:
	case VCS_MERCURIAL:
	case VCS_MONOTONE:
	case VCS_SUBVERSION:
	default:
		return fmt.Errorf("%s: '%s'", ERROR_VCS_TYPE_BAD, me.Type)
	}

	if me.URL == nil {
		return fmt.Errorf("%s: 'nil'", ERROR_VCS_URL_BAD)
	}

	if me.URL.String() == "" {
		return fmt.Errorf("%s: '%s'",
			ERROR_VCS_URL_BAD,
			me.URL.String(),
		)
	}

	if me.Browser != nil {
		if me.Browser.String() == "" {
			return fmt.Errorf("%s (URL): '%s'",
				ERROR_VCS_BROWSER_BAD,
				me.Browser.String(),
			)
		}
	}

	return nil
}

// Strings creates the DEBIAN/control compatible VCS field.
//
// It will panic if the VCS was not sanitized before use and an error occurred.
//
// An example output (VCS_GIT) would be:
//   Vcs-Git: 'https://example.org/repo -b debian [p/package]'
func (me *VCS) String() (s string) {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	if me.Browser != nil {
		s += _FIELD_VCS_BROWSER + me.Browser.String() + "\n"
	}

	// VCS-X
	s += string(me.Type) + ": " + me.URL.String()

	if me.Branch != "" {
		s += " -b " + me.Branch
	}

	if me.Path != "" {
		s += " [" + me.Path + "]"
	}

	return s
}
