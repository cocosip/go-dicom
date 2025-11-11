// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// This tool generates standard DICOM UID constants from DicomUIDGenerated.cs

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// mustWrite wraps writer operations and panics on error (for code generation)
func mustWrite(writer *bufio.Writer, s string) {
	if _, err := writer.WriteString(s); err != nil {
		panic(fmt.Sprintf("failed to write output: %v", err))
	}
}

// mustFlush wraps writer.Flush and panics on error (for code generation)
func mustFlush(writer *bufio.Writer) {
	if err := writer.Flush(); err != nil {
		panic(fmt.Sprintf("failed to flush output: %v", err))
	}
}

func run() error {
	// Process standard UIDs
	if err := processFile("fo-dicom-code/DicomUIDGenerated.cs", "pkg/dicom/uid/uids_generated.go", "Standard"); err != nil {
		return err
	}

	// Process private UIDs
	if err := processFile("fo-dicom-code/DicomUIDPrivate.cs", "pkg/dicom/uid/uids_private.go", "Private"); err != nil {
		return err
	}

	return nil
}

func processFile(inputFile, outputFile, category string) error {
	// Open the C# file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", inputFile, err)
	}
	defer func() { _ = file.Close() }()

	// Create output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() { _ = outFile.Close() }()

	writer := bufio.NewWriter(outFile)
	defer mustFlush(writer)

	// Determine source file name for header comment
	sourceFile := "DicomUIDGenerated.cs"
	if category == "Private" {
		sourceFile = "DicomUIDPrivate.cs"
	}

	// Write header
	mustWrite(writer, fmt.Sprintf(`// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from %s. DO NOT EDIT.

package uid

// %s DICOM UID constants
var (
`, sourceFile, category))

	// Regex to match UID definitions
	// Format 1: new DicomUID("uid", "name", DicomUidType.XXX, true|false)
	// Format 2: new DicomUID("uid", "name", DicomUidType.XXX) - retired defaults to false
	uidRegex := regexp.MustCompile(`public static readonly DicomUID\s+(\w+)\s*=\s*new DicomUID\("([^"]+)",\s*"([^"]+)",\s*DicomUidType\.(\w+)(?:,\s*(true|false))?\)`)

	scanner := bufio.NewScanner(file)
	uidCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if matches := uidRegex.FindStringSubmatch(line); matches != nil {
			varName := matches[1]
			uidValue := matches[2]
			name := matches[3]
			uidType := matches[4]
			// matches[5] will be empty string if no retired parameter, or "true"/"false" if present
			retired := matches[5] == "true"

			// Map C# UID type to Go type
			goType := mapUIDType(uidType)

			// Escape quotes in name
			name = strings.ReplaceAll(name, `"`, `\"`)

			// Write UID constant
			mustWrite(writer, fmt.Sprintf("\t// %s %s\n", varName, name))
			mustWrite(writer, fmt.Sprintf("\t%s = New(\"%s\", \"%s\", %s, %v)\n\n",
				varName, uidValue, name, goType, retired))

			uidCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Write footer
	mustWrite(writer, ")\n\n")
	mustWrite(writer, "func init() {\n")
	mustWrite(writer, fmt.Sprintf("\t// Register all %s UIDs\n", strings.ToLower(category)))

	// Re-read file to generate registration code
	_, _ = file.Seek(0, 0) // Ignore seek errors in code generator
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if matches := uidRegex.FindStringSubmatch(line); matches != nil {
			varName := matches[1]
			mustWrite(writer, fmt.Sprintf("\tRegister(%s)\n", varName))
		}
	}

	mustWrite(writer, "}\n")

	fmt.Printf("Generated %d %s UID constants in %s\n", uidCount, category, outputFile)
	return nil
}

func mapUIDType(csType string) string {
	switch csType {
	case "TransferSyntax":
		return "TypeTransferSyntax"
	case "SOPClass":
		return "TypeSOPClass"
	case "MetaSOPClass":
		return "TypeMetaSOPClass"
	case "ServiceClass":
		return "TypeServiceClass"
	case "SOPInstance":
		return "TypeSOPInstance"
	case "ApplicationContextName":
		return "TypeApplicationContextName"
	case "ApplicationHostingModel":
		return "TypeApplicationHostingModel"
	case "CodingScheme":
		return "TypeCodingScheme"
	case "FrameOfReference":
		return "TypeFrameOfReference"
	case "LDAP":
		return "TypeLDAP"
	case "MappingResource":
		return "TypeMappingResource"
	case "ContextGroupName":
		return "TypeContextGroupName"
	default:
		return "TypeUnknown"
	}
}
