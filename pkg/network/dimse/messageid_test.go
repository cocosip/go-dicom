// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import (
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestMessageIDGenerator_Next(t *testing.T) {
	gen := NewMessageIDGenerator()

	// Generate first few IDs
	id1 := gen.Next()
	if id1 != 1 {
		t.Errorf("Expected first ID to be 1, got %d", id1)
	}

	id2 := gen.Next()
	if id2 != 2 {
		t.Errorf("Expected second ID to be 2, got %d", id2)
	}

	id3 := gen.Next()
	if id3 != 3 {
		t.Errorf("Expected third ID to be 3, got %d", id3)
	}
}

func TestMessageIDGenerator_Wraparound(t *testing.T) {
	gen := NewMessageIDGenerator()

	// Set counter near overflow
	gen.counter = 0xFFFE

	id1 := gen.Next()
	if id1 != 0xFFFF {
		t.Errorf("Expected ID 0xFFFF, got 0x%04X", id1)
	}

	// Should wrap to 1 (skipping 0)
	id2 := gen.Next()
	if id2 != 1 {
		t.Errorf("Expected ID to wrap to 1, got %d", id2)
	}
}

func TestMessageIDGenerator_Reset(t *testing.T) {
	gen := NewMessageIDGenerator()

	// Generate some IDs
	gen.Next()
	gen.Next()
	gen.Next()

	// Reset
	gen.Reset()

	// Should start from 1 again
	id := gen.Next()
	if id != 1 {
		t.Errorf("Expected ID 1 after reset, got %d", id)
	}
}

func TestMessageIDGenerator_Concurrent(t *testing.T) {
	gen := NewMessageIDGenerator()
	numGoroutines := 100
	idsPerGoroutine := 100

	var wg sync.WaitGroup
	ids := make(chan uint16, numGoroutines*idsPerGoroutine)

	// Launch concurrent goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < idsPerGoroutine; j++ {
				ids <- gen.Next()
			}
		}()
	}

	wg.Wait()
	close(ids)

	// Collect all IDs
	seen := make(map[uint16]bool)
	for id := range ids {
		if seen[id] {
			t.Errorf("Duplicate ID generated: %d", id)
		}
		seen[id] = true
	}

	// Should have generated unique IDs
	if len(seen) != numGoroutines*idsPerGoroutine {
		t.Errorf("Expected %d unique IDs, got %d", numGoroutines*idsPerGoroutine, len(seen))
	}
}

func TestMessageIDGenerator_AssignMessageID(t *testing.T) {
	gen := NewMessageIDGenerator()

	// Test with C-ECHO request (messageID = 0)
	req := NewCEchoRequest()

	msgID, err := gen.AssignMessageID(req)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}

	if msgID != 1 {
		t.Errorf("Expected MessageID 1, got %d", msgID)
	}

	if req.MessageID() != 1 {
		t.Errorf("Request MessageID not updated, got %d", req.MessageID())
	}
}

func TestMessageIDGenerator_AssignMessageID_AlreadySet(t *testing.T) {
	gen := NewMessageIDGenerator()

	// Test with C-ECHO request and manually set MessageID
	req := NewCEchoRequest()
	req.SetMessageID(999) // Manually set MessageID

	msgID, err := gen.AssignMessageID(req)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}

	// Should keep the original MessageID
	if msgID != 999 {
		t.Errorf("Expected MessageID 999, got %d", msgID)
	}

	if req.MessageID() != 999 {
		t.Errorf("Request MessageID changed, got %d", req.MessageID())
	}
}

func TestMessageIDGenerator_AssignMessageID_CStoreRequest(t *testing.T) {
	gen := NewMessageIDGenerator()

	// Create a C-STORE request with messageID = 0
	ds := createTestDataset()
	req, err := NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("Failed to create C-STORE request: %v", err)
	}

	msgID, err := gen.AssignMessageID(req)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}

	if msgID != 1 {
		t.Errorf("Expected MessageID 1, got %d", msgID)
	}

	if req.MessageID() != 1 {
		t.Errorf("Request MessageID not updated, got %d", req.MessageID())
	}
}

func TestMessageIDGenerator_AssignMessageID_CFindRequest(t *testing.T) {
	gen := NewMessageIDGenerator()

	// Create a C-FIND request with messageID = 0
	query := createTestDataset()
	req := NewCFindRequestPatientRoot(QueryRetrieveLevelPatient, query)

	msgID, err := gen.AssignMessageID(req)
	if err != nil {
		t.Fatalf("Failed to assign MessageID: %v", err)
	}

	if msgID != 1 {
		t.Errorf("Expected MessageID 1, got %d", msgID)
	}

	if req.MessageID() != 1 {
		t.Errorf("Request MessageID not updated, got %d", req.MessageID())
	}
}

// Helper function to create a test dataset
func createTestDataset() *dataset.Dataset {
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))
	return ds
}
