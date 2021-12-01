//+build unix !windows

package term

import (
	"syscall"
	"unsafe"
)

// Size is to get a terminal row (height) and column (width) sizes at a given
// instance. If there is an error occurs, both row and column becomes 0.
func (t *Terminal) Size() (row uint, column uint) {
	size := &struct {
		Row    uint16
		Column uint16
	}{}

	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(size)))

	row = uint(size.Row)
	if t.row != 0 {
		row = t.row
	}

	column = uint(size.Column)
	if t.column != 0 {
		column = t.column
	}

	return row, column
}
