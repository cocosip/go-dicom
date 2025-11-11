// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

func TestEncodeDIMSEMessage(t *testing.T) {
	// Create a C-ECHO request
	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Encode the message
	commandData, datasetData, err := EncodeDIMSEMessage(req, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("EncodeDIMSEMessage failed: %v", err)
	}

	// Command data should not be empty
	if len(commandData) == 0 {
		t.Error("Command data is empty")
	}

	// C-ECHO should not have dataset
	if datasetData != nil {
		t.Error("C-ECHO should not have dataset")
	}
}

func TestEncodeDIMSEMessageWithDataset(t *testing.T) {
	t.Skip("Skipping test - CreateTestDataset helper not yet implemented")
	// TODO: Implement this test once we have a test dataset creation helper
	// Create a C-STORE request with dataset
	// ds := dimse.CreateTestDataset() // Helper to create test dataset
	// req, err := dimse.NewCStoreRequest(ds)
	// if err != nil {
	// 	t.Skip("Skipping test - NewCStoreRequest requires valid dataset")
	// 	return
	// }
	//
	// req.SetMessageID(1)
	//
	// // Encode the message
	// commandData, datasetData, err := EncodeDIMSEMessage(req, transfer.ExplicitVRLittleEndian)
	// if err != nil {
	// 	t.Fatalf("EncodeDIMSEMessage failed: %v", err)
	// }
	//
	// // Both command and dataset should be present
	// if len(commandData) == 0 {
	// 	t.Error("Command data is empty")
	// }
	//
	// if len(datasetData) == 0 {
	// 	t.Error("Dataset data is empty for C-STORE")
	// }
}

func TestServiceStart(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service
	service := NewService(client, nil)
	defer func() { _ = service.Close() }()

	// Start service
	err := service.Start()
	if err != nil {
		t.Fatalf("Start failed: %v", err)
	}

	// Give goroutines time to start
	time.Sleep(50 * time.Millisecond)

	// Close service
	_ = service.Close()

	// Give goroutines time to stop
	time.Sleep(50 * time.Millisecond)
}

func TestSendMessage(t *testing.T) {
	// Create a pipe connection
    server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// Create an association with a presentation context for C-ECHO
	assoc := createTestAssociation()

	// Create service with association
    service := NewService(client, assoc)
    defer func() { _ = service.Close() }()

	// Create a send request
    req := dimse.NewCEchoRequest()
    if err := req.SetMessageID(1); err != nil {
        t.Fatalf("SetMessageID failed: %v", err)
    }

	sendReq := &sendRequest{
		message:  req,
		resultCh: make(chan error, 1),
	}

	// Start a goroutine to read from server side
	go func() {
		// Just read and discard data to prevent blocking
		buf := make([]byte, 4096)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	// Send message
	err := service.sendMessage(sendReq)
	if err != nil {
		t.Fatalf("sendMessage failed: %v", err)
	}
}

// createTestAssociation creates a test association with common presentation contexts
func createTestAssociation() *association.Association {
	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")

	// Add presentation context for C-ECHO (Verification SOP Class)
	pc := association.NewPresentationContext(1, "1.2.840.10008.1.1", transfer.ExplicitVRLittleEndian)
	pc.AcceptedTransferSyntax = transfer.ExplicitVRLittleEndian
	pc.Result = association.ResultAcceptance
	_ = assoc.AddPresentationContext(pc)

	// Add presentation context for CT Image Storage (for C-STORE tests)
	pc2 := association.NewPresentationContext(3, "1.2.840.10008.5.1.4.1.1.2", transfer.ExplicitVRLittleEndian)
	pc2.AcceptedTransferSyntax = transfer.ExplicitVRLittleEndian
	pc2.Result = association.ResultAcceptance
	_ = assoc.AddPresentationContext(pc2)

	assoc.SetEstablished(true)
	return assoc
}

func TestSendLoop_ContextCancellation(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service
    service := NewService(client, nil)
    defer func() { _ = service.Close() }()

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start send loop in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.sendLoop(ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

	// Cancel context
	cancel()

	// Wait for loop to exit
	select {
	case err := <-errCh:
		if err != context.Canceled {
			t.Errorf("Expected context.Canceled, got %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("sendLoop did not exit after context cancellation")
	}
}

func TestGroupPDVsIntoPDUs(t *testing.T) {
	assoc := createTestAssociation()
	service := NewService(nil, assoc, WithMaxPDULength(1000))

	tests := []struct {
		name           string
		pdvDataSizes   []int
		maxPDULength   uint32
		expectedGroups int
	}{
		{
			name:           "Single small PDV",
			pdvDataSizes:   []int{100},
			maxPDULength:   1000,
			expectedGroups: 1,
		},
		{
			name:           "Multiple small PDVs fit in one PDU",
			pdvDataSizes:   []int{100, 100, 100},
			maxPDULength:   1000,
			expectedGroups: 1,
		},
		{
			name:           "PDVs require multiple PDUs",
			pdvDataSizes:   []int{400, 400, 400},
			maxPDULength:   1000,
			expectedGroups: 2, // First: 6 (PDU) + 406 (PDV1) + 406 (PDV2) = 818; Second: 6 + 406 = 412
		},
		{
			name:           "Some PDVs can be packed",
			pdvDataSizes:   []int{300, 300, 100, 100},
			maxPDULength:   1000,
			expectedGroups: 1, // All fit: 6 (PDU) + 306*2 + 106*2 = 830
		},
		{
			name:           "Empty PDV list",
			pdvDataSizes:   []int{},
			maxPDULength:   1000,
			expectedGroups: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create PDVs with specified data sizes
			var pdvs []*PDV
			for _, size := range tt.pdvDataSizes {
				pdv := &PDV{
					PresentationContextID: 1,
					IsCommand:             false,
					IsLastFragment:        false,
					Data:                  make([]byte, size),
				}
				pdvs = append(pdvs, pdv)
			}

			// Update service maxPDULength for this test
			service.config.maxPDULength = tt.maxPDULength

			// Group PDVs
			groups := service.groupPDVsIntoPDUs(pdvs)

			if len(groups) != tt.expectedGroups {
				t.Errorf("Expected %d groups, got %d", tt.expectedGroups, len(groups))
			}

			// Verify each group doesn't exceed max PDU length
			const pduHeaderSize = 6
			for i, group := range groups {
				totalSize := pduHeaderSize
				for _, pdv := range group {
					pdvSize := 4 + 1 + 1 + len(pdv.Data) // length + context ID + header + data
					totalSize += pdvSize
				}
				if totalSize > int(tt.maxPDULength) {
					t.Errorf("Group %d exceeds max PDU length: %d > %d", i, totalSize, tt.maxPDULength)
				}
			}

			// Verify all PDVs are included
			totalPDVs := 0
			for _, group := range groups {
				totalPDVs += len(group)
			}
			if totalPDVs != len(pdvs) {
				t.Errorf("Expected %d PDVs total, got %d", len(pdvs), totalPDVs)
			}
		})
	}
}

func TestSendLoop_ServiceClose(t *testing.T) {
	// Create a pipe connection
    server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// Create service
	service := NewService(client, nil)

	// Start send loop in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.sendLoop(service.ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

    // Close service
    _ = service.Close()

	// Wait for loop to exit
	select {
	case err := <-errCh:
		if err != nil && err != context.Canceled {
			t.Errorf("sendLoop returned unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("sendLoop did not exit after service close")
	}
}
