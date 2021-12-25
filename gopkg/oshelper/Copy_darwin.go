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

//go:build darwin
// +build darwin

package oshelper

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func _copyPipe(source string, dest string, fi os.FileInfo) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s: '%v'", ERROR_FILE_STAT, fi)
		}
	}()

	stat := fi.Sys().(*syscall.Stat_t)

	// create pipe file
	err = unix.Mkfifo(dest, uint32(stat.Mode))
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PIPE_CREATE, err)
	}

	// restore file permission
	err = os.Chmod(dest, fi.Mode())
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PIPE_PERM, err)
	}

	// restore timestamp
	return _restoreTimestamp(dest, fi)
}
