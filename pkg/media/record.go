// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import "github.com/cocosip/go-dicom/pkg/dicom/dataset"

// Record represents one item in a DICOM Directory Record Sequence.
type Record struct {
	recordType RecordType
	dataset    *dataset.Dataset
	offset     uint32
	children   []*Record
	key        string
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
