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

package archiver

import (
	"context"
	"os"
	"path/filepath"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/checksum"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/oshelper"
)

type job struct {
	ch chan conductor.Message

	// Checksum is for generating checksum hasher
	Checksum *checksum.Hasher

	// Version is for categorizing the release using version numbers
	Version string

	// Filename is the filename of the release
	Filename string

	// SourcePath is the filepath to the package file.
	SourcePath string

	// DestPath is the directory path to house the released package.
	DestPath string
}

// Run is to execute the release job.
func (me *job) Run(ctx context.Context, ch chan conductor.Message) {
	var err error
	var out, version, destPath string

	// register channel
	me.ch = ch

	// formulate version
	version = VersionToDir(me.Version)
	me.reportStatus("%s: %s", "process version pathing", version)

	// formulate pathing
	destPath = filepath.Join(me.DestPath, version, me.Filename)

	// making housing directories
	me.reportStatus("%s: %s",
		"making housing directory",
		filepath.Dir(destPath),
	)
	err = os.MkdirAll(filepath.Dir(destPath), PERMISSION_DIR)
	if err != nil {
		me.reportError("%s: %s", ERROR_TARGET_DIR_CREATE, err)
		return
	}

	// copy Source to Dest
	err = oshelper.Copy(me.SourcePath, destPath)
	if err != nil {
		me.reportError("%s: %s", ERROR_TARGET_COPY, err)
		return
	}

	// checksum Dest
	err = me.Checksum.HashFile(destPath)
	if err != nil {
		me.reportError("%s: %s", ERROR_TARGET_CHECKSUM, err)
		return
	}

	// get checksum hex value
	out, err = me.Checksum.ToHex()
	if err != nil {
		me.reportError("%s: %s", ERROR_TARGET_CHECKSUM, err)
		return
	}

	// report output
	me.reportOutput("%s %s", out, destPath)

	// report done
	me.reportDone()
}

// Name generates the job name.
func (me *job) Name() string {
	return me.Filename
}

func (me *job) reportOutput(format string, args ...interface{}) {
	if me.ch != nil {
		me.ch <- conductor.CreateOutput(me.Filename, format, args...)
	}
}

func (me *job) reportStatus(format string, args ...interface{}) {
	if me.ch != nil {
		me.ch <- conductor.CreateStatus(me.Filename, format, args...)
	}
}

//nolint:unparam
func (me *job) reportError(format string, args ...interface{}) {
	if me.ch != nil {
		me.ch <- conductor.CreateError(me.Filename, format, args...)
	}
}

func (me *job) reportDone() {
	if me.ch != nil {
		me.ch <- conductor.CreateDone(me.Filename)
	}
}
