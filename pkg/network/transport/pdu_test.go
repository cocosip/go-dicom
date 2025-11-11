// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transport

import (
	"bytes"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

func TestReadWritePDU_Basic(t *testing.T) {
	// Create a pair of connected pipes
	server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// Test data
	testPDU := &pdu.RawPDU{
		Type: pdu.TypeAAssociateRQ,
		Data: []byte("test data"),
	}

	// Write PDU from client
	go func() {
		if err := WritePDU(client, 0, testPDU); err != nil {
			t.Errorf("WritePDU failed: %v", err)
		}
	}()

	// Read PDU on server
	receivedPDU, err := ReadPDU(server, 0)
	if err != nil {
		t.Fatalf("ReadPDU failed: %v", err)
	}

	// Verify
	if receivedPDU.Type != testPDU.Type {
		t.Errorf("Expected PDU type %d, got %d", testPDU.Type, receivedPDU.Type)
	}
	if !bytes.Equal(receivedPDU.Data, testPDU.Data) {
		t.Errorf("Expected data %v, got %v", testPDU.Data, receivedPDU.Data)
	}
}

func TestReadWritePDU_EmptyData(t *testing.T) {
    server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// PDU with empty data
	testPDU := &pdu.RawPDU{
		Type: pdu.TypeAAssociateAC,
		Data: []byte{},
	}

	go func() {
		if err := WritePDU(client, 0, testPDU); err != nil {
			t.Errorf("WritePDU failed: %v", err)
		}
	}()

	receivedPDU, err := ReadPDU(server, 0)
	if err != nil {
		t.Fatalf("ReadPDU failed: %v", err)
	}

	if receivedPDU.Type != testPDU.Type {
		t.Errorf("Expected PDU type %d, got %d", testPDU.Type, receivedPDU.Type)
	}
	if len(receivedPDU.Data) != 0 {
		t.Errorf("Expected empty data, got %d bytes", len(receivedPDU.Data))
	}
}

func TestReadWritePDU_LargeData(t *testing.T) {
    server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// Large PDU (1 MB)
	largeData := make([]byte, 1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	testPDU := &pdu.RawPDU{
		Type: pdu.TypePDataTF,
		Data: largeData,
	}

	go func() {
		if err := WritePDU(client, 5*time.Second, testPDU); err != nil {
			t.Errorf("WritePDU failed: %v", err)
		}
	}()

	receivedPDU, err := ReadPDU(server, 5*time.Second)
	if err != nil {
		t.Fatalf("ReadPDU failed: %v", err)
	}

	if receivedPDU.Type != testPDU.Type {
		t.Errorf("Expected PDU type %d, got %d", testPDU.Type, receivedPDU.Type)
	}
	if !bytes.Equal(receivedPDU.Data, testPDU.Data) {
		t.Errorf("Large data mismatch")
	}
}

func TestReadPDU_Timeout(t *testing.T) {
    server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// Don't write anything, just try to read with timeout
	errChan := make(chan error, 1)
	go func() {
		_, err := ReadPDU(server, 100*time.Millisecond)
		errChan <- err
	}()

	// Should timeout
	select {
	case err := <-errChan:
		if err == nil {
			t.Fatal("Expected timeout error, got nil")
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("Test itself timed out")
	}
}

func TestWritePDU_Timeout(t *testing.T) {
    server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// Fill up the pipe buffer to cause write to block
	largeData := make([]byte, 10*1024*1024) // 10 MB

	testPDU := &pdu.RawPDU{
		Type: pdu.TypePDataTF,
		Data: largeData,
	}

	// Don't read on the other end, so write should eventually block
	errChan := make(chan error, 1)
	go func() {
		errChan <- WritePDU(client, 100*time.Millisecond, testPDU)
	}()

	// Should timeout
	select {
	case err := <-errChan:
		if err == nil {
			t.Fatal("Expected timeout error, got nil")
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("Test itself timed out")
	}
}

func TestReadWritePDU_MultipleTypes(t *testing.T) {
    server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// Test different PDU types
	testPDUs := []*pdu.RawPDU{
		{Type: pdu.TypeAAssociateRQ, Data: []byte("associate request")},
		{Type: pdu.TypeAAssociateAC, Data: []byte("associate accept")},
		{Type: pdu.TypeAAssociateRJ, Data: []byte("associate reject")},
		{Type: pdu.TypePDataTF, Data: []byte("data transfer")},
		{Type: pdu.TypeAReleaseRQ, Data: []byte("release request")},
		{Type: pdu.TypeAReleaseRP, Data: []byte("release response")},
		{Type: pdu.TypeAAbort, Data: []byte("abort")},
	}

	// Write all PDUs
	go func() {
		for _, p := range testPDUs {
			if err := WritePDU(client, 0, p); err != nil {
				t.Errorf("WritePDU failed for type %d: %v", p.Type, err)
				return
			}
		}
	}()

	// Read all PDUs
	for i, expected := range testPDUs {
		received, err := ReadPDU(server, 0)
		if err != nil {
			t.Fatalf("ReadPDU %d failed: %v", i, err)
		}

		if received.Type != expected.Type {
			t.Errorf("PDU %d: expected type %d, got %d", i, expected.Type, received.Type)
		}
		if !bytes.Equal(received.Data, expected.Data) {
			t.Errorf("PDU %d: data mismatch", i)
		}
	}
}

func TestReadPDU_ExcessiveLength(t *testing.T) {
    server, client := net.Pipe()
    defer func() { _ = server.Close() }()
    defer func() { _ = client.Close() }()

	// Manually write a PDU with excessive length
	go func() {
		// Write PDU header with length > 100MB
        header := make([]byte, 6)
        header[0] = pdu.TypePDataTF
        header[1] = 0
		// Set length to 200MB (exceeds limit)
		length := uint32(200 * 1024 * 1024)
		header[2] = byte(length >> 24)
		header[3] = byte(length >> 16)
		header[4] = byte(length >> 8)
		header[5] = byte(length)
        _, _ = client.Write(header)
	}()

	// Should fail with length validation error
	_, err := ReadPDU(server, 0)
	if err == nil {
		t.Fatal("Expected error for excessive PDU length, got nil")
	}
}
