// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
// Copyright 2021 Sebastiaan van Stijin (github@gone.nl)
// Copyright 2018 Daniel Nephin (dnephin@gmail.com)
// Copyright 2017 Christopher Jones (ophj@linux.vnet.ibm.com)
// Copyright 2016 Stefan J. Wernli (swernli@microsoft.com)
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

//go:build windows
// +build windows

package commander

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows"
)

func ___setUtimesNano(dest string, ts []syscall.Timespec) (err error) {
	return fmt.Errorf("utime is not supported")
}

func ___setPlatformTime(dest string, modified time.Time) (err error) {
	path, err = windows.UTFPtrFromString(dest)
	if err != nil {
		return fmt.Errorf("%s: %s",
			"failed to set platform time",
			err,
		)
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
		return fmt.Errorf("%s: %s",
			"failed to create platform file",
			err,
		)
	}
	defer windows.Close(file)

	timeSpec := windows.NsecToTimespec(modified.UnixNano())
	timestamp := windows.NsecToFiletime(windows.TimespecToNsec(timeSpec))

	return windows.SetFileTime(file, &timestamp, nil, nil)
}
