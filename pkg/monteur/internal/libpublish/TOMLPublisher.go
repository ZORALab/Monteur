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

type TOMLPublisher struct {
	Metadata     *_tomlMetadata
	Variables    map[string]interface{}
	Dependencies []*commander.Dependency
	CMD          []*commander.Action
}

func (fx *TOMLPublisher) Parse(path string) (err error) {
	fx.Variables = map[string]interface{}{}

	err = toml.DecodeFile(path, &fx, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	return nil
}

func (fx *TOMLPublisher) Process() (p *Publisher, err error) {
	p = &Publisher{}

	return p, nil
}