// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
// Copyright 2021 Sebastiaan van Stijin (github@gone.nl)
// Copyright 2018 Daniel Nephin (dnephin@gmail.com)
// Copyright 2016 Allen Sun (shlallen1990@gmail.com)
// Copyright 2015 Darren Stahl (darst@microsoft.com)
// Copyright 2015 Zhang Wei (zhangwei_cs@qq.com)
// Copyright 2014 Victor Vieux (victorvieux@gmail.com)
// Copyright 2014 Kazuyoshi Kato (kato.kazuyoshi@gmail.com)
// Copyright 2014 Alexander Larsson (alexander.larsson@gmail.com)
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

//go:build !(windows || linux || freebsd)
// +build !windows,!linux,!freebsd

package commander

import (
	"time"
)

func ___setUtimesNano(dest string, ts []syscall.Timespec) (err error) {
	return fmt.Errorf("utime is not supported")
}

func ___setPlatformTime(dest string, modified time.Time) (err error) {
	return nil
}
