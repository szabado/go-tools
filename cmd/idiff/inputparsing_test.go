package main

import (
	"bufio"
	"strings"
	"testing"
	"time"

	a "github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	assert := a.New(t)
	joinResult := join([]string{"a", "b"})

	assert.Equal("a\nb", joinResult)
}

func TestReadLine(t *testing.T) {
	assert := a.New(t)

	reader := strings.NewReader("aaaa\nbbb\ncc\nd")
	scanner := bufio.NewScanner(reader)
	firstLine, elapsed := readLine(scanner)

	assert.Equal("aaaa", firstLine)
	assert.Less(elapsed, 100*time.Millisecond)
}

func TestReadLineEmpty(t *testing.T) {
	assert := a.New(t)
	reader := strings.NewReader("")
	scanner := bufio.NewScanner(reader)
	firstLine, elapsed := readLine(scanner)
	assert.Equal("", firstLine)
	assert.Equal(0*time.Millisecond, elapsed)
}
