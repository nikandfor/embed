package tests

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	d := EmptyReadAll()
	assert.Empty(t, d)

	r := EmptyReader()
	_, err := r.Read(nil)
	assert.Equal(t, io.EOF, err)
}
