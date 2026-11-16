// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package service

import (
	"context"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// TestRegisterActiveOperationDecoratesEveryRegistration guards the operation ID
// that links an inbound C-FIND/C-MOVE/C-GET handler to the outbound
// sub-operations it sends. The request dispatcher registers the operation once
// before dispatching and the handler registers it again to own the cancelable
// context; both calls must return a context carrying the operation ID.
func TestRegisterActiveOperationDecoratesEveryRegistration(t *testing.T) {
	service := NewService(nil, nil)
	req := dimse.NewCFindRequest(dimse.QueryRetrieveLevelStudy, dataset.New())
	if err := req.SetMessageID(7); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	if _, err := service.registerInboundRequest(req); err != nil {
		t.Fatalf("registerInboundRequest() error = %v", err)
	}

	first, unregisterFirst := service.registerActiveOperation(context.Background(), req.MessageID())
	defer unregisterFirst()
	want := operationIDFromContext(first)
	if want == 0 {
		t.Fatal("first registration context has no operation ID")
	}

	second, unregisterSecond := service.registerActiveOperation(context.Background(), req.MessageID())
	defer unregisterSecond()
	if got := operationIDFromContext(second); got != want {
		t.Fatalf("second registration operation ID = %d, want %d", got, want)
	}
}
