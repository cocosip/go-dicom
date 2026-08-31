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
	attrs := make([]slog.Attr, 0, 14)
	if event.Association.ConnectionID != 0 {
		attrs = append(attrs, slog.Uint64("connection_id", uint64(event.Association.ConnectionID)))
	}
	if event.Association.AssociationID != 0 {
		attrs = append(attrs, slog.Uint64("association_id", uint64(event.Association.AssociationID)))
	}
	if event.Association.CallingAE != "" {
		attrs = append(attrs, slog.String("calling_ae", event.Association.CallingAE))
	}
	if event.Association.CalledAE != "" {
		attrs = append(attrs, slog.String("called_ae", event.Association.CalledAE))
	}
	if event.Direction != DirectionNone {
		attrs = append(attrs, slog.String("direction", string(event.Direction)))
	}
	if event.MessageID != 0 {
		attrs = append(attrs, slog.Uint64("message_id", uint64(event.MessageID)))
	}
	if event.Command != 0 {
		attrs = append(attrs, slog.Uint64("command", uint64(event.Command)))
	}
	if event.StatusState != "" || event.StatusCode != 0 {
		attrs = append(attrs, slog.Uint64("status", uint64(event.StatusCode)))
	}
	if event.StatusState != "" {
		attrs = append(attrs, slog.String("status_state", event.StatusState))
	}
	if event.Duration != 0 {
		attrs = append(attrs, slog.Duration("duration", event.Duration))
	}
	if event.Outcome != OutcomeNone {
		attrs = append(attrs, slog.String("outcome", string(event.Outcome)))
	}
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
