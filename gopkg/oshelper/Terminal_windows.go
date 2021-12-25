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

//go:build windows
// +build windows

package oshelper

import (
	"syscall"
	"unsafe"
)

func _termSize() (row uint, column uint) {
	kernel := syscall.NewLazyDLL("kernel32.dll")
	process := kernel.NewProc("GetConsoleScreenBufferInfo")

	// get syscall handler
	handle, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return row, column
	}

	// construct syscall return data
	info := &struct {
		Size struct {
			X int16
			Y int16
		}
		CursorPosition struct {
			X int16
			Y int16
		}
		Attributes uint16
		Windows    struct {
			Left   int16
			Top    int16
			Right  int16
			Bottom int16
		}
		MaximumWindowSize struct {
			X int16
			Y int16
		}
	}{}

	// perform syscall
	uERR, _, _ := process.Call(uintptr(handle),
		uintptr(unsafe.Pointer(info)),
		0,
	)
	if uERR == 0 {
		return row, column // error occurred
	}

	// calculate row and column
	row = uint(info.Windows.Right - info.Windows.Left + 1)
	column = uint(info.Windows.Bottom - info.Windows.Top + 1)
	return row, column
}
