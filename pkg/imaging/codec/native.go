// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

// NativeCodec handles uncompressed pixel data with various transfer syntaxes.
// This codec supports:
// - Implicit VR Little Endian
// - Explicit VR Little Endian
// - Explicit VR Big Endian
type NativeCodec struct {
	transferSyntax *transfer.TransferSyntax
	isBigEndian    bool
}

// NewNativeCodec creates a new Native (uncompressed) codec.
func NewNativeCodec(ts *transfer.TransferSyntax, isBigEndian bool) *NativeCodec {
	return &NativeCodec{
		transferSyntax: ts,
		isBigEndian:    isBigEndian,
	}
}

// NewImplicitVRLittleEndianCodec creates a codec for Implicit VR Little Endian.
func NewImplicitVRLittleEndianCodec() *NativeCodec {
	return &NativeCodec{
		transferSyntax: transfer.ImplicitVRLittleEndian,
		isBigEndian:    false,
	}
}

// NewExplicitVRLittleEndianCodec creates a codec for Explicit VR Little Endian.
func NewExplicitVRLittleEndianCodec() *NativeCodec {
	return &NativeCodec{
		transferSyntax: transfer.ExplicitVRLittleEndian,
		isBigEndian:    false,
	}
}

// NewExplicitVRBigEndianCodec creates a codec for Explicit VR Big Endian.
func NewExplicitVRBigEndianCodec() *NativeCodec {
	return &NativeCodec{
		transferSyntax: transfer.ExplicitVRBigEndian,
		isBigEndian:    true,
	}
}

// Name returns the codec name.
func (c *NativeCodec) Name() string {
	if c.isBigEndian {
		return "Native Big Endian"
	}
	return "Native Little Endian"
}

// TransferSyntax returns the transfer syntax this codec handles.
func (c *NativeCodec) TransferSyntax() *transfer.TransferSyntax {
	return c.transferSyntax
}

// Encode copies and optionally swaps byte order of pixel data.
// For uncompressed data, encoding is primarily a byte-order conversion.
func (c *NativeCodec) Encode(src *PixelData, dst *PixelData, params Parameters) error {
	if src == nil || dst == nil {
		return fmt.Errorf("source and destination pixel data must not be nil")
	}

	// For uncompressed data, we need to handle byte swapping if necessary
	bytesPerSample := int(src.BytesAllocated())

	// If single-byte samples, no swapping needed
	if bytesPerSample == 1 {
		dst.Data = make([]byte, len(src.Data))
		copy(dst.Data, src.Data)
		return nil
	}

	// Multi-byte samples may need byte swapping
	shouldSwap := false
	if params != nil {
		if swapParam := params.GetParameter("swap_bytes"); swapParam != nil {
			if swap, ok := swapParam.(bool); ok {
				shouldSwap = swap
			}
		}
	}

	dst.Data = make([]byte, len(src.Data))

	if !shouldSwap {
		// No swapping, just copy
		copy(dst.Data, src.Data)
		return nil
	}

	// Swap bytes based on sample size
	return c.swapBytes(src.Data, dst.Data, bytesPerSample)
}

// Decode copies and optionally swaps byte order of pixel data.
// For uncompressed data, decoding is primarily a byte-order conversion.
func (c *NativeCodec) Decode(src *PixelData, dst *PixelData, params Parameters) error {
	if src == nil || dst == nil {
		return fmt.Errorf("source and destination pixel data must not be nil")
	}

	// For uncompressed data, we need to handle byte swapping if necessary
	bytesPerSample := int(src.BytesAllocated())

	// If single-byte samples, no swapping needed
	if bytesPerSample == 1 {
		dst.Data = make([]byte, len(src.Data))
		copy(dst.Data, src.Data)
		return nil
	}

	// Multi-byte samples may need byte swapping
	shouldSwap := c.isBigEndian // Swap if source is big endian (convert to little endian)
	if params != nil {
		if swapParam := params.GetParameter("swap_bytes"); swapParam != nil {
			if swap, ok := swapParam.(bool); ok {
				shouldSwap = swap
			}
		}
	}

	dst.Data = make([]byte, len(src.Data))

	if !shouldSwap {
		// No swapping, just copy
		copy(dst.Data, src.Data)
		return nil
	}

	// Swap bytes based on sample size
	return c.swapBytes(src.Data, dst.Data, bytesPerSample)
}

// swapBytes swaps the byte order of multi-byte samples.
func (c *NativeCodec) swapBytes(src, dst []byte, bytesPerSample int) error {
	if len(src) != len(dst) {
		return fmt.Errorf("source and destination buffers must be same length")
	}

	if len(src)%bytesPerSample != 0 {
		return fmt.Errorf("data length %d not divisible by bytes per sample %d",
			len(src), bytesPerSample)
	}

	switch bytesPerSample {
	case 2:
		return c.swap16(src, dst)
	case 4:
		return c.swap32(src, dst)
	case 8:
		return c.swap64(src, dst)
	default:
		return fmt.Errorf("unsupported bytes per sample: %d", bytesPerSample)
	}
}

// swap16 swaps 16-bit values between little and big endian.
func (c *NativeCodec) swap16(src, dst []byte) error {
	for i := 0; i < len(src); i += 2 {
		if i+1 >= len(src) {
			return fmt.Errorf("incomplete 16-bit sample at offset %d", i)
		}
		dst[i] = src[i+1]
		dst[i+1] = src[i]
	}
	return nil
}

// swap32 swaps 32-bit values between little and big endian.
func (c *NativeCodec) swap32(src, dst []byte) error {
	for i := 0; i < len(src); i += 4 {
		if i+3 >= len(src) {
			return fmt.Errorf("incomplete 32-bit sample at offset %d", i)
		}
		dst[i] = src[i+3]
		dst[i+1] = src[i+2]
		dst[i+2] = src[i+1]
		dst[i+3] = src[i]
	}
	return nil
}

// swap64 swaps 64-bit values between little and big endian.
func (c *NativeCodec) swap64(src, dst []byte) error {
	for i := 0; i < len(src); i += 8 {
		if i+7 >= len(src) {
			return fmt.Errorf("incomplete 64-bit sample at offset %d", i)
	}
		dst[i] = src[i+7]
		dst[i+1] = src[i+6]
		dst[i+2] = src[i+5]
		dst[i+3] = src[i+4]
		dst[i+4] = src[i+3]
		dst[i+5] = src[i+2]
		dst[i+6] = src[i+1]
		dst[i+7] = src[i]
	}
	return nil
}

// ConvertEndianness converts pixel data between little and big endian.
// This is a utility function for external use.
func ConvertEndianness(data []byte, bytesPerSample int) ([]byte, error) {
	result := make([]byte, len(data))
	codec := NewNativeCodec(nil, false)
	err := codec.swapBytes(data, result, bytesPerSample)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadUint16 reads a uint16 value from the buffer at the specified offset,
// respecting the codec's endianness.
func (c *NativeCodec) ReadUint16(data []byte, offset int) (uint16, error) {
	if offset+2 > len(data) {
		return 0, fmt.Errorf("offset %d out of range for buffer length %d", offset, len(data))
	}

	if c.isBigEndian {
		return binary.BigEndian.Uint16(data[offset:]), nil
	}
	return binary.LittleEndian.Uint16(data[offset:]), nil
}

// ReadUint32 reads a uint32 value from the buffer at the specified offset,
// respecting the codec's endianness.
func (c *NativeCodec) ReadUint32(data []byte, offset int) (uint32, error) {
	if offset+4 > len(data) {
		return 0, fmt.Errorf("offset %d out of range for buffer length %d", offset, len(data))
	}

	if c.isBigEndian {
		return binary.BigEndian.Uint32(data[offset:]), nil
	}
	return binary.LittleEndian.Uint32(data[offset:]), nil
}

// WriteUint16 writes a uint16 value to the buffer at the specified offset,
// respecting the codec's endianness.
func (c *NativeCodec) WriteUint16(data []byte, offset int, value uint16) error {
	if offset+2 > len(data) {
		return fmt.Errorf("offset %d out of range for buffer length %d", offset, len(data))
	}

	if c.isBigEndian {
		binary.BigEndian.PutUint16(data[offset:], value)
	} else {
		binary.LittleEndian.PutUint16(data[offset:], value)
	}
	return nil
}

// WriteUint32 writes a uint32 value to the buffer at the specified offset,
// respecting the codec's endianness.
func (c *NativeCodec) WriteUint32(data []byte, offset int, value uint32) error {
	if offset+4 > len(data) {
		return fmt.Errorf("offset %d out of range for buffer length %d", offset, len(data))
	}

	if c.isBigEndian {
		binary.BigEndian.PutUint32(data[offset:], value)
	} else {
		binary.LittleEndian.PutUint32(data[offset:], value)
	}
	return nil
}
