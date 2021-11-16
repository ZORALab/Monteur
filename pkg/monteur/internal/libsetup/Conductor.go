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
	"fmt"
	"os"

	"gitlab.com/zoralab/monteur/pkg/monteur/internal/chmsg"
	"gitlab.com/zoralab/monteur/pkg/monteur/internal/libmonteur"
)

const (
	jobNotDone      uint = 0
	jobDone         uint = 1
	jobAllCompleted uint = 2
)

type Conductor struct {
	ctx     context.Context
	channel chan chmsg.Message

	stop func()

	Runners map[string]*Program

	hasInitialized bool
}

func (me *Conductor) init() {
	if me.hasInitialized {
		return
	}

	chLength := len(me.Runners)
	if chLength == 0 {
		chLength = 2
	}

	me.channel = make(chan chmsg.Message, chLength*2)
	me.ctx, me.stop = context.WithCancel(context.Background())

	me.hasInitialized = true
}

func (me *Conductor) Run() {
	if len(me.Runners) == 0 {
		return
	}

	me.init()

	for _, program := range me.Runners {
		// setup configurations for program
		program.ReportUp = me.channel

		// begin the run asynchonously
		go program.Shop(me.ctx)
	}
}

func (me *Conductor) Coordinate() (err error) {
	var msg chmsg.Message
	var ok bool

	if len(me.Runners) == 0 {
		return nil
	}

	for {
		select {
		case <-me.ctx.Done():
			return nil
		case msg, ok = <-me.channel:
			if !ok {
				return nil // channel was closed
			}

			switch me.checkDone(msg) {
			case jobDone:
				continue
			case jobAllCompleted:
				return nil
			case jobNotDone:
				fallthrough
			default:
			}

			err = me.checkError(msg)
			if err != nil {
				me.stop()
				return err
			}

			me.checkStatus(msg)
		}
	}
}

func (me *Conductor) checkStatus(msg chmsg.Message) {
	var status string

	rmsg, ok := msg.Get(libmonteur.CHMSG_STATUS)
	if !ok {
		return
	}

	// handle coordinated display printout
	status, ok = rmsg.(string)
	if !ok {
		return
	}

	fmt.Fprintf(os.Stdout, "%s\n", status)
}

func (me *Conductor) checkError(msg chmsg.Message) (err error) {
	var ok bool
	var rmsg interface{}

	rmsg, ok = msg.Get(libmonteur.CHMSG_ERROR)
	if !ok {
		return nil
	}

	err, ok = rmsg.(error)
	if !ok {
		return nil
	}

	return err
}

func (me *Conductor) checkDone(msg chmsg.Message) (state uint) {
	var name string
	var ok, isDone bool
	var rname, rmsg interface{}

	state = jobNotDone

	rmsg, ok = msg.Get(libmonteur.CHMSG_DONE)
	if !ok {
		return state
	}

	rname, ok = msg.Get(libmonteur.CHMSG_OWNER)
	if !ok {
		return state
	}

	isDone, ok = rmsg.(bool)
	if !ok {
		return state
	}

	if !isDone {
		return state
	}

	name, ok = rname.(string)
	if !ok {
		return state
	}

	// a job is done
	state = jobDone
	delete(me.Runners, name)

	if len(me.Runners) == 0 {
		state = jobAllCompleted
	}

	return state
}
