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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
)

type TOMLChangelog struct {
	Entries   []string
	LineBreak string
	Regex     string
	CMD       []*TOMLAction
}

type TOMLMetadata struct {
	Name        string
	Description string
	Type        string
}

func (me *TOMLMetadata) Sanitize(path string) (err error) {
	if me.Name == "" {
		return fmt.Errorf("%s: Name for %s",
			ERROR_PUBLISH_METADATA_MISSING,
			path,
		)
	}

	return nil
}

type TOMLDependency struct {
	Name      string
	Condition string
	Command   string
	Type      commander.ActionID
}

type TOMLAction struct {
	Name       string
	Type       commander.ActionID
	Location   string
	Source     string
	Target     string
	Save       string
	SaveRegex  string
	ToSTDERR   string
	ToSTDOUT   string
	Condition  []string
	SaveStderr bool
}

func (base *TOMLAction) Sanitize() (err error) {
	if base.Name == "" {
		return fmt.Errorf("%s: %s",
			ERROR_COMMAND_BAD,
			"Command.Name is empty",
		)
	}

	if base.Type == "" {
		return fmt.Errorf("%s: %s",
			ERROR_COMMAND_BAD,
			"Command.Type is empty",
		)
	}

	if len(base.Condition) == 0 {
		return fmt.Errorf("%s: %s",
			ERROR_COMMAND_BAD,
			"Command.Condition is empty",
		)
	}

	return nil
}

func (base *TOMLAction) ParseExec(in string) (out string) {
	in = strings.TrimRight(in, "\r\n")

	if base.SaveRegex != "" {
		r, _ := regexp.Compile(base.SaveRegex)

		list := r.FindStringSubmatch(in)
		for i, s := range list {
			if i == 0 {
				continue
			}

			out += s
		}
	} else {
		out = in
	}

	return out
}

type TOMLChecksum struct {
	Type   string
	Format string
	Value  string
}

type TOMLPackage struct {
	Files        map[string]string
	Name         string
	Changelog    string
	Source       string
	Target       string
	OS           []string
	Arch         []string
	Distribution []string
	BuildSource  bool
}

type TOMLRelease struct {
	Data     *TOMLReleaseData
	Packages map[string]*TOMLPackage
	Target   string
	Checksum string
}

type TOMLReleaseData struct {
	Path   string
	Format string
}

type TOMLSource struct {
	Checksum    *TOMLChecksum
	Headers     map[string]string
	Archive     string
	Format      string
	URL         string
	Method      string
	Destination string
}

func (base *TOMLSource) Merge(in *TOMLSource) {
	if in.Archive != "" {
		base.Archive = in.Archive
	}

	if in.Format != "" {
		base.Format = in.Format
	}

	if in.URL != "" {
		base.URL = in.URL
	}

	if in.Method != "" {
		base.Method = in.Method
	}

	if len(in.Headers) > 0 {
		if len(base.Headers) == 0 {
			base.Headers = map[string]string{}
		}

		for k, v := range in.Headers {
			base.Headers[k] = v
		}
	}

	base.mergeChecksum(in)
}

func (base *TOMLSource) mergeChecksum(in *TOMLSource) {
	if in.Checksum == nil {
		return
	}

	if base.Checksum == nil {
		base.Checksum = &TOMLChecksum{}
	}

	// if both checksum value is not available, disable checksum entirely
	if base.Checksum.Value == "" && in.Checksum.Value == "" {
		base.Checksum = nil
		return
	}

	if in.Checksum.Value != "" {
		base.Checksum.Value = in.Checksum.Value
	}

	if in.Checksum.Type != "" {
		base.Checksum.Type = in.Checksum.Type
	}

	if in.Checksum.Format != "" {
		base.Checksum.Format = in.Checksum.Format
	}
}

type TOMLSourceConfig map[string]string

func AcceptTOML(path string, info os.FileInfo, err error) (bool, error) {
	// return if err occurred
	if err != nil {
		return false, fmt.Errorf("%s: %s", ERROR_TOML_PARSE_FAILED, err)
	}

	// ensures we only accepts regular file with .toml extension
	if filepath.Ext(path) != EXTENSION_TOML || !info.Mode().IsRegular() {
		return false, nil
	}

	// accepted
	return true, nil
}
