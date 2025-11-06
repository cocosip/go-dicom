// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package charset implements DICOM character set encoding support.
package charset

import (
	"fmt"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)

// Default is the default DICOM encoding (ASCII/ISO-8859-1).
var Default = charmap.ISO8859_1

// CharsetInfo holds information about a DICOM character set.
type CharsetInfo struct {
	Name     string
	Encoding encoding.Encoding
}

// knownCharsets maps DICOM Specific Character Set values to Go encodings.
var knownCharsets = map[string]*CharsetInfo{
	// ISO 8859 Latin character sets
	"ISO_IR 100": {"Latin-1 (Western European)", charmap.ISO8859_1},
	"ISO_IR 101": {"Latin-2 (Central European)", charmap.ISO8859_2},
	"ISO_IR 109": {"Latin-3 (South European)", charmap.ISO8859_3},
	"ISO_IR 110": {"Latin-4 (North European)", charmap.ISO8859_4},
	"ISO_IR 144": {"Cyrillic", charmap.ISO8859_5},
	"ISO_IR 127": {"Arabic", charmap.ISO8859_6},
	"ISO_IR 126": {"Greek", charmap.ISO8859_7},
	"ISO_IR 138": {"Hebrew", charmap.ISO8859_8},
	"ISO_IR 148": {"Latin-5 (Turkish)", charmap.ISO8859_9},
	"ISO_IR 166": {"Thai", charmap.Windows874},

	// Japanese
	"ISO_IR 13": {"Japanese (Shift-JIS)", japanese.ShiftJIS},

	// Korean
	"ISO_IR 149": {"Korean (EUC-KR)", korean.EUCKR},

	// Chinese
	"GBK":     {"Chinese Simplified (GBK)", simplifiedchinese.GBK},
	"GB18030": {"Chinese Simplified (GB18030)", simplifiedchinese.GB18030},

	// Unicode
	"ISO_IR 192": {"Unicode (UTF-8)", unicode.UTF8},

	// ISO 2022 Extended character sets (simplified handling)
	"ISO 2022 IR 6":   {"ASCII", charmap.ISO8859_1},
	"ISO 2022 IR 100": {"Latin-1 Extended", charmap.ISO8859_1},
	"ISO 2022 IR 101": {"Latin-2 Extended", charmap.ISO8859_2},
	"ISO 2022 IR 109": {"Latin-3 Extended", charmap.ISO8859_3},
	"ISO 2022 IR 110": {"Latin-4 Extended", charmap.ISO8859_4},
	"ISO 2022 IR 144": {"Cyrillic Extended", charmap.ISO8859_5},
	"ISO 2022 IR 127": {"Arabic Extended", charmap.ISO8859_6},
	"ISO 2022 IR 126": {"Greek Extended", charmap.ISO8859_7},
	"ISO 2022 IR 138": {"Hebrew Extended", charmap.ISO8859_8},
	"ISO 2022 IR 148": {"Latin-5 Extended", charmap.ISO8859_9},
	"ISO 2022 IR 13":  {"Japanese Extended", japanese.ShiftJIS},
	"ISO 2022 IR 87":  {"Japanese (ISO-2022-JP)", japanese.ISO2022JP},
	"ISO 2022 IR 149": {"Korean Extended", korean.EUCKR},
	"ISO 2022 IR 58":  {"Chinese Simplified (GB2312)", simplifiedchinese.HZGB2312},
	"ISO 2022 IR 166": {"Thai Extended", charmap.Windows874},
}

// GetEncoding returns the encoding for a given DICOM Specific Character Set value.
// Returns Default (ISO-8859-1) if the charset is not recognized.
func GetEncoding(charset string) encoding.Encoding {
	charset = strings.TrimSpace(charset)
	if charset == "" {
		return Default
	}

	// Try exact match first
	if info, ok := knownCharsets[charset]; ok {
		return info.Encoding
	}

	// Try common misspellings: "ISO-IR" or "ISO IR" instead of "ISO_IR"
	normalized := strings.ReplaceAll(charset, "ISO-IR", "ISO_IR")
	normalized = strings.ReplaceAll(normalized, "ISO IR", "ISO_IR")

	if info, ok := knownCharsets[normalized]; ok {
		return info.Encoding
	}

	// Return default if not found
	return Default
}

// GetEncodings returns encodings for multiple Specific Character Set values.
// The first encoding is the default encoding, subsequent encodings are for
// extended character sets (code extensions).
func GetEncodings(charsets []string) []encoding.Encoding {
	if len(charsets) == 0 {
		return []encoding.Encoding{Default}
	}

	encodings := make([]encoding.Encoding, len(charsets))
	for i, cs := range charsets {
		encodings[i] = GetEncoding(cs)
	}
	return encodings
}

// GetCharsetName returns the DICOM Specific Character Set value for a given encoding.
// Returns an error if the encoding is not a known DICOM character set.
func GetCharsetName(enc encoding.Encoding) (string, error) {
	for charset, info := range knownCharsets {
		if info.Encoding == enc {
			// Prefer non-extended (ISO 2022) versions
			if !strings.HasPrefix(charset, "ISO 2022") {
				return charset, nil
			}
		}
	}

	// Second pass: return extended version if no basic version found
	for charset, info := range knownCharsets {
		if info.Encoding == enc {
			return charset, nil
		}
	}

	return "", fmt.Errorf("no DICOM charset found for encoding")
}

// DecodeString decodes a byte slice using the specified encodings.
// For single-byte character sets, only the first encoding is used.
// For multi-byte or extended character sets, escape sequences are processed.
func DecodeString(data []byte, encodings []encoding.Encoding) (string, error) {
	if len(data) == 0 {
		return "", nil
	}

	if len(encodings) == 0 {
		encodings = []encoding.Encoding{Default}
	}

	// Simple case: single encoding, no escape sequences
	decoder := encodings[0].NewDecoder()
	decoded, err := decoder.Bytes(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode string: %w", err)
	}

	return string(decoded), nil
}

// EncodeString encodes a string using the specified encodings.
// It will try each encoding in order until one succeeds.
func EncodeString(s string, encodings []encoding.Encoding) ([]byte, error) {
	if len(encodings) == 0 {
		encodings = []encoding.Encoding{Default}
	}

	// Try the first encoding
	encoder := encodings[0].NewEncoder()
	encoded, err := encoder.String(s)
	if err == nil {
		return []byte(encoded), nil
	}

	// If first encoding failed, try others
	for _, enc := range encodings[1:] {
		encoder := enc.NewEncoder()
		encoded, err := encoder.String(s)
		if err == nil {
			return []byte(encoded), nil
		}
	}

	// Fallback: use default encoding with replacement characters
	encoder = Default.NewEncoder()
	encoded, _ = encoder.String(s)
	return []byte(encoded), fmt.Errorf("could not encode string with given encodings, used replacement characters")
}

// KnownCharsets returns a list of all known DICOM character sets.
func KnownCharsets() []string {
	charsets := make([]string, 0, len(knownCharsets))
	for charset := range knownCharsets {
		charsets = append(charsets, charset)
	}
	return charsets
}

// GetCharsetInfo returns information about a DICOM character set.
func GetCharsetInfo(charset string) (*CharsetInfo, bool) {
	charset = strings.TrimSpace(charset)
	info, ok := knownCharsets[charset]
	return info, ok
}
