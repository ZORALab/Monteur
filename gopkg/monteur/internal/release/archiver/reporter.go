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
	"path/filepath"
	"strings"
	"sync"
)

type reporter struct {
	// assign by manager
	Log      Logger
	DestPath string

	// internal usage
	mutex     *sync.RWMutex
	checksums []*metadata
}

func (me *reporter) Init() {
	me.mutex = &sync.RWMutex{}
	me.checksums = []*metadata{}
}

func (me *reporter) Info(format string, args ...interface{}) {
	var err error

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// Log output
	err = me.IsHealthy()
	if err == nil {
		me.Log.Info(format, args...)
	}
}

func (me *reporter) Success(format string, args ...interface{}) {
	var err error

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// Log output
	err = me.IsHealthy()
	if err == nil {
		me.Log.Success(format, args...)
	}
}

func (me *reporter) Warning(format string, args ...interface{}) {
	var err error

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// Log output
	err = me.IsHealthy()
	if err == nil {
		me.Log.Warning(format, args...)
	}
}

func (me *reporter) Error(format string, args ...interface{}) {
	var err error

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// Log output
	err = me.IsHealthy()
	if err == nil {
		me.Log.Error(format, args...)
	}
}

func (me *reporter) Output(format string, args ...interface{}) {
	var list []string
	var out, filename, hash, url string
	var err error
	var meta *metadata

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// Log output
	err = me.IsHealthy()
	if err == nil {
		me.Log.Output(format, args...)
	}

	// process output string
	out = fmt.Sprintf(format, args...)
	list = strings.Split(out, "➤ ")
	out = list[len(list)-1]
	out = strings.TrimRight(out, "\r\n ")
	out = strings.TrimLeft(out, "\r\n ")
	list = strings.Split(out, " ")
	filename = list[1]
	hash = list[0]

	// process url
	url, _ = filepath.Rel(me.DestPath, filename)
	filename = filepath.Base(filename)

	// construct metadata
	meta = &metadata{
		Filename: filename,
		Hash:     hash,
		URL:      url,
	}

	// append output to checksum
	me.checksums = append(me.checksums, meta)
}

func (me *reporter) IsHealthy() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Logger is absent")
		}
	}()

	// test run IsHealthy()
	err = me.Log.IsHealthy()

	return err //nolint:wrapcheck
}
