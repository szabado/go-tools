package main

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbose bool
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
		err := executeDiff(input, output)
		logrus.Info(err)
		return err
	},
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		osExit(1)
	}
}
