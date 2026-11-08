// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

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
	return PresentationLUTWithVRMode(ds, LUTCompatible)
}

// PresentationLUTWithVRMode reads the presentation LUT using the requested
// LUT Data VR policy.
func PresentationLUTWithVRMode(ds *dataset.Dataset, mode LUTVRMode) (lut.LUT, bool, error) {
	if mode != LUTCompatible && mode != LUTStandard {
		return nil, false, fmt.Errorf("unsupported LUT VR mode: %d", mode)
	}
	if ds == nil {
		return nil, false, nil
	}
	hasSequence := ds.Contains(tag.PresentationLUTSequence)
	hasShape := ds.Contains(tag.PresentationLUTShape)
	if hasSequence && hasShape {
		return nil, true, fmt.Errorf("presentation LUT Sequence and Presentation LUT Shape are mutually exclusive")
	}
	if hasSequence {
		table, err := presentationSequenceLUT(ds, mode)
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

func presentationSequenceLUT(ds *dataset.Dataset, mode LUTVRMode) (lut.LUT, error) {
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
	if mode == LUTStandard && (descriptor.BitsPerEntry < 10 || descriptor.BitsPerEntry > 16) {
		return nil, fmt.Errorf("read Presentation LUT descriptor: bits per entry must be 10 through 16, got %d", descriptor.BitsPerEntry)
	}
	var values []uint16
	if mode == LUTStandard {
		values, err = dicomlut.ReadDataStrict(item, tag.LUTData, descriptor, dicomlut.ByteOrder(ds))
	} else {
		values, err = dicomlut.ReadData(item, tag.LUTData, descriptor, dicomlut.ByteOrder(ds))
	}
	if err != nil {
		return nil, fmt.Errorf("read Presentation LUT Data: %w", err)
	}
	table := &voiTableLUT{values: values, first: descriptor.FirstMappedValue}
	maximumOutput := float64(uint32(1)<<descriptor.BitsPerEntry - 1)
	// PS3.3 C.11.6.1: "there is an implicit linear scaling of the output range of
	// the preceding transformation (such as the VOI LUT transformation) so that
	// it is always mapped to the specified input range of the Presentation LUT."
	// The preceding stages emit this library's 0..255 intermediate, so map that
	// range onto the LUT's declared input range before indexing the table.
	inputMinimum := float64(descriptor.FirstMappedValue)
	inputMaximum := inputMinimum + float64(len(values)-1)
	composite := lut.NewCompositeLUT()
	if inputMinimum != 0 || inputMaximum != 255 {
		composite.Add(newPresentationInputScaleLUT(inputMinimum, inputMaximum))
	}
	composite.Add(table)
	composite.Add(newVOIOutputScaleLUT(0, maximumOutput))
	return composite, nil
}

// presentationInputScaleLUT maps the preceding transformation's 0..255 output
// range onto a Presentation LUT's declared input range (PS3.3 C.11.6.1).
type presentationInputScaleLUT struct {
	minimum float64
	maximum float64
}

func newPresentationInputScaleLUT(minimum, maximum float64) *presentationInputScaleLUT {
	if maximum <= minimum {
		maximum = minimum + 1
	}
	return &presentationInputScaleLUT{minimum: minimum, maximum: maximum}
}

func (p *presentationInputScaleLUT) IsValid() bool               { return true }
func (p *presentationInputScaleLUT) MinimumOutputValue() float64 { return p.minimum }
func (p *presentationInputScaleLUT) MaximumOutputValue() float64 { return p.maximum }
func (p *presentationInputScaleLUT) Recalculate()                {}
func (p *presentationInputScaleLUT) Transform(input float64) float64 {
	if input <= 0 {
		return p.minimum
	}
	if input >= 255 {
		return p.maximum
	}
	return p.minimum + input*(p.maximum-p.minimum)/255
}
