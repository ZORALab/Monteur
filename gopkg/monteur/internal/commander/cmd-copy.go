// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
// Copyright 2020 Tobias Klauser (tklauser@distanz.ch)
// Copyright 2019 Kir Kolyshkin (kolyshkin@gmail.com)
// Copyright 2019 Dominic Yin (hi@ydcool.me)
// Copyright 2019 Tõnis Tiigi (tonistiigi@gmail.com)
// Copyright 2018 Maxim Ivanov
// Copyright 2017 Sargun Dhillon (sargun@sargun.me)
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

package commander

import (
	"fmt"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/oshelper"
)

func cmdCopy(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return nil, fmt.Errorf("source is empty")
	}

	if action.Target == "" {
		return nil, fmt.Errorf("target is empty")
	}

	err = oshelper.Copy(action.Source, action.Target)
	if err != nil {
		err = fmt.Errorf("copy failed with error: %s", err)
	}

	return nil, err
}

func cmdCopyQuiet(action *Action) (out interface{}, err error) {
	out, _ = cmdCopy(action)
	return out, nil
}
