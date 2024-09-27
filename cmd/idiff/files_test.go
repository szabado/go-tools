package main

import (
	"os"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	assert := a.New(t)
	fileName, err := writeToTempFile("a\nb")
	assert.NoError(err)

	fileContents, err := os.ReadFile(fileName)
	assert.NoError(err)
	assert.Equal("a\nb", string(fileContents))
}
