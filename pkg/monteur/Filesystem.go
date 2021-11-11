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

import (
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
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
	BASEDIR_TAG = libmonteur.BASEDIR_TAG

	// WORKINGDIR_TAG is the location where Monteur will be working in.
	//
	// The common location would be a `tmp/` directory at the root
	// repository. This is mainly a safe location for Monteur to operate
	// its testing, building, packaging, and etc.
	WORKINGDIR_TAG = libmonteur.WORKINGDIR_TAG

	// BUILDDIR_TAG is the location where Monteur will park built outputs.
	//
	// The common location would be an `image/` directory at the root
	// repository. This is mainly used for parking all the built outputs.
	BUILDDIR_TAG = libmonteur.BUILDDIR_TAG

	// SCRIPTSDIR_TAG is the location of all localized executable scripts.
	//
	// The common location would be `scripts/` or `.scripts/` directory at
	// root repository. This is mainly for Monteur to search for executable
	// commands using script. Beware that not all operating system uses
	// the same script intepreter.
	SCRIPTDIR_TAG = libmonteur.SCRIPTDIR_TAG

	// BINDIR_TAG is the location for housing build commands binaries.
	//
	// The common location would be `.bin/` directory at root repository.
	// This is mainly for Monteur to search for build commands and also
	// the location where Monteur setup all the build binaries.
	BINDIR_TAG = libmonteur.BINDIR_TAG

	// BINCFG_TAG is the location of all local binary configurations.
	//
	// The common location would be `.bin/config.d/` directory inside
	// BINDIR_TAG. This is mainly to store all binary configurations for
	// sourcing.
	BINCFG_TAG = libmonteur.BINCFG_TAG

	// DOCDIR_TAG is the location for housing repo's documentations.
	//
	// The common location would be `public/` directory at root repository.
	// This is mainly for Monteur to park all compiled documentations
	// outputs ready for publishing.
	DOCDIR_TAG = libmonteur.DOCDIR_TAG
)
