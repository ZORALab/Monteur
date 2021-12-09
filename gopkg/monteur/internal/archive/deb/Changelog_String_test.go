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

package deb

import (
	"testing"
)

func TestChangelogString(t *testing.T) {
	var panicError interface{}

	for i, s := range getTestScenarios() {
		if s.TestType != testChangelogString {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)
		entity := s.createEntity()
		timestamp := s.createTimestamp()
		version := s.createVersion()
		app := s.createAppName()
		urgency := s.createChangelogUrgency()
		changes := s.createChangelogChanges()
		distribution := s.createChangelogDistro()
		path := s.createPath()
		statFx := s.createStatFx()
		panicError = noPanicFlag
		str := ""

		// test
		subject := &Changelog{
			Version:      version,
			Maintainer:   entity,
			Timestamp:    timestamp,
			Package:      app,
			Urgency:      urgency,
			Changes:      changes,
			Distribution: distribution,
			Path:         path,
			statFx:       statFx,
		}

		t.Run(s.stringUID(), func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					panicError = err
				}
			}()
			str = subject.String()
		})

		// assert
		th.ExpectUIDCorrectness(i, s.UID, false)
		s.assertChangelog(th, str)
		s.assertPanic(th, panicError)
		s.log(th, map[string]interface{}{
			"given package":    app,
			"given version":    version,
			"given maintainer": entity,
			"given urgency":    urgency,
			"given changes":    changes,
			"given distro":     distribution,
			"given path":       path,
			"given statFx":     statFx,
			"created subject":  subject,
			"got":              str,
			"panic error":      panicError,
		})
		th.Logf(`TESTNOTE: Changelog Entry can only be fully tested
visually
═══════ START ══════
%s
═══════  END  ══════`, str)
		th.Conclude()
	}
}
