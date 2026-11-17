// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package pixeldata

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// applyWindowTo8bit maps pixel data to 8-bit frames using the default LINEAR VOI window.
func applyWindowTo8bit(pd *Data, center, width float64, ignorePadding bool) ([][]byte, error) {
	return applyWindowTo8bitWithModality(pd, center, width, ignorePadding, nil, lut.VOILUTFunctionLinear)
}

func applyWindowTo8bitWithModality(pd *Data, center, width float64, ignorePadding bool, modality lut.LUT, function lut.VOILUTFunction) ([][]byte, error) {
	if pd.Info == nil {
		return nil, fmt.Errorf("pixel data info is nil")
	}
	if pd.Info.Encapsulated {
		return nil, fmt.Errorf("cannot apply window on encapsulated data; decode first")
	}

	// Build the VOI LUT using shared lut package to avoid duplicated window math.
	// Function-specific width rules (LINEAR >= 1, LINEAR_EXACT/SIGMOID > 0) are
	// validated here so dataset-driven callers cannot silently fall back.
	if function == "" {
		function = lut.VOILUTFunctionLinear
	}
	voiLUT, err := lut.CreateValidatedVOILUT(function, center, width)
	if err != nil {
		return nil, err
	}

	bytesPerSample := pd.Info.BytesAllocated()
	if bytesPerSample != 1 && bytesPerSample != 2 && bytesPerSample != 4 {
		return nil, fmt.Errorf("unsupported BytesAllocated=%d for windowing", bytesPerSample)
	}

	// Build the mapping chain once: the modality LUT (if any) shifts the stored
	// values, then the VOI LUT maps them to 0..255.
	transform := lut.LUT(voiLUT)
	if modality != nil {
		composite := lut.NewCompositeLUT()
		composite.Add(modality)
		composite.Add(voiLUT)
		transform = composite
	} else if minimum, maximum, ok := storedValueTableRange(pd.Info); ok {
		// Precalculate over the stored-value range for faster mapping, but only
		// when the range fits a table: clamping a larger range would mis-map the
		// upper part (e.g. BitsStored > 16).
		transform = lut.NewPrecalculatedLUT(voiLUT, minimum, maximum)
	}

	return mapSamplesTo8bit(pd, transform, ignorePadding), nil
}

// storedValueTableRange returns the stored-value domain and reports whether it
// is small enough to precalculate a lookup table. A modality LUT changes the
// domain, so callers only use it when no modality LUT is applied.
func storedValueTableRange(info *Info) (minimum, maximum int, ok bool) {
	if info.BitsStored == 0 || info.BitsStored > 16 {
		return 0, 0, false
	}
	if info.PixelRepresentation == pixel.SignedPixels {
		maximum = (1 << (info.BitsStored - 1)) - 1
		minimum = -maximum - 1
	} else {
		minimum = 0
		maximum = (1 << info.BitsStored) - 1
	}
	return minimum, maximum, maximum-minimum <= 65536
}

// mapSamplesTo8bit maps every sample of every frame to 8 bits through transform.
// Windowing and LUT-sequence mapping share this loop; the caller has already
// validated the pixel data and assembled the transform chain, including any
// output scaling.
func mapSamplesTo8bit(pd *Data, transform lut.LUT, ignorePadding bool) [][]byte {
	bytesPerSample := pd.Info.BytesAllocated()

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

			out[idx] = clampByte(int(transform.Transform(float64(val)) + 0.5))
		}

		result[fi] = out
	}

	return result
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
			off := sampleOffset(sampleIndex, pd.Info)
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
