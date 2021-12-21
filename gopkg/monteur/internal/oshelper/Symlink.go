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

package oshelper

import (
	"os"
	"syscall"
)

// SymlinkTimestamps obtain the timestamp from the link itself and not the file.
//
// It shall return accessed time, changed time, and modified time.
//
// Should an error occurs, all timestamps are return as `nil`.
func SymlinkTimestamps(fi os.FileInfo) (aTime, cTime, mTime *syscall.Timespec) {
	defer func() {
		if r := recover(); r != nil {
			aTime = nil
			cTime = nil
			mTime = nil
		}
	}()

	stat := fi.Sys().(*syscall.Stat_t)

	aTime = &stat.Atim
	cTime = &stat.Ctim
	mTime = &stat.Mtim

	return aTime, cTime, mTime
}
