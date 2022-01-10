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

package libcmd

import (
	"context"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/conductor"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libsecrets"
)

type Task interface {
	Run(context.Context, chan conductor.Message)
	Name() string
}

type Manager struct {
	task Task

	Metadata  *libmonteur.TOMLMetadata
	Variables map[string]interface{}
	Job       string
}

func (me *Manager) Parse(path string, secret *libsecrets.Secrets) (err error) {
	system, ok := me.Variables[libmonteur.VAR_COMPUTE].(string)
	if !ok {
		panic("MONTEUR DEV: please assign VAR_COMPUTE before Parse()!")
	}

	// parse data file
	switch me.Job {
	case libmonteur.JOB_TEST:
		err = me.runTest(system, path, secret)
	case libmonteur.JOB_PREPARE:
		err = me.runPrepare(system, path, secret)
	case libmonteur.JOB_BUILD:
		err = me.runBuild(system, path, secret)
	case libmonteur.JOB_PACKAGE:
		err = me.runPackage(system, path, secret)
	case libmonteur.JOB_RELEASE:
		err = me.runRelease(system, path, secret)
	case libmonteur.JOB_COMPOSE:
		err = me.runCompose(system, path, secret)
	case libmonteur.JOB_PUBLISH:
		err = me.runPublish(system, path, secret)
	case libmonteur.JOB_CLEAN:
		err = me.runClean(system, path, secret)
	case libmonteur.JOB_SETUP:
		err = me.runSetup(system, path, secret)
	default:
		panic("MONTEUR DEV: What kind of job is this? ➤ " + me.Job)
	}

	return err
}

func (me *Manager) runTest(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	subject := &basicCMD{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject
	return nil
}

func (me *Manager) runPrepare(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	subject := &preparer{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject
	return nil
}

func (me *Manager) runBuild(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	subject := &basicCMD{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject
	return nil
}

func (me *Manager) runPackage(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	subject := &packager{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject
	return nil
}

func (me *Manager) runRelease(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	subject := &releaser{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject
	return nil
}

func (me *Manager) runCompose(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	subject := &basicCMD{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject
	return nil
}

func (me *Manager) runPublish(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	subject := &basicCMD{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject
	return nil
}

func (me *Manager) runClean(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	subject := &basicCMD{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject
	return nil
}

func (me *Manager) runSetup(system string,
	path string, secret *libsecrets.Secrets) (err error) {
	err = initializeMonteurFS(me.Variables)
	if err != nil {
		return err
	}

	subject := &setup{
		thisSystem: system,
		variables:  me.Variables,
	}

	err = subject.Parse(path, secret)
	if err != nil {
		return err
	}

	me.task = subject

	return nil
}

// Name is for generating the program Metadata.Name when used as in interface.
//
// This should only be called after the Manager is initialized successfully.
func (me *Manager) Name() string {
	return me.task.Name()
}

// Run is to execute the publisher's commands sequence.
//
// Everything must be setup properly before calling this function. It was meant
// for Monteur's commands-driven API(s).
//
// All errors generated in this method shall use `me.reportError` instead of
// returning `fmt.Errorf` since it will be executed in parallel with others
// in an asynchonous manner.
//
// This should only be called after the Manager is initialized successfully.
func (me *Manager) Run(ctx context.Context, ch chan conductor.Message) {
	me.task.Run(ctx, ch)
}
