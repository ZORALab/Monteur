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

package targz

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitlab.com/zoralab/monteur/gopkg/oshelper"
)

const (
	EXTENSION       = ".tar.gz"
	PERMISSION_FILE = 0655
	PERMISSION_DIR  = 0755
)

type Archiver struct {
	reader *tar.Reader
	writer *tar.Writer

	// Archive is the filepath to the `.tar.gz` archive file.
	//
	// The value is **STRICTLY** a filepath with filename and `.tar.gz`
	// file extension.
	Archive string

	// Raw is the directory path for the uncompressed data directory.
	//
	// The value **MUST** be a directory holding the archive content for
	// compression or an empty directory for decompression.
	Raw string

	// CreateDirectory decides on creating missing directory for Archive.
	//
	// When set to `true`, Archiver shall create the directory for housing
	// the `.tar.gz` archived file if it is not exist.
	//
	// Default is returning an error (`false`).
	CreateDirectory bool

	// Overwrite decides on overwriting RAW data directory.
	//
	// When set to `true`, Archiver shall delete and recreate directory
	// path given in RAW field.
	//
	// Default is returning an error (`false`).
	Overwrite bool

	// ReliefExtension decides to relax file extension checking.
	//
	// When set to `true`, Archiver shall not throw an error when checking
	// Archive for strict `.tar.gz` file extension.
	//
	// Default is returning an error (`false`).
	ReliefExtension bool

	// FollowSymlink decides to resolve symlink and archive package file.
	//
	// When set to `true`, Archiver shall resolve the symlink recursively
	// until the target file is saved.
	//
	// Default (`false`) is to save a relative symlink to the `Raw`
	// directory.
	FollowSymlink bool
}

// Sanitize initializes and check all input data are correct before executions.
//
// This function shall returns error if any data is not compliant.
func (me *Archiver) Sanitize() (err error) {
	var info os.FileInfo
	var path string

	// sanitize empty Archive and Raw
	if me.Archive == "" {
		return fmt.Errorf("%s: %s", ERROR_PATH_EMPTY, "Archive")
	}

	if me.Raw == "" {
		return fmt.Errorf("%s: %s", ERROR_PATH_EMPTY, "Raw")
	}

	// sanitize Archive - absolute path
	me.Archive, err = filepath.Abs(me.Archive)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PATH_ABS_FAILED, me.Archive)
	}

	// sanitize Archive - extension
	if !strings.HasSuffix(me.Archive, EXTENSION) &&
		!me.ReliefExtension {
		return fmt.Errorf("%s: %s", ERROR_EXTENSION_MISSING, me.Archive)
	}

	// sanitize Archive - housing directory
	path = filepath.Dir(me.Archive)
	info, err = os.Stat(path)
	switch {
	case os.IsNotExist(err) && !me.CreateDirectory:
		return fmt.Errorf("%s: %s", ERROR_PATH_DIR_MISSING, path)
	case !info.IsDir():
		return fmt.Errorf("%s: %s", ERROR_PATH_NOT_DIR, path)
	}

	// sanitize Archive - pathing
	info, err = os.Stat(me.Archive)
	switch {
	case err == nil && !info.IsDir(), os.IsExist(err), os.IsNotExist(err):
		// Compress() and Extract() shall check internally
	case info.IsDir():
		return fmt.Errorf("%s: %s", ERROR_PATH_IS_DIR, me.Archive)
	default:
		return fmt.Errorf("%s: %s", ERROR_PATH_ARCHIVE, err)
	}

	// sanitize Raw
	me.Raw, err = filepath.Abs(me.Raw)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PATH_ABS_FAILED, me.Raw)
	}

	info, err = os.Stat(me.Raw)
	switch {
	case os.IsNotExist(err):
		return fmt.Errorf("%s: %s", ERROR_PATH_EMPTY, me.Raw)
	case !info.IsDir():
		return fmt.Errorf("%s: %s", ERROR_PATH_NOT_DIR, me.Raw)
	}

	return nil
}

// Compress is to compress the Raw data directory into Archive file.
//
// It called Archiver.Sanitize() internally.
func (me *Archiver) Compress() (err error) {
	var f *os.File
	var gz *gzip.Writer

	err = me.Sanitize()
	if err != nil {
		return err
	}

	// create base directory if requested
	if me.CreateDirectory {
		err = os.MkdirAll(filepath.Dir(me.Archive), PERMISSION_DIR)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_CREATE_FAILED,
				me.Archive,
			)
		}
	}

	// check for overwrite
	_, err = os.Stat(me.Archive)
	if err == nil { // file exists
		if me.Overwrite {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_EXISTS,
				me.Archive,
			)
		}

		err = os.RemoveAll(me.Archive)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_OVERWRITE_FAILED,
				me.Archive,
			)
		}
	}

	// begin archiving
	f, err = os.Create(me.Archive)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_DEST_CREATE_FAILED, me.Archive)
	}
	defer f.Close()

	gz = gzip.NewWriter(f)
	defer gz.Close()
	me.writer = tar.NewWriter(gz)
	defer me.writer.Close()

	// walk through Raw data directory and compress it
	return filepath.Walk(me.Raw, me.compress)
}

func (me *Archiver) compress(path string, info os.FileInfo, err error) error {
	var mode os.FileMode

	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_READ_FAILED,
			path,
			err,
		)
	}

	mode = info.Mode()
	switch {
	case mode.IsRegular():
		return me.compressRegular(path, info)
	case mode.IsDir():
		return me.compressDir(path, info)
	case mode&os.ModeSymlink != 0:
		return me.compressSymlink(path, info)
	default:
		return fmt.Errorf("%s: %s", ERROR_FILE_UNSUPPORTED, path)
	}
}

func (me *Archiver) compressSymlink(path string, info os.FileInfo) (err error) {
	var header *tar.Header
	var targetInfo os.FileInfo
	var target string

	// evaluate target location
	target, err = filepath.EvalSymlinks(path)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_SYMLINK_READ_FAILED, path)
	}

	target, err = filepath.Abs(target)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PATH_ABS_FAILED, path)
	}

	// check symlink relative pathing
	if !strings.HasPrefix(target, me.Raw) {
		return fmt.Errorf("%s: %s", ERROR_SYMLINK_OUT_OF_BOUND, path)
	}

	// check follow decision
	if me.FollowSymlink {
		targetInfo, err = os.Stat(target)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_SYMLINK_UNRESOLVED,
				path,
			)
		}

		// resolve recursively
		return me.compress(target, targetInfo, nil)
	}

	// create relative symlink
	header, err = tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_HEADER_READ_FAILED, path)
	}

	// restore Name with relative pathing for preserving directory tree
	header.Name, err = filepath.Rel(me.Raw, path)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PATH_REL_FAILED, path)
	}

	// perform write
	err = me.writer.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_HEADER_WRITE_FAILED,
			path,
			err,
		)
	}

	return nil
}

func (me *Archiver) compressDir(path string, info os.FileInfo) (err error) {
	var header *tar.Header

	header, err = tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_HEADER_READ_FAILED, path)
	}

	// restore Name with relative pathing for preserving directory tree
	header.Name, err = filepath.Rel(me.Raw, path)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PATH_REL_FAILED, path)
	}

	// perform write
	err = me.writer.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_HEADER_WRITE_FAILED,
			path,
			err,
		)
	}

	return nil
}

func (me *Archiver) compressRegular(path string, info os.FileInfo) (err error) {
	var f *os.File
	var header *tar.Header

	header, err = tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_HEADER_READ_FAILED, path)
	}

	// restore Name with relative pathing for preserving directory tree
	header.Name, err = filepath.Rel(me.Raw, path)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PATH_REL_FAILED, path)
	}

	// perform write
	err = me.writer.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_HEADER_WRITE_FAILED,
			path,
			err,
		)
	}

	// open file to read
	f, err = os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_READ_FAILED,
			path,
			err,
		)
	}
	defer f.Close()

	// copy file data
	_, err = io.Copy(me.writer, f)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_WRITE_FAILED,
			path,
			err,
		)
	}

	return nil
}

// Extract is to extract a `.tar.gz` archived file into the data directory.
//
// It called Archiver.Sanitize() internally.
func (me *Archiver) Extract() (err error) {
	var f *os.File
	var gz *gzip.Reader

	err = me.Sanitize()
	if err != nil {
		return err
	}

	// create directory if requested
	if me.CreateDirectory {
		err = os.MkdirAll(filepath.Dir(me.Raw), PERMISSION_DIR)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_CREATE_FAILED,
				me.Raw,
			)
		}
	}

	// overwrite Raw if requested
	_, err = os.Stat(me.Raw)
	if err == nil {
		if me.Overwrite {
			return fmt.Errorf("%s: %s", ERROR_DEST_EXISTS, me.Raw)
		}

		err = os.RemoveAll(me.Raw)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_OVERWRITE_FAILED,
				me.Raw,
			)
		}
	}

	// open Archive for extractions
	f, err = os.Open(me.Archive)
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_DEST_CREATE_FAILED,
			me.Archive,
		)
	}
	defer f.Close()

	gz, err = gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_ARCHIVE_READ_FAILED, err)
	}
	defer gz.Close()

	me.reader = tar.NewReader(gz)

	return me.extract()
}

func (me *Archiver) extract() (err error) {
	var path string
	var ok bool
	var header *tar.Header

	now := time.Now()

	for {
		header, err = me.reader.Next()
		switch err {
		case io.EOF:
			return nil
		case nil:
		default:
			return fmt.Errorf("%s: %s",
				ERROR_ARCHIVE_FILE_HEADER_FAILED,
				err,
			)
		}

		path, ok = me.sanitizeExtractedPath(me.Raw, header.Name)
		if !ok {
			return fmt.Errorf("%s: %s",
				ERROR_ARCHIVE_TAINTED_HEADER_PATH,
				header.Name,
			)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			err = me.extractDirectory(path, header, now)
			if err != nil {
				return err
			}
		case tar.TypeReg:
			err = me.extractRegular(path, header, now)
			if err != nil {
				return err
			}
		case tar.TypeSymlink:
			err = me.extractSymlink(path, header, now)
			if err != nil {
				return err
			}
		}
	}
}

func (me *Archiver) extractSymlink(path string,
	header *tar.Header, now time.Time) (err error) {
	err = os.Symlink(header.Linkname, path)
	if err != nil {
		return fmt.Errorf("%s (%s -> %s): %s",
			ERROR_SYMLINK_CREATE_FAILED,
			path,
			header.Linkname,
			err,
		)
	}

	return me.restoreMetadata(path, header, PERMISSION_FILE, now)
}

func (me *Archiver) extractDirectory(path string,
	header *tar.Header, now time.Time) (err error) {
	err = os.Mkdir(path, PERMISSION_DIR)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_ARCHIVE_FILE_HEADER_FAILED,
			err,
		)
	}

	return me.restoreMetadata(path, header, PERMISSION_DIR, now)
}

func (me *Archiver) extractRegular(path string,
	header *tar.Header, now time.Time) (err error) {
	var f *os.File

	err = os.MkdirAll(filepath.Dir(path), PERMISSION_DIR)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			ERROR_DEST_CREATE_FAILED,
			path,
			err,
		)
	}

	f, err = os.Create(path)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			ERROR_DEST_CREATE_FAILED,
			path,
			err,
		)
	}

	_, err = io.Copy(f, me.reader)
	_ = f.Sync()
	f.Close()
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			ERROR_FILE_WRITE_FAILED,
			path,
			err,
		)
	}

	return me.restoreMetadata(path, header, PERMISSION_FILE, now)
}

func (me *Archiver) restoreMetadata(path string,
	header *tar.Header, backupMode os.FileMode, now time.Time) (err error) {
	var aTime, mTime time.Time

	// restore chmod
	mode := os.FileMode(header.Mode)
	if mode < 0000 || mode > 0777 {
		mode = backupMode
	}

	err = os.Chmod(path, mode)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			ERROR_FILE_CHMOD_FAILED,
			path,
			err,
		)
	}

	// restore timestamp
	mTime = header.ModTime
	if mTime.IsZero() {
		mTime = now
	}

	aTime = header.AccessTime
	if aTime.IsZero() {
		aTime = now
	}

	switch header.Typeflag {
	case tar.TypeSymlink:
		err = oshelper.SymlinkChtimes(path, aTime, mTime)
	default:

		err = os.Chtimes(path, aTime, mTime)
	}

	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_FILE_CHTIMES_FAILED,
			err,
		)
	}

	return nil
}

func (me *Archiver) sanitizeExtractedPath(root, f string) (v string, ok bool) {
	// G305: Zip Slip Vulnerability
	v = filepath.Join(root, f)
	if strings.HasPrefix(v, filepath.Clean(root)) {
		return v, true
	}

	return "", false
}
