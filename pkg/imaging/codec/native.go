// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

const contextCopyChunkSize = 64 * 1024

// NativeCodec handles uncompressed pixel data with various transfer syntaxes.
// This codec supports:
// - Implicit VR Little Endian
// - Explicit VR Little Endian
// - Explicit VR Big Endian
type NativeCodec struct {
	transferSyntax *transfer.Syntax
	isBigEndian    bool
}

// NewNativeCodec creates a new Native (uncompressed) codec.
func NewNativeCodec(ts *transfer.Syntax, isBigEndian bool) *NativeCodec {
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
func (c *NativeCodec) TransferSyntax() *transfer.Syntax {
	return c.transferSyntax
}

// DefaultParameters returns default parameters for this codec.
func (c *NativeCodec) DefaultParameters() Parameters {
	return NativeParameters{}
}

// Encode encodes pixel data from oldPixelData to newPixelData.
// For native (uncompressed) codec, this is essentially a copy operation with potential byte swapping.
func (c *NativeCodec) Encode(ctx context.Context, oldPixelData FrameSource, newPixelData FrameSink, parameters Parameters) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if oldPixelData == nil || newPixelData == nil {
		return fmt.Errorf("source and destination pixel data must not be nil")
	}
	params, err := nativeParameters(parameters)
	if err != nil {
		return err
	}

	frameInfo := oldPixelData.FrameInfo()
	frameCount := oldPixelData.FrameCount()

	// Process each frame
	for i := 0; i < frameCount; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		srcFrame, err := oldPixelData.Frame(ctx, i)
		if err != nil {
			return fmt.Errorf("failed to get frame %d: %w", i, err)
		}

		var dstFrame []byte
		if err := c.encodeFrame(ctx, srcFrame, &dstFrame, frameInfo, params); err != nil {
			return fmt.Errorf("failed to encode frame %d: %w", i, err)
		}

		if err := newPixelData.AddFrame(ctx, dstFrame); err != nil {
			return fmt.Errorf("failed to add frame %d: %w", i, err)
		}
	}

	return nil
}

// Decode decodes pixel data from oldPixelData to newPixelData.
// For native (uncompressed) codec, this is essentially a copy operation with potential byte swapping.
func (c *NativeCodec) Decode(ctx context.Context, oldPixelData FrameSource, newPixelData FrameSink, parameters Parameters) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if oldPixelData == nil || newPixelData == nil {
		return fmt.Errorf("source and destination pixel data must not be nil")
	}
	params, err := nativeParameters(parameters)
	if err != nil {
		return err
	}

	frameInfo := oldPixelData.FrameInfo()
	frameCount := oldPixelData.FrameCount()

	// Process each frame
	for i := 0; i < frameCount; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		srcFrame, err := oldPixelData.Frame(ctx, i)
		if err != nil {
			return fmt.Errorf("failed to get frame %d: %w", i, err)
		}

		var dstFrame []byte
		if err := c.decodeFrame(ctx, srcFrame, &dstFrame, frameInfo, params); err != nil {
			return fmt.Errorf("failed to decode frame %d: %w", i, err)
		}

		if err := newPixelData.AddFrame(ctx, dstFrame); err != nil {
			return fmt.Errorf("failed to add frame %d: %w", i, err)
		}
	}

	return nil
}

// encodeFrame encodes a single frame (internal helper method).
func (c *NativeCodec) encodeFrame(ctx context.Context, src []byte, dst *[]byte, info FrameInfo, params NativeParameters) error {
	if len(src) == 0 {
		return fmt.Errorf("source frame data must not be empty")
	}
	// Bits Allocated selects whether native Pixel Data is byte- or word-based.
	bytesAllocated := info.BitDepth.BytesAllocated()

	// If single-byte samples, no swapping needed
	if bytesAllocated == 1 {
		return copyFrame(ctx, src, dst)
	}

	// Multi-byte native Pixel Data uses OW and may need word byte swapping.
	shouldSwap := c.isBigEndian
	switch params.ByteSwap {
	case ByteSwapDisabled:
		shouldSwap = false
	case ByteSwapEnabled:
		shouldSwap = true
	}

	*dst = make([]byte, len(src))

	if !shouldSwap {
		// No swapping, just copy
		return copyFrameInto(ctx, src, *dst)
	}

	return c.swapOWWords(ctx, src, *dst)
}

// decodeFrame decodes a single frame (internal helper method).
func (c *NativeCodec) decodeFrame(ctx context.Context, src []byte, dst *[]byte, info FrameInfo, params NativeParameters) error {
	if len(src) == 0 {
		return fmt.Errorf("source frame data must not be empty")
	}
	// Bits Allocated selects whether native Pixel Data is byte- or word-based.
	bytesAllocated := info.BitDepth.BytesAllocated()

	// If single-byte samples, no swapping needed
	if bytesAllocated == 1 {
		return copyFrame(ctx, src, dst)
	}

	// Multi-byte native Pixel Data uses OW and may need word byte swapping.
	shouldSwap := c.isBigEndian // Swap if source is big endian (convert to little endian)
	switch params.ByteSwap {
	case ByteSwapDisabled:
		shouldSwap = false
	case ByteSwapEnabled:
		shouldSwap = true
	}

	*dst = make([]byte, len(src))

	if !shouldSwap {
		// No swapping, just copy
		return copyFrameInto(ctx, src, *dst)
	}

	return c.swapOWWords(ctx, src, *dst)
}

func copyFrame(ctx context.Context, src []byte, dst *[]byte) error {
	*dst = make([]byte, len(src))
	return copyFrameInto(ctx, src, *dst)
}

func copyFrameInto(ctx context.Context, src, dst []byte) error {
	for offset := 0; offset < len(src); offset += contextCopyChunkSize {
		if err := ctx.Err(); err != nil {
			return err
		}
		end := min(offset+contextCopyChunkSize, len(src))
		copy(dst[offset:end], src[offset:end])
	}
	return ctx.Err()
}

// swapOWWords swaps the bytes within each 16-bit OW word. A pixel sample may
// occupy multiple words; the order of those words is not reversed.
func (c *NativeCodec) swapOWWords(ctx context.Context, src, dst []byte) error {
	if len(src) != len(dst) {
		return fmt.Errorf("source and destination buffers must be same length")
	}

	if len(src)%2 != 0 {
		return fmt.Errorf("data length %d is not aligned to 16-bit OW words", len(src))
	}
	return c.swap16(ctx, src, dst)
}

// swap16 swaps bytes within each 16-bit OW word.
func (c *NativeCodec) swap16(ctx context.Context, src, dst []byte) error {
	for i := 0; i < len(src); i += 2 {
		if i%contextCopyChunkSize == 0 {
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if i+1 >= len(src) {
			return fmt.Errorf("incomplete 16-bit sample at offset %d", i)
		}
		dst[i] = src[i+1]
		dst[i+1] = src[i]
	}
	return nil
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
