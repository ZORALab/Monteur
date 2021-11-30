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

type ActionID string

const (
	ACTION_PLACEHOLDER            ActionID = "placeholder"
	ACTION_COMMAND                ActionID = "command"
	ACTION_COMMAND_QUIET          ActionID = "command-quiet"
	ACTION_COPY                   ActionID = "copy"
	ACTION_COPY_RECURSIVE         ActionID = "copy-recursive"
	ACTION_COPY_RECURSIVE_QUIET   ActionID = "copy-recursive-quiet"
	ACTION_COPY_QUIET             ActionID = "copy-quiet"
	ACTION_CREATE_DIR             ActionID = "create-dir"
	ACTION_CREATE_PATH            ActionID = "create-path"
	ACTION_DELETE                 ActionID = "delete"
	ACTION_DELETE_RECURSIVE       ActionID = "delete-recursive"
	ACTION_DELETE_RECURSIVE_QUIET ActionID = "delete-recursive-quiet"
	ACTION_DELETE_QUIET           ActionID = "delete-quiet"
	ACTION_IS_EXISTS              ActionID = "is-exists"
	ACTION_MOVE                   ActionID = "move"
	ACTION_MOVE_QUIET             ActionID = "move-quiet"
)
