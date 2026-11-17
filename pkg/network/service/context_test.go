// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/status"
)

// TestAssociationFromContextReportsCallingAE verifies that an SCP handler can
// read the negotiated association (here the calling AE title) from its context.
func TestAssociationFromContextReportsCallingAE(t *testing.T) {
	assoc := createTestAssociation()
	captured := make(chan *association.Association, 1)

	service := NewService(nil, assoc,
		WithCStoreHandler(func(ctx context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
			captured <- AssociationFromContext(ctx)
			return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
		}))
	t.Cleanup(func() { _ = service.Close() })
	drainSuccessfulSends(service)

	if err := service.handleRequest(context.Background(), newRoleTestCStoreRequest(t)); err != nil {
		t.Fatalf("handleRequest() error = %v", err)
	}

	select {
	case got := <-captured:
		if got != assoc {
			t.Fatalf("AssociationFromContext() = %p, want %p", got, assoc)
		}
		if got.CallingAE != "TEST-SCU" {
			t.Fatalf("CallingAE = %q, want TEST-SCU", got.CallingAE)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("C-STORE handler was not called")
	}
}

func TestAssociationFromContextWithoutServiceHandler(t *testing.T) {
	if got := AssociationFromContext(context.Background()); got != nil {
		t.Fatalf("AssociationFromContext() = %p, want nil", got)
	}
	//nolint:staticcheck // an explicit nil context must not panic.
	if got := AssociationFromContext(nil); got != nil {
		t.Fatalf("AssociationFromContext(nil) = %p, want nil", got)
	}
}
