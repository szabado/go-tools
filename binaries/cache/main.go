package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/alessio/shellescape"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/szabado/go-tools/pkg/cache-persistence"
)

var usage = `cache: A Cache for slow shell commands.

Querying log clusters or curling API endpoints can have a latency that can
make it annoying to build up a pipe pipeline iteratively. This tool caches
those results for you so you iterate quickly.

cache runs the command for you and stores the result, and then returns the
output to you. Any data stored has a TTL of 1 hour, and subsequent calls of
the same command will return the stored result. cache will only store the
results of successful commands: if your bash command has a non-zero exit
code, then it will be uncached.

Usage:
  cache [flags] [command]

Flags:
  -c, --clear, --clean   Clear the cache.
  -o, --overwrite        Overwrite any cache entry for this command.
  -v, --verbose          Verbose logging.

Examples

  cache curl -X GET example.com
`

func main() {
	err := runRoot(os.Args, os.Stdout)

	if err != nil {
		switch tErr := err.(type) {
		case *exec.ExitError:
			os.Exit(tErr.ExitCode())
		case *UsageError:
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			fmt.Fprint(os.Stderr, usage)
			os.Exit(1)
		default:
			fmt.Fprint(os.Stderr, "\n")
			logrus.Errorf("Error: %s\n", err)
		}
		os.Exit(1)
	}
}

func parseArgs(args []string) (verbose bool, clearCache bool, overwrite bool, command string, err error) {
argparse:
	for i := 1; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--clean", "--clear", "-c":
			clearCache = true
		case "--verbose", "-v":
			verbose = true
		case "--overwrite", "-o":
			overwrite = true
		default:
			if strings.HasPrefix(arg, "-") {
				return false, false, false, "", errors.Errorf("unknown flag: %s", arg)
			}
			command = escapeAndJoin(args[i:])
			break argparse
		}
	}

	if clearCache || command != "" {
		return verbose, clearCache, overwrite, command, nil
	} else {
		return verbose, clearCache, overwrite, command, errors.New("command not specified")
	}
}

func runRoot(args []string, output io.Writer) error {
	verbose, clearCache, overwrite, command, err := parseArgs(args)
	if err != nil {
		return NewUsageError(err)
	}

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	if len(command) == 0 && !clearCache {
		return NewUsageError(errors.New("No arguments provided"))
	}

	logrus.Infof("Command: %s", command)
	persister := persistence.NewFsPersister()

	if clearCache {
		logrus.Info("Deleting database")
		return persister.Wipe()
	}

	return runCommand(persister, command, output, overwrite)
}

func runCommand(persister *persistence.FsPersister, command string, output io.Writer, overwrite bool) error {
	if !overwrite {
		err := persister.ReadInto(command, output)
		if err == persistence.ErrKeyNotFound {
			logrus.Debug("No cached result found")
		} else if err != nil && err != persistence.ErrKeyNotFound {
			logrus.WithError(err).Errorf("Unknown error fetching cached result")
		} else if err == nil {
			logrus.Debug("Found cached result, exiting early")
			return nil
		}
	}

	record, err := persister.GetWriterForKey(command)
	if err != nil {
		logrus.WithError(err).Warn("Failed to open record for writing")
	} else {
		output = io.MultiWriter(output, record)
		defer record.Close()
	}

	logrus.Debugf("Executing command")
	return errors.Wrapf(executeCommand(command, output), "error running command")
}

func executeCommand(command string, output io.Writer) error {
	logrus.Infof("Executing command: %s", command)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = output

	return cmd.Run()
}

func escapeAndJoin(cmdSegments []string) string {
	var builder strings.Builder
	for _, seg := range cmdSegments {
		builder.WriteString(shellescape.Quote(seg))
		builder.WriteByte(' ')
	}

	return builder.String()
}
