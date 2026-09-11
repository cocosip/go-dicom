// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"fmt"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/internal/dicomlut"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
)

// PresentationLUT reads the mutually exclusive Presentation LUT Sequence or
// Presentation LUT Shape. The bool result reports whether either is present.
func PresentationLUT(ds *dataset.Dataset) (lut.LUT, bool, error) {
	if ds == nil {
		return nil, false, nil
	}
	hasSequence := ds.Contains(tag.PresentationLUTSequence)
	hasShape := ds.Contains(tag.PresentationLUTShape)
	if hasSequence && hasShape {
		return nil, true, fmt.Errorf("presentation LUT Sequence and Presentation LUT Shape are mutually exclusive")
	}
	if hasSequence {
		table, err := presentationSequenceLUT(ds)
		return table, true, err
	}
	if !hasShape {
		return nil, false, nil
	}
	valueElement, _ := ds.Get(tag.PresentationLUTShape)
	value, ok := valueElement.(*element.String)
	if !ok || value.ValueRepresentation() != vr.CS {
		return nil, true, fmt.Errorf("presentation LUT Shape must use CS VR")
	}
	if value.Count() != 1 {
		return nil, true, fmt.Errorf("presentation LUT Shape must contain exactly one value")
	}
	switch strings.ToUpper(strings.TrimSpace(value.GetValue(0))) {
	case "IDENTITY":
		return nil, true, nil
	case "INVERSE":
		return lut.NewInvertLUT(0, 255), true, nil
	default:
		return nil, true, fmt.Errorf("unsupported Presentation LUT Shape %q", value.GetValue(0))
	}
}

func presentationSequenceLUT(ds *dataset.Dataset) (lut.LUT, error) {
	sequence, err := ds.GetSequence(tag.PresentationLUTSequence)
	if err != nil {
		return nil, fmt.Errorf("read Presentation LUT Sequence: %w", err)
	}
	if sequence.Count() != 1 || sequence.GetItem(0) == nil {
		return nil, fmt.Errorf("presentation LUT Sequence must contain exactly one item, got %d", sequence.Count())
	}
	item := sequence.GetItem(0)
	descriptor, err := dicomlut.ReadDescriptor(item, tag.LUTDescriptor, false)
	if err != nil {
		return nil, fmt.Errorf("read Presentation LUT descriptor: %w", err)
	}
	if descriptor.BitsPerEntry != 8 && descriptor.BitsPerEntry != 16 {
		return nil, fmt.Errorf("read Presentation LUT descriptor: bits per entry must be 8 or 16, got %d", descriptor.BitsPerEntry)
	}
	values, err := dicomlut.ReadData(item, tag.LUTData, descriptor, dicomlut.ByteOrder(ds))
	if err != nil {
		return nil, fmt.Errorf("read Presentation LUT Data: %w", err)
	}
	table := &voiTableLUT{values: values, first: descriptor.FirstMappedValue}
	maximumOutput := float64(uint32(1)<<descriptor.BitsPerEntry - 1)
	return scaledVOITable(table, 0, maximumOutput), nil
}
