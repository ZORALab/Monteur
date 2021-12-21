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
	"fmt"
	"regexp"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

type changelog struct {
	fxSTDOUT  func(string, ...interface{})
	fxSTDERR  func(string, ...interface{})
	variables *map[string]interface{}
	changelog *libmonteur.TOMLChangelog
	log       *liblog.Logger
}

// Exec instructs the changelog to run all the comamnds and generate its
// entries.
func (me *changelog) Exec() (err error) {
	var task *executive
	var r *regexp.Regexp
	var entries, out, s, line string
	var i int
	var list, outList []string

	me.log.Info("Executing Changelog Processing...")
	if len(me.changelog.Entries) > 0 {
		me.log.Info("Found Manual Entires: '%#v'", me.changelog.Entries)
		goto done
	}

	me.log.Info("Checking LineBreak...")
	if me.changelog.LineBreak == "" {
		err = fmt.Errorf(libmonteur.ERROR_CHANGELOG_LINE_BREAK_MISSING)
		goto done
	}
	me.log.Info("Got: %#v", me.changelog.LineBreak)

	me.log.Info("Executing Changelog Task Commands...")
	task = &executive{
		log:       me.log,
		variables: *me.variables,
		orders:    me.changelog.CMD,
		fxSTDOUT:  me.fxSTDOUT,
		fxSTDERR:  me.fxSTDERR,
	}

	err = task.Exec()
	if err != nil {
		goto done
	}

	me.log.Info("Getting %s ...", libmonteur.VAR_CHANGELOG_ENTRIES)
	entries = (*me.variables)[libmonteur.VAR_CHANGELOG_ENTRIES].(string)
	if entries == "" {
		err = fmt.Errorf(libmonteur.ERROR_CHANGELOG_ENTIRES_MISSING)
		goto done
	}
	me.changelog.Entries = strings.Split(entries, me.changelog.LineBreak)

	// filter regex line by line if available
	if me.changelog.Regex == "" {
		err = nil
		goto save
	}

	me.log.Info("Found Regex Filter '%s', cleaning...", me.changelog.Regex)
	r, err = regexp.Compile(me.changelog.Regex)
	if err != nil {
		err = fmt.Errorf("%s: %s",
			libmonteur.ERROR_CHANGELOG_REGEX_BAD,
			err,
		)

		me.log.Info("Regex Filtering ➤ FAILED")
		goto save
	}

	list = me.changelog.Entries
	me.changelog.Entries = []string{}
	for _, line = range list {
		outList = r.FindStringSubmatch(line)
		out = ""

		for i, s = range outList {
			if i == 0 {
				continue
			}

			out += s
		}

		me.changelog.Entries = append(me.changelog.Entries, out)
	}

save:
	(*me.variables)[libmonteur.VAR_CHANGELOG_ENTRIES] = me.changelog.Entries
	me.log.Info("Got: '%#v'", me.changelog.Entries)
done:
	me.log.Info("Executing Changelog Processing ➤ DONE\n\n")
	return err
}
