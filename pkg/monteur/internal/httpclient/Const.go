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

package httpclient

//nolint:stylecheck,revive
const (
	// DOWNLOAD_EXTENSION is the common extension indicating download status
	DOWNLOAD_EXTENSION = ".download"

	// TIMEOUT is the default timing for timeout a download in seconds.
	TIMEOUT = 60

	// FILE_PERMISSION is the default access permission for downloaded file.
	FILE_PERMISSION = 0600

	// DIR_PERMISSION is the default access permission for directory.
	DIR_PERMISSION = 0700
)
