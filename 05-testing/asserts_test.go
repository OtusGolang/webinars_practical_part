package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertReporting(t *testing.T) {

	var err error = io.EOF
	// ...

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.True(t, err == nil)
}
