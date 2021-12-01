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
	"fmt"
	"sync"
)

const (
	CHMSG_DONE   = "done"
	CHMSG_ERROR  = "error"
	CHMSG_OWNER  = "owner"
	CHMSG_STATUS = "status"
	CHMSG_OUTPUT = "output"
)

// Message is the interface for message payload used in Go channel tramissions.
type Message interface {
	Add(k string, v interface{})
	Get(k string) (v interface{}, ok bool)
	Update(k string, v interface{})
	Delete(k string)
}

type message struct {
	payload map[string]interface{}
	lock    *sync.Mutex
}

// Add is to add a key-value into Message interfacing object.
func (ch *message) Add(k string, v interface{}) {
	ch.lock.Lock()
	defer ch.lock.Unlock()

	ch.payload[k] = v
}

// Get is to get a value from a given key.
//
// It returns the value and the status of query. The status shall be checked
// before using the value for failure handling.
func (ch *message) Get(k string) (v interface{}, ok bool) {
	ch.lock.Lock()
	defer ch.lock.Unlock()

	v, ok = ch.payload[k]
	return v, ok
}

// Update is to update or add a key-value into Message interfacing object.
//
// This is a mirror to Add() function.
func (ch *message) Update(k string, v interface{}) {
	ch.lock.Lock()
	defer ch.lock.Unlock()

	ch.Add(k, v)
}

// Delete is to delete a key-value from the Message interfacing object.
//
// The function deletes the key-value regardless of its existence.
func (ch *message) Delete(k string) {
	ch.lock.Lock()
	defer ch.lock.Unlock()

	delete(ch.payload, k)
}

// NewMessage is to create the initialized Message interface object.
//
// You need to construct your own message data that can or cannot be understood
// by Conductor. This is useful when you use Message outside of Conductor
// (e.g. you develop your own synchonization mechanism).
func NewMessage() Message {
	return &message{
		payload: map[string]interface{}{},
		lock:    &sync.Mutex{},
	}
}

// CreateStatus creates a standard Message object for Conductor.
//
// It takes 2 inputs: the Job owner name and the fmt.Printf like message.
func CreateStatus(owner string, format string, a ...interface{}) Message {
	m := NewMessage()

	m.Add(CHMSG_OWNER, owner)
	m.Add(CHMSG_STATUS, fmt.Sprintf(format, a...))

	return m
}

// CreateOutput creates a standard Message object for Conductor.
//
// It takes 2 inputs: the Job owner name and the fmt.Printf like message.
func CreateOutput(owner string, format string, a ...interface{}) Message {
	m := NewMessage()

	m.Add(CHMSG_OWNER, owner)
	m.Add(CHMSG_OUTPUT, fmt.Sprintf(format, a...))

	return m
}

// CreateError creates an error Message object for Conductor.
//
// It takes 2 inputs: the Job owner name and the fmt.Errorf like error message.
func CreateError(owner string, format string, a ...interface{}) Message {
	m := NewMessage()

	m.Add(CHMSG_OWNER, owner)
	m.Add(CHMSG_ERROR, fmt.Errorf(format, a...))

	return m
}

// CreateDone creates a done Message object for Conductor.
//
// It takes 1 input: the Job owner name.
func CreateDone(owner string) Message {
	m := NewMessage()

	m.Add(CHMSG_OWNER, owner)
	m.Add(CHMSG_DONE, true)

	return m
}
