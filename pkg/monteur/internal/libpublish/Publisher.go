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

package libpublish

import (
	"fmt"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
)

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

type Publisher struct {
	Metadata     *_tomlMetadata
	thisSystem   string
	omniSystem   string
	Variables    map[string]interface{}
	Dependencies []*commander.Dependency
	CMD          []*commander.Action
}

func (fx *Publisher) Parse(path string) (err error) {
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

	// construct TOML file data structure
	s := struct {
		Metadata     *_tomlMetadata
		Variables    map[string]interface{}
		Dependencies *[]*_tomlDependency
		CMD          *[]*_tomlAction
	}{
		Metadata:     fx.Metadata,
		Variables:    fx.Variables,
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

	err = fx.sanitizeMetadata(path)
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

func (fx *Publisher) sanitizeDependencies(in []*_tomlDependency) (err error) {
	// initialize all variables
	fx.Dependencies = []*commander.Dependency{}

	// scan conditions for building commands list
	for _, dep := range in {
		if !fx._supportedSystem(dep.Condition) {
			continue
		}

		s := &commander.Dependency{
			Name:    dep.Name,
			Type:    dep.Type,
			Command: dep.Command,
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

func (fx *Publisher) sanitizeCMD(in []*_tomlAction) (err error) {
	// initialize all variables
	fx.CMD = []*commander.Action{}

	// scan conditions for building commands list
	for _, cmd := range in {
		if !fx._supportedSystem(cmd.Condition) {
			continue
		}

		a := &commander.Action{
			Name:     cmd.Name,
			Type:     cmd.Type,
			Location: cmd.Location,
			Source:   cmd.Source,
			Target:   cmd.Target,
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

func (fx *Publisher) sanitizeMetadata(path string) (err error) {
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

func (fx *Publisher) _saveFx(key string, output interface{}) (err error) {
	fx.Variables[key] = output

	return nil
}

func (fx *Publisher) _supportedSystem(condition string) bool {
	switch condition {
	case fx.thisSystem:
		return true
	case fx.omniSystem:
		return true
	default:
		return false
	}
}

// Publish is to execute the publisher's publishing sequences.
func (fx *Publisher) Publish() (err error) {
	return fx.run(true)
}

// Build is to execute the publisher's publication material building sequences.
func (fx *Publisher) Build() (err error) {
	return fx.run(false)
}

func (fx *Publisher) run(isPublishing bool) (err error) {
	for i, cmd := range fx.CMD {
		err = cmd.Run()
		if err != nil {
			fx.__reportError("failed to execute CMD: (%d) %s",
				i,
				err,
			)
		}
	}

	return nil
}

func (fx *Publisher) __reportError(format string, args ...interface{}) error {
	if fx.Metadata == nil || fx.Metadata.Name == "" {
		return fmt.Errorf("publisher '' ➤ "+format, args...)
	}

	args = append([]interface{}{fx.Metadata.Name}, args...)
	return fmt.Errorf("'%s' ➤ "+format, args...)
}
