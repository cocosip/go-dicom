// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package server

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/service"
)

type serverObservationRecorder struct {
	mu      sync.Mutex
	events  []observability.Event
	metrics []observability.Metric
	logs    []observability.LogRecord
}

func (r *serverObservationRecorder) metric(_ context.Context, metric observability.Metric) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.metrics = append(r.metrics, metric)
}

func (r *serverObservationRecorder) log(_ context.Context, record observability.LogRecord) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.logs = append(r.logs, record)
}

func (r *serverObservationRecorder) event(_ context.Context, event observability.Event) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events = append(r.events, event)
}

func TestObservabilityListenFailureIsStructured(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer func() { _ = listener.Close() }()
	port := listener.Addr().(*net.TCPAddr).Port
	recorder := &serverObservationRecorder{}
	server := New(
		WithPort(port),
		WithLogger(observability.LoggerFunc(recorder.log)),
		WithMetricsObserver(observability.MetricsObserverFunc(recorder.metric)),
	)

	if err := server.ListenAndServe(context.Background()); err == nil {
		t.Fatal("ListenAndServe() error = nil, want occupied-port failure")
	}

	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	if len(recorder.logs) != 1 ||
		recorder.logs[0].Level != observability.LevelError ||
		recorder.logs[0].Message != "listener_failed" ||
		recorder.logs[0].Event.Outcome != observability.OutcomeFailure {
		t.Fatalf("logs = %#v, want one structured listener failure", recorder.logs)
	}
	if len(recorder.metrics) != 1 ||
		recorder.metrics[0].Kind != observability.MetricError ||
		recorder.metrics[0].ErrorKind != observability.ErrorConnection ||
		recorder.metrics[0].Value != 1 {
		t.Fatalf("metrics = %#v, want one connection error", recorder.metrics)
	}
}

func TestObservabilityHooksReachAcceptedService(t *testing.T) {
	recorder := &serverObservationRecorder{}
	server := New(
		WithAssociationTimeout(time.Second),
		WithEventObserver(observability.EventObserverFunc(recorder.event)),
	)
	server.ctx, server.cancel = context.WithCancel(context.Background())
	defer server.cancel()

	clientConn, serverConn := net.Pipe()
	done := make(chan struct{})
	server.wg.Add(1)
	go func() {
		server.handleConnection(serverConn)
		close(done)
	}()

	clientService := service.NewService(clientConn, nil, service.WithAssociationRequestor(true))
	rq := &pdu.AAssociateRQ{
		ProtocolVersion:    1,
		CalledAETitle:      "OBS-SCP",
		CallingAETitle:     "OBS-SCU",
		ApplicationContext: "1.2.840.10008.3.1.1.1",
		PresentationContexts: []pdu.PresentationContextRQ{{
			ID:               1,
			AbstractSyntax:   testVerificationSOPClassUID,
			TransferSyntaxes: []string{"1.2.840.10008.1.2.1"},
		}},
		UserInformation: &pdu.UserInformation{MaximumLength: 16384},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := clientService.SendAssociationRequest(ctx, rq); err != nil {
		t.Fatalf("SendAssociationRequest() error = %v", err)
	}
	if _, err := clientService.ReceiveAssociationResponse(ctx); err != nil {
		t.Fatalf("ReceiveAssociationResponse() error = %v", err)
	}
	_ = clientService.Close()
	select {
	case <-done:
	case <-ctx.Done():
		t.Fatal("server connection did not stop")
	}

	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	var opened, requested, accepted bool
	for _, event := range recorder.events {
		switch event.Kind {
		case observability.EventConnectionOpened:
			opened = true
		case observability.EventAssociationRequested:
			requested = event.Direction == observability.DirectionInbound
		case observability.EventAssociationAccepted:
			accepted = event.Direction == observability.DirectionOutbound
		}
	}
	if !opened || !requested || !accepted {
		t.Fatalf("events = %#v, want opened, inbound requested, outbound accepted", recorder.events)
	}
}
