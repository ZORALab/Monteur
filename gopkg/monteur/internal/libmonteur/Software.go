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
	"strconv"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/styler"
)

// Software is the software information consolidated in a common data structure.
type Software struct {
	// Name is the name of the software.
	Name string

	// Command is the terminal command name of the software.
	Command string

	// ID is the store identification of the software.
	ID string

	// Version is the version instance of the Software.
	Version string

	// Category is the main category where the software belongs.
	Category string

	// Abstract is the short description summary.
	//
	// Abstract cannot be longer than 80 characters.
	Abstract string

	// Description is the long description.
	Description string

	// Suite is the name of the software where it belongs.
	//
	// Example: 'Microsoft Excel' belongs to 'Microsoft Office' Suite.
	Suite string

	// Contact is the general contact channel.
	Contact *Entity

	// Help is the software help data.
	Help *SoftwareHelp

	// Copyrights is the list of legal copyright data for the software.
	Copyrights []*Copyright

	// Debian is the .deb packaging data
	Debian *DEB

	// Time is the current operating timestamp
	Time *Timestamp

	// Maintainers are the list of Entity maintaining the software.
	Maintainers map[string]*Entity

	// Contributors are the list of Entity maintaining the software.
	Contributors map[string]*Entity

	// Sponsors are the list of Entity financially sponsoring the software.
	Sponsors map[string]*Entity

	// OS is the supported operating system.
	//
	// Example: 'Windows 7', 'OSX 10.6', 'Android 1.6'.
	OS []string

	// ARCH is the supported CPU Architecture.
	//
	// Example: `amd64`, `arm64`, `i386`, `arm`, `riscv64`.
	ARCH []string
}

func (me *Software) String() (s string) {
	s += styler.PortraitKV("Timestamp", me.Time.String())
	s += styler.PortraitKV("Name", me.Name)
	s += styler.PortraitKV("Command", me.Command)
	s += styler.PortraitKV("ID", me.Command)
	s += styler.PortraitKV("Version", me.Version)
	s += styler.PortraitKV("Category", me.Category)
	s += styler.PortraitKV("Suite", me.Suite)
	s += styler.PortraitKV("Abstract", me.Abstract)
	s += styler.PortraitKV("Description", me.Description)
	s += styler.PortraitKV("Contact", me.Contact.String())
	s += me.strAppEntities("Maintainers", me.Maintainers)
	s += me.strAppEntities("Contributors", me.Contributors)
	s += me.strAppEntities("Sponsors", me.Sponsors)
	s += me.strAppCopyrights("Copyrights", me.Copyrights)
	s += me.Debian.String()
	s += me.Help.String()

	return s
}

func (me *Software) strAppEntities(title string,
	list map[string]*Entity) (s string) {
	s = strings.ToUpper(title) + "\n"

	for _, v := range list {
		s += v.String()
	}
	s += "\n"

	return s
}

func (me *Software) strAppCopyrights(title string,
	list []*Copyright) (s string) {
	s = strings.ToUpper(title) + "\n"

	for i, v := range list {
		s += "[ DOC" + strconv.Itoa(i+1) + " ]\n"
		s += v.String()
	}
	s += "\n"

	return s
}
