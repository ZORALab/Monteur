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

// Package monteur is the Go package interface to run Monteur functions.
//
// These functions are the package services offered by Monteur project where it
// is friendly to Go import. The objective is to ensure the availability where
// any interested Go developer can easily integrate/spin Monteur into his/her
// specific CI needs.
package monteur

import (
	"fmt"
)

// Purge is the function to remove all setup and data from the repository.
//
// The result is a clean slate git repository with all the programs, caches,
// build artifacts, test results, etc removed.
func Purge() int {
	c := &cleaner{CleanAll: true}
	return c.Run()
}

// Setup is the function to download all dependencies as per configurations.
//
// The action shall download all the dependencies as stated by all the
// configuration files inside a repository's ./.configs/monteur/setup/
// directory.
func Setup() (statusCode int) {
	s := &setup{}
	return s.Run()
}

// Develop is the function to configure the terminal matching local tools.
//
// The action is to ensure anyone who tries to develop the repository has the
// same setup without needing to manually setup the workspace. Another objective
// is to ensure everyone speaks the common language when it comes to
// development.
func Develop() int {
	fmt.Println("Placeholder: Develop function called")
	return STATUS_ERROR
}

// Test is the function to execute the autonomous test job for the repository.
//
// This action is to ensure the test sequences are called uniformly during
// development or a continuous improvement autonomous run. That way, anyone
// including the CI infrastructure can run testing for the repository both
// manually and autonomously at any given time.
func Test() int {
	fmt.Println("Placeholder: Test function called")
	return STATUS_ERROR
}

// Clean is the function to clear up the repository for the next run.
//
// This action is to clean up the repository from a previous run, allowing a
// fresh run on the next round. Unlike Purge() function, does not remove all the
// downloaded dependencies done by Setup() function.
func Clean() int {
	c := &cleaner{CleanAll: false}
	return c.Run()
}

// Release is the function to update repository for releasing a next version.
//
// This action is to update all necessary documents like changelog, version
// numbers, build configurations as programmed for the next release. This
// function should be done before building the next version release.
func Release() int {
	fmt.Println("Placeholder: Release function called")
	return STATUS_OK
}

// Build is the function to build the software with current configurations.
//
// This action is to build the release version software into many of its
// variants such as but not limited to operating system, CPU types, packaging
// types (e.g. plugins).
func Build() int {
	fmt.Println("Placeholder: Build function called")
	return STATUS_ERROR
}

// Package is the function to package the built software into distributions.
//
// This action packages the built software into many distributions channel
// formats like .msi for Microsoft Windows OS, .deb for Debian-based Linux OS,
// .rpm for RPM-based Linux OS, .dmg for MacOS, .appImage for AppImage.
func Package() int {
	fmt.Println("Placeholder: Package function called")
	return STATUS_ERROR
}

// Publish is the function to update and publish the documentations.
//
// this action generates the documentations artifact and publish it to its
// reading channels such as web, file server for PDF files, and etc.
func Publish() int {
	fmt.Println("Placeholder: Publish function called")
	return STATUS_ERROR
}
