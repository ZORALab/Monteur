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
			me.Timestamp.Format(time.RFC1123Z),
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

	if len(me.Changes) == 0 {
		return fmt.Errorf("%s (%s): %s '%d'",
			ERROR_CHANGELOG_ENTRY_BAD,
			"Changes",
			"list is empty",
			len(me.Changes),
		)
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

	// header
	s += me.Package
	s += " (" + me.Version.String() + ")"
	s += " " + strings.Join(me.Distribution, " ")
	s += "; urgency=" + string(me.Urgency) + "\n"

	// optional line spacing
	s += "\n"

	// change data
	for _, v := range me.Changes {
		s += "  * " + v + "\n"
	}

	// optional line spacing
	s += "\n"

	// signed-off
	s += " --" +
		me.Maintainer.Tag() +
		"  " +
		me.Timestamp.Format(time.RFC1123Z)

	return s
}
