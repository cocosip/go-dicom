// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// recvLoop runs the receive goroutine that processes incoming PDUs after association establishment.
//
// Design Note:
// This implementation uses a two-phase approach (different from fo-dicom):
//   Phase 1: Association negotiation (before recvLoop starts)
//     - Client: Connect() → SendAssociationRequest → ReceiveAssociationResponse
//     - Server: ReceiveAssociationRequest → SendAssociationAccept/Reject
//   Phase 2: Data transfer (recvLoop running)
//     - P-DATA-TF: DIMSE message exchange
//     - A-RELEASE-RQ/RP: Association release
//     - A-ABORT: Abnormal termination
//
// PDU handling in recvLoop:
//   - P-DATA-TF (0x04): Parse PDVs, reassemble and decode DIMSE messages ✓
//   - A-RELEASE-RQ (0x05): Handle release request, send response ✓
//   - A-RELEASE-RP (0x06): Handle release response, close connection ✓
//   - A-ABORT (0x07): Handle abort notification ✓
//   - A-ASSOCIATE-RQ/AC/RJ (0x01/0x02/0x03): Protocol error (already handled in Phase 1)
//   - Unknown PDU: Send A-ABORT with UnrecognizedPDU reason
//
// The loop continues until:
// - The context is cancelled
// - The closeCh channel is closed
// - An error occurs (including receiving A-ABORT or A-RELEASE-RP)
func (s *Service) recvLoop(ctx context.Context) error {
	// Message reassembly state
	var commandFragments []byte
	var datasetFragments []byte
	var currentContextID byte

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-s.closeCh:
			return nil

		default:
			// Read PDU from connection
			rawPDU, err := transport.ReadPDU(s.conn, s.config.readTimeout)
			if err != nil {
				return fmt.Errorf("failed to read PDU: %w", err)
			}

			// Handle different PDU types
			switch rawPDU.Type {
			case pdu.TypePDataTF:
				// Handle P-DATA-TF: Parse PDVs and process DIMSE messages
				if err := s.handlePDataTF(rawPDU, &commandFragments, &datasetFragments, &currentContextID); err != nil {
					return fmt.Errorf("failed to handle P-DATA-TF: %w", err)
				}

			case pdu.TypeAReleaseRQ:
				// Handle A-RELEASE-RQ
				if err := s.handleReleaseRequest(ctx); err != nil {
					return fmt.Errorf("failed to handle A-RELEASE-RQ: %w", err)
				}

			case pdu.TypeAReleaseRP:
				// Handle A-RELEASE-RP (response to our release request)
				// This means the peer accepted our release request
				if err := s.handleReleaseResponse(ctx); err != nil {
					return fmt.Errorf("failed to handle A-RELEASE-RP: %w", err)
				}
				return nil

			case pdu.TypeAAbort:
				// Handle A-ABORT
				abort := &pdu.AAbort{}
				if err := abort.Decode(rawPDU); err != nil {
					return fmt.Errorf("failed to decode A-ABORT: %w", err)
				}
				s.handleAbort(ctx, abort)
				return fmt.Errorf("received A-ABORT: source=%d, reason=%d", abort.Source, abort.Reason)

			case pdu.TypeAAssociateRQ:
				// Handle A-ASSOCIATE-RQ (typically for server side)
				// This can happen in recvLoop if server is listening
				// The current implementation uses ReceiveAssociationRequest() before starting recvLoop,
				// so receiving it here might indicate a protocol error or re-negotiation attempt.
				// For now, treat as unexpected (following current architecture)
				_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderUnexpectedPDU)
				return fmt.Errorf("unexpected A-ASSOCIATE-RQ in recvLoop (association already established)")

			case pdu.TypeAAssociateAC, pdu.TypeAAssociateRJ:
				// Handle A-ASSOCIATE-AC/RJ (typically for client side)
				// The current implementation uses ReceiveAssociationResponse() before starting recvLoop,
				// so receiving it here might indicate a protocol error.
				// For now, treat as unexpected (following current architecture)
				_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderUnexpectedPDU)
				return fmt.Errorf("unexpected A-ASSOCIATE-AC/RJ in recvLoop (association already established)")

			default:
				// Unknown PDU type - send A-ABORT
				_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderUnrecognizedPDU)
				return fmt.Errorf("unexpected PDU type in recvLoop: 0x%02X", rawPDU.Type)
			}
		}
	}
}

// getTransferSyntaxForContext gets the transfer syntax for a presentation context ID.
func (s *Service) getTransferSyntaxForContext(contextID byte) *transfer.TransferSyntax {
	assoc := s.GetAssociation()
	if assoc == nil {
		return transfer.ExplicitVRLittleEndian // Default fallback
	}

	pc := assoc.FindPresentationContextByID(contextID)
	if pc == nil || pc.AcceptedTransferSyntax == nil {
		return transfer.ExplicitVRLittleEndian // Default fallback
	}

	return pc.AcceptedTransferSyntax
}

// processReceivedMessage decodes and dispatches a received DIMSE message.
func (s *Service) processReceivedMessage(commandData, datasetData []byte, transferSyntax *transfer.TransferSyntax, contextID byte) error {
	// Decode command and data datasets
	commandDS, dataDS, err := DecodeDIMSEMessage(commandData, datasetData, transferSyntax)
	if err != nil {
		return fmt.Errorf("failed to decode DIMSE message: %w", err)
	}

	// Create DIMSE message object from datasets
	msg, err := createMessageFromDatasets(commandDS, dataDS)
	if err != nil {
		return fmt.Errorf("failed to create message from datasets: %w", err)
	}

	// Dispatch to handler
	if err := s.handleReceivedMessage(s.ctx, msg); err != nil {
		return fmt.Errorf("failed to handle received message: %w", err)
	}

	return nil
}

// handlePDataTF processes a P-DATA-TF PDU.
// It parses PDVs, accumulates fragments, and decodes complete DIMSE messages.
func (s *Service) handlePDataTF(rawPDU *pdu.RawPDU, commandFragments, datasetFragments *[]byte, currentContextID *byte) error {
	// Parse PDVs from P-DATA-TF PDU
	pdvs, err := ParsePDataTFPDU(rawPDU)
	if err != nil {
		return fmt.Errorf("failed to parse P-DATA-TF PDU: %w", err)
	}

	// Process each PDV
	for _, pdv := range pdvs {
		// If this is the first PDV or context ID changed, reset state
		if *currentContextID == 0 {
			*currentContextID = pdv.PresentationContextID
		} else if *currentContextID != pdv.PresentationContextID {
			return fmt.Errorf("presentation context ID changed mid-message: %d -> %d",
				*currentContextID, pdv.PresentationContextID)
		}

		// Accumulate fragments
		if pdv.IsCommand {
			*commandFragments = append(*commandFragments, pdv.Data...)
		} else {
			*datasetFragments = append(*datasetFragments, pdv.Data...)
		}

		// If this is the last fragment, decode the message
		if pdv.IsLastFragment {
			// Get transfer syntax from presentation context
			transferSyntax := s.getTransferSyntaxForContext(*currentContextID)

			// Decode the complete message
			err := s.processReceivedMessage(*commandFragments, *datasetFragments, transferSyntax, *currentContextID)
			if err != nil {
				return fmt.Errorf("failed to process received message: %w", err)
			}

			// Reset state for next message
			*commandFragments = nil
			*datasetFragments = nil
			*currentContextID = 0
		}
	}

	return nil
}

// handleReleaseRequest processes an A-RELEASE-RQ PDU.
// It calls the OnAssociationRelease callback (if set) and sends an A-RELEASE-RP response.
func (s *Service) handleReleaseRequest(ctx context.Context) error {
	// Get release handler
	s.callbacksMu.RLock()
	releaseHandler := s.associationReleaseHandler
	s.callbacksMu.RUnlock()

	// Call OnAssociationRelease callback if set
	if releaseHandler != nil {
		if err := releaseHandler.OnAssociationRelease(ctx); err != nil {
			// Handler rejected the release, send abort instead
			return s.Abort(ctx, pdu.AbortSourceServiceUser, pdu.AbortReasonServiceUserNotSpecified)
		}
	}

	// Send A-RELEASE-RP
	if err := s.SendReleaseResponse(ctx); err != nil {
		return fmt.Errorf("failed to send A-RELEASE-RP: %w", err)
	}

	return nil
}

// handleReleaseResponse processes an A-RELEASE-RP PDU.
// This is received in response to our A-RELEASE-RQ, indicating the peer accepted the release.
// We update the service state and notify the lifecycle handler.
func (s *Service) handleReleaseResponse(ctx context.Context) error {
	// Update state to closed
	if err := s.setState(StateClosed); err != nil {
		// State transition error, but continue with cleanup
		// (we may already be in a closing state)
	}

	// Get lifecycle handler
	s.callbacksMu.RLock()
	lifecycleHandler := s.connectionLifecycleHandler
	s.callbacksMu.RUnlock()

	// Call OnConnectionClosed callback if set (nil error = normal close)
	if lifecycleHandler != nil {
		lifecycleHandler.OnConnectionClosed(ctx, nil)
	}

	return nil
}

// handleAbort processes an A-ABORT PDU.
// It calls the OnAbort callback (if set) for notification.
func (s *Service) handleAbort(ctx context.Context, abort *pdu.AAbort) {
	// Get lifecycle handler
	s.callbacksMu.RLock()
	lifecycleHandler := s.connectionLifecycleHandler
	s.callbacksMu.RUnlock()

	// Call OnAbort callback if set
	if lifecycleHandler != nil {
		lifecycleHandler.OnAbort(ctx, abort.Source, abort.Reason)
	}
}
