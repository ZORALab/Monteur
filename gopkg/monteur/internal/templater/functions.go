// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by datalicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templater

import (
	"fmt"

	txtTemplate "text/template"
)

func textTemplate(name string,
	funcMap map[string]interface{}) *txtTemplate.Template {
	t := txtTemplate.New(name)

	funcMap["string"] = stringify
	funcMap["printf"] = printf

	return t.Funcs(funcMap)
}

func stringify(input interface{}) string {
	return fmt.Sprintf("%#s", input)
}

func printf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
