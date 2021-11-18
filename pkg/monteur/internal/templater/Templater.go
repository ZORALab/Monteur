// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by datalicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templater

import (
	"bytes"
)

// String is to template a given string with the variables.
//
// Should the text/template return an `error`, the `out` value is always set to
// the original `text` value.
//
// Only a no `error` situation (`nil`) shall the `out` value shall be the
// processed version.
func String(text string, variables interface{}) (out string, err error) {
	// initialize varibles
	out = text

	// initialize text template
	t := textTemplate("Value")

	// parse the text input for processing
	t, err = t.Parse(text)
	if err != nil {
		return out, err //nolint:wrapcheck
	}

	// process the string
	var b bytes.Buffer
	if err := t.Execute(&b, variables); err != nil {
		return out, err //nolint:wrapcheck
	}
	out = b.String()

	return out, nil
}
