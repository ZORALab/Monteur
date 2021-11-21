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

package commander

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func __copyFileUTimestampNano(dest string,
	info os.FileInfo, stat *syscall.Stat_t) (err error) {
	var ok bool

	if stat == nil {
		stat, ok = info.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("unable to obtain stat from FileInfo")
		}
	}

	timespec := []syscall.Timespec{stat.Atim, stat.Mtim}

	err = ___setUtimesNano(dest, timespec)
	if err != nil {
		return err
	}

	return nil
}

func __copyFileTimestamp(dest string,
	info os.FileInfo, stat *syscall.Stat_t) (err error) {
	var ok bool

	if stat == nil {
		stat, ok = info.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("unable to obtain stat from FileInfo")
		}
	}

	unixMinTime := time.Unix(0, 0)
	unixMaxTime := unixMinTime.Add(1<<63 - 1)
	accessed := time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
	modified := time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec)

	if accessed.Before(unixMinTime) || accessed.After(unixMaxTime) {
		accessed = unixMinTime
	}

	if modified.Before(unixMinTime) || modified.After(unixMaxTime) {
		modified = unixMinTime
	}

	err = os.Chtimes(dest, accessed, modified)
	if err != nil {
		return fmt.Errorf("%s: %s", "error changing time", err)
	}

	err = ___setPlatformTime(dest, modified)
	if err != nil {
		return err
	}

	return nil
}
