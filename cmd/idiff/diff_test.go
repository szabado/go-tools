package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	a "github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	assert := a.New(t)

	reader := io.MultiReader(
		delayedReader{50 * time.Millisecond, strings.NewReader("a\nb\nc\n")},
		delayedReader{50 * time.Millisecond, strings.NewReader("a\nb\nc\nd\n")},
		delayedReader{50 * time.Millisecond, strings.NewReader("")},
	)

	var output bytes.Buffer
	err := executeDiff(reader, &output, 40*time.Millisecond)
	assert.NoError(err)

	assert.Equal("3a4\n> d\n", output.String())
}
