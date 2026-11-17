// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package media

import (
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// Record represents one item in a DICOM Directory Record Sequence.
type Record struct {
	recordType RecordType
	dataset    *dataset.Dataset
	offset     uint32
	children   []*Record
	key        string
}

// recordKey returns the value AddFile merges records by: the Patient ID
// (falling back to the normalized name), the Study or Series Instance UID, or
// the SOP Instance UID. Records loaded from an existing DICOMDIR use the same
// key, so adding a file merges into the loaded hierarchy instead of appending a
// second PATIENT/STUDY/SERIES record for the same entity.
func recordKey(recordType RecordType, ds *dataset.Dataset) string {
	switch recordType {
	case RecordPatient:
		if patientID, ok := stringValue(ds, tag.PatientID); ok && patientID != "" {
			return patientID
		}
		// Fall back to the name only when the (non-conformant) Patient ID is empty.
		name, _ := stringValue(ds, tag.PatientName)
		return "\x00" + normalizePersonName(name)
	case RecordStudy:
		key, _ := stringValue(ds, tag.StudyInstanceUID)
		return key
	case RecordSeries:
		key, _ := stringValue(ds, tag.SeriesInstanceUID)
		return key
	default:
		key, _ := stringValue(ds, tag.SOPInstanceUID)
		return key
	}
}

func (r *Record) addChild(child *Record) {
	r.children = append(r.children, child)
}

func newRecord(recordType RecordType, ds *dataset.Dataset, offset uint32) *Record {
	return &Record{
		recordType: recordType,
		dataset:    ds,
		offset:     offset,
	}
}

// Type returns the Directory Record Type.
func (r *Record) Type() RecordType {
	if r == nil {
		return ""
	}
	return r.recordType
}

// Dataset returns the Dataset stored in this directory record.
func (r *Record) Dataset() *dataset.Dataset {
	if r == nil {
		return nil
	}
	return r.dataset
}

// Offset returns the absolute byte offset of this record's Item Tag.
func (r *Record) Offset() uint32 {
	if r == nil {
		return 0
	}
	return r.offset
}

// Children returns a defensive copy of the record's direct children.
func (r *Record) Children() []*Record {
	if r == nil {
		return nil
	}
	return append([]*Record(nil), r.children...)
}
