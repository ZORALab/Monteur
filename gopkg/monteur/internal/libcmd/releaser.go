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
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libarchiver"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

type releaser struct {
	reportUp   chan conductor.Message
	log        *liblog.Logger
	thisSystem string

	variables map[string]interface{}
	metadata  *libmonteur.TOMLMetadata
	releases  *libmonteur.TOMLRelease
	cmd       []*libmonteur.TOMLAction
}

func (me *releaser) Parse(path string) (err error) {
	// init temporary raw input variables
	dependencies := []*commander.Dependency{}
	dep := []*libmonteur.TOMLDependency{}
	fmtVar := map[string]interface{}{}
	cmd := []*libmonteur.TOMLAction{}

	// init all important variables
	me.metadata = &libmonteur.TOMLMetadata{}
	me.cmd = []*libmonteur.TOMLAction{}
	me.releases = &libmonteur.TOMLRelease{
		Data:     &libmonteur.TOMLReleaseData{},
		Packages: map[string]*libmonteur.TOMLPackage{},
	}

	// construct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
		Dependencies *[]*libmonteur.TOMLDependency
		Releases     *libmonteur.TOMLRelease
		CMD          *[]*libmonteur.TOMLAction
	}{
		Metadata:     me.metadata,
		Variables:    me.variables,
		FMTVariables: &fmtVar,
		Dependencies: &dep,
		Releases:     me.releases,
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

	err = sanitizeDeps(dep, &dependencies, me.thisSystem, me.variables)
	if err != nil {
		return err
	}

	err = checkDependencies(&dependencies)
	if err != nil {
		return err
	}

	err = sanitizeCMD(cmd, &me.cmd, me.thisSystem)
	if err != nil {
		return err
	}

	err = sanitizeRelease(me.releases, me.variables)
	if err != nil {
		return err
	}

	// init
	err = initializeLogger(&me.log, me.metadata.Name, me.variables)
	if err != nil {
		return err
	}

	return err
}

// Run executes the full run-job.
func (me *releaser) Run(ctx context.Context, ch chan conductor.Message) {
	var init, releasePkg, conclude func()
	var variables map[string]interface{}
	var pkg *libmonteur.TOMLPackage
	var manager interface{}
	var err error

	me.log.Info(libmonteur.LOG_JOB_START + "\n\n")
	me.reportUp = ch

	switch me.metadata.Type {
	case libmonteur.RELEASE_ARCHIVE:
		init = func() {
			manager, err = libarchiver.Init(me.log,
				me.releases,
				me.variables,
			)
		}

		releasePkg = func() {
			err = libarchiver.Release(pkg, manager, variables)
		}

		conclude = func() {
			err = libarchiver.Conclude(manager)
		}
	case libmonteur.RELEASE_MANUAL:
		init = nil

		releasePkg = func() {
			err = me.runManually(variables)
		}

		conclude = nil
	default:
		me.reportError(fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_RELEASER_TYPE_UNSUPPORTED,
			me.metadata.Type,
		))
		return
	}

	me.runFx(init, "Executing release initialization function now...")
	if err != nil {
		me.reportError(err)
		return
	}

	for _, pkg = range me.releases.Packages {
		// copy the original variables into the new variable list
		variables = map[string]interface{}{}
		for k, v := range me.variables {
			variables[k] = v
		}

		// process package variables
		err = processPackageVariables(pkg, &variables)
		if err != nil {
			me.reportError(err)
			return
		}

		// process source
		pkg.Source, err = libmonteur.ProcessString(pkg.Source,
			variables,
		)
		if err != nil {
			me.reportError(err)
			return
		}
		variables[libmonteur.VAR_SOURCE] = pkg.Source

		// process target
		pkg.Target, err = me.processPkgTarget(pkg.Target,
			me.releases.Target,
			variables,
		)
		if err != nil {
			me.reportError(err)
			return
		}
		variables[libmonteur.VAR_TARGET] = pkg.Target

		me.runFx(releasePkg, "Executing release package function now...")
		if err != nil {
			me.reportError(err)
			return
		}
	}

	me.runFx(conclude, "Executing release conclusion function now...")
	if err != nil {
		me.reportError(err)
		return
	}

	me.reportDone()
}

func (me *releaser) processPkgTarget(in string,
	def string, variables map[string]interface{}) (out string, err error) {
	out, err = libmonteur.ProcessString(in, variables)
	if err != nil {
		return "", err //nolint:wrapcheck
	}

	if out != "" {
		return out, nil
	}

	// use def as a replacement since out is empty
	return libmonteur.ProcessString(def, variables) //nolint:wrapcheck
}

func (me *releaser) runFx(fx func(), name string) {
	if fx == nil {
		return
	}

	me.log.Info(name)
	fx()
	me.log.Info(libmonteur.LOG_OK)
}

func (me *releaser) runManually(variables map[string]interface{}) (err error) {
	me.log.Info("Executing Manual Release Commands now...")
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

	me.log.Info("Executing Manual Release Commands ➤ DONE\n\n")
	return nil
}

// Name is to return the job name
func (me *releaser) Name() string {
	return me.metadata.Name
}

func (me *releaser) reportStatus(format string, args ...interface{}) {
	reportStatus(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *releaser) reportError(err error) {
	reportError(me.log, me.reportUp, me.metadata.Name, "%s", err)
}

func (me *releaser) reportOutput(format string, args ...interface{}) {
	reportOutput(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *releaser) reportDone() {
	reportDone(me.log, me.reportUp, me.metadata.Name)
}
