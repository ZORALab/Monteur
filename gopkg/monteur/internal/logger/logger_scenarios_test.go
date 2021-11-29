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

func getTestScenarios() []testScenario {
	return []testScenario{
		{
			UID:      1,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      2,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      3,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      4,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      5,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      6,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      7,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      8,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      9,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      10,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      11,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      12,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      13,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      14,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      15,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      16,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      17,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      18,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      19,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      20,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      21,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      22,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      23,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      24,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      25,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      26,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      27,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      28,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      29,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      30,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      31,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      32,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      33,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      34,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      35,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      36,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      37,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      38,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      39,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      40,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      41,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      42,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      43,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      44,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      45,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      46,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      47,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      48,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      49,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      50,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      51,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      52,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      53,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      54,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      55,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      56,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      57,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      58,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      59,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      60,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      61,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      62,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      63,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      64,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      65,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      66,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      67,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      68,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      69,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      70,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      71,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      72,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      73,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      74,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      75,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      76,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      77,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      78,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      79,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      80,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      81,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      82,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      83,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      84,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      85,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      86,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      87,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      88,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideArgument:    true,
			},
		}, {
			UID:      89,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      90,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      91,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      92,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      93,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      94,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      95,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      96,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideArgument:    true,
			},
		}, {
			UID:      97,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      98,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      99,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      100,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      101,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      102,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      103,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      104,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useMessageFormat:     true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      105,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      106,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      107,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      108,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      109,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      110,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      111,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      112,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useEmptyFormat:       true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      113,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      114,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   false,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      115,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      116,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   false,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      117,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      118,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     false,
				useCustomTimestamp:   true,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      119,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useEmptyLevel:        true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      120,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. empty arguments are given.
`,
			Switches: map[string]bool{
				useTypeOutput:        true,
				useProperLevel:       true,
				usePreProcessing:     true,
				useCustomTimestamp:   true,
				useStatementFormat:   true,
				provideEmptyArgument: true,
			},
		}, {
			UID:      121,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      122,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      123,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      124,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      125,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      126,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      127,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      128,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideNilArgument: true,
			},
		}, {
			UID:      129,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      130,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      131,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      132,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      133,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      134,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      135,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      136,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. empty format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useEmptyFormat:     true,
				provideNilArgument: true,
			},
		}, {
			UID:      137,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      138,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      139,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      140,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      141,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      142,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      143,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. empty level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useEmptyLevel:      true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      144,
			TestType: testLogf,
			Description: `
logger.Logf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. statement format is given.
6. nil argument is given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useStatementFormat: true,
				provideNilArgument: true,
			},
		}, {
			UID:      145,
			TestType: testLogln,
			Description: `
logger.Logln should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				provideArgument:    true,
			},
		}, {
			UID:      146,
			TestType: testLogln,
			Description: `
logger.Logln should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				provideArgument:    true,
			},
		}, {
			UID:      147,
			TestType: testLogln,
			Description: `
logger.Logln should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				provideArgument:    true,
			},
		}, {
			UID:      148,
			TestType: testLogln,
			Description: `
logger.Logln should work properly when:
1. TYPE_STATUS is selected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				provideArgument:    true,
			},
		}, {
			UID:      149,
			TestType: testLogln,
			Description: `
logger.Logln should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				provideArgument:    true,
			},
		}, {
			UID:      150,
			TestType: testLogln,
			Description: `
logger.Logln should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				provideArgument:    true,
			},
		}, {
			UID:      151,
			TestType: testLogln,
			Description: `
logger.Logln should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				provideArgument:    true,
			},
		}, {
			UID:      152,
			TestType: testLogln,
			Description: `
logger.Logln should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeOutput:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				provideArgument:    true,
			},
		}, {
			UID:      153,
			TestType: testPrintln,
			Description: `
logger.Println should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				provideArgument:    true,
			},
		}, {
			UID:      154,
			TestType: testPrintln,
			Description: `
logger.Println should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				provideArgument:    true,
			},
		}, {
			UID:      155,
			TestType: testPrintln,
			Description: `
logger.Println should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				provideArgument:    true,
			},
		}, {
			UID:      156,
			TestType: testPrintln,
			Description: `
logger.Println should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				provideArgument:    true,
			},
		}, {
			UID:      157,
			TestType: testPrintln,
			Description: `
logger.Println should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				provideArgument:    true,
			},
		}, {
			UID:      158,
			TestType: testPrintln,
			Description: `
logger.Println should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				provideArgument:    true,
			},
		}, {
			UID:      159,
			TestType: testPrintln,
			Description: `
logger.Println should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				provideArgument:    true,
			},
		}, {
			UID:      160,
			TestType: testPrintln,
			Description: `
logger.Println should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				provideArgument:    true,
			},
		}, {
			UID:      161,
			TestType: testPrintf,
			Description: `
logger.Printf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      162,
			TestType: testPrintf,
			Description: `
logger.Printf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments are given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      163,
			TestType: testPrintf,
			Description: `
logger.Printf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      164,
			TestType: testPrintf,
			Description: `
logger.Printf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      165,
			TestType: testPrintf,
			Description: `
logger.Printf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      166,
			TestType: testPrintf,
			Description: `
logger.Printf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. no custom timestamp is given.
5. proper format is given.
6. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: false,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      167,
			TestType: testPrintf,
			Description: `
logger.Printf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. no preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   false,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      168,
			TestType: testPrintf,
			Description: `
logger.Printf should work properly when:
1. TYPE_STATUS is expected.
2. proper level tag is given.
3. preprocessing is given.
4. custom timestamp is given.
5. proper format is given.
6. proper arguments given.
`,
			Switches: map[string]bool{
				useTypeStatus:      true,
				useProperLevel:     true,
				usePreProcessing:   true,
				useCustomTimestamp: true,
				useMessageFormat:   true,
				provideArgument:    true,
			},
		}, {
			UID:      169,
			TestType: testAdd,
			Description: `
logger.Add should work properly when:
1. TYPE_STATUS is selected.
2. proper label is given.
3. not expecting error.
`,
			Switches: map[string]bool{
				useTypeStatus:  true,
				useProperLabel: true,
			},
		}, {
			UID:      170,
			TestType: testAdd,
			Description: `
logger.Add should work properly when:
1. TYPE_STATUS is selected.
2. empty label is given.
3. expecting error.
`,
			Switches: map[string]bool{
				useTypeStatus: true,
				useEmptyLabel: true,
				expectError:   true,
			},
		}, {
			UID:      171,
			TestType: testAdd,
			Description: `
logger.Add should work properly when:
1. TYPE_OUTPUT is selected.
2. proper label is given.
3. not expecting error.
`,
			Switches: map[string]bool{
				useTypeOutput:  true,
				useProperLabel: true,
				expectError:    false,
			},
		}, {
			UID:      172,
			TestType: testAdd,
			Description: `
logger.Add should work properly when:
1. TYPE_OUTPUT is selected.
2. empty label is given.
3. expecting error.
`,
			Switches: map[string]bool{
				useTypeOutput: true,
				useEmptyLabel: true,
				expectError:   true,
			},
		}, {
			UID:      173,
			TestType: testRemove,
			Description: `
logger.Remove should work properly when:
1. TYPE_STATUS is selected.
2. proper label is given.
3. not expecting retain.
`,
			Switches: map[string]bool{
				useTypeStatus:  true,
				useProperLabel: true,
				expectRetain:   false,
			},
		}, {
			UID:      174,
			TestType: testRemove,
			Description: `
logger.Remove should work properly when:
1. TYPE_STATUS is selected.
2. empty label is given.
3. expecting retain.
`,
			Switches: map[string]bool{
				useTypeStatus: true,
				useEmptyLabel: true,
				expectRetain:  true,
			},
		}, {
			UID:      175,
			TestType: testRemove,
			Description: `
logger.Remove should work properly when:
1. TYPE_OUTPUT is selected.
2. proper label is given.
3. not expecting retain.
`,
			Switches: map[string]bool{
				useTypeOutput:  true,
				useProperLabel: true,
				expectRetain:   false,
			},
		}, {
			UID:      176,
			TestType: testRemove,
			Description: `
logger.Remove should work properly when:
1. TYPE_OUTPUT is selected.
2. empty label is given.
3. expecting retain.
`,
			Switches: map[string]bool{
				useTypeOutput: true,
				useEmptyLabel: true,
				expectRetain:  true,
			},
		}, {
			UID:      177,
			TestType: testCreateFile,
			Description: `
CreateFile should work properly when:
1. proper filepath is given.
2. not expecting error.
`,
			Switches: map[string]bool{
				useProperFilepath: true,
				expectError:       false,
			},
		}, {
			UID:      178,
			TestType: testCreateFile,
			Description: `
CreateFile should work properly when:
1. empty filepath is given.
2. expecting error.
`,
			Switches: map[string]bool{
				useEmptyFilepath: true,
				expectError:      true,
			},
		}, {
			UID:      179,
			TestType: testCreateFile,
			Description: `
CreateFile should work properly when:
1. directory filepath is given.
2. expecting error.
`,
			Switches: map[string]bool{
				useDirectoryFilepath: true,
				expectError:          true,
			},
		}, {
			UID:      180,
			TestType: testCreateStderr,
			Description: `
CreateStderr should work properly and:
1. file is expected.
`,
			Switches: map[string]bool{
				useProperFilepath: true,
			},
		}, {
			UID:      181,
			TestType: testCreateStdout,
			Description: `
CreateStdout should work properly and:
1. file is expected.
`,
			Switches: map[string]bool{
				useProperFilepath: true,
			},
		},
	}
}
