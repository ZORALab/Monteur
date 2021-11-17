// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by datalicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libsetup

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/checksum"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
)

type _programMetadata struct {
	Name        string
	Description string
	Type        string
}

type _programChecksum struct {
	Type   string
	Format string
	Value  string
}

type _programSource struct {
	Checksum *_programChecksum
	Headers  map[string]string
	Archive  string
	Format   string
	URL      string
	Method   string
}

type _programSetup struct {
	Source    string
	Target    string
	Type      string
	Condition string
}

type TOMLProgram struct {
	Metadata  *_programMetadata
	Variables map[string]interface{}
	Sources   map[string]*_programSource
	Config    map[string]string
	Setup     []*_programSetup
}

// Parse is to extract the TOML data from the given file pathing.
func (data *TOMLProgram) Parse(pathing string) (err error) {
	err = toml.DecodeFile(pathing, data, nil)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_TOML_PARSE_FAILED,
			err,
		)
	}

	return nil
}

// Process is to create the operable Program object based on its TOML data.
//
// This function shall create the working `*libsetup.Program` object file
// for later operation. This TOML program data structure can then be disposed
// off.
func (data *TOMLProgram) Process() (app *Program, err error) {
	app = &Program{
		Metadata: &Metadata{},
		Source: &Source{
			Headers: map[string]string{},
		},
		Setup:         []*Setup{},
		WorkspacePath: data.Variables[libmonteur.VAR_TMP].(string),
		InstallPath:   data.Variables[libmonteur.VAR_BIN].(string),
		ConfigPath:    data.Variables[libmonteur.VAR_CFG].(string),
	}

	err = data.processMetadata(app)
	if err != nil {
		return nil, err
	}

	err = data.processSource(app)
	if err != nil {
		return nil, err
	}

	err = data.processSetup(app)
	if err != nil {
		return nil, err
	}

	err = data.processConfig(app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (data *TOMLProgram) processMetadata(app *Program) (err error) {
	data.Metadata.Type, err = data.__processVar(data.Metadata.Type)
	if err != nil {
		return data._error(libmonteur.ERROR_PROGRAM_TYPE_BAD, err)
	}

	switch data.Metadata.Type {
	case libmonteur.PROGRAM_TYPE_HTTPS_DOWNLOAD:
		app.Source.Get = app.SourceHTTPS
	case libmonteur.PROGRAM_TYPE_LOCAL_SYSTEM:
		app.Source.Get = app.SourceLocal
	default:
		return data._error(libmonteur.ERROR_PROGRAM_TYPE_UNKNOWN,
			data.Metadata.Type)
	}

	data.Metadata.Name, err = data.__processVar(data.Metadata.Name)
	if err != nil {
		return data._error(libmonteur.ERROR_PROGRAM_META_NAME_BAD, err)
	}
	app.Metadata.Name = data.Metadata.Name

	//nolint:lll
	app.Metadata.Description, err = data.__processVar(data.Metadata.Description)
	if err != nil {
		return data._error(libmonteur.ERROR_PROGRAM_META_DESC_BAD, err)
	}

	return nil
}

func (data *TOMLProgram) processSource(app *Program) (err error) {
	// get all-all source and merge accordingly
	system := libmonteur.ALL_OS + libmonteur.COMPUTE_SYSTEM_SEPARATOR +
		libmonteur.ALL_ARCH
	actual := data.Sources[system]

	// get supported source
	system = data.Variables[libmonteur.VAR_COMPUTE].(string)
	ret := data.Sources[system]

	// check we actually have a working actual data
	actual = data._mergeProgramSource(actual, ret)
	if actual == nil {
		return data._error(libmonteur.ERROR_PROGRAM_UNSUPPORTED,
			system,
		)
	}

	// process format
	actual.Format, err = data.__processVar(actual.Format)
	actual.Format = strings.ToLower(actual.Format)
	switch {
	case err != nil:
		return data._error(libmonteur.ERROR_PROGRAM_ARCHIVE_FORMAT_BAD,
			actual.Format,
		)
	case actual.Format == libmonteur.PROGRAM_FORMAT_TAR_GZ:
		app.Source.Unpack = app.UnarchiveTarGz
	case actual.Format == libmonteur.PROGRAM_FORMAT_ZIP:
		app.Source.Unpack = app.UnarchiveZip
	default:
		return data._error(
			libmonteur.ERROR_PROGRAM_ARCHIVE_FORMAT_UNKNOWN,
			actual.Format,
		)
	}
	data.Variables[libmonteur.VAR_FORMAT] = actual.Format
	defer delete(data.Variables, libmonteur.VAR_FORMAT)

	// process archive
	app.Source.Archive, err = data.__processVar(actual.Archive)
	if err != nil || app.Source.Archive == "" {
		return data._error(libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			actual.Archive,
		)
	}
	data.Variables[libmonteur.VAR_ARCHIVE] = app.Source.Archive
	defer delete(data.Variables, libmonteur.VAR_ARCHIVE)

	// process method
	app.Source.Method, err = data.__processVar(actual.Method)
	if err != nil || app.Source.Method == "" {
		return data._error(libmonteur.ERROR_PROGRAM_HTTPS_METHOD_BAD,
			actual.Method,
		)
	}
	data.Variables[libmonteur.VAR_METHOD] = app.Source.Method
	defer delete(data.Variables, libmonteur.VAR_METHOD)

	// process url
	app.Source.URL, err = data.__processVar(actual.URL)
	if err != nil || app.Source.URL == "" {
		return data._error(libmonteur.ERROR_PROGRAM_URL_BAD,
			actual.URL,
		)
	}
	data.Variables[libmonteur.VAR_URL] = app.Source.URL
	defer delete(data.Variables, libmonteur.VAR_URL)

	// process headers
	app.Source.Headers = map[string]string{}
	for k, v := range actual.Headers {
		v, err = data.__processVar(v)
		if err != nil {
			//nolint: lll
			return data._error(libmonteur.ERROR_PROGRAM_HTTPS_HEADER_BAD,
				v,
			)
		}

		app.Source.Headers[k] = v
	}

	// process Checksum
	app.Source.Checksum, err = data._getChecksum(actual.Checksum)
	if err != nil {
		return err
	}

	return nil
}

func (data *TOMLProgram) _mergeProgramSource(base *_programSource,
	actual *_programSource) (out *_programSource) {
	switch {
	case base == nil && actual == nil:
		return nil
	case base == nil:
		return actual
	case actual == nil:
		return base
	}

	// both base and actual are valid. Perform merging
	if actual.Archive != "" {
		base.Archive = actual.Archive
	}

	if actual.Format != "" {
		base.Format = actual.Format
	}

	if actual.URL != "" {
		base.URL = actual.URL
	}

	if actual.Method != "" {
		base.Method = actual.Method
	}

	if actual.Checksum != nil {
		if base.Checksum == nil {
			base.Checksum = actual.Checksum
		}

		if actual.Checksum.Type != "" || base.Checksum.Type == "" {
			base.Checksum.Type = actual.Checksum.Type
		}

		if actual.Checksum.Format != "" || base.Checksum.Format == "" {
			base.Checksum.Format = actual.Checksum.Format
		}

		if actual.Checksum.Value != "" || base.Checksum.Value == "" {
			base.Checksum.Value = actual.Checksum.Value
		}
	}

	if len(actual.Headers) > 0 {
		base.Headers = map[string]string{}
		for k, v := range actual.Headers {
			base.Headers[k] = v
		}
	}

	return base
}

func (data *TOMLProgram) _getChecksum(x *_programChecksum) (h *checksum.Hasher,
	err error) {
	if x == nil {
		return nil, nil
	}

	if x.Value == "" {
		return nil, data._error(libmonteur.ERROR_CHECKSUM_BAD,
			x.Value)
	}

	if x.Type == "" {
		return nil, data._error(libmonteur.ERROR_CHECKSUM_ALGO_UNKNOWN,
			"''",
		)
	}

	if x.Format == "" {
		return nil, data._error(libmonteur.ERROR_CHECKSUM_FORMAT_BAD,
			"''",
		)
	}

	x.Type = strings.ToLower(x.Type)
	x.Format = strings.ToLower(x.Format)
	h = &checksum.Hasher{}

	switch x.Format {
	case libmonteur.CHECKSUM_FORMAT_BASE64:
		err = h.ParseBase64(x.Value)
	case libmonteur.CHECKSUM_FORMAT_HEX:
		err = h.ParseHex(x.Value)
	case libmonteur.CHECKSUM_FORMAT_BASE64_URL:
		err = h.ParseBase64URL(x.Value)
	default:
		//nolint:lll
		return nil, data._error(libmonteur.ERROR_CHECKSUM_FORMAT_UNKNOWN,
			x.Format,
		)
	}

	if err != nil {
		return nil, data._error(libmonteur.ERROR_CHECKSUM_BAD,
			err,
		)
	}

	switch x.Type {
	case libmonteur.CHECKSUM_ALGO_SHA512:
		_ = h.SetAlgo(checksum.HASHER_SHA512)
	case libmonteur.CHECKSUM_ALGO_SHA256:
		_ = h.SetAlgo(checksum.HASHER_SHA256)
	case libmonteur.CHECKSUM_ALGO_MD5:
		_ = h.SetAlgo(checksum.HASHER_MD5)
	default:
		return nil, data._error(libmonteur.ERROR_CHECKSUM_ALGO_UNKNOWN,
			x.Type,
		)
	}

	return h, nil
}

func (data *TOMLProgram) processSetup(app *Program) (err error) {
	var os, arch string
	var skip bool

	os = data.Variables[libmonteur.VAR_OS].(string)
	arch = data.Variables[libmonteur.VAR_ARCH].(string)

	for i, v := range data.Setup {
		skip, err = data._processSetupCondition(i, v, os, arch)
		if err != nil {
			return err
		}

		if skip {
			continue
		}

		x := &Setup{}
		err = data._processSetupType(i, x, v)
		if err != nil {
			return err
		}

		err = data._processSetupSource(i, x, v)
		if err != nil {
			return err
		}

		err = data._processSetupTarget(i, x, v)
		if err != nil {
			return err
		}

		// save to instruction list
		app.Setup = append(app.Setup, x)
	}

	return nil
}

func (data *TOMLProgram) _processSetupCondition(step int,
	value *_programSetup, os string, arch string) (skip bool, err error) {
	value.Condition, err = data.__processVar(value.Condition)
	if err != nil {
		return false, data._errorStep(step,
			libmonteur.ERROR_PROGRAM_INST_CONDITION_BAD,
			value.Condition,
		)
	}

	list := strings.Split(value.Condition,
		libmonteur.COMPUTE_SYSTEM_SEPARATOR)
	if len(list) != 2 {
		return false, data._errorStep(step,
			libmonteur.ERROR_PROGRAM_INST_CONDITION_BAD,
			value.Condition,
		)
	}

	switch {
	case list[0] == libmonteur.ALL_OS && list[1] == libmonteur.ALL_ARCH:
	case list[0] == libmonteur.ALL_OS && list[1] == arch:
	case list[0] == os && list[1] == libmonteur.ALL_ARCH:
	case list[0] == os && list[1] == arch:
	default:
		return true, nil // step is not for this platform. Skip.
	}

	return false, nil
}

func (data *TOMLProgram) _processSetupType(step int,
	x *Setup, value *_programSetup) (err error) {
	value.Type, err = data.__processVar(value.Type)
	if err != nil {
		return data._errorStep(step,
			libmonteur.ERROR_PROGRAM_INST_TYPE_BAD,
			err,
		)
	}

	value.Type = strings.ToLower(value.Type)
	switch value.Type {
	case libmonteur.PROGRAM_SETUP_INST_MOVE:
		x.Type = iNST_MOVE
	case libmonteur.PROGRAM_SETUP_INST_SCRIPT:
		x.Type = iNST_SCRIPT
	default:
		return data._errorStep(step,
			libmonteur.ERROR_PROGRAM_INST_TYPE_UNKNOWN,
			value.Type,
		)
	}

	return nil
}

func (data *TOMLProgram) _processSetupSource(step int,
	x *Setup, value *_programSetup) (err error) {
	x.Source, err = data.__processVar(value.Source)
	if err != nil {
		return data._errorStep(step,
			libmonteur.ERROR_PROGRAM_INST_SOURCE_BAD,
			value.Source,
		)
	}

	return nil
}

func (data *TOMLProgram) _processSetupTarget(step int,
	x *Setup, value *_programSetup) (err error) {
	x.Target, err = data.__processVar(value.Target)
	if err != nil {
		return data._errorStep(step,
			libmonteur.ERROR_PROGRAM_INST_TARGET_BAD,
			value.Target,
		)
	}

	return nil
}

func (data *TOMLProgram) processConfig(app *Program) (err error) {
	var config, system string
	var ok bool

	// initialize important variables
	system = data.Variables[libmonteur.VAR_OS].(string)

	// obtain the supported config data
	config, ok = data.Config[system]
	if !ok {
		return data._error(libmonteur.ERROR_PROGRAM_CONFIG_BAD, "''")
	}

	// process variables
	app.Config, err = data.__processVar(config)
	if err != nil {
		return data._error(libmonteur.ERROR_PROGRAM_CONFIG_BAD, err)
	}

	return nil
}

func (data *TOMLProgram) __processVar(text string) (string, error) {
	t, err := template.New("value").Parse(text)
	if err != nil {
		return text, err //nolint:wrapcheck
	}

	var b bytes.Buffer
	if err := t.Execute(&b, data.Variables); err != nil {
		return text, err //nolint:wrapcheck
	}

	return b.String(), nil
}

func (data *TOMLProgram) _error(tag string, message interface{}) (err error) {
	err = fmt.Errorf("%s for %s", tag, data.Metadata.Name)
	if message != nil {
		err = fmt.Errorf("%s: %#v", err, message)
	}

	return err
}

func (data *TOMLProgram) _errorStep(step int,
	tag string, message interface{}) (err error) {
	err = fmt.Errorf("%s for %s [Step %d]", tag, data.Metadata.Name, step)
	if message != nil {
		err = fmt.Errorf("%s: %#v", err, message)
	}

	return err
}
