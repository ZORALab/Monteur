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

package styler

import (
	"strings"
)

// BoxString is to box a given content with various bordered unicode.
//
// If the content is empty, the returning string shall be emptied. The available
// `borderType` are defined in the constant list of "Border Types". Default is
// `BORDER_SINGLE`.
func BoxString(content string, borderType uint) string {
	if content == "" {
		return ""
	}

	l := len(content)
	tl := "┌"
	tr := "┐"
	bl := "└"
	br := "┘"
	h := "─"
	v := "│"

	switch borderType {
	case BORDER_DOUBLE:
		tl = "╔"
		tr = "╗"
		bl = "╚"
		br = "╝"
		h = "═"
		v = "║"
	case BORDER_BOLD:
		tl = "┏"
		tr = "┓"
		bl = "┗"
		br = "┛"
		h = "━"
		v = "┃"
	}

	return tl + strings.Repeat(h, l) + tr + "\n" +
		v + content + v + "\n" +
		bl + strings.Repeat(h, l) + br + "\n"
}
