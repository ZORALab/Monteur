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

// Error messages are the package standardized messages
const (
	ERROR_CHECKSUM               = "error when checksum"
	ERROR_CHECKSUM_BAD_FILE      = "failed to open file for checksum"
	ERROR_CHECKSUM_DELETE_FAILED = "failed to remove bad checksum file"
	ERROR_CHECKSUM_MISMATCHED    = "checksum mismatched"
	ERROR_FILE_EXISTS            = "destination file already exists"
	ERROR_FILE_OVERWRITE_FAILED  = "failed to overwrite destination file"
	ERROR_FILE_STAT              = "failed to obtain file stat locally"
	ERROR_FILE_RENAMED_FAILED    = "failed to rename downloaded file"
	ERROR_FILENAME_MISSING       = "failed to obtain filename remotely"
	ERROR_FILESIZE_MISSING       = "failed to obtain filesize remotely"
	ERROR_HASHER_UNHEALTHY       = "given checksum hasher is not healthy"
	ERROR_METHOD_MISSING         = "request method is missing"
	ERROR_PATH_INVALID           = "given Destination pathing is invalid"
	ERROR_PATH_MISSING           = "given Destination pathing is missing"
	ERROR_REQUEST_FAILED         = "failed to perform request remotely"
	ERROR_REQUEST_INIT_FAILED    = "failed to initialize request"
	ERROR_RESPONSE_BAD           = "bad response"
)
