// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const maximumDirectoryIconDimension = 128

// NewIconImageSequence creates a one-frame 8-bit MONOCHROME2 DICOMDIR icon.
func NewIconImageSequence(width, height int, pixels []byte) (*dataset.Sequence, error) {
	if width < 1 || width > maximumDirectoryIconDimension || height < 1 || height > maximumDirectoryIconDimension {
		return nil, fmt.Errorf("DICOMDIR icon dimensions must be between 1 and %d pixels", maximumDirectoryIconDimension)
	}
	if len(pixels) != width*height {
		return nil, fmt.Errorf("DICOMDIR icon pixel length is %d, want %d", len(pixels), width*height)
	}

	item := dataset.New()
	elements := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{uint16(height)}),   // #nosec G115 -- dimensions are bounded above
		element.NewUnsignedShort(tag.Columns, []uint16{uint16(width)}), // #nosec G115 -- dimensions are bounded above
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}),
		element.NewString(tag.NumberOfFrames, vr.IS, []string{"1"}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewOtherByte(tag.PixelData, append([]byte(nil), pixels...)),
	}
	for _, value := range elements {
		if err := item.Add(value); err != nil {
			return nil, fmt.Errorf("add DICOMDIR icon element %s: %w", value.Tag(), err)
		}
	}
	return dataset.NewSequenceWithItems(tag.IconImageSequence, []*dataset.Dataset{item}), nil
}
