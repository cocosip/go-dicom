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

// handleRequest dispatches a request to the appropriate handler.
//
// All requests are dispatched in a goroutine so the recv loop stays free to
// route incoming responses. This is essential for C-GET, whose handler sends
// C-STORE sub-operations back over the same association and then waits for
// the C-STORE responses — blocking the recv loop would cause a deadlock.
// Dispatching all request types the same way keeps the design consistent.
func (s *Service) handleRequest(ctx context.Context, req dimse.Request) error {
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
	var responses []*dimse.CFindResponse

	// Use custom handler if available
	if handlers != nil && handlers.CFindHandler != nil {
		var err error
		responses, err = handlers.CFindHandler(ctx, req)
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
	if handlers != nil && handlers.CMoveHandler != nil {
		op := newCMoveOperation(req, func(resp *dimse.CMoveResponse) error {
			return s.Send(ctx, resp)
		})
		if err := handlers.CMoveHandler(ctx, op); err != nil {
			return s.Send(ctx, dimse.NewCMoveResponseFromRequest(req, status.CMoveFailedUnableToProcess))
		}
		return nil
	}
	// Default: no handler registered — return success with no operations.
	return s.Send(ctx, dimse.NewCMoveResponseFromRequest(req, status.Success))
}

// handleCGetRequest handles a C-GET request.
func (s *Service) handleCGetRequest(ctx context.Context, req *dimse.CGetRequest, handlers *Handlers) error {
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
		if err := handlers.CGetHandler(ctx, op); err != nil {
			return s.Send(ctx, dimse.NewCGetResponseFromRequest(req, status.CGetFailedUnableToProcess))
		}
		return nil
	}
	// Default: no handler registered — return success with no operations.
	return s.Send(ctx, dimse.NewCGetResponseFromRequest(req, status.Success))
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
