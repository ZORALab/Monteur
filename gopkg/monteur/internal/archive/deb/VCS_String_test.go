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

func TestVCSString(t *testing.T) {
	var panicError interface{}

	for i, s := range getTestScenarios() {
		if s.TestType != testVCSString {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)
		vcsType := s.createVCSType()
		url := s.createURL()
		branch := s.createBranch()
		path := s.createPath()
		browser := s.createBrowser()
		panicError = noPanicFlag
		str := ""

		// test
		subject := &VCS{
			Browser: browser,
			Type:    vcsType,
			URL:     url,
			Branch:  branch,
			Path:    path,
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
		s.assertVCS(th, str)
		s.assertPanic(th, panicError)
		s.log(th, map[string]interface{}{
			"given VCS type":  vcsType,
			"given url":       url,
			"given browser":   browser,
			"given branch":    branch,
			"given path":      path,
			"created subject": subject,
			"got":             str,
			"panic error":     panicError,
		})
		th.Conclude()
	}
}
