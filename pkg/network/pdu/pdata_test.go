// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pdu

import (
	"bytes"
	"testing"
)

func TestNewPDataTF(t *testing.T) {
	pdata := NewPDataTF()
	if pdata == nil {
		t.Error("Expected non-nil P-DATA-TF")
	}
	if len(pdata.PDVs) != 0 {
		t.Errorf("Expected 0 PDVs, got %d", len(pdata.PDVs))
	}
}

func TestNewPDV(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04}
	pdv := NewPDV(1, true, true, data)

	if pdv.PresentationContextID != 1 {
		t.Errorf("Expected context ID 1, got %d", pdv.PresentationContextID)
	}
	if !pdv.IsCommand {
		t.Error("Expected IsCommand to be true")
	}
	if !pdv.IsLastFragment {
		t.Error("Expected IsLastFragment to be true")
	}
	if !bytes.Equal(pdv.Data, data) {
		t.Errorf("Data mismatch")
	}
}

func TestPDataTF_EncodeDecodeSinglePDV(t *testing.T) {
	// Create P-DATA-TF with single PDV
	original := NewPDataTF()
	data := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
	pdv := NewPDV(1, true, true, data)
	original.AddPDV(pdv)

	// Encode
	pdu, err := original.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	if pdu.Type != TypePDataTF {
		t.Errorf("Expected PDU type 0x04, got 0x%02X", pdu.Type)
	}

	// Decode
	decoded := &PDataTF{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.PDVs) != 1 {
		t.Fatalf("Expected 1 PDV, got %d", len(decoded.PDVs))
	}

	decodedPDV := decoded.PDVs[0]
	if decodedPDV.PresentationContextID != pdv.PresentationContextID {
		t.Errorf("Context ID mismatch: expected %d, got %d", pdv.PresentationContextID, decodedPDV.PresentationContextID)
	}
	if decodedPDV.IsCommand != pdv.IsCommand {
		t.Errorf("IsCommand mismatch: expected %v, got %v", pdv.IsCommand, decodedPDV.IsCommand)
	}
	if decodedPDV.IsLastFragment != pdv.IsLastFragment {
		t.Errorf("IsLastFragment mismatch: expected %v, got %v", pdv.IsLastFragment, decodedPDV.IsLastFragment)
	}
	if !bytes.Equal(decodedPDV.Data, pdv.Data) {
		t.Errorf("Data mismatch")
	}
}

func TestPDataTF_EncodeDecodeMultiplePDVs(t *testing.T) {
	// Create P-DATA-TF with multiple PDVs
	original := NewPDataTF()

	// First PDV: command data, not last
	original.AddPDV(NewPDV(1, true, false, []byte{0x01, 0x02}))

	// Second PDV: command data, last
	original.AddPDV(NewPDV(1, true, true, []byte{0x03, 0x04}))

	// Third PDV: dataset data, last
	original.AddPDV(NewPDV(1, false, true, []byte{0x05, 0x06, 0x07, 0x08}))

	// Encode
	pdu, err := original.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// Decode
	decoded := &PDataTF{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify
	if len(decoded.PDVs) != 3 {
		t.Fatalf("Expected 3 PDVs, got %d", len(decoded.PDVs))
	}

	// Check first PDV
	if decoded.PDVs[0].IsCommand != true {
		t.Error("PDV 0: expected command")
	}
	if decoded.PDVs[0].IsLastFragment != false {
		t.Error("PDV 0: expected not last fragment")
	}

	// Check second PDV
	if decoded.PDVs[1].IsCommand != true {
		t.Error("PDV 1: expected command")
	}
	if decoded.PDVs[1].IsLastFragment != true {
		t.Error("PDV 1: expected last fragment")
	}

	// Check third PDV
	if decoded.PDVs[2].IsCommand != false {
		t.Error("PDV 2: expected data (not command)")
	}
	if decoded.PDVs[2].IsLastFragment != true {
		t.Error("PDV 2: expected last fragment")
	}
}

func TestPDataTF_MessageControlHeader(t *testing.T) {
	tests := []struct {
		name           string
		isCommand      bool
		isLastFragment bool
		expectedHeader byte
	}{
		{"Command, last", true, true, 0x03},      // 0b00000011
		{"Command, not last", true, false, 0x01}, // 0b00000001
		{"Data, last", false, true, 0x02},        // 0b00000010
		{"Data, not last", false, false, 0x00},   // 0b00000000
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdata := NewPDataTF()
			pdata.AddPDV(NewPDV(1, tt.isCommand, tt.isLastFragment, []byte{0xAA, 0xBB}))

			// Encode
			pdu, err := pdata.Encode()
			if err != nil {
				t.Fatalf("Encode failed: %v", err)
			}

			// Check message control header byte
			// PDU data format: [PDV Length (4)] [Context ID (1)] [Header (1)] [Data...]
			if len(pdu.Data) < 6 {
				t.Fatalf("PDU data too short: %d bytes", len(pdu.Data))
			}

			actualHeader := pdu.Data[5] // Byte 5 is the message control header
			if actualHeader != tt.expectedHeader {
				t.Errorf("Header mismatch: expected 0x%02X, got 0x%02X", tt.expectedHeader, actualHeader)
			}

			// Decode and verify
			decoded := &PDataTF{}
			if err := decoded.Decode(pdu); err != nil {
				t.Fatalf("Decode failed: %v", err)
			}

			if decoded.PDVs[0].IsCommand != tt.isCommand {
				t.Errorf("IsCommand mismatch")
			}
			if decoded.PDVs[0].IsLastFragment != tt.isLastFragment {
				t.Errorf("IsLastFragment mismatch")
			}
		})
	}
}

func TestPDataTF_DecodeInvalidPDUType(t *testing.T) {
	pdu := NewRawPDU(TypeAAssociateRQ, []byte{0x00})

	pdata := &PDataTF{}
	err := pdata.Decode(pdu)
	if err == nil {
		t.Error("Expected error for invalid PDU type, got nil")
	}
}

func TestPDataTF_EncodeNoPDVs(t *testing.T) {
	pdata := NewPDataTF()

	_, err := pdata.Encode()
	if err == nil {
		t.Error("Expected error when encoding P-DATA-TF with no PDVs")
	}
}

func TestPDataTF_DecodeNoPDVs(t *testing.T) {
	// Create PDU with empty data
	pdu := NewRawPDU(TypePDataTF, []byte{})

	pdata := &PDataTF{}
	err := pdata.Decode(pdu)
	if err == nil {
		t.Error("Expected error when decoding P-DATA-TF with no PDVs")
	}
}

func TestPDataTF_TotalDataSize(t *testing.T) {
	pdata := NewPDataTF()
	pdata.AddPDV(NewPDV(1, true, false, make([]byte, 100)))
	pdata.AddPDV(NewPDV(1, true, true, make([]byte, 200)))
	pdata.AddPDV(NewPDV(1, false, true, make([]byte, 300)))

	expected := 100 + 200 + 300
	if pdata.TotalDataSize() != expected {
		t.Errorf("TotalDataSize: expected %d, got %d", expected, pdata.TotalDataSize())
	}
}

func TestPDataTF_GetCommandPDVs(t *testing.T) {
	pdata := NewPDataTF()
	pdata.AddPDV(NewPDV(1, true, false, []byte{0x01}))  // Command
	pdata.AddPDV(NewPDV(1, true, true, []byte{0x02}))   // Command
	pdata.AddPDV(NewPDV(1, false, true, []byte{0x03}))  // Data

	commandPDVs := pdata.GetCommandPDVs()
	if len(commandPDVs) != 2 {
		t.Errorf("Expected 2 command PDVs, got %d", len(commandPDVs))
	}
}

func TestPDataTF_GetDataPDVs(t *testing.T) {
	pdata := NewPDataTF()
	pdata.AddPDV(NewPDV(1, true, true, []byte{0x01}))   // Command
	pdata.AddPDV(NewPDV(1, false, false, []byte{0x02})) // Data
	pdata.AddPDV(NewPDV(1, false, true, []byte{0x03}))  // Data

	dataPDVs := pdata.GetDataPDVs()
	if len(dataPDVs) != 2 {
		t.Errorf("Expected 2 data PDVs, got %d", len(dataPDVs))
	}
}

func TestFragmentData(t *testing.T) {
	// Test data fragmentation
	data := make([]byte, 1000)
	for i := range data {
		data[i] = byte(i % 256)
	}

	// Fragment into 300-byte chunks
	pdvs := FragmentData(1, true, data, 300)

	// Should create 4 PDVs: 300 + 300 + 300 + 100
	if len(pdvs) != 4 {
		t.Fatalf("Expected 4 PDVs, got %d", len(pdvs))
	}

	// Check sizes
	if len(pdvs[0].Data) != 300 {
		t.Errorf("PDV 0: expected 300 bytes, got %d", len(pdvs[0].Data))
	}
	if len(pdvs[3].Data) != 100 {
		t.Errorf("PDV 3: expected 100 bytes, got %d", len(pdvs[3].Data))
	}

	// Check last fragment flags
	for i := 0; i < 3; i++ {
		if pdvs[i].IsLastFragment {
			t.Errorf("PDV %d: should not be last fragment", i)
		}
	}
	if !pdvs[3].IsLastFragment {
		t.Error("PDV 3: should be last fragment")
	}

	// Check all are command type
	for i, pdv := range pdvs {
		if !pdv.IsCommand {
			t.Errorf("PDV %d: should be command type", i)
		}
		if pdv.PresentationContextID != 1 {
			t.Errorf("PDV %d: wrong context ID", i)
		}
	}

	// Verify data integrity
	var reassembled bytes.Buffer
	for _, pdv := range pdvs {
		reassembled.Write(pdv.Data)
	}
	if !bytes.Equal(reassembled.Bytes(), data) {
		t.Error("Reassembled data doesn't match original")
	}
}

func TestFragmentDataNoFragmentation(t *testing.T) {
	// Test with maxPDVSize = 0 (no fragmentation)
	data := []byte{0x01, 0x02, 0x03, 0x04}
	pdvs := FragmentData(1, false, data, 0)

	if len(pdvs) != 1 {
		t.Fatalf("Expected 1 PDV, got %d", len(pdvs))
	}

	if !pdvs[0].IsLastFragment {
		t.Error("Single PDV should be last fragment")
	}
	if pdvs[0].IsCommand {
		t.Error("PDV should be data type")
	}
	if !bytes.Equal(pdvs[0].Data, data) {
		t.Error("Data mismatch")
	}
}

func TestReassembleData(t *testing.T) {
	// Create fragmented data
	originalData := make([]byte, 1000)
	for i := range originalData {
		originalData[i] = byte(i % 256)
	}

	pdvs := FragmentData(1, true, originalData, 300)

	// Reassemble
	reassembled, err := ReassembleData(pdvs)
	if err != nil {
		t.Fatalf("ReassembleData failed: %v", err)
	}

	// Verify
	if !bytes.Equal(reassembled, originalData) {
		t.Error("Reassembled data doesn't match original")
	}
}

func TestReassembleDataNoPDVs(t *testing.T) {
	_, err := ReassembleData([]PDV{})
	if err == nil {
		t.Error("Expected error for empty PDV list")
	}
}

func TestReassembleDataIncomplete(t *testing.T) {
	// Create PDVs without last fragment
	pdvs := []PDV{
		NewPDV(1, true, false, []byte{0x01, 0x02}),
		NewPDV(1, true, false, []byte{0x03, 0x04}), // Missing last fragment flag
	}

	_, err := ReassembleData(pdvs)
	if err == nil {
		t.Error("Expected error for incomplete message")
	}
}

func TestReassembleDataMixedContexts(t *testing.T) {
	// Create PDVs with different context IDs
	pdvs := []PDV{
		NewPDV(1, true, false, []byte{0x01, 0x02}),
		NewPDV(3, true, true, []byte{0x03, 0x04}), // Different context ID
	}

	_, err := ReassembleData(pdvs)
	if err == nil {
		t.Error("Expected error for mixed context IDs")
	}
}

func TestReassembleDataMixedTypes(t *testing.T) {
	// Create PDVs with different types (command vs data)
	pdvs := []PDV{
		NewPDV(1, true, false, []byte{0x01, 0x02}),  // Command
		NewPDV(1, false, true, []byte{0x03, 0x04}),  // Data
	}

	_, err := ReassembleData(pdvs)
	if err == nil {
		t.Error("Expected error for mixed PDV types")
	}
}

func TestPDataTF_String(t *testing.T) {
	pdata := NewPDataTF()
	pdata.AddPDV(NewPDV(1, true, true, []byte{0x01}))
	pdata.AddPDV(NewPDV(1, false, true, []byte{0x02}))

	str := pdata.String()
	expected := "P-DATA-TF (2 PDVs)"

	if str != expected {
		t.Errorf("String() mismatch: expected %s, got %s", expected, str)
	}
}

func TestPDataTF_LargeMessage(t *testing.T) {
	// Test with a large message (10KB)
	largeData := make([]byte, 10000)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	// Fragment into 1KB chunks
	pdvs := FragmentData(1, false, largeData, 1024)

	// Create P-DATA-TF
	pdata := NewPDataTF()
	for _, pdv := range pdvs {
		pdata.AddPDV(pdv)
	}

	// Encode
	pdu, err := pdata.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// Decode
	decoded := &PDataTF{}
	if err := decoded.Decode(pdu); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Reassemble
	reassembled, err := ReassembleData(decoded.PDVs)
	if err != nil {
		t.Fatalf("ReassembleData failed: %v", err)
	}

	// Verify
	if !bytes.Equal(reassembled, largeData) {
		t.Error("Large message data integrity check failed")
	}
}
