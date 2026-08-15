// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
)

// applyVOILUT applies VOI LUT Sequence if present; otherwise falls back to windowing.
// For now, LUT output is clamped to 8-bit; LUT entry values are assumed to be 8 or 16 bit.
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

	// LUT Descriptor: [#entries, first mapped pixel value, bits per entry]
	descVal, err := item.GetUInt16(tag.LUTDescriptor, 0)
	if err != nil {
		return nil, fmt.Errorf("missing LUT Descriptor: %w", err)
	}
	numEntries := int(descVal)
	firstMap := int16(item.TryGetUInt16(tag.LUTDescriptor, 1))
	bitsPerEntry := item.TryGetUInt16(tag.LUTDescriptor, 2)
	if numEntries == 0 {
		numEntries = 65536
	}
	if bitsPerEntry == 0 {
		bitsPerEntry = 16
	}
	if bitsPerEntry < 8 || bitsPerEntry > 16 {
		return nil, fmt.Errorf("invalid VOI LUT bits per entry: %d", bitsPerEntry)
	}

	lutDataElem, ok := item.Get(tag.LUTData)
	if !ok {
		return nil, fmt.Errorf("missing LUT Data")
	}
	var lutRaw []byte
	switch v := lutDataElem.(type) {
	case *element.OtherByte:
		lutRaw = v.GetData()
	case *element.OtherWord:
		lutRaw = v.GetData()
	default:
		return nil, fmt.Errorf("unsupported LUT Data element %T", lutDataElem)
	}

	values := make([]uint16, numEntries)
	byteOrder := datasetByteOrder(ds)
	if bitsPerEntry <= 8 {
		for i := 0; i < numEntries && i < len(lutRaw); i++ {
			values[i] = uint16(lutRaw[i])
		}
	} else {
		for i := 0; i < numEntries && (i*2+1) < len(lutRaw); i++ {
			values[i] = byteOrder.Uint16(lutRaw[i*2:])
		}
	}

	table := &voiTableLUT{
		values: values,
		first:  int(firstMap),
	}
	maximumOutput := float64(uint32(1)<<bitsPerEntry - 1)
	return mapThroughLUT(pd, scaledVOITable(table, 0, maximumOutput), ignorePadding)
}

// voiTableLUT adapts VOI LUT Data to the shared LUT interface.
type voiTableLUT struct {
	values []uint16
	first  int
}

func (v *voiTableLUT) IsValid() bool {
	return false
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

func normalizedVOITable(table *voiTableLUT) lut.LUT {
	return scaledVOITable(table, table.MinimumOutputValue(), table.MaximumOutputValue())
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
