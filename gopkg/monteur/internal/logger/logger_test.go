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
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"gitlab.com/zoralab/cerigo/testing/thelper"
)

const (
	testAdd          = "testAdd"
	testCreateFile   = "testCreateFile"
	testCreateStderr = "testCreateStderr"
	testCreateStdout = "testCreateStdout"
	testRemove       = "testRemove"
	testLogf         = "testLogf"
	testLogln        = "testLogln"
	testPrintf       = "testPrintf"
	testPrintln      = "testPrintln"
)

const (
	expectError  = "expectError"
	expectRetain = "expectRetain"

	provideArgument      = "provideArgument"
	provideNilArgument   = "provideNilArgument"
	provideEmptyArgument = "provideEmptyArgument"

	useCustomTimestamp = "useCustomTimestamp"

	useEmptyFormat     = "useEmptyFormat"
	useMessageFormat   = "useMessageFormat"
	useStatementFormat = "useStatementFormat"

	useEmptyLabel  = "useEmptyLabel"
	useProperLabel = "useProperLabel"

	useEmptyLevel  = "useEmptyLevel"
	useProperLevel = "useProperLevel"

	usePreProcessing = "usePreProcessing"

	useTypeOutput = "useTypeOutput"
	useTypeStatus = "useTypeStatus"

	useEmptyFilepath     = "useEmptyFilepath"
	useProperFilepath    = "useProperFilepath"
	useDirectoryFilepath = "useDirectoryFilepath"
)

const (
	argument          = "CuStOm-ARgUM3nt"
	argumentFormat    = "the value is: %v\n"
	customTimestamp   = "customTimestamp"
	directoryFilepath = "."
	emptyFormat       = ""
	emptyLevel        = ""
	outLabel          = "outLog"
	preprocessor      = "preprocessed"
	properFilepath    = "test.log"
	properLevel       = "test-info"
	statusLabel       = "statusLog"
	statementFormat   = "this is a statement"
)

type testScenario thelper.Scenario

func (s *testScenario) prepareTHelper(t *testing.T) *thelper.THelper {
	return thelper.NewTHelper(t)
}

func (s *testScenario) log(th *thelper.THelper,
	data map[string]interface{}) {
	th.LogScenario(thelper.Scenario(*s), data)
}

func (s *testScenario) assertTag(th *thelper.THelper,
	status, output, tag string) {
	x := strings.ToUpper(tag)

	if s.Switches[useTypeStatus] {
		th.ExpectStringHasKeywords("statusWriter", status, "tag", x)
	}

	if s.Switches[useTypeOutput] {
		th.ExpectStringHasKeywords("outputWriter", output, "tag", x)
	}
}

func (s *testScenario) assertArguments(th *thelper.THelper,
	log *Logger,
	outputLabel string,
	statusLabel string,
	arguments []interface{}) {
	out := log.outputWriters[outputLabel].(*bytes.Buffer).String()
	status := log.statusWriters[statusLabel].(*bytes.Buffer).String()

	switch {
	case s.Switches[provideArgument]:
		s._assertArguments(th, out, status, arguments)
	case s.Switches[provideEmptyArgument]:
		s._assertEmptyArgument(th, out, status)
	case s.Switches[provideNilArgument]:
		s._assertNilArgument(th, out, status)
	default:
	}
}

func (s *testScenario) _assertArguments(th *thelper.THelper,
	out string, status string, arguments []interface{}) {
	var arg string

	for _, argx := range arguments {
		arg = fmt.Sprintf("%v", argx)
		if s.Switches[useTypeOutput] && !strings.Contains(out, arg) {
			th.Errorf("output missing args: %s", arg)
		}

		if s.Switches[useTypeStatus] && !strings.Contains(status, arg) {
			th.Errorf("status missing args: %s", arg)
		}
	}
}

func (s *testScenario) _assertEmptyArgument(th *thelper.THelper,
	out string, status string) {
	if s.Switches[useTypeOutput] && strings.Contains(out, argument) {
		th.Errorf("output args contains %v", argument)
	}

	if s.Switches[useTypeStatus] && strings.Contains(status, argument) {
		th.Errorf("status args contains %v", argument)
	}
}

func (s *testScenario) _assertNilArgument(th *thelper.THelper,
	out string, status string) {
	if s.Switches[useTypeOutput] && strings.Contains(out, argument) {
		th.Errorf("output args contains %v", argument)
	}

	if s.Switches[useTypeStatus] && strings.Contains(status, argument) {
		th.Errorf("status args contains %v", argument)
	}
}

func (s *testScenario) assertPreProcessing(th *thelper.THelper,
	log *Logger,
	outputLabel string,
	statusLabel string) {
	out := log.outputWriters[outputLabel].(*bytes.Buffer).String()
	status := log.statusWriters[statusLabel].(*bytes.Buffer).String()

	if !s.Switches[usePreProcessing] {
		if s.Switches[useTypeOutput] &&
			strings.Contains(out, preprocessor) {
			th.Errorf("output preprocessed despite no to")
		}

		if s.Switches[useTypeStatus] &&
			strings.Contains(status, preprocessor) {
			th.Errorf("status preprocessed despite no to")
		}

		return
	}

	if s.Switches[useTypeOutput] {
		th.ExpectStringHasKeywords("output preprocessor", out,
			"preprocessor", preprocessor)
	}

	if s.Switches[useTypeStatus] {
		th.ExpectStringHasKeywords("status preprocessor", status,
			"preprocessor", preprocessor)
	}
}

func (s *testScenario) assertTimestamp(th *thelper.THelper,
	log *Logger,
	outputLabel string,
	statusLabel string,
	format string) {
	out := log.outputWriters[outputLabel].(*bytes.Buffer).String()
	status := log.statusWriters[statusLabel].(*bytes.Buffer).String()

	if !s.Switches[useCustomTimestamp] {
		if s.Switches[useTypeOutput] &&
			strings.Contains(out, format) {
			th.Errorf("output custom timestamped despite no to")
		}

		if s.Switches[useTypeStatus] &&
			strings.Contains(status, format) {
			th.Errorf("status custom timestamped despite no to")
		}

		return
	}

	if s.Switches[useTypeOutput] {
		th.ExpectStringHasKeywords("output timestamp", out,
			"timestamp", format)
	}

	if s.Switches[useTypeStatus] {
		th.ExpectStringHasKeywords("status timestamp", status,
			"timestamp", format)
	}
}

func (s *testScenario) assertAdd(th *thelper.THelper,
	log *Logger, status StatusType, label string) {
	var ok, exist bool

	if !s.Switches[expectError] {
		exist = true
	}

	switch {
	case status == TYPE_STATUS:
		_, ok = log.outputWriters[label]
		if ok {
			th.Errorf("writer registered into outputWriters")
		}

		_, ok = log.statusWriters[label]
		if !ok && exist {
			th.Errorf("writer did not register into statusWriters")
		} else if ok && !exist {
			th.Errorf("writer did register into statusWriters")
		}
	case status == TYPE_OUTPUT:
		_, ok = log.statusWriters[label]
		if ok {
			th.Errorf("writer registered into statusWriters")
		}

		_, ok = log.outputWriters[label]
		if !ok && exist {
			th.Errorf("writer did not register into outputWriters")
		} else if ok && !exist {
			th.Errorf("writer did register into outputWriters")
		}
	default:
		th.Errorf("unknown StatusType for this test")
	}
}

func (s *testScenario) assertRemove(th *thelper.THelper,
	log *Logger, status StatusType, label string) {
	var ok, exist bool

	if s.Switches[expectRetain] {
		exist = true
	}

	switch {
	case status == TYPE_STATUS:
		_, ok = log.outputWriters[label]
		if !ok {
			th.Errorf("writer deleted from the wrong outputWriters")
		}

		_, ok = log.statusWriters[label]
		if ok && !exist {
			th.Errorf("writer retained in statusWriters")
		} else if !ok && exist {
			th.Errorf("writer missing in statusWriters")
		}
	case status == TYPE_OUTPUT:
		_, ok = log.statusWriters[label]
		if !ok {
			th.Errorf("writer deleted from the wrong statusWriters")
		}

		_, ok = log.outputWriters[label]
		if ok && !exist {
			th.Errorf("writer retained in outputWriters")
		} else if !ok && exist {
			th.Errorf("writer missing in outputWriters")
		}
	default:
		th.Errorf("unknown StatusType for this test")
	}
}

func (s *testScenario) assertCreateFile(th *thelper.THelper, f *os.File) {
	var exist bool

	if !s.Switches[expectError] {
		exist = true
	}

	switch {
	case s.Switches[useProperFilepath]:
		if f == nil && exist {
			th.Errorf("failed to create file")
		}
	case s.Switches[useEmptyFilepath]:
		if f != nil && !exist {
			th.Errorf("unknown file is created")
		}
	case s.Switches[useDirectoryFilepath]:
		if f != nil && !exist {
			th.Errorf("unknown file is created with directory path")
		}
	}
}

func (s *testScenario) expectError() bool {
	return s.Switches[expectError]
}

func (s *testScenario) generateStatusType() StatusType {
	switch {
	case s.Switches[useTypeOutput]:
		return TYPE_OUTPUT
	case s.Switches[useTypeStatus]:
		fallthrough
	default:
		return TYPE_STATUS
	}
}

func (s *testScenario) generateTag() string {
	switch {
	case s.Switches[useEmptyLevel]:
		return emptyLevel
	case s.Switches[useProperLevel]:
		fallthrough
	default:
		return properLevel
	}
}

func (s *testScenario) generateMessageFormat() string {
	switch {
	case s.Switches[useMessageFormat]:
		return argumentFormat
	case s.Switches[useEmptyFormat]:
		return emptyFormat
	case s.Switches[useStatementFormat]:
		fallthrough
	default:
		return statementFormat
	}
}

func (s *testScenario) generateMessageArguments() []interface{} {
	switch {
	case s.Switches[provideNilArgument]:
		return nil
	case s.Switches[provideEmptyArgument]:
		return []interface{}{""}
	case s.Switches[provideArgument]:
		fallthrough
	default:
		return []interface{}{argument}
	}
}

func (s *testScenario) generateWriter() *bytes.Buffer {
	return &bytes.Buffer{}
}

func (s *testScenario) generateLabel() (out, status string) {
	switch {
	case s.Switches[useEmptyLabel]:
		return "", ""
	case s.Switches[useProperLabel]:
		fallthrough
	default:
		return outLabel, statusLabel
	}
}

func (s *testScenario) generatePreProcessor() func(string) string {
	if !s.Switches[usePreProcessing] {
		return nil
	}

	return func(s string) string {
		return strings.TrimSuffix(s, "\n") + " " + preprocessor + "\n"
	}
}

func (s *testScenario) generateTimestamp() string {
	if !s.Switches[useCustomTimestamp] {
		return ISO8601
	}

	return customTimestamp
}

func (s *testScenario) generateFilepath() string {
	switch {
	case s.Switches[useEmptyFilepath]:
		return ""
	case s.Switches[useDirectoryFilepath]:
		return directoryFilepath
	case s.Switches[useProperFilepath]:
		fallthrough
	default:
	}

	return properFilepath
}

func (s *testScenario) cleanUp() {
	_ = os.RemoveAll(properFilepath)
}
