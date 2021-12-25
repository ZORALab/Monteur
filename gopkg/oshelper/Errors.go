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

package oshelper

const (
	ERROR_DEST_EMPTY       = "dest is empty"
	ERROR_DIRECTORY_CREATE = "error creating directory"
	ERROR_DIRECTORY        = "error encountered while walking directory"
	ERROR_CHOWN            = "error changing ownership"
	ERROR_CHTIMES          = "error changing timestamps"
	ERROR_FILE_COPY        = "error copying file"
	ERROR_FILE_OPEN        = "error opening file"
	ERROR_FILE_PERM        = "error setting file permission"
	ERROR_FILE_SYNC        = "error syncing file"
	ERROR_FILE_STAT        = "error obtaining file stat"
	ERROR_SOURCE_EMPTY     = "source is empty"
	ERROR_PIPE_CREATE      = "error creating pipe"
	ERROR_PIPE_PERM        = "error setting pipe permission"
	ERROR_PIPE_UNSUPPORTED = "pipe is unsupported"
	ERROR_SOURCE_UNKNOWN   = "unknown source type"
	ERROR_SYMLINK_READ     = "error reading symlink"
	ERROR_SYMLINK_CREATE   = "error creating symlink"
)
