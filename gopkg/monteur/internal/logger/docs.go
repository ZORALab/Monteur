// Copyright 2020 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2020 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
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

// Package logger is a logging service for building and managing logs.
//
// Logger was designed to work safely and concurrently in order to avoid any
// race conditions.
//
// ## Key Features
//
// Among the key features we have are:
//
//  1. Capable of multi-Writer outputs.
//  2. Supports level logging (free specified by users).
//  3. Supports pre-processing function before writing.
package logger
