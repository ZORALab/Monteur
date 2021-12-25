// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
// Copyright 2021 Sebastiaan van Stijin (github@gone.nl)
// Copyright 2018 Daniel Nephin (dnephin@gmail.com)
// Copyright 2017 Christopher Jones (ophj@linux.vnet.ibm.com)
// Copyright 2016 Stefan J. Wernli (swernli@microsoft.com)
// Copyright 2015 Dustin H (https://github.com/djherbis)
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

package oshelper

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

const (
	newLine = "\r\n"
)

func _fileOwners(fi os.FileInfo) (uid int, gid int) {
	return MAX_UID, MAX_GID
}

func _fileTimestamps(fi os.FileInfo) (accessed, changed, modified time.Time) {
	unixMinTime := time.Unix(0, 0)
	unixMaxTime := unixMinTime.Add(1<<63 - 1)

	defer func() {
		if r := recover(); r != nil {
			accessed = unixMinTime
			changed = unixMinTime
			modified = unixMinTime
		}
	}()

	stat := fi.Sys().(*syscall.Win32FileAttributeData)

	accessed = time.Unix(0, stat.LastAccessTime.Nanoseconds())
	switch {
	case accessed.Before(unixMinTime):
		accessed = unixMinTime
	case accessed.After(unixMaxTime):
		accessed = unixMaxTime
	}

	changed = time.Unix(0, stat.CreationTime.Nanoseconds())
	switch {
	case changed.Before(unixMinTime):
		changed = unixMinTime
	case changed.After(unixMaxTime):
		changed = unixMaxTime
	}

	modified = time.Unix(0, stat.LastWriteTime.Nanoseconds())
	switch {
	case modified.Before(unixMinTime):
		modified = unixMinTime
	case modified.After(unixMaxTime):
		modified = unixMaxTime
	}

	return accessed, changed, modified
}

func _fileSetPlatformTime(dest string, mTime time.Time) (err error) {
	var file windows.Handle
	var path *uint16

	path, err = windows.UTF16PtrFromString(dest)
	if err != nil {
		return fmt.Errorf("%s: %s", "error processing file", err)
	}

	file, err = windows.CreateFile(path,
		windows.FILE_WRITE_ATTRIBUTES,
		windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_FLAG_BACKUP_SEMANTICS,
		0,
	)
	if err != nil {
		return fmt.Errorf("%s: %s", "error creating platform file", err)
	}
	defer windows.Close(file)

	timeSpec := windows.NsecToTimespec(mTime.UnixNano())
	timestamp := windows.NsecToFiletime(windows.TimespecToNsec(timeSpec))

	return windows.SetFileTime(file, &timestamp, nil, nil)
}
