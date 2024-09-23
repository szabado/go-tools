package main

import (
	"bufio"
	"io"
	"strings"
	"time"
)

func readFiles(inputReader io.Reader) (string, string) {
	scanner := bufio.NewScanner(inputReader)
	results := greedyReader(scanner)
	return strings.Join(*results[0], "\n"), strings.Join(*results[1], "\n")
	// fmt.Printf("%+v\n", *results[0])
	// fmt.Printf("%+v\n", *results[1])
}

func greedyReader(inputScanner *bufio.Scanner) []*[]string {
	println("reading first line")
	firstLine, _, _ := readLine(inputScanner)
	println("read first line")
	doc := []string{firstLine}
	docs := []*[]string{&doc}
	timeout := 50 * time.Millisecond
loop:
	for {
		if len(docs) == 1 {
			println("if")
			line, duration, _ := readLine(inputScanner)
			println("read line")
			if duration >= timeout {
				doc = []string{}
				docs = append(docs, &doc)
			}
			doc = append(doc, line)
		} else {
			println("else")
			line, duration := readLineWithTimeout(inputScanner, timeout)
			doc = append(doc, line)
			if duration >= timeout {
				break loop
			}
		}
	}

	return docs
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

// func readFile(input io.Reader) {
// 	inputChan := make(chan string, 1)
// 	inputScanner := bufio.NewScanner(input)
// 	go readFileWithoutTimeout(inputChan)

// 	i := 0
// 	var allFiles [][]string
// 	var file []string
// 	scanner.ScanStrings()
// 	allFiles = append(allFiles, file)
// inputloop:
// 	for {
// 		select {
// 		case line := <-inputChan:
// 			input = append(input, line)
// 		case <-time.After(200 * time.Millisecond):
// 			var newInput []string
// 			if i == 1 {
// 				break inputloop
// 			} else {
// 				i++
// 			}
// 			input = newInput
// 			allInputs = append(allInputs, input)
// 		}
// 	}

// 	log.Info(allInputs)
// }

// func readFileWithoutTimeout(input chan string) error {
// 	for {
// 		in := bufio.NewReader(os.Stdin)
// 		result, err := in.ReadString('\n')
// 		if err != nil {
// 			return err
// 		}

// 		input <- result
// 	}
// }
