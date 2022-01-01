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
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/checksum"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

const (
	PERMISSION_FILE = 0644
	PERMISSION_DIR  = 0755
)

type Manager struct {
	mutex         *sync.Mutex
	jobs          map[string]conductor.Job
	dataTemplate  string
	dataExtension string

	// Log is the Input/Output processor interface.
	//
	// If this is unassigned, the Manager will operate silently.
	Log Logger

	// Path is the output directory pathing.
	//
	// If path does not exists, Manager shall create it on-behalf.
	//
	// This field is **MANDATORY**.
	Path string

	// DataPath is the output directory of data files.
	//
	// If path does not exists, Manager shall only process the release
	// artifacts without generating its corresponding data file.
	DataPath string

	// Version is the version ID for organizing the release artifacts.
	//
	// This field is **MANDATORY**.
	Version string

	// Format is the format for generating the data file.
	//
	// It is only used if a valid DataPath is present.
	Format FormatID

	// Checksum is the checksum type being used for this release.
	//
	// It is use for generating the final checksum output list.
	Checksum checksum.HashType
}

// Sanitize checks and validates given input data are valid and ready for use.
//
// It shall return error should there be any non-compliant data.
func (me *Manager) Sanitize() (err error) {
	if me.mutex == nil {
		me.mutex = &sync.Mutex{}
	}

	me.mutex.Lock()
	defer me.mutex.Unlock()

	if me.Version == "" {
		return fmt.Errorf(ERROR_VERSION_MISSING)
	}

	err = me.sanitizePath()
	if err != nil {
		return err
	}

	err = me.sanitizeDataPath()
	if err != nil {
		return err
	}

	err = me.sanitizeChecksum()
	if err != nil {
		return err
	}

	return nil
}

func (me *Manager) sanitizePath() (err error) {
	var info os.FileInfo

	if me.Path == "" {
		return fmt.Errorf(ERROR_PATH_MISSING)
	}

	info, err = os.Stat(me.Path)
	switch {
	case err == nil && info.IsDir():
	case err == nil && !info.IsDir():
		return fmt.Errorf(ERROR_PATH_INVALID)
	case os.IsNotExist(err):
		err = os.MkdirAll(me.Path, PERMISSION_DIR)
		if err != nil {
			return fmt.Errorf("%s: %s", ERROR_PATH_CREATE, err)
		}
	default:
		return fmt.Errorf(ERROR_PATH)
	}

	return nil
}

func (me *Manager) sanitizeDataPath() (err error) {
	var info os.FileInfo

	if me.DataPath == "" {
		return nil
	}

	info, err = os.Stat(me.DataPath)
	switch {
	case err == nil && info.IsDir():
	case err == nil && !info.IsDir():
		return fmt.Errorf(ERROR_DATAPATH_INVALID)
	case os.IsNotExist(err):
		err = os.MkdirAll(me.DataPath, PERMISSION_DIR)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DATAPATH_CREATE,
				err,
			)
		}
	default:
		return fmt.Errorf(ERROR_PATH)
	}

	// sanitize Format
	switch me.Format {
	case FORMAT_NONE:
	case FORMAT_CSV:
		me.dataTemplate = _TEMPLATE_CSV
		me.dataExtension = _EXTENSION_CSV
	case FORMAT_TXT:
		me.dataTemplate = _TEMPLATE_TXT
		me.dataExtension = _EXTENSION_TXT
	case FORMAT_TOML:
		me.dataTemplate = _TEMPLATE_TOML
		me.dataExtension = _EXTENSION_TOML
	default:
		return fmt.Errorf("%s: %d", ERROR_FORMAT_UNSUPPORTED, me.Format)
	}

	return nil
}

func (me *Manager) sanitizeChecksum() (err error) {
	switch me.Checksum {
	case checksum.HASHER_SHA256:
	case checksum.HASHER_SHA512:
	case checksum.HASHER_SHA512_TO_SHA256:
	default:
		return fmt.Errorf(ERROR_CHECKSUM_UNSUPPORTED)
	}

	return nil
}

// Add stages a given package filepath into the release jobs list.
func (me *Manager) Add(target string) (err error) {
	var info os.FileInfo
	var mode os.FileMode
	var j *job

	err = me.Sanitize()
	if err != nil {
		return err
	}

	// acquire lock after sanitization
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// sanitize input path
	info, err = os.Stat(target)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("%s: %s", ERROR_TARGET_MISSING, err)
	}

	if err == nil {
		mode = info.Mode()
		switch {
		case mode.IsRegular():
		case !mode.IsRegular():
			return fmt.Errorf(ERROR_TARGET_INVALID)
		default:
			return fmt.Errorf(ERROR_PATH)
		}
	}

	// create target data structure
	j = &job{
		Version:    me.Version,
		DestPath:   me.Path,
		SourcePath: target,
		Filename:   filepath.Base(target),
		Checksum:   me.createHasher(),
	}

	// add to jobs
	if me.jobs == nil {
		me.jobs = map[string]conductor.Job{}
	}
	me.jobs[j.Name()] = j

	return nil
}

func (me *Manager) createHasher() (out *checksum.Hasher) {
	var err error

	out = &checksum.Hasher{}

	err = out.SetAlgo(me.Checksum)
	if err != nil {
		_ = out.SetAlgo(checksum.HASHER_SHA256)
	}

	return out
}

// Release executes the release sequences in parallel.
func (me *Manager) Release() (err error) {
	var c *conductor.Conductor
	var logger *reporter

	err = me.Sanitize()
	if err != nil {
		return err
	}

	// acquire lock after sanitization
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// validate length of data
	if len(me.jobs) == 0 {
		return nil
	}

	// create reporter
	logger = me.createLogger()

	// create conductor
	c = &conductor.Conductor{
		Runners: me.jobs,
		Log:     logger,
	}

	// orchestrate
	err = c.Run()
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = c.Coordinate()
	if err != nil {
		return err //nolint:wrapcheck
	}

	// write overall checksum text file
	err = me.writeChecksumFile(logger.checksums)
	if err != nil {
		return err
	}

	// write checksum data files
	err = me.writeChecksumDataFile(logger.checksums)
	if err != nil {
		return err
	}

	return nil
}

func (me *Manager) writeChecksumDataFile(data []*metadata) (err error) {
	var f *os.File
	var dataPath, out string

	if me.DataPath == "" {
		return nil
	}

	if len(data) == 0 {
		return fmt.Errorf(ERROR_CHECKSUM_EMPTY)
	}

	// formulate pathing
	dataPath = filepath.Join(me.DataPath,
		VersionToDir(me.Version)+me.dataExtension,
	)

	// create housing directory to ensure data file is workable
	err = os.MkdirAll(filepath.Dir(dataPath), PERMISSION_DIR)
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_TARGET_DATA_DIR_CREATE,
			err,
		)
	}

	// delete previous data file if exists
	_ = os.RemoveAll(dataPath)

	// open data file for write
	f, err = os.OpenFile(dataPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		PERMISSION_FILE,
	)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_TARGET_CHECKSUM, err)
	}
	defer func() {
		_ = f.Sync()
		f.Close()
	}()

	// write data into data file
	for i, v := range data {
		// insert index
		v.Index = i

		// template text
		out, err = templater.String(me.dataTemplate, v)
		if err != nil {
			return fmt.Errorf("%s (%s): %s",
				ERROR_TARGET_CHECKSUM,
				v.Filename,
				err,
			)
		}

		if i != 0 {
			out = "\n\n\n\n\n" + out
		}

		// write into file
		_, err = f.WriteString(out)
		if err != nil {
			return fmt.Errorf("%s (%s): %s",
				ERROR_TARGET_CHECKSUM,
				v.Filename,
				err,
			)
		}
	}

	return nil
}

func (me *Manager) writeChecksumFile(data []*metadata) (err error) {
	var path, out string
	var f *os.File

	// sanitize input
	if len(data) == 0 {
		return fmt.Errorf("%s: missing data", ERROR_TARGET_CHECKSUM)
	}

	for i, v := range data {
		if i != 0 {
			out += "\n"
		}

		out += v.Hash + " " + v.Filename
	}

	// formulate checksum file pathing
	path = filepath.Join(me.Path, VersionToDir(me.Version),
		_CHECKSUM_FILENAME,
	)

	// delete checksum file
	_ = os.RemoveAll(path)

	// open checksum file
	f, err = os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		PERMISSION_FILE,
	)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_TARGET_CHECKSUM, err)
	}

	// write checksum data into file
	_, err = f.WriteString(out)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_TARGET_CHECKSUM, err)
		_ = f.Close()

		return err
	}

	// sync and safe close
	_ = f.Sync()
	f.Close()

	return nil
}

func (me *Manager) createLogger() (out *reporter) {
	var err error

	defer func() {
		if r := recover(); r != nil {
			out = nil
		}
	}()

	out = &reporter{
		Log:      me.Log,
		DestPath: me.Path,
	}

	err = out.IsHealthy()
	if err != nil {
		return nil
	}

	out.Init()

	return out
}
