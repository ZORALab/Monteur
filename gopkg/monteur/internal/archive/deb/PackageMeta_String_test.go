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

func TestPackageMetaString(t *testing.T) {
	var panicError interface{}

	for i, s := range getTestScenarios() {
		if s.TestType != testPackageMetaString {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)
		name := s.createAppName()
		version := s.createVersion()
		control := s.createVERControl()
		arch := s.createArchitectures()
		panicError = noPanicFlag
		str := ""

		// test
		subject := &PackageMeta{
			Name:          name,
			Version:       version,
			VERControl:    control,
			Architectures: arch,
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
		s.assertPackageMeta(th, str)
		s.assertPanic(th, panicError)
		s.log(th, map[string]interface{}{
			"given name":          name,
			"given version":       version,
			"given VERControl":    control,
			"given architectures": arch,
			"created subject":     subject,
			"got":                 str,
			"panic error":         panicError,
		})
		th.Conclude()
	}
}
