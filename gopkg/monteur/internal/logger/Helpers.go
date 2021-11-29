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
	"os"
)

// CreateFile is for creating a writable os.File with given path.
//
// The opened file can:
//   1. create if file does not exists.
//   2. append if file exists.
//   3. only open for write only.
//
// The UNIX file permission is 600 (owner with read write permission only) for
// safety control.
//
// This helper function does not perform any automated file close. Hence, it is
// entirely your responsibility to call its closure (e.g. `defer f.Close()`) in
// your program.
func CreateFile(path string) (f *os.File, err error) {
	// validate the given input
	if path == "" {
		return nil, fmt.Errorf(ERROR_FILE_MISSING)
	}

	i, err := os.Stat(path)
	if !os.IsNotExist(err) && !i.Mode().IsRegular() {
		return nil, fmt.Errorf(ERROR_FILE_INCOMPATIBLE)
	}

	// open file for read
	f, err = os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		PERMISSION_FILE_LOG,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", ERROR_FILE_OPEN_FAILED, err)
	}

	// return opened file
	return f, nil
}

// CreateStderr is for getting the STDERR os.File for terminal printing.
func CreateStderr() (f *os.File) {
	return os.Stderr
}

// CreateStdout is for getting the STDOUT os.File for terminal printing.
func CreateStdout() (f *os.File) {
	return os.Stdout
}
