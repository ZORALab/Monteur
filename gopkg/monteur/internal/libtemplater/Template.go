// Copyright 2022 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2022 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
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

package libtemplater

import (
	"fmt"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

func Template(in string, v map[string]interface{}) (out string, err error) {
	var sec *libsecrets.Secrets
	var ok bool
	var funcMap map[string]interface{}

	sec, ok = v[libmonteur.VAR_SECRETS].(*libsecrets.Secrets)
	if !ok {
		panic("MONTEUR DEV: why is VAR_SECRETS missing?")
	}

	funcMap = map[string]interface{}{
		"GetSecret": sec.Query,
	}

	out, err = templater.String(in, v, funcMap)
	if err != nil {
		return "", fmt.Errorf("%s: %s",
			libmonteur.ERROR_PACKAGER_FMT_BAD,
			err,
		)
	}

	return out, nil
}

func TemplateRaw(in string, v map[string]interface{}) (out string, err error) {
	out, err = templater.String(in, v, map[string]interface{}{})
	if err != nil {
		return "", fmt.Errorf("%s: %s",
			libmonteur.ERROR_PACKAGER_FMT_BAD,
			err,
		)
	}

	return out, nil
}

func TemplateVariables(list, fmtVar *map[string]interface{}) (err error) {
	var val interface{}

	switch {
	case list == nil, *list == nil:
		panic("MONTEUR DEV: given Variables list is nil")
	case fmtVar == nil, *fmtVar == nil:
		panic("MONTEUR DEV: given FMTVariables list is nil")
	}

	for key, value := range *fmtVar {
		switch v := value.(type) {
		case string:
			val, err = Template(v, *list)
		default:
			val = v
		}

		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_VARIABLES_FMT_BAD,
				err,
			)
		}

		(*list)[key] = val
	}

	return nil
}

func TemplateVariablesRaw(list, fmtVar *map[string]interface{}) (err error) {
	var val interface{}

	switch {
	case list == nil, *list == nil:
		panic("MONTEUR DEV: given Variables list is nil")
	case fmtVar == nil, *fmtVar == nil:
		panic("MONTEUR DEV: given FMTVariables list is nil")
	}

	for key, value := range *fmtVar {
		switch v := value.(type) {
		case string:
			val, err = TemplateRaw(v, *list)
		default:
			val = v
		}

		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_VARIABLES_FMT_BAD,
				err,
			)
		}

		(*list)[key] = val
	}

	return nil
}

// ToFilepath processes a given string to a filepath friendly output.
//
// It shall replaces the following symbols into dash (`-`):
//   1. period (`.`)
//   2. percent (`%`)
//   3. dollar (`$`)
//   4. forward slash (`/`)
//   5. backslash(`\`)
//   6. space (` `)
//   7. underscore (`_`)
func ToFilepath(in string) (out string) {
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

// ToDigitLedVersion processes given version to digit-led string.
//
// If the given string is already digit-led, the function shall return the
// version string as it is.
func ToDigitLedVersion(in string) string {
	for i, c := range in {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if i == 0 {
				return in
			}

			return in[i:]
		}
	}

	return in
}
