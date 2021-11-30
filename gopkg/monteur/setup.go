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

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsetup"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libworkspace"
)

type setup struct {
	workspace *libworkspace.Workspace
	settings  *libsetup.Run
	logger    *liblog.Logger
	workers   map[string]conductor.Job
	secrets   map[string]interface{}
}

// Run is to execute the setup service function in Monteur.
func (fx *setup) Run() int {
	var err error

	err = fx._init()
	if err != nil {
		return fx._reportError(err)
	}

	// parse all programs
	err = filepath.Walk(fx.workspace.Filesystem.SetupProgramConfigDir,
		fx.__filterProgramMetadata)
	if err != nil {
		return fx._reportError(err)
	}

	// setup repository
	fx.logger.Info("Cleaning %s...", fx.workspace.Filesystem.SetupTMPDir)
	err = fx.settings.CleanDir(fx.workspace.Filesystem.SetupTMPDir)
	if err != nil {
		return fx._reportError(err)
	}
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	fx.logger.Info("Cleaning %s...", fx.workspace.Filesystem.BinCfgDir)
	err = fx.settings.CleanDir(fx.workspace.Filesystem.BinCfgDir)
	if err != nil {
		return fx._reportError(err)
	}
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	fx.logger.Info("Cleaning %s...", fx.workspace.Filesystem.BinConfigdDir)
	err = fx.settings.CleanDir(fx.workspace.Filesystem.BinConfigdDir)
	if err != nil {
		return fx._reportError(err)
	}
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	fx.logger.Info("Cleaning %s...", fx.workspace.Filesystem.BinDir)
	err = fx.settings.CleanDir(fx.workspace.Filesystem.BinDir)
	if err != nil {
		return fx._reportError(err)
	}
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	fx.logger.Info("Setup %s...", fx.workspace.Filesystem.BinConfigFile)
	err = fx.settings.SetupConfig(fx.workspace.OS,
		fx.workspace.Filesystem.BinConfigFile,
		fx.workspace.Filesystem.BinDir,
		fx.workspace.Filesystem.BinConfigdDir,
	)
	if err != nil {
		return fx._reportError(err)
	}
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	// execute all tasks
	c := &conductor.Conductor{
		Runners: fx.workers,
		Log:     fx.logger,
	}

	err = c.Run()
	if err != nil {
		return fx._reportError(err)
	}

	err = c.Coordinate()
	if err != nil {
		return fx._reportError(err)
	}

	// safely close the logs and exit as completion
	fx.logger.Sync()
	fx.logger.Close()
	return STATUS_OK
}

func (fx *setup) __filterProgramMetadata(path string,
	info os.FileInfo, err error) error {
	var s *libsetup.Manager

	// return err if occurred
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// ensure we only accept the correct regular file with .toml extension
	if filepath.Ext(path) != libmonteur.EXTENSION_TOML || info.IsDir() {
		return nil
	}

	fx.logger.Info("Processing %s...", path)

	// initialize the TOML Program object
	//nolint:lll
	s = &libsetup.Manager{
		Variables: map[string]interface{}{
			libmonteur.VAR_OS:      fx.workspace.OS,
			libmonteur.VAR_ARCH:    fx.workspace.ARCH,
			libmonteur.VAR_COMPUTE: fx.workspace.ComputeSystem,
			libmonteur.VAR_TMP:     fx.workspace.Filesystem.SetupTMPDir,
			libmonteur.VAR_BIN:     fx.workspace.Filesystem.BinDir,
			libmonteur.VAR_CFG:     fx.workspace.Filesystem.BinConfigdDir,
			libmonteur.VAR_LOG:     fx.workspace.Filesystem.WorkspaceLogDir,
			libmonteur.VAR_ROOT:    fx.workspace.Filesystem.RootDir,
			libmonteur.VAR_HOME:    fx.workspace.Filesystem.CurrentDir,
			libmonteur.VAR_SECRETS: fx.secrets,
		},
	}

	fx.logger.Info("Inserting Task Variables...")
	for k, v := range s.Variables {
		if k == libmonteur.VAR_SECRETS {
			fx.logger.Info("\"%s\": %v",
				k,
				libmonteur.LOG_FORMAT_REDACTED,
			)
			continue
		}

		fx.logger.Info("\"%s\": %#v", k, v)
	}
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	fx.logger.Info("Decode Task Data from config file...")
	err = s.Parse(path)
	if err != nil {
		return err //nolint:wrapcheck
	}
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	fx.logger.Info("Register task into job list...")
	fx.workers[s.Metadata.Name] = s
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	return nil
}

func (fx *setup) _init() (err error) {
	fx.settings = &libsetup.Run{}
	fx.workers = map[string]conductor.Job{}
	fx.workspace = &libworkspace.Workspace{}

	// initialize workspace
	err = fx.workspace.Init()
	if err != nil {
		return err //nolint:wrapcheck
	}

	// initialize logger
	fx.logger = &liblog.Logger{
		ToTerminal: true,
	}
	fx.logger.Init()
	fx.workspace.Filesystem.WorkspaceLogDir = filepath.Join(
		fx.workspace.Filesystem.LogDir,
		libmonteur.DIRECTORY_SETUP,
		fx.workspace.Filesystem.WorkspaceLogDir,
	)

	err = fx.logger.Add(liblog.TYPE_STATUS, filepath.Join(
		fx.workspace.Filesystem.WorkspaceLogDir,
		libmonteur.FILE_LOG_PREFIX_JOB+libmonteur.FILE_LOG_STATUS,
	))
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = fx.logger.Add(liblog.TYPE_OUTPUT, filepath.Join(
		fx.workspace.Filesystem.WorkspaceLogDir,
		libmonteur.FILE_LOG_PREFIX_JOB+libmonteur.FILE_LOG_OUTPUT,
	))
	if err != nil {
		return err //nolint:wrapcheck
	}

	fx.logger.Info("\n%s", fx.workspace.String())
	fx.logger.Info("CURRENT CI JOB:\n%s\n", "setup")

	fx.logger.Info("Parsing secrets...")
	fx.secrets = libsecrets.GetSecrets(fx.workspace.Filesystem.SecretsDir)
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	fx.logger.Info("Initialize settings...")
	err = fx.settings.Parse(fx.workspace.Filesystem.SetupTOMLFile)
	if err != nil {
		return err //nolint:wrapcheck
	}
	fx.logger.Success(libmonteur.LOG_SUCCESS)

	return nil
}

func (fx *setup) _reportError(err error) int {
	fx.logger.Error("%s %s\n", libmonteur.ERROR_SETUP, err)
	fx.logger.Sync()
	fx.logger.Close()
	return STATUS_ERROR
}
