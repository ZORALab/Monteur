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

package checksum

import (
	"bytes"
	"crypto/md5" //nolint:gosec
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
)

// Hasher is the data structure to checksum a downloaded file.
//
// It is safe to create using the conventional `&checksum.Hasher{}` method.
// Hasher has safety functions to ensure it will be running smoothly.
//
// To use the `Hasher`, you need to call its setting methods for configurations.
type Hasher struct {
	hash    hash.Hash
	value   []byte
	healthy bool
}

// Compare is to use the `Hasher` and checksum with the parsed value.
//
// This function will call `IsHealthy()` function if the latter function was not
// executed beforehand. Should the `Hasher` is found not healthy, `Compare`
// function shall return an error.
//
// For positive matching value, the `ok` is set to `true` with no error.
//
// For negative matching value, the `ok` is set to `false` with no error.
//
// Should there be any error, the `ok` is always `false`.
func (hasher *Hasher) Compare(target io.Reader) (ok bool, err error) {
	// ensure the hasher is healthy before use
	if !hasher.healthy {
		err = hasher.IsHealthy()
		if err != nil {
			return false, err
		}
	}

	// consume the health status so that we do not re-use the same hasher
	// by accident
	hasher.healthy = false

	// copy the reader
	_, err = io.Copy(hasher.hash, target)
	if err != nil {
		return false, fmt.Errorf("%s: %s", ERROR_READ_FILE, err)
	}

	// compare results and return nil if true
	if bytes.Equal(hasher.value, hasher.hash.Sum(nil)) {
		return true, nil
	}

	// the value is mismatched so return an error
	return false, nil
}

// IsHealthy is a function to ensure the `Hasher` is ready for use.
//
// This function is designed for early checking and proper settings in certain
// implementations like long/large download time. The use is to remove all
// human and fixable error before sending it for the tough implementations which
// can take time and resources.
func (hasher *Hasher) IsHealthy() (err error) {
	defer func() {
		if recover() != nil {
			err = fmt.Errorf(ERROR_INIT_FAILED)
		}
	}()

	// try resetting the hasher
	hasher.hash.Reset()

	// verify user had assigned the checksum value
	if len(hasher.value) == 0 && err == nil {
		err = fmt.Errorf(ERROR_MISSING_VALUE)
	}

	if err == nil {
		hasher.healthy = true
	}

	return err
}

// ParseBase64 is for parsing a standard base64 string checksum value.
//
// It shall return error as value should there be any decoding problem occurs.
// Otherwise, it will always be nil.
func (hasher *Hasher) ParseBase64(raw string) (err error) {
	hasher.value, err = base64.StdEncoding.DecodeString(raw)
	if err == nil {
		return nil
	}

	hasher.value = nil

	return fmt.Errorf(ERROR_PARSE_BAD)
}

// ParseBase64URL is for parsing an URL-base64 encoded checksum `string` value.
//
// The difference from `ParseBase64` function is the checksum value has unpadded
// alternate base64 encoding usually used in URLs and filenames.
//
// It shall return error as value should there be any decoding problem occurs.
// Otherwise, it will always be nil.
func (hasher *Hasher) ParseBase64URL(raw string) (err error) {
	hasher.value, err = base64.URLEncoding.DecodeString(raw)
	if err == nil {
		return nil
	}

	hasher.value = nil

	return fmt.Errorf(ERROR_PARSE_BAD)
}

// ParseHex is for parsing a hexadecimal encoded checksum `string` value.
//
// It shall return error as value should there be any decoding problem occurs.
// Otherwise, it will always be nil.
func (hasher *Hasher) ParseHex(raw string) (err error) {
	hasher.value, err = hex.DecodeString(raw)
	if err == nil {
		return nil
	}

	hasher.value = nil

	return fmt.Errorf(ERROR_PARSE_BAD)
}

// ParseBytes is for parsing raw checksum value in `[]byte` data type.
//
// The data in the byte slice **SHALL be the RAW value** without any encoding
// like Base64 or hex.
func (hasher *Hasher) ParseBytes(raw []byte) (err error) {
	if len(raw) == 0 {
		return fmt.Errorf(ERROR_PARSE_BAD)
	}

	hasher.value = raw

	return nil
}

// SetAlgo is to set the hasher algorithm based on supported list of HashType.
//
// See "Supported Hashing Algorithms" Constants list for supported algorithms.
func (hasher *Hasher) SetAlgo(label HashType) (err error) {
	switch label {
	case HASHER_MD5:
		hasher.hash = md5.New() //nolint:gosec
	case HASHER_SHA256:
		hasher.hash = sha256.New()
	case HASHER_SHA512:
		hasher.hash = sha512.New()
	case HASHER_UNSET:
		fallthrough
	default:
		return fmt.Errorf(ERROR_ALGO_BAD)
	}

	return nil
}
