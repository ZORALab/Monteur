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

package oshelper

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Copy is a function to duplicate a given source file/directory to destination.
//
// When given a file (be it regular, symlink, or pipe file across various
// operating systems) as `source`, it will be copied to the given `dest`
// filepath.
//
// When given a directory path as `source`, Copy shall recursively duplicate
// that directory alongside its contents to the `dest` directory path.
//
// Should the `source` and `destination` be a file, **the filename with its
// extensions SHALL be included inside the path**.
//
// Pipe file copying is not available on Windows and it shall return unsupported
// error.
func Copy(source string, dest string) (err error) {
	var fi os.FileInfo
	var mode fs.FileMode

	if source == "" {
		return fmt.Errorf(ERROR_SOURCE_EMPTY)
	}

	if dest == "" {
		return fmt.Errorf(ERROR_DEST_EMPTY)
	}

	fi, err = os.Lstat(source)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_OPEN, err)
	}

	mode = fi.Mode()
	switch {
	case mode.IsRegular():
		err = copyFile(source, dest, fi)
	case mode.IsDir():
		err = copyDir(source, dest)
	case mode&fs.ModeSymlink != 0:
		err = copySymlink(source, dest, fi)
	case mode&fs.ModeNamedPipe != 0:
		err = copyPipe(source, dest, fi)
	default:
		err = fmt.Errorf(ERROR_SOURCE_UNKNOWN)
	}

	return err
}

// CopyPath creates dest's housing directory before performing `Copy(...)`.
//
// This is to ensure a seamless copy without needing to manually create housing
// directory.
func CopyPath(source string, dest string) (err error) {
	err = os.MkdirAll(filepath.Dir(dest), 0755)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_DIRECTORY_CREATE, err)
	}

	return Copy(source, dest)
}

func copyFile(source string, dest string, fi os.FileInfo) (err error) {
	var sourceFile, destFile *os.File

	// open source file for read
	sourceFile, err = os.OpenFile(source, os.O_RDONLY, os.ModePerm)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_OPEN, err)
		goto done
	}

	// open dest file for write
	destFile, err = os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_OPEN, err)
		goto endSource
	}

	// copy file
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_COPY, err)
		_ = destFile.Sync()
		goto endDest
	}

	// sync dest file for good write
	err = destFile.Sync()
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_SYNC, err)
		goto endDest
	}

	// close all files for inodes restorations
	destFile.Close()
	sourceFile.Close()

	// restore file permission
	err = os.Chmod(dest, fi.Mode())
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_PERM, err)
		goto done
	}

	// restore file ownership
	err = _restoreOwnership(dest, fi)
	if err != nil {
		goto done
	}

	// restore timestamp
	err = _restoreTimestamp(dest, fi)
	if err != nil {
		goto done
	}

endDest:
	destFile.Close()
endSource:
	sourceFile.Close()
done:
	return err
}

func copyDir(source string, dest string) (err error) {
	err = filepath.Walk(source, func(path string,
		info os.FileInfo, err error) error {
		// if err is found, return now
		if err != nil {
			return fmt.Errorf("%s: %s", ERROR_DIRECTORY, err)
		}

		// ensure we are not working on the given source path and only
		// works with the content.
		if path == source {
			return nil
		}

		// generate relative path
		destPath, _ := filepath.Rel(source, path)
		destPath = filepath.Join(dest, destPath)

		// take the next action according to its nature recursively
		mode := info.Mode()
		switch {
		case mode.IsRegular():
			err = os.MkdirAll(filepath.Dir(destPath),
				PERMISSION_DIR,
			)
			if err != nil {
				return fmt.Errorf("%s: %s",
					ERROR_DIRECTORY_CREATE,
					err,
				)
			}

			err = copyFile(path, destPath, info)
		case mode.IsDir():
			err = os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return fmt.Errorf("%s: %s",
					ERROR_DIRECTORY_CREATE,
					err,
				)
			}

			err = copyDir(path, destPath)
		case mode&os.ModeSymlink != 0:
			err = os.MkdirAll(filepath.Dir(destPath),
				PERMISSION_DIR,
			)
			if err != nil {
				return fmt.Errorf("%s: %s",
					ERROR_DIRECTORY_CREATE,
					err,
				)
			}

			err = copySymlink(path, destPath, info)
		case mode&fs.ModeNamedPipe != 0:
			err = os.MkdirAll(filepath.Dir(destPath),
				PERMISSION_DIR,
			)
			if err != nil {
				return fmt.Errorf("%s: %s",
					ERROR_DIRECTORY_CREATE,
					err,
				)
			}

			err = copyPipe(path, destPath, info)
		default:
			err = fmt.Errorf("%s: '%s'", ERROR_SOURCE_UNKNOWN, path)
		}

		return err
	})

	return err
}

func copySymlink(source string, dest string, fi os.FileInfo) (err error) {
	var link string

	// read symlink data at source
	link, err = os.Readlink(source)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_SYMLINK_READ, err)
	}

	// create new symlink at dest
	err = os.Symlink(dest, link)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_SYMLINK_CREATE, err)
	}

	err = _restoreSymlinkTimestamp(dest, fi)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_CHTIMES, err)
	}

	return nil
}

func copyPipe(source string, dest string, fi os.FileInfo) (err error) {
	return _copyPipe(source, dest, fi) // OS specific
}

func _restoreOwnership(dest string, fi os.FileInfo) (err error) {
	uid, gid := FileOwners(fi)

	err = os.Lchown(dest, uid, gid)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_CHOWN, err)
	}

	return err
}

func _restoreTimestamp(dest string, fi os.FileInfo) (err error) {
	accessed, _, modified := FileTimestamps(fi)

	err = os.Chtimes(dest, accessed, modified)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_CHTIMES, err)
	}

	return FileSetPlatformTime(dest, modified)
}

func _restoreSymlinkTimestamp(dest string, fi os.FileInfo) (err error) {
	accessed, _, modified := SymlinkTimestamps(fi)

	err = SymlinkChtimes(dest, accessed, modified)
	if err != nil {
		return err
	}

	return nil
}
