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

package libmonteur

import (
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/styler"
)

// SoftwareHelp is the help data for generating all help files.
type SoftwareHelp struct {
	// Manpage is the list of UNIX manual pages data.
	Manpage *SoftwareManpage

	// Command is the app program to call for help.
	Command string

	// Description is a general description for `Command`.
	Description string

	// Resources is the data for extended assistances.
	Resources string
}

func (me *SoftwareHelp) String() (s string) {
	s = styler.PortraitKV("Help.Command", me.Command)
	s += styler.PortraitKV("Help.Description", me.Description)
	s += styler.PortraitKV("Help.Resources", me.Resources)

	if me.Manpage != nil {
		s += me.Manpage.String()
	}

	return s
}

// SoftwareManpage is the UNIX manual page document data.
//
// These pages' data will be formatted by the Software data.
type SoftwareManpage struct {
	Lv1 string // general commands
	Lv2 string // system calls
	Lv3 string // library functions
	Lv4 string // special files (devices, drivers, etc)
	Lv5 string // file formats and conventions
	Lv6 string // games and screensavers
	Lv7 string // michellaneous
	Lv8 string // system admin commands and daemons
}

func (me *SoftwareManpage) String() (s string) {
	s = styler.PortraitKV("Help.Manpage.Lv1", me.Lv1)
	s += styler.PortraitKV("Help.Manpage.Lv2", me.Lv2)
	s += styler.PortraitKV("Help.Manpage.Lv3", me.Lv3)
	s += styler.PortraitKV("Help.Manpage.Lv4", me.Lv4)
	s += styler.PortraitKV("Help.Manpage.Lv5", me.Lv5)
	s += styler.PortraitKV("Help.Manpage.Lv6", me.Lv6)
	s += styler.PortraitKV("Help.Manpage.Lv7", me.Lv7)
	s += styler.PortraitKV("Help.Manpage.Lv8", me.Lv8)

	return s
}
