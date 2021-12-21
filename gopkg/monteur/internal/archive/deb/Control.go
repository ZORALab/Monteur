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
	"net/url"
	"strings"
)

// StandardsVersions are the Debian supported .deb standards version values.
//
// See: https://www.debian.org/doc/debian-policy/upgrading-checklist.html
//
// In any cases, raise a ticket to increase standards versions matching to the
// latest Debian.
type StandardsVersion string

const (
	STANDARD_VER_4_6_0 StandardsVersion = "4.6.0"
	STANDARD_VER_4_5_1 StandardsVersion = "4.5.1"
	STANDARD_VER_4_5_0 StandardsVersion = "4.5.0"
	STANDARD_VER_4_4_1 StandardsVersion = "4.4.1"
	STANDARD_VER_4_3_0 StandardsVersion = "4.3.0"
	STANDARD_VER_4_2_1 StandardsVersion = "4.2.1"
)

// Priority is the package installation priority.
//
// More info:
//   https://www.debian.org/doc/debian-policy/ch-archive.html#priorities
type Priority string

const (
	PRIORITY_REQUIRED  Priority = "required"
	PRIORITY_IMPORTANT Priority = "important"
	PRIORITY_STANDARD  Priority = "standard"
	PRIORITY_OPTIONAL  Priority = "optional"
)

// PackageType is the package type.
//
// More info:
//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-package-type
type PackageType string

const (
	PACKAGE_TYPE_DEB  = "deb"
	PACKAGE_TYPE_UDEB = "udeb"
)

// Rules-Requires-Root values list are the prefixed value for that fields.
//
// More info:
//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#rules-requires-root
const (
	RULES_ROOT_NO             = "no"
	RULES_ROOT_BINARY_TARGETS = "binary-targets"
)

// Control is the DEBIAN/control file data.
type Control struct {
	// Maintainer is the Entity who is responsible for this package.
	//
	// This field is **MANDATORY**.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#maintainer
	Maintainer *Entity

	// Packages are the relational packages to this package.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-relationships.html
	Packages map[PackageListType]*PackageList

	// Version is the .deb version data.
	Version *Version

	// Homepage is the website of the software.
	Homepage *url.URL

	//nolint:lll
	// Description is the long description with the strict Synopsis.
	//
	// This field is **MANDATORY**.
	//
	// More info:
	//   https://www.debian.org/doc/manuals/developers-reference/best-pkging-practices.html#bpp-pkg-synopsis
	Description *Description

	// VCS is the `VCS-X:` field data.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-vcs-fields
	VCS *VCS

	// Testsuite is the `Testsuite:` field data.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-vcs-fields
	Testsuite *Testsuite

	// Name is the name of the package used for `Source` and `Package`.
	//
	// This field is **MANDATORY**.
	//
	// It can only contain [A-Z], [a-z], [0-9], and [+.-] characters only.
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#source
	Name string

	// Section is the classification of the package (category).
	//
	// You should get the catalog list from your targeted operating system.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-archive.html#s-subsections
	Section string

	// StandardsVersion is the Debian package standards version to comply.
	//
	// This field is **MANDATORY**.
	StandardsVersion StandardsVersion

	// PackageType is the `Package-Type:` field.
	//
	// It accepts `deb` or `udeb`.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-package-type
	PackageType PackageType

	// RulesRequiresRoot is the `Rules-Requires-Root` field.
	//
	// If custom values are provided, a forward slash (`/`) is required
	// such that (`<namespace>/<case>`) pattern is complied and
	// multiple values are seaprated by comma.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#rules-requires-root
	RulesRequiresRoot string

	// Priority is the package installation priority.
	//
	// Only accepts `required`, `important`, `standard`, and `optional`.
	// When in doubts, `optional` is the right value. The other 3 have
	// specific definitions according to the operating system.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-archive.html#priorities
	Priority Priority

	// Architecture is the package supported CPU architecture.
	//
	// It **MUST** match Debian machine architecture strings like `any`,
	// `amd64`, and etc.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-customized-programs.html#s-arch-spec
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#architecture
	Architecture string

	// Uploaders are the list of co-maintainers.
	//
	// This field is optional.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#uploaders
	Uploaders []*Entity

	// Essential is to set the package to be prohibited from uninstall.
	//
	// This field is optional.
	//
	// More info:
	//   https://www.debian.org/doc/debian-policy/ch-controlfields.html#essential
	Essential bool

	// BuildSource is the flag to indicate for building source package.
	//
	// The default (`false`) is to build binary package.
	//
	// This is used by this Go's deb package internally as it was a missed
	// out data.
	BuildSource bool
}

// Sanitize() ensures the filled data is complying to the DEBIAN/control format.
//
// It will return an error if any data is not compliant and a warning report in
// case there are acceptable deviant.
func (me *Control) Sanitize() (warn string, err error) {
	var x string

	me.sanitizeHomepage()

	if me.Name == "" {
		return warn, fmt.Errorf("%s: %s", ERROR_CONTROL_NAME_BAD, "''")
	}
	me.Name = strings.ToLower(me.Name)

	err = me.sanitizeMaintainer()
	if err != nil {
		return warn, err
	}

	err = me.sanitizeUploaders()
	if err != nil {
		return warn, err
	}

	err = me.sanitizePriority()
	if err != nil {
		return warn, err
	}

	x, err = me.sanitizeArchitecture()
	if x != "" {
		warn += x + "\n"
	}

	if err != nil {
		return warn, err
	}

	err = me.sanitizePackages()
	if err != nil {
		return warn, err
	}

	err = me.sanitizeStandardsVersion()
	if err != nil {
		return warn, err
	}

	err = me.sanitizeVersion()
	if err != nil {
		return warn, err
	}

	err = me.sanitizeRulesRequiresRoot()
	if err != nil {
		return warn, err
	}

	err = me.sanitizeDescription()
	if err != nil {
		return warn, err
	}

	err = me.sanitizePackageType()
	if err != nil {
		return warn, err
	}

	err = me.sanitizeVCS()
	if err != nil {
		return warn, err
	}

	err = me.sanitizeTestsuite()
	if err != nil {
		return warn, err
	}

	return "", nil
}

func (me *Control) sanitizeRulesRequiresRoot() (err error) {
	switch me.RulesRequiresRoot {
	case "":
		me.RulesRequiresRoot = RULES_ROOT_BINARY_TARGETS
		return nil
	case RULES_ROOT_NO:
		return nil
	case RULES_ROOT_BINARY_TARGETS:
		return nil
	default:
	}

	list := strings.Split(me.RulesRequiresRoot, ",")
	for _, v := range list {
		if !strings.Contains(v, "/") {
			return fmt.Errorf("%s: %s",
				ERROR_CONTROL_RULES_ROOT_BAD,
				err,
			)
		}
	}

	return nil
}

func (me *Control) sanitizeTestsuite() (err error) {
	if me.Testsuite == nil {
		return nil
	}

	err = me.Testsuite.Sanitize()
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_CONTROL_TESTSUITE_BAD,
			err,
		)
	}

	return nil
}

func (me *Control) sanitizeVCS() (err error) {
	if me.VCS == nil {
		return nil
	}

	err = me.VCS.Sanitize()
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_CONTROL_VCS_BAD,
			err,
		)
	}

	return nil
}

func (me *Control) sanitizePackageType() (err error) {
	switch me.PackageType {
	case "":
	case PACKAGE_TYPE_DEB:
	case PACKAGE_TYPE_UDEB:
	default:
		return fmt.Errorf("%s: '%s'",
			ERROR_CONTROL_PACKAGE_TYPE_BAD,
			me.PackageType,
		)
	}

	return nil
}

func (me *Control) sanitizeDescription() (err error) {
	if me.Description == nil {
		return fmt.Errorf("%s: %s",
			ERROR_CONTROL_DESCRIPTION_BAD,
			"nil",
		)
	}

	err = me.Description.Sanitize()
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_CONTROL_DESCRIPTION_BAD,
			err,
		)
	}

	return nil
}

func (me *Control) sanitizeHomepage() {
	if me.Homepage == nil {
		me.Homepage = &url.URL{}
	}
}

func (me *Control) sanitizeVersion() (err error) {
	if me.Version == nil {
		return fmt.Errorf("%s: '%s'",
			ERROR_CONTROL_VERSION_BAD,
			"nil",
		)
	}

	err = me.Version.Sanitize()
	if err != nil {
		return fmt.Errorf("%s: '%s'",
			ERROR_CONTROL_VERSION_BAD,
			err,
		)
	}

	return nil
}

func (me *Control) sanitizeStandardsVersion() (err error) {
	switch me.StandardsVersion {
	case STANDARD_VER_4_6_0:
	case STANDARD_VER_4_5_1:
	case STANDARD_VER_4_5_0:
	case STANDARD_VER_4_4_1:
	case STANDARD_VER_4_3_0:
	case STANDARD_VER_4_2_1:
	default:
		return fmt.Errorf("%s: '%s'",
			ERROR_CONTROL_STANDARDS_VERSION_BAD,
			me.StandardsVersion,
		)
	}

	return nil
}

func (me *Control) sanitizePackages() (err error) {
	if me.Packages == nil {
		me.Packages = map[PackageListType]*PackageList{}
		return nil
	}

	// sanitize package list
	list := map[PackageListType]*PackageList{}
	for k, v := range me.Packages {
		if v == nil {
			continue
		}

		err = v.Sanitize()
		if err != nil {
			return fmt.Errorf("%s (%s): '%s'",
				ERROR_CONTROL_PACKAGE_LIST_BAD,
				v.Name,
				err,
			)
		}

		list[k] = v
	}
	me.Packages = list

	return nil
}

func (me *Control) sanitizeArchitecture() (warn string, err error) {
	switch me.Architecture {
	case "":
		return "", fmt.Errorf("%s: '%s'",
			ERROR_CONTROL_ARCHITECTURE_BAD,
			me.Architecture,
		)
	case "any":
	case "amd64":
	case "i386":
	case "armel":
	case "armhf":
	case "arm64":
	case "mips64el":
	case "mipsel":
	case "ppc64el":
	case "s390x":
	case "hurd-i386":
	case "kfreebsd-amd64":
	case "kfreebsd-i386":
	case "m68k":
	case "riscv64":
	case "sparc64":
	case "sh4":
	case "x32":
	default:
		return fmt.Sprintf("WARNING: unrecognized CPU: '%s'",
			me.Architecture,
		), nil
	}

	return "", nil
}

func (me *Control) sanitizePriority() (err error) {
	switch me.Priority {
	case PRIORITY_REQUIRED:
	case PRIORITY_IMPORTANT:
	case PRIORITY_STANDARD:
	case PRIORITY_OPTIONAL:
	case "":
	default:
		return fmt.Errorf("%s: '%s'",
			ERROR_CONTROL_PRIORITY_BAD,
			me.Priority,
		)
	}

	return nil
}

func (me *Control) sanitizeUploaders() (err error) {
	if me.Uploaders == nil {
		me.Uploaders = []*Entity{}
	}

	if len(me.Uploaders) == 0 {
		return nil
	}

	mList := map[string]bool{me.Maintainer.Name: true}
	list := []*Entity{}
	for _, v := range me.Uploaders {
		if v == nil {
			continue
		}

		err = v.Sanitize()
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_CONTROL_UPLOADERS_BAD,
				err,
			)
		}

		if !mList[v.Name] {
			mList[v.Name] = true
			list = append(list, v)
		}
	}
	me.Uploaders = list

	return nil
}

func (me *Control) sanitizeMaintainer() (err error) {
	if me.Maintainer == nil {
		return fmt.Errorf("%s: %s", ERROR_CONTROL_MAINTAINER_BAD, "nil")
	}

	err = me.Maintainer.Sanitize()
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_CONTROL_MAINTAINER_BAD, err)
	}

	return nil
}

// String generates the DEBIAN/control file data.
func (me *Control) String() (s string) {
	_, err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	// generate control file
	s += me.stringSourceParagraph()
	s += "\n"
	s += me.stringPackageParagraph()

	// trim tailing whitespaces
	s = strings.TrimRight(s, "\n")
	s = strings.TrimRight(s, "\r")
	s = strings.TrimRight(s, " ")

	return s
}

func (me *Control) stringSourceParagraph() (s string) {
	s += _FIELD_SOURCE + me.Name + "\n"
	s += _FIELD_MAINTAINER + me.Maintainer.Tag() + "\n"
	s += me.stringUploaders()

	if me.Section != "" {
		s += _FIELD_SECTION + me.Section + "\n"
	}

	if me.Priority != "" {
		s += _FIELD_PRIORITY + string(me.Priority) + "\n"
	}

	s += me.stringPackageList(PKG_LIST_BUILD_DEPENDS)
	s += me.stringPackageList(PKG_LIST_BUILD_DEPENDS_INDEP)
	s += me.stringPackageList(PKG_LIST_BUILD_CONFLICTS)
	s += me.stringPackageList(PKG_LIST_BUILD_CONFLICTS_INDEP)
	s += me.stringPackageList(PKG_LIST_BUILD_CONFLICTS_ARCH)

	s += _FIELD_STANDARDS_VERSION + string(me.StandardsVersion) + "\n"

	if me.Homepage != nil {
		s += _FIELD_HOMEPAGE + me.Homepage.String() + "\n"
	}

	if me.VCS != nil {
		s += me.VCS.String() + "\n"
	}

	if me.Testsuite != nil {
		s += _FIELD_TESTSUITE + me.Testsuite.String() + "\n"
	}

	if me.RulesRequiresRoot != "" {
		s += _FIELD_RULES_REQUIRES_ROOT + me.RulesRequiresRoot + "\n"
	}

	return s
}

func (me *Control) stringPackageParagraph() (s string) {
	s += _FIELD_PACKAGE + me.Name + "\n"

	if me.BuildSource {
		s += _FIELD_ARCHITECTURE + me.Architecture + "\n"
	}

	if !me.BuildSource {
		s += _FIELD_VERSION + me.Version.String() + "\n"
	}

	if me.Section != "" {
		s += _FIELD_SECTION + me.Section + "\n"
	}

	if me.Priority != "" {
		s += _FIELD_PRIORITY + string(me.Priority) + "\n"
	}

	if !me.BuildSource {
		s += _FIELD_ARCHITECTURE + me.Architecture + "\n"
	}

	if me.Essential {
		s += _FIELD_ESSENTIAL + "yes \n"
	}

	s += me.stringPackageList(PKG_LIST_DEPENDS)
	s += me.stringPackageList(PKG_LIST_PRE_DEPENDS)
	s += me.stringPackageList(PKG_LIST_RECOMMENDS)
	s += me.stringPackageList(PKG_LIST_SUGGESTS)
	s += me.stringPackageList(PKG_LIST_ENHANCES)
	s += me.stringPackageList(PKG_LIST_BREAKS)
	s += me.stringPackageList(PKG_LIST_CONFLICTS)
	s += me.stringPackageList(PKG_LIST_REPLACES)
	s += me.stringPackageList(PKG_LIST_PROVIDES)

	if !me.BuildSource {
		s += _FIELD_MAINTAINER + me.Maintainer.Tag() + "\n"
	}

	s += me.Description.String() + "\n"

	if me.Homepage != nil {
		s += _FIELD_HOMEPAGE + me.Homepage.String() + "\n"
	}

	s += me.stringPackageList(PKG_LIST_BUILT_USING)

	if me.BuildSource {
		switch me.PackageType {
		case PACKAGE_TYPE_UDEB:
			s += _FIELD_PACKAGE_TYPE + string(me.PackageType) + "\n"
		case PACKAGE_TYPE_DEB:
			// linter: explicit-default-in-package-type
		default:
		}
	}

	return s
}

func (me *Control) stringPackageList(id PackageListType) string {
	var ok bool
	var list *PackageList
	var x string

	// try to obtain the list
	list, ok = me.Packages[id]
	if !ok || list == nil {
		return ""
	}

	// got the list, process the string
	x = list.String()
	if x != "" {
		return x + "\n"
	}

	return ""
}

func (me *Control) stringUploaders() (s string) {
	x := ""

	for i, v := range me.Uploaders {
		if i > 0 {
			x += ",\n" + _FIELD_UPLOADERS_FOLD
		}

		x += v.Tag()
	}

	if x != "" {
		s += _FIELD_UPLOADERS + x + "\n"
	}

	return s
}
