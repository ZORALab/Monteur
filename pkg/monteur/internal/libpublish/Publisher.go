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

type Publisher struct {
	Metadata     *_tomlMetadata
	Variables    map[string]interface{}
	Dependencies []*commander.Dependency
	CMD          []*commander.Action
}

func (fx *Publisher) Parse(path string) (err error) {
	fx.Variables = map[string]interface{}{}

	err = toml.DecodeFile(path, &fx, nil)
	if err != nil {
		return fx.__reportError("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	err = fx.sanitize(path)
	if err != nil {
		return err
	}

	return nil
}

func (fx *Publisher) sanitize(path string) (err error) {
	err = fx._sanitizeMetadata(path)
	if err != nil {
		return err
	}

	err = fx._sanitizeDependencies()
	if err != nil {
		return err
	}

	err = fx._sanitizeCMD()
	if err != nil {
		return err
	}

	return nil
}

func (fx *Publisher) _sanitizeDependencies() (err error) {
	for _, dep := range fx.Dependencies {
		// begin initialization
		err = dep.Init()
		if err != nil {
			return fx.__reportError("%s", err)
		}
	}

	return nil
}

func (fx *Publisher) _sanitizeCMD() (err error) {
	for i, cmd := range fx.CMD {
		// set all settings
		cmd.SaveFx = fx._saveFx

		// begin initialization
		err = cmd.Init()
		if err != nil {
			return fx.__reportError("(CMD %d) %s", i, err)
		}
	}

	return nil
}

func (fx *Publisher) _sanitizeMetadata(path string) (err error) {
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

// Run is to execute the publisher's publishing sequences.
func (fx *Publisher) Run() (err error) {
	return nil
}

func (fx *Publisher) __reportError(format string, args ...interface{}) error {
	if fx.Metadata == nil || fx.Metadata.Name == "" {
		return fmt.Errorf("publisher '' ➤ "+format, args...)
	}

	args = append([]interface{}{fx.Metadata.Name}, args...)
	return fmt.Errorf("'%s' ➤ "+format, args...)
}
