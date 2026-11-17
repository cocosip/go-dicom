// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transcode"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
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
			// Process send request. Re-check the state here because a message
			// may have been queued before A-RELEASE started (PS3.8 Sta7). The
			// sendPause lock keeps a release request from being injected between
			// two PDUs of a multi-PDU message.
			s.sendPause.Lock()
			err := s.checkSendState()
			if err == nil {
				err = s.sendMessage(req)
			}
			s.sendPause.Unlock()
			s.observeSentMessage(req.message, req.lifecycle, err)
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

// checkSendState rejects DIMSE writes outside the established state.
func (s *Service) checkSendState() error {
	if state := s.GetState(); state != StateAssociationAccepted {
		return fmt.Errorf("cannot send DIMSE in state %s", state)
	}
	return nil
}

// sendMessage sends a single DIMSE message over the network.
// Steps:
//  1. Get presentation context and transfer syntax from association
//  2. Encode the message into command and data datasets
//  3. Fragment the datasets into PDVs
//  4. Create P-DATA-TF PDU(s) from the PDVs
//  5. Write the PDU(s) to the network connection
func (s *Service) sendMessage(req *sendRequest) error {
	if req.ctx != nil {
		select {
		case <-req.ctx.Done():
			return req.ctx.Err()
		default:
		}
	}
	// Reject requests with missing required values before they reach the wire;
	// an empty SOP Class/Instance UID or AE title would otherwise be encoded as
	// a zero-length Type-1 element the peer has to reject.
	if request, ok := req.message.(dimse.Request); ok {
		if err := dimse.ValidateRequestForSend(request); err != nil {
			return err
		}
	}
	// Get association
	assoc := s.GetAssociation()
	if assoc == nil {
		return fmt.Errorf("no association available")
	}

	// Find presentation context.
	// A response must be sent on the presentation context of its request
	// (PS3.7 9.3), so an explicit ID that is not part of the association is an
	// error instead of an invitation to switch contexts. Requests may carry a
	// stale ID from a previous association, so they may still fall back to the
	// abstract-syntax lookup; the resolved ID is written before sending.
	message := req.message
	var pc *association.PresentationContext
	if req.message.CommandField() == uint16(dimse.CommandCStoreRQ) {
		ctx := req.ctx
		if ctx == nil {
			ctx = context.Background()
		}
		var err error
		pc, message, err = s.prepareCStoreMessage(ctx, assoc, req.message)
		if err != nil {
			return err
		}
	}
	if pc == nil {
		if contextID := message.PresentationContextID(); contextID != 0 {
			pc = assoc.FindPresentationContextByID(contextID)
			if pc == nil && dimse.CommandField(message.CommandField()).IsResponse() {
				return fmt.Errorf("presentation context ID %d is not part of the association", contextID)
			}
		}
		if pc == nil {
			// Fall back: look up by SOP Class UID (abstract syntax).
			sopClassUID, ok := message.CommandDataset().GetString(tag.AffectedSOPClassUID)
			if !ok || sopClassUID == "" {
				// Some messages use RequestedSOPClassUID instead (N-GET, N-SET, N-DELETE, N-ACTION)
				sopClassUID, ok = message.CommandDataset().GetString(tag.RequestedSOPClassUID)
				if !ok || sopClassUID == "" {
					return fmt.Errorf("message has no SOP Class UID (neither AffectedSOPClassUID nor RequestedSOPClassUID)")
				}
			}
			pc = assoc.FindPresentationContextByAbstractSyntax(sopClassUID)
			if pc == nil {
				return fmt.Errorf("no accepted presentation context found for SOP Class UID: %s", sopClassUID)
			}
		}
	}
	if !pc.IsAccepted() {
		return fmt.Errorf("presentation context ID %d is not accepted", pc.ID)
	}
	if dimse.CommandField(message.CommandField()).IsRequest() {
		if err := s.requireLocalRole(pc, true); err != nil {
			return err
		}
	}

	// Get transfer syntax from presentation context
	transferSyntax := pc.AcceptedTransferSyntax
	if transferSyntax == nil {
		return fmt.Errorf("presentation context ID %d has no accepted transfer syntax", pc.ID)
	}
	message.SetPresentationContextID(pc.ID)

	// Encode the message. A C-STORE data set is streamed into PDUs while it is
	// serialized, so a large data set is never held as a second, full-size copy
	// before being fragmented; other messages are encoded and fragmented whole.
	dataDS := message.DataDataset()
	streamCStore := dimse.CommandField(message.CommandField()) == dimse.CommandCStoreRQ &&
		dataDS != nil && len(dataDS.Elements()) > 0
	var commandData, datasetData []byte
	var err error
	if streamCStore {
		commandData, err = encodeCommandSetForStreaming(message)
	} else {
		commandData, datasetData, err = EncodeDIMSEMessage(message, transferSyntax)
	}
	if err != nil {
		return fmt.Errorf("failed to encode message: %w", err)
	}

	// Use presentation context ID from association negotiation
	contextID := pc.ID
	maxPDULength := s.outgoingMaxPDULength()

	return s.writeCommandAndDataSet(req.ctx, message, transferSyntax, contextID,
		maxPDULength, commandData, datasetData, streamCStore)
}

// writeCommandAndDataSet writes the command set and then the data set, either
// streamed from the data-set encoder (C-STORE) or fragmented from its buffered
// encoding.
func (s *Service) writeCommandAndDataSet(
	ctx context.Context,
	message dimse.Message,
	transferSyntax *transfer.Syntax,
	contextID byte,
	maxPDULength uint32,
	commandData, datasetData []byte,
	streamCStore bool,
) error {
	if err := s.writePDVFragment(ctx, commandData, contextID, true, maxPDULength); err != nil {
		return err
	}
	if !streamCStore {
		return s.writePDVFragment(ctx, datasetData, contextID, false, maxPDULength)
	}
	if err := s.streamDataSet(ctx, message.DataDataset(), transferSyntax, contextID, maxPDULength); err != nil {
		// The command set is already on the wire, so the peer cannot reassemble
		// this message: terminate the association instead of leaving it waiting
		// for data fragments that never arrive.
		_ = s.abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderNotSpecified, err, false)
		return fmt.Errorf("failed to send data set: %w", err)
	}
	return nil
}

// writePDVFragment fragments one encoded DIMSE message part into PDVs and
// writes them. Empty parts are not sent at all (PS3.8 E.1).
func (s *Service) writePDVFragment(ctx context.Context, data []byte, contextID byte, isCommand bool, maxPDULength uint32) error {
	if len(data) == 0 {
		return nil
	}
	pdvs := FragmentData(data, contextID, isCommand, maxPDULength)
	if len(pdvs) == 0 {
		part := "dataset"
		if isCommand {
			part = "command"
		}
		return fmt.Errorf("maximum PDU length %d is too small to hold a %s PDV", maxPDULength, part)
	}
	return s.writePDVGroups(ctx, pdvs)
}

// writePDVGroups packs PDVs into P-DATA-TF PDUs and writes them in order.
func (s *Service) writePDVGroups(ctx context.Context, pdvs []*PDV) error {
	for _, pdvGroup := range s.groupPDVsIntoPDUs(pdvs) {
		pduData, err := CreatePDataTFPDU(pdvGroup)
		if err != nil {
			return fmt.Errorf("failed to create P-DATA-TF PDU: %w", err)
		}

		// Write PDU to connection (serialized against release/abort writers)
		if err := s.writePDU(ctx, pduData); err != nil {
			return fmt.Errorf("failed to write PDU: %w", err)
		}
		s.emitPDUBytes(ctx, observability.DirectionOutbound, pduData)
	}
	return nil
}

// streamDataSet serializes a data set directly into one-PDV P-DATA-TF PDUs.
func (s *Service) streamDataSet(
	ctx context.Context,
	dataDS *dataset.Dataset,
	transferSyntax *transfer.Syntax,
	contextID byte,
	maxPDULength uint32,
) error {
	if transferSyntax == nil {
		transferSyntax = transfer.ExplicitVRLittleEndian
	}
	stream, err := newDatasetPDUWriter(ctx, s, contextID, maxPDULength)
	if err != nil {
		return err
	}
	if err := writer.Write(stream, dataDS,
		writer.WithTransferSyntax(transferSyntax),
		writer.WithoutPreamble()); err != nil {
		return fmt.Errorf("failed to encode data dataset: %w", err)
	}
	return stream.Finish()
}

// datasetPDUWriter streams an encoded data set into P-DATA-TF PDUs. Each PDU
// carries exactly one data PDV whose payload is written straight into the PDU
// buffer that goes to the connection, so the encoded bytes are copied once
// instead of once into a whole-data-set buffer and again into a PDU buffer
// (PS3.8 E.2). A full fragment is flushed by the next Write or by Finish so the
// final fragment can be marked as the last one.
type datasetPDUWriter struct {
	service   *Service
	ctx       context.Context
	contextID byte
	maxData   int // payload bytes per fragment; 0 means unbounded
	buf       []byte
	filled    int
}

const (
	// pdvHeaderSize is the presentation context ID plus the message control
	// header that precede every PDV payload (PS3.8 E.2).
	pdvHeaderSize = 6
	// pdvUnboundedInitialCapacity sizes the single-PDV buffer used when no
	// maximum PDU length was negotiated.
	pdvUnboundedInitialCapacity = 64 * 1024
)

// newDatasetPDUWriter returns a writer for maxPDULength, or an error when that
// length cannot hold a single data PDV (PS3.8 E.1).
func newDatasetPDUWriter(ctx context.Context, service *Service, contextID byte, maxPDULength uint32) (*datasetPDUWriter, error) {
	if maxPDULength == 0 {
		// No negotiated limit: buffer the whole data set into one PDV.
		return &datasetPDUWriter{
			service:   service,
			ctx:       ctx,
			contextID: contextID,
			buf:       make([]byte, pdvHeaderSize, pdvHeaderSize+pdvUnboundedInitialCapacity),
		}, nil
	}
	// The PDV header is part of the negotiated PDU length, and every fragment
	// holds an even number of bytes (PS3.8 E.1).
	if maxPDULength < pdvHeaderSize+2 {
		return nil, fmt.Errorf("maximum PDU length %d is too small to hold a dataset PDV", maxPDULength)
	}
	maxData := (maxPDULength - pdvHeaderSize) &^ 1
	return &datasetPDUWriter{
		service:   service,
		ctx:       ctx,
		contextID: contextID,
		maxData:   int(maxData),
		buf:       make([]byte, int(uint32(pdvHeaderSize)+maxData)),
	}, nil
}

// Write appends p to the current fragment, flushing it when it is full.
func (w *datasetPDUWriter) Write(p []byte) (int, error) {
	written := 0
	for len(p) > 0 {
		if w.maxData > 0 && w.filled == w.maxData {
			if err := w.flush(false); err != nil {
				return written, err
			}
		}
		take := len(p)
		if w.maxData > 0 {
			if space := w.maxData - w.filled; take > space {
				take = space
			}
			copy(w.buf[pdvHeaderSize+w.filled:], p[:take])
		} else {
			w.buf = append(w.buf, p...)
		}
		w.filled += take
		written += take
		p = p[take:]
	}
	return written, nil
}

// Finish writes the last fragment, padded to an even length per PS3.8 E.1.
func (w *datasetPDUWriter) Finish() error {
	if w.filled == 0 {
		return nil
	}
	if w.filled%2 != 0 {
		if len(w.buf) > pdvHeaderSize+w.filled {
			w.buf[pdvHeaderSize+w.filled] = 0
		} else {
			w.buf = append(w.buf, 0)
		}
		w.filled++
	}
	return w.flush(true)
}

func (w *datasetPDUWriter) flush(last bool) error {
	payloadLength := w.filled
	// PDV header: length (context ID + control header + payload), context ID,
	// control header with the last-fragment bit (PS3.8 E.2).
	binary.BigEndian.PutUint32(w.buf[0:4], uint32(1+1+payloadLength)) // #nosec G115 -- bounded by the negotiated PDU length
	w.buf[4] = w.contextID
	if last {
		w.buf[5] = 0x02
	} else {
		w.buf[5] = 0
	}

	// The PDU header is added by pdu.RawPDU.Write, so Data holds the PDV header
	// followed by the payload that was written straight into this buffer.
	raw := &pdu.RawPDU{Type: pdu.TypePDataTF, Data: w.buf[:pdvHeaderSize+payloadLength]}
	if err := w.service.writePDU(w.ctx, raw); err != nil {
		return fmt.Errorf("failed to write PDU: %w", err)
	}
	w.service.emitPDUBytes(w.ctx, observability.DirectionOutbound, raw)
	w.filled = 0
	return nil
}

func (s *Service) prepareCStoreMessage(ctx context.Context, assoc *association.Association, message dimse.Message) (*association.PresentationContext, dimse.Message, error) {
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

	manager := s.config.transcodeManager
	if manager == nil {
		return nil, nil, fmt.Errorf("%w: cannot convert C-STORE SOP Class %s from %s",
			ErrTranscodeManagerUnavailable, sopClassUID, sourceSyntax.UID().UID())
	}
	for _, candidate := range candidates {
		if !manager.CanTranscode(sourceSyntax, candidate.AcceptedTransferSyntax) {
			continue
		}
		transcoder, err := manager.NewTranscoder(sourceSyntax, candidate.AcceptedTransferSyntax)
		if err != nil {
			return nil, nil, fmt.Errorf("create C-STORE transcoder from %s to %s: %w",
				sourceSyntax.UID().UID(), candidate.AcceptedTransferSyntax.UID().UID(), err)
		}
		transcoded, err := transcoder.Transcode(ctx, dataDS)
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
		"%w: no accepted transfer syntax for C-STORE SOP Class %s is directly usable or transcodable from %s; accepted transfer syntaxes: %s",
		transcode.ErrCodecUnavailable,
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

	var groups [][]*PDV
	var currentGroup []*PDV
	// MaximumLengthReceived applies to the PDU variable field and excludes the
	// six-byte PDU header (PS3.8).
	currentSize := 0
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
			currentSize = pdvSize
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
