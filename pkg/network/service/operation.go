// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// SubOperationResponder is the common interface shared by CMoveOperation and
// CGetOperation for reporting sub-operation progress to the SCU.
type SubOperationResponder interface {
	// SendPending sends an intermediate response (status 0xFF00) with current counts.
	// Call this after each sub-operation completes.
	SendPending(remaining, completed, failed, warning uint16) error

	// SendSuccess sends the final response with Success status (0x0000).
	SendSuccess() error

	// SendWarning sends the final response with Warning status (0xB000),
	// used when sub-operations completed but some failed.
	SendWarning() error

	// SendFailure sends a final failure response with the given status.
	// Common statuses: status.CMoveRefusedMoveDestinationUnknown (0xA801),
	// status.CMoveRefusedOutOfResourcesSubOps (0xA702),
	// status.CGetRefusedOutOfResourcesSubOps (0xA702).
	SendFailure(s *status.Status) error
}

// subOperationBase holds the shared state and implements SubOperationResponder.
// It is embedded by CMoveOperation and CGetOperation.
type subOperationBase struct {
	sendPending func(remaining, completed, failed, warning uint16) error
	sendFinal   func(s *status.Status) error
}

func (b *subOperationBase) SendPending(remaining, completed, failed, warning uint16) error {
	return b.sendPending(remaining, completed, failed, warning)
}

func (b *subOperationBase) SendSuccess() error                 { return b.sendFinal(status.Success) }
func (b *subOperationBase) SendWarning() error                 { return b.sendFinal(status.CMoveWarningSubOperationsComplete) }
func (b *subOperationBase) SendFailure(s *status.Status) error { return b.sendFinal(s) }

// ── C-MOVE ────────────────────────────────────────────────────────────────────

// CMoveOperation is the interface passed to a C-MOVE handler.
// It embeds SubOperationResponder for progress reporting and exposes
// read-only accessors for the original request.
type CMoveOperation interface {
	SubOperationResponder

	// Request returns the underlying C-MOVE request.
	Request() *dimse.CMoveRequest

	// QueryLevel returns the Query/Retrieve level.
	QueryLevel() dimse.QueryRetrieveLevel

	// MoveDestination returns the destination AE title.
	MoveDestination() string

	// Identifier returns the query identifier dataset.
	Identifier() *dataset.Dataset
}

// cMoveOperation is the concrete implementation of CMoveOperation.
type cMoveOperation struct {
	subOperationBase
	req *dimse.CMoveRequest
}

func (op *cMoveOperation) Request() *dimse.CMoveRequest           { return op.req }
func (op *cMoveOperation) QueryLevel() dimse.QueryRetrieveLevel   { return op.req.QueryLevel() }
func (op *cMoveOperation) MoveDestination() string                { return op.req.MoveDestination() }
func (op *cMoveOperation) Identifier() *dataset.Dataset           { return op.req.DataDataset() }

func newCMoveOperation(req *dimse.CMoveRequest, send func(*dimse.CMoveResponse) error) CMoveOperation {
	return &cMoveOperation{
		subOperationBase: subOperationBase{
			sendPending: func(remaining, completed, failed, warning uint16) error {
				return send(dimse.NewCMoveResponsePending(
					req.MessageID(), req.AffectedSOPClassUID(),
					remaining, completed, failed, warning,
				))
			},
			sendFinal: func(s *status.Status) error {
				return send(dimse.NewCMoveResponse(req.MessageID(), s, req.AffectedSOPClassUID()))
			},
		},
		req: req,
	}
}

// ── C-GET ─────────────────────────────────────────────────────────────────────

// CGetOperation is the interface passed to a C-GET handler.
// It embeds SubOperationResponder for progress reporting, exposes read-only
// accessors for the original request, and provides SendCStore to push each
// file back to the SCU over the same association.
type CGetOperation interface {
	SubOperationResponder

	// Request returns the underlying C-GET request.
	Request() *dimse.CGetRequest

	// QueryLevel returns the Query/Retrieve level.
	QueryLevel() dimse.QueryRetrieveLevel

	// Identifier returns the query identifier dataset.
	Identifier() *dataset.Dataset

	// SendCStore sends a C-STORE sub-operation back to the SCU over the same
	// association and returns the SCU's C-STORE response.
	SendCStore(ctx context.Context, ds *dataset.Dataset) (*dimse.CStoreResponse, error)
}

// cGetOperation is the concrete implementation of CGetOperation.
type cGetOperation struct {
	subOperationBase
	req        *dimse.CGetRequest
	sendCStore func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error)
}

func (op *cGetOperation) Request() *dimse.CGetRequest           { return op.req }
func (op *cGetOperation) QueryLevel() dimse.QueryRetrieveLevel  { return op.req.QueryLevel() }
func (op *cGetOperation) Identifier() *dataset.Dataset          { return op.req.DataDataset() }

func (op *cGetOperation) SendCStore(ctx context.Context, ds *dataset.Dataset) (*dimse.CStoreResponse, error) {
	if ds == nil {
		return nil, fmt.Errorf("dataset is nil")
	}
	req, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		return nil, fmt.Errorf("failed to create C-STORE request: %w", err)
	}
	return op.sendCStore(ctx, req)
}

func newCGetOperation(
	req *dimse.CGetRequest,
	sendCStore func(context.Context, *dimse.CStoreRequest) (*dimse.CStoreResponse, error),
	sendResp func(*dimse.CGetResponse) error,
) CGetOperation {
	return &cGetOperation{
		subOperationBase: subOperationBase{
			sendPending: func(remaining, completed, failed, warning uint16) error {
				return sendResp(dimse.NewCGetResponsePending(
					req.MessageID(), req.AffectedSOPClassUID(),
					remaining, completed, failed, warning,
				))
			},
			sendFinal: func(s *status.Status) error {
				return sendResp(dimse.NewCGetResponse(req.MessageID(), s, req.AffectedSOPClassUID()))
			},
		},
		req:        req,
		sendCStore: sendCStore,
	}
}
