// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import "encoding/binary"

func decodePixelSampleLE(data []byte, offset int, info *PixelDataInfo) (int32, bool) {
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
		return int32(sample), true
	}

	signBit := uint32(1 << (bitsStored - 1))
	if sample&signBit == 0 {
		return int32(sample), true
	}
	return int32(sample) - int32(1<<bitsStored), true
}

func normalizeNativePixelDataToLittleEndian(data []byte, info *PixelDataInfo) []byte {
	if info == nil || info.BytesAllocated() != 2 {
		return data
	}
	out := make([]byte, len(data))
	copy(out, data)
	for i := 0; i+1 < len(out); i += 2 {
		out[i], out[i+1] = out[i+1], out[i]
	}
	return out
}
