// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// handleReceivedMessage processes a received DIMSE message.
// It dispatches the message to the appropriate handler or sends a default response.
func (s *Service) handleReceivedMessage(ctx context.Context, msg dimse.Message) error {
	// Check if this is a response to a pending request
	if resp, ok := msg.(dimse.Response); ok {
		return s.handleResponse(resp)
	}

	// Otherwise, it's a request that needs handling
	if req, ok := msg.(dimse.Request); ok {
		return s.handleRequest(ctx, req)
	}

	return fmt.Errorf("received message is neither request nor response: %T", msg)
}

// handleResponse routes a response to the corresponding pending request.
func (s *Service) handleResponse(resp dimse.Response) error {
	msgID := resp.MessageIDBeingRespondedTo()

	s.pendingRequestsMu.Lock()
	pending, exists := s.pendingRequests[msgID]
	s.pendingRequestsMu.Unlock()

	if !exists {
		// The request may have timed out or been cancelled locally. Late or
		// duplicate responses should not abort the whole association.
		return nil
	}

	// Send response to the waiting goroutine
	select {
	case pending.responseCh <- resp:
		// Successfully delivered
	case <-pending.cancelCh:
		// Request was cancelled
	case <-s.closeCh:
		// Service is closing
	}

	return nil
}

func (s *Service) registerActiveOperation(parent context.Context, messageID uint16) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	token := &struct{}{}

	s.activeOperationsMu.Lock()
	s.activeOperations[messageID] = &activeOperation{token: token, cancel: cancel}
	s.activeOperationsMu.Unlock()

	unregister := func() {
		s.activeOperationsMu.Lock()
		if op, exists := s.activeOperations[messageID]; exists && op.token == token {
			delete(s.activeOperations, messageID)
		}
		s.activeOperationsMu.Unlock()
		cancel()
	}

	return ctx, unregister
}

func (s *Service) cancelActiveOperation(messageID uint16) bool {
	s.activeOperationsMu.Lock()
	op, exists := s.activeOperations[messageID]
	if exists {
		delete(s.activeOperations, messageID)
	}
	s.activeOperationsMu.Unlock()

	if exists {
		op.cancel()
	}
	return exists
}

// handleRequest dispatches a request to the appropriate handler.
//
// All requests are dispatched in a goroutine so the recv loop stays free to
// route incoming responses. This is essential for C-GET, whose handler sends
// C-STORE sub-operations back over the same association and then waits for
// the C-STORE responses — blocking the recv loop would cause a deadlock.
// Dispatching all request types the same way keeps the design consistent.
func (s *Service) handleRequest(ctx context.Context, req dimse.Request) error {
	assoc := s.GetAssociation()
	if assoc != nil {
		pc, err := presentationContextForRequest(assoc, req)
		if err != nil {
			return err
		}
		if err := s.requireLocalRole(pc, false); err != nil {
			return err
		}
	}

	s.handlersMu.RLock()
	handlers := s.handlers
	s.handlersMu.RUnlock()

	s.requestWg.Add(1)
	go func() {
		defer s.requestWg.Done()
		if err := s.dispatchRequest(ctx, req, handlers); err != nil {
			_ = s.abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderNotSpecified, err, false)
		}
	}()
	return nil
}

// dispatchRequest routes a request to the appropriate handler by command field.
func (s *Service) dispatchRequest(ctx context.Context, req dimse.Request, handlers *Handlers) error {
	cmdField := dimse.CommandField(req.CommandField())
	switch cmdField {
	case dimse.CommandCEchoRQ:
		r, ok := req.(*dimse.CEchoRequest)
		if !ok {
			return fmt.Errorf("expected *CEchoRequest for command %s, got %T", cmdField, req)
		}
		return s.handleCEchoRequest(ctx, r, handlers)
	case dimse.CommandCStoreRQ:
		r, ok := req.(*dimse.CStoreRequest)
		if !ok {
			return fmt.Errorf("expected *CStoreRequest for command %s, got %T", cmdField, req)
		}
		return s.handleCStoreRequest(ctx, r, handlers)
	case dimse.CommandCFindRQ:
		r, ok := req.(*dimse.CFindRequest)
		if !ok {
			return fmt.Errorf("expected *CFindRequest for command %s, got %T", cmdField, req)
		}
		return s.handleCFindRequest(ctx, r, handlers)
	case dimse.CommandCMoveRQ:
		r, ok := req.(*dimse.CMoveRequest)
		if !ok {
			return fmt.Errorf("expected *CMoveRequest for command %s, got %T", cmdField, req)
		}
		return s.handleCMoveRequest(ctx, r, handlers)
	case dimse.CommandCGetRQ:
		r, ok := req.(*dimse.CGetRequest)
		if !ok {
			return fmt.Errorf("expected *CGetRequest for command %s, got %T", cmdField, req)
		}
		return s.handleCGetRequest(ctx, r, handlers)

	// DIMSE-N services
	case dimse.CommandNEventReportRQ:
		r, ok := req.(*dimse.NEventReportRequest)
		if !ok {
			return fmt.Errorf("expected *NEventReportRequest for command %s, got %T", cmdField, req)
		}
		return s.handleNEventReportRequest(ctx, r, handlers)
	case dimse.CommandNGetRQ:
		r, ok := req.(*dimse.NGetRequest)
		if !ok {
			return fmt.Errorf("expected *NGetRequest for command %s, got %T", cmdField, req)
		}
		return s.handleNGetRequest(ctx, r, handlers)
	case dimse.CommandNSetRQ:
		r, ok := req.(*dimse.NSetRequest)
		if !ok {
			return fmt.Errorf("expected *NSetRequest for command %s, got %T", cmdField, req)
		}
		return s.handleNSetRequest(ctx, r, handlers)
	case dimse.CommandNActionRQ:
		r, ok := req.(*dimse.NActionRequest)
		if !ok {
			return fmt.Errorf("expected *NActionRequest for command %s, got %T", cmdField, req)
		}
		return s.handleNActionRequest(ctx, r, handlers)
	case dimse.CommandNCreateRQ:
		r, ok := req.(*dimse.NCreateRequest)
		if !ok {
			return fmt.Errorf("expected *NCreateRequest for command %s, got %T", cmdField, req)
		}
		return s.handleNCreateRequest(ctx, r, handlers)
	case dimse.CommandNDeleteRQ:
		r, ok := req.(*dimse.NDeleteRequest)
		if !ok {
			return fmt.Errorf("expected *NDeleteRequest for command %s, got %T", cmdField, req)
		}
		return s.handleNDeleteRequest(ctx, r, handlers)

	// C-CANCEL-RQ — no response per PS3.7 §9.3.5
	case dimse.CommandCCancelRQ:
		r, ok := req.(*dimse.CCancelRequest)
		if !ok {
			return fmt.Errorf("expected *CCancelRequest for command %s, got %T", cmdField, req)
		}
		return s.handleCCancelRequest(r)

	default:
		return fmt.Errorf("unsupported DIMSE command: %s (0x%04X)", cmdField.String(), cmdField)
	}
}

// handleCEchoRequest handles a C-ECHO request.
func (s *Service) handleCEchoRequest(ctx context.Context, req *dimse.CEchoRequest, handlers *Handlers) error {
	var resp *dimse.CEchoResponse

	// Use custom handler if available
	if handlers != nil && handlers.CEchoHandler != nil {
		var err error
		resp, err = handlers.CEchoHandler(ctx, req)
		if err != nil {
			// Handler returned error - send failure response
			resp = dimse.NewCEchoResponseFromRequest(req, status.RefusedOutOfResources)
		}
	} else {
		// Default handler - always return success
		resp = dimse.NewCEchoResponseFromRequest(req, status.Success)
	}

	// Send response
	return s.Send(ctx, resp)
}

// handleCStoreRequest handles a C-STORE request.
func (s *Service) handleCStoreRequest(ctx context.Context, req *dimse.CStoreRequest, handlers *Handlers) error {
	var resp *dimse.CStoreResponse

	// Use custom handler if available
	if handlers != nil && handlers.CStoreHandler != nil {
		var err error
		resp, err = handlers.CStoreHandler(ctx, req)
		if err != nil {
			// Handler returned error - send failure response
			resp = dimse.NewCStoreResponseFromRequest(req, status.CStoreRefusedOutOfResources)
		}
	} else {
		// Default handler - return success (but don't actually store anything)
		resp = dimse.NewCStoreResponseFromRequest(req, status.Success)
	}

	// Send response
	return s.Send(ctx, resp)
}

// handleCFindRequest handles a C-FIND request.
func (s *Service) handleCFindRequest(ctx context.Context, req *dimse.CFindRequest, handlers *Handlers) error {
	handlerCtx, unregister := s.registerActiveOperation(ctx, req.MessageID())
	defer unregister()

	var responses []*dimse.CFindResponse

	// Use custom handler if available
	if handlers != nil && handlers.CFindHandler != nil {
		var err error
		responses, err = handlers.CFindHandler(handlerCtx, req)
		if err != nil {
			// Handler returned error - send failure response
			resp := dimse.NewCFindResponseFromRequest(req, status.CFindRefusedOutOfResources, nil)
			return s.Send(ctx, resp)
		}
	} else {
		// Default handler - return success with no results
		resp := dimse.NewCFindResponseFromRequest(req, status.Success, nil)
		responses = []*dimse.CFindResponse{resp}
	}

	// Send all responses
	for _, resp := range responses {
		if err := s.Send(ctx, resp); err != nil {
			return fmt.Errorf("failed to send C-FIND response: %w", err)
		}
	}

	return nil
}

// handleCMoveRequest handles a C-MOVE request.
func (s *Service) handleCMoveRequest(ctx context.Context, req *dimse.CMoveRequest, handlers *Handlers) error {
	handlerCtx, unregister := s.registerActiveOperation(ctx, req.MessageID())
	defer unregister()

	if handlers != nil && handlers.CMoveHandler != nil {
		op := newCMoveOperation(req, func(resp *dimse.CMoveResponse) error {
			return s.Send(ctx, resp)
		})
		if err := handlers.CMoveHandler(handlerCtx, op); err != nil {
			return s.Send(ctx, dimse.NewCMoveResponseFromRequest(req, status.CMoveFailedUnableToProcess))
		}
		return nil
	}
	// Default: no handler registered — return success with no operations.
	return s.Send(ctx, dimse.NewCMoveResponseFromRequest(req, status.Success))
}

// handleCGetRequest handles a C-GET request.
func (s *Service) handleCGetRequest(ctx context.Context, req *dimse.CGetRequest, handlers *Handlers) error {
	handlerCtx, unregister := s.registerActiveOperation(ctx, req.MessageID())
	defer unregister()

	if handlers != nil && handlers.CGetHandler != nil {
		op := newCGetOperation(
			req,
			func(storeCtx context.Context, storeReq *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
				return s.SendCStore(storeCtx, storeReq)
			},
			func(resp *dimse.CGetResponse) error {
				return s.Send(ctx, resp)
			},
		)
		if err := handlers.CGetHandler(handlerCtx, op); err != nil {
			return s.Send(ctx, dimse.NewCGetResponseFromRequest(req, status.CGetFailedUnableToProcess))
		}
		return nil
	}
	// Default: no handler registered — return success with no operations.
	return s.Send(ctx, dimse.NewCGetResponseFromRequest(req, status.Success))
}

// handleNEventReportRequest handles an N-EVENT-REPORT request.
func (s *Service) handleNEventReportRequest(ctx context.Context, req *dimse.NEventReportRequest, handlers *Handlers) error {
	var resp *dimse.NEventReportResponse
	if handlers != nil && handlers.NEventReportHandler != nil {
		var err error
		resp, err = handlers.NEventReportHandler(ctx, req)
		if err != nil {
			resp = dimse.NewNEventReportResponse(req.MessageID(), status.NEventReportFailureProcessingFailure,
				req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), req.EventTypeID(), nil)
		}
	} else {
		resp = dimse.NewNEventReportResponseSuccess(req.MessageID(),
			req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), req.EventTypeID(), nil)
	}
	return s.Send(ctx, resp)
}

// handleNGetRequest handles an N-GET request.
func (s *Service) handleNGetRequest(ctx context.Context, req *dimse.NGetRequest, handlers *Handlers) error {
	var resp *dimse.NGetResponse
	if handlers != nil && handlers.NGetHandler != nil {
		var err error
		resp, err = handlers.NGetHandler(ctx, req)
		if err != nil {
			resp = dimse.NewNGetResponse(req.MessageID(), status.NGetFailureProcessingFailure,
				req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), nil)
		}
	} else {
		resp = dimse.NewNGetResponseSuccess(req.MessageID(),
			req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), nil)
	}
	return s.Send(ctx, resp)
}

// handleNSetRequest handles an N-SET request.
func (s *Service) handleNSetRequest(ctx context.Context, req *dimse.NSetRequest, handlers *Handlers) error {
	var resp *dimse.NSetResponse
	if handlers != nil && handlers.NSetHandler != nil {
		var err error
		resp, err = handlers.NSetHandler(ctx, req)
		if err != nil {
			resp = dimse.NewNSetResponse(req.MessageID(), status.NSetFailureProcessingFailure,
				req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), nil)
		}
	} else {
		resp = dimse.NewNSetResponseSuccess(req.MessageID(),
			req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), nil)
	}
	return s.Send(ctx, resp)
}

// handleNActionRequest handles an N-ACTION request.
func (s *Service) handleNActionRequest(ctx context.Context, req *dimse.NActionRequest, handlers *Handlers) error {
	var resp *dimse.NActionResponse
	if handlers != nil && handlers.NActionHandler != nil {
		var err error
		resp, err = handlers.NActionHandler(ctx, req)
		if err != nil {
			resp = dimse.NewNActionResponse(req.MessageID(), status.NActionFailureProcessingFailure,
				req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), req.ActionTypeID(), nil)
		}
	} else {
		resp = dimse.NewNActionResponseSuccess(req.MessageID(),
			req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID(), req.ActionTypeID(), nil)
	}
	return s.Send(ctx, resp)
}

// handleNCreateRequest handles an N-CREATE request.
func (s *Service) handleNCreateRequest(ctx context.Context, req *dimse.NCreateRequest, handlers *Handlers) error {
	var resp *dimse.NCreateResponse
	if handlers != nil && handlers.NCreateHandler != nil {
		var err error
		resp, err = handlers.NCreateHandler(ctx, req)
		if err != nil {
			resp = dimse.NewNCreateResponse(req.MessageID(), status.NCreateFailureProcessingFailure,
				req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), nil)
		}
	} else {
		resp = dimse.NewNCreateResponseSuccess(req.MessageID(),
			req.AffectedSOPClassUID(), req.AffectedSOPInstanceUID(), nil)
	}
	return s.Send(ctx, resp)
}

// handleNDeleteRequest handles an N-DELETE request.
func (s *Service) handleNDeleteRequest(ctx context.Context, req *dimse.NDeleteRequest, handlers *Handlers) error {
	var resp *dimse.NDeleteResponse
	if handlers != nil && handlers.NDeleteHandler != nil {
		var err error
		resp, err = handlers.NDeleteHandler(ctx, req)
		if err != nil {
			resp = dimse.NewNDeleteResponse(req.MessageID(), status.NDeleteFailureProcessingFailure,
				req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID())
		}
	} else {
		resp = dimse.NewNDeleteResponseSuccess(req.MessageID(),
			req.RequestedSOPClassUID(), req.RequestedSOPInstanceUID())
	}
	return s.Send(ctx, resp)
}

// SetHandlers sets the DIMSE message handlers for this service.
// This should be called before starting the service.
//
// Example:
//
//	handlers := &service.Handlers{
//	    CEchoHandler: func(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
//	        // Custom C-ECHO handling
//	        return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
//	    },
//	    CStoreHandler: func(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
//	        // Store the DICOM dataset
//	        dataset := req.DataDataset()
//	        // ... save to disk or database ...
//	        return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil
//	    },
//	}
//	service.SetHandlers(handlers)
func (s *Service) SetHandlers(handlers *Handlers) {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()
	s.handlers = handlers
}

// GetHandlers returns the current DIMSE message handlers.
func (s *Service) GetHandlers() *Handlers {
	s.handlersMu.RLock()
	defer s.handlersMu.RUnlock()
	return s.handlers
}

// SetAssociationNegotiator sets the association negotiation callbacks.
// This is typically used in server mode to control which associations are accepted.
//
// Example:
//
//	negotiator := &MyNegotiator{}
//	service.SetAssociationNegotiator(negotiator)
func (s *Service) SetAssociationNegotiator(negotiator AssociationNegotiator) {
	s.callbacksMu.Lock()
	defer s.callbacksMu.Unlock()
	s.associationNegotiator = negotiator
}

// GetAssociationNegotiator returns the current association negotiator.
func (s *Service) GetAssociationNegotiator() AssociationNegotiator {
	s.callbacksMu.RLock()
	defer s.callbacksMu.RUnlock()
	return s.associationNegotiator
}

// SetAssociationReleaseHandler sets the association release callback.
// This is called when an A-RELEASE-RQ is received.
//
// Example:
//
//	handler := &MyReleaseHandler{}
//	service.SetAssociationReleaseHandler(handler)
func (s *Service) SetAssociationReleaseHandler(handler AssociationReleaseHandler) {
	s.callbacksMu.Lock()
	defer s.callbacksMu.Unlock()
	s.associationReleaseHandler = handler
}

// GetAssociationReleaseHandler returns the current association release handler.
func (s *Service) GetAssociationReleaseHandler() AssociationReleaseHandler {
	s.callbacksMu.RLock()
	defer s.callbacksMu.RUnlock()
	return s.associationReleaseHandler
}

// SetConnectionLifecycleHandler sets the connection lifecycle callbacks.
// This is used to monitor connection events (abort, close).
//
// Example:
//
//	handler := &MyLifecycleHandler{}
//	service.SetConnectionLifecycleHandler(handler)
func (s *Service) SetConnectionLifecycleHandler(handler ConnectionLifecycleHandler) {
	s.callbacksMu.Lock()
	defer s.callbacksMu.Unlock()
	s.connectionLifecycleHandler = handler
}

// GetConnectionLifecycleHandler returns the current connection lifecycle handler.
func (s *Service) GetConnectionLifecycleHandler() ConnectionLifecycleHandler {
	s.callbacksMu.RLock()
	defer s.callbacksMu.RUnlock()
	return s.connectionLifecycleHandler
}

// handleCCancelRequest handles a C-CANCEL-RQ request.
// Per DICOM PS3.7 §9.3.5, no response is sent for a cancel request.
// The method finds the pending request being cancelled and closes its
// cancel channel to signal the waiting goroutine.
func (s *Service) handleCCancelRequest(req *dimse.CCancelRequest) error {
	messageID := req.MessageIDBeingRespondedTo()
	if messageID == 0 {
		return fmt.Errorf("C-CANCEL-RQ has no MessageIDBeingRespondedTo")
	}

	s.pendingRequestsMu.Lock()
	pending, exists := s.pendingRequests[messageID]
	if exists {
		close(pending.cancelCh)
		delete(s.pendingRequests, messageID)
	}
	s.pendingRequestsMu.Unlock()
	s.cancelActiveOperation(messageID)

	// Unknown message IDs may already be complete or locally cancelled; that is
	// not an error per the DICOM standard.

	return nil
}
