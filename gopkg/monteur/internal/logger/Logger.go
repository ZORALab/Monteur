// Copyright 2020 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2020 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
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

package logger

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

// Logger is the core structure for offering logging services.
//
// It has important unexported fields requires initializations. Hence, this
// structure is not safe to use the conventional `&struct{...}` creation. Please
// use `New()` function instead.
type Logger struct {
	mutex         *sync.Mutex
	outputWriters map[string]io.Writer
	statusWriters map[string]io.Writer

	preprocessor    func(string) string // assigned by user
	timestampFormat string              // assigned by user
}

// New is to create a Logger object safely with initialized internal field.
//
// It returns a Logger's struct pointer with initialized fields.
func New() *Logger {
	return &Logger{
		timestampFormat: ISO8601,
		preprocessor:    nil,
		mutex:           &sync.Mutex{},
		outputWriters:   map[string]io.Writer{},
		statusWriters:   map[string]io.Writer{},
	}
}

// IsHealthy is to check the logger is healthy for operations
//
// This is useful for situation where one needs a sit-rep of the logger status
// and checking the logger has been assigned into an interface object that is
// `nil` assignable. The latter is usually used for checking its existence
// inside that interface assignment.
func (s *Logger) IsHealthy() (err error) {
	if s.mutex == nil {
		return fmt.Errorf(ERROR_UNHEALTHY)
	}

	if s.outputWriters == nil {
		return fmt.Errorf(ERROR_UNHEALTHY)
	}

	if s.statusWriters == nil {
		return fmt.Errorf(ERROR_UNHEALTHY)
	}

	return nil
}

// Add adds a given writer to the Logger for simultenous logging.
//
// If there is an existing writer, it will be closed and overwritten with the
// newly given writer.
func (s *Logger) Add(w io.Writer, statusType StatusType, label string) error {
	// validate critical inputs
	if label == "" {
		return fmt.Errorf(ERROR_LABEL_EMPTY)
	}

	// lock to ensure there is no transaction
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// get the existing one
	var list *map[string]io.Writer

	switch {
	case statusType == TYPE_OUTPUT:
		list = &s.outputWriters
	default:
		list = &s.statusWriters
	}

	// create or overwrite writer
	(*list)[label] = w

	return nil
}

// Remove safely removes a Writer from the Logger.
//
// This function will do nothing if the given label is an empty string or
// invalid.
func (s *Logger) Remove(statusType StatusType, label string) {
	// validate critical inputs
	if label == "" {
		return
	}

	// lock to ensure there is no other transaction
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// get the existing one
	var list *map[string]io.Writer

	switch {
	case statusType == TYPE_OUTPUT:
		list = &s.outputWriters
	default:
		list = &s.statusWriters
	}

	// create or overwrite writer
	delete(*list, label)
}

// SetPreprocessor set a given preprocessing function for Logger.
//
// This is especially useful for those who wants to apply additional
// validations, filtering, or encryptions. The default is `nil` function which
// does nothing.
//
// The output must be a single string which will later be prepended with the
// standard timestamp and level tag. Therefore, if your output is non-printable
// data bytes (e.g. encrypted payload), you are **RECOMMENDED** to encode it
// into Base64 format.
//
// The input is the formatted string without standard timestamp and level tags.
//
// To disable this function, provide `nil` as its input.
func (s *Logger) SetPreprocessor(input func(string) string) {
	// lock to ensure there is no other transaction
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// set preprocessor function
	s.preprocessor = input
}

// SetTimestamp set a custom timestamp format for Logger.
//
// To restore back to default, simply do `SetTimestamp(logger.ISO8601)` again.
func (s *Logger) SetTimestamp(format string) {
	// validate input
	if format == "" {
		return
	}

	// lock to ensure there is no other transaction
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.timestampFormat = format
}

// Printf is the conventional print interface for Logger.
//
// This interface is fully compatible with `fmt.Printf(...)`. It will always
// return 0 and nil.
//
// This function does not use any level tag and shall always print to
// TYPE_STATUS log type.
func (s *Logger) Printf(format string, a ...interface{}) (n int, err error) {
	s.Logf(TYPE_STATUS, TAG_NO, format, a...)
	return 0, nil
}

// Println is the conventional print interface for Logger.
//
// This interface is fully compatible with `fmt.Println(...)`. It always return
// 0 and nil.
//
// This function does not use any level tag and shall always print to
// TYPE_STATUS log type.
func (s *Logger) Println(args ...interface{}) (n int, err error) {
	s.Logf(TYPE_STATUS, TAG_NO, fmt.Sprintln(args...)+"\n")
	return 0, nil
}

// Logln prints given log statement.
//
// The usage is similar to `fmt.Println(...)` with additional `statusType` and
// `tag level` fields.
func (s *Logger) Logln(statusType StatusType,
	level string, args ...interface{}) {
	s.Logf(statusType, level, fmt.Sprintln(args...)+"\n")
}

// Logf prints formatted log output.
//
// The usage is similar to `fmt.Printf(...)` with additional `statusType` and
// `tag level` fields.
func (s *Logger) Logf(statusType StatusType,
	level string,
	format string,
	a ...interface{}) {
	// lock to ensure there is no other transaction
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// generate output string
	output := fmt.Sprintf(format, a...)

	// run preprocessor if available
	if s.preprocessor != nil {
		output = s.preprocessor(output)
	}

	// prepend standard level tag
	if level != "" {
		output = "[ " + strings.ToUpper(level) + " ] " + output
	}

	// prepend timestamp
	output = time.Now().UTC().Format(s.timestampFormat) + " " + output

	// add file into the list
	switch {
	case statusType == TYPE_OUTPUT:
		for _, w := range s.outputWriters {
			_, _ = io.WriteString(w, output)
		}
	default:
		for _, w := range s.statusWriters {
			_, _ = io.WriteString(w, output)
		}
	}
}
