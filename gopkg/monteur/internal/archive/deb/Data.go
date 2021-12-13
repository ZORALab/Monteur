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
	"strings"
)

// ShellScriptType are the executable shell scripts timing type for dpkg to run.
type ShellScriptType string

const (
	SHELL_PRE_INSTALL  ShellScriptType = "preinst"
	SHELL_POST_INSTALL ShellScriptType = "postinst"
	SHELL_PRE_REMOVE   ShellScriptType = "prerm"
	SHELL_POST_REMOVE  ShellScriptType = "postrm"
)

// Data is the full Debian specific package metadata for archiving purposes.
type Data struct {
	// Control is the control data saved into DEBIAN/control.
	//
	// This field is **MANDATORY**.
	Control *Control

	// Copyright is the copyright data saved into DEBIAN/copyright
	//
	// This field is **MANDATORY**.
	Copyright *Copyright

	// Changelog are the changelog data saved into DEBIAN/changelog
	//
	// This field is **MANDATORY**.
	Changelog *Changelog

	// Source is the deb file format saved into DEBIAN/source directory.
	//
	// This field is **MANDATORY**.
	Source *Source

	// Manpage is the manpage data saved into DEBIAN/app.manpages
	//
	// This field is **MANDATORY**.
	//
	// The key will be the manpage extension while the value is the content
	// of the manpage.
	//
	// Each key:value dataset shall generate its own file.
	Manpage map[string]string

	// Scripts are the shell scripts for dpkg to run.
	//
	// The key will be used as the filename while the value is the content
	// of the script.
	//
	// Each key:value dataset shall generate its own script file with
	// executable permission.
	Scripts map[ShellScriptType]string

	// Install are the list of files for DEBIAN/install.
	//
	// This field is **MANDATORY**.
	//
	// The `key` is used to designate the installation path while the
	// `value` is the pathing to the file for installation.
	//
	// Spacing between the key:value shall be formatted automatically so
	// there is no need to manually include it.
	Install map[string]string

	// Rules are the package unpack operations data for DEBIAN/rules file.
	//
	// This field is **MANDATORY**.
	Rules string

	// Compat is the debhelper compatibility level for DEBIAN/compat file.
	//
	// This field is **MANDATORY**.
	Compat uint
}

// Sanitize is to ensure all Data is complying to the .deb strict format.
//
// This function shall return error should any data is not compliant.
func (me *Data) Sanitize() (err error) {
	err = me.sanitizeControl()
	if err != nil {
		return err
	}

	err = me.sanitizeCopyright()
	if err != nil {
		return err
	}

	err = me.sanitizeChangelog()
	if err != nil {
		return err
	}

	err = me.sanitizeSource()
	if err != nil {
		return err
	}

	err = me.sanitizeManpage()
	if err != nil {
		return err
	}

	err = me.sanitizeScripts()
	if err != nil {
		return err
	}

	err = me.sanitizeInstall()
	if err != nil {
		return err
	}

	err = me.sanitizeRules()
	if err != nil {
		return err
	}

	err = me.sanitizeCompat()
	if err != nil {
		return err
	}

	return nil
}

func (me *Data) sanitizeCompat() (err error) {
	if me.Compat == 0 {
		return fmt.Errorf(ERROR_COMPAT_UNSET)
	}

	return nil
}

func (me *Data) sanitizeRules() (err error) {
	if me.Rules == "" {
		return fmt.Errorf(ERROR_RULES_EMPTY)
	}

	return nil
}

func (me *Data) sanitizeInstall() (err error) {
	if me.Install == nil {
		me.Install = map[string]string{}
	}

	if len(me.Install) == 0 {
		return fmt.Errorf("%s: list is 'nil'", ERROR_INSTALL_BAD)
	}

	list := map[string]string{}
	checklist := map[string][]string{}
	for k, v := range me.Install {
		if k == "" {
			continue
		}

		// check against Debian checklist
		err = me._scanDebianChecklist3p7p0(k)
		if err != nil {
			return err
		}

		err = me._scanDebianChecklist4p6p0(k)
		if err != nil {
			return err
		}

		// register the installation path first
		list[k] = v

		// add program's installation path to checklist
		array, ok := checklist[v]
		if !ok {
			checklist[v] = []string{k}
		} else {
			checklist[v] = append(array, k)
		}
	}

	err = me._scanDebianChecklist4p0p0(checklist)
	if err != nil {
		return err
	}

	me.Install = list

	return nil
}

func (me *Data) _scanDebianChecklist4p6p0(k string) (err error) {
	if strings.HasPrefix(k, "usr/lib64") ||
		strings.HasPrefix(k, "/usr/lib64") {
		return fmt.Errorf("%s: %s",
			ERROR_CHECKLIST_4_6_0_LIB64,
			k,
		)
	}

	switch {
	}

	return nil
}

func (me *Data) _scanDebianChecklist3p7p0(k string) (err error) {
	if strings.HasPrefix(k, "usr/X11R6") ||
		strings.HasPrefix(k, "/usr/X11R6") {
		return fmt.Errorf("%s: %s",
			ERROR_CHECKLIST_3_7_0_XORG,
			k,
		)
	}

	return nil
}

func (me *Data) _scanDebianChecklist4p0p0(l map[string][]string) (err error) {
	var isPath, isUsrPath bool

	for k, v := range l {
		isPath = false
		isUsrPath = false

		for _, pathing := range v {
			switch {
			case strings.HasPrefix(pathing, "/usr/"),
				strings.HasPrefix(pathing, "usr/"):
				isUsrPath = true
			default:
				isPath = true
			}

			if isPath && isUsrPath {
				return fmt.Errorf("%s: '%s'",
					ERROR_CHECKLIST_4_0_0_PATHS,
					k,
				)
			}
		}
	}

	return nil
}

func (me *Data) sanitizeScripts() (err error) {
	if me.Scripts == nil {
		me.Scripts = map[ShellScriptType]string{}
	}

	if len(me.Scripts) == 0 {
		return nil
	}

	list := map[ShellScriptType]string{}
	for k, v := range me.Scripts {
		if v == "" {
			continue
		}

		switch k {
		case SHELL_PRE_INSTALL:
		case SHELL_POST_INSTALL:
		case SHELL_PRE_REMOVE:
		case SHELL_POST_REMOVE:
		default:
			return fmt.Errorf("%s: '%s'",
				ERROR_SHELL_SCRIPT_TYPE_UNKNOWN,
				string(k),
			)
		}

		list[k] = v
	}
	me.Scripts = list

	return nil
}

func (me *Data) sanitizeManpage() (err error) {
	if me.Manpage == nil {
		me.Manpage = map[string]string{}
	}

	if len(me.Manpage) == 0 {
		return fmt.Errorf("%s: list are empty", ERROR_MANPAGE_BAD)
	}

	list := map[string]string{}
	for k, v := range me.Manpage {
		if v == "" {
			continue
		}

		list[k] = v
	}
	me.Manpage = list

	return nil
}

func (me *Data) sanitizeSource() (err error) {
	if me.Source == nil {
		return fmt.Errorf("%s: Source is 'nil'", ERROR_SOURCE_BAD)
	}

	return me.Source.Sanitize()
}

func (me *Data) sanitizeChangelog() (err error) {
	if me.Changelog == nil {
		return fmt.Errorf("%s: Changelog is 'nil'",
			ERROR_CHANGELOG_BAD,
		)
	}

	return me.Changelog.Sanitize()
}

func (me *Data) sanitizeCopyright() (err error) {
	if me.Copyright == nil {
		return fmt.Errorf("%s: Copyright is 'nil'",
			ERROR_CONTROL_BAD,
		)
	}

	return me.Copyright.Sanitize()
}

func (me *Data) sanitizeControl() (err error) {
	if me.Control == nil {
		return fmt.Errorf("%s: Control is 'nil'",
			ERROR_CONTROL_BAD,
		)
	}

	_, err = me.Control.Sanitize()

	return err
}
