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
	"fmt"
	"strconv"
)

// Processor is the data structure processing secrets data.
//
// It is safe to create using the conventional `&Processor{}` method.
type Processor struct {
	// DecodeFx is a function for decoding the secret data used in `Decode`.
	//
	// This function **SHALL NOT** be `nil`.
	DecodeFx func(in string, out interface{}, cfg interface{}) (err error)
}

// Decode is the processor method to decode a given secret file using DecodeFx.
//
// If the `Processor`'s `DecodeFx` is not assigned (`nil`), the `Decode` method
// shall return ERROR_MISSING_DECODE_FX without executing any decoding
// sequences.
//
// For `input` and `config` parameters, `Decode` shall return the `DecodeFx`
// generated error message. Hence, the method permits the `nil` or `""` as
// value.
//
// Decode allows overwriting elements from previous decoded data structure via
// the optional `data`.
//
// The method shall generates the following output:
//   1. `nil`, `err` - any error is found
//   2. `out`, `nil` - decoded data without any error
//
// The format of the query is using period (`.`) to join multiple depth key.
// This also applies to array values. Example:
//     v, err := unit.Decode(...)
//     if err != nil {
//             // handle error
//     }
//
//     data, ok := v["development.database.0.Password"]
//
// The flattening is for quick and easy grab of information. If there is a need
// to restore an array or a map structure, see `QueryArray` and `QueryMap`
// functions.
func (unit *Processor) Decode(data map[string]interface{}, input string,
	config interface{}) (out map[string]interface{}, err error) {
	var rawD map[string]interface{}

	if unit.DecodeFx == nil {
		return nil, fmt.Errorf(ERROR_MISSING_DECODE_FX)
	}

	// Decode the secrets from the given input
	rawD = map[string]interface{}{}
	err = unit.DecodeFx(input, &rawD, nil)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// post process the data into a string friendly query method
	out = data
	if data == nil {
		out = map[string]interface{}{}
	}

	for k, v := range rawD {
		unit.postProcessing("", k, v, out)
	}

	return out, nil
}

func (unit *Processor) postProcessing(pre string, key string,
	value interface{}, d map[string]interface{}) {
	if pre != "" {
		key = pre + "." + key
	}

	switch value.(type) {
	case []interface{}:
		for i, val := range value.([]interface{}) {
			unit.postProcessing(key, strconv.Itoa(i), val, d)
		}
	case map[string]interface{}:
		for k, val := range value.(map[string]interface{}) {
			unit.postProcessing(key, k, val, d)
		}
	default:
		d[key] = value
	}
}
