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

//nolint:lll
const (
	changelogFile = `testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogBadHeaderFile = `testapp (5:1.0.0-1alpha) stable: urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogBadPackageHeaderFile = `testapp [5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogBadVersionHeaderFile = `testapp (5:1.0.0-1alpha] stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogChangesLeadFile = `  * changed feature A
testapp (5:1.0.0-1alpha) stable: urgency=low

  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogDoubleHeaderFile = `testapp (5:1.0.0-1alpha) stable; urgency=low
testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogSigHeaderFile = `testapp (5:1.0.0-1alpha) stable; urgency=low
 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogSigLeadFile = ` -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300
testapp (5:1.0.0-1alpha) stable: urgency=low

  * changed feature A
  * changed feature B`
	changelogTailingHeaderFile = `testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300
testapp (1.0.0) stable; urgency=low`
	changelogTailingChangesFile = `testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300
  * changed feature A
`
	changelogTailingSigFile = `testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300
 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogBadVersionEpochFile = `testapp (a:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogBadVersionUpstreamFile = `testapp (5:v1.0.0-alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogMissingDistroFile = `testapp (5:1.0.0-1alpha) ; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogBadNameFile = `testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov [alexi@testing.corp>  Tue, 25 May 2021 16:01:22 +0300`
	changelogBadEmailFile = `testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp]  Tue, 25 May 2021 16:01:22 +0300`
	changelogBadEmailDelimiterFile = `testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp> Tue, 25 May 2021 16:01:22 +0300`
	changelogBadTimestampFile = `testapp (5:1.0.0-1alpha) stable; urgency=low

  * changed feature A
  * changed feature B

 -- Alexi Romanov <alexi@testing.corp>  ReadThisAlready!`
)
