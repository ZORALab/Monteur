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

package libmonteur

import (
	"fmt"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

// ProcessString processes a given string with variables templating.
//
// It will return error if the templater fails.
//
// The generated output are not HTMLSafe.
func ProcessString(given string, v map[string]interface{}) (out string,
	err error) {
	out, err = templater.String(given, v)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_PACKAGER_FMT_BAD, err)
		out = ""
	}

	return out, err
}

// ProcessToFilepath processes a given string to a filepath friendly output.
//
// It shall replaces the following symbols into dash (`-`):
//   1. period (`.`)
//   2. percent (`%`)
//   3. dollar (`$`)
//   4. forward slash (`/`)
//   5. backslash(`\`)
//   6. space (` `)
//   7. underscore (`_`)
func ProcessToFilepath(in string) (out string) {
	out = strings.ToLower(in)
	out = strings.ReplaceAll(out, ".", "-")
	out = strings.ReplaceAll(out, "%", "-")
	out = strings.ReplaceAll(out, "$", "-")
	out = strings.ReplaceAll(out, "/", "-")
	out = strings.ReplaceAll(out, "\\", "-")
	out = strings.ReplaceAll(out, " ", "-")
	out = strings.ReplaceAll(out, "_", "-")

	return out
}
