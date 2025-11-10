// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

func TestPDVEncode(t *testing.T) {
	pdv := &PDV{
		PresentationContextID: 1,
		IsCommand:             true,
		IsLastFragment:        true,
		Data:                  []byte("test data"),
	}

	encoded := pdv.Encode()

	// Verify length (4 bytes length + 1 byte context ID + 1 byte header + data)
	expectedLength := 4 + 1 + 1 + len(pdv.Data)
	if len(encoded) != expectedLength {
		t.Errorf("Expected length %d, got %d", expectedLength, len(encoded))
	}

	// Verify presentation context ID
	if encoded[4] != 1 {
		t.Errorf("Expected context ID 1, got %d", encoded[4])
	}

	// Verify header (command=1, last=1 -> 0x03)
	expectedHeader := byte(0x03)
	if encoded[5] != expectedHeader {
		t.Errorf("Expected header 0x%02x, got 0x%02x", expectedHeader, encoded[5])
	}

	// Verify data
	if !bytes.Equal(encoded[6:], pdv.Data) {
		t.Error("Data mismatch")
	}
}

func TestPDVDecode(t *testing.T) {
	// Create test data
	data := make([]byte, 15)
	// Length: 11 (1 + 1 + 9)
	data[0] = 0x00
	data[1] = 0x00
	data[2] = 0x00
	data[3] = 0x0B
	// Context ID: 1
	data[4] = 0x01
	// Header: command + last fragment
	data[5] = 0x03
	// Data: "test data"
	copy(data[6:], []byte("test data"))

	pdv, err := DecodePDV(data)
	if err != nil {
		t.Fatalf("DecodePDV failed: %v", err)
	}

	if pdv.PresentationContextID != 1 {
		t.Errorf("Expected context ID 1, got %d", pdv.PresentationContextID)
	}

	if !pdv.IsCommand {
		t.Error("Expected IsCommand=true")
	}

	if !pdv.IsLastFragment {
		t.Error("Expected IsLastFragment=true")
	}

	if !bytes.Equal(pdv.Data, []byte("test data")) {
		t.Errorf("Data mismatch: got %v", pdv.Data)
	}
}

func TestPDVEncodeDecodeRoundTrip(t *testing.T) {
	tests := []struct {
		name              string
		presentationCtxID byte
		isCommand         bool
		isLastFragment    bool
		data              []byte
	}{
		{"Command Last", 1, true, true, []byte("command")},
		{"Command Not Last", 1, true, false, []byte("command part 1")},
		{"Dataset Last", 3, false, true, []byte("dataset")},
		{"Dataset Not Last", 3, false, false, []byte("dataset part 1")},
		{"Empty Data", 1, true, true, []byte{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := &PDV{
				PresentationContextID: tt.presentationCtxID,
				IsCommand:             tt.isCommand,
				IsLastFragment:        tt.isLastFragment,
				Data:                  tt.data,
			}

			encoded := original.Encode()
			decoded, err := DecodePDV(encoded)
			if err != nil {
				t.Fatalf("DecodePDV failed: %v", err)
			}

			if decoded.PresentationContextID != original.PresentationContextID {
				t.Errorf("Context ID mismatch: got %d, want %d",
					decoded.PresentationContextID, original.PresentationContextID)
			}

			if decoded.IsCommand != original.IsCommand {
				t.Errorf("IsCommand mismatch: got %v, want %v",
					decoded.IsCommand, original.IsCommand)
			}

			if decoded.IsLastFragment != original.IsLastFragment {
				t.Errorf("IsLastFragment mismatch: got %v, want %v",
					decoded.IsLastFragment, original.IsLastFragment)
			}

			if !bytes.Equal(decoded.Data, original.Data) {
				t.Errorf("Data mismatch: got %v, want %v", decoded.Data, original.Data)
			}
		})
	}
}

func TestCreatePDataTFPDU(t *testing.T) {
	pdvs := []*PDV{
		{
			PresentationContextID: 1,
			IsCommand:             true,
			IsLastFragment:        false,
			Data:                  []byte("part1"),
		},
		{
			PresentationContextID: 1,
			IsCommand:             true,
			IsLastFragment:        true,
			Data:                  []byte("part2"),
		},
	}

	pduData, err := CreatePDataTFPDU(pdvs)
	if err != nil {
		t.Fatalf("CreatePDataTFPDU failed: %v", err)
	}

	if pduData.Type != pdu.TypePDataTF {
		t.Errorf("Expected PDU type P-DATA-TF, got %d", pduData.Type)
	}

	// Verify we can parse it back
	parsedPDVs, err := ParsePDataTFPDU(pduData)
	if err != nil {
		t.Fatalf("ParsePDataTFPDU failed: %v", err)
	}

	if len(parsedPDVs) != len(pdvs) {
		t.Errorf("Expected %d PDVs, got %d", len(pdvs), len(parsedPDVs))
	}

	for i, pdv := range parsedPDVs {
		if !bytes.Equal(pdv.Data, pdvs[i].Data) {
			t.Errorf("PDV %d data mismatch", i)
		}
	}
}

func TestParsePDataTFPDU(t *testing.T) {
	// Create a P-DATA-TF PDU with multiple PDVs
	pdvs := []*PDV{
		{
			PresentationContextID: 1,
			IsCommand:             true,
			IsLastFragment:        true,
			Data:                  []byte("command data"),
		},
		{
			PresentationContextID: 3,
			IsCommand:             false,
			IsLastFragment:        true,
			Data:                  []byte("dataset data"),
		},
	}

	pduData, err := CreatePDataTFPDU(pdvs)
	if err != nil {
		t.Fatalf("CreatePDataTFPDU failed: %v", err)
	}

	parsed, err := ParsePDataTFPDU(pduData)
	if err != nil {
		t.Fatalf("ParsePDataTFPDU failed: %v", err)
	}

	if len(parsed) != 2 {
		t.Fatalf("Expected 2 PDVs, got %d", len(parsed))
	}

	// Verify first PDV
	if parsed[0].PresentationContextID != 1 {
		t.Errorf("PDV 0: wrong context ID")
	}
	if !parsed[0].IsCommand {
		t.Errorf("PDV 0: should be command")
	}
	if !bytes.Equal(parsed[0].Data, []byte("command data")) {
		t.Errorf("PDV 0: data mismatch")
	}

	// Verify second PDV
	if parsed[1].PresentationContextID != 3 {
		t.Errorf("PDV 1: wrong context ID")
	}
	if parsed[1].IsCommand {
		t.Errorf("PDV 1: should not be command")
	}
	if !bytes.Equal(parsed[1].Data, []byte("dataset data")) {
		t.Errorf("PDV 1: data mismatch")
	}
}

func TestFragmentData(t *testing.T) {
	tests := []struct {
		name          string
		data          []byte
		maxPDULength  uint32
		wantFragments int
	}{
		{"Small data", []byte("small"), 16384, 1},
		{"Empty data", []byte{}, 16384, 1},
		{"Large data", make([]byte, 50000), 16384, 4}, // Should split into multiple fragments
		{"Exact fit", make([]byte, 16384-12), 16384, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdvs := FragmentData(tt.data, 1, true, tt.maxPDULength)

			if len(pdvs) != tt.wantFragments {
				t.Errorf("Expected %d fragments, got %d", tt.wantFragments, len(pdvs))
			}

			// Verify all except last have IsLastFragment=false
			for i := 0; i < len(pdvs)-1; i++ {
				if pdvs[i].IsLastFragment {
					t.Errorf("Fragment %d should not be last", i)
				}
			}

			// Verify last fragment has IsLastFragment=true
			if len(pdvs) > 0 && !pdvs[len(pdvs)-1].IsLastFragment {
				t.Error("Last fragment should have IsLastFragment=true")
			}

			// Verify data integrity
			var reconstructed []byte
			for _, pdv := range pdvs {
				reconstructed = append(reconstructed, pdv.Data...)
			}

			if !bytes.Equal(reconstructed, tt.data) {
				t.Error("Reconstructed data doesn't match original")
			}
		})
	}
}

func TestFragmentDataMultiplePDUs(t *testing.T) {
	// Test with data larger than one PDU
	data := make([]byte, 100000) // 100KB
	for i := range data {
		data[i] = byte(i % 256)
	}

	pdvs := FragmentData(data, 1, false, 16384)

	// Should create multiple fragments
	if len(pdvs) <= 1 {
		t.Errorf("Expected multiple fragments for 100KB data, got %d", len(pdvs))
	}

	// Verify all fragments have correct context ID and isCommand
	for i, pdv := range pdvs {
		if pdv.PresentationContextID != 1 {
			t.Errorf("Fragment %d: wrong context ID", i)
		}
		if pdv.IsCommand {
			t.Errorf("Fragment %d: should not be command", i)
		}
	}

	// Reassemble and verify
	var reconstructed []byte
	for _, pdv := range pdvs {
		reconstructed = append(reconstructed, pdv.Data...)
	}

	if !bytes.Equal(reconstructed, data) {
		t.Error("Data corruption during fragmentation")
	}
}
