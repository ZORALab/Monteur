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
	"strconv"
)

// Delete function is to zero all secret data before sending to GC.
func Delete(data map[string]interface{}) (err error) {
	for key, value := range data {
		zeroClean(key, value, data)
	}

	return nil
}

func zeroClean(key string, value interface{}, data map[string]interface{}) {
	switch ivalue := value.(type) {
	case []interface{}:
		for i, val := range ivalue {
			zeroClean(key+QUERY_CONNECTOR+strconv.Itoa(i),
				val,
				data)
		}
	case map[string]interface{}:
		for k, val := range ivalue {
			zeroClean(key+QUERY_CONNECTOR+k,
				val,
				data)
		}
	case []byte:
		for i := range ivalue {
			ivalue[i] = 0x00
		}
	case int, int8, int16, int32, int64:
		data[key] = 0
	case uint, uint8, uint16, uint32, uint64:
		data[key] = 0
	case float32, float64:
		data[key] = 0
	case bool:
		data[key] = false
	case string:
		panic("avoid using string for secret: " + key)
	default:
		panic("unknown data type: " + key)
	}
}
