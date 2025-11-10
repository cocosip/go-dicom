// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package sr

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// ReferencedSOP represents a reference to a SOP Instance in DICOM SR.
//
// A referenced SOP consists of:
//   - Referenced SOP Instance UID: The unique identifier of the referenced instance
//   - Referenced SOP Class UID: The SOP class of the referenced instance
//
// This is used in IMAGE, COMPOSITE, and WAVEFORM content items to reference
// other DICOM instances.
//
// Reference: DICOM Part 3, Section 10.4
type ReferencedSOP struct {
	dataset *dataset.Dataset
}

// NewReferencedSOP creates a new referenced SOP
func NewReferencedSOP(instanceUID, classUID string) *ReferencedSOP {
	ds := dataset.New()

	// Add Referenced SOP Instance UID (0008,1155) VR=UI
	ds.AddOrUpdate(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{instanceUID}))

	// Add Referenced SOP Class UID (0008,1150) VR=UI
	ds.AddOrUpdate(element.NewString(tag.ReferencedSOPClassUID, vr.UI, []string{classUID}))

	return &ReferencedSOP{dataset: ds}
}

// NewReferencedSOPFromDataset creates a ReferencedSOP from an existing dataset
func NewReferencedSOPFromDataset(ds *dataset.Dataset) *ReferencedSOP {
	if ds == nil {
		return nil
	}
	return &ReferencedSOP{dataset: ds}
}

// NewReferencedSOPFromSequence creates a ReferencedSOP from the first item in a sequence
func NewReferencedSOPFromSequence(t *tag.Tag, ds *dataset.Dataset) (*ReferencedSOP, error) {
	seq, err := ds.GetSequence(t)
	if err != nil {
		return nil, WrapError("sequence not found", err)
	}

	if seq.Count() == 0 {
		return nil, ErrEmptySequence
	}

	return NewReferencedSOPFromDataset(seq.GetItem(0)), nil
}

// InstanceUID returns the Referenced SOP Instance UID
func (r *ReferencedSOP) InstanceUID() string {
	if r == nil || r.dataset == nil {
		return ""
	}
	val, _ := r.dataset.GetString(tag.ReferencedSOPInstanceUID)
	return val
}

// ClassUID returns the Referenced SOP Class UID
func (r *ReferencedSOP) ClassUID() string {
	if r == nil || r.dataset == nil {
		return ""
	}
	val, _ := r.dataset.GetString(tag.ReferencedSOPClassUID)
	return val
}

// Dataset returns the underlying dataset
func (r *ReferencedSOP) Dataset() *dataset.Dataset {
	if r == nil {
		return nil
	}
	return r.dataset
}
