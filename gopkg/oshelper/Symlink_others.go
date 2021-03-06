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

//go:build !(linux || freebsd)
// +build !linux,!freebsd

package oshelper

import (
	"os"
	"time"
)

func _symlinkTimestamps(fi os.FileInfo) (aTim, cTim, mTim time.Time) {
	return time.Time{}, time.Time{}, time.Time{} // not supported
}

func _symlinkChtimes(dest string, aTim time.Time, mTim time.Time) (err error) {
	return nil // not supported
}
