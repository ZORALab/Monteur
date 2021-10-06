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
)

type Processor struct {
	reader *tar.Reader
	writer *tar.Writer

	// Archive is the pathing to the tar.gz archive file.
	//
	// The value **MUST** be a file and not a directory.
	//
	// You **MUST** supply `.tar.gz` file extension in this pathing.
	// Otherwise, it will throw an error.
	Archive string

	// Raw is the pathing for the uncompressed data directory.
	//
	// The value **MUST** be a directory holding the archive structure or
	// an empty directory for decompression.
	Raw string

	// CreateDirectory is a decision to create directory for Archive.
	//
	// When true, Processor will create the directory for housing Archive
	// content if it does not exist.
	//
	// Otherwise, the default is returning an error.
	CreateDirectory bool

	// Overwrite is a decision to overwrite the destination file/directory.
	//
	// When true, Processor shall delete and re-create the destination.
	//
	// Otherwise, the default is returning an error.
	Overwrite bool

	// ReliefExtension is a decision not to strictly check file extension.
	//
	// When true, Processor shall not throw an error when archive pathing
	// does not have `.tar.gz` file extension.
	//
	// Otherwise, the default is returning an error.
	ReliefExtension bool

	isSanitized bool
}

// Init initializes the tar.gz Processor with the given Processor parameters.
//
// This function setups and sanitizes the Processor given parameters, generates
// necessary components for other functions' comsumption.
func (tg *Processor) Init() (err error) {
	var info os.FileInfo
	var path, filename, extension string

	// reset sanitization status
	tg.isSanitized = false

	// sanitize Archive
	tg.Archive, err = filepath.Abs(tg.Archive)

	filename = filepath.Base(tg.Archive)
	extension = filepath.Ext(filename)
	filename = filename[0 : len(filename)-len(extension)]
	extension = filepath.Ext(filename) + extension

	switch {
	case err != nil:
		return fmt.Errorf("%s: %s", ERROR_PATH_ABS_FAILED, tg.Archive)
	case tg.Archive == "":
		return fmt.Errorf("%s: %s", ERROR_PATH_EMPTY, "Archive")
	case extension != ARCHIVE_EXTENSION && !tg.ReliefExtension:
		return fmt.Errorf("%s: %s", ERROR_EXTENSION_MISSING, tg.Archive)
	}

	info, err = os.Stat(tg.Archive)
	switch {
	case err == nil && !info.IsDir(), os.IsExist(err), os.IsNotExist(err):
		// does not matter as compress() and extract() will check
		// it on demand.
	case info.IsDir():
		return fmt.Errorf("%s: %s", ERROR_PATH_IS_DIR, tg.Archive)
	default:
		return fmt.Errorf("%s: %s", ERROR_PATH_ARCHIVE, err)
	}

	path = filepath.Dir(tg.Archive)
	info, err = os.Stat(path)
	switch {
	case os.IsNotExist(err):
		if !tg.CreateDirectory {
			return fmt.Errorf("%s: %s",
				ERROR_PATH_DIR_MISSING,
				path,
			)
		}
	case !info.IsDir():
		return fmt.Errorf("%s: %s", ERROR_PATH_NOT_DIR, path)
	}

	// sanitize Raw
	tg.Raw, err = filepath.Abs(tg.Raw)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PATH_ABS_FAILED, tg.Raw)
	}

	info, err = os.Stat(tg.Raw)
	switch {
	case os.IsNotExist(err):
		return fmt.Errorf("%s: %s", ERROR_PATH_EMPTY, tg.Raw)
	case !info.IsDir():
		return fmt.Errorf("%s: %s", ERROR_PATH_NOT_DIR, tg.Raw)
	}

	// successful initialization
	tg.isSanitized = true

	return nil
}

// Compress is to compress the given Raw directory into the Archive file.
//
// pathing is optional where it was facilitated for specific files or
// directories instead of the entire Destination directory. If it is empty,
// the entire Destination directory will be used.
func (tg *Processor) Compress(pathing ...string) (err error) {
	if !tg.isSanitized {
		err = tg.Init()
		if err != nil {
			return err
		}
	}

	// sanitize destination
	_, err = os.Stat(tg.Archive)
	if !os.IsNotExist(err) {
		if !tg.Overwrite {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_EXISTS,
				tg.Archive,
			)
		}

		err = os.RemoveAll(tg.Archive)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_OVERWRITE_FAILED,
				tg.Archive,
			)
		}
	}

	// create directory if requested
	if tg.CreateDirectory {
		err = os.MkdirAll(filepath.Dir(tg.Archive), DIR_PERMISSION)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_CREATE_FAILED,
				tg.Archive,
			)
		}
	}

	// create destination file
	f, err := os.Create(tg.Archive)
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_DEST_CREATE_FAILED,
			tg.Archive,
		)
	}
	defer f.Close()

	// create tar and gz writers
	gz := gzip.NewWriter(f)
	defer gz.Close()
	tg.writer = tar.NewWriter(gz)
	defer tg.writer.Close()

	// walk through and compress
	return filepath.Walk(tg.Raw, tg._compressFx) //nolint:wrapcheck
}

func (tg *Processor) _compressFx(path string,
	info os.FileInfo, err error) error {
	// return error if found
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_READ_FAILED,
			path,
			err,
		)
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILE_HEADER_READ_FAILED, path)
	}

	// set the name as fullpath to preserve directory name
	header.Name = path

	// perform write
	err = tg.writer.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_HEADER_WRITE_FAILED,
			path,
			err,
		)
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_READ_FAILED,
			path,
			err,
		)
	}
	defer f.Close()

	_, err = io.Copy(tg.writer, f)
	if err != nil {
		return fmt.Errorf("%s: (%s) %s",
			ERROR_FILE_WRITE_FAILED,
			path,
			err,
		)
	}

	return nil
}

// Extract is to decompress the given Archive file back into Raw directory.
//
//
// pathing is optional where it was facilitated for specific files or
// directories instead of the entire Destination directory. If it is empty,
// the entire Destination directory will be used.
func (tg *Processor) Extract(pathing ...string) (err error) {
	if !tg.isSanitized {
		err = tg.Init()
		if err != nil {
			return err
		}
	}

	// sanitize destination
	_, err = os.Stat(tg.Raw)
	if err != nil && !os.IsNotExist(err) {
		if !tg.Overwrite {
			return fmt.Errorf("%s: %s", ERROR_DEST_EXISTS, tg.Raw)
		}

		err = os.RemoveAll(tg.Raw)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_OVERWRITE_FAILED,
				tg.Raw,
			)
		}
	}

	// create directory if requested
	if tg.CreateDirectory {
		err = os.MkdirAll(filepath.Dir(tg.Raw), DIR_PERMISSION)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_DEST_CREATE_FAILED,
				tg.Archive,
			)
		}
	}

	// open destination file
	f, err := os.Open(tg.Archive)
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_DEST_CREATE_FAILED,
			tg.Archive,
		)
	}
	defer f.Close()

	// create tar and gz reader
	gz, err := gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_ARCHIVE_READ_FAILED, err)
	}
	defer gz.Close()

	tg.reader = tar.NewReader(gz)

	return tg._extractFiles()
}

func (tg *Processor) _extractFiles() (err error) {
	var path string
	var ok bool

	for {
		header, err := tg.reader.Next()
		if err == nil {
			path, ok = sanitizeArchivePath(tg.Raw, header.Name)
			if !ok {
				return fmt.Errorf("%s: %s",
					ERROR_ARCHIVE_TAINTED_HEADER_PATH,
					header.Name)
			}
		}

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return fmt.Errorf("%s: %s",
				ERROR_ARCHIVE_FILE_HEADER_FAILED,
				err,
			)
		case header.Typeflag == tar.TypeDir:
			err = os.Mkdir(path, DIR_PERMISSION)
			if err != nil {
				return fmt.Errorf("%s: %s",
					ERROR_ARCHIVE_FILE_HEADER_FAILED,
					err,
				)
			}
		case header.Typeflag == tar.TypeReg:
			f, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("%s: %s",
					ERROR_ARCHIVE_FILE_HEADER_FAILED,
					err,
				)
			}

			_, err = io.Copy(f, tg.reader)
			f.Close()

			if err != nil {
				return fmt.Errorf("%s: %s",
					ERROR_ARCHIVE_FILE_HEADER_FAILED,
					err,
				)
			}
		default:
			return fmt.Errorf("%s", ERROR_ARCHIVE_EXTRACT_FAILED)
		}
	}
}

// sanitize archive file pathing from
//   1. "G305: Zip Slip vulnerability"
func sanitizeArchivePath(d, t string) (v string, ok bool) {
	v = filepath.Join(d, t)
	if strings.HasPrefix(v, filepath.Clean(d)) {
		return v, true
	}

	return "", false
}
