package main

import (
	"bufio"
	"io"
	"strings"
	"time"
)

func readPastedInput(inputReader io.Reader) (string, string) {
	inputScanner := bufio.NewScanner(inputReader)
	firstLine, _ := readLine(inputScanner)
	doc := []string{firstLine}
	doc2 := []string{}
	timeout := 200 * time.Millisecond
	timeouts := 0
	for timeouts < 2 {
		if timeouts == 0 {
			line, duration := readLine(inputScanner)
			if duration >= timeout {
				doc = append(doc, "")
				doc2 = append(doc2, line)
				timeouts++
			} else {
				doc = append(doc, line)
			}
		} else {
			line, duration := readLineWithTimeout(inputScanner, timeout)
			if duration >= timeout {
				timeouts++
			}
			doc2 = append(doc2, line)
		}
	}

	return strings.Join(doc, "\n"), strings.Join(doc2, "\n")
}

func readLine(inputReader *bufio.Scanner) (string, time.Duration) {
	start := time.Now()
	if !inputReader.Scan() {
		return "", 0
	}
	text := inputReader.Text()
	elapsed := time.Since(start)

	return text, elapsed
}

func readLineWithTimeout(inputScanner *bufio.Scanner, timeout time.Duration) (string, time.Duration) {
	resultChan := make(chan stringWithDuration, 1)
	go func() {
		value, duration := readLine(inputScanner)
		resultChan <- stringWithDuration{value, duration}
	}()
	select {
	case result := <-resultChan:
		return result.value, result.duration
	case <-time.After(timeout):
		return "", timeout
	}
}

type stringWithDuration struct {
	value    string
	duration time.Duration
}
