// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package main provides a command-line tool for converting DICOM files to JSON format.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cocosip/go-dicom/cmd/internal/cli"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/serialization"
)

var (
	outputFile      = flag.String("o", "", "Output file (default: stdout)")
	useKeywords     = flag.Bool("keywords", true, "Use tag keywords as JSON keys")
	compactOutput   = flag.Bool("compact", false, "Compact output (no indentation)")
	includeMetaInfo = flag.Bool("meta", false, "Include file meta information")
)

func main() {
	os.Exit(cli.Command{
		Name:    "dicom2json",
		Usage:   "dicom2json [options] <dicom-file>",
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

	// Prepare JSON options
	var opts []serialization.JSONOption
	if *useKeywords {
		opts = append(opts, serialization.WithWriteTagsAsKeywords(true))
	}
	if !*compactOutput {
		opts = append(opts, serialization.WithIndent("  "))
	}

	// Convert to JSON
	var jsonData []byte
	if *includeMetaInfo && result.FileMetaInformation != nil {
		// Create combined dataset with both meta info and main dataset
		combinedDS := result.FileMetaInformation.Dataset().Clone()
		for _, elem := range result.Dataset.Elements() {
			if err := combinedDS.AddOrUpdate(elem); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to add element %s: %v\n", elem.Tag(), err)
			}
		}
		jsonData, err = serialization.ToJSON(combinedDS, opts...)
	} else {
		jsonData, err = serialization.ToJSON(result.Dataset, opts...)
	}

	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	// Output
	if *outputFile != "" {
		err = os.WriteFile(*outputFile, jsonData, 0600)
		if err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		fmt.Printf("JSON written to %s\n", *outputFile)
	} else {
		fmt.Println(string(jsonData))
	}
	return nil
}
