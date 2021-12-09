// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, VCS 2.0 (the "License");
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
	"strings"

	"gitlab.com/zoralab/cerigo/strings/strhelper"
)

func foldSPDXText(text string) (s string) {
	// trim trailing whitespaces before processing
	text = strings.TrimRight(text, "\n")
	text = strings.TrimRight(text, "\r")
	text = strings.TrimRight(text, " ")

	// split paragraphs
	paragraphs := strings.Split(text, "\n\n")

	// style each paragraph
	styler := strhelper.Styler{}
	for i, paragraph := range paragraphs {
		list := styler.ContentWrap(paragraph, 79, "\n")

		if i != 0 {
			s += " .\n"
		}

		for _, v := range list {
			s += " " + v
		}
	}

	// remove double line spacing
	s = strings.ReplaceAll(s, " \n \n", "\n")

	// trim trailing whitespaces before processing
	s = strings.TrimRight(s, "\n")
	s = strings.TrimRight(s, "\r")
	s = strings.TrimRight(s, " ")

	return s
}
