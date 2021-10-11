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
	"runtime"

	//nolint:typecheck
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/endec/toml"
	//nolint:typecheck
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/schema"
	//nolint:typecheck
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/styler"
)

// Workspace is the Monteur continuous integration main data sructure.
//
// This data structure is responsible for running through all Monteur operations
type Workspace struct {
	version       string
	os            string
	arch          string
	computeSystem string

	language   *schema.Language
	filesystem *pathing
	app        *schema.Software
}

func (w *Workspace) Init() error {
	if err := w._parseWorkspaceData(); err != nil {
		return err
	}

	if err := w._parseAppData(); err != nil {
		return err
	}

	w.os = runtime.GOOS
	w.arch = runtime.GOARCH
	w.computeSystem = w.os + COMPUTE_SYSTEM_SEPARATOR + w.arch

	return nil
}

func (w *Workspace) _parseWorkspaceData() (err error) {
	w.version = VERSION
	w.filesystem = &pathing{}
	w.language = &schema.Language{}

	err = w.filesystem.Init()
	if err != nil {
		return err
	}

	s := struct {
		Language   *schema.Language
		Filesystem *pathing
	}{
		Language:   w.language,
		Filesystem: w.filesystem,
	}

	err = toml.DecodeFile(w.filesystem.WorkspaceTOMLFile, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FAILED_CONFIG_DECODE, err)
	}

	err = w.filesystem.Update()
	if err != nil {
		return err
	}

	return nil
}

func (w *Workspace) _parseAppData() (err error) {
	if w.app == nil {
		w.app = &schema.Software{}
	}

	p := w.filesystem.Join(w.filesystem.AppConfigDir,
		w.language.AlternateName+TOML_EXTENSION)

	s := &struct {
		Metadata *schema.Software
	}{
		Metadata: w.app,
	}

	err = toml.DecodeFile(p, s, nil)
	if err != nil {
		return err
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
		w.version,
		w.language.Name, w.language.AlternateName,
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

func (w *Workspace) _stringApp() string {
	return styler.BoxString("Product Metadata", styler.BORDER_SINGLE) +
		fmt.Sprintf(`NAME
%s
`,
			w.app.Name,
		)
}
