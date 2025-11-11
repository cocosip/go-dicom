// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
)

// RLECodec implements DICOM RLE (Run-Length Encoding) compression and decompression.
// This codec conforms to DICOM Part 5, Annex G: RLE Compression.
type RLECodec struct{}

// NewRLECodec creates a new RLE codec instance.
func NewRLECodec() *RLECodec {
	return &RLECodec{}
}

// Name returns the codec name.
func (c *RLECodec) Name() string {
	return "RLE Lossless"
}

// TransferSyntax returns the RLE Lossless transfer syntax.
func (c *RLECodec) TransferSyntax() *transfer.Syntax {
	return transfer.RLELossless
}

// Encode compresses pixel data using RLE compression.
func (c *RLECodec) Encode(src *PixelData, dst *PixelData, _ Parameters) error {
	if src == nil || dst == nil {
		return fmt.Errorf("source and destination pixel data must not be nil")
	}

	pixelCount := int(src.Width) * int(src.Height)
	numberOfSegments := src.BytesAllocated() * int(src.SamplesPerPixel)
	frameSize := src.UncompressedFrameSize()

	var result bytes.Buffer

	for frame := 0; frame < src.NumberOfFrames; frame++ {
		frameStart := frame * frameSize
		frameEnd := frameStart + frameSize
		if frameEnd > len(src.Data) {
			return fmt.Errorf("frame %d exceeds data length", frame)
		}
		frameData := src.Data[frameStart:frameEnd]

		encoder := newRLEEncoder()

		for s := 0; s < numberOfSegments; s++ {
			encoder.NextSegment()

			sample := s / src.BytesAllocated()
			sabyte := s % src.BytesAllocated()

			var pos, offset int

			if src.IsInterleaved() {
				pos = sample * src.BytesAllocated()
				offset = numberOfSegments
			} else {
				pos = sample * src.BytesAllocated() * pixelCount
				offset = src.BytesAllocated()
			}

			pos += src.BytesAllocated() - sabyte - 1

			for p := 0; p < pixelCount; p++ {
				if pos >= len(frameData) {
					return fmt.Errorf("read position %d exceeds frame buffer length %d", pos, len(frameData))
				}
				encoder.Encode(frameData[pos])
				pos += offset
			}
			encoder.Flush()
		}

		encoder.MakeEvenLength()
		frameBytes := encoder.GetBuffer()
		result.Write(frameBytes)
	}

	dst.Data = result.Bytes()
	return nil
}

// Decode decompresses RLE-compressed pixel data.
func (c *RLECodec) Decode(src *PixelData, dst *PixelData, _ Parameters) error {
	if src == nil || dst == nil {
		return fmt.Errorf("source and destination pixel data must not be nil")
	}

	pixelCount := int(src.Width) * int(src.Height)
	numberOfSegments := src.BytesAllocated() * int(src.SamplesPerPixel)

	var result bytes.Buffer
	dataOffset := 0

	for frame := 0; frame < src.NumberOfFrames; frame++ {
		// Create new frame data of even length
		frameSize := dst.UncompressedFrameSize()
		if (frameSize & 1) == 1 {
			frameSize++
		}

		frameData := make([]byte, frameSize)

		// Decode RLE data for this frame
		if dataOffset >= len(src.Data) {
			return fmt.Errorf("insufficient data for frame %d", frame)
		}

		decoder, bytesRead, err := newRLEDecoder(src.Data[dataOffset:])
		if err != nil {
			return fmt.Errorf("failed to create RLE decoder for frame %d: %w", frame, err)
		}
		dataOffset += bytesRead

		if decoder.NumberOfSegments != numberOfSegments {
			return fmt.Errorf("unexpected number of RLE segments: got %d, expected %d",
				decoder.NumberOfSegments, numberOfSegments)
		}

		for s := 0; s < numberOfSegments; s++ {
			sample := s / dst.BytesAllocated()
			sabyte := s % dst.BytesAllocated()

			var pos, offset int

			if dst.IsInterleaved() {
				pos = sample * dst.BytesAllocated()
				offset = int(dst.SamplesPerPixel) * dst.BytesAllocated()
			} else {
				pos = sample * dst.BytesAllocated() * pixelCount
				offset = dst.BytesAllocated()
			}

			pos += dst.BytesAllocated() - sabyte - 1

			if err := decoder.DecodeSegment(s, frameData, pos, offset); err != nil {
				return fmt.Errorf("failed to decode segment %d: %w", s, err)
			}
		}

		result.Write(frameData)
	}

	dst.Data = result.Bytes()
	return nil
}

// rleEncoder implements the RLE encoding algorithm.
type rleEncoder struct {
	count      int
	offsets    [15]uint32
	buffer     bytes.Buffer
	tempBuffer [132]byte
	prevByte   int
	repeatCnt  int
	bufferPos  int
}

func newRLEEncoder() *rleEncoder {
	enc := &rleEncoder{
		prevByte: -1,
	}

	// Write header (will be rewritten later)
	_ = binary.Write(&enc.buffer, binary.LittleEndian, uint32(enc.count))
	for i := 0; i < 15; i++ {
		_ = binary.Write(&enc.buffer, binary.LittleEndian, enc.offsets[i])
	}

	return enc
}

func (e *rleEncoder) NextSegment() {
	e.Flush()
	if (e.buffer.Len() & 1) == 1 {
		e.buffer.WriteByte(0x00)
	}
	e.offsets[e.count] = uint32(e.buffer.Len())
	e.count++
}

func (e *rleEncoder) Encode(b byte) {
	if int(b) == e.prevByte {
		e.repeatCnt++

		if e.repeatCnt > 2 && e.bufferPos > 0 {
			// Starting a run, flush out the buffer
			for e.bufferPos > 0 {
				count := min(128, e.bufferPos)
				e.buffer.WriteByte(byte(count - 1))
				e.buffer.Write(e.tempBuffer[:count])
				// Shift buffer
				copy(e.tempBuffer[:], e.tempBuffer[count:e.bufferPos])
				e.bufferPos -= count
			}
		} else if e.repeatCnt > 128 {
			count := min(e.repeatCnt, 128)
			e.buffer.WriteByte(byte(257 - count))
			e.buffer.WriteByte(byte(e.prevByte))
			e.repeatCnt -= count
		}
	} else {
		switch e.repeatCnt {
		case 0:
			// Do nothing
		case 1:
			e.tempBuffer[e.bufferPos] = byte(e.prevByte)
			e.bufferPos++
		case 2:
			e.tempBuffer[e.bufferPos] = byte(e.prevByte)
			e.bufferPos++
			e.tempBuffer[e.bufferPos] = byte(e.prevByte)
			e.bufferPos++
		default:
			for e.repeatCnt > 0 {
				count := min(e.repeatCnt, 128)
				e.buffer.WriteByte(byte(257 - count))
				e.buffer.WriteByte(byte(e.prevByte))
				e.repeatCnt -= count
			}
		}

		for e.bufferPos > 128 {
			count := min(128, e.bufferPos)
			e.buffer.WriteByte(byte(count - 1))
			e.buffer.Write(e.tempBuffer[:count])
			// Shift buffer
			copy(e.tempBuffer[:], e.tempBuffer[count:e.bufferPos])
			e.bufferPos -= count
		}

		e.prevByte = int(b)
		e.repeatCnt = 1
	}
}

func (e *rleEncoder) Flush() {
	if e.repeatCnt < 2 {
		for e.repeatCnt > 0 {
			e.tempBuffer[e.bufferPos] = byte(e.prevByte)
			e.bufferPos++
			e.repeatCnt--
		}
	}

	for e.bufferPos > 0 {
		count := min(128, e.bufferPos)
		e.buffer.WriteByte(byte(count - 1))
		e.buffer.Write(e.tempBuffer[:count])
		// Shift buffer
		copy(e.tempBuffer[:], e.tempBuffer[count:e.bufferPos])
		e.bufferPos -= count
	}

	if e.repeatCnt >= 2 {
		for e.repeatCnt > 0 {
			count := min(e.repeatCnt, 128)
			e.buffer.WriteByte(byte(257 - count))
			e.buffer.WriteByte(byte(e.prevByte))
			e.repeatCnt -= count
		}
	}

	e.prevByte = -1
	e.repeatCnt = 0
	e.bufferPos = 0
}

func (e *rleEncoder) MakeEvenLength() {
	if (e.buffer.Len() & 1) == 1 {
		e.buffer.WriteByte(0x00)
	}
}

func (e *rleEncoder) GetBuffer() []byte {
	e.Flush()

	// Rewrite header
	result := e.buffer.Bytes()
	buf := bytes.NewBuffer(result[:0])
	_ = binary.Write(buf, binary.LittleEndian, uint32(e.count))
	for i := 0; i < 15; i++ {
		_ = binary.Write(buf, binary.LittleEndian, e.offsets[i])
	}

	return result
}

// rleDecoder implements the RLE decoding algorithm.
type rleDecoder struct {
	NumberOfSegments int
	offsets          [15]int
	data             []byte
}

func newRLEDecoder(data []byte) (*rleDecoder, int, error) {
	if len(data) < 64 {
		return nil, 0, fmt.Errorf("RLE data too short: need at least 64 bytes, got %d", len(data))
	}

	dec := &rleDecoder{
		data: data,
	}

	reader := bytes.NewReader(data)

	var numSegments uint32
	if err := binary.Read(reader, binary.LittleEndian, &numSegments); err != nil {
		return nil, 0, fmt.Errorf("failed to read number of segments: %w", err)
	}
	dec.NumberOfSegments = int(numSegments)

	for i := 0; i < 15; i++ {
		var offset uint32
		if err := binary.Read(reader, binary.LittleEndian, &offset); err != nil {
			return nil, 0, fmt.Errorf("failed to read offset %d: %w", i, err)
		}
		dec.offsets[i] = int(offset)
	}

	// Calculate total bytes consumed for this frame
	// RLE segments are variable length, conservatively use all remaining data
	bytesRead := len(data)

	return dec, bytesRead, nil
}

func (d *rleDecoder) DecodeSegment(segment int, buffer []byte, start int, sampleOffset int) error {
	if segment < 0 || segment >= d.NumberOfSegments {
		return fmt.Errorf("segment number %d out of range [0, %d)", segment, d.NumberOfSegments)
	}

	offset := d.getSegmentOffset(segment)
	length := d.getSegmentLength(segment)

	return d.decode(buffer, start, sampleOffset, d.data, offset, length)
}

func (d *rleDecoder) getSegmentOffset(segment int) int {
	return d.offsets[segment]
}

func (d *rleDecoder) getSegmentLength(segment int) int {
	offset := d.getSegmentOffset(segment)
	if segment < (d.NumberOfSegments - 1) {
		next := d.getSegmentOffset(segment + 1)
		return next - offset
	}
	return len(d.data) - offset
}

func (d *rleDecoder) decode(buffer []byte, start int, sampleOffset int, rleData []byte, offset int, count int) error {
	pos := start
	end := offset + count
	bufferLength := len(buffer)

	for i := offset; i < end && pos < bufferLength; {
		if i >= len(rleData) {
			break
		}

		control := int8(rleData[i])
		i++

		if control >= 0 {
			// Literal run
			length := int(control) + 1

			if (end - i) < length {
				return fmt.Errorf("RLE literal run exceeds input buffer length")
			}
			if (pos + ((length - 1) * sampleOffset)) >= bufferLength {
				return fmt.Errorf("RLE literal run exceeds output buffer length")
			}

			if sampleOffset == 1 {
				copy(buffer[pos:], rleData[i:i+length])
				i += length
				pos += length
			} else {
				for j := 0; j < length; j++ {
					buffer[pos] = rleData[i]
					i++
					pos += sampleOffset
				}
			}
		} else if control >= -127 {
			// Repeat run
			length := int(-control)

			if (pos + ((length - 1) * sampleOffset)) >= bufferLength {
				return fmt.Errorf("RLE repeat run exceeds output buffer length")
			}

			if i >= len(rleData) {
				return io.ErrUnexpectedEOF
			}

			b := rleData[i]
			i++

			if sampleOffset == 1 {
				for j := 0; j <= length; j++ {
					buffer[pos] = b
					pos++
				}
			} else {
				for j := 0; j <= length; j++ {
					buffer[pos] = b
					pos += sampleOffset
				}
			}
		}

		if (i + 1) >= end {
			break
		}
	}

	return nil
}
