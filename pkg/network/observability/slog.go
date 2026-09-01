// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package observability

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/cocosip/go-dicom/pkg/logging"
)

// EmitSlog writes a network lifecycle record through go-dicom's configured
// slog logger. It emits protocol metadata only and never serializes a DIMSE
// dataset or PDU payload.
func EmitSlog(ctx context.Context, component string, record LogRecord) {
	level := slogLevel(record.Level)
	if !logging.Enabled(ctx, level) {
		return
	}

	event := record.Event
	attrs := make([]slog.Attr, 0, 28)
	attrs = appendAssociationAttrs(attrs, event.Association)
	attrs = appendDIMSEAttrs(attrs, event)
	attrs = appendStatusAttrs(attrs, event)
	if record.Code != "" {
		attrs = append(attrs, slog.String("code", record.Code))
	}
	if record.ItemType != 0 {
		attrs = append(attrs, slog.Uint64("item_type", uint64(record.ItemType)))
	}
	if event.Error != nil {
		attrs = append(attrs,
			slog.String("failure_stage", string(event.Kind)),
			slog.String("error_type", fmt.Sprintf("%T", event.Error)),
		)
	}

	logging.Emit(ctx, logging.Record{
		Level: level, Component: component, Event: record.Message, Message: record.Message, Attrs: attrs,
	})
}

func appendAssociationAttrs(attrs []slog.Attr, association Association) []slog.Attr {
	if association.ConnectionID != 0 {
		attrs = append(attrs, slog.Uint64("connection_id", uint64(association.ConnectionID)))
	}
	if association.AssociationID != 0 {
		attrs = append(attrs, slog.Uint64("association_id", uint64(association.AssociationID)))
	}
	if association.CallingAE != "" {
		attrs = append(attrs, slog.String("calling_ae", association.CallingAE))
	}
	if association.CalledAE != "" {
		attrs = append(attrs, slog.String("called_ae", association.CalledAE))
	}
	if association.LocalAddr != "" {
		attrs = append(attrs, slog.String("local_addr", association.LocalAddr))
	}
	if association.RemoteAddr != "" {
		attrs = append(attrs, slog.String("remote_addr", association.RemoteAddr))
	}
	if association.RemoteHost != "" {
		attrs = append(attrs, slog.String("remote_host", association.RemoteHost))
	}
	if association.RemotePort != 0 {
		attrs = append(attrs, slog.Int("remote_port", association.RemotePort))
	}
	if association.ImplementationClassUID != "" {
		attrs = append(attrs, slog.String("implementation_class_uid", association.ImplementationClassUID))
	}
	if association.ImplementationVersionName != "" {
		attrs = append(attrs, slog.String("implementation_version", association.ImplementationVersionName))
	}
	if association.MaxPDULength != 0 {
		attrs = append(attrs, slog.Uint64("max_pdu_length", uint64(association.MaxPDULength)))
	}
	if association.AsyncOpsInvoked != 0 {
		attrs = append(attrs, slog.Uint64("async_ops_invoked", uint64(association.AsyncOpsInvoked)))
	}
	if association.AsyncOpsPerformed != 0 {
		attrs = append(attrs, slog.Uint64("async_ops_performed", uint64(association.AsyncOpsPerformed)))
	}
	if association.PresentationContextCount != 0 {
		attrs = append(attrs, slog.Uint64("presentation_context_count", uint64(association.PresentationContextCount)))
	}
	if association.AcceptedPresentationContextCount != 0 {
		attrs = append(attrs, slog.Uint64("accepted_presentation_context_count", uint64(association.AcceptedPresentationContextCount)))
	}
	if association.PresentationContextSummary != "" {
		attrs = append(attrs, slog.String("presentation_contexts", association.PresentationContextSummary))
	}
	return attrs
}

func appendDIMSEAttrs(attrs []slog.Attr, event Event) []slog.Attr {
	if event.Direction != DirectionNone {
		attrs = append(attrs, slog.String("direction", string(event.Direction)))
	}
	if event.MessageID != 0 {
		attrs = append(attrs, slog.Uint64("message_id", uint64(event.MessageID)))
	}
	if event.Command != 0 {
		attrs = append(attrs, slog.Uint64("command", uint64(event.Command)))
	}
	if event.CommandName != "" {
		attrs = append(attrs, slog.String("command_name", event.CommandName))
	}
	if event.Operation != "" {
		attrs = append(attrs, slog.String("operation", event.Operation))
	}
	if event.OperationID != 0 {
		attrs = append(attrs, slog.Uint64("operation_id", event.OperationID))
	}
	if event.ParentOperationID != 0 {
		attrs = append(attrs, slog.Uint64("parent_operation_id", event.ParentOperationID))
	}
	if event.SOPClassUID != "" {
		attrs = append(attrs, slog.String("sop_class_uid", event.SOPClassUID))
	}
	if event.TransferSyntax != "" {
		attrs = append(attrs, slog.String("transfer_syntax", event.TransferSyntax))
	}
	if event.QueryRetrieveLevel != "" {
		attrs = append(attrs, slog.String("query_retrieve_level", event.QueryRetrieveLevel))
	}
	if event.MoveDestinationAE != "" {
		attrs = append(attrs, slog.String("move_destination_ae", event.MoveDestinationAE))
	}
	if event.PresentationContextID != 0 {
		attrs = append(attrs, slog.Uint64("presentation_context_id", uint64(event.PresentationContextID)))
	}
	if event.SubOperationCounts {
		attrs = append(attrs,
			slog.Uint64("remaining_suboperations", uint64(event.RemainingSubOperations)),
			slog.Uint64("completed_suboperations", uint64(event.CompletedSubOperations)),
			slog.Uint64("failed_suboperations", uint64(event.FailedSubOperations)),
			slog.Uint64("warning_suboperations", uint64(event.WarningSubOperations)),
		)
	}
	if event.ResultCount != 0 {
		attrs = append(attrs, slog.Uint64("result_count", event.ResultCount))
	}
	return attrs
}

func appendStatusAttrs(attrs []slog.Attr, event Event) []slog.Attr {
	if event.StatusState != "" || event.StatusCode != 0 {
		attrs = append(attrs,
			slog.Uint64("status", uint64(event.StatusCode)),
			slog.Uint64("status_code", uint64(event.StatusCode)),
			slog.String("status_hex", fmt.Sprintf("0x%04X", event.StatusCode)),
		)
	}
	if event.StatusState != "" {
		attrs = append(attrs,
			slog.String("status_state", event.StatusState),
			slog.String("status_name", event.StatusState),
		)
	}
	if event.Duration != 0 {
		attrs = append(attrs, slog.Duration("duration", event.Duration))
	}
	if event.Outcome != OutcomeNone {
		attrs = append(attrs, slog.String("outcome", string(event.Outcome)))
	}
	return attrs
}

func slogLevel(level Level) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
