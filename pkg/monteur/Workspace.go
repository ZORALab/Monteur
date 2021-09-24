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
	"path/filepath"

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
	version    string
	language   *schema.Language
	filesystem *pathing
	app        *schema.Software
}

func (w *Workspace) Init() error {
	if err := w.parseWorkspaceData(); err != nil {
		return err
	}

	if err := w.parseAppData(); err != nil {
		return err
	}

	return nil
}

func (w *Workspace) parseWorkspaceData() (err error) {
	w.version = VERSION
	w.filesystem = &pathing{}
	w.language = &schema.Language{}

	err = w.filesystem.Init()
	if err != nil {
		return err
	}

	p := filepath.Join(w.filesystem.ConfigDir, WORKSPACE_TOML_FILE)
	s := struct {
		Language   *schema.Language
		Filesystem *pathing
	}{
		Language:   w.language,
		Filesystem: w.filesystem,
	}

	err = toml.DecodeFile(p, &s, nil)
	if err != nil {
		return fmt.Errorf("%s", ERROR_FAILED_CONFIG_DECODE)
	}

	err = w.filesystem.Update()
	if err != nil {
		return err
	}

	return nil
}

func (w *Workspace) parseAppData() (err error) {
	if w.app == nil {
		w.app = &schema.Software{}
	}

	p := filepath.Join(w.filesystem.ConfigDir,
		APP_DATA_DIR,
		w.language.AlternateName+".toml")

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
		w.stringCIBasic() + "\n" +
		w.stringCILocation() + "\n" +
		w.stringApp()
}

func (w *Workspace) stringCIBasic() string {
	return fmt.Sprintf(`VERSION
%s

LANGUAGE
%s (%s)
`,
		w.version,
		w.language.Name, w.language.AlternateName,
	)
}

func (w *Workspace) stringCILocation() string {
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

func (w *Workspace) stringApp() string {
	return styler.BoxString("Product Metadata", styler.BORDER_SINGLE) +
		fmt.Sprintf(`NAME
%s
`,
			w.app.Name,
		)
}
