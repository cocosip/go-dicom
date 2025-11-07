// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import "errors"

var (
	// ErrServiceClosed indicates the service has been closed.
	ErrServiceClosed = errors.New("service closed")

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

	// ErrMessageTooLarge indicates a message exceeds the maximum PDU length.
	ErrMessageTooLarge = errors.New("message exceeds maximum PDU length")

	// ErrInvalidPDU indicates a received PDU is invalid.
	ErrInvalidPDU = errors.New("invalid PDU")

	// ErrUnexpectedPDU indicates an unexpected PDU type was received.
	ErrUnexpectedPDU = errors.New("unexpected PDU type")
)
