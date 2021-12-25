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

package oshelper

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

type TermType uint

const (
	TERM_NONE TermType = iota
	TERM_BASH
	TERM_DOS
	TERM_SH
)

type Terminal struct {
	Stdout io.Writer
	Stderr io.Writer

	Type TermType
}

func (me *Terminal) Exec(cmd string, timeout uint64) (err error) {
	var executive *exec.Cmd

	executive, err = me.Start(cmd)
	if err != nil {
		goto done
	}

	err = me.Wait(executive, timeout)
done:
	return err
}

func (me *Terminal) Start(cmd string) (command *exec.Cmd, err error) {
	command = me.createCommand(cmd)

	err = command.Start()
	if err != nil {
		err = fmt.Errorf("cmd exec failed: %s", err)
	}

	return command, err
}

func (me *Terminal) Wait(executive *exec.Cmd, timeout uint64) (err error) {
	var ret chan error
	var ok bool

	// validate input
	if executive == nil {
		err = fmt.Errorf("missing executive for Wait()")
		goto done
	}

	// wait without timeout
	if timeout == 0 {
		err = executive.Wait()
		goto done
	}

	// wait with timeout
	ret = make(chan error, 1)

	go func() {
		ret <- executive.Wait()
		close(ret)
	}()

	for {
		select {
		case <-time.After(time.Duration(timeout)):
			err = executive.Process.Kill()
			err = fmt.Errorf("timeout with error: '%s'", err)
			goto done
		case err, ok = <-ret:
			if ok {
				goto done
			}
		}
	}

done:
	return err
}

func (me *Terminal) IsRoot() bool {
	return os.Getuid() == 0
}

func TermSize() (row uint, column uint) {
	return _termSize() // os-specific
}

func (me *Terminal) createCommand(cmd string) (out *exec.Cmd) {
	//nolint:gosec
	switch me.Type {
	case TERM_BASH:
		out = exec.Command("bash", "-c", cmd)
	case TERM_SH:
		out = exec.Command("sh", "-c", cmd)
	case TERM_DOS:
		out = exec.Command("cmd", "/c", cmd)
	default:
		args := strings.Split(cmd, " ")
		out = exec.Command(args[0], args[1:]...)
	}

	out.Stdout = me.Stdout
	out.Stderr = me.Stderr

	return out
}
