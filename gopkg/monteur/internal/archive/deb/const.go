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

const (
	_FIELD_ARCHITECTURE        = "Architecture: "
	_FIELD_COMMENT             = "Comment: "
	_FIELD_COPYRIGHT           = "Copyright: "
	_FIELD_COPYRIGHT_FOLD      = "           "
	_FIELD_DESCRIPTION         = "Description: "
	_FIELD_DISCLAIMER          = "Disclaimer: "
	_FIELD_ESSENTIAL           = "Essential: "
	_FIELD_FILES               = "Files: "
	_FIELD_FORMAT              = "Format: "
	_FIELD_HOMEPAGE            = "Homepage: "
	_FIELD_LICENSE             = "License: "
	_FIELD_MAINTAINER          = "Maintainer: "
	_FIELD_PACKAGE             = "Package: "
	_FIELD_PRIORITY            = "Priority: "
	_FIELD_RULES_REQUIRES_ROOT = "Rules-Requires-Root: "
	_FIELD_SECTION             = "Section: "
	_FIELD_SOURCE              = "Source: "
	_FIELD_STANDARDS_VERSION   = "Standards-Version: "
	_FIELD_TESTSUITE           = "Testsuite: "
	_FIELD_UPLOADERS           = "Uploaders: "
	_FIELD_UPLOADERS_FOLD      = "           "
	_FIELD_UPSTREAM_NAME       = "Upstream-Name: "
	_FIELD_UPSTREAM_CONTACT    = "Upstream-Contact: "
	_FIELD_VCS_BROWSER         = "Vcs-Browser: "
	_FIELD_VERSION             = "Version: "
	_FIELD_PACKAGE_TYPE        = "Package-Type: "
)

const (
	_FIELD_CHANGELOG_DELIMIT_PACKAGE   = " ("
	_FIELD_CHANGELOG_DELIMIT_DISTRO    = ") "
	_FIELD_CHANGELOG_DELIMIT_URGENCY   = "; urgency="
	_FIELD_CHANGELOG_DELIMIT_CHANGE    = "  * "
	_FIELD_CHANGELOG_DELIMIT_SIGNATURE = " -- "
	_FIELD_CHANGELOG_DELIMIT_EMAIL     = " <"
	_FIELD_CHANGELOG_DELIMIT_TIMESTAMP = "  "
)

const (
	_PERMISSION_DIR  = 0755
	_PERMISSION_FILE = 0644
	_PERMISSION_EXEC = 0755
)
