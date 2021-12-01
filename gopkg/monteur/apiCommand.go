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
	"os"
	"path/filepath"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libcmd"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libworkspace"
)

type apiCommand struct {
	secrets   map[string]interface{}
	workers   map[string]conductor.Job
	workspace *libworkspace.Workspace
	settings  *libcmd.Run
	logger    *liblog.Logger

	Job      string
	ErrorTag string
}

// Run is to execute the apiCommand algorithm.
func (api *apiCommand) Run() (statusCode int) {
	err := api._init()
	if err != nil {
		return _reportError(api.logger, api.ErrorTag, err)
	}

	// parse all apiCommands
	err = filepath.Walk(api.workspace.ConfigDir, api._filter)
	if err != nil {
		return _reportError(api.logger, api.ErrorTag, err)
	}

	// execute each task in parallel
	c := &conductor.Conductor{
		Runners: api.workers,
		Log:     api.logger,
	}

	err = c.Run()
	if err != nil {
		return _reportError(api.logger, api.ErrorTag, err)
	}

	err = c.Coordinate()
	if err != nil {
		return _reportError(api.logger, api.ErrorTag, err)
	}

	// safely close the logs and exit as completion
	api.logger.Sync()
	api.logger.Close()
	return STATUS_OK
}

func (api *apiCommand) _filter(path string, info os.FileInfo, err error) error {
	var ok bool
	var s *libcmd.Manager

	ok, err = _acceptanceFilter(path, info, err)
	if !ok {
		return err
	}

	s, err = _createCMDManager(api.logger,
		api.workspace,
		&api.secrets,
		path,
	)
	if err != nil {
		return err
	}

	api.logger.Info("Register task into job list...")
	api.workers[s.Metadata.Name] = s
	api.logger.Success(libmonteur.LOG_SUCCESS)

	return nil
}

func (api *apiCommand) _init() (err error) {
	api.workers = map[string]conductor.Job{}

	err = _initWorkspace(api.Job, &api.workspace)
	if err != nil {
		return err
	}

	err = _initLogger(&api.logger, api.workspace)
	if err != nil {
		return err
	}

	_initSecrets(&api.secrets,
		api.logger,
		api.workspace.Filesystem.SecretsDir,
	)

	err = _initCMDSettings(&api.settings,
		api.logger,
		api.workspace.JobTOMLFile,
	)
	if err != nil {
		return err
	}

	return nil
}
