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

package libsetup

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/chmsg"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/httpclient"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/targz"
)

// Program is the supported setup program to source
type Program struct {
	Metadata      *Metadata
	WorkspacePath string
	InstallPath   string
	ConfigPath    string

	Config string
	Source *Source

	ReportUp chan chmsg.Message

	Setup []*Setup
}

// Shop is to get the Program from its source
//
// Everything must be setup properly before calling this function. It was meant
// for Monteur's Setup API.
func (app *Program) Shop(ctx context.Context) {
	app.Source.Get(ctx)
	app.Source.Unpack(ctx)
	app.Install(ctx)
	app.PostConfigure(ctx)
	app.ReportDone()
}

func (app *Program) SourceHTTPS(ctx context.Context) {
	d := &httpclient.Downloader{
		HandleError:   app.ReportError,
		HandleSuccess: app.Source.HandleSuccess,
		Destination:   app.WorkspacePath,
	}

	d.HandleProgress = func(downloaded, total int64) {
		percent := float64(downloaded) / float64(total) * 100

		app.ReportStatus(fmt.Sprintf("%-10s: %d / %d B (%.0f%%)",
			app.Metadata.Name,
			downloaded,
			total,
			percent,
		))
	}

	d.Download(ctx, app.Source.Method, app.Source.URL, app.Source.Checksum)
}

func (app *Program) SourceLocal(ctx context.Context) {
	//TODO: os.LookPath
}

func (app *Program) Install(ctx context.Context) {
	for i, step := range app.Setup {
		switch step.Type {
		case INST_MOVE:
			app.move(step.Source, step.Target)
		case INST_SCRIPT:
			app.script(step.Source, step.Target)
		case INST_UNKNOWN:
			fallthrough
		default:
			app.ReportError(fmt.Errorf("%s: [Step %d] type %d",
				"unsupported setup instruction",
				i,
				step.Type,
			))

			return
		}
	}
}

func (app *Program) PostConfigure(ctx context.Context) {
	var pathing string
	var err error

	// process pathing
	pathing = strings.ToLower(app.Metadata.Name)
	pathing = strings.ReplaceAll(pathing, " ", "-")
	pathing = strings.ReplaceAll(pathing, "_", "-")
	pathing = strings.ReplaceAll(pathing, "%", "-")
	pathing = strings.ReplaceAll(pathing, "!", "-")
	pathing = filepath.Join(app.ConfigPath, pathing)

	// write into config directory
	err = os.WriteFile(pathing, []byte(app.Config), CONFIG_PERMISSION)
	if err != nil {
		app.ReportError(fmt.Errorf("%s: %s",
			"unable to write config file",
			err,
		))
	}
}

func (app *Program) move(source string, target string) {
	var err error

	source = filepath.Join(app.WorkspacePath, source)
	target = filepath.Join(app.InstallPath, target)

	// check source is available to move
	if _, err = os.Stat(source); os.IsNotExist(err) {
		app.ReportError(fmt.Errorf("%s: %s",
			"source file is missing for move",
			source,
		))
		return
	}

	// remove target regardlessly
	_ = os.RemoveAll(target)

	// move the source to target
	err = os.Rename(source, target)
	if err != nil {
		app.ReportError(fmt.Errorf("%s: %s",
			"setup move failed",
			err,
		))
	}
}

func (app *Program) script(source string, target string) {
	var err error

	target = filepath.Join(app.WorkspacePath, target)

	// remove the target regardlessly
	_ = os.RemoveAll(target)

	// create script from source
	err = os.WriteFile(target, []byte(source), EXECUTABLE_PERMISSION)
	if err != nil {
		app.ReportError(fmt.Errorf("%s: %s",
			"setup script failed",
			err,
		))
	}
}

func (app *Program) UnarchiveTarGz(ctx context.Context) {
	var err error

	processor := &targz.Processor{
		Archive: filepath.Join(app.WorkspacePath, app.Source.Archive),
		Raw:     app.WorkspacePath,
	}

	err = processor.Extract()
	if err != nil {
		app.ReportError(err)
		return
	}
}

func (app *Program) UnarchiveZip(ctx context.Context) {
	//TODO: zip archives
}

func (app *Program) ReportError(err error) {
	if app.ReportUp == nil {
		return
	}

	msg := chmsg.New()
	msg.Add(libmonteur.CHMSG_OWNER, app.Metadata.Name)
	msg.Add(libmonteur.CHMSG_ERROR, err)
	app.ReportUp <- msg
}

func (app *Program) ReportStatus(message string) {
	if app.ReportUp == nil {
		return
	}

	msg := chmsg.New()
	msg.Add(libmonteur.CHMSG_OWNER, app.Metadata.Name)
	msg.Add(libmonteur.CHMSG_STATUS, message)
	app.ReportUp <- msg
}

func (app *Program) ReportDone() {
	if app.ReportUp == nil {
		return
	}

	msg := chmsg.New()
	msg.Add(libmonteur.CHMSG_OWNER, app.Metadata.Name)
	msg.Add(libmonteur.CHMSG_DONE, true)
	app.ReportUp <- msg
}
