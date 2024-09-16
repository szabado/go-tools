package main

import (
	"os"
	"strings"
	"unicode"

	"github.com/szabado/go-tools/pkg/ssh"
)

func split(hostList string) []string {
	return strings.FieldsFunc(hostList, func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	})
}

func parseHostsArg(hostsArg string) ([]*ssh.Host, error) {
	hosts := make([]*ssh.Host, 0)
	for _, hostArg := range split(hostsArg) {
		hosts = append(hosts, ssh.ParseHostString(hostArg))
	}
	return hosts, nil
}

func loadFileContents(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
