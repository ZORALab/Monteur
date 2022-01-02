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
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/oshelper"
)

// Package type
const (
	PACKAGE_DEB_MANUAL = "deb-manual"
	PACKAGE_MANUAL     = "manual"
	PACKAGE_TARGZ      = "targz"
	PACKAGE_ZIP        = "zip"
)

func UpdatePackagePath(variables *map[string]interface{},
	pkg *TOMLPackage,
	packagingType string,
	log func(string, ...interface{})) (packagePath string, err error) {
	var name string

	// get package name and process it
	name, err = ProcessString(pkg.Name, *variables)
	if err != nil {
		return "", err
	}

	log("Updating package path...")
	packagePath = (*variables)[VAR_PACKAGE].(string)
	packagePath = filepath.Join(packagePath, packagingType, name)
	(*variables)[VAR_PACKAGE] = packagePath
	log("➤ Got: '%s'", packagePath)

	log("Creating package path %s ...", packagePath)
	_ = os.RemoveAll(packagePath)
	err = os.MkdirAll(packagePath, PERMISSION_DIRECTORY)
	if err != nil {
		return "", fmt.Errorf("%s: %s", ERROR_PACKAGER_MKDIR, err)
	}
	log(strings.TrimSuffix(LOG_SUCCESS, "\n"))

	return packagePath, nil
}

func AssemblePackage(pkg *TOMLPackage,
	variables map[string]interface{},
	log func(string, ...interface{})) (err error) {
	log("Assembling package contents now...")

	for k, v := range pkg.Files {
		k, err = ProcessString(k, variables)
		if err != nil {
			return err
		}

		v, err = ProcessString(v, variables)
		if err != nil {
			return err
		}

		log("Placing merchandise...")
		log("From: %s", v)
		log("To  : %s", k)

		_, err = os.Stat(v)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_PACKAGER_FILE_MISSING,
				v,
			)
		}

		err = oshelper.CopyPath(v, k)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_PACKAGER_FILES_COPY_FAILED,
				err,
			)
		}

		log(strings.TrimSuffix(LOG_SUCCESS, "\n"))
	}

	return nil
}
