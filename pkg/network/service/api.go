// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package service provides DICOM Service Class User (SCU) and Service Class Provider (SCP) implementations.
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
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
		message:   msg,
		resultCh:  resultCh,
		ctx:       ctx,
		lifecycle: s.lifecycleForSend(msg),
	}

	// Send to queue
	select {
	case s.sendQueue <- req:
		// Successfully queued
	case <-ctx.Done():
		return ctx.Err()
	case <-s.closeCh:
		return s.CloseError()
	}

	// Wait for result
	select {
	case err := <-resultCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-s.closeCh:
		return s.CloseError()
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

// SendCCancel sends a C-CANCEL-RQ for a pending C-FIND, C-MOVE, or C-GET request.
// C-CANCEL has no response; callers must pass the presentation context ID used
// by the original request because the cancel command has no SOP Class UID.
func (s *Service) SendCCancel(ctx context.Context, messageIDBeingRespondedTo uint16, presentationContextID byte) error {
	assoc := s.GetAssociation()
	if assoc == nil {
		return fmt.Errorf("no association available")
	}
	if presentationContextID == 0 {
		return fmt.Errorf("presentation context ID is required for C-CANCEL")
	}
	req := dimse.NewCCancelRequest(messageIDBeingRespondedTo)
	req.SetPresentationContextID(presentationContextID)
	return s.Send(ctx, req)
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
	return sendSimpleRequest[*dimse.CEchoRequest, *dimse.CEchoResponse](ctx, s, req)
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
	return sendSimpleRequest[*dimse.CStoreRequest, *dimse.CStoreResponse](ctx, s, req)
}

// SendNCreate sends an N-CREATE request and waits for its response.
func (s *Service) SendNCreate(ctx context.Context, req *dimse.NCreateRequest) (*dimse.NCreateResponse, error) {
	return sendSimpleRequest[*dimse.NCreateRequest, *dimse.NCreateResponse](ctx, s, req)
}

// SendNSet sends an N-SET request and waits for its response.
func (s *Service) SendNSet(ctx context.Context, req *dimse.NSetRequest) (*dimse.NSetResponse, error) {
	return sendSimpleRequest[*dimse.NSetRequest, *dimse.NSetResponse](ctx, s, req)
}

// SendNAction sends an N-ACTION request and waits for its response.
func (s *Service) SendNAction(ctx context.Context, req *dimse.NActionRequest) (*dimse.NActionResponse, error) {
	return sendSimpleRequest[*dimse.NActionRequest, *dimse.NActionResponse](ctx, s, req)
}

// SendNDelete sends an N-DELETE request and waits for its response.
func (s *Service) SendNDelete(ctx context.Context, req *dimse.NDeleteRequest) (*dimse.NDeleteResponse, error) {
	return sendSimpleRequest[*dimse.NDeleteRequest, *dimse.NDeleteResponse](ctx, s, req)
}

// registerPendingRequest registers a pending request to receive its response.
func (s *Service) registerPendingRequest(messageID uint16, request dimse.Request, respCh chan dimse.Response) *pendingRequest {
	s.pendingRequestsMu.Lock()
	defer s.pendingRequestsMu.Unlock()

	pending := &pendingRequest{
		request:    request,
		responseCh: respCh,
		cancelCh:   make(chan struct{}),
		lifecycle:  newRequestLifecycle(s, observability.DirectionOutbound, request),
	}
	s.pendingRequests[messageID] = pending
	return pending
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

// sendSimpleRequest is a generic helper for sending simple request-response DIMSE operations (C-ECHO, C-STORE).
// It handles message ID assignment, response waiting, and type assertion.
func sendSimpleRequest[Req dimse.Request, Resp dimse.Response](
	ctx context.Context,
	s *Service,
	req Req,
) (Resp, error) {
	var zero Resp
	assoc := s.GetAssociation()
	if assoc == nil {
		return zero, fmt.Errorf("no association available")
	}
	release, err := s.acquireAsyncOperation(ctx)
	if err != nil {
		return zero, err
	}
	defer release()
	msgID, err := assoc.AssignMessageID(req)
	if err != nil {
		return zero, fmt.Errorf("failed to assign message ID: %w", err)
	}
	respCh := make(chan dimse.Response, 1)
	pending := s.registerPendingRequest(msgID, req, respCh)
	defer s.unregisterPendingRequest(msgID)
	if err := s.Send(ctx, req); err != nil {
		pending.lifecycle.finishContext(ctx, err)
		return zero, err
	}
	select {
	case respMsg := <-respCh:
		resp, ok := respMsg.(Resp)
		if !ok {
			return zero, fmt.Errorf("unexpected response type: %T", respMsg)
		}
		return resp, nil
	case <-ctx.Done():
		pending.lifecycle.finishContext(ctx, ctx.Err())
		return zero, ctx.Err()
	case <-s.closeCh:
		err := s.CloseError()
		pending.lifecycle.finishError(ctx, err)
		return zero, err
	}
}

// pendingResponse is an interface for DIMSE responses that can be pending or final.
type pendingResponse interface {
	dimse.Response
	IsPending() bool
}

// sendRequestWithProgress is a generic helper for sending requests that return progressive responses (C-FIND, C-GET, C-MOVE).
// It handles message ID assignment, response collection, and channel management.
func sendRequestWithProgress[Req dimse.Request, Resp pendingResponse](
	ctx context.Context,
	s *Service,
	req Req,
	errMsg string,
) (<-chan Resp, error) {
	assoc := s.GetAssociation()
	if assoc == nil {
		return nil, fmt.Errorf("no association available")
	}
	release, err := s.acquireAsyncOperation(ctx)
	if err != nil {
		return nil, err
	}
	msgID, err := assoc.AssignMessageID(req)
	if err != nil {
		release()
		return nil, fmt.Errorf("failed to assign message ID: %w", err)
	}
	resultCh := make(chan Resp, 10)
	respCh := make(chan dimse.Response, 10)
	pending := s.registerPendingRequest(msgID, req, respCh)

	if err := s.Send(ctx, req); err != nil {
		pending.lifecycle.finishContext(ctx, err)
		s.unregisterPendingRequest(msgID)
		close(resultCh)
		release()
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}

	go func() {
		defer close(resultCh)
		defer s.unregisterPendingRequest(msgID)
		defer release()
		for {
			select {
			case respMsg, ok := <-respCh:
				if !ok {
					return
				}
				resp, ok := respMsg.(Resp)
				if !ok {
					return
				}
				select {
				case resultCh <- resp:
				case <-ctx.Done():
					pending.lifecycle.finishContext(ctx, ctx.Err())
					return
				case <-s.closeCh:
					pending.lifecycle.finishError(ctx, s.CloseError())
					return
				}
				if !resp.IsPending() {
					return
				}
			case <-ctx.Done():
				pending.lifecycle.finishContext(ctx, ctx.Err())
				return
			case <-s.closeCh:
				pending.lifecycle.finishError(ctx, s.CloseError())
				return
			}
		}
	}()
	return resultCh, nil
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
	release, err := s.acquireAsyncOperation(ctx)
	if err != nil {
		return nil, err
	}

	// Assign message ID from association
	msgID, err := assoc.AssignMessageID(req)
	if err != nil {
		release()
		return nil, fmt.Errorf("failed to assign message ID: %w", err)
	}

	// Create result channel
	resultCh := make(chan *dimse.CFindResponse, 10)

	// Register pending request for multiple responses
	respCh := make(chan dimse.Response, 10)
	pending := s.registerPendingRequest(msgID, req, respCh)

	// Send request before starting the response goroutine. The buffered response
	// channel still accepts fast responses, and send failures cannot race with
	// the goroutine's channel close.
	if err := s.Send(ctx, req); err != nil {
		pending.lifecycle.finishContext(ctx, err)
		s.unregisterPendingRequest(msgID)
		close(resultCh)
		release()
		return nil, err
	}

	// Start goroutine to handle responses
	go func() {
		defer close(resultCh)
		defer s.unregisterPendingRequest(msgID)
		defer release()

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
					pending.lifecycle.finishContext(ctx, ctx.Err())
					return
				case <-s.closeCh:
					pending.lifecycle.finishError(ctx, s.CloseError())
					return
				}

				// If this is the final response, stop
				if !resp.IsPending() {
					return
				}

			case <-ctx.Done():
				pending.lifecycle.finishContext(ctx, ctx.Err())
				return
			case <-s.closeCh:
				pending.lifecycle.finishError(ctx, s.CloseError())
				return
			}
		}
	}()

	return resultCh, nil
}

// SendCMove sends a C-MOVE request and returns a channel that will receive responses.
// C-MOVE is used to request that an SCP move DICOM instances to a specified destination.
//
// The returned channel will receive multiple responses:
//   - Pending (0xFF00) responses with sub-operation progress updates
//   - Final response (0x0000 = success, 0xB000 = warning, or error codes)
//
// The channel will be closed after the final response is received.
// The caller should read from the channel until it's closed.
//
// Common error codes:
//   - 0xA701: Out of resources - unable to calculate number of matches
//   - 0xA702: Out of resources - unable to perform sub-operations
//   - 0xA801: Move destination unknown
//   - 0xA900: Identifier does not match SOP Class
//   - 0xC000: Unable to process
func (s *Service) SendCMove(ctx context.Context, req *dimse.CMoveRequest) (<-chan *dimse.CMoveResponse, error) {
	return sendRequestWithProgress[*dimse.CMoveRequest, *dimse.CMoveResponse](ctx, s, req, "failed to send C-MOVE request")
}

// SendCGet sends a C-GET request and returns a channel that will receive responses.
// C-GET is used to request that an SCP retrieve DICOM instances and send them back
// to the requestor over the same association.
//
// The returned channel will receive multiple responses:
//   - Pending (0xFF00) responses with sub-operation progress updates
//   - Final response (0x0000 = success, 0xB000 = warning, or error codes)
//
// The channel will be closed after the final response is received.
// The caller should read from the channel until it's closed.
//
// Note: C-GET responses are accompanied by C-STORE sub-operations where the SCP
// sends the requested instances. The caller must have a C-STORE handler configured
// to receive the instances.
//
// Common error codes:
//   - 0xA701: Out of resources - unable to calculate number of matches
//   - 0xA702: Out of resources - unable to perform sub-operations
//   - 0xA900: Identifier does not match SOP Class
//   - 0xC000: Unable to process
func (s *Service) SendCGet(ctx context.Context, req *dimse.CGetRequest) (<-chan *dimse.CGetResponse, error) {
	return sendRequestWithProgress[*dimse.CGetRequest, *dimse.CGetResponse](ctx, s, req, "failed to send C-GET request")
}
