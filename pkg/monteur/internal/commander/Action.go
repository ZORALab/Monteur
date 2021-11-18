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

package commander

type ActionID uint

const (
	ACTION_PLACEHOLDER ActionID = iota
	ACTION_COMMAND
	ACTION_COMMAND_QUIET
	ACTION_COPY
	ACTION_COPY_RECURSIVE
	ACTION_COPY_RECURSIVE_QUIET
	ACTION_COPY_QUIET
	ACTION_CREATE_DIR
	ACTION_CREATE_PATH
	ACTION_DELETE
	ACTION_DELETE_RECURSIVE
	ACTION_DELETE_RECURSIVE_QUIET
	ACTION_DELETE_QUIET
	ACTION_IS_EXISTS
)

type Action struct {
	// Name is for the action naming used in logging and identification
	Name string

	// Location is where the directory shall change to for execution
	Location string

	// Source is the input of the action in general.
	//
	// See 'Type' documentations for the action's specification.
	Source string

	// Target is the output of the action in general.
	//
	// See 'Type' documentations for the action's specification.
	Target string

	// Save is to save the output into variables.
	//
	// The value shall be the name (or 'key') of the variable.
	Save string

	// Type is the action type ID.
	Type ActionID
}
