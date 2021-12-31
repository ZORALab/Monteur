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
	"strings"
	"sync"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/checksum"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

const (
	PERMISSION_FILE = 0655
	PERMISSION_DIR  = 0755
)

const (
	CHECKSUM_SHA256        = checksum.HASHER_SHA256
	CHECKSUM_SHA512        = checksum.HASHER_SHA512
	CHECKSUM_SHA512_TO_256 = checksum.HASHER_SHA512_TO_SHA256
)

type Manager struct {
	mutex        *sync.Mutex
	jobs         map[string]conductor.Job
	dataTemplate string
	dataFilename string

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
		me.dataFilename = _FILENAME_CSV
	case FORMAT_TXT:
		me.dataTemplate = _TEMPLATE_TXT
		me.dataFilename = _FILENAME_TXT
	case FORMAT_TOML:
		me.dataTemplate = _TEMPLATE_TOML
		me.dataFilename = _FILENAME_TOML
	default:
		return fmt.Errorf("%s: %d", ERROR_FORMAT_UNSUPPORTED, me.Format)
	}

	return nil
}

func (me *Manager) sanitizeChecksum() (err error) {
	switch me.Checksum {
	case CHECKSUM_SHA256:
	case CHECKSUM_SHA512:
	case CHECKSUM_SHA512_TO_256:
	default:
		return fmt.Errorf("%s: %d",
			ERROR_CHECKSUM_UNSUPPORTED,
			me.Checksum,
		)
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
	mode = info.Mode()
	switch {
	case err == nil && mode.IsRegular():
	case err == nil && !mode.IsRegular():
		return fmt.Errorf(ERROR_TARGET_INVALID)
	case os.IsNotExist(err):
		return fmt.Errorf("%s: %s", ERROR_TARGET_MISSING, err)
	default:
		return fmt.Errorf(ERROR_PATH)
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
	err = me.writeChecksumFile(strings.Join(logger.checksums, "\n"))
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

func (me *Manager) writeChecksumDataFile(data []string) (err error) {
	var f *os.File
	var dataPath, url, out string
	var list []string
	var variables map[string]interface{}

	if me.DataPath == "" {
		return nil
	}

	if len(data) == 0 {
		return fmt.Errorf(ERROR_CHECKSUM_EMPTY)
	}

	url = filepath.Join(VersionToDir(me.Version), me.dataFilename)
	dataPath = filepath.Join(me.DataPath, url)

	err = os.MkdirAll(filepath.Dir(dataPath), PERMISSION_DIR)
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_TARGET_DATA_DIR_CREATE,
			err,
		)
	}

	f, err = os.OpenFile(dataPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		PERMISSION_FILE,
	)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_TARGET_CHECKSUM, err)
	}

	for _, v := range data {
		list = strings.Split(v, " ")

		variables = map[string]interface{}{
			"Filename": list[0],
			"Hash":     list[1],
			"URL":      url,
		}

		// template text
		out, err = templater.String(me.dataTemplate, variables)
		if err != nil {
			return fmt.Errorf("%s (%s): %s",
				ERROR_TARGET_CHECKSUM,
				list[0],
				err,
			)
		}

		// write into file
		_, err = f.WriteString(out)
		if err != nil {
			return fmt.Errorf("%s (%s): %s",
				ERROR_TARGET_CHECKSUM,
				list[0],
				err,
			)
		}
	}

	_ = f.Sync()
	f.Close()

	return nil
}

func (me *Manager) writeChecksumFile(data string) (err error) {
	var path string
	var f *os.File

	path = filepath.Join(me.Path,
		VersionToDir(me.Version),
		_FILENAME_TXT,
	)

	f, err = os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		PERMISSION_FILE,
	)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_TARGET_CHECKSUM, err)
	}

	_, err = f.WriteString(data)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_TARGET_CHECKSUM, err)
		_ = f.Close()

		return err
	}

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

	out = &reporter{log: me.Log}

	err = out.IsHealthy()
	if err != nil {
		return nil
	}

	out.Init()

	return out
}
