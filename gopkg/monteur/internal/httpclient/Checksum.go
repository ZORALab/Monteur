package httpclient

import (
	"io"
)

// Checksum is an interface for httpclient to perform Checksum operations.
//
// Any checksum hasher supplied to Downloder **MUST** have the following
// interface methods for proper operations. Otherwise, you can define or create
// simple structure to wrap around your desired hasher.
type Checksum interface {
	IsHealthy() error
	Compare(file io.Reader) (bool, error)
}

func checksumOK(hasher Checksum) (verdict bool) {
	var err error

	defer func() {
		if err := recover(); err != nil {
			verdict = false
		}
	}()

	err = hasher.IsHealthy()
	if err == nil {
		verdict = true
	}

	return verdict
}
