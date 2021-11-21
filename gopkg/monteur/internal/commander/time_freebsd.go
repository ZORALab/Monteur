// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
// Copyright 2021 Sebastiaan van Stijin (github@gone.nl)
// Copyright 2019 Tobias Klauser (tklauser@distanz.ch)
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

//go:build freebsd
// +build freebsd

package commander

import (
	"fmt"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func ___setUtimesNano(dest string, ts []syscall.Timespec) (err error) {
	uts := []unix.Timespec{
		unix.NsecToTimespec(syscall.TimespecToNsec(ts[0])),
		unix.NsecToTimespec(syscall.TimespecToNsec(ts[1])),
	}

	err = unix.UtimesNanoAt(unix.AT_FDCWD,
		desc,
		uts,
		unix.AT_SYMLINK_NOFOLLOW,
	)
	if err != nil && err != unix.ENOSYS {
		return fmt.Errorf("failed to set utime: %s", err)
	}

	return nil
}

func ___setPlatformTime(dest string, modified time.Time) (err error) {
	return nil
}
