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

package targz

//nolint:stylecheck,revive
const (
	ERROR_ARCHIVE_EXTRACT_FAILED      = "failed to extract archive file"
	ERROR_ARCHIVE_FILE_HEADER_FAILED  = "failed to read header from archive"
	ERROR_ARCHIVE_READ_FAILED         = "failed to read archive file"
	ERROR_ARCHIVE_TAINTED_HEADER_PATH = "tainted content header pathing"
	ERROR_PATH_ARCHIVE                = "error in archive metadata"
	ERROR_PATH_ABS_FAILED             = "failed to absolute pathing"
	ERROR_PATH_DIR_MISSING            = "path directory does not exist"
	ERROR_PATH_EMPTY                  = "path is empty"
	ERROR_PATH_IS_DIR                 = "path is a directory"
	ERROR_PATH_NOT_DIR                = "path is not directory"
	ERROR_EXTENSION_MISSING           = "missing .tar.gz extension"
	ERROR_FILE_READ_FAILED            = "failed to read file"
	ERROR_FILE_WRITE_FAILED           = "failed to write file"
	ERROR_FILE_HEADER_READ_FAILED     = "failed to read file header"
	ERROR_FILE_HEADER_WRITE_FAILED    = "failed to write file header"
	ERROR_DEST_EXISTS                 = "destination exists"
	ERROR_DEST_OVERWRITE_FAILED       = "failed to overwrite"
	ERROR_DEST_CREATE_FAILED          = "failed to create destination"
)
