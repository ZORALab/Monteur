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

// HashType is the ID number for hashing algorithm
//
// To ensure the values are specific to checksum package only, we declare a
// new type over the intended `uint` data type.
type HashType uint

// Supported Hashing Algorithm is the list of constants ID for selecting algo.
//nolint:stylecheck,revive
const (
	HASHER_UNSET  HashType = 0
	HASHER_MD5    HashType = 1
	HASHER_SHA256 HashType = 2
	HASHER_SHA512 HashType = 3
)
