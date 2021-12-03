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
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/styler"
)

// Entity is the merged data structure from Person and Organization.
//
// The merging observes common fields mainly used for legal identification
// without needing to differentiate a Person and an Organization.
//
// Refers to:
//   1. https://schema.org/Person
//   2. https://schema.org/Organization
type Entity struct {
	// Name is the name of the Entity.
	Name string

	// Aliases are another names of the Entity.
	Aliases []string

	// ID is the unique identifier of the Entity.
	//
	// Example:
	//   1. Nationally Registered Identification Card
	//   2. Social Security Number
	//   3. Business Registration Number
	ID map[string]string

	// CopyrightTime is the timestamp where the entity had asserted.
	CopyrightTime *Timestamp

	// JointTime is the timestamp where the entity had joint.
	JointTime *Timestamp

	// Addresses are the physical location for interacting with the Entity.
	Addresses []string

	// Email is the email address for contacting the Entity.
	Email []string

	// Telephone is the phone number for reaching the Entity directly.
	Telephone []string

	// Fax is the phone number dedicated for faxing to Entity.
	Fax []string

	// Website is the website URL of the Entity.
	Website []string
}

func (me *Entity) String() (s string) {
	s += "---------------------------------------------------------------\n"
	s += styler.PortraitKV("Name", me.Name)
	s += me.strTimestamp("Joint Since", me.JointTime)
	s += me.strTimestamp("Copyright Since", me.CopyrightTime)
	s += styler.PortraitKArray("Aliases", me.Aliases)
	s += styler.PortraitKMap("ID", me.ID)
	s += styler.PortraitKArray("Addresses", me.Addresses)
	s += styler.PortraitKArray("Email", me.Email)
	s += styler.PortraitKArray("Telephone", me.Telephone)
	s += styler.PortraitKArray("Fax", me.Fax)
	s += styler.PortraitKArray("Website", me.Website)
	s = strings.TrimRight(s, "\n") + "\n"
	s += "---------------------------------------------------------------\n"

	return s
}

func (me *Entity) strTimestamp(title string, x *Timestamp) string {
	if x == nil {
		return ""
	}

	return styler.PortraitKV(title, x.String())
}
