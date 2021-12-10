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

func getTestScenarios() []testScenario {
	return []testScenario{
		{
			UID:      1,
			TestType: testEntityString,
			Description: `
Entity.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:  true,
				useProperEmail: true,
				useProperYear:  true,
				expectPanic:    false,
			},
		}, {
			UID:      2,
			TestType: testEntityString,
			Description: `
Entity.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is not given.
4. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:  true,
				useProperEmail: true,
				useProperYear:  false,
				expectPanic:    false,
			},
		}, {
			UID:      3,
			TestType: testEntityString,
			Description: `
Entity.String() should work properly when:
1. Name is not given.
2. Email is given.
3. Year is not given.
4. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:  false,
				useProperEmail: true,
				useProperYear:  false,
				expectPanic:    true,
			},
		}, {
			UID:      4,
			TestType: testEntityString,
			Description: `
Entity.String() should work properly when:
1. Name is given.
2. Email is not given.
3. Year is not given.
4. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:  true,
				useProperEmail: false,
				useProperYear:  false,
				expectPanic:    true,
			},
		}, {
			UID:      5,
			TestType: testEntityString,
			Description: `
Entity.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is badly given.
4. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:  true,
				useProperEmail: true,
				useBadYear:     true,
				expectPanic:    false,
			},
		}, {
			UID:      6,
			TestType: testDescriptionString,
			Description: `
Description.String() should work properly when:
1. Synopsis is properly given.
2. Body is properly given.
3. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperSynopsis:        true,
				useProperDescriptionText: true,
				expectPanic:              false,
			},
		}, {
			UID:      7,
			TestType: testDescriptionString,
			Description: `
Description.String() should work properly when:
1. Synopsis is not given.
2. Body is properly given.
3. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperSynopsis:        false,
				useProperDescriptionText: true,
				expectPanic:              true,
			},
		}, {
			UID:      8,
			TestType: testDescriptionString,
			Description: `
Description.String() should work properly when:
1. Synopsis is properly given.
2. Body is not given.
3. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperSynopsis:        true,
				useProperDescriptionText: false,
				expectPanic:              false,
			},
		}, {
			UID:      9,
			TestType: testDescriptionString,
			Description: `
Description.String() should work properly when:
1. Synopsis is badly given.
2. Body is given.
3. Panic is expected.
`,
			Switches: map[string]bool{
				useLongSynopsis:          true,
				useProperDescriptionText: true,
				expectPanic:              true,
			},
		}, {
			UID:      10,
			TestType: testDescriptionString,
			Description: `
Description.String() should work properly when:
1. Synopsis is badly given.
2. Body is not given.
3. Panic is expected.
`,
			Switches: map[string]bool{
				useLongSynopsis:          true,
				useProperDescriptionText: false,
				expectPanic:              true,
			},
		}, {
			UID:      11,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is set.
2. Digit-led upstream is set.
3. Upstream value is properly given.
4. Dash version upstream is given.
5. Digit-led revision is set.
6. Revision value is properly given.
7. Dash version revision is not given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      true,
				useAlphaLedVersionUpstream: false,
				useProperVersionUpstream:   true,
				useDashVersionUpstream:     true,
				useAlphaLedVersionRevision: false,
				useProperVersionRevision:   true,
				useDashVersionRevision:     false,
				expectPanic:                false,
			},
		}, {
			UID:      12,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is not set.
2. Digit-led upstream is set.
3. Upstream value is properly given.
4. Dash version upstream is given.
5. Digit-led revision is set.
6. Revision value is properly given.
7. Dash version revision is not given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      false,
				useAlphaLedVersionUpstream: false,
				useProperVersionUpstream:   true,
				useDashVersionUpstream:     true,
				useAlphaLedVersionRevision: false,
				useProperVersionRevision:   true,
				useDashVersionRevision:     false,
				expectPanic:                false,
			},
		}, {
			UID:      13,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is set.
2. Digit-led upstream is not set.
3. Upstream value is properly given.
4. Dash version upstream is given.
5. Digit-led revision is set.
6. Revision value is properly given.
7. Dash version revision is not given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      true,
				useAlphaLedVersionUpstream: true,
				useProperVersionUpstream:   true,
				useDashVersionUpstream:     true,
				useAlphaLedVersionRevision: false,
				useProperVersionRevision:   true,
				useDashVersionRevision:     false,
				expectPanic:                true,
			},
		}, {
			UID:      14,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is set.
2. Digit-led upstream is set.
3. Upstream value is badly given.
4. Dash version upstream is given.
5. Digit-led revision is set.
6. Revision value is properly given.
7. Dash version revision is not given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      true,
				useAlphaLedVersionUpstream: true,
				useIllegalVersionUpstream:  true,
				useDashVersionUpstream:     true,
				useAlphaLedVersionRevision: false,
				useProperVersionRevision:   true,
				useDashVersionRevision:     false,
				expectPanic:                true,
			},
		}, {
			UID:      15,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is set.
2. Digit-led upstream is set set.
3. Upstream value is properly given.
4. Dash version upstream is given.
5. Digit-led revision is set.
6. Revision value is not given.
7. Dash version revision is not given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      true,
				useAlphaLedVersionUpstream: false,
				useProperVersionUpstream:   true,
				useDashVersionUpstream:     true,
				useAlphaLedVersionRevision: false,
				useProperVersionRevision:   false,
				useDashVersionRevision:     false,
				expectPanic:                true,
			},
		}, {
			UID:      16,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is set.
2. Digit-led upstream is set.
3. Upstream value is properly given.
4. Dash version upstream is given.
5. Digit-led revision is not set.
6. Revision value is properly given.
7. Dash version revision is not given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      true,
				useAlphaLedVersionUpstream: false,
				useProperVersionUpstream:   true,
				useDashVersionUpstream:     true,
				useAlphaLedVersionRevision: true,
				useProperVersionRevision:   true,
				useDashVersionRevision:     false,
				expectPanic:                true,
			},
		}, {
			UID:      17,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is set.
2. Digit-led upstream is set.
3. Upstream value is properly given.
4. Dash version upstream is given.
5. Digit-led revision is set.
6. Revision value is properly given.
7. Dash version revision is given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      true,
				useAlphaLedVersionUpstream: false,
				useProperVersionUpstream:   true,
				useDashVersionUpstream:     true,
				useAlphaLedVersionRevision: false,
				useProperVersionRevision:   true,
				useDashVersionRevision:     true,
				expectPanic:                true,
			},
		}, {
			UID:      18,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is set.
2. Digit-led upstream is set.
3. Upstream value is not given.
4. Dash version upstream is given.
5. Digit-led revision is set.
6. Revision value is properly given.
7. Dash version revision is not given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      true,
				useAlphaLedVersionUpstream: false,
				useProperVersionUpstream:   false,
				useDashVersionUpstream:     true,
				useAlphaLedVersionRevision: false,
				useProperVersionRevision:   true,
				useDashVersionRevision:     false,
				expectPanic:                true,
			},
		}, {
			UID:      19,
			TestType: testVersionString,
			Description: `
Version.String() should work properly when:
1. Epoch is set.
2. Digit-led upstream is set.
3. Upstream value is properly given.
4. Dash version upstream is not given.
5. Digit-led revision is set.
6. Revision value is not given.
7. Dash version revision is not given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperVersionEpoch:      true,
				useAlphaLedVersionUpstream: false,
				useProperVersionUpstream:   true,
				useDashVersionUpstream:     false,
				useAlphaLedVersionRevision: false,
				useProperVersionRevision:   false,
				useDashVersionRevision:     false,
				expectPanic:                false,
			},
		}, {
			UID:      20,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to ARCH.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSArch:      true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     false,
			},
		}, {
			UID:      21,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Bazaar.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSBazaar:    true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     false,
			},
		}, {
			UID:      22,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to CVS.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSCVS:       true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     false,
			},
		}, {
			UID:      23,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to DARCS.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSDarcs:     true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     false,
			},
		}, {
			UID:      24,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Git.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSGit:       true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     false,
			},
		}, {
			UID:      25,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Mercurial.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSMercurial: true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     false,
			},
		}, {
			UID:      26,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Monotone.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSMonotone:  true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     false,
			},
		}, {
			UID:      27,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Subversion.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSSubversion: true,
				useProperURL:     true,
				useProperBranch:  true,
				useProperPath:    true,
				expectPanic:      false,
			},
		}, {
			UID:      28,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Git.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is not given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSGit:       true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   false,
				expectPanic:     false,
			},
		}, {
			UID:      29,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Git.
2. URL is properly set to HTTP type.
3. Branch is not given.
4. Path is not given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSGit:       true,
				useProperURL:    true,
				useProperBranch: false,
				useProperPath:   false,
				expectPanic:     false,
			},
		}, {
			UID:      30,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Git.
2. URL is properly set to NON-HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useVCSGit:       true,
				useNonHTTPURL:   true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     false,
			},
		}, {
			UID:      31,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Git.
2. URL is not given.
3. Branch is properly given.
4. Path is properly given.
5. Panic is expected.
`,
			Switches: map[string]bool{
				useVCSGit:       true,
				useProperURL:    false,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     true,
			},
		}, {
			UID:      32,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Unknown.
2. URL is properly set to HTTP type.
3. Branch is properly given.
4. Path is properly given.
5. Panic is expected.
`,
			Switches: map[string]bool{
				useVCSUnknown:   true,
				useProperURL:    true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     true,
			},
		}, {
			UID:      33,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Git.
2. URL is properly set to HTTP type.
3. Vcs-Browser URL is set to empty.
4. Branch is properly given.
5. Path is properly given.
6. Panic is expected.
`,
			Switches: map[string]bool{
				useVCSGit:       true,
				useProperURL:    true,
				useEmptyBrowser: true,
				useProperBranch: true,
				useProperPath:   true,
				expectPanic:     true,
			},
		}, {
			UID:      34,
			TestType: testVCSString,
			Description: `
VCS.String() should work properly when:
1. VCS type is set to Git.
2. URL is properly set to empty type.
3. Vcs-Browser URL is set to empty.
4. Branch is properly given.
5. Path is properly given.
6. Panic is expected.
`,
			Switches: map[string]bool{
				useVCSGit:        true,
				useEmptyURL:      true,
				useProperBrowser: true,
				useProperBranch:  true,
				useProperPath:    true,
				expectPanic:      true,
			},
		}, {
			UID:      35,
			TestType: testTestsuiteString,
			Description: `
Testsuite.String() should work properly when:
1. Paths is properly given.
2. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperPath: true,
				expectPanic:   false,
			},
		}, {
			UID:      36,
			TestType: testTestsuiteString,
			Description: `
Testsuite.String() should work properly when:
1. Paths is badly given.
2. Panic is expected.
`,
			Switches: map[string]bool{
				useGhostPath: true,
				expectPanic:  true,
			},
		}, {
			UID:      37,
			TestType: testTestsuiteString,
			Description: `
Testsuite.String() should work properly when:
1. Paths is emptiedly given.
2. Panic is expected.
`,
			Switches: map[string]bool{
				useEmptyPath: true,
				expectPanic:  true,
			},
		}, {
			UID:      38,
			TestType: testTestsuiteString,
			Description: `
Testsuite.String() should work properly when:
1. Paths is not given.
2. Panic is expected.
`,
			Switches: map[string]bool{
				useProperPath: false,
				expectPanic:   true,
			},
		}, {
			UID:      39,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App Name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Strictly Earlier.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlSTER:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      40,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App Name is not given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Strictly Earlier.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperAppName:         false,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlSTER:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              true,
			},
		}, {
			UID:      41,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is not given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Strictly Earlier.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    false,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlSTER:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      42,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Strictly Earlier.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: false,
				useVERControlSTER:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      43,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is not given.
4. Version revision is properly given.
5. VERControl is given with Strictly Earlier.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: false,
				useProperVersionRevision: true,
				useVERControlSTER:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              true,
			},
		}, {
			UID:      44,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Strictly Earlier.
6. Architecture value is badly given.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlSTER:        true,
				useProperArch:            false,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      45,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Strictly Earlier.
6. Architecture value is properly given.
7. Architecture list badly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlSTER:        true,
				useProperArch:            true,
				useProperArchList:        false,
				expectPanic:              false,
			},
		}, {
			UID:      46,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Earlier or Equal.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEAEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      47,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Exactly Equal.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      48,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Later or Equal.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlLAEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      49,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Strictly Later.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlSTLA:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      50,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Unknown.
6. Architecture value is properly given.
7. Architecture list properly given.
8. Panic is expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlUnknown:     true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              true,
			},
		}, {
			UID:      51,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. VERControl is given with Strictly Later.
6. Architecture value is properly given with any.
7. Architecture list properly given.
8. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlSTLA:        true,
				useAnyArch:               true,
				useProperArchList:        true,
				expectPanic:              false,
			},
		}, {
			UID:      52,
			TestType: testPackageMetaString,
			Description: `
PackageMeta.String() should work properly when:
1. App name is properly given.
2. Version epoch is properly given.
3. Version upstream is properly given.
4. Version revision is properly given.
5. Version is set to nil.
6. VERControl is given with Exact Equal.
7. Architecture value is properly given.
8. Architecture list properly given.
9. Panic is expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useNilVersion:            true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				expectPanic:              true,
			},
		}, {
			UID:      53,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Depends.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListBuildDepends:   true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      54,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Depends.
10. Package list is emptily given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListBuildDepends:   true,
				useEmptyPackageListing:   true,
				expectPanic:              false,
			},
		}, {
			UID:      55,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Depends.
10. Package list is given as nil.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListBuildDepends:   true,
				useProperPackageListing:  false,
				expectPanic:              false,
			},
		}, {
			UID:      56,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Depends Independent.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:            true,
				useProperVersionEpoch:       true,
				useProperVersionUpstream:    true,
				useProperVersionRevision:    true,
				useVERControlEXEQ:           true,
				useProperArch:               true,
				useProperArchList:           true,
				usePkgListBuildDependsIndep: true,
				useProperPackageListing:     true,
				expectPanic:                 false,
			},
		}, {
			UID:      57,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Depends Architecture.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:           true,
				useProperVersionEpoch:      true,
				useProperVersionUpstream:   true,
				useProperVersionRevision:   true,
				useVERControlEXEQ:          true,
				useProperArch:              true,
				useProperArchList:          true,
				usePkgListBuildDependsArch: true,
				useProperPackageListing:    true,
				expectPanic:                false,
			},
		}, {
			UID:      58,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Conflicts.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListBuildConflicts: true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      59,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Conflicts Independent.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:              true,
				useProperVersionEpoch:         true,
				useProperVersionUpstream:      true,
				useProperVersionRevision:      true,
				useVERControlEXEQ:             true,
				useProperArch:                 true,
				useProperArchList:             true,
				usePkgListBuildConflictsIndep: true,
				useProperPackageListing:       true,
				expectPanic:                   false,
			},
		}, {
			UID:      60,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Conflicts Architecture.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:             true,
				useProperVersionEpoch:        true,
				useProperVersionUpstream:     true,
				useProperVersionRevision:     true,
				useVERControlEXEQ:            true,
				useProperArch:                true,
				useProperArchList:            true,
				usePkgListBuildConflictsArch: true,
				useProperPackageListing:      true,
				expectPanic:                  false,
			},
		}, {
			UID:      61,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Pre Depends.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListPreDepends:     true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      62,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Depends.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListDepends:        true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      63,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Recommends.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListRecommends:     true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      64,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Suggests.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListSuggests:       true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      65,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Enhances.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListEnhances:       true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      66,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Breaks.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListBreaks:         true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      67,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Conflicts.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListConflicts:      true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      68,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Provides.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListProvides:       true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      69,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Replaces.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListReplaces:       true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      70,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Unknown.
10. Package list is properly given.
11. Panic is expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListUnknown:        true,
				useProperPackageListing:  true,
				expectPanic:              true,
			},
		}, {
			UID:      71,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is badly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Build Conflicts.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: false,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListBuildConflicts: true,
				useProperPackageListing:  true,
				expectPanic:              true,
			},
		}, {
			UID:      72,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				expectPanic:              false,
			},
		}, {
			UID:      73,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   true,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				expectPanic:              true,
			},
		}, {
			UID:      74,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is set.
8. License without comment is not set.
9. License without body is not set.
10. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     true,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				expectPanic:              true,
			},
		}, {
			UID:      75,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is set.
9. License without body is not set.
10. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: true,
				useLicenseWithoutBody:    false,
				expectPanic:              false,
			},
		}, {
			UID:      76,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is set.
10. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    true,
				expectPanic:              false,
			},
		}, {
			UID:      77,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are not given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        false,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				expectPanic:              true,
			},
		}, {
			UID:      78,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. License is set to use 'and'.
11. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				useLicenseTagAnd:         true,
				expectPanic:              true,
			},
		}, {
			UID:      79,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. License is set to use 'or'.
11. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				useLicenseTagOr:          true,
				expectPanic:              true,
			},
		}, {
			UID:      80,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. License is set to use 'symbol'.
11. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				useLicenseTagSymbol:      true,
				expectPanic:              true,
			},
		}, {
			UID:      81,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is badly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. License is set to use 'symbol'.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          false,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				useLicenseTagSymbol:      false,
				expectPanic:              false,
			},
		}, {
			UID:      82,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is given.
2. Email is given.
3. Year is given.
4. Entity is badly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. License with files set to nil.
11. License is set to use 'symbol'.
12. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            true,
				useProperEntity:          false,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				useLicenseWithNilFiles:   true,
				useLicenseTagSymbol:      false,
				expectPanic:              true,
			},
		}, {
			UID:      83,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is properly given.
11. Copyright Disclaimer is properly given.
12. App Name is properly given.
13. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  false,
			},
		}, {
			UID:      84,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is badly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is properly given.
11. Copyright Disclaimer is properly given.
12. App name is properly provided.
13. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              false,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  true,
			},
		}, {
			UID:      85,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with "".
10. Copyright Source is properly given.
11. Copyright Disclaimer is properly given.
12. App name is properly provided.
13. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        false,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  true,
			},
		}, {
			UID:      86,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is badly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is properly given.
11. Copyright Disclaimer is properly given.
12. App name is properly provided.
13. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               false,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  true,
			},
		}, {
			UID:      87,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is badly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is properly given.
11. Copyright Disclaimer is properly given.
12. App name is properly provided.
13. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:                false,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  true,
			},
		}, {
			UID:      88,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is without tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is properly given.
11. Copyright Disclaimer is properly given.
12. App name is properly provided.
13. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         true,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            false,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  true,
			},
		}, {
			UID:      89,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are given as nil.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is properly given.
11. Copyright Disclaimer is properly given.
12. App name is properly provided.
13. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            false,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  false,
			},
		}, {
			UID:      90,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is badly given.
11. Copyright Disclaimer is properly given.
12. App name is properly provided.
13. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     false,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  false,
			},
		}, {
			UID:      91,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is properly given.
11. Copyright Disclaimer is badly given.
12. App name is properly provided.
13. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: false,
				useProperAppName:             true,
				expectPanic:                  false,
			},
		}, {
			UID:      92,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. Upstream License tag is without comment.
8. License Entities are properly given.
9. Licenses are properly created.
10. Copyright Format is given with 1.0.
11. Copyright Source is properly given.
12. Copyright Disclaimer is badly given.
13. App name is properly provided.
14. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     true,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: false,
				useProperAppName:             true,
				expectPanic:                  false,
			},
		}, {
			UID:      93,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is badly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. Upstream License tag is without comment.
8. License Entities are properly given.
9. Licenses are properly created.
10. Copyright Format is given with 1.0.
11. Copyright Source is properly given.
12. Copyright Disclaimer is badly given.
13. App name is properly provided.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                true,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     true,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             true,
				expectPanic:                  false,
			},
		}, {
			UID:      94,
			TestType: testCopyrightString,
			Description: `
Copyright.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Upstream Entity is properly given.
5. Upstream License Entities are properly given.
6. Upstream License tag is with tag.
7. License Entities are properly given.
8. Licenses are properly created.
9. Copyright Format is given with 1.0.
10. Copyright Source is properly given.
11. Copyright Disclaimer is properly given.
12. App Name is badly given.
13. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperEntities:            true,
				useLicenseWithoutComment:     false,
				useLicenseWithoutTag:         false,
				useLicenseTagSymbol:          false,
				useLicenseTagOr:              false,
				useLicenseTagAnd:             false,
				useProperLicenses:            true,
				useCopyrightFormat1p0:        true,
				useProperCopyrightSource:     true,
				useProperCopyrightDisclaimer: true,
				useProperAppName:             false,
				expectPanic:                  true,
			},
		}, {
			UID:      95,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               false,
			},
		}, {
			UID:      96,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is badly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             false,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      97,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is badly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            false,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      98,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is badly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             true,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               false,
			},
		}, {
			UID:      99,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is badly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           false,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      100,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is badly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          false,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      101,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is badly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  false,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      102,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is badly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  false,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               false,
			},
		}, {
			UID:      103,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is badly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
12. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        false,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      104,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is badly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           false,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      105,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is badly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: false,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      106,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to medium.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyMedium: true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               false,
			},
		}, {
			UID:      107,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to high.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyHigh:   true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               false,
			},
		}, {
			UID:      108,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to emergency.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:                true,
				useProperEmail:               true,
				useProperYear:                false,
				useProperEntity:              true,
				useProperAppName:             true,
				useProperVersionUpstream:     true,
				useProperVersionRevision:     true,
				useProperTimestamp:           true,
				useChangelogUrgencyEmergency: true,
				useProperDistro:              true,
				useProperChangelogChanges:    true,
				useNilVersion:                false,
				useProperPath:                true,
				expectPanic:                  false,
			},
		}, {
			UID:      109,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to critical.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:               true,
				useProperEmail:              true,
				useProperYear:               false,
				useProperEntity:             true,
				useProperAppName:            true,
				useProperVersionUpstream:    true,
				useProperVersionRevision:    true,
				useProperTimestamp:          true,
				useChangelogUrgencyCritical: true,
				useProperDistro:             true,
				useProperChangelogChanges:   true,
				useNilVersion:               false,
				useProperPath:               true,
				expectPanic:                 false,
			},
		}, {
			UID:      110,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to unknown.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
12. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:              true,
				useProperEmail:             true,
				useProperYear:              false,
				useProperEntity:            true,
				useProperAppName:           true,
				useProperVersionUpstream:   true,
				useProperVersionRevision:   true,
				useProperTimestamp:         true,
				useChangelogUrgencyUnknown: true,
				useProperDistro:            true,
				useProperChangelogChanges:  true,
				useNilVersion:              false,
				useProperPath:              true,
				expectPanic:                true,
			},
		}, {
			UID:      111,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             true,
				useProperPath:             false,
				expectPanic:               true,
			},
		}, {
			UID:      112,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is emptily given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useEmptyTimestamp:         true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPath:             true,
				expectPanic:               true,
			},
		}, {
			UID:      113,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is emptily given.
12. Version is not set to nil.
13. Path is properly given.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useProperYear:            false,
				useProperEntity:          true,
				useProperAppName:         true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useProperTimestamp:       true,
				useChangelogUrgencyLow:   true,
				useProperDistro:          true,
				useEmptyChangelogChanges: true,
				useNilVersion:            false,
				useProperPath:            true,
				expectPanic:              true,
			},
		}, {
			UID:      114,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given as directory.
14. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPathDirectory:    true,
				expectPanic:               true,
			},
		}, {
			UID:      115,
			TestType: testChangelogString,
			Description: `
Changelog.String() should work properly when:
1. Entity Name is properly given.
2. Entity Email is properly given.
3. Entity Year is properly given.
4. Maintainer Entity is properly given.
5. App Name is properly given.
6. Version upstream is properly given.
7. Version revision is properly given.
8. Timestamp is properly given.
9. Changelog urgency is set to low.
10. Changelog distribution is properly given.
11. Changelog changes is properly given.
12. Version is not set to nil.
13. Path is properly given.
14. StatFx to simulate error is set.
15. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperAppName:          true,
				useProperVersionUpstream:  true,
				useProperVersionRevision:  true,
				useProperTimestamp:        true,
				useChangelogUrgencyLow:    true,
				useProperDistro:           true,
				useProperChangelogChanges: true,
				useNilVersion:             false,
				useProperPathDirectory:    false,
				simulateStatFxError:       true,
				expectPanic:               true,
			},
		}, {
			UID:      116,
			TestType: testPackageListString,
			Description: `
PackageList.String() should work properly when:
1. Package Name is properly given.
2. Package Version epoch is properly given.
3. Package Version upstream is properly given.
4. Package Version revision is properly given.
6. Package VERControl is given with Exact Equal.
7. Package Architecture value is properly given.
8. Package Architecture list properly given.
9. Package List Type is set to Built Using.
10. Package list is properly given.
11. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperAppName:         true,
				useProperVersionEpoch:    true,
				useProperVersionUpstream: true,
				useProperVersionRevision: true,
				useVERControlEXEQ:        true,
				useProperArch:            true,
				useProperArchList:        true,
				usePkgListBuiltUsing:     true,
				useProperPackageListing:  true,
				expectPanic:              false,
			},
		}, {
			UID:      117,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      118,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is badly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             false,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      119,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is badly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            false,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      120,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is badly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             true,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      121,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is badly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           false,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      122,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is emptly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useEmptyPackageList:       true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      123,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is badly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useFaultyPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      124,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is badly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  false,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      125,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is badly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              false,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      126,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      false,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      127,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is badly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              false,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      128,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is badly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        false,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      129,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is badly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          false,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      130,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is badly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          false,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      131,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is badly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useFaultyStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      132,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to no.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRNo:                  true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      133,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to unknown.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRUnknown:             true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      134,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to none.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           false,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      135,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to UDEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeUDEB:        true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      136,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to none.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         false,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      137,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to Unknown.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeUnknown:     true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      138,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to none.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         false,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      139,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Important.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityImportant:      true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      140,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Standard.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityStandard:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      141,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Optional.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityOptional:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      142,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to unknown.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityUnknown:        true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      143,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to none.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       false,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      144,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is emptily given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             false,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      145,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is not set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          false,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      146,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is not set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        false,
				expectPanic:               false,
			},
		}, {
			UID:      147,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is not given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useNilPackageList:         true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      148,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is not given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useNilVersion:             true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      149,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is weirdly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useUnknownArch:            true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      150,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is faultily given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useFaultyDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      151,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is faultily given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useFaultyVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      152,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is faultily given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useFaultyTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      153,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to unknown.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeUnknown:     true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      154,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      155,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are emptily given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useEmptyUploaders:         true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      156,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are not given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        false,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      157,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to bin-target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are faultily given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useFaultyUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               true,
			},
		}, {
			UID:      158,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is properly given.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to compliant custom.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useProperPackageList:      true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRCustom:              true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      159,
			TestType: testLicenseString,
			Description: `
License.String() should work properly when:
1. Name is properly given.
2. Email is properly given.
3. Year is badly given.
4. Entity is properly set.
5. Entities are properly given.
6. License without files is not set.
7. License without tag is not set.
8. License without comment is not set.
9. License without body is not set.
10. Panic is expected.
`,
			Switches: map[string]bool{
				useProperName:            true,
				useProperEmail:           true,
				useBadYear:               true,
				useProperEntity:          true,
				useProperEntities:        true,
				useLicenseWithoutFiles:   false,
				useLicenseWithoutTag:     false,
				useLicenseWithoutComment: false,
				useLicenseWithoutBody:    false,
				expectPanic:              true,
			},
		}, {
			UID:      160,
			TestType: testControlString,
			Description: `
1. Control.String() should work properly when:
2. Maintainer's Name is properly given.
3. Maintainer's Email is properly given.
4. Maintainer's Year is properly given.
5. Maintainer's data is properly given.
6. Package List is given with nil content.
7. Version is properly given.
8. Homepage is properly given.
9. Description is properly given.
10. VCS is properly given.
11. Testsuite is properly given.
12. App name is properly given.
13. Section is properly given.
14. Standards Version is properly given.
15. Package Type set to DEB.
16. Rules Requires Root is set to Bin Target.
17. Priority is set to Required.
18. Architecture is properly given.
19. Uploaders are properly given.
20. Essential Flag is set.
21. Build Source Flag is set.
22. Panic is not expected.
`,
			Switches: map[string]bool{
				useProperName:             true,
				useProperEmail:            true,
				useProperYear:             false,
				useProperEntity:           true,
				useNilContentPackageList:  true,
				useProperVersionEpoch:     true,
				useProperVersionUpstream:  true,
				useProperURL:              true,
				useProperDescription:      true,
				useProperVCS:              true,
				useProperTestsuite:        true,
				useProperAppName:          true,
				useProperSection:          true,
				useProperStandardsVersion: true,
				usePackageTypeDEB:         true,
				useRRRBinTarget:           true,
				usePriorityRequired:       true,
				useProperArch:             true,
				useProperUploaders:        true,
				useEssentialFlag:          true,
				useBuildSourceFlag:        true,
				expectPanic:               false,
			},
		}, {
			UID:      161,
			TestType: testSourceSanitize,
			Description: `
Source.Sanitize() should work properly when:
1. Format is properly given as Native 3.0.
2. Local Options are properly given.
3. Options are properly given.
4. Error is not expected.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useSourceFormatNative3p0:    true,
				useProperSourceLocalOptions: true,
				useProperSourceOptions:      true,
				expectError:                 false,
				expectPanic:                 false,
			},
		}, {
			UID:      162,
			TestType: testSourceSanitize,
			Description: `
Source.Sanitize() should work properly when:
1. Format is properly given as Quilt 3.0.
2. Local Options are properly given.
3. Options are properly given.
4. Error is not expected.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useSourceFormatNative3p0:    true,
				useProperSourceLocalOptions: true,
				useProperSourceOptions:      true,
				expectError:                 false,
				expectPanic:                 false,
			},
		}, {
			UID:      163,
			TestType: testSourceSanitize,
			Description: `
Source.Sanitize() should work properly when:
1. Format is badly given.
2. Local Options are properly given.
3. Options are properly given.
4. Error is expected.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useSourceFormatNative3p0:    false,
				useProperSourceLocalOptions: true,
				useProperSourceOptions:      true,
				expectError:                 true,
				expectPanic:                 false,
			},
		}, {
			UID:      164,
			TestType: testSourceSanitize,
			Description: `
Source.Sanitize() should work properly when:
1. Format is properly given as Native 3.0.
2. Local Options are badly given.
3. Options are properly given.
4. Error is not expected.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useSourceFormatNative3p0:    true,
				useProperSourceLocalOptions: false,
				useProperSourceOptions:      true,
				expectError:                 false,
				expectPanic:                 false,
			},
		}, {
			UID:      165,
			TestType: testSourceSanitize,
			Description: `
Source.Sanitize() should work properly when:
1. Format is properly given as Native 3.0.
2. Local Options are properly given.
3. Options are badly given.
4. Error is not expected.
5. Panic is not expected.
`,
			Switches: map[string]bool{
				useSourceFormatNative3p0:    true,
				useProperSourceLocalOptions: true,
				useProperSourceOptions:      false,
				expectError:                 false,
				expectPanic:                 false,
			},
		},
	}
}
