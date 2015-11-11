package testutils

import "errors"

// ErrorReaderDefaultError is the default error used for ErrorReader.
var ErrorReaderDefaultError = errors.New("Expected error when reading from ErrorReader")

// ErrorReader is an io.Reader that will always return an error when read from.
type ErrorReader struct {
	Error error
}

// NewErrorReader returns a reader that will return the ErrorReaderDefaultError.
func NewErrorReader() *ErrorReader {
	return &ErrorReader{
		Error: ErrorReaderDefaultError,
	}
}

// Read will return the predefined error.
//
// The argument will be untouched and zero will be the amount of bytes read.
func (r *ErrorReader) Read([]byte) (int, error) {
	return 0, r.Error
}
