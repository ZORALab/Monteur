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
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/filesystem"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

type UserPath struct {
	Home string
}

// Pathing is the data structure holding the Monteur working filesystem.
type Pathing struct {
	timestampDir string

	CurrentDir string
	RootDir    string
	ConfigDir  string

	// user input paths
	BaseDir    string
	WorkingDir string
	BuildDir   string
	ScriptDir  string
	BinDir     string
	BinCfgDir  string
	LogDir     string
	DataDir    string
	ReleaseDir string

	// workspace Pathing
	WorkspaceTOMLFile string
	WorkspaceLogDir   string

	// app
	AppConfigDir      string
	AppMetaTOMLFile   string
	AppHelpTOMLFile   string
	AppDebianTOMLFile string
	AppCopyrightsDir  string

	// sub-directories for setup fx
	SetupTMPDir    string
	SetupConfigDir string
	SetupTOMLFile  string

	// sub-directories for publish fx
	PublishTMPDir    string
	PublishConfigDir string
	PublishTOMLFile  string

	// sub-directories for compose fx
	ComposeConfigDir string
	ComposeTMPDir    string
	ComposeTOMLFile  string

	// sub-directories for test fx
	TestConfigDir string
	TestTMPDir    string
	TestTOMLFile  string

	// sub-directories for prepare fx
	PrepareConfigDir string
	PrepareTMPDir    string
	PrepareTOMLFile  string

	// sub-directories for build fx
	BuildConfigDir string
	BuildTMPDir    string
	BuildTOMLFile  string

	// sub-directories for package fx
	PackageConfigDir string
	PackageTMPDir    string
	PackageTOMLFile  string

	// sub-directories for release fx
	ReleaseConfigDir string
	ReleaseTMPDir    string
	ReleaseTOMLFile  string

	// sub-directories for clean fx
	CleanConfigDir string
	CleanTMPDir    string
	CleanTOMLFile  string

	// user
	User *UserPath

	// secrets - user input paths
	SecretsDir []string
}

// Init is the method to initialize critical Pathing for Monteur operations.
//
// If any error is found, it needs to be prompted to the repository developers
// as Monteur fails to detect any git repository with Monteur support throughout
// the current directory Pathing or failed to obtain current directory Pathing.
func (fp *Pathing) Init() (err error) {
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

	err = fp._initUserDir()
	if err != nil {
		return err
	}

	fp.WorkspaceTOMLFile = libmonteur.FILE_TOML_WORKSPACE
	err = fp._initConfigSubPath(&fp.WorkspaceTOMLFile, "WorkspaceTOMLFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) _initConfigDir() (err error) {
	fp.ConfigDir = filepath.Join(fp.RootDir,
		libmonteur.DIRECTORY_MONTEUR_CONFIG,
	)
	return nil
}

func (fp *Pathing) _initRootDir() (err error) {
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

func (fp *Pathing) _initCurrentDir(g func(string) (string, error)) (err error) {
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

func (fp *Pathing) _initConfigSubPath(p *string, name string) (err error) {
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

func (fp *Pathing) _initUserDir() (err error) {
	fp.User = &UserPath{}

	fp.User.Home, err = os.UserHomeDir()
	if err != nil {
		fp.User.Home = ""
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_BAD,
			"HOME",
		)
	}

	return nil
}

// `Update` is the method to update all dependent Pathings to full path.
//
// If any error is found, it needs to be prompted to the repository developers
// to resolve the Pathing issues in the workspace.toml configuration file.
//
// If panic is found, it means Monteur developers did something funny and
// screwed things up. Monteur developers shall not release Monteur that causes
// panics.
//
// This function should only be called after all the relative paths are filled
// into Pathing data structure AND the Pathing was initialized successfully via
// `Init()` function.
func (fp *Pathing) Update(langCode string) (err error) {
	err = fp.updateBasePaths(langCode)
	if err != nil {
		return err
	}

	err = fp.updateSetupPaths()
	if err != nil {
		return err
	}

	err = fp.updatePublishPaths()
	if err != nil {
		return err
	}

	err = fp.updateComposePaths()
	if err != nil {
		return err
	}

	err = fp.updateTestPaths()
	if err != nil {
		return err
	}

	err = fp.updateBuildPaths()
	if err != nil {
		return err
	}

	err = fp.updatePackagePaths()
	if err != nil {
		return err
	}

	err = fp.updateReleasePaths()
	if err != nil {
		return err
	}

	err = fp.updateCleanPaths()
	if err != nil {
		return err
	}

	err = fp.updatePreparePaths()
	if err != nil {
		return err
	}

	err = fp.updateLogPaths()
	if err != nil {
		return err
	}

	err = fp.updateAppPaths()
	if err != nil {
		return err
	}

	err = fp._initSecretsDir()
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateBasePaths(langCode string) (err error) {
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

	err = fp._initDependentDir(&fp.LogDir, "LogDir")
	if err != nil {
		return err
	}

	err = fp._initDependentDir(&fp.DataDir, "DataDir")
	if err != nil {
		return err
	}

	err = fp._initDependentDir(&fp.ReleaseDir, "ReleaseDir")
	if err != nil {
		return err
	}

	fp.AppConfigDir = filepath.Join(libmonteur.DIRECTORY_APP_CONFIG,
		langCode)
	err = fp._initConfigSubPath(&fp.AppConfigDir, "AppConfigDir")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateSetupPaths() (err error) {
	fp.SetupConfigDir = filepath.Join(libmonteur.DIRECTORY_SETUP,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.SetupConfigDir, "SetupConfigDir")
	if err != nil {
		return err
	}

	fp.SetupTMPDir = libmonteur.DIRECTORY_SETUP
	err = fp._initWorkingSubPath(&fp.SetupTMPDir, "SetupTMPDir")
	if err != nil {
		return err
	}

	fp.SetupTOMLFile = filepath.Join(libmonteur.DIRECTORY_SETUP,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.SetupTOMLFile,
		"SetupTOMLConfigFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updatePublishPaths() (err error) {
	fp.PublishConfigDir = filepath.Join(libmonteur.DIRECTORY_PUBLISH,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.PublishConfigDir, "PublishConfigDir")
	if err != nil {
		return err
	}

	fp.PublishTMPDir = libmonteur.DIRECTORY_PUBLISH
	err = fp._initWorkingSubPath(&fp.PublishTMPDir, "PublishTMPDir")
	if err != nil {
		return err
	}

	fp.PublishTOMLFile = filepath.Join(libmonteur.DIRECTORY_PUBLISH,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.PublishTOMLFile,
		"PublishTOMLConfigFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateComposePaths() (err error) {
	fp.ComposeConfigDir = filepath.Join(libmonteur.DIRECTORY_COMPOSE,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.ComposeConfigDir, "ComposeConfigDir")
	if err != nil {
		return err
	}

	fp.ComposeTMPDir = libmonteur.DIRECTORY_COMPOSE
	err = fp._initWorkingSubPath(&fp.ComposeTMPDir, "ComposeTMPDir")
	if err != nil {
		return err
	}

	fp.ComposeTOMLFile = filepath.Join(libmonteur.DIRECTORY_COMPOSE,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.ComposeTOMLFile,
		"ComposeTOMLConfigFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateTestPaths() (err error) {
	fp.TestConfigDir = filepath.Join(libmonteur.DIRECTORY_TEST,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.TestConfigDir, "TestConfigDir")
	if err != nil {
		return err
	}

	fp.TestTMPDir = filepath.Join(libmonteur.DIRECTORY_TEST,
		fp.timestampDir,
	)
	err = fp._initWorkingSubPath(&fp.TestTMPDir, "TestTMPDir")
	if err != nil {
		return err
	}

	fp.TestTOMLFile = filepath.Join(libmonteur.DIRECTORY_TEST,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.TestTOMLFile, "TestTOMLConfigFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateBuildPaths() (err error) {
	fp.BuildConfigDir = filepath.Join(libmonteur.DIRECTORY_BUILD,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.BuildConfigDir, "BuildConfigDir")
	if err != nil {
		return err
	}

	fp.BuildTMPDir = libmonteur.DIRECTORY_BUILD
	err = fp._initWorkingSubPath(&fp.BuildTMPDir, "BuildTMPDir")
	if err != nil {
		return err
	}

	fp.BuildTOMLFile = filepath.Join(libmonteur.DIRECTORY_BUILD,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.BuildTOMLFile, "BuildTOMLFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updatePackagePaths() (err error) {
	fp.PackageConfigDir = filepath.Join(libmonteur.DIRECTORY_PACKAGE,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.PackageConfigDir, "PackageConfigDir")
	if err != nil {
		return err
	}

	fp.PackageTMPDir = libmonteur.DIRECTORY_PACKAGE
	err = fp._initWorkingSubPath(&fp.PackageTMPDir, "PackageTMPDir")
	if err != nil {
		return err
	}

	fp.PackageTOMLFile = filepath.Join(libmonteur.DIRECTORY_PACKAGE,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.PackageTOMLFile, "PackageTOMLFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateReleasePaths() (err error) {
	fp.ReleaseConfigDir = filepath.Join(libmonteur.DIRECTORY_RELEASE,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.ReleaseConfigDir, "ReleaseConfigDir")
	if err != nil {
		return err
	}

	fp.ReleaseTMPDir = libmonteur.DIRECTORY_RELEASE
	err = fp._initWorkingSubPath(&fp.ReleaseTMPDir, "ReleaseTMPDir")
	if err != nil {
		return err
	}

	fp.ReleaseTOMLFile = filepath.Join(libmonteur.DIRECTORY_PACKAGE,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.ReleaseTOMLFile, "ReleaseTOMLFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateCleanPaths() (err error) {
	fp.CleanConfigDir = filepath.Join(libmonteur.DIRECTORY_CLEAN,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.CleanConfigDir, "CleanConfigDir")
	if err != nil {
		return err
	}

	fp.CleanTMPDir = libmonteur.DIRECTORY_CLEAN
	err = fp._initWorkingSubPath(&fp.CleanTMPDir, "CleanTMPDir")
	if err != nil {
		return err
	}

	fp.CleanTOMLFile = filepath.Join(libmonteur.DIRECTORY_CLEAN,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.CleanTOMLFile, "CleanTOMLFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateAppPaths() (err error) {
	fp.AppMetaTOMLFile = libmonteur.FILE_TOML_APP_METADATA
	err = fp._initAppConfigPath(&fp.AppMetaTOMLFile, "AppMetaTOMLFile")
	if err != nil {
		return err
	}

	fp.AppCopyrightsDir = libmonteur.DIRECTORY_APP_COPYRIGHT
	err = fp._initAppConfigPath(&fp.AppCopyrightsDir,
		"AppCopyrightsTOMLFile",
	)
	if err != nil {
		return err
	}

	fp.AppHelpTOMLFile = libmonteur.FILE_TOML_APP_HELP
	err = fp._initAppConfigPath(&fp.AppHelpTOMLFile, "AppHelpTOMLFile")
	if err != nil {
		return err
	}

	fp.AppDebianTOMLFile = libmonteur.FILE_TOML_APP_DEBIAN
	err = fp._initAppConfigPath(&fp.AppDebianTOMLFile, "AppDebianTOMLFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updatePreparePaths() (err error) {
	fp.PrepareConfigDir = filepath.Join(libmonteur.DIRECTORY_PREPARE,
		libmonteur.DIRECTORY_JOBS,
	)
	err = fp._initConfigSubPath(&fp.PrepareConfigDir, "PrepareConfigDir")
	if err != nil {
		return err
	}

	fp.PrepareTMPDir = libmonteur.DIRECTORY_PREPARE
	err = fp._initWorkingSubPath(&fp.PrepareTMPDir, "PrepareTMPDir")
	if err != nil {
		return err
	}

	fp.PrepareTOMLFile = filepath.Join(libmonteur.DIRECTORY_PREPARE,
		libmonteur.FILE_TOML,
	)
	err = fp._initConfigSubPath(&fp.PrepareTOMLFile, "PrepareTOMLFile")
	if err != nil {
		return err
	}

	return nil
}

func (fp *Pathing) updateLogPaths() (err error) {
	fp.WorkspaceLogDir = fp.timestampDir

	return nil
}

func (fp *Pathing) _initAppConfigPath(p *string, name string) (err error) {
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

	if fp.AppConfigDir == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_MISSING,
			"AppConfigDir",
		)
	}

	*p = filepath.Join(fp.AppConfigDir, *p)

	return nil
}

func (fp *Pathing) _initWorkingSubPath(p *string, name string) (err error) {
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

func (fp *Pathing) _initDependentDir(p *string, name string) (err error) {
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

func (fp *Pathing) _initSecretsDir() (err error) {
	if fp.SecretsDir == nil {
		fp.SecretsDir = []string{}
		return nil
	}

	if len(fp.SecretsDir) == 0 {
		return nil
	}

	data := map[string]string{
		libmonteur.BASEDIR_TAG:     fp.BaseDir,
		libmonteur.WORKINGDIR_TAG:  fp.WorkingDir,
		libmonteur.BUILDDIR_TAG:    fp.BuildDir,
		libmonteur.BINDIR_TAG:      fp.BinDir,
		libmonteur.BINCFG_TAG:      fp.BinCfgDir,
		libmonteur.USERHOMEDIR_TAG: fp.User.Home,
		libmonteur.ROOTDIR_TAG:     fp.RootDir,
	}
	list := []string{}

	for _, path := range fp.SecretsDir {
		t, err := template.New("Value").Parse(path)
		if err != nil {
			return fmt.Errorf("%s: %s for %s",
				libmonteur.ERROR_DIR_BAD,
				err,
				path,
			)
		}

		var b bytes.Buffer
		if err := t.Execute(&b, data); err != nil {
			return fmt.Errorf("%s: %s for %s",
				libmonteur.ERROR_DIR_BAD,
				err,
				path,
			)
		}

		list = append(list, b.String())
	}

	fp.SecretsDir = list

	return nil
}

// Join is a public API for joining multiple Pathing into one.
//
// It returns a fully compatible filepath.Join(...) string value.
func (fp *Pathing) Join(paths ...string) string {
	return filepath.Join(paths...)
}
