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

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libworkspace"
)

type publisher struct {
	workspace *libworkspace.Workspace
	secrets   map[string]interface{}
}

func (fx *publisher) Build() (statusCode int) {
	err := fx._init()
	if err != nil {
		fx._reportError(err)
	}

	return STATUS_OK
}

func (fx *publisher) Publish() (statusCode int) {
	err := fx._init()
	if err != nil {
		fx._reportError(err)
	}

	return STATUS_OK
}

func (fx *publisher) _reportError(err error) int {
	fmt.Fprintf(os.Stdout, "%s %s\n", libmonteur.ERROR_PUBLISH, err)
	return STATUS_ERROR
}

func (fx *publisher) _init() (err error) {
	fx.workspace = &libworkspace.Workspace{}

	// initialize workspace
	err = fx.workspace.Init()
	if err != nil {
		return err //nolint:wrapcheck
	}

	// initialize secrets and parse every one of them
	fx.secrets = libsecrets.GetSecrets(fx.workspace.Filesystem.SecretsDir)

	return nil
}
