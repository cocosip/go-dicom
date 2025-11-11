// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"
	"fmt"
	"io"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

// serviceResponder implements AssociationResponder and provides methods
// for the negotiator to send association responses.
type serviceResponder struct {
	service *Service
}

// SendAccept sends an A-ASSOCIATE-AC (Association Accept) PDU.
func (r *serviceResponder) SendAccept(ctx context.Context, assoc *association.Association) error {
	// Convert Association to A-ASSOCIATE-AC PDU
	ac := association.ToAAssociateAC(assoc)

	// Send the PDU
	if err := r.service.SendAssociationAccept(ctx, ac); err != nil {
		return fmt.Errorf("failed to send A-ASSOCIATE-AC: %w", err)
	}

	// Store the association
	r.service.SetAssociation(assoc)

	return nil
}

// SendReject sends an A-ASSOCIATE-RJ (Association Reject) PDU.
func (r *serviceResponder) SendReject(ctx context.Context, result, source, reason byte) error {
	return r.service.SendAssociationReject(ctx, result, source, reason)
}

// serializeRawPDU converts a RawPDU to complete PDU bytes (header + data).
func serializeRawPDU(rawPDU *pdu.RawPDU) []byte {
	pduBytes := make([]byte, 6+len(rawPDU.Data))
	pduBytes[0] = rawPDU.Type
	pduBytes[1] = rawPDU.Reserved
	pduBytes[2] = byte(rawPDU.Length >> 24)
	pduBytes[3] = byte(rawPDU.Length >> 16)
	pduBytes[4] = byte(rawPDU.Length >> 8)
	pduBytes[5] = byte(rawPDU.Length)
	copy(pduBytes[6:], rawPDU.Data)
	return pduBytes
}

// pduEncoder is an interface for PDUs that can be encoded.
type pduEncoder interface {
	Encode() (*pdu.RawPDU, error)
}

// sendAssociationPDU is a helper that sends an association-related PDU and transitions state.
func (s *Service) sendAssociationPDU(ctx context.Context, pduData pduEncoder, pduName string, newState State) error {
	if s.IsClosed() {
		return ErrServiceClosed
	}
	currentState := s.GetState()
	if currentState != StateIdle {
		return fmt.Errorf("cannot send %s in state %s", pduName, currentState)
	}
	rawPDU, err := pduData.Encode()
	if err != nil {
		return fmt.Errorf("failed to encode %s: %w", pduName, err)
	}
	if s.config.writeTimeout > 0 {
		if err := s.conn.SetWriteDeadline(deadlineFromContext(ctx, s.config.writeTimeout)); err != nil {
			return fmt.Errorf("failed to set write deadline: %w", err)
		}
	}
	if _, err := s.conn.Write(serializeRawPDU(rawPDU)); err != nil {
		return fmt.Errorf("failed to write %s: %w", pduName, err)
	}
	if err := s.setState(newState); err != nil {
		return fmt.Errorf("failed to transition state: %w", err)
	}
	return nil
}

// SendAssociationRequest sends an A-ASSOCIATE-RQ PDU to request an association.
// This is typically called by a client (SCU) to initiate a DICOM association.
func (s *Service) SendAssociationRequest(ctx context.Context, req *pdu.AAssociateRQ) error {
	return s.sendAssociationPDU(ctx, req, "A-ASSOCIATE-RQ", StateAssociationRequested)
}

// SendAssociationAccept sends an A-ASSOCIATE-AC PDU to accept an association.
// This is typically called by a server (SCP) in response to an A-ASSOCIATE-RQ.
func (s *Service) SendAssociationAccept(ctx context.Context, ac *pdu.AAssociateAC) error {
	return s.sendAssociationPDU(ctx, ac, "A-ASSOCIATE-AC", StateAssociationAccepted)
}

// SendAssociationReject sends an A-ASSOCIATE-RJ PDU to reject an association.
// This is typically called by a server (SCP) when it cannot accept the association.
//
// Parameters:
//   - result: The result code (1=rejected-permanent, 2=rejected-transient)
//   - source: The source of rejection (1=service-user, 2=service-provider)
//   - reason: The reason for rejection (depends on source)
func (s *Service) SendAssociationReject(ctx context.Context, result, source, reason byte) error {
	// Check if service is closed
	if s.IsClosed() {
		return ErrServiceClosed
	}

	// Create A-ASSOCIATE-RJ PDU
	rj := &pdu.AAssociateRJ{
		Result: result,
		Source: source,
		Reason: reason,
	}

	// Encode PDU
	rawPDU, err := rj.Encode()
	if err != nil {
		return fmt.Errorf("failed to encode A-ASSOCIATE-RJ: %w", err)
	}

	// Write to connection with timeout
	if s.config.writeTimeout > 0 {
		if err := s.conn.SetWriteDeadline(deadlineFromContext(ctx, s.config.writeTimeout)); err != nil {
			return fmt.Errorf("failed to set write deadline: %w", err)
		}
	}

	if _, err := s.conn.Write(serializeRawPDU(rawPDU)); err != nil {
		return fmt.Errorf("failed to write A-ASSOCIATE-RJ: %w", err)
	}

	// Transition to Closed state after rejection
	if err := s.setState(StateClosed); err != nil {
		return fmt.Errorf("failed to transition state: %w", err)
	}

	return nil
}

// SendReleaseRequest sends an A-RELEASE-RQ PDU to request graceful release of the association.
// This can be called by either SCU or SCP when they want to close the association normally.
//
// The method:
// 1. Validates the current state (must be AssociationAccepted or Transferring)
// 2. Encodes the release request PDU
// 3. Writes it to the connection
// 4. Transitions to ReleaseRequested state
func (s *Service) SendReleaseRequest(ctx context.Context) error {
	// Check if service is closed
	if s.IsClosed() {
		return ErrServiceClosed
	}

	// Validate state - must be in active association
	currentState := s.GetState()
	if currentState != StateAssociationAccepted && currentState != StateTransferring {
		return fmt.Errorf("cannot send A-RELEASE-RQ in state %s", currentState)
	}

	// Create A-RELEASE-RQ PDU
	releaseRQ := &pdu.AReleaseRQ{}

	// Encode PDU
	rawPDU, err := releaseRQ.Encode()
	if err != nil {
		return fmt.Errorf("failed to encode A-RELEASE-RQ: %w", err)
	}

	// Write to connection with timeout
	if s.config.writeTimeout > 0 {
		if err := s.conn.SetWriteDeadline(deadlineFromContext(ctx, s.config.writeTimeout)); err != nil {
			return fmt.Errorf("failed to set write deadline: %w", err)
		}
	}

	if _, err := s.conn.Write(serializeRawPDU(rawPDU)); err != nil {
		return fmt.Errorf("failed to write A-RELEASE-RQ: %w", err)
	}

	// Transition to ReleaseRequested state
	if err := s.setState(StateReleaseRequested); err != nil {
		return fmt.Errorf("failed to transition state: %w", err)
	}

	return nil
}

// SendReleaseResponse sends an A-RELEASE-RP PDU to confirm release of the association.
// This is sent in response to an A-RELEASE-RQ from the peer.
//
// The method:
// 1. Validates the current state
// 2. Encodes the release response PDU
// 3. Writes it to the connection
// 4. Transitions to Closed state
func (s *Service) SendReleaseResponse(ctx context.Context) error {
	// Check if service is closed
	if s.IsClosed() {
		return ErrServiceClosed
	}

	// Create A-RELEASE-RP PDU
	releaseRP := &pdu.AReleaseRP{}

	// Encode PDU
	rawPDU, err := releaseRP.Encode()
	if err != nil {
		return fmt.Errorf("failed to encode A-RELEASE-RP: %w", err)
	}

	// Write to connection with timeout
	if s.config.writeTimeout > 0 {
		if err := s.conn.SetWriteDeadline(deadlineFromContext(ctx, s.config.writeTimeout)); err != nil {
			return fmt.Errorf("failed to set write deadline: %w", err)
		}
	}

	if _, err := s.conn.Write(serializeRawPDU(rawPDU)); err != nil {
		return fmt.Errorf("failed to write A-RELEASE-RP: %w", err)
	}

	// Transition to Closed state
	if err := s.setState(StateClosed); err != nil {
		return fmt.Errorf("failed to transition state: %w", err)
	}

	return nil
}

// SendAbort sends an A-ABORT PDU to immediately terminate the association.
// This is used for abnormal termination (errors, protocol violations, etc.).
//
// Parameters:
//   - source: The source of abort (0=service-user, 2=service-provider)
//   - reason: The reason for abort (0=not-specified, 1=unrecognized-pdu, etc.)
func (s *Service) SendAbort(ctx context.Context, source, reason byte) error {
	// Check if service is closed (but allow abort even if closed)
	// This is because abort is used for error handling

	// Create A-ABORT PDU
	abort := &pdu.AAbort{
		Source: source,
		Reason: reason,
	}

	// Encode PDU
	rawPDU, err := abort.Encode()
	if err != nil {
		return fmt.Errorf("failed to encode A-ABORT: %w", err)
	}

	// Write to connection with timeout (ignore errors if connection is broken)
	if s.conn != nil {
		if s.config.writeTimeout > 0 {
			_ = s.conn.SetWriteDeadline(deadlineFromContext(ctx, s.config.writeTimeout))
		}
		_, _ = s.conn.Write(serializeRawPDU(rawPDU))
	}

	// Transition to Aborted state (ignore errors)
	_ = s.setState(StateAborted)

	return nil
}

// ReceiveAssociationResponse waits for and processes an A-ASSOCIATE-AC or A-ASSOCIATE-RJ response.
// This is typically called by a client after sending an A-ASSOCIATE-RQ.
//
// Returns:
//   - *pdu.AAssociateAC if association was accepted
//   - error if association was rejected or other error occurred
func (s *Service) ReceiveAssociationResponse(ctx context.Context) (*pdu.AAssociateAC, error) {
	// Check if service is closed
	if s.IsClosed() {
		return nil, ErrServiceClosed
	}

	// Validate state - should be AssociationRequested
	currentState := s.GetState()
	if currentState != StateAssociationRequested {
		return nil, fmt.Errorf("cannot receive association response in state %s", currentState)
	}

	// Read PDU header
	header := make([]byte, 6)
	if s.config.readTimeout > 0 {
		if err := s.conn.SetReadDeadline(deadlineFromContext(ctx, s.config.readTimeout)); err != nil {
			return nil, fmt.Errorf("failed to set read deadline: %w", err)
		}
	}

	if _, err := io.ReadFull(s.conn, header); err != nil {
		return nil, fmt.Errorf("failed to read PDU header: %w", err)
	}

	pduType := header[0]
	pduLength := uint32(header[2])<<24 | uint32(header[3])<<16 | uint32(header[4])<<8 | uint32(header[5])

	// Read PDU body
	body := make([]byte, pduLength)
	if _, err := io.ReadFull(s.conn, body); err != nil {
		return nil, fmt.Errorf("failed to read PDU body: %w", err)
	}

	// Create RawPDU (Data should contain only the body, not header)
	rawPDU := &pdu.RawPDU{
		Type:   pduType,
		Length: pduLength,
		Data:   body,
	}

	// Check PDU type
	switch pduType {
	case pdu.TypeAAssociateAC:
		// Decode A-ASSOCIATE-AC
		ac := &pdu.AAssociateAC{}
		if err := ac.Decode(rawPDU); err != nil {
			return nil, fmt.Errorf("failed to decode A-ASSOCIATE-AC: %w", err)
		}

		// Transition to AssociationAccepted state
		if err := s.setState(StateAssociationAccepted); err != nil {
			return nil, fmt.Errorf("failed to transition state: %w", err)
		}

		return ac, nil

	case pdu.TypeAAssociateRJ:
		// Decode A-ASSOCIATE-RJ
		rj := &pdu.AAssociateRJ{}
		if err := rj.Decode(rawPDU); err != nil {
			return nil, fmt.Errorf("failed to decode A-ASSOCIATE-RJ: %w", err)
		}

		// Transition to Closed state
		_ = s.setState(StateClosed)

		return nil, fmt.Errorf("association rejected: result=%d, source=%d, reason=%d",
			rj.Result, rj.Source, rj.Reason)

	default:
		return nil, fmt.Errorf("unexpected PDU type: expected A-ASSOCIATE-AC or RJ, got 0x%02X", pduType)
	}
}

// ReceiveAssociationRequest waits for and processes an A-ASSOCIATE-RQ.
// This is typically called by a server after accepting a TCP connection.
//
// Returns:
//   - *pdu.AAssociateRQ if successful
//   - error if failed or wrong PDU type received
func (s *Service) ReceiveAssociationRequest(ctx context.Context) (*pdu.AAssociateRQ, error) {
	// Check if service is closed
	if s.IsClosed() {
		return nil, ErrServiceClosed
	}

	// Validate state - should be Idle
	currentState := s.GetState()
	if currentState != StateIdle {
		return nil, fmt.Errorf("cannot receive association request in state %s", currentState)
	}

	// Read PDU header
	header := make([]byte, 6)
	if s.config.readTimeout > 0 {
		if err := s.conn.SetReadDeadline(deadlineFromContext(ctx, s.config.readTimeout)); err != nil {
			return nil, fmt.Errorf("failed to set read deadline: %w", err)
		}
	}

	if _, err := io.ReadFull(s.conn, header); err != nil {
		return nil, fmt.Errorf("failed to read PDU header: %w", err)
	}

	pduType := header[0]
	pduLength := uint32(header[2])<<24 | uint32(header[3])<<16 | uint32(header[4])<<8 | uint32(header[5])

	// Validate PDU type
	if pduType != pdu.TypeAAssociateRQ {
		return nil, fmt.Errorf("unexpected PDU type: expected A-ASSOCIATE-RQ (0x01), got 0x%02X", pduType)
	}

	// Read PDU body
	body := make([]byte, pduLength)
	if _, err := io.ReadFull(s.conn, body); err != nil {
		return nil, fmt.Errorf("failed to read PDU body: %w", err)
	}

	// Create RawPDU (Data should contain only the body, not header)
	rawPDU := &pdu.RawPDU{
		Type:   pduType,
		Length: pduLength,
		Data:   body,
	}

	// Decode A-ASSOCIATE-RQ
	rq := &pdu.AAssociateRQ{}
	if err := rq.Decode(rawPDU); err != nil {
		return nil, fmt.Errorf("failed to decode A-ASSOCIATE-RQ: %w", err)
	}

	// Build Association from A-ASSOCIATE-RQ
	assoc := association.FromAAssociateRQ(rq)

	// Create responder for the negotiator
	responder := &serviceResponder{
		service: s,
	}

	// Call OnAssociationRequest callback if negotiator is set (server mode)
	// The negotiator MUST call responder.SendAccept() or responder.SendReject()
	s.callbacksMu.RLock()
	negotiator := s.associationNegotiator
	s.callbacksMu.RUnlock()

	if negotiator != nil {
		// Let the negotiator handle the request and send the response
		if err := negotiator.OnAssociationRequest(ctx, assoc, responder); err != nil {
			return nil, fmt.Errorf("association negotiation failed: %w", err)
		}
	} else {
		// No negotiator - use default behavior: accept all contexts with first transfer syntax
		defaultNegotiator := &DefaultAssociationNegotiator{}
		if err := defaultNegotiator.OnAssociationRequest(ctx, assoc, responder); err != nil {
			return nil, fmt.Errorf("default association negotiation failed: %w", err)
		}
	}

	return rq, nil
}
