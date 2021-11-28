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

func TestPrintln(t *testing.T) {
	scenarios := getTestScenarios()

	for i, s := range scenarios {
		if s.TestType != testPrintln {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)
		arguments := s.generateMessageArguments()
		outputWriter := s.generateWriter()
		statusWriter := s.generateWriter()
		outputLabel, statusLabel := s.generateLabel()
		psFx := s.generatePreProcessor()
		ts := s.generateTimestamp()

		log := New()
		log.statusWriters[statusLabel] = statusWriter
		log.outputWriters[outputLabel] = outputWriter

		// test
		log.SetPreprocessor(psFx)
		log.SetTimestamp(ts)
		ret, err := log.Println(arguments...)

		// assert
		th.ExpectUIDCorrectness(i, s.UID, false)
		s.assertArguments(th,
			log,
			outputLabel,
			statusLabel,
			arguments,
		)
		s.assertPreProcessing(th,
			log,
			outputLabel,
			statusLabel,
		)
		s.assertTimestamp(th,
			log,
			outputLabel,
			statusLabel,
			ts,
		)
		th.ExpectError(err, false)
		th.ExpectSameBool("ret == 0", ret == 0,
			"expect", true,
		)
		s.log(th, map[string]interface{}{
			"argument":          arguments,
			"print output":      outputWriter.String(),
			"print status":      statusWriter.String(),
			"output label":      outputLabel,
			"status label":      statusLabel,
			"preprocessor":      psFx,
			"timestamp":         ts,
			"log":               log,
			"log.statusWriters": log.statusWriters,
			"log.outputWriters": log.outputWriters,
		})
		th.Conclude()
	}
}
