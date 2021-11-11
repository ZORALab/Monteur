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

package libsetup

import (
	"context"
	"net/http"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/checksum"
)

type Source struct {
	Archive  string
	URL      string
	Method   string
	Checksum *checksum.Hasher

	Get    func(ctx context.Context)
	Unpack func(ctx context.Context)

	HandleError    func(err error)
	HandleProgress func(progress, total int64)
	HandleSuccess  func()
	HandleRedirect func(req *http.Request, via []*http.Request) error

	Headers []string
}
