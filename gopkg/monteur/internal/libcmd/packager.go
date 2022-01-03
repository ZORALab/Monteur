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
	"context"
	"fmt"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libdeb"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libtargz"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libzip"
)

type packager struct {
	reportUp   chan conductor.Message
	log        *liblog.Logger
	thisSystem string

	variables    map[string]interface{}
	metadata     *libmonteur.TOMLMetadata
	dependencies []*commander.Dependency
	packages     map[string]*libmonteur.TOMLPackage
	cmd          []*libmonteur.TOMLAction
}

func (me *packager) Parse(path string) (err error) {
	// init temporary raw input variables
	dep := []*libmonteur.TOMLDependency{}
	fmtVar := map[string]interface{}{}
	cmd := []*libmonteur.TOMLAction{}

	// initialize all important variables
	me.metadata = &libmonteur.TOMLMetadata{}
	me.dependencies = []*commander.Dependency{}
	me.cmd = []*libmonteur.TOMLAction{}
	me.packages = map[string]*libmonteur.TOMLPackage{}

	// construct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
		Dependencies *[]*libmonteur.TOMLDependency
		Packages     map[string]*libmonteur.TOMLPackage
		CMD          *[]*libmonteur.TOMLAction
	}{
		Metadata:     me.metadata,
		Variables:    me.variables,
		FMTVariables: &fmtVar,
		Dependencies: &dep,
		Packages:     me.packages,
		CMD:          &cmd,
	}

	// decode
	err = toml.DecodeFile(path, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// sanitize
	err = sanitizeMetadata(me.metadata, path)
	if err != nil {
		return err
	}

	err = libmonteur.SanitizeVariables(&me.variables, &fmtVar)
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = sanitizeDeps(dep, &me.dependencies, me.thisSystem, me.variables)
	if err != nil {
		return err
	}

	err = sanitizeCMD(cmd, &me.cmd, me.thisSystem)
	if err != nil {
		return err
	}

	// init
	err = initializeLogger(&me.log, me.metadata.Name, me.variables)
	if err != nil {
		return err
	}

	return nil
}

// Run executes the full run-job.
func (me *packager) Run(ctx context.Context, ch chan conductor.Message) {
	var err error

	me.log.Info("Run Task Now: " + libmonteur.LOG_SUCCESS + "\n")
	me.reportUp = ch

	err = me.runPackaging()
	if err != nil {
		me.reportError(err)
		return
	}

	me.reportDone()
}

func (me *packager) runPackaging() (err error) {
	for _, pkg := range me.packages {
		// copy the original variables into the new variable list
		variables := map[string]interface{}{}
		for k, v := range me.variables {
			variables[k] = v
		}

		err = processPackageVariables(pkg, &variables)
		if err != nil {
			return err
		}

		// ensure package has files
		if len(pkg.Files) == 0 {
			return fmt.Errorf(libmonteur.ERROR_PACKAGER_FILES_MISSING)
		}

		// execute built-in package preprations
		err = me._preparePackage(pkg, variables)
		if err != nil {
			return err
		}

		// execute given commands
		err = me._runCMD(variables)
		if err != nil {
			return err
		}
	}

	return nil
}

func (me *packager) _runCMD(variables map[string]interface{}) (err error) {
	me.log.Info("Executing Packaging CMD now...")
	task := &executive{
		log:       me.log,
		variables: variables,
		orders:    me.cmd,
		fxSTDOUT:  me.reportOutput,
		fxSTDERR:  me.reportStatus,
	}

	err = task.Exec()
	if err != nil {
		return err
	}

	me.log.Info("Executing Packaging CMD ➤ DONE\n\n")
	return nil
}

func (me *packager) _preparePackage(pkg *libmonteur.TOMLPackage,
	variables map[string]interface{}) (err error) {
	me.log.Info("Executing Packaging Preparations now...")

	switch me.metadata.Type {
	case libmonteur.PACKAGE_TARGZ:
		err = libtargz.Package(pkg, &variables, me.log)
	case libmonteur.PACKAGE_ZIP:
		err = libzip.Package(pkg, &variables, me.log)
	case libmonteur.PACKAGE_DEB_MANUAL:
		_, err = libdeb.Prepare(pkg, &variables, me.log)
	case libmonteur.PACKAGE_MANUAL:
	default:
		err = fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_PACKAGER_TYPE_UNKNOWN,
			me.metadata.Type,
		)
	}

	if err != nil {
		return err
	}

	// end the preparations
	me.log.Info("Executing Packaging Preparations ➤ DONE\n\n")
	return nil
}

// Name is to return the task name
func (me *packager) Name() string {
	return me.metadata.Name
}

func (me *packager) reportStatus(format string, args ...interface{}) {
	reportStatus(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *packager) reportError(err error) {
	reportError(me.log, me.reportUp, me.metadata.Name, "%s", err)
}

func (me *packager) reportOutput(format string, args ...interface{}) {
	reportOutput(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *packager) reportDone() {
	reportDone(me.log, me.reportUp, me.metadata.Name)
}
