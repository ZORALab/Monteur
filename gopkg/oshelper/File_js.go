// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
// Copyright 2020 Tobias Klauser (tklauser@distanz.ch)
// Copyright 2019 Kir Kolyshkin (kolyshkin@gmail.com)
// Copyright 2019 Dominic Yin (hi@ydcool.me)
// Copyright 2019 Tõnis Tiigi (tonistiigi@gmail.com)
// Copyright 2018 Maxim Ivanov
// Copyright 2017 Sargun Dhillon (sargun@sargun.me)
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
	"os"
	"syscall"
	"time"
)

const (
	newLine = "\n"
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

	stat := fi.Sys().(*syscall.Stat_t)

	accessed = time.Unix(stat.Atime, stat.AtimeNsec)
	switch {
	case accessed.Before(unixMinTime):
		accessed = unixMinTime
	case accessed.After(unixMaxTime):
		accessed = unixMaxTime
	default:
	}

	changed = time.Unix(stat.Ctime, stat.CtimeNsec)
	switch {
	case changed.Before(unixMinTime):
		changed = unixMinTime
	case changed.After(unixMaxTime):
		changed = unixMaxTime
	default:
	}

	modified = time.Unix(stat.Mtime, stat.MtimeNsec)
	switch {
	case modified.Before(unixMinTime):
		modified = unixMinTime
	case modified.After(unixMaxTime):
		modified = unixMaxTime
	default:
	}

	return accessed, changed, modified
}

func _fileSetPlatformTime(dest string, mTime time.Time) (err error) {
	return nil
}
