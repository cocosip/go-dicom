// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parseable"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// SOP Class UIDs for testing
const (
	testSOPClassUID = "1.2.840.10008.5.1.4.1.1.2" // CT Image Storage
	testTsUID       = "1.2.840.10008.1.2.1"       // Explicit VR Little Endian
)

func TestDataset_GetUID(t *testing.T) {
	t.Parallel()
	ds := New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testSOPClassUID}))

	result, err := ds.GetUID(tag.SOPClassUID)
	if err != nil {
		t.Fatalf("GetUID() error = %v", err)
	}
	if result.UID() != testSOPClassUID {
		t.Errorf("UID() = %q, want %q", result.UID(), testSOPClassUID)
	}
}

func TestDataset_GetUID_NotFound(t *testing.T) {
	t.Parallel()
	ds := New()
	_, err := ds.GetUID(tag.SOPClassUID)
	if err == nil {
		t.Error("GetUID() expected error for missing element")
	}
}

func TestDataset_GetTransferSyntax(t *testing.T) {
	t.Parallel()
	ds := New()
	_ = ds.Add(element.NewString(tag.TransferSyntaxUID, vr.UI, []string{testTsUID}))

	result, err := ds.GetTransferSyntax()
	if err != nil {
		t.Fatalf("GetTransferSyntax() error = %v", err)
	}
	if result.UID().UID() != testTsUID {
		t.Errorf("UID() = %q, want %q", result.UID().UID(), testTsUID)
	}
}

func TestDataset_GetParseable(t *testing.T) {
	t.Parallel()
	ds := New()
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testSOPClassUID}))

	result, err := GetParseable(ds, tag.SOPClassUID, parseable.ParserFor(func() *uid.UID { return &uid.UID{} }))
	if err != nil {
		t.Fatalf("GetParseable() error = %v", err)
	}
	if result.UID() != testSOPClassUID {
		t.Errorf("UID() = %q, want %q", result.UID(), testSOPClassUID)
	}
}

func TestDataset_GetParseable_TransferSyntax(t *testing.T) {
	t.Parallel()
	ds := New()
	_ = ds.Add(element.NewString(tag.TransferSyntaxUID, vr.UI, []string{testTsUID}))

	result, err := GetParseable(ds, tag.TransferSyntaxUID, parseable.ParserFor(func() *transfer.Syntax { return &transfer.Syntax{} }))
	if err != nil {
		t.Fatalf("GetParseable() error = %v", err)
	}
	if result.UID().UID() != testTsUID {
		t.Errorf("UID() = %q, want %q", result.UID().UID(), testTsUID)
	}
}
