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

func TestControlString(t *testing.T) {
	var panicError interface{}

	for i, s := range getTestScenarios() {
		if s.TestType != testControlString {
			continue
		}

		// prepare
		th := s.prepareTHelper(t)
		panicError = noPanicFlag
		str := ""

		entity := s.createEntity()
		packages := s.createRelPackages()
		version := s.createVersion()
		website := s.createURL()
		description := s.createDescription()
		vcs := s.createVCS()
		testsuite := s.createTestsuite()
		name := s.createAppName()
		section := s.createSection()
		standards := s.createStandardsVersion()
		packageType := s.createPackageType()
		rrr := s.createRulesRequiresRoot()
		priority := s.createPriority()
		arch := s.createArchitecture()
		uploaders := s.createUploaders()
		essentials := s.createEssential()
		buildSource := s.createBuildSource()

		// test
		subject := &Control{
			Maintainer:        entity,
			Packages:          packages,
			Version:           version,
			Homepage:          website,
			Description:       description,
			VCS:               vcs,
			Testsuite:         testsuite,
			Name:              name,
			Section:           section,
			StandardsVersion:  standards,
			PackageType:       packageType,
			RulesRequiresRoot: rrr,
			Priority:          priority,
			Architecture:      arch,
			Uploaders:         uploaders,
			Essential:         essentials,
			BuildSource:       buildSource,
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
		s.assertControl(th, str)
		s.assertPanic(th, panicError)
		s.log(th, map[string]interface{}{
			"given maintainer":          entity,
			"given relational packages": packages,
			"given version":             version,
			"given homepage":            website,
			"given description":         description,
			"given vcs":                 vcs,
			"given testsuite":           testsuite,
			"given name":                name,
			"given section":             section,
			"given standards version":   standards,
			"given package type":        packageType,
			"given rules requires root": rrr,
			"given priority":            priority,
			"given architecture":        arch,
			"given uploaders":           uploaders,
			"given essential flag":      essentials,
			"given build source flag":   buildSource,
			"created subject":           subject,
			"got":                       str,
			"panic error":               panicError,
		})
		th.Logf(`TESTNOTE: DEBIAN/control can only be fully tested
visually
═══════ START ══════
%s
═══════  END  ══════`, str)
		th.Conclude()
	}
}
