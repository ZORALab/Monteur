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

package libarchiver

import (
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/release/archiver"
)

func Init(logger *liblog.Logger,
	release *libmonteur.TOMLRelease,
	variables map[string]interface{}) (out *archiver.Manager, err error) {
	var app *libmonteur.Software
	var ok bool

	app, ok = variables[libmonteur.VAR_APP].(*libmonteur.Software)
	if !ok {
		panic("MONTEUR DEV: why is VAR_APP not assigned?")
	}

	out = &archiver.Manager{
		Log:      logger,
		Path:     release.Target,
		DataPath: release.Data.Path,
		Version:  app.Version,
	}

	switch release.Data.Format {
	case libmonteur.FORMAT_TOML:
		out.Format = archiver.FORMAT_TOML
	case libmonteur.FORMAT_CSV:
		out.Format = archiver.FORMAT_CSV
	case libmonteur.FORMAT_TXT:
		fallthrough
	default:
		out.Format = archiver.FORMAT_TXT
	}

	switch release.Checksum {
	case libmonteur.CHECKSUM_ALGO_SHA512_TO_SHA256:
		out.Checksum = archiver.CHECKSUM_SHA512_TO_256
	case libmonteur.CHECKSUM_ALGO_SHA512:
		out.Checksum = archiver.CHECKSUM_SHA512
	case libmonteur.CHECKSUM_ALGO_SHA256:
		fallthrough
	default:
		out.Checksum = archiver.CHECKSUM_SHA256
	}

	return out, nil
}

func Release(pkg *libmonteur.TOMLReleasePackage,
	controller interface{},
	variables map[string]interface{}) (err error) {
	var ok bool
	var manager *archiver.Manager

	manager, ok = controller.(*archiver.Manager)
	if !ok {
		panic("MONTEUR DEV: libarchiver's controller misconfigured!")
	}

	return manager.Add(pkg.Source) //nolint: wrapcheck
}

func Conclude(controller interface{}) (err error) {
	var ok bool
	var manager *archiver.Manager

	manager, ok = controller.(*archiver.Manager)
	if !ok {
		panic("MONTEUR DEV: libarchiver's controller misconfigured!")
	}

	return manager.Release() //nolint:wrapcheck
}
