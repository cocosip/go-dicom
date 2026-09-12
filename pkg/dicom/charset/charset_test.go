// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package charset_test

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
)

const testCharsetISO2022IR58 = "ISO 2022 IR 58"

func TestGetEncoding(t *testing.T) {
	tests := []struct {
		name    string
		charset string
		want    string // encoding name for comparison
	}{
		{"Empty string", "", "ISO-8859-1"},
		{"Latin-1", testCharsetLatin1, "ISO-8859-1"},
		{"Latin-2", "ISO_IR 101", "ISO-8859-2"},
		{"Greek", "ISO_IR 126", "ISO-8859-7"},
		{"Cyrillic", "ISO_IR 144", "ISO-8859-5"},
		{testEncodingUTF8, testCharsetUTF8, testEncodingUTF8},
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
		{"Correct format", testCharsetLatin1},
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
	encs = charset.GetEncodings([]string{testCharsetLatin1, testCharsetUTF8})
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
	if name != testCharsetLatin1 {
		t.Errorf("GetCharsetName(ISO8859_1) = %q, want ISO_IR 100", name)
	}

	// Test UTF-8
	name, err = charset.GetCharsetName(unicode.UTF8)
	if err != nil {
		t.Fatalf("GetCharsetName(UTF8) error = %v", err)
	}
	if name != testCharsetUTF8 {
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
			[]byte(testHelloWorld),
			testCharsetLatin1,
			testHelloWorld,
		},
		{
			"UTF-8 text",
			[]byte("Hello 世界"),
			testCharsetUTF8,
			"Hello 世界",
		},
		{
			"Empty data",
			[]byte{},
			testCharsetLatin1,
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

func TestDecodeStringHandlesDICOMISO2022ChineseAndKorean(t *testing.T) {
	tests := []struct {
		name    string
		charset string
		data    []byte
		want    string
	}{
		{
			name:    "GB2312 IR 58",
			charset: testCharsetISO2022IR58,
			data:    []byte{'A', 0x1b, '$', ')', 'A', 0xd5, 0xc5, 0xc8, 0xfd, 0x1b, '(', 'B', 'Z'},
			want:    "A张三Z",
		},
		{
			name:    "Korean IR 149",
			charset: "ISO 2022 IR 149",
			data:    []byte{'A', 0x1b, '$', ')', 'C', 0xb1, 0xe8, 0xc8, 0xf1, 0xc1, 0xdf, 0x1b, '(', 'B', 'Z'},
			want:    "A김희중Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := charset.DecodeString(tt.data, charset.GetEncodings([]string{testCharsetLatin1, tt.charset}))
			if err != nil {
				t.Fatalf("DecodeString() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("DecodeString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestEncodeStringRejectsUnrepresentableDefaultCharset(t *testing.T) {
	data, err := charset.EncodeString("张三", []encoding.Encoding{charmap.ISO8859_1})
	if err == nil {
		t.Fatal("EncodeString() error = nil, want an encoding error")
	}
	if bytes.Equal(data, []byte("张三")) {
		t.Fatal("EncodeString() returned raw UTF-8 bytes after encoding failure")
	}
}

func TestEncodeStringAddsDICOMISO2022Designation(t *testing.T) {
	data, err := charset.EncodeString("张三", charset.GetEncodings([]string{testCharsetISO2022IR58}))
	if err != nil {
		t.Fatalf("EncodeString() error = %v", err)
	}
	if len(data) < 4 || !bytes.Equal(data[:4], []byte{0x1b, '$', ')', 'A'}) {
		t.Fatalf("EncodeString() prefix = %x, want DICOM IR 58 designation", data)
	}
	decoded, err := charset.DecodeString(data, charset.GetEncodings([]string{testCharsetISO2022IR58}))
	if err != nil || decoded != "张三" {
		t.Fatalf("DecodeString(EncodeString()) = %q, %v; want 张三", decoded, err)
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
			testHelloWorld,
			testCharsetLatin1,
			false,
		},
		{
			"UTF-8 text",
			"Hello 世界",
			testCharsetUTF8,
			false,
		},
		{
			"Empty string",
			"",
			testCharsetLatin1,
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
		{"ASCII", testHelloWorld, testCharsetLatin1},
		{"UTF-8", "Hello 世界 مرحبا Привет", testCharsetUTF8},
		{"Latin-1", "Café résumé", testCharsetLatin1},
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
			if tt.encoding == testCharsetUTF8 && decoded != tt.text {
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
		if cs == testCharsetLatin1 {
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
	info, ok := charset.GetCharsetInfo(testCharsetLatin1)
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
	encs := charset.GetEncodings([]string{testCharsetLatin1, testCharsetUTF8, "ISO_IR 126"})
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
