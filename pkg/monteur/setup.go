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

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libsetup"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libworkspace"
)

type setup struct {
	workspace *libworkspace.Workspace
	settings  *libsetup.Run
	programs  map[string]*libsetup.Program
	secrets   map[string]interface{}
}

func (fx *setup) _reportError(err error) int {
	fmt.Fprintf(os.Stdout, "%s %s\n", libmonteur.ERROR_SETUP, err)
	return STATUS_ERROR
}

// Run is to execute the setup service function in Monteur.
func (fx *setup) Run() int {
	var err error

	err = fx._init()
	if err != nil {
		return fx._reportError(err)
	}

	err = fx._parseProgramsMetadata()
	if err != nil {
		return fx._reportError(err)
	}

	err = fx._cleanUp()
	if err != nil {
		return fx._reportError(err)
	}

	err = fx._shopPrograms()
	if err != nil {
		return fx._reportError(err)
	}

	return STATUS_OK
}

func (fx *setup) _init() (err error) {
	fx.settings = &libsetup.Run{}
	fx.programs = map[string]*libsetup.Program{}
	fx.workspace = &libworkspace.Workspace{}

	// initialize workspace
	err = fx.workspace.Init()
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = fx.settings.Parse(fx.workspace.Filesystem.SetupTOMLFile)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// initialize secrets and parse every one of them
	fx.secrets = libsecrets.GetSecrets(fx.workspace.Filesystem.SecretsDir)

	return nil
}

func (fx *setup) _parseProgramsMetadata() (err error) {
	err = filepath.Walk(fx.workspace.Filesystem.SetupProgramConfigDir,
		fx.__filterProgramMetadata)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}

func (fx *setup) __filterProgramMetadata(pathing string,
	info os.FileInfo, err error) error {
	var s *libsetup.TOMLProgram
	var app *libsetup.Program

	// return err if occurred
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// ensure we only accept the correct regular file with .toml extension
	if filepath.Ext(pathing) != libmonteur.EXTENSION_TOML || info.IsDir() {
		return nil
	}

	// initialize the TOML Program object
	//nolint:lll
	s = &libsetup.TOMLProgram{
		Variables: map[string]interface{}{
			libmonteur.VAR_OS:      fx.workspace.OS,
			libmonteur.VAR_ARCH:    fx.workspace.ARCH,
			libmonteur.VAR_COMPUTE: fx.workspace.ComputeSystem,
			libmonteur.VAR_TMP:     fx.workspace.Filesystem.SetupTMPDir,
			libmonteur.VAR_BIN:     fx.workspace.Filesystem.BinDir,
			libmonteur.VAR_CFG:     fx.workspace.Filesystem.BinCfgDir,
			libmonteur.VAR_ROOT:    fx.workspace.Filesystem.RootDir,
			libmonteur.VAR_HOME:    fx.workspace.Filesystem.CurrentDir,
			libmonteur.VAR_SECRETS: fx.secrets,
		},
	}

	// decode the program's toml file
	err = s.Parse(pathing)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// process the data and generate the program object for operation
	app, err = s.Process()
	if err != nil {
		return err //nolint:wrapcheck
	}

	fx.programs[app.Metadata.Name] = app

	return nil
}

func (fx *setup) _cleanUp() (err error) {
	var data []byte

	// remove all
	_ = os.RemoveAll(fx.workspace.Filesystem.SetupTMPDir)
	_ = os.RemoveAll(fx.workspace.Filesystem.BinCfgDir)
	_ = os.RemoveAll(fx.workspace.Filesystem.BinDir)

	// create all
	err = os.MkdirAll(fx.workspace.Filesystem.SetupTMPDir,
		libmonteur.PERMISSION_DIRECTORY)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_CREATE_FAILED,
			err,
		)
	}

	err = os.MkdirAll(fx.workspace.Filesystem.BinDir,
		libmonteur.PERMISSION_DIRECTORY)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_CREATE_FAILED,
			err,
		)
	}

	err = os.MkdirAll(fx.workspace.Filesystem.BinCfgDir,
		libmonteur.PERMISSION_DIRECTORY)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_CREATE_FAILED,
			err,
		)
	}

	// create config
	switch {
	case fx.workspace.OS == "windows":
		data = []byte(``)
	default:
		data = []byte(`#!/bin/sh
export LOCAL_BIN="` + fx.workspace.Filesystem.BinDir + `"
config_dir="` + fx.workspace.Filesystem.BinCfgDir + `"

stop() {
	PATH=:${PATH}:
	PATH=${PATH//:$LOCAL_BIN:/:}

	for cfg in "$config_dir"/*; do
		source "$cfg" --stop
	done
}

case $1 in
--stop)
	stop
	;;
*)
	export PATH="${PATH}:$LOCAL_BIN"
	for cfg in "$config_dir"/*; do
		source $cfg
	done
esac`)
	}

	err = os.WriteFile(fx.workspace.Filesystem.BinConfigFile,
		data,
		libmonteur.PERMISSION_CONFIG)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_FAILED,
			err,
		)
	}

	return nil
}

func (fx *setup) _shopPrograms() (err error) {
	conductor := &libsetup.Conductor{
		Runners: fx.programs,
	}

	conductor.Run()
	return conductor.Coordinate() //nolint:wrapcheck
}
