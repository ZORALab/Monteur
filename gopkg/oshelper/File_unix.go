// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
// Copyright 2020 Tobias Klauser (tklauser@distanz.ch)
// Copyright 2019 Kir Kolyshkin (kolyshkin@gmail.com)
// Copyright 2019 Dominic Yin (hi@ydcool.me)
// Copyright 2019 Tõnis Tiigi (tonistiigi@gmail.com)
// Copyright 2018 Maxim Ivanov
// Copyright 2017 Sargun Dhillon (sargun@sargun.me)
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

//go:build !(windows || darwin)
// +build !windows,!darwin

package oshelper

import (
	"os"
	"syscall"
	"time"
)

const (
	newLine = "\n"
)

// FileOwners get the UID and GID from a given FileInfo.
//
// Should an error is found, both UID and GID are set to MAX_UID and MAX_GID.
//
// For Windows, this function always return MAX_UID and MAX_GID.
func FileOwners(fi os.FileInfo) (uid int, gid int) {
	defer func() {
		if r := recover(); r != nil {
			uid = MAX_UID
			gid = MAX_GID
		}
	}()

	stat := fi.Sys().(*syscall.Stat_t)

	uid = int(stat.Uid)
	gid = int(stat.Gid)

	return uid, gid
}

// FileTimestamps get the file timestamps from FileInfo.
//
// It gets:
//   1. accessed time
//   2. changed time
//   3. modified time.
//
// Should any of the timestamp is invalid (outside of UNIX Epoch), the intital
// UNIX timestamp Epoch is set to 0.
//
// Should an error is found, all timestamps are set to UNIX timestamp Epoch 0.
func FileTimestamps(fi os.FileInfo) (accessed, changed, modified time.Time) {
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

	// supporting 32-bits
	//nolint:unconvert
	accessed = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	switch {
	case accessed.Before(unixMinTime):
		accessed = unixMinTime
	case accessed.After(unixMaxTime):
		accessed = unixMaxTime
	default:
	}

	// supporting 32-bits
	//nolint:unconvert
	changed = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	switch {
	case changed.Before(unixMinTime):
		changed = unixMinTime
	case changed.After(unixMaxTime):
		changed = unixMaxTime
	default:
	}

	// supporting 32-bits
	//nolint:unconvert
	modified = time.Unix(int64(stat.Mtim.Sec), int64(stat.Mtim.Nsec))
	switch {
	case modified.Before(unixMinTime):
		modified = unixMinTime
	case modified.After(unixMaxTime):
		modified = unixMaxTime
	default:
	}

	return accessed, changed, modified
}

// FileSetPlatformTime is to set timestamp for platform file.
//
// This function is only supported on Windows operating system. It will return
// `nil` for unsupported ones.
func FileSetPlatformTime(dest string, mTime time.Time) (err error) {
	return nil
}
