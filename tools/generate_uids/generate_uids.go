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
	defer file.Close()

	// Create output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// Determine source file name for header comment
	sourceFile := "DicomUIDGenerated.cs"
	if category == "Private" {
		sourceFile = "DicomUIDPrivate.cs"
	}

	// Write header
	writer.WriteString(fmt.Sprintf(`// Copyright (c) 2025 go-dicom contributors.
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
			writer.WriteString(fmt.Sprintf("\t// %s %s\n", varName, name))
			writer.WriteString(fmt.Sprintf("\t%s = New(\"%s\", \"%s\", %s, %v)\n\n",
				varName, uidValue, name, goType, retired))

			uidCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Write footer
	writer.WriteString(")\n\n")
	writer.WriteString("func init() {\n")
	writer.WriteString(fmt.Sprintf("\t// Register all %s UIDs\n", strings.ToLower(category)))

	// Re-read file to generate registration code
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if matches := uidRegex.FindStringSubmatch(line); matches != nil {
			varName := matches[1]
			writer.WriteString(fmt.Sprintf("\tRegister(%s)\n", varName))
		}
	}

	writer.WriteString("}\n")

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
