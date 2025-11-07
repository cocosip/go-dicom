// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"context"

	"github.com/cocosip/go-dicom/pkg/network/association"
)

// AssociationResponder provides methods to respond to an association request.
// This interface is passed to AssociationNegotiator.OnAssociationRequest to allow
// the negotiator to explicitly send the association response (accept or reject).
//
// This follows the fo-dicom pattern where the provider explicitly calls
// SendAssociationAcceptAsync or SendAssociationRejectAsync.
type AssociationResponder interface {
	// SendAccept sends an A-ASSOCIATE-AC (Association Accept) PDU.
	// The association parameter should contain the negotiated presentation contexts.
	//
	// Example:
	//   return responder.SendAccept(ctx, assoc)
	SendAccept(ctx context.Context, assoc *association.Association) error

	// SendReject sends an A-ASSOCIATE-RJ (Association Reject) PDU.
	//
	// Parameters:
	//   - result: Rejection result (Permanent=1, Transient=2)
	//   - source: Rejection source (ServiceUser=1, ServiceProviderACSE=2, ServiceProviderPresentation=3)
	//   - reason: Rejection reason (depends on source, see DICOM spec)
	//
	// Example:
	//   return responder.SendReject(ctx, pdu.ResultRejectedPermanent, pdu.SourceServiceUser, pdu.ReasonServiceUserCalledAETitleNotRecognized)
	SendReject(ctx context.Context, result, source, reason byte) error
}

// AssociationResponderFuncs holds optional callback functions for responding to association requests.
// This provides a convenient way to implement AssociationResponder by passing functions.
// This is particularly useful for testing or creating mock responders.
//
// Example:
//
//	responder := &service.AssociationResponderFuncs{
//	    SendAcceptFunc: func(ctx context.Context, assoc *association.Association) error {
//	        log.Println("Association accepted")
//	        return nil
//	    },
//	    SendRejectFunc: func(ctx context.Context, result, source, reason byte) error {
//	        log.Printf("Association rejected: result=%d, source=%d, reason=%d", result, source, reason)
//	        return nil
//	    },
//	}
type AssociationResponderFuncs struct {
	SendAcceptFunc func(ctx context.Context, assoc *association.Association) error
	SendRejectFunc func(ctx context.Context, result, source, reason byte) error
}

// SendAccept implements AssociationResponder.
func (r *AssociationResponderFuncs) SendAccept(ctx context.Context, assoc *association.Association) error {
	if r.SendAcceptFunc != nil {
		return r.SendAcceptFunc(ctx, assoc)
	}
	return nil
}

// SendReject implements AssociationResponder.
func (r *AssociationResponderFuncs) SendReject(ctx context.Context, result, source, reason byte) error {
	if r.SendRejectFunc != nil {
		return r.SendRejectFunc(ctx, result, source, reason)
	}
	return nil
}

// AssociationNegotiator provides callbacks for DICOM association negotiation.
// This interface is typically implemented by server applications to control
// which association requests are accepted and how presentation contexts are negotiated.
//
// This follows the fo-dicom pattern where the negotiator receives the full Association
// object, modifies the PresentationContexts, and explicitly sends the response.
//
// Example implementation:
//
//	type MyNegotiator struct {}
//
//	func (n *MyNegotiator) OnAssociationRequest(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
//	    // Validate AE titles
//	    if assoc.CallingAE != "EXPECTED_AE" {
//	        return responder.SendReject(ctx, pdu.ResultRejectedPermanent, pdu.SourceServiceUser, pdu.ReasonServiceUserCalledAETitleNotRecognized)
//	    }
//
//	    // Negotiate presentation contexts
//	    for _, pc := range assoc.PresentationContexts {
//	        if pc.AbstractSyntax == uid.Verification.UID() {
//	            // Accept C-ECHO with any proposed transfer syntax
//	            pc.AcceptTransferSyntaxes(false, pc.ProposedTransferSyntaxes...)
//	        } else if pc.AbstractSyntax == uid.CTImageStorage.UID() {
//	            // Accept CT images with specific transfer syntaxes
//	            ts, _ := transfer.Lookup(uid.ExplicitVRLittleEndian)
//	            pc.AcceptTransferSyntaxes(true, ts)
//	        } else {
//	            // Reject unsupported abstract syntaxes
//	            pc.SetResult(association.ResultAbstractSyntaxNotSupported, nil)
//	        }
//	    }
//
//	    // Explicitly send association accept
//	    return responder.SendAccept(ctx, assoc)
//	}
type AssociationNegotiator interface {
	// OnAssociationRequest is called when an A-ASSOCIATE-RQ is received (server mode).
	//
	// The negotiator receives the complete Association object built from the A-ASSOCIATE-RQ PDU
	// and must explicitly send a response using the provided responder.
	//
	// It can inspect and modify the PresentationContexts to accept or reject individual contexts
	// using methods like:
	//   - pc.AcceptTransferSyntaxes() - accept with matching transfer syntax
	//   - pc.SetResult() - set result and transfer syntax manually
	//   - pc.Accept() or pc.Reject() - simple accept/reject
	//
	// After modifying the association, the negotiator MUST call one of:
	//   - responder.SendAccept(ctx, assoc) - to send A-ASSOCIATE-AC
	//   - responder.SendReject(ctx, result, source, reason) - to send A-ASSOCIATE-RJ
	//
	// Parameters:
	//   - ctx: Request context (can be used for timeout/cancellation)
	//   - assoc: The association built from the A-ASSOCIATE-RQ PDU
	//            The negotiator can modify the PresentationContexts
	//   - responder: Interface to send the association response
	//
	// Returns:
	//   - nil if response was sent successfully
	//   - error if failed to send response
	OnAssociationRequest(ctx context.Context, assoc *association.Association, responder AssociationResponder) error
}

// AssociationReleaseHandler provides callbacks for DICOM association release.
// This interface is typically implemented by applications that need to perform
// cleanup or validation when an association is being released.
//
// Example implementation:
//
//	type MyReleaseHandler struct {}
//
//	func (h *MyReleaseHandler) OnAssociationRelease(ctx context.Context) error {
//	    // Perform cleanup before accepting release
//	    // ...
//	    return nil // Accept the release
//	}
type AssociationReleaseHandler interface {
	// OnAssociationRelease is called when an A-RELEASE-RQ is received.
	// Return nil to accept the release request, or an error to reject it.
	// If rejected, the service will send an A-ABORT instead of A-RELEASE-RP.
	//
	// Parameters:
	//   - ctx: Request context (can be used for timeout/cancellation)
	//
	// Returns:
	//   - nil to accept the release (will send A-RELEASE-RP)
	//   - error to reject the release (will send A-ABORT)
	OnAssociationRelease(ctx context.Context) error
}

// ConnectionLifecycleHandler provides callbacks for connection lifecycle events.
// This interface is typically implemented by applications that need to monitor
// connection status and perform cleanup when connections are closed or aborted.
//
// This follows the fo-dicom IDicomService interface pattern.
//
// Example implementation:
//
//	type MyLifecycleHandler struct {}
//
//	func (h *MyLifecycleHandler) OnAbort(ctx context.Context, source, reason byte) {
//	    log.Printf("Association aborted: source=%d, reason=%d", source, reason)
//	    // Perform cleanup...
//	}
//
//	func (h *MyLifecycleHandler) OnConnectionClosed(ctx context.Context, err error) {
//	    if err != nil {
//	        log.Printf("Connection closed with error: %v", err)
//	    }
//	    // Perform cleanup...
//	}
type ConnectionLifecycleHandler interface {
	// OnAbort is called when an A-ABORT is received or sent.
	// This is an informational callback that cannot prevent the abort.
	// The abort has already occurred when this callback is invoked.
	//
	// This follows fo-dicom's OnReceiveAbort(DicomAbortSource source, DicomAbortReason reason).
	//
	// Parameters:
	//   - ctx: Request context
	//   - source: Abort source (0=ServiceUser, 2=ServiceProvider)
	//   - reason: Abort reason (see pdu.AbortReason* constants)
	OnAbort(ctx context.Context, source, reason byte)

	// OnConnectionClosed is called when the connection is closed.
	// This is an informational callback that is invoked after the
	// connection has already been closed.
	//
	// This follows fo-dicom's OnConnectionClosed(Exception exception).
	//
	// Parameters:
	//   - ctx: Request context
	//   - err: The error that caused the closure (nil if closed normally)
	OnConnectionClosed(ctx context.Context, err error)
}

// DefaultAssociationNegotiator is a simple implementation that accepts all presentation contexts.
// It accepts each proposed presentation context with the first proposed transfer syntax.
type DefaultAssociationNegotiator struct{}

// OnAssociationRequest accepts all association requests and accepts all presentation contexts
// with their first proposed transfer syntax.
func (n *DefaultAssociationNegotiator) OnAssociationRequest(ctx context.Context, assoc *association.Association, responder AssociationResponder) error {
	// Accept all presentation contexts with first transfer syntax
	for _, pc := range assoc.PresentationContexts {
		if len(pc.ProposedTransferSyntaxes) > 0 {
			pc.Accept(pc.ProposedTransferSyntaxes[0])
		} else {
			// No transfer syntaxes proposed, reject
			pc.Reject(association.ResultTransferSyntaxesNotSupported)
		}
	}

	// Send association accept
	return responder.SendAccept(ctx, assoc)
}

// DefaultAssociationReleaseHandler is a simple implementation that accepts all release requests.
type DefaultAssociationReleaseHandler struct{}

// OnAssociationRelease accepts all release requests.
func (h *DefaultAssociationReleaseHandler) OnAssociationRelease(ctx context.Context) error {
	return nil // Accept all
}

// DefaultConnectionLifecycleHandler is a no-op implementation of ConnectionLifecycleHandler.
type DefaultConnectionLifecycleHandler struct{}

// OnAbort does nothing (no-op).
func (h *DefaultConnectionLifecycleHandler) OnAbort(ctx context.Context, source, reason byte) {
	// No-op
}

// OnConnectionClosed does nothing (no-op).
func (h *DefaultConnectionLifecycleHandler) OnConnectionClosed(ctx context.Context, err error) {
	// No-op
}

// FuncAssociationNegotiator is a simple function-based implementation of AssociationNegotiator.
// It allows using a simple function as an AssociationNegotiator without defining a new type.
//
// Example:
//
//	negotiator := service.FuncAssociationNegotiator(func(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
//	    // Accept all presentation contexts with first transfer syntax
//	    for _, pc := range assoc.PresentationContexts {
//	        if len(pc.ProposedTransferSyntaxes) > 0 {
//	            pc.Accept(pc.ProposedTransferSyntaxes[0])
//	        }
//	    }
//	    return responder.SendAccept(ctx, assoc)
//	})
type FuncAssociationNegotiator func(ctx context.Context, assoc *association.Association, responder AssociationResponder) error

// OnAssociationRequest implements AssociationNegotiator.
func (f FuncAssociationNegotiator) OnAssociationRequest(ctx context.Context, assoc *association.Association, responder AssociationResponder) error {
	return f(ctx, assoc, responder)
}

// FuncAssociationReleaseHandler is a simple function-based implementation of AssociationReleaseHandler.
// It allows using a simple function as an AssociationReleaseHandler without defining a new type.
//
// Example:
//
//	handler := service.FuncAssociationReleaseHandler(func(ctx context.Context) error {
//	    // Perform cleanup
//	    return nil // Accept release
//	})
type FuncAssociationReleaseHandler func(ctx context.Context) error

// OnAssociationRelease implements AssociationReleaseHandler.
func (f FuncAssociationReleaseHandler) OnAssociationRelease(ctx context.Context) error {
	return f(ctx)
}

// ConnectionLifecycleHandlerFuncs holds optional callback functions for connection lifecycle events.
// This provides a convenient way to implement ConnectionLifecycleHandler by passing functions.
//
// Example:
//
//	handler := &service.ConnectionLifecycleHandlerFuncs{
//	    OnAbortFunc: func(ctx context.Context, source, reason byte) {
//	        log.Printf("Connection aborted: source=%d, reason=%d", source, reason)
//	    },
//	    OnConnectionClosedFunc: func(ctx context.Context, err error) {
//	        if err != nil {
//	            log.Printf("Connection closed with error: %v", err)
//	        }
//	    },
//	}
type ConnectionLifecycleHandlerFuncs struct {
	OnAbortFunc            func(ctx context.Context, source, reason byte)
	OnConnectionClosedFunc func(ctx context.Context, err error)
}

// OnAbort implements ConnectionLifecycleHandler.
func (h *ConnectionLifecycleHandlerFuncs) OnAbort(ctx context.Context, source, reason byte) {
	if h.OnAbortFunc != nil {
		h.OnAbortFunc(ctx, source, reason)
	}
}

// OnConnectionClosed implements ConnectionLifecycleHandler.
func (h *ConnectionLifecycleHandlerFuncs) OnConnectionClosed(ctx context.Context, err error) {
	if h.OnConnectionClosedFunc != nil {
		h.OnConnectionClosedFunc(ctx, err)
	}
}
