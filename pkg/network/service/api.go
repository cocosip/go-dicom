// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// Send sends a DIMSE message synchronously and waits for completion.
// Returns an error if the message fails to send or if the service is closed.
//
// This is a blocking call that will wait until:
// - The message is successfully sent
// - An error occurs
// - The context is cancelled
// - The service is closed
//
// Example:
//
//	req := dimse.NewCEchoRequest()
//	req.SetMessageID(1)
//	err := service.Send(ctx, req)
//	if err != nil {
//	    log.Fatalf("Failed to send C-ECHO: %v", err)
//	}
func (s *Service) Send(ctx context.Context, msg dimse.Message) error {
	// Check if service is running
	if s.IsClosed() {
		return ErrServiceClosed
	}

	// Check state allows sending
	if !s.GetState().CanSendDIMSE() {
		return fmt.Errorf("cannot send DIMSE message in state %s", s.GetState())
	}

	// Create result channel
	resultCh := make(chan error, 1)

	// Create send request
	req := &sendRequest{
		message:  msg,
		resultCh: resultCh,
	}

	// Send to queue
	select {
	case s.sendQueue <- req:
		// Successfully queued
	case <-ctx.Done():
		return ctx.Err()
	case <-s.closeCh:
		return ErrServiceClosed
	}

	// Wait for result
	select {
	case err := <-resultCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-s.closeCh:
		return ErrServiceClosed
	}
}

// SendWithTimeout sends a DIMSE message with a timeout.
// This is a convenience wrapper around Send() that creates a context with timeout.
//
// Example:
//
//	req := dimse.NewCEchoRequest()
//	req.SetMessageID(1)
//	err := service.SendWithTimeout(req, 30*time.Second)
func (s *Service) SendWithTimeout(msg dimse.Message, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.Send(ctx, msg)
}

// SendCEcho sends a C-ECHO request and returns the response.
// C-ECHO is used to verify connectivity to a DICOM node.
//
// The request's MessageID will be automatically assigned if not already set.
//
// Example:
//
//	req := dimse.NewCEchoRequest()
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	resp, err := service.SendCEcho(ctx, req)
//	if err != nil {
//	    log.Fatalf("C-ECHO failed: %v", err)
//	}
//	if resp.IsSuccess() {
//	    log.Println("C-ECHO successful")
//	}
func (s *Service) SendCEcho(ctx context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
	// Get association to assign message ID
	assoc := s.GetAssociation()
	if assoc == nil {
		return nil, fmt.Errorf("no association available")
	}

	// Assign message ID from association
	msgID, err := assoc.AssignMessageID(req)
	if err != nil {
		return nil, fmt.Errorf("failed to assign message ID: %w", err)
	}

	// Register pending request to receive response
	respCh := make(chan dimse.Response, 1)
	s.registerPendingRequest(msgID, req, respCh)
	defer s.unregisterPendingRequest(msgID)

	// Send request
	if err := s.Send(ctx, req); err != nil {
		return nil, err
	}

	// Wait for response
	select {
	case respMsg := <-respCh:
		// Type assert to CEchoResponse
		resp, ok := respMsg.(*dimse.CEchoResponse)
		if !ok {
			return nil, fmt.Errorf("unexpected response type: %T", respMsg)
		}
		return resp, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.closeCh:
		return nil, ErrServiceClosed
	}
}

// SendCStore sends a C-STORE request to store a DICOM dataset.
//
// The request's MessageID will be automatically assigned if not already set.
//
// Example:
//
//	req, err := dimse.NewCStoreRequest(dicomDataset)
//	if err != nil {
//	    log.Fatalf("Failed to create C-STORE request: %v", err)
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//	defer cancel()
//	resp, err := service.SendCStore(ctx, req)
//	if err != nil {
//	    log.Fatalf("C-STORE failed: %v", err)
//	}
//	if resp.IsSuccess() {
//	    log.Println("C-STORE successful")
//	}
func (s *Service) SendCStore(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
	// Get association to assign message ID
	assoc := s.GetAssociation()
	if assoc == nil {
		return nil, fmt.Errorf("no association available")
	}

	// Assign message ID from association
	msgID, err := assoc.AssignMessageID(req)
	if err != nil {
		return nil, fmt.Errorf("failed to assign message ID: %w", err)
	}

	// Register pending request to receive response
	respCh := make(chan dimse.Response, 1)
	s.registerPendingRequest(msgID, req, respCh)
	defer s.unregisterPendingRequest(msgID)

	// Send request
	if err := s.Send(ctx, req); err != nil {
		return nil, err
	}

	// Wait for response
	select {
	case respMsg := <-respCh:
		// Type assert to CStoreResponse
		resp, ok := respMsg.(*dimse.CStoreResponse)
		if !ok {
			return nil, fmt.Errorf("unexpected response type: %T", respMsg)
		}
		return resp, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.closeCh:
		return nil, ErrServiceClosed
	}
}

// registerPendingRequest registers a pending request to receive its response.
func (s *Service) registerPendingRequest(messageID uint16, request dimse.Request, respCh chan dimse.Response) {
	s.pendingRequestsMu.Lock()
	defer s.pendingRequestsMu.Unlock()

	s.pendingRequests[messageID] = &pendingRequest{
		request:    request,
		responseCh: respCh,
		cancelCh:   make(chan struct{}),
	}
}

// unregisterPendingRequest removes a pending request.
func (s *Service) unregisterPendingRequest(messageID uint16) {
	s.pendingRequestsMu.Lock()
	defer s.pendingRequestsMu.Unlock()

	if pending, exists := s.pendingRequests[messageID]; exists {
		close(pending.cancelCh)
		delete(s.pendingRequests, messageID)
	}
}

// SendCFind sends a C-FIND request to query for DICOM objects.
//
// Returns a channel that will receive all C-FIND responses (both pending and final).
// The channel will be closed when the final response is received or an error occurs.
//
// The request's MessageID will be automatically assigned if not already set.
//
// Example:
//
//	query := dataset.New()
//	// Add query parameters to query
//	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, query)
//	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//	defer cancel()
//	resultCh, err := service.SendCFind(ctx, req)
//	if err != nil {
//	    log.Fatalf("C-FIND failed: %v", err)
//	}
//	for resp := range resultCh {
//	    if resp.IsPending() {
//	        // Process pending result
//	    } else {
//	        // Final response
//	    }
//	}
func (s *Service) SendCFind(ctx context.Context, req *dimse.CFindRequest) (<-chan *dimse.CFindResponse, error) {
	// Get association to assign message ID
	assoc := s.GetAssociation()
	if assoc == nil {
		return nil, fmt.Errorf("no association available")
	}

	// Assign message ID from association
	msgID, err := assoc.AssignMessageID(req)
	if err != nil {
		return nil, fmt.Errorf("failed to assign message ID: %w", err)
	}

	// Create result channel
	resultCh := make(chan *dimse.CFindResponse, 10)

	// Register pending request for multiple responses
	respCh := make(chan dimse.Response, 10)
	s.registerPendingRequest(msgID, req, respCh)

	// Start goroutine to handle responses
	go func() {
		defer close(resultCh)
		defer s.unregisterPendingRequest(msgID)

		for {
			select {
			case respMsg := <-respCh:
				// Type assert to CFindResponse
				resp, ok := respMsg.(*dimse.CFindResponse)
				if !ok {
					// Invalid response type - stop processing
					return
				}

				// Send to result channel
				select {
				case resultCh <- resp:
				case <-ctx.Done():
					return
				case <-s.closeCh:
					return
				}

				// If this is the final response, stop
				if !resp.IsPending() {
					return
				}

			case <-ctx.Done():
				return
			case <-s.closeCh:
				return
			}
		}
	}()

	// Send request
	if err := s.Send(ctx, req); err != nil {
		s.unregisterPendingRequest(msgID)
		return nil, err
	}

	return resultCh, nil
}
