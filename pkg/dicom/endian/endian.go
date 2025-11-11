// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package endian provides DICOM endianness representation and utilities.
package endian

import "encoding/binary"

// Endian represents the byte order (endianness) of DICOM data.
type Endian int

const (
	// Little represents little-endian byte order.
	// This is the most common byte order in DICOM.
	Little Endian = iota

	// Big represents big-endian byte order.
	// This is less common and mostly used in retired transfer syntaxes.
	Big
)

var (
	// LocalMachine represents the endianness of the local machine.
	LocalMachine = func() Endian {
		buf := [2]byte{0x01, 0x00}
		if binary.LittleEndian.Uint16(buf[:]) == 1 {
			return Little
		}
		return Big
	}()

	// Network represents network byte order (big-endian).
	Network = Big
)

// String returns the string representation of the endianness.
func (e Endian) String() string {
	switch e {
	case Little:
		return "Little Endian"
	case Big:
		return "Big Endian"
	default:
		return "Unknown Endian"
	}
}

// IsBig returns true if the endianness is big-endian.
func (e Endian) IsBig() bool {
	return e == Big
}

// IsLittle returns true if the endianness is little-endian.
func (e Endian) IsLittle() bool {
	return e == Little
}

// ByteOrder returns the binary.ByteOrder corresponding to this endianness.
func (e Endian) ByteOrder() binary.ByteOrder {
	if e == Big {
		return binary.BigEndian
	}
	return binary.LittleEndian
}

// SwapUint16 swaps the byte order of a uint16 value.
func SwapUint16(value uint16) uint16 {
	return (value >> 8) | (value << 8)
}

// SwapUint32 swaps the byte order of a uint32 value.
func SwapUint32(value uint32) uint32 {
	return ((value & 0x000000ff) << 24) |
		((value & 0x0000ff00) << 8) |
		((value & 0x00ff0000) >> 8) |
		((value & 0xff000000) >> 24)
}

// SwapUint64 swaps the byte order of a uint64 value.
func SwapUint64(value uint64) uint64 {
	return ((value & 0x00000000000000ff) << 56) |
		((value & 0x000000000000ff00) << 40) |
		((value & 0x0000000000ff0000) << 24) |
		((value & 0x00000000ff000000) << 8) |
		((value & 0x000000ff00000000) >> 8) |
		((value & 0x0000ff0000000000) >> 24) |
		((value & 0x00ff000000000000) >> 40) |
		((value & 0xff00000000000000) >> 56)
}

// SwapInt16 swaps the byte order of an int16 value.
func SwapInt16(value int16) int16 {
	return int16(SwapUint16(uint16(value))) // #nosec G115 -- safe bit pattern conversion
}

// SwapInt32 swaps the byte order of an int32 value.
func SwapInt32(value int32) int32 {
	return int32(SwapUint32(uint32(value))) // #nosec G115 -- safe bit pattern conversion
}

// SwapInt64 swaps the byte order of an int64 value.
func SwapInt64(value int64) int64 {
	return int64(SwapUint64(uint64(value))) // #nosec G115 -- safe bit pattern conversion
}

// SwapBytes swaps bytes in a byte slice in groups of bytesToSwap bytes.
// For example, SwapBytes(2, data) swaps every pair of bytes.
func SwapBytes(bytesToSwap int, data []byte) {
	SwapBytesN(bytesToSwap, data, len(data))
}

// SwapBytesN swaps bytes in a byte slice in groups of bytesToSwap bytes,
// processing up to count bytes.
func SwapBytesN(bytesToSwap int, data []byte, count int) {
	if count > len(data) {
		panic("count exceeds data length")
	}

	if bytesToSwap == 1 || count == 0 {
		return
	}

	if bytesToSwap == 2 {
		swapBytes2(data, count)
		return
	}

	if bytesToSwap == 4 {
		swapBytes4(data, count)
		return
	}

	// General case for other byte sizes
	limit := count - (count % bytesToSwap)
	for i := 0; i < limit; i += bytesToSwap {
		// Reverse bytes in this group
		for j := 0; j < bytesToSwap/2; j++ {
			data[i+j], data[i+bytesToSwap-1-j] = data[i+bytesToSwap-1-j], data[i+j]
		}
	}
}

// swapBytes2 swaps bytes in pairs (optimized for 2-byte swaps).
func swapBytes2(data []byte, count int) {
	limit := count - (count % 2)
	for i := 0; i < limit; i += 2 {
		data[i], data[i+1] = data[i+1], data[i]
	}
}

// swapBytes4 swaps bytes in groups of 4 (optimized for 4-byte swaps).
func swapBytes4(data []byte, count int) {
	limit := count - (count % 4)
	for i := 0; i < limit; i += 4 {
		data[i], data[i+3] = data[i+3], data[i]
		data[i+1], data[i+2] = data[i+2], data[i+1]
	}
}
