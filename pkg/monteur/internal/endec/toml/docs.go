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

// Package toml is the encoding package for parsing or generating TOML format.
//
// This package warps an embedded third-party TOML endec to ensure consistency
// across package usage. The design was to expand the encoding packages from
// Go standard package to reach TOML format.
//
// Depending on the embedded third-party TOML endec, this package may use
// Go-only interface marshall and unmarshall convention to keep its codes clean
// and simple.
package toml
