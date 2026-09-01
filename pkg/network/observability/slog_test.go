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

const (
	testSlogImplementationClassUID = "1.2.3"
	testSlogImplementationVersion  = "TEST_1"
	testSlogSuccessStatus          = "Success"
)

func TestEmitSlogWritesSafeNetworkLifecycleRecord(t *testing.T) {
	handler := &slogRecordHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	EmitSlog(context.Background(), "network.service", LogRecord{
		Level:   LevelInfo,
		Message: string(EventRequestCompleted),
		Event: Event{
			Timestamp: time.Unix(1, 0),
			Kind:      EventRequestCompleted,
			Association: Association{
				ConnectionID:                     7,
				AssociationID:                    9,
				CallingAE:                        "SCU",
				CalledAE:                         "SCP",
				LocalAddr:                        "127.0.0.1:1111",
				RemoteAddr:                       "127.0.0.1:2222",
				RemoteHost:                       "127.0.0.1",
				RemotePort:                       2222,
				ImplementationClassUID:           testSlogImplementationClassUID,
				ImplementationVersionName:        testSlogImplementationVersion,
				MaxPDULength:                     32768,
				AsyncOpsInvoked:                  4,
				AsyncOpsPerformed:                2,
				PresentationContextCount:         3,
				AcceptedPresentationContextCount: 2,
				PresentationContextSummary:       `[{"id":3,"abstract_syntax":"1.2.3.4","transfer_syntaxes":["1.2.840.10008.1.2.1"],"result":0}]`,
			},
			Direction:             DirectionOutbound,
			MessageID:             3,
			Command:               0x0030,
			CommandName:           "C-ECHO-RSP",
			Operation:             "C-ECHO",
			OperationID:           11,
			SOPClassUID:           "1.2.840.10008.1.1",
			TransferSyntax:        "1.2.840.10008.1.2.1",
			PresentationContextID: 3,
			StatusCode:            0,
			StatusState:           testSlogSuccessStatus,
			Duration:              12 * time.Millisecond,
			Outcome:               OutcomeSuccess,
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
	contexts, ok := attrs["presentation_contexts"].(string)
	if !ok || contexts != `[{"id":3,"abstract_syntax":"1.2.3.4","transfer_syntaxes":["1.2.840.10008.1.2.1"],"result":0}]` {
		t.Fatalf("presentation_contexts = %#v", attrs["presentation_contexts"])
	}
	for key, want := range map[string]any{
		"component":                           "network.service",
		"event":                               "request_completed",
		"connection_id":                       uint64(7),
		"association_id":                      uint64(9),
		"calling_ae":                          "SCU",
		"called_ae":                           "SCP",
		"local_addr":                          "127.0.0.1:1111",
		"remote_addr":                         "127.0.0.1:2222",
		"remote_host":                         "127.0.0.1",
		"remote_port":                         int64(2222),
		"implementation_class_uid":            testSlogImplementationClassUID,
		"implementation_version":              testSlogImplementationVersion,
		"max_pdu_length":                      uint64(32768),
		"async_ops_invoked":                   uint64(4),
		"async_ops_performed":                 uint64(2),
		"presentation_context_count":          uint64(3),
		"accepted_presentation_context_count": uint64(2),
		"direction":                           "outbound",
		"message_id":                          uint64(3),
		"command":                             uint64(0x0030),
		"command_name":                        "C-ECHO-RSP",
		"operation":                           "C-ECHO",
		"operation_id":                        uint64(11),
		"sop_class_uid":                       "1.2.840.10008.1.1",
		"transfer_syntax":                     "1.2.840.10008.1.2.1",
		"presentation_context_id":             uint64(3),
		"status":                              uint64(0),
		"status_state":                        testSlogSuccessStatus,
		"status_name":                         testSlogSuccessStatus,
		"outcome":                             "success",
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

func TestEmitSlogWritesAssociationNegotiationSummary(t *testing.T) {
	handler := &slogRecordHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	EmitSlog(context.Background(), "network.service", LogRecord{
		Level:   LevelInfo,
		Message: string(EventAssociationAccepted),
		Event: Event{
			Kind: EventAssociationAccepted,
			Association: Association{
				ImplementationClassUID:           testSlogImplementationClassUID,
				ImplementationVersionName:        testSlogImplementationVersion,
				MaxPDULength:                     16384,
				AsyncOpsInvoked:                  2,
				AsyncOpsPerformed:                1,
				PresentationContextCount:         4,
				AcceptedPresentationContextCount: 3,
			},
			Outcome: OutcomeSuccess,
		},
	})

	records := handler.snapshot()
	if len(records) != 1 {
		t.Fatalf("slog records = %d, want 1", len(records))
	}
	attrs := recordAttrs(records[0])
	for key, want := range map[string]any{
		"implementation_class_uid":            testSlogImplementationClassUID,
		"implementation_version":              testSlogImplementationVersion,
		"max_pdu_length":                      uint64(16384),
		"async_ops_invoked":                   uint64(2),
		"async_ops_performed":                 uint64(1),
		"presentation_context_count":          uint64(4),
		"accepted_presentation_context_count": uint64(3),
	} {
		if got, ok := attrs[key]; !ok || got != want {
			t.Errorf("association attribute %q = %#v, want %#v", key, got, want)
		}
	}
}

func TestEmitSlogWritesQueryRetrieveProgressAndParentOperation(t *testing.T) {
	handler := &slogRecordHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

	EmitSlog(context.Background(), "network.service", LogRecord{
		Level:   LevelDebug,
		Message: string(EventRequestPending),
		Event: Event{
			Kind:                   EventRequestPending,
			Command:                0x8010,
			CommandName:            "C-GET-RSP",
			Operation:              "C-GET",
			OperationID:            21,
			ParentOperationID:      0,
			QueryRetrieveLevel:     "STUDY",
			SubOperationCounts:     true,
			RemainingSubOperations: 2,
			CompletedSubOperations: 1,
			FailedSubOperations:    0,
			WarningSubOperations:   0,
			Outcome:                OutcomePending,
		},
	})
	EmitSlog(context.Background(), "network.service", LogRecord{
		Level:   LevelInfo,
		Message: string(EventRequestCompleted),
		Event: Event{
			Kind:              EventRequestCompleted,
			Command:           0x0001,
			CommandName:       "C-STORE-RQ",
			Operation:         "C-STORE",
			OperationID:       22,
			ParentOperationID: 21,
			Outcome:           OutcomeSuccess,
		},
	})

	records := handler.snapshot()
	if len(records) != 2 {
		t.Fatalf("slog records = %d, want 2", len(records))
	}
	progressAttrs := recordAttrs(records[0])
	for key, want := range map[string]any{
		"operation":               "C-GET",
		"command_name":            "C-GET-RSP",
		"query_retrieve_level":    "STUDY",
		"remaining_suboperations": uint64(2),
		"completed_suboperations": uint64(1),
	} {
		if got, ok := progressAttrs[key]; !ok || got != want {
			t.Errorf("progress attribute %q = %#v, want %#v", key, got, want)
		}
	}
	if records[0].Level != slog.LevelDebug {
		t.Errorf("progress level = %s, want DEBUG", records[0].Level)
	}
	childAttrs := recordAttrs(records[1])
	if childAttrs["parent_operation_id"] != uint64(21) {
		t.Errorf("parent_operation_id = %#v, want 21", childAttrs["parent_operation_id"])
	}
}

func TestEmitSlogUsesRecordMessageWhenEventKindIsUnavailable(t *testing.T) {
	handler := &slogRecordHandler{}
	if err := logging.Configure(logging.Config{Handler: handler}); err != nil {
		t.Fatalf("Configure() error = %v", err)
	}
	t.Cleanup(logging.Disable)

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
