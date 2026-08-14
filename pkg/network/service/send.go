// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

type dataDatasetMessage struct {
	dimse.Message
	data *dataset.Dataset
}

func (m *dataDatasetMessage) DataDataset() *dataset.Dataset {
	return m.data
}

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

	// Find presentation context.
	// For responses (or any message with an explicit presentation context ID set),
	// use the stored ID directly — the DICOM standard requires that a response is
	// sent on the same presentation context as its request.
	// For outgoing requests, fall back to an abstract-syntax lookup.
	message := req.message
	var pc *association.PresentationContext
	if req.message.CommandField() == uint16(dimse.CommandCStoreRQ) {
		var err error
		pc, message, err = prepareCStoreMessage(assoc, req.message)
		if err != nil {
			return err
		}
	}
	if pc == nil && req.message.PresentationContextID() != 0 {
		contextID := req.message.PresentationContextID()
		pc = assoc.FindPresentationContextByID(contextID)
	}
	if pc == nil {
		// Fall back: look up by SOP Class UID (abstract syntax).
		sopClassUID, ok := req.message.CommandDataset().GetString(tag.AffectedSOPClassUID)
		if !ok || sopClassUID == "" {
			// Some messages use RequestedSOPClassUID instead (N-GET, N-SET, N-DELETE, N-ACTION)
			sopClassUID, ok = req.message.CommandDataset().GetString(tag.RequestedSOPClassUID)
			if !ok || sopClassUID == "" {
				return fmt.Errorf("message has no SOP Class UID (neither AffectedSOPClassUID nor RequestedSOPClassUID)")
			}
		}
		pc = assoc.FindPresentationContextByAbstractSyntax(sopClassUID)
		if pc == nil {
			return fmt.Errorf("no accepted presentation context found for SOP Class UID: %s", sopClassUID)
		}
	}
	if dimse.CommandField(message.CommandField()).IsRequest() {
		if err := s.requireLocalRole(pc, true); err != nil {
			return err
		}
	}

	// Get transfer syntax from presentation context
	transferSyntax := pc.AcceptedTransferSyntax
	if transferSyntax == nil {
		// Fall back to Explicit VR Little Endian
		transferSyntax = transfer.ExplicitVRLittleEndian
	}

	// Encode message
	commandData, datasetData, err := EncodeDIMSEMessage(message, transferSyntax)
	if err != nil {
		return fmt.Errorf("failed to encode message: %w", err)
	}

	// Use presentation context ID from association negotiation
	contextID := pc.ID

	// Fragment command data
	maxPDULength := s.outgoingMaxPDULength()
	commandPDVs := FragmentData(commandData, contextID, true, maxPDULength)
	if len(commandPDVs) == 0 {
		return fmt.Errorf("maximum PDU length %d is too small to hold a command PDV", maxPDULength)
	}

	// Fragment dataset data (if present)
	var datasetPDVs []*PDV
	if len(datasetData) > 0 {
		datasetPDVs = FragmentData(datasetData, contextID, false, maxPDULength)
		if len(datasetPDVs) == 0 {
			return fmt.Errorf("maximum PDU length %d is too small to hold a dataset PDV", maxPDULength)
		}
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

func prepareCStoreMessage(assoc *association.Association, message dimse.Message) (*association.PresentationContext, dimse.Message, error) {
	dataDS := message.DataDataset()
	if dataDS == nil {
		return nil, message, nil
	}

	sopClassUID, _ := message.CommandDataset().GetString(tag.AffectedSOPClassUID)
	var candidates []*association.PresentationContext
	if contextID := message.PresentationContextID(); contextID != 0 {
		candidate := assoc.FindPresentationContextByID(contextID)
		if candidate == nil || !candidate.IsAccepted() || candidate.AcceptedTransferSyntax == nil {
			return nil, nil, fmt.Errorf("C-STORE presentation context ID %d is not accepted", contextID)
		}
		if candidate.AbstractSyntax != sopClassUID {
			return nil, nil, fmt.Errorf(
				"C-STORE presentation context ID %d is for SOP Class %s, not %s",
				contextID, candidate.AbstractSyntax, sopClassUID)
		}
		candidates = append(candidates, candidate)
	} else {
		for _, candidate := range assoc.GetAcceptedPresentationContexts() {
			if candidate.AbstractSyntax == sopClassUID && candidate.AcceptedTransferSyntax != nil {
				candidates = append(candidates, candidate)
			}
		}
	}

	sourceSyntax := dataDS.InternalTransferSyntax()
	if sourceSyntax == nil {
		if dataDS.Contains(tag.PixelData) {
			return nil, nil, fmt.Errorf(
				"cannot send C-STORE SOP Class %s with Pixel Data because its source transfer syntax is unknown",
				sopClassUID)
		}
		if len(candidates) > 0 {
			return candidates[0], message, nil
		}
		return nil, message, nil
	}

	for _, candidate := range candidates {
		if candidate.AcceptedTransferSyntax.UID().UID() == sourceSyntax.UID().UID() {
			return candidate, message, nil
		}
	}
	if !dataDS.Contains(tag.PixelData) && len(candidates) > 0 {
		return candidates[0], message, nil
	}

	manager := codec.GetDefaultManager()
	for _, candidate := range candidates {
		if !manager.CanTranscode(sourceSyntax, candidate.AcceptedTransferSyntax) {
			continue
		}
		transcoder, err := manager.CreateTranscoder(sourceSyntax, candidate.AcceptedTransferSyntax)
		if err != nil {
			return nil, nil, fmt.Errorf("create C-STORE transcoder from %s to %s: %w",
				sourceSyntax.UID().UID(), candidate.AcceptedTransferSyntax.UID().UID(), err)
		}
		transcoded, err := transcoder.Transcode(dataDS)
		if err != nil {
			return nil, nil, fmt.Errorf("transcode C-STORE dataset from %s to %s: %w",
				sourceSyntax.UID().UID(), candidate.AcceptedTransferSyntax.UID().UID(), err)
		}
		return candidate, &dataDatasetMessage{Message: message, data: transcoded}, nil
	}

	acceptedSyntaxes := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		acceptedSyntaxes = append(acceptedSyntaxes, candidate.AcceptedTransferSyntax.UID().UID())
	}
	return nil, nil, fmt.Errorf(
		"no accepted transfer syntax for C-STORE SOP Class %s is directly usable or transcodable from %s; accepted transfer syntaxes: %s",
		sopClassUID, sourceSyntax.UID().UID(), strings.Join(acceptedSyntaxes, ", "))
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
	maxPDULength := int(s.outgoingMaxPDULength())

	for _, pdv := range pdvs {
		// Calculate size of this PDV when encoded
		// PDV encoding: 4 bytes length + 1 byte context ID + 1 byte header + data
		pdvSize := 4 + 1 + 1 + len(pdv.Data)

		// Check if adding this PDV would exceed max PDU length
		if currentSize+pdvSize > maxPDULength && len(currentGroup) > 0 {
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

func (s *Service) outgoingMaxPDULength() uint32 {
	maxPDULength := s.config.maxPDULength
	if assoc := s.GetAssociation(); assoc != nil && assoc.MaxPDULength > 0 &&
		(maxPDULength == 0 || assoc.MaxPDULength < maxPDULength) {
		maxPDULength = assoc.MaxPDULength
	}
	if maxPDULength == 0 {
		// 0 means unlimited per DICOM standard (MaximumLengthReceived = 0 in A-ASSOCIATE-RQ).
		// Use 128 KB to avoid creating excessive numbers of tiny PDV fragments.
		return 128 * 1024
	}
	return maxPDULength
}
