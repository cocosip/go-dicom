// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// This tool generates standard DICOM dictionary entries from DICOM Dictionary.xml

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Dictionary struct {
	Version string `xml:"version,attr"`
	Tags    []Tag  `xml:"tag"`
}

type Tag struct {
	Group   string `xml:"group,attr"`
	Element string `xml:"element,attr"`
	Keyword string `xml:"keyword,attr"`
	VR      string `xml:"vr,attr"`
	VM      string `xml:"vm,attr"`
	Retired string `xml:"retired,attr"`
	Name    string `xml:",chardata"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Open the XML dictionary file
	file, err := os.Open("fo-dicom-code/Dictionaries/DICOM Dictionary.xml")
	if err != nil {
		return fmt.Errorf("failed to open DICOM Dictionary.xml: %w", err)
	}
	defer file.Close()

	// Read all XML data
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read XML data: %w", err)
	}

	// Parse XML
	var dict Dictionary
	if err := xml.Unmarshal(data, &dict); err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}

	// Create output file
	outFile, err := os.Create("pkg/dicom/dict/dictionary_data.go")
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	// Write header
	fmt.Fprintf(outFile, `// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DICOM Dictionary.xml. DO NOT EDIT.

package dict

import (
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// loadStandardEntries loads all standard DICOM dictionary entries.
func loadStandardEntries(d *Dictionary) {
`)

	// Sort tags by group and element for consistent output
	sort.Slice(dict.Tags, func(i, j int) bool {
		if dict.Tags[i].Group != dict.Tags[j].Group {
			return dict.Tags[i].Group < dict.Tags[j].Group
		}
		return dict.Tags[i].Element < dict.Tags[j].Element
	})

	// Track masked tags (e.g., group length)
	maskedTags := make(map[string]bool)

	// First pass: identify masked tags
	for _, t := range dict.Tags {
		if strings.Contains(t.Group, "x") || strings.Contains(t.Element, "x") {
			maskedTags[t.Keyword] = true
		}
	}

	// Generate dictionary entries
	count := 0
	for _, t := range dict.Tags {
		retired := t.Retired == "true"

		// Escape quotes in name
		name := strings.ReplaceAll(t.Name, `"`, `\"`)

		// Handle VR - can be multiple separated by " or " or "/"
		// Some special tags (like Item delimiters) don't have VR
		var vrCode string
		if t.VR == "" {
			// For tags without VR, use None
			vrCode = "vr.None"
		} else {
			// Try splitting by " or " first, then by "/"
			vrList := strings.Split(t.VR, " or ")
			if len(vrList) == 1 {
				vrList = strings.Split(t.VR, "/")
			}

			if len(vrList) == 1 {
				vrCode = fmt.Sprintf("vr.%s", vrList[0])
			} else {
				vrCodes := make([]string, len(vrList))
				for i, v := range vrList {
					vrCodes[i] = fmt.Sprintf("vr.%s", strings.TrimSpace(v))
				}
				vrCode = strings.Join(vrCodes, ", ")
			}
		}

		// Handle VM - clean up "or" cases (e.g., "1-n or 1" -> "1-n")
		vm := strings.TrimSpace(t.VM)
		if strings.Contains(vm, " or ") {
			// Take the first (more general) part
			parts := strings.Split(vm, " or ")
			vm = strings.TrimSpace(parts[0])
		}
		vmCode := getVMCode(vm)

		// Check if this is a masked tag
		if maskedTags[t.Keyword] {
			// Generate masked tag entry
			maskPattern := fmt.Sprintf("(%s,%s)", t.Group, t.Element)
			fmt.Fprintf(outFile, "\td.Add(NewEntryWithMask(\n")
			fmt.Fprintf(outFile, "\t\ttag.MustParseMaskedTag(%q),\n", maskPattern)
			fmt.Fprintf(outFile, "\t\t%q, // %s\n", name, t.Keyword)
			fmt.Fprintf(outFile, "\t\t%q,\n", t.Keyword)
			fmt.Fprintf(outFile, "\t\t%s,\n", vmCode)
			fmt.Fprintf(outFile, "\t\t%v,\n", retired)
			fmt.Fprintf(outFile, "\t\t%s,\n", vrCode)
			fmt.Fprintf(outFile, "\t))\n")
		} else {
			// Generate normal tag entry
			fmt.Fprintf(outFile, "\td.Add(NewEntry(\n")
			fmt.Fprintf(outFile, "\t\ttag.New(0x%s, 0x%s),\n", t.Group, t.Element)
			fmt.Fprintf(outFile, "\t\t%q, // %s\n", name, t.Keyword)
			fmt.Fprintf(outFile, "\t\t%q,\n", t.Keyword)
			fmt.Fprintf(outFile, "\t\t%s,\n", vmCode)
			fmt.Fprintf(outFile, "\t\t%v,\n", retired)
			fmt.Fprintf(outFile, "\t\t%s,\n", vrCode)
			fmt.Fprintf(outFile, "\t))\n")
		}

		count++
	}

	// Write footer
	fmt.Fprintf(outFile, "}\n")

	fmt.Printf("Generated %d dictionary entries\n", count)
	return nil
}

// getVMCode converts VM string to Go code
func getVMCode(vmStr string) string {
	switch vmStr {
	case "1":
		return "vm.VM1"
	case "2":
		return "vm.VM2"
	case "3":
		return "vm.VM3"
	case "4":
		return "vm.VM4"
	case "6":
		return "vm.VM6"
	case "16":
		return "vm.VM16"
	case "1-2":
		return "vm.VM1_2"
	case "1-3":
		return "vm.VM1_3"
	case "1-8":
		return "vm.VM1_8"
	case "1-32":
		return "vm.VM1_32"
	case "1-99":
		return "vm.VM1_99"
	case "1-n":
		return "vm.VM1_n"
	case "2-n":
		return "vm.VM2_n"
	case "2-2n":
		return "vm.VM2_2n"
	case "3-n":
		return "vm.VM3_n"
	case "3-3n":
		return "vm.VM3_3n"
	default:
		// For unknown VM, try to construct the code
		return fmt.Sprintf("vm.MustParse(%q)", vmStr)
	}
}
