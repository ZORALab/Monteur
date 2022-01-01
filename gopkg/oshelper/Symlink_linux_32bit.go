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

//go:build linux && (386 || arm)
// +build linux
// +build 386 arm

package oshelper

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func _symlinkTimestamps(fi os.FileInfo) (aTim, cTim, mTim time.Time) {
	defer func() {
		if r := recover(); r != nil {
			aTim = time.Time{}
			cTim = time.Time{}
			mTim = time.Time{}
		}
	}()

	stat := fi.Sys().(*syscall.Stat_t)

	aTim = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	cTim = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	mTim = time.Unix(int64(stat.Mtim.Sec), int64(stat.Mtim.Nsec))

	return aTim, cTim, mTim
}

func _symlinkChtimes(dest string, aTim time.Time, mTim time.Time) (err error) {
	uts := []unix.Timespec{
		unix.NsecToTimespec(aTim.Unix()),
		unix.NsecToTimespec(mTim.Unix()),
	}

	err = unix.UtimesNanoAt(unix.AT_FDCWD,
		dest,
		uts,
		unix.AT_SYMLINK_NOFOLLOW,
	)
	if err != nil && err != unix.ENOSYS {
		err = fmt.Errorf("failed to set utime: %s", err)
	}

	return err
}
