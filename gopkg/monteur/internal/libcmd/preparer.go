// Copyright 2022 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2022 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
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
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libtemplater"
)

type preparer struct {
	reportUp   chan conductor.Message
	log        *liblog.Logger
	thisSystem string

	variables map[string]interface{}
	metadata  *libmonteur.TOMLMetadata
	changelog *libmonteur.TOMLChangelog
	packages  map[string]*libmonteur.TOMLPackage
	cmd       []*libmonteur.TOMLAction
}

func (me *preparer) Parse(path string,
	secrets *libsecrets.Secrets) (err error) {
	// init temporary raw input variables
	dependencies := []*commander.Dependency{}
	dep := []*libmonteur.TOMLDependency{}
	fmtVar := map[string]interface{}{}
	cmd := []*libmonteur.TOMLAction{}

	// initialize all important variables
	me.metadata = &libmonteur.TOMLMetadata{}
	me.changelog = &libmonteur.TOMLChangelog{
		CMD: []*libmonteur.TOMLAction{},
	}
	me.cmd = []*libmonteur.TOMLAction{}
	me.packages = map[string]*libmonteur.TOMLPackage{}

	// construct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
		Dependencies *[]*libmonteur.TOMLDependency
		Changelog    *libmonteur.TOMLChangelog
		Packages     map[string]*libmonteur.TOMLPackage
		CMD          *[]*libmonteur.TOMLAction
	}{
		Metadata:     me.metadata,
		Variables:    me.variables,
		FMTVariables: &fmtVar,
		Dependencies: &dep,
		Changelog:    me.changelog,
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

	err = libtemplater.TemplateVariables(&me.variables, &fmtVar)
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

	err = sanitizeChangelog(me.changelog, me.thisSystem)
	if err != nil {
		return err
	}

	err = sanitizeCMD(cmd, &me.cmd, me.thisSystem)
	if err != nil {
		return err
	}

	// init
	err = initializeLogger(&me.log, me.metadata.Name, me.variables, secrets)
	if err != nil {
		return err
	}

	return nil
}

// Run executes the full run-job.
func (me *preparer) Run(ctx context.Context, ch chan conductor.Message) {
	var err error

	me.log.Info(libmonteur.LOG_JOB_START + "\n\n")
	me.reportUp = ch

	err = me.sourceChangelogEntries()
	if err != nil {
		me.reportError(err)
		return
	}

	err = me.runUpdates()
	if err != nil {
		me.reportError(err)
		return
	}

	me.reportDone()
}

func (me *preparer) runUpdates() (err error) {
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

		// execute changelog update
		err = me._updateChangelog(pkg, &variables)
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

func (me *preparer) sourceChangelogEntries() (err error) {
	me.log.Info("Executing latest changelog entries sourcing now...")

	task := &changelog{
		fxSTDOUT:  me.reportOutput,
		fxSTDERR:  me.reportStatus,
		variables: &me.variables,
		changelog: me.changelog,
		log:       me.log,
	}

	err = task.Exec()
	if err != nil {
		return err
	}

	me.log.Info("Executing latest changelog entries sourcing ➤ DONE\n\n")
	return nil
}

func (me *preparer) _updateChangelog(pkg *libmonteur.TOMLPackage,
	variables *map[string]interface{}) (err error) {
	me.log.Info("Executing changelog update now...")

	switch me.metadata.Type {
	case libmonteur.CHANGELOG_MARKDOWN:
	case libmonteur.CHANGELOG_MANUAL:
	case libmonteur.CHANGELOG_DEB:
		err = libdeb.Changelog(pkg, variables, me.log)
	default:
		err = fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_PREPARER_TYPE_UNKNOWN,
			me.metadata.Type,
		)
	}

	if err != nil {
		return err
	}

	me.log.Info("Executing changelog update ➤ DONE\n\n")
	return nil
}

func (me *preparer) _runCMD(variables map[string]interface{}) (err error) {
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

// Name is to return the task name
func (me *preparer) Name() string {
	return me.metadata.Name
}

func (me *preparer) reportStatus(format string, args ...interface{}) {
	reportStatus(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *preparer) reportError(err error) {
	reportError(me.log, me.reportUp, me.metadata.Name, "%s", err)
}

func (me *preparer) reportOutput(format string, args ...interface{}) {
	reportOutput(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *preparer) reportDone() {
	reportDone(me.log, me.reportUp, me.metadata.Name)
}
