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

package monteur

import (
	"fmt"
	"os"
	"path/filepath"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libworkspace"
)

func _logVariables(l *liblog.Logger, list *map[string]interface{}) {
	l.Info("Inserting Task Variables...")
	for k, v := range *list {
		if k == libmonteur.VAR_SECRETS {
			l.Info("\"%s\": %v", k, libmonteur.LOG_FORMAT_REDACTED)
			continue
		}

		l.Info("\"%s\": %#v", k, v)
	}
	l.Info(libmonteur.LOG_SUCCESS)
}

func _initLogger(l **liblog.Logger, w *libworkspace.Workspace) (err error) {
	*l = &liblog.Logger{ToTerminal: true}
	(*l).Init()

	err = (*l).Add(liblog.TYPE_STATUS, filepath.Join(
		w.Filesystem.WorkspaceLogDir,
		libmonteur.FILE_LOG_PREFIX_JOB+libmonteur.FILE_LOG_STATUS,
	))
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = (*l).Add(liblog.TYPE_OUTPUT, filepath.Join(
		w.Filesystem.WorkspaceLogDir,
		libmonteur.FILE_LOG_PREFIX_JOB+libmonteur.FILE_LOG_OUTPUT,
	))
	if err != nil {
		return err //nolint:wrapcheck
	}

	(*l).Info("\n%s", w.String())

	return nil
}

func _initWorkspace(job string, w **libworkspace.Workspace) (err error) {
	*w = &libworkspace.Workspace{Job: job}

	err = (*w).Init()
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}

func _reportError(l *liblog.Logger, tag string, err error) int {
	if l != nil {
		l.Error("%s %s\n", tag, err)
		l.Sync()
		l.Close()
	} else {
		fmt.Fprintf(os.Stderr, "%s %s\n", tag, err)
	}
	return STATUS_ERROR
}
