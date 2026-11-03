// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// DefaultValueLimit is the default maximum number of runes rendered for one
// element value by Summary and FormatValue when the caller passes zero.
const DefaultValueLimit = 64

const emptyValue = "(empty)"

// Summary returns a bounded, single-line tag/VR/value representation of an
// element. It is safe to use in logs, error messages, and %v formatting:
// container and bulk binary values are summarized and never decoded.
func Summary(elem Element) string {
	if elem == nil {
		return "<nil element>"
	}
	result := formatValueTag(elem.Tag()) + " " + formatValueVR(elem.ValueRepresentation())
	if value := FormatValue(elem, DefaultValueLimit); value != "" {
		result += " " + value
	}
	return result
}

// FormatValue renders an element value as bounded, single-line text for
// diagnostics and logs. Text values are bracketed, multi-values are separated
// by backslashes, sequences return an empty string (rendered by the Dataset
// formatter), and bulk binary values or fragments are summarized by size.
func FormatValue(elem Element, maxLen int) string {
	if elem == nil {
		return "(nil element)"
	}
	if maxLen <= 0 {
		maxLen = DefaultValueLimit
	}

	if fragments, ok := elem.(interface {
		FragmentCount() int
		OffsetTable() []uint32
	}); ok {
		fragmentCount := fragments.FragmentCount()
		fragmentLabel := "fragments"
		if fragmentCount == 1 {
			fragmentLabel = "fragment"
		}
		offsetCount := len(fragments.OffsetTable())
		offsetLabel := "offsets"
		if offsetCount == 1 {
			offsetLabel = "offset"
		}
		return fmt.Sprintf("(%d %s, %d %s)", fragmentCount, fragmentLabel, offsetCount, offsetLabel)
	}

	vrCode := ""
	if value := elem.ValueRepresentation(); value != nil {
		vrCode = value.Code()
	}
	switch vrCode {
	case vr.CodeSQ:
		return ""
	case vr.CodeOB, vr.CodeOW, vr.CodeOL, vr.CodeOD, vr.CodeOF, vr.CodeOV, vr.CodeUN:
		if elem.Length() == 0 {
			return emptyValue
		}
		return fmt.Sprintf("(%d bytes)", elem.Length())
	}

	values, err := CanonicalStrings(elem)
	if err != nil || len(values) == 0 {
		if elem.Length() == 0 {
			return emptyValue
		}
		return fmt.Sprintf("(%d bytes)", elem.Length())
	}

	text, truncated := joinBounded(values, maxLen, '\\')
	if text == "" {
		return emptyValue
	}
	if isDisplayTextVR(vrCode) {
		if truncated {
			return "[" + text + "...]"
		}
		return "[" + text + "]"
	}
	if truncated {
		return text + "..."
	}
	return text
}

// isDisplayTextVR reports whether values of the VR are rendered as DICOM text.
func isDisplayTextVR(code string) bool {
	switch code {
	case "AE", "AS", "CS", "DA", "DS", "DT", "IS", "LO", "LT", "PN",
		"SH", "ST", "TM", "UC", "UI", "UR", "UT":
		return true
	default:
		return false
	}
}

func joinBounded(values []string, maxLen int, separator byte) (string, bool) {
	var builder strings.Builder
	count := 0
	for index, value := range values {
		if index > 0 {
			if count >= maxLen {
				return builder.String(), true
			}
			builder.WriteByte(separator)
			count++
		}
		for _, current := range value {
			if count >= maxLen {
				return builder.String(), true
			}
			builder.WriteRune(current)
			count++
		}
	}
	return builder.String(), false
}
