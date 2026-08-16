// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"errors"
	"io"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func startObservedServicePair(t *testing.T, clientRecorder, serverRecorder *observationRecorder, handlers *Handlers) *Service {
	t.Helper()
	clientConn, serverConn := net.Pipe()
	clientService := NewService(clientConn, observedTestAssociation(t),
		WithAssociationRequestor(true),
		WithEventObserver(observability.EventObserverFunc(clientRecorder.observeEvent)),
		WithMetricsObserver(observability.MetricsObserverFunc(clientRecorder.observeMetric)),
	)
	serverService := NewService(serverConn, observedTestAssociation(t),
		WithAssociationRequestor(false),
		WithEventObserver(observability.EventObserverFunc(serverRecorder.observeEvent)),
		WithMetricsObserver(observability.MetricsObserverFunc(serverRecorder.observeMetric)),
	)
	serverService.SetHandlers(handlers)
	t.Cleanup(func() {
		_ = clientService.Close()
		_ = serverService.Close()
	})
	if err := clientService.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("client setState() error = %v", err)
	}
	if err := serverService.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("server setState() error = %v", err)
	}
	if err := serverService.Start(); err != nil {
		t.Fatalf("server Start() error = %v", err)
	}
	if err := clientService.Start(); err != nil {
		t.Fatalf("client Start() error = %v", err)
	}
	return clientService
}

func observedTestAssociation(t *testing.T) *association.Association {
	t.Helper()
	assoc := createTestAssociation()
	pc := association.NewPresentationContext(5, "1.2.840.10008.5.1.4.1.2.2.1", transfer.ExplicitVRLittleEndian)
	pc.AcceptedTransferSyntax = transfer.ExplicitVRLittleEndian
	pc.Result = association.ResultAcceptance
	if err := assoc.AddPresentationContext(pc); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}
	return assoc
}

func requestEvents(recorder *observationRecorder, command dimse.CommandField) []observability.Event {
	events, _, _ := recorder.snapshot()
	result := make([]observability.Event, 0, len(events))
	for _, event := range events {
		if event.Command == uint16(command) && event.MessageID != 0 {
			result = append(result, event)
		}
	}
	return result
}

func eventKinds(events []observability.Event) []observability.EventKind {
	kinds := make([]observability.EventKind, len(events))
	for i, event := range events {
		kinds[i] = event.Kind
	}
	return kinds
}

func waitForRequestEventCount(t *testing.T, recorder *observationRecorder, command dimse.CommandField, want int) []observability.Event {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for {
		events := requestEvents(recorder, command)
		if len(events) >= want {
			return events
		}
		if time.Now().After(deadline) {
			t.Fatalf("request event count = %d, want at least %d: %#v", len(events), want, events)
		}
		time.Sleep(time.Millisecond)
	}
}

func TestOutboundRequestObservabilitySuccessAndInboundLifecycle(t *testing.T) {
	clientRecorder := &observationRecorder{}
	serverRecorder := &observationRecorder{}
	clientService := startObservedServicePair(t, clientRecorder, serverRecorder, &Handlers{
		CEchoHandler: func(_ context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
			return dimse.NewCEchoResponseFromRequest(req, status.Success), nil
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := clientService.SendCEcho(ctx, dimse.NewCEchoRequest())
	if err != nil {
		t.Fatalf("SendCEcho() error = %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("response status = %v, want success", resp.Status())
	}

	clientEvents := requestEvents(clientRecorder, dimse.CommandCEchoRQ)
	wantClient := []observability.EventKind{observability.EventRequestSent, observability.EventRequestCompleted}
	if got := eventKinds(clientEvents); !equalEventKinds(got, wantClient) {
		t.Fatalf("client event kinds = %v, want %v", got, wantClient)
	}
	serverEvents := waitForRequestEventCount(t, serverRecorder, dimse.CommandCEchoRQ, 2)
	wantServer := []observability.EventKind{observability.EventRequestReceived, observability.EventRequestCompleted}
	if got := eventKinds(serverEvents); !equalEventKinds(got, wantServer) {
		t.Fatalf("server event kinds = %v, want %v", got, wantServer)
	}
	for _, event := range append(clientEvents, serverEvents...) {
		if event.Association.AssociationID == 0 || event.MessageID == 0 || event.Duration < 0 {
			t.Errorf("incomplete request event: %#v", event)
		}
	}
	if clientEvents[1].StatusCode != 0 || clientEvents[1].Outcome != observability.OutcomeSuccess {
		t.Fatalf("client completion = %#v", clientEvents[1])
	}
}

func TestOutboundRequestObservabilityPendingOrder(t *testing.T) {
	clientRecorder := &observationRecorder{}
	serverRecorder := &observationRecorder{}
	clientService := startObservedServicePair(t, clientRecorder, serverRecorder, &Handlers{
		CFindHandler: func(_ context.Context, req *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
			return []*dimse.CFindResponse{
				dimse.NewCFindResponseFromRequest(req, status.Pending, nil),
				dimse.NewCFindResponseFromRequest(req, status.PendingWarning, nil),
				dimse.NewCFindResponseFromRequest(req, status.Success, nil),
			}, nil
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	responses, err := clientService.SendCFind(ctx, dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New()))
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	for response := range responses {
		if response == nil {
			t.Fatal("SendCFind() returned a nil response")
		}
	}

	events := requestEvents(clientRecorder, dimse.CommandCFindRQ)
	want := []observability.EventKind{
		observability.EventRequestSent,
		observability.EventRequestPending,
		observability.EventRequestPending,
		observability.EventRequestCompleted,
	}
	if got := eventKinds(events); !equalEventKinds(got, want) {
		t.Fatalf("event kinds = %v, want %v", got, want)
	}
	if events[1].StatusCode != status.Pending.Code || events[2].StatusCode != status.PendingWarning.Code {
		t.Fatalf("pending status codes = [%04X, %04X]", events[1].StatusCode, events[2].StatusCode)
	}
}

func TestOutboundRequestObservabilityTimeout(t *testing.T) {
	recorder := &observationRecorder{}
	clientConn, peerConn := net.Pipe()
	service := NewService(clientConn, createTestAssociation(),
		WithAssociationRequestor(true),
		WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
		WithMetricsObserver(observability.MetricsObserverFunc(recorder.observeMetric)),
	)
	t.Cleanup(func() {
		_ = service.Close()
		_ = peerConn.Close()
	})
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	if err := service.Start(); err != nil {
		t.Fatalf("Start() error = %v", err)
	}
	go func() {
		buffer := make([]byte, 4096)
		for {
			if _, readErr := peerConn.Read(buffer); readErr != nil {
				return
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Millisecond)
	defer cancel()
	_, err := service.SendCEcho(ctx, dimse.NewCEchoRequest())
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("SendCEcho() error = %v, want deadline exceeded", err)
	}

	events := requestEvents(recorder, dimse.CommandCEchoRQ)
	want := []observability.EventKind{observability.EventRequestSent, observability.EventRequestTimedOut}
	if got := eventKinds(events); !equalEventKinds(got, want) {
		t.Fatalf("event kinds = %v, want %v", got, want)
	}
	if events[1].Outcome != observability.OutcomeTimeout || events[1].Duration <= 0 {
		t.Fatalf("timeout event = %#v", events[1])
	}
	assertRequestErrorMetric(t, recorder, dimse.CommandCEchoRQ, observability.OutcomeTimeout)
}

func TestOutboundRequestObservabilityLocalCancellation(t *testing.T) {
	recorder := &observationRecorder{}
	clientConn, peerConn := net.Pipe()
	service := NewService(clientConn, createTestAssociation(),
		WithAssociationRequestor(true),
		WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
	)
	t.Cleanup(func() {
		_ = service.Close()
		_ = peerConn.Close()
	})
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	if err := service.Start(); err != nil {
		t.Fatalf("Start() error = %v", err)
	}
	go func() {
		_, _ = io.Copy(io.Discard, peerConn)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan error, 1)
	go func() {
		_, err := service.SendCEcho(ctx, dimse.NewCEchoRequest())
		result <- err
	}()
	waitForRequestEventCount(t, recorder, dimse.CommandCEchoRQ, 1)
	cancel()
	if err := <-result; !errors.Is(err, context.Canceled) {
		t.Fatalf("SendCEcho() error = %v, want context canceled", err)
	}

	events := requestEvents(recorder, dimse.CommandCEchoRQ)
	want := []observability.EventKind{observability.EventRequestSent, observability.EventRequestCancelled}
	if got := eventKinds(events); !equalEventKinds(got, want) {
		t.Fatalf("event kinds = %v, want %v", got, want)
	}
}

func TestInboundRequestObservabilityPeerCancellation(t *testing.T) {
	clientRecorder := &observationRecorder{}
	serverRecorder := &observationRecorder{}
	handlerStarted := make(chan struct{})
	handlerCancelled := make(chan struct{})
	clientService := startObservedServicePair(t, clientRecorder, serverRecorder, &Handlers{
		CFindHandler: func(ctx context.Context, _ *dimse.CFindRequest) ([]*dimse.CFindResponse, error) {
			close(handlerStarted)
			<-ctx.Done()
			close(handlerCancelled)
			return nil, ctx.Err()
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	responses, err := clientService.SendCFind(ctx, req)
	if err != nil {
		t.Fatalf("SendCFind() error = %v", err)
	}
	<-handlerStarted
	if err := clientService.SendCCancel(ctx, req.MessageID(), 5); err != nil {
		t.Fatalf("SendCCancel() error = %v", err)
	}
	<-handlerCancelled
	for response := range responses {
		if response == nil {
			t.Fatal("SendCFind() returned a nil response")
		}
	}

	events := requestEvents(serverRecorder, dimse.CommandCFindRQ)
	want := []observability.EventKind{observability.EventRequestReceived, observability.EventRequestCancelled}
	if got := eventKinds(events); !equalEventKinds(got, want) {
		t.Fatalf("event kinds = %v, want %v", got, want)
	}
	if events[1].Outcome != observability.OutcomeCancelled {
		t.Fatalf("cancel event = %#v", events[1])
	}
	clientCancelEvents := requestEvents(clientRecorder, dimse.CommandCCancelRQ)
	if got := eventKinds(clientCancelEvents); !equalEventKinds(got, []observability.EventKind{
		observability.EventRequestSent,
		observability.EventRequestCompleted,
	}) {
		t.Fatalf("client C-CANCEL event kinds = %v, want sent/completed", got)
	}
	serverCancelEvents := requestEvents(serverRecorder, dimse.CommandCCancelRQ)
	if got := eventKinds(serverCancelEvents); !equalEventKinds(got, []observability.EventKind{
		observability.EventRequestReceived,
		observability.EventRequestCompleted,
	}) {
		t.Fatalf("server C-CANCEL event kinds = %v, want received/completed", got)
	}
	for _, event := range append(clientCancelEvents, serverCancelEvents...) {
		if event.MessageID != req.MessageID() {
			t.Errorf("C-CANCEL event message ID = %d, want target %d", event.MessageID, req.MessageID())
		}
	}
}

type failingWriteConn struct {
	closed    chan struct{}
	closeOnce sync.Once
}

func newFailingWriteConn() *failingWriteConn {
	return &failingWriteConn{closed: make(chan struct{})}
}

func (c *failingWriteConn) Read([]byte) (int, error) {
	<-c.closed
	return 0, net.ErrClosed
}

func (*failingWriteConn) Write([]byte) (int, error) { return 0, errors.New("transport write failed") }
func (c *failingWriteConn) Close() error {
	c.closeOnce.Do(func() { close(c.closed) })
	return nil
}
func (*failingWriteConn) LocalAddr() net.Addr              { return nil }
func (*failingWriteConn) RemoteAddr() net.Addr             { return nil }
func (*failingWriteConn) SetDeadline(time.Time) error      { return nil }
func (*failingWriteConn) SetReadDeadline(time.Time) error  { return nil }
func (*failingWriteConn) SetWriteDeadline(time.Time) error { return nil }

func TestOutboundRequestObservabilityTransportFailure(t *testing.T) {
	recorder := &observationRecorder{}
	service := NewService(newFailingWriteConn(), createTestAssociation(),
		WithAssociationRequestor(true),
		WithEventObserver(observability.EventObserverFunc(recorder.observeEvent)),
		WithMetricsObserver(observability.MetricsObserverFunc(recorder.observeMetric)),
	)
	t.Cleanup(func() { _ = service.Close() })
	if err := service.setState(StateAssociationAccepted); err != nil {
		t.Fatalf("setState() error = %v", err)
	}
	if err := service.Start(); err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if _, err := service.SendCEcho(ctx, dimse.NewCEchoRequest()); err == nil {
		t.Fatal("SendCEcho() error = nil, want transport failure")
	}

	events := requestEvents(recorder, dimse.CommandCEchoRQ)
	want := []observability.EventKind{observability.EventRequestFailed}
	if got := eventKinds(events); !equalEventKinds(got, want) {
		t.Fatalf("event kinds = %v, want %v", got, want)
	}
	assertRequestErrorMetric(t, recorder, dimse.CommandCEchoRQ, observability.OutcomeFailure)
}

func assertRequestErrorMetric(
	t *testing.T,
	recorder *observationRecorder,
	command dimse.CommandField,
	outcome observability.Outcome,
) {
	t.Helper()
	_, metrics, _ := recorder.snapshot()
	for _, metric := range metrics {
		if metric.Kind == observability.MetricError &&
			metric.Command == uint16(command) &&
			metric.Outcome == outcome &&
			metric.Value == 1 {
			return
		}
	}
	t.Fatalf("request error metric for %s/%s not found: %#v", command, outcome, metrics)
}

func equalEventKinds(left, right []observability.EventKind) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
