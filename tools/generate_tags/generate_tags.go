// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// This tool generates standard DICOM tag constants from DicomTagGenerated.cs

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Open the C# file
	file, err := os.Open("fo-dicom-code/DicomTagGenerated.cs")
	if err != nil {
		return fmt.Errorf("failed to open DicomTagGenerated.cs: %w", err)
	}
	defer file.Close()

	// Create output file
	outFile, err := os.Create("pkg/dicom/tag/tags_generated.go")
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// Write header
	writer.WriteString(`// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DicomTagGenerated.cs. DO NOT EDIT.

package tag

// Standard DICOM tag constants
var (
`)

	// Regex to match tag definitions
	// ///<summary>(0000,0000) VR=UL VM=1 Command Group Length</summary>
	// public readonly static DicomTag CommandGroupLength = new DicomTag(0x0000, 0x0000);
	commentRegex := regexp.MustCompile(`^\s*///<summary>\(([0-9A-Fa-f]{4}),([0-9A-Fa-f]{4})\)\s+VR=(\w+)\s+VM=([\w\-n]+)\s+(.+)</summary>`)
	tagRegex := regexp.MustCompile(`^\s*public readonly static DicomTag\s+(\w+)\s*=\s*new DicomTag\(0x([0-9A-Fa-f]{4}),\s*0x([0-9A-Fa-f]{4})\);`)

	scanner := bufio.NewScanner(file)
	var currentComment string

	tagCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Check for comment line
		if matches := commentRegex.FindStringSubmatch(line); matches != nil {
			// Extract info from comment
			group := matches[1]
			element := matches[2]
			vr := matches[3]
			vm := matches[4]
			description := matches[5]

			currentComment = fmt.Sprintf("(%s,%s) VR=%s VM=%s %s",
				group, element, vr, vm, description)
			continue
		}

		// Check for tag definition
		if matches := tagRegex.FindStringSubmatch(line); matches != nil {
			tagName := matches[1]
			group := matches[2]
			element := matches[3]

			// Write Go constant
			if currentComment != "" {
				writer.WriteString(fmt.Sprintf("\t// %s %s\n", tagName, currentComment))
			}
			writer.WriteString(fmt.Sprintf("\t%s = New(0x%s, 0x%s)\n\n", tagName, group, element))

			tagCount++
			currentComment = ""
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Write footer
	writer.WriteString(")\n")

	fmt.Printf("Generated %d tag constants\n", tagCount)
	return nil
}
