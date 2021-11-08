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
	"context"
	"fmt"
	"os"
	"path/filepath"

	//nolint:typecheck
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/endec/toml"

	//nolint:typecheck
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/chmsg"
)

type setup struct {
	workspace *Workspace
	settings  *downloadSettings
	programs  []*binProgram
}

// Run is to execute the setup service function in Monteur.
func (fx *setup) Run() int {
	var err error

	err = fx._init()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return STATUS_ERROR
	}

	err = fx._parseProgramsMetadata()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return STATUS_ERROR
	}

	err = fx._cleanUp()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return STATUS_ERROR
	}

	err = fx._shopPrograms()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return STATUS_ERROR
	}

	return STATUS_OK
}

func (fx *setup) _init() (err error) {
	fx.programs = []*binProgram{}
	fx.workspace = &Workspace{}

	if err := fx.workspace.Init(); err != nil {
		return err
	}

	if err := fx.__parseSettings(); err != nil {
		return err
	}

	return nil
}

func (fx *setup) __parseSettings() (err error) {
	// construct TOML data structure for parsing
	s := struct {
		Downloads *downloadSettings
	}{
		Downloads: fx.settings,
	}

	// parse setting TOML data file
	err = toml.DecodeFile(fx.workspace.filesystem.SetupTOMLFile, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FAILED_CONFIG_DECODE, err)
	}

	return nil
}

func (fx *setup) _parseProgramsMetadata() (err error) {
	err = filepath.Walk(fx.workspace.filesystem.SetupProgramConfigDir,
		fx.__filterProgramMetadata)
	if err != nil {
		return err
	}

	return nil
}

func (fx *setup) __filterProgramMetadata(pathing string,
	info os.FileInfo, err error) error {
	// ensure we only accept the correct file
	if filepath.Ext(pathing) != TOML_EXTENSION {
		return nil
	}

	if info.IsDir() {
		return nil
	}

	// decode the program toml file
	s := &binProgram{}

	err = toml.DecodeFile(pathing, s, nil)
	if err != nil {
		return err
	}

	// set compulsory variables into the data structure
	s.Variables[BIN_PROGRAM_VAR_OS] = fx.workspace.os
	s.Variables[BIN_PROGRAM_VAR_ARCH] = fx.workspace.arch
	s.Variables[BIN_PROGRAM_VAR_COMPUTE] = fx.workspace.computeSystem
	s.Variables[BIN_PROGRAM_VAR_TMP] = fx.workspace.filesystem.SetupTMPDir
	s.Variables[BIN_PROGRAM_VAR_BIN] = fx.workspace.filesystem.BinDir
	s.Variables[BIN_PROGRAM_VAR_CFG] = fx.workspace.filesystem.BinCfgDir

	// sanitize the data structure before acceptance
	err = s.Sanitize()
	if err != nil {
		return err
	}

	fx.programs = append(fx.programs, s)

	return nil
}

func (fx *setup) _cleanUp() (err error) {
	var data []byte

	// remove all
	_ = os.RemoveAll(fx.workspace.filesystem.SetupTMPDir)
	_ = os.RemoveAll(fx.workspace.filesystem.BinCfgDir)
	_ = os.RemoveAll(fx.workspace.filesystem.BinDir)

	// create all
	err = os.MkdirAll(fx.workspace.filesystem.SetupTMPDir,
		SETUP_DIRECTORY_PERMISSION)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_SETUP_DIR_CREATE_FAILED, err)
	}

	err = os.MkdirAll(fx.workspace.filesystem.BinDir,
		SETUP_DIRECTORY_PERMISSION)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_SETUP_DIR_CREATE_FAILED, err)
	}

	err = os.MkdirAll(fx.workspace.filesystem.BinCfgDir,
		SETUP_DIRECTORY_PERMISSION)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_SETUP_DIR_CREATE_FAILED, err)
	}

	// create config
	switch {
	case fx.workspace.os == "windows":
		data = []byte(`
`)
	default:
		data = []byte(`#!/bin/sh
export LOCAL_BIN="` + fx.workspace.filesystem.BinDir + `"
config_dir="` + fx.workspace.filesystem.BinCfgDir + `"

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
	for cfg in "$config_dir"/*; do
		source $cfg
	done
	export PATH="${PATH}:$LOCAL_BIN"
esac`)
	}

	err = os.WriteFile(fx.workspace.filesystem.BinConfigFile,
		data,
		SETUP_CONFIG_PERMISSION,
	)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_SETUP_DIR_CREATE_FAILED, err)
	}

	return nil
}

func (fx *setup) _shopPrograms() (err error) {
	var ok bool
	var rmsg interface{}
	var status string
	var msg chmsg.Message
	var ctx context.Context
	var cancel func()
	var rx chan chmsg.Message
	var isDone bool
	var count uint

	ctx, cancel = context.WithCancel(context.Background())
	rx = make(chan chmsg.Message, len(fx.programs)*2)

	// initiate all programs to get the program from its source
	for _, v := range fx.programs {
		go v.Get(ctx, rx)
		count += 1
	}

	// process all statuses
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok = <-rx:
			if !ok {
				return // rx channel closed
			}

			rmsg, ok = msg.Get(chmsg_DONE)
			if ok {
				isDone = rmsg.(bool)
				if !isDone {
					continue
				}

				count -= 1

				if count == 0 {
					return nil
				}
			}

			rmsg, ok = msg.Get(chmsg_ERROR)
			if ok {
				err = rmsg.(error)
				if err != nil {
					cancel()
					return err
				}
			}

			rmsg, ok = msg.Get(chmsg_STATUS)
			if ok {
				status = rmsg.(string)
				fmt.Printf("%s\n", status)
			}
		}
	}

	return nil
}
