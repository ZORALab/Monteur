// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by melicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libsetup

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/checksum"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/httpclient"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/targz"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

// Manager is the supported setup program to source
type Manager struct {
	Variables map[string]interface{}

	Metadata *libmonteur.TOMLMetadata
	log      *liblog.Logger
	source   *libmonteur.TOMLSource
	checksum *checksum.Hasher

	getFx    func(context.Context)
	unpackFx func(context.Context)

	dependencies []*commander.Dependency
	setup        []*commander.Action

	reportUp chan conductor.Message

	config        string
	workspacePath string
	installPath   string
	configPath    string

	thisSystem string
}

// Parse is to parse a Setup API's Program TOML file.
//
// This API should be called before the Conductor's orchestrations as it does
// not have any logging features.
//
// All errors generated from this method shall be returned
// **WITHOUT USING `me.ReportError(...)`** method.
func (me *Manager) Parse(path string) (err error) {
	var ok bool

	// initialize all important variables
	me.Metadata = &libmonteur.TOMLMetadata{}
	me.dependencies = []*commander.Dependency{}
	me.setup = []*commander.Action{}

	me.thisSystem, ok = me.Variables[libmonteur.VAR_COMPUTE].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_COMPUTE before Parse()!")
	}

	me.workspacePath, ok = me.Variables[libmonteur.VAR_TMP].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_TMP before Parse()!")
	}

	me.installPath, ok = me.Variables[libmonteur.VAR_BIN].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_BIN before Parse()!")
	}

	me.configPath, ok = me.Variables[libmonteur.VAR_CFG].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_CFG before Parse()!")
	}

	fmtVar := map[string]interface{}{}
	dep := []*libmonteur.TOMLDependency{}
	cmd := []*libmonteur.TOMLAction{}
	cfg := map[string]string{}
	sources := map[string]*libmonteur.TOMLSource{}

	// consturct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
		Sources      *map[string]*libmonteur.TOMLSource
		Dependencies *[]*libmonteur.TOMLDependency
		Setup        *[]*libmonteur.TOMLAction
		Config       *map[string]string
	}{
		Metadata:     me.Metadata,
		Variables:    me.Variables,
		FMTVariables: &fmtVar,
		Sources:      &sources,
		Dependencies: &dep,
		Setup:        &cmd,
		Config:       &cfg,
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
	err = me.sanitizeMetadata(path)
	if err != nil {
		return err
	}

	err = me.sanitizeFMTVariables(fmtVar)
	if err != nil {
		return err
	}

	err = me.sanitizeDeps(dep)
	if err != nil {
		return err
	}

	err = me.sanitizeSource(sources)
	if err != nil {
		return err
	}

	err = me.sanitizeSetup(cmd)
	if err != nil {
		return err
	}

	err = me.sanitizeConfig(cfg)
	if err != nil {
		return err
	}

	err = me.initializeLogger()
	if err != nil {
		return err
	}

	return nil
}

func (me *Manager) initializeLogger() (err error) {
	var sRet, name string
	var ok bool

	// initialize variables
	sRet, ok = me.Variables[libmonteur.VAR_LOG].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_LOG before Parse()!")
	}
	me.log = &liblog.Logger{}
	me.log.Init()

	name = strings.ToLower(me.Metadata.Name)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ReplaceAll(name, "!", "")
	name = strings.ReplaceAll(name, "$", "")

	// create status log
	err = me.log.Add(liblog.TYPE_STATUS, filepath.Join(
		sRet,
		name+"-"+libmonteur.FILE_LOG_STATUS,
	))
	if err != nil {
		return err //nolint:wrapcheck
	}

	// create output log
	err = me.log.Add(liblog.TYPE_OUTPUT, filepath.Join(
		sRet,
		name+"-"+libmonteur.FILE_LOG_OUTPUT,
	))
	if err != nil {
		me.log.Close()
		return err //nolint:wrapcheck
	}

	me.log.Info(libmonteur.LOG_JOB_INIT_SUCCESS)

	return nil
}

func (me *Manager) sanitizeConfig(in map[string]string) (err error) {
	os, ok := me.Variables[libmonteur.VAR_OS].(string)
	if !ok {
		panic("Monteur DEV: please assign VAR_OS before Parse()!")
	}

	me.config, ok = in[os]
	if !ok {
		me.config, _ = in[libmonteur.ALL_OS]
	}

	if me.config == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_BAD,
			"missing config for "+os,
		)
	}

	me.config, err = templater.String(me.config, me.Variables)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_BAD,
			err,
		)
	}

	return nil
}

func (me *Manager) sanitizeSetup(in []*libmonteur.TOMLAction) (err error) {
	for _, cmd := range in {
		if !libmonteur.IsComputeSystemSupported(me.thisSystem,
			cmd.Condition) {
			continue
		}

		s := &commander.Action{
			Name:     cmd.Name,
			Type:     cmd.Type,
			Location: cmd.Location,
			Source:   cmd.Source,
			Target:   cmd.Target,
			Save:     cmd.Save,
			SaveFx:   me._saveFx,
		}

		me.setup = append(me.setup, s)
	}

	// sanitize each of them
	for i, cmd := range me.setup {
		err = cmd.Init()
		if err != nil {
			return fmt.Errorf("%s (Setup %d): %s",
				libmonteur.ERROR_COMMAND_BAD,
				i+1,
				err,
			)
		}
	}

	return nil
}

func (me *Manager) sanitizeSource(in map[string]*libmonteur.TOMLSource) (err error) {
	// get 'all-all' omni-compute platform and merge accordingly
	actual := in[libmonteur.COMPUTE_SYSTEM_OMNI]

	// get platform specific source
	specific := in[me.thisSystem]

	// merge the 2 sources' data
	switch {
	case actual == nil && specific == nil:
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_UNSUPPORTED,
			me.thisSystem,
		)
	case actual == nil:
		actual = specific
	}

	// both base and actual are valid. Perform the merging...
	if specific != nil {
		actual.Merge(specific)
	}
	me.source = actual

	// sanitize each data fields
	err = me._sanitizesourceData()
	if err != nil {
		me.source = nil
		return err
	}

	return nil
}

func (me *Manager) _sanitizesourceData() (err error) {
	err = me.__sanitizesourceFormat()
	if err != nil {
		return err
	}
	me.Variables[libmonteur.VAR_FORMAT] = me.source.Format
	defer delete(me.Variables, libmonteur.VAR_FORMAT)

	err = me.__sanitizesourceArchive()
	if err != nil {
		return err
	}
	me.Variables[libmonteur.VAR_ARCHIVE] = me.source.Archive
	defer delete(me.Variables, libmonteur.VAR_ARCHIVE)

	err = me.__sanitizesourceMethod()
	if err != nil {
		return err
	}
	me.Variables[libmonteur.VAR_METHOD] = me.source.Method
	defer delete(me.Variables, libmonteur.VAR_METHOD)

	err = me.__sanitizesourceURL()
	if err != nil {
		return err
	}
	me.Variables[libmonteur.VAR_URL] = me.source.URL
	defer delete(me.Variables, libmonteur.VAR_URL)

	err = me.__sanitizesourceHeaders()
	if err != nil {
		return err
	}

	err = me.__sanitizesourceChecksum()
	if err != nil {
		return err
	}

	return nil
}

func (me *Manager) __sanitizesourceChecksum() (err error) {
	var h *checksum.Hasher

	if me.source.Checksum == nil {
		return nil
	}

	if me.source.Checksum.Value == "" {
		return fmt.Errorf("%s: ''", libmonteur.ERROR_CHECKSUM_BAD)
	}

	h = &checksum.Hasher{}
	switch strings.ToLower(me.source.Checksum.Format) {
	case libmonteur.CHECKSUM_FORMAT_BASE64:
		err = h.ParseBase64(me.source.Checksum.Value)
	case libmonteur.CHECKSUM_FORMAT_BASE64_URL:
		err = h.ParseBase64URL(me.source.Checksum.Value)
	case libmonteur.CHECKSUM_FORMAT_HEX:
		err = h.ParseHex(me.source.Checksum.Value)
	default:
		return fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_CHECKSUM_FORMAT_UNKNOWN,
			me.source.Checksum.Format,
		)
	}

	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_CHECKSUM_BAD,
			err,
		)
	}

	switch strings.ToLower(me.source.Checksum.Type) {
	case libmonteur.CHECKSUM_ALGO_SHA512:
		_ = h.SetAlgo(checksum.HASHER_SHA512)
	case libmonteur.CHECKSUM_ALGO_SHA256:
		_ = h.SetAlgo(checksum.HASHER_SHA256)
	case libmonteur.CHECKSUM_ALGO_MD5:
		_ = h.SetAlgo(checksum.HASHER_MD5)
	default:
		return fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_CHECKSUM_ALGO_UNKNOWN,
			me.source.Checksum.Type,
		)
	}

	me.checksum = h

	return nil
}

func (me *Manager) __sanitizesourceHeaders() (err error) {
	headers := map[string]string{}

	for k, v := range me.source.Headers {
		v, err = templater.String(v, me.Variables)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_PROGRAM_HTTPS_HEADER_BAD,
				err,
			)
		}

		headers[k] = v
	}

	me.source.Headers = headers

	return nil
}

func (me *Manager) __sanitizesourceURL() (err error) {
	me.source.URL, err = templater.String(me.source.URL, me.Variables)
	if err != nil || me.source.Method == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			me.source.URL,
		)
	}

	return nil
}

func (me *Manager) __sanitizesourceMethod() (err error) {
	me.source.Method, err = templater.String(me.source.Method, me.Variables)
	if err != nil || me.source.Method == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			me.source.Method,
		)
	}

	return nil
}

func (me *Manager) __sanitizesourceArchive() (err error) {
	me.source.Archive, err = templater.String(me.source.Archive,
		me.Variables)
	if err != nil || me.source.Archive == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			me.source.Archive,
		)
	}

	return nil
}

func (me *Manager) __sanitizesourceFormat() (err error) {
	me.source.Format, err = templater.String(me.source.Format, me.Variables)
	me.source.Format = strings.ToLower(me.source.Format)

	switch {
	case err != nil:
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_FORMAT_BAD,
			me.source.Format,
		)
	case me.source.Format == libmonteur.PROGRAM_FORMAT_TAR_GZ:
		me.unpackFx = me.UnarchiveTarGz
	case me.source.Format == libmonteur.PROGRAM_FORMAT_ZIP:
		me.unpackFx = me.UnarchiveZip
	case me.source.Format == libmonteur.PROGRAM_FORMAT_RAW:
		me.unpackFx = me.UnarchiveRaw
	default:
	}

	if me.unpackFx == nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_FORMAT_UNKNOWN,
			me.source.Format,
		)
	}

	return nil
}

func (me *Manager) sanitizeDeps(in []*libmonteur.TOMLDependency) (err error) {
	var val string

	for _, dep := range in {
		if !libmonteur.IsComputeSystemSupported(me.thisSystem,
			[]string{dep.Condition}) {
			continue
		}

		val, err = templater.String(dep.Command, me.Variables)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_COMMAND_DEPENDENCY_FMT_BAD,
				err,
			)
		}

		s := &commander.Dependency{
			Name:    dep.Name,
			Type:    dep.Type,
			Command: val,
		}

		me.dependencies = append(me.dependencies, s)
	}

	// sanitize each of them
	for _, dep := range me.dependencies {
		err = dep.Init()
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_DEPENDENCY_BAD,
				err,
			)
		}
	}

	return nil
}

func (me *Manager) sanitizeFMTVariables(in map[string]interface{}) (err error) {
	var val interface{}

	if in == nil {
		return nil
	}

	for key, value := range in {
		switch v := value.(type) {
		case string:
			val, err = templater.String(v, me.Variables)
		default:
			val = v
		}

		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_VARIABLES_FMT_BAD,
				err,
			)
		}

		me.Variables[key] = val
	}

	return nil
}

func (me *Manager) sanitizeMetadata(path string) (err error) {
	err = me.Metadata.Sanitize(path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	switch strings.ToLower(me.Metadata.Type) {
	case libmonteur.PROGRAM_TYPE_HTTPS_DOWNLOAD:
		me.getFx = me.sourceHTTPS
	case libmonteur.PROGRAM_TYPE_LOCAL_SYSTEM:
		me.getFx = me.sourceLocal
	}

	if me.getFx == nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_TYPE_UNKNOWN,
			me.Metadata.Type,
		)
	}

	return nil
}

// Name is for generating the program Metadata.Name when used as interface
//
// This should only be called after the Manager is initialized successfully.
func (me *Manager) Name() string {
	if me.Metadata == nil {
		return ""
	}

	return me.Metadata.Name
}

// Run is to get the Program from its source.
//
// Everything must be setup properly before calling this function. It was meant
// for Monteur's Setup API.
//
// All errors generated in this method shall use `me.reportError` instead of
// returning `fmt.Errorf` since it will be executed in parallel with others
// in an asynchonous manner.
//
// This should only be called after the Manager is initialized successfully.
func (me *Manager) Run(ctx context.Context, ch chan conductor.Message) {
	me.reportUp = ch
	me.log.Success(libmonteur.LOG_SUCCESS)

	me.getFx(ctx)
	me.unpackFx(ctx)
	me.Install(ctx)
	me.PostConfigure(ctx)
	me.reportDone()
}

func (me *Manager) sourceHTTPS(ctx context.Context) {
	me.log.Info("Sourcing %s using HTTPS download...", me.Metadata.Name)

	d := &httpclient.Downloader{
		Destination: filepath.Join(me.workspacePath,
			me.source.Archive),
		Headers: me.source.Headers,
	}

	d.HandleError = func(err error) {
		me.reportError("%s", err)
	}

	d.HandleSuccess = func() {
		me.log.Info("%s ➤ download completed", me.Metadata.Name)
	}

	d.HandleProgress = func(downloaded, total int64) {
		percent := float64(downloaded) / float64(total) * 100

		me.log.Info("%d / %d Bytes (%.0f%%)",
			downloaded,
			total,
			percent,
		)
	}

	me.log.Info("Downloader Destination: %v", d.Destination)
	me.log.Info("Downloader HandleSuccess: %v", d.HandleSuccess)
	me.log.Info("Downloader HandleProgress: %v", d.HandleProgress)
	me.log.Info("Downloader Method: %v", me.source.Method)
	me.log.Info("Downloader URL: %v", me.source.URL)
	me.log.Info("Downloader Checksum: %v", me.checksum)

	if len(d.Headers) == 0 {
		me.log.Info("Downloader Headers: {}")
	} else {
		for k, v := range d.Headers {
			me.log.Info("  '%s': '%s'", k, v)
		}
	}

	me.log.Info("Begin downloading...")
	d.Download(ctx, me.source.Method, me.source.URL, me.checksum)
}

func (me *Manager) sourceLocal(ctx context.Context) {
	me.log.Info("looking for local executable: %v", me.source.URL)

	_, err := exec.LookPath(me.source.URL)
	if err != nil {
		me.reportError("%s: %s",
			libmonteur.ERROR_PROGRAM_MISSING,
			me.source.URL,
		)

		return
	}

	me.log.Info("%s ➤ search completed", me.Metadata.Name)
}

func (me *Manager) Install(ctx context.Context) {
	var err error

	for i, cmd := range me.setup {
		me.log.Info("Executing Setup Commands...")
		me.log.Info("Name: '%s'", cmd.Name)
		me.log.Info("Save: '%s'", cmd.Save)
		me.log.Info("SaveFx: '%v'", cmd.SaveFx)
		me.log.Info("Type: '%v'", cmd.Type)

		me.log.Info("Formatting cmd.Location...")
		cmd.Location, err = templater.String(cmd.Location, me.Variables)
		if err != nil {
			me.reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)

			return
		}
		me.log.Info("Got: '%s'", cmd.Location)

		me.log.Info("Formatting cmd.Source...")
		cmd.Source, err = templater.String(cmd.Source, me.Variables)
		if err != nil {
			me.reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)

			return
		}
		me.log.Info("Got: '%s'", cmd.Source)

		me.log.Info("Formatting cmd.Target...")
		cmd.Target, err = templater.String(cmd.Target, me.Variables)
		if err != nil {
			me.reportError("%s: %s",
				libmonteur.ERROR_COMMAND_FMT_BAD,
				err,
			)

			return
		}
		me.log.Info("Got: '%s'", cmd.Target)

		me.log.Info("Running cmd...")
		if cmd.Save == "" {
			cmd.Save = libmonteur.COMMAND_SAVE_NONE
		}

		err = cmd.Run()
		if err != nil {
			me.reportError("%s: (Step %d) %s",
				libmonteur.ERROR_COMMAND_FAILED,
				i+1,
				err,
			)

			return
		}
	}
}

func (me *Manager) _saveFx(key string, output interface{}) {
	switch v := output.(type) {
	case *commander.ExecOutput:
		me.log.Info("Reading STDERR...")
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stderr))
		me.log.Info("Reading STDOUT...")
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stdout))

		if key != libmonteur.COMMAND_SAVE_NONE {
			val := strings.TrimRight(string(v.Stdout), "\r\n")
			me.Variables[key] = val
			me.log.Info("Saving '%v' to '%s'...", output, key)
		}
	case commander.ExecOutput:
		me.log.Info("Reading STDERR...")
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stderr))
		me.log.Info("Reading STDOUT...")
		me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, string(v.Stdout))

		if key != libmonteur.COMMAND_SAVE_NONE {
			val := strings.TrimRight(string(v.Stdout), "\r\n")
			me.Variables[key] = val
			me.log.Info("Saving '%v' to '%s'...", output, key)
		}
	default:
		me.log.Info("Reading output...")
		if v == nil {
			me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, "nil\n")
		} else {
			me.log.Info(libmonteur.LOG_FORMAT_OUTPUT_LONG, output)
		}

		if key != libmonteur.COMMAND_SAVE_NONE {
			me.Variables[key] = output
			me.log.Info("Saving '%v' to '%s'...", output, key)
		}
	}

	me.log.Info(libmonteur.LOG_SUCCESS)
}

func (me *Manager) PostConfigure(ctx context.Context) {
	var pathing string
	var err error

	// process pathing
	pathing = strings.ToLower(me.Metadata.Name)
	pathing = strings.ReplaceAll(pathing, " ", "-")
	pathing = strings.ReplaceAll(pathing, "_", "-")
	pathing = strings.ReplaceAll(pathing, "%", "-")
	pathing = strings.ReplaceAll(pathing, "!", "-")
	pathing = filepath.Join(me.configPath, pathing)

	// write into config directory
	me.log.Info("Post-configuring '%s'...", pathing)
	me.log.Info("File:\n%s", me.config)
	err = os.WriteFile(pathing,
		[]byte(me.config),
		libmonteur.PERMISSION_CONFIG,
	)
	if err != nil {
		me.reportError("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_FAILED,
			err,
		)
	}

	me.log.Info(libmonteur.LOG_SUCCESS)
}

func (me *Manager) UnarchiveTarGz(ctx context.Context) {
	var pathing string
	var err error

	pathing = filepath.Join(me.workspacePath, me.source.Archive)

	me.log.Info("Unarchive .tar.gz %s...", pathing)
	processor := &targz.Processor{
		Archive: pathing,
		Raw:     me.workspacePath,
	}

	err = processor.Extract()
	if err != nil {
		me.reportError("%s", err)
		return
	}

	me.log.Info(libmonteur.LOG_SUCCESS)
}

func (me *Manager) UnarchiveZip(ctx context.Context) {
	//TODO: zip archives
}

func (me *Manager) UnarchiveRaw(ctx context.Context) {
	me.log.Info("Unarchive raw %s...",
		filepath.Join(me.workspacePath, me.source.Archive),
	)
	me.log.Info(libmonteur.LOG_SUCCESS)
}

func (me *Manager) reportError(format string, args ...interface{}) {
	if me.Metadata == nil || me.Metadata.Name == "" {
		format = "Task '' ➤ " + format
	} else {
		format = "Task '%s' ➤ " + format
		args = append([]interface{}{me.Metadata.Name}, args...)
	}

	if me.log != nil {
		me.log.Error(format, args...)
		me.log.Sync()
		me.log.Close()
	}

	if me.reportUp != nil {
		me.reportUp <- conductor.CreateError(me.Metadata.Name,
			format,
			args...,
		)
	}
}

func (me *Manager) reportStatus(format string, args ...interface{}) {
	if me.Metadata == nil || me.Metadata.Name == "" {
		format = "Task '' ➤ " + format
	} else {
		format = "Task '%s' ➤ " + format
		args = append([]interface{}{me.Metadata.Name}, args...)
	}

	if me.log != nil {
		me.log.Info(format, args...)
	}

	if me.reportUp != nil {
		me.reportUp <- conductor.CreateStatus(me.Metadata.Name,
			format,
			args...,
		)
	}
}

func (me *Manager) reportDone() {
	if me.log != nil {
		me.log.Success(libmonteur.LOG_SUCCESS)
		me.log.Sync()
		me.log.Close()
	}

	if me.reportUp != nil {
		me.reportUp <- conductor.CreateDone(me.Metadata.Name)
	}
}
