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
	"strconv"
	"strings"
)

const (
	CHOWN_SEPARATOR = ":"
)

func cmdChown(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return nil, fmt.Errorf("source is empty")
	}

	if action.Target == "" {
		return nil, fmt.Errorf("target is empty")
	}

	if _, err = os.Stat(action.Source); os.IsNotExist(err) {
		return nil, fmt.Errorf("source does not exist")
	}

	// split action.Target to for uid and gid parsing
	users := strings.Split(action.Target, CHOWN_SEPARATOR)
	if len(users) != 2 {
		return nil, fmt.Errorf("target should be 'UID:GID' int format")
	}

	// get uid and gid
	uid, err := strconv.Atoi(users[0])
	if err != nil {
		return nil, fmt.Errorf("bad uid: %s", err)
	}

	gid, err := strconv.Atoi(users[1])
	if err != nil {
		return nil, fmt.Errorf("bad gid: %s", err)
	}

	// chown source file
	err = os.Chown(action.Source, uid, gid)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "error while chown", err)
	}

	return nil, nil
}

func cmdChownQuiet(action *Action) (out interface{}, err error) {
	out, _ = cmdChown(action)
	return out, nil
}
