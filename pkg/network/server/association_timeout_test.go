// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package server

import (
	"context"
	"testing"
	"time"
)

// TestAssociationContextZeroMeansNoTimeout verifies that AssociationTimeout=0
// disables the limit instead of producing an already-expired context.
func TestAssociationContextZeroMeansNoTimeout(t *testing.T) {
	s := New(WithAssociationTimeout(0))
	s.ctx = context.Background()
	ctx, cancel := s.associationContext()
	defer cancel()

	if err := ctx.Err(); err != nil {
		t.Fatalf("associationContext() is already done with timeout 0: %v", err)
	}
	select {
	case <-ctx.Done():
		t.Fatal("associationContext() is already done with timeout 0")
	default:
	}
}

// TestAssociationContextPositiveTimeout verifies the deadline is applied.
func TestAssociationContextPositiveTimeout(t *testing.T) {
	s := New(WithAssociationTimeout(time.Second))
	s.ctx = context.Background()
	ctx, cancel := s.associationContext()
	defer cancel()

	if _, ok := ctx.Deadline(); !ok {
		t.Fatal("associationContext() has no deadline with a positive timeout")
	}
}
