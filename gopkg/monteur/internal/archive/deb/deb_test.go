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
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"gitlab.com/zoralab/cerigo/testing/thelper"
)

const (
	testChangelogParseForErrorPanic = "testChangelogParseForErrorPanic"
	testChangelogString             = "testChangelogString"
	testControlString               = "testControlString"
	testCopyrightString             = "testCopyrightString"
	testDataGenerate                = "testDataGenerate"
	testDataSanitize                = "testDataSanitize"
	testDescriptionString           = "testDescriptionString"
	testEntityString                = "testEntityString"
	testLicenseString               = "testLicenseString"
	testPackageListString           = "testPackageListString"
	testPackageMetaString           = "testPackageMetaString"
	testSourceSanitize              = "testSourceSanitize"
	testTestsuiteString             = "testTestsuiteString"
	testVCSString                   = "testVCSString"
	testVersionString               = "testVersionString"
)

//nolint:lll
const (
	expectError = "expectError"
	expectPanic = "expectPanic"

	simulateStatFxError = "simulateStatFxError"

	createBadHeaderChangelogFile                = "createBadHeaderChangelogFile"
	createBadNameChangelogFile                  = "createBadNameChangelogFile"
	createBadEmailChangelogFile                 = "createBadEmailChangelogFile"
	createBadEmailDelimiterChangelogFile        = "createBadEmailDelimiterChangelogFile"
	createBadTimestampChangelogFile             = "createBadTimestampChangelogFile"
	createBadVersionHeaderChangelogFile         = "createBadVersionHeaderChangelogFile"
	createBadVersionEpochHeaderChangelogFile    = "createBadVersionEpochHeaderChangelogFile"
	createBadVersionUpstreamHeaderChangelogFile = "createBadVersionUpstreamHeaderChangelogFile"
	createChangesLeadChangelogFile              = "createChangesLeadChangelogFile"
	createDoubleHeaderChangelogFile             = "createDoubleHeaderChangelogFile"
	createSigHeaderChangelogFile                = "createSigHeaderChangelogFile"
	createSigLeadChangelogFile                  = "createSigLeadChangelogFile"
	createMissingDistroHeaderChangelogFile      = "createMissingDistroChangelogFile"
	createNoChangelogFile                       = "createNoChangelogFile"
	createBadPackageHeaderChangelogFile         = "createBadPackageHeaderChangelogFile"
	createTailingHeaderChangelogFile            = "createTailingHeaderChangelogFile"
	createTailingChangesChangelogFile           = "createTailingChangesChangelogFile"
	createTailingSigChangelogFile               = "createTailingSigChangelogFile"

	use3p7p0Install                  = "use3p7p0Install"
	use4p0p0Install                  = "use4p0p0Install"
	use4p6p0Install                  = "use4p6p0Install"
	useAlphaLedVersionRevision       = "useAlphaLedVersionRevision"
	useAlphaLedVersionUpstream       = "useAlphaLedVersionUpstream"
	useAnyArch                       = "useAnyArch"
	useBuildSourceFlag               = "useBuildSourceFlag"
	useChangelogUrgencyCritical      = "useChangelogUrgencyCritical"
	useChangelogUrgencyEmergency     = "useChangelogUrgencyEmergency"
	useChangelogUrgencyHigh          = "useChangelogUrgencyHigh"
	useChangelogUrgencyLow           = "useChangelogUrgencyLow"
	useChangelogUrgencyMedium        = "useChangelogUrgencyMedium"
	useChangelogUrgencyUnknown       = "useChangelogUrgencyUnknown"
	useCompletedChangelogParseStatus = "useCompletedChangelogParseStatus"
	useCopyrightFormat1p0            = "useCopyrightFormat1p0"
	useDashVersionRevision           = "useDashVersionRevision"
	useDashVersionUpstream           = "useDashVersionUpstream"
	useEmptyArchList                 = "useEmptyArchList"
	useEmptyBrowser                  = "useEmptyBrowser"
	useEmptyChangelogChanges         = "useEmptyChangelogChanges"
	useEmptyDescription              = "useEmptyDescription"
	useEmptyDistro                   = "useEmptyDistro"
	useEmptyEntities                 = "useEmptyEntities"
	useEmptyInstall                  = "useEmptyInstall"
	useEmptyLicenses                 = "useEmptyLicenses"
	useEmptyManpage                  = "useEmptyManpage"
	useEmptyPackageList              = "useEmptyPackageList"
	useEmptyPackageListing           = "useEmptyPackageListing"
	useEmptyPath                     = "useEmptyPath"
	useEmptyScripts                  = "useEmptyScripts"
	useEmptySection                  = "useEmptySection"
	useEmptyTestsuite                = "useEmptyTestsuite"
	useEmptyTimestamp                = "useEmptyTimestamp"
	useEmptyUploaders                = "useEmptyUploaders"
	useEmptyURL                      = "useEmptyURL"
	useEmptyVCS                      = "useEmptyVCS"
	useEssentialFlag                 = "useEssentialFlag"
	useFaultyDescription             = "useFaultyDescription"
	useFaultyPackageList             = "useFaultyPackageList"
	useFaultyScripts                 = "useFaultyScripts"
	useFaultySection                 = "useFaultySection"
	useFaultyStandardsVersion        = "useFaultyStandardsVersion"
	useFaultyTestedDataControl       = "useFaultyTestedDataControl"
	useFaultyTestedDataCopyright     = "useFaultyTestedDataCopyright"
	useFaultyTestedDataChangelog     = "useFaultyTestedDataChangelog"
	useFaultyTestedDataSource        = "useFaultyTestedDataSource"
	useFaultyTestsuite               = "useFaultyTestsuite"
	useFaultyUploaders               = "useFaultyUploaders"
	useFaultyVCS                     = "useFaultyVCS"
	useFaultyWorkDir                 = "useFaultyWorkDir"
	useGhostPath                     = "useGhostPath"
	useIllegalVersionUpstream        = "useIllegalVersionUpstream"
	useIllegalVersionRevision        = "useIllegalVersionRevision"
	useLicenseTagAnd                 = "useLicenseTagAnd"
	useLicenseTagOr                  = "useLicenseTagOr"
	useLicenseTagSymbol              = "useLicenseTagSymbol"
	useLicenseWithNilFiles           = "useLicenseWithNilFiles"
	useLicenseWithoutBody            = "useLicenseWithoutBody"
	useLicenseWithoutComment         = "useLicenseWithoutComment"
	useLicenseWithoutFiles           = "useLicenseWithoutFiles"
	useLicenseWithoutTag             = "useLicenseWithoutTag"
	useLongSynopsis                  = "useLongSynopsis"
	useNilChangelog                  = "useNilChangelog"
	useNilContentPackageList         = "useNilContentPackageList"
	useNilControl                    = "useNilControl"
	useNilCopyright                  = "useNilCopyright"
	useNilPackageList                = "useNilPackageList"
	useNilSource                     = "useNilSource"
	useNilVersion                    = "useNilVersion"
	useNonHTTPURL                    = "useNonHTTPURL"
	usePackageTypeDEB                = "usePackageTypeDEB"
	usePackageTypeUDEB               = "usePackageTypeUDEB"
	usePackageTypeUnknown            = "usePackageTypeUnknown"
	usePkgListBuildDepends           = "usePkgListBuildDepends"
	usePkgListBuildDependsIndep      = "usePkgListBuildDependsIndep"
	usePkgListBuildDependsArch       = "usePkgListBuildDependesArch"
	usePkgListBuildConflicts         = "usePkgListBuildConflicts"
	usePkgListBuildConflictsIndep    = "usePkgListBuildConflictsIndep"
	usePkgListBuildConflictsArch     = "usePkgListBuildConflictsArch"
	usePkgListBuiltUsing             = "usePkgListBuiltUsing"
	usePkgListPreDepends             = "usePkgListPreDepends"
	usePkgListDepends                = "usePkgListDepends"
	usePkgListRecommends             = "usePkgListRecommends"
	usePkgListSuggests               = "usePkgListSuggests"
	usePkgListEnhances               = "usePkgListEnhances"
	usePkgListBreaks                 = "usePkgListBreaks"
	usePkgListConflicts              = "usePkgListConflicts"
	usePkgListProvides               = "usePkgListProvides"
	usePkgListReplaces               = "usePkgListReplaces"
	usePkgListUnknown                = "usePkgListUnknown"
	usePriorityImportant             = "usePriorityImportant"
	usePriorityOptional              = "usePriorityOptional"
	usePriorityRequired              = "usePriorityRequired"
	usePriorityStandard              = "usePriorityStandard"
	usePriorityUnknown               = "usePriorityUnknown"
	useProperAppName                 = "useProperAppName"
	useProperArch                    = "useProperArch"
	useProperArchList                = "useProperArchList"
	useProperBranch                  = "useProperBranch"
	useProperBrowser                 = "useProperBrowser"
	useProperChangelogChanges        = "useProperChangelogChanges"
	useProperCompat                  = "useProperCompat"
	useProperCopyrightDisclaimer     = "useProperCopyrightDisclaimer"
	useProperCopyrightSource         = "useProperCopyrightSource"
	useProperDescription             = "useProperDescription"
	useProperDescriptionText         = "useProperDescriptionText"
	useProperDistro                  = "useProperDistro"
	useProperEntities                = "useProperEntities"
	useProperEntity                  = "useProperEntity"
	useProperEmail                   = "useProperEmail"
	useProperLicenses                = "useProperLicenses"
	useProperManpage                 = "useProperManpage"
	useProperName                    = "useProperName"
	useProperPackageList             = "useProperPackageList"
	useProperPackageListing          = "useProperPackageListing"
	useProperPath                    = "useProperPath"
	useProperPathDirectory           = "useProperPathDirectory"
	useProperInstall                 = "useProperInstall"
	useProperRules                   = "useProperRules"
	useProperScripts                 = "useProperScript"
	useProperSection                 = "useProperSection"
	useProperSourceLocalOptions      = "useProperSourceLocalOptions"
	useProperSourceOptions           = "useProperSourceOptions"
	useProperStandardsVersion        = "useProperStandardsVersion"
	useProperSynopsis                = "useProperSynopsis"
	useProperTestsuite               = "useProperTestsuite"
	useProperTimestamp               = "useProperTimestmap"
	useProperUploaders               = "useProperUploaders"
	useProperURL                     = "useProperURL"
	useProperVCS                     = "useProperVCS"
	useProperVersionEpoch            = "useProperVersionEpoch"
	useProperVersionRevision         = "useProperVersionRevision"
	useProperVersionUpstream         = "useProperVersionUpstream"
	useProperWorkDir                 = "useProperWorkDir"
	useProperYear                    = "useProperYear"
	useSourceFormatNative3p0         = "useSourceFormatNative3p0"
	useSourceFormatQuilt3p0          = "useSourceFormatQuilt3p0"
	useRRRBinTarget                  = "useRRRBinTarget"
	useRRRCustom                     = "useRRRCustom"
	useRRRNo                         = "useRRRNo"
	useRRRUnknown                    = "useRRRUnknown"
	useUnknownArch                   = "useUnknownArch"
	useUnknownChangelogParseStatus   = "useUnknownChangelogParseStatus"
	useVCSArch                       = "useVCSArch"
	useVCSBazaar                     = "useVCSBazaar"
	useVCSCVS                        = "useVCSCVS"
	useVCSDarcs                      = "useVCSDarcs"
	useVCSGit                        = "useVCSGit"
	useVCSMercurial                  = "useVCSMercurial"
	useVCSMonotone                   = "useVCSMonotone"
	useVCSSubversion                 = "useVCSSubversion"
	useVCSUnknown                    = "useVCSUnknown"
	useVERControlSTER                = "useVERControlSTER"
	useVERControlEAEQ                = "useVERControlEAEQ"
	useVERControlEXEQ                = "useVERControlEXEQ"
	useVERControlLAEQ                = "useVERControlLAEQ"
	useVERControlSTLA                = "useVERControlSTLA"
	useVERControlUnknown             = "useVERControlUnknown"
	useBadYear                       = "useBadYear"
)

//nolint:lll
const (
	app               = "TestApp"
	appControl        = "ControlApp"
	changelogA        = "changed feature A"
	changelogB        = "changed feature B"
	compat            = 9
	customRRR         = "namespace/case1"
	branch            = "next"
	distroDebian      = "debian"
	distroUbuntu      = "ubuntu"
	epoch             = 5
	email             = "john.smith@testing.email"
	emailControl      = "aoi.fujimura@testing.corp"
	emailControl2     = "nasuki.aoi@testing.corp"
	ghostPath         = "p/package"
	installPath       = "usr/local/testapp/program"
	installPath2      = "usr/bin/program"
	installProgram    = "testapp"
	keySynopsis       = "sys"
	keyDescription    = "body"
	name              = "John Smith"
	nameControl       = "Aoi Fujimura"
	nameControl2      = "Nasuki Aoi"
	nonHTTPURL        = "file:///home/u0/Documents/testfile"
	path              = "testsuite/changelog"
	testsuitePath     = "testsuite/.gitkeep"
	directory         = "testsuite"
	arch              = "amd64"
	archANY           = "any"
	properHTTPURL     = "https://www.example.com/path/to/dir?query=language"
	restrictedWorkDir = "./restricted"
	revision          = "0.50.0~upstream"
	revisionDash      = "-dbg"
	revisionIllegal   = "my/VersionX.51"
	revisionPrefix    = "u"
	rulesFile         = `#!/usr/bin/make -f

# Uncomment this to turn on verbose mode.
#export DH_VERBOSE=1

%:
        dh $@`
	section     = "contrib/devel"
	shellScript = `#!/bin/sh
1>&2 printf "shell script executed"
`
	sourceLocalOptions = `unapply-patches
abort-on-upstream-changes`
	sourceOptions   = `extend-diff-ignore = "(^|/)(config\.sub|config\.guess|Makefile)$`
	timestamp       = "Tue, 07 Dec 2021 16:08:49 +0000"
	timestampZero   = "Mon, 01 Jan 0001 00:00:00 +0000"
	upstream        = "0.50.0~testapp"
	upstreamDash    = "-alpha"
	upstreamIllegal = "my/VersionX.51"
	upstreamPrefix  = "v"
	unknown         = "unknown"
	urgencyUnknown  = "unknownUrgency"
	workDir         = "./workdir/testapp"
	year            = 2021
	yearControl     = 2020
	yearControl2    = 2019

	noPanicFlag = "panic did not occur"

	licenseBody       = licenseApache2Body
	licenseDisclaimer = "Use at your own risk"
	licenseComment    = "For testing rendering purpose only"
	licenseFile1      = "*.go"
	licenseFile2      = "*.pdf"
	licenseTag        = "Apache-2"
)

type testScenario thelper.Scenario

func (s *testScenario) prepareTHelper(t *testing.T) *thelper.THelper {
	return thelper.NewTHelper(t)
}

func (s *testScenario) log(th *thelper.THelper, data map[string]interface{}) {
	th.LogScenario(thelper.Scenario(*s), data)
}

func (s *testScenario) stringUID() string {
	return strconv.Itoa(s.UID)
}

func (s *testScenario) assertChangelogParseForErrorPanic(th *thelper.THelper,
	err error) {
	switch {
	case s.Switches[useCompletedChangelogParseStatus]:
		if err == nil {
			th.Errorf("error is expected but was not raised")
		}

		return
	default:
		if err != nil {
			th.Errorf("unexpected error raised")
		}

		return
	}
}

func (s *testScenario) assertDataGenerate(th *thelper.THelper, err error) {
	switch {
	case s.Switches[useFaultyTestedDataControl],
		s.Switches[useNilControl],
		s.Switches[useFaultyTestedDataCopyright],
		s.Switches[useNilCopyright],
		s.Switches[useFaultyTestedDataChangelog],
		s.Switches[useNilChangelog],
		s.Switches[useFaultyTestedDataSource],
		s.Switches[useNilSource],
		!s.Switches[useProperManpage],
		s.Switches[useEmptyManpage],
		!s.Switches[useProperScripts],
		!s.Switches[useProperRules],
		!s.Switches[useProperCompat],
		s.Switches[useEmptyInstall],
		s.Switches[use3p7p0Install],
		s.Switches[use4p6p0Install],
		s.Switches[use4p0p0Install]:
		if err == nil {
			th.Errorf("expected error but none was raised")
		}

		return
	case !s.Switches[useProperInstall]:
		if err == nil {
			th.Errorf("expected error but none was raised")
		}

		return
	case s.Switches[expectError]:
		if err == nil {
			th.Errorf("expected error but none was raised")
		}

		return
	}

	if err != nil {
		th.Errorf("unexpected error was raised")
	}
}

func (s *testScenario) logFiles(th *thelper.THelper, err error) {
	if err != nil {
		th.Logf("test has error, skipping files logging.")
		return
	}

	// build filepath
	workPath := filepath.Join(workDir, "DEBIAN")
	if s.Switches[useBuildSourceFlag] {
		workPath = filepath.Join(workDir, "debian")
	}

	s._assertFile(th, filepath.Join(workPath, "control"))
	s._assertFile(th, filepath.Join(workPath, "copyright"))
	s._assertFile(th, filepath.Join(workPath, "changelog"))
	s._assertFile(th, path)
	s._assertFile(th, filepath.Join(workPath, "source", "format"))
	if s.Switches[useProperSourceLocalOptions] {
		s._assertFile(th, filepath.Join(workPath,
			"source",
			"local-options"),
		)
	}
	if s.Switches[useProperSourceOptions] {
		s._assertFile(th, filepath.Join(workPath, "source", "options"))
	}

	if s.Switches[useProperManpage] {
		s._assertFile(th, filepath.Join(workPath, app+".manpages"))
		s._assertFile(th, filepath.Join(workPath, app+"."+manPageTag1))
	}

	if s.Switches[useProperScripts] {
		s._assertFile(th, filepath.Join(workPath,
			string(SHELL_PRE_INSTALL)))
		s._assertFile(th, filepath.Join(workPath,
			string(SHELL_POST_INSTALL)))
		s._assertFile(th, filepath.Join(workPath,
			string(SHELL_PRE_REMOVE)))
		s._assertFile(th, filepath.Join(workPath,
			string(SHELL_POST_REMOVE)))
	}

	s._assertFile(th, filepath.Join(workPath, "rules"))
	s._assertFile(th, filepath.Join(workPath, "compat"))
	s._assertFile(th, filepath.Join(workPath, "install"))
}

func (s *testScenario) _assertFile(th *thelper.THelper, path string) {
	var out string
	var err error
	var scanner *bufio.Scanner
	var f *os.File

	f, err = os.OpenFile(path, os.O_RDONLY, _PERMISSION_FILE)
	if err != nil {
		th.Errorf("failed to open asserted file: '%s'", path)
		return
	}

	scanner = bufio.NewScanner(f)
	out = ""
	isFirst := true
	for scanner.Scan() {
		if !isFirst {
			out += "\n"
		}

		out += scanner.Text()
		isFirst = false
	}

	f.Close()

	th.Logf("File '%s' HAS:\n╔═══ BEGIN ═══╗\n%s\n╚═══  END  ═══╝\n",
		path,
		out,
	)
}

func (s *testScenario) assertSource(th *thelper.THelper, err error) {
	switch {
	case s.Switches[expectError]:
		if err == nil {
			th.Errorf("expected error is not raised.")
		}
	default:
		if err != nil {
			th.Errorf("unexpected error was raised.")
		}
	}
}

//nolint:lll,gocognit
func (s *testScenario) assertControl(th *thelper.THelper, str string) {
	// assert all empty output and panics
	switch {
	case !s.Switches[useProperName],
		!s.Switches[useProperEmail],
		!s.Switches[useProperEntity]:
		if str != "" {
			th.Errorf("control generated despite faulty maintainer")
		}

		return
	case s.Switches[useFaultyPackageList]:
		if str != "" {
			th.Errorf("control generated despite faulty package list")
		}

		return
	case !s.Switches[useProperVersionUpstream],
		s.Switches[useNilVersion]:
		if str != "" {
			th.Errorf("control generated despite faulty version")
		}

		return
	case s.Switches[useFaultyDescription],
		s.Switches[useEmptyDescription],
		!s.Switches[useProperDescription]:
		if str != "" {
			th.Errorf("control generated despite faulty description")
		}

		return
	case !s.Switches[useProperAppName]:
		if str != "" {
			th.Errorf("control generated despite faulty app name")
		}

		return
	case s.Switches[useFaultyStandardsVersion],
		!s.Switches[useProperStandardsVersion]:
		if str != "" {
			th.Errorf("control generated despite faulty std version")
		}

		return
	case s.Switches[useFaultyVCS]:
		if str != "" {
			th.Errorf("control generated despite faulty vcs")
		}

		return
	case s.Switches[useFaultyTestsuite]:
		if str != "" {
			th.Errorf("control generated despite faulty testsuite")
		}

		return
	case s.Switches[usePackageTypeUnknown]:
		if str != "" {
			th.Errorf("control generated despite faulty pkg type")
		}

		return
	case s.Switches[usePriorityUnknown]:
		if str != "" {
			th.Errorf("control generated despite priority unknown")
		}

		return
	case !s.Switches[useProperArch] && !s.Switches[useUnknownArch]:
		if str != "" {
			th.Errorf("control generated despite unknown arch")
		}

		return
	case s.Switches[useRRRUnknown]:
		if str != "" {
			th.Errorf("control generated despite unknown RRR")
		}

		return
	case s.Switches[useFaultyUploaders]:
		if str != "" {
			th.Errorf("control generated despite faulty uploaders")
		}

		return
	default:
	}

	if str == "" {
		th.Errorf("control not generated despite ok")
	}
}

//nolint:gocognit
func (s *testScenario) assertChangelog(th *thelper.THelper, str string) {
	// assert all empty output and panics
	switch {
	case !s.Switches[useProperName],
		!s.Switches[useProperEmail],
		!s.Switches[useProperEntity]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog entry generated despite bad Maintainer!")
		}

		return
	case !s.Switches[useProperAppName]:
		if str != "" {
			th.Errorf("Changelog entry generated despite bad Name!")
		}

		return
	case !s.Switches[useProperVersionUpstream], s.Switches[useNilVersion]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog entry generated despite bad Version!")
		}

		return
	case !s.Switches[useProperTimestamp]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog entry generated despite bad timestamp!")
		}

		return
	case s.Switches[useEmptyTimestamp]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog entry generated despite empty timestamp!")
		}

		return
	case !s.Switches[useProperDistro]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog entry generated despite bad distro!")
		}

		return
	case !s.Switches[useProperChangelogChanges]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog entry generated despite bad changes!")
		}

		return
	case s.Switches[useEmptyChangelogChanges]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog entry generated despite empty Version!")
		}

		return
	case s.Switches[useChangelogUrgencyUnknown]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog entry generated despite unknown urgency!")
		}

		return
	case s.Switches[useProperPathDirectory]:
		if str != "" {
			th.Errorf("Changelog generated despite directory path")
		}

		return
	case s.Switches[simulateStatFxError]:
		if str != "" {
			//nolint:lll
			th.Errorf("Changelog generated despite path state error")
		}

		return
	}

	// header line
	expect := app + _FIELD_CHANGELOG_DELIMIT_PACKAGE + upstream
	if s.Switches[useProperVersionRevision] {
		expect += "-" + revision
	}
	expect += _FIELD_CHANGELOG_DELIMIT_DISTRO + distroDebian + " " +
		distroUbuntu

	expect += _FIELD_CHANGELOG_DELIMIT_URGENCY
	switch {
	case s.Switches[useChangelogUrgencyMedium]:
		expect += string(CHANGELOG_URGENCY_MEDIUM)
	case s.Switches[useChangelogUrgencyHigh]:
		expect += string(CHANGELOG_URGENCY_HIGH)
	case s.Switches[useChangelogUrgencyEmergency]:
		expect += string(CHANGELOG_URGENCY_EMERGENCY)
	case s.Switches[useChangelogUrgencyCritical]:
		expect += string(CHANGELOG_URGENCY_CRITICAL)
	default:
		expect += string(CHANGELOG_URGENCY_LOW)
	}

	expect += "\n"

	// line spacing
	expect += "\n"

	// changelog contents
	expect += _FIELD_CHANGELOG_DELIMIT_CHANGE + changelogA + "\n"
	expect += _FIELD_CHANGELOG_DELIMIT_CHANGE + changelogB + "\n"

	// line spacing
	expect += "\n"

	// signed-off
	expect += _FIELD_CHANGELOG_DELIMIT_SIGNATURE + name
	expect += _FIELD_CHANGELOG_DELIMIT_EMAIL + email
	expect += ">" + _FIELD_CHANGELOG_DELIMIT_TIMESTAMP + timestamp

	// assert
	if expect != str {
		th.Errorf("Changelog entry unexpected. Expect: '%#v'", expect)
	}
}

//nolint:gocognit
func (s *testScenario) assertCopyright(th *thelper.THelper, str string) {
	switch {
	case !s.Switches[useProperName],
		!s.Switches[useProperEmail],
		!s.Switches[useProperEntity]:
		if str != "" {
			th.Errorf("Copyright created without upstream contact")
		}

		return
	case s.Switches[useProperYear]:
		if str == "" {
			th.Errorf("Copyright not created with Contact.Year")

			return
		}
	case !s.Switches[useProperEntities] &&
		s.Switches[useLicenseWithoutTag]:
		fallthrough
	case !s.Switches[useProperLicenses] &&
		s.Switches[useLicenseWithoutTag]:
		if str != "" {
			th.Errorf("Copyright created despite no licenses")
		}

		return
	case !(s.Switches[useCopyrightFormat1p0]):
		if str != "" {
			th.Errorf("Copyright created despite no format")
		}

		return
	case !(s.Switches[useProperAppName]):
		if str != "" {
			th.Errorf("Copyright created despite no upstream name")
		}

		return
	}

	// build format
	expect := _FIELD_FORMAT
	switch {
	case s.Switches[useCopyrightFormat1p0]:
		expect += string(COPYRIGHT_FORMAT_1_0) + "\n"
	default:
	}

	// build upstream name
	expect += _FIELD_UPSTREAM_NAME + app + "\n"

	// build upstream contact
	expect += _FIELD_UPSTREAM_CONTACT + name + " <" + email + ">\n"

	// build source
	if s.Switches[useProperCopyrightSource] {
		expect += _FIELD_SOURCE + properHTTPURL + "\n"
	}

	// build disclaimer
	if s.Switches[useProperCopyrightDisclaimer] {
		expect += _FIELD_DISCLAIMER + licenseDisclaimer + "\n"
	}

	// build comment
	if !s.Switches[useLicenseWithoutComment] {
		expect += _FIELD_COMMENT + licenseComment + "\n"
	}

	// build license
	if !s.Switches[useLicenseWithoutTag] {
		expect += _FIELD_LICENSE + licenseTag + "\n"
	}

	// build copyright
	if s.Switches[useProperEntities] {
		expect += s._generateCopyrightAssert() + "\n"
	}

	// assert copyright header paragraph
	if !strings.HasPrefix(str, expect) {
		th.Errorf("copyright does not have expected prefix: '%#v'",
			expect,
		)
	}

	// assert file paragraph
	if !s.Switches[useLicenseWithoutBody] {
		if !strings.Contains(str, licenseApache2BodyExpect) {
			th.Errorf("copyright has missing file paragraph")
		}
	}
}

func (s *testScenario) assertLicense(th *thelper.THelper, str string) {
	expect := ""

	// assert compulsory fields
	switch {
	case s.Switches[useLicenseWithoutFiles],
		s.Switches[useLicenseWithNilFiles]:
		if expect != str {
			th.Errorf("license created despite no files")
		}

		return
	case s.Switches[useLicenseWithoutTag],
		s.Switches[useLicenseTagAnd],
		s.Switches[useLicenseTagOr],
		s.Switches[useLicenseTagSymbol]:
		if expect != str {
			th.Errorf("license created despite no license tag")
		}

		return
	case s.Switches[useBadYear],
		!s.Switches[useProperEntities]:
		if expect != str {
			th.Errorf("license created despite no entities")
		}

		return
	default:
	}

	// assert files
	expect = _FIELD_FILES + licenseFile1 + ", " + licenseFile2
	if !strings.Contains(str, expect) {
		th.Errorf("license files did not show up")
	}

	// assert copyright
	expect = s._generateCopyrightAssert()
	if s.Switches[useProperEntities] && !strings.Contains(str, expect) {
		th.Errorf("copyright entities did not show up")
	}

	// assert comment
	expect = _FIELD_COMMENT + licenseComment
	switch {
	case !s.Switches[useLicenseWithoutComment] &&
		!strings.Contains(str, expect):
		th.Errorf("comments did not show up")
	case s.Switches[useLicenseWithoutComment] &&
		strings.Contains(str, expect):
		th.Errorf("comments did shown up despite not to")
	}

	// assert body
	switch {
	case !s.Switches[useLicenseWithoutBody] &&
		!strings.Contains(str, licenseApache2BodyExpect):
		th.Errorf("license body did not show up")
	case s.Switches[useLicenseWithoutBody] &&
		strings.Contains(str, licenseApache2BodyExpect):
		th.Errorf("license body did shown up despite not to")
	}
}

func (s *testScenario) _generateCopyrightAssert() string {
	expect := _FIELD_COPYRIGHT
	if s.Switches[useProperEntity] {
		expect += s._generateEntityAssert() + "\n" +
			_FIELD_COPYRIGHT_FOLD
	}

	expect += strconv.Itoa(yearControl) + " " + nameControl +
		" <" + emailControl + ">"

	return expect
}

//nolint:gocognit
func (s *testScenario) assertPackageList(th *thelper.THelper, str string) {
	var x, y string
	var joint bool

	expect := string(s.createPackageListName())
	switch {
	case expect == "":
		goto assert
	case !s.Switches[useProperVersionUpstream]:
		expect = ""
		goto assert
	case s.Switches[usePkgListUnknown]:
		expect = ""
		goto assert
	default:
	}
	expect += ":\n"

	switch {
	case s.Switches[useProperPackageListing]:
	case s.Switches[useEmptyPackageListing]:
		expect = ""
		goto assert
	default:
		expect = ""
		goto assert
	}

	x = " " + s._generateVersionAssert()

	y = " " + s._generateVersionAssert()
	y = strings.Replace(y, app, appControl, 1)

assert:
	if expect != "" && !strings.HasPrefix(str, expect) {
		th.Errorf("Package list did not render header: '%#v'", expect)
	}

	if x != "" && !strings.Contains(str, x) {
		th.Errorf("Package list missing '%#v'", x)
	}

	if y != "" && !strings.Contains(str, y) {
		th.Errorf("Package list missing '%#v'", y)
	}

	if x != "" && y != "" {
		expect = x + ",\n" + y
		joint = strings.Contains(str, expect)
		if !joint {
			expect = y + ",\n" + x
			joint = strings.Contains(str, expect)
		}

		if !joint {
			th.Errorf("Package list does not join 2 packages!")
		}
	}
}

func (s *testScenario) assertPackageMeta(th *thelper.THelper, str string) {
	expect := s._generateVersionAssert()

	if expect != str {
		th.Logf("PackageMeta string is unexpected. Expect: '%#v'",
			expect,
		)
	}
}

func (s *testScenario) _generateVersionAssert() string {
	expect := ""

	switch {
	case s.Switches[useProperAppName]:
		expect += app
	default:
		goto assert
	}

	expect += " ("

	switch {
	case s.Switches[useVERControlUnknown]:
		expect = ""
		goto assert
	case s.Switches[useVERControlSTER]:
		expect += string(VERCONTROL_STRICTLY_EARLIER)
	case s.Switches[useVERControlEAEQ]:
		expect += string(VERCONTROL_EARLIER_OR_EQUAL)
	case s.Switches[useVERControlEXEQ]:
		expect += string(VERCONTROL_EXACTLY_EQUAL)
	case s.Switches[useVERControlLAEQ]:
		expect += string(VERCONTROL_LATER_OR_EQUAL)
	case s.Switches[useVERControlSTLA]:
		expect += string(VERCONTROL_STRICTLY_LATER)
	default:
		expect = ""
		goto assert
	}

	expect += " "

	switch {
	case s.Switches[useProperVersionUpstream]:
		if s.Switches[useProperVersionEpoch] {
			expect += strconv.Itoa(epoch) + ":"
		}

		expect += upstream
	default:
		expect = ""
		goto assert
	}

	if s.Switches[useProperVersionRevision] {
		expect += "-"
		expect += revision
	}

	expect += ")"

	if !s.Switches[useProperArchList] {
		goto assert
	}

	switch {
	case s.Switches[useProperArch]:
		expect += " [" + arch + "]"
	case s.Switches[useAnyArch]:
		expect += " [" + archANY + "]"
	default:
	}

assert:
	return expect
}

func (s *testScenario) assertTestsuite(th *thelper.THelper, str string) {
	expect := ""

	if s.Switches[useProperPath] {
		expect = testsuitePath
	}

	if str != expect {
		th.Errorf("Testsuite string is unexpected. Expect: '%#v'",
			expect)
	}
}

func (s *testScenario) assertVCS(th *thelper.THelper, str string) {
	expect := ""

	switch {
	case s.Switches[useEmptyBrowser]:
		goto assert
	case s.Switches[useProperURL]:
		expect += _FIELD_VCS_BROWSER + properHTTPURL + "\n"
	case s.Switches[useNonHTTPURL]:
		expect += _FIELD_VCS_BROWSER + nonHTTPURL + "\n"
	}

	switch {
	case s.Switches[useVCSArch]:
		expect += string(VCS_ARCH) + ":"
	case s.Switches[useVCSBazaar]:
		expect += string(VCS_BAZAAR) + ":"
	case s.Switches[useVCSCVS]:
		expect += string(VCS_CVS) + ":"
	case s.Switches[useVCSDarcs]:
		expect += string(VCS_DARCS) + ":"
	case s.Switches[useVCSGit]:
		expect += string(VCS_GIT) + ":"
	case s.Switches[useVCSMercurial]:
		expect += string(VCS_MERCURIAL) + ":"
	case s.Switches[useVCSMonotone]:
		expect += string(VCS_MONOTONE) + ":"
	case s.Switches[useVCSSubversion]:
		expect += string(VCS_SUBVERSION) + ":"
	default:
		expect = ""
		goto assert
	}

	switch {
	case s.Switches[useProperURL]:
		expect += " " + properHTTPURL
	case s.Switches[useNonHTTPURL]:
		expect += " " + nonHTTPURL
	default:
		expect = ""
		goto assert
	}

	if s.Switches[useProperBranch] {
		expect += " -b " + branch
	}

	if s.Switches[useProperPath] {
		expect += " [" + path + "]"
	}

assert:
	if str != expect {
		th.Errorf("VCS is not within expected value: '%#v'", expect)
	}
}

func (s *testScenario) assertVersion(th *thelper.THelper, str string) {
	expect := ""

	if s.Switches[useProperVersionEpoch] {
		expect += strconv.Itoa(epoch) + ":"
	}

	expect += upstream

	if s.Switches[useDashVersionUpstream] &&
		s.Switches[useProperVersionRevision] {
		expect += upstreamDash
	}

	if s.Switches[useProperVersionRevision] {
		expect += "-" + revision
	}

	switch {
	case !s.Switches[useProperVersionUpstream]:
		expect = ""
	case s.Switches[useAlphaLedVersionUpstream]:
		expect = ""
	case s.Switches[useAlphaLedVersionRevision]:
		expect = ""
	case s.Switches[useDashVersionRevision]:
		expect = ""
	case !s.Switches[useProperVersionRevision] &&
		s.Switches[useDashVersionUpstream]:
		expect = ""
	case s.Switches[useIllegalVersionUpstream]:
		expect = ""
	case s.Switches[useIllegalVersionRevision]:
		expect = ""
	default:
	}

	if str != expect {
		th.Errorf("version output mismatched: '%s'", expect)
	}
}

func (s *testScenario) assertDescription(th *thelper.THelper, str string) {
	hasSynop := strings.Contains(str, keySynopsis)
	hasBody := strings.Contains(str, keyDescription)

	switch {
	case s.Switches[useProperSynopsis] && hasSynop:
	case !s.Switches[useLongSynopsis] && hasSynop:
		th.Errorf("has description despite bad synopsis")
	case !s.Switches[useProperSynopsis] && hasSynop:
		th.Errorf("synopsis appears even none is given")
	case !s.Switches[useProperSynopsis] && hasBody:
		th.Errorf("body appears even synopsis is bad")
	case !s.Switches[useProperSynopsis] && !hasBody:
	case s.Switches[useProperDescriptionText] && hasBody:
	case s.Switches[useProperDescriptionText] && !hasBody:
		th.Errorf("missing body description")
	case !s.Switches[useProperDescriptionText] && hasBody:
		th.Errorf("body description appears when none is given")
	}
}

func (s *testScenario) assertPanic(th *thelper.THelper, ret interface{}) {
	didPanic := true

	out, ok := ret.(string)
	if ok && out == noPanicFlag {
		didPanic = false
	}

	switch {
	case s.Switches[expectPanic] && didPanic:
	case s.Switches[expectPanic] && !didPanic:
		th.Errorf("panic did not occur")
	case !s.Switches[expectPanic] && didPanic:
		th.Errorf("panic occurred")
	case !s.Switches[expectPanic] && !didPanic:
	}
}

func (s *testScenario) assertEntity(th *thelper.THelper, str string) {
	expect := s._generateEntityAssert()

	if str != expect {
		th.Errorf("unexpected str output")
	}
}

func (s *testScenario) _generateEntityAssert() string {
	expect := ""

	if s.Switches[useProperYear] {
		expect += strconv.Itoa(year) + " "
	}

	if s.Switches[useProperName] {
		expect += name
	} else {
		expect = ""
	}

	if s.Switches[useProperEmail] {
		if !s.Switches[useProperName] {
			expect = ""
		} else {
			expect += " <" + email + ">"
		}
	} else {
		expect = ""
	}

	return expect
}

func (s *testScenario) createName() string {
	switch {
	case s.Switches[useProperName]:
		return name
	default:
		return ""
	}
}

func (s *testScenario) createYear() int {
	switch {
	case s.Switches[useBadYear]:
		return year * -1
	case s.Switches[useProperYear]:
		return year
	default:
		return 0
	}
}

func (s *testScenario) createEmail() string {
	switch {
	case s.Switches[useProperEmail]:
		return email
	default:
		return ""
	}
}

func (s *testScenario) createSynopsis() string {
	switch {
	case s.Switches[useProperSynopsis]:
		return strings.Repeat(keySynopsis, 65/len(keySynopsis))
	case s.Switches[useLongSynopsis]:
		return strings.Repeat(keySynopsis, 300/len(keySynopsis))
	default:
		return ""
	}
}

func (s *testScenario) createDescription() *Description {
	switch {
	case s.Switches[useProperDescription]:
		return &Description{
			Synopsis: strings.Repeat(keySynopsis,
				65/len(keySynopsis)),
			Body: strings.Repeat(keyDescription+"A ", 10) +
				"\n\n" +
				strings.Repeat(keyDescription+"B ", 10) +
				"\n\n" +
				strings.Repeat(keyDescription+"C ", 10) +
				"\n\n",
		}
	case s.Switches[useFaultyDescription]:
		return &Description{
			Synopsis: strings.Repeat(keySynopsis,
				300/len(keySynopsis)),
			Body: strings.Repeat(keyDescription+"A ", 10) +
				"\n\n" +
				strings.Repeat(keyDescription+"B ", 10) +
				"\n\n" +
				strings.Repeat(keyDescription+"C ", 10) +
				"\n\n",
		}
	case s.Switches[useEmptyDescription]:
		return &Description{
			Body: strings.Repeat(keyDescription+"A ", 10) +
				"\n\n" +
				strings.Repeat(keyDescription+"B ", 10) +
				"\n\n" +
				strings.Repeat(keyDescription+"C ", 10) +
				"\n\n",
		}
	}
	return nil
}

func (s *testScenario) createDescriptionText() string {
	switch {
	case s.Switches[useProperDescriptionText]:
		return strings.Repeat(keyDescription+"A ", 80) + "\n" +
			strings.Repeat(keyDescription+"B ", 80) + "\n" +
			strings.Repeat(keyDescription+"C ", 80) + "\n"
	default:
		return ""
	}
}

func (s *testScenario) createVCS() *VCS {
	address, _ := url.Parse(properHTTPURL)

	switch {
	case s.Switches[useProperVCS]:
		return &VCS{
			Browser: address,
			Type:    VCS_GIT,
			URL:     address,
			Branch:  branch,
			Path:    path,
		}
	case s.Switches[useEmptyVCS]:
		return &VCS{}
	case s.Switches[useFaultyVCS]:
		return &VCS{
			Browser: address,
			Branch:  branch,
			Path:    path,
		}
	}

	return nil
}

func (s *testScenario) createVersionUpstream() (ret string) {
	if s.Switches[useAlphaLedVersionUpstream] {
		ret += upstreamPrefix
	}

	if s.Switches[useProperVersionUpstream] {
		ret += upstream
	} else if s.Switches[useIllegalVersionUpstream] {
		ret += upstreamIllegal
	}

	if s.Switches[useDashVersionUpstream] {
		ret += upstreamDash
	}

	if !s.Switches[useProperVersionUpstream] &&
		!s.Switches[useIllegalVersionUpstream] {
		ret = ""
	}

	return ret
}

func (s *testScenario) createVersionRevision() (ret string) {
	if s.Switches[useAlphaLedVersionRevision] {
		ret += revisionPrefix
	}

	if s.Switches[useProperVersionRevision] {
		ret += revision
	} else if s.Switches[useIllegalVersionRevision] {
		ret += revisionIllegal
	}

	if s.Switches[useDashVersionRevision] {
		ret += revisionDash
	}

	if !s.Switches[useProperVersionRevision] &&
		!s.Switches[useIllegalVersionRevision] {
		ret = ""
	}

	return ret
}

func (s *testScenario) createVersionEpoch() (ret uint) {
	if s.Switches[useProperVersionEpoch] {
		return epoch
	}

	return 0
}

func (s *testScenario) createVCSType() VCSType {
	switch {
	case s.Switches[useVCSArch]:
		return VCS_ARCH
	case s.Switches[useVCSBazaar]:
		return VCS_BAZAAR
	case s.Switches[useVCSCVS]:
		return VCS_CVS
	case s.Switches[useVCSDarcs]:
		return VCS_DARCS
	case s.Switches[useVCSGit]:
		return VCS_GIT
	case s.Switches[useVCSMercurial]:
		return VCS_MERCURIAL
	case s.Switches[useVCSMonotone]:
		return VCS_MONOTONE
	case s.Switches[useVCSSubversion]:
		return VCS_SUBVERSION
	case s.Switches[useVCSUnknown]:
		return "Vcs-unknown"
	default:
		return ""
	}
}

func (s *testScenario) createURL() (x *url.URL) {
	switch {
	case s.Switches[useProperURL]:
		x, _ = url.Parse(properHTTPURL)
	case s.Switches[useNonHTTPURL]:
		x, _ = url.Parse(nonHTTPURL)
	case s.Switches[useEmptyURL]:
		x = new(url.URL)
	default:
		x = nil
	}

	return x
}

func (s *testScenario) createBrowser() (x *url.URL) {
	switch {
	case s.Switches[useEmptyBrowser]:
		x = new(url.URL)
	case s.Switches[useProperURL], s.Switches[useProperBrowser]:
		x, _ = url.Parse(properHTTPURL)
	case s.Switches[useNonHTTPURL]:
		x, _ = url.Parse(nonHTTPURL)
	default:
		x = nil
	}

	return x
}

func (s *testScenario) createBranch() string {
	switch {
	case s.Switches[useProperBranch]:
		return branch
	default:
		return ""
	}
}

func (s *testScenario) createTestsuite() *Testsuite {
	switch {
	case s.Switches[useProperTestsuite]:
		return &Testsuite{
			Paths: []string{testsuitePath,
				testsuitePath,
				testsuitePath},
		}
	case s.Switches[useEmptyTestsuite]:
		return &Testsuite{}
	case s.Switches[useFaultyTestsuite]:
		return &Testsuite{
			Paths: []string{ghostPath, ghostPath, ghostPath},
		}
	default:
		return nil
	}
}

func (s *testScenario) createPath() string {
	s._createChangelogFile()

	switch {
	case s.Switches[useProperPath] && s.TestType == testTestsuiteString:
		return testsuitePath
	case s.Switches[useProperPath]:
		return path
	case s.Switches[useProperPathDirectory]:
		return directory
	default:
		return ""
	}
}

func (s *testScenario) createChangelogParseStatus() changelogParseStatus {
	switch {
	case s.Switches[useCompletedChangelogParseStatus]:
		return _CHANGELOG_PARSE_SIGNATURE
	case s.Switches[useUnknownChangelogParseStatus]:
		return 100000
	default:
		return _CHANGELOG_PARSE_HEADER
	}
}

func (s *testScenario) createChangelogData() []string {
	data := s.__createChangelogData()
	if data == "" {
		return []string{}
	}

	return strings.Split(data, "\n")
}

func (s *testScenario) _createChangelogFile() {
	var err error

	data := s.__createChangelogData()
	if data == "" {
		return
	}

	_ = os.RemoveAll(path)

	err = s.__writeFile(path, _PERMISSION_FILE, data)
	if err != nil {
		panic("failed to setup changelog file for testing")
	}
}

func (s *testScenario) __createChangelogData() string {
	switch {
	case s.Switches[createNoChangelogFile]:
		return ""
	case s.Switches[createBadHeaderChangelogFile]:
		return changelogBadHeaderFile
	case s.Switches[createBadPackageHeaderChangelogFile]:
		return changelogBadPackageHeaderFile
	case s.Switches[createBadVersionHeaderChangelogFile]:
		return changelogBadVersionHeaderFile
	case s.Switches[createDoubleHeaderChangelogFile]:
		return changelogDoubleHeaderFile
	case s.Switches[createSigHeaderChangelogFile]:
		return changelogSigHeaderFile
	case s.Switches[createChangesLeadChangelogFile]:
		return changelogChangesLeadFile
	case s.Switches[createSigLeadChangelogFile]:
		return changelogSigLeadFile
	case s.Switches[createTailingHeaderChangelogFile]:
		return changelogTailingHeaderFile
	case s.Switches[createTailingChangesChangelogFile]:
		return changelogTailingChangesFile
	case s.Switches[createTailingSigChangelogFile]:
		return changelogTailingSigFile
	case s.Switches[createBadVersionEpochHeaderChangelogFile]:
		return changelogBadVersionEpochFile
	case s.Switches[createBadVersionUpstreamHeaderChangelogFile]:
		return changelogBadVersionUpstreamFile
	case s.Switches[createMissingDistroHeaderChangelogFile]:
		return changelogMissingDistroFile
	case s.Switches[createBadNameChangelogFile]:
		return changelogBadNameFile
	case s.Switches[createBadEmailChangelogFile]:
		return changelogBadEmailFile
	case s.Switches[createBadEmailDelimiterChangelogFile]:
		return changelogBadEmailDelimiterFile
	case s.Switches[createBadTimestampChangelogFile]:
		return changelogBadTimestampFile
	default:
		return changelogFile
	}
}

func (s *testScenario) __writeFile(path string,
	perm os.FileMode, data string) (err error) {
	var f *os.File

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, _PERMISSION_FILE)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_OPEN_FAILED, err)
	}

	_, err = f.WriteString(data)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_WRITE_FAILED, err)
		goto closeWriter
	}

	err = f.Chmod(perm)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_CHMOD_FAILED, err)
	}

closeWriter:
	_ = f.Sync()
	f.Close()
	return err
}

func (s *testScenario) createPaths() (x []string) {
	switch {
	case s.Switches[useEmptyPath]:
		return []string{}
	case s.Switches[useProperPath] && s.TestType == testTestsuiteString:
		return []string{testsuitePath, testsuitePath, testsuitePath}
	case s.Switches[useProperPath]:
		return []string{path, path, path}
	case s.Switches[useGhostPath]:
		return []string{ghostPath, ghostPath, ghostPath}
	default:
		return nil
	}
}

func (s *testScenario) createAppName() string {
	if s.Switches[useProperAppName] {
		return app
	}

	return ""
}

func (s *testScenario) createVersion() *Version {
	if s.Switches[useNilVersion] {
		return nil
	}

	return &Version{
		Upstream: s.createVersionUpstream(),
		Revision: s.createVersionRevision(),
		Epoch:    s.createVersionEpoch(),
	}
}

func (s *testScenario) createVERControl() VERControl {
	switch {
	case s.Switches[useVERControlUnknown]:
		return upstream
	case s.Switches[useVERControlSTER]:
		return VERCONTROL_STRICTLY_EARLIER
	case s.Switches[useVERControlEAEQ]:
		return VERCONTROL_EARLIER_OR_EQUAL
	case s.Switches[useVERControlEXEQ]:
		return VERCONTROL_EXACTLY_EQUAL
	case s.Switches[useVERControlLAEQ]:
		return VERCONTROL_LATER_OR_EQUAL
	case s.Switches[useVERControlSTLA]:
		return VERCONTROL_STRICTLY_LATER
	default:
		return ""
	}
}

func (s *testScenario) createArchitecture() string {
	switch {
	case s.Switches[useProperArch]:
		return arch
	case s.Switches[useAnyArch]:
		return archANY
	case s.Switches[useUnknownArch]:
		return unknown
	default:
		return ""
	}
}

func (s *testScenario) createArchitectures() []string {
	switch {
	case s.Switches[useProperArchList]:
		return []string{s.createArchitecture()}
	case s.Switches[useEmptyArchList]:
		return []string{}
	default:
		return nil
	}
}

func (s *testScenario) createPackageListName() PackageListType {
	switch {
	case s.Switches[usePkgListBuildDepends]:
		return PKG_LIST_BUILD_DEPENDS
	case s.Switches[usePkgListBuildDependsIndep]:
		return PKG_LIST_BUILD_DEPENDS_INDEP
	case s.Switches[usePkgListBuildDependsArch]:
		return PKG_LIST_BUILD_DEPENDS_ARCH
	case s.Switches[usePkgListBuildConflicts]:
		return PKG_LIST_BUILD_CONFLICTS
	case s.Switches[usePkgListBuildConflictsIndep]:
		return PKG_LIST_BUILD_CONFLICTS_INDEP
	case s.Switches[usePkgListBuildConflictsArch]:
		return PKG_LIST_BUILD_CONFLICTS_ARCH
	case s.Switches[usePkgListPreDepends]:
		return PKG_LIST_PRE_DEPENDS
	case s.Switches[usePkgListDepends]:
		return PKG_LIST_DEPENDS
	case s.Switches[usePkgListRecommends]:
		return PKG_LIST_RECOMMENDS
	case s.Switches[usePkgListSuggests]:
		return PKG_LIST_SUGGESTS
	case s.Switches[usePkgListEnhances]:
		return PKG_LIST_ENHANCES
	case s.Switches[usePkgListBreaks]:
		return PKG_LIST_BREAKS
	case s.Switches[usePkgListConflicts]:
		return PKG_LIST_CONFLICTS
	case s.Switches[usePkgListProvides]:
		return PKG_LIST_PROVIDES
	case s.Switches[usePkgListReplaces]:
		return PKG_LIST_REPLACES
	case s.Switches[usePkgListBuiltUsing]:
		return PKG_LIST_BUILT_USING
	case s.Switches[usePkgListUnknown]:
		return unknown
	default:
		return ""
	}
}

func (s *testScenario) createRelPackages() map[PackageListType]*PackageList {
	var meta map[string]*PackageMeta

	switch {
	case s.Switches[useNilPackageList]:
		return nil
	case s.Switches[useNilContentPackageList]:
		return map[PackageListType]*PackageList{
			PKG_LIST_BUILD_DEPENDS:         nil,
			PKG_LIST_BUILD_DEPENDS_INDEP:   nil,
			PKG_LIST_BUILD_DEPENDS_ARCH:    nil,
			PKG_LIST_BUILD_CONFLICTS:       nil,
			PKG_LIST_BUILD_CONFLICTS_INDEP: nil,
			PKG_LIST_BUILD_CONFLICTS_ARCH:  nil,
			PKG_LIST_PRE_DEPENDS:           nil,
			PKG_LIST_DEPENDS:               nil,
			PKG_LIST_RECOMMENDS:            nil,
			PKG_LIST_SUGGESTS:              nil,
			PKG_LIST_ENHANCES:              nil,
			PKG_LIST_BREAKS:                nil,
			PKG_LIST_CONFLICTS:             nil,
			PKG_LIST_PROVIDES:              nil,
			PKG_LIST_REPLACES:              nil,
			PKG_LIST_BUILT_USING:           nil,
		}
	case s.Switches[useProperPackageList]:
		x := &PackageMeta{
			Name: app,
			Version: &Version{
				Upstream: upstream,
			},
			VERControl:    VERCONTROL_EXACTLY_EQUAL,
			Architectures: []string{arch},
		}

		y := &PackageMeta{
			Name: appControl,
			Version: &Version{
				Upstream: upstream,
			},
			VERControl:    VERCONTROL_EXACTLY_EQUAL,
			Architectures: []string{arch},
		}

		meta = map[string]*PackageMeta{
			x.Name: x,
			y.Name: y,
		}
	case s.Switches[useEmptyPackageList]:
		meta = map[string]*PackageMeta{}
	case s.Switches[useFaultyPackageList]:
		x := &PackageMeta{
			Version: &Version{
				Upstream: upstream,
			},
			Architectures: []string{arch},
		}

		y := &PackageMeta{
			Version: &Version{
				Upstream: upstream,
			},
			Architectures: []string{arch},
		}

		meta = map[string]*PackageMeta{
			x.Name: x,
			y.Name: y,
		}
	default:
		meta = nil
	}

	return map[PackageListType]*PackageList{
		PKG_LIST_BUILD_DEPENDS: {
			Name: PKG_LIST_BUILD_DEPENDS,
			List: meta,
		},
		PKG_LIST_BUILD_DEPENDS_INDEP: {
			Name: PKG_LIST_BUILD_DEPENDS_INDEP,
			List: meta,
		},
		PKG_LIST_BUILD_DEPENDS_ARCH: {
			Name: PKG_LIST_BUILD_DEPENDS_ARCH,
			List: meta,
		},
		PKG_LIST_BUILD_CONFLICTS: {
			Name: PKG_LIST_BUILD_CONFLICTS,
			List: meta,
		},
		PKG_LIST_BUILD_CONFLICTS_INDEP: {
			Name: PKG_LIST_BUILD_CONFLICTS_INDEP,
			List: meta,
		},
		PKG_LIST_BUILD_CONFLICTS_ARCH: {
			Name: PKG_LIST_BUILD_CONFLICTS_ARCH,
			List: meta,
		},
		PKG_LIST_PRE_DEPENDS: {
			Name: PKG_LIST_PRE_DEPENDS,
			List: meta,
		},
		PKG_LIST_DEPENDS: {
			Name: PKG_LIST_DEPENDS,
			List: meta,
		},
		PKG_LIST_RECOMMENDS: {
			Name: PKG_LIST_RECOMMENDS,
			List: meta,
		},
		PKG_LIST_SUGGESTS: {
			Name: PKG_LIST_SUGGESTS,
			List: meta,
		},
		PKG_LIST_ENHANCES: {
			Name: PKG_LIST_ENHANCES,
			List: meta,
		},
		PKG_LIST_BREAKS: {
			Name: PKG_LIST_BREAKS,
			List: meta,
		},
		PKG_LIST_CONFLICTS: {
			Name: PKG_LIST_CONFLICTS,
			List: meta,
		},
		PKG_LIST_PROVIDES: {
			Name: PKG_LIST_PROVIDES,
			List: meta,
		},
		PKG_LIST_REPLACES: {
			Name: PKG_LIST_REPLACES,
			List: meta,
		},
		PKG_LIST_BUILT_USING: {
			Name: PKG_LIST_BUILT_USING,
			List: meta,
		},
	}
}

func (s *testScenario) createPackageListing() map[string]*PackageMeta {
	switch {
	case s.Switches[useProperPackageListing]:
		x := &PackageMeta{
			Name:          s.createAppName(),
			Version:       s.createVersion(),
			VERControl:    s.createVERControl(),
			Architectures: s.createArchitectures(),
		}

		y := &PackageMeta{
			Name:          appControl,
			Version:       s.createVersion(),
			VERControl:    s.createVERControl(),
			Architectures: s.createArchitectures(),
		}

		return map[string]*PackageMeta{
			x.Name: x,
			y.Name: y,
		}
	case s.Switches[useEmptyPackageListing]:
		return map[string]*PackageMeta{}
	default:
		return nil
	}
}

func (s *testScenario) createEntity() (x *Entity) {
	switch {
	case s.Switches[useProperEntity]:
		x = &Entity{
			Name:  s.createName(),
			Email: s.createEmail(),
			Year:  s.createYear(),
		}
	default:
		x = nil
	}

	return x
}

func (s *testScenario) createEntities() (x []*Entity) {
	switch {
	case s.Switches[useProperEntities]:
		x = []*Entity{}
		x = append(x, s.createEntity(), &Entity{
			Name:  nameControl,
			Email: emailControl,
			Year:  yearControl,
		})
	case s.Switches[useEmptyEntities]:
		x = []*Entity{}
	default:
	}

	return x
}

func (s *testScenario) createLicense() (f []string,
	l string, b string, c string) {
	// create the large data first
	f = []string{licenseFile1, licenseFile2}
	l = licenseTag
	c = licenseComment
	b = licenseBody

	if s.Switches[useLicenseWithoutFiles] {
		f = []string{}
	}

	switch {
	case s.Switches[useLicenseWithoutTag]:
		l = ""
	case s.Switches[useLicenseTagSymbol]:
		l = l + ", " + l
	case s.Switches[useLicenseTagOr]:
		l = l + " or " + l
	case s.Switches[useLicenseTagAnd]:
		l = l + " and " + l
	}

	if s.Switches[useLicenseWithoutComment] {
		c = ""
	}

	if s.Switches[useLicenseWithoutBody] {
		b = ""
	}

	if s.Switches[useLicenseWithNilFiles] {
		f = nil
	}

	return f, l, b, c
}

func (s *testScenario) createLicenses() (x []*License) {
	switch {
	case s.Switches[useProperLicenses]:
		x = []*License{
			{
				License: licenseTag,
				Body:    licenseBody,
				Comment: licenseComment,
				Copyright: []*Entity{
					{
						Name:  nameControl,
						Email: emailControl,
						Year:  yearControl,
					},
				},
				Files: []string{licenseFile1, licenseFile2},
			},
		}
	case s.Switches[useEmptyLicenses]:
		x = []*License{}
	default:
	}

	return x
}

func (s *testScenario) createCopyrightFormat() CopyrightFormat {
	switch {
	case s.Switches[useCopyrightFormat1p0]:
		return COPYRIGHT_FORMAT_1_0
	default:
		return ""
	}
}

func (s *testScenario) createCopyrightSource() string {
	switch {
	case s.Switches[useProperCopyrightSource]:
		return properHTTPURL
	default:
		return ""
	}
}

func (s *testScenario) createCopyrightDisclaimer() string {
	switch {
	case s.Switches[useProperCopyrightDisclaimer]:
		return licenseDisclaimer
	default:
		return ""
	}
}

func (s *testScenario) createChangelogUrgency() ChangelogUrgency {
	switch {
	case s.Switches[useChangelogUrgencyLow]:
		return CHANGELOG_URGENCY_LOW
	case s.Switches[useChangelogUrgencyMedium]:
		return CHANGELOG_URGENCY_MEDIUM
	case s.Switches[useChangelogUrgencyHigh]:
		return CHANGELOG_URGENCY_HIGH
	case s.Switches[useChangelogUrgencyEmergency]:
		return CHANGELOG_URGENCY_EMERGENCY
	case s.Switches[useChangelogUrgencyCritical]:
		return CHANGELOG_URGENCY_CRITICAL
	case s.Switches[useChangelogUrgencyUnknown]:
		return urgencyUnknown
	default:
		return ""
	}
}

func (s *testScenario) createChangelogChanges() (x []string) {
	switch {
	case s.Switches[useProperChangelogChanges]:
		x = []string{changelogA, changelogB}
	case s.Switches[useEmptyChangelogChanges]:
		x = []string{}
	default:
		x = nil
	}

	return x
}

func (s *testScenario) createChangelogDistro() []string {
	switch {
	case s.Switches[useProperDistro]:
		return []string{distroDebian, distroUbuntu}
	case s.Switches[useEmptyDistro]:
		return []string{}
	default:
		return nil
	}
}

func (s *testScenario) createTimestamp() *time.Time {
	switch {
	case s.Switches[useProperTimestamp]:
		x, _ := time.Parse(_CHANGELOG_TIMESTAMP_FORMAT, timestamp)
		return &x
	case s.Switches[useEmptyTimestamp]:
		x, _ := time.Parse(_CHANGELOG_TIMESTAMP_FORMAT, timestampZero)
		return &x
	default:
		return nil
	}
}

func (s *testScenario) createStatFx() func(string) (os.FileInfo, error) {
	if s.Switches[simulateStatFxError] {
		return func(path string) (i os.FileInfo, err error) {
			return nil, fmt.Errorf(simulateStatFxError)
		}
	}

	return nil
}

func (s *testScenario) createSection() string {
	switch {
	case s.Switches[useProperSection]:
		return section
	case s.Switches[useEmptySection]:
	case s.Switches[useFaultySection]:
		return unknown
	}

	return ""
}

func (s *testScenario) createStandardsVersion() StandardsVersion {
	switch {
	case s.Switches[useProperStandardsVersion]:
		return STANDARD_VER_4_6_0
	case s.Switches[useFaultyStandardsVersion]:
		return unknown
	default:
		return ""
	}
}

func (s *testScenario) createPackageType() PackageType {
	switch {
	case s.Switches[usePackageTypeDEB]:
		return PACKAGE_TYPE_DEB
	case s.Switches[usePackageTypeUDEB]:
		return PACKAGE_TYPE_UDEB
	case s.Switches[usePackageTypeUnknown]:
		return unknown
	default:
		return ""
	}
}

func (s *testScenario) createRulesRequiresRoot() string {
	switch {
	case s.Switches[useRRRBinTarget]:
		return RULES_ROOT_BINARY_TARGETS
	case s.Switches[useRRRNo]:
		return RULES_ROOT_NO
	case s.Switches[useRRRCustom]:
		return customRRR
	case s.Switches[useRRRUnknown]:
		return unknown
	default:
		return ""
	}
}

func (s *testScenario) createPriority() Priority {
	switch {
	case s.Switches[usePriorityRequired]:
		return PRIORITY_REQUIRED
	case s.Switches[usePriorityImportant]:
		return PRIORITY_IMPORTANT
	case s.Switches[usePriorityStandard]:
		return PRIORITY_STANDARD
	case s.Switches[usePriorityOptional]:
		return PRIORITY_OPTIONAL
	case s.Switches[usePriorityUnknown]:
		return unknown
	default:
		return ""
	}
}

func (s *testScenario) createEssential() bool {
	switch {
	case s.Switches[useEssentialFlag]:
		return true
	default:
		return false
	}
}

func (s *testScenario) createBuildSource() bool {
	switch {
	case s.Switches[useBuildSourceFlag]:
		return true
	default:
		return false
	}
}

func (s *testScenario) createUploaders() (x []*Entity) {
	switch {
	case s.Switches[useProperUploaders]:
		x = []*Entity{}
		x = append(x, s.createEntity(),
			nil,
			&Entity{
				Name:  nameControl,
				Email: emailControl,
				Year:  yearControl,
			},
			&Entity{
				Name:  nameControl2,
				Email: emailControl2,
				Year:  yearControl2,
			},
		)
	case s.Switches[useFaultyUploaders]:
		x = []*Entity{}
		x = append(x, s.createEntity(),
			nil,
			&Entity{
				Email: emailControl,
				Year:  yearControl,
			},
			&Entity{
				Name: nameControl2,
				Year: yearControl2,
			},
		)
	case s.Switches[useEmptyUploaders]:
		x = []*Entity{}
	default:
		x = nil
	}

	return x
}

func (s *testScenario) createSourceFormat() SourceFormatType {
	switch {
	case s.Switches[useSourceFormatNative3p0]:
		return SOURCE_FORMAT_NATIVE_3_0
	case s.Switches[useSourceFormatQuilt3p0]:
		return SOURCE_FORMAT_QUILT_3_0
	default:
		return ""
	}
}

func (s *testScenario) createSourceLocalOptions() string {
	switch {
	case s.Switches[useProperSourceLocalOptions]:
		return sourceLocalOptions
	default:
		return ""
	}
}

func (s *testScenario) createSourceOptions() string {
	switch {
	case s.Switches[useProperSourceOptions]:
		return sourceOptions
	default:
		return ""
	}
}

func (s *testScenario) preConfigureTestedData() {
	// control, copyright, changelog
	s.Switches[useProperName] = true
	s.Switches[useProperEmail] = true
	s.Switches[useProperYear] = true
	s.Switches[useProperEntity] = true
	s.Switches[useProperAppName] = true

	// control, changelog
	s.Switches[useProperVersionUpstream] = true

	// control
	s.Switches[useProperPackageList] = true
	s.Switches[useProperURL] = true
	s.Switches[useProperDescription] = true
	s.Switches[useProperVCS] = true
	s.Switches[useProperTestsuite] = true
	s.Switches[useProperSection] = true
	s.Switches[usePackageTypeDEB] = true
	s.Switches[useRRRBinTarget] = true
	s.Switches[usePriorityRequired] = true
	s.Switches[useProperArch] = true
	s.Switches[useProperUploaders] = true
	s.Switches[useEssentialFlag] = true
	if !s.Switches[useFaultyTestedDataControl] {
		s.Switches[useProperStandardsVersion] = true
	} else {
		s.Switches[useProperStandardsVersion] = false
	}

	// copyright
	s.Switches[useProperEntities] = true
	s.Switches[useProperLicenses] = true
	s.Switches[useProperCopyrightSource] = true
	s.Switches[useProperCopyrightDisclaimer] = true
	if !s.Switches[useFaultyTestedDataCopyright] {
		s.Switches[useCopyrightFormat1p0] = true
	} else {
		s.Switches[useCopyrightFormat1p0] = false
	}

	// changelog
	s.Switches[useChangelogUrgencyLow] = true
	s.Switches[useProperDistro] = true
	s.Switches[useProperChangelogChanges] = true
	if !s.Switches[useFaultyTestedDataChangelog] {
		s.Switches[useProperTimestamp] = true
	} else {
		s.Switches[useProperTimestamp] = false
	}

	// Source
	if !s.Switches[useFaultyTestedDataSource] {
		s.Switches[useSourceFormatNative3p0] = true
	} else {
		s.Switches[useSourceFormatNative3p0] = false
	}
}

func (s *testScenario) createControl() *Control {
	// depends on: preConfigureTestedData()
	if s.Switches[useNilControl] {
		return nil
	}

	return &Control{
		Maintainer:        s.createEntity(),
		Packages:          s.createRelPackages(),
		Version:           s.createVersion(),
		Homepage:          s.createURL(),
		Description:       s.createDescription(),
		VCS:               s.createVCS(),
		Testsuite:         s.createTestsuite(),
		Name:              s.createAppName(),
		Section:           s.createSection(),
		StandardsVersion:  s.createStandardsVersion(),
		PackageType:       s.createPackageType(),
		RulesRequiresRoot: s.createRulesRequiresRoot(),
		Priority:          s.createPriority(),
		Architecture:      s.createArchitecture(),
		Uploaders:         s.createUploaders(),
		Essential:         s.createEssential(),
		BuildSource:       s.createBuildSource(),
	}
}

func (s *testScenario) createCopyright() *Copyright {
	// depends on: preConfigureTestedData()
	if s.Switches[useNilCopyright] {
		return nil
	}

	_, license, _, comment := s.createLicense()

	return &Copyright{
		Format:     s.createCopyrightFormat(),
		Name:       s.createAppName(),
		Contact:    s.createEntity(),
		Source:     s.createCopyrightSource(),
		Disclaimer: s.createCopyrightDisclaimer(),
		Comment:    comment,
		License:    license,
		Copyright:  s.createEntities(),
		Licenses:   s.createLicenses(),
	}
}

func (s *testScenario) createChangelog() *Changelog {
	// depends on: preConfigureTestedData()
	if s.Switches[useNilChangelog] {
		return nil
	}

	return &Changelog{
		Version:      s.createVersion(),
		Maintainer:   s.createEntity(),
		Timestamp:    s.createTimestamp(),
		Package:      s.createAppName(),
		Urgency:      s.createChangelogUrgency(),
		Changes:      s.createChangelogChanges(),
		Distribution: s.createChangelogDistro(),
		Path:         s.createPath(),
	}
}

func (s *testScenario) createSource() *Source {
	// depends on: preConfigureTestedData()
	if s.Switches[useNilSource] {
		return nil
	}

	return &Source{
		Format:       s.createSourceFormat(),
		LocalOptions: s.createSourceLocalOptions(),
		Options:      s.createSourceOptions(),
	}
}

func (s *testScenario) createManpage() (x map[string]string) {
	switch {
	case s.Switches[useProperManpage]:
		x = map[string]string{
			manPageTag1: manPage1Content,
			"":          manPage1Content,
			manPageTagX: "",
		}
	case s.Switches[useEmptyManpage]:
		x = map[string]string{}
	}

	return x
}

func (s *testScenario) createScripts() (x map[ShellScriptType]string) {
	switch {
	case s.Switches[useProperScripts]:
		x = map[ShellScriptType]string{
			SHELL_PRE_INSTALL:  shellScript,
			SHELL_POST_INSTALL: shellScript,
			SHELL_PRE_REMOVE:   shellScript,
			SHELL_POST_REMOVE:  shellScript,
		}
	case s.Switches[useFaultyScripts]:
		x = map[ShellScriptType]string{
			SHELL_POST_REMOVE: "",
			unknown:           shellScript,
		}
	case s.Switches[useEmptyScripts]:
		x = map[ShellScriptType]string{}
	}

	return x
}

func (s *testScenario) createInstalls() (x map[string]string) {
	switch {
	case s.Switches[useProperInstall]:
		x = map[string]string{
			installPath:  installProgram,
			"":           installProgram,
			installPath2: installProgram,
		}
	case s.Switches[use3p7p0Install]:
		x = map[string]string{
			"":                   installProgram,
			"usr/X11R6/testapp":  installProgram,
			"/usr/X11R6/testapp": installProgram,
		}
	case s.Switches[use4p0p0Install]:
		x = map[string]string{
			"":                      installProgram,
			"usr/bin/testapp/app":   installProgram,
			"usr/share/testapp/app": installProgram,
			"testapp/app":           installProgram,
		}
	case s.Switches[use4p6p0Install]:
		x = map[string]string{
			"":                   installProgram,
			"usr/lib64/testapp":  installProgram,
			"/usr/lib64/testapp": installProgram,
		}
	case s.Switches[useEmptyInstall]:
		x = map[string]string{}
	}

	return x
}

func (s *testScenario) createRules() (x string) {
	switch {
	case s.Switches[useProperRules]:
		x = rulesFile
	default:
		x = ""
	}

	return x
}

func (s *testScenario) createCompat() (x uint) {
	switch {
	case s.Switches[useProperCompat]:
		x = compat
	default:
		x = 0
	}

	return x
}

func (s *testScenario) createWorkDir() (x string) {
	switch {
	case s.Switches[useProperWorkDir]:
		x = workDir
	case s.Switches[useFaultyWorkDir]:
		x = restrictedWorkDir
	default:
		x = ""
	}

	return x
}

func (s *testScenario) cleanUpDataGenerate() {
	// remove all working directories
	_ = os.RemoveAll(workDir)
	_ = os.Chmod(restrictedWorkDir, 0777)
	_ = os.RemoveAll(restrictedWorkDir)

	// remove all changelog file
	_ = os.RemoveAll(path)
}

func (s *testScenario) setupDataGenerate() {
	// call remove to purge previous test leftovers
	s.cleanUpDataGenerate()

	// create all working directories for testing
	_ = os.MkdirAll(restrictedWorkDir, 0755)
	_ = os.MkdirAll(workDir, 0755)
	_ = os.Chmod(restrictedWorkDir, 0600)
}
