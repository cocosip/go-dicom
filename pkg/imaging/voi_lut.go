// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
)

// applyVOILUT applies VOI LUT Sequence if present; otherwise falls back to windowing.
// For now, LUT output is clamped to 8-bit; LUT entry values are assumed to be 8 or 16 bit.
func applyVOILUT(pd *DicomPixelData, ds *dataset.Dataset, center, width float64, ignorePadding bool) ([][]byte, error) {
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
	if bitsPerEntry <= 8 {
		for i := 0; i < numEntries && i < len(lutRaw); i++ {
			values[i] = uint16(lutRaw[i])
		}
	} else {
		for i := 0; i < numEntries && (i*2+1) < len(lutRaw); i++ {
			values[i] = binary.LittleEndian.Uint16(lutRaw[i*2:])
		}
	}

	table := &voiTableLUT{
		values: values,
		first:  int(firstMap),
	}
	precalc := lut.NewPrecalculatedLUT(table, table.minInput(), table.maxInput())
	return mapThroughLUT(pd, precalc, ignorePadding)
}

// voiTableLUT adapts VOI LUT Data to the shared LUT interface.
type voiTableLUT struct {
	values []uint16
	first  int
}

func (v *voiTableLUT) IsValid() bool {
	return false
}

func (v *voiTableLUT) MinimumOutputValue() float64 { return 0 }
func (v *voiTableLUT) MaximumOutputValue() float64 { return 255 }
func (v *voiTableLUT) Recalculate()                {}

func (v *voiTableLUT) Transform(input float64) float64 {
	idx := int(input) - v.first
	if idx < 0 {
		idx = 0
	}
	if idx >= len(v.values) {
		idx = len(v.values) - 1
	}
	return float64(clampByte(int(v.values[idx])))
}

func (v *voiTableLUT) minInput() int { return v.first }
func (v *voiTableLUT) maxInput() int { return v.first + len(v.values) - 1 }

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
	var padMin, padMax int32
	if hasPadding {
		padMin = *pd.Info.PixelPaddingValue
		if pd.Info.PixelPaddingRangeLimit != nil {
			padMax = *pd.Info.PixelPaddingRangeLimit
		} else {
			padMax = padMin
		}
	}

	result := make([][]byte, len(pd.frames))

	for fi, frame := range pd.frames {
		out := make([]byte, len(frame)/bytesPerSample)

		for idx, off := 0, 0; off+bytesPerSample <= len(frame); off, idx = off+bytesPerSample, idx+1 {
			var val int32
			if bytesPerSample == 1 {
				if pd.Info.PixelRepresentation == SignedPixels {
					val = int32(int8(frame[off]))
				} else {
					val = int32(frame[off])
				}
			} else {
				if pd.Info.PixelRepresentation == SignedPixels {
					val = int32(int16(binary.LittleEndian.Uint16(frame[off:])))
				} else {
					val = int32(binary.LittleEndian.Uint16(frame[off:]))
				}
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
