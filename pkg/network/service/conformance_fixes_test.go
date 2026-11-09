// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

func captureSentMessages(service *Service) <-chan dimse.Message {
	sent := make(chan dimse.Message, 16)
	go func() {
		for {
			select {
			case request := <-service.sendQueue:
				sent <- request.message
				request.resultCh <- nil
			case <-service.closeCh:
				return
			}
		}
	}()
	return sent
}

func waitForSentMessage(t *testing.T, sent <-chan dimse.Message, match func(dimse.Message) bool) dimse.Message {
	t.Helper()
	deadline := time.After(2 * time.Second)
	for {
		select {
		case message := <-sent:
			if match(message) {
				return message
			}
		case <-deadline:
			t.Fatal("timed out waiting for the expected DIMSE message")
		}
	}
}

func expectNoSentMessage(t *testing.T, sent <-chan dimse.Message, window time.Duration) {
	t.Helper()
	select {
	case message := <-sent:
		t.Fatalf("unexpected extra DIMSE message %T after C-CANCEL", message)
	case <-time.After(window):
	}
}

func TestRequestTimeoutSendsCCancel(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(30*time.Millisecond))
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	sent := captureSentMessages(service)

	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	request.SetPresentationContextID(1)
	events, err := service.SendCFind(context.Background(), request)
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	waitForSentMessage(t, sent, func(message dimse.Message) bool {
		_, ok := message.(*dimse.CFindRequest)
		return ok
	})

	cancel := waitForSentMessage(t, sent, func(message dimse.Message) bool {
		_, ok := message.(*dimse.CCancelRequest)
		return ok
	}).(*dimse.CCancelRequest)
	if cancel.MessageIDBeingRespondedTo() != request.MessageID() {
		t.Fatalf("C-CANCEL MessageIDBeingRespondedTo = %d, want %d", cancel.MessageIDBeingRespondedTo(), request.MessageID())
	}
	if cancel.PresentationContextID() != request.PresentationContextID() || cancel.PresentationContextID() == 0 {
		t.Fatalf("C-CANCEL presentation context ID = %d, want %d", cancel.PresentationContextID(), request.PresentationContextID())
	}
	// The timeout and a context cancellation can race; the request must still
	// produce exactly one C-CANCEL.
	expectNoSentMessage(t, sent, 100*time.Millisecond)

	event := <-events
	if !errors.Is(event.Err, ErrRequestTimeout) {
		t.Fatalf("terminal event = %#v, want ErrRequestTimeout", event)
	}
}

func TestContextCancellationSendsCCancel(t *testing.T) {
	service := newAcceptedTestService(t)
	sent := captureSentMessages(service)

	ctx, cancel := context.WithCancel(context.Background())
	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	request.SetPresentationContextID(1)
	events, err := service.SendCFind(ctx, request)
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	waitForSentMessage(t, sent, func(message dimse.Message) bool {
		_, ok := message.(*dimse.CFindRequest)
		return ok
	})

	cancel()
	waitForSentMessage(t, sent, func(message dimse.Message) bool {
		_, ok := message.(*dimse.CCancelRequest)
		return ok
	})
	expectNoSentMessage(t, sent, 100*time.Millisecond)

	event := <-events
	if !errors.Is(event.Err, context.Canceled) {
		t.Fatalf("terminal event = %#v, want context.Canceled", event)
	}
}

func TestNonCancellableRequestTimeoutDoesNotSendCCancel(t *testing.T) {
	service := NewService(nil, createTestAssociation(), WithRequestTimeout(20*time.Millisecond))
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	sent := captureSentMessages(service)

	_, err := service.SendCEcho(context.Background(), dimse.NewCEchoRequest())
	if !errors.Is(err, ErrRequestTimeout) {
		t.Fatalf("SendCEcho() error = %v, want ErrRequestTimeout", err)
	}

	// Drain the C-ECHO itself first; a C-CANCEL sent later would otherwise never
	// be observed by a single select.
	waitForSentMessage(t, sent, func(message dimse.Message) bool {
		_, ok := message.(*dimse.CEchoRequest)
		return ok
	})
	expectNoSentMessage(t, sent, 100*time.Millisecond)
}

func TestSendRejectsEmptyRequiredSOPInstanceUID(t *testing.T) {
	service := newAcceptedTestService(t)
	// Validation runs in sendMessage, so drive the real encode path instead of
	// a mock that acknowledges every message.
	go func() {
		for {
			select {
			case request := <-service.sendQueue:
				request.resultCh <- service.sendMessage(request)
			case <-service.closeCh:
				return
			}
		}
	}()

	_, err := service.SendNGet(context.Background(), dimse.NewNGetRequest("1.2.840.10008.5.1.1.1", "", nil))
	if err == nil {
		t.Fatal("SendNGet() = nil, want a validation error for an empty SOP Instance UID")
	}
	if !strings.Contains(err.Error(), "RequestedSOPInstanceUID is required") {
		t.Fatalf("SendNGet() error = %v, want a required-field error", err)
	}
}

func TestPeerAbortLeavesAssociationAborted(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer func() { _ = serverConn.Close() }()

	service := NewService(clientConn, createTestAssociation())
	defer func() { _ = service.Close() }()
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	if err := service.Start(); err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	abort := pdu.NewAAbort(pdu.AbortSourceServiceUser, pdu.AbortReasonServiceUserNotSpecified)
	raw, err := abort.Encode()
	if err != nil {
		t.Fatalf("AAbort.Encode() error = %v", err)
	}
	if err := transport.WritePDU(serverConn, time.Second, raw); err != nil {
		t.Fatalf("WritePDU() error = %v", err)
	}

	select {
	case <-service.closeCh:
	case <-time.After(2 * time.Second):
		t.Fatal("service did not close after A-ABORT")
	}
	if state := service.GetState(); state != StateAborted {
		t.Fatalf("state = %s, want %s (peer A-ABORT must not be reported as Closed)", state, StateAborted)
	}
	if err := service.CloseError(); err == nil || !strings.Contains(err.Error(), "A-ABORT") {
		t.Fatalf("CloseError() = %v, want the received A-ABORT recorded", err)
	}
}
