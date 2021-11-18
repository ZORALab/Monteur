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

package libpublish

import (
	"fmt"
)

type Publisher struct {
	Name string
}

func (fx *Publisher) Build() (err error) {
	fmt.Printf("Placeholder BUILD called\n")
	return nil
}

func (fx *Publisher) Publish() (err error) {
	fmt.Printf("Placeholder PUBLISH called\n")
	return nil
}
