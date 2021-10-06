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

package chmsg

import (
	"sync"
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

// New is to create the initialized Message interface object.
func New() Message {
	return &message{
		payload: map[string]interface{}{},
		lock:    &sync.Mutex{},
	}
}
