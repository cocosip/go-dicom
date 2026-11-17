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
	useKeywords     = flag.Bool("keywords", false, "Use tag keywords as JSON keys instead of PS3.18 8-character tags")
	compactOutput   = flag.Bool("compact", false, "Compact output (no indentation)")
	includeMetaInfo = flag.Bool("meta", false, "Merge file meta information (Group 0002) into the instance object")
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
		jsonData, err = toJSONWithMetaInfo(result, opts)
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

// toJSONWithMetaInfo merges the Group 0002 elements into the main dataset
// object and serializes the result. Note that PS3.18 DICOM JSON has no
// representation for the file meta information, so the merged output is a
// non-standard extension. A failed merge is returned as an error instead of
// emitting a partial document.
func toJSONWithMetaInfo(result *parser.ParseResult, opts []serialization.JSONOption) ([]byte, error) {
	combined := result.FileMetaInformation.Dataset().Clone()
	for _, elem := range result.Dataset.Elements() {
		if err := combined.AddOrUpdate(elem); err != nil {
			return nil, fmt.Errorf("merge element %s: %w", elem.Tag(), err)
		}
	}
	return serialization.ToJSON(combined, opts...)
}
