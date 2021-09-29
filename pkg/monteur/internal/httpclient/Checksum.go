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
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
)

// Checksum is the data structure to checksum a downloaded file.
//
// You need to provide all necessary values for the checksum to work.
type Checksum struct {
	// Hash is the checksum algorithm for calculating the file checksum.
	Hash hash.Hash

	value []byte

	// RetainOnError is the decision for retaining the mismatched file.
	//
	// By default, should a checksum mismatched or error is found, it shall
	// be removed automatically.
	RetainOnError bool
}

func (hasher *Checksum) init() (err error) {
	defer func() {
		if recover() != nil {
			err = fmt.Errorf("%s%s",
				ERROR_TAG,
				ERROR_CHECKSUM_INVALID,
			)
		}
	}()

	// try reset the hasher
	hasher.Hash.Reset()

	// verify user had assigned the checksum value
	if len(hasher.value) == 0 && err == nil {
		err = fmt.Errorf("%s%s", ERROR_TAG, ERROR_CHECKSUM_BAD)
	}

	return err
}

func (hasher *Checksum) compare() (err error) {
	if bytes.Equal(hasher.value, hasher.Hash.Sum(nil)) {
		return nil
	}

	return fmt.Errorf("%s%s", ERROR_TAG, ERROR_CHECKSUM_MISMATCHED)
}

// ParseBase64 is for parsing a standard base64 string checksum value.
func (hasher *Checksum) ParseBase64(raw string) (err error) {
	hasher.value, err = base64.StdEncoding.DecodeString(raw)
	if err == nil {
		return nil
	}

	hasher.value = nil

	return fmt.Errorf("%s%s", ERROR_TAG, ERROR_CHECKSUM_BAD)
}

// ParseBase64URL is for parsing a URL-friendly base64 string checksum value.
//
// The difference is the string value has unpadded alternate base64 encoding
// typically used in URLs and filenames.
func (hasher *Checksum) ParseBase64URL(raw string) (err error) {
	hasher.value, err = base64.URLEncoding.DecodeString(raw)
	if err == nil {
		return nil
	}

	hasher.value = nil

	return fmt.Errorf("%s%s", ERROR_TAG, ERROR_CHECKSUM_BAD)
}

// ParseHex is for parsing a hexadecimal encoded string checksum value.
func (hasher *Checksum) ParseHex(raw string) (err error) {
	hasher.value, err = hex.DecodeString(raw)
	if err == nil {
		return nil
	}

	hasher.value = nil

	return fmt.Errorf("%s%s", ERROR_TAG, ERROR_CHECKSUM_BAD)
}

// ParseBytes is for parsing a byte slice checksum value.
//
// The data in the byte slice **SHALL be the RAW value** without any additional
// encoding like Base64 or hex.
func (hasher *Checksum) ParseBytes(raw []byte) (err error) {
	if len(raw) == 0 {
		return fmt.Errorf("%s%s", ERROR_TAG, ERROR_CHECKSUM_BAD)
	}

	hasher.value = raw

	return nil
}
