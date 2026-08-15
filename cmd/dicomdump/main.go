// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main provides a command-line tool for dumping DICOM files.
package main

import (
	"flag"
	"fmt"
	"io"
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
	return dataset.Walk(ds, func(event dataset.WalkEvent) (dataset.WalkAction, error) {
		depth := walkDepth(event.Path)
		switch event.Kind {
		case dataset.WalkElement, dataset.WalkFragmentBegin:
			if err := writeDumpElement(output, event.Element, depth, showValues, compact); err != nil {
				return dataset.WalkContinue, err
			}

		case dataset.WalkSequenceBegin:
			sequence := event.Element.(*dataset.Sequence)
			if err := writeDumpSequence(output, sequence, depth, compact); err != nil {
				return dataset.WalkContinue, err
			}
			if maxDepth >= 0 && depth >= maxDepth {
				return dataset.WalkSkipChildren, nil
			}

		case dataset.WalkSequenceItemBegin:
			itemIndex := event.Path[len(event.Path)-1].ItemIndex
			if itemIndex == nil {
				return dataset.WalkContinue, fmt.Errorf("sequence item path has no item index")
			}
			indent := strings.Repeat("  ", depth)
			if compact {
				_, err := fmt.Fprintf(output, "%sItem #%d:\n", indent, *itemIndex+1)
				return dataset.WalkContinue, err
			}
			_, err := fmt.Fprintf(output, "%s--- Item #%d ---\n", indent, *itemIndex+1)
			return dataset.WalkContinue, err
		}
		return dataset.WalkContinue, nil
	})
}

func walkDepth(path dataset.Path) int {
	depth := 0
	for _, segment := range path {
		if segment.ItemIndex != nil {
			depth++
		}
	}
	return depth
}

func writeDumpSequence(output io.Writer, sequence *dataset.Sequence, depth int, compact bool) error {
	indent := strings.Repeat("  ", depth)
	if compact {
		_, err := fmt.Fprintf(output, "%s%s %s (Sequence with %d item(s))\n",
			indent, sequence.Tag(), sequence.ValueRepresentation().Code(), sequence.Count())
		return err
	}
	_, err := fmt.Fprintf(output, "%s%-20s %s  (Sequence with %d item(s))\n",
		indent, sequence.Tag(), sequence.ValueRepresentation().Code(), sequence.Count())
	return err
}

func writeDumpElement(output io.Writer, elem element.Element, depth int, showValues bool, compact bool) error {
	indent := strings.Repeat("  ", depth)
	if compact {
		if _, err := fmt.Fprintf(output, "%s%s %s ", indent, elem.Tag(), elem.ValueRepresentation().Code()); err != nil {
			return err
		}
	} else if _, err := fmt.Fprintf(output, "%s%-20s %s  ", indent, elem.Tag(), elem.ValueRepresentation().Code()); err != nil {
		return err
	}
	if !showValues {
		_, err := fmt.Fprintln(output)
		return err
	}
	value := formatElementValue(elem, compact)
	if compact {
		_, err := fmt.Fprintln(output, value)
		return err
	}
	_, err := fmt.Fprintf(output, "%-40s\n", value)
	return err
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
	if !isTextVR(elem.ValueRepresentation().Code()) {
		return ""
	}
	values, err := element.CanonicalStrings(elem)
	if err != nil {
		return ""
	}
	value := strings.Join(values, "\\")
	if compact && len(value) > 60 {
		return value[:60] + "..."
	}
	return value
}

func isTextVR(code string) bool {
	switch code {
	case "AE", "AS", "CS", "DA", "DS", "DT", "IS", "LO", "LT", "PN",
		"SH", "ST", "TM", "UC", "UI", "UR", "UT":
		return true
	default:
		return false
	}
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
