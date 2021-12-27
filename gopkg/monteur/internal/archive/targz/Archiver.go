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
	"time"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/archive"
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
	var extension string

	if !me.ReliefExtension {
		extension = EXTENSION
	}

	me.Archive, err = archive.SanitizeArchive(me.Archive, extension)
	if err != nil {
		return err //nolint:wrapcheck
	}

	me.Raw, err = archive.SanitizeRaw(me.Raw)
	if err != nil {
		return err //nolint:wrapcheck
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

	// check for overwrite
	err = archive.Overwrite(me.Archive, me.Overwrite)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// create base directory when requested
	err = archive.MkdirAll(filepath.Dir(me.Archive),
		archive.PERMISSION_DIR,
		true,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// begin archiving
	f, err = os.Create(me.Archive)
	if err != nil {
		return fmt.Errorf("%s: %s",
			archive.ERROR_ARCHIVE_CREATE,
			me.Archive,
		)
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
			archive.ERROR_FILE_READ,
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
		return fmt.Errorf("%s: %s",
			archive.ERROR_FILE_UNSUPPORTED,
			path,
		)
	}
}

func (me *Archiver) compressSymlink(path string, info os.FileInfo) (err error) {
	var header *tar.Header
	var targetInfo os.FileInfo
	var target string

	// ensure symlink is only targeting within data directory
	target, targetInfo, err = archive.EvalSymlink(me.Raw, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// check follow decision and resolve recursively if set
	if me.FollowSymlink {
		return me.compress(target, targetInfo, nil)
	}

	// create header from info
	header, err = tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_READ,
			path,
			err,
		)
	}

	// configure pathing
	header.Name, err = archive.RelPath(me.Raw, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// write symlink
	err = me.writer.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_WRITE,
			path,
			err,
		)
	}

	return nil
}

func (me *Archiver) compressDir(path string, info os.FileInfo) (err error) {
	var header *tar.Header

	// create header from info
	header, err = tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_READ,
			path,
			err,
		)
	}

	// configure pathing
	header.Name, err = archive.RelPath(me.Raw, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// write to archive
	err = me.writer.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_WRITE,
			path,
			err,
		)
	}

	return nil
}

func (me *Archiver) compressRegular(path string, info os.FileInfo) (err error) {
	var f *os.File
	var header *tar.Header

	// create header from info
	header, err = tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_READ,
			path,
			err,
		)
	}

	// configure pathing
	header.Name, err = archive.RelPath(me.Raw, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// open file to read
	f, err = os.Open(path)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_FILE_WRITE,
			path,
			err,
		)
	}
	defer f.Close()

	// create writer from header
	err = me.writer.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_WRITE,
			path,
			err,
		)
	}

	// copy file data
	_, err = io.Copy(me.writer, f)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_FILE_WRITE,
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
	err = archive.MkdirAll(filepath.Dir(me.Raw),
		archive.PERMISSION_DIR,
		me.CreateDirectory,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// open Archive for extractions
	f, err = os.Open(me.Archive)
	if err != nil {
		return fmt.Errorf("%s: %s",
			archive.ERROR_ARCHIVE_READ,
			me.Archive,
		)
	}
	defer f.Close()

	gz, err = gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("%s: %s",
			archive.ERROR_ARCHIVE_READ,
			err,
		)
	}
	defer gz.Close()

	me.reader = tar.NewReader(gz)

	return me.extract()
}

func (me *Archiver) extract() (err error) {
	var path string
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
				archive.ERROR_HEADER_READ,
				err,
			)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			err = me.extractDirectory(path, header, now)
		case tar.TypeReg:
			err = me.extractRegular(path, header, now)
		case tar.TypeSymlink:
			err = me.extractSymlink(path, header, now)
		default:
			err = fmt.Errorf("%s (%s): %s",
				archive.ERROR_FILE_UNSUPPORTED,
				path,
				header.Name,
			)
		}

		if err != nil {
			return err
		}
	}
}

func (me *Archiver) extractSymlink(path string,
	header *tar.Header, now time.Time) (err error) {
	var mode os.FileMode
	var aTime, mTime time.Time
	var target string

	// create housing directory
	err = archive.MkdirAll(filepath.Dir(path), archive.PERMISSION_DIR, true)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// overwrite Raw if requested
	err = archive.Overwrite(path, me.Overwrite)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// check target path is valid
	target, err = archive.SanitizeCompressionPath(me.Raw, header.Linkname)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// create symlink
	err = archive.CreateSymlink(target, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// restore metadata
	mTime = header.ModTime
	if mTime.IsZero() {
		mTime = now
	}

	aTime = header.AccessTime
	if aTime.IsZero() {
		aTime = now
	}

	mode = os.FileMode(header.Mode)
	if mode < 0000 || mode > 0777 {
		mode = archive.PERMISSION_FILE
	}

	return archive.RestoreMetadata(path, //nolint:wrapcheck
		aTime,
		mTime,
		mode,
		true,
	)
}

func (me *Archiver) extractDirectory(path string,
	header *tar.Header, now time.Time) (err error) {
	var mode os.FileMode
	var aTime, mTime time.Time

	// overwrite Raw if requested
	err = archive.Overwrite(path, me.Overwrite)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// create target directory
	err = archive.MkdirAll(path, archive.PERMISSION_DIR, true)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// restore metadata
	mTime = header.ModTime
	if mTime.IsZero() {
		mTime = now
	}

	aTime = header.AccessTime
	if aTime.IsZero() {
		aTime = now
	}

	mode = os.FileMode(header.Mode)
	if mode < 0000 || mode > 0777 {
		mode = archive.PERMISSION_DIR
	}

	return archive.RestoreMetadata(path, //nolint:wrapcheck
		aTime,
		mTime,
		mode,
		true,
	)
}

func (me *Archiver) extractRegular(path string,
	header *tar.Header, now time.Time) (err error) {
	var mode os.FileMode
	var aTime, mTime time.Time
	var f *os.File

	// create housing directory
	err = archive.MkdirAll(filepath.Dir(path), archive.PERMISSION_DIR, true)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// overwrite Raw if requested
	err = archive.Overwrite(path, me.Overwrite)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// create destination for write
	f, err = os.Create(path)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_FILE_WRITE,
			path,
			err,
		)
	}

	// perform file extraction
	err = archive.ExtractCopy(f, me.reader, archive.COPY_SIZE, path)
	_ = f.Sync()
	f.Close()
	if err != nil {
		return err //nolint:wrapcheck
	}

	// restore metadata
	mTime = header.ModTime
	if mTime.IsZero() {
		mTime = now
	}

	aTime = header.AccessTime
	if aTime.IsZero() {
		aTime = now
	}

	mode = os.FileMode(header.Mode)
	if mode < 0000 || mode > 0777 {
		mode = archive.PERMISSION_FILE
	}

	return archive.RestoreMetadata(path, //nolint:wrapcheck
		aTime,
		mTime,
		mode,
		true,
	)
}
