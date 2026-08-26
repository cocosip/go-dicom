// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

var (
	// ErrServiceClosed indicates the service has been closed.
	ErrServiceClosed = errors.New("service closed")

	// ErrServiceAlreadyStarted indicates Start/Run was called more than once.
	ErrServiceAlreadyStarted = errors.New("service already started")

	// ErrInvalidState indicates an operation was attempted in an invalid state.
	ErrInvalidState = errors.New("invalid state for this operation")

	// ErrNoAssociation indicates no association is established.
	ErrNoAssociation = errors.New("no association established")

	// ErrAssociationRejected indicates the association was rejected.
	ErrAssociationRejected = errors.New("association rejected")

	// ErrAssociationAborted indicates the association was aborted.
	ErrAssociationAborted = errors.New("association aborted")

	// ErrTimeout indicates an operation timed out.
	ErrTimeout = errors.New("operation timed out")

	// ErrRequestTimeout indicates that an individual DIMSE request did not
	// receive a response within its configured response idle timeout.
	ErrRequestTimeout = errors.New("DIMSE request timeout")

	// ErrMessageTooLarge indicates a message exceeds the maximum PDU length.
	ErrMessageTooLarge = errors.New("message exceeds maximum PDU length")

	// ErrInvalidPDU indicates a received PDU is invalid.
	ErrInvalidPDU = errors.New("invalid PDU")

	// ErrUnexpectedPDU indicates an unexpected PDU type was received.
	ErrUnexpectedPDU = errors.New("unexpected PDU type")

	// ErrCFindOperationCompleted indicates that a C-FIND handler attempted to
	// send another response after its final response.
	ErrCFindOperationCompleted = errors.New("C-FIND operation already completed")

	// ErrCFindHandlerConflict indicates that both legacy and streaming C-FIND
	// handlers were configured for the same service.
	ErrCFindHandlerConflict = errors.New("C-FIND stream and legacy handlers cannot both be configured")

	// ErrHandlerShutdownTimeout indicates that Service.Close stopped waiting
	// because an inbound request handler did not exit before the configured
	// handler shutdown timeout.
	ErrHandlerShutdownTimeout = errors.New("request handler shutdown timeout")
)

// RequestTimeoutError identifies the timed-out outgoing DIMSE request.
type RequestTimeoutError struct {
	MessageID uint16
	Command   dimse.CommandField
	Timeout   time.Duration
}

func (e *RequestTimeoutError) Error() string {
	return fmt.Sprintf("%s request %d timed out after %s", e.Command, e.MessageID, e.Timeout)
}

// Unwrap makes request timeouts available through errors.Is.
func (e *RequestTimeoutError) Unwrap() error {
	return ErrRequestTimeout
}
