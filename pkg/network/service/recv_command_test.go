// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"bytes"
	"encoding/binary"
	"strings"
	"testing"
)

func TestCommandHasDataset(t *testing.T) {
	t.Run("no dataset present", func(t *testing.T) {
		var command bytes.Buffer
		if err := binary.Write(&command, binary.LittleEndian, uint16(0x0000)); err != nil {
			t.Fatalf("failed to write group: %v", err)
		}
		if err := binary.Write(&command, binary.LittleEndian, uint16(0x0800)); err != nil {
			t.Fatalf("failed to write element: %v", err)
		}
		if err := binary.Write(&command, binary.LittleEndian, uint32(2)); err != nil {
			t.Fatalf("failed to write length: %v", err)
		}
		if err := binary.Write(&command, binary.LittleEndian, uint16(0x0101)); err != nil {
			t.Fatalf("failed to write CommandDataSetType: %v", err)
		}

		hasData := new(Service).commandHasDataset(command.Bytes())
		if hasData {
			t.Fatal("expected no dataset to be present")
		}
	})

	t.Run("truncated command data returns error", func(t *testing.T) {
		hasData := new(Service).commandHasDataset([]byte{0x00, 0x00, 0x00})
		if hasData {
			t.Fatal("expected malformed command data to be treated as not having a dataset")
		}
	})
}

func TestHandlePDataTFReturnsErrorForMalformedCommandFragments(t *testing.T) {
	rawPDU, err := CreatePDataTFPDU([]*PDV{
		{
			PresentationContextID: 1,
			IsCommand:             true,
			IsLastFragment:        true,
			Data:                  []byte{0x00, 0x00, 0x00},
		},
	})
	if err != nil {
		t.Fatalf("CreatePDataTFPDU failed: %v", err)
	}

	var commandFragments []byte
	var datasetFragments []byte
	var currentContextID byte

	err = new(Service).handlePDataTF(rawPDU, &commandFragments, &datasetFragments, &currentContextID)
	if err == nil {
		t.Fatal("expected malformed command fragments to return an error")
	}
	if !strings.Contains(err.Error(), "CommandDataSetType") {
		t.Fatalf("expected CommandDataSetType error, got %v", err)
	}
}
