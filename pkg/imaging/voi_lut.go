// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
)

// applyVOILUT applies VOI LUT Sequence and maps its declared output range to 8-bit.
func applyVOILUT(pd *DicomPixelData, ds *dataset.Dataset, _, _ float64, ignorePadding bool) ([][]byte, error) {
	if ds == nil {
		return nil, fmt.Errorf("dataset is nil for VOI LUT")
	}

	seqElem, ok := ds.Get(tag.VOILUTSequence)
	if !ok {
		return nil, fmt.Errorf("VOI LUT Sequence not present")
	}
	seq, ok := seqElem.(*dataset.Sequence)
	if !ok || seq.Count() == 0 {
		return nil, fmt.Errorf("VOI LUT Sequence empty or invalid")
	}

	// Take first LUT item (fo-dicom also picks first)
	item := seq.GetItem(0)

	signed := pd != nil && pd.Info != nil && pd.Info.PixelRepresentation == SignedPixels
	descriptor, err := readLUTDescriptor(item, tag.LUTDescriptor, signed)
	if err != nil {
		return nil, fmt.Errorf("missing LUT Descriptor: %w", err)
	}
	if err := validateImageVOILUTDescriptor(descriptor); err != nil {
		return nil, fmt.Errorf("invalid LUT Descriptor: %w", err)
	}
	values, err := readLUTData(item, tag.LUTData, descriptor, datasetByteOrder(ds))
	if err != nil {
		return nil, fmt.Errorf("read LUT Data: %w", err)
	}

	table := &voiTableLUT{
		values: values,
		first:  descriptor.firstMappedValue,
	}
	maximumOutput := float64(uint32(1)<<descriptor.bitsPerEntry - 1)
	return mapThroughLUT(pd, scaledVOITable(table, 0, maximumOutput), ignorePadding)
}

// voiTableLUT adapts VOI LUT Data to the shared LUT interface.
type voiTableLUT struct {
	values []uint16
	first  int
}

func validateImageVOILUTDescriptor(descriptor lutDescriptor) error {
	if descriptor.bitsPerEntry != 8 && descriptor.bitsPerEntry != 16 {
		return fmt.Errorf("bits per entry must be 8 or 16 for an Image IOD, got %d", descriptor.bitsPerEntry)
	}
	return nil
}

func (v *voiTableLUT) IsValid() bool {
	return len(v.values) > 0
}

func (v *voiTableLUT) MinimumOutputValue() float64 {
	if len(v.values) == 0 {
		return 0
	}
	minimum := v.values[0]
	for _, value := range v.values[1:] {
		if value < minimum {
			minimum = value
		}
	}
	return float64(minimum)
}

func (v *voiTableLUT) MaximumOutputValue() float64 {
	if len(v.values) == 0 {
		return 0
	}
	maximum := v.values[0]
	for _, value := range v.values[1:] {
		if value > maximum {
			maximum = value
		}
	}
	return float64(maximum)
}
func (v *voiTableLUT) Recalculate() {}

func (v *voiTableLUT) Transform(input float64) float64 {
	idx := int(input) - v.first
	if idx < 0 {
		idx = 0
	}
	if idx >= len(v.values) {
		idx = len(v.values) - 1
	}
	return float64(v.values[idx])
}

type voiOutputScaleLUT struct {
	minimum float64
	maximum float64
}

func newVOIOutputScaleLUT(minimum, maximum float64) *voiOutputScaleLUT {
	if maximum <= minimum {
		maximum = minimum + 1
	}
	return &voiOutputScaleLUT{minimum: minimum, maximum: maximum}
}

func (v *voiOutputScaleLUT) IsValid() bool               { return true }
func (v *voiOutputScaleLUT) MinimumOutputValue() float64 { return 0 }
func (v *voiOutputScaleLUT) MaximumOutputValue() float64 { return 255 }
func (v *voiOutputScaleLUT) Recalculate()                {}
func (v *voiOutputScaleLUT) Transform(input float64) float64 {
	if input <= v.minimum {
		return 0
	}
	if input >= v.maximum {
		return 255
	}
	return (input - v.minimum) * 255 / (v.maximum - v.minimum)
}

func scaledVOITable(table *voiTableLUT, minimum, maximum float64) lut.LUT {
	composite := lut.NewCompositeLUT()
	composite.Add(table)
	composite.Add(newVOIOutputScaleLUT(minimum, maximum))
	return composite
}

func mapThroughLUT(pd *DicomPixelData, table lut.LUT, ignorePadding bool) ([][]byte, error) {
	if pd.Info == nil {
		return nil, fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return nil, fmt.Errorf("cannot apply LUT on encapsulated data; decode first")
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return nil, fmt.Errorf("unsupported BytesAllocated=%d for LUT mapping", bytesPerSample)
	}

	hasPadding := pd.Info.PixelPaddingValue != nil
	var padMin, padMax int64
	if hasPadding {
		padMin = int64(*pd.Info.PixelPaddingValue)
		if pd.Info.PixelPaddingRangeLimit != nil {
			padMax = int64(*pd.Info.PixelPaddingRangeLimit)
		} else {
			padMax = padMin
		}
	}

	result := make([][]byte, len(pd.frames))

	for fi, frame := range pd.frames {
		out := make([]byte, len(frame)/bytesPerSample)

		for idx, off := 0, 0; off+bytesPerSample <= len(frame); off, idx = off+bytesPerSample, idx+1 {
			val, ok := decodePixelSampleLE(frame, off, pd.Info)
			if !ok {
				continue
			}

			if ignorePadding && hasPadding && val >= padMin && val <= padMax {
				out[idx] = 0
				continue
			}

			mapped := table.Transform(float64(val))
			out[idx] = clampByte(int(mapped + 0.5))
		}

		result[fi] = out
	}

	return result, nil
}
