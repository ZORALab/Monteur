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

// PackageListType are the package list relationship for DEBIAN/control
type PackageListType string

const (
	PKG_LIST_BUILD_DEPENDS         PackageListType = "Build-Depends"
	PKG_LIST_BUILD_DEPENDS_INDEP   PackageListType = "Build-Depends-Indep"
	PKG_LIST_BUILD_DEPENDS_ARCH    PackageListType = "Build-Depends-Arch"
	PKG_LIST_BUILD_CONFLICTS       PackageListType = "Build-Conflicts"
	PKG_LIST_BUILD_CONFLICTS_INDEP PackageListType = "Build-Conflicts-Indep"
	PKG_LIST_BUILD_CONFLICTS_ARCH  PackageListType = "Build-Conflicts-Arch"
	PKG_LIST_PRE_DEPENDS           PackageListType = "Pre-Depends"
	PKG_LIST_DEPENDS               PackageListType = "Depends"
	PKG_LIST_RECOMMENDS            PackageListType = "Recommends"
	PKG_LIST_SUGGESTS              PackageListType = "Suggests"
	PKG_LIST_ENHANCES              PackageListType = "Enhances"
	PKG_LIST_BREAKS                PackageListType = "Breaks"
	PKG_LIST_CONFLICTS             PackageListType = "Conflicts"
	PKG_LIST_PROVIDES              PackageListType = "Provides"
	PKG_LIST_REPLACES              PackageListType = "Replaces"
	PKG_LIST_BUILT_USING           PackageListType = "Built-Using"
)

// PackageList is the list of related packages.
//
// More info:
//    https://www.debian.org/doc/debian-policy/ch-relationships.html
type PackageList struct {
	// List are the list of PackageMeta data in the list.
	//
	// This field is **MANDATORY**.
	//
	// Map data type is used over Array data type mainly for removing
	// duplications.
	List map[string]*PackageMeta

	// Name is the package list relationship in DEBIAN/control.
	//
	// This field is **MANDATORY**.
	Name PackageListType
}

// Parse is to parse a list of packages into its List.
//
// It returns error should any list becomes incomprehensible.
func (me *PackageList) Parse(name string, packages []string) (err error) {
	var pkg *PackageMeta

	me.Name = PackageListType(name)
	me.List = map[string]*PackageMeta{}

	for _, meta := range packages {
		pkg = &PackageMeta{}

		err = pkg.Parse(meta)
		if err != nil {
			return err
		}

		me.List[pkg.Name] = pkg
	}

	return me.Sanitize()
}

// Sanitize is to check the PackageList is ready for use.
func (me *PackageList) Sanitize() (err error) {
	switch me.Name {
	case PKG_LIST_BUILD_DEPENDS:
	case PKG_LIST_BUILD_DEPENDS_INDEP:
	case PKG_LIST_BUILD_DEPENDS_ARCH:
	case PKG_LIST_BUILD_CONFLICTS:
	case PKG_LIST_BUILD_CONFLICTS_INDEP:
	case PKG_LIST_BUILD_CONFLICTS_ARCH:
	case PKG_LIST_PRE_DEPENDS:
	case PKG_LIST_DEPENDS:
	case PKG_LIST_RECOMMENDS:
	case PKG_LIST_SUGGESTS:
	case PKG_LIST_ENHANCES:
	case PKG_LIST_BREAKS:
	case PKG_LIST_CONFLICTS:
	case PKG_LIST_PROVIDES:
	case PKG_LIST_REPLACES:
	case PKG_LIST_BUILT_USING:
	default:
		return fmt.Errorf("%s: '%s'",
			ERROR_PACKAGE_LIST_NAME_BAD,
			me.Name,
		)
	}

	if me.List == nil {
		me.List = map[string]*PackageMeta{}
	}

	list := map[string]*PackageMeta{}
	for _, v := range me.List {
		list[v.Name] = v

		err = v.Sanitize()
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_PACKAGE_LIST_BAD,
				err,
			)
		}
	}
	me.List = list

	return nil
}

// String will generate a full package list with its relation tag.
//
// The format is solely for DEBIAN/control file creation.
//
// An example of generated output for `PKG_LIST_BUILD_DEPENDS`:
//
//   Build-Depends:
//    libluajit5.1-dev [i386 amd64 kfreebsd-i386 armel armhf powerpc mips],
//    liblua5.1-dev [hurd-i386 ia64 kfreebsd-amd64 s390x sparc],
func (me *PackageList) String() (s string) {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	if len(me.List) == 0 {
		return ""
	}

	// add name
	s += string(me.Name) + ":\n"

	// add title
	first := true
	for _, v := range me.List {
		if !first {
			s += ",\n"
		}

		s += " " + v.String()
		first = false
	}

	return s
}
