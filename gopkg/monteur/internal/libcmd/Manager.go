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
	thisSystem string

	Job       string
	Variables map[string]interface{}
	Metadata  *libmonteur.TOMLMetadata

	reportUp     chan conductor.Message
	log          *liblog.Logger
	dependencies []*commander.Dependency
	cmd          []*libmonteur.TOMLAction
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

	// initialize raw input variables
	dep := []*libmonteur.TOMLDependency{}
	fmtVar := map[string]interface{}{}
	cmd := []*libmonteur.TOMLAction{}

	// parse data file
	switch me.Job {
	case libmonteur.JOB_TEST:
		err = me.parseTOMLCMD(path, &dep, &fmtVar, &cmd)
	case libmonteur.JOB_BUILD:
		err = me.parseTOMLBuildCMD(path, &dep, &fmtVar, &cmd)
	case libmonteur.JOB_COMPOSE:
		err = me.parseTOMLCMD(path, &dep, &fmtVar, &cmd)
	case libmonteur.JOB_PUBLISH:
		err = me.parseTOMLCMD(path, &dep, &fmtVar, &cmd)
	default:
		panic("MONTEUR DEV: What kind of job is this? ➤ " + me.Job)
	}

	if err != nil {
		return err
	}

	// sanitize
	err = me.sanitizeMetadata(path)
	if err != nil {
		return err
	}

	err = libmonteur.SanitizeVariables(&me.Variables, &fmtVar)
	if err != nil {
		return err //nolint:wrapcheck
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

func (me *Manager) parseTOMLBuildCMD(path string,
	dep *[]*libmonteur.TOMLDependency,
	fmtVar *map[string]interface{},
	cmd *[]*libmonteur.TOMLAction) (err error) {
	// construct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
		BuildDeps    *[]*libmonteur.TOMLDependency
		BuildCMD     *[]*libmonteur.TOMLAction
	}{
		Metadata:     me.Metadata,
		Variables:    me.Variables,
		FMTVariables: fmtVar,
		BuildDeps:    dep,
		BuildCMD:     cmd,
	}

	// decode
	err = toml.DecodeFile(path, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	return nil
}

func (me *Manager) parseTOMLCMD(path string,
	dep *[]*libmonteur.TOMLDependency,
	fmtVar *map[string]interface{},
	cmd *[]*libmonteur.TOMLAction) (err error) {
	// construct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
		Dependencies *[]*libmonteur.TOMLDependency
		CMD          *[]*libmonteur.TOMLAction
	}{
		Metadata:     me.Metadata,
		Variables:    me.Variables,
		FMTVariables: fmtVar,
		Dependencies: dep,
		CMD:          cmd,
	}

	// decode
	err = toml.DecodeFile(path, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
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

func (me *Manager) sanitizeMetadata(path string) (err error) {
	err = me.Metadata.Sanitize(path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}

func (me *Manager) _saveFx(key string, variable, output interface{}) {
	cmd, ok := variable.(*libmonteur.TOMLAction)
	if !ok {
		panic("MONTEUR DEV: why is TOMLAction owner not assigned here?")
	}

	// process commander output
	switch v := output.(type) {
	case *commander.ExecOutput:
		me.__saveExecOutput(key, v, cmd)
	case commander.ExecOutput:
		me.__saveExecOutput(key, &v, cmd)
	case string:
		me.__saveString(key, v, cmd)
	case *string:
		me.__saveString(key, *v, cmd)
	default:
		me.__save(key, v)
	}
}

func (me *Manager) __save(key string, v interface{}) {
	me.log.Info("Reading output...")
	if v == nil {
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, "nil\n")
	} else {
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, v)
	}

	if key == libmonteur.COMMAND_SAVE_NONE {
		goto completed
	}

	me.log.Info("Saving '%v' to '%s'...", v, key)
	me.Variables[key] = v

completed:
	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
}

func (me *Manager) __saveString(key string,
	v string,
	cmd *libmonteur.TOMLAction) {
	me.log.Info("Reading output...")
	me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, v)

	if key == libmonteur.COMMAND_SAVE_NONE {
		goto completed
	}

	v = cmd.ParseExec(v)

	me.log.Info("Saving '%v' to '%s'...", v, key)
	me.Variables[key] = v

completed:
	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
}

func (me *Manager) __saveExecOutput(key string,
	v *commander.ExecOutput,
	cmd *libmonteur.TOMLAction) {
	var val string

	me.log.Info("Reading STDERR...")
	me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stderr))
	me.log.Info("Reading STDOUT...")
	me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stdout))

	// save
	if key == libmonteur.COMMAND_SAVE_NONE {
		me.log.Info(libmonteur.LOG_SUCCESS)
		goto completed
	}

	if cmd.SaveStderr {
		val = cmd.ParseExec(string(v.Stderr))
		me.log.Info("Requested to save STDERR instead...")
	} else {
		val = cmd.ParseExec(string(v.Stdout))
	}

	me.log.Info("Saving '%v' to '%s'...", val, key)
	me.Variables[key] = val

completed:
	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
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
	me.log.Info("Run Task Now: " + libmonteur.LOG_SUCCESS + "\n")

	for i, order := range me.cmd {
		me.log.Info("Executing Command...")
		x := me.createCMD(order)

		err = me.processCMDLocation(order, x)
		if err != nil {
			return
		}

		err = me.processCMDSource(order, x)
		if err != nil {
			return
		}

		err = me.processCMDTarget(order, x)
		if err != nil {
			return
		}

		me.processCMDSave(order, x)

		err = me.initCMD(x)
		if err != nil {
			return
		}

		err = me.runCMD(i+1, x)
		if err != nil {
			return
		}

		err = me.formatToSTDERR(order)
		if err != nil {
			return
		}

		err = me.formatToSTDOUT(order)
		if err != nil {
			return
		}

		me.log.Info("Execute Command ➤ DONE\n\n")
	}

	me.reportDone()
}

func (me *Manager) formatToSTDOUT(order *libmonteur.TOMLAction) (err error) {
	me.log.Info("Formatting order.ToSTDOUT...")

	order.ToSTDOUT, err = templater.String(order.ToSTDOUT,
		me.Variables)
	if err != nil {
		me.reportError("%s: %s",
			libmonteur.ERROR_COMMAND_FMT_BAD,
			err,
		)

		return err //nolint:wrapcheck
	}
	me.log.Info("Got: '%s'", order.ToSTDOUT)
	if order.ToSTDOUT != "" {
		me.reportOutput(order.ToSTDOUT)
	}

	return nil
}

func (me *Manager) formatToSTDERR(order *libmonteur.TOMLAction) (err error) {
	me.log.Info("Formatting order.ToSTDERR...")

	order.ToSTDERR, err = templater.String(order.ToSTDERR,
		me.Variables)
	if err != nil {
		me.reportError("%s: %s",
			libmonteur.ERROR_COMMAND_FMT_BAD,
			err,
		)

		return
	}

	me.log.Info("Got: '%s'", order.ToSTDERR)
	if order.ToSTDERR != "" {
		me.reportStatus(order.ToSTDERR)
	}

	return nil
}

func (me *Manager) runCMD(i int, x *commander.Action) (err error) {
	me.log.Info("Running cmd...")

	err = x.Run()
	if err != nil {
		me.reportError("%s: (Step %d) %s",
			libmonteur.ERROR_COMMAND_FAILED,
			i,
			err,
		)

		return err //nolint:wrapcheck
	}

	return nil
}

func (me *Manager) initCMD(x *commander.Action) (err error) {
	me.log.Info("Initialize cmd...")

	err = x.Init()
	if err != nil {
		me.reportError("%s: %s",
			libmonteur.ERROR_COMMAND_BAD,
			err,
		)

		return err //nolint:wrapcheck
	}

	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
	return nil
}

func (me *Manager) processCMDSave(order *libmonteur.TOMLAction,
	x *commander.Action) {
	me.log.Info("Processing cmd.Save...")

	if order.Save == "" {
		x.Save = libmonteur.COMMAND_SAVE_NONE
	} else {
		x.Save = order.Save
	}

	me.log.Info("Got: '%s'", x.Save)

	// always input x.SaveFx
	me.log.Info("Processing cmd.SaveFx...")
	x.SaveFx = me._saveFx
	me.log.Info("Got: '%s'", x.SaveFx)

	// always input x.SaveVar as the order for _saveFx operation or panic
	me.log.Info("Processing cmd.SaveVar...")
	x.SaveVar = order
	me.log.Info("Got: '%s'", x.SaveVar)
}

func (me *Manager) processCMDTarget(order *libmonteur.TOMLAction,
	x *commander.Action) (err error) {
	me.log.Info("Formatting cmd.Target...")

	order.Target, err = templater.String(order.Target, me.Variables)
	if err != nil {
		me.reportError("%s: %s",
			libmonteur.ERROR_COMMAND_FMT_BAD,
			err,
		)

		return err //nolint:wrapcheck
	}
	x.Target = order.Target

	me.log.Info("Got: '%s'", x.Target)
	return nil
}

func (me *Manager) processCMDSource(order *libmonteur.TOMLAction,
	x *commander.Action) (err error) {
	me.log.Info("Formatting cmd.Source...")

	order.Source, err = templater.String(order.Source, me.Variables)
	if err != nil {
		me.reportError("%s: %s",
			libmonteur.ERROR_COMMAND_FMT_BAD,
			err,
		)

		return err //nolint:wrapcheck
	}
	x.Source = order.Source

	me.log.Info("Got: '%s'", x.Source)
	return nil
}

func (me *Manager) processCMDLocation(order *libmonteur.TOMLAction,
	x *commander.Action) (err error) {
	me.log.Info("Formatting cmd.Location...")

	order.Location, err = templater.String(order.Location,
		me.Variables,
	)
	if err != nil {
		me.reportError("%s: %s",
			libmonteur.ERROR_COMMAND_FMT_BAD,
			err,
		)

		return err //nolint:wrapcheck
	}
	x.Location = order.Location

	me.log.Info("Got: '%s'", x.Location)
	return nil
}

func (me *Manager) createCMD(cmd *libmonteur.TOMLAction) (x *commander.Action) {
	x = &commander.Action{
		Name: cmd.Name,
		Type: cmd.Type,
	}

	me.log.Info("Name: '%s'", x.Name)
	me.log.Info("Type: '%v'", x.Type)
	return x
}

func (me *Manager) reportError(format string, args ...interface{}) {
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

func (me *Manager) reportOutput(format string, args ...interface{}) {
	if me.log != nil {
		me.log.Output(format, args...)
	}

	if me.reportUp != nil {
		me.reportUp <- conductor.CreateOutput(me.Metadata.Name,
			format,
			args...,
		)
	}
}

func (me *Manager) reportStatus(format string, args ...interface{}) {
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
