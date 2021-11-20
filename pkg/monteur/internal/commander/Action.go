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

package commander

import (
	"fmt"
)

type Action struct {
	// Name is for the action naming used in logging and identification
	Name string

	// Location is where the directory shall change to for execution
	Location string

	// Source is the input of the action in general.
	//
	// See 'Type' documentations for the action's specification.
	Source string

	// Target is the output of the action in general.
	//
	// See 'Type' documentations for the action's specification.
	Target string

	// Save is to save the output into variables.
	//
	// The value shall be the name (or 'key') of the variable.
	Save string

	// SaveFx is the function to operate value storing for `Save` key.
	//
	// This function **MUST** be set if `Save` is set.
	SaveFx func(key string, output interface{}) (err error)

	actionFx func(action *Action) (output interface{}, err error)

	// Type is the action type ID.
	Type ActionID
}

// Init is a method to ensure Action is sanitized and ready for execution.
//
// It validates all known configurations before executing the commands.
func (action *Action) Init() (err error) {
	if action.Name == "" {
		return fmt.Errorf("action's Name is empty")
	}

	err = action._initMeta()
	if err != nil {
		return err
	}

	err = action._initType()
	if err != nil {
		return err
	}

	return nil
}

func (action *Action) _initMeta() (err error) {
	if action.Location == "" {
		return action.__reportError("Location is empty")
	}

	if action.Type == "" {
		return action.__reportError("Type is empty")
	}

	if action.Source == "" {
		return action.__reportError("Source is empty")
	}

	if action.Save != "" && action.SaveFx == nil {
		return action.__reportError("%s: '%s'",
			"SaveFx is missing for save",
			action.Save,
		)
	}

	return nil
}

func (action *Action) _initType() (err error) {
	switch action.Type {
	case ACTION_PLACEHOLDER:
		action.actionFx = cmdPlaceholder
	case ACTION_COMMAND:
		action.actionFx = cmdPlaceholder
	case ACTION_COMMAND_QUIET:
		action.actionFx = cmdPlaceholder
	case ACTION_COPY:
		action.actionFx = cmdPlaceholder
	case ACTION_COPY_RECURSIVE:
		action.actionFx = cmdPlaceholder
	case ACTION_COPY_RECURSIVE_QUIET:
		action.actionFx = cmdPlaceholder
	case ACTION_COPY_QUIET:
		action.actionFx = cmdPlaceholder
	case ACTION_CREATE_DIR:
		action.actionFx = cmdPlaceholder
	case ACTION_CREATE_PATH:
		action.actionFx = cmdPlaceholder
	case ACTION_DELETE:
		action.actionFx = cmdPlaceholder
	case ACTION_DELETE_RECURSIVE:
		action.actionFx = cmdPlaceholder
	case ACTION_DELETE_RECURSIVE_QUIET:
		action.actionFx = cmdPlaceholder
	case ACTION_DELETE_QUIET:
		action.actionFx = cmdPlaceholder
	case ACTION_IS_EXISTS:
		action.actionFx = cmdPlaceholder
	default:
		return action.__reportError("%s: %s",
			"unknown 'Type'",
			action.Type,
		)
	}

	return nil
}

// Run is to instruct the Action to execute its commands.
//
// This function only return `error` value when the instructed action has an
// error.
//
// If `Action.Save` and `Action.SaveFx` are properly set, this method shall
// pass the output of the command and `Save` as Key-Value parameters into
// `Action.SaveFx` and execute it accordingly.
func (action *Action) Run() (err error) {
	output, err := action.actionFx(action)
	if err != nil {
		return err
	}

	if action.Save == "" {
		return nil
	}

	return action.SaveFx(action.Save, output)
}

func (action *Action) __reportError(format string, args ...interface{}) error {
	if action.Name == "" {
		return fmt.Errorf("action '' - "+format, args...)
	}

	args = append([]interface{}{action.Name}, args...)

	return fmt.Errorf("action '%s' - "+format, args...)
}
