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

package archive

const (
	ERROR_ARCHIVE             = "error in archive pathing"
	ERROR_ARCHIVE_ABS_FAILED  = "failed to absolute archive pathing"
	ERROR_ARCHIVE_CREATE      = "error creating archive file"
	ERROR_ARCHIVE_DIR_INVALID = "archive housing pathing is not a directory"
	ERROR_ARCHIVE_DIR_MISSING = "archive housing directory is missing"
	ERROR_ARCHIVE_EMPTY       = "given archive pathing is empty"
	ERROR_ARCHIVE_EXT_MISSING = "archive pathing has missing file extension"
	ERROR_ARCHIVE_NOT_DIR     = "archive pathing is not a directory"
	ERROR_ARCHIVE_READ        = "error reading archive file"

	ERROR_COMPRESSION_INVALID = "invalid compression level"

	ERROR_DIR         = "error with directory"
	ERROR_DIR_INVALID = "path is not a directory"

	ERROR_EXTRACT_COPY = "error while extract-copying"

	ERROR_FILE_ATIME_INVALID  = "invalid file access time"
	ERROR_FILE_CHMOD_FAILED   = "error while chmod file"
	ERROR_FILE_CHTIMES_FAILED = "error while chtimes file"
	ERROR_FILE_INFO_READ      = "error while reading file info"
	ERROR_FILE_MODE_INVALID   = "invalid file mode"
	ERROR_FILE_MTIME_INVALID  = "invalid file modified time"
	ERROR_FILE_READ           = "error reading file"
	ERROR_FILE_WRITE          = "error writing file"
	ERROR_FILE_UNSUPPORTED    = "unsupported file type like socket, pipe..."

	ERROR_HEADER_READ  = "error reading file header"
	ERROR_HEADER_WRITE = "error writing file header"

	ERROR_MKDIR_FAILED = "failed to make directory pathing"

	ERROR_OVERWRITE_FAILED = "error while overwriting"

	ERROR_PATH_BASE_EMPTY   = "base path is empty"
	ERROR_PATH_EMPTY        = "path is empty"
	ERROR_PATH_OUT_OF_BOUND = "given compressing path is outside of raw pathing"
	ERROR_PATH_ABS_FAILED   = "failed to absolute compressing path"
	ERROR_PATH_REL_FAILED   = "failed to relative compressing path"

	ERROR_SYMLINK_CREATE            = "error when creating symlink"
	ERROR_SYMLINK_EVAL              = "error evaluating symlink"
	ERROR_SYMLINK_EVAL_FAILED       = "failed to evaluate symlink's target"
	ERROR_SYMLINK_PATH_EMPTY        = "given symlink pathing is empty"
	ERROR_SYMLINK_READ              = "error reading symlink's target"
	ERROR_SYMLINK_TARGET_ABS_FAILED = "failed to absolute symlink's target path"
	ERROR_SYMLINK_TARGET_EMPTY      = "given symlink target pathing is empty"
	ERROR_SYMLINK_UNRESOLVABLE      = "error resolving symlink"

	ERROR_RAW            = "error in raw pathing"
	ERROR_RAW_EMPTY      = "given raw pathing is empty"
	ERROR_RAW_ABS_FAILED = "failed to absolute raw pathing"
	ERROR_RAW_NOT_DIR    = "raw pathing is not a directory"
	ERROR_RAW_MISSING    = "raw directory is missing"

	ERROR_TARGET_EXISTS = "target already exists"
)
