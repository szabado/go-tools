package main

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	timeout time.Duration  = 200 * time.Millisecond
	output  io.Writer      = os.Stdout
	input   io.Reader      = os.Stdin
	osExit  func(code int) = os.Exit
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging")
}

var rootCmd = &cobra.Command{
	Use:   "idiff",
	Short: "A tool for diffing input pasted in.",
	PersistentPreRunE: func(c *cobra.Command, args []string) error {
		if verbose {
			logrus.SetLevel(logrus.InfoLevel)
		} else {
			logrus.SetLevel(logrus.ErrorLevel)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeDiff(input, output, timeout)
	},
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		osExit(1)
	}
	osExit(0)
}
