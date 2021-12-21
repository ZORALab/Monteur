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

package libworkspace

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsecrets"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/styler"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

// Workspace is the Monteur continuous integration main data sructure.
//
// This data structure is responsible for running through all Monteur operations
type Workspace struct {
	Timestamp  *time.Time
	Filesystem *Pathing
	Language   *libmonteur.Language
	App        *libmonteur.Software
	Variables  *map[string]interface{}
	secrets    *map[string]interface{}

	Job           string
	Version       string
	OS            string
	ARCH          string
	ComputeSystem string
	ConfigDir     string
	JobTOMLFile   string
}

// Init is to initialize the workspace for usage
func (me *Workspace) Init() error {
	x := time.Now().UTC()
	me.Timestamp = &x

	if err := me.parseWorkspaceData(); err != nil {
		return err
	}

	if err := me.parseAppData(); err != nil {
		return err
	}

	me.processSecrets()
	me.processDataByJob()

	return nil
}

func (me *Workspace) parseWorkspaceData() (err error) {
	me.Language = &libmonteur.Language{}
	me.Variables = &map[string]interface{}{}
	me.OS = runtime.GOOS
	me.ARCH = runtime.GOARCH
	me.Version = libmonteur.VERSION
	me.ComputeSystem = me.OS + libmonteur.COMPUTE_SYSTEM_SEPARATOR + me.ARCH

	// initialize pathing for parsing
	me.Filesystem = &Pathing{
		timestampDir: me.Timestamp.Format("2006-Jan-02T15-04-05UTC"),
	}

	err = me.Filesystem.Init()
	if err != nil {
		return err
	}

	// initialize processing data
	fmtVar := map[string]interface{}{}

	// parse workspace TOML data
	s := struct {
		Language     *libmonteur.Language
		Filesystem   *Pathing
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
	}{
		Language:     me.Language,
		Filesystem:   me.Filesystem,
		Variables:    *me.Variables,
		FMTVariables: &fmtVar,
	}

	err = toml.DecodeFile(me.Filesystem.WorkspaceTOMLFile, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	err = libmonteur.SanitizeVariables(me.Variables, &fmtVar)
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = me._sanitizeLanguage()
	if err != nil {
		return err
	}

	// update filesystem pathing
	err = me.Filesystem.Update(me.Language.Code)
	if err != nil {
		return err
	}

	// init app
	me.App = &libmonteur.Software{
		Time: me.__createTimestamp(),
	}
	(*me.Variables)[libmonteur.VAR_APP] = me.App

	return nil
}

func (me *Workspace) _sanitizeLanguage() (err error) {
	if me.Language.Code == "" {
		return fmt.Errorf(libmonteur.ERROR_LANGUAGE_CODE_MISSING)
	}

	if me.Language.Name == "" {
		return fmt.Errorf(libmonteur.ERROR_LANGUAGE_NAME_MISSING)
	}

	return nil
}

func (me *Workspace) parseAppData() (err error) {
	err = me._parseAppSpec()
	if err != nil {
		return err
	}

	err = me._parseAppDebian()
	if err != nil {
		return err
	}

	err = me._parseAppCopyrights()
	if err != nil {
		return err
	}

	err = me._parseAppHelp()
	if err != nil {
		return err
	}

	return nil
}

func (me *Workspace) _parseAppCopyrights() (err error) {
	//nolint:wrapcheck
	return filepath.Walk(me.Filesystem.AppCopyrightsDir,
		me.__filterAppCopyright,
	)
}

func (me *Workspace) __filterAppCopyright(path string,
	info os.FileInfo, err error) error {
	var ok bool

	ok, err = libmonteur.AcceptTOML(path, info, err)
	if !ok {
		return err //nolint:wrapcheck
	}

	// create data type for parsing
	c := &libmonteur.Copyright{}

	// parse workspace TOML data
	s := struct {
		Copyright *libmonteur.Copyright
	}{
		Copyright: c,
	}

	err = toml.DecodeFile(path, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// save into copyright list
	me.App.Copyrights = append(me.App.Copyrights, c)

	return nil
}

func (me *Workspace) _parseAppDebian() (err error) {
	me.App.Debian = &libmonteur.DEB{}

	// parse workspace TOML data
	s := struct {
		DEB *libmonteur.DEB
	}{
		DEB: me.App.Debian,
	}

	err = toml.DecodeFile(me.Filesystem.AppDebianTOMLFile, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	return nil
}

func (me *Workspace) _parseAppHelp() (err error) {
	var ret string

	me.App.Help = &libmonteur.SoftwareHelp{
		Manpage: map[string]string{},
	}

	// parse workspace TOML data
	s := struct {
		Help *libmonteur.SoftwareHelp
	}{
		Help: me.App.Help,
	}

	err = toml.DecodeFile(me.Filesystem.AppHelpTOMLFile, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// sanitize all help fields
	err = me.__formatAppString(&me.App.Help.Command)
	if err != nil {
		return err
	}

	err = me.__formatAppString(&me.App.Help.Description)
	if err != nil {
		return err
	}

	err = me.__formatAppString(&me.App.Help.Resources)
	if err != nil {
		return err
	}

	if me.App.Help.Manpage == nil {
		return nil
	}

	for k, v := range me.App.Help.Manpage {
		ret = v
		err = me.__formatAppString(&ret)
		if err != nil {
			return err
		}

		me.App.Help.Manpage[k] = ret
	}

	return nil
}

func (me *Workspace) __formatAppString(s *string) (err error) {
	*s, err = templater.String(*s, *me.Variables)
	if err != nil {
		return fmt.Errorf("%s: %s", libmonteur.ERROR_APP_FMT_BAD, err)
	}

	return nil
}

func (me *Workspace) _parseAppSpec() (err error) {
	// parse workspace TOML data
	s := struct {
		Software *libmonteur.Software
	}{
		Software: me.App,
	}

	err = toml.DecodeFile(me.Filesystem.AppMetaTOMLFile, &s, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	// restore all default data to prevent user overwriting
	me.App.Time = me.__createTimestamp()

	return nil
}

func (me *Workspace) __createTimestamp() *libmonteur.Timestamp {
	return &libmonteur.Timestamp{
		Year:   strconv.Itoa(me.Timestamp.Year()),
		Month:  strconv.Itoa(int(me.Timestamp.Month())),
		Day:    strconv.Itoa(me.Timestamp.Day()),
		Hour:   strconv.Itoa(me.Timestamp.Hour()),
		Minute: strconv.Itoa(me.Timestamp.Minute()),
		Second: strconv.Itoa(me.Timestamp.Second()),
		Zone:   "00:00",
	}
}

func (me *Workspace) processSecrets() {
	me.secrets = &map[string]interface{}{}

	*(me).secrets = libsecrets.GetSecrets(me.Filesystem.SecretsDir)
}

func (me *Workspace) processDataByJob() {
	(*me.Variables)[libmonteur.VAR_OS] = me.OS
	(*me.Variables)[libmonteur.VAR_ARCH] = me.ARCH
	(*me.Variables)[libmonteur.VAR_COMPUTE] = me.ComputeSystem
	(*me.Variables)[libmonteur.VAR_HOME] = me.Filesystem.CurrentDir
	(*me.Variables)[libmonteur.VAR_ROOT] = me.Filesystem.RootDir
	(*me.Variables)[libmonteur.VAR_BASE] = me.Filesystem.BaseDir
	(*me.Variables)[libmonteur.VAR_CFG] = me.Filesystem.BinCfgDir
	(*me.Variables)[libmonteur.VAR_BIN] = me.Filesystem.BinDir
	(*me.Variables)[libmonteur.VAR_BUILD] = me.Filesystem.BuildTMPDir
	(*me.Variables)[libmonteur.VAR_DOC] = me.Filesystem.ComposeTMPDir
	(*me.Variables)[libmonteur.VAR_SECRETS] = *(me).secrets
	(*me.Variables)[libmonteur.VAR_TIMESTAMP] = me.Timestamp
	(*me.Variables)[libmonteur.VAR_DATA] = me.Filesystem.DataDir

	switch me.Job {
	case libmonteur.JOB_SETUP:
	case libmonteur.JOB_CLEAN:
	case libmonteur.JOB_TEST:
		me.ConfigDir = me.Filesystem.TestConfigDir
		me.JobTOMLFile = me.Filesystem.TestTOMLFile
		me.Filesystem.WorkspaceLogDir = filepath.Join(
			me.Filesystem.LogDir,
			libmonteur.DIRECTORY_TEST,
			me.Filesystem.WorkspaceLogDir,
		)

		// assign specific variables
		(*me.Variables)[libmonteur.VAR_TMP] = me.Filesystem.TestTMPDir
	case libmonteur.JOB_BUILD:
		me.ConfigDir = me.Filesystem.BuildConfigDir
		me.JobTOMLFile = me.Filesystem.BuildTOMLFile
		me.Filesystem.WorkspaceLogDir = filepath.Join(
			me.Filesystem.LogDir,
			libmonteur.DIRECTORY_BUILD,
			me.Filesystem.WorkspaceLogDir,
		)

		// assign specific variables
		(*me.Variables)[libmonteur.VAR_TMP] = me.Filesystem.BuildTMPDir
	case libmonteur.JOB_PACKAGE:
		me.ConfigDir = me.Filesystem.PackageConfigDir
		me.JobTOMLFile = me.Filesystem.PackageTOMLFile
		me.Filesystem.WorkspaceLogDir = filepath.Join(
			me.Filesystem.LogDir,
			libmonteur.DIRECTORY_PACKAGE,
			me.Filesystem.WorkspaceLogDir,
		)

		// assign specific variables
		(*me.Variables)[libmonteur.VAR_TMP] = me.Filesystem.PackageTMPDir
	case libmonteur.JOB_RELEASE:
	case libmonteur.JOB_COMPOSE:
		me.ConfigDir = me.Filesystem.ComposeConfigDir
		me.JobTOMLFile = me.Filesystem.ComposeTOMLFile
		me.Filesystem.WorkspaceLogDir = filepath.Join(
			me.Filesystem.LogDir,
			libmonteur.DIRECTORY_COMPOSE,
			me.Filesystem.WorkspaceLogDir,
		)

		// assign specific variables
		(*me.Variables)[libmonteur.VAR_TMP] = me.Filesystem.ComposeTMPDir
	case libmonteur.JOB_PUBLISH:
		me.ConfigDir = me.Filesystem.PublishConfigDir
		me.JobTOMLFile = me.Filesystem.PublishTOMLFile
		me.Filesystem.WorkspaceLogDir = filepath.Join(
			me.Filesystem.LogDir,
			libmonteur.DIRECTORY_PUBLISH,
			me.Filesystem.WorkspaceLogDir,
		)

		// assign specific variables
		(*me.Variables)[libmonteur.VAR_TMP] = me.Filesystem.PublishTMPDir
	default:
		panic("Monteur DEV: what kind of CI Job is this? ➤ " + me.Job)
	}

	// assign log directory after job specific processing
	(*me.Variables)[libmonteur.VAR_LOG] = me.Filesystem.WorkspaceLogDir
}

// String is the standard string interface for printing out Workspace data.
//
// Since Workspace is a complicated data structure, its String() method has to
// be uniquely constructed.
func (me *Workspace) String() (s string) {
	s = styler.BoxString("Monteur", styler.BORDER_DOUBLE)
	s += "One manufacturing automation app ➤ Do more with less noises!\n\n"
	s += styler.PortraitKV("Monteur Version", me.Version)
	s += me.stringCIJob() + "\n"
	s += me.stringCILocation() + "\n"
	s += me.stringApp()
	s += styler.BoxString("Execution Log", styler.BORDER_SINGLE)
	s = strings.TrimRight(s, "\n")

	return s
}

func (me *Workspace) stringCIJob() (s string) {
	s = styler.BoxString("CI Job", styler.BORDER_SINGLE)
	s += styler.PortraitKV("Job Name", me.Job)
	s += styler.PortraitKV("Language Name", me.Language.Name)
	s += styler.PortraitKV("Language Code", me.Language.Code)
	s += styler.PortraitKV("Job Timestamp", me.Timestamp.String())

	return s
}

func (me *Workspace) stringCILocation() (s string) {
	s = styler.BoxString("CI Pathing", styler.BORDER_SINGLE)
	s += styler.PortraitKV("Current Directory", me.Filesystem.CurrentDir)
	s += styler.PortraitKV("Root Directory", me.Filesystem.RootDir)
	s += styler.PortraitKV("Config Directory", me.Filesystem.ConfigDir)
	s += styler.PortraitKV("Base Directory", me.Filesystem.BaseDir)
	s += styler.PortraitKV("Working Directory", me.Filesystem.WorkingDir)
	s += styler.PortraitKV("Build Directory", me.Filesystem.BuildDir)
	s += styler.PortraitKV("Script Directory", me.Filesystem.ScriptDir)
	s += styler.PortraitKV("Bin Directory", me.Filesystem.BinDir)
	s += styler.PortraitKV("Log Directory", me.Filesystem.LogDir)
	s += styler.PortraitKV("Data Directory", me.Filesystem.DataDir)
	s += styler.PortraitKV("App Config Directory",
		me.Filesystem.AppConfigDir)

	return s
}

func (me *Workspace) stringApp() (s string) {
	s = styler.BoxString("Product Metadata", styler.BORDER_SINGLE)
	s += me.App.String()

	return s
}
