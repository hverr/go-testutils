package testutils

import "errors"

// ErrorWriterDefaultError is the default error used for ErrorWriter.
var ErrorWriterDefaultError = errors.New("Expected error when writing to ErrorWriter")

// ErrorWriter is an io.Writer that will always return an error when written to.
type ErrorWriter struct {
	Error error
}

// NewErrorWriter returns a writer that will return ErrorWriterDefaultError.
func NewErrorWriter() *ErrorWriter {
	return &ErrorWriter{
		Error: ErrorWriterDefaultError,
	}
}

// Write will return the predifined error.
//
// Zero will be the amount of bytes written.
func (w *ErrorWriter) Write([]byte) (int, error) {
	return 0, w.Error
}
