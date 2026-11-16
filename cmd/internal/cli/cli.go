// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

// Package cli centralises flag parsing, usage text and exit codes for the
// command-line tools in cmd/ so every tool reports help and failures alike.
package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/cocosip/go-dicom/cmd/internal/version"
)

// Command describes one command-line tool.
type Command struct {
	// Name is the binary name reported by -version.
	Name string
	// Usage is the one-line synopsis shown above the flag defaults.
	Usage string
	// MinArgs and MaxArgs bound the positional arguments. A negative MaxArgs
	// means unlimited.
	MinArgs int
	MaxArgs int
}

// Run parses the global flag set and runs run, returning the process exit
// status. Help (-h/-help) and -version exit 0; a missing or extra positional
// argument prints usage to stderr and returns 2; an error returned by run
// prints to stderr and returns 1.
func (c Command) Run(run func() error) int {
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		_, _ = fmt.Fprintf(out, "Usage: %s\n\n", c.Usage)
		flag.PrintDefaults()
	}
	// flag.Parse exits 0 for -h/-help and 2 for an unknown flag, so those
	// cases never reach the checks below.
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s version %s\n", c.Name, version.String())
		return 0
	}
	if err := c.checkArgs(flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		flag.Usage()
		return 2
	}
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}
	return 0
}

func (c Command) checkArgs(args []string) error {
	if len(args) < c.MinArgs {
		return fmt.Errorf("missing required argument")
	}
	if c.MaxArgs >= 0 && len(args) > c.MaxArgs {
		return fmt.Errorf("unexpected argument %q", args[c.MaxArgs])
	}
	return nil
}
