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

package libworkspace

import (
	"fmt"
	"path/filepath"
	"runtime"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/schema"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/styler"
)

// Workspace is the Monteur continuous integration main data sructure.
//
// This data structure is responsible for running through all Monteur operations
type Workspace struct {
	Filesystem *Pathing
	Language   *schema.Language
	App        *schema.Software

	Job           string
	Version       string
	OS            string
	ARCH          string
	ComputeSystem string
	ConfigDir     string
	JobTOMLFile   string
}

// Init is to initialize the workspace for usage
func (w *Workspace) Init() error {
	if err := w._parseWorkspaceData(); err != nil {
		return err
	}

	if err := w._parseAppData(); err != nil {
		return err
	}

	w.OS = runtime.GOOS
	w.ARCH = runtime.GOARCH
	w.Version = libmonteur.VERSION
	w.ComputeSystem = w.OS + libmonteur.COMPUTE_SYSTEM_SEPARATOR + w.ARCH
	w._processPathingByJob()

	return nil
}

func (w *Workspace) _parseWorkspaceData() (err error) {
	w.Filesystem = &Pathing{}
	w.Language = &schema.Language{}

	err = w.Filesystem.Init()
	if err != nil {
		return err
	}

	s := struct {
		Language   *schema.Language
		Filesystem *Pathing
	}{
		Language:   w.Language,
		Filesystem: w.Filesystem,
	}

	err = toml.DecodeFile(w.Filesystem.WorkspaceTOMLFile, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	err = w.Filesystem.Update()
	if err != nil {
		return err
	}

	return nil
}

func (w *Workspace) _parseAppData() (err error) {
	if w.App == nil {
		w.App = &schema.Software{}
	}

	p := w.Filesystem.Join(w.Filesystem.AppConfigDir,
		w.Language.AlternateName+libmonteur.EXTENSION_TOML)

	s := &struct {
		Metadata *schema.Software
	}{
		Metadata: w.App,
	}

	err = toml.DecodeFile(p, s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	return nil
}

func (w *Workspace) _processPathingByJob() {
	switch w.Job {
	case libmonteur.JOB_SETUP:
	case libmonteur.JOB_CLEAN:
	case libmonteur.JOB_TEST:
		w.ConfigDir = w.Filesystem.TestConfigDir

		w.Filesystem.WorkspaceLogDir = filepath.Join(
			w.Filesystem.LogDir,
			libmonteur.DIRECTORY_TEST,
			w.Filesystem.WorkspaceLogDir,
		)

		w.JobTOMLFile = w.Filesystem.TestTOMLFile
	case libmonteur.JOB_BUILD:
		w.ConfigDir = w.Filesystem.BuildConfigDir

		w.Filesystem.WorkspaceLogDir = filepath.Join(
			w.Filesystem.LogDir,
			libmonteur.DIRECTORY_BUILD,
			w.Filesystem.WorkspaceLogDir,
		)

		w.JobTOMLFile = w.Filesystem.BuildTOMLFile
	case libmonteur.JOB_PACKAGE:
	case libmonteur.JOB_RELEASE:
	case libmonteur.JOB_COMPOSE:
		w.ConfigDir = w.Filesystem.ComposeConfigDir

		w.Filesystem.WorkspaceLogDir = filepath.Join(
			w.Filesystem.LogDir,
			libmonteur.DIRECTORY_COMPOSE,
			w.Filesystem.WorkspaceLogDir,
		)

		w.JobTOMLFile = w.Filesystem.ComposeTOMLFile
	case libmonteur.JOB_PUBLISH:
		w.ConfigDir = w.Filesystem.PublishConfigDir

		w.Filesystem.WorkspaceLogDir = filepath.Join(
			w.Filesystem.LogDir,
			libmonteur.DIRECTORY_PUBLISH,
			w.Filesystem.WorkspaceLogDir,
		)

		w.JobTOMLFile = w.Filesystem.PublishTOMLFile
	default:
		panic("Monteur DEV: what kind of CI Job is this? ➤ " + w.Job)
	}
}

// String is the standard string interface for printing out Workspace data.
//
// Since Workspace is a complicated data structure, its String() method has to
// be uniquely constructed.
func (w *Workspace) String() string {
	return styler.BoxString("Das Monteur", styler.BORDER_DOUBLE) +
		w._stringCIBasic() + "\n" +
		w._stringCILocation() + "\n" +
		w._stringApp()
}

func (w *Workspace) _stringCIBasic() string {
	return fmt.Sprintf(`VERSION
%s

LANGUAGE
%s (%s)
`,
		w.Version,
		w.Language.Name, w.Language.AlternateName,
	)
}

func (w *Workspace) _stringCILocation() string {
	return styler.BoxString("CI Pathing", styler.BORDER_SINGLE) +
		fmt.Sprintf(`CURRENT DIRECTORY LOCATION
%s

ROOT REPOSITORY LOCATION
%s

MONTEUR CONFIG LOCATION
%s

SOURCE CODE BASE LOCATION
%s

WORKING LOCATION
%s

BUILD LOCATION
%s

SCRIPT LOCATION
%s

BIN LOCATION
%s

LOG LOCATION
%s
`,
			w.Filesystem.CurrentDir,
			w.Filesystem.RootDir,
			w.Filesystem.ConfigDir,
			w.Filesystem.BaseDir,
			w.Filesystem.WorkingDir,
			w.Filesystem.BuildDir,
			w.Filesystem.ScriptDir,
			w.Filesystem.BinDir,
			w.Filesystem.LogDir,
		)
}

func (w *Workspace) _stringApp() string {
	return styler.BoxString("Product Metadata", styler.BORDER_SINGLE) +
		fmt.Sprintf(`NAME
%s
`,
			w.App.Name,
		)
}
