// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package observability

import (
	"context"
	"log/slog"
	"sync"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/logging"
)

func TestEmitSlogWritesSafeNetworkLifecycleRecord(t *testing.T) {
	handler := &slogRecordHandler{}
	previous := logging.Logger()
	logging.SetLogger(slog.New(handler))
	t.Cleanup(func() { logging.SetLogger(previous) })

	EmitSlog(context.Background(), "network.service", LogRecord{
		Level:   LevelInfo,
		Message: string(EventRequestCompleted),
		Event: Event{
			Timestamp: time.Unix(1, 0),
			Kind:      EventRequestCompleted,
			Association: Association{
				ConnectionID:  7,
				AssociationID: 9,
				CallingAE:     "SCU",
				CalledAE:      "SCP",
			},
			Direction:   DirectionOutbound,
			MessageID:   3,
			Command:     0x0030,
			StatusCode:  0,
			StatusState: "Success",
			Duration:    12 * time.Millisecond,
			Outcome:     OutcomeSuccess,
		},
	})

	records := handler.snapshot()
	if len(records) != 1 {
		t.Fatalf("slog records = %d, want 1", len(records))
	}
	record := records[0]
	if record.Level != slog.LevelInfo || record.Message != "request_completed" {
		t.Fatalf("record = level=%s message=%q, want info/request_completed", record.Level, record.Message)
	}
	attrs := recordAttrs(record)
	for key, want := range map[string]any{
		"component":      "network.service",
		"event":          "request_completed",
		"connection_id":  uint64(7),
		"association_id": uint64(9),
		"calling_ae":     "SCU",
		"called_ae":      "SCP",
		"direction":      "outbound",
		"message_id":     uint64(3),
		"command":        uint64(0x0030),
		"status":         uint64(0),
		"status_state":   "Success",
		"outcome":        "success",
	} {
		if got, ok := attrs[key]; !ok || got != want {
			t.Errorf("attribute %q = %#v, want %#v", key, got, want)
		}
	}
	for _, forbidden := range []string{"dataset", "pdu", "pixel_data", "patient_name", "patient_id"} {
		if _, exists := attrs[forbidden]; exists {
			t.Errorf("unexpected sensitive attribute %q in %#v", forbidden, attrs)
		}
	}
}

func TestEmitSlogUsesRecordMessageWhenEventKindIsUnavailable(t *testing.T) {
	handler := &slogRecordHandler{}
	previous := logging.Logger()
	logging.SetLogger(slog.New(handler))
	t.Cleanup(func() { logging.SetLogger(previous) })

	EmitSlog(context.Background(), "network.server", LogRecord{
		Level:   LevelError,
		Message: "listener_failed",
		Event: Event{
			Outcome: OutcomeFailure,
		},
	})

	records := handler.snapshot()
	if len(records) != 1 {
		t.Fatalf("slog records = %d, want 1", len(records))
	}
	attrs := recordAttrs(records[0])
	if attrs["event"] != "listener_failed" {
		t.Errorf("event = %#v, want listener_failed", attrs["event"])
	}
}

type slogRecordHandler struct {
	mu      sync.Mutex
	records []slog.Record
}

func (h *slogRecordHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *slogRecordHandler) Handle(_ context.Context, record slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.records = append(h.records, record.Clone())
	return nil
}

func (h *slogRecordHandler) WithAttrs([]slog.Attr) slog.Handler { return h }

func (h *slogRecordHandler) WithGroup(string) slog.Handler { return h }

func (h *slogRecordHandler) snapshot() []slog.Record {
	h.mu.Lock()
	defer h.mu.Unlock()
	return append([]slog.Record(nil), h.records...)
}

func recordAttrs(record slog.Record) map[string]any {
	attrs := make(map[string]any)
	record.Attrs(func(attr slog.Attr) bool {
		attrs[attr.Key] = attr.Value.Any()
		return true
	})
	return attrs
}
