//+build windows !unix

package term

import (
	"syscall"
	"unsafe"
)

type (
	short int16
	word  uint16

	small_rect struct {
		Left   short
		Top    short
		Right  short
		Bottom short
	}

	coord struct {
		X short
		Y short
	}

	console_screen_buffer_info struct {
		Size              coord
		CursorPosition    coord
		Attributes        word
		Window            small_rect
		MaximumWindowSize coord
	}
)

// Size is to get a terminal row (height) and column (width) sizes at a given
// instance. If there is an error occurs, both row and column becomes 0.
func (t *Terminal) Size() (row uint, column uint) {
	kernel := syscall.NewLazyDLL("kernel32.dll")
	process := kernel.NewProc("GetConsoleScreenBufferInfo")

	handle, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return row, column
	}

	info := console_screen_buffer_info{}

	uERR, _, _ := process.Call(uintptr(handle),
		uintptr(unsafe.Pointer(&info)),
		0)

	if uERR == 0 {
		// an error has occured
		return row, column
	}

	row = uint(info.Window.Right - info.Window.Left + 1)
	column = uint(info.Window.Bottom - info.Window.Top + 1)
	return row, column
}
