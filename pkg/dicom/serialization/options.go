// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package serialization provides options for DICOM JSON serialization and deserialization and xml serialization.
package serialization

// NumberSerializationMode defines how DICOM numbers (IS, DS, SV, UV) should be serialized
type NumberSerializationMode int

const (
	// AsNumber always serializes DICOM numbers as JSON numbers.
	// This will return an error when a number can't be parsed.
	// Example: "00081160":{"vr":"IS","Value":[76]}
	AsNumber NumberSerializationMode = iota

	// AsString always serializes DICOM numbers as JSON strings.
	// Example: "00081160":{"vr":"IS","Value":["76"]}
	AsString

	// PreferablyAsNumber tries to serialize as JSON numbers, but falls back to strings
	// if parsing fails. This won't error on unparsable values.
	// Example: "00081160":{"vr":"IS","Value":[76]} or {"vr":"IS","Value":["invalid"]}
	PreferablyAsNumber
)

// jsonConfig configures JSON serialization/deserialization behavior
type jsonConfig struct {
	// writeTagsAsKeywords writes DICOM keywords instead of tags as JSON keys.
	// This makes the JSON non-compliant to DICOM JSON standard.
	// Example: "PatientName" instead of "00100010"
	writeTagsAsKeywords bool

	// writeKeyword adds a "keyword" attribute to each element.
	// Note: This is non-standard and may break parsers!
	writeKeyword bool

	// writeName adds a "name" attribute to each element.
	// Note: This is non-standard and may break parsers!
	writeName bool

	// autoValidate enables validation when deserializing.
	// If true, elements will be validated according to DICOM rules.
	autoValidate bool

	// numberSerializationMode defines how numeric values (IS, DS, SV, UV) are serialized.
	numberSerializationMode NumberSerializationMode

	// indent specifies the indentation string for pretty-printing.
	// Empty string means no indentation (compact format).
	indent string
}

// defaultJSONConfig returns the default configuration for DICOM JSON serialization.
// These options produce standard-compliant DICOM JSON.
func defaultJSONConfig() jsonConfig {
	return jsonConfig{
		writeTagsAsKeywords:     false,
		writeKeyword:            false,
		writeName:               false,
		autoValidate:            true,
		numberSerializationMode: AsNumber,
		indent:                  "",
	}
}

// JSONOption is a functional option for configuring JSON serialization
type JSONOption func(*jsonConfig)

// WithWriteTagsAsKeywords configures whether to write tags as keywords (non-standard)
func WithWriteTagsAsKeywords(enabled bool) JSONOption {
	return func(c *jsonConfig) {
		c.writeTagsAsKeywords = enabled
	}
}

// WithWriteKeyword configures whether to write keyword attribute (non-standard)
func WithWriteKeyword(enabled bool) JSONOption {
	return func(c *jsonConfig) {
		c.writeKeyword = enabled
	}
}

// WithWriteName configures whether to write name attribute (non-standard)
func WithWriteName(enabled bool) JSONOption {
	return func(c *jsonConfig) {
		c.writeName = enabled
	}
}

// WithAutoValidate configures whether to validate during deserialization
func WithAutoValidate(enabled bool) JSONOption {
	return func(c *jsonConfig) {
		c.autoValidate = enabled
	}
}

// WithNumberSerializationMode configures how numbers are serialized
func WithNumberSerializationMode(mode NumberSerializationMode) JSONOption {
	return func(c *jsonConfig) {
		c.numberSerializationMode = mode
	}
}

// WithIndent configures the indentation for pretty-printing
func WithIndent(indent string) JSONOption {
	return func(c *jsonConfig) {
		c.indent = indent
	}
}

// xmlConfig configures XML serialization behavior
type xmlConfig struct {
	// indent specifies the indentation string for pretty-printing.
	// Empty string means no indentation (compact format).
	// Common values: "" (compact), "  " (2 spaces), "\t" (tab)
	indent string
}

// defaultXMLConfig returns the default configuration for DICOM XML serialization.
// These options produce standard-compliant DICOM NativeDicomModel XML.
func defaultXMLConfig() xmlConfig {
	return xmlConfig{
		indent: "  ", // Default to 2-space indentation for readability
	}
}

// XMLOption is a functional option for configuring XML serialization
type XMLOption func(*xmlConfig)

// WithXMLIndent configures the indentation for XML pretty-printing
func WithXMLIndent(indent string) XMLOption {
	return func(c *xmlConfig) {
		c.indent = indent
	}
}

// WithXMLCompact disables indentation for compact XML output
func WithXMLCompact(indent string) XMLOption {
	return func(c *xmlConfig) {
		c.indent = indent
	}
}
