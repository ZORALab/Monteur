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
