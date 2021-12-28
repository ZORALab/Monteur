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
	"io"
	"os"
	"path/filepath"
	"time"

	"gitlab.com/zoralab/monteur/gopkg/oshelper"
)

const (
	COPY_SIZE      int64 = 32 * 1024
	PATH_SEPARATOR       = "/"
)

const (
	PERMISSION_ALL  = 0777
	PERMISSION_EXEC = 0750
	PERMISSION_FILE = 0655
	PERMISSION_DIR  = 0755
)

// MkdirAll is to make directory with a given decision.
func MkdirAll(path string, mode os.FileMode, decision bool) (err error) {
	var info os.FileInfo

	if !decision {
		info, err = os.Stat(path)
		if err != nil {
			return fmt.Errorf("%s (%s): %s", ERROR_DIR, path, err)
		}

		if !info.IsDir() {
			return fmt.Errorf("%s (%s): %s",
				ERROR_DIR_INVALID,
				path,
				err,
			)
		}

		return nil
	}

	err = os.MkdirAll(path, PERMISSION_DIR)
	if err != nil {
		err = fmt.Errorf("%s (%s): %s",
			ERROR_MKDIR_FAILED,
			path,
			err,
		)
	}

	return err
}

// EvalSymlink is to evaluate a given symlink's target path is ready for use.
//
// It calls `os.Stat(...)` against the target's path to ensure the target exists
// and healthy.
//
// It also calls `SanitizeCompressionPath(...)` internally for checking the
// target's path validity for compression.
func EvalSymlink(base string,
	path string) (out string, info os.FileInfo, err error) {
	out, err = filepath.EvalSymlinks(path)
	if err != nil {
		return "", nil, fmt.Errorf("%s (%s): %s",
			ERROR_SYMLINK_EVAL_FAILED,
			path,
			err,
		)
	}

	info, err = os.Stat(out)
	if err != nil {
		return "", nil, fmt.Errorf("%s (%s): %s",
			ERROR_SYMLINK_UNRESOLVABLE,
			path,
			err,
		)
	}

	out, err = SanitizeCompressionPath(base, out)
	if err != nil {
		return "", nil, fmt.Errorf("%s where %s",
			ERROR_SYMLINK_EVAL,
			err,
		)
	}

	return out, info, nil
}

// CreateSymlink is to create a symlink.
func CreateSymlink(target string, path string) (err error) {
	if target == "" {
		return fmt.Errorf(ERROR_SYMLINK_TARGET_EMPTY)
	}

	if path == "" {
		return fmt.Errorf(ERROR_SYMLINK_PATH_EMPTY)
	}

	err = os.Symlink(target, path)
	if err != nil {
		return fmt.Errorf("%s (%s -> %s): %s",
			ERROR_SYMLINK_CREATE,
			path,
			target,
			err,
		)
	}

	return nil
}

// RelPath is to make relative path from given `base` and `path` pathings.
//
// It will convert to forward slash (`/`) before generating output.
//
// Should there be any error, output (`out`) shall be empty.
func RelPath(base string, path string) (out string, err error) {
	out, err = filepath.Rel(base, path)
	if err != nil {
		return "", fmt.Errorf("%s: %s",
			ERROR_PATH_REL_FAILED,
			path,
		)
	}

	return filepath.ToSlash(out), nil
}

// Overwrite is to remove the target path for overwriting.
//
// This function only performs the removal of target path should the decision
// is right.
//
// If the target is already missing, the function shall return `nil`.
func Overwrite(path string, decision bool) (err error) {
	// sanitize input
	if path == "" {
		return nil
	}

	// check target existence
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("%s (%s): %s",
			ERROR_FILE_INFO_READ,
			path,
			err,
		)
	}

	if !decision {
		return fmt.Errorf("%s (%s): %s",
			ERROR_TARGET_EXISTS,
			path,
			err,
		)
	}

	err = os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			ERROR_OVERWRITE_FAILED,
			path,
			err,
		)
	}

	return nil
}

// ExtractCopy is to perform io.Copy() in a controlled stream packet.
//
// This copy method is to counter the given compress file being a large file
// and causes denial-of-service vulnerability.
//
// All Archiver's file restoration (`Extract(...)`) shall use this function
// in order to have the same security benefits unilaterally.
//
// If the byteSize is unset (`0`), it will use the default size which is
// `COPY_SIZE` constant.
func ExtractCopy(w io.Writer,
	r io.Reader, byteSize int64, path string) (err error) {
	if byteSize <= 0 {
		byteSize = COPY_SIZE
	}

	for {
		_, err = io.CopyN(w, r, byteSize)
		if err != nil {
			if err == io.EOF {
				break
			}

			return fmt.Errorf("%s (%s): %s",
				ERROR_EXTRACT_COPY,
				path,
				err,
			)
		}
	}

	return nil
}

// RestoreMetadata is to restore target file metadata.
//
// This function currently restores access time, modified time, and file mode.
//
// Sysmlink metadata restoration only works on supporting operating system.
func RestoreMetadata(path string,
	aTime time.Time, mTime time.Time,
	mode os.FileMode, isSymlink bool) (err error) {
	// validate input
	if mode < 0000 || mode > 0777 {
		return fmt.Errorf("%s (%s): %o",
			ERROR_FILE_MODE_INVALID,
			path,
			mode,
		)
	}

	if aTime.IsZero() {
		return fmt.Errorf("%s (%s): %o",
			ERROR_FILE_ATIME_INVALID,
			path,
			mode,
		)
	}

	if mTime.IsZero() {
		return fmt.Errorf("%s (%s): %o",
			ERROR_FILE_MTIME_INVALID,
			path,
			mode,
		)
	}

	// restore timestamp
	switch {
	case isSymlink:
		err = oshelper.SymlinkChtimes(path, aTime, mTime)
	default:
		err = os.Chtimes(path, aTime, mTime)
	}

	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_CHTIMES_FAILED, err)
	}

	// restore file mode
	err = os.Chmod(path, mode)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			ERROR_FILE_CHMOD_FAILED,
			path,
			err,
		)
	}

	return nil
}
