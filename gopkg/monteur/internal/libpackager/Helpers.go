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

package libpackager

import (
	"fmt"
	"os"
	"path/filepath"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libtemplater"
	"gitlab.com/zoralab/monteur/gopkg/oshelper"
)

func UpdatePackagePath(variables *map[string]interface{},
	pkg *libmonteur.TOMLPackage,
	packagingType string,
	log func(string, ...interface{})) (packagePath string, err error) {
	var name string

	log("Updating package path...")
	name = (*variables)[libmonteur.VAR_PACKAGE_NAME].(string)
	name += "-" + (*variables)[libmonteur.VAR_PACKAGE_VERSION].(string)
	name += "-" + (*variables)[libmonteur.VAR_PACKAGE_OS].(string)
	name += "-" + (*variables)[libmonteur.VAR_PACKAGE_ARCH].(string)

	packagePath = (*variables)[libmonteur.VAR_PACKAGE].(string)
	packagePath = filepath.Join(packagePath, packagingType, name)
	(*variables)[libmonteur.VAR_PACKAGE] = packagePath
	log("➤ Got: '%s'", packagePath)

	log("Creating package path %s ...", packagePath)
	_ = os.RemoveAll(packagePath)
	err = os.MkdirAll(packagePath, libmonteur.PERMISSION_DIRECTORY)
	if err != nil {
		return "", fmt.Errorf("%s: %s",
			libmonteur.ERROR_PACKAGER_MKDIR,
			err,
		)
	}
	log(libmonteur.LOG_OK)

	return packagePath, nil
}

func AssemblePackage(pkg *libmonteur.TOMLPackage,
	variables map[string]interface{},
	log func(string, ...interface{})) (err error) {
	log("Assembling package contents now...")

	for k, v := range pkg.Files {
		k, err = libtemplater.Template(k, variables)
		if err != nil {
			return err //nolint:wrapcheck
		}

		v, err = libtemplater.Template(v, variables)
		if err != nil {
			return err //nolint:wrapcheck
		}

		log("Placing merchandise...")
		log("From: %s", v)
		log("To  : %s", k)

		_, err = os.Stat(v)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_PACKAGER_FILE_MISSING,
				v,
			)
		}

		err = oshelper.CopyPath(v, k)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_PACKAGER_FILES_COPY_FAILED,
				err,
			)
		}

		log(libmonteur.LOG_OK)
	}

	return nil
}
