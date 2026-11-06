// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package main provides a command-line tool for dumping DICOM files.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
)

const version = "0.1.0"

func main() {
	// Define flags
	showVersion := flag.Bool("version", false, "Show version information")
	showHelp := flag.Bool("help", false, "Show help information")
	maxDepth := flag.Int("depth", -1, "Maximum depth for nested sequences (-1 for unlimited)")
	showValues := flag.Bool("values", true, "Show element values")
	compact := flag.Bool("compact", false, "Compact output (one line per element)")

	flag.Parse()

	// Show version
	if *showVersion {
		fmt.Printf("dicomdump version %s\n", version)
		os.Exit(0)
	}

	// Show help
	if *showHelp || flag.NArg() == 0 {
		printUsage()
		os.Exit(0)
	}

	// Get input file
	inputFile := flag.Arg(0)

	// Parse DICOM file
	result, err := parser.ParseFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to parse DICOM file: %v\n", err)
		os.Exit(1)
	}

	// Dump file meta information
	if result.FileMetaInformation != nil {
		fmt.Println("# File Meta Information")
		if err := dumpDataset(os.Stdout, result.FileMetaInformation.Dataset(), *maxDepth, *showValues, *compact); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to dump File Meta Information: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()
	}

	// Dump main dataset
	fmt.Println("# Main Dataset")
	if err := dumpDataset(os.Stdout, result.Dataset, *maxDepth, *showValues, *compact); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to dump Dataset: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: dicomdump [options] <dicom-file>")
	fmt.Println()
	fmt.Println("Dump all DICOM tags in a file")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -version       Show version information")
	fmt.Println("  -help          Show this help message")
	fmt.Println("  -depth N       Maximum depth for nested sequences (-1 for unlimited, default: -1)")
	fmt.Println("  -values        Show element values (default: true)")
	fmt.Println("  -compact       Compact output, one line per element (default: false)")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  dicomdump image.dcm")
	fmt.Println("  dicomdump -depth 2 -compact image.dcm")
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
