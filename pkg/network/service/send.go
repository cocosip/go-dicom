// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// sendLoop runs the send goroutine that processes outgoing DIMSE messages.
// It reads from the sendQueue channel, encodes messages, fragments them into PDVs,
// creates P-DATA-TF PDUs, and writes them to the network connection.
//
// The loop continues until:
// - The context is cancelled
// - The closeCh channel is closed
// - An error occurs
func (s *Service) sendLoop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-s.closeCh:
			return nil

		case req := <-s.sendQueue:
			// Process send request
			err := s.sendMessage(req)
			if err != nil {
				// Send error back to caller
				select {
				case req.resultCh <- err:
				case <-ctx.Done():
					return ctx.Err()
				case <-s.closeCh:
					return nil
				}
			} else {
				// Send success
				select {
				case req.resultCh <- nil:
				case <-ctx.Done():
					return ctx.Err()
				case <-s.closeCh:
					return nil
				}
			}
		}
	}
}

// sendMessage sends a single DIMSE message over the network.
// Steps:
//  1. Get presentation context and transfer syntax from association
//  2. Encode the message into command and data datasets
//  3. Fragment the datasets into PDVs
//  4. Create P-DATA-TF PDU(s) from the PDVs
//  5. Write the PDU(s) to the network connection
func (s *Service) sendMessage(req *sendRequest) error {
	// Get association
	assoc := s.GetAssociation()
	if assoc == nil {
		return fmt.Errorf("no association available")
	}

	// Get SOP Class UID (Abstract Syntax) from message command dataset
	// Try AffectedSOPClassUID first (used by most operations)
	sopClassUID, ok := req.message.CommandDataset().GetString(tag.AffectedSOPClassUID)
	if !ok || sopClassUID == "" {
		// Some messages use RequestedSOPClassUID instead (N-GET, N-SET, N-DELETE, N-ACTION)
		sopClassUID, ok = req.message.CommandDataset().GetString(tag.RequestedSOPClassUID)
		if !ok || sopClassUID == "" {
			return fmt.Errorf("message has no SOP Class UID (neither AffectedSOPClassUID nor RequestedSOPClassUID)")
		}
	}

	// Find presentation context for this SOP Class
	pc := assoc.FindPresentationContextByAbstractSyntax(sopClassUID)
	if pc == nil {
		return fmt.Errorf("no accepted presentation context found for SOP Class UID: %s", sopClassUID)
	}

	// Get transfer syntax from presentation context
	transferSyntax := pc.AcceptedTransferSyntax
	if transferSyntax == nil {
		// Fall back to Explicit VR Little Endian
		transferSyntax = transfer.ExplicitVRLittleEndian
	}

	// Encode message
	commandData, datasetData, err := EncodeDIMSEMessage(req.message, transferSyntax)
	if err != nil {
		return fmt.Errorf("failed to encode message: %w", err)
	}

	// Use presentation context ID from association negotiation
	contextID := pc.ID

	// Fragment command data
	commandPDVs := FragmentData(commandData, contextID, true, s.config.maxPDULength)

	// Fragment dataset data (if present)
	var datasetPDVs []*PDV
	if len(datasetData) > 0 {
		datasetPDVs = FragmentData(datasetData, contextID, false, s.config.maxPDULength)
	}

	// Combine all PDVs
	allPDVs := append(commandPDVs, datasetPDVs...)

	// Group PDVs into P-DATA-TF PDUs
	// Pack multiple PDVs into one PDU when possible to reduce overhead
	pdvGroups := s.groupPDVsIntoPDUs(allPDVs)

	// Write each PDU to the connection
	for _, pdvGroup := range pdvGroups {
		pduData, err := CreatePDataTFPDU(pdvGroup)
		if err != nil {
			return fmt.Errorf("failed to create P-DATA-TF PDU: %w", err)
		}

		// Write PDU to connection
		err = transport.WritePDU(s.conn, s.config.writeTimeout, pduData)
		if err != nil {
			return fmt.Errorf("failed to write PDU: %w", err)
		}
	}

	return nil
}

// groupPDVsIntoPDUs groups PDVs into PDUs, packing multiple PDVs per PDU when possible.
// This reduces network overhead by minimizing the number of PDUs sent.
//
// Strategy:
//   - Pack as many PDVs as possible into each PDU without exceeding maxPDULength
//   - Each PDV already accounts for PDV header (6 bytes)
//   - PDU header adds 6 bytes (type + reserved + length)
//
// Note: This is a simple greedy packing algorithm. A more sophisticated approach
// could optimize for better packing efficiency.
func (s *Service) groupPDVsIntoPDUs(pdvs []*PDV) [][]*PDV {
	if len(pdvs) == 0 {
		return nil
	}

	const pduHeaderSize = 6 // PDU header: 1 byte type + 1 byte reserved + 4 bytes length

	var groups [][]*PDV
	var currentGroup []*PDV
	currentSize := pduHeaderSize

	for _, pdv := range pdvs {
		// Calculate size of this PDV when encoded
		// PDV encoding: 4 bytes length + 1 byte context ID + 1 byte header + data
		pdvSize := 4 + 1 + 1 + len(pdv.Data)

		// Check if adding this PDV would exceed max PDU length
		if currentSize+pdvSize > int(s.config.maxPDULength) && len(currentGroup) > 0 {
			// Start a new PDU
			groups = append(groups, currentGroup)
			currentGroup = []*PDV{pdv}
			currentSize = pduHeaderSize + pdvSize
		} else {
			// Add to current PDU
			currentGroup = append(currentGroup, pdv)
			currentSize += pdvSize
		}
	}

	// Add the last group
	if len(currentGroup) > 0 {
		groups = append(groups, currentGroup)
	}

	return groups
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
func (s *Service) Start() error {
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
func (s *Service) Err() <-chan error {
	return s.errCh
}
