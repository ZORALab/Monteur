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

	// Relationships list all the dependent packages and their relation.
	Relationships map[string][]string

	// VCS contains the debian/control VCS data.
	VCS *DEBVCS

	// Testsuite is the debian/control Testsuite data.
	Testsuite *DEBTestsuite

	// Changelog is the debian/changelog data.
	Changelog *DEBChangelog

	// Source is the DEBIAN/source files data.
	Source *DEBSource

	// Scripts is the DEBIAN/{pre,post}{inst,rm} files data
	Scripts map[string]string

	// Install is the DEBIAN/install files data.
	Install map[string]string

	// Rules is the DEBIAN/rules file data.
	Rules string

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

	if me.VCS == nil {
		s += styler.PortraitKV("DEB.VCS", "nil")
	} else {
		s += me.VCS.String()
	}

	if me.Testsuite == nil {
		s += styler.PortraitKV("DEB.Testsuite", "nil")
	} else {
		s += me.Testsuite.String()
	}

	if me.Changelog == nil {
		s += styler.PortraitKV("DEB.Changelog", "nil")
	} else {
		s += me.Changelog.String()
	}

	if me.Source == nil {
		s += styler.PortraitKV("DEB.Source", "nil")
	} else {
		s += me.Source.String()
	}

	for k, v := range me.Relationships {
		s += styler.PortraitKArray("DEB.Relationships."+k, v)
	}

	s += styler.PortraitKMap("DEB.Scripts", me.Scripts)
	s += styler.PortraitKMap("DEB.Install", me.Install)
	s += styler.PortraitKV("DEB.Rules", me.Rules)
	s += styler.PortraitKV("DEB.Compat", strconv.Itoa(int(me.Compat)))

	return s
}

// DEBSource is the Debian/source file data.
//
// This version was abstracted across libmonteur data structure for "Don't
// Repeat Yourself" practice.
type DEBSource struct {
	// Format is the source format.
	//
	// Either `3.0 (native)` or `3.0 (quilt)`.
	Format string

	// LocalOptions is the debian/source/local-options file data.
	LocalOptions string

	// Options is the debian/source/options file data.
	Options string

	// LintianOverrides is the debian/source/lintian-overrides file data.
	LintianOverrides string
}

func (me *DEBSource) String() (s string) {
	s = styler.PortraitKV("DEB.Source.Format", me.Format)
	s += styler.PortraitKV("DEB.Source.LocalOptions", me.LocalOptions)
	s += styler.PortraitKV("DEB.Source.Options", me.Options)
	s += styler.PortraitKV("DEB.Source.LintianOverrides",
		me.LintianOverrides)

	return s
}

// DEBChangelog is the Debian/control's Changelog data.
//
// This version was abstracted across libmonteur data structure for "Don't
// Repeat Yourself" practice.
type DEBChangelog struct {
	// Urgency sets the current .deb debian/changelog entry's urgency.
	Urgency string
}

func (me *DEBChangelog) String() (s string) {
	s = styler.PortraitKV("DEB.Changelog.Urgency", me.Urgency)

	return s
}

// DEBTestsuite is the Debian/control's Testsuite data.
//
// This version was abstracted across libmonteur data structure for "Don't
// Repeat Yourself" practice.
type DEBTestsuite struct {
	Paths []string
}

func (me *DEBTestsuite) String() (s string) {
	s = styler.PortraitKArray("DEB.Testsuite", me.Paths)

	return s
}

// DEBVCS is the DEBIAN/control's VCS data.
//
// This version was abstracted across libmonteur data structure for "Don't
// Repeat Yourself" practice.
//
// More info at:
// https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-vcs-fields
type DEBVCS struct {
	Type   string
	URL    string
	Branch string
	Path   string
}

func (me *DEBVCS) String() (s string) {
	s = styler.PortraitKV("DEB.Control.VCS.Type", me.Type)
	s += styler.PortraitKV("DEB.Control.VCS.URL", me.URL)
	s += styler.PortraitKV("DEB.Control.VCS.Branch", me.Branch)
	s += styler.PortraitKV("DEB.Control.VCS.Path", me.Path)

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

	// RulesRequiresRoot is the deb root permission for installation.
	RulesRequiresRoot string

	// PackageType is the .deb package type.
	//
	// Either 'deb' or 'udeb'.
	PackageType string

	// Essential is to decide the .deb package is essential.
	Essential bool
}

func (me *DEBControl) String() (s string) {
	s = styler.PortraitKV("DEB.Control.Standards", me.Standards)
	s += styler.PortraitKV("DEB.Control.Section", me.Section)
	s += styler.PortraitKV("DEB.Control.Priority", me.Priority)
	s += styler.PortraitKV("DEB.Control.RulesRequiresRoot",
		me.RulesRequiresRoot)
	s += styler.PortraitKV("DEB.Control.PackageType", me.PackageType)

	if me.Essential {
		s += styler.PortraitKV("DEB.Control.Essential", "true")
	} else {
		s += styler.PortraitKV("DEB.Control.Essential", "false")
	}

	return s
}

type DEBCopyrights struct {
	// The Debian copyright format.
	//
	// More info:
	//   1. https://www.debian.org/doc/debian-policy/ch-docs.html
	Format string

	// The Debian copyright disclaimer.
	Disclaimer string

	// The Debian copyright comment.
	Comment string
}

func (me *DEBCopyrights) String() (s string) {
	s = styler.PortraitKV("DEB.Copyrights.Format", me.Format)
	s += styler.PortraitKV("DEB.Copyrights.Disclaimer", me.Disclaimer)
	s += styler.PortraitKV("DEB.Copyrights.Comment", me.Comment)

	return s
}
