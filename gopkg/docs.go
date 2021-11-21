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

// gopkg is the root directory of all Monteur Go packages.
//
// Unlike the conventional Go pathing and structure, Monteur project aims to
// operate continuous integration not just remotely but also locally. This
// removes the most critical vender locked-in dependency: the CI infrastructure.
//
// Also, Das Monteur should also allow anyone to easily customize and build
// the software when he/she has access to the source codes without getting into
// cracking their head to solve all the dependencies nightmare for both build
// tools and the software dependencies as well.
package gopkg
