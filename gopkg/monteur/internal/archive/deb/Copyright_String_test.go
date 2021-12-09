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

func TestCopyrightString(t *testing.T) {
	var panicError interface{}

	for i, s := range getTestScenarios() {
		if s.TestType != testCopyrightString {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)
		format := s.createCopyrightFormat()
		name := s.createAppName()
		entity := s.createEntity()
		source := s.createCopyrightSource()
		disclaimer := s.createCopyrightDisclaimer()
		_, license, _, comment := s.createLicense()
		copyright := s.createEntities()
		licenses := s.createLicenses()
		panicError = noPanicFlag
		str := ""

		// test
		subject := &Copyright{
			Format:     format,
			Name:       name,
			Contact:    entity,
			Source:     source,
			Disclaimer: disclaimer,
			Comment:    comment,
			License:    license,
			Copyright:  copyright,
			Licenses:   licenses,
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
		s.assertCopyright(th, str)
		s.assertPanic(th, panicError)
		s.log(th, map[string]interface{}{
			"given format":           format,
			"given name":             name,
			"given upstream contact": entity,
			"given source":           source,
			"given disclaimer":       disclaimer,
			"given comment":          comment,
			"given license":          license,
			"given copyrights":       copyright,
			"entity":                 subject,
			"got":                    str,
			"panic error":            panicError,
		})
		th.Logf(`TESTNOTE: Copyright full body can only be fully tested
visually
═══════ START ══════
%s
═══════  END  ══════`, str)
		th.Conclude()
	}
}
