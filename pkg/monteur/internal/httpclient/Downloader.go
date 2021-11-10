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

package httpclient

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/checksum"
)

type Downloader struct {
	request   *http.Request
	indicator *indicator
	checksum  *checksum.Hasher

	// Headers are the additional headers to add into the request.
	Headers map[string]string

	// HandleError is a function handler for handling error the user way.
	HandleError func(err error)

	// HandleProgress is a function handler for handling download progress.
	HandleProgress func(downloaded int64, total int64)

	// HandleSuccess is a function handler to execute after the download
	// is completed.
	HandleSuccess func()

	// HandleRedirect is a function handler to execute upon receiving a
	// redirect instruction from the server.
	HandleRedirect func(req *http.Request, via []*http.Request) error

	// Destination is the directory + (optionally) filename for file saving.
	//
	// If filename is not given. Downloader tries to determine the filename
	// from header or URL.
	//
	// If it still fails, an error is raised. For such edge cases, it's
	// better you provide an optional filename on your next retry.
	Destination string

	// Timeout sets the Downloader to timeout when stuck with waiting.
	//
	// If zero or the value is not provided, Downloader will use the default
	// TIMEOUT value (defined in package constant).
	Timeout time.Duration

	// Overwrite is the decision for Downloader to overwrite existing
	// destination file if found.
	//
	// This variable is ignored when the destination is missing which is
	// a fresh download.
	//
	// By default (false), Downloader tries to detect local resumable
	// artifacts and tries to resume it. Otherwise, it will throw an error.
	Overwrite bool

	// CreateDirectory is the decision to create directory when not exist.
	//
	// By default (false), Downloader will not create the directory for the
	// destination pathing and will throw an error.
	CreateDirectory bool

	// RetainOnError is the decision to retain download artifact post error.
	//
	// By default (false), Downloader will not retain the artifact and
	// delete immediately after any error occurred (e.g. mismatched
	// checksum).
	RetainOnError bool
}

// Download is to initiate a download with a given URL to destination location.
//
// Download accepts a context for cancellation or timeout controls over it. The
// `ctx` shall not be nil or by minimum, provide context.Background() instead.
//
// Checksum is the data structure for checking the file integrity.
//
// This data is optional. When provided, Downloader will perform checksum before
// releasing the process as completed. If there is an error with the checksum,
// the downloaded file shall be deleted by default (unless
// Checksum.RetainOnError is set to `true`).
func (d *Downloader) Download(ctx context.Context,
	method string,
	urlstr string,
	hasher *checksum.Hasher) {
	var response *http.Response
	var client *http.Client
	var inReader io.Reader
	var err error

	client, err = d.init(ctx, method, urlstr, hasher)
	if err != nil {
		d.handleError(err)
		return
	}

	err = d.obtainMetadata(client)
	if err != nil {
		d.handleError(err)
		return
	}

	err = d.tryResume()
	if err != nil {
		d.handleError(err)
		return
	}

	// open the destination file for download
	f, err := os.OpenFile(d.Destination+DOWNLOAD_EXTENSION,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		FILE_PERMISSION,
	)
	if err != nil {
		d.handleError(err)
		return
	}

	defer func() {
		f.Close()

		if err != nil {
			return
		}

		err = d.checksumArtifact()
		if err != nil {
			d.handleError(err)
			return
		}

		err = d.renameArtifact()
		if err != nil {
			d.handleError(err)
			return
		}

		d.handleSuccess()
	}()

	// make the download request
	response, err = client.Do(d.request)
	if err != nil {
		d.handleError(err)
		return
	}

	defer func() {
		response.Body.Close()
	}()

	// read the content
	inReader = io.TeeReader(response.Body, d.indicator)

	_, err = io.Copy(f, inReader)
	if err != nil {
		d.handleError(err)
	}
}

func (d *Downloader) init(ctx context.Context,
	method string,
	urlStr string,
	hasher *checksum.Hasher) (client *http.Client, err error) {
	// validate saving location
	if d.Destination == "" {
		d.Destination = "."
	}

	d.Destination, err = filepath.Abs(d.Destination)
	if err != nil {
		return nil, fmt.Errorf("%s: %s",
			ERROR_PATH_INVALID,
			d.Destination,
		)
	}

	// validate request method
	if method == "" {
		return nil, fmt.Errorf(ERROR_METHOD_MISSING)
	}

	// configure progress bar indicator
	d.indicator = &indicator{
		total:          0,
		downloaded:     0,
		handleProgress: d.handleProgress,
	}

	// inspect checksum is usable before acceptance
	if hasher != nil {
		err = hasher.IsHealthy()
		if err != nil {
			return nil, fmt.Errorf("%s: %s",
				ERROR_HASHER_UNHEALTHY,
				err,
			)
		}

		d.checksum = hasher
	}

	// configure new http request for the downloader
	d.request, err = http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", ERROR_REQUEST_INIT_FAILED, err)
	}

	// add headers if available
	if d.Headers != nil {
		for k, v := range d.Headers {
			d.request.Header.Set(k, v)
		}
	}

	// add given context into request
	d.request = d.request.WithContext(ctx)

	// set timeout to default TIMEOUT seconds for bad or 0 value
	if d.Timeout <= 0 {
		d.Timeout = TIMEOUT * time.Second
	}

	// setup the http client
	client = &http.Client{
		Timeout: d.Timeout,
	}

	// insert redirect function if available
	if d.HandleRedirect != nil {
		client.CheckRedirect = d.HandleRedirect
	}

	return client, nil
}

func (d *Downloader) checksumArtifact() (err error) {
	var f *os.File
	var ok bool

	if d.checksum == nil {
		return nil
	}

	// obtain the checksum value from the downloaded file
	f, err = os.Open(d.Destination + DOWNLOAD_EXTENSION)
	if err != nil {
		return fmt.Errorf("%s: %s",
			ERROR_CHECKSUM_BAD_FILE,
			d.Destination+DOWNLOAD_EXTENSION,
		)
	}

	// initiate checksum comparison
	ok, err = d.checksum.Compare(f)
	f.Close()

	// no error
	if err == nil {
		if ok {
			return nil // matching checksum
		}

		return fmt.Errorf(ERROR_CHECKSUM_MISMATCHED)
	}

	err = fmt.Errorf("%s: [%s] %s",
		ERROR_CHECKSUM,
		d.Destination+DOWNLOAD_EXTENSION,
		err,
	)

	if d.RetainOnError {
		return err // requested to retain artifact
	}

	// otherwise attempting to deleting it
	if os.Remove(d.Destination+DOWNLOAD_EXTENSION) != nil {
		err = fmt.Errorf("%s: %s", ERROR_CHECKSUM_DELETE_FAILED, err)
	}

	return err
}

func (d *Downloader) renameArtifact() (err error) {
	err = os.Rename(d.Destination+DOWNLOAD_EXTENSION, d.Destination)
	if err != nil {
		err = fmt.Errorf("%s: %s", ERROR_FILE_RENAMED_FAILED, err)
	}

	return err
}

func (d *Downloader) tryResume() (err error) {
	var fi os.FileInfo

	// check for any existing download artifacts
	fi, err = os.Stat(d.Destination + DOWNLOAD_EXTENSION)

	switch {
	case os.IsNotExist(err):
		return nil
	case err != nil:
		return fmt.Errorf("%s: %s", ERROR_FILE_STAT, err)
	case !fi.Mode().IsRegular():
		return nil
	}

	// if overwrite, execute overwrite and let go
	if d.Overwrite {
		err = os.Remove(d.Destination + DOWNLOAD_EXTENSION)
		if err != nil {
			return fmt.Errorf("%s: %s",
				ERROR_FILE_OVERWRITE_FAILED,
				err,
			)
		}

		return nil
	}

	// get current download size
	d.indicator.downloaded = fi.Size()

	// reject invalid data
	if d.indicator.downloaded >= d.indicator.total {
		d.indicator.downloaded = 0
		return nil
	}

	// successfully resumed
	d.request.Header.Set("Range",
		fmt.Sprintf("bytes=%d-", d.indicator.downloaded),
	)

	return nil
}

func (d *Downloader) obtainMetadata(client *http.Client) (err error) {
	var response *http.Response
	var length, disposition string

	// request header and extract all target metadata
	response, err = client.Head(d.request.URL.String())
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_REQUEST_FAILED, err)
	}

	length = response.Header.Get("content-length")
	disposition = response.Header.Get("content-Disposition")
	response.Body.Close()

	err = d.processSize(length)
	if err != nil {
		return err
	}

	err = d.processFilepath(disposition)
	if err != nil {
		return err
	}

	return nil
}

func (d *Downloader) processFilepath(disposition string) (err error) {
	var params map[string]string
	var fi os.FileInfo
	var target *url.URL

	fi, err = os.Stat(d.Destination)

	switch {
	case os.IsNotExist(err):
		// new full directory + filename
		if filepath.Ext(filepath.Base(d.Destination)) != "" {
			goto create_directory
		}

		// new full directory: proceed for getting filename
	case fi.Mode().IsDir():
		// existing full directory: proceed for getting filename
	case fi.Mode().IsRegular():
		// existing full directory + filename
		goto check_overwrite
	default:
		// non-proceeding errors
		return fmt.Errorf("%s: %s", ERROR_FILE_STAT, err)
	}

	// parsing filenames from header first then url as last resort
	_, params, err = mime.ParseMediaType(disposition)
	if err != nil {
		goto get_filename_from_url
	}

	if params["filename"] != "" {
		d.Destination = filepath.Join(d.Destination, params["filename"])
		goto create_directory
	}

get_filename_from_url:
	target, err = url.Parse(d.request.URL.String())

	if err != nil {
		return fmt.Errorf(ERROR_FILENAME_MISSING)
	}

	d.Destination = filepath.Join(d.Destination, path.Base(target.Path))

create_directory:
	// destination is now directory + filename
	err = d.createDirectory(filepath.Dir(d.Destination))
	if err != nil {
		return err
	}

	// destination is now directory (exist) + filename
	fi, err = os.Stat(d.Destination)

	switch {
	case os.IsNotExist(err):
		return nil
	case err != nil:
		return fmt.Errorf("%s: %s", ERROR_PATH_INVALID, err)
	case fi.Mode().IsRegular():
		// permit to proceed for checking overwrite setting
	default:
		return fmt.Errorf("%s: %s", ERROR_PATH_INVALID, d.Destination)
	}

check_overwrite:
	// destination is now directory (exist) + filename (exist)
	if !d.Overwrite {
		return fmt.Errorf("%s: %s", ERROR_FILE_EXISTS, d.Destination)
	}

	return nil
}

func (d *Downloader) createDirectory(pathing string) (err error) {
	_, err = os.Stat(pathing)
	if os.IsNotExist(err) && !d.CreateDirectory {
		return fmt.Errorf("%s: %s", ERROR_PATH_MISSING, pathing)
	}

	err = os.MkdirAll(pathing, DIR_PERMISSION)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_PATH_INVALID, err)
	}

	return nil
}

func (d *Downloader) processSize(length string) (err error) {
	if length == "" {
		return fmt.Errorf(ERROR_FILESIZE_MISSING)
	}

	d.indicator.total, err = strconv.ParseInt(length, 10, 64)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FILESIZE_MISSING, err)
	}

	return nil
}

func (d *Downloader) handleError(err error) {
	if d.HandleError != nil {
		d.HandleError(err)
	}
}

func (d *Downloader) handleSuccess() {
	if d.HandleSuccess != nil {
		d.HandleSuccess()
	}
}

func (d *Downloader) handleProgress(downloaded int64, total int64) {
	if d.HandleProgress != nil {
		d.HandleProgress(downloaded, total)
	}
}

// Context is to generate the background context for Downloader.
//
// It is made available so that user does not need to import "context" package
// just to do simple downloading.
func Context() context.Context {
	return context.Background()
}
