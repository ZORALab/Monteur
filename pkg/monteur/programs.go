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
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/chmsg"      //nolint:typecheck
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/httpclient" //nolint:typecheck
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/targz"      //nolint:typecheck
)

type binProgram struct {
	Metadata *struct {
		Name        string
		Description string
		Type        string
	}

	Variables map[string]interface{}

	Sources map[string]*struct {
		Archive string
		Format  string
		URL     string
		Method  string
		Headers []string
	}

	Setup []*struct {
		Source    string
		Target    string
		Type      string
		Condition string
	}

	Config map[string]string

	_sourceFx    func(ctx context.Context, tx chan chmsg.Message)
	_unarchiveFx func(ctx context.Context, tx chan chmsg.Message)

	hadSanitized bool
}

// Sanitize is to ensure the parsed program data is usable.
//
// This function also processes all variables placeholders in every fields.
func (app *binProgram) Sanitize() (err error) {
	app.hadSanitized = false

	err = app._sanitizeMetadata()
	if err != nil {
		return err
	}

	err = app._sanitizeSources()
	if err != nil {
		return err
	}

	err = app._sanitizeSetupInstruction()
	if err != nil {
		return err
	}

	err = app._sanitizeConfig()
	if err != nil {
		return err
	}

	err = app._isSupported()
	if err != nil {
		return err
	}

	app.hadSanitized = true
	return nil
}

func (app *binProgram) _sanitizeMetadata() (err error) {
	// process Name
	app.Metadata.Name, err = app.__processVar(app.Metadata.Name)
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_SETUP_METADATA_NAME_BAD,
			err,
		)
	}

	// process Description
	app.Metadata.Description, err = app.__processVar(app.Metadata.Description)
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_SETUP_METADATA_DESC_BAD,
			err,
		)
	}

	// process Type
	app.Metadata.Type, err = app.__processVar(app.Metadata.Type)

	switch {
	case err != nil:
		return fmt.Errorf("%s: %s", ERROR_SETUP_TYPE_BAD, err)
	case app.Metadata.Type == BIN_PROGRAM_TYPE_HTTPS_DOWNLOAD:
		app._sourceFx = app.__sourceHTTPS
	case app.Metadata.Type == BIN_PROGRAM_TYPE_LOCAL_SYSTEM:
		app._sourceFx = app.__sourceLocal
	default:
		return fmt.Errorf("%s: %s",
			ERROR_SETUP_TYPE_UNKNOWN,
			app.Metadata.Type,
		)
	}

	return nil
}

func (app *binProgram) _sanitizeSources() (err error) {
	var headers []string

	for k, v := range app.Sources {
		// process Format
		v.Format, err = app.__processVar(v.Format)
		if err != nil {
			return fmt.Errorf("%s%s: [%s] %s for %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_SOURCE_ARCHIVE_FORMAT_BAD,
				k,
				v.Format,
				app.Metadata.Name,
			)
		}

		switch v.Format {
		case BIN_PROGRAM_FORMAT_TAR_GZ:
			app._unarchiveFx = app.__unarchiveTarGz
		case BIN_PROGRAM_FORMAT_ZIP:
			app._unarchiveFx = app.__unarchiveZip
		default:
			return fmt.Errorf("%s%s: [%s] %s for %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_SOURCE_ARCHIVE_FORMAT_UNKNOWN,
				k,
				v.Format,
				app.Metadata.Name,
			)
		}

		v.Format = strings.ToLower(v.Format)
		app.Variables[BIN_PROGRAM_VAR_FORMAT] = v.Format
		defer delete(app.Variables, BIN_PROGRAM_VAR_FORMAT)

		// process Archive
		v.Archive, err = app.__processVar(v.Archive)
		if err != nil {
			delete(app.Variables, BIN_PROGRAM_VAR_FORMAT)

			return fmt.Errorf("%s%s: [%s] %s for %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_SOURCE_ARCHIVE_BAD,
				k,
				v.Archive,
				app.Metadata.Name,
			)
		}

		app.Variables[BIN_PROGRAM_VAR_ARCHIVE] = v.Archive
		defer delete(app.Variables, BIN_PROGRAM_VAR_ARCHIVE)

		// process Method
		v.Method, err = app.__processVar(v.Method)
		if err != nil {
			return fmt.Errorf("%s%s: [%s] %s for %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_SOURCE_METHOD_BAD,
				k,
				v.Method,
				app.Metadata.Name,
			)
		}

		app.Variables[BIN_PROGRAM_VAR_METHOD] = v.Method
		defer delete(app.Variables, BIN_PROGRAM_VAR_METHOD)

		// process URL
		v.URL, err = app.__processVar(v.URL)
		if err != nil {
			return fmt.Errorf("%s%s: [%s] %s for %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_SOURCE_BAD,
				k,
				v.URL,
				app.Metadata.Name,
			)
		}

		app.Variables[BIN_PROGRAM_VAR_URL] = v.URL
		defer delete(app.Variables, BIN_PROGRAM_VAR_URL)

		// process Headers
		headers = []string{}
		for _, h := range v.Headers {
			h, err = app.__processVar(h)
			if err != nil {
				return fmt.Errorf("%s%s: [%s] %s for %s",
					ERROR_SETUP_TAG,
					ERROR_SETUP_SOURCE_HEADER_BAD,
					h,
					app.Metadata.Name,
				)
			}

			headers = append(headers, h)
		}
		v.Headers = headers

		// update structure to the latest
		app.Sources[k] = v
	}

	// delete localized variables
	delete(app.Variables, BIN_PROGRAM_VAR_FORMAT)
	delete(app.Variables, BIN_PROGRAM_VAR_ARCHIVE)

	return nil
}

func (app *binProgram) _sanitizeSetupInstruction() (err error) {
	for k, v := range app.Setup {
		// process Type
		v.Type, err = app.__processVar(v.Type)
		if err != nil || v.Type == "" {
			return fmt.Errorf("%s%s: [%s] %s for %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_INSTRUCTION_TYPE_BAD,
				k,
				v.Type,
				app.Metadata.Name,
			)
		}

		v.Type = strings.ToLower(v.Type)

		// process Conditions
		v.Condition, err = app.__processVar(v.Condition)
		if err != nil {
			return err
		}

		list := strings.Split(v.Condition, COMPUTE_SYSTEM_SEPARATOR)
		if len(list) != 2 {
			return fmt.Errorf("%s%s: %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_INSTRUCTION_CONDITION_BAD,
				v.Condition,
			)
		}

		if strings.ToLower(list[0]) == "all" {
			list[0] = app.Variables[BIN_PROGRAM_VAR_OS].(string)
		}

		if strings.ToLower(list[1]) == "all" {
			list[1] = app.Variables[BIN_PROGRAM_VAR_ARCH].(string)
		}

		v.Condition = strings.ToLower(list[0]) +
			COMPUTE_SYSTEM_SEPARATOR +
			strings.ToLower(list[1])

		// process Source
		v.Source, err = app.__processVar(v.Source)
		if err != nil {
			return fmt.Errorf("%s%s: %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_INSTRUCTION_SOURCE_BAD,
				v.Source,
			)
		}

		// process Target
		v.Target, err = app.__processVar(v.Target)
		if err != nil {
			return fmt.Errorf("%s%s: %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_INSTRUCTION_TARGET_BAD,
				v.Target,
			)
		}

		// update structure to the latest
		app.Setup[k] = v
	}

	return nil
}

func (app *binProgram) _sanitizeConfig() (err error) {
	for k, v := range app.Config {
		v, err = app.__processVar(v)
		if err != nil {
			return fmt.Errorf("%s%s: %s for %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_POSTCONFIG_BAD,
				k,
				v,
			)
		}

		// update structure to the latest
		app.Config[k] = v
	}

	// check local config file is available for current OS
	_, ok := app.Config[app.Variables[BIN_PROGRAM_VAR_OS].(string)]
	if !ok {
		return fmt.Errorf("%s%s: %s",
			ERROR_SETUP_TAG,
			ERROR_SETUP_CONFIG_MISSING,
			app.Metadata.Name,
		)
	}

	return nil
}

func (app *binProgram) _isSupported() (err error) {
	if app.Sources[app.Variables[BIN_PROGRAM_VAR_COMPUTE].(string)] == nil {
		return fmt.Errorf("%s%s: %s for %s",
			ERROR_SETUP_TAG,
			ERROR_SETUP_PROGRAM_NOT_SUPPORTED,
			app.Metadata.Name,
			app.Variables[BIN_PROGRAM_VAR_COMPUTE],
		)
	}

	return nil
}

func (app *binProgram) __processVar(text string) (string, error) {
	t, err := template.New("value").Parse(text)
	if err != nil {
		return text, err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, app.Variables); err != nil {
		return text, err
	}

	return b.String(), nil
}

// Get is to obtain the program from its supported source.
//
// This should be done after the binProgram has been sanitized. The function
// is designed to be multi-threaded or executed parallelly.
func (app *binProgram) Get(ctx context.Context, tx chan chmsg.Message) {
	var rx chan chmsg.Message
	var msg chmsg.Message

	if !app.hadSanitized {
		msg = chmsg.New()
		msg.Add(chmsg_ERROR, fmt.Errorf("%s%s: %s\n",
			ERROR_SETUP_TAG,
			ERROR_SETUP_PROGRAM_NOT_SANITIZED,
			app.Metadata.Name,
		))
		tx <- msg

		return
	}

	// run the sourcing in background so that we can cancel per context
	rx = make(chan chmsg.Message)
	go func() {
		app._sourceFx(ctx, rx)
		app._unarchiveFx(ctx, rx)
		app._setupFx(ctx, rx)
		app._configFx(ctx, rx)

		// send completed message and close the channel
		msg = chmsg.New()
		msg.Add(chmsg_DONE, true)
		rx <- msg
		close(rx)
	}()

	// wait for signals
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-rx:
			if ok {
				tx <- msg

				v, ok := msg.Get(chmsg_DONE)
				if ok && v.(bool) {
					return
				}
			}
		}
	}
}

func (app *binProgram) _setupFx(ctx context.Context, tx chan chmsg.Message) {
	var msg chmsg.Message

	for k, step := range app.Setup {
		if step.Condition != app.Variables[BIN_PROGRAM_VAR_COMPUTE] {
			continue // skip unsupported platforms
		}

		switch step.Type {
		case BIN_PROGRAM_SETUP_INSTRUCTION_MOVE:
			app.__setupMove(ctx, tx, step.Source, step.Target)
		case BIN_PROGRAM_SETUP_INSTRUCTION_SCRIPT:
			app.__setupScript(ctx, tx, step.Source, step.Target)
		default:
			// mercifully report the error instead of stopping the
			// entire process.
			msg = chmsg.New()
			msg.Add(chmsg_ERROR, fmt.Errorf("%s%s: [%s] %s for %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_INSTRUCTION_TYPE_UNKNOWN,
				k,
				step.Type,
				app.Metadata.Name,
			))
			tx <- msg
		}
	}
}

func (app *binProgram) _configFx(ctx context.Context, tx chan chmsg.Message) {
	var msg chmsg.Message
	var pathing string
	var err error

	// obtain the corresponding config file
	data := app.Config[app.Variables[BIN_PROGRAM_VAR_OS].(string)]

	// process the file pathing
	pathing = strings.ToLower(app.Metadata.Name)
	pathing = strings.Replace(pathing, " ", "-", -1)
	pathing = strings.Replace(pathing, "_", "-", -1)
	pathing = strings.Replace(pathing, "%", "-", -1)
	pathing = strings.Replace(pathing, "!", "-", -1)
	pathing = filepath.Join(app.Variables[BIN_PROGRAM_VAR_CFG].(string),
		pathing)

	// attempt to write into config directory
	err = os.WriteFile(pathing, []byte(data), SETUP_CONFIG_PERMISSION)
	if err != nil {
		msg = chmsg.New()
		msg.Add(chmsg_ERROR, fmt.Errorf("%s%s: [%s] %s",
			ERROR_SETUP_TAG,
			ERROR_SETUP_CONFIG_FAILED,
			pathing,
			err,
		))
		tx <- msg
		return
	}
}

func (app *binProgram) __setupMove(ctx context.Context,
	tx chan chmsg.Message,
	source string,
	target string) {
	var err error
	var msg chmsg.Message

	// formulate pathing
	source = filepath.Join(app.Variables[BIN_PROGRAM_VAR_TMP].(string),
		source)
	target = filepath.Join(app.Variables[BIN_PROGRAM_VAR_BIN].(string),
		target)

	// check source is available
	if _, err = os.Stat(source); os.IsNotExist(err) {
		msg = chmsg.New()
		msg.Add(chmsg_ERROR, fmt.Errorf("%s%s: %s",
			ERROR_SETUP_TAG,
			ERROR_SETUP_INSTALL_MOVE_FAILED,
			err,
		))
		tx <- msg
		return
	}

	// remove the target regardlessly
	_ = os.RemoveAll(target)

	// move the source to the target
	err = os.Rename(source, target)
	if err != nil {
		msg = chmsg.New()
		msg.Add(chmsg_ERROR, fmt.Errorf("%s%s: %s -> %s %s",
			ERROR_SETUP_TAG,
			ERROR_SETUP_INSTALL_MOVE_FAILED,
			source,
			target,
			err,
		))
		tx <- msg
		return
	}
}

func (app *binProgram) __setupScript(ctx context.Context,
	tx chan chmsg.Message,
	source string,
	target string) {
	var err error
	var msg chmsg.Message

	// formulate pathing
	target = filepath.Join(app.Variables[BIN_PROGRAM_VAR_BIN].(string),
		target)

	// remove the target regardlessly
	_ = os.RemoveAll(target)

	// create script from source
	err = os.WriteFile(target, []byte(source), SETUP_PROGRAMS_PERMISSION)
	if err != nil {
		msg = chmsg.New()
		msg.Add(chmsg_ERROR, fmt.Errorf("%s%s: [%s] %s",
			ERROR_SETUP_TAG,
			ERROR_SETUP_INSTALL_SCRIPT_FAILED,
			target,
			err,
		))
		tx <- msg
		return
	}
}

func (app *binProgram) __sourceHTTPS(ctx context.Context,
	tx chan chmsg.Message) {
	var d *httpclient.Downloader
	var msg chmsg.Message
	var percent float64

	// get sources
	source := app.Sources[app.Variables[BIN_PROGRAM_VAR_COMPUTE].(string)]

	// create downloader
	d = &httpclient.Downloader{
		HandleError: func(err error) {
			msg = chmsg.New()
			msg.Add(chmsg_ERROR, fmt.Errorf("%s%s: %s",
				ERROR_SETUP_TAG,
				ERROR_SETUP_HTTPS_DOWNLOAD_FAILED,
				err,
			))

			tx <- msg
		},
		HandleProgress: func(p, t int64) {
			percent = float64(p) / float64(t) * 100

			msg = chmsg.New()
			msg.Add(chmsg_STATUS,
				fmt.Sprintf("%-10s: %d / %d Bytes (%.0f%%)",
					app.Metadata.Name,
					p,
					t,
					percent,
				))

			tx <- msg
		},
		HandleSuccess: func() {
		},
		Destination: app.Variables[BIN_PROGRAM_VAR_TMP].(string),
	}

	// download the content
	if source.Method == "" {
		source.Method = "GET"
	}

	d.Download(ctx, source.Method, source.URL, nil)
}

func (app *binProgram) __sourceLocal(ctx context.Context,
	tx chan chmsg.Message) {
	// os.LookPath
}

func (app *binProgram) __unarchiveTarGz(ctx context.Context,
	tx chan chmsg.Message) {
	var err error
	var msg chmsg.Message

	// populate all sources
	source := app.Sources[app.Variables[BIN_PROGRAM_VAR_COMPUTE].(string)]
	dest := app.Variables[BIN_PROGRAM_VAR_TMP].(string)
	src := filepath.Join(app.Variables[BIN_PROGRAM_VAR_TMP].(string),
		source.Archive)

	// create processor
	processor := &targz.Processor{
		Archive: src,
		Raw:     dest,
	}

	// proceed to extract payload
	err = processor.Extract()
	if err != nil {
		msg = chmsg.New()
		msg.Add(chmsg_ERROR, fmt.Errorf("%s%s: %s",
			ERROR_SETUP_TAG,
			ERROR_SETUP_HTTPS_DOWNLOAD_FAILED,
			err,
		))

		tx <- msg
	}
}

func (app *binProgram) __unarchiveZip(ctx context.Context,
	tx chan chmsg.Message) {
}
