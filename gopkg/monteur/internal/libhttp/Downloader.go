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

package libhttp

import (
	"context"
	"os"
	"path/filepath"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/httpclient"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libchecksum"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

func Source(ctx context.Context, source *libmonteur.TOMLSource,
	variables map[string]interface{},
	log *liblog.Logger,
	cs libchecksum.Hasher) (err error) {
	var ok bool
	var destination string

	log.Info("Sourcing %s using HTTPS download...", source.URL)

	// extract Raw directory pathing
	destination, ok = variables[libmonteur.VAR_TMP].(string)
	if !ok {
		panic("MONTEUR DEV: why is VAR_TMP not assigned?")
	}

	// process destination pathing
	destination = filepath.Join(destination, source.Archive)

	// clean up destination pathing
	_ = os.RemoveAll(destination)
	_ = os.MkdirAll(filepath.Dir(destination),
		libmonteur.PERMISSION_DIRECTORY,
	)

	// setup downloader
	d := &httpclient.Downloader{
		Destination: destination,
		Headers:     source.Headers,
	}

	d.HandleError = func(e error) {
		err = e
	}

	d.HandleSuccess = func() {
		log.Info("%s ➤ download completed", source.URL)
	}

	d.HandleProgress = func(downloaded, total int64) {
		percent := float64(downloaded) / float64(total) * 100

		log.Info("%d / %d Bytes (%.0f%%)",
			downloaded,
			total,
			percent,
		)
	}

	log.Info("Downloader Destination: %v", d.Destination)
	log.Info("Downloader HandleSuccess: %v", d.HandleSuccess)
	log.Info("Downloader HandleProgress: %v", d.HandleProgress)
	log.Info("Downloader Method: %v", source.Method)
	log.Info("Downloader URL: %v", source.URL)
	log.Info("Downloader Checksum: %v", cs)

	if len(d.Headers) == 0 {
		log.Info("Downloader Headers: {}")
	} else {
		for k, v := range d.Headers {
			log.Info("  '%s': '%s'", k, v)
		}
	}

	log.Info("Begin downloading...")
	d.Download(ctx, source.Method, source.URL, cs)

	return err
}
