package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func writeToTempFile(contents string) (string, error) {
	dir := os.TempDir()
	logrus.Info("temp dir: ", dir)
	file, err := os.CreateTemp(dir, "idiff-*")
	if err != nil {
		return "", err
	}
	defer file.Close()
	file.WriteString(contents)
	fileName := file.Name()
	logrus.Info("temp file: ", fileName)
	return fileName, nil
}
