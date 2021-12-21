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
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libdeb"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/oshelper"
)

type packager struct {
	reportUp   chan conductor.Message
	log        *liblog.Logger
	thisSystem string

	variables    map[string]interface{}
	metadata     *libmonteur.TOMLMetadata
	dependencies []*commander.Dependency
	changelog    *libmonteur.TOMLChangelog
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

	err = sanitizeChangelog(me.changelog, me.thisSystem)
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

	err = me.runChangelogging()
	if err != nil {
		me.reportError(err)
		return
	}

	me.runPackaging()

	me.reportDone()
}

func (me *packager) runChangelogging() (err error) {
	x := &changelog{
		fxSTDOUT:  me.reportOutput,
		fxSTDERR:  me.reportStatus,
		variables: &me.variables,
		changelog: me.changelog,
		log:       me.log,
	}

	return x.Exec()
}

func (me *packager) runPackaging() {
	var err error

	for _, pkg := range me.packages {
		// copy the original variables into the new variable list
		variables := map[string]interface{}{}
		for k, v := range me.variables {
			variables[k] = v
		}

		err = me._processPackageVariables(pkg, &variables)
		if err != nil {
			me.reportError(err)
			return
		}

		// process key data
		err = me._processPackageFields(pkg, &variables)
		if err != nil {
			me.reportError(err)
			return
		}

		// execute built-in package preprations
		err = me._preparePackage(pkg, variables)
		if err != nil {
			me.reportError(err)
			return
		}

		// execute given commands
		err = me._runPackage(variables)
		if err != nil {
			me.reportError(err)
			return
		}
	}
}

func (me *packager) _runPackage(variables map[string]interface{}) (err error) {
	me.log.Info("Executing Packaging Processing now...")
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

	me.log.Info("Executing Packaging Processing ➤ DONE\n\n")
	return nil
}

func (me *packager) _preparePackage(pkg *libmonteur.TOMLPackage,
	variables map[string]interface{}) (err error) {
	me.log.Info("Executing Packaging Preparations now...")
	// copy all files
	for k, v := range pkg.Files {
		me.log.Info("Placing merchandise...")
		me.log.Info("From: %s", v)
		me.log.Info("To  : %s", k)

		err = oshelper.CopyPath(v, k)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_PACKAGER_FILES_COPY_FAILED,
				err,
			)
		}

		me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))
	}

	// run the preparations
	switch me.metadata.Type {
	case libmonteur.PACKAGE_DEB_MANUAL:
		me.log.Info("Preparing %s packaging...",
			libmonteur.PACKAGE_DEB_MANUAL,
		)

		_, err = libdeb.Prepare(pkg, variables)
	case libmonteur.PACKAGE_MANUAL:
		me.log.Info("Preparing %s packaging...",
			libmonteur.PACKAGE_MANUAL,
		)

		err = me.__prepareManualPackaging(pkg, variables)
	default:
		err = fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_PACKAGER_TYPE_UNKNOWN,
			me.metadata.Type,
		)
	}

	if err != nil {
		return err
	}
	me.log.Info(strings.TrimSuffix(libmonteur.LOG_SUCCESS, "\n"))

	// end the preparations
	me.log.Info("Executing Packaging Preparations ➤ DONE\n\n")
	return nil
}

func (me *packager) __prepareManualPackaging(pkg *libmonteur.TOMLPackage,
	variables map[string]interface{}) (err error) {
	return nil
}

func (me *packager) _processPackageFields(pkg *libmonteur.TOMLPackage,
	variables *map[string]interface{}) (err error) {
	var key, value string

	pkg.Changelog, err = libmonteur.ProcessString(pkg.Changelog,
		*variables,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// process pkg.Files
	if len(pkg.Files) == 0 {
		return fmt.Errorf(libmonteur.ERROR_PACKAGER_FILES_MISSING)
	}

	list := map[string]string{}
	for k, v := range pkg.Files {
		key, err = libmonteur.ProcessString(k, *variables)
		if err != nil {
			return err //nolint:wrapcheck
		}

		value, err = libmonteur.ProcessString(v, *variables)
		if err != nil {
			return err //nolint:wrapcheck
		}

		list[key] = value
	}
	pkg.Files = list

	return nil
}

func (me *packager) _processPackageVariables(pkg *libmonteur.TOMLPackage,
	variables *map[string]interface{}) (err error) {
	var ok bool
	var app *libmonteur.Software
	var packagePath string

	// process operating system
	(*variables)[libmonteur.VAR_PACKAGE_OS] = pkg.OS

	// process architecture
	(*variables)[libmonteur.VAR_PACKAGE_ARCH] = pkg.Arch

	// extract app data from variable list
	app, ok = (*variables)[libmonteur.VAR_APP].(*libmonteur.Software)
	if !ok {
		return fmt.Errorf("%s", libmonteur.ERROR_PACKAGER_APP_MISSING)
	}

	// clean up VAR_PACKAGE (PackageDir) base directory
	packagePath, ok = (*variables)[libmonteur.VAR_TMP].(string)
	if !ok {
		panic("MONTEUR DEV: why is libmonteur.VAR_TMP missing?")
	}
	packagePath = filepath.Join(packagePath,
		strings.ToLower(app.Name)+"-"+pkg.OS[0]+"-"+pkg.Arch[0],
	)

	_ = os.RemoveAll(packagePath)
	_ = os.MkdirAll(packagePath, libmonteur.PERMISSION_DIRECTORY)

	// process VAR_PACKAGE (PackageDir)
	packagePath = filepath.Join(packagePath, "workspace")
	(*variables)[libmonteur.VAR_PACKAGE] = packagePath

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
