// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package charset_test

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
)

const testCharsetISO2022IR87 = "ISO 2022 IR 87"

// TestEncodeValueSetResetsAroundPersonNameDelimiters verifies PS3.5 6.1.2.5.3:
// the Value 1 character set is active around the 5EH and 3DH delimiters, and an
// extension character set is designated inside each name component that needs
// it.
func TestEncodeValueSetResetsAroundPersonNameDelimiters(t *testing.T) {
	encodings := charset.GetEncodings([]string{testCharsetLatin1, testCharsetISO2022IR87})
	const personName = "Doe^John=山田^太郎"

	encoded, err := charset.EncodeValueSet(personName, "PN", encodings)
	if err != nil {
		t.Fatalf("EncodeValueSet() error = %v", err)
	}
	if !bytes.HasPrefix(encoded, []byte("Doe^John=")) {
		t.Fatalf("encoded alphabetic group = %x, want the unescaped value 1 bytes", encoded)
	}
	// Both ideographic components must be closed with ESC ( B before the
	// following delimiter, so the default repertoire is active at ^ and at the
	// end of the value.
	reset := []byte{0x1b, '(', 'B'}
	if count := bytes.Count(encoded, reset); count < 2 {
		t.Fatalf("encoded value has %d default-repertoire resets, want at least 2: %x", count, encoded)
	}
	// No extension character may run directly into a delimiter: the reset
	// sequence has to stand between them.
	for index := 0; index < len(encoded)-1; index++ {
		next := encoded[index+1]
		if (next == '^' || next == '=') && encoded[index] >= 0x80 {
			t.Fatalf("byte %#02x runs into the %q delimiter without a reset: %x", encoded[index], next, encoded)
		}
	}

	decoded, err := charset.DecodeString(encoded, encodings)
	if err != nil {
		t.Fatalf("DecodeString() error = %v", err)
	}
	if decoded != personName {
		t.Fatalf("DecodeString(EncodeValueSet()) = %q, want %q", decoded, personName)
	}
}

// TestEncodeValueSetResetsBeforeValueSeparator verifies PS3.5 6.1.2.5.3: the
// Value 1 character set is active at the beginning of every value.
func TestEncodeValueSetResetsBeforeValueSeparator(t *testing.T) {
	encodings := charset.GetEncodings([]string{testCharsetLatin1, testCharsetISO2022IR87})
	const values = "山田\\太郎"

	encoded, err := charset.EncodeValueSet(values, "LO", encodings)
	if err != nil {
		t.Fatalf("EncodeValueSet() error = %v", err)
	}
	// Each value carries its own designation, so the second value starts with
	// the ISO 2022 IR 87 escape right after the separator.
	if !bytes.Contains(encoded, []byte{'\\', 0x1b, '$', 'B'}) {
		t.Fatalf("second value does not designate its character set: %x", encoded)
	}

	decoded, err := charset.DecodeString(encoded, encodings)
	if err != nil {
		t.Fatalf("DecodeString() error = %v", err)
	}
	if decoded != values {
		t.Fatalf("DecodeString(EncodeValueSet()) = %q, want %q", decoded, values)
	}
}

// TestEncodeValueSetKeepsBackslashInSingleValuedText verifies that ST, LT and UT
// keep a 05/12 byte as text: those VRs are single-valued, so it is not a value
// separator (PS3.5 6.1.2.3).
func TestEncodeValueSetKeepsBackslashInSingleValuedText(t *testing.T) {
	const text = "line one\\line two"
	for _, vrCode := range []string{"ST", "LT", "UT"} {
		encoded, err := charset.EncodeValueSet(text, vrCode, charset.GetEncodings([]string{testCharsetUTF8}))
		if err != nil {
			t.Fatalf("EncodeValueSet(%s) error = %v", vrCode, err)
		}
		if string(encoded) != text {
			t.Fatalf("EncodeValueSet(%s) = %q, want %q", vrCode, encoded, text)
		}
	}
}

// TestEncodeValueSetRejectsBytesOutsideTIS620 verifies PS3.3 C.12.1.1.3: ISO_IR
// 166 is TIS 620-2533, which leaves 0x80-0x9F undefined even though the
// underlying code page assigns characters there.
func TestEncodeValueSetRejectsBytesOutsideTIS620(t *testing.T) {
	if _, err := charset.EncodeValueSet("€", "LO", charset.GetEncodings([]string{"ISO_IR 166"})); err == nil {
		t.Fatal("EncodeValueSet() encoded a TIS 620-2533 byte position that TIS 620 leaves undefined")
	}
}

// TestEncodeValueSetRejectsGBKExtensionForIR58 verifies PS3.3 C.12.1.1.5:
// ISO 2022 IR 58 is GB 2312, so GBK-only code positions must not be emitted.
func TestEncodeValueSetRejectsGBKExtensionForIR58(t *testing.T) {
	if _, err := charset.EncodeValueSet("€", "LO", charset.GetEncodings([]string{testCharsetISO2022IR58})); err == nil {
		t.Fatal("EncodeValueSet() encoded a GBK code position under ISO 2022 IR 58")
	}
}
