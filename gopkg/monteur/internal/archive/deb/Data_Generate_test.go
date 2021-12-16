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

func TestDataGenerate(t *testing.T) {
	var panicError interface{}
	var err error

	for i, s := range getTestScenarios() {
		if s.TestType != testDataGenerate {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)

		s.preConfigureTestedData()
		s.setupDataGenerate()

		control := s.createControl()
		copyright := s.createCopyright()
		changelog := s.createChangelog()
		source := s.createSource()
		manpage := s.createManpage()
		scripts := s.createScripts()
		installs := s.createInstalls()
		rules := s.createRules()
		compat := s.createCompat()
		workPath := s.createWorkDir()

		// test
		subject := &Data{
			Control:   control,
			Copyright: copyright,
			Changelog: changelog,
			Source:    source,
			Manpage:   manpage,
			Scripts:   scripts,
			Install:   installs,
			Rules:     rules,
			Compat:    compat,
		}

		panicError = noPanicFlag
		err = nil
		t.Run(s.stringUID(), func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					panicError = err
				}
			}()

			err = subject.Generate(workPath)
		})

		// assert
		th.ExpectUIDCorrectness(i, s.UID, false)
		s.assertDataGenerate(th, err)
		s.assertPanic(th, panicError)
		s.log(th, map[string]interface{}{
			"given control":   control,
			"given copyright": copyright,
			"given changelog": changelog,
			"given source":    source,
			"given manpage":   manpage,
			"given scripts":   scripts,
			"given installs":  installs,
			"given rules":     rules,
			"given compat":    compat,
			"got error":       err,
			"got panic error": panicError,
			"created subject": subject,
		})
		s.logFiles(th, err)
		th.Conclude()
		s.cleanUpDataGenerate()
	}
}
