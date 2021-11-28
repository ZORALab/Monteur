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

func TestRemove(t *testing.T) {
	scenarios := getTestScenarios()

	for i, s := range scenarios {
		if s.TestType != testRemove {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)

		writer := s.generateWriter()
		_, label := s.generateLabel()
		statusType := s.generateStatusType()

		log := New()
		log.statusWriters[label] = writer
		log.outputWriters[label] = writer

		// test
		log.Remove(statusType, label)

		// assert
		th.ExpectUIDCorrectness(i, s.UID, false)
		s.assertRemove(th, log, statusType, label)
		s.log(th, map[string]interface{}{
			"writer":            writer,
			"label":             label,
			"status type":       statusType,
			"log":               log,
			"log.statusWriters": log.statusWriters,
			"log.outputWriters": log.outputWriters,
		})
		th.Conclude()
	}
}
