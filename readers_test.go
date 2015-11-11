package testutils

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorReader(t *testing.T) {
	var reader io.Reader

	r := NewErrorReader()
	reader = r
	n, err := reader.Read(nil)
	assert.EqualValues(t, 0, n)
	assert.Equal(t, ErrorReaderDefaultError, err)
}

func TestErrorReaderWithCustomError(t *testing.T) {
	testErr := errors.New("Custom test error.")
	r := NewErrorReader()
	r.Error = testErr
	n, err := r.Read(nil)
	assert.EqualValues(t, 0, n)
	assert.Equal(t, testErr, err)
}
