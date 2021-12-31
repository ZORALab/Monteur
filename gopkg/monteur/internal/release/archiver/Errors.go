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

package archiver

const (
	ERROR_CHECKSUM_UNSUPPORTED = "unsupported checksum type"
	ERROR_CHECKSUM_EMPTY       = "checksum list is empty"

	ERROR_DATAPATH_INVALID = "given DataPath is invalid"
	ERROR_DATAPATH_CREATE  = "error creating DataPath"

	ERROR_FORMAT_UNSUPPORTED = "given Format is unsupported"

	ERROR_PATH         = "error with Path"
	ERROR_PATH_CREATE  = "error creating path"
	ERROR_PATH_INVALID = "given Path is not a directory"
	ERROR_PATH_MISSING = "given Path is missing"

	ERROR_TARGET_CHECKSUM        = "error checksumming target"
	ERROR_TARGET_DATA_DIR_CREATE = "error creating target data path's directory"
	ERROR_TARGET_DIR_CREATE      = "error creating target path's directory"
	ERROR_TARGET_INVALID         = "package file is invalid"
	ERROR_TARGET_MISSING         = "package file is missing"
	ERROR_TARGET_COPY            = "error copying package file"

	ERROR_VERSION_MISSING = "given Version is missing"
)
