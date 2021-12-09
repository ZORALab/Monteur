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
	"fmt"
	"strconv"
)

// Entity is the legal entity data for person or organization.
type Entity struct {
	// Name is the legally registered full name of the entity.
	//
	// This field is **COMPULSORY**.
	Name string

	// Email is the contact email address.
	//
	// This field is **COMPULSORY**.
	Email string

	// Year is where the copyright year is applied to.
	//
	// If Year is more than 0, it will be prefixed before the Name.
	Year int
}

// Sanitize is to ensure Entity's data is compliant.
func (me *Entity) Sanitize() (err error) {
	if me.Year < 0 {
		me.Year = 0
	}

	if me.Name == "" {
		return fmt.Errorf("%s: ''", ERROR_ENTITY_NAME_BAD)
	}

	if me.Email == "" {
		return fmt.Errorf("%s: ''", ERROR_ENTITY_EMAIL_BAD)
	}

	return nil
}

// Tag generates the Entity's name and email string.
//
// If year is provided, it shall be prefixed. Some example outputs with year
// and without year.
//   John Smith <john.smith@testing.email>
//
// This function will panic if Entity is not sanitized prior use and an error
// has occurred.
func (me *Entity) Tag() string {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	return me.Name + " <" + me.Email + ">"
}

// String generates the Entity output string.
//
// If year is provided, it shall be prefixed. Some example outputs with year
// and without year.
//   2021 John Smith <john.smith@testing.email>
//   John Smith <john.smith@testing.email>
//
// This function will panic if Entity is not sanitized prior use and an error
// has occurred.
func (me *Entity) String() (s string) {
	s = me.Tag()

	if me.Year > 0 {
		s = strconv.Itoa(me.Year) + " " + s
	}

	return s
}
