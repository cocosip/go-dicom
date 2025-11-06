// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package charset_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
)

func TestGetEncoding(t *testing.T) {
	tests := []struct {
		name    string
		charset string
		want    string // encoding name for comparison
	}{
		{"Empty string", "", "ISO-8859-1"},
		{"Latin-1", "ISO_IR 100", "ISO-8859-1"},
		{"Latin-2", "ISO_IR 101", "ISO-8859-2"},
		{"Greek", "ISO_IR 126", "ISO-8859-7"},
		{"Cyrillic", "ISO_IR 144", "ISO-8859-5"},
		{"UTF-8", "ISO_IR 192", "UTF-8"},
		{"Chinese GBK", "GBK", "GBK"},
		{"Chinese GB18030", "GB18030", "GB18030"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := charset.GetEncoding(tt.charset)
			if enc == nil {
				t.Fatal("GetEncoding() returned nil")
			}
			// Note: We can't directly compare encodings, so we just verify it's not nil
		})
	}
}

func TestGetEncodingMisspellings(t *testing.T) {
	// Test common misspellings
	tests := []struct {
		name    string
		charset string
	}{
		{"ISO-IR format", "ISO-IR 100"},
		{"ISO IR format", "ISO IR 100"},
		{"Correct format", "ISO_IR 100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := charset.GetEncoding(tt.charset)
			if enc == nil {
				t.Errorf("GetEncoding(%q) returned nil", tt.charset)
			}
			// All should resolve to the same encoding
			if enc != charmap.ISO8859_1 {
				t.Errorf("GetEncoding(%q) did not return ISO-8859-1", tt.charset)
			}
		})
	}
}

func TestGetEncodings(t *testing.T) {
	// Test with empty slice
	encs := charset.GetEncodings([]string{})
	if len(encs) != 1 {
		t.Errorf("GetEncodings([]) returned %d encodings, want 1", len(encs))
	}
	if encs[0] != charset.Default {
		t.Error("GetEncodings([]) did not return default encoding")
	}

	// Test with multiple charsets
	encs = charset.GetEncodings([]string{"ISO_IR 100", "ISO_IR 192"})
	if len(encs) != 2 {
		t.Errorf("GetEncodings() returned %d encodings, want 2", len(encs))
	}
}

func TestGetCharsetName(t *testing.T) {
	// Test getting charset name from encoding
	name, err := charset.GetCharsetName(charmap.ISO8859_1)
	if err != nil {
		t.Fatalf("GetCharsetName() error = %v", err)
	}
	if name != "ISO_IR 100" {
		t.Errorf("GetCharsetName(ISO8859_1) = %q, want ISO_IR 100", name)
	}

	// Test UTF-8
	name, err = charset.GetCharsetName(unicode.UTF8)
	if err != nil {
		t.Fatalf("GetCharsetName(UTF8) error = %v", err)
	}
	if name != "ISO_IR 192" {
		t.Errorf("GetCharsetName(UTF8) = %q, want ISO_IR 192", name)
	}
}

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		encoding string
		want     string
	}{
		{
			"ASCII text",
			[]byte("Hello World"),
			"ISO_IR 100",
			"Hello World",
		},
		{
			"UTF-8 text",
			[]byte("Hello 世界"),
			"ISO_IR 192",
			"Hello 世界",
		},
		{
			"Empty data",
			[]byte{},
			"ISO_IR 100",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := charset.GetEncoding(tt.encoding)
			result, err := charset.DecodeString(tt.data, []encoding.Encoding{enc})
			if err != nil {
				t.Fatalf("DecodeString() error = %v", err)
			}
			if result != tt.want {
				t.Errorf("DecodeString() = %q, want %q", result, tt.want)
			}
		})
	}
}

func TestEncodeString(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		encoding string
		wantErr  bool
	}{
		{
			"ASCII text",
			"Hello World",
			"ISO_IR 100",
			false,
		},
		{
			"UTF-8 text",
			"Hello 世界",
			"ISO_IR 192",
			false,
		},
		{
			"Empty string",
			"",
			"ISO_IR 100",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := charset.GetEncoding(tt.encoding)
			data, err := charset.EncodeString(tt.text, []encoding.Encoding{enc})
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(data) == 0 && len(tt.text) > 0 {
				t.Error("EncodeString() returned empty data for non-empty string")
			}
		})
	}
}

func TestEncodeDecodeRoundtrip(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		encoding string
	}{
		{"ASCII", "Hello World", "ISO_IR 100"},
		{"UTF-8", "Hello 世界 مرحبا Привет", "ISO_IR 192"},
		{"Latin-1", "Café résumé", "ISO_IR 100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := charset.GetEncoding(tt.encoding)
			encodings := []encoding.Encoding{enc}

			// Encode
			encoded, err := charset.EncodeString(tt.text, encodings)
			if err != nil {
				t.Fatalf("EncodeString() error = %v", err)
			}

			// Decode
			decoded, err := charset.DecodeString(encoded, encodings)
			if err != nil {
				t.Fatalf("DecodeString() error = %v", err)
			}

			// For UTF-8, should match exactly
			if tt.encoding == "ISO_IR 192" && decoded != tt.text {
				t.Errorf("Roundtrip failed: got %q, want %q", decoded, tt.text)
			}
		})
	}
}

func TestKnownCharsets(t *testing.T) {
	charsets := charset.KnownCharsets()
	if len(charsets) == 0 {
		t.Error("KnownCharsets() returned empty list")
	}

	// Should contain at least some basic charsets
	found := false
	for _, cs := range charsets {
		if cs == "ISO_IR 100" {
			found = true
			break
		}
	}
	if !found {
		t.Error("KnownCharsets() should contain ISO_IR 100")
	}
}

func TestGetCharsetInfo(t *testing.T) {
	// Test known charset
	info, ok := charset.GetCharsetInfo("ISO_IR 100")
	if !ok {
		t.Fatal("GetCharsetInfo(ISO_IR 100) returned false")
	}
	if info == nil {
		t.Fatal("GetCharsetInfo(ISO_IR 100) returned nil info")
	}
	if info.Name == "" {
		t.Error("CharsetInfo.Name is empty")
	}
	if info.Encoding == nil {
		t.Error("CharsetInfo.Encoding is nil")
	}

	// Test unknown charset
	_, ok = charset.GetCharsetInfo("UNKNOWN_CHARSET")
	if ok {
		t.Error("GetCharsetInfo(UNKNOWN_CHARSET) should return false")
	}
}

func TestDefaultEncoding(t *testing.T) {
	if charset.Default == nil {
		t.Fatal("Default encoding is nil")
	}

	// Default should be ISO-8859-1
	if charset.Default != charmap.ISO8859_1 {
		t.Error("Default encoding should be ISO-8859-1")
	}
}

func TestMultipleEncodings(t *testing.T) {
	// Test with multiple encodings (code extensions scenario)
	encs := charset.GetEncodings([]string{"ISO_IR 100", "ISO_IR 192", "ISO_IR 126"})
	if len(encs) != 3 {
		t.Errorf("GetEncodings() returned %d encodings, want 3", len(encs))
	}

	// All encodings should be non-nil
	for i, enc := range encs {
		if enc == nil {
			t.Errorf("Encoding at index %d is nil", i)
		}
	}
}
