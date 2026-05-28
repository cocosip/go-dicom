// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging/lut"
)

// applyWindowTo8bit maps pixel data to 8-bit frames using a VOI window.
func applyWindowTo8bit(pd *DicomPixelData, center, width float64, ignorePadding bool) ([][]byte, error) {
	if pd.Info == nil {
		return nil, fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return nil, fmt.Errorf("cannot apply window on encapsulated data; decode first")
	}
	if width < 1 {
		width = 1
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return nil, fmt.Errorf("unsupported BytesAllocated=%d for windowing", bytesPerSample)
	}

	// Build VOI LUT using shared lut package to avoid duplicated window math.
	voiLUT := lut.CreateVOILUT(lut.VOILUTFunctionLinear, center, width)

	// Precalculate within valid input range for faster mapping.
	// Cap at 16-bit range (65536 entries) to prevent excessive memory allocation
	// when BitsStored is very large (e.g., 32-bit).
	var minInput, maxInput int
	if pd.Info.PixelRepresentation == SignedPixels {
		maxInput = (1 << (pd.Info.BitsStored - 1)) - 1
		minInput = -maxInput - 1
	} else {
		minInput = 0
		maxInput = (1 << pd.Info.BitsStored) - 1
	}
	if maxInput-minInput > 65536 {
		maxInput = minInput + 65536
	}
	precalc := lut.NewPrecalculatedLUT(voiLUT, minInput, maxInput)

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
			val, ok := decodePixelSampleLE(frame, off, pd.Info)
			if !ok {
				continue
			}

			if ignorePadding && hasPadding && val >= padMin && val <= padMax {
				out[idx] = 0
				continue
			}

			mapped := precalc.Transform(float64(val))
			out[idx] = clampByte(int(mapped + 0.5))
		}

		result[fi] = out
	}

	return result, nil
}

// minMaxSamples returns min/max sample values across frames, optionally ignoring padding.
func minMaxSamples(pd *DicomPixelData, ignorePadding bool) (float64, float64, error) {
	if pd.Info == nil {
		return 0, 0, fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return 0, 0, fmt.Errorf("min/max on encapsulated data requires decode first")
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 {
		return 0, 0, fmt.Errorf("unsupported BytesAllocated=%d for min/max", bytesPerSample)
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

	var (
		found bool
		vmin  float64
		vmax  float64
	)

	for _, frame := range pd.frames {
		for off := 0; off+bytesPerSample <= len(frame); off += bytesPerSample {
			val, ok := decodePixelSampleLE(frame, off, pd.Info)
			if !ok {
				continue
			}

			if ignorePadding && hasPadding && val >= padMin && val <= padMax {
				continue
			}

			if !found {
				vmin, vmax = float64(val), float64(val)
				found = true
			} else {
				if float64(val) < vmin {
					vmin = float64(val)
				}
				if float64(val) > vmax {
					vmax = float64(val)
				}
			}
		}
	}

	if !found {
		return 0, 0, fmt.Errorf("no samples found (possibly all padding)")
	}

	return vmin, vmax, nil
}
