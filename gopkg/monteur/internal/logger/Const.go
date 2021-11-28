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

package logger

// These are general constants used in this package
const (
	// ISO8601 is the timestamp format used to generate log entry.
	ISO8601 = "2006-01-02T15:04:05Z07:00"
)

// `StatusType`s are the output controller for `Printf(...)` and similar.
//
// These controls are mainly to tell logger to output the data to either
// `STDERR` or `STDOUT` channels.
//
// When in doubt, one should use and always stick to `TypeStatus`.
type StatusType uint

const (
	// TYPE_STATUS is to output for STDERR in terminal. (Default)
	TYPE_STATUS StatusType = iota

	// TypeOutput is to output for STDOUT in terminal.
	TYPE_OUTPUT
)

// `Level`s are the log severity level available to use.
const (
	TAG_ERROR   = "ERROR"
	TAG_WARNING = "WARNING"
	TAG_INFO    = "INFO"
	TAG_SUCCESS = "SUCCESS"
	TAG_DEBUG   = "DEBUG"
	TAG_NO      = ""
)

// `Permission`s are the list of filesystem permissions applicable to UNIX OS
const (
	PERMISSION_FILE_LOG = 0600
)
