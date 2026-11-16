// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package main provides a command-line tool for dumping DICOM files.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cocosip/go-dicom/cmd/internal/cli"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
)

var (
	maxDepth   = flag.Int("depth", -1, "Maximum depth for nested sequences (-1 for unlimited)")
	showValues = flag.Bool("values", true, "Show element values")
	compact    = flag.Bool("compact", false, "Compact output (one line per element)")
)

func main() {
	os.Exit(cli.Command{
		Name:    "dicomdump",
		Usage:   "dicomdump [options] <dicom-file>",
		MinArgs: 1,
		MaxArgs: 1,
	}.Run(run))
}

func run() error {
	// Get input file
	inputFile := flag.Arg(0)

	// Parse DICOM file
	result, err := parser.ParseFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to parse DICOM file: %w", err)
	}

	// Dump file meta information
	if result.FileMetaInformation != nil {
		fmt.Println("# File Meta Information")
		if err := dumpDataset(os.Stdout, result.FileMetaInformation.Dataset(), *maxDepth, *showValues, *compact); err != nil {
			return fmt.Errorf("failed to dump File Meta Information: %w", err)
		}
		fmt.Println()
	}

	// Dump main dataset
	fmt.Println("# Main Dataset")
	if err := dumpDataset(os.Stdout, result.Dataset, *maxDepth, *showValues, *compact); err != nil {
		return fmt.Errorf("failed to dump Dataset: %w", err)
	}
	return nil
}

func dumpDataset(output io.Writer, ds *dataset.Dataset, maxDepth int, showValues bool, compact bool) error {
	layout := dataset.LayoutDump
	if compact {
		layout = dataset.LayoutRows
	}
	text := ds.Format(dataset.FormatOptions{
		Layout:     layout,
		MaxDepth:   maxDepth,
		HideValues: !showValues,
	})
	if text == "" {
		return nil
	}
	_, err := fmt.Fprintln(output, text)
	return err
}
