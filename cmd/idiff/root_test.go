package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	a "github.com/stretchr/testify/assert"
)

func TestExecuteRoot(t *testing.T) {
	assert := a.New(t)
	exitCode := -1
	osExit = func(value int) {
		exitCode = value
	}

	input = io.MultiReader(
		delayedReader{50 * time.Millisecond, strings.NewReader("a\nb\nc\n")},
		delayedReader{50 * time.Millisecond, strings.NewReader("a\nb\nc\nd\n")},
		delayedReader{50 * time.Millisecond, strings.NewReader("")},
	)

	var outputBuffer bytes.Buffer
	output = &outputBuffer
	timeout = 40 * time.Millisecond

	execute()

	assert.Equal("3a4\n> d\n", outputBuffer.String())
	assert.Equal(exitCode, 0)
}
