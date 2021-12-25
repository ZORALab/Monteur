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

package oshelper

import (
	"os"
	"time"
)

const (
	MAX_UID = int(^uint(0) >> 1)
	MAX_GID = int(^uint(0) >> 1)
)

// FileOwners get the UID and GID from a given FileInfo.
//
// Should an error is found, both UID and GID are set to MAX_UID and MAX_GID.
//
// For Windows, this function always return MAX_UID and MAX_GID.
func FileOwners(fi os.FileInfo) (uid int, gid int) {
	return _fileOwners(fi)
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
	return _fileTimestamps(fi)
}

// FileSetPlatformTime is to set timestamp for platform file.
//
// This function is only supported on Windows operating system. It will return
// `nil` for unsupported ones.
func FileSetPlatformTime(dest string, mTime time.Time) (err error) {
	return _fileSetPlatformTime(dest, mTime)
}
