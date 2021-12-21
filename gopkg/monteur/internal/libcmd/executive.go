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
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

type executive struct {
	fxSTDOUT  func(string, ...interface{})
	fxSTDERR  func(string, ...interface{})
	variables map[string]interface{}
	log       *liblog.Logger
	orders    []*libmonteur.TOMLAction
}

// Exec instructs the executive to run all the given commands.
func (me *executive) Exec() (err error) {
	for i, order := range me.orders {
		me.log.Info("Executing Command...")

		cmd := me.create(order)
		me.log.Info("Name: '%s'", cmd.Name)
		me.log.Info("Type: '%v'", cmd.Type)

		me.log.Info("Formatting cmd.Location...")
		cmd.Location, err = libmonteur.ProcessString(order.Location,
			me.variables,
		)
		if err != nil {
			return err //nolint:wrapcheck
		}
		me.log.Info("Got: '%s'", cmd.Location)

		me.log.Info("Formatting cmd.Source...")
		cmd.Source, err = libmonteur.ProcessString(order.Source,
			me.variables,
		)
		if err != nil {
			return err //nolint:wrapcheck
		}
		me.log.Info("Got: '%s'", cmd.Source)

		me.log.Info("Formatting cmd.Target...")
		cmd.Target, err = libmonteur.ProcessString(order.Target,
			me.variables,
		)
		if err != nil {
			return err //nolint:wrapcheck
		}
		me.log.Info("Got: '%s'", cmd.Target)

		me.log.Info("Processing cmd.Save...")
		me.processSave(cmd, order)
		me.log.Info("Got cmd.Save   : '%s'", cmd.Save)
		me.log.Info("Got cmd.SaveFx : '%s'", cmd.SaveFx)
		me.log.Info("Got cmd.SaveVar: '%s'", cmd.SaveVar)

		me.log.Info("Initialize cmd...")
		err = me.initCMD(cmd)
		if err != nil {
			return err
		}
		me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))

		me.log.Info("Run cmd...")
		err = me.exec(cmd, i+1)
		if err != nil {
			return err
		}

		me.log.Info("formatting ToSTDOUT...")
		order.ToSTDOUT, err = libmonteur.ProcessString(order.ToSTDOUT,
			me.variables,
		)
		if err != nil {
			return err //nolint:wrapcheck
		}
		me.log.Info("Got: '%s'", order.ToSTDOUT)
		me.report(me.fxSTDOUT, "STDOUT", order.ToSTDOUT)

		me.log.Info("formatting ToSTDERR...")
		order.ToSTDERR, err = libmonteur.ProcessString(order.ToSTDERR,
			me.variables,
		)
		if err != nil {
			return err //nolint:wrapcheck
		}
		me.log.Info("Got: '%s'", order.ToSTDERR)
		me.report(me.fxSTDERR, "STDERR", order.ToSTDERR)

		me.log.Info("Execute Command ➤ DONE\n\n")
	}

	return nil
}

func (me *executive) report(fx func(string, ...interface{}),
	name string,
	data string) {
	if data == "" {
		return
	}

	me.log.Info("reporting %s...", name)
	if fx != nil {
		fx(data)
	}

	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS,
		"\n"))
}

func (me *executive) exec(cmd *commander.Action, step int) (err error) {
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: (Step %d) %s",
			libmonteur.ERROR_COMMAND_FAILED,
			step,
			err,
		)
	}

	return nil
}

func (me *executive) initCMD(cmd *commander.Action) (err error) {
	err = cmd.Init()
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_COMMAND_BAD,
			err,
		)
	}

	return nil
}

func (me *executive) processSave(cmd *commander.Action,
	order *libmonteur.TOMLAction) {
	cmd.Save = libmonteur.COMMAND_SAVE_NONE
	if order.Save != "" {
		cmd.Save = order.Save
	}

	cmd.SaveFx = me.fxSave
	cmd.SaveVar = order
}

func (me *executive) create(cmd *libmonteur.TOMLAction) *commander.Action {
	return &commander.Action{
		Name: cmd.Name,
		Type: cmd.Type,
	}
}

func (me *executive) fxSave(key string, variable, output interface{}) {
	cmd, ok := variable.(*libmonteur.TOMLAction)
	if !ok {
		panic("MONTEUR DEV: why is TOMLAction owner not assigned here?")
	}

	me.log.Info("Reading output...")
	switch v := output.(type) {
	case *commander.ExecOutput:
		me._saveExecOutput(key, v, cmd)
	case commander.ExecOutput:
		me._saveExecOutput(key, &v, cmd)
	case string:
		me._saveString(key, v, cmd)
	case *string:
		me._saveString(key, *v, cmd)
	case bool:
		me._saveBool(key, v)
	case *bool:
		me._saveBool(key, *v)
	default:
		me._save(key, v)
	}
}

func (me *executive) _save(key string, v interface{}) {
	if v == nil {
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, "nil\n")
	} else {
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, v)
	}

	if key == libmonteur.COMMAND_SAVE_NONE {
		goto completed
	}

	me.log.Info("Saving '%v' to '%s'...", v, key)
	me.variables[key] = v

completed:
	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
}

func (me *executive) _saveBool(key string, v bool) {
	data := "false\n"
	if v {
		data = "true\n"
	}
	me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, data)

	if key == libmonteur.COMMAND_SAVE_NONE {
		goto completed
	}

	me.log.Info("Saving '%v' to '%s'...", v, key)
	me.variables[key] = v

completed:
	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
}

func (me *executive) _saveString(key string,
	v string, cmd *libmonteur.TOMLAction) {
	me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, v)

	if key == libmonteur.COMMAND_SAVE_NONE {
		goto completed
	}

	v = cmd.ParseExec(v)

	me.log.Info("Saving '%v' to '%s'...", v, key)
	me.variables[key] = v

completed:
	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
}

func (me *executive) _saveExecOutput(key string,
	v *commander.ExecOutput, cmd *libmonteur.TOMLAction) {
	var val string

	me.log.Info("Reading command.STDERR...")
	me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stderr))
	me.log.Info("Reading command.STDOUT...")
	me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stdout))

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
	me.variables[key] = val

completed:
	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
}
