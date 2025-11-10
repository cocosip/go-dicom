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
)

// String represents a DICOM string element.
//
// String elements can contain single or multiple values separated by backslashes.
// They handle character set encoding and provide methods for accessing values as strings.
//
// Applicable VRs: AE, AS, CS, DA, DS, DT, IS, LO, LT, PN, SH, ST, TM, UC, UI, UR, UT
type String struct {
	*base
	encoding encoding.Encoding
}

// NewString creates a new string element with the given tag, VR, and values.
// Values are joined with backslash separators and encoded into a buffer.
func NewString(t *tag.Tag, v *vr.VR, values []string) *String {
	return NewStringWithEncoding(t, v, values, charset.Default)
}

// NewStringWithEncoding creates a new string element with a specific encoding.
func NewStringWithEncoding(t *tag.Tag, v *vr.VR, values []string, enc encoding.Encoding) *String {
	// Join values with backslash separator
	joined := strings.Join(values, "\\")

	// Encode to bytes
	var data []byte
	if enc != nil {
		encoder := enc.NewEncoder()
		encoded, err := encoder.Bytes([]byte(joined))
		if err != nil {
			// Fallback to raw bytes on encoding error
			data = []byte(joined)
		} else {
			data = encoded
		}
	} else {
		data = []byte(joined)
	}

	// Remove trailing spaces (DICOM requirement)
	data = bytes.TrimRight(data, " ")

	// Create buffer
	buf := buffer.NewMemory(data)

	return &String{
		base:     newBase(t, v, buf),
		encoding: enc,
	}
}

// NewStringFromBuffer creates a string element from an existing buffer.
// This is used when parsing DICOM data.
func NewStringFromBuffer(t *tag.Tag, v *vr.VR, buf buffer.ByteBuffer, enc encoding.Encoding) *String {
	if enc == nil {
		enc = charset.Default
	}
	return &String{
		base:     newBase(t, v, buf),
		encoding: enc,
	}
}

// Count returns the number of values in this element.
// Values are separated by backslashes.
func (s *String) Count() int {
	str := s.GetString()
	if str == "" {
		return 0
	}
	return strings.Count(str, "\\") + 1
}

// GetString returns the complete string value (all values joined by backslash).
func (s *String) GetString() string {
	if s.buffer == nil || s.buffer.Size() == 0 {
		return ""
	}

	data := s.buffer.Data()
	if s.encoding != nil {
		decoder := s.encoding.NewDecoder()
		decoded, err := decoder.Bytes(data)
		if err != nil {
			// Fallback to raw string on decoding error
			return string(data)
		}
		return string(decoded)
	}
	return string(data)
}

// GetValue returns the value at the specified index.
// If index is -1 or out of range, returns the complete string.
func (s *String) GetValue(index int) string {
	str := s.GetString()
	if str == "" {
		return ""
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

	values := strings.Split(str, "\\")
	// Trim spaces from each value
	for i, v := range values {
		values[i] = strings.TrimSpace(v)
	}
	return values
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

	// Check maximum length based on VR
	maxLen := s.vr.MaximumLength()
	if maxLen > 0 && s.Length() > maxLen {
		return fmt.Errorf("value length %d exceeds maximum %d for VR %s", s.Length(), maxLen, s.vr.Code())
	}

	// Perform VR-specific validation on each value
	values := s.GetValues()
	for i, value := range values {
		if err := s.vr.ValidateString(value); err != nil {
			return fmt.Errorf("value[%d] validation failed for VR %s: %w", i, s.vr.Code(), err)
		}
	}

	return nil
}
