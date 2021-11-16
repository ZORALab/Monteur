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

package libsecrets

import (
	"os"
	"path/filepath"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/secrets"
)

func GetSecrets(pathings []string) (sec map[string]interface{}) {
	sec = map[string]interface{}{}

	for _, path := range pathings {
		sec = _parseSecrets(sec, path)
	}

	return sec
}

func _parseSecrets(data map[string]interface{},
	pathing string) (sec map[string]interface{}) {
	if i, err := os.Stat(pathing); !os.IsNotExist(err) && !i.IsDir() {
		if filepath.Ext(pathing) != libmonteur.EXTENSION_TOML {
			return data // not a toml data file
		}

		sec = __parseSecretsFile(data, pathing)
		return sec
	}

	sec = data
	_ = filepath.Walk(pathing, func(path string, info os.FileInfo,
		err error) error {
		if path == pathing {
			return nil
		}

		sec = _parseSecrets(sec, path)
		return nil
	})

	return sec
}

func __parseSecretsFile(data map[string]interface{},
	path string) (sec map[string]interface{}) {
	var err error

	if filepath.Ext(path) != libmonteur.EXTENSION_TOML {
		return data // not a correct file
	}

	s := &secrets.Processor{
		DecodeFx: toml.DecodeFile,
	}

	sec, err = s.Decode(data, path, nil)
	if err != nil {
		return data
	}

	return sec
}
