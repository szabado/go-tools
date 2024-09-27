package main

import (
	"bufio"
	"io"
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

func TestReadLineWithTimeoutTimedOut(t *testing.T) {
	assert := a.New(t)

	reader := delayedReader{100 * time.Millisecond, strings.NewReader("a\nb")}
	scanner := bufio.NewScanner(reader)
	firstLine, elapsed := readLineWithTimeout(scanner, 50*time.Millisecond)
	assert.Equal("", firstLine)
	assert.LessOrEqual(50*time.Millisecond, elapsed)
	assert.Less(elapsed, 75*time.Millisecond)
}

func TestReadLineWithTimeoutNotTimedOut(t *testing.T) {
	assert := a.New(t)

	reader := delayedReader{50 * time.Millisecond, strings.NewReader("a\nb")}
	scanner := bufio.NewScanner(reader)
	firstLine, elapsed := readLineWithTimeout(scanner, 100*time.Millisecond)
	assert.Equal("a", firstLine)
	assert.LessOrEqual(50*time.Millisecond, elapsed)
	assert.Less(elapsed, 75*time.Millisecond)
}

func TestReadPastedInputs(t *testing.T) {
	assert := a.New(t)

	reader := io.MultiReader(
		// The actual delay at the start isn't measured
		delayedReader{100 * time.Millisecond, strings.NewReader("a\nb\nc\n")},
		delayedReader{50 * time.Millisecond, strings.NewReader("d\ne\nf\n")},
		delayedReader{50 * time.Millisecond, strings.NewReader("")},
	)
	firstInput, secondInput := readPastedInput(reader, 40*time.Millisecond)
	assert.Equal("a\nb\nc\n", firstInput)
	assert.Equal("d\ne\nf\n", secondInput)
}

type delayedReader struct {
	delay            time.Duration
	underlyingReader io.Reader
}

var _ io.Reader = (*delayedReader)(nil)

func (re delayedReader) Read(p []byte) (n int, err error) {
	time.Sleep(re.delay)
	return re.underlyingReader.Read(p)
}
