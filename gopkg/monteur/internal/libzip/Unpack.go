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
	"path/filepath"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/archive/zip"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

func Unpack(source *libmonteur.TOMLSource,
	variables map[string]interface{}) (err error) {
	var ok bool
	var raw, archive string

	// extract Raw directory pathing
	raw, ok = variables[libmonteur.VAR_TMP].(string)
	if !ok {
		panic("MONTEUR DEV: why is VAR_TMP not assigned?")
	}

	// process destination pathing
	archive = filepath.Join(raw, source.Archive)

	// setup Archiver
	archiver := &zip.Archiver{
		Archive:         archive,
		Raw:             raw,
		CreateDirectory: true,
		Overwrite:       true,
	}

	// extract Raw data from Archive
	err = archiver.Extract()
	if err != nil {
		err = fmt.Errorf("%s (%s -> %s): %s",
			"error unpack .zip program",
			archive,
			raw,
			err,
		)
	}

	return err
}
