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
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
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

//nolint:gocognit
// Generate is to create all the debian / DEBIAN files and directories.
func (me *Data) Generate(workingDir string) (err error) {
	var path string

	err = me.Sanitize()
	if err != nil {
		goto done
	}

	if workingDir == "" {
		err = fmt.Errorf("%s: workingDir = ''", ERROR_DIR_MISSING)
		goto done
	}

	// identify debian or DEBIAN directory by Control/BuildSource.
	path = filepath.Join(workingDir, "debian")

	err = me.createBaseDir(path)
	if err != nil {
		goto done
	}

	err = me.createControl(path)
	if err != nil {
		goto done
	}

	err = me.createCopyright(path)
	if err != nil {
		goto done
	}

	err = me.CreateChangelog(path)
	if err != nil {
		goto done
	}

	err = me.createSourceDir(path)
	if err != nil {
		goto done
	}

	err = me.createSourceFormat(path)
	if err != nil {
		goto done
	}

	err = me.createSourceLocalOptions(path)
	if err != nil {
		goto done
	}

	err = me.createSourceOptions(path)
	if err != nil {
		goto done
	}

	err = me.createSourceLintianOverrides(path)
	if err != nil {
		goto done
	}

	err = me.createManpages(path)
	if err != nil {
		goto done
	}

	err = me.createScripts(path)
	if err != nil {
		goto done
	}

	err = me.createRules(path)
	if err != nil {
		goto done
	}

	err = me.createCompat(path)
	if err != nil {
		goto done
	}

	err = me.createInstall(path)
	if err != nil {
		goto done
	}

done:
	return err
}

func (me *Data) createInstall(path string) (err error) {
	target := filepath.Join(path, "install")

	s := ""
	isFirst := true
	for k, v := range me.Install {
		if !isFirst {
			s += "\n"
		}

		s += v + " " + k
		isFirst = false
	}

	return me._writeFile(target, _PERMISSION_FILE, s)
}

func (me *Data) createCompat(path string) (err error) {
	target := filepath.Join(path, "compat")

	return me._writeFile(target,
		_PERMISSION_FILE,
		strconv.Itoa(int(me.Compat)))
}

func (me *Data) createRules(path string) (err error) {
	target := filepath.Join(path, "rules")

	return me._writeFile(target, _PERMISSION_EXEC, me.Rules)
}

func (me *Data) createManpages(path string) (err error) {
	var manpages, name, fPath string
	list := []string{}

	for k, v := range me.Manpage {
		// generate the manpage path
		name = me.Control.Name + "." + k
		fPath = filepath.Join(path, name)

		// write to file
		err = me._writeFile(fPath, _PERMISSION_FILE, v)
		if err != nil {
			return err
		}

		// add manpage into list of manpages
		list = append(list, "debian/"+name)
	}

	// write the final debian/manpages
	manpages = ""
	for _, v := range list {
		manpages += v + "\n"
	}

	fPath = filepath.Join(path, me.Control.Name+".manpages")
	err = me._writeFile(fPath, _PERMISSION_FILE, manpages)
	if err != nil {
		return err
	}

	return nil
}

func (me *Data) createScripts(path string) (err error) {
	for k, v := range me.Scripts {
		target := filepath.Join(path, string(k))

		err = me._writeFile(target, _PERMISSION_EXEC, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (me *Data) createSourceLintianOverrides(path string) (err error) {
	var target string

	if me.Source.LintianOverrides == "" {
		return nil
	}

	target = filepath.Join(path, "source", "lintian-overrides")
	err = me._writeFile(target, _PERMISSION_FILE,
		me.Source.LintianOverrides)
	if err != nil {
		return err
	}

	target = filepath.Join(path, me.Control.Name+".lintian-overrides")
	return me._writeFile(target, _PERMISSION_FILE,
		me.Source.LintianOverrides)
}

func (me *Data) createSourceOptions(path string) (err error) {
	if me.Source.Options == "" {
		return nil
	}

	target := filepath.Join(path, "source", "options")

	return me._writeFile(target, _PERMISSION_FILE, me.Source.Options)
}

func (me *Data) createSourceLocalOptions(path string) (err error) {
	if me.Source.LocalOptions == "" {
		return nil
	}

	target := filepath.Join(path, "source", "local-options")

	return me._writeFile(target, _PERMISSION_FILE, me.Source.LocalOptions)
}

func (me *Data) createSourceFormat(path string) (err error) {
	target := filepath.Join(path, "source", "format")

	return me._writeFile(target, _PERMISSION_FILE, string(me.Source.Format))
}

func (me *Data) createSourceDir(path string) (err error) {
	target := filepath.Join(path, "source")

	return me._createDir(target)
}

// CreateChangelog creates the changelog file into given directory path.
//
// The given path **MUST** be a directory path and the generated changelog file
// is named `changelog` with no extension (same as `debian/changelog`).
//
// If the Changelog.Path is available, this function shall automatically update
// the persistent Changelog data file in Changelog.Path.
func (me *Data) CreateChangelog(path string) (err error) {
	var info os.FileInfo

	// formulate the filepath
	target := filepath.Join(path, "changelog")

	// write latest changes if available
	if len(me.Changelog.Changes) > 0 {
		err = me._writeFile(target,
			_PERMISSION_FILE,
			me.Changelog.String(),
		)

		if err != nil {
			return err
		}
	}

	// check path for preprend file
	if me.Changelog.Path == "" {
		goto end
	}

	info, err = os.Stat(me.Changelog.Path)
	switch {
	case os.IsNotExist(err):
		goto updatePersistentDataFile
	case err != nil:
		return fmt.Errorf("%s: %s", ERROR_CHANGELOG_PATH_BAD, err)
	case info != nil && info.Mode().IsRegular():
		err = me._appendChangelog(target, me.Changelog.Path)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf(ERROR_CHANGELOG_NOT_A_FILE)
	}

updatePersistentDataFile:
	if len(me.Changelog.Changes) == 0 {
		goto end // data file was the updated version
	}

	err = me._overwriteFile(me.Changelog.Path, target)
	if err != nil {
		return err
	}

end:
	return nil
}

func (me *Data) _overwriteFile(out string, in string) (err error) {
	var f, s *os.File

	f, err = os.OpenFile(out, os.O_RDWR|os.O_CREATE, _PERMISSION_FILE)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_OPEN_FAILED, err)
		goto done
	}

	s, err = os.OpenFile(in, os.O_RDONLY, _PERMISSION_FILE)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_OPEN_FAILED, err)
		goto doneWriter
	}

	_, err = io.Copy(f, s)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_OVERWRITE_FAILED, err)
		goto doneReader
	}

doneReader:
	s.Close()
doneWriter:
	_ = f.Sync()
	f.Close()
done:
	return err
}

func (me *Data) _appendChangelog(target string, source string) (err error) {
	var ok bool
	var scanner *bufio.Scanner
	var f, s *os.File
	var out string
	var notFirst bool
	var entry *Changelog

	// open changelog writer
	f, err = os.OpenFile(target,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		_PERMISSION_FILE,
	)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_OPEN_FAILED, err)
	}

	// open changelog reader
	s, err = os.OpenFile(source, os.O_RDONLY, _PERMISSION_FILE)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_OPEN_FAILED, err)
		goto closeWriter
	}
	scanner = bufio.NewScanner(s)

	// parse each entries
	entry = &Changelog{}

	if len(me.Changelog.Changes) > 0 {
		notFirst = true
	}

	for scanner.Scan() {
		// parse the line
		ok, err = entry.Parse(scanner.Text())
		if err != nil {
			goto closeReader
		}

		// append to changelog file
		if ok {
			out = entry.String()

			if notFirst {
				out = "\n\n" + out
			}

			_, err = f.WriteString(out)
			if err != nil {
				err = fmt.Errorf("%s: %s",
					ERROR_FILE_WRITE_FAILED,
					err,
				)
				goto closeReader
			}

			notFirst = true
			entry = &Changelog{}
		}
	}

closeReader:
	s.Close()
closeWriter:
	_ = f.Sync()
	f.Close()
	return err
}

func (me *Data) createCopyright(path string) (err error) {
	target := filepath.Join(path, "copyright")

	return me._writeFile(target, _PERMISSION_FILE, me.Copyright.String())
}

func (me *Data) createControl(path string) (err error) {
	target := filepath.Join(path, "control")

	return me._writeFile(target, _PERMISSION_FILE, me.Control.String())
}

func (me *Data) createBaseDir(path string) (err error) {
	return me._createDir(path)
}

func (me *Data) _createDir(path string) (err error) {
	_ = os.RemoveAll(path)

	err = os.MkdirAll(path, _PERMISSION_DIR)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_DIR_CREATE_FAILED, err)
	}

	return err
}

func (me *Data) _writeFile(path string,
	perm os.FileMode, data string) (err error) {
	var f *os.File

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, _PERMISSION_FILE)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_OPEN_FAILED, err)
	}

	_, err = f.WriteString(data)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_WRITE_FAILED, err)
		goto closeWriter
	}

	err = f.Chmod(perm)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_CHMOD_FAILED, err)
	}

closeWriter:
	_ = f.Sync()
	f.Close()
	return err
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
