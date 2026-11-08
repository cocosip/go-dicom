// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package pixeldata

import (
	"fmt"
	"math"

	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// applyWindowTo8bit maps pixel data to 8-bit frames using a VOI window.
func applyWindowTo8bit(pd *Data, center, width float64, ignorePadding bool) ([][]byte, error) {
	return applyWindowTo8bitWithModality(pd, center, width, ignorePadding, nil)
}

func applyWindowTo8bitWithModality(pd *Data, center, width float64, ignorePadding bool, modality lut.LUT) ([][]byte, error) {
	if pd.Info == nil {
		return nil, fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return nil, fmt.Errorf("cannot apply window on encapsulated data; decode first")
	}
	if width < 1 || math.IsNaN(width) || math.IsInf(width, 0) {
		return nil, fmt.Errorf("window width must be finite and at least 1, got %v", width)
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 && bytesPerSample != 4 {
		return nil, fmt.Errorf("unsupported BytesAllocated=%d for windowing", bytesPerSample)
	}

	// Build VOI LUT using shared lut package to avoid duplicated window math.
	voiLUT := lut.CreateVOILUT(lut.VOILUTFunctionLinear, center, width)

	// Precalculate within the valid input range for faster mapping, but only when
	// the range fits a table: clamping a larger range would mis-map the upper part
	// (e.g. BitsStored > 16), so fall back to direct transform in that case.
	var minInput, maxInput int
	if pd.Info.PixelRepresentation == pixel.SignedPixels {
		maxInput = (1 << (pd.Info.BitsStored - 1)) - 1
		minInput = -maxInput - 1
	} else {
		minInput = 0
		maxInput = (1 << pd.Info.BitsStored) - 1
	}
	usePrecalc := maxInput-minInput <= 65536
	var precalc *lut.PrecalculatedLUT
	if usePrecalc {
		precalc = lut.NewPrecalculatedLUT(voiLUT, minInput, maxInput)
	}

	hasPadding := pd.Info.PixelPaddingValue != nil
	padMin, padMax, _ := pixelPaddingRange(pd.Info)

	result := make([][]byte, len(pd.frames))

	for fi, frame := range pd.frames {
		out := make([]byte, frameSampleCount(frame, pd.Info))

		for idx := 0; idx < len(out); idx++ {
			off := sampleOffset(idx, pd.Info)
			if pd.Info.BitsAllocated != 1 && off+bytesPerSample > len(frame) {
				break
			}
			val, ok := decodePixelSampleLE(frame, off, pd.Info)
			if !ok {
				continue
			}

			if ignorePadding && hasPadding && val >= padMin && val <= padMax {
				out[idx] = 0
				continue
			}

			modalityValue := float64(val)
			if modality != nil {
				modalityValue = modality.Transform(modalityValue)
			}
			mapped := float64(0)
			if modality != nil || !usePrecalc {
				// The precomputed table is indexed in stored-value space; a rescale
				// can shift or reverse that domain, so transform directly. Likewise,
				// skip the table when the stored range is too large to precalculate.
				mapped = voiLUT.Transform(modalityValue)
			} else {
				mapped = precalc.Transform(modalityValue)
			}
			out[idx] = clampByte(int(mapped + 0.5))
		}

		result[fi] = out
	}

	return result, nil
}

// minMaxSamples returns min/max sample values across frames, optionally ignoring padding.
func minMaxSamples(pd *Data, ignorePadding bool) (float64, float64, error) {
	if pd.Info == nil {
		return 0, 0, fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return 0, 0, fmt.Errorf("min/max on encapsulated data requires decode first")
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 && bytesPerSample != 4 {
		return 0, 0, fmt.Errorf("unsupported BytesAllocated=%d for min/max", bytesPerSample)
	}

	hasPadding := pd.Info.PixelPaddingValue != nil
	padMin, padMax, _ := pixelPaddingRange(pd.Info)

	var (
		found bool
		vmin  float64
		vmax  float64
	)

	for _, frame := range pd.frames {
		sampleCount := len(frame) / bytesPerSample
		if pd.Info.BitsAllocated == 1 {
			sampleCount = int(pd.Info.Width) * int(pd.Info.Height) * int(pd.Info.SamplesPerPixel)
		}
		for sampleIndex := 0; sampleIndex < sampleCount; sampleIndex++ {
			off := sampleIndex * bytesPerSample
			if pd.Info.BitsAllocated == 1 {
				off = sampleIndex
			}
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
