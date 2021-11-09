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

//nolint:stylecheck,revive
package monteur

// Critical Object Names are the file or directory names critical for Monteur.
//
// These critical object names are mainly to locate root repository with Monteur
// supports.
const (
	TOML_EXTENSION           = ".toml"
	GIT_DIRECTORY_NAME       = ".git"
	MONTEUR_CFG_NAME         = ".configs/monteur"
	WORKSPACE_TOML_FILE      = "workspace" + TOML_EXTENSION
	COMPUTE_SYSTEM_SEPARATOR = "-"
	BIN_CONFIG_FILENAME      = "config"
)

// Critical object names are the file or directory names critical for Setup Fx.
const (
	SETUP_CONFIG_TOML_FILE     = "setup/config" + TOML_EXTENSION
	SETUP_PROGRAMS_DIRECTORY   = "setup/programs/"
	SETUP_PROGRAMS_PERMISSION  = 0700
	SETUP_CONFIG_PERMISSION    = 0600
	SETUP_DIRECTORY_PERMISSION = 0755
)

// Standard config files pathing defined for Monteur to operate consistently.
const (
	APP_DATA_DIR = "app"
)

// Checksum constants for setting the checksum hasher and parser types.
//
// These are the supported checksum formats and algo types.
const (
	CHECKSUM_FORMAT_BASE64     = "base64"
	CHECKSUM_FORMAT_HEX        = "hex"
	CHECKSUM_FORMAT_BASE64_URL = "base64-url"

	CHECKSUM_ALGO_SHA512 = "sha512"
	CHECKSUM_ALGO_SHA256 = "sha256"
	CHECKSUM_ALGO_MD5    = "md5"
)

// Dependencies sourcing type enumerated value for setup function.
//
// It is used in every toml config file inside setup/programs/ config directory
// for identifying how to source your dependency program.
const (
	BIN_PROGRAM_TYPE_HTTPS_DOWNLOAD = "https-download"
	BIN_PROGRAM_TYPE_LOCAL_SYSTEM   = "local-system"
)

// Supported operating function types enumerated IDs.
//
// It is used in every toml config file inside setup/program/ config directory
// for compatible function types across different stages.
const (
	BIN_PROGRAM_FORMAT_TAR_GZ = "tar.gz"
	BIN_PROGRAM_FORMAT_ZIP    = "zip"

	BIN_PROGRAM_SETUP_INSTRUCTION_MOVE   = "move"
	BIN_PROGRAM_SETUP_INSTRUCTION_SCRIPT = "script"
)

// Supported variables keys in key:value variables placholders.
//
// It is used in every toml config file inside setup/program/ config directory
// for placeholding variable elements in the fields' value.
const (
	BIN_PROGRAM_VAR_OS      = "OS"
	BIN_PROGRAM_VAR_ARCH    = "Arch"
	BIN_PROGRAM_VAR_COMPUTE = "ComputeSystem"
	BIN_PROGRAM_VAR_TMP     = "WorkingDir"
	BIN_PROGRAM_VAR_BIN     = "BinDir"
	BIN_PROGRAM_VAR_CFG     = "ConfigDir"
	BIN_PROGRAM_VAR_ARCHIVE = "Archive"
	BIN_PROGRAM_VAR_FORMAT  = "Format"
	BIN_PROGRAM_VAR_METHOD  = "Method"
	BIN_PROGRAM_VAR_URL     = "URL"
)

// Channel messages key for key:value identifications in chmsg.Message object.
const (
	chmsg_ERROR  = "error"
	chmsg_STATUS = "status"
	chmsg_DONE   = "done"
)

// Filesystem tags are common keys for setting various working directories
// pathing.
const (
	// BASEDIR_TAG is the source codes based location.
	//
	// The common location would be at the root repository. Otherwise, for
	// repository is operating as a workspace, the source codes base
	// location can be elsewhere inside the repo. Example, Monteur uses
	// pkg/ directory as its BASEDIR_TAG since it needs the root repository
	// location to work.
	BASEDIR_TAG = "BaseDir"

	// WORKINGDIR_TAG is the location where Monteur will be working in.
	//
	// The common location would be a `tmp/` directory at the root
	// repository. This is mainly a safe location for Monteur to operate
	// its testing, building, packaging, and etc.
	WORKINGDIR_TAG = "WorkingDir"

	// BUILDDIR_TAG is the location where Monteur will park built outputs.
	//
	// The common location would be an `image/` directory at the root
	// repository. This is mainly used for parking all the built outputs.
	BUILDDIR_TAG = "BuildDir"

	// SCRIPTSDIR_TAG is the location of all localized executable scripts.
	//
	// The common location would be `scripts/` or `.scripts/` directory at
	// root repository. This is mainly for Monteur to search for executable
	// commands using script. Beware that not all operating system uses
	// the same script intepreter.
	SCRIPTDIR_TAG = "ScriptDir"

	// BINDIR_TAG is the location for housing build commands binaries.
	//
	// The common location would be `.bin/` directory at root repository.
	// This is mainly for Monteur to search for build commands and also
	// the location where Monteur setup all the build binaries.
	BINDIR_TAG = "BinDir"

	// BINCFG_TAG is the location of all local binary configurations.
	//
	// The common location would be `.bin/config.d/` directory inside
	// BINDIR_TAG. This is mainly to store all binary configurations for
	// sourcing.
	BINCFG_TAG = "BinCfgDir"

	// DOCDIR_TAG is the location for housing repo's documentations.
	//
	// The common location would be `public/` directory at root repository.
	// This is mainly for Monteur to park all compiled documentations
	// outputs ready for publishing.
	DOCDIR_TAG = "DocsDir"
)
