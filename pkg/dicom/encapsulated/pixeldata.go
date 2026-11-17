// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

// Package encapsulated centralizes access to encapsulated DICOM Pixel Data.
package encapsulated

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

type offsetSource uint8

const (
	noOffsets offsetSource = iota
	basicOffsets
	extendedOffsets
)

type frameOffsets struct {
	source  offsetSource
	offsets []uint64
	lengths []uint64
}

// PixelData binds a Pixel Data Fragment Sequence to the optional Extended
// Offset Table elements stored alongside it in a Dataset.
type PixelData struct {
	element  element.Element
	sequence *element.FragmentSequence
	offsets  frameOffsets
	// syntax is the Transfer Syntax of the dataset the fragments came from; it
	// identifies the codestream when frame boundaries are implicit.
	syntax *transfer.Syntax
}

// OpenDataset opens encapsulated Pixel Data without changing its original VR
// or the Dataset representation of BOT and EOT values.
func OpenDataset(ds *dataset.Dataset) (*PixelData, error) {
	if ds == nil {
		return nil, fmt.Errorf("dataset cannot be nil")
	}
	pixelDataElement, ok := ds.Get(tag.PixelData)
	if !ok {
		return nil, fmt.Errorf("missing Pixel Data tag")
	}

	var sequence *element.FragmentSequence
	switch value := pixelDataElement.(type) {
	case *element.OtherByteFragment:
		sequence = value.FragmentSequence
	case *element.OtherWordFragment:
		sequence = value.FragmentSequence
	default:
		return nil, fmt.Errorf("pixel data is %T, want encapsulated OB or OW", pixelDataElement)
	}

	offsets, err := readFrameOffsets(ds, sequence.OffsetTable())
	if err != nil {
		return nil, err
	}
	return &PixelData{element: pixelDataElement, sequence: sequence, offsets: offsets, syntax: ds.InternalTransferSyntax()}, nil
}

// ValueRepresentation returns the original Pixel Data VR.
func (p *PixelData) ValueRepresentation() *vr.VR {
	if p == nil || p.element == nil {
		return nil
	}
	return p.element.ValueRepresentation()
}

// FragmentSequence returns the underlying raw Fragment Sequence.
func (p *PixelData) FragmentSequence() *element.FragmentSequence {
	if p == nil {
		return nil
	}
	return p.sequence
}

// Frames assembles every encoded frame using EOT, BOT, or an unambiguous
// no-table layout.
func (p *PixelData) Frames(frameCount int) ([][]byte, error) {
	ranges, err := p.frameRanges(frameCount)
	if err != nil {
		return nil, err
	}
	frames := make([][]byte, 0, len(ranges))
	for frameIndex, frameRange := range ranges {
		frame, err := p.assembleFrame(frameRange, frameIndex)
		if err != nil {
			return nil, err
		}
		frames = append(frames, frame)
	}
	return frames, nil
}

// Frame assembles one encoded frame without loading fragments belonging to
// other frames.
func (p *PixelData) Frame(frameCount, frameIndex int) ([]byte, error) {
	ranges, err := p.frameRanges(frameCount)
	if err != nil {
		return nil, err
	}
	if frameIndex < 0 || frameIndex >= len(ranges) {
		return nil, fmt.Errorf("frame index %d out of range [0, %d)", frameIndex, len(ranges))
	}
	return p.assembleFrame(ranges[frameIndex], frameIndex)
}

type fragmentRange struct {
	start int
	end   int
	// length is the unpadded frame length declared by the Extended Offset
	// Table Lengths; 0 means no authoritative length is available.
	length uint64
}

func (p *PixelData) frameRanges(frameCount int) ([]fragmentRange, error) {
	if p == nil || p.sequence == nil {
		return nil, fmt.Errorf("encapsulated Pixel Data cannot be nil")
	}
	fragments := p.sequence.Fragments()
	if len(fragments) == 0 {
		return nil, fmt.Errorf("no fragments available")
	}
	for i, fragment := range fragments {
		if fragment == nil {
			return nil, fmt.Errorf("fragment %d is nil", i)
		}
	}
	// Zero-length fragment items carry no data; the parser and writer tolerate
	// them (PS3.5 A.4 permits empty fragment items), so they contribute no bytes
	// to a frame instead of failing the whole read.

	if p.offsets.source == noOffsets {
		return rangesWithoutOffsets(fragments, frameCount, p.syntax)
	}
	if frameCount < 1 {
		frameCount = len(p.offsets.offsets)
	}
	if frameCount != len(p.offsets.offsets) {
		return nil, fmt.Errorf("offset table frames mismatch: expected %d, got %d entries", frameCount, len(p.offsets.offsets))
	}
	if len(p.offsets.offsets) == 0 || p.offsets.offsets[0] != 0 {
		return nil, fmt.Errorf("offset table must start at zero")
	}

	fragmentByOffset := make(map[uint64]int, len(fragments))
	var encodedLength uint64
	for i, fragment := range fragments {
		fragmentByOffset[encodedLength] = i
		valueLength := uint64(fragment.Size())
		if valueLength%2 != 0 {
			valueLength++
		}
		itemLength := uint64(8) + valueLength
		if encodedLength > math.MaxUint64-itemLength {
			return nil, fmt.Errorf("fragment offset overflow at index %d", i)
		}
		encodedLength += itemLength
	}

	ranges := make([]fragmentRange, frameCount)
	for frameIndex, startOffset := range p.offsets.offsets {
		start, ok := fragmentByOffset[startOffset]
		if !ok {
			return nil, fmt.Errorf("offset %d for frame %d does not align with a Fragment Item", startOffset, frameIndex)
		}
		endOffset := encodedLength
		if frameIndex+1 < len(p.offsets.offsets) {
			endOffset = p.offsets.offsets[frameIndex+1]
		}
		end := len(fragments)
		if endOffset != encodedLength {
			var found bool
			end, found = fragmentByOffset[endOffset]
			if !found {
				return nil, fmt.Errorf("offset %d for frame %d does not align with a Fragment Item", endOffset, frameIndex+1)
			}
		}
		if start >= end {
			return nil, fmt.Errorf("frame %d derived from offset table is empty", frameIndex)
		}
		var frameLength uint64
		if p.offsets.source == extendedOffsets {
			length := p.offsets.lengths[frameIndex]
			var payloadLength uint64
			for fragmentIndex := start; fragmentIndex < end; fragmentIndex++ {
				// Extended Offset Table Lengths exclude the trailing pad byte of an
				// odd-length fragment (PS3.3 Table C.7-11a), so do not round up.
				fragmentLength := uint64(fragments[fragmentIndex].Size())
				if payloadLength > math.MaxUint64-fragmentLength {
					return nil, fmt.Errorf("extended offset table length overflow for frame %d", frameIndex)
				}
				payloadLength += fragmentLength
			}
			// A padded last fragment stores one byte more than the frame length.
			paddedLastFragment := payloadLength == length+1 &&
				uint64(fragments[end-1].Size())%2 == 0
			if length == 0 || (length != payloadLength && !paddedLastFragment) {
				return nil, fmt.Errorf("extended offset table length %d for frame %d does not match Fragment Items", length, frameIndex)
			}
			frameLength = length
		}
		ranges[frameIndex] = fragmentRange{start: start, end: end, length: frameLength}
	}
	return ranges, nil
}

func rangesWithoutOffsets(fragments []buffer.ByteBuffer, frameCount int, syntax *transfer.Syntax) ([]fragmentRange, error) {
	if frameCount < 1 {
		frameCount = len(fragments)
	}
	if frameCount == 1 {
		return []fragmentRange{{start: 0, end: len(fragments)}}, nil
	}
	if frameCount != len(fragments) {
		if ranges, ok := rangesFromFrameEndMarkers(fragments, frameCount, syntax); ok {
			return ranges, nil
		}
		return nil, fmt.Errorf(
			"cannot determine frame boundaries without BOT or EOT: %d frames use %d fragments",
			frameCount,
			len(fragments),
		)
	}
	ranges := make([]fragmentRange, frameCount)
	for i := range ranges {
		ranges[i] = fragmentRange{start: i, end: i + 1}
	}
	return ranges, nil
}

// rangesFromFrameEndMarkers splits fragments into frames at the end-of-stream
// marker of the compressed bit stream. PS3.5 A.4 allows an empty Basic Offset
// Table, so a sender may leave the frame boundaries implicit; codestreams in
// the JPEG transfer syntax family all terminate with FF D9, which cannot occur
// unescaped inside them, so the marker identifies the frame boundaries.
func rangesFromFrameEndMarkers(fragments []buffer.ByteBuffer, frameCount int, syntax *transfer.Syntax) ([]fragmentRange, bool) {
	if !usesFrameEndMarker(syntax) {
		return nil, false
	}
	ranges := make([]fragmentRange, 0, frameCount)
	start := 0
	previousEndedWithMarkerPrefix := false
	for index, fragment := range fragments {
		size := fragment.Size()
		data := make([]byte, size)
		if size == 0 || fragment.GetByteRange(0, size, data) != nil {
			return nil, false
		}
		marker := previousEndedWithMarkerPrefix && data[0] == frameEndMarker[1]
		for position := 0; position+1 < len(data); position++ {
			if data[position] == frameEndMarker[0] && data[position+1] == frameEndMarker[1] {
				marker = true
			}
		}
		previousEndedWithMarkerPrefix = data[len(data)-1] == frameEndMarker[0]
		if marker {
			ranges = append(ranges, fragmentRange{start: start, end: index + 1})
			start = index + 1
		}
	}
	if len(ranges) != frameCount {
		return nil, false
	}
	if start != len(fragments) {
		// Trailing fragments after the last marker complete the last frame.
		ranges[len(ranges)-1].end = len(fragments)
	}
	return ranges, true
}

// frameEndMarker is the end-of-stream marker of the JPEG and JPEG 2000
// codestreams (JPEG EOI, JPEG 2000 EOC).
var frameEndMarker = [2]byte{0xff, 0xd9}

// usesFrameEndMarker reports whether the transfer syntax carries a codestream
// that ends with FF D9.
func usesFrameEndMarker(syntax *transfer.Syntax) bool {
	if syntax == nil {
		return false
	}
	return strings.HasPrefix(syntax.UID().UID(), "1.2.840.10008.1.2.4.")
}

func (p *PixelData) assembleFrame(frameRange fragmentRange, frameIndex int) ([]byte, error) {
	fragments := p.sequence.Fragments()
	var frame []byte
	for _, fragment := range fragments[frameRange.start:frameRange.end] {
		size := fragment.Size()
		if size == 0 {
			continue
		}
		chunk := make([]byte, size)
		if err := fragment.GetByteRange(0, size, chunk); err != nil {
			return nil, fmt.Errorf("read frame %d fragment: %w", frameIndex, err)
		}
		frame = append(frame, chunk...)
	}
	if len(frame) == 0 {
		return nil, fmt.Errorf("frame %d is empty", frameIndex)
	}
	// Extended Offset Table Lengths exclude the fragment pad byte, so drop it.
	if frameRange.length > 0 && frameRange.length < uint64(len(frame)) {
		frame = frame[:frameRange.length]
	}
	return frame, nil
}

func readFrameOffsets(ds *dataset.Dataset, basic []uint32) (frameOffsets, error) {
	offsetsElement, hasOffsets := ds.Get(tag.ExtendedOffsetTable)
	lengthsElement, hasLengths := ds.Get(tag.ExtendedOffsetTableLengths)
	if hasOffsets || hasLengths {
		if !hasOffsets || !hasLengths {
			return frameOffsets{}, fmt.Errorf("extended offset table and Extended Offset Table Lengths must both be present")
		}
		offsets, err := readUint64Values(ds, offsetsElement, "Extended Offset Table")
		if err != nil {
			return frameOffsets{}, err
		}
		lengths, err := readUint64Values(ds, lengthsElement, "Extended Offset Table Lengths")
		if err != nil {
			return frameOffsets{}, err
		}
		if len(offsets) == 0 || len(offsets) != len(lengths) {
			return frameOffsets{}, fmt.Errorf(
				"extended offset table entry mismatch: got %d offsets and %d lengths",
				len(offsets),
				len(lengths),
			)
		}
		return frameOffsets{source: extendedOffsets, offsets: offsets, lengths: lengths}, nil
	}
	if len(basic) == 0 {
		return frameOffsets{source: noOffsets}, nil
	}
	offsets := make([]uint64, len(basic))
	for i, offset := range basic {
		offsets[i] = uint64(offset)
	}
	return frameOffsets{source: basicOffsets, offsets: offsets}, nil
}

func readUint64Values(ds *dataset.Dataset, valueElement element.Element, name string) ([]uint64, error) {
	value, ok := valueElement.(*element.OtherVeryLong)
	if !ok {
		return nil, fmt.Errorf("%s has unexpected VR/type %T", name, valueElement)
	}
	data := value.GetData()
	if len(data)%8 != 0 {
		return nil, fmt.Errorf("%s length %d is not divisible by 8", name, len(data))
	}
	order, known := element.NumericByteOrder(valueElement)
	if !known {
		order = binary.LittleEndian
		if syntax := ds.InternalTransferSyntax(); syntax != nil {
			order = syntax.Endian().ByteOrder()
		}
	}
	values := make([]uint64, len(data)/8)
	for i := range values {
		values[i] = order.Uint64(data[i*8:])
	}
	return values, nil
}

// IsPixelDataMetadataTag reports whether t belongs to the encapsulated Pixel
// Data representation and must be replaced or removed together.
func IsPixelDataMetadataTag(t *tag.Tag) bool {
	if t == nil {
		return false
	}
	value := t.ToUint32()
	return value == tag.PixelData.ToUint32() ||
		value == tag.ExtendedOffsetTable.ToUint32() ||
		value == tag.ExtendedOffsetTableLengths.ToUint32()
}
