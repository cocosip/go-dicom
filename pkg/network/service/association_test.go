// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

// mockConn is a mock net.Conn for testing
type mockConn struct {
	readData  []byte
	readPos   int
	writeData []byte
	closed    bool
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// serializeRawPDUForTest converts a RawPDU to complete PDU bytes (header + data) for test mocking.
func serializeRawPDUForTest(rawPDU *pdu.RawPDU) []byte {
	pduBytes := make([]byte, 6+len(rawPDU.Data))
	pduBytes[0] = rawPDU.Type
	pduBytes[1] = rawPDU.Reserved
	pduBytes[2] = byte(rawPDU.Length >> 24)
	pduBytes[3] = byte(rawPDU.Length >> 16)
	pduBytes[4] = byte(rawPDU.Length >> 8)
	pduBytes[5] = byte(rawPDU.Length)
	copy(pduBytes[6:], rawPDU.Data)
	return pduBytes
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	if m.readPos >= len(m.readData) {
		return 0, net.ErrClosed
	}
	n = copy(b, m.readData[m.readPos:])
	m.readPos += n
	return n, nil
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	m.writeData = append(m.writeData, b...)
	return len(b), nil
}

func (m *mockConn) Close() error {
	m.closed = true
	return nil
}

func (m *mockConn) LocalAddr() net.Addr                { return nil }
func (m *mockConn) RemoteAddr() net.Addr               { return nil }
func (m *mockConn) SetDeadline(t time.Time) error      { return nil }
func (m *mockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *mockConn) SetWriteDeadline(t time.Time) error { return nil }

func TestSendAssociationRequest(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil)
	defer service.Close()

	// Verify initial state is Idle
	if service.GetState() != StateIdle {
		t.Fatalf("Expected initial state Idle, got %s", service.GetState())
	}

	// Create A-ASSOCIATE-RQ
	rq := &pdu.AAssociateRQ{
		ProtocolVersion:    1,
		CalledAETitle:      "CALLED_AE",
		CallingAETitle:     "CALLING_AE",
		ApplicationContext: "1.2.840.10008.3.1.1.1",
	}

	ctx := context.Background()
	err := service.SendAssociationRequest(ctx, rq)
	if err != nil {
		t.Fatalf("SendAssociationRequest failed: %v", err)
	}

	// Verify state transitioned to AssociationRequested
	if service.GetState() != StateAssociationRequested {
		t.Errorf("Expected state AssociationRequested, got %s", service.GetState())
	}

	// Verify data was written
	if len(conn.writeData) == 0 {
		t.Error("Expected data to be written to connection")
	}

	// Verify PDU type is A-ASSOCIATE-RQ (0x01)
	if conn.writeData[0] != pdu.TypeAAssociateRQ {
		t.Errorf("Expected PDU type 0x01, got 0x%02X", conn.writeData[0])
	}
}

func TestSendAssociationRequest_WrongState(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil)
	defer service.Close()

	// Set state to something other than Idle
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	rq := &pdu.AAssociateRQ{
		ProtocolVersion: 1,
		CalledAETitle:   "CALLED_AE",
		CallingAETitle:  "CALLING_AE",
	}

	ctx := context.Background()
	err := service.SendAssociationRequest(ctx, rq)
	if err == nil {
		t.Error("Expected error when sending A-ASSOCIATE-RQ in wrong state")
	}
}

func TestSendAssociationAccept(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil)
	defer service.Close()

	// Create A-ASSOCIATE-AC with minimal required fields
	ac := pdu.NewAAssociateAC()
	ac.CalledAETitle = "CALLED_AE"
	ac.CallingAETitle = "CALLING_AE"

	ctx := context.Background()
	err := service.SendAssociationAccept(ctx, ac)
	if err != nil {
		t.Fatalf("SendAssociationAccept failed: %v", err)
	}

	// Verify state transitioned to AssociationAccepted
	if service.GetState() != StateAssociationAccepted {
		t.Errorf("Expected state AssociationAccepted, got %s", service.GetState())
	}

	// Verify data was written
	if len(conn.writeData) == 0 {
		t.Error("Expected data to be written to connection")
	}

	// Debug: print first 10 bytes
	t.Logf("Written data (first 10 bytes): % X", conn.writeData[:min(10, len(conn.writeData))])

	// Verify PDU type is A-ASSOCIATE-AC (0x02)
	if conn.writeData[0] != pdu.TypeAAssociateAC {
		t.Errorf("Expected PDU type 0x02, got 0x%02X", conn.writeData[0])
	}
}

func TestSendAssociationReject(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil)
	defer service.Close()

	ctx := context.Background()
	err := service.SendAssociationReject(ctx, 1, 1, 1)
	if err != nil {
		t.Fatalf("SendAssociationReject failed: %v", err)
	}

	// Verify state transitioned to Closed
	if service.GetState() != StateClosed {
		t.Errorf("Expected state Closed, got %s", service.GetState())
	}

	// Verify data was written
	if len(conn.writeData) == 0 {
		t.Error("Expected data to be written to connection")
	}

	// Verify PDU type is A-ASSOCIATE-RJ (0x03)
	if conn.writeData[0] != pdu.TypeAAssociateRJ {
		t.Errorf("Expected PDU type 0x03, got 0x%02X", conn.writeData[0])
	}
}

func TestSendReleaseRequest(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil)
	defer service.Close()

	// Set state to AssociationAccepted first
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	ctx := context.Background()
	err := service.SendReleaseRequest(ctx)
	if err != nil {
		t.Fatalf("SendReleaseRequest failed: %v", err)
	}

	// Verify state transitioned to ReleaseRequested
	if service.GetState() != StateReleaseRequested {
		t.Errorf("Expected state ReleaseRequested, got %s", service.GetState())
	}

	// Verify data was written
	if len(conn.writeData) == 0 {
		t.Error("Expected data to be written to connection")
	}

	// Verify PDU type is A-RELEASE-RQ (0x05)
	if conn.writeData[0] != pdu.TypeAReleaseRQ {
		t.Errorf("Expected PDU type 0x05, got 0x%02X", conn.writeData[0])
	}
}

func TestSendReleaseRequest_WrongState(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil)
	defer service.Close()

	// State is Idle by default, which is not valid for release
	ctx := context.Background()
	err := service.SendReleaseRequest(ctx)
	if err == nil {
		t.Error("Expected error when sending A-RELEASE-RQ in wrong state")
	}
}

func TestSendReleaseResponse(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil)
	defer service.Close()

	ctx := context.Background()
	err := service.SendReleaseResponse(ctx)
	if err != nil {
		t.Fatalf("SendReleaseResponse failed: %v", err)
	}

	// Verify state transitioned to Closed
	if service.GetState() != StateClosed {
		t.Errorf("Expected state Closed, got %s", service.GetState())
	}

	// Verify data was written
	if len(conn.writeData) == 0 {
		t.Error("Expected data to be written to connection")
	}

	// Verify PDU type is A-RELEASE-RP (0x06)
	if conn.writeData[0] != pdu.TypeAReleaseRP {
		t.Errorf("Expected PDU type 0x06, got 0x%02X", conn.writeData[0])
	}
}

func TestSendAbort(t *testing.T) {
	conn := &mockConn{}
	service := NewService(conn, nil)
	defer service.Close()

	ctx := context.Background()
	err := service.SendAbort(ctx, 0, 0)
	if err != nil {
		t.Fatalf("SendAbort failed: %v", err)
	}

	// Verify state transitioned to Aborted
	if service.GetState() != StateAborted {
		t.Errorf("Expected state Aborted, got %s", service.GetState())
	}

	// Verify data was written
	if len(conn.writeData) == 0 {
		t.Error("Expected data to be written to connection")
	}

	// Verify PDU type is A-ABORT (0x07)
	if conn.writeData[0] != pdu.TypeAAbort {
		t.Errorf("Expected PDU type 0x07, got 0x%02X", conn.writeData[0])
	}
}

func TestReceiveAssociationResponse_Accept(t *testing.T) {
	// Create a valid A-ASSOCIATE-AC PDU
	ac := pdu.NewAAssociateAC()
	ac.CalledAETitle = "CALLED_AE"
	ac.CallingAETitle = "CALLING_AE"

	rawPDU, err := ac.Encode()
	if err != nil {
		t.Fatalf("Failed to encode A-ASSOCIATE-AC: %v", err)
	}

	// Build complete PDU with header for reading
	pduBytes := serializeRawPDUForTest(rawPDU)

	conn := &mockConn{
		readData: pduBytes,
	}

	service := NewService(conn, nil)
	defer service.Close()

	// Set state to AssociationRequested
	if err := service.setState(StateAssociationRequested); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	ctx := context.Background()
	receivedAC, err := service.ReceiveAssociationResponse(ctx)
	if err != nil {
		t.Fatalf("ReceiveAssociationResponse failed: %v", err)
	}

	if receivedAC == nil {
		t.Fatal("Expected A-ASSOCIATE-AC, got nil")
	}

	// Verify state transitioned to AssociationAccepted
	if service.GetState() != StateAssociationAccepted {
		t.Errorf("Expected state AssociationAccepted, got %s", service.GetState())
	}

	// Verify AC content
	if receivedAC.CalledAETitle != ac.CalledAETitle {
		t.Errorf("Expected CalledAETitle %s, got %s", ac.CalledAETitle, receivedAC.CalledAETitle)
	}
}

func TestReceiveAssociationResponse_Reject(t *testing.T) {
	// Create a valid A-ASSOCIATE-RJ PDU
	rj := &pdu.AAssociateRJ{
		Result: 1,
		Source: 1,
		Reason: 1,
	}

	rawPDU, err := rj.Encode()
	if err != nil {
		t.Fatalf("Failed to encode A-ASSOCIATE-RJ: %v", err)
	}

	// Build complete PDU with header for reading
	pduBytes := serializeRawPDUForTest(rawPDU)

	conn := &mockConn{
		readData: pduBytes,
	}

	service := NewService(conn, nil)
	defer service.Close()

	// Set state to AssociationRequested
	if err := service.setState(StateAssociationRequested); err != nil {
		t.Fatalf("Failed to set state: %v", err)
	}

	ctx := context.Background()
	receivedAC, err := service.ReceiveAssociationResponse(ctx)
	if err == nil {
		t.Error("Expected error when receiving A-ASSOCIATE-RJ")
	}

	if receivedAC != nil {
		t.Error("Expected nil A-ASSOCIATE-AC when rejected")
	}

	// Verify state transitioned to Closed
	if service.GetState() != StateClosed {
		t.Errorf("Expected state Closed, got %s", service.GetState())
	}
}

func TestReceiveAssociationRequest(t *testing.T) {
	// Create a valid A-ASSOCIATE-RQ PDU
	rq := pdu.NewAAssociateRQ()
	rq.CalledAETitle = "CALLED_AE"
	rq.CallingAETitle = "CALLING_AE"

	rawPDU, err := rq.Encode()
	if err != nil {
		t.Fatalf("Failed to encode A-ASSOCIATE-RQ: %v", err)
	}

	// Build complete PDU with header for reading
	pduBytes := serializeRawPDUForTest(rawPDU)

	conn := &mockConn{
		readData: pduBytes,
	}

	service := NewService(conn, nil)
	defer service.Close()

	// State should be Idle (default)
	ctx := context.Background()
	receivedRQ, err := service.ReceiveAssociationRequest(ctx)
	if err != nil {
		t.Fatalf("ReceiveAssociationRequest failed: %v", err)
	}

	if receivedRQ == nil {
		t.Fatal("Expected A-ASSOCIATE-RQ, got nil")
	}

	// Verify RQ content
	if receivedRQ.CalledAETitle != rq.CalledAETitle {
		t.Errorf("Expected CalledAETitle %s, got %s", rq.CalledAETitle, receivedRQ.CalledAETitle)
	}
	if receivedRQ.CallingAETitle != rq.CallingAETitle {
		t.Errorf("Expected CallingAETitle %s, got %s", rq.CallingAETitle, receivedRQ.CallingAETitle)
	}
}

func TestDeadlineFromContext(t *testing.T) {
	tests := []struct {
		name           string
		ctx            context.Context
		timeout        time.Duration
		expectDeadline bool
	}{
		{
			name:           "No timeout",
			ctx:            context.Background(),
			timeout:        0,
			expectDeadline: false,
		},
		{
			name:           "With timeout, no context deadline",
			ctx:            context.Background(),
			timeout:        1 * time.Second,
			expectDeadline: true,
		},
		{
			name: "Context deadline earlier than timeout",
			ctx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
				defer cancel()
				return ctx
			}(),
			timeout:        1 * time.Second,
			expectDeadline: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deadline := deadlineFromContext(tt.ctx, tt.timeout)
			if tt.expectDeadline && deadline.IsZero() {
				t.Error("Expected deadline to be set")
			}
			if !tt.expectDeadline && !deadline.IsZero() {
				t.Error("Expected no deadline")
			}
		})
	}
}
