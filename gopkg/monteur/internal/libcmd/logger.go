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

package libcmd

import (
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

func reportError(log *liblog.Logger,
	ch chan conductor.Message,
	name string,
	format string, args ...interface{}) {
	if log != nil {
		log.Error(format, args...)
		log.Sync()
		log.Close()
	}

	if ch != nil {
		ch <- conductor.CreateError(name, format, args...)
	}
}

func reportOutput(log *liblog.Logger,
	ch chan conductor.Message,
	name string,
	format string, args ...interface{}) {
	if log != nil {
		log.Output(format, args...)
		log.Sync()
		log.Close()
	}

	if ch != nil {
		ch <- conductor.CreateOutput(name, format, args...)
	}
}

func reportStatus(log *liblog.Logger,
	ch chan conductor.Message,
	name string,
	format string, args ...interface{}) {
	if log != nil {
		log.Info(format, args...)
		log.Sync()
		log.Close()
	}

	if ch != nil {
		ch <- conductor.CreateStatus(name, format, args...)
	}
}

func reportDone(log *liblog.Logger, ch chan conductor.Message, name string) {
	if log != nil {
		log.Success(libmonteur.LOG_DONE + "\n")
		log.Sync()
		log.Close()
	}

	if ch != nil {
		ch <- conductor.CreateDone(name)
	}
}
