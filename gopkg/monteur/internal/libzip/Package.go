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

package libzip

import (
	"fmt"
	"os"
	"path/filepath"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/archive/zip"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

func Package(pkg *libmonteur.TOMLPackage,
	variables map[string]interface{}) (err error) {
	var archivePath string

	// process targz-specific variables
	packagePath := variables[libmonteur.VAR_PACKAGE].(string)
	packagePath = filepath.Join(packagePath, libmonteur.PACKAGE_ZIP)
	variables[libmonteur.VAR_PACKAGE] = packagePath
	_ = os.MkdirAll(packagePath, libmonteur.PERMISSION_DIRECTORY)

	// process necessary internal variables
	variables[libmonteur.VAR_PACKAGE_OS] = pkg.OS[0]
	variables[libmonteur.VAR_PACKAGE_ARCH] = pkg.Arch[0]

	pkg.Name, err = libmonteur.ProcessString(pkg.Name, variables)
	if err != nil {
		return err //nolint:wrapcheck
	}

	pkg.Name = libmonteur.ProcessToFilepath(pkg.Name)

	packagePath = variables[libmonteur.VAR_PACKAGE].(string)
	archivePath = filepath.Join(filepath.Dir(packagePath),
		pkg.Name+zip.EXTENSION,
	)

	archiver := &zip.Archiver{
		Archive:         archivePath,
		Raw:             packagePath,
		CreateDirectory: true,
		Overwrite:       true,
		Compression:     zip.COMPRESSION_DEFLATE,
	}

	err = archiver.Compress()
	if err != nil {
		err = fmt.Errorf("%s (%s): %s",
			"error packaging zip package",
			archivePath,
			err,
		)
	}

	return err
}
