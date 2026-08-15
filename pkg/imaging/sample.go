// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import "encoding/binary"

// decodePixelSampleLE decodes a single pixel sample from data at the given byte offset.
// BitsAllocated values of 8, 16, and 32 are supported.
func decodePixelSampleLE(data []byte, offset int, info *PixelDataInfo) (int64, bool) {
	if info == nil {
		return 0, false
	}
	bytesPerSample := info.BytesAllocated()
	if offset < 0 || offset+bytesPerSample > len(data) {
		return 0, false
	}

	var raw uint32
	switch bytesPerSample {
	case 1:
		raw = uint32(data[offset])
	case 2:
		raw = uint32(binary.LittleEndian.Uint16(data[offset:]))
	case 4:
		raw = binary.LittleEndian.Uint32(data[offset:])
	default:
		return 0, false
	}

	bitsStored := info.BitsStored
	if bitsStored == 0 || bitsStored > info.BitsAllocated {
		bitsStored = info.BitsAllocated
	}
	highBit := info.HighBit
	if highBit+1 < bitsStored {
		highBit = bitsStored - 1
	}
	shift := int(highBit + 1 - bitsStored)
	mask := uint32(1<<bitsStored) - 1
	sample := (raw >> shift) & mask

	if info.PixelRepresentation != SignedPixels {
		return int64(sample), true
	}

	signBit := uint32(1 << (bitsStored - 1))
	if sample&signBit == 0 {
		return int64(sample), true
	}
	return int64(sample) - int64(1<<bitsStored), true
}

func swapPixelDataBytes(data []byte, info *PixelDataInfo) []byte {
	if info == nil {
		return data
	}
	bytesPerSample := info.BytesAllocated()
	if bytesPerSample != 2 && bytesPerSample != 4 {
		return data
	}
	out := make([]byte, len(data))
	copy(out, data)
	for offset := 0; offset+bytesPerSample <= len(out); offset += bytesPerSample {
		for left, right := offset, offset+bytesPerSample-1; left < right; left, right = left+1, right-1 {
			out[left], out[right] = out[right], out[left]
		}
	}
	return out
}
