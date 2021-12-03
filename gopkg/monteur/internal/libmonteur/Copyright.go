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

// Copyright is a created data structure to hold all copyright information.
type Copyright struct {
	// Name is the name of the Copyright/License.
	Name string

	// ID is the unique identifier of the Copyright/License.
	ID string

	// Notice is the short license notice of the Copyright without holders.
	Notice string

	// Text is the full license body of the Copyright without holders.
	Text string

	// Materials are the affected targets by the Copyright/License.
	Materials []string
}

func (me *Copyright) String() (s string) {
	s = styler.PortraitKV("Name", me.Name)
	s += styler.PortraitKV("ID", me.ID)
	s += styler.PortraitKArray("Materials", me.Materials)
	s += styler.PortraitKV("Notice", me.Notice)
	s += styler.PortraitKV("Text", me.Text)

	return s
}
