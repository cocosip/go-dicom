// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// Directory represents a DICOM media storage directory.
// Directory mutations are not safe for concurrent use.
type Directory struct {
	dataset        *dataset.Dataset
	fileMeta       *dataset.Dataset
	transferSyntax *transfer.Syntax
	roots          []*Record
	diagnostics    []Diagnostic
	iconGenerator  IconGenerator
	imageIcons     bool
}

// NewDirectory creates an empty DICOMDIR.
func NewDirectory(options ...Option) (*Directory, error) {
	config := config{transferSyntax: transfer.ExplicitVRLittleEndian}
	for _, option := range options {
		option(&config)
	}
	if err := validateDirectoryTransferSyntax(config.transferSyntax); err != nil {
		return nil, err
	}

	sopInstanceUID := uid.GenerateDerivedFromUUID().UID()
	fmi := dataset.NewDefaultFileMetaInformation()
	if err := fmi.SetMediaStorageSOPClassUID(uid.MediaStorageDirectoryStorage.UID()); err != nil {
		return nil, fmt.Errorf("set Media Storage SOP Class UID: %w", err)
	}
	if err := fmi.SetMediaStorageSOPInstanceUID(sopInstanceUID); err != nil {
		return nil, fmt.Errorf("set Media Storage SOP Instance UID: %w", err)
	}
	if err := fmi.SetTransferSyntax(config.transferSyntax); err != nil {
		return nil, fmt.Errorf("set DICOMDIR transfer syntax: %w", err)
	}

	ds := dataset.NewWithTransferSyntax(config.transferSyntax)
	for _, item := range []element.Element{
		element.NewString(tag.FileSetID, vr.CS, []string{""}),
		element.NewUnsignedLong(tag.OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity, []uint32{0}),
		element.NewUnsignedLong(tag.OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity, []uint32{0}),
		element.NewUnsignedShort(tag.FileSetConsistencyFlag, []uint16{0}),
		dataset.NewSequence(tag.DirectoryRecordSequence),
	} {
		if err := ds.Add(item); err != nil {
			return nil, fmt.Errorf("add DICOMDIR element %s: %w", item.Tag(), err)
		}
	}

	return &Directory{
		dataset:        ds,
		fileMeta:       fmi.Dataset(),
		transferSyntax: config.transferSyntax,
		iconGenerator:  config.iconGenerator,
		imageIcons:     config.imageIcons,
	}, nil
}

func validateDirectoryTransferSyntax(ts *transfer.Syntax) error {
	if ts == nil {
		return fmt.Errorf("DICOMDIR transfer syntax cannot be nil")
	}
	if ts != transfer.ExplicitVRLittleEndian && ts != transfer.ImplicitVRLittleEndian {
		return fmt.Errorf("unsupported DICOMDIR transfer syntax %s", ts.UID())
	}
	return nil
}

// Dataset returns the DICOMDIR main Dataset.
func (d *Directory) Dataset() *dataset.Dataset {
	if d == nil {
		return nil
	}
	return d.dataset
}

// FileMetaInformation returns the DICOMDIR File Meta Information Dataset.
func (d *Directory) FileMetaInformation() *dataset.Dataset {
	if d == nil {
		return nil
	}
	return d.fileMeta
}

// RootRecords returns a defensive copy of the root directory records.
func (d *Directory) RootRecords() []*Record {
	if d == nil {
		return nil
	}
	return append([]*Record(nil), d.roots...)
}

// Diagnostics returns a defensive copy of non-fatal DICOMDIR diagnostics.
func (d *Directory) Diagnostics() []Diagnostic {
	if d == nil {
		return nil
	}
	return append([]Diagnostic(nil), d.diagnostics...)
}

// Walk visits directory records in depth-first pre-order.
func (d *Directory) Walk(visitor func(*Record) error) error {
	if d == nil {
		return fmt.Errorf("directory cannot be nil")
	}
	if visitor == nil {
		return fmt.Errorf("record visitor cannot be nil")
	}
	var walk func([]*Record) error
	walk = func(records []*Record) error {
		for _, record := range records {
			if err := visitor(record); err != nil {
				return err
			}
			if err := walk(record.children); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(d.roots)
}
