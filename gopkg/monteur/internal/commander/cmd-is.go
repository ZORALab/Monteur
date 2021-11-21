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

func cmdIsExists(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return false, fmt.Errorf("source is empty")
	}

	_, err = os.Stat(action.Source)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, fmt.Errorf("%s: %s", "target not exist", err)
	}

	return false, fmt.Errorf("%s: %s",
		"error finding target existence",
		err,
	)
}
