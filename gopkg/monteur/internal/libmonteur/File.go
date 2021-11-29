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

// Critical Object Names are the file or directory names critical for Monteur.
//
// These critical object names are mainly to locate root repository with Monteur
// supports.
const (
	EXTENSION_LOG   = ".log"
	EXTENSION_TOML  = ".toml"
	EXTENSION_TARGZ = ".tar.gz"
	EXTENSION_ZIP   = ".zip"
)

const (
	DIRECTORY_APP              = "app"
	DIRECTORY_GIT              = ".git"
	DIRECTORY_MONTEUR_CONFIG_D = "config.d"
	DIRECTORY_MONTEUR_CONFIG   = ".configs/monteur"

	DIRECTORY_PUBLISH   = "publish"
	DIRECTORY_PUBLISHER = DIRECTORY_PUBLISH + "/publishers"

	DIRECTORY_COMPOSE  = "compose"
	DIRECTORY_COMPOSER = DIRECTORY_COMPOSE + "/composers"

	DIRECTORY_SETUP          = "setup"
	DIRECTORY_SETUP_PROGRAMS = DIRECTORY_SETUP + "/programs"
)

const (
	FILE_TOML_COMPOSE   = DIRECTORY_COMPOSE + "/config" + EXTENSION_TOML
	FILE_TOML_PUBLISH   = DIRECTORY_PUBLISH + "/config" + EXTENSION_TOML
	FILE_TOML_SETUP     = DIRECTORY_SETUP + "/config" + EXTENSION_TOML
	FILE_TOML_WORKSPACE = "workspace" + EXTENSION_TOML

	FILE_LOG_JOB_STATUS = "job-status" + EXTENSION_LOG
	FILE_LOG_JOB_OUTPUT = "job-output" + EXTENSION_LOG
)

const (
	FILENAME_BIN_CONFIG_BASH = "BASH"
)

// File permissions are the designed permission flags used in Monteur
const (
	PERMISSION_EXECUTABLE = 0700
	PERMISSION_CONFIG     = 0600
	PERMISSION_DIRECTORY  = 0755
)
