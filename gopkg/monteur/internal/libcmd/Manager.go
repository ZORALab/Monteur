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

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

type _tomlMetadata struct {
	Type string
	Name string
}

type _tomlDependency struct {
	Name      string
	Condition string
	Command   string
	Type      commander.ActionID
}

type _tomlAction struct {
	Name      string
	Location  string
	Source    string
	Target    string
	Save      string
	Condition string
	Type      commander.ActionID
}

type Manager struct {
	Metadata     *_tomlMetadata
	thisSystem   string
	omniSystem   string
	Variables    map[string]interface{}
	Dependencies []*commander.Dependency
	CMD          []*commander.Action
}

func (fx *Manager) Parse(path string) (err error) {
	var fmtVar map[string]interface{}
	var dep []*_tomlDependency
	var cmd []*_tomlAction
	var ok bool

	// initialize all important variables
	fx.thisSystem, ok = fx.Variables[libmonteur.VAR_COMPUTE].(string)
	if !ok {
		panic("MONTEUR DEV: please assign variables before Parse()!")
	}

	fx.omniSystem = libmonteur.ALL_OS +
		libmonteur.COMPUTE_SYSTEM_SEPARATOR +
		libmonteur.ALL_ARCH

	fx.Metadata = &_tomlMetadata{}
	dep = []*_tomlDependency{}
	fmtVar = map[string]interface{}{}

	// construct TOML file data structure
	s := struct {
		Metadata     *_tomlMetadata
		Variables    map[string]interface{}
		Dependencies *[]*_tomlDependency
		CMD          *[]*_tomlAction
		FMTVariables *map[string]interface{}
	}{
		Metadata:     fx.Metadata,
		Variables:    fx.Variables,
		FMTVariables: &fmtVar,
		Dependencies: &dep,
		CMD:          &cmd,
	}

	// decode
	err = toml.DecodeFile(path, &s, nil)
	if err != nil {
		return fx.__reportError("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// sanitize
	err = fx.sanitizeMetadata(path)
	if err != nil {
		return err
	}

	err = fx.sanitizeFMTVariables(fmtVar)
	if err != nil {
		return err
	}

	err = fx.sanitizeDependencies(dep)
	if err != nil {
		return err
	}

	err = fx.sanitizeCMD(cmd)
	if err != nil {
		return err
	}

	return nil
}

func (fx *Manager) sanitizeDependencies(in []*_tomlDependency) (err error) {
	var val string

	// initialize all variables
	fx.Dependencies = []*commander.Dependency{}

	// scan conditions for building commands list
	for _, dep := range in {
		if !fx._supportedSystem(dep.Condition) {
			continue
		}

		val, err = templater.String(dep.Command, fx.Variables)
		if err != nil {
			return fx.__reportError("%s: %s",
				libmonteur.ERROR_COMMAND_DEPENDENCY_FMT_BAD,
				err,
			)
		}

		s := &commander.Dependency{
			Name:    dep.Name,
			Type:    dep.Type,
			Command: val,
		}

		fx.Dependencies = append(fx.Dependencies, s)
	}

	// sanitize each commands for validity
	for _, dep := range fx.Dependencies {
		err = dep.Init()
		if err != nil {
			return fx.__reportError("%s", err)
		}
	}

	return nil
}

func (fx *Manager) sanitizeFMTVariables(in map[string]interface{}) (err error) {
	var val interface{}

	for key, value := range in {
		switch v := value.(type) {
		case string:
			val, err = templater.String(v, fx.Variables)
		default:
			val = v
		}

		if err != nil {
			return fx.__reportError("%s: %s",
				libmonteur.ERROR_VARIABLES_FMT_BAD,
				err,
			)
		}

		fx.Variables[key] = val
	}

	return nil
}

func (fx *Manager) sanitizeCMD(in []*_tomlAction) (err error) {
	var location, source, target string

	// initialize all variables
	fx.CMD = []*commander.Action{}

	// scan conditions for building commands list
	for _, cmd := range in {
		if !fx._supportedSystem(cmd.Condition) {
			continue
		}

		location, err = templater.String(cmd.Location, fx.Variables)
		if err != nil {
			return fx.__reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)
		}

		source, err = templater.String(cmd.Source, fx.Variables)
		if err != nil {
			return fx.__reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)
		}

		target, err = templater.String(cmd.Target, fx.Variables)
		if err != nil {
			return fx.__reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)
		}

		a := &commander.Action{
			Name:     cmd.Name,
			Type:     cmd.Type,
			Location: location,
			Source:   source,
			Target:   target,
			Save:     cmd.Save,
			SaveFx:   fx._saveFx,
		}

		fx.CMD = append(fx.CMD, a)
	}

	// sanitize each commands for validity
	for i, cmd := range fx.CMD {
		err = cmd.Init()
		if err != nil {
			return fx.__reportError("(CMD %d) %s", i, err)
		}
	}

	return nil
}

func (fx *Manager) sanitizeMetadata(path string) (err error) {
	if fx.Metadata == nil {
		return fx.__reportError("%s: %s",
			libmonteur.ERROR_PUBLISH_METADATA_MISSING,
			path,
		)
	}

	if fx.Metadata.Name == "" {
		return fx.__reportError("%s: '%s' for %s",
			libmonteur.ERROR_PUBLISH_METADATA_MISSING,
			"Name",
			path,
		)
	}

	if fx.Metadata.Type == "" {
		return fx.__reportError("%s: '%s' for %s",
			libmonteur.ERROR_PUBLISH_METADATA_MISSING,
			path,
		)
	}

	return nil
}

func (fx *Manager) _saveFx(key string, output interface{}) (err error) {
	fx.Variables[key] = output

	return nil
}

func (fx *Manager) _supportedSystem(condition string) bool {
	switch condition {
	case fx.thisSystem:
		return true
	case fx.omniSystem:
		return true
	default:
		return false
	}
}

// Run is to execute the publisher's commands sequence.
func (fx *Manager) Run() (err error) {
	for i, cmd := range fx.CMD {
		err = cmd.Run()
		if err != nil {
			return fx.__reportError("%s: (Step %d) %s",
				libmonteur.ERROR_COMMAND_FAILED,
				i+1,
				err,
			)
		}
	}

	return nil
}

func (fx *Manager) __reportError(format string, args ...interface{}) error {
	if fx.Metadata == nil || fx.Metadata.Name == "" {
		return fmt.Errorf("publisher '' ➤ "+format, args...)
	}

	args = append([]interface{}{fx.Metadata.Name}, args...)
	return fmt.Errorf("'%s' ➤ "+format, args...)
}
