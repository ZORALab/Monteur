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

	ok, err = libmonteur.AcceptTOML(path, info, err)
	if !ok {
		return err //nolint:wrapcheck
	}

	api.logger.Info("Processing %s...", path)
	s = &libcmd.Manager{
		Job:       api.workspace.Job,
		Variables: map[string]interface{}{},
	}

	for k, v := range *api.workspace.Variables {
		s.Variables[k] = v
	}

	_logVariables(api.logger, &s.Variables)

	api.logger.Info("Decode Task Data from config file...")
	err = s.Parse(path)
	if err != nil {
		return err //nolint:wrapcheck
	}
	api.logger.Info(libmonteur.LOG_SUCCESS)

	api.logger.Info("Register task into job list...")
	api.workers[s.Name()] = s
	api.logger.Info(libmonteur.LOG_SUCCESS)

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

	api.logger.Info("Initialize settings...")
	api.settings = &libcmd.Run{}

	err = api.settings.Parse(api.workspace.JobTOMLFile,
		api.workspace.Variables,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}
	api.logger.Info(libmonteur.LOG_SUCCESS)

	return nil
}
