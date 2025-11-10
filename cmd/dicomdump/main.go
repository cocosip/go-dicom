// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main provides a command-line tool for dumping DICOM files.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
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
		dumpDataset(result.FileMetaInformation.Dataset(), 0, *maxDepth, *showValues, *compact)
		fmt.Println()
	}

	// Dump main dataset
	fmt.Println("# Main Dataset")
	dumpDataset(result.Dataset, 0, *maxDepth, *showValues, *compact)
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

func dumpDataset(ds *dataset.Dataset, depth int, maxDepth int, showValues bool, compact bool) {
	if maxDepth >= 0 && depth > maxDepth {
		return
	}

	indent := strings.Repeat("  ", depth)

	for _, elem := range ds.Elements() {
		tag := elem.Tag()
		vr := elem.ValueRepresentation()

		// Print tag and VR
		if compact {
			fmt.Printf("%s%s %s ", indent, tag.String(), vr.Code())
		} else {
			fmt.Printf("%s%-20s %s  ", indent, tag.String(), vr.Code())
		}

		// Check if it's a sequence
		if seq, ok := elem.(*dataset.Sequence); ok {
			itemCount := seq.Count()
			if compact {
				fmt.Printf("(Sequence with %d item(s))\n", itemCount)
			} else {
				fmt.Printf("(Sequence with %d item(s))\n", itemCount)
			}

			// Dump sequence items
			if maxDepth < 0 || depth < maxDepth {
				for i := 0; i < itemCount; i++ {
					item := seq.GetItem(i)
					if compact {
						fmt.Printf("%s  Item #%d:\n", indent, i+1)
					} else {
						fmt.Printf("%s  --- Item #%d ---\n", indent, i+1)
					}
					dumpDataset(item, depth+1, maxDepth, showValues, compact)
				}
			}
		} else {
			// Print value
			if showValues {
				value := formatElementValue(elem, compact)
				if compact {
					fmt.Printf("%s\n", value)
				} else {
					fmt.Printf("%-40s\n", value)
				}
			} else {
				fmt.Printf("\n")
			}
		}
	}
}

func formatElementValue(elem element.Element, compact bool) string {
	// Try string-based elements first
	if val := tryFormatString(elem, compact); val != "" {
		return val
	}

	// Try numeric elements
	if val := tryFormatNumeric(elem, compact); val != "" {
		return val
	}

	// For binary data, show size
	if val := tryFormatBinary(elem); val != "" {
		return val
	}

	return "(empty)"
}

func tryFormatString(elem element.Element, compact bool) string {
	if strElem, ok := elem.(interface{ GetValue() (string, error) }); ok {
		if val, err := strElem.GetValue(); err == nil {
			if compact && len(val) > 60 {
				return val[:60] + "..."
			}
			return val
		}
	}
	return ""
}

func tryFormatNumeric(elem element.Element, compact bool) string {
	vr := elem.ValueRepresentation()
	vrCode := vr.Code()

	switch vrCode {
	case "US": // Unsigned Short
		if usElem, ok := elem.(interface{ GetValues() ([]uint16, error) }); ok {
			if vals, err := usElem.GetValues(); err == nil {
				return formatNumericArray(vals, compact)
			}
		}
	case "UL": // Unsigned Long
		if ulElem, ok := elem.(interface{ GetValues() ([]uint32, error) }); ok {
			if vals, err := ulElem.GetValues(); err == nil {
				return formatNumericArray(vals, compact)
			}
		}
	case "SS": // Signed Short
		if ssElem, ok := elem.(interface{ GetValues() ([]int16, error) }); ok {
			if vals, err := ssElem.GetValues(); err == nil {
				return formatNumericArray(vals, compact)
			}
		}
	case "SL": // Signed Long
		if slElem, ok := elem.(interface{ GetValues() ([]int32, error) }); ok {
			if vals, err := slElem.GetValues(); err == nil {
				return formatNumericArray(vals, compact)
			}
		}
	case "FL": // Float
		if flElem, ok := elem.(interface{ GetValues() ([]float32, error) }); ok {
			if vals, err := flElem.GetValues(); err == nil {
				return formatNumericArray(vals, compact)
			}
		}
	case "FD": // Double
		if fdElem, ok := elem.(interface{ GetValues() ([]float64, error) }); ok {
			if vals, err := fdElem.GetValues(); err == nil {
				return formatNumericArray(vals, compact)
			}
		}
	}
	return ""
}

func tryFormatBinary(elem element.Element) string {
	buf := elem.Buffer()
	if buf != nil {
		size := buf.Size()
		if size > 0 {
			return fmt.Sprintf("(Binary data, %d bytes)", size)
		}
	}
	return ""
}

func formatNumericArray[T any](vals []T, compact bool) string {
	if len(vals) == 0 {
		return "(empty)"
	}
	if len(vals) == 1 {
		return fmt.Sprintf("%v", vals[0])
	}
	if compact && len(vals) > 5 {
		return fmt.Sprintf("[%v, %v, %v, ... (%d values)]", vals[0], vals[1], vals[2], len(vals))
	}
	return fmt.Sprintf("%v", vals)
}
