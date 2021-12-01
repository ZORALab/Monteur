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

package liblog

import (
	"fmt"
	"os"
	"path/filepath"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/logger"
)

const (
	TYPE_STATUS = logger.TYPE_STATUS
	TYPE_OUTPUT = logger.TYPE_OUTPUT
)

type Logger struct {
	executor      *logger.Logger
	statusWriters map[string]*os.File
	outputWriters map[string]*os.File

	ToTerminal bool
	DebugMode  bool
}

// Init is to initialize the logger for use.
func (log *Logger) Init() {
	log.executor = logger.New()
	log.statusWriters = map[string]*os.File{}
	log.outputWriters = map[string]*os.File{}
}

// IsHealthy is to check the status of the logger.
func (log *Logger) IsHealthy() (err error) {
	if log.executor == nil {
		return fmt.Errorf(libmonteur.ERROR_LOG_UNHEALTHY)
	}

	if log.statusWriters == nil {
		return fmt.Errorf(libmonteur.ERROR_LOG_UNHEALTHY)
	}

	if log.outputWriters == nil {
		return fmt.Errorf(libmonteur.ERROR_LOG_UNHEALTHY)
	}

	return nil
}

// Add is to add a given filepath for creating the io.Writer for logger.
func (log *Logger) Add(logType logger.StatusType, path string) (err error) {
	if path == "" {
		return fmt.Errorf(libmonteur.ERROR_LOG_PATH_EMPTY)
	}

	_ = os.MkdirAll(filepath.Dir(path), 0755)

	f, err := logger.CreateFile(path)
	if err != nil {
		return fmt.Errorf("%s: %s", libmonteur.ERROR_LOG_PREPARE, err)
	}

	err = log.executor.Add(f, logType, path)
	if err != nil {
		return fmt.Errorf("%s: %s", libmonteur.ERROR_LOG_PREPARE, err)
	}

	switch logType {
	case TYPE_OUTPUT:
		log.outputWriters[path] = f
	case TYPE_STATUS:
		fallthrough
	default:
		log.statusWriters[path] = f
	}

	return nil
}

// Sync is to sync all the Writers to ensure all data are written to files.
func (log *Logger) Sync() {
	for _, f := range log.outputWriters {
		_ = f.Sync()
	}

	for _, f := range log.statusWriters {
		_ = f.Sync()
	}
}

// Close is to close all the Writers for the logger safely.
func (log *Logger) Close() {
	for _, f := range log.outputWriters {
		_ = f.Sync()
		f.Close()
	}

	for _, f := range log.statusWriters {
		_ = f.Sync()
		f.Close()
	}
}

// Error is to log an error statement straight to status and output types logs.
func (log *Logger) Error(format string, a ...interface{}) {
	format += "\n"
	log.executor.Logf(logger.TYPE_STATUS, logger.TAG_ERROR, format, a...)
	log.executor.Logf(logger.TYPE_OUTPUT, logger.TAG_ERROR, format, a...)

	if !log.ToTerminal {
		return
	}

	fmt.Fprintf(os.Stdout, format, a...)
}

// Warning is to log a warning statement straight to status type logs.
func (log *Logger) Warning(format string, a ...interface{}) {
	format += "\n"
	log.executor.Logf(logger.TYPE_STATUS, logger.TAG_WARNING, format, a...)

	if !log.ToTerminal {
		return
	}

	fmt.Fprintf(os.Stderr, format, a...)
}

// Info is to log an info statement straight to status type logs.
func (log *Logger) Info(format string, a ...interface{}) {
	format += "\n"
	log.executor.Logf(logger.TYPE_STATUS, logger.TAG_INFO, format, a...)

	if !log.ToTerminal {
		return
	}

	fmt.Fprintf(os.Stderr, format, a...)
}

// Success is to log a success statement straight to status type logs.
func (log *Logger) Success(format string, a ...interface{}) {
	format += "\n"
	log.executor.Logf(logger.TYPE_STATUS, logger.TAG_SUCCESS, format, a...)

	if !log.ToTerminal {
		return
	}

	fmt.Fprintf(os.Stderr, format, a...)
}

// Debug is to log a debug statement straight to status type logs.
func (log *Logger) Debug(format string, a ...interface{}) {
	if !log.DebugMode {
		return
	}

	format += "\n"
	log.executor.Logf(logger.TYPE_STATUS, logger.TAG_DEBUG, format, a...)

	if !log.ToTerminal {
		return
	}

	fmt.Fprintf(os.Stderr, format, a...)
}

// Output is to log an output statements straight to status and output logs.
func (log *Logger) Output(format string, a ...interface{}) {
	format = "[ OUTPUT ] " + format + "\n"
	log.executor.Logf(logger.TYPE_STATUS, logger.TAG_NO, format, a...)
	log.executor.Logf(logger.TYPE_OUTPUT, logger.TAG_NO, format, a...)

	if !log.ToTerminal {
		return
	}

	fmt.Fprintf(os.Stdout, format, a...)
}

// Logf is to log a raw statement straight to the selected logs' type.
func (log *Logger) Logf(logType logger.StatusType,
	format string, a ...interface{}) {
	log.executor.Logf(logType, logger.TAG_NO, format, a...)

	if !log.ToTerminal {
		return
	}

	if logType == logger.TYPE_OUTPUT {
		fmt.Fprintf(os.Stdout, format, a...)
		return
	}

	fmt.Fprintf(os.Stderr, format, a...)
}
