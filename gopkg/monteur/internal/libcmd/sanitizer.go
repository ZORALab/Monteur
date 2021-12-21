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

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

func sanitizeChangelog(changelog *libmonteur.TOMLChangelog,
	system string) (err error) {
	// sanitize CMD
	cmd := changelog.CMD
	changelog.CMD = []*libmonteur.TOMLAction{}
	err = sanitizeCMD(cmd, &changelog.CMD, system)
	if err != nil {
		return err
	}

	return nil
}

func sanitizeCMD(in []*libmonteur.TOMLAction,
	out *[]*libmonteur.TOMLAction,
	system string) (err error) {
	for _, cmd := range in {
		if !libmonteur.IsComputeSystemSupported(system, cmd.Condition) {
			continue
		}

		err = cmd.Sanitize()
		if err != nil {
			return err //nolint:wrapcheck
		}

		*out = append(*out, cmd)
	}

	return nil
}

func sanitizeDeps(in []*libmonteur.TOMLDependency,
	out *[]*commander.Dependency,
	system string,
	variables map[string]interface{}) (err error) {
	var val string

	// scan conditions for building commands list
	for _, dep := range in {
		if !libmonteur.IsComputeSystemSupported(system,
			[]string{dep.Condition}) {
			continue
		}

		val, err = templater.String(dep.Command, variables)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_COMMAND_DEPENDENCY_FMT_BAD,
				err,
			)
		}

		s := &commander.Dependency{
			Name:    dep.Name,
			Type:    dep.Type,
			Command: val,
		}

		*out = append(*out, s)
	}

	// sanitize each commands for validity
	for _, dep := range *out {
		err = dep.Init()
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_DEPENDENCY_BAD,
				err,
			)
		}
	}

	return nil
}

func sanitizeMetadata(meta *libmonteur.TOMLMetadata, path string) (err error) {
	err = meta.Sanitize(path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
