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

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/checksum"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libhttp"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblocal"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libtargz"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libzip"
)

type setup struct {
	thisSystem string
	config     string

	reportUp chan conductor.Message
	log      *liblog.Logger

	variables    map[string]interface{}
	metadata     *libmonteur.TOMLMetadata
	source       *libmonteur.TOMLSource
	dependencies []*commander.Dependency
	cmd          []*libmonteur.TOMLAction
}

func (me *setup) Parse(path string) (err error) {
	// init temporary raw input variables
	dep := []*libmonteur.TOMLDependency{}
	fmtVar := map[string]interface{}{}
	cmd := []*libmonteur.TOMLAction{}
	cfg := map[string]string{}
	sources := map[string]*libmonteur.TOMLSource{}

	// initialize all important variables
	me.metadata = &libmonteur.TOMLMetadata{}
	me.dependencies = []*commander.Dependency{}
	me.cmd = []*libmonteur.TOMLAction{}
	me.source = &libmonteur.TOMLSource{}

	// construct TOML file data structure
	s := struct {
		Metadata     *libmonteur.TOMLMetadata
		Variables    map[string]interface{}
		FMTVariables *map[string]interface{}
		Dependencies *[]*libmonteur.TOMLDependency
		Sources      *map[string]*libmonteur.TOMLSource
		CMD          *[]*libmonteur.TOMLAction
		Config       *map[string]string
	}{
		Metadata:     me.metadata,
		Variables:    me.variables,
		FMTVariables: &fmtVar,
		Dependencies: &dep,
		Sources:      &sources,
		CMD:          &cmd,
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

	err = sanitizeSources(sources, &me.source, &me.variables, me.thisSystem)
	if err != nil {
		return err
	}

	err = sanitizeCMD(cmd, &me.cmd, me.thisSystem)
	if err != nil {
		return err
	}

	err = sanitizeSourceConfig(cfg, &me.config, me.variables)
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
func (me *setup) Run(ctx context.Context, ch chan conductor.Message) {
	var unpackFx func(*libmonteur.TOMLSource, map[string]interface{}) error
	var sourceFx func(context.Context,
		*libmonteur.TOMLSource,
		map[string]interface{}, *liblog.Logger, *checksum.Hasher) error
	var cs *checksum.Hasher
	var err error
	var task *executive

	me.log.Info("Run Task Now: " + libmonteur.LOG_SUCCESS + "\n")
	me.reportUp = ch

	unpackFx, err = me.prepareUnpackFx()
	if err != nil {
		me.reportError(err)
		return
	}

	sourceFx, err = me.prepareSourceFx()
	if err != nil {
		me.reportError(err)
	}

	cs, err = me.prepareChecksumFx()
	if err != nil {
		me.reportError(err)
	}

	err = sourceFx(ctx, me.source, me.variables, me.log, cs)
	if err != nil {
		me.reportError(err)
		return
	}

	me.log.Info("executing unpack function now...")
	if unpackFx != nil {
		err = unpackFx(me.source, me.variables)
		if err != nil {
			me.reportError(err)
			return
		}
	}
	me.log.Info("Executing unpack function now ➤ DONE\n\n")

	me.log.Info("executing cmd now...")
	task = &executive{
		log:       me.log,
		variables: me.variables,
		orders:    me.cmd,
		fxSTDOUT:  me.reportOutput,
		fxSTDERR:  me.reportStatus,
	}

	err = task.Exec()
	if err != nil {
		me.reportError(err)
		return
	}
	me.log.Info("Executing CMD now ➤ DONE\n\n")

	me.log.Info("Executing config scripting now...")
	err = me.processConfig()
	if err != nil {
		me.reportError(err)
		return
	}
	me.log.Info("Executing config scripting now ➤ DONE\n\n")

	me.reportDone()
}

func (me *setup) processConfig() (err error) {
	var ok bool
	var configPath, pathing string

	// get config path
	configPath, ok = me.variables[libmonteur.VAR_CFG].(string)
	if !ok {
		panic("MONTEUR_DEV: why is VAR_CFG missing?")
	}

	// process pathing
	pathing = strings.ToLower(me.metadata.Name)
	pathing = strings.ReplaceAll(pathing, " ", "-")
	pathing = strings.ReplaceAll(pathing, "_", "-")
	pathing = strings.ReplaceAll(pathing, "%", "-")
	pathing = strings.ReplaceAll(pathing, "!", "-")
	pathing = filepath.Join(configPath,
		libmonteur.DIRECTORY_MONTEUR_CONFIG_D,
		pathing,
	)

	// write into config directory
	me.log.Info("Post-configuring into '%s'", pathing)
	err = os.WriteFile(pathing,
		[]byte(me.config),
		libmonteur.PERMISSION_CONFIG,
	)
	if err != nil {
		err = fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_FAILED,
			err,
		)
	}

	return err
}

func (me *setup) prepareSourceFx() (out func(context.Context,
	*libmonteur.TOMLSource,
	map[string]interface{}, *liblog.Logger, *checksum.Hasher) error,
	err error) {
	switch strings.ToLower(me.metadata.Type) {
	case libmonteur.PROGRAM_TYPE_HTTPS_DOWNLOAD:
		out = libhttp.Source
	case libmonteur.PROGRAM_TYPE_LOCAL_SYSTEM:
		out = liblocal.Source
	default:
		err = fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_TYPE_UNKNOWN,
			me.metadata.Type,
		)
	}

	return out, err
}

func (me *setup) prepareChecksumFx() (out *checksum.Hasher, err error) {
	if me.source.Checksum == nil {
		return nil, nil
	}

	out = &checksum.Hasher{}
	switch strings.ToLower(me.source.Checksum.Format) {
	case libmonteur.CHECKSUM_FORMAT_BASE64:
		err = out.ParseBase64(me.source.Checksum.Value)
	case libmonteur.CHECKSUM_FORMAT_BASE64_URL:
		err = out.ParseBase64URL(me.source.Checksum.Value)
	case libmonteur.CHECKSUM_FORMAT_HEX:
		err = out.ParseHex(me.source.Checksum.Value)
	default:
		return nil, fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_CHECKSUM_FORMAT_UNKNOWN,
			me.source.Checksum.Format,
		)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %s",
			libmonteur.ERROR_CHECKSUM_BAD,
			err,
		)
	}

	switch strings.ToLower(me.source.Checksum.Type) {
	case libmonteur.CHECKSUM_ALGO_SHA512:
		_ = out.SetAlgo(checksum.HASHER_SHA512)
	case libmonteur.CHECKSUM_ALGO_SHA256:
		_ = out.SetAlgo(checksum.HASHER_SHA256)
	case libmonteur.CHECKSUM_ALGO_MD5:
		_ = out.SetAlgo(checksum.HASHER_MD5)
	default:
		return nil, fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_CHECKSUM_ALGO_UNKNOWN,
			me.source.Checksum.Type,
		)
	}

	return out, nil
}

func (me *setup) prepareUnpackFx() (out func(*libmonteur.TOMLSource,
	map[string]interface{}) error, err error) {
	switch me.source.Format {
	case libmonteur.PROGRAM_FORMAT_TAR_GZ:
		out = libtargz.Unpack
	case libmonteur.PROGRAM_FORMAT_ZIP:
		out = libzip.Unpack
	case libmonteur.PROGRAM_FORMAT_RAW:
		out = nil
	default:
		err = fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_FORMAT_UNKNOWN,
			me.source.Format,
		)
	}

	return out, err
}

// Name is to return the task name
func (me *setup) Name() string {
	return me.metadata.Name
}

func (me *setup) reportStatus(format string, args ...interface{}) {
	reportStatus(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *setup) reportError(err error) {
	reportError(me.log, me.reportUp, me.metadata.Name, "%s", err)
}

func (me *setup) reportOutput(format string, args ...interface{}) {
	reportOutput(me.log, me.reportUp, me.metadata.Name, format, args...)
}

func (me *setup) reportDone() {
	reportDone(me.log, me.reportUp, me.metadata.Name)
}
