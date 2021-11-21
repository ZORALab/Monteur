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

	Version       string
	OS            string
	ARCH          string
	ComputeSystem string
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

DOC LOCATION
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
			w.Filesystem.DocDir,
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
