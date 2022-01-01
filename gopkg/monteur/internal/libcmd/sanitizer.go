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

package libcmd

import (
	"fmt"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/commander"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/templater"
)

func sanitizeChangelog(changelog *libmonteur.TOMLChangelog,
	system string) (err error) {
	// sanitize CMD
	cmd := changelog.CMD
	changelog.CMD = []*libmonteur.TOMLAction{}
	err = sanitizeCMD(cmd, &changelog.CMD, system)
	if err != nil {
		return err
	}

	return nil
}

func sanitizeCMD(in []*libmonteur.TOMLAction,
	out *[]*libmonteur.TOMLAction,
	system string) (err error) {
	for _, cmd := range in {
		if !libmonteur.IsComputeSystemSupported(system, cmd.Condition) {
			continue
		}

		err = cmd.Sanitize()
		if err != nil {
			return err //nolint:wrapcheck
		}

		*out = append(*out, cmd)
	}

	return nil
}

func sanitizeDeps(in []*libmonteur.TOMLDependency,
	out *[]*commander.Dependency,
	system string,
	variables map[string]interface{}) (err error) {
	var val string

	// scan conditions for building commands list
	for _, dep := range in {
		if dep == nil {
			continue
		}

		if !libmonteur.IsComputeSystemSupported(system,
			[]string{dep.Condition}) {
			continue
		}

		val, err = templater.String(dep.Command, variables)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_COMMAND_DEPENDENCY_FMT_BAD,
				err,
			)
		}

		s := &commander.Dependency{
			Name:    dep.Name,
			Type:    dep.Type,
			Command: val,
		}

		*out = append(*out, s)
	}

	// sanitize each commands for validity
	for _, dep := range *out {
		err = dep.Init()
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_DEPENDENCY_BAD,
				err,
			)
		}
	}

	return nil
}

func sanitizeMetadata(meta *libmonteur.TOMLMetadata, path string) (err error) {
	err = meta.Sanitize(path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}

func sanitizeRelease(release *libmonteur.TOMLRelease,
	variables map[string]interface{}) (err error) {
	var k string
	var v *libmonteur.TOMLReleasePackage
	var packages map[string]*libmonteur.TOMLReleasePackage

	// sanitize Target
	release.Target, err = libmonteur.ProcessString(release.Target,
		variables,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// sanitize Data
	if release.Data == nil {
		release.Data = &libmonteur.TOMLReleaseData{}
	}

	// sanitize Data.Path
	release.Data.Path, err = libmonteur.ProcessString(release.Data.Path,
		variables,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// sanitize Packages
	packages = map[string]*libmonteur.TOMLReleasePackage{}
	for k, v = range release.Packages {
		if k == "" {
			continue
		}

		// sanitize individual source
		if v.Source == "" {
			continue
		}

		// sanitize individual target
		switch {
		case v.Target != "":
		case release.Target != "":
			v.Target = release.Target
		default:
			return fmt.Errorf("%s: '%s'",
				libmonteur.ERROR_RELEASER_TARGET_MISSING,
				k,
			)
		}

		packages[k] = v
	}
	release.Packages = packages

	return nil
}

func sanitizeSources(sources map[string]*libmonteur.TOMLSource,
	out **libmonteur.TOMLSource,
	variables *map[string]interface{},
	thisSystem string) (err error) {
	// get 'all-all' omni-platform first and merge accordingly
	actual := sources[libmonteur.COMPUTE_SYSTEM_OMNI]

	// get platform specific source
	specific := sources[thisSystem]

	// merge both into one
	if actual == nil {
		if specific == nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_PROGRAM_UNSUPPORTED,
				thisSystem,
			)
		}

		actual = specific
	} else if specific != nil {
		actual.Merge(specific)
	}

	// save to output
	*out = actual

	// sanitize
	err = sanitizeSourceFormat(*out, variables)
	if err != nil {
		return err
	}

	err = sanitizeSourceArchive(*out, variables)
	if err != nil {
		return err
	}

	err = sanitizeSourceMethod(*out, variables)
	if err != nil {
		return err
	}

	err = sanitizeSourceURL(*out, variables)
	if err != nil {
		return err
	}

	err = sanitizeSourceHeaders(*out, variables)
	if err != nil {
		return err
	}

	err = sanitizeSourceChecksum(*out)
	if err != nil {
		return err
	}

	return err
}

func sanitizeSourceChecksum(out *libmonteur.TOMLSource) (err error) {
	// exit if there is no checksum. It's optional anyway.
	if out.Checksum == nil {
		return nil
	}

	if out.Checksum.Value == "" {
		return fmt.Errorf("%s: ''",
			libmonteur.ERROR_CHECKSUM_BAD,
		)
	}

	if out.Checksum.Format == "" {
		return fmt.Errorf("%s: ''",
			libmonteur.ERROR_CHECKSUM_FORMAT_BAD,
		)
	}

	if out.Checksum.Type == "" {
		return fmt.Errorf("%s: ''",
			libmonteur.ERROR_CHECKSUM_TYPE_BAD,
		)
	}

	return nil
}

func sanitizeSourceHeaders(out *libmonteur.TOMLSource,
	variables *map[string]interface{}) (err error) {
	// loop through each of them and process accordingly
	headers := map[string]string{}
	for k, v := range out.Headers {
		v, err = libmonteur.ProcessString(v, *variables)
		if err != nil {
			return fmt.Errorf("%s: %s",
				libmonteur.ERROR_PROGRAM_HTTPS_HEADER_BAD,
				err,
			)
		}

		headers[k] = v
	}

	out.Headers = headers

	return nil
}

func sanitizeSourceURL(out *libmonteur.TOMLSource,
	variables *map[string]interface{}) (err error) {
	out.URL, err = libmonteur.ProcessString(out.URL, *variables)
	if err != nil {
		return fmt.Errorf("%s: URL error = %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			err,
		)
	}

	if out.URL == "" {
		return fmt.Errorf("%s: URL = '%s'",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			out.URL,
		)
	}

	(*variables)[libmonteur.VAR_URL] = out.URL

	return nil
}

func sanitizeSourceMethod(out *libmonteur.TOMLSource,
	variables *map[string]interface{}) (err error) {
	out.Method, err = libmonteur.ProcessString(out.Method, *variables)
	if err != nil {
		return fmt.Errorf("%s: Method error = %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			err,
		)
	}

	if out.Method == "" {
		return fmt.Errorf("%s: Method = '%s'",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			out.Method,
		)
	}

	(*variables)[libmonteur.VAR_METHOD] = out.Method

	return nil
}

func sanitizeSourceArchive(out *libmonteur.TOMLSource,
	variables *map[string]interface{}) (err error) {
	out.Archive, err = libmonteur.ProcessString(out.Archive, *variables)
	if err != nil {
		return fmt.Errorf("%s: Archive error = %s",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			err,
		)
	}

	if out.Archive == "" {
		return fmt.Errorf("%s: Archive = '%s'",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			out.Archive,
		)
	}

	(*variables)[libmonteur.VAR_ARCHIVE] = out.Archive

	return nil
}

func sanitizeSourceFormat(out *libmonteur.TOMLSource,
	variables *map[string]interface{}) (err error) {
	out.Format, err = libmonteur.ProcessString(out.Format, *variables)
	if err != nil {
		return fmt.Errorf("%s: Format error = '%s'",
			libmonteur.ERROR_PROGRAM_ARCHIVE_BAD,
			out.Format,
		)
	}

	out.Format = strings.ToLower(out.Format)
	(*variables)[libmonteur.VAR_FORMAT] = out.Format

	return nil
}

func sanitizeSourceConfig(cfg map[string]string,
	out *string,
	variables map[string]interface{}) (err error) {
	var ok bool
	var os string

	// get os for query
	os, ok = variables[libmonteur.VAR_OS].(string)
	if !ok {
		panic("Monteur DEV: why is VAR_OS missing?")
	}

	*out, ok = cfg[os]
	if !ok {
		*out = cfg[libmonteur.ALL_OS]
	}

	// check config contents
	if *out == "" {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_BAD,
			"missing config for "+os,
		)
	}

	*out, err = libmonteur.ProcessString(*out, variables)
	if err != nil {
		return fmt.Errorf("%s: %s",
			libmonteur.ERROR_PROGRAM_CONFIG_BAD,
			err,
		)
	}

	return nil
}
