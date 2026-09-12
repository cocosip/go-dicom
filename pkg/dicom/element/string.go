// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// Compile-time check to ensure String implements Element interface
var _ Element = (*String)(nil)

// String represents a DICOM string element.
//
// String elements can contain single or multiple values separated by backslashes.
// They handle character set encoding and provide methods for accessing values as strings.
//
// Applicable VRs: AE, AS, CS, DA, DS, DT, IS, LO, LT, PN, SH, ST, TM, UC, UI, UR, UT
type String struct {
	*base
	encoding  encoding.Encoding
	encodings []encoding.Encoding
	encodeErr error
}

// NewString creates a new string element with the given tag, VR, and values.
// Values are joined with backslash separators and encoded into a buffer.
func NewString(t *tag.Tag, v *vr.VR, values []string) *String {
	result := NewStringWithEncoding(t, v, values, charset.Default)
	if result.encodeErr != nil {
		// This legacy constructor cannot return an error. Preserve the Go text
		// losslessly as UTF-8 instead of exposing UTF-8 bytes as Latin-1.
		return NewStringWithEncoding(t, v, values, unicode.UTF8)
	}
	return result
}

// NewStringWithEncoding creates a new string element with a specific encoding.
func NewStringWithEncoding(t *tag.Tag, v *vr.VR, values []string, enc encoding.Encoding) *String {
	if enc == nil {
		enc = charset.Default
	}

	// Join values with backslash separator
	joined := strings.Join(values, "\\")

	// Encode to bytes
	var data []byte
	var encodeErr error
	if enc != nil {
		encoded, err := charset.EncodeString(joined, []encoding.Encoding{enc})
		if err != nil {
			encodeErr = fmt.Errorf("encode string using %T: %w", enc, err)
			data = nil
		} else {
			data = encoded
		}
	} else {
		data = []byte(joined)
	}

	// Remove trailing spaces (DICOM requirement)
	data = bytes.TrimRight(data, " ")

	// Add padding if required by VR
	// UI (Unique Identifier) must be even length, pad with null byte if odd
	if v == vr.UI && len(data)%2 == 1 {
		data = append(data, 0x00)
	}

	// Create buffer
	buf := buffer.NewMemory(data)

	return &String{
		base:      newBase(t, v, buf),
		encoding:  enc,
		encodings: []encoding.Encoding{enc},
		encodeErr: encodeErr,
	}
}

// NewStringFromBuffer creates a string element from an existing buffer.
// This is used when parsing DICOM data.
func NewStringFromBuffer(t *tag.Tag, v *vr.VR, buf buffer.ByteBuffer, enc encoding.Encoding) *String {
	if enc == nil {
		enc = charset.Default
	}
	return NewStringFromBufferWithEncodings(t, v, buf, []encoding.Encoding{enc})
}

// NewStringFromBufferWithEncodings creates a string element from an existing buffer.
// Multiple encodings are used for DICOM SpecificCharacterSet code extensions.
func NewStringFromBufferWithEncodings(t *tag.Tag, v *vr.VR, buf buffer.ByteBuffer, encodings []encoding.Encoding) *String {
	if len(encodings) == 0 {
		encodings = []encoding.Encoding{charset.Default}
	}
	copied := make([]encoding.Encoding, len(encodings))
	copy(copied, encodings)
	if copied[0] == nil {
		copied[0] = charset.Default
	}

	return &String{
		base:      newBase(t, v, buf),
		encoding:  copied[0],
		encodings: copied,
	}
}

// Count returns the number of values in this element.
// Values are separated by backslashes.
func (s *String) Count() int {
	str := s.GetString()
	if str == "" {
		return 0
	}
	if s.hasLiteralBackslash() {
		return 1
	}
	return strings.Count(str, "\\") + 1
}

// GetString returns the complete string value (all values joined by backslash).
func (s *String) GetString() string {
	decoded, err := s.GetStringWithError()
	if err != nil {
		return ""
	}
	return decoded
}

// GetStringWithError returns the decoded string and reports encoding errors
// instead of exposing undecoded bytes as a different character set.
func (s *String) GetStringWithError() (string, error) {
	if s == nil {
		return "", fmt.Errorf("string element is nil")
	}
	if s.encodeErr != nil {
		return "", s.encodeErr
	}
	if s.buffer == nil || s.buffer.Size() == 0 {
		return "", nil
	}

	data := s.buffer.Data()

	decoded, err := charset.DecodeString(data, s.encodings)
	if err != nil {
		return "", err
	}
	strData := []byte(decoded)

	// Remove trailing null bytes and spaces (DICOM padding)
	// UI VR uses null byte (0x00) for padding, other VRs may use space
	strData = bytes.TrimRight(strData, "\x00 ")

	return string(strData), nil
}

// GetValue returns the value at the specified index.
// If index is -1 or out of range, returns the complete string.
func (s *String) GetValue(index int) string {
	str := s.GetString()
	if str == "" {
		return ""
	}
	if s.hasLiteralBackslash() {
		return str
	}

	values := strings.Split(str, "\\")
	if index < 0 || index >= len(values) {
		return str
	}

	return strings.TrimSpace(values[index])
}

// GetValues returns all values as a slice of strings.
func (s *String) GetValues() []string {
	str := s.GetString()
	if str == "" {
		return nil
	}
	if s.hasLiteralBackslash() {
		return []string{str}
	}

	values := strings.Split(str, "\\")
	// Trim spaces from each value
	for i, v := range values {
		values[i] = strings.TrimSpace(v)
	}
	return values
}

func (s *String) hasLiteralBackslash() bool {
	return s.vr == vr.LT || s.vr == vr.ST || s.vr == vr.UT || s.vr == vr.UR
}

// Encoding returns the character set encoding used by this element.
func (s *String) Encoding() encoding.Encoding {
	return s.encoding
}

// String returns a string representation of the element.
func (s *String) String() string {
	value := s.GetString()
	if len(value) > 64 {
		value = value[:64] + "..."
	}
	return fmt.Sprintf("Tag: %s, VR: %s, Value: %q", s.tag.String(), s.vr.Code(), value)
}

// Validate performs DICOM validation on the string element.
func (s *String) Validate() error {
	// Skip validation if disabled
	if !vr.PerformValidation {
		return nil
	}
	return s.validateValue()
}

func (s *String) validateValue() error {
	// Check the maximum length of each encoded value, excluding separators and
	// the Value Field's trailing padding.
	maxLen := s.vr.MaximumLength()
	if maxLen > 0 {
		data := []byte(nil)
		if s.buffer != nil {
			data = bytes.TrimRight(s.buffer.Data(), string(s.vr.PaddingValue()))
		}
		values := bytes.Split(data, []byte{'\\'})
		if s.hasLiteralBackslash() {
			values = [][]byte{data}
		}
		for i, value := range values {
			if uint32(len(value)) > maxLen {
				return fmt.Errorf("value[%d] length %d exceeds maximum %d for VR %s", i, len(value), maxLen, s.vr.Code())
			}
		}
	}

	// Perform VR-specific validation on each value
	values := s.GetValues()
	for i, value := range values {
		if err := s.vr.ValidateStringValue(value); err != nil {
			return fmt.Errorf("value[%d] validation failed for VR %s: %w", i, s.vr.Code(), err)
		}
	}

	return nil
}
