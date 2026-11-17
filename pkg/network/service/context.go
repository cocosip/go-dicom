// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"

	"github.com/cocosip/go-dicom/pkg/network/association"
)

// associationContextKey is the context key for the association that carried an
// inbound DIMSE request.
type associationContextKey struct{}

// withAssociation attaches the association to a handler context.
func withAssociation(ctx context.Context, assoc *association.Association) context.Context {
	if assoc == nil {
		return ctx
	}
	return context.WithValue(ctx, associationContextKey{}, assoc)
}

// AssociationFromContext returns the association that carried the inbound DIMSE
// request being handled, or nil when ctx does not come from a Service handler.
//
// Handlers use it to read the negotiated association parameters, for example
// the calling AE title, the called AE title or the accepted presentation
// contexts (fo-dicom exposes the same information as DicomService.Association).
func AssociationFromContext(ctx context.Context) *association.Association {
	if ctx == nil {
		return nil
	}
	assoc, _ := ctx.Value(associationContextKey{}).(*association.Association)
	return assoc
}
