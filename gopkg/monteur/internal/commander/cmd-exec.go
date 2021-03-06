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
	"bytes"
	"fmt"
)

type ExecOutput struct {
	Stdout []byte
	Stderr []byte
}

func cmdExec(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return nil, fmt.Errorf("source is empty")
	}

	// construct all necessary data
	t := _createTerminal()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	t.Stdout = stdout
	t.Stderr = stderr
	x := &ExecOutput{}

	err = t.Exec(action.Source, 0)

	// process output
	x.Stdout = stdout.Bytes()
	x.Stderr = stderr.Bytes()
	if err != nil {
		err = fmt.Errorf("%s: %s", "failed to execute command", err)
	}

	return x, err
}

func cmdExecQuiet(action *Action) (out interface{}, err error) {
	out, _ = cmdExec(action)
	return out, nil
}
