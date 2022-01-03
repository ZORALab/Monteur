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
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ChangelogUrgency are the strict urgency values.
type ChangelogUrgency string

const (
	CHANGELOG_URGENCY_LOW       ChangelogUrgency = "low"
	CHANGELOG_URGENCY_MEDIUM    ChangelogUrgency = "medium"
	CHANGELOG_URGENCY_HIGH      ChangelogUrgency = "high"
	CHANGELOG_URGENCY_EMERGENCY ChangelogUrgency = "emergency"
	CHANGELOG_URGENCY_CRITICAL  ChangelogUrgency = "critical"
)

// changelogParseStatus are the status of the Changelog parser.
type changelogParseStatus uint

const (
	_CHANGELOG_PARSE_HEADER changelogParseStatus = iota
	_CHANGELOG_PARSE_CHANGE
	_CHANGELOG_PARSE_SIGNATURE
)

const (
	_CHANGELOG_TIMESTAMP_FORMAT = time.RFC1123Z
)

// Changelog is the DEBIAN/changelog file data.
//
// To prevent memory hogging with large changelog file, Changelog only process
// the latest changelog entry and prepend to the changelog file with a given
// filepath.
//
// The design is to hold all necessary data for rending that entry accurately.
type Changelog struct {
	statFx func(string) (os.FileInfo, error)

	// Version is the .deb package version number of release
	//
	// This field is **MANDATORY**.
	Version *Version

	// Maintainer is the contact person for packaging this change entry
	//
	// This field is **MANDATORY**.
	Maintainer *Entity

	// Timestamp is the timestamp of the changelog entry
	//
	// This field is **MANDATORY**.
	Timestamp *time.Time

	// Package is the name of the package.
	//
	// This field is **MANDATORY**.
	Package string

	// Path is the path to the changelog file for reading its history.
	//
	// If the filepath does not exist, Changelog shall create one.
	Path string

	// Urgency is the urgency of the changes without header.
	//
	// This field is **MANDATORY**.
	//
	// Header will be added automatically during rendering.
	Urgency ChangelogUrgency

	// Changes are the list of changes without spacing and asterisk.
	//
	// This field is **MANDATORY**.
	//
	// Spacing and asterisk will be added automatically during rendering.
	Changes []string

	// Distribution are the list of supported distributions.
	//
	// This field is **MANDATORY**.
	//
	// Space separator will be added automatically during rendering.
	Distribution []string

	parseStatus changelogParseStatus
}

func (me *Changelog) Parse(line string) (ready bool, err error) {
	// format the input line
	line = strings.TrimLeft(line, "\n")
	line = strings.TrimLeft(line, "\r")
	line = strings.TrimRight(line, "\n")
	line = strings.TrimRight(line, "\r")
	line = strings.TrimRight(line, " ")

	// line is empty. Ignore it
	if line == "" {
		goto done
	}

	// check line expectation
	switch me.parseStatus {
	case _CHANGELOG_PARSE_HEADER:
		err = me.parseHeader(line)
	case _CHANGELOG_PARSE_CHANGE:
		err = me.parseChanges(line)
	case _CHANGELOG_PARSE_SIGNATURE:
		err = fmt.Errorf("%s: Unexpected line '%s'",
			ERROR_CHANGELOG_PARSE_COMPLETED,
			line,
		)
	default:
		panic("MONTEUR DEV: what is this changelog parse status? " +
			strconv.Itoa(int(me.parseStatus)))
	}

done:
	return (me.Sanitize() == nil), err
}

func (me *Changelog) parseSignature(line string) (err error) {
	line = strings.TrimPrefix(line, _FIELD_CHANGELOG_DELIMIT_SIGNATURE)

	// parse maintainer name
	me.Maintainer = &Entity{}
	fragments := strings.Split(line, _FIELD_CHANGELOG_DELIMIT_EMAIL)
	if len(fragments) != 2 {
		return fmt.Errorf("%s: unknown maintainer name. Got: '%s'",
			ERROR_CHANGELOG_PARSE_BAD,
			line,
		)
	}
	me.Maintainer.Name = fragments[0]
	line = fragments[1]

	// parse maintainer email
	fragments = strings.Split(line, _FIELD_CHANGELOG_DELIMIT_TIMESTAMP)
	if len(fragments) != 2 {
		return fmt.Errorf("%s: unknown maintainer email. Got: '%s'",
			ERROR_CHANGELOG_PARSE_BAD,
			line,
		)
	}
	me.Maintainer.Email = fragments[0]
	me.Maintainer.Email = strings.TrimLeft(me.Maintainer.Email, "> ")
	me.Maintainer.Email = strings.TrimRight(me.Maintainer.Email, "> ")
	line = fragments[1]

	// parse timestamp
	line = strings.TrimLeft(line, " ")
	line = strings.TrimRight(line, "\n")
	line = strings.TrimRight(line, "\r")
	line = strings.TrimRight(line, " ")
	t, err := time.Parse(_CHANGELOG_TIMESTAMP_FORMAT, line)
	if err != nil || t.IsZero() || line == "" {
		return fmt.Errorf("%s: unknown timestamp. Got: '%s'",
			ERROR_CHANGELOG_PARSE_BAD,
			line,
		)
	}

	me.Timestamp = &t
	me.Maintainer.Year = me.Timestamp.Year()

	// perform checking and sanitization
	err = me.Maintainer.Sanitize()
	if err != nil {
		return fmt.Errorf("%s: bad maintainer data. Got: '%s'",
			ERROR_CHANGELOG_PARSE_BAD,
			err,
		)
	}

	// secure the parsing lock so that we do not expect parsing activity
	me.parseStatus = _CHANGELOG_PARSE_SIGNATURE
	return nil
}

func (me *Changelog) parseChanges(line string) (err error) {
	switch {
	case strings.Contains(line, _FIELD_CHANGELOG_DELIMIT_URGENCY):
		return fmt.Errorf("%s: expect change entry. Got: '%s'",
			ERROR_CHANGELOG_PARSE_BAD,
			line,
		)
	case strings.HasPrefix(line, _FIELD_CHANGELOG_DELIMIT_SIGNATURE):
		if len(me.Changes) == 0 {
			return fmt.Errorf("%s: expect change entry. Got: '%s'",
				ERROR_CHANGELOG_PARSE_BAD,
				line,
			)
		}

		// call signature since this line is a signature itself
		me.parseStatus = _CHANGELOG_PARSE_SIGNATURE
		return me.parseSignature(line)
	}

	if me.Changes == nil {
		me.Changes = []string{}
	}

	line = strings.TrimPrefix(line, _FIELD_CHANGELOG_DELIMIT_CHANGE)
	me.Changes = append(me.Changes, line)

	return nil
}

func (me *Changelog) parseHeader(line string) (err error) {
	switch {
	case strings.HasPrefix(line, _FIELD_CHANGELOG_DELIMIT_CHANGE):
		return fmt.Errorf("%s: expected header. Got: '%s'",
			ERROR_CHANGELOG_PARSE_BAD,
			line,
		)
	case strings.HasPrefix(line, _FIELD_CHANGELOG_DELIMIT_SIGNATURE):
		return fmt.Errorf("%s: expected header. Got: '%s'",
			ERROR_CHANGELOG_PARSE_BAD,
			line,
		)
	}

	// perform header parsing
	line, err = me._parseUrgency(line)
	if err != nil {
		return err
	}

	line, err = me._parsePackage(line)
	if err != nil {
		return err
	}

	line, err = me._parseVersion(line)
	if err != nil {
		return err
	}

	err = me._parseDistro(line)
	if err != nil {
		return err
	}

	// switch to parsing changes
	me.parseStatus = _CHANGELOG_PARSE_CHANGE

	return nil
}

func (me *Changelog) _parseDistro(line string) (err error) {
	line = strings.TrimLeft(line, " ")
	line = strings.TrimRight(line, " ")
	me.Distribution = strings.Split(line, " ")

	if line == "" || len(me.Distribution) == 0 {
		return fmt.Errorf("%s: missing distro '%s'",
			ERROR_CHANGELOG_PARSE_DISTRO_FAILED,
			line,
		)
	}

	return nil
}

func (me *Changelog) _parseVersion(input string) (line string, err error) {
	var epoch int

	lines := strings.Split(input, _FIELD_CHANGELOG_DELIMIT_DISTRO)
	if len(lines) != 2 {
		return input, fmt.Errorf("%s: broken format '%s'",
			ERROR_CHANGELOG_PARSE_VERSION_FAILED,
			input,
		)
	}

	ver := lines[0]
	line = lines[1]

	me.Version = &Version{}

	// obtain upstream and revision
	fragments := strings.Split(ver, "-")
	if len(fragments) == 1 {
		me.Version.Upstream = ver
	} else {
		me.Version.Revision = fragments[len(fragments)-1]
		me.Version.Upstream = strings.TrimSuffix(ver,
			"-"+me.Version.Revision,
		)
	}

	// strip epoch
	fragments = strings.Split(me.Version.Upstream, ":")
	if len(fragments) == 2 {
		// strip epoch
		epoch, err = strconv.Atoi(fragments[0])
		if err != nil {
			return input, fmt.Errorf(
				"%s: error in epoch conversion: %s",
				ERROR_CHANGELOG_PARSE_VERSION_FAILED,
				err,
			)
		}
		me.Version.Epoch = uint(epoch)

		// update upstream
		me.Version.Upstream = fragments[1]
	}

	// sanitize version
	err = me.Version.Sanitize()
	if err != nil {
		return input, fmt.Errorf("%s: %s",
			ERROR_CHANGELOG_PARSE_VERSION_FAILED,
			err,
		)
	}

	return line, nil
}

func (me *Changelog) _parsePackage(input string) (line string, err error) {
	lines := strings.Split(input, _FIELD_CHANGELOG_DELIMIT_PACKAGE)
	if len(lines) != 2 {
		return input, fmt.Errorf("%s: bad format '%s'",
			ERROR_CHANGELOG_PARSE_PACKAGE_FAILED,
			input,
		)
	}

	me.Package = lines[0]
	line = lines[1]

	return line, err
}

func (me *Changelog) _parseUrgency(input string) (line string, err error) {
	lines := strings.Split(input, _FIELD_CHANGELOG_DELIMIT_URGENCY)
	if len(lines) != 2 {
		return input, fmt.Errorf("%s: bad format '%s'",
			ERROR_CHANGELOG_PARSE_URGENCY_FAILED,
			input,
		)
	}

	me.Urgency = ChangelogUrgency(lines[1])
	line = lines[0]

	return line, err
}

// Sanitize is to ensure Changelog data is compliant to the strict format.
//
// It shall return error if any data is not compliant.
func (me *Changelog) Sanitize() (err error) {
	err = me.sanitizePackage()
	if err != nil {
		return err
	}

	err = me.sanitizeVersion()
	if err != nil {
		return err
	}

	err = me.sanitizeDistribution()
	if err != nil {
		return err
	}

	err = me.sanitizeUrgency()
	if err != nil {
		return err
	}

	err = me.sanitizeChanges()
	if err != nil {
		return err
	}

	err = me.sanitizeMaintainer()
	if err != nil {
		return err
	}

	err = me.sanitizeTimestamp()
	if err != nil {
		return err
	}

	err = me.sanitizePath()
	if err != nil {
		return err
	}

	return nil
}

func (me *Changelog) sanitizePath() (err error) {
	if me.statFx == nil {
		me.statFx = os.Stat
	}

	info, err := me.statFx(me.Path)
	switch {
	case err != nil && os.IsNotExist(err):
	case err != nil:
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Path",
			err,
		)
	case info.IsDir():
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Path",
			"path is directory: "+me.Path,
		)
	default:
	}

	return nil
}

func (me *Changelog) sanitizeTimestamp() (err error) {
	if me.Timestamp == nil {
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Timestamp",
			me.Timestamp,
		)
	}

	if me.Timestamp.IsZero() {
		return fmt.Errorf("%s (%s): not set. Got '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Timestamp",
			me.Timestamp.Format(_CHANGELOG_TIMESTAMP_FORMAT),
		)
	}

	return nil
}

func (me *Changelog) sanitizeMaintainer() (err error) {
	if me.Maintainer == nil {
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Maintainer",
			me.Maintainer,
		)
	}

	err = me.Maintainer.Sanitize()
	if err != nil {
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Maintainer",
			err,
		)
	}

	return nil
}

func (me *Changelog) sanitizeChanges() (err error) {
	if me.Changes == nil {
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Changes",
			me.Changes,
		)
	}

	if len(me.Changes) == 0 && me.Path == "" {
		return fmt.Errorf("%s (%s): %s '%d'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Changes",
			"list is empty",
			len(me.Changes),
		)
	}

	for i, v := range me.Changes {
		if len(v) > 80 {
			return fmt.Errorf("%s (line %d): '%s'",
				ERROR_CHANGELOG_ENTRY_TOO_LONG,
				i+1,
				v,
			)
		}
	}

	return nil
}

func (me *Changelog) sanitizeUrgency() (err error) {
	switch me.Urgency {
	case CHANGELOG_URGENCY_LOW:
	case CHANGELOG_URGENCY_MEDIUM:
	case CHANGELOG_URGENCY_HIGH:
	case CHANGELOG_URGENCY_EMERGENCY:
	case CHANGELOG_URGENCY_CRITICAL:
	default:
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Urgency",
			me.Urgency,
		)
	}

	return nil
}

func (me *Changelog) sanitizeDistribution() (err error) {
	if me.Distribution == nil {
		me.Distribution = []string{}
	}

	if len(me.Distribution) == 0 {
		return fmt.Errorf("%s (%s): '%v'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Distribution",
			me.Distribution,
		)
	}

	return nil
}

func (me *Changelog) sanitizeVersion() (err error) {
	if me.Version == nil {
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Version",
			me.Version,
		)
	}

	err = me.Version.Sanitize()
	if err != nil {
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Version",
			err,
		)
	}

	return nil
}

func (me *Changelog) sanitizePackage() (err error) {
	if me.Package == "" {
		return fmt.Errorf("%s (%s): '%s'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Package",
			me.Package,
		)
	}

	me.Package = strings.ToLower(me.Package)

	return nil
}

// String generates the DEBIAN/changelog latest entry for preprend.
//
// It will panic if Changelog is not sanitized before use and an error has
// occurred.
func (me *Changelog) String() (s string) {
	err := me.Sanitize()
	if err != nil {
		panic(err)
	}

	if len(me.Changes) == 0 {
		return ""
	}

	// header
	s += me.Package
	s += _FIELD_CHANGELOG_DELIMIT_PACKAGE + me.Version.String()
	s += _FIELD_CHANGELOG_DELIMIT_DISTRO +
		strings.Join(me.Distribution, " ")
	s += _FIELD_CHANGELOG_DELIMIT_URGENCY + string(me.Urgency) + "\n"

	// optional line spacing
	s += "\n"

	// change data
	for _, v := range me.Changes {
		s += _FIELD_CHANGELOG_DELIMIT_CHANGE + v + "\n"
	}

	// optional line spacing
	s += "\n"

	// signed-off
	s += _FIELD_CHANGELOG_DELIMIT_SIGNATURE +
		me.Maintainer.Tag() +
		_FIELD_CHANGELOG_DELIMIT_TIMESTAMP +
		me.Timestamp.Format(_CHANGELOG_TIMESTAMP_FORMAT)

	return s
}
