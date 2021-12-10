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

// Data is the full Debian specific package metadata for archiving purposes.
type Data struct {
	// Control is the control data saved into DEBIAN/control
	Control *Control

	// Copyright is the copyright data saved into DEBIAN/copyright
	Copyright *Copyright

	// Changelog are the changelog data saved into DEBIAN/changelog
	Changelog *Changelog

	// Source is the deb file format saved into DEBIAN/source directory.
	Source *Source

	// Manpage is the manpage data saved into DEBIAN/app.manpages
	Manpage map[string]string

	// Install is the installation list of files saved into DEBIAN/install
	Install map[string]string

	// Rules are the package unpack operations data saved into DEBIAN/rules
	Rules string

	// Compat is the debhelper compatibility level saved into DEBIAN/compat
	Compat uint
}
