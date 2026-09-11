// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"fmt"
	"strconv"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// NewForDataset creates empty pixel data using image metadata and the internal
// transfer syntax of a Dataset. The Dataset does not need to contain Pixel
// Data yet, and the Pixel Data VR is derived without caller input.
func NewForDataset(ds *dataset.Dataset) (*Data, error) {
	if ds == nil {
		return nil, fmt.Errorf("dataset cannot be nil")
	}
	syntax := ds.InternalTransferSyntax()
	if syntax == nil {
		return nil, fmt.Errorf("cannot determine Pixel Data VR: missing Transfer Syntax")
	}

	rows, err := ds.GetUInt16(tag.Rows, 0)
	if err != nil {
		return nil, fmt.Errorf("missing or invalid Rows tag: %w", err)
	}
	columns, err := ds.GetUInt16(tag.Columns, 0)
	if err != nil {
		return nil, fmt.Errorf("missing or invalid Columns tag: %w", err)
	}
	bitsAllocated, err := ds.GetUInt16(tag.BitsAllocated, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot determine Pixel Data VR: missing or invalid Bits Allocated: %w", err)
	}
	if syntax.IsEncapsulated() && bitsAllocated > 16 {
		return nil, fmt.Errorf("cannot represent encapsulated Pixel Data with Bits Allocated %d", bitsAllocated)
	}

	bitsStored := ds.TryGetUInt16(tag.BitsStored, 0)
	if bitsStored == 0 {
		bitsStored = bitsAllocated
	}
	highBit := ds.TryGetUInt16(tag.HighBit, 0)
	if highBit == 0 && bitsStored > 1 {
		highBit = bitsStored - 1
	}
	samplesPerPixel := ds.TryGetUInt16(tag.SamplesPerPixel, 0)
	if samplesPerPixel == 0 {
		samplesPerPixel = 1
	}

	photometricValue, _ := ds.GetString(tag.PhotometricInterpretation)
	if photometricValue == "" {
		photometricValue = pixel.Monochrome2.Value
	}
	photometric, err := pixel.ParsePhotometricInterpretation(photometricValue)
	if err != nil {
		return nil, fmt.Errorf("invalid Photometric Interpretation %q: %w", photometricValue, err)
	}

	info := &Info{
		Width:                     columns,
		Height:                    rows,
		NumberOfFrames:            1,
		BitsAllocated:             bitsAllocated,
		BitsStored:                bitsStored,
		HighBit:                   highBit,
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       pixel.Representation(ds.TryGetUInt16(tag.PixelRepresentation, 0)),
		PlanarConfiguration:       pixel.PlanarConfiguration(ds.TryGetUInt16(tag.PlanarConfiguration, 0)),
		PhotometricInterpretation: photometric,
		Encapsulated:              syntax.IsEncapsulated(),
		TransferSyntaxUID:         syntax.UID().UID(),
	}
	if info.Encapsulated {
		info.VRCode = vr.CodeOB
	} else {
		info.VRCode = nativePixelDataVR(info)
	}

	return New(info)
}

// WriteToDataset materializes the current frames as a DICOM Pixel Data
// element, derives OB/OW/fragment representation, and updates Number of
// Frames. The target Dataset transfer syntax must match the pixel data.
func (pd *Data) WriteToDataset(ds *dataset.Dataset) error {
	if pd == nil || pd.Info == nil {
		return fmt.Errorf("pixel data info is nil")
	}
	if ds == nil {
		return fmt.Errorf("dataset cannot be nil")
	}
	syntax := ds.InternalTransferSyntax()
	if syntax == nil {
		return fmt.Errorf("cannot write Pixel Data: missing Transfer Syntax")
	}
	if pd.Info.TransferSyntaxUID != "" && pd.Info.TransferSyntaxUID != syntax.UID().UID() {
		return fmt.Errorf(
			"pixel data transfer syntax %s does not match Dataset Transfer Syntax %s",
			pd.Info.TransferSyntaxUID,
			syntax.UID().UID(),
		)
	}
	if len(pd.frames) == 0 {
		return fmt.Errorf("pixel data has no frames")
	}
	bitsAllocated, err := ds.GetUInt16(tag.BitsAllocated, 0)
	if err != nil {
		return fmt.Errorf("missing or invalid Bits Allocated: %w", err)
	}
	if bitsAllocated != pd.Info.BitsAllocated {
		return fmt.Errorf(
			"pixel data Bits Allocated %d does not match Dataset Bits Allocated %d",
			pd.Info.BitsAllocated,
			bitsAllocated,
		)
	}

	pd.Info.TransferSyntaxUID = syntax.UID().UID()
	pd.Info.NumberOfFrames = len(pd.frames)
	pixelElement, err := pd.ToElement()
	if err != nil {
		return err
	}
	frameCountElement := element.NewIntegerString(tag.NumberOfFrames, []string{strconv.Itoa(len(pd.frames))})
	if err := ds.AddOrUpdate(frameCountElement); err != nil {
		return fmt.Errorf("update Number of Frames: %w", err)
	}
	if err := ds.AddOrUpdate(pixelElement); err != nil {
		return fmt.Errorf("update Pixel Data: %w", err)
	}
	return nil
}
