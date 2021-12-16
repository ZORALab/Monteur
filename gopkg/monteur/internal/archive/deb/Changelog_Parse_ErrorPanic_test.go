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

func TestChangelogParseForErrorPanic(t *testing.T) {
	var panicError interface{}
	var err error

	for i, s := range getTestScenarios() {
		if s.TestType != testChangelogParseForErrorPanic {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)
		statFx := s.createStatFx()
		status := s.createChangelogParseStatus()
		lines := s.createChangelogData()

		// test
		subject := &Changelog{
			statFx:      statFx,
			parseStatus: status,
		}

		panicError = noPanicFlag
		err = nil
		for _, line := range lines {
			t.Run(s.stringUID(), func(t *testing.T) {
				defer func() {
					if err := recover(); err != nil {
						panicError = err
					}
				}()
				_, err = subject.Parse(line)
			})
		}

		// assert
		th.ExpectUIDCorrectness(i, s.UID, false)
		s.assertChangelogParseForErrorPanic(th, err)
		s.assertPanic(th, panicError)
		s.log(th, map[string]interface{}{
			"given status":    status,
			"given lines":     lines,
			"given statFx":    statFx,
			"created subject": subject,
			"got error":       err,
			"panic error":     panicError,
		})
		th.Conclude()
	}
}
