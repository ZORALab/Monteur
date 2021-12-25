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

package libdeb

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/archive/deb"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

func Prepare(pkg *libmonteur.TOMLPackage,
	variables map[string]interface{}) (d *deb.Data, err error) {
	// process deb-specific variables
	packagePath := variables[libmonteur.VAR_PACKAGE].(string)
	packagePath = filepath.Join(packagePath, "deb")
	variables[libmonteur.VAR_PACKAGE] = packagePath
	_ = os.MkdirAll(packagePath, libmonteur.PERMISSION_DIRECTORY)

	// get app data as needed across multiple internal functions
	app := variables[libmonteur.VAR_APP].(*libmonteur.Software)

	d, err = createAppData(app, variables, pkg, packagePath)
	if err != nil || d == nil {
		d = nil
		goto done
	}

	err = prepareWorkspace(d, packagePath)
	if err != nil {
		d = nil
		goto done
	}

done:
	return d, err
}

func prepareWorkspace(d *deb.Data, path string) (err error) {
	err = d.Generate(path)
	if err != nil {
		err = fmt.Errorf("%s: %s",
			libmonteur.ERROR_PACKAGER_PREPARE_FAILED,
			err,
		)
	}

	return err
}

func createAppData(app *libmonteur.Software,
	variables map[string]interface{},
	pkg *libmonteur.TOMLPackage,
	packagePath string) (d *deb.Data, err error) {
	var timestamp *time.Time
	var ok, buildSource bool
	var archList []string
	var control *deb.Control
	var copyright *deb.Copyright
	var changelog *deb.Changelog
	var source *deb.Source
	var version *deb.Version
	var packages map[deb.PackageListType]*deb.PackageList

	// extract important data
	timestamp, ok = variables[libmonteur.VAR_TIMESTAMP].(*time.Time)
	if !ok {
		panic("MONTEUR DEV: why is libmonteur.VAR_TIMESTAMP missing?")
	}

	archList, ok = variables[libmonteur.VAR_PACKAGE_ARCH].([]string)
	if !ok {
		err = fmt.Errorf("%s", libmonteur.ERROR_PACKAGER_ARCH_MISSING)
		goto done
	}

	packages, err = _createPackagesList(app.Debian.Relationships)
	if err != nil {
		goto done
	}

	// IMPORTANT NOTE:
	// monteur's BuildSource has a different meaning from .deb's
	// `BuildSource. Montuer's means build source package while `.deb`
	// is building the binary package from source data. Hence, the inverse
	// is needed since we want to build package from `.deb`, source data.
	//
	// We will revisit this decision after native build is supported.
	buildSource = !pkg.BuildSource

	version = &deb.Version{
		Epoch:    _createVersionEpoch(app.Version),
		Upstream: _createVersionUpstream(app.Version),
		Revision: _createVersionRevision(app.Version),
	}

	err = version.Sanitize()
	if err != nil {
		err = fmt.Errorf("%s: %s",
			libmonteur.ERROR_PACKAGER_VERSION_INCOMPATIBLE,
			err,
		)
		goto done
	}

	variables[libmonteur.VAR_PACKAGE_VERSION] = version.String()

	// construct the .deb Data
	control = &deb.Control{
		Maintainer: &deb.Entity{
			Name:  _createEntityName(app.Contact),
			Email: _createEntityEmail(app.Contact),
			Year:  timestamp.Year(),
		},
		Packages: packages,
		Version:  version,
		Homepage: _createURL(app.Website),
		Description: &deb.Description{
			Synopsis: app.Abstract,
			Body:     app.Description,
		},
		VCS: &deb.VCS{
			Type:    deb.VCSType(app.Debian.VCS.Type),
			Browser: _createURL(app.Website),
			URL:     _createURL(app.Debian.VCS.URL),
			Branch:  app.Debian.VCS.Branch,
			Path:    app.Debian.VCS.Path,
		},
		Testsuite:         _createTestsuite(app, packagePath),
		Name:              app.Name,
		Section:           app.Debian.Control.Section,
		StandardsVersion:  deb.StandardsVersion(app.Debian.Control.Standards),
		PackageType:       deb.PackageType(app.Debian.Control.PackageType),
		Priority:          deb.Priority(app.Debian.Control.Priority),
		RulesRequiresRoot: app.Debian.Control.RulesRequiresRoot,
		Architecture:      _createArch(archList),
		Uploaders:         _createUploaders(app, timestamp.Year()),
		Essential:         app.Debian.Control.Essential,
		BuildSource:       buildSource,
	}

	copyright = &deb.Copyright{
		Format:     deb.CopyrightFormat(app.Debian.Copyright.Format),
		Name:       app.Name,
		Contact:    control.Maintainer,
		Source:     app.Website,
		Disclaimer: app.Debian.Copyright.Disclaimer,
		Comment:    app.Debian.Copyright.Comment,
		Copyright:  control.Uploaders,
		Licenses: _createLicenses(control.Uploaders,
			app.Copyrights,
			timestamp.Year(),
		),
	}

	changelog = &deb.Changelog{
		Version:      control.Version,
		Maintainer:   control.Maintainer,
		Timestamp:    timestamp,
		Package:      control.Name,
		Path:         pkg.Changelog,
		Urgency:      deb.ChangelogUrgency(app.Debian.Changelog.Urgency),
		Changes:      _createChangelogEntries(variables),
		Distribution: pkg.Distribution,
	}

	source = &deb.Source{
		Format:           deb.SourceFormatType(app.Debian.Source.Format),
		LocalOptions:     app.Debian.Source.LocalOptions,
		Options:          app.Debian.Source.Options,
		LintianOverrides: app.Debian.Source.LintianOverrides,
	}

	d = &deb.Data{
		Control:   control,
		Copyright: copyright,
		Changelog: changelog,
		Source:    source,
		Manpage:   app.Help.Manpage,
		Scripts:   _createScripts(app.Debian.Scripts),
		Install:   _createInstall(app.Debian.Install, variables),
		Rules:     app.Debian.Rules,
		Compat:    app.Debian.Compat,
	}

	err = d.Sanitize()
	if err != nil {
		d = nil
		err = fmt.Errorf("%s: %s",
			libmonteur.ERROR_PACKAGER_DEB_BAD,
			err,
		)
	}

done:
	return d, err
}

func _createInstall(in map[string]string,
	variables map[string]interface{}) (out map[string]string) {
	var target, source, k, v string
	out = map[string]string{}

	for k, v = range in {
		target, _ = libmonteur.ProcessString(k, variables)
		source, _ = libmonteur.ProcessString(v, variables)
		out[target] = source
	}

	return out
}

func _createPackagesList(in map[string][]string) (out map[deb.PackageListType]*deb.PackageList,
	err error) {
	var pkgList *deb.PackageList

	out = map[deb.PackageListType]*deb.PackageList{}
	for rel, list := range in {
		pkgList = &deb.PackageList{}

		err = pkgList.Parse(rel, list)
		if err != nil {
			return nil, fmt.Errorf("%s: %s",
				libmonteur.ERROR_PACKAGER_RELATIONSHIPS,
				err,
			)
		}

		out[pkgList.Name] = pkgList
	}

	return out, nil
}

func _createChangelogEntries(variables map[string]interface{}) (out []string) {
	var ret string
	var ok bool

	out, ok = variables[libmonteur.VAR_CHANGELOG_ENTRIES].([]string)
	if !ok {
		return []string{}
	}

	// There are entries. Let's truncate it.
	list := []string{}
	for _, v := range out {
		ret = __truncateChangelogEntry(v)
		list = append(list, ret)
	}

	return list
}

func __truncateChangelogEntry(in string) string {
	i := 0
	for j := range in {
		if i == 70 {
			return in[:j] + "..."
		}
		i++
	}

	return in
}

func _createTestsuite(app *libmonteur.Software, base string) *deb.Testsuite {
	if app.Debian.Testsuite == nil {
		return nil
	}

	if len(app.Debian.Testsuite.Paths) == 0 {
		return nil
	}

	return &deb.Testsuite{
		Basepath: base,
		Paths:    app.Debian.Testsuite.Paths,
	}
}

func _createScripts(in map[string]string) (out map[deb.ShellScriptType]string) {
	out = map[deb.ShellScriptType]string{}

	for k, v := range in {
		out[deb.ShellScriptType(k)] = v
	}

	return out
}

func _createLicenses(mainList []*deb.Entity,
	licenses []*libmonteur.Copyright,
	copyrightYear int) (output []*deb.License) {
	output = []*deb.License{}

	for _, v := range licenses {
		x := &deb.License{
			License: v.ID,
			Body:    v.Text,
			Comment: v.Comment,
			Copyright: _createLicenseHolders(mainList,
				v.Holders,
				copyrightYear,
			),
			Files: v.Materials,
		}

		output = append(output, x)
	}

	return output
}

func _createLicenseHolders(mainList []*deb.Entity,
	holders []string,
	copyrightYear int) (output []*deb.Entity) {
	if len(holders) == 0 {
		return mainList
	}

	// specific holders listed for overwriting
	output = []*deb.Entity{}
	for _, v := range holders {
		v = strings.TrimLeft(v, "\r\n ")
		v = strings.TrimRight(v, "\r\n ")

		list := strings.Split(v, " ")
		if len(list) < 3 {
			return mainList // assume everyone in due to error
		}

		year, err := strconv.Atoi(list[0])
		if err != nil {
			year = copyrightYear
		}

		name := strings.Join(list[1:len(list)-1], " ")

		e := &deb.Entity{
			Name:  name,
			Email: list[len(list)-1],
			Year:  year,
		}

		output = append(output, e)
	}

	return output
}

func _createArch(list []string) string {
	return list[0]
}

func _createURL(link string) *url.URL {
	if link == "" {
		return nil
	}

	x, err := url.Parse(link)
	if err != nil {
		return nil
	}

	return x
}

func _createVersionUpstream(version string) (upstream string) {
	// trim off revision if found
	revision := _createVersionRevision(version)
	if revision != "" {
		version = strings.TrimSuffix(version, "-"+revision)
	}

	// trim off epoch
	list := strings.Split(version, ":")
	if len(list) == 1 {
		upstream = list[0]
	} else {
		upstream = list[len(list)-1]
	}

	// trim all leading alphabets to be digit-led
	prefix := ""
	for _, c := range upstream {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			break
		default:
			prefix += string(c)
		}
	}

	return strings.TrimLeft(upstream, prefix)
}

func _createVersionEpoch(version string) uint {
	x := strings.Split(version, ":")
	if len(x) == 1 {
		return 0
	}

	epoch, err := strconv.Atoi(x[0])
	if err != nil {
		return 0
	}

	return uint(epoch)
}

func _createVersionRevision(version string) string {
	x := strings.Split(version, "-")
	if len(x) == 1 {
		return ""
	}

	return x[len(x)-1]
}

func _createUploaders(app *libmonteur.Software,
	copyrightYear int) (final []*deb.Entity) {
	var contributors map[string]*deb.Entity
	var x *deb.Entity

	// filter all maintainers, contributors, and sponsors
	contributors = map[string]*deb.Entity{}
	for k, v := range app.Maintainers {
		x = &deb.Entity{
			Name:  _createEntityName(v),
			Email: _createEntityEmail(v),
			Year:  _createCopyrightYear(v, copyrightYear),
		}

		contributors[k] = x
	}

	for k, v := range app.Contributors {
		x = &deb.Entity{
			Name:  _createEntityName(v),
			Email: _createEntityEmail(v),
			Year:  _createCopyrightYear(v, copyrightYear),
		}

		contributors[k] = x
	}

	for k, v := range app.Sponsors {
		x = &deb.Entity{
			Name:  _createEntityName(v),
			Email: _createEntityEmail(v),
			Year:  _createCopyrightYear(v, copyrightYear),
		}

		contributors[k] = x
	}

	// generate the array list
	final = []*deb.Entity{}
	for _, v := range contributors {
		final = append(final, v)
	}

	return final
}

func _createCopyrightYear(entity *libmonteur.Entity, copyrightYear int) int {
	if entity == nil {
		return copyrightYear
	}

	if entity.CopyrightTime == nil {
		return copyrightYear
	}

	year, err := strconv.Atoi(entity.CopyrightTime.Year)
	if err != nil {
		return copyrightYear
	}

	return year
}

func _createEntityEmail(entity *libmonteur.Entity) string {
	return entity.Email[0]
}

func _createEntityName(entity *libmonteur.Entity) string {
	return entity.Name
}
