// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/internal/dicomlut"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// DatasetValueRange reads and validates the smallest and largest image pixel values.
func DatasetValueRange(ds *dataset.Dataset) (float64, float64, error) {
	minimum, err := dicomlut.ShortValue(ds, tag.SmallestImagePixelValue)
	if err != nil {
		return 0, 0, err
	}
	maximum, err := dicomlut.ShortValue(ds, tag.LargestImagePixelValue)
	if err != nil {
		return 0, 0, err
	}
	if minimum >= maximum {
		return 0, 0, fmt.Errorf("invalid image pixel value range %d..%d", minimum, maximum)
	}
	return float64(minimum), float64(maximum), nil
}

// VOILUT reads the first VOI LUT Sequence item from ds.
func VOILUT(ds *dataset.Dataset, signed bool) (lut.LUT, error) {
	return VOILUTAt(ds, signed, 0)
}

// VOILUTAt reads one VOI LUT Sequence item from ds.
func VOILUTAt(ds *dataset.Dataset, signed bool, index int) (lut.LUT, error) {
	sequence, err := ds.GetSequence(tag.VOILUTSequence)
	if err != nil || sequence.Count() == 0 {
		return nil, fmt.Errorf("VOI LUT Sequence is missing or empty")
	}
	if index < 0 || index >= sequence.Count() {
		return nil, fmt.Errorf("VOI LUT Sequence item index %d is out of range [0, %d)", index, sequence.Count())
	}
	item := sequence.GetItem(index)
	if item == nil {
		return nil, fmt.Errorf("VOI LUT Sequence item %d is nil", index)
	}
	descriptor, err := dicomlut.ReadDescriptor(item, tag.LUTDescriptor, signed)
	if err != nil {
		return nil, fmt.Errorf("read VOI LUT descriptor: %w", err)
	}
	if err := validateImageVOILUTDescriptor(descriptor); err != nil {
		return nil, fmt.Errorf("read VOI LUT descriptor: %w", err)
	}
	values, err := dicomlut.ReadData(item, tag.LUTData, descriptor, dicomlut.ByteOrder(ds))
	if err != nil {
		return nil, fmt.Errorf("read VOI LUT Data: %w", err)
	}
	table := &voiTableLUT{values: values, first: descriptor.FirstMappedValue}
	maximumOutput := float64(uint32(1)<<descriptor.BitsPerEntry - 1)
	return scaledVOITable(table, 0, maximumOutput), nil
}

// VOILUTFrom reads a VOI LUT from primary or falls back to fallback.
func VOILUTFrom(primary, fallback *dataset.Dataset, signed bool, index int) (lut.LUT, error) {
	if primary != nil && primary.Contains(tag.VOILUTSequence) {
		return VOILUTAt(primary, signed, index)
	}
	if fallback != nil && fallback.Contains(tag.VOILUTSequence) {
		return VOILUTAt(fallback, signed, index)
	}
	return nil, fmt.Errorf("VOI LUT Sequence is missing")
}

// ModalityLUT reads the single Modality LUT Sequence item from ds.
func ModalityLUT(ds *dataset.Dataset, signed bool) (lut.LUT, error) {
	sequence, err := ds.GetSequence(tag.ModalityLUTSequence)
	if err != nil {
		return nil, fmt.Errorf("modality LUT Sequence is missing: %w", err)
	}
	if sequence.Count() != 1 || sequence.GetItem(0) == nil {
		return nil, fmt.Errorf("modality LUT Sequence must contain exactly one item, got %d", sequence.Count())
	}
	item := sequence.GetItem(0)
	descriptor, err := dicomlut.ReadDescriptor(item, tag.LUTDescriptor, signed)
	if err != nil {
		return nil, fmt.Errorf("read Modality LUT descriptor: %w", err)
	}
	if descriptor.BitsPerEntry != 8 && descriptor.BitsPerEntry != 16 {
		return nil, fmt.Errorf("read Modality LUT descriptor: bits per entry must be 8 or 16, got %d", descriptor.BitsPerEntry)
	}
	entries, err := dicomlut.ReadData(item, tag.LUTData, descriptor, dicomlut.ByteOrder(ds))
	if err != nil {
		return nil, fmt.Errorf("read Modality LUT Data: %w", err)
	}
	values := make([]float64, len(entries))
	for index, entry := range entries {
		values[index] = float64(entry)
	}
	return lut.NewModalitySequenceLUT(values, descriptor.FirstMappedValue, signed), nil
}

// applyVOILUT applies VOI LUT Sequence and maps its declared output range to 8-bit.
func applyVOILUT(pd *Data, ds *dataset.Dataset, _, _ float64, ignorePadding bool) ([][]byte, error) {
	signed := pd != nil && pd.Info != nil && pd.Info.PixelRepresentation == pixel.SignedPixels
	table, err := VOILUT(ds, signed)
	if err != nil {
		return nil, err
	}
	return mapThroughLUT(pd, table, ignorePadding)
}

// voiTableLUT adapts VOI LUT Data to the shared LUT interface.
type voiTableLUT struct {
	values []uint16
	first  int
}

func validateImageVOILUTDescriptor(descriptor dicomlut.Descriptor) error {
	if descriptor.BitsPerEntry != 8 && descriptor.BitsPerEntry != 16 {
		return fmt.Errorf("bits per entry must be 8 or 16 for an Image IOD, got %d", descriptor.BitsPerEntry)
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

func mapThroughLUT(pd *Data, table lut.LUT, ignorePadding bool) ([][]byte, error) {
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
