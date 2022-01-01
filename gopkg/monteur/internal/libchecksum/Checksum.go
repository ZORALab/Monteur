// Copyright 2022 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2022 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
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

package libchecksum

import (
	"fmt"
	"io"
	"strings"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/checksum"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
)

// Hasher is an interface to isolate checkum package away from monteur
type Hasher interface {
	Hash(*[]byte) error
	HashFile(path string) error
	Compare(io.Reader) (bool, error)
	Reset() error
	IsHealthy() error
	ParseBase64(string) error
	ToBase64() (string, error)
	ParseBase64URL(string) error
	ToBase64URL() (string, error)
	ParseHex(string) error
	ToHex() (string, error)
	ParseBytes([]byte) error
	ToBytes() ([]byte, error)
	SetAlgo(checksum.HashType) error
}

// CreateChecksum create a ready-to-use checksum hasher based on given type.
//
// If the given type is unknown or something went wrong during reset, err shall
// be returned.
func CreateChecksum(checksumType string) (out Hasher, err error) {
	out = &checksum.Hasher{}

	switch strings.ToLower(checksumType) {
	case libmonteur.CHECKSUM_ALGO_SHA512:
		_ = out.SetAlgo(checksum.HASHER_SHA512)
	case libmonteur.CHECKSUM_ALGO_SHA256:
		_ = out.SetAlgo(checksum.HASHER_SHA256)
	case libmonteur.CHECKSUM_ALGO_MD5:
		_ = out.SetAlgo(checksum.HASHER_MD5)
	default:
		return nil, fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_CHECKSUM_ALGO_UNKNOWN,
			checksumType,
		)
	}

	err = out.Reset()
	if err != nil {
		return nil, fmt.Errorf("%s: %s",
			libmonteur.ERROR_CHECKSUM_BAD,
			err,
		)
	}

	return out, nil
}

// ArchiverID returns the Checksum.HashType of a given ID.
func ArchiverID(checksumType string) (out checksum.HashType, err error) {
	switch strings.ToLower(checksumType) {
	case libmonteur.CHECKSUM_ALGO_SHA512:
		return checksum.HASHER_SHA512, nil
	case libmonteur.CHECKSUM_ALGO_SHA256:
		return checksum.HASHER_SHA256, nil
	default:
		return 0, fmt.Errorf("%s: '%s'",
			libmonteur.ERROR_CHECKSUM_ALGO_UNKNOWN,
			checksumType,
		)
	}
}
