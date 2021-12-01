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

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libcmd"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libworkspace"
)

func _acceptanceFilter(path string, info os.FileInfo, err error) (bool, error) {
	// return if err occurred
	if err != nil {
		return false, fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// ensures we only accepts regular file with .toml extension
	if filepath.Ext(path) != libmonteur.EXTENSION_TOML ||
		!info.Mode().IsRegular() {
		return false, nil
	}

	// accepted
	return true, nil
}

func _createCMDManager(l *liblog.Logger,
	w *libworkspace.Workspace,
	secrets *map[string]interface{},
	path string) (s *libcmd.Manager, err error) {
	l.Info("Processing %s...", path)

	s = &libcmd.Manager{
		Variables: map[string]interface{}{
			libmonteur.VAR_OS:      w.OS,
			libmonteur.VAR_ARCH:    w.ARCH,
			libmonteur.VAR_COMPUTE: w.ComputeSystem,
			libmonteur.VAR_HOME:    w.Filesystem.CurrentDir,
			libmonteur.VAR_ROOT:    w.Filesystem.RootDir,
			libmonteur.VAR_BASE:    w.Filesystem.BaseDir,
			libmonteur.VAR_LOG:     w.Filesystem.WorkspaceLogDir,
			libmonteur.VAR_CFG:     w.Filesystem.BinCfgDir,
			libmonteur.VAR_BIN:     w.Filesystem.BinDir,
			libmonteur.VAR_SECRETS: *secrets,
		},
	}

	switch w.Job {
	case "testing":
		s.Variables[libmonteur.VAR_TMP] = w.Filesystem.TestTMPDir
	case "publish":
		s.Variables[libmonteur.VAR_TMP] = w.Filesystem.PublishTMPDir
		s.Variables[libmonteur.VAR_DOC] = w.Filesystem.ComposeTMPDir
	case "compose":
		s.Variables[libmonteur.VAR_TMP] = w.Filesystem.ComposeTMPDir
	default:
		s.Variables[libmonteur.VAR_TMP] = w.Filesystem.WorkingDir
	}

	_logVariables(l, &s.Variables)

	l.Info("Decode Task Data from config file...")
	err = s.Parse(path)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	l.Success(libmonteur.LOG_SUCCESS)
	return s, nil
}

func _logVariables(l *liblog.Logger, list *map[string]interface{}) {
	l.Info("Inserting Task Variables...")
	for k, v := range *list {
		if k == libmonteur.VAR_SECRETS {
			l.Info("\"%s\": %v", k, libmonteur.LOG_FORMAT_REDACTED)
			continue
		}

		l.Info("\"%s\": %#v", k, v)
	}
	l.Success(libmonteur.LOG_SUCCESS)
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

func _initSecrets(list *map[string]interface{},
	l *liblog.Logger,
	path []string) {
	l.Info("Parsing secrets...")
	*list = libsecrets.GetSecrets(path)
	l.Success(libmonteur.LOG_SUCCESS)
}

func _initCMDSettings(x **libcmd.Run,
	l *liblog.Logger, path string) (err error) {
	l.Info("Initialize settings...")

	*x = &libcmd.Run{}
	err = (*x).Parse(path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	l.Success(libmonteur.LOG_SUCCESS)

	return nil
}

func _reportError(l *liblog.Logger, tag string, err error) int {
	l.Error("%s %s\n", tag, err)
	l.Sync()
	l.Close()
	return STATUS_ERROR
}
