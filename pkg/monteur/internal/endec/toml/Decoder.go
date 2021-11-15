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

package toml

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

// DecodeFile is to decode a TOML file from a given path to a data structure.
//
// This function is to simplify and to warp a third-party TOML endec for simple
// utilization.
func DecodeFile(path string,
	data interface{}, config interface{}) (err error) {
	var c *Config
	var ok bool

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FAILED_CONFIG, path)
	}

	if c, ok = config.(*Config); ok {
		c = config.(*Config)
	}

	return decode(f, data, c)
}

// DecodeString is to decode an in-memory TOML string data to a data structure.
//
// This function is to simplify and to warp a third-party TOML endec for simple
// utilization.
func DecodeString(input string,
	data interface{}, config interface{}) (err error) {
	var c *Config
	var ok bool

	f := strings.NewReader(input)

	if c, ok = config.(*Config); ok {
		c = config.(*Config)
	}

	return decode(f, data, c)
}

// DecodeBytes is to decode an in-memory TOML bytes data to a data structure.
//
// This funcion is to simplify and to warp a third-party TOML endec for simple
// utilization.
func DecodeBytes(input []byte,
	data interface{}, config interface{}) (err error) {
	var c *Config
	var ok bool

	f := bytes.NewReader(input)

	if c, ok = config.(*Config); ok {
		c = config.(*Config)
	}

	return decode(f, data, c)
}

func decode(input io.Reader, data interface{}, config *Config) (err error) {
	decoder := toml.NewDecoder(input)

	if config != nil {
		if config.Strict {
			decoder.SetStrict(true)
		}
	}

	err = decoder.Decode(data)
	if err != nil {
		return fmt.Errorf("%s: %s", ERROR_FAILED_DECODE, err)
	}

	return nil
}
