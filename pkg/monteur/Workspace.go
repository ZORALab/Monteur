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
	"fmt"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/filesystem" //nolint:typecheck
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

// Workspace is the Monteur continuous integration main data sructure.
//
// This data structure is responsible for running through all Monteur operations
type Workspace struct {
	filesystem *filesystem.Filepath
}

func (w *Workspace) Init() error {
	w.filesystem = &filesystem.Filepath{}
	if err := w.filesystem.Init(); err != nil {
		return err
	}

	return nil
}

func (w *Workspace) String() string {
	return fmt.Sprintf(`DIRECTORY LIST
--------------
CURRENT
%s

ROOT REPO
%s

MONTEUR CONFIG
%s

SOURCE CODE BASE
%s

WORKING
%s

BUILD
%s

SCRIPT
%s

BIN
%s

DOC
%s
`,
		w.filesystem.CurrentDir,
		w.filesystem.RootDir,
		w.filesystem.ConfigDir,
		w.filesystem.BaseDir,
		w.filesystem.WorkingDir,
		w.filesystem.BuildDir,
		w.filesystem.ScriptDir,
		w.filesystem.BinDir,
		w.filesystem.DocDir,
	)
}
