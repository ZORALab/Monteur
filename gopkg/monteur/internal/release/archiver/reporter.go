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
	"sync"
)

type reporter struct {
	// assign by manager
	log Logger

	// internal usage
	mutex     *sync.RWMutex
	checksums []string
}

func (me *reporter) Init() {
	me.mutex = &sync.RWMutex{}
	me.checksums = []string{}
}

func (me *reporter) Info(format string, args ...interface{}) {
	var err error

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// log output
	err = me.IsHealthy()
	if err == nil {
		me.log.Info(format, args...)
	}
}

func (me *reporter) Success(format string, args ...interface{}) {
	var err error

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// log output
	err = me.IsHealthy()
	if err == nil {
		me.log.Success(format, args...)
	}
}

func (me *reporter) Warning(format string, args ...interface{}) {
	var err error

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// log output
	err = me.IsHealthy()
	if err == nil {
		me.log.Warning(format, args...)
	}
}

func (me *reporter) Error(format string, args ...interface{}) {
	var err error

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// log output
	err = me.IsHealthy()
	if err == nil {
		me.log.Error(format, args...)
	}
}

func (me *reporter) Output(format string, args ...interface{}) {
	var err error

	out := fmt.Sprintf(format, args...)

	// acquire lock
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// log output
	err = me.IsHealthy()
	if err == nil {
		me.log.Output(format, args...)
	}

	// append to checksum
	me.checksums = append(me.checksums, out)
}

func (me *reporter) IsHealthy() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("logger is absent")
		}
	}()

	// test run IsHealthy()
	err = me.log.IsHealthy()

	return err //nolint:wrapcheck
}
