package testutils

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorWriter(t *testing.T) {
	var writer io.Writer

	w := NewErrorWriter()
	writer = w
	n, err := writer.Write(nil)
	assert.EqualValues(t, 0, n)
	assert.Equal(t, ErrorWriterDefaultError, err)
}

func TestErrorWriterWithCustomError(t *testing.T) {
	testErr := errors.New("Custom test error.")
	w := NewErrorWriter()
	w.Error = testErr
	n, err := w.Write(nil)
	assert.EqualValues(t, 0, n)
	assert.Equal(t, testErr, err)
}
