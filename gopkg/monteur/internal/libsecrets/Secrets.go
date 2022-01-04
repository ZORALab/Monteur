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

package libsecrets

import (
	"fmt"
	"strings"
	"sync"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/endec/toml"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/secrets"
)

type Secrets struct {
	data  map[string]interface{}
	mutex *sync.RWMutex
}

func (me *Secrets) Parse(pathings []string) (err error) {
	me.data = map[string]interface{}{}
	if me.mutex == nil {
		me.mutex = &sync.RWMutex{}
	}

	me.mutex.Lock()
	defer me.mutex.Unlock()

	s := &secrets.Processor{DecodeFx: toml.SilentDecodeFile}

	me.data, err = s.DecodeMultiPath(me.data, pathings, nil)
	if err != nil {
		err = fmt.Errorf("%s: %s",
			libmonteur.ERROR_SECRET_PARSE,
			err,
		)
	}

	return err
}

func (me *Secrets) Filter(s string) string {
	me.mutex.RLock()
	defer me.mutex.RUnlock()

	for _, v := range me.data {
		s = strings.ReplaceAll(s,
			fmt.Sprintf("%v", v),
			libmonteur.SECRET_REDACTED,
		)
	}

	return s
}

func (me *Secrets) Query(s string) (out interface{}) {
	var ok bool

	me.mutex.RLock()
	defer me.mutex.RUnlock()

	out, ok = me.data[s]
	if !ok {
		out = libmonteur.SECRET_NO_DATA
	}

	return out
}
