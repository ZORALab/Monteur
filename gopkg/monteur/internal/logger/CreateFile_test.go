// Copyright 2020 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2020 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
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

package logger

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	scenarios := getTestScenarios()

	for i, s := range scenarios {
		if s.TestType != testCreateFile {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)
		path := s.generateFilepath()

		// test
		f, err := CreateFile(path)
		if f != nil {
			f.Close()
		}

		// assert
		th.ExpectUIDCorrectness(i, s.UID, false)
		th.ExpectError(err, s.expectError())
		s.assertCreateFile(th, f)
		s.cleanUp()
		s.log(th, map[string]interface{}{
			"error": err,
			"file":  f,
			"path":  path,
		})
		th.Conclude()
	}
}
