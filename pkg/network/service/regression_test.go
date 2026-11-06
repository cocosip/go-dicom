// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// TestEncodeDIMSEMessageDerivesCommandDataSetType verifies that the command
// set flag always matches the bytes actually sent (PS3.7 6.3).
func TestEncodeDIMSEMessageDerivesCommandDataSetType(t *testing.T) {
	t.Run("no dataset", func(t *testing.T) {
		req := dimse.NewCEchoRequest()
		if err := req.SetMessageID(1); err != nil {
			t.Fatalf("SetMessageID() error = %v", err)
		}
		commandData, datasetData, err := EncodeDIMSEMessage(req, nil)
		if err != nil {
			t.Fatalf("EncodeDIMSEMessage() error = %v", err)
		}
		if len(datasetData) != 0 {
			t.Fatalf("datasetData = %d bytes, want none", len(datasetData))
		}
		if got := commandDataSetType(t, commandData); got != 0x0101 {
			t.Fatalf("CommandDataSetType = 0x%04X, want 0x0101", got)
		}
	})

	t.Run("empty dataset is treated as absent", func(t *testing.T) {
		// A pending C-FIND-RSP may carry an (empty) identifier; encoding it
		// produces no bytes, so the flag must say "no dataset".
		resp := dimse.NewCFindResponse(2, status.CFindPending, verificationSOPClassUID, dataset.New())
		commandData, datasetData, err := EncodeDIMSEMessage(resp, nil)
		if err != nil {
			t.Fatalf("EncodeDIMSEMessage() error = %v", err)
		}
		if len(datasetData) != 0 {
			t.Fatalf("datasetData = %d bytes, want none", len(datasetData))
		}
		if got := commandDataSetType(t, commandData); got != 0x0101 {
			t.Fatalf("CommandDataSetType = 0x%04X, want 0x0101", got)
		}
	})

	t.Run("non-empty dataset", func(t *testing.T) {
		ds := dataset.New()
		if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"TEST"})); err != nil {
			t.Fatalf("Add() error = %v", err)
		}
		req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, ds)
		if err := req.SetMessageID(3); err != nil {
			t.Fatalf("SetMessageID() error = %v", err)
		}
		commandData, datasetData, err := EncodeDIMSEMessage(req, nil)
		if err != nil {
			t.Fatalf("EncodeDIMSEMessage() error = %v", err)
		}
		if len(datasetData) == 0 {
			t.Fatal("datasetData is empty, want encoded identifier")
		}
		if got := commandDataSetType(t, commandData); got != 0x0001 {
			t.Fatalf("CommandDataSetType = 0x%04X, want 0x0001", got)
		}
	})
}

func commandDataSetType(t *testing.T, commandData []byte) uint16 {
	t.Helper()
	ds, err := decodeRawDataset(commandData, transfer.ImplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("decodeRawDataset() error = %v", err)
	}
	value, err := ds.GetUInt16(tag.CommandDataSetType, 0)
	if err != nil {
		t.Fatalf("GetUInt16(CommandDataSetType) error = %v", err)
	}
	return value
}

// TestSendAcceptIdentifiesAcceptor verifies the AC carries the local AE's
// Implementation Class UID, never the requestor's (PS3.7 D.3.3.2.2).
func TestSendAcceptIdentifiesAcceptor(t *testing.T) {
	conn := &mockConn{}
	svc := NewService(conn, nil,
		WithImplementationClassUID("1.2.3.4.5"),
		WithImplementationVersionName("LOCAL-1.0"),
	)
	defer func() { _ = svc.Close() }()

	assoc := association.NewAssociation(testCalledAE, testCallingAE)
	assoc.ImplementationClassUID = "9.9.9.9" // requestor's value must not leak into the AC
	pc := association.NewPresentationContext(1, verificationSOPClassUID, transfer.ImplicitVRLittleEndian)
	pc.Accept(transfer.ImplicitVRLittleEndian)
	if err := assoc.AddPresentationContext(pc); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}

	responder := &serviceResponder{service: svc}
	if err := responder.SendAccept(context.Background(), assoc); err != nil {
		t.Fatalf("SendAccept() error = %v", err)
	}

	raw, err := pdu.ReadPDU(bytes.NewReader(conn.writeData))
	if err != nil {
		t.Fatalf("ReadPDU() error = %v", err)
	}
	ac := &pdu.AAssociateAC{}
	if err := ac.Decode(raw); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if ac.UserInformation == nil {
		t.Fatal("A-ASSOCIATE-AC has no user information")
	}
	if ac.UserInformation.ImplementationClassUID != "1.2.3.4.5" {
		t.Fatalf("ImplementationClassUID = %q, want the acceptor's 1.2.3.4.5",
			ac.UserInformation.ImplementationClassUID)
	}
	if ac.UserInformation.ImplementationVersionName != "LOCAL-1.0" {
		t.Fatalf("ImplementationVersionName = %q, want LOCAL-1.0",
			ac.UserInformation.ImplementationVersionName)
	}
}

// TestAsyncOperationsNegotiation verifies PS3.7 D.3.3.3: omit when not
// requested, clamp accepted values to the offered values.
func TestAsyncOperationsNegotiation(t *testing.T) {
	tests := []struct {
		name     string
		offered  *pdu.AsynchronousOperationsWindow
		config   [2]uint16
		want     *pdu.AsynchronousOperationsWindow
		wantNone bool
	}{
		{
			name:     "not requested is omitted",
			offered:  nil,
			config:   [2]uint16{1, 1},
			wantNone: true,
		},
		{
			name:    "accepted equals configured when below offered",
			offered: &pdu.AsynchronousOperationsWindow{MaximumNumberOperationsInvoked: 2, MaximumNumberOperationsPerformed: 3},
			config:  [2]uint16{1, 1},
			want:    &pdu.AsynchronousOperationsWindow{MaximumNumberOperationsInvoked: 1, MaximumNumberOperationsPerformed: 1},
		},
		{
			name:    "clamped to offered when configured is higher",
			offered: &pdu.AsynchronousOperationsWindow{MaximumNumberOperationsInvoked: 1, MaximumNumberOperationsPerformed: 2},
			config:  [2]uint16{4, 5},
			want:    &pdu.AsynchronousOperationsWindow{MaximumNumberOperationsInvoked: 1, MaximumNumberOperationsPerformed: 2},
		},
		{
			name:    "zero offered means unlimited",
			offered: &pdu.AsynchronousOperationsWindow{MaximumNumberOperationsInvoked: 0, MaximumNumberOperationsPerformed: 0},
			config:  [2]uint16{2, 3},
			want:    &pdu.AsynchronousOperationsWindow{MaximumNumberOperationsInvoked: 2, MaximumNumberOperationsPerformed: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := newTestRQ()
			rq.UserInformation.AsynchronousOperations = tt.offered
			raw, err := rq.Encode()
			if err != nil {
				t.Fatalf("Encode() error = %v", err)
			}

			conn := &mockConn{readData: serializeRawPDUForTest(raw)}
			svc := NewService(conn, nil, WithAsynchronousOperations(tt.config[0], tt.config[1]))
			defer func() { _ = svc.Close() }()

			if _, err := svc.ReceiveAssociationRequest(context.Background()); err != nil {
				t.Fatalf("ReceiveAssociationRequest() error = %v", err)
			}

			rawAC, err := pdu.ReadPDU(bytes.NewReader(conn.writeData))
			if err != nil {
				t.Fatalf("ReadPDU() error = %v", err)
			}
			ac := &pdu.AAssociateAC{}
			if err := ac.Decode(rawAC); err != nil {
				t.Fatalf("Decode() error = %v", err)
			}
			got := ac.UserInformation.AsynchronousOperations
			if tt.wantNone {
				if got != nil {
					t.Fatalf("AC async ops = %+v, want omitted", got)
				}
				return
			}
			if got == nil {
				t.Fatal("AC async ops missing")
			}
			if got.MaximumNumberOperationsInvoked != tt.want.MaximumNumberOperationsInvoked ||
				got.MaximumNumberOperationsPerformed != tt.want.MaximumNumberOperationsPerformed {
				t.Fatalf("AC async ops = %d/%d, want %d/%d",
					got.MaximumNumberOperationsInvoked, got.MaximumNumberOperationsPerformed,
					tt.want.MaximumNumberOperationsInvoked, tt.want.MaximumNumberOperationsPerformed)
			}
		})
	}
}

// TestInboundCCancelDoesNotTouchOutboundPending verifies the C-CANCEL Message ID
// is resolved in the inbound direction only.
func TestInboundCCancelDoesNotTouchOutboundPending(t *testing.T) {
	svc := NewService(&mockConn{}, nil)
	defer func() { _ = svc.Close() }()

	pending := &pendingRequest{
		responseCh: make(chan dimse.Response, 1),
		doneCh:     make(chan error, 1),
		cancelCh:   make(chan struct{}),
	}
	svc.pendingRequests[7] = pending

	if err := svc.handleCCancelRequest(dimse.NewCCancelRequest(7)); err != nil {
		t.Fatalf("handleCCancelRequest() error = %v", err)
	}
	if _, ok := svc.pendingRequests[7]; !ok {
		t.Fatal("inbound C-CANCEL removed an outbound pending request")
	}
	select {
	case <-pending.cancelCh:
		t.Fatal("inbound C-CANCEL closed the outbound cancel channel")
	default:
	}
}

// TestProtocolErrorSendsAbort verifies PS3.8 9.3.8: a P-DATA protocol error is
// reported with A-ABORT instead of a silent disconnect.
func TestProtocolErrorSendsAbort(t *testing.T) {
	raw, err := CreatePDataTFPDU([]*PDV{{
		PresentationContextID: 2, // even context IDs are invalid
		IsCommand:             true,
		IsLastFragment:        true,
		Data:                  []byte{0x00, 0x00},
	}})
	if err != nil {
		t.Fatalf("CreatePDataTFPDU() error = %v", err)
	}
	conn := &mockConn{readData: serializeRawPDUForTest(raw)}
	svc := NewService(conn, nil)
	defer func() { _ = svc.Close() }()
	if err := svc.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	if err := svc.recvLoop(context.Background()); err == nil {
		t.Fatal("recvLoop() error = nil, want protocol error")
	}
	if len(conn.writeData) < 6 || conn.writeData[0] != pdu.TypeAAbort {
		t.Fatalf("written PDU = % X, want A-ABORT", conn.writeData)
	}
}

// TestCommandDatasetValidation verifies group 0000 elements only.
func TestCommandDatasetValidation(t *testing.T) {
	if err := validateCommandDataset(dataset.New()); err == nil {
		t.Fatal("missing CommandField was accepted")
	}

	ds := dataset.New()
	if err := ds.Add(element.NewUnsignedShort(tag.CommandField, []uint16{uint16(dimse.CommandCEchoRQ)})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"X"})); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if err := validateCommandDataset(ds); err == nil {
		t.Fatal("non-command element was accepted in a command dataset")
	}
}

// TestSendAssignsMessageID verifies a request queued through Send always carries
// a Message ID (PS3.7 9.3.1).
func TestSendAssignsMessageID(t *testing.T) {
	svc := NewService(&mockConn{}, nil)
	defer func() { _ = svc.Close() }()
	if err := svc.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case req := <-svc.sendQueue:
			req.resultCh <- nil
		case <-svc.closeCh:
		}
	}()

	req := dimse.NewCEchoRequest()
	if err := svc.Send(context.Background(), req); err != nil {
		t.Fatalf("Send() error = %v", err)
	}
	<-done
	if req.MessageID() == 0 {
		t.Fatal("Send() left the request without a Message ID")
	}
}

// TestCMoveHandlerWithoutFinalResponseIsCompleted verifies the service sends a
// final response when the handler forgets it, and carries the last counts.
func TestCMoveHandlerWithoutFinalResponseIsCompleted(t *testing.T) {
	svc := NewService(&mockConn{}, nil)
	defer func() { _ = svc.Close() }()
	if err := svc.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	sent := make(chan dimse.Message, 8)
	go func() {
		for {
			select {
			case req := <-svc.sendQueue:
				sent <- req.message
				req.resultCh <- nil
			case <-svc.closeCh:
				return
			}
		}
	}()

	req := dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, "DEST", dataset.New())
	handlers := &Handlers{CMoveHandler: func(_ context.Context, op CMoveOperation) error {
		return op.SendPending(2, 1, 0, 0)
	}}
	if err := svc.handleCMoveRequest(context.Background(), req, handlers); err != nil {
		t.Fatalf("handleCMoveRequest() error = %v", err)
	}

	pending := (<-sent).(*dimse.CMoveResponse)
	if pending.StatusCode() != 0xFF00 {
		t.Fatalf("first response status = 0x%04X, want pending", pending.StatusCode())
	}
	final := (<-sent).(*dimse.CMoveResponse)
	if final.StatusCode() != 0x0000 {
		t.Fatalf("final response status = 0x%04X, want success", final.StatusCode())
	}
	if !final.HasSubOperationCounts() {
		t.Fatal("final response has no sub-operation counts")
	}
	if final.NumberOfCompletedSubOperations() != 1 {
		t.Fatalf("completed = %d, want 1", final.NumberOfCompletedSubOperations())
	}
	if final.NumberOfRemainingSubOperations() != 0 {
		t.Fatalf("remaining = %d, want 0 on a successful final response",
			final.NumberOfRemainingSubOperations())
	}
}

// TestCGetHandlerWithoutFinalResponseIsCompleted mirrors the C-MOVE test for
// C-GET.
func TestCGetHandlerWithoutFinalResponseIsCompleted(t *testing.T) {
	svc := NewService(&mockConn{}, nil)
	defer func() { _ = svc.Close() }()
	if err := svc.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	sent := make(chan dimse.Message, 8)
	go func() {
		for {
			select {
			case req := <-svc.sendQueue:
				sent <- req.message
				req.resultCh <- nil
			case <-svc.closeCh:
				return
			}
		}
	}()

	req := dimse.NewCGetRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	handlers := &Handlers{CGetHandler: func(_ context.Context, op CGetOperation) error {
		return op.SendPending(2, 1, 0, 0)
	}}
	if err := svc.handleCGetRequest(context.Background(), req, handlers); err != nil {
		t.Fatalf("handleCGetRequest() error = %v", err)
	}

	<-sent // pending
	final := (<-sent).(*dimse.CGetResponse)
	if final.StatusCode() != 0x0000 {
		t.Fatalf("final response status = 0x%04X, want success", final.StatusCode())
	}
	if final.NumberOfCompletedSubOperations() != 1 || final.NumberOfRemainingSubOperations() != 0 {
		t.Fatalf("final counts = remaining %d, completed %d; want 0/1",
			final.NumberOfRemainingSubOperations(), final.NumberOfCompletedSubOperations())
	}
}

// TestPDataRejectedAfterReleaseRequest verifies that a queued DIMSE message is
// not written after A-RELEASE-RQ started (PS3.8 Sta7).
func TestPDataRejectedAfterReleaseRequest(t *testing.T) {
	conn := &mockConn{}
	svc := NewService(conn, nil)
	defer func() { _ = svc.Close() }()
	if err := svc.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	if err := svc.setState(StateReleaseRequested); err != nil {
		t.Fatalf("setState() error = %v", err)
	}

	raw := pdu.NewRawPDU(pdu.TypePDataTF, []byte{0x00, 0x00, 0x00, 0x02, 0x01, 0x03})
	if err := svc.writePDU(context.Background(), raw); err == nil {
		t.Fatal("writePDU(P-DATA-TF) succeeded after A-RELEASE-RQ")
	}
	if len(conn.writeData) != 0 {
		t.Fatalf("wrote %d bytes after A-RELEASE-RQ", len(conn.writeData))
	}
}

// TestInvalidAETitleGetsReject verifies an all-space peer AE Title produces an
// A-ASSOCIATE-RJ instead of a silent close.
func TestInvalidAETitleGetsReject(t *testing.T) {
	rq := newTestRQ()
	raw, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	data := append([]byte(nil), raw.Data...)
	called := bytes.Repeat([]byte{' '}, 16)
	copy(called, testCalledAE)
	idx := bytes.Index(data, called)
	if idx < 0 {
		t.Fatal("called AE Title not found in encoded RQ")
	}
	for i := 0; i < 16; i++ {
		data[idx+i] = ' '
	}

	conn := &mockConn{readData: serializeRawPDUForTest(pdu.NewRawPDU(raw.Type, data))}
	svc := NewService(conn, nil)
	defer func() { _ = svc.Close() }()

	_, err = svc.ReceiveAssociationRequest(context.Background())
	if !errors.Is(err, ErrAssociationRejected) {
		t.Fatalf("ReceiveAssociationRequest() error = %v, want ErrAssociationRejected", err)
	}
	if len(conn.writeData) < 6 || conn.writeData[0] != pdu.TypeAAssociateRJ {
		t.Fatalf("written PDU = % X, want A-ASSOCIATE-RJ", conn.writeData)
	}
}

// TestDefaultNegotiatorRejectsUnknownTransferSyntax verifies PS3.8 9.3.3.2.
func TestDefaultNegotiatorRejectsUnknownTransferSyntax(t *testing.T) {
	n := &DefaultAssociationNegotiator{Registry: transfer.NewRegistry()}
	unknown := transfer.New(uid.Parse("1.2.3.4.5", "", uid.TypeTransferSyntax))
	if n.supportsTransferSyntax(unknown) {
		t.Fatal("unknown transfer syntax was reported as supported")
	}
	if !n.supportsTransferSyntax(transfer.ImplicitVRLittleEndian) {
		t.Fatal("standard transfer syntax was reported as unsupported")
	}
}

// TestRejectedAssociationReturnsError verifies the caller can tell a rejection
// from success.
func TestRejectedAssociationReturnsError(t *testing.T) {
	rq := newTestRQ()
	raw, err := rq.Encode()
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	conn := &mockConn{readData: serializeRawPDUForTest(raw)}
	svc := NewService(conn, nil, WithAssociationNegotiator(FuncAssociationNegotiator(
		func(ctx context.Context, _ *association.Association, responder AssociationResponder) error {
			return responder.SendReject(ctx, pdu.ResultRejectedPermanent, pdu.SourceServiceUser, pdu.ReasonServiceUserNoReasonGiven)
		},
	)))
	defer func() { _ = svc.Close() }()

	_, err = svc.ReceiveAssociationRequest(context.Background())
	if !errors.Is(err, ErrAssociationRejected) {
		t.Fatalf("ReceiveAssociationRequest() error = %v, want ErrAssociationRejected", err)
	}
}

// TestSendReleaseResponseRejectsInvalidState verifies the state guard.
func TestSendReleaseResponseRejectsInvalidState(t *testing.T) {
	svc := NewService(&mockConn{}, nil)
	defer func() { _ = svc.Close() }()
	if err := svc.SendReleaseResponse(context.Background()); err == nil {
		t.Fatal("SendReleaseResponse() succeeded in Idle state")
	}
}

// TestSendRejectsNilMessage verifies a nil message is rejected instead of
// panicking inside the send loop.
func TestSendRejectsNilMessage(t *testing.T) {
	svc := NewService(&mockConn{}, nil)
	defer func() { _ = svc.Close() }()
	if err := svc.Send(context.Background(), nil); err == nil {
		t.Fatal("Send(nil) succeeded")
	}
}

// TestDefaultMaxMessageSizeIsBounded verifies the reassembly bound is on by default.
func TestDefaultMaxMessageSizeIsBounded(t *testing.T) {
	if got := defaultServiceConfig().maxMessageSize; got == 0 {
		t.Fatal("default maxMessageSize is unlimited")
	}
}
