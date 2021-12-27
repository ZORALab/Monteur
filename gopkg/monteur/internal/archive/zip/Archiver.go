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

package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/archive"
)

type CompressionID uint16

const (
	COMPRESS_NONE    uint16 = 0
	COMPRESS_DEFLATE uint16 = 8
)

const (
	EXTENSION = ".zip"
)

type Archiver struct {
	writer *zip.Writer

	// Archive is the filepath to the `.zip` archive file.
	//
	// The value is **STRICTLY** a filepath with filename and `.zip` file
	// extension.
	Archive string

	// Raw is the directory path for the uncompressed data directory.
	//
	// The value **MUST** be a directory holding the archive content for
	// compression or empty directory for decompression.
	Raw string

	// Compression is the compression methods.
	Compression CompressionID

	// CreateDirectory decides on creating missing directory for Archive.
	//
	// When set to `true`, Archiver shall create the directory for housing
	// the `.tar.gz` archived file if it is not exist.
	//
	// Default is returning an error (`false`).
	CreateDirectory bool

	// Overwrite decides on overwriting RAW data directory.
	//
	// When set to `true`, Archiver shall delete and recreate directory path
	// given in RAW field.
	//
	// Default is returning an error (`false`).
	Overwrite bool

	// Relief Extension decides to relax file extension checking.
	//
	// When set to `true`, Archiver shall not throw an error when checking
	// Archive for strict `.zip` file extension.
	//
	// Default is returning an error (`false`).
	ReliefExtension bool

	// FollowSymlink decides to resolve symlink and archive package file.
	//
	// When set to `true`, Archiver shall resolve the symlink recursively
	// until the target file is saved.
	//
	// Default `false`) is to save a relative symlink to the `Raw`
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
// It calls `Archiver.Sanitize()` internally.
func (me *Archiver) Compress() (err error) {
	var f *os.File

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
			err,
		)
	}
	defer f.Close()

	me.writer = zip.NewWriter(f)
	defer me.writer.Close()

	return filepath.Walk(me.Raw, me.compress)
}

func (me *Archiver) compress(path string, info os.FileInfo, err error) error {
	var mode os.FileMode

	if err != nil {
		return fmt.Errorf("%s (%s): %s",
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
	var header *zip.FileHeader
	var target string
	var targetInfo os.FileInfo
	var zipFile io.Writer

	target, targetInfo, err = archive.EvalSymlink(me.Raw, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// check follow decision and resolve recursively if set
	if me.FollowSymlink {
		return me.compress(target, targetInfo, nil)
	}

	// create header from info
	header, err = zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_READ,
			path,
			err,
		)
	}

	// configure compression mode
	header.Method = uint16(me.Compression)

	// configure pathing
	header.Name, err = archive.RelPath(me.Raw, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// write symlink
	zipFile, err = me.writer.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_WRITE,
			path,
			err,
		)
	}

	_, err = zipFile.Write([]byte(target))
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_SYMLINK_CREATE,
			path,
			err,
		)
	}

	return nil
}

func (me *Archiver) compressDir(path string, info os.FileInfo) (err error) {
	var header *zip.FileHeader

	// create header from info
	header, err = zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_READ,
			path,
			err,
		)
	}

	// configure compression mode
	header.Method = uint16(me.Compression)

	// configure pathing
	header.Name, err = archive.RelPath(me.Raw, path)
	if err != nil {
		return err //nolint:wrapcheck
	}
	header.Name += archive.PATH_SEPARATOR

	// write to archive
	_, err = me.writer.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("%s: %s",
			archive.ERROR_HEADER_WRITE,
			path,
		)
	}

	return nil
}

func (me *Archiver) compressRegular(path string, info os.FileInfo) (err error) {
	var f *os.File
	var header *zip.FileHeader
	var zipFile io.Writer

	// create header from info
	header, err = zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_READ,
			path,
			err,
		)
	}

	// configure compression mode
	header.Method = uint16(me.Compression)

	// configure pathing
	header.Name, err = archive.RelPath(me.Raw, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// open file for reading
	f, err = os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			archive.ERROR_FILE_READ,
			path,
			err,
		)
	}
	defer f.Close()

	// create file from header
	zipFile, err = me.writer.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_HEADER_WRITE,
			path,
			err,
		)
	}

	// copy file data
	_, err = io.Copy(zipFile, f)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			archive.ERROR_FILE_WRITE,
			path,
			err,
		)
	}

	return nil
}

// Extract is to extract a `.zip` Archiver File into the Raw data directory.
//
// It calls `Archiver.Sanitize()` internally.
func (me *Archiver) Extract() (err error) {
	var reader *zip.ReadCloser
	var now time.Time

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
	reader, err = zip.OpenReader(me.Archive)
	if err != nil {
		return fmt.Errorf("%s: %s",
			archive.ERROR_ARCHIVE_READ,
			me.Archive,
		)
	}
	defer reader.Close()

	// iterate through file in the archives
	now = time.Now()
	for _, f := range reader.File {
		err = me.extract(f, now)
		if err != nil {
			return err
		}
	}

	return nil
}

func (me *Archiver) extract(header *zip.File, now time.Time) (err error) {
	var path string
	var mode fs.FileMode

	path, err = archive.SanitizeCompressionPath(me.Raw, header.Name)
	if err != nil {
		return err //nolint:wrapcheck
	}

	fi := header.FileInfo()
	mode = fi.Mode()
	switch {
	case mode.IsRegular():
		err = me.extractRegular(path, header, now)
	case mode.IsDir():
		err = me.extractDir(path, header, now)
	case mode&fs.ModeSymlink != 0:
		err = me.extractSymlink(path, header, now)
	default:
		err = fmt.Errorf("%s (%s): %s",
			archive.ERROR_FILE_UNSUPPORTED,
			path,
			header.Name,
		)
	}

	return err
}

func (me *Archiver) extractSymlink(path string,
	header *zip.File, now time.Time) (err error) {
	var mTime time.Time
	var buf *strings.Builder
	var target string
	var r io.ReadCloser

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

	// get symlink target from body
	r, err = header.Open()
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_SYMLINK_READ,
			path,
			err,
		)
	}

	buf = new(strings.Builder)
	err = archive.ExtractCopy(buf, r, archive.COPY_SIZE, path)
	r.Close()
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_SYMLINK_READ,
			path,
			err,
		)
	}
	target = buf.String()

	// check target path is valid
	target, err = archive.SanitizeCompressionPath(me.Raw, target)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// create symlink
	err = archive.CreateSymlink(target, path)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// restore matadata
	mTime = header.Modified
	if mTime.IsZero() {
		mTime = now
	}

	return archive.RestoreMetadata(path, //nolint:wrapcheck
		now,
		mTime,
		archive.PERMISSION_DIR,
		true,
	)
}

func (me *Archiver) extractDir(path string,
	header *zip.File, now time.Time) (err error) {
	var mTime time.Time

	// overwrite Raw if requested
	err = archive.Overwrite(path, me.Overwrite)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// create the target directory
	err = archive.MkdirAll(path, archive.PERMISSION_DIR, true)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// restore matadata
	mTime = header.Modified
	if mTime.IsZero() {
		mTime = now
	}

	return archive.RestoreMetadata(path, //nolint:wrapcheck
		now,
		mTime,
		archive.PERMISSION_DIR,
		false,
	)
}

func (me *Archiver) extractRegular(path string,
	header *zip.File, now time.Time) (err error) {
	var mTime time.Time
	var r io.ReadCloser
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

	r, err = header.Open()
	if err != nil {
		return fmt.Errorf("%s (%s): %s",
			archive.ERROR_FILE_READ,
			path,
			err,
		)
	}
	defer r.Close()

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
	err = archive.ExtractCopy(f, r, archive.COPY_SIZE, path)
	_ = f.Sync()
	f.Close()
	if err != nil {
		return err //nolint:wrapcheck
	}

	// restore metadata
	mTime = header.Modified
	if mTime.IsZero() {
		mTime = now
	}

	return archive.RestoreMetadata(path, //nolint:wrapcheck
		now,
		mTime,
		archive.PERMISSION_FILE,
		false,
	)
}
