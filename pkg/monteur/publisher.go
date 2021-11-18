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

package monteur

import (
	"fmt"
)

type publisher struct {
	OnlyBuild bool
}

func (unit *publisher) Run() (statusCode int) {
	fmt.Println("Placeholder: build publications called")

	if unit.OnlyBuild {
		return STATUS_OK
	}

	fmt.Println("Placeholder: publish function called")
	return STATUS_OK
}
