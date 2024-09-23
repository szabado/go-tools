package main

import (
	"io"
	"os/exec"
)

func executeDiff(input io.Reader) error {
	file1, file2 := readFiles(input)
	file1Name, err := writeToTempFile(file1)
	if err != nil {
		return err
	}
	file2Name, err := writeToTempFile(file2)
	if err != nil {
		return err
	}

	cmd := exec.Command("diff", file1Name, file2Name)
	cmd.Wait()

	return nil
}
