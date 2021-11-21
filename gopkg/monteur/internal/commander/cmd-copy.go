// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
// Copyright 2020 Tobias Klauser (tklauser@distanz.ch)
// Copyright 2019 Kir Kolyshkin (kolyshkin@gmail.com)
// Copyright 2019 Dominic Yin (hi@ydcool.me)
// Copyright 2019 Tõnis Tiigi (tonistiigi@gmail.com)
// Copyright 2018 Maxim Ivanov
// Copyright 2017 Sargun Dhillon (sargun@sargun.me)
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
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/sys/unix"
)

func cmdCopy(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return nil, fmt.Errorf("source is empty")
	}

	if action.Target == "" {
		return nil, fmt.Errorf("target is empty")
	}

	// get src file mode and execute copy accordingly
	fi, err := os.Lstat(action.Source)
	if err != nil {
		return nil, fmt.Errorf("%s: %s",
			"problem opening file",
			err,
		)
	}

	mode := fi.Mode()
	switch {
	case mode.IsRegular():
		err = _copyRegular(action.Source, action.Target, fi)
	case mode.IsDir():
		err = fmt.Errorf("%s: %s",
			"copy failed as .Source is a directory",
			action.Source,
		)
	case mode&fs.ModeSymlink != 0:
		err = _copySymlink(action.Source, action.Target, fi)
	case mode&fs.ModeNamedPipe != 0:
		err = _copyPipe(action.Target, fi)
	default:
		err = fmt.Errorf("unidentifiable .Source")
	}

	return nil, err
}

func cmdCopyQuiet(action *Action) (out interface{}, err error) {
	out, _ = cmdCopy(action)
	return out, nil
}

func cmdCopyRecursive(action *Action) (out interface{}, err error) {
	if action.Source == "" {
		return nil, fmt.Errorf("source is empty")
	}

	if action.Target == "" {
		return nil, fmt.Errorf("target is empty")
	}

	return nil, _copyDir(action.Source, action.Target)
}

func cmdCopyRecursiveQuiet(action *Action) (out interface{}, err error) {
	out, _ = cmdCopyRecursive(action)
	return out, nil
}

func _copyDir(src string, dest string) (err error) {
	err = filepath.Walk(src, func(path string,
		info os.FileInfo, err error) error {
		// error obtaining file information
		if err != nil {
			return fmt.Errorf("error with source file: %s", err)
		}

		// filepath.Walk does return its own path. Ignore it as we only
		// want the contents
		if path == src {
			return nil
		}

		// generate relative path
		destPath, _ := filepath.Rel(src, path)
		destPath = filepath.Join(dest, destPath)

		// take action according to the target nature
		mode := info.Mode()
		switch {
		case mode.IsRegular():
			err = _copyRegular(path, destPath, info)
		case mode.IsDir():
			err = os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return fmt.Errorf("%s: %s",
					"failed to create directory",
					err,
				)
			}

			err = _copyDir(path, destPath)
		case mode&os.ModeSymlink != 0:
			err = _copySymlink(path, destPath, info)
		case mode&os.ModeNamedPipe != 0:
			fallthrough
		case mode&os.ModeSocket != 0:
			err = _copyPipe(destPath, info)
		default:
			return fmt.Errorf("unidentifiable Source")
		}

		return err
	})

	return err //nolint:wrapcheck
}

func _copyPipe(dest string, info os.FileInfo) (err error) {
	// get stat from Source info
	stat := info.Sys().(*syscall.Stat_t)

	// create the pipe file
	err = unix.Mkfifo(dest, stat.Mode)
	if err != nil {
		return fmt.Errorf("%s: %s",
			"failed to create pipe file",
			err,
		)
	}

	// restore owner
	err = os.Chmod(dest, info.Mode())
	if err != nil {
		return fmt.Errorf("%s: %s",
			"failed to set file permission",
			err,
		)
	}

	// restore timestamp
	err = __copyFileTimestamp(dest, info, stat)
	if err != nil {
		return err
	}

	return nil
}

func _copySymlink(src string, dest string, info os.FileInfo) (err error) {
	// read the symlink
	link, err := os.Readlink(src)
	if err != nil {
		return fmt.Errorf("%s: %s",
			"failed to read symlink",
			err,
		)
	}

	// create the symlink
	err = os.Symlink(dest, link)
	if err != nil {
		return fmt.Errorf("%s: %s",
			"error creating symlink",
			err,
		)
	}

	// restore timestamp
	err = __copyFileUTimestampNano(dest, info, nil)
	if err != nil {
		return err
	}

	return nil
}

func _copyRegular(src string, dest string, info os.FileInfo) (err error) {
	// open source file
	fileSrc, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("%s: %s",
			"error opening source file",
			err,
		)
	}

	// open destination file
	fileDest, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fileSrc.Close()
		return fmt.Errorf("%s: %s",
			"error opening destination file",
			err,
		)
	}

	// perform body copy
	_, err = io.Copy(fileDest, fileSrc)
	if err != nil {
		fileSrc.Close()
		fileDest.Close()
		return fmt.Errorf("%s: %s",
			"error copying file",
			err,
		)
	}

	// ensure good read
	err = fileDest.Sync()
	if err != nil {
		fileSrc.Close()
		fileDest.Close()
		return fmt.Errorf("%s: %s",
			"error sync file",
			err,
		)
	}

	// close all the files
	fileSrc.Close()
	fileDest.Close()

	// restore file permission
	err = os.Chmod(dest, info.Mode())
	if err != nil {
		return fmt.Errorf("%s: %s",
			"failed to set file permission",
			err,
		)
	}

	// restore file ownership
	stat := info.Sys().(*syscall.Stat_t)
	uid := int(stat.Uid)
	gid := int(stat.Gid)
	err = os.Lchown(dest, uid, gid)
	if err != nil {
		return fmt.Errorf("%s: %s",
			"error changing file permission",
			err,
		)
	}

	// restore timestamp
	err = __copyFileTimestamp(dest, info, stat)
	if err != nil {
		return err
	}

	return nil
}
