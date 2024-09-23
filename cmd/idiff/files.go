package main

import (
	"os"
	"path/filepath"
)

func writeToTempFile(contents string) (string, error) {
	dir := os.TempDir()
	file, err := os.CreateTemp(dir, "idiff-*")
	if err != nil {
		return "", err
	}
	defer file.Close()
	file.WriteString(contents)
	return filepath.Join(dir, file.Name()), nil
}
