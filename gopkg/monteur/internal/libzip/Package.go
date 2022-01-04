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
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libpackager"
)

func Package(pkg *libmonteur.TOMLPackage,
	variables *map[string]interface{},
	log *liblog.Logger) (err error) {
	var archivePath, packagePath string

	log.Info("Preparing %s packing ...", libmonteur.PACKAGE_ZIP)

	// process package pathing
	packagePath, err = libpackager.UpdatePackagePath(variables,
		pkg,
		libmonteur.PACKAGE_ZIP,
		log.Info,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// copy all files into workspace
	err = libpackager.AssemblePackage(pkg, *variables, log.Info)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// begin archive
	archivePath = filepath.Join(filepath.Dir(packagePath),
		filepath.Base(packagePath)+zip.EXTENSION,
	)
	_ = os.RemoveAll(archivePath)

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
