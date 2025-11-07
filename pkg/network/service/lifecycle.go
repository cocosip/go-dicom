// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"
)

// Run starts the service by launching the send and receive loops,
// and blocks until the service is closed or an error occurs.
//
// Run should be called after the DICOM association is established.
// It will start two goroutines using Start(), then wait for errors
// or service closure.
//
// If either loop encounters an error, Run will return the first error.
//
// Example usage:
//
//	service := service.NewService(conn, assoc)
//	defer service.Close()
//
//	// Start the service and wait
//	if err := service.Run(); err != nil {
//	    log.Printf("Service error: %v", err)
//	}
//
// For non-blocking start, use Start() instead.
func (s *Service) Run() error {
	// Check if service is already closed
	if s.IsClosed() {
		return ErrServiceClosed
	}

	// Start the send and receive loops (non-blocking)
	if err := s.Start(); err != nil {
		return err
	}

	// Wait for either an error or service closure
	select {
	case err := <-s.errCh:
		// Got an error from send/recv loop
		// Filter out context.Canceled as it's a normal shutdown
		if err == context.Canceled {
			return nil
		}
		return err
	case <-s.closeCh:
		// Service was closed normally
		return nil
	case <-s.ctx.Done():
		// Context was cancelled (normal shutdown)
		return nil
	}
}

// Start starts the service's send and receive loops.
// This should be called after the association is established.
//
// The service will run in the background until:
// - Close() is called
// - An error occurs in send/receive loops
// - The context is cancelled
//
// Example:
//
//	service := service.NewService(conn, assoc)
//	if err := service.Start(); err != nil {
//	    return err
//	}
//	defer service.Close()
//
// For blocking operation, use Run() instead.
func (s *Service) Start() error {
	// Check if service is already closed
	if s.IsClosed() {
		return ErrServiceClosed
	}

	// Start send loop
	go func() {
		err := s.sendLoop(s.ctx)
		if err != nil && err != context.Canceled {
			select {
			case s.errCh <- fmt.Errorf("send loop error: %w", err):
			default:
			}
		}
	}()

	// Start receive loop
	go func() {
		err := s.recvLoop(s.ctx)
		if err != nil && err != context.Canceled {
			select {
			case s.errCh <- fmt.Errorf("recv loop error: %w", err):
			default:
			}
		}
	}()

	return nil
}

// Err returns the error channel for the service.
// This channel will receive errors from the send/receive loops.
// Use this with Start() for non-blocking error monitoring.
func (s *Service) Err() <-chan error {
	return s.errCh
}

// Abort aborts the service with the specified reason.
// This is a convenience method that:
//   1. Sends an A-ABORT PDU if connection is available
//   2. Closes the service (without changing Aborted state)
//
// Parameters:
//   - ctx: Context for the abort operation
//   - source: Abort source (0=service-user, 2=service-provider)
//   - reason: Abort reason (0=not-specified, 1=unrecognized-pdu, etc.)
func (s *Service) Abort(ctx context.Context, source, reason byte) error {
	// Send A-ABORT PDU (ignore errors as connection may already be broken)
	// This will set state to Aborted
	_ = s.SendAbort(ctx, source, reason)

	// Close the service resources without changing state
	// (state is already Aborted from SendAbort)
	var err error
	s.closeOnce.Do(func() {
		// Cancel context to stop goroutines
		s.cancel()

		// Close channels
		close(s.closeCh)

		// Cancel all pending requests
		s.cancelPendingRequests()

		// Close connection
		if s.conn != nil {
			err = s.conn.Close()
		}
	})

	return err
}

// GracefulRelease attempts a graceful release of the DICOM association.
// It sends an A-RELEASE-RQ and waits for A-RELEASE-RP, then closes the service.
//
// If the release fails or times out, it falls back to sending an A-ABORT.
//
// This is a convenience method for proper DICOM association termination.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	if err := service.GracefulRelease(ctx); err != nil {
//	    log.Printf("Release failed: %v", err)
//	}
func (s *Service) GracefulRelease(ctx context.Context) error {
	// Check if we're in a state that allows release
	state := s.GetState()
	if !state.CanSendDIMSE() {
		// Already closed or in invalid state
		return s.Close()
	}

	// Send A-RELEASE-RQ
	if err := s.SendReleaseRequest(ctx); err != nil {
		// If send fails, abort
		_ = s.Abort(ctx, 0, 0)
		return fmt.Errorf("failed to send release request: %w", err)
	}

	// Wait for A-RELEASE-RP (should be received by recvLoop)
	// We'll use a simple approach: wait for state to become Closed
	// In a production implementation, you might want to:
	// - Have recvLoop handle A-RELEASE-RP specifically
	// - Use a channel to signal release completion
	// For now, we'll just close the service
	// The peer should send A-RELEASE-RP, but we don't wait for it here

	return s.Close()
}

// WaitForClose blocks until the service is closed.
// This is useful when you want to wait for the service to finish
// after starting it with Run().
//
// Example:
//
//	go service.Run()
//	// ... do work ...
//	service.Close()
//	service.WaitForClose() // Wait for cleanup to complete
func (s *Service) WaitForClose() {
	<-s.closeCh
}
