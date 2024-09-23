package main

import (
	"io"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func executeDiff(input io.Reader, output io.Writer) error {
	logrus.Info("Getting input")
	file1, file2 := readFiles(input)
	logrus.Info("file1: ", file1)
	logrus.Info("file2: ", file2)
	logrus.Info("Input fetched")
	file1Name, err := writeToTempFile(file1)
	if err != nil {
		logrus.Warn("error writing file 1: ", err)
		return err
	}
	file2Name, err := writeToTempFile(file2)
	if err != nil {
		logrus.Warn("error writing file 2: ", err)
		return err
	}

	cmd := exec.Command("diff", file1Name, file2Name)
	logrus.Info("cmd created")

	outputBytes, err := cmd.Output()
	if err != nil {
		logrus.Warn("error running command: ", err)
		return err
	}
	output.Write(outputBytes)

	return err
}
