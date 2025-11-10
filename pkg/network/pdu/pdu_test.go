// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"bytes"
	"io"
	"testing"
)

func TestPDUTypeString(t *testing.T) {
	tests := []struct {
		name     string
		pduType  byte
		expected string
	}{
		{"A-ASSOCIATE-RQ", TypeAAssociateRQ, "A-ASSOCIATE-RQ"},
		{"A-ASSOCIATE-AC", TypeAAssociateAC, "A-ASSOCIATE-AC"},
		{"A-ASSOCIATE-RJ", TypeAAssociateRJ, "A-ASSOCIATE-RJ"},
		{"P-DATA-TF", TypePDataTF, "P-DATA-TF"},
		{"A-RELEASE-RQ", TypeAReleaseRQ, "A-RELEASE-RQ"},
		{"A-RELEASE-RP", TypeAReleaseRP, "A-RELEASE-RP"},
		{"A-ABORT", TypeAAbort, "A-ABORT"},
		{"Unknown", 0xFF, "Unknown(0xFF)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PDUTypeString(tt.pduType)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestNewRawPDU(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04}
	pdu := NewRawPDU(TypeAAssociateRQ, data)

	if pdu.Type != TypeAAssociateRQ {
		t.Errorf("Expected type 0x01, got 0x%02X", pdu.Type)
	}
	if pdu.Reserved != 0x00 {
		t.Errorf("Expected reserved 0x00, got 0x%02X", pdu.Reserved)
	}
	if pdu.Length != 4 {
		t.Errorf("Expected length 4, got %d", pdu.Length)
	}
	if !bytes.Equal(pdu.Data, data) {
		t.Errorf("Expected data %v, got %v", data, pdu.Data)
	}
}

func TestRawPDU_String(t *testing.T) {
	pdu := NewRawPDU(TypeAAssociateRQ, []byte{0x01, 0x02})
	expected := "A-ASSOCIATE-RQ (length=2)"
	if pdu.String() != expected {
		t.Errorf("Expected %s, got %s", expected, pdu.String())
	}
}

func TestRawPDU_TotalLength(t *testing.T) {
	tests := []struct {
		name          string
		dataLength    int
		expectedTotal uint32
	}{
		{"Empty data", 0, 6},
		{"4 bytes", 4, 10},
		{"100 bytes", 100, 106},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]byte, tt.dataLength)
			pdu := NewRawPDU(TypeAAssociateRQ, data)
			if pdu.TotalLength() != tt.expectedTotal {
				t.Errorf("Expected total length %d, got %d", tt.expectedTotal, pdu.TotalLength())
			}
		})
	}
}

func TestRawPDU_WriteRead(t *testing.T) {
	// Test roundtrip: Write -> Read
	tests := []struct {
		name    string
		pduType byte
		data    []byte
	}{
		{"A-ASSOCIATE-RQ with data", TypeAAssociateRQ, []byte{0x01, 0x02, 0x03, 0x04}},
		{"A-RELEASE-RQ empty", TypeAReleaseRQ, []byte{}},
		{"P-DATA-TF with large data", TypePDataTF, make([]byte, 1000)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create original PDU
			original := NewRawPDU(tt.pduType, tt.data)

			// Write to buffer
			var buf bytes.Buffer
			if err := original.Write(&buf); err != nil {
				t.Fatalf("Write failed: %v", err)
			}

			// Read back
			read := &RawPDU{}
			if err := read.Read(&buf); err != nil {
				t.Fatalf("Read failed: %v", err)
			}

			// Compare
			if read.Type != original.Type {
				t.Errorf("Type mismatch: expected 0x%02X, got 0x%02X", original.Type, read.Type)
			}
			if read.Reserved != original.Reserved {
				t.Errorf("Reserved mismatch: expected 0x%02X, got 0x%02X", original.Reserved, read.Reserved)
			}
			if read.Length != original.Length {
				t.Errorf("Length mismatch: expected %d, got %d", original.Length, read.Length)
			}
			if !bytes.Equal(read.Data, original.Data) {
				t.Errorf("Data mismatch")
			}
		})
	}
}

func TestRawPDU_Read_InvalidPDUType(t *testing.T) {
	// Invalid PDU type (0xFF)
	data := []byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x04, 0x01, 0x02, 0x03, 0x04}
	buf := bytes.NewReader(data)

	pdu := &RawPDU{}
	err := pdu.Read(buf)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestRawPDU_Read_IncompleteHeader(t *testing.T) {
	// Only 3 bytes (incomplete header)
	data := []byte{0x01, 0x00, 0x00}
	buf := bytes.NewReader(data)

	pdu := &RawPDU{}
	err := pdu.Read(buf)
	if err == nil {
		t.Error("Expected error for incomplete header, got nil")
	}
	if err != io.ErrUnexpectedEOF && err != io.EOF {
		// Check if error is wrapped
		if !bytes.Contains([]byte(err.Error()), []byte("unexpected EOF")) &&
			!bytes.Contains([]byte(err.Error()), []byte("EOF")) {
			t.Errorf("Expected EOF error, got: %v", err)
		}
	}
}

func TestRawPDU_Read_IncompleteData(t *testing.T) {
	// Header says 4 bytes of data, but only 2 bytes provided
	data := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x01, 0x02}
	buf := bytes.NewReader(data)

	pdu := &RawPDU{}
	err := pdu.Read(buf)
	if err == nil {
		t.Error("Expected error for incomplete data, got nil")
	}
}

func TestReadPDU(t *testing.T) {
	// Test convenience function
	data := []byte{0x01, 0x02, 0x03, 0x04}
	original := NewRawPDU(TypeAAssociateRQ, data)

	var buf bytes.Buffer
	if err := original.Write(&buf); err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	pdu, err := ReadPDU(&buf)
	if err != nil {
		t.Fatalf("ReadPDU failed: %v", err)
	}

	if pdu.Type != original.Type {
		t.Errorf("Type mismatch: expected 0x%02X, got 0x%02X", original.Type, pdu.Type)
	}
	if pdu.Length != original.Length {
		t.Errorf("Length mismatch: expected %d, got %d", original.Length, pdu.Length)
	}
}

func TestWritePDU(t *testing.T) {
	// Test convenience function
	data := []byte{0x01, 0x02, 0x03, 0x04}
	pdu := NewRawPDU(TypeAAssociateRQ, data)

	var buf bytes.Buffer
	if err := WritePDU(&buf, pdu); err != nil {
		t.Fatalf("WritePDU failed: %v", err)
	}

	// Verify header
	header := buf.Bytes()[:6]
	if header[0] != TypeAAssociateRQ {
		t.Errorf("Expected type 0x01, got 0x%02X", header[0])
	}
	if header[1] != 0x00 {
		t.Errorf("Expected reserved 0x00, got 0x%02X", header[1])
	}
	// Length should be 4 in big-endian
	if header[2] != 0x00 || header[3] != 0x00 || header[4] != 0x00 || header[5] != 0x04 {
		t.Errorf("Expected length bytes [0x00, 0x00, 0x00, 0x04], got %v", header[2:6])
	}
}

func TestIsValidPDUType(t *testing.T) {
	validTypes := []byte{
		TypeAAssociateRQ, TypeAAssociateAC, TypeAAssociateRJ,
		TypePDataTF, TypeAReleaseRQ, TypeAReleaseRP, TypeAAbort,
	}

	for _, pduType := range validTypes {
		if !isValidPDUType(pduType) {
			t.Errorf("Expected 0x%02X to be valid", pduType)
		}
	}

	invalidTypes := []byte{0x00, 0x08, 0x09, 0xFF}
	for _, pduType := range invalidTypes {
		if isValidPDUType(pduType) {
			t.Errorf("Expected 0x%02X to be invalid", pduType)
		}
	}
}
