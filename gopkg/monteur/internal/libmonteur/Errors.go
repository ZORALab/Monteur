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

const (
	ERROR_BAD_REPO          = "current path is not a git-repo + monteur"
	ERROR_DIR_BAD           = "failed to get directory path"
	ERROR_DIR_MISSING       = "missing directory"
	ERROR_DIR_CREATE_FAILED = "failed to create directory"
)

const (
	ERROR_FILE_DECODE_FAILED = "failed to decode file"
	ERROR_TOML_PARSE_FAILED  = "failed to parse toml file"
)

const (
	ERROR_SETUP   = "[ ERROR - Setup   ]"
	ERROR_TEST    = "[ ERROR - Test    ]"
	ERROR_BUILD   = "[ ERROR - Build   ]"
	ERROR_PACKAGE = "[ ERROR - Package ]"
	ERROR_COMPOSE = "[ ERROR - Compose ]"
	ERROR_PUBLISH = "[ ERROR - Publish ]"
	ERROR_RELEASE = "[ ERROR - Release ]"
	ERROR_CLEAN   = "[ ERROR - Clean   ]"
)

const (
	ERROR_APP_FMT_BAD   = "bad app data formatting"
	ERROR_APP_DATA      = "error processing app data"
	ERROR_APP_COPYRIGHT = "error processing app copyright data"
)

const (
	ERROR_CHANGELOG_ENTIRES_MISSING    = "missing current changelog entries"
	ERROR_CHANGELOG_LINE_BREAK_MISSING = "missing changelog entries' line break"
	ERROR_CHANGELOG_REGEX_BAD          = "bad regex for changelog entries"
)

const (
	ERROR_CHECKSUM_ALGO_UNKNOWN   = "unsupported checksum value"
	ERROR_CHECKSUM_BAD            = "bad checksum value"
	ERROR_CHECKSUM_FORMAT_BAD     = "bad checksum format"
	ERROR_CHECKSUM_FORMAT_UNKNOWN = "unsupported checksum format"
	ERROR_CHECKSUM_TYPE_BAD       = "bad checksum type"
)

const (
	ERROR_COMMAND_BAD                = "bad command"
	ERROR_COMMAND_DEPENDENCY_FMT_BAD = "bad command's dependency formatting"
	ERROR_COMMAND_FAILED             = "failed to execute command"
	ERROR_COMMAND_FMT_BAD            = "bad command formatting"
)

const (
	ERROR_DEPENDENCY_BAD = "bad dependency"
)

const (
	ERROR_VARIABLES_FMT_BAD = "bad variable formatting"
)

const (
	ERROR_LANGUAGE_CODE_MISSING = "missing language code"
	ERROR_LANGUAGE_NAME_MISSING = "missing language name"
)

const (
	ERROR_LOG_PATH_EMPTY = "given path is empty"
	ERROR_LOG_PREPARE    = "failed to open and prepare log file"
	ERROR_LOG_UNHEALTHY  = "logger is unhealthy"
)

const (
	ERROR_PACKAGER_APP_MISSING          = "missing app data for package"
	ERROR_PACKAGER_ARCH_MISSING         = "missing package architecture(s) value"
	ERROR_PACKAGER_DEB_BAD              = "bad deb data"
	ERROR_PACKAGER_FMT_BAD              = "bad variable formatting"
	ERROR_PACKAGER_FILE_MISSING         = "missing file for packaging"
	ERROR_PACKAGER_FILES_COPY_FAILED    = "failed to copy package's file"
	ERROR_PACKAGER_FILES_MISSING        = "missing package's Files"
	ERROR_PACKAGER_OS_MISSING           = "missing package os(es) value"
	ERROR_PACKAGER_PREPARE_FAILED       = "failed to prepare package"
	ERROR_PACKAGER_RELATIONSHIPS        = "error parsing related dep packages"
	ERROR_PACKAGER_TYPE_UNKNOWN         = "unknown package type (Metadata.Type)"
	ERROR_PACKAGER_VERSION_INCOMPATIBLE = "incompatible package version data"
)

const (
	ERROR_RELEASER_TARGET_MISSING   = "Releases.Target directory path is missing"
	ERROR_RELEASER_TYPE_UNSUPPORTED = "Metadata.Type is not supported"
)

const (
	ERROR_PROGRAM_ARCHIVE_BAD            = "bad archived program's name"
	ERROR_PROGRAM_ARCHIVE_FORMAT_BAD     = "bad archived program's format"
	ERROR_PROGRAM_ARCHIVE_FORMAT_UNKNOWN = "unsupported archived program's format"

	ERROR_PROGRAM_CONFIG_BAD    = "bad program's config data"
	ERROR_PROGRAM_CONFIG_FAILED = "failed to create program's config file"

	ERROR_PROGRAM_HTTPS_HEADER_BAD = "bad program's HTTPS source header"
	ERROR_PROGRAM_HTTPS_METHOD_BAD = "bad program's HTTPS method"

	ERROR_PROGRAM_INST_CONDITION_BAD  = "bad program's setup inst. condition"
	ERROR_PROGRAM_INST_FAILED         = "failed to setup program"
	ERROR_PROGRAM_INST_SOURCE_BAD     = "bad program's inst. source"
	ERROR_PROGRAM_INST_SOURCE_MISSING = "missing source file"
	ERROR_PROGRAM_INST_TARGET_BAD     = "bad program's inst. target"
	ERROR_PROGRAM_INST_TYPE_BAD       = "bad program's setup instruction"
	ERROR_PROGRAM_INST_TYPE_UNKNOWN   = "unknown program's setup instruction"

	ERROR_PROGRAM_META_NAME_BAD = "bad program name"
	ERROR_PROGRAM_META_DESC_BAD = "bad program description"

	ERROR_PROGRAM_TYPE_BAD     = "bad program sourcing type"
	ERROR_PROGRAM_TYPE_UNKNOWN = "unknown program sourcing type"

	ERROR_PROGRAM_MISSING     = "program is missing from local system"
	ERROR_PROGRAM_UNSUPPORTED = "unsupported program for this hardware"

	ERROR_PROGRAM_URL_BAD = "bad program's source URL"
)

const (
	ERROR_PUBLISH_METADATA_MISSING = "missing metadata"
)
