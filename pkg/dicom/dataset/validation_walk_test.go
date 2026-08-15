// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"errors"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

const referencedImageSequencePath = "(0008,1140)"

func TestDatasetValidateReturnsStructuralErrorForTypedNilSequence(t *testing.T) {
	ds := New()
	var sequence *Sequence
	ds.items[tag.ReferencedImageSequence.ToUint32()] = sequence

	err := ds.Validate()
	var validationErr *ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("Validate() error = %v, want *ValidationError", err)
	}
	if validationErr.Kind != ValidationStructural {
		t.Fatalf("validation kind = %q, want %q", validationErr.Kind, ValidationStructural)
	}
	if got := FormatPath(validationErr.Path); got != referencedImageSequencePath {
		t.Fatalf("validation path = %q, want %s", got, referencedImageSequencePath)
	}
}
