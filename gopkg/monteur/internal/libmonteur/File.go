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
	FORMAT_TOML = "toml"
	FORMAT_CSV  = "csv"
	FORMAT_TXT  = "txt"
)

const (
	DIRECTORY_GIT              = ".git"
	DIRECTORY_MONTEUR_CONFIG_D = "config.d"
	DIRECTORY_MONTEUR_CONFIG   = ".configs/monteur"
	DIRECTORY_JOBS             = "jobs"

	DIRECTORY_APP           = "app"
	DIRECTORY_APP_CONFIG    = DIRECTORY_APP + "/config"
	DIRECTORY_APP_COPYRIGHT = "copyrights" // has lang prefix

	DIRECTORY_PUBLISH = "publish"
	DIRECTORY_CLEAN   = "clean"
	DIRECTORY_COMPOSE = "compose"
	DIRECTORY_SETUP   = "setup"
	DIRECTORY_TEST    = "test"
	DIRECTORY_BUILD   = "build"
	DIRECTORY_PACKAGE = "package"
	DIRECTORY_RELEASE = "release"
)

const (
	FILE_TOML               = "config" + EXTENSION_TOML
	FILE_TOML_WORKSPACE     = "workspace" + EXTENSION_TOML
	FILE_TOML_APP_METADATA  = "metadata" + EXTENSION_TOML
	FILE_TOML_APP_COPYRIGHT = "copyrights" + EXTENSION_TOML
	FILE_TOML_APP_HELP      = "help" + EXTENSION_TOML
	FILE_TOML_APP_DEBIAN    = "debian" + EXTENSION_TOML

	FILE_LOG_PREFIX_JOB = "job-"

	FILE_LOG_OUTPUT = "output" + EXTENSION_LOG
	FILE_LOG_STATUS = "status" + EXTENSION_LOG
)

const (
	FILENAME_BIN_CONFIG_MAIN = "main"
)

// File permissions are the designed permission flags used in Monteur
const (
	PERMISSION_EXECUTABLE = 0700
	PERMISSION_CONFIG     = 0600
	PERMISSION_DIRECTORY  = 0755
)
