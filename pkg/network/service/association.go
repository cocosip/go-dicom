// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/observability"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
	"github.com/cocosip/go-dicom/pkg/network/transport"
)

// serviceResponder implements AssociationResponder and provides methods
// for the negotiator to send association responses.
type serviceResponder struct {
	service *Service
}

// maxAssociationPDULength bounds an A-ASSOCIATE-RQ/AC received before any
// Maximum Length has been negotiated. It prevents a peer from making the
// process allocate a huge buffer from a 6-byte header (PS3.8 9.3.2 limits the
// User Information item to 65533 bytes, so 4 MiB is already generous).
const maxAssociationPDULength = 4 << 20

// SendAccept sends an A-ASSOCIATE-AC (Association Accept) PDU.
func (r *serviceResponder) SendAccept(ctx context.Context, assoc *association.Association) error {
	// The AC identifies the acceptor, never the requestor (PS3.7 D.3.3.2.2).
	assoc.LocalImplementationClassUID = r.service.config.implementationClassUID
	assoc.LocalImplementationVersionName = r.service.config.implementationVersionName

	// Convert Association to A-ASSOCIATE-AC PDU. The AC advertises the
	// acceptor's own receive limit; assoc.MaxPDULength is the requestor's limit
	// and is retained for outgoing PDU fragmentation.
	ac := association.ToAAssociateACWithMaxPDULength(assoc, r.service.config.maxPDULength)
	// Keep the negotiated request-side contexts available while the AC log is emitted.
	r.service.SetAssociation(assoc)

	// Send the PDU
	if err := r.service.SendAssociationAccept(ctx, ac); err != nil {
		return fmt.Errorf("failed to send A-ASSOCIATE-AC: %w", err)
	}

	return nil
}

// SendReject sends an A-ASSOCIATE-RJ (Association Reject) PDU.
func (r *serviceResponder) SendReject(ctx context.Context, result, source, reason byte) error {
	return r.service.SendAssociationReject(ctx, result, source, reason)
}

// pduEncoder is an interface for PDUs that can be encoded.
type pduEncoder interface {
	Encode() (*pdu.RawPDU, error)
}

// writePDU serializes all writes to the connection and applies the configured
// write timeout. Serialization is required because a DICOM PDU is emitted with
// several Write calls (header then body); concurrent writers would interleave.
func (s *Service) writePDU(ctx context.Context, raw *pdu.RawPDU) error {
	if s.conn == nil {
		return fmt.Errorf("connection is nil")
	}
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	// P-DATA-TF is only valid in the established state (PS3.8 Sta6). Checking
	// under writeMu guarantees that an A-RELEASE-RQ cannot slip in front of a
	// DIMSE message that was already being encoded.
	if raw.Type == pdu.TypePDataTF && s.GetState() != StateAssociationAccepted {
		return fmt.Errorf("cannot send P-DATA-TF in state %s", s.GetState())
	}
	if deadline := deadlineFromContext(ctx, s.config.writeTimeout); !deadline.IsZero() {
		if err := s.conn.SetWriteDeadline(deadline); err != nil {
			return fmt.Errorf("failed to set write deadline: %w", err)
		}
		defer func() { _ = s.conn.SetWriteDeadline(time.Time{}) }()
	}
	return raw.Write(s.conn)
}

// sendAssociationPDU is a helper that sends an association-related PDU and transitions state.
func (s *Service) sendAssociationPDU(ctx context.Context, pduData pduEncoder, pduName string, newState State) error {
	if s.conn == nil {
		return fmt.Errorf("cannot send %s: connection is nil", pduName)
	}
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
	if err := s.writePDU(ctx, rawPDU); err != nil {
		return fmt.Errorf("failed to write %s: %w", pduName, err)
	}
	s.emitPDUBytes(ctx, observability.DirectionOutbound, rawPDU)
	if err := s.setState(newState); err != nil {
		return fmt.Errorf("failed to transition state: %w", err)
	}
	return nil
}

// SendAssociationRequest sends an A-ASSOCIATE-RQ PDU to request an association.
// This is typically called by a client (SCU) to initiate a DICOM association.
func (s *Service) SendAssociationRequest(ctx context.Context, req *pdu.AAssociateRQ) error {
	s.ensureAssociationIdentity(req.CallingAETitle, req.CalledAETitle)
	associationInfo := s.associationObservationForRQ(req)
	if err := s.sendAssociationPDU(ctx, req, "A-ASSOCIATE-RQ", StateAssociationRequested); err != nil {
		return err
	}
	s.emitEvent(ctx, observability.Event{
		Kind:        observability.EventAssociationRequested,
		Association: associationInfo,
		Direction:   observability.DirectionOutbound,
		Outcome:     observability.OutcomeSuccess,
	})
	s.emitMetric(ctx, observability.Metric{
		Kind:        observability.MetricAssociation,
		Association: associationInfo,
		Direction:   observability.DirectionOutbound,
		Outcome:     observability.OutcomeSuccess,
		Value:       1,
	})
	return nil
}

// SendAssociationAccept sends an A-ASSOCIATE-AC PDU to accept an association.
// This is typically called by a server (SCP) in response to an A-ASSOCIATE-RQ.
func (s *Service) SendAssociationAccept(ctx context.Context, ac *pdu.AAssociateAC) error {
	if err := s.sendAssociationPDU(ctx, ac, "A-ASSOCIATE-AC", StateAssociationAccepted); err != nil {
		return err
	}
	s.emitAssociationObservationWithInfo(ctx, observability.EventAssociationAccepted, observability.DirectionOutbound, observability.OutcomeSuccess, nil, s.associationObservationForAC(ac))
	return nil
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
	if s.conn == nil {
		return fmt.Errorf("cannot send A-ASSOCIATE-RJ: connection is nil")
	}
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
	if err := s.writePDU(ctx, rawPDU); err != nil {
		return fmt.Errorf("failed to write A-ASSOCIATE-RJ: %w", err)
	}
	s.emitPDUBytes(ctx, observability.DirectionOutbound, rawPDU)

	// Transition to Closed state after rejection
	if err := s.setState(StateClosed); err != nil {
		return fmt.Errorf("failed to transition state: %w", err)
	}
	s.emitAssociationObservation(ctx, observability.EventAssociationRejected, observability.DirectionOutbound, observability.OutcomeRejected, nil)

	return nil
}

// SendReleaseRequest sends an A-RELEASE-RQ PDU to request graceful release of the association.
// This can be called by either SCU or SCP when they want to close the association normally.
//
// The method:
// 1. Validates the current state (must be AssociationAccepted or Transferring)
// 2. Encodes the release request PDU
// 3. Transitions to ReleaseRequested state
// 4. Writes it to the connection
func (s *Service) SendReleaseRequest(ctx context.Context) error {
	// Check if service is closed
	if s.conn == nil {
		return fmt.Errorf("cannot send A-RELEASE-RQ: connection is nil")
	}
	if s.IsClosed() {
		return ErrServiceClosed
	}

	// Validate state - must be in an active association
	currentState := s.GetState()
	if currentState != StateAssociationAccepted {
		return fmt.Errorf("cannot send A-RELEASE-RQ in state %s", currentState)
	}

	// Create A-RELEASE-RQ PDU
	releaseRQ := &pdu.AReleaseRQ{}

	// Encode PDU
	rawPDU, err := releaseRQ.Encode()
	if err != nil {
		return fmt.Errorf("failed to encode A-RELEASE-RQ: %w", err)
	}

	// Publish the release state before writing. A fast peer can send its
	// A-RELEASE-RP as soon as it receives this request, while recvLoop runs
	// concurrently with this method.
	if err := s.setState(StateReleaseRequested); err != nil {
		return fmt.Errorf("failed to transition state: %w", err)
	}

	// Write to connection with timeout
	if err := s.writePDU(ctx, rawPDU); err != nil {
		return fmt.Errorf("failed to write A-RELEASE-RQ: %w", err)
	}
	s.emitPDUBytes(ctx, observability.DirectionOutbound, rawPDU)

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
	if s.conn == nil {
		return fmt.Errorf("cannot send A-RELEASE-RP: connection is nil")
	}
	if s.IsClosed() {
		return ErrServiceClosed
	}

	// A-RELEASE-RP is only valid while an association is established or while
	// the peer's release request is being answered (PS3.8 Sta7/Sta8).
	switch currentState := s.GetState(); currentState {
	case StateAssociationAccepted, StateReleaseRequested:
	default:
		return fmt.Errorf("cannot send A-RELEASE-RP in state %s", currentState)
	}

	if err := s.writeReleaseResponse(ctx); err != nil {
		return err
	}

	// Transition to Closed state
	if err := s.setState(StateClosed); err != nil {
		return fmt.Errorf("failed to transition state: %w", err)
	}
	s.emitAssociationObservation(ctx, observability.EventAssociationReleased, observability.DirectionOutbound, observability.OutcomeSuccess, nil)

	return nil
}

// writeReleaseResponse encodes and writes an A-RELEASE-RP without changing the
// service state. It is used by SendReleaseResponse and by the compatibility
// path that keeps the connection open after a peer release.
func (s *Service) writeReleaseResponse(ctx context.Context) error {
	// Create A-RELEASE-RP PDU
	releaseRP := &pdu.AReleaseRP{}

	// Encode PDU
	rawPDU, err := releaseRP.Encode()
	if err != nil {
		return fmt.Errorf("failed to encode A-RELEASE-RP: %w", err)
	}

	// Write to connection with timeout
	if err := s.writePDU(ctx, rawPDU); err != nil {
		return fmt.Errorf("failed to write A-RELEASE-RP: %w", err)
	}
	s.emitPDUBytes(ctx, observability.DirectionOutbound, rawPDU)
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
	var writeErr error
	if s.conn != nil {
		writeErr = s.writePDU(ctx, rawPDU)
		if writeErr == nil {
			s.emitPDUBytes(ctx, observability.DirectionOutbound, rawPDU)
		}
	}

	// Transition to Aborted state (ignore errors)
	_ = s.setState(StateAborted)
	s.emitAssociationObservation(ctx, observability.EventAssociationAborted, observability.DirectionOutbound, observability.OutcomeAborted, writeErr)

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
	if s.conn == nil {
		return nil, fmt.Errorf("cannot receive association response: connection is nil")
	}
	if s.IsClosed() {
		return nil, ErrServiceClosed
	}

	// Validate state - should be AssociationRequested
	currentState := s.GetState()
	if currentState != StateAssociationRequested {
		return nil, fmt.Errorf("cannot receive association response in state %s", currentState)
	}

	rawPDU, err := transport.ReadPDUMaxLength(s.conn, readTimeoutFromContext(ctx, s.config.readTimeout), maxAssociationPDULength)
	if err != nil {
		if errors.Is(err, transport.ErrUnrecognizedPDUType) {
			_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderUnrecognizedPDU)
		}
		return nil, fmt.Errorf("failed to read PDU: %w", err)
	}
	s.emitPDUBytes(ctx, observability.DirectionInbound, rawPDU)
	pduType := rawPDU.Type

	// Check PDU type
	switch pduType {
	case pdu.TypeAAssociateAC:
		// Decode A-ASSOCIATE-AC
		ac := &pdu.AAssociateAC{}
		if err := ac.Decode(rawPDU); err != nil {
			_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderInvalidPDUParam)
			return nil, fmt.Errorf("failed to decode A-ASSOCIATE-AC: %w", err)
		}
		if err := ac.ValidateStructure(); err != nil {
			_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderInvalidPDUParam)
			return nil, fmt.Errorf("invalid A-ASSOCIATE-AC structure: %w", err)
		}

		// Transition to AssociationAccepted state
		if err := s.setState(StateAssociationAccepted); err != nil {
			return nil, fmt.Errorf("failed to transition state: %w", err)
		}
		s.emitAssociationObservationWithInfo(ctx, observability.EventAssociationAccepted, observability.DirectionInbound, observability.OutcomeSuccess, nil, s.associationObservationForAC(ac))

		return ac, nil

	case pdu.TypeAAssociateRJ:
		// Decode A-ASSOCIATE-RJ
		rj := &pdu.AAssociateRJ{}
		if err := rj.Decode(rawPDU); err != nil {
			_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderInvalidPDUParam)
			return nil, fmt.Errorf("failed to decode A-ASSOCIATE-RJ: %w", err)
		}

		// Transition to Closed state
		_ = s.setState(StateClosed)
		s.emitAssociationObservation(ctx, observability.EventAssociationRejected, observability.DirectionInbound, observability.OutcomeRejected, nil)

		return nil, fmt.Errorf("%w: result=%d, source=%d, reason=%d",
			ErrAssociationRejected, rj.Result, rj.Source, rj.Reason)

	default:
		_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderUnexpectedPDU)
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
	if s.conn == nil {
		return nil, fmt.Errorf("cannot receive association request: connection is nil")
	}
	if s.IsClosed() {
		return nil, ErrServiceClosed
	}

	// Validate state - should be Idle
	currentState := s.GetState()
	if currentState != StateIdle {
		return nil, fmt.Errorf("cannot receive association request in state %s", currentState)
	}

	rawPDU, err := transport.ReadPDUMaxLength(s.conn, readTimeoutFromContext(ctx, s.config.readTimeout), maxAssociationPDULength)
	if err != nil {
		// An unrecognized PDU type must be reported with A-ABORT (PS3.8 9.3.8);
		// plain I/O errors just close the connection.
		if errors.Is(err, transport.ErrUnrecognizedPDUType) {
			_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderUnrecognizedPDU)
		}
		return nil, fmt.Errorf("failed to read PDU: %w", err)
	}
	s.emitPDUBytes(ctx, observability.DirectionInbound, rawPDU)
	pduType := rawPDU.Type

	// Validate PDU type
	if pduType != pdu.TypeAAssociateRQ {
		_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderUnexpectedPDU)
		return nil, fmt.Errorf("unexpected PDU type: expected A-ASSOCIATE-RQ (0x01), got 0x%02X", pduType)
	}

	// Decode A-ASSOCIATE-RQ
	rq := &pdu.AAssociateRQ{}
	var warnings []pdu.DecodeWarning
	if err := rq.DecodeWithWarnings(rawPDU, func(warning pdu.DecodeWarning) {
		warnings = append(warnings, warning)
	}); err != nil {
		_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderInvalidPDUParam)
		return nil, fmt.Errorf("failed to decode A-ASSOCIATE-RQ: %w", err)
	}
	if err := rq.ValidateStructure(); err != nil {
		_ = s.Abort(ctx, pdu.AbortSourceServiceProvider, pdu.AbortReasonServiceProviderInvalidPDUParam)
		return nil, fmt.Errorf("invalid A-ASSOCIATE-RQ structure: %w", err)
	}
	// Reject an invalid peer AE Title with A-ASSOCIATE-RJ (PS3.8 Table 9-11)
	// instead of failing later while encoding the A-ASSOCIATE-AC.
	if err := pdu.ValidateAETitle("called", rq.CalledAETitle); err != nil {
		_ = s.SendAssociationReject(ctx, pdu.ResultRejectedPermanent, pdu.SourceServiceUser, pdu.ReasonServiceUserCalledAETitleNotRecognized)
		return nil, fmt.Errorf("%w: %v", ErrAssociationRejected, err)
	}
	if err := pdu.ValidateAETitle("calling", rq.CallingAETitle); err != nil {
		_ = s.SendAssociationReject(ctx, pdu.ResultRejectedPermanent, pdu.SourceServiceUser, pdu.ReasonServiceUserCallingAETitleNotRecognized)
		return nil, fmt.Errorf("%w: %v", ErrAssociationRejected, err)
	}
	s.ensureAssociationIdentity(rq.CallingAETitle, rq.CalledAETitle)
	associationInfo := s.associationObservationForRQ(rq)
	for _, warning := range warnings {
		s.emitDecodeWarning(ctx, associationInfo, warning)
	}
	s.emitEvent(ctx, observability.Event{
		Kind:        observability.EventAssociationRequested,
		Association: associationInfo,
		Direction:   observability.DirectionInbound,
		Outcome:     observability.OutcomeSuccess,
	})
	s.emitMetric(ctx, observability.Metric{
		Kind:        observability.MetricAssociation,
		Association: associationInfo,
		Direction:   observability.DirectionInbound,
		Outcome:     observability.OutcomeSuccess,
		Value:       1,
	})
	if rq.ProtocolVersion&0x0001 == 0 {
		if err := s.SendAssociationReject(
			ctx,
			pdu.ResultRejectedPermanent,
			pdu.SourceServiceProviderACSE,
			pdu.ReasonServiceProviderACSEProtocolVersionNotSupported,
		); err != nil {
			return nil, fmt.Errorf("failed to reject unsupported protocol version 0x%04X: %w", rq.ProtocolVersion, err)
		}
		return nil, fmt.Errorf("%w: unsupported protocol version 0x%04X", ErrAssociationRejected, rq.ProtocolVersion)
	}

	// Build Association from A-ASSOCIATE-RQ
	assoc := association.FromAAssociateRQ(rq, s.config.transferSyntaxRegistry)

	// PS3.7 D.3.3.3: the response shall only contain the Asynchronous Operations
	// Window when the request contained it, and the accepted values shall be
	// equal to or less than the proposed values (0 means unlimited).
	if requested := assoc.RequestedAsynchronousOperations; requested != nil {
		var configuredInvoked, configuredPerformed uint16
		if configured := s.config.asynchronousOperations; configured != nil {
			configuredInvoked = configured.MaxInvokedOperations
			configuredPerformed = configured.MaxPerformedOperations
		}
		assoc.AsynchronousOperations = association.NewAsynchronousOperationsWindow(
			clampAsyncOperations(configuredInvoked, requested.MaxInvokedOperations),
			clampAsyncOperations(configuredPerformed, requested.MaxPerformedOperations),
		)
	} else {
		assoc.AsynchronousOperations = nil
	}

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
		// No negotiator - use default behavior: accept all contexts with the
		// first proposed transfer syntax that the registry knows about.
		defaultNegotiator := &DefaultAssociationNegotiator{Registry: s.config.transferSyntaxRegistry}
		if err := defaultNegotiator.OnAssociationRequest(ctx, assoc, responder); err != nil {
			return nil, fmt.Errorf("default association negotiation failed: %w", err)
		}
	}

	// The negotiator must have sent exactly one response. Returning nil from a
	// rejection path previously left callers unable to tell success from
	// failure, so report the outcome explicitly.
	switch s.GetState() {
	case StateAssociationAccepted:
		return rq, nil
	case StateClosed, StateAborted:
		return nil, ErrAssociationRejected
	default:
		return nil, fmt.Errorf("association negotiator returned without sending A-ASSOCIATE-AC or RJ")
	}
}

// clampAsyncOperations returns the accepted value for one Asynchronous
// Operations Window field. Zero means "unlimited" in PS3.7 D.3.3.3; the
// response value must be equal to or less than the offered value.
func clampAsyncOperations(configured, offered uint16) uint16 {
	if offered == 0 {
		return configured
	}
	if configured == 0 || configured > offered {
		return offered
	}
	return configured
}
