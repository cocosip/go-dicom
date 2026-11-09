// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"fmt"
	"time"
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
		return s.shutdownError()
	case <-s.ctx.Done():
		// Context was cancelled (normal shutdown)
		return s.shutdownError()
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

	s.startMu.Lock()
	if s.started {
		s.startMu.Unlock()
		return ErrServiceAlreadyStarted
	}
	s.started = true
	s.startMu.Unlock()

	// Start send loop
	go func() {
		err := s.sendLoop(s.ctx)
		if err != nil && err != context.Canceled {
			loopErr := fmt.Errorf("send loop error: %w", err)
			select {
			case s.errCh <- loopErr:
			default:
			}
			// Record the error and close so all pending goroutines unblock.
			_ = s.initiateClose(StateClosed, loopErr)
		}
	}()

	// Start receive loop
	go func() {
		err := s.recvLoop(s.ctx)
		if err != nil && err != context.Canceled {
			loopErr := fmt.Errorf("recv loop error: %w", err)
			select {
			case s.errCh <- loopErr:
			default:
			}
			// Record the error and close so all pending goroutines unblock.
			_ = s.initiateClose(StateClosed, loopErr)
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
//  1. Sends an A-ABORT PDU if connection is available
//  2. Closes the service (without changing Aborted state)
//
// Parameters:
//   - ctx: Context for the abort operation
//   - source: Abort source (0=service-user, 2=service-provider)
//   - reason: Abort reason (0=not-specified, 1=unrecognized-pdu, etc.)
func (s *Service) Abort(ctx context.Context, source, reason byte) error {
	return s.abort(ctx, source, reason, nil, true)
}

func (s *Service) abort(ctx context.Context, source, reason byte, recordErr error, wait bool) error {
	// Send A-ABORT PDU (ignore errors as connection may already be broken)
	// This will set state to Aborted
	_ = s.SendAbort(ctx, source, reason)

	// Close the service resources without changing state
	// (state is already Aborted from SendAbort)
	err := s.initiateClose(StateAborted, recordErr)
	if wait {
		s.waitForShutdown()
	}
	return err
}

// GracefulRelease attempts a graceful release of the DICOM association.
// It sends an A-RELEASE-RQ and waits for A-RELEASE-RP, then closes the service.
//
// The provided context bounds the whole shutdown: if it expires while waiting
// for the peer's A-RELEASE-RP or for in-flight request handlers to drain, the
// service is closed immediately and an error is returned. This keeps callers
// such as Client.Close from blocking for the full handler shutdown timeout.
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
		// Already closed (or closing): trigger shutdown without waiting for the
		// finalizer past the caller's budget.
		_ = s.initiateClose(StateClosed, nil)
		return s.awaitShutdown(ctx, "waiting for handler shutdown")
	}

	// Send A-RELEASE-RQ. SendReleaseRequest acquires sendPause internally, so no
	// DIMSE message is being written concurrently and the release cannot truncate
	// a multi-PDU message (PS3.8 requires release only after outstanding
	// operations complete). The acquisition respects ctx so a stalled peer cannot
	// block shutdown forever.
	err := s.SendReleaseRequest(ctx)
	if err != nil {
		// If send fails, abort without waiting for the handler drain: the send
		// already failed, so blocking here would defeat the bounded contract.
		_ = s.abort(ctx, 0, 0, nil, false)
		return fmt.Errorf("failed to send release request: %w", err)
	}

	select {
	case <-s.releaseCh:
		return s.awaitShutdown(ctx, "waiting for handler shutdown")
	case <-s.closeCh:
		return s.awaitShutdown(ctx, "waiting for handler shutdown")
	case <-ctx.Done():
		// Trigger shutdown without waiting for the finalizer: the caller's
		// budget is exhausted, so keep the same bounded-latency contract as
		// awaitShutdown. Close() would block for up to handlerShutdownTimeout.
		_ = s.initiateClose(StateClosed, nil)
		return fmt.Errorf("timed out waiting for A-RELEASE-RP: %w", ctx.Err())
	}
}

// awaitShutdown waits for the shutdown finalizer, but never past ctx: when the
// context expires the service is closed and the wait is abandoned so callers
// with a bounded shutdown budget (e.g. Client.Close) return promptly.
func (s *Service) awaitShutdown(ctx context.Context, operation string) error {
	select {
	case <-s.shutdownCh:
		return s.shutdownError()
	case <-ctx.Done():
		// Trigger shutdown without waiting for the finalizer (Close would block
		// for up to handlerShutdownTimeout and defeat the bounded wait).
		_ = s.initiateClose(StateClosed, nil)
		return fmt.Errorf("timed out %s: %w", operation, ctx.Err())
	}
}

// acquireSendPause acquires sendPause, waiting until any in-flight DIMSE message
// finishes or ctx is done. The non-blocking fast path avoids overhead in the
// common (idle send loop) case; the bounded poll only runs under contention.
func (s *Service) acquireSendPause(ctx context.Context) error {
	if s.sendPause.TryLock() {
		return nil
	}
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if s.sendPause.TryLock() {
				return nil
			}
		}
	}
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
