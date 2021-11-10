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

package httpclient

type indicator struct {
	handleProgress func(downloaded int64, total int64)
	total          int64
	downloaded     int64
}

func (p *indicator) Write(data []byte) (n int, err error) {
	n = len(data)
	p.downloaded += int64(n)
	p.handleProgress(p.downloaded, p.total)
	return n, nil
}
