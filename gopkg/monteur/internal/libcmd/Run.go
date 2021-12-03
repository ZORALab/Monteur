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

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

type Run struct {
}

func (fx *Run) Parse(path string, varList *map[string]interface{}) (err error) {
	// initiate working variables
	fmtVar := &map[string]interface{}{}

	// construct TOML file data structure
	s := struct {
		Variables    *map[string]interface{}
		FMTVariables *map[string]interface{}
	}{
		Variables:    varList,
		FMTVariables: fmtVar,
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
	err = libmonteur.SanitizeVariables(varList, fmtVar)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
