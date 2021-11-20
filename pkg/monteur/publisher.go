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

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libcmd"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libworkspace"
)

type publisher struct {
	workspace *libworkspace.Workspace
	secrets   map[string]interface{}
	settings  *libcmd.Run
	workers   map[string]*libcmd.Manager
}

func (fx *publisher) Build() (statusCode int) {
	err := fx._init()
	if err != nil {
		return fx._reportError(err)
	}

	err = filepath.Walk(fx.workspace.Filesystem.PublishBuilderConfigDir,
		fx._filterPublisher)
	if err != nil {
		return fx._reportError(err)
	}

	for _, p := range fx.workers {
		err = p.Run()
		if err != nil {
			return fx._reportError(err)
		}
	}

	return STATUS_OK
}

func (fx *publisher) Publish() (statusCode int) {
	err := fx._init()
	if err != nil {
		return fx._reportError(err)
	}

	err = filepath.Walk(fx.workspace.Filesystem.PublishConfigDir,
		fx._filterPublisher)
	if err != nil {
		return fx._reportError(err)
	}

	for _, p := range fx.workers {
		err = p.Run()
		if err != nil {
			return fx._reportError(err)
		}
	}

	return STATUS_OK
}

func (fx *publisher) _filterPublisher(path string,
	info os.FileInfo, err error) error {
	var s *libcmd.Manager

	// return if err occurred
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// ensures we only accept the correct regular file with .toml extension
	if filepath.Ext(path) != libmonteur.EXTENSION_TOML || info.IsDir() {
		return nil
	}

	// initialize TOML Parser object
	//nolint:lll
	s = &libcmd.Manager{
		Variables: map[string]interface{}{
			libmonteur.VAR_OS:      fx.workspace.OS,
			libmonteur.VAR_ARCH:    fx.workspace.ARCH,
			libmonteur.VAR_COMPUTE: fx.workspace.ComputeSystem,
			libmonteur.VAR_TMP:     fx.workspace.Filesystem.PublishTMPDir,
			libmonteur.VAR_BIN:     fx.workspace.Filesystem.BinDir,
			libmonteur.VAR_CFG:     fx.workspace.Filesystem.BinCfgDir,
			libmonteur.VAR_ROOT:    fx.workspace.Filesystem.RootDir,
			libmonteur.VAR_HOME:    fx.workspace.Filesystem.CurrentDir,
			libmonteur.VAR_SECRETS: fx.secrets,
		},
	}

	// decode the publisher toml file
	err = s.Parse(path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// save successful publisher data into list for further processig
	fx.workers[s.Metadata.Name] = s

	return nil
}

func (fx *publisher) _init() (err error) {
	fx.settings = &libcmd.Run{}
	fx.workspace = &libworkspace.Workspace{}
	fx.workers = map[string]*libcmd.Manager{}

	// initialize workspace
	err = fx.workspace.Init()
	if err != nil {
		return err //nolint:wrapcheck
	}

	// initialize secrets and parse every one of them
	fx.secrets = libsecrets.GetSecrets(fx.workspace.Filesystem.SecretsDir)

	// initialize settings
	err = fx.settings.Parse(fx.workspace.Filesystem.PublishTOMLFile)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}

func (fx *publisher) _reportError(err error) int {
	fmt.Fprintf(os.Stdout, "%s %s\n", libmonteur.ERROR_PUBLISH, err)
	return STATUS_ERROR
}
