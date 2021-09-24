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

//nolint:typecheck
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/filesystem"
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
	DocDir     string
}

// Init is the method to initialize critical pathing for Monteur operations.
//
// If any error is found, it needs to be prompted to the repository developers
// as Monteur fails to detect any git repository with Monteur support throughout
// the current directory pathing or failed to obtain current directory pathing.
func (fp *pathing) Init() error {
	if err := fp.initCurrentDir(); err != nil {
		return err
	}

	if err := fp.initRootDir(); err != nil {
		return err
	}

	if err := fp.initConfigDir(); err != nil {
		return err
	}

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
// the heavy-duty reflect package to do such a simple job for error reporting.
func (fp *pathing) initDependentDir(p *string, name string) (err error) {
	if *p == "" {
		return fmt.Errorf("%s: %s", ERROR_MISSING_DIR, name)
	}

	*p = filepath.Join(fp.RootDir, *p)

	return nil
}

func (fp *pathing) initConfigDir() (err error) {
	fp.ConfigDir = filepath.Join(fp.RootDir, MONTEUR_CFG_NAME)
	return nil
}

func (fp *pathing) initRootDir() (err error) {
	gitDir := filepath.Join(fp.CurrentDir, GIT_DIRECTORY_NAME)
	monteurConfig := filepath.Join(fp.CurrentDir, MONTEUR_CFG_NAME)

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
		gitDir = filepath.Join(dir, GIT_DIRECTORY_NAME)
		monteurConfig = filepath.Join(dir, MONTEUR_CFG_NAME)

		if filesystem.IsDirExists(monteurConfig) &&
			filesystem.IsDirExists(gitDir) {
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

func (fp *pathing) initCurrentDir() (err error) {
	fp.CurrentDir, err = filepath.Abs(".")
	if err != nil {
		fp.CurrentDir = ""
		return fmt.Errorf(ERROR_CURRENT_DIRECTORY)
	}

	return nil
}
