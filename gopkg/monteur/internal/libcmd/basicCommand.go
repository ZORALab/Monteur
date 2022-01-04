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
	"context"
	"fmt"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libtemplater"
)

type basicCMD struct {
	reportUp chan conductor.Message

	thisSystem string

	variables map[string]interface{}
	metadata  *libmonteur.TOMLMetadata

	log *liblog.Logger
	cmd []*libmonteur.TOMLAction
}

// Parse is to parse the given data filepath into basicCMD data type.
func (me *basicCMD) Parse(path string,
	secrets *libsecrets.Secrets) (err error) {
	// initialize raw input variables
	dependencies := []*commander.Dependency{}
	dep := []*libmonteur.TOMLDependency{}
	fmtVar := map[string]interface{}{}
	cmd := []*libmonteur.TOMLAction{}

	// initialize all important variables
	me.metadata = &libmonteur.TOMLMetadata{}
	me.cmd = []*libmonteur.TOMLAction{}

	// construct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
		Dependencies *[]*libmonteur.TOMLDependency
		CMD          *[]*libmonteur.TOMLAction
	}{
		Metadata:     me.metadata,
		Variables:    me.variables,
		FMTVariables: &fmtVar,
		Dependencies: &dep,
		CMD:          &cmd,
	}

	// decode
	err = toml.DecodeFile(path, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// sanitize
	err = sanitizeMetadata(me.metadata, path)
	if err != nil {
		return err
	}

	err = libtemplater.TemplateVariables(&me.variables, &fmtVar)
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = sanitizeDeps(dep, &dependencies, me.thisSystem, me.variables)
	if err != nil {
		return err
	}

	err = checkDependencies(&dependencies)
	if err != nil {
		return err
	}

	err = sanitizeCMD(cmd, &me.cmd, me.thisSystem)
	if err != nil {
		return err
	}

	// init
	err = initializeLogger(&me.log, me.metadata.Name, me.variables, secrets)
	if err != nil {
		return err
	}

	return nil
}

// Run is the universal interface for Manager to execute its run.
func (me *basicCMD) Run(ctx context.Context, ch chan conductor.Message) {
	var err error

	me.log.Info(libmonteur.LOG_JOB_START + "\n\n")
	me.reportUp = ch

	task := &executive{
		log:       me.log,
		variables: me.variables,
		orders:    me.cmd,
		fxSTDOUT:  me.reportOutput,
		fxSTDERR:  me.reportStatus,
	}

	err = task.Exec()
	if err != nil {
		me.reportError("%s", err)
		return
	}

	me.reportDone()
}

// Name is to return the task name
func (me *basicCMD) Name() string {
	return me.metadata.Name
}

func (me *basicCMD) reportStatus(format string, args ...interface{}) {
	reportStatus(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *basicCMD) reportError(format string, args ...interface{}) {
	reportError(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *basicCMD) reportOutput(format string, args ...interface{}) {
	reportOutput(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *basicCMD) reportDone() {
	reportDone(me.log, me.reportUp, me.metadata.Name)
}
