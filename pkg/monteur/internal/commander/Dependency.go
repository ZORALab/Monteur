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
	"os/exec"
)

type Dependency struct {
	Name string
	Type ActionID
}

// Init is a method to ensure Dependency is sanitized and ready for execution.
//
// It validates all known configurations before executing the commands.
func (dep *Dependency) Init() (err error) {
	if dep.Name == "" {
		return dep.__reportError("Name is empty")
	}

	switch dep.Type {
	case ACTION_COMMAND, ACTION_COMMAND_QUIET:
		return dep.checkCommandDependency()
	default:
	}

	return nil
}

func (dep *Dependency) checkCommandDependency() (err error) {
	path, err := exec.LookPath(dep.Name)

	if err != nil {
		return dep.__reportError("missing")
	}

	dep.Name = path

	return nil
}

func (dep *Dependency) __reportError(format string, args ...interface{}) error {
	if dep.Name == "" {
		return fmt.Errorf("dependency '' - "+format, args...)
	}

	args = append([]interface{}{dep.Name}, args...)

	return fmt.Errorf("dependency '%s' - "+format, args...)
}
