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

package secrets

import (
	"strings"
)

// QueryArray restores the array data structure from the decoded data.
//
// It takes a decoded data and the query `key` to the array (right before the
// numbering sequences). Example: for `development.databases.0`,
// `development.databases.1`, ..., the key is `development.databases`.
//
// If the key is missing, the function shall return an empty array.
func QueryArray(data map[string]interface{}, key string) (out []interface{}) {
	out = []interface{}{}

	// check the key and bail if it is empty
	if key == "" {
		return out
	}
	key += QUERY_CONNECTOR

	// process the query
	for k, val := range data {
		if !strings.HasPrefix(k, key) {
			continue
		}

		out = append(out, val)
	}

	return out
}

// QueryMap restore the map data structure from the decoded data.
//
// It takes a decoded data and the query `key` to the map. Example: for
// `development.database.Username`, `development.database.Password`, ..., the
// key is `development.database`.
//
// If the key is missing, the function shall return an empty map.
//
// QueryMap retains the tailing string query. If we based on the example above
// and the given key is `development`, the generated list can be:
// `database.Username`, `database.Password`, ...
func QueryMap(data map[string]interface{},
	key string) (out map[string]interface{}) {
	out = map[string]interface{}{}

	// check the key and bail if it is empty
	if key == "" {
		return out
	}
	key += QUERY_CONNECTOR

	// process the query
	for k, val := range data {
		if !strings.HasPrefix(k, key) {
			continue
		}

		k = strings.TrimPrefix(k, key)

		out[k] = val
	}

	return out
}
