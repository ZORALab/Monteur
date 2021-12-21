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
)

// SourceFormatType is the list of debian/source/format data.
type SourceFormatType string

const (
	SOURCE_FORMAT_NATIVE_3_0 SourceFormatType = "3.0 (native)"
	SOURCE_FORMAT_QUILT_3_0  SourceFormatType = "3.0 (quilt)"
)

// Source are the data for debian/source directory.
//
// More info:
//   https://www.debian.org/doc/manuals/maint-guide/dother.en.html
type Source struct {
	// Format are the debian/source/format string data.
	//
	// This field is **MANDATORY**.
	Format SourceFormatType

	// LocalOptions are the debian/source/local-options data.
	//
	// This field is usually used for quilt version control system. Leaving
	// it empty shall not generate the debian/source/local-options file.
	LocalOptions string

	// Options are the debian/source/options data.
	//
	// This field is used by dh_autoreconf. More info:
	//   https://www.debian.org/doc/manuals/maint-guide/dother.en.html
	//   https://www.debian.org/doc/manuals/maint-guide/dreq.en.html#customrules
	Options string

	// LintianOverrides are the debian/source/lintian-overrides data.
	//
	// This field is used for overriding false-positive reports.
	LintianOverrides string
}

// Sanitize is to check all Source data are compliant to .deb strict format.
//
// It shall return error if any data is found not compliant.
func (me *Source) Sanitize() (err error) {
	switch me.Format {
	case SOURCE_FORMAT_NATIVE_3_0:
	case SOURCE_FORMAT_QUILT_3_0:
	default:
		return fmt.Errorf("%s: '%s'",
			ERROR_SOURCE_FORMAT_UNKNOWN,
			me.Format,
		)
	}

	return nil
}
