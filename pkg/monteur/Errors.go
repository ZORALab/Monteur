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

package monteur

//nolint:stylecheck,revive,lll
const (
	ERROR_BAD_REPO             = "current path is not a git-repo + monteur"
	ERROR_CURRENT_DIRECTORY    = "failed to get current directory pathing"
	ERROR_MISSING_DIR          = "missing directory"
	ERROR_FAILED_CONFIG_DECODE = "failed to decode file"

	ERROR_SETUP_TAG = "setup(): "

	ERROR_SETUP_CHECKSUM_UNSUPPORTED_ALGO   = "unsupported checksum algo"
	ERROR_SETUP_CHECKSUM_UNSUPPORTED_FORMAT = "unsupported checksum format"
	ERROR_SETUP_CHECKSUM_BAD                = "bad/empty checksum value"
	ERROR_SETUP_CHECKSUM_BAD_ALGO           = "bad/empty checksum algo"
	ERROR_SETUP_CHECKSUM_BAD_FORMAT         = "bad/empty checksum format"

	ERROR_SETUP_CONFIG_MISSING                = "missing config data"
	ERROR_SETUP_CONFIG_FAILED                 = "failed to create config"
	ERROR_SETUP_DIR_CREATE_FAILED             = "create directory failed"
	ERROR_SETUP_HTTPS_DOWNLOAD_FAILED         = "https download failed"
	ERROR_SETUP_INSTALL_MOVE_FAILED           = "failed to move"
	ERROR_SETUP_INSTALL_SCRIPT_FAILED         = "failed to script"
	ERROR_SETUP_INSTRUCTION_SOURCE_BAD        = "bad instruction source"
	ERROR_SETUP_INSTRUCTION_TARGET_BAD        = "bad instruction target"
	ERROR_SETUP_INSTRUCTION_CONDITION_BAD     = "bad instruction condition"
	ERROR_SETUP_INSTRUCTION_TYPE_BAD          = "bad instruction type"
	ERROR_SETUP_INSTRUCTION_TYPE_UNKNOWN      = "unknown instruction type"
	ERROR_SETUP_METADATA_NAME_BAD             = "bad name"
	ERROR_SETUP_METADATA_DESC_BAD             = "bad description"
	ERROR_SETUP_POSTCONFIG_BAD                = "bad post config"
	ERROR_SETUP_PROGRAM_NOT_SANITIZED         = "program not sanitized"
	ERROR_SETUP_PROGRAM_NOT_SUPPORTED         = "program not supported"
	ERROR_SETUP_SOURCE_ARCHIVE_BAD            = "bad archive name"
	ERROR_SETUP_SOURCE_ARCHIVE_FORMAT_BAD     = "bad archive format"
	ERROR_SETUP_SOURCE_ARCHIVE_FORMAT_UNKNOWN = "unsupported archive format"
	ERROR_SETUP_SOURCE_HEADER_BAD             = "bad header value"
	ERROR_SETUP_SOURCE_METHOD_BAD             = "bad method for url"
	ERROR_SETUP_SOURCE_BAD                    = "bad url address"
	ERROR_SETUP_TYPE_UNKNOWN                  = "unknown source type"
	ERROR_SETUP_TYPE_BAD                      = "bad source type"
)
