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

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

// Supported variables keys in key:value variables placholders.
//
// It is used in every toml config file inside setup/program/ config directory
// for placeholding variable elements in the fields' value.
const (
	VAR_ARCH    = "Arch"
	VAR_ARCHIVE = "Archive"
	VAR_APP     = "App"
	VAR_BASE    = "BaseDir"
	VAR_BUILD   = "BuildDir"
	VAR_BIN     = "BinDir"
	VAR_CFG     = "ConfigDir"
	VAR_COMPUTE = "ComputeSystem"
	VAR_DOC     = "DocsDir"
	VAR_FORMAT  = "Format"
	VAR_HOME    = "HomeDir"
	VAR_LOG     = "LogDir"
	VAR_METHOD  = "Method"
	VAR_OS      = "OS"
	VAR_ROOT    = "RootDir"
	VAR_SECRETS = "Secrets"
	VAR_TMP     = "WorkingDir"
	VAR_URL     = "URL"
)

func SanitizeVariables(list, fmtVar *map[string]interface{}) (err error) {
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
			val, err = templater.String(v, *list)
		default:
			val = v
		}

		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_VARIABLES_FMT_BAD,
				err,
			)
		}

		(*list)[key] = val
	}

	return nil
}
