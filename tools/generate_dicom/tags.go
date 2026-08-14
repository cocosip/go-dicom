// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"

	"github.com/cocosip/go-dicom/tools/internal/dicomxml"
)

func generateTags(dictionary *dicomxml.Dictionary) ([]byte, error) {
	var output bytes.Buffer
	fmt.Fprintf(&output, `// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DICOM Dictionary.xml (version %s). DO NOT EDIT.

package tag

// Standard DICOM tag constants
var (
`, dictionary.Version)

	keywords := make(map[string]struct{}, len(dictionary.Tags))
	for _, dictionaryTag := range dictionary.Tags {
		if dictionaryTag.Keyword == "" {
			return nil, fmt.Errorf("tag (%s,%s) has no keyword", dictionaryTag.Group, dictionaryTag.Element)
		}
		identifier := generatedIdentifier(dictionaryTag.Keyword, dictionaryTag.Retired)
		if _, exists := keywords[identifier]; exists {
			return nil, fmt.Errorf("duplicate tag identifier %q", identifier)
		}
		keywords[identifier] = struct{}{}

		description := strings.Join(strings.Fields(dictionaryTag.Name), " ")
		fmt.Fprintf(
			&output,
			"\t// %s (%s,%s) VR=%s VM=%s %s\n\t%s = New(0x%s, 0x%s)\n\n",
			identifier,
			dictionaryTag.Group,
			dictionaryTag.Element,
			dictionaryTag.VR,
			dictionaryTag.VM,
			description,
			identifier,
			strictNumericTagPart(dictionaryTag.Group),
			strictNumericTagPart(dictionaryTag.Element),
		)
	}
	output.WriteString(")\n")
	return formatGenerated("tags", output.Bytes())
}

func generatedIdentifier(keyword, retired string) string {
	if strings.EqualFold(retired, "true") {
		return keyword + "RETIRED"
	}
	return keyword
}

func strictNumericTagPart(value string) string {
	return strings.NewReplacer("x", "0", "X", "0").Replace(value)
}

func formatGenerated(name string, source []byte) ([]byte, error) {
	formatted, err := format.Source(source)
	if err != nil {
		return nil, fmt.Errorf("format generated %s: %w", name, err)
	}
	return formatted, nil
}
