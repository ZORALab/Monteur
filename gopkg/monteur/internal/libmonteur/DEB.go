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

package libmonteur

import (
	"strconv"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/styler"
)

// DEB is the information needed for compiling a .deb package file.
//
// More info:
//   1. https://www.debian.org/doc/debian-policy/
type DEB struct {
	// Control are the data for DEBIAN/control
	Control *DEBControl

	// Copyright are the data for DEBIAN/copyright
	Copyright *DEBCopyrights

	// Depends list all the dependencies across different lifecycles.
	Depends *DEBDepends

	// Executions list all the package execution works.
	Executions *DEBExecutions

	// Source is the deb file format saved in DEBIAN/source file.
	Source string

	// Compat is the debhelper compatibility level.
	Compat uint
}

func (me *DEB) String() (s string) {
	if me.Control == nil {
		s = styler.PortraitKV("DEB.Control", "nil")
	} else {
		s = me.Control.String()
	}

	if me.Copyright == nil {
		s += styler.PortraitKV("DEB.Copyright", "nil")
	} else {
		s += me.Copyright.String()
	}

	if me.Depends == nil {
		s += styler.PortraitKV("DEB.Depends", "nil")
	} else {
		s += me.Depends.String()
	}

	if me.Executions == nil {
		s += styler.PortraitKV("DEB.Executions", "nil")
	} else {
		s += me.Executions.String()
	}

	s += styler.PortraitKV("DEB.Source", me.Source)
	s += styler.PortraitKV("DEB.Compat", strconv.Itoa(int(me.Compat)))

	return s
}

// DEBControl is the DEBIAN/control file data.
//
// This version was abstracted across libmonteur data structure for "Don't
// Repeat Yourself" practice.
//
// More info: https://www.debian.org/doc/debian-policy/ch-controlfields.html
type DEBControl struct {
	// Standards is the deb policy version that will be built with.
	//
	// More info:
	//  1. https://www.debian.org/doc/debian-policy/upgrading-checklist.html
	Standards string

	// Section is the list of section of a operating system.
	//
	// You need to refer the operating system manual (e.g. Debian, Ubuntu,
	// Manjaro, etc) for their catalog listings.
	//
	// More info:
	//  1. https://www.debian.org/doc/debian-policy/ch-archive.html
	Section string

	// Priority is the level of importances of installing the .deb package.
	//
	// Some good values:
	//   1. `required`
	//   2. `important`
	//   3. `standard`
	//   4. `optional`
	//
	// More info:
	//  1. https://www.debian.org/doc/debian-policy/ch-archive.html
	Priority string
}

func (me *DEBControl) String() (s string) {
	s = styler.PortraitKV("DEB.Control.Standards", me.Standards)
	s += styler.PortraitKV("DEB.Control.Section", me.Section)
	s += styler.PortraitKV("DEB.Control.Priority", me.Priority)

	return s
}

type DEBCopyrights struct {
	// The Debian copyright format.
	//
	// More info:
	//   1. https://www.debian.org/doc/debian-policy/ch-docs.html
	Format string
}

func (me *DEBCopyrights) String() (s string) {
	s = styler.PortraitKV("DEB.Copyrights.Format", me.Format)

	return s
}

type DEBDepends struct {
	// Build list all the dependencies needed for building the .deb package.
	Build []string

	// Dependencies list all the dependencies needed for the .deb package.
	Dependencies []string
}

func (me *DEBDepends) String() (s string) {
	s = styler.PortraitKArray("DEB.Depends.Build", me.Build)
	s += styler.PortraitKArray("DEB.Depends.Dependencies", me.Dependencies)

	return s
}

type DEBExecutions struct {
	// Rules are the contents for DEBIAN/rules
	Rules string

	// Install are the list of files install paths for DEBIAN/install
	Install []string
}

func (me *DEBExecutions) String() (s string) {
	s = styler.PortraitKV("DEB.Rules", me.Rules)
	s += styler.PortraitKArray("DEB.Install", me.Install)

	return s
}
