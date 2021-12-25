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

// SymlinkTimestamps obtain the timestamp from the link itself and not the file.
//
// It shall return accessed time, changed time, and modified time.
//
// Should an error occurs, all timestamps are return as `nil`.
func SymlinkTimestamps(fi os.FileInfo) (aTime, cTime, mTime time.Time) {
	return _symlinkTimestamps(fi)
}

// SymlinkChTimes changes the timestamp on a symlink itself and not destination.
//
// It is only supported on FreeBSD and Linux operating systems. Other
// unsupported OS shall return `nil`.
func SymlinkChtimes(dest string, aTime time.Time, mTime time.Time) (err error) {
	return _symlinkChtimes(dest, aTime, mTime)
}
