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

import (
	"fmt"
	"os"
	"path/filepath"
)

func cmdDelete(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return nil, fmt.Errorf("source is empty")
	}

	err = _delete(action.Source)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func cmdDeleteQuiet(action *Action) (out interface{}, err error) {
	_, _ = cmdCopy(action)
	return nil, nil
}

func cmdDeleteRecursive(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return nil, fmt.Errorf("source is empty")
	}

	return nil, _deleteDir(action.Source)
}

func cmdDeleteRecursiveQuiet(action *Action) (out interface{}, err error) {
	out, _ = cmdDeleteRecursive(action)
	return out, nil
}

func _deleteDir(src string) (err error) {
	err = filepath.Walk(src, func(path string,
		info os.FileInfo, err error) error {
		// error obtaining file information
		if err != nil {
			return fmt.Errorf("%s: %s",
				"error with source file",
				err,
			)
		}

		// take action according to the target nature
		mode := info.Mode()
		switch {
		case mode.IsDir():
			err = _deleteDir(path)
		default:
			err = _delete(path)
		}

		return err
	})

	if err != nil {
		return err //nolint:wrapcheck
	}

	return _delete(src)
}

func _delete(src string) (err error) {
	err = os.Remove(src)
	if err != nil {
		err = fmt.Errorf("%s: %s", "failed to delete target", err)
	}

	return err
}
