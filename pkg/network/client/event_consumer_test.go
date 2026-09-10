// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

func TestConsumeCFindEventsDeliversPendingIdentifiersAndAcceptsFinal(t *testing.T) {
	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	first := dataset.New()
	second := dataset.New()
	events := bufferedEvents(
		service.ResponseEvent[*dimse.CFindResponse]{
			Response: dimse.NewCFindResponseFromRequest(request, status.CFindPending, first),
		},
		service.ResponseEvent[*dimse.CFindResponse]{
			Response: dimse.NewCFindResponseFromRequest(
				request, status.CFindPendingWarningOptionalKeysNotSupported, second,
			),
		},
		service.ResponseEvent[*dimse.CFindResponse]{
			Response: dimse.NewCFindResponseFromRequest(request, status.Success, nil),
		},
	)

	var got []*dataset.Dataset
	err := consumeCFindEvents(context.Background(), events, func(identifier *dataset.Dataset) bool {
		got = append(got, identifier)
		return true
	}, func() error {
		t.Fatal("normal C-FIND sent C-CANCEL")
		return nil
	})
	if err != nil {
		t.Fatalf("consumeCFindEvents() error = %v", err)
	}
	if len(got) != 2 || got[0] != first || got[1] != second {
		t.Fatalf("identifiers = %#v, want first and second in order", got)
	}
}

func TestConsumeCFindEventsReturnsTerminalErrorAfterPendingResponse(t *testing.T) {
	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	identifier := dataset.New()
	wantErr := service.ErrRequestTimeout
	events := bufferedEvents(
		service.ResponseEvent[*dimse.CFindResponse]{
			Response: dimse.NewCFindResponseFromRequest(request, status.CFindPending, identifier),
		},
		service.ResponseEvent[*dimse.CFindResponse]{Err: wantErr},
	)

	count := 0
	err := consumeCFindEvents(context.Background(), events, func(got *dataset.Dataset) bool {
		if got != identifier {
			t.Fatalf("identifier = %p, want %p", got, identifier)
		}
		count++
		return true
	}, func() error { return nil })
	if !errors.Is(err, wantErr) {
		t.Fatalf("consumeCFindEvents() error = %v, want %v", err, wantErr)
	}
	if count != 1 {
		t.Fatalf("callback count = %d, want 1", count)
	}
}

func TestConsumeCFindEventsCancelsOnceAndWaitsForFinalCancel(t *testing.T) {
	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	events := bufferedEvents(
		service.ResponseEvent[*dimse.CFindResponse]{
			Response: dimse.NewCFindResponseFromRequest(request, status.CFindPending, dataset.New()),
		},
		service.ResponseEvent[*dimse.CFindResponse]{
			Response: dimse.NewCFindResponseFromRequest(request, status.Cancel, nil),
		},
	)

	cancels := 0
	err := consumeCFindEvents(context.Background(), events, func(*dataset.Dataset) bool {
		return false
	}, func() error {
		cancels++
		return nil
	})
	if err != nil {
		t.Fatalf("consumeCFindEvents() error = %v", err)
	}
	if cancels != 1 {
		t.Fatalf("C-CANCEL count = %d, want 1", cancels)
	}
}

func TestConsumeCFindEventsRejectsSilentClose(t *testing.T) {
	events := make(chan service.ResponseEvent[*dimse.CFindResponse])
	close(events)

	err := consumeCFindEvents(context.Background(), events, func(*dataset.Dataset) bool { return true }, func() error { return nil })
	if err == nil || !strings.Contains(err.Error(), "without a final response or terminal error") {
		t.Fatalf("consumeCFindEvents() error = %v, want incomplete stream error", err)
	}
}

func TestConsumeCFindEventsReturnsCancelSendFailure(t *testing.T) {
	request := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	events := bufferedEvents(
		service.ResponseEvent[*dimse.CFindResponse]{
			Response: dimse.NewCFindResponseFromRequest(request, status.CFindPending, dataset.New()),
		},
	)
	wantErr := errors.New("cancel write failed")

	err := consumeCFindEvents(context.Background(), events, func(*dataset.Dataset) bool {
		return false
	}, func() error { return wantErr })
	if !errors.Is(err, wantErr) {
		t.Fatalf("consumeCFindEvents() error = %v, want cancel failure", err)
	}
}

func TestConsumeCMoveEventsDeliversCountsAndAcceptsWarning(t *testing.T) {
	request := dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, "DEST", dataset.New())
	events := bufferedEvents(
		service.ResponseEvent[*dimse.CMoveResponse]{
			Response: dimse.NewCMoveResponsePending(1, request.AffectedSOPClassUID(), 3, 2, 1, 4),
		},
		service.ResponseEvent[*dimse.CMoveResponse]{
			Response: dimse.NewCMoveResponseFromRequest(request, status.CMoveWarningSubOperationsComplete),
		},
	)

	var got [4]uint16
	err := consumeCMoveEvents(context.Background(), events, func(remaining, completed, failed, warning uint16) bool {
		got = [4]uint16{remaining, completed, failed, warning}
		return true
	}, func() error { return nil })
	if err != nil {
		t.Fatalf("consumeCMoveEvents() error = %v", err)
	}
	if got != [4]uint16{3, 2, 1, 4} {
		t.Fatalf("counts = %v, want [3 2 1 4]", got)
	}
}

func TestConsumeCGetEventsReturnsTerminalAssociationClose(t *testing.T) {
	events := bufferedEvents(
		service.ResponseEvent[*dimse.CGetResponse]{Err: service.ErrServiceClosed},
	)

	err := consumeCGetEvents(context.Background(), events, nil, func() error { return nil })
	if !errors.Is(err, service.ErrServiceClosed) {
		t.Fatalf("consumeCGetEvents() error = %v, want ErrServiceClosed", err)
	}
}

func TestConsumeCMoveEventsContextCancellationSendsOneCancel(t *testing.T) {
	events := make(chan service.ResponseEvent[*dimse.CMoveResponse])
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	cancels := 0

	err := consumeCMoveEvents(ctx, events, nil, func() error {
		cancels++
		return nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("consumeCMoveEvents() error = %v, want context.Canceled", err)
	}
	if cancels != 1 {
		t.Fatalf("C-CANCEL count = %d, want 1", cancels)
	}
}

func TestConsumeCMoveEventsBufferedFinalWinsOverCanceledContext(t *testing.T) {
	request := dimse.NewCMoveRequest(dimse.QueryRetrieveLevelStudy, "DEST", dataset.New())
	events := bufferedEvents(
		service.ResponseEvent[*dimse.CMoveResponse]{
			Response: dimse.NewCMoveResponseFromRequest(request, status.Success),
		},
	)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := consumeCMoveEvents(ctx, events, nil, func() error {
		t.Fatal("buffered final response caused C-CANCEL")
		return nil
	})
	if err != nil {
		t.Fatalf("consumeCMoveEvents() error = %v, want buffered Success", err)
	}
}

func bufferedEvents[T dimse.Response](events ...service.ResponseEvent[T]) <-chan service.ResponseEvent[T] {
	result := make(chan service.ResponseEvent[T], len(events))
	for _, event := range events {
		result <- event
	}
	close(result)
	return result
}
