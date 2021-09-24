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

//nolint:typecheck
import (
	"fmt"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/endec/toml"
)

// App is the main software information consolidated in a data structure.
//
// This data structure is unsafe to create using the conventional &struct{}
// manner. It contains elements that requires independent initialization. Hence,
// please use NewApp() function to create one.
type App struct {
	Metadata *Metadata
}

// NewApp() is to create an initialized App data structure returning as pointer.
//
// This function initializes the *App structure so that it's safe to use upon
// creations.
func NewApp() *App {
	return &App{
		Metadata: &Metadata{},
	}
}

// ParseWorkspace() is the method to parse the repository workspace settings.
//
// This single method shall parse and validate the entire repository before
// executing any CI jobs.
func (a *App) ParseWorkspace() error {
	err := a.parseTOML("./../.configs/monteur/appdata/en.toml", true)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) parseTOML(path string, update bool) (err error) {
	if path == "" {
		return fmt.Errorf("missing app toml data file")
	}

	err = toml.DecodeFile(path, &a, nil)
	if err != nil {
		return err
	}

	if update {
		fmt.Printf("TESTING: %#v\n", a.Metadata.Name)
	}

	return nil
}
