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

package libmonteur

// Language is the data structure for defining a language metadata.
//
// The fields are compliant to Schema.org definitions available at:
// https://schema.org/Language
//
// Not all fields are made available due to uncommon usage. Request only when
// it is necessary.
type Language struct {
	// Name houses the name of the language in its native tongue.
	//
	// The reason to stay in its native tongue is purely common sense where
	// people in its native language would want to quickly identifies the
	// language without depending on another language translations.
	Name string

	// Code is ISO-639-1 (optionally +ISO 31661-1 Alpha 2) code.
	//
	// These codes are something like `en`, `en-us`, `en-gb`, `en-ca`, and
	// so on.
	Code string
}
