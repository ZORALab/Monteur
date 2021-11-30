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

package libsetup

import (
	"fmt"
	"os"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

type Run struct {
	Limit uint64
}

func (me *Run) Parse(path string) (err error) {
	s := struct {
		Downloads *Run
	}{
		Downloads: me,
	}

	err = toml.DecodeFile(path, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	return nil
}

func (me *Run) CleanDir(path string) (err error) {
	// delete
	_ = os.RemoveAll(path)

	// create
	err = os.MkdirAll(path, libmonteur.PERMISSION_DIRECTORY)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_DIR_CREATE_FAILED,
			err,
		)
	}

	return nil
}

func (me *Run) SetupConfig(currentOS string,
	path string, binDir string, configDir string) (err error) {
	var data []byte

	switch currentOS {
	case "windows":
		return nil // yet to support
	default:
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
	}

	err = os.WriteFile(path, data, libmonteur.PERMISSION_CONFIG)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_FAILED,
			err,
		)
	}

	return nil
}
