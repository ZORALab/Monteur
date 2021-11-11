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
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/filesystem"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
)

// pathing is the data structure holding the Monteur working filesystem.
type pathing struct {
	CurrentDir string
	RootDir    string
	ConfigDir  string
	BaseDir    string
	WorkingDir string
	BuildDir   string
	ScriptDir  string
	BinDir     string
	BinCfgDir  string
	DocDir     string

	// workspace pathing
	WorkspaceTOMLFile string

	// app
	AppConfigDir string

	// bin
	BinConfigFile string

	// sub-directories for setup fx
	SetupTMPDir           string
	SetupProgramConfigDir string
	SetupTOMLFile         string
}

// Init is the method to initialize critical pathing for Monteur operations.
//
// If any error is found, it needs to be prompted to the repository developers
// as Monteur fails to detect any git repository with Monteur support throughout
// the current directory pathing or failed to obtain current directory pathing.
func (fp *pathing) Init() (err error) {
	err = fp._initCurrentDir(filepath.Abs)
	if err != nil {
		return err
	}

	err = fp._initRootDir()
	if err != nil {
		return err
	}

	err = fp._initConfigDir()
	if err != nil {
		return err
	}

	fp.WorkspaceTOMLFile = libmonteur.FILE_WORKSPACE_TOML
	err = fp._initConfigSubPath(&fp.WorkspaceTOMLFile, "WorkspaceTOMLFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *pathing) _initConfigDir() (err error) {
	fp.ConfigDir = filepath.Join(fp.RootDir,
		libmonteur.DIRECTORY_MONTEUR_CONFIG,
	)
	return nil
}

func (fp *pathing) _initRootDir() (err error) {
	gitDir := filepath.Join(fp.CurrentDir, libmonteur.DIRECTORY_GIT)
	monteurConfig := filepath.Join(fp.CurrentDir,
		libmonteur.DIRECTORY_MONTEUR_CONFIG,
	)

	if filesystem.IsDirExists(monteurConfig) &&
		filesystem.IsDirExists(gitDir) {
		fp.RootDir = fp.CurrentDir
	}

	// scan for possible root directory
	dir := fp.CurrentDir

	last := strings.Split(fp.CurrentDir, string(os.PathSeparator))[0]
	if last == "" {
		last = string(os.PathSeparator)
	}

	for {
		dir = filepath.Dir(dir)
		gitDir = filepath.Join(dir, libmonteur.DIRECTORY_GIT)
		monteurConfig = filepath.Join(dir,
			libmonteur.DIRECTORY_MONTEUR_CONFIG,
		)

		if filesystem.IsDirExists(monteurConfig) &&
			filesystem.IsDirExists(gitDir) {
			fp.RootDir = dir
		}

		if dir == last {
			break
		}
	}

	if fp.RootDir == "" {
		return fmt.Errorf(libmonteur.ERROR_BAD_REPO)
	}

	return nil
}

func (fp *pathing) _initCurrentDir(g func(string) (string, error)) (err error) {
	fp.CurrentDir, err = g(".")
	if err != nil {
		fp.CurrentDir = ""
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_BAD,
			"./",
		)
	}

	return nil
}

func (fp *pathing) _initConfigSubPath(p *string, name string) (err error) {
	// NOTE:
	// 1) `name` is mainly for specifying the variable name without needing
	//    to import the heavy-duty reflect package to do such a simple job
	//    for error reporting.
	if *p == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_MISSING,
			name,
		)
	}

	if fp.ConfigDir == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_MISSING,
			"ConfigDir",
		)
	}

	*p = filepath.Join(fp.ConfigDir, *p)

	return nil
}

// `Update` is the method to update all dependent pathings to full path.
//
// If any error is found, it needs to be prompted to the repository developers
// to resolve the pathing issues in the workspace.toml configuration file.
//
// If panic is found, it means Monteur developers did something funny and
// screwed things up. Monteur developers shall not release Monteur that causes
// panics.
//
// This function should only be called after all the relative paths are filled
// into pathing data structure AND the pathing was initialized successfully via
// `Init()` function.
func (fp *pathing) Update() (err error) {
	err = fp._initDependentDir(&fp.BaseDir, "BaseDir")
	if err != nil {
		return err
	}

	err = fp._initDependentDir(&fp.WorkingDir, "WorkingDir")
	if err != nil {
		return err
	}

	err = fp._initDependentDir(&fp.BuildDir, "BuildDir")
	if err != nil {
		return err
	}

	err = fp._initDependentDir(&fp.ScriptDir, "ScriptDir")
	if err != nil {
		return err
	}

	err = fp._initDependentDir(&fp.BinDir, "BinDir")
	if err != nil {
		return err
	}

	err = fp._initDependentDir(&fp.BinCfgDir, "BinCfgDir")
	if err != nil {
		return err
	}

	err = fp._initDependentDir(&fp.DocDir, "DocDir")
	if err != nil {
		return err
	}

	fp.BinConfigFile = libmonteur.FILENAME_BIN_CONFIG
	err = fp._initBinSubPath(&fp.BinConfigFile, "BinConfigFile")
	if err != nil {
		return err
	}

	fp.AppConfigDir = libmonteur.DIRECTORY_APP
	err = fp._initConfigSubPath(&fp.AppConfigDir, "WorkspaceAppDir")
	if err != nil {
		return err
	}

	fp.SetupTMPDir = libmonteur.DIRECTORY_SETUP_PROGRAMS
	err = fp._initWorkingSubPath(&fp.SetupTMPDir, "SetupTMPDir")
	if err != nil {
		return err
	}

	fp.SetupProgramConfigDir = libmonteur.DIRECTORY_SETUP_PROGRAMS
	err = fp._initConfigSubPath(&fp.SetupProgramConfigDir,
		"SetupProgramConfigDir")
	if err != nil {
		return err
	}

	fp.SetupTOMLFile = libmonteur.FILE_SETUP_TOML
	err = fp._initConfigSubPath(&fp.SetupTOMLFile,
		"SetupTOMLConfigFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *pathing) _initBinSubPath(p *string, name string) (err error) {
	// NOTE:
	// 1) `name` is mainly for specifying the variable name without needing
	//    to import the heavy-duty reflect package to do such a simple job
	//    for error reporting.
	if *p == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_MISSING,
			name,
		)
	}

	if fp.BinDir == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_MISSING,
			"BinDir",
		)
	}

	*p = filepath.Join(fp.BinDir, *p)

	return nil
}

func (fp *pathing) _initWorkingSubPath(p *string, name string) (err error) {
	// NOTE:
	// 1) `name` is mainly for specifying the variable name without needing
	//    to import the heavy-duty reflect package to do such a simple job
	//    for error reporting.
	if *p == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_MISSING,
			name,
		)
	}

	if fp.WorkingDir == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_MISSING,
			"WorkingDir",
		)
	}

	*p = filepath.Join(fp.WorkingDir, *p)

	return nil
}

func (fp *pathing) _initDependentDir(p *string, name string) (err error) {
	// NOTE:
	// 1) `name` is mainly for specifying the variable name without needing
	//    to import the heavy-duty reflect package to do such a simple job
	//    for error reporting.
	if *p == "" {
		return fmt.Errorf("%s: %s", libmonteur.ERROR_DIR_MISSING, name)
	}

	if fp.RootDir == "" {
		return fmt.Errorf(libmonteur.ERROR_BAD_REPO)
	}

	*p = filepath.Join(fp.RootDir, *p)

	return nil
}

// Join is a public API for joining multiple pathing into one.
//
// It returns a fully compatible filepath.Join(...) string value.
func (fp *pathing) Join(paths ...string) string {
	return filepath.Join(paths...)
}
