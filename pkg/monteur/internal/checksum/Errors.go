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

package checksum

// Error messages are the package standardized messages
//nolint:stylecheck,revive
const (
	ERROR_ALGO_BAD      = "bad hashing algorithm"
	ERROR_INIT_FAILED   = "failed to initialize checksum"
	ERROR_MISMATCHED    = "checksum mismatched"
	ERROR_MISSING_VALUE = "missing checksum value"
	ERROR_PARSE_BAD     = "bad value for parser"
	ERROR_READ_FILE     = "failed to read file"
)
