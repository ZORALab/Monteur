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

package conductor

import (
	"context"
	"fmt"
)

const (
	jobNotDone      uint = 0
	jobDone         uint = 1
	jobAllCompleted uint = 2
)

// Conductor is the coordinators for executing multiple Jobs in parallel.
//
// This is similar to conducting an orchestra in a theater where the main
// conductor coordinates various musicians to make a good presentation.
//
// Conductor is safe to be created using the standard `&struct{}` method.
type Conductor struct {
	ctx     context.Context
	channel chan Message

	stop func()

	// Log is the control object for logging purposes
	Log Logger

	// Runners are the list of Jobs to be executed in parallel
	Runners map[string]Job

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

	me.channel = make(chan Message, chLength*2)
	me.ctx, me.stop = context.WithCancel(context.Background())

	me.hasInitialized = true
}

// Run is to start the parallel executions
//
// If the Conductor.Runners is empty, this function will return an error as
// there is nothing for execution to begin with.
func (me *Conductor) Run() (err error) {
	if len(me.Runners) == 0 {
		return fmt.Errorf(ERROR_JOBLESS)
	}

	me.init()

	for _, program := range me.Runners {
		me.logInfo("Starting Job '%s' in background...", program.Name())
		go program.Run(me.ctx, me.channel)
		me.logSuccess("➤ OK\n")
	}

	return nil
}

// Coordinate is to manage the coordinatations between the jobs
//
// If the Conductor.Runners is empty, this function will return an error as
// there is nothing for execution to begin with.
//
// Otherwise, should any of the job returns an error, Conductor will stop the
// orchestra entirely and report the error.
func (me *Conductor) Coordinate() (err error) {
	var msg Message
	var ok bool

	if len(me.Runners) == 0 {
		me.logWarning(ERROR_JOBLESS)
		return nil
	}

	me.logInfo("Coordinating running jobs orchestra...")
	for {
		select {
		case <-me.ctx.Done():
			return nil
		case msg, ok = <-me.channel:
			if !ok {
				me.logWarning(ERROR_CHANNEL_CLOSED)
				return nil
			}

			switch me.checkDone(msg) {
			case jobDone:
				continue
			case jobAllCompleted:
				me.logSuccess("➤ DONE")
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

func (me *Conductor) checkStatus(msg Message) {
	var name string
	var ok bool
	var rmsg, rname interface{}
	var status string

	// obtain message and owner from Message
	rmsg, ok = msg.Get(CHMSG_STATUS)
	if !ok {
		return
	}

	rname, ok = msg.Get(CHMSG_OWNER)
	if ok {
		name, ok = rname.(string)
		if !ok {
			name = ""
		}
	}

	// handle coordinated display printout
	status, ok = rmsg.(string)
	if !ok {
		return
	}

	// log the status
	if name != "" {
		me.logError("Status from Job '%s' ➤ %s", name, status)
	} else {
		me.logError("Status ➤ %s", status)
	}
}

func (me *Conductor) checkError(msg Message) (err error) {
	var name string
	var ok bool
	var rmsg, rname interface{}

	// obtain message and owner from Message
	rmsg, ok = msg.Get(CHMSG_ERROR)
	if !ok {
		return nil
	}

	rname, ok = msg.Get(CHMSG_OWNER)
	if ok {
		name, ok = rname.(string)
		if !ok {
			name = ""
		}
	}

	// obtain error object
	err, ok = rmsg.(error)
	if !ok {
		return nil
	}

	// log the output before returning error
	if name != "" {
		me.logError("Error from Job '%s' ➤ %s", name, err)
	} else {
		me.logError("Error ➤ %s", err)
	}

	// return error for conductor to stop the orchestra
	return err
}

func (me *Conductor) checkDone(msg Message) (state uint) {
	var name string
	var ok, isDone bool
	var rname, rmsg interface{}

	state = jobNotDone

	rmsg, ok = msg.Get(CHMSG_DONE)
	if !ok {
		return state
	}

	rname, ok = msg.Get(CHMSG_OWNER)
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
	me.logInfo("Job '%s' ➤ COMPLETED", name)

	if len(me.Runners) == 0 {
		state = jobAllCompleted
	}

	return state
}

func (me *Conductor) logError(format string, a ...interface{}) {
	if !loggerAvailable(me.Log) {
		return
	}

	me.Log.Error(format, a...)
}

func (me *Conductor) logWarning(format string, a ...interface{}) {
	if !loggerAvailable(me.Log) {
		return
	}

	me.Log.Warning(format, a...)
}

func (me *Conductor) logSuccess(format string, a ...interface{}) {
	if !loggerAvailable(me.Log) {
		return
	}

	me.Log.Success(format, a...)
}

func (me *Conductor) logInfo(format string, a ...interface{}) {
	if !loggerAvailable(me.Log) {
		return
	}

	me.Log.Info(format, a...)
}
