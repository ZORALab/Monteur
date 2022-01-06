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

package archiver

// FormatID generates the Data format into DataPath
type FormatID uint

const (
	FORMAT_NONE FormatID = iota
	FORMAT_TOML
	FORMAT_TXT
	FORMAT_CSV
)

const (
	_EXTENSION_TOML = ".toml"
	_TEMPLATE_TOML  = `[{{- .Index -}}]
Filename = '{{- .Filename -}}'
Hash = '{{- .Hash -}}'
Format = 'Hex'
URL = '{{- .URL -}}'`
)
const (
	_EXTENSION_CSV = ".csv"
	_TEMPLATE_CSV  = `{{- .Hash }},{{ .Format }},{{ .Filename }},{{ .URL }}`
)

const (
	_EXTENSION_TXT = ".txt"
	_TEMPLATE_TXT  = `{{- .Hash }} {{ .Format }} {{ .Filename }} {{ .URL }}`
)

const (
	_CHECKSUM_FILENAME  = "checksum"
	_CHECKSUM_EXTENSION = ".txt"
)
