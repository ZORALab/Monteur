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

//nolint:lll
const (
	ERROR_CHANGELOG_BAD                  = "bad debian/changelog file"
	ERROR_CHANGELOG_ENTRY_BAD            = "bad debian/changelog entry"
	ERROR_CHANGELOG_FILE_READ            = "error while reading changelog path"
	ERROR_CHANGELOG_FILEPATH_BAD         = "bad debian/changelog filepath"
	ERROR_CHANGELOG_NOT_A_FILE           = "given debian/changelog path is not a file"
	ERROR_CHANGELOG_PARSE_BAD            = "bad debian/changelog path parsing"
	ERROR_CHANGELOG_PARSE_COMPLETED      = "parse changelog entry work was completed"
	ERROR_CHANGELOG_PARSE_DISTRO_FAILED  = "failed to parse changelog entry's Distribution"
	ERROR_CHANGELOG_PARSE_PACKAGE_FAILED = "failed to parse changelog entry's Package"
	ERROR_CHANGELOG_PARSE_URGENCY_FAILED = "failed to parse changelog entry's Urgency"
	ERROR_CHANGELOG_PARSE_VERSION_FAILED = "failed to parse changelog entry's Version"
	ERROR_CHANGELOG_PATH_BAD             = "bad changelog path for append"
	ERROR_CHANGELOG_ENTRY_TOO_LONG       = "changelog entry is longer than 80 characters"

	ERROR_CHECKLIST_4_6_0_LIB64 = "Error Checklist 4.6.0: no package allowed to install files into /usr/lib64/"
	ERROR_CHECKLIST_4_0_0_PATHS = "Error Checklist 4.0.0: same programs cannot be installed into both /path and /usr/path"
	ERROR_CHECKLIST_3_7_0_XORG  = "Error Checklist 3.7.0: /usr/X11R6 is gone forever."

	ERROR_COMPAT_UNSET = "DEBIAN/compat cannot be left unset"

	ERROR_CONTROL_ARCHITECTURE_BAD      = "bad debian/control Architecture"
	ERROR_CONTROL_BAD                   = "bad debian/control file"
	ERROR_CONTROL_CHANGEDBY_BAD         = "bad debian/control Changed-By"
	ERROR_CONTROL_DESCRIPTION_BAD       = "bad debian/control Description"
	ERROR_CONTROL_MAINTAINER_BAD        = "bad debian/control Maintainer"
	ERROR_CONTROL_NAME_BAD              = "bad debian/control Name"
	ERROR_CONTROL_PACKAGE_LIST_BAD      = "bad debian/control package list"
	ERROR_CONTROL_PACKAGE_TYPE_BAD      = "bad debian/control Package Type"
	ERROR_CONTROL_PRIORITY_BAD          = "bad debian/control Priority"
	ERROR_CONTROL_RULES_ROOT_BAD        = "bad debian/control Rules-Requires-Root"
	ERROR_CONTROL_SECTION_BAD           = "bad debian/control Section"
	ERROR_CONTROL_STANDARDS_VERSION_BAD = "bad debian/control Standard-Version"
	ERROR_CONTROL_UPLOADERS_BAD         = "bad debian/control Uploaders"
	ERROR_CONTROL_VCS_BAD               = "bad debian/control VCS-XXX"
	ERROR_CONTROL_VERSION_BAD           = "bad debian/control Version"
	ERROR_CONTROL_TESTSUITE_BAD         = "bad debian/control Testsuite"

	ERROR_COPYRIGHT_BAD                  = "bad debian/copyright file"
	ERROR_COPYRIGHT_FORMAT_BAD           = "bad debian/copyright Format"
	ERROR_COPYRIGHT_LICENSE_MISSING      = "missing debian/copyright License"
	ERROR_COPYRIGHT_UPSTREAM_CONTACT_BAD = "bad debian/copyright Upstream-Contact"
	ERROR_COPYRIGHT_UPSTREAM_NAME_BAD    = "bad debian/copyright Upstream-Name"

	ERROR_DEP_ADD_FAILED = "failed to add package"
	ERROR_DEP_BAD        = "bad dependency status"
	ERROR_DEP_NAME_BAD   = "bad dependency name"

	ERROR_DIR_MISSING           = "directory is missing"
	ERROR_DIR_CREATE_FAILED     = "failed to create directory"
	ERROR_FILE_CHMOD_FAILED     = "failed to change file permission"
	ERROR_FILE_OPEN_FAILED      = "failed to open file"
	ERROR_FILE_WRITE_FAILED     = "failed to write file"
	ERROR_FILE_OVERWRITE_FAILED = "failed to overwrite file"

	ERROR_PACKAGE_LIST_BAD      = "bad package list"
	ERROR_PACKAGE_LIST_NAME_BAD = "bad package list name"

	ERROR_PACKAGE_NAME_MISSING         = "missing dep-package Name"
	ERROR_PACKAGE_PARSE_ARCH_BAD       = "incomprehensible deb-package Architecture"
	ERROR_PACKAGE_PARSE_NAME_BAD       = "incomprehensible deb-package Name"
	ERROR_PACKAGE_PARSE_VERCONTROL_BAD = "incomprehensible deb-package VERControl"
	ERROR_PACKAGE_VER_BAD              = "bad dep-package VERControl"
	ERROR_PACKAGE_VER_MISSING          = "missing dep-package VERControl"
	ERROR_PACKAGE_VER_CONTROL_UNKNOWN  = "unknown dep-package VERControl"

	ERROR_PROCESSOR_DESTDIR_BAD        = "bad DestDir"
	ERROR_PROCESSOR_DESTDIR_MISSING    = "missing DestDir"
	ERROR_PROCESSOR_WORKINGDIR_BAD     = "bad WorkingDir"
	ERROR_PROCESSOR_WORKINGDIR_MISSING = "missing WorkingDir"
	ERROR_PROCESSOR_DEBIAN_MISSING     = "missing debian or DEBIAN directory"

	ERROR_ENTITY_NAME_BAD  = "entity has bad name"
	ERROR_ENTITY_EMAIL_BAD = "entity has bad email"

	ERROR_EXTRACT_UNSUPPORTED = "Extract(...) is unsupported"

	ERROR_INSTALL_BAD = "bad DEBIAN/install data"

	ERROR_LICENSE_COPYRIGHT_EMPTY    = "copyright cannot be empty for file license"
	ERROR_LICENSE_COPYRIGHT_YEAR_BAD = "Copyright.Year must be > 0"
	ERROR_LICENSE_EMPTY              = "license cannot be empty for file license"
	ERROR_LICENSE_FILE_EMPTY         = "files cannot be empty for file license"
	ERROR_LICENSE_MULTI              = "only 1 license for file license"

	ERROR_MANPAGE_BAD = "bad manpages"

	ERROR_RULES_EMPTY = "DEBIAN/rules cannot be left empty"

	ERROR_SHELL_SCRIPT_TYPE_UNKNOWN = "unknown debian/script type"

	ERROR_SOURCE_BAD            = "bad debian/source data"
	ERROR_SOURCE_FORMAT_UNKNOWN = "unknown debian/source/format"

	ERROR_TESTSUITE_BAD = "bad testsuite path"

	ERROR_VCS_TYPE_BAD    = "bad VCS type"
	ERROR_VCS_URL_BAD     = "bad VCS URL"
	ERROR_VCS_BROWSER_BAD = "bad VCS browser URL"

	ERROR_VERSION_PARSE_EPOCH_FAILED          = "incomprehensible epoch parsing"
	ERROR_VERSION_PARSE_EPOCH_UPSTREAM_FAILED = "incomprehensible epoch+upstream parsing"
	ERROR_VERSION_REVISION_DIGIT_FIRST        = "revision must start with digit"
	ERROR_VERSION_REVISION_ILLEGAL            = "revision has illegal char"
	ERROR_VERSION_UPSTREAM_EMPTY              = "upstream is empty"
	ERROR_VERSION_UPSTREAM_DIGIT_FIRST        = "upstream must start with digit"
	ERROR_VERSION_UPSTREAM_ILLEGAL            = "upstream has illegal char"
	ERROR_VERSION_UPSTREAM_NO_DASH            = "upstream cannot have dash"
)
