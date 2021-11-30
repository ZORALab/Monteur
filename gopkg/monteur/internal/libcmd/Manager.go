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
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

type Manager struct {
	Metadata     *libmonteur.TOMLMetadata
	thisSystem   string
	Variables    map[string]interface{}
	Dependencies []*commander.Dependency
	CMD          []*commander.Action

	log *liblog.Logger
}

func (me *Manager) Parse(path string) (err error) {
	var fmtVar map[string]interface{}
	var dep []*libmonteur.TOMLDependency
	var cmd []*libmonteur.TOMLAction
	var ok bool

	// initialize all important variables
	me.Metadata = &libmonteur.TOMLMetadata{}
	me.Dependencies = []*commander.Dependency{}
	me.CMD = []*commander.Action{}

	me.thisSystem, ok = me.Variables[libmonteur.VAR_COMPUTE].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_COMPUTE before Parse()!")
	}

	dep = []*libmonteur.TOMLDependency{}
	fmtVar = map[string]interface{}{}

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
	err = me.Metadata.Sanitize(path)
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

func (me *Manager) sanitizeDeps(in []*libmonteur.TOMLDependency) (err error) {
	var val string

	// initialize all variables
	me.Dependencies = []*commander.Dependency{}

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

		me.Dependencies = append(me.Dependencies, s)
	}

	// sanitize each commands for validity
	for _, dep := range me.Dependencies {
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

func (me *Manager) sanitizeCMD(in []*libmonteur.TOMLAction) (err error) {
	// initialize all variables
	me.CMD = []*commander.Action{}

	// scan conditions for building commands list
	for _, cmd := range in {
		if !libmonteur.IsComputeSystemSupported(me.thisSystem,
			cmd.Condition) {
			continue
		}

		a := &commander.Action{
			Name:     cmd.Name,
			Type:     cmd.Type,
			Location: cmd.Location,
			Source:   cmd.Source,
			Target:   cmd.Target,
			Save:     cmd.Save,
			SaveFx:   me._saveFx,
		}

		me.CMD = append(me.CMD, a)
	}

	// sanitize each of them
	for i, cmd := range me.CMD {
		err = cmd.Init()
		if err != nil {
			return fmt.Errorf("%s (CMD %d) %s",
				libmonteur.ERROR_COMMAND_BAD,
				i+1,
				err,
			)
		}
	}

	return nil
}

func (me *Manager) sanitizeMetadata(path string) (err error) {
	if me.Metadata == nil {
		return me.__reportError("%s: %s",
			libmonteur.ERROR_PUBLISH_METADATA_MISSING,
			path,
		)
	}

	if me.Metadata.Name == "" {
		return me.__reportError("%s: '%s' for %s",
			libmonteur.ERROR_PUBLISH_METADATA_MISSING,
			"Name",
			path,
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

	me.log.Info("Task initialized successfully. Standing By...")

	return nil
}

func (me *Manager) _saveFx(key string, output interface{}) (err error) {
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

	me.log.Success(libmonteur.LOG_SUCCESS)
	return nil
}

// Run is to execute the publisher's commands sequence.
func (me *Manager) Run() (err error) {
	me.log.Success(libmonteur.LOG_SUCCESS)

	for i, cmd := range me.CMD {
		me.log.Info("Executing Command...")
		me.log.Info("Name: '%s'", cmd.Name)
		me.log.Info("Save: '%s'", cmd.Save)
		me.log.Info("SaveFx: '%v'", cmd.SaveFx)
		me.log.Info("Type: '%v'", cmd.Type)

		me.log.Info("Formatting cmd.Location...")
		cmd.Location, err = templater.String(cmd.Location, me.Variables)
		if err != nil {
			return me.__reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)
		}
		me.log.Info("Got: '%s'", cmd.Location)

		me.log.Info("Formatting cmd.Source...")
		cmd.Source, err = templater.String(cmd.Source, me.Variables)
		if err != nil {
			return me.__reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)
		}
		me.log.Info("Got: '%s'", cmd.Source)

		me.log.Info("Formatting cmd.Target...")
		cmd.Target, err = templater.String(cmd.Target, me.Variables)
		if err != nil {
			return me.__reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)
		}
		me.log.Info("Got: '%s'", cmd.Target)

		me.log.Info("Running cmd...")
		if cmd.Save == "" {
			cmd.Save = libmonteur.COMMAND_SAVE_NONE
		}

		err = cmd.Run()
		if err != nil {
			return me.__reportError("%s: (Step %d) %s",
				libmonteur.ERROR_COMMAND_FAILED,
				i+1,
				err,
			)
		}
	}

	me.log.Sync()
	me.log.Close()
	return nil
}

func (me *Manager) __reportError(format string,
	args ...interface{}) (err error) {
	if me.Metadata == nil || me.Metadata.Name == "" {
		me.log.Error("Task '' ➤ "+format, args...)
		err = fmt.Errorf("Task '' ➤ "+format, args...)
		goto endReporting
	}

	args = append([]interface{}{me.Metadata.Name}, args...)
	me.log.Error("Task '%s' ➤ "+format, args...)
	err = fmt.Errorf("Task '%s' ➤ "+format, args...)

endReporting:
	me.log.Sync()
	me.log.Close()
	return err
}
