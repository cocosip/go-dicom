// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

// Package charset implements DICOM character set encoding support.
package charset

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/cocosip/go-dicom/pkg/logging"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)

// asciiEncoding marks the DICOM default repertoire so the encode/decode
// helpers can enforce its 7-bit restriction; the embedded encoding.Encoding
// supplies decoder and encoder construction unchanged.
type asciiEncoding struct{ encoding.Encoding }

// Default is the DICOM default repertoire (ISO-IR 6 / 7-bit ASCII).
// The underlying x/text map is used for ASCII-compatible decoding, while the
// public encode/decode helpers enforce the 7-bit restriction.
var Default encoding.Encoding = &asciiEncoding{Encoding: charmap.ISO8859_1}

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
	// validate rejects byte sequences the DICOM character set does not define,
	// so the encoder falls back instead of writing characters outside the
	// repertoire.
	validate func([]byte) error
}

// dicomIR13Encoding wraps JIS X 0201. designation is set only for the ISO 2022
// extension, which must emit ESC 02/09 04/09 (G1 katakana) before the value
// (PS3.5 6.1.2.5.3); the plain ISO_IR 13 charset carries no escape sequence.
type dicomIR13Encoding struct {
	encoding.Encoding
	designation []byte
}

// validatedSingleByteEncoding applies the repertoire restriction of an 8-bit
// code table that has no ISO 2022 designation.
type validatedSingleByteEncoding struct {
	encoding.Encoding
	validate func([]byte) error
}

func (e *dicomISO2022Encoding) validateEncoded(data []byte) error {
	if e.validate == nil {
		return nil
	}
	return e.validate(data)
}

func (e *dicomIR13Encoding) validateEncoded(data []byte) error {
	if !validIR13Bytes(data) {
		return fmt.Errorf("ISO_IR 13 cannot encode characters outside JIS X 0201")
	}
	return nil
}

func (e *validatedSingleByteEncoding) validateEncoded(data []byte) error {
	return e.validate(data)
}

// validTIS620Bytes rejects the positions that TIS 620-2533 leaves undefined:
// charmap.Windows874 assigns graphic characters in 0x80-0x9F, which TIS 620
// does not (PS3.3 C.12.1.1.3 maps ISO_IR 166 to TIS 620-2533).
func validTIS620Bytes(data []byte) error {
	for _, value := range data {
		if value >= 0x80 && value <= 0x9f {
			return fmt.Errorf("ISO_IR 166 (TIS 620-2533) has no character at byte %#02x", value)
		}
	}
	return nil
}

// validGB2312Bytes rejects the GBK positions that GB 2312 does not define
// (PS3.3 C.12.1.1.5 maps ISO 2022 IR 58 to GB 2312); GBK is used as the
// underlying codec because it is code point backward compatible.
func validGB2312Bytes(data []byte) error {
	for index := 0; index < len(data); index++ {
		lead := data[index]
		if lead < 0x80 {
			continue
		}
		if lead < 0xa1 || lead > 0xf7 {
			return fmt.Errorf("ISO 2022 IR 58 (GB 2312) has no character at byte %#02x", lead)
		}
		index++
		if index >= len(data) || data[index] < 0xa1 || data[index] > 0xfe {
			return fmt.Errorf("ISO 2022 IR 58 (GB 2312) has no character with lead byte %#02x", lead)
		}
	}
	return nil
}

const iso2022ASCII = "ascii"
const iso2022IR13 = "ir13"
const iso2022Japanese = "japanese"

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
	"ISO_IR 166": {"Thai (TIS 620-2533)", &validatedSingleByteEncoding{Encoding: charmap.Windows874, validate: validTIS620Bytes}},

	// Japanese
	"ISO_IR 13": {"Japanese (JIS X 0201)", &dicomIR13Encoding{Encoding: japanese.ShiftJIS}},

	// Korean
	"ISO_IR 149": {"Korean (EUC-KR)", korean.EUCKR},

	// Chinese
	"GBK":     {"Chinese Simplified (GBK)", simplifiedchinese.GBK},
	"GB18030": {"Chinese Simplified (GB18030)", simplifiedchinese.GB18030},

	// Unicode
	"ISO_IR 192": {"Unicode (UTF-8)", unicode.UTF8},

	// ISO 2022 Extended character sets (simplified handling)
	"ISO 2022 IR 6":   {"ASCII", Default},
	"ISO 2022 IR 100": {"Latin-1 Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_1, designation: []byte{0x1b, '-', 'A'}}},
	"ISO 2022 IR 101": {"Latin-2 Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_2, designation: []byte{0x1b, '-', 'B'}}},
	"ISO 2022 IR 109": {"Latin-3 Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_3, designation: []byte{0x1b, '-', 'C'}}},
	"ISO 2022 IR 110": {"Latin-4 Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_4, designation: []byte{0x1b, '-', 'D'}}},
	"ISO 2022 IR 144": {"Cyrillic Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_5, designation: []byte{0x1b, '-', 'L'}}},
	"ISO 2022 IR 127": {"Arabic Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_6, designation: []byte{0x1b, '-', 'G'}}},
	"ISO 2022 IR 126": {"Greek Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_7, designation: []byte{0x1b, '-', 'F'}}},
	"ISO 2022 IR 138": {"Hebrew Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_8, designation: []byte{0x1b, '-', 'H'}}},
	"ISO 2022 IR 148": {"Latin-5 Extended", &dicomISO2022Encoding{Encoding: charmap.ISO8859_9, designation: []byte{0x1b, '-', 'M'}}},
	"ISO 2022 IR 13":  {"Japanese Extended", &dicomIR13Encoding{Encoding: japanese.ShiftJIS, designation: []byte{0x1b, ')', 'I'}}},
	"ISO 2022 IR 87":  {"Japanese (ISO-2022-JP)", japanese.ISO2022JP},
	"ISO 2022 IR 149": {"Korean Extended", &dicomISO2022Encoding{Encoding: korean.EUCKR, designation: []byte{0x1b, '$', ')', 'C'}}},
	"ISO 2022 IR 58":  {"Chinese Simplified (GB 2312)", &dicomISO2022Encoding{Encoding: simplifiedchinese.GBK, designation: []byte{0x1b, '$', ')', 'A'}, validate: validGB2312Bytes}},
	"ISO 2022 IR 159": {"Japanese (JIS X 0212)", unsupportedCharset{name: "ISO 2022 IR 159"}},
	"ISO 2022 IR 166": {"Thai Extended (TIS 620-2533)", &dicomISO2022Encoding{Encoding: charmap.Windows874, designation: []byte{0x1b, '-', 'T'}, validate: validTIS620Bytes}},
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
//
// Values using the DICOM default repertoire (no Specific Character Set declared)
// that contain bytes above 0x7F are decoded leniently as ISO-8859-1 with a
// warning instead of failing. Use DecodeStringStrict to reject such bytes.
func DecodeString(data []byte, encodings []encoding.Encoding) (string, error) {
	return decodeString(data, encodings, false)
}

// DecodeStringStrict is like DecodeString but enforces the 7-bit DICOM default
// repertoire: bytes above 0x7F with no declared Specific Character Set are
// reported as an error instead of being decoded as ISO-8859-1.
func DecodeStringStrict(data []byte, encodings []encoding.Encoding) (string, error) {
	return decodeString(data, encodings, true)
}

func decodeString(data []byte, encodings []encoding.Encoding, strict bool) (string, error) {
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
	if _, isDefault := enc.(*asciiEncoding); isDefault {
		if !isASCII(data) {
			if strict {
				return "", fmt.Errorf("DICOM default character set only permits 7-bit ASCII")
			}
			logging.Emit(context.Background(), logging.Record{
				Level:     slog.LevelWarn,
				Component: "dicom.charset",
				Event:     "default_charset_extended_bytes",
				Message:   "DICOM default character set contains bytes above 0x7F; decoding leniently as ISO-8859-1",
				Attrs:     []slog.Attr{slog.Int("byte_count", len(data))},
			})
			decoded, err := charmap.ISO8859_1.NewDecoder().Bytes(data)
			if err != nil {
				return "", fmt.Errorf("failed to decode string as ISO-8859-1: %w", err)
			}
			return string(decoded), nil
		}
		return string(data), nil
	}
	if ir13, isIR13 := enc.(*dicomIR13Encoding); isIR13 && len(ir13.designation) == 0 && !validIR13Bytes(data) {
		// PS3.3 C.12.1.1.1 defines ISO_IR 13 as the single-byte JIS X 0201 code
		// table; the Shift_JIS multi-byte bytes accepted here are not part of
		// it. Decode them leniently, but report the deviation.
		logging.Emit(context.Background(), logging.Record{
			Level:     slog.LevelWarn,
			Component: "dicom.charset",
			Event:     "ir13_multibyte_bytes",
			Message:   "ISO_IR 13 value contains bytes outside JIS X 0201; decoding leniently as Shift_JIS",
			Attrs:     []slog.Attr{slog.Int("byte_count", len(data))},
		})
	}
	decoded, err := enc.NewDecoder().Bytes(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode string: %w", err)
	}
	return string(decoded), nil
}

// initialISO2022State returns the ISO 2022 state that is active at the start of
// a value: Value 1 of Specific Character Set, or the default repertoire when
// Value 1 is missing (PS3.5 6.1.2.5.3, 6.1.2.5.4). Value 1 is not limited to
// ISO 2022 IR 13; any single-byte set can be Value 1 when code extensions are
// used, and its bytes may appear before the first escape sequence.
func initialISO2022State(encodings []encoding.Encoding) string {
	if len(encodings) == 0 || encodings[0] == nil {
		return iso2022ASCII
	}
	switch enc := encodings[0].(type) {
	case *dicomIR13Encoding:
		return iso2022IR13
	case *dicomISO2022Encoding:
		if len(enc.designation) == 3 && enc.designation[0] == 0x1b &&
			(enc.designation[1] == '-' || enc.designation[1] == '.') {
			if _, ok := iso2022SingleByte[enc.designation[2]]; ok {
				return "sb" + string(enc.designation[2])
			}
		}
	default:
		for letter, single := range iso2022SingleByte {
			if enc == single {
				return "sb" + string(letter)
			}
		}
	}
	return iso2022ASCII
}

func decodeDICOMISO2022(data []byte, encodings []encoding.Encoding) (string, error) {
	initial := initialISO2022State(encodings)
	current := initial
	var designation []byte
	var segment []byte
	var result bytes.Buffer
	reset := func() {
		current = initial
		designation = nil
	}
	flush := func() error {
		if len(segment) == 0 {
			return nil
		}
		var input []byte
		var enc encoding.Encoding
		switch current {
		case iso2022ASCII:
			if !isASCII(segment) {
				return fmt.Errorf("DICOM default character set only permits 7-bit ASCII")
			}
			enc = Default
			input = segment
		case "shiftjis", iso2022IR13:
			enc = japanese.ShiftJIS
			if current == "ir13" && !validIR13Bytes(segment) {
				return fmt.Errorf("ISO_IR 13 contains bytes outside JIS X 0201")
			}
			input = segment
		case "gb":
			enc = simplifiedchinese.GB18030
			input = segment
		case "korean":
			enc = korean.EUCKR
			input = segment
		case iso2022Japanese:
			enc = japanese.ISO2022JP
			input = append(append([]byte(nil), designation...), segment...)
			input = append(input, 0x1b, '(', 'B')
		case kindIR13Katakana:
			enc = japanese.ShiftJIS
			input = segment
		default:
			if len(current) == 3 && current[:2] == "sb" {
				singleEnc, exists := iso2022SingleByte[current[2]]
				if !exists {
					return fmt.Errorf("unsupported DICOM ISO-2022 state %q", current)
				}
				enc = singleEnc
				input = segment
				break
			}
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
		if data[index] == '\\' && current == iso2022ASCII {
			if err := flush(); err != nil {
				return "", err
			}
			result.WriteByte(data[index])
			// The next value starts with Value 1 of Specific Character Set
			// active again (PS3.5 6.1.2.5.3).
			reset()
			index++
			continue
		}
		if data[index] == '\r' || data[index] == '\n' || data[index] == '\f' {
			if err := flush(); err != nil {
				return "", err
			}
			result.WriteByte(data[index])
			// Value 1 (or the default repertoire) is active at the beginning of
			// each line and page (PS3.5 6.1.2.5.3).
			reset()
			index++
			continue
		}
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
		if !iso2022EscapeDeclared(kind, seq, encodings) {
			return "", fmt.Errorf("DICOM ISO-2022 escape %x is not declared", seq)
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

func iso2022EscapeDeclared(kind string, sequence []byte, encodings []encoding.Encoding) bool {
	if kind == iso2022ASCII {
		return true
	}
	if len(kind) == 3 && kind[:2] == "sb" {
		singleEnc, exists := iso2022SingleByte[kind[2]]
		if !exists {
			return false
		}
		for _, enc := range encodings {
			if enc == singleEnc {
				return true
			}
			// An ISO 2022 extension wraps the character-map encoding and carries
			// the escape designation that produces the bytes being decoded.
			if dicom, ok := enc.(*dicomISO2022Encoding); ok &&
				dicom.Encoding == singleEnc && bytes.Equal(dicom.designation, sequence) {
				return true
			}
		}
		return false
	}
	if kind == kindIR13Katakana {
		for _, enc := range encodings {
			if enc == japanese.ShiftJIS {
				return true
			}
			if dicom, ok := enc.(*dicomIR13Encoding); ok && dicom.Encoding == japanese.ShiftJIS {
				return true
			}
		}
		return false
	}
	for _, enc := range encodings {
		if enc == japanese.ISO2022JP && kind == iso2022Japanese {
			return true
		}
		if dicom, ok := enc.(*dicomISO2022Encoding); ok && bytes.Equal(dicom.designation, sequence) {
			return true
		}
	}
	return false
}

// iso2022SingleByte holds the G1/G2 single-byte designations used by ISO 2022
// escape sequences of the form ESC - F / ESC . F (PS3.5 6.1.2.5.3 and
// Annex A). The designations are fixed by the standard; e.g. ESC - 47 is
// ISO-IR 127 (Arabic), not ISO-IR 138 (Hebrew).
var iso2022SingleByte = map[byte]encoding.Encoding{
	'A': charmap.ISO8859_1,
	'B': charmap.ISO8859_2,
	'C': charmap.ISO8859_3,
	'D': charmap.ISO8859_4,
	'L': charmap.ISO8859_5,
	'G': charmap.ISO8859_6,
	'F': charmap.ISO8859_7,
	'H': charmap.ISO8859_8,
	'M': charmap.ISO8859_9,
	'T': charmap.Windows874,
}

const kindIR13Katakana = "ir13katakana"

func parseISO2022Escape(data []byte) (kind string, sequence []byte, consumed int, ok bool) {
	if len(data) >= 3 && data[0] == 0x1b && data[1] == '(' {
		switch data[2] {
		case 'B', 'J', 'H':
			return iso2022ASCII, nil, 3, true
		case 'I':
			// Tolerated spelling of the IR 13 katakana designation.
			return kindIR13Katakana, append([]byte(nil), data[:3]...), 3, true
		}
	}
	if len(data) >= 3 && data[0] == 0x1b && data[1] == ')' && data[2] == 'I' {
		// ESC ) I: JIS X 0201 katakana (half-width) in G1, the designation
		// DICOM requires for ISO 2022 IR 13 (PS3.5 6.1.2.5.3).
		return kindIR13Katakana, append([]byte(nil), data[:3]...), 3, true
	}
	if len(data) >= 3 && data[0] == 0x1b && (data[1] == '-' || data[1] == '.') {
		if _, exists := iso2022SingleByte[data[2]]; exists {
			return "sb" + string(data[2]), append([]byte(nil), data[:3]...), 3, true
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
//
// The string is treated as a multi-valued text attribute: the value separator
// (5CH) is encoded as a boundary between values, per PS3.5 6.1.2.5.3.
func EncodeString(s string, encodings []encoding.Encoding) ([]byte, error) {
	return EncodeValueSet(s, "", encodings)
}

// EncodeValueSet encodes the Value Field of one textual Data Element.
//
// PS3.5 6.1.2.5.3 requires the character set of Specific Character Set Value 1
// (or the DICOM default repertoire) to be active at the beginning of every
// value and name component, to be restored before the value separator (5CH)
// and, for PN, before the name component and component group delimiters (5EH,
// 3DH), and requires an explicit designation before the first use of an
// extension character set in each of those. Every such segment is therefore
// encoded on its own, with the first encoding that can represent it.
//
// vrCode is the two-letter VR code of the Data Element. ST, LT, UT and UR are
// single-valued and may contain 05/12 as text, so they are never split.
func EncodeValueSet(s, vrCode string, encodings []encoding.Encoding) ([]byte, error) {
	if len(encodings) == 0 {
		encodings = []encoding.Encoding{Default}
	}
	delimiters := textValueDelimiters(vrCode)
	if delimiters == "" {
		return encodeSegment(s, encodings)
	}

	var result []byte
	rest := s
	for {
		index := strings.IndexAny(rest, delimiters)
		segment := rest
		if index >= 0 {
			segment = rest[:index]
		}
		encoded, err := encodeSegment(segment, encodings)
		if err != nil {
			return nil, err
		}
		result = append(result, encoded...)
		if index < 0 {
			return result, nil
		}
		result = append(result, rest[index])
		rest = rest[index+1:]
	}
}

// textValueDelimiters returns the characters that separate the segments of a
// text attribute which must each start in the Value 1 character set.
func textValueDelimiters(vrCode string) string {
	switch vrCode {
	case vrCodePersonName:
		return "\\^="
	case "ST", "LT", "UT", "UR":
		return ""
	default:
		return "\\"
	}
}

const vrCodePersonName = "PN"

// encodeSegment encodes one value, name component or component group.
func encodeSegment(segment string, encodings []encoding.Encoding) ([]byte, error) {
	var lastErr error
	for _, enc := range encodings {
		encoded, err := encodeWithEncoding(segment, enc)
		if err == nil {
			return encoded, nil
		}
		lastErr = err
	}
	return nil, fmt.Errorf("could not encode string with given encodings: %w", lastErr)
}

func encodeWithEncoding(s string, enc encoding.Encoding) ([]byte, error) {
	if enc == nil {
		enc = Default
	}
	if _, strict := enc.(*asciiEncoding); strict {
		if !isASCII([]byte(s)) {
			return nil, fmt.Errorf("DICOM default character set only permits 7-bit ASCII")
		}
		return []byte(s), nil
	}
	encoded, err := enc.NewEncoder().String(s)
	if err != nil {
		return nil, err
	}
	if checker, ok := enc.(interface{ validateEncoded([]byte) error }); ok {
		if err := checker.validateEncoded([]byte(encoded)); err != nil {
			return nil, err
		}
	}
	designation, reset := designationAndReset(enc)
	if len(designation) == 0 {
		return []byte(encoded), nil
	}
	result := make([]byte, 0, len(designation)+len(encoded)+len(reset))
	result = append(result, designation...)
	result = append(result, encoded...)
	return append(result, reset...), nil
}

// designationAndReset returns the ISO 2022 escape sequence that invokes the
// character set and the sequence that restores the default repertoire.
func designationAndReset(enc encoding.Encoding) ([]byte, []byte) {
	switch typed := enc.(type) {
	case *dicomIR13Encoding:
		if len(typed.designation) == 0 {
			return nil, nil
		}
		return typed.designation, []byte{0x1b, '(', 'B'}
	case *dicomISO2022Encoding:
		return typed.designation, []byte{0x1b, '(', 'B'}
	default:
		return nil, nil
	}
}

func validIR13Bytes(data []byte) bool {
	for _, value := range data {
		if (value >= 0x80 && value <= 0xa0) || value >= 0xe0 {
			return false
		}
	}
	return true
}

func isASCII(data []byte) bool {
	for _, value := range data {
		if value > 0x7f {
			return false
		}
	}
	return true
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
	if !ok {
		return nil, false
	}
	infoCopy := *info
	return &infoCopy, true
}

// unsupportedCharset is a placeholder encoding for character sets that are
// recognized by name but not backed by an available encoder (e.g. ISO 2022
// IR 159 / JIS X 0212). Using it produces an explicit error instead of silently
// decoding with the default repertoire.
type unsupportedCharset struct{ name string }

func (u unsupportedCharset) NewDecoder() *encoding.Decoder {
	return &encoding.Decoder{Transformer: unsupportedTransformer(u.name)}
}

func (u unsupportedCharset) NewEncoder() *encoding.Encoder {
	return &encoding.Encoder{Transformer: unsupportedTransformer(u.name)}
}

type unsupportedTransformer string

func (u unsupportedTransformer) Transform(_, _ []byte, _ bool) (int, int, error) {
	return 0, 0, fmt.Errorf("DICOM character set %s is not supported", string(u))
}

func (unsupportedTransformer) Reset() {}
