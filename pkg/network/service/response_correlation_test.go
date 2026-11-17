// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// runCEchoWithMismatchedResponse sends a C-ECHO, lets respond answer it with a
// message that reuses the Message ID but violates the correlation rules, and
// asserts that the association is aborted and the caller receives the protocol
// error instead of a response. It returns the caller's error.
func runCEchoWithMismatchedResponse(t *testing.T, respond func(net.Conn, uint16) error) error {
	t.Helper()

	clientConn, peerConn := net.Pipe()
	t.Cleanup(func() {
		_ = clientConn.Close()
		_ = peerConn.Close()
	})

	service := acceptedService(t, clientConn, createTestAssociation(), WithAssociationRequestor(true))
	t.Cleanup(func() { _ = service.Close() })
	if err := service.Start(); err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	peerDone := make(chan error, 1)
	go func() {
		command, err := readNextDIMSECommand(peerConn, 2*time.Second)
		if err != nil {
			peerDone <- fmt.Errorf("read C-ECHO-RQ: %w", err)
			return
		}
		commandField, err := command.GetUInt16(tag.CommandField, 0)
		if err != nil {
			peerDone <- fmt.Errorf("read Command Field: %w", err)
			return
		}
		if dimse.CommandField(commandField) != dimse.CommandCEchoRQ {
			peerDone <- fmt.Errorf("received %s, want C-ECHO-RQ", dimse.CommandField(commandField))
			return
		}
		messageID, err := command.GetUInt16(tag.MessageID, 0)
		if err != nil {
			peerDone <- fmt.Errorf("read Message ID: %w", err)
			return
		}
		if err := respond(peerConn, messageID); err != nil {
			peerDone <- err
			return
		}

		rawAbort, err := transport.ReadPDU(peerConn, 2*time.Second)
		if err != nil {
			peerDone <- fmt.Errorf("read A-ABORT: %w", err)
			return
		}
		if rawAbort.Type != pdu.TypeAAbort {
			peerDone <- fmt.Errorf("PDU type = 0x%02X, want A-ABORT", rawAbort.Type)
			return
		}
		peerDone <- nil
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := service.SendCEcho(ctx, dimse.NewCEchoRequest())
	if resp != nil || err == nil {
		t.Fatalf("SendCEcho() = (%v, %v), want a protocol error and no response", resp, err)
	}
	if errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("SendCEcho() waited for a response the receive loop should have rejected: %v", err)
	}

	select {
	case err := <-peerDone:
		if err != nil {
			t.Fatalf("peer error = %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("peer did not observe the A-ABORT")
	}
	return err
}

// TestResponseWithWrongCommandFieldAbortsAssociation verifies that a response is
// correlated to its request by Message ID and the operation it answers: a
// pending C-ECHO must not accept a C-STORE-RSP that reuses its Message ID.
func TestResponseWithWrongCommandFieldAbortsAssociation(t *testing.T) {
	err := runCEchoWithMismatchedResponse(t, func(conn net.Conn, messageID uint16) error {
		resp := dimse.NewCStoreResponseSuccess(messageID, testCTImageStorageUID, "2.25.1")
		resp.SetPresentationContextID(1)
		return writeDIMSEPDU(conn, resp, 1, transfer.ExplicitVRLittleEndian)
	})
	if !strings.Contains(err.Error(), "Command Field 0x8030") {
		t.Fatalf("SendCEcho() error = %v, want the Command Field protocol error", err)
	}
}

// TestResponseOnWrongPresentationContextAbortsAssociation verifies the other
// half of PS3.7 9.3: the response must arrive on the presentation context the
// request was sent on, even when the Command Field matches.
func TestResponseOnWrongPresentationContextAbortsAssociation(t *testing.T) {
	err := runCEchoWithMismatchedResponse(t, func(conn net.Conn, messageID uint16) error {
		// Context 3 is accepted (CT Image Storage), but the C-ECHO was sent on
		// context 1 (Verification).
		resp := dimse.NewCEchoResponseSuccess(messageID)
		resp.SetPresentationContextID(3)
		return writeDIMSEPDU(conn, resp, 3, transfer.ExplicitVRLittleEndian)
	})
	if !strings.Contains(err.Error(), "presentation context 3, want 1") {
		t.Fatalf("SendCEcho() error = %v, want the presentation context protocol error", err)
	}
}
