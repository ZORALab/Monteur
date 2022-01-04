// Copyright 2022 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2022 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
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
	"strings"
)

// Init is for Init API to setup Monteur into a repository.
type Init struct {
	pwd string
}

// Run is for Init to execute all the repo initialization for Monteur.
func (me *Init) Run() (out int) {
	me.reportStatus("Initializing repo with Monteur..." + "\n")
	me.getPWD()

	out = me.isWorkspaceExists()
	if out != STATUS_OK {
		return out
	}

	out = me.createWorkspaceTOML()
	if out != STATUS_OK {
		return out
	}

	out = me.createCIJobsConfigs()
	if out != STATUS_OK {
		return out
	}

	out = me.createAppMetadata()
	if out != STATUS_OK {
		return out
	}

	out = me.createAppHelp()
	if out != STATUS_OK {
		return out
	}

	out = me.createAppDebian()
	if out != STATUS_OK {
		return out
	}

	out = me.createAppCopyright()
	if out != STATUS_OK {
		return out
	}

	me.reportStatus("Initializing repo with Monteur " + LOG_SUCCESS + "\n")
	return STATUS_OK
}

func (me *Init) createAppCopyright() int {
	var path string
	var err error

	path = filepath.Join(me.pwd,
		DIRECTORY_MONTEUR_CONFIG,
		DIRECTORY_APP_CONFIG,
		LANG_CODE_DEFAULT,
		DIRECTORY_APP_COPYRIGHT,
		FILE_TOML_APP_COPYRIGHT,
	)
	me.reportStatus("Creating: '%s'\n", path)

	err = me._createFile(path, `[Copyright]
Name = 'License Name' # license full name
ID = 'license-name'   # identifiable license ID
Comment = ''
Materials = [ '*' ]
Holders = [
        # A string with year, then name, then email with diamond braces. Ex.:
        #
        #   '2021 John S. Smith <john.s.smith@email.com>',
        #
        # If the list is empty, all maintainers, contributors, and sponsors
        # shall be attached to it.
]
Notice = """
License notice usually prepend in front of source codes.
"""
Text = """
Full license text body.
"""
`)
	if err != nil {
		me.reportError("Error creating file: %s\n", err)
		return STATUS_ERROR
	}

	me.reportStatus(LOG_OK + "\n")
	return STATUS_OK
}

func (me *Init) createAppDebian() int {
	var path string
	var err error

	path = filepath.Join(me.pwd,
		DIRECTORY_MONTEUR_CONFIG,
		DIRECTORY_APP_CONFIG,
		LANG_CODE_DEFAULT,
		FILE_TOML_APP_DEBIAN,
	)
	me.reportStatus("Creating: '%s'\n", path)

	err = me._createFile(path, `[DEB]
Compat = 11
Rules = """
#!/usr/bin/make -f

# Uncomment this to turn on verbose mode.
#export DH_VERBOSE=1

%:
	dh $@

override_dh_auto_build:
	echo "nothing to override"
"""

[DEB.Control]
Essential = false
PackageType = 'deb'
Priority = 'optional'
RulesRequiresRoot = 'binary-targets'
Standards = '4.6.0'
Section = 'devel'

[DEB.Relationships]
'Build-Depends' = [
	'debhelper (>= 11)',
]
'Depends' = [
]

# More Info:
#  https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-vcs-fields
[DEB.VCS]
Type = 'Vcs-Git'
URL = 'https://example.com/git'
Branch = 'main'
#Path = '.'

[DEB.Testsuite]
Paths = [
	# 'relative/path/to/debTestScript',
]

[DEB.Copyright]
Format = 'https://www.debian.org/doc/packaging-manuals/copyright-format/1.0/'
Disclaimer = ''
Comment = ''

[DEB.Changelog]
Urgency = 'low'

[DEB.Source]
Format = '3.0 (native)'
LocalOptions = """
"""
Options = """
"""
LintianOverrides = """
# supply the license file in case OS did not supply one

# we are go package so it is okay not to specify Depends:
"""

[DEB.Install]`)
	if err != nil {
		me.reportError("Error creating file: %s\n", err)
		return STATUS_ERROR
	}

	me.reportStatus(LOG_OK + "\n")
	return STATUS_OK
}

func (me *Init) createAppHelp() int {
	var path string
	var err error

	path = filepath.Join(me.pwd,
		DIRECTORY_MONTEUR_CONFIG,
		DIRECTORY_APP_CONFIG,
		LANG_CODE_DEFAULT,
		FILE_TOML_APP_HELP,
	)
	me.reportStatus("Creating: '%s'\n", path)

	//nolint:lll
	err = me._createFile(path, `[Help]
Command = '$ {{ .App.Command }} help'
Description = '{{ .App.Abstract }}'
Resources = 'Visit {{ .App.Website }} for detailed documentations.'

[Help.Manpage]
Lv1 = """
.\\" {{ .App.Name }} - Lv1 Manpage
.\\" Contact {{ index .App.Contact.Email 0 }} for errors or typos.
.TH man 1 "{{- .App.Time.Day }} {{ .App.Time.Month }} {{ .App.Time.Year -}}" "{{- .App.Version -}}" "{{- .App.ID }} man page"

.SH NAME
{{ .App.Name }} - {{ .App.Abstract }}

.SH SYNOPSIS
{{ .App.Help.Command }}

.SH DESCRIPTION
{{ .App.Description -}}

.SH OPTIONS
{{ .App.Help.Description }}

.SH SEE ALSO
{{ .App.Help.Resources }}

#.SH AUTHORS
#{{ range $index, $owner := .App.Maintainers -}}
#       {{- $owner.Name }} ({{- index $owner.Email 0 -}})
#{{ end -}}
"""`)
	if err != nil {
		me.reportError("Error creating file: %s\n", err)
		return STATUS_ERROR
	}

	me.reportStatus(LOG_OK + "\n")
	return STATUS_OK
}

func (me *Init) createAppMetadata() int {
	var appName, path string
	var err error

	appName = filepath.Base(me.pwd)

	path = filepath.Join(me.pwd,
		DIRECTORY_MONTEUR_CONFIG,
		DIRECTORY_APP_CONFIG,
		LANG_CODE_DEFAULT,
		FILE_TOML_APP_METADATA,
	)
	me.reportStatus("Creating: '%s'\n", path)

	err = me._createFile(path, `[Software]
Name = '`+strings.Title(appName)+`'
Command = '`+ProcessToFilepath(appName)+`'
ID = '`+ProcessToFilepath(appName)+`'
Version = 'v0.0.1' # version number
Category = 'devel' # software category
Suite = ''         # software suite (e.g. MSFT Office for MSFT Word)
Abstract = 'software XYZ tool in one single app.'
Description = """
Describe your software as much as you want. Please make sure you end it with
its unique selling pitch.
"""
Website = 'https://example.com/'

[Software.Contact]
Name = 'ACME Entity'
Email = [ 'hello@example.com' ]




[Software.Maintainers.ACME]
Name = 'ACME Entity'
Email = [ 'tech@example.com' ]
[Software.Maintainers.ACME.JointTime]
Year = '2021'




[Software.Contributors.ACME]
Name = 'ACME Entity'
Email = [ 'tech@example.com' ]
[Software.Contributors.ACME.JointTime]
Year = '2021'




[Software.Sponsors.ACME]
Name = 'ACME Entity'
Email = [ 'tech@example.com' ]
[Software.Sponsors.ACME.JointTime]
Year = '2021'
`)
	if err != nil {
		me.reportError("Error creating file: %s\n", err)
		return STATUS_ERROR
	}

	me.reportStatus(LOG_OK + "\n")
	return STATUS_OK
}

func (me *Init) createCIJobsConfigs() int {
	var jobs []string
	var path string
	var err error

	jobs = []string{
		DIRECTORY_SETUP,
		DIRECTORY_CLEAN,
		DIRECTORY_TEST,
		DIRECTORY_PREPARE,
		DIRECTORY_BUILD,
		DIRECTORY_PACKAGE,
		DIRECTORY_RELEASE,
		DIRECTORY_COMPOSE,
		DIRECTORY_PUBLISH,
	}

	for _, job := range jobs {
		// [1] create job-wide config file
		path = filepath.Join(me.pwd,
			DIRECTORY_MONTEUR_CONFIG,
			job,
			FILE_TOML,
		)
		me.reportStatus("Creating: '%s'\n", path)

		err = me._createFile(path, `[Variables]

[FMTVariables]`)
		if err != nil {
			me.reportError("Error creating file: %s\n", err)
			return STATUS_ERROR
		}
		me.reportStatus(LOG_OK + "\n")

		// [2] create job directory
		path = filepath.Join(me.pwd,
			DIRECTORY_MONTEUR_CONFIG,
			job,
			DIRECTORY_JOBS,
			EXTENSION_GITKEEP,
		)
		me.reportStatus("Creating: '%s'\n", path)

		err = me._touch(path)
		if err != nil {
			me.reportError("Error creating file: %s\n", err)
			return STATUS_ERROR
		}
		me.reportStatus(LOG_OK + "\n")
	}

	return STATUS_OK
}

func (me *Init) createWorkspaceTOML() int {
	var path string
	var err error

	path = filepath.Join(me.pwd,
		DIRECTORY_MONTEUR_CONFIG,
		FILE_TOML_WORKSPACE,
	)
	me.reportStatus("Creating: '%s'\n", path)

	err = me._createFile(path, `[Filesystem]
BaseDir = '.'
WorkingDir = '.monteurFS/tmp/'
BuildDir = '.monteurFS/build/'
ScriptDir = '.scripts/'
BinDir = '.monteurFS/bin/'
BinCfgDir = '.monteurFS/config'
LogDir = '.monteurFS/log'
DataDir = '.configs/monteur/app/data'
ReleaseDir = '.monteurFS/releases'
SecretsDir = [
        '{{ .HomeDir }}/.secrets',
        '{{ .RootDir }}/.configs/monteur/secrets',
]

[Language]
Name = '`+LANG_NAME_DEFAULT+`'
Code = '`+LANG_CODE_DEFAULT+`'

[Variables]

[FMTVariables]
`)
	if err != nil {
		// report error
		me.reportError("error creating file: %s\n", err)
		return STATUS_ERROR
	}

	me.reportStatus(LOG_OK + "\n")
	return STATUS_OK
}

func (me *Init) isWorkspaceExists() int {
	var path, pwd string
	var err error

	pwd = me.pwd
	for {
		if pwd == "/" {
			break
		}

		path = filepath.Join(pwd,
			DIRECTORY_MONTEUR_CONFIG,
			FILE_TOML_WORKSPACE,
		)
		me.reportStatus("Checking existing Monteur at '%s'\n", path)

		_, err = os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				me.reportStatus(LOG_OK + "\n")
				pwd = filepath.Dir(pwd)
				continue
			}

			me.reportError("error locating Monteur: %s\n", err)
			return STATUS_ERROR
		}

		me.reportError("Existing Monteur detected!\n")
		return STATUS_ERROR
	}

	return STATUS_OK
}

func (me *Init) _createFile(path string, data string) (err error) {
	var f *os.File

	err = os.MkdirAll(filepath.Dir(path), PERMISSION_DIRECTORY)
	if err != nil {
		return err //nolint:wrapcheck
	}

	f, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, PERMISSION_FILE)
	if err != nil {
		return err //nolint:wrapcheck
	}

	_, err = f.WriteString(data)
	_ = f.Sync()
	f.Close()

	return err //nolint:wrapcheck
}

func (me *Init) _touch(path string) (err error) {
	var f *os.File

	err = os.MkdirAll(filepath.Dir(path), PERMISSION_DIRECTORY)
	if err != nil {
		return err //nolint:wrapcheck
	}

	f, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, PERMISSION_FILE)
	if err != nil {
		return err //nolint:wrapcheck
	}

	f.Close()
	return nil
}

func (me *Init) reportStatus(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func (me *Init) reportError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "[ ERROR ] "+format, args...)
}

func (me *Init) getPWD() {
	me.pwd, _ = os.Getwd()
}
