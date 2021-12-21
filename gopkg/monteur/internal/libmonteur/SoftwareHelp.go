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
	Manpage map[string]string

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

	if len(me.Manpage) == 0 {
		s += styler.PortraitKV("Help.Manpage", "nil")
	} else {
		for k, v := range me.Manpage {
			s += styler.PortraitKV("Help.Manpage."+k, v)
		}
	}

	return s
}
