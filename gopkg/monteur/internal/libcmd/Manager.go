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
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

type Manager struct {
	Metadata *libmonteur.TOMLMetadata

	thisSystem string
	Variables  map[string]interface{}

	dependencies []*commander.Dependency
	cmd          []*libmonteur.TOMLAction

	reportUp chan conductor.Message

	log *liblog.Logger
}

func (me *Manager) Parse(path string) (err error) {
	var ok bool

	// initialize all important variables
	me.Metadata = &libmonteur.TOMLMetadata{}
	me.dependencies = []*commander.Dependency{}
	me.cmd = []*libmonteur.TOMLAction{}

	me.thisSystem, ok = me.Variables[libmonteur.VAR_COMPUTE].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_COMPUTE before Parse()!")
	}

	dep := []*libmonteur.TOMLDependency{}
	fmtVar := map[string]interface{}{}
	cmd := []*libmonteur.TOMLAction{}

	// construct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		Dependencies *[]*libmonteur.TOMLDependency
		CMD          *[]*libmonteur.TOMLAction
		FMTVariables *map[string]interface{}
	}{
		Metadata:     me.Metadata,
		Variables:    me.Variables,
		Dependencies: &dep,
		CMD:          &cmd,
		FMTVariables: &fmtVar,
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
	err = me.sanitizeMetadata(path)
	if err != nil {
		return err
	}

	err = me.sanitizeFMTVariables(fmtVar)
	if err != nil {
		return err
	}

	err = me.sanitizeDeps(dep)
	if err != nil {
		return err
	}

	err = me.sanitizeCMD(cmd)
	if err != nil {
		return err
	}

	// initialize logger
	err = me.initializeLogger()
	if err != nil {
		return err
	}

	return nil
}

func (me *Manager) initializeLogger() (err error) {
	var sRet, name string
	var ok bool

	// initialize logger
	sRet, ok = me.Variables[libmonteur.VAR_LOG].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_LOG before Parse()!")
	}
	me.log = &liblog.Logger{}
	me.log.Init()

	name = strings.ToLower(me.Metadata.Name)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ReplaceAll(name, "+", "-")
	name = strings.ReplaceAll(name, "!", "")
	name = strings.ReplaceAll(name, "$", "")

	err = me.log.Add(liblog.TYPE_STATUS, filepath.Join(
		sRet,
		name+"-"+libmonteur.FILE_LOG_STATUS,
	))
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = me.log.Add(liblog.TYPE_OUTPUT, filepath.Join(
		sRet,
		name+"-"+libmonteur.FILE_LOG_OUTPUT,
	))
	if err != nil {
		me.log.Close()
		return err //nolint:wrapcheck
	}

	me.log.Info(libmonteur.LOG_JOB_INIT_SUCCESS)

	return nil
}

func (me *Manager) sanitizeCMD(in []*libmonteur.TOMLAction) (err error) {
	for _, cmd := range in {
		if !libmonteur.IsComputeSystemSupported(me.thisSystem,
			cmd.Condition) {
			continue
		}

		err = cmd.Sanitize()
		if err != nil {
			return err //nolint:wrapcheck
		}

		me.cmd = append(me.cmd, cmd)
	}

	return nil
}

func (me *Manager) sanitizeDeps(in []*libmonteur.TOMLDependency) (err error) {
	var val string

	// scan conditions for building commands list
	for _, dep := range in {
		if !libmonteur.IsComputeSystemSupported(me.thisSystem,
			[]string{dep.Condition}) {
			continue
		}

		val, err = templater.String(dep.Command, me.Variables)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_COMMAND_DEPENDENCY_FMT_BAD,
				err,
			)
		}

		s := &commander.Dependency{
			Name:    dep.Name,
			Type:    dep.Type,
			Command: val,
		}

		me.dependencies = append(me.dependencies, s)
	}

	// sanitize each commands for validity
	for _, dep := range me.dependencies {
		err = dep.Init()
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_DEPENDENCY_BAD,
				err,
			)
		}
	}

	return nil
}

func (me *Manager) sanitizeFMTVariables(in map[string]interface{}) (err error) {
	var val interface{}

	if in == nil {
		return nil
	}

	for key, value := range in {
		switch v := value.(type) {
		case string:
			val, err = templater.String(v, me.Variables)
		default:
			val = v
		}

		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_VARIABLES_FMT_BAD,
				err,
			)
		}

		me.Variables[key] = val
	}

	return nil
}

func (me *Manager) sanitizeMetadata(path string) (err error) {
	err = me.Metadata.Sanitize(path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}

func (me *Manager) _saveFx(key string, output interface{}) {
	switch v := output.(type) {
	case *commander.ExecOutput:
		me.log.Info("Reading STDERR...")
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stderr))
		me.log.Info("Reading STDOUT...")
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stdout))

		if key != libmonteur.COMMAND_SAVE_NONE {
			val := strings.TrimRight(string(v.Stdout), "\r\n")
			me.Variables[key] = val
			me.log.Info("Saving '%v' to '%s'...", output, key)
		}
	case commander.ExecOutput:
		me.log.Info("Reading STDERR...")
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stderr))
		me.log.Info("Reading STDOUT...")
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stdout))

		if key != libmonteur.COMMAND_SAVE_NONE {
			val := strings.TrimRight(string(v.Stdout), "\r\n")
			me.Variables[key] = val
			me.log.Info("Saving '%v' to '%s'...", output, key)
		}
	default:
		me.log.Info("Reading output...")
		if v == nil {
			me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, "nil\n")
		} else {
			me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, output)
		}

		if key != libmonteur.COMMAND_SAVE_NONE {
			me.Variables[key] = output
			me.log.Info("Saving '%v' to '%s'...", output, key)
		}
	}

	me.log.Info(libmonteur.LOG_SUCCESS)
	return
}

// Name is for generating the program Metadata.Name when used as in interface.
//
// This should only be called after the Manager is initialized successfully.
func (me *Manager) Name() string {
	if me.Metadata == nil {
		return ""
	}

	return me.Metadata.Name
}

// Run is to execute the publisher's commands sequence.
//
// Everything must be setup properly before calling this function. It was meant
// for Monteur's commands-driven API(s).
//
// All errors generated in this method shall use `me.reportError` instead of
// returning `fmt.Errorf` since it will be executed in parallel with others
// in an asynchonous manner.
//
// This should only be called after the Manager is initialized successfully.
func (me *Manager) Run(ctx context.Context, ch chan conductor.Message) {
	var err error

	me.reportUp = ch
	me.log.Success(libmonteur.LOG_SUCCESS)

	for i, order := range me.cmd {
		x := &commander.Action{
			Name:   order.Name,
			Save:   order.Save,
			SaveFx: me._saveFx,
			Type:   order.Type,
		}

		me.log.Info("Executing Command...")
		me.log.Info("Name: '%s'", x.Name)
		me.log.Info("SaveFx: '%v'", x.SaveFx)
		me.log.Info("Type: '%v'", x.Type)

		me.log.Info("Formatting cmd.Location...")
		order.Location, err = templater.String(order.Location,
			me.Variables,
		)
		if err != nil {
			me.reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)

			return
		}
		x.Location = order.Location
		me.log.Info("Got: '%s'", x.Location)

		me.log.Info("Formatting cmd.Source...")
		order.Source, err = templater.String(order.Source,
			me.Variables)
		if err != nil {
			me.reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)

			return
		}
		x.Source = order.Source
		me.log.Info("Got: '%s'", x.Source)

		me.log.Info("Formatting cmd.Target...")
		order.Target, err = templater.String(order.Target, me.Variables)
		if err != nil {
			me.reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)

			return
		}
		x.Target = order.Target
		me.log.Info("Got: '%s'", x.Target)

		me.log.Info("Processing cmd.Save...")
		if x.Save == "" {
			x.Save = libmonteur.COMMAND_SAVE_NONE
		}
		me.log.Info("Got: '%s'", x.Save)

		me.log.Info("Initialize cmd...")
		err = x.Init()
		if err != nil {
			me.reportError("%s: %s",
				libmonteur.ERROR_COMMAND_BAD,
				err,
			)
		}
		me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))

		me.log.Info("Running cmd...")

		err = x.Run()
		if err != nil {
			me.reportError("%s: (Step %d) %s",
				libmonteur.ERROR_COMMAND_FAILED,
				i+1,
				err,
			)

			return
		}
	}

	me.reportDone()
	return
}

func (me *Manager) reportError(format string, args ...interface{}) {
	if me.Metadata == nil || me.Metadata.Name == "" {
		format = "Task '' ➤ " + format
	} else {
		format = "Task '%s' ➤ " + format
		args = append([]interface{}{me.Metadata.Name}, args...)
	}

	if me.log != nil {
		me.log.Error(format, args...)
		me.log.Sync()
		me.log.Close()
	}

	if me.reportUp != nil {
		me.reportUp <- conductor.CreateError(me.Metadata.Name,
			format,
			args...,
		)
	}
}

func (me *Manager) reportStatus(format string, args ...interface{}) {
	if me.Metadata == nil || me.Metadata.Name == "" {
		format = "Task '' ➤ " + format
	} else {
		format = "Task '%s' ➤ " + format
		args = append([]interface{}{me.Metadata.Name}, args...)
	}

	if me.log != nil {
		me.log.Info(format, args...)
	}

	if me.reportUp != nil {
		me.reportUp <- conductor.CreateStatus(me.Metadata.Name,
			format,
			args...,
		)
	}
}

func (me *Manager) reportDone() {
	if me.log != nil {
		me.log.Success(libmonteur.LOG_SUCCESS)
		me.log.Sync()
		me.log.Close()
	}

	if me.reportUp != nil {
		me.reportUp <- conductor.CreateDone(me.Metadata.Name)
	}
}
