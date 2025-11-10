// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main provides a command-line tool for converting DICOM files to JSON format.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/serialization"
)

const version = "0.1.0"

func main() {
	// Define flags
	showVersion := flag.Bool("version", false, "Show version information")
	showHelp := flag.Bool("help", false, "Show help information")
	outputFile := flag.String("o", "", "Output file (default: stdout)")
	useKeywords := flag.Bool("keywords", true, "Use tag keywords as JSON keys")
	compact := flag.Bool("compact", false, "Compact output (no indentation)")
	includeMetaInfo := flag.Bool("meta", false, "Include file meta information")

	flag.Parse()

	// Show version
	if *showVersion {
		fmt.Printf("dicom2json version %s\n", version)
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

	// Prepare JSON options
	var opts []serialization.JSONOption
	if *useKeywords {
		opts = append(opts, serialization.WithWriteTagsAsKeywords(true))
	}
	if !*compact {
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
		fmt.Fprintf(os.Stderr, "Error: Failed to convert to JSON: %v\n", err)
		os.Exit(1)
	}

	// Output
	if *outputFile != "" {
		err = os.WriteFile(*outputFile, jsonData, 0600)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to write output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("JSON written to %s\n", *outputFile)
	} else {
		fmt.Println(string(jsonData))
	}
}

func printUsage() {
	fmt.Println("Usage: dicom2json [options] <dicom-file>")
	fmt.Println()
	fmt.Println("Convert a DICOM file to JSON format (DICOM JSON Model)")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -version       Show version information")
	fmt.Println("  -help          Show this help message")
	fmt.Println("  -o FILE        Output to file instead of stdout")
	fmt.Println("  -keywords      Use tag keywords as JSON keys (default: true)")
	fmt.Println("  -compact       Compact output, no indentation (default: false)")
	fmt.Println("  -meta          Include file meta information (default: false)")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  dicom2json image.dcm")
	fmt.Println("  dicom2json -o output.json image.dcm")
	fmt.Println("  dicom2json -compact -meta image.dcm > output.json")
}
