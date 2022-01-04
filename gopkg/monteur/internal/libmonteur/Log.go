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

package libmonteur

// Supported variables keys in key:value variables placholders.
//
// It is used in every toml config file inside setup/program/ config directory
// for placeholding variable elements in the fields' value.
const (
	LOG_FORMAT_OUTPUT = "Got:\n╔═══ BEGIN ═══╗\n%v╚═══  END  ═══╝"

	LOG_JOB_INIT_SUCCESS = "Task initialized successfully. Standing By..."
	LOG_JOB_START        = "Run Task Now: " + LOG_OK

	LOG_SUCCESS = "➤ SUCCESS"
	LOG_OK      = "➤ OK"
	LOG_DONE    = "➤ DONE"
)
