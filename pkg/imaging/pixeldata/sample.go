// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"context"
	"encoding/binary"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/io/endian"
)

const contextCopyChunkSize = 64 * 1024

func copyWithContext(ctx context.Context, destination, source []byte) error {
	for offset := 0; offset < len(source); offset += contextCopyChunkSize {
		if err := ctx.Err(); err != nil {
			return err
		}
		end := min(offset+contextCopyChunkSize, len(source))
		copy(destination[offset:end], source[offset:end])
	}
	return ctx.Err()
}

// decodePixelSampleLE decodes a single pixel sample from data at the given byte offset.
// BitsAllocated values of 8, 16, and 32 are supported.
func decodePixelSampleLE(data []byte, offset int, info *Info) (int64, bool) {
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

	if info.PixelRepresentation != pixel.SignedPixels {
		return int64(sample), true
	}

	signBit := uint32(1 << (bitsStored - 1))
	if sample&signBit == 0 {
		return int64(sample), true
	}
	return int64(sample) - int64(1<<bitsStored), true
}

func swapPixelDataBytes(data []byte, info *Info) []byte {
	if info == nil {
		return data
	}
	if info.BytesAllocated() == 1 {
		return data
	}
	out := make([]byte, len(data))
	copy(out, data)
	// DICOM OW values are sequences of 16-bit words, regardless of the native
	// sample width represented by those words.
	for offset := 0; offset+1 < len(out); offset += 2 {
		out[offset], out[offset+1] = out[offset+1], out[offset]
	}
	return out
}

// pixelDataNeedsByteSwap reports whether native Pixel Data is encoded in a
// byte order different from the library's internal little-endian form.
// Some transfer syntaxes keep dataset encoding little-endian while declaring
// big-endian pixel words (for example GE's private syntax), so Endian alone is
// not sufficient here.
func pixelDataNeedsByteSwap(info *Info) bool {
	if info == nil || info.BytesAllocated() == 1 {
		return false
	}
	syntax, err := transfer.Parse(info.TransferSyntaxUID)
	if err == nil {
		return syntax.Endian() == endian.Big || syntax.SwapPixelData()
	}
	return info.TransferSyntaxUID == transfer.ExplicitVRBigEndian.UID().UID() ||
		info.TransferSyntaxUID == transfer.ImplicitVRBigEndian.UID().UID() ||
		info.TransferSyntaxUID == transfer.GEPrivateImplicitVRBigEndian.UID().UID()
}

func clampByte(value int) byte {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return byte(value)
}
