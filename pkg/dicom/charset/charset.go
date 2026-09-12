// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package charset implements DICOM character set encoding support.
package charset

import (
	"bytes"
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

// Info holds information about a DICOM character set.
type Info struct {
	Name     string
	Encoding encoding.Encoding
}

// dicomISO2022Encoding retains the DICOM designation needed when encoding a
// value. The underlying x/text encoding only handles the character bytes.
type dicomISO2022Encoding struct {
	encoding.Encoding
	designation []byte
}

const iso2022ASCII = "ascii"

// knownCharsets maps DICOM Specific Character Set values to Go encodings.
var knownCharsets = map[string]*Info{
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
	"ISO 2022 IR 149": {"Korean Extended", &dicomISO2022Encoding{Encoding: korean.EUCKR, designation: []byte{0x1b, '$', ')', 'C'}}},
	"ISO 2022 IR 58":  {"Chinese Simplified (GB2312)", &dicomISO2022Encoding{Encoding: simplifiedchinese.GB18030, designation: []byte{0x1b, '$', ')', 'A'}}},
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

	if bytes.Contains(data, []byte{0x1b}) {
		return decodeDICOMISO2022(data, encodings)
	}

	enc := encodings[0]
	if enc == nil {
		enc = Default
	}
	decoded, err := enc.NewDecoder().Bytes(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode string: %w", err)
	}
	return string(decoded), nil
}

func decodeDICOMISO2022(data []byte, encodings []encoding.Encoding) (string, error) {
	current := iso2022ASCII
	var initial encoding.Encoding = Default
	if len(encodings) > 0 && encodings[0] != nil {
		initial = encodings[0]
	}
	if initial == Default && len(encodings) > 1 && encodings[1] != nil {
		initial = encodings[1]
	}
	if initial == japanese.ShiftJIS {
		current = "shiftjis"
	}
	var designation []byte
	var segment []byte
	var result bytes.Buffer
	flush := func() error {
		if len(segment) == 0 {
			return nil
		}
		var input []byte
		var enc encoding.Encoding
		switch current {
		case iso2022ASCII:
			enc = Default
			input = segment
		case "shiftjis":
			enc = japanese.ShiftJIS
			input = segment
		case "gb":
			enc = simplifiedchinese.GB18030
			input = segment
		case "korean":
			enc = korean.EUCKR
			input = segment
		case "japanese":
			enc = japanese.ISO2022JP
			input = append(append([]byte(nil), designation...), segment...)
			input = append(input, 0x1b, '(', 'B')
		default:
			return fmt.Errorf("unsupported DICOM ISO-2022 state %q", current)
		}
		decoded, err := enc.NewDecoder().Bytes(input)
		if err != nil {
			return fmt.Errorf("failed to decode ISO-2022 %s segment: %w", current, err)
		}
		result.Write(decoded)
		segment = segment[:0]
		return nil
	}

	for index := 0; index < len(data); {
		if data[index] != 0x1b {
			segment = append(segment, data[index])
			index++
			continue
		}
		if err := flush(); err != nil {
			return "", err
		}
		kind, seq, consumed, ok := parseISO2022Escape(data[index:])
		if !ok {
			return "", fmt.Errorf("unsupported or truncated DICOM ISO-2022 escape at byte %d", index)
		}
		current = kind
		designation = seq
		index += consumed
	}
	if err := flush(); err != nil {
		return "", err
	}
	return result.String(), nil
}

func parseISO2022Escape(data []byte) (kind string, sequence []byte, consumed int, ok bool) {
	if len(data) >= 3 && data[0] == 0x1b && data[1] == '(' {
		switch data[2] {
		case 'B', 'J', 'H':
			return iso2022ASCII, nil, 3, true
		}
	}
	if len(data) >= 3 && data[0] == 0x1b && data[1] == '$' {
		if data[2] == 'B' {
			return "japanese", append([]byte(nil), data[:3]...), 3, true
		}
		if len(data) >= 4 && data[2] == ')' {
			switch data[3] {
			case 'A':
				return "gb", append([]byte(nil), data[:4]...), 4, true
			case 'C':
				return "korean", append([]byte(nil), data[:4]...), 4, true
			}
		}
	}
	return "", nil, 0, false
}

// EncodeString encodes a string using the specified encodings.
// It will try each encoding in order until one succeeds.
func EncodeString(s string, encodings []encoding.Encoding) ([]byte, error) {
	if len(encodings) == 0 {
		encodings = []encoding.Encoding{Default}
	}

	// Try the first encoding
	encoded, err := encodeWithEncoding(s, encodings[0])
	if err == nil {
		return encoded, nil
	}

	// If first encoding failed, try others
	for _, enc := range encodings[1:] {
		encoded, err := encodeWithEncoding(s, enc)
		if err == nil {
			return encoded, nil
		}
	}

	return nil, fmt.Errorf("could not encode string with given encodings")
}

func encodeWithEncoding(s string, enc encoding.Encoding) ([]byte, error) {
	if enc == nil {
		enc = Default
	}
	encoded, err := enc.NewEncoder().String(s)
	if err != nil {
		return nil, err
	}
	if dicom, ok := enc.(*dicomISO2022Encoding); ok {
		result := append([]byte(nil), dicom.designation...)
		result = append(result, []byte(encoded)...)
		result = append(result, 0x1b, '(', 'B')
		return result, nil
	}
	return []byte(encoded), nil
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
func GetCharsetInfo(charset string) (*Info, bool) {
	charset = strings.TrimSpace(charset)
	info, ok := knownCharsets[charset]
	return info, ok
}
