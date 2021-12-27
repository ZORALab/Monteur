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

package archive

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SanitizeArchive ensures the Archive filepath is ready for use.
//
// It will absolute the given pathing and return it as output should no error
// occur.
//
// In the event of error, no output (`out`) is always empty.
func SanitizeArchive(path, ext string) (out string, err error) {
	var info os.FileInfo

	// sanitize input parameters
	if path == "" {
		return "", fmt.Errorf(ERROR_ARCHIVE_EMPTY)
	}

	// sanitize absolute pathing
	out, err = filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("%s: %s",
			ERROR_ARCHIVE_ABS_FAILED,
			path,
		)
	}

	// sanitize file extension
	if ext != "" && !strings.HasSuffix(out, ext) {
		return "", fmt.Errorf("%s: '%s'",
			ERROR_ARCHIVE_EXT_MISSING,
			ext,
		)
	}

	// sanitize housing directory
	path = filepath.Dir(out)
	info, err = os.Stat(path)
	switch {
	case os.IsNotExist(err):
		return "", fmt.Errorf("%s: %s", ERROR_ARCHIVE_DIR_MISSING, path)
	case !info.IsDir():
		return "", fmt.Errorf("%s: %s", ERROR_ARCHIVE_DIR_INVALID, path)
	}

	// sanitize archive file
	info, err = os.Stat(out)
	switch {
	case err == nil && !info.IsDir(), os.IsExist(err), os.IsNotExist(err):
		return out, nil // Compress() and Extract() shall handle
	case info.IsDir():
		return "", fmt.Errorf("%s: %s", ERROR_ARCHIVE_ABS_FAILED, out)
	default:
		return "", fmt.Errorf("%s: %s", ERROR_ARCHIVE, err)
	}
}

// SanitizeRaw ensures the Raw data directory is ready for use.
//
// It will absolute the given pathing and return it as output should no error
// occur.
//
// In the event of error, no output (`out`) is always empty.
func SanitizeRaw(path string) (out string, err error) {
	var info os.FileInfo

	// sanitize input parameters
	if path == "" {
		return "", fmt.Errorf(ERROR_RAW_EMPTY)
	}

	// sanitize absolute pathing
	out, err = filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("%s: %s", ERROR_RAW_ABS_FAILED, path)
	}

	// sanitize archive file
	info, err = os.Stat(out)
	switch {
	case err == nil:
		if info.IsDir() {
			return out, nil
		}

		return "", fmt.Errorf("%s: %s", ERROR_RAW_NOT_DIR, out)
	case os.IsNotExist(err):
		return "", fmt.Errorf("%s: %s", ERROR_RAW_MISSING, out)
	default:
		return "", fmt.Errorf("%s: %s", ERROR_RAW, err)
	}
}

// SanitizeCompressionPath ensures the compression target path is correct.
//
// This function shall scans the compressing target path for various
// vulnerabilities and makes sure it is safe before use.
func SanitizeCompressionPath(base string, path string) (v string, err error) {
	// Sanitize input
	if base == "" {
		return "", fmt.Errorf(ERROR_PATH_BASE_EMPTY)
	}

	if path == "" {
		return "", fmt.Errorf(ERROR_PATH_EMPTY)
	}

	// absolute base pathing
	base, err = filepath.Abs(base)
	if err != nil {
		return "", fmt.Errorf("%s: %s", ERROR_PATH_ABS_FAILED, base)
	}

	// Scan for G305: Zip Slip Vulnerability
	v = filepath.Join(base, path)
	if !strings.HasPrefix(v, filepath.Clean(base)) {
		return "", fmt.Errorf("%s: %s", ERROR_PATH_OUT_OF_BOUND, path)
	}

	return v, nil
}
