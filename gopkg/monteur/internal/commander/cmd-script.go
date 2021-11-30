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
	"os"
)

const (
	SCRIPT_PERMISSION = 0600
)

func cmdScript(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return nil, fmt.Errorf("source is empty")
	}

	if action.Target == "" {
		return nil, fmt.Errorf("target is empty")
	}

	if _, err = os.Stat(action.Target); !os.IsNotExist(err) {
		return nil, fmt.Errorf("target exists")
	}

	// remove target regardlessly
	_ = os.RemoveAll(action.Target)

	// move source to target
	err = os.WriteFile(action.Target,
		[]byte(action.Source),
		SCRIPT_PERMISSION,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "error while scripting", err)
	}

	return nil, nil
}

func cmdScriptQuiet(action *Action) (out interface{}, err error) {
	out, _ = cmdMove(action)
	return out, nil
}
