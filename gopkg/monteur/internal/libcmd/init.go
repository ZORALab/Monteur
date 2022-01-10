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

package libcmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsecrets"
)

func initializeLogger(logger **liblog.Logger,
	name string,
	variables map[string]interface{},
	secrets *libsecrets.Secrets) (err error) {
	var sRet string
	var ok bool

	// gather inputs
	sRet, ok = variables[libmonteur.VAR_LOG].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_LOG before Parse()!")
	}

	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ReplaceAll(name, "+", "-")
	name = strings.ReplaceAll(name, "!", "")
	name = strings.ReplaceAll(name, "$", "")

	// initialize logger
	*logger = &liblog.Logger{}
	(*logger).Init(secrets)

	err = (*logger).Add(liblog.TYPE_STATUS, filepath.Join(
		sRet,
		name+"-"+libmonteur.FILE_LOG_STATUS,
	))
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = (*logger).Add(liblog.TYPE_OUTPUT, filepath.Join(
		sRet,
		name+"-"+libmonteur.FILE_LOG_OUTPUT,
	))
	if err != nil {
		(*logger).Close()
		return err //nolint:wrapcheck
	}

	(*logger).Info(libmonteur.LOG_JOB_INIT_SUCCESS)

	return nil
}

func initializeMonteurFS(variables map[string]interface{}) (err error) {
	var list []string
	var data []byte
	var thisOS, configPath, configDir, binDir string
	var ok bool

	// create list of directories
	list = []string{}

	// process critical directories
	binDir, ok = variables[libmonteur.VAR_BIN].(string)
	if !ok {
		panic("MONTEUR_DEV: why is VAR_BIN missing?")
	}
	list = append(list, binDir)

	configDir, ok = variables[libmonteur.VAR_CFG].(string)
	if !ok {
		panic("MONTEUR_DEV: why is VAR_CFG missing?")
	}
	configPath = filepath.Join(configDir,
		libmonteur.FILENAME_BIN_CONFIG_MAIN,
	)
	configDir = filepath.Join(configDir,
		libmonteur.DIRECTORY_MONTEUR_CONFIG_D,
	)
	list = append(list, configDir)

	for _, path := range list {
		err = os.MkdirAll(path, libmonteur.PERMISSION_DIRECTORY)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_DIR_CREATE_FAILED,
				err,
			)
		}
	}

	// process thisOS
	thisOS, ok = variables[libmonteur.VAR_OS].(string)
	if !ok || thisOS == "" {
		panic("MONTEUR_DEV: why is VAR_OS missing?")
	}

	switch thisOS {
	case "linux",
		"freebsd",
		"openbsd",
		"plan9",
		"dragonfly",
		"android",
		"netbsd",
		"solaris",
		"darwin":
		data = []byte(`#!/bin/sh
export LOCAL_BIN="` + binDir + `"
config_dir="` + configDir + `"

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
	default:
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_OS_UNSUPPORTED,
			thisOS,
		)
	}

	// remove previous file regardlessly
	_ = os.RemoveAll(configPath)

	// create file
	err = os.WriteFile(configPath, data, libmonteur.PERMISSION_CONFIG)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_FAILED,
			err,
		)
	}

	return nil
}
