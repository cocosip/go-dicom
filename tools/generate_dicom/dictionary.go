// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/cocosip/go-dicom/tools/internal/dicomxml"
)

func generateDictionary(dictionary *dicomxml.Dictionary) ([]byte, error) {
	var output bytes.Buffer
	fmt.Fprintf(&output, `// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DICOM Dictionary.xml (version %s). DO NOT EDIT.

package dict

import (
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// loadStandardEntries loads all standard DICOM dictionary entries.
func loadStandardEntries(d *Dictionary) {
`, dictionary.Version)

	for _, dictionaryTag := range dictionary.Tags {
		if err := writeDictionaryEntry(&output, dictionaryTag, "d"); err != nil {
			return nil, err
		}
	}
	output.WriteString("}\n")
	return formatGenerated("dictionary", output.Bytes())
}

func generatePrivateDictionary(dictionaries *dicomxml.PrivateDictionaries) ([]byte, int, error) {
	var output bytes.Buffer
	output.WriteString(`// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from Private Dictionary.xml. DO NOT EDIT.

package dict

import (
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vm"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// loadPrivateEntries loads all bundled private DICOM dictionary entries.
func loadPrivateEntries(d *Dictionary) {
`)

	creatorNames := make(map[string]struct{}, len(dictionaries.Dictionaries))
	entryCount := 0
	for _, privateDictionary := range dictionaries.Dictionaries {
		creator := strings.TrimSpace(privateDictionary.Creator)
		if creator == "" {
			return nil, 0, fmt.Errorf("private dictionary has no creator")
		}
		if _, exists := creatorNames[creator]; exists {
			return nil, 0, fmt.Errorf("duplicate private creator %q", creator)
		}
		creatorNames[creator] = struct{}{}

		output.WriteString("\t{\n")
		fmt.Fprintf(&output, "\t\tprivateDictionary := d.GetPrivateDictionary(%s)\n", strconv.Quote(creator))
		for _, dictionaryTag := range privateDictionary.Tags {
			if err := writeDictionaryEntry(&output, dictionaryTag, "privateDictionary"); err != nil {
				return nil, 0, fmt.Errorf("private creator %q: %w", creator, err)
			}
			entryCount++
		}
		output.WriteString("\t}\n")
	}
	output.WriteString("}\n")

	formatted, err := formatGenerated("private dictionary", output.Bytes())
	if err != nil {
		return nil, 0, err
	}
	return formatted, entryCount, nil
}

func writeDictionaryEntry(output *bytes.Buffer, dictionaryTag dicomxml.Tag, dictionaryVariable string) error {
	if dictionaryTag.Group == "" || dictionaryTag.Element == "" {
		return fmt.Errorf("tag is missing group or element")
	}
	vmCode := getVMCode(dictionaryTag.VM)
	vrCodes, err := getVRCodes(dictionaryTag.VR)
	if err != nil {
		return fmt.Errorf("tag (%s,%s): %w", dictionaryTag.Group, dictionaryTag.Element, err)
	}
	retired := strings.EqualFold(dictionaryTag.Retired, "true")

	if strings.ContainsAny(dictionaryTag.Group+dictionaryTag.Element, "xX") {
		fmt.Fprintf(output, "\t%s.Add(NewEntryWithMask(\n", dictionaryVariable)
		fmt.Fprintf(output, "\t\ttag.MustParseMaskedTag(%s),\n", strconv.Quote(fmt.Sprintf("(%s,%s)", dictionaryTag.Group, dictionaryTag.Element)))
	} else {
		fmt.Fprintf(output, "\t%s.Add(NewEntry(\n", dictionaryVariable)
		fmt.Fprintf(output, "\t\ttag.New(0x%s, 0x%s),\n", dictionaryTag.Group, dictionaryTag.Element)
	}
	fmt.Fprintf(output, "\t\t%s,\n", strconv.Quote(strings.TrimSpace(dictionaryTag.Name)))
	fmt.Fprintf(output, "\t\t%s,\n", strconv.Quote(strings.TrimSpace(dictionaryTag.Keyword)))
	fmt.Fprintf(output, "\t\t%s,\n", vmCode)
	fmt.Fprintf(output, "\t\t%t,\n", retired)
	fmt.Fprintf(output, "\t\t%s,\n", strings.Join(vrCodes, ", "))
	output.WriteString("\t))\n")
	return nil
}

func getVRCodes(value string) ([]string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return []string{"vr.None"}, nil
	}
	value = strings.ReplaceAll(value, " or ", "/")
	parts := strings.FieldsFunc(value, func(character rune) bool {
		return character == '_' || character == '/' || character == '\\' ||
			character == ',' || character == '|' || unicode.IsSpace(character)
	})
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid VR %q", value)
	}
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if len(part) != 2 {
			return nil, fmt.Errorf("invalid VR %q", part)
		}
		result = append(result, "vr."+part)
	}
	return result, nil
}

func getVMCode(value string) string {
	value = strings.TrimSpace(strings.Split(value, " or ")[0])
	switch value {
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
		return "vm.VM12"
	case "1-3":
		return "vm.VM13"
	case "1-8":
		return "vm.VM18"
	case "1-32":
		return "vm.VM132"
	case "1-99":
		return "vm.VM199"
	case "1-n":
		return "vm.VM1N"
	case "2-n":
		return "vm.VM2N"
	case "2-2n":
		return "vm.VM22N"
	case "3-n":
		return "vm.VM3N"
	case "3-3n":
		return "vm.VM33N"
	default:
		return fmt.Sprintf("vm.MustParse(%s)", strconv.Quote(value))
	}
}
