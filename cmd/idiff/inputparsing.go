package main

import (
	"bufio"
	"io"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func readFiles(inputReader io.Reader) (string, string) {
	scanner := bufio.NewScanner(inputReader)
	results := greedyReader(scanner)
	return strings.Join(results[0], "\n"), strings.Join(results[1], "\n")
}

func greedyReader(inputScanner *bufio.Scanner) [][]string {
	firstLine, _, _ := readLine(inputScanner)
	logrus.Info("First line read")
	doc := []string{firstLine}
	doc2 := []string{}
	logrus.Infof("doc1: %v", doc)
	timeout := 200 * time.Millisecond
	timeouts := 0
loop:
	for {
		logrus.Infof("start: %v", doc)
		if timeouts == 0 {
			line, duration, _ := readLine(inputScanner)
			logrus.Info("line: ", line)
			if duration >= timeout {
				doc = append(doc, "")
				logrus.Info("Timed out - switching to second doc")
				doc2 = append(doc2, line)
				logrus.Info("Doc2", doc2)
				timeouts++
			} else {
				logrus.Info("Not timed out - first doc")
				doc = append(doc, line)
			}
		} else {
			line, duration := readLineWithTimeout(inputScanner, timeout)
			logrus.Info("line: ", line)
			if duration >= timeout {
				logrus.Info("timed out")
				if timeouts >= 2 {
					logrus.Infof("end: %v", doc)
					break loop
				} else {
					timeouts++
				}
			}
			doc2 = append(doc2, line)
			logrus.Infof("end: %v", doc)
		}
		logrus.Infof("end: %v", &doc)
	}

	return [][]string{doc, doc2}
}

func readLine(inputReader *bufio.Scanner) (string, time.Duration, bool) {
	start := time.Now()
	if !inputReader.Scan() {
		return "", 0, false
	}
	text := inputReader.Text()

	return text, time.Since(start), true
}

func readLineWithTimeout(inputScanner *bufio.Scanner, timeout time.Duration) (string, time.Duration) {
	resultChan := make(chan stringWithDuration, 1)
	go readLineToChannel(resultChan, inputScanner)
	select {
	case result := <-resultChan:
		return result.value, result.duration
	case <-time.After(timeout):
		return "", timeout
	}
}

func readLineToChannel(resultChan chan stringWithDuration, inputScanner *bufio.Scanner) {
	value, duration, _ := readLine(inputScanner)
	resultChan <- stringWithDuration{value, duration}
}

type stringWithDuration struct {
	value    string
	duration time.Duration
}

