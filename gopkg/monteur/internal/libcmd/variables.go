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

package libcmd

import (
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

func processPackageVariables(pkg *libmonteur.TOMLPackage,
	variables *map[string]interface{}) (err error) {
	var ok bool
	var packagePath string
	var app *libmonteur.Software

	// process critical variables
	app, ok = (*variables)[libmonteur.VAR_APP].(*libmonteur.Software)
	if !ok {
		panic("MONTEUR DEV: why is VAR_APP missing?")
	}

	(*variables)[libmonteur.VAR_PACKAGE_NAME] =
		libmonteur.ProcessToFilepath(app.Name)
	(*variables)[libmonteur.VAR_PACKAGE_VERSION] =
		libmonteur.ProcessToFilepath(app.Version)
	(*variables)[libmonteur.VAR_PACKAGE_VERSION_DIGIT_LED] =
		libmonteur.ProcessDigitLedVersion(app.Version)
	(*variables)[libmonteur.VAR_PACKAGE_OS] = pkg.OS[0]
	(*variables)[libmonteur.VAR_PACKAGE_ARCH] = pkg.Arch[0]

	// process VAR_PACKAGE (PackageDir) base directory
	packagePath, ok = (*variables)[libmonteur.VAR_PACKAGE].(string)
	if !ok {
		packagePath, ok = (*variables)[libmonteur.VAR_TMP].(string)
		if !ok {
			panic("MONTEUR DEV: why is libmonteur.VAR_TMP missing?")
		}
	}
	(*variables)[libmonteur.VAR_PACKAGE] = packagePath

	// process pkg.Changelog pathing
	pkg.Changelog, err = libmonteur.ProcessString(pkg.Changelog,
		*variables,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
