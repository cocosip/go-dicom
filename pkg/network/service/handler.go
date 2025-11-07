// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
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
		// No pending request for this response - may have been cancelled or timed out
		return fmt.Errorf("received response for unknown MessageID: %d", msgID)
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
func (s *Service) handleRequest(ctx context.Context, req dimse.Request) error {
	// Get handlers
	s.handlersMu.RLock()
	handlers := s.handlers
	s.handlersMu.RUnlock()

	// Dispatch based on message type
	switch r := req.(type) {
	case *dimse.CEchoRequest:
		return s.handleCEchoRequest(ctx, r, handlers)
	case *dimse.CStoreRequest:
		return s.handleCStoreRequest(ctx, r, handlers)
	case *dimse.CFindRequest:
		return s.handleCFindRequest(ctx, r, handlers)
	default:
		return fmt.Errorf("unsupported request type: %T", req)
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
			resp = dimse.NewCEchoResponseFromRequest(req, 0xA700) // Refused: Out of Resources
		}
	} else {
		// Default handler - always return success
		resp = dimse.NewCEchoResponseFromRequest(req, 0x0000)
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
			resp = dimse.NewCStoreResponseFromRequest(req, 0xA700) // Refused: Out of Resources
		}
	} else {
		// Default handler - return success (but don't actually store anything)
		resp = dimse.NewCStoreResponseFromRequest(req, 0x0000)
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
			resp := dimse.NewCFindResponseFromRequest(req, 0xA700, nil) // Refused: Out of Resources
			return s.Send(ctx, resp)
		}
	} else {
		// Default handler - return success with no results
		resp := dimse.NewCFindResponseFromRequest(req, 0x0000, nil)
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

// SetHandlers sets the message handlers for this service.
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

// GetHandlers returns the current message handlers.
func (s *Service) GetHandlers() *Handlers {
	s.handlersMu.RLock()
	defer s.handlersMu.RUnlock()
	return s.handlers
}
