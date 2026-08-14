// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/tools/internal/dicomxml"
)

func generateUIDs(dictionary *dicomxml.Dictionary) ([]byte, error) {
	var output bytes.Buffer
	fmt.Fprintf(&output, `// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Code generated from DICOM Dictionary.xml (version %s). DO NOT EDIT.

package uid

// Standard DICOM UID constants
var (
`, dictionary.Version)

	keywords := make(map[string]struct{}, len(dictionary.UIDs))
	for _, dictionaryUID := range dictionary.UIDs {
		if dictionaryUID.Keyword == "" {
			return nil, fmt.Errorf("UID %q has no keyword", dictionaryUID.Value)
		}
		identifier := generatedIdentifier(dictionaryUID.Keyword, dictionaryUID.Retired)
		if _, exists := keywords[identifier]; exists {
			return nil, fmt.Errorf("duplicate UID identifier %q", identifier)
		}
		keywords[identifier] = struct{}{}

		uidType, err := mapUIDType(dictionaryUID.Type)
		if err != nil {
			return nil, fmt.Errorf("UID %s (%s): %w", dictionaryUID.Value, dictionaryUID.Keyword, err)
		}
		name := strings.Join(strings.Fields(dictionaryUID.Name), " ")
		fmt.Fprintf(&output, "\t// %s %s\n", identifier, name)
		fmt.Fprintf(
			&output,
			"\t%s = New(%s, %s, %s, %t)\n\n",
			identifier,
			strconv.Quote(dictionaryUID.Value),
			strconv.Quote(name),
			uidType,
			strings.EqualFold(dictionaryUID.Retired, "true"),
		)
	}

	output.WriteString(")\n\nfunc init() {\n")
	for _, dictionaryUID := range dictionary.UIDs {
		fmt.Fprintf(
			&output,
			"\tRegister(%s)\n",
			generatedIdentifier(dictionaryUID.Keyword, dictionaryUID.Retired),
		)
	}
	output.WriteString("}\n")
	return formatGenerated("UIDs", output.Bytes())
}

func mapUIDType(value string) (string, error) {
	switch value {
	case "Transfer Syntax":
		return "TypeTransferSyntax", nil
	case "SOP Class":
		return "TypeSOPClass", nil
	case "Meta SOP Class":
		return "TypeMetaSOPClass", nil
	case "Service Class":
		return "TypeServiceClass", nil
	case "Well-known SOP Instance":
		return "TypeSOPInstance", nil
	case "Application Context Name":
		return "TypeApplicationContextName", nil
	case "Application Hosting Model":
		return "TypeApplicationHostingModel", nil
	case "Coding Scheme", "DICOM UIDs as a Coding Scheme":
		return "TypeCodingScheme", nil
	case "Synchronization Frame of Reference":
		return "TypeFrameOfReference", nil
	case "LDAP OID":
		return "TypeLDAP", nil
	case "Mapping Resource":
		return "TypeMappingResource", nil
	case "Context Group Name":
		return "TypeContextGroupName", nil
	default:
		return "", fmt.Errorf("unsupported UID type %q", value)
	}
}
