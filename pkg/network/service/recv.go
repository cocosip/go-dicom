// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// recvLoop runs the receive goroutine that processes incoming PDUs after association establishment.
//
// Design Note:
// This implementation uses a two-phase approach (different from fo-dicom):
//
//	Phase 1: Association negotiation (before recvLoop starts)
//	  - Client: Connect() → SendAssociationRequest → ReceiveAssociationResponse
//	  - Server: ReceiveAssociationRequest → SendAssociationAccept/Reject
//	Phase 2: Data transfer (recvLoop running)
//	  - P-DATA-TF: DIMSE message exchange
//	  - A-RELEASE-RQ/RP: Association release
//	  - A-ABORT: Abnormal termination
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
	fragments := make(map[byte]*dimseMessageFragments)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-s.closeCh:
			return nil

		default:
			// Read PDU from connection
			rawPDU, err := transport.ReadPDUMaxLength(s.conn, readTimeoutFromContext(ctx, s.config.readTimeout), s.config.maxPDULength)
			if err != nil {
				switch {
				case errors.Is(err, transport.ErrUnrecognizedPDUType):
					// PS3.8 9.3.8: unrecognized PDU -> A-ABORT(unrecognized-PDU).
					_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderUnrecognizedPDU)
				case errors.Is(err, transport.ErrPDULengthExceeded):
					_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderNotSpecified)
				}
				return fmt.Errorf("failed to read PDU: %w", err)
			}
			s.emitPDUBytes(ctx, observability.DirectionInbound, rawPDU)

			// Handle different PDU types
			switch rawPDU.Type {
			case pdu.TypePDataTF:
				// Handle P-DATA-TF: Parse PDVs and process DIMSE messages
				if err := s.handlePDataTF(rawPDU, fragments); err != nil {
					if !s.GetState().IsTerminal() {
						_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderInvalidPDUParam)
					}
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
				s.handleReleaseResponse(ctx)
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

type dimseMessageFragments struct {
	command         []byte
	dataset         []byte
	commandComplete bool
	expectsDataset  bool
}

// validateCommandDataset enforces the PS3.7 rule that the command set only
// contains group 0000 elements and that the mandatory Command Field is present.
func validateCommandDataset(ds *dataset.Dataset) error {
	if ds == nil {
		return fmt.Errorf("command dataset is nil")
	}
	if _, ok := ds.GetString(tag.CommandField); !ok {
		return fmt.Errorf("missing Command Field (0000,0001)")
	}
	for _, elem := range ds.Elements() {
		if elem.Tag().Group() != 0x0000 {
			return fmt.Errorf("element %s is not a command element (group 0000)", elem.Tag())
		}
	}
	return nil
}

// messageSizeLimit returns the maximum reassembled DIMSE message size.
// Zero means unlimited.
func (s *Service) messageSizeLimit() uint32 {
	return s.config.maxMessageSize
}

// getTransferSyntaxForContext gets the transfer syntax for a presentation context ID.
func (s *Service) getTransferSyntaxForContext(contextID byte) *transfer.Syntax {
	assoc := s.GetAssociation()
	if assoc == nil {
		return nil
	}

	pc := assoc.FindPresentationContextByID(contextID)
	if pc == nil || !pc.IsAccepted() {
		return nil
	}

	return pc.AcceptedTransferSyntax
}

func (s *Service) acceptedTransferSyntaxForContext(contextID byte) (*transfer.Syntax, error) {
	assoc := s.GetAssociation()
	if assoc == nil {
		return nil, fmt.Errorf("presentation context ID %d is not negotiated", contextID)
	}
	pc := assoc.FindPresentationContextByID(contextID)
	if pc == nil || !pc.IsAccepted() {
		return nil, fmt.Errorf("presentation context ID %d is not accepted", contextID)
	}
	return pc.AcceptedTransferSyntax, nil
}

// processReceivedMessage decodes and dispatches a received DIMSE message.
func (s *Service) processReceivedMessage(commandData, datasetData []byte, transferSyntax *transfer.Syntax, contextID byte) error {
	// Decode command and data datasets
	commandDS, dataDS, err := DecodeDIMSEMessage(commandData, datasetData, transferSyntax)
	if err != nil {
		return fmt.Errorf("failed to decode DIMSE message: %w", err)
	}
	if err := validateCommandDataset(commandDS); err != nil {
		return fmt.Errorf("invalid command dataset: %w", err)
	}

	// Create DIMSE message object from datasets
	msg, err := createMessageFromDatasets(commandDS, dataDS)
	if err != nil {
		return fmt.Errorf("failed to create message from datasets: %w", err)
	}
	// Preserve optional diagnostic command elements (Error Comment/ID,
	// Offending Element) that the rebuilt response command set would drop.
	if resp, ok := msg.(dimse.Response); ok {
		copyOptionalCommandElements(commandDS, resp.CommandDataset())
	}

	// Store the presentation context ID so responses can be sent on the same context.
	// This is required by the DICOM standard: a response must use the same
	// presentation context ID as the corresponding request.
	msg.SetPresentationContextID(contextID)

	// Dispatch to handler
	if err := s.handleReceivedMessage(s.ctx, msg); err != nil {
		// If handling failed (e.g., couldn't send response due to missing presentation context),
		// send A-ABORT to notify the peer instead of silently failing. Do not
		// send a second abort when the handler already terminated the association.
		if !s.GetState().IsTerminal() {
			_ = s.Abort(s.ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderNotSpecified)
		}
		return fmt.Errorf("failed to handle received message: %w", err)
	}

	return nil
}

// handlePDataTF processes a P-DATA-TF PDU.
// It parses PDVs, accumulates fragments, and decodes complete DIMSE messages.
func (s *Service) handlePDataTF(rawPDU *pdu.RawPDU, fragments map[byte]*dimseMessageFragments) error {
	// Parse PDVs from P-DATA-TF PDU
	pdvs, err := ParsePDataTFPDU(rawPDU)
	if err != nil {
		return fmt.Errorf("failed to parse P-DATA-TF PDU: %w", err)
	}

	// Process each PDV
	for _, pdv := range pdvs {
		if pdv.PresentationContextID == 0 || pdv.PresentationContextID%2 == 0 {
			return fmt.Errorf("invalid presentation context ID %d", pdv.PresentationContextID)
		}
		transferSyntax, err := s.acceptedTransferSyntaxForContext(pdv.PresentationContextID)
		if err != nil {
			return err
		}
		state := fragments[pdv.PresentationContextID]
		if state == nil {
			state = &dimseMessageFragments{}
			fragments[pdv.PresentationContextID] = state
		}
		// Bound the total reassembled message size when configured, so a peer
		// cannot exhaust memory by streaming unlimited fragments for one message.
		// Unlimited by default: a single instance may legitimately be very large.
		if limit := s.messageSizeLimit(); limit > 0 &&
			uint64(len(state.command))+uint64(len(state.dataset))+uint64(len(pdv.Data)) > uint64(limit) {
			return fmt.Errorf("reassembled DIMSE message exceeds maximum size %d bytes on presentation context %d", limit, pdv.PresentationContextID)
		}
		if pdv.IsCommand {
			if state.commandComplete {
				return fmt.Errorf("received a new command before dataset completed on presentation context %d", pdv.PresentationContextID)
			}
			state.command = append(state.command, pdv.Data...)
		} else {
			if !state.commandComplete || !state.expectsDataset {
				return fmt.Errorf("received dataset before command completed on presentation context %d", pdv.PresentationContextID)
			}
			state.dataset = append(state.dataset, pdv.Data...)
		}
		if pdv.IsLastFragment {
			if pdv.IsCommand {
				hasData, err := s.inspectCommandHasDataset(state.command)
				if err != nil {
					return fmt.Errorf("failed to inspect CommandDataSetType: %w", err)
				}
				if !hasData {
					if err := s.processReceivedMessage(state.command, state.dataset, transferSyntax, pdv.PresentationContextID); err != nil {
						return fmt.Errorf("failed to process received message: %w", err)
					}
					delete(fragments, pdv.PresentationContextID)
				} else {
					state.commandComplete = true
					state.expectsDataset = true
				}
			} else {
				if err := s.processReceivedMessage(state.command, state.dataset, transferSyntax, pdv.PresentationContextID); err != nil {
					return fmt.Errorf("failed to process received message: %w", err)
				}
				delete(fragments, pdv.PresentationContextID)
			}
		}
	}

	return nil
}

// commandHasDataset checks the CommandDataSetType field to determine if a data dataset is present.
// Returns true if data is expected, false if CommandDataSetType = 0x0101 (NO_DATA_SET_PRESENT).
func (s *Service) commandHasDataset(commandData []byte) bool {
	hasData, err := s.inspectCommandHasDataset(commandData)
	return err == nil && hasData
}

func (s *Service) inspectCommandHasDataset(commandData []byte) (bool, error) {
	// Quick check: decode just enough to read CommandDataSetType (0000,0800)
	// This is a simplified check using Implicit VR Little Endian (command dataset encoding)
	r := bytes.NewReader(commandData)
	byteOrder := binary.LittleEndian

	for r.Len() > 0 {
		if r.Len() < 8 {
			return false, fmt.Errorf("CommandDataSetType not readable: truncated command header")
		}

		var group, elem uint16
		if err := binary.Read(r, byteOrder, &group); err != nil {
			return false, fmt.Errorf("CommandDataSetType not readable: %w", err)
		}
		if err := binary.Read(r, byteOrder, &elem); err != nil {
			return false, fmt.Errorf("CommandDataSetType not readable: %w", err)
		}

		var valueLength uint32
		if err := binary.Read(r, byteOrder, &valueLength); err != nil {
			return false, fmt.Errorf("CommandDataSetType not readable: %w", err)
		}

		if group == 0x0000 && elem == 0x0800 {
			// Found CommandDataSetType
			if valueLength < 2 || r.Len() < 2 {
				return false, fmt.Errorf("CommandDataSetType not readable: truncated value")
			}
			var dataSetType uint16
			if err := binary.Read(r, byteOrder, &dataSetType); err != nil {
				return false, fmt.Errorf("CommandDataSetType not readable: %w", err)
			}
			// PS3.7 Table E.1-1: 0x0101 means no Data Set; any other value
			// indicates a Data Set is present.
			return dataSetType != 0x0101, nil
		}

		// Skip value
		if valueLength > 0 && valueLength != 0xFFFFFFFF {
			if int64(valueLength) > int64(r.Len()) {
				return false, fmt.Errorf("CommandDataSetType not readable: truncated element value")
			}
			if _, err := r.Seek(int64(valueLength), 1); err != nil {
				return false, fmt.Errorf("CommandDataSetType not readable: %w", err)
			}
		} else if valueLength == 0xFFFFFFFF {
			return false, fmt.Errorf("CommandDataSetType not readable: undefined length in command set")
		}
	}

	return false, fmt.Errorf("CommandDataSetType not found")
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
	if s.config.keepConnectionOnPeerRelease {
		// Compatibility mode: reply with A-RELEASE-RP but leave the association
		// usable (state stays AssociationAccepted on the live socket).
		if err := s.acquireSendPause(ctx); err != nil {
			return fmt.Errorf("failed to wait for in-flight send: %w", err)
		}
		err := s.writeReleaseResponse(ctx)
		s.sendPause.Unlock()
		if err != nil {
			return fmt.Errorf("failed to send A-RELEASE-RP: %w", err)
		}
		s.emitAssociationObservation(ctx, observability.EventAssociationReleased, observability.DirectionOutbound, observability.OutcomeSuccess, nil)
		return nil
	}
	if err := s.SendReleaseResponse(ctx); err != nil {
		return fmt.Errorf("failed to send A-RELEASE-RP: %w", err)
	}

	return s.initiateClose(StateClosed, nil)
}

// handleReleaseResponse processes an A-RELEASE-RP PDU.
// This is received in response to our A-RELEASE-RQ, indicating the peer accepted the release.
// We update the state and notify the waiting graceful release path.
func (s *Service) handleReleaseResponse(ctx context.Context) {
	s.emitAssociationObservation(ctx, observability.EventAssociationReleased, observability.DirectionInbound, observability.OutcomeSuccess, nil)

	select {
	case s.releaseCh <- struct{}{}:
	default:
	}

	_ = s.initiateClose(StateClosed, nil)
}

// handleAbort processes an A-ABORT PDU.
// It calls the OnAbort callback (if set) for notification.
func (s *Service) handleAbort(ctx context.Context, abort *pdu.AAbort) {
	s.emitAssociationObservation(ctx, observability.EventAssociationAborted, observability.DirectionInbound, observability.OutcomeAborted, nil)
	// PS3.8: receiving A-ABORT moves the association into the Aborted terminal
	// state (not Closed).
	_ = s.setState(StateAborted)
	// Get lifecycle handler
	s.callbacksMu.RLock()
	lifecycleHandler := s.connectionLifecycleHandler
	s.callbacksMu.RUnlock()

	// Call OnAbort callback if set
	if lifecycleHandler != nil {
		lifecycleHandler.OnAbort(ctx, abort.Source, abort.Reason)
	}
}
