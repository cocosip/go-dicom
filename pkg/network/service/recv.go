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

// recvLoop runs the receive goroutine that processes incoming DIMSE messages.
// It reads P-DATA-TF PDUs from the network, extracts and reassembles PDVs,
// decodes them into DIMSE messages, and dispatches them to handlers or the receive queue.
//
// The loop continues until:
// - The context is cancelled
// - The closeCh channel is closed
// - An error occurs
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

			// Only process P-DATA-TF PDUs in recvLoop
			// Other PDU types (A-ASSOCIATE, A-RELEASE, A-ABORT) are handled elsewhere
			if rawPDU.Type != pdu.TypePDataTF {
				// TODO: Handle other PDU types or pass to appropriate handler
				continue
			}

			// Parse PDVs from P-DATA-TF PDU
			pdvs, err := ParsePDataTFPDU(rawPDU)
			if err != nil {
				return fmt.Errorf("failed to parse P-DATA-TF PDU: %w", err)
			}

			// Process each PDV
			for _, pdv := range pdvs {
				// If this is the first PDV or context ID changed, reset state
				if currentContextID == 0 {
					currentContextID = pdv.PresentationContextID
				} else if currentContextID != pdv.PresentationContextID {
					return fmt.Errorf("presentation context ID changed mid-message: %d -> %d",
						currentContextID, pdv.PresentationContextID)
				}

				// Accumulate fragments
				if pdv.IsCommand {
					commandFragments = append(commandFragments, pdv.Data...)
				} else {
					datasetFragments = append(datasetFragments, pdv.Data...)
				}

				// If this is the last fragment, decode the message
				if pdv.IsLastFragment {
					// Get transfer syntax from presentation context
					transferSyntax := s.getTransferSyntaxForContext(currentContextID)

					// Decode the complete message
					err := s.processReceivedMessage(commandFragments, datasetFragments, transferSyntax, currentContextID)
					if err != nil {
						return fmt.Errorf("failed to process received message: %w", err)
					}

					// Reset state for next message
					commandFragments = nil
					datasetFragments = nil
					currentContextID = 0
				}
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
