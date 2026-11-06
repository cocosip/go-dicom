// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"bytes"
	"testing"
)

// TestHandlePDataTFEnforcesConfiguredMaxMessageSize verifies the optional
// per-message reassembly bound rejects an oversized fragmented message.
func TestHandlePDataTFEnforcesConfiguredMaxMessageSize(t *testing.T) {
	svc := NewService(nil, createTestAssociation(), WithMaxMessageSize(8))
	defer func() { _ = svc.Close() }()

	raw, err := CreatePDataTFPDU([]*PDV{{
		PresentationContextID: 1,
		IsCommand:             true,
		IsLastFragment:        false,
		Data:                  bytes.Repeat([]byte{0}, 32),
	}})
	if err != nil {
		t.Fatalf("CreatePDataTFPDU() error = %v", err)
	}

	if err := svc.handlePDataTF(raw, map[byte]*dimseMessageFragments{}); err == nil {
		t.Fatal("handlePDataTF() error = nil, want reassembly size error")
	}
}

// TestHandlePDataTFUnlimitedByDefault verifies that a large fragmented message is
// not rejected when WithMaxMessageSize is unset, so big instances keep working.
func TestHandlePDataTFUnlimitedByDefault(t *testing.T) {
	svc := NewService(nil, createTestAssociation())
	defer func() { _ = svc.Close() }()

	raw, err := CreatePDataTFPDU([]*PDV{{
		PresentationContextID: 1,
		IsCommand:             true,
		IsLastFragment:        false,
		Data:                  bytes.Repeat([]byte{0}, 32),
	}})
	if err != nil {
		t.Fatalf("CreatePDataTFPDU() error = %v", err)
	}

	// A single non-final command PDV is an incomplete message, not an error.
	if err := svc.handlePDataTF(raw, map[byte]*dimseMessageFragments{}); err != nil {
		t.Fatalf("handlePDataTF() error = %v, want nil (unlimited by default)", err)
	}
}
