package filesystem

import (
	"os"
)

// IsDirExists is a quick check for directory existence.
//
// It will not only check for directory existence but also verifies that the
// given path is also a directory.
func IsDirExists(path string) bool {
	if i, err := os.Stat(path); !os.IsNotExist(err) && i.IsDir() {
		return true
	}

	return false
}
