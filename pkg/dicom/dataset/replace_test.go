// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestDatasetReplaceFromPreservesReceiverValidationModeAndCopiesState(t *testing.T) {
	target := New()
	target.SetAutoValidate(false)
	if err := target.Add(element.NewString(tag.PatientName, vr.PN, []string{"Old^Name"})); err != nil {
		t.Fatal(err)
	}
	targetPointer := target

	source := NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	if err := source.Add(element.NewString(tag.PatientName, vr.PN, []string{"New^Name"})); err != nil {
		t.Fatal(err)
	}
	if err := target.ReplaceFrom(source); err != nil {
		t.Fatal(err)
	}

	if target != targetPointer {
		t.Fatal("ReplaceFrom changed the receiver pointer")
	}
	if target.AutoValidate() {
		t.Fatal("ReplaceFrom changed the receiver automatic-validation mode")
	}
	if target.InternalTransferSyntax() != transfer.ExplicitVRBigEndian {
		t.Fatal("ReplaceFrom did not transfer internal transfer syntax")
	}
	if got, ok := target.GetString(tag.PatientName); !ok || got != "New^Name" {
		t.Fatalf("PatientName = %q, %t", got, ok)
	}

	if err := source.AddOrUpdate(element.NewString(tag.PatientName, vr.PN, []string{"Later^Mutation"})); err != nil {
		t.Fatal(err)
	}
	if got, _ := target.GetString(tag.PatientName); got != "New^Name" {
		t.Fatalf("target aliased source after ReplaceFrom: %q", got)
	}
}

func TestDatasetReplaceFromRejectsNilInputsWithoutMutation(t *testing.T) {
	target := New()
	if err := target.Add(element.NewString(tag.PatientID, vr.LO, []string{"123"})); err != nil {
		t.Fatal(err)
	}

	if err := target.ReplaceFrom(nil); err == nil {
		t.Fatal("ReplaceFrom(nil) succeeded")
	}
	var nilTarget *Dataset
	if err := nilTarget.ReplaceFrom(target); err == nil {
		t.Fatal("nil target ReplaceFrom succeeded")
	}
	if got, _ := target.GetString(tag.PatientID); got != "123" {
		t.Fatalf("target changed after failed ReplaceFrom: %q", got)
	}
}
