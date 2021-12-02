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

//nolint:lll
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
	ERROR_COMPOSE = "[ ERROR - Compose ]"
	ERROR_PUBLISH = "[ ERROR - Publish ]"
)

const (
	ERROR_CHECKSUM_ALGO_UNKNOWN   = "unsupported checksum value"
	ERROR_CHECKSUM_BAD            = "bad checksum value"
	ERROR_CHECKSUM_FORMAT_BAD     = "bad checksum format"
	ERROR_CHECKSUM_FORMAT_UNKNOWN = "unsupported checksum format"
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
	ERROR_LOG_PATH_EMPTY = "given path is empty"
	ERROR_LOG_PREPARE    = "failed to open and prepare log file"
	ERROR_LOG_UNHEALTHY  = "logger is unhealthy"
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
