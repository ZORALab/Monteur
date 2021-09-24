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

package filesystem

//nolint:typecheck
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/endec/toml"
)

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

// Filepath is the data structure holding the Monteur working filesystem.
//
// This data structure is part of App{} data structure.
type Filepath struct {
	CurrentDir string
	RootDir    string
	ConfigDir  string
	BaseDir    string
	WorkingDir string
	BuildDir   string
	ScriptDir  string
	BinDir     string
	DocDir     string
}

// Init is the method to initialize Filepath data structure via Monteur configs.
//
// If any error is found, it needs to be prompted to the repository developers.
//
// This method was designed not to panic. If it does, it means Monteur's
// developers did something funny on Monteur side and Monteur developers must
// fix the problem immediately.
//nolint:lll
func (fp *Filepath) Init() error {
	if err := fp.initCurrentDir(); err != nil {
		return err
	}

	if err := fp.initRootDir(); err != nil {
		return err
	}

	if err := fp.initConfigDir(); err != nil {
		return err
	}

	if err := fp.parseTOML(); err != nil {
		return err
	}

	if err := fp.initDependentDir(&fp.BaseDir, "BaseDir"); err != nil {
		return err
	}

	if err := fp.initDependentDir(&fp.WorkingDir, "WorkingDir"); err != nil {
		return err
	}

	if err := fp.initDependentDir(&fp.BuildDir, "BuildDir"); err != nil {
		return err
	}

	if err := fp.initDependentDir(&fp.ScriptDir, "ScriptDir"); err != nil {
		return err
	}

	if err := fp.initDependentDir(&fp.BinDir, "BinDir"); err != nil {
		return err
	}

	if err := fp.initDependentDir(&fp.DocDir, "DocDir"); err != nil {
		return err
	}

	return nil
}

// `name` is mainly for specifying the variable name without needing to import
// the heavy reflect package to do the job.
func (fp *Filepath) initDependentDir(p *string, name string) (err error) {
	if *p == "" {
		return fmt.Errorf("%s: %s", ERROR_MISSING_DIR, name)
	}

	*p = filepath.Join(fp.RootDir, *p)

	return nil
}

func (fp *Filepath) parseTOML() (err error) {
	configFile := filepath.Join(fp.ConfigDir, WORKSPACE_TOML_FILE)
	s := struct{ Filesystem *Filepath }{Filesystem: fp}

	err = toml.DecodeFile(configFile, &s, nil)
	if err != nil {
		return fmt.Errorf("%s", ERROR_FAILED_CONFIG_DECODE)
	}

	return nil
}

func (fp *Filepath) initConfigDir() (err error) {
	fp.ConfigDir = filepath.Join(fp.RootDir, MONTEUR_CFG_NAME)
	return nil
}

func (fp *Filepath) initRootDir() (err error) {
	gitDir := filepath.Join(fp.CurrentDir, GIT_DIRECTORY_NAME)
	monteurConfig := filepath.Join(fp.CurrentDir, MONTEUR_CFG_NAME)

	if IsDirExists(monteurConfig) && IsDirExists(gitDir) {
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
		gitDir = filepath.Join(dir, GIT_DIRECTORY_NAME)
		monteurConfig = filepath.Join(dir, MONTEUR_CFG_NAME)

		if IsDirExists(monteurConfig) && IsDirExists(gitDir) {
			fp.RootDir = dir
		}

		if dir == last {
			break
		}
	}

	if fp.RootDir == "" {
		return fmt.Errorf(ERROR_BAD_REPO)
	}

	return nil
}

func (fp *Filepath) initCurrentDir() (err error) {
	fp.CurrentDir, err = filepath.Abs(".")
	if err != nil {
		fp.CurrentDir = ""
		return fmt.Errorf(ERROR_CURRENT_DIRECTORY)
	}

	return nil
}
