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

package monteur

//nolint:stylecheck,revive
// Critical Object Names are the file or directory names critical for Monteur.
//
// These critical object names are mainly to locate root repository with Monteur
// supports.
const (
	GIT_DIRECTORY_NAME  = ".git"
	MONTEUR_CFG_NAME    = ".configs/monteur"
	WORKSPACE_TOML_FILE = "workspace.toml"
)

// Standard config files pathing defined for Monteur to operate consistently.
const (
	APP_DATA_DIR = "app"
)

//nolint:stylecheck,revive
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

	// DOCDIR_TAG is the location for housing repo's documentations.
	//
	// The common location would be `public/` directory at root repository.
	// This is mainly for Monteur to park all compiled documentations
	// outputs ready for publishing.
	DOCDIR_TAG = "DocsDir"
)
