// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"fmt"
	"math"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/types"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// Transcoder handles transcoding of DICOM datasets between different transfer syntaxes.
// It can compress, decompress, and convert pixel data formats.
type Transcoder struct {
	inputSyntax   *transfer.Syntax
	outputSyntax  *transfer.Syntax
	inputCodec    Codec
	outputCodec   Codec
	inputParams   Parameters
	outputParams  Parameters
	codecRegistry *Registry
}

// TranscoderOption is a functional option for creating a Transcoder.
type TranscoderOption func(*Transcoder)

// NewTranscoder creates a new Transcoder.
func NewTranscoder(inputSyntax, outputSyntax *transfer.Syntax, opts ...TranscoderOption) *Transcoder {
	t := &Transcoder{
		inputSyntax:   inputSyntax,
		outputSyntax:  outputSyntax,
		codecRegistry: GetGlobalRegistry(),
	}

	// Apply options
	for _, opt := range opts {
		opt(t)
	}

	// Get codecs from registry if not provided
	if t.inputCodec == nil && inputSyntax.IsEncapsulated() {
		if codec, ok := t.codecRegistry.GetCodec(inputSyntax); ok {
			t.inputCodec = codec
		}
	}

	if t.outputCodec == nil && outputSyntax.IsEncapsulated() {
		if codec, ok := t.codecRegistry.GetCodec(outputSyntax); ok {
			t.outputCodec = codec
		}
	}

	return t
}

// WithInputCodec sets the input codec explicitly.
func WithInputCodec(codec Codec) TranscoderOption {
	return func(t *Transcoder) {
		t.inputCodec = codec
	}
}

// WithOutputCodec sets the output codec explicitly.
func WithOutputCodec(codec Codec) TranscoderOption {
	return func(t *Transcoder) {
		t.outputCodec = codec
	}
}

// WithInputParameters sets the input codec parameters.
func WithInputParameters(params Parameters) TranscoderOption {
	return func(t *Transcoder) {
		t.inputParams = params
	}
}

// WithOutputParameters sets the output codec parameters.
func WithOutputParameters(params Parameters) TranscoderOption {
	return func(t *Transcoder) {
		t.outputParams = params
	}
}

// WithCodecRegistry sets a custom codec registry.
func WithCodecRegistry(registry *Registry) TranscoderOption {
	return func(t *Transcoder) {
		t.codecRegistry = registry
	}
}

// InputSyntax returns the input transfer syntax.
func (t *Transcoder) InputSyntax() *transfer.Syntax {
	return t.inputSyntax
}

// OutputSyntax returns the output transfer syntax.
func (t *Transcoder) OutputSyntax() *transfer.Syntax {
	return t.outputSyntax
}

// Transcode converts a dataset from input transfer syntax to output transfer syntax.
func (t *Transcoder) Transcode(ds *dataset.Dataset) (*dataset.Dataset, error) {
	// Check if dataset contains pixel data
	if !ds.Contains(tag.PixelData) {
		// No pixel data - just clone and update transfer syntax
		newDS := ds.Clone()
		// Note: In a full implementation, we would set InternalTransferSyntax
		// For now, return the cloned dataset
		return newDS, nil
	}

	// Get current transfer syntax from dataset (simplified - would come from file meta info)
	inputEncapsulated := t.inputSyntax.IsEncapsulated()
	outputEncapsulated := t.outputSyntax.IsEncapsulated()

	if !inputEncapsulated && !outputEncapsulated {
		// Uncompressed to uncompressed
		return t.transcodeUncompressedToUncompressed(ds)
	}

	if inputEncapsulated && outputEncapsulated {
		// Compressed to compressed (decompress then compress)
		tempDS, err := t.decode(ds, transfer.ExplicitVRLittleEndian)
		if err != nil {
			return nil, fmt.Errorf("failed to decode: %w", err)
		}
		return t.encode(tempDS, t.outputSyntax)
	}

	if inputEncapsulated {
		// Compressed to uncompressed
		return t.decode(ds, t.outputSyntax)
	}

	if outputEncapsulated {
		// Uncompressed to compressed
		return t.encode(ds, t.outputSyntax)
	}

	return nil, fmt.Errorf("unable to determine transcode path from %s to %s",
		t.inputSyntax.UID().UID(), t.outputSyntax.UID().UID())
}

// DecodeFrame decodes a single frame from compressed pixel data.
// This follows fo-dicom's pattern of creating a temporary PixelData with one frame,
// then using the high-level Decode method.
func (t *Transcoder) DecodeFrame(ds *dataset.Dataset, frameIndex int) ([]byte, error) {
	// Check if pixel data exists
	pixelDataElem, exists := ds.Get(tag.PixelData)
	if !exists {
		return nil, fmt.Errorf("dataset does not contain pixel data")
	}

	// If already uncompressed, extract the frame directly
	if !t.inputSyntax.IsEncapsulated() {
		return t.extractUncompressedFrame(ds, frameIndex)
	}

	// Get fragment sequence
	var fragments []buffer.ByteBuffer
	switch pd := pixelDataElem.(type) {
	case *element.OtherByteFragment:
		fragments = pd.Fragments()
	case *element.OtherWordFragment:
		fragments = pd.Fragments()
	default:
		return nil, fmt.Errorf("expected fragment sequence for encapsulated transfer syntax")
	}

	if frameIndex >= len(fragments) {
		return nil, fmt.Errorf("frame index %d out of range (0-%d)", frameIndex, len(fragments)-1)
	}

	// Get the compressed frame
	compressedFrame := fragments[frameIndex].Data()

	// Decode using input codec
	if t.inputCodec == nil {
		return nil, fmt.Errorf("no codec available for input transfer syntax %s",
			t.inputSyntax.UID().UID())
	}

	// Build frame info from dataset
	frameInfo, err := t.buildFrameInfoFromDataset(ds)
	if err != nil {
		return nil, fmt.Errorf("failed to build frame info: %w", err)
	}

	// Create temporary PixelData with only this frame (fo-dicom pattern)
	oldPixelData := newSimplePixelData(frameInfo, true) // encapsulated=true
	if err := oldPixelData.AddFrame(compressedFrame); err != nil {
		return nil, fmt.Errorf("failed to add frame: %w", err)
	}

	newPixelData := newSimplePixelData(frameInfo, false) // encapsulated=false

	// Decode using high-level codec method
	if err := t.inputCodec.Decode(oldPixelData, newPixelData, t.inputParams); err != nil {
		return nil, fmt.Errorf("failed to decode frame: %w", err)
	}

	// Return the decoded frame
	return newPixelData.GetFrame(0)
}

// transcodeUncompressedToUncompressed handles conversion between uncompressed formats.
func (t *Transcoder) transcodeUncompressedToUncompressed(ds *dataset.Dataset) (*dataset.Dataset, error) {
	// Clone dataset
	newDS := ds.Clone()

	// In a full implementation, would handle:
	// - Byte order conversion
	// - Planar configuration conversion
	// - Photometric interpretation conversion

	// For now, just clone
	return newDS, nil
}

// decode decompresses pixel data from a dataset using the high-level codec.Decode method.
// This follows the fo-dicom pattern of creating complete PixelData objects and using
// the codec's Decode method instead of frame-by-frame processing.
//
// Note: This method requires types.PixelData implementations that also support conversion
// back to DICOM elements (e.g., DicomPixelData from the imaging package).
func (t *Transcoder) decode(ds *dataset.Dataset, _ *transfer.Syntax) (*dataset.Dataset, error) {
	if t.inputCodec == nil {
		return nil, fmt.Errorf("no codec available for decoding %s", t.inputSyntax.UID().UID())
	}

	// Build frame info from dataset
	frameInfo, err := t.buildFrameInfoFromDataset(ds)
	if err != nil {
		return nil, err
	}

	// Create temporary PixelData objects using testPixelData helper (same as DecodeFrame pattern)
	// This allows us to use the high-level codec.Decode without a circular dependency

	// Extract compressed pixel data
	pixelDataElem, _ := ds.Get(tag.PixelData)
	var fragments []buffer.ByteBuffer
	var offsetTable []uint32

	switch pd := pixelDataElem.(type) {
	case *element.OtherByteFragment:
		fragments = pd.Fragments()
		offsetTable = pd.OffsetTable()
	case *element.OtherWordFragment:
		fragments = pd.Fragments()
		offsetTable = pd.OffsetTable()
	default:
		return nil, fmt.Errorf("unexpected pixel data element type for encapsulated transfer syntax")
	}

	// Determine frame count
	frameCount := 1
	if nf, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil { 
		frameCount = int(nf)
	}
	if frameCount < 1 {
		frameCount = len(fragments)
	}
	if frameCount < 1 {
		frameCount = 1
	}

	// Build compressed frames from offset table or one-fragment-per-frame
	compressedFrames, err := framesFromFragments(fragments, offsetTable, frameCount)
	if err != nil {
		return nil, err
	}

	// Create old PixelData with all compressed frames
	oldPixelData := newSimplePixelData(frameInfo, true) // encapsulated=true
	for _, frame := range compressedFrames {
		if err := oldPixelData.AddFrame(frame); err != nil {
			return nil, fmt.Errorf("failed to add frame to oldPixelData: %w", err)
		}
	}

	// Create new PixelData for decoded output
	newPixelData := newSimplePixelData(frameInfo, false) // encapsulated=false

	// Use the high-level codec.Decode method
	if err := t.inputCodec.Decode(oldPixelData, newPixelData, t.inputParams); err != nil {
		return nil, fmt.Errorf("failed to decode pixel data: %w", err)
	}

	// Create new dataset with decoded pixel data
	newDS := ds.Clone()
	newDS.Remove(tag.PixelData)

	// Reconstruct uncompressed pixel data from decoded frames
	var uncompressedData []byte
	for i := 0; i < newPixelData.FrameCount(); i++ {
		frameData, err := newPixelData.GetFrame(i)
		if err != nil {
			return nil, fmt.Errorf("failed to get decoded frame %d: %w", i, err)
		}
		uncompressedData = append(uncompressedData, frameData...)
	}

	// Add uncompressed pixel data
	if frameInfo.BitsAllocated <= 8 {
		_ = newDS.Add(element.NewOtherByte(tag.PixelData, uncompressedData))
	} else {
		_ = newDS.Add(element.NewOtherWord(tag.PixelData, uncompressedData))
	}

	return newDS, nil
}

// encode compresses pixel data from a dataset using the high-level codec.Encode method.
// This follows the fo-dicom pattern of creating complete PixelData objects and using
// the codec's Encode method instead of frame-by-frame processing.
func (t *Transcoder) encode(ds *dataset.Dataset, outputTS *transfer.Syntax) (*dataset.Dataset, error) {
	if t.outputCodec == nil {
		return nil, fmt.Errorf("no codec available for encoding to %s", outputTS.UID().UID())
	}

	// Build frame info from dataset
	frameInfo, err := t.buildFrameInfoFromDataset(ds)
	if err != nil {
		return nil, err
	}

	// Extract uncompressed pixel data
	pixelDataElem, _ := ds.Get(tag.PixelData)
	var pixelData []byte
	switch elem := pixelDataElem.(type) {
	case *element.OtherByte:
		pixelData = elem.GetData()
	case *element.OtherWord:
		pixelData = elem.GetData()
	default:
		return nil, fmt.Errorf("unexpected pixel data element type for uncompressed data")
	}

	// Determine frame count
	frameCount := 1
	if nf, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil {
		frameCount = int(nf)
	}

	// Calculate frame size
	bytesAllocated := int((frameInfo.BitsAllocated - 1) / 8 + 1)
	frameSize := bytesAllocated * int(frameInfo.SamplesPerPixel) * int(frameInfo.Width) * int(frameInfo.Height)
	if frameSize <= 0 {
		return nil, fmt.Errorf("invalid frame size calculated")
	}

	// Split uncompressed data into frames
	oldPixelData := newSimplePixelData(frameInfo, false) // encapsulated=false
	for i := 0; i < frameCount; i++ {
		start := i * frameSize
		end := start + frameSize
		if end > len(pixelData) {
			end = len(pixelData)
		}
		if start < len(pixelData) {
			frameData := pixelData[start:end]
			if err := oldPixelData.AddFrame(frameData); err != nil {
				return nil, fmt.Errorf("failed to add frame to oldPixelData: %w", err)
			}
		}
	}

	// Create new PixelData for encoded output
	newPixelData := newSimplePixelData(frameInfo, true) // encapsulated=true

	// Use the high-level codec.Encode method
	if err := t.outputCodec.Encode(oldPixelData, newPixelData, t.outputParams); err != nil {
		return nil, fmt.Errorf("failed to encode pixel data: %w", err)
	}

	// Collect encoded frames and build fragment sequence
	var frameFragments [][]byte
	for i := 0; i < newPixelData.FrameCount(); i++ {
		frameData, err := newPixelData.GetFrame(i)
		if err != nil {
			return nil, fmt.Errorf("failed to get encoded frame %d: %w", i, err)
		}
		frameFragments = append(frameFragments, frameData)
	}

	obf, err := buildFragmentSequence(frameFragments)
	if err != nil {
		return nil, err
	}

	// Create new dataset with encoded pixel data
	newDS := ds.Clone()
	newDS.Remove(tag.PixelData)
	_ = newDS.Add(obf)

	return newDS, nil
}

// framesFromFragments builds per-frame compressed data using fragments and an optional BOT.
// If offsetTable is present, it slices the concatenated fragments using offsets.
// Otherwise, it assumes one fragment per frame (best-effort fallback).
func framesFromFragments(fragments []buffer.ByteBuffer, offsetTable []uint32, frameCount int) ([][]byte, error) {
	if len(fragments) == 0 {
		return nil, fmt.Errorf("no fragments available to decode")
	}

	for i, frag := range fragments {
		if len(frag.Data()) == 0 {
			return nil, fmt.Errorf("fragment %d is empty", i)
		}
	}

	if frameCount < 1 {
		frameCount = len(offsetTable)
	}
	if frameCount < 1 {
		frameCount = len(fragments)
	}
	if frameCount < 1 {
		frameCount = 1
	}

	// BOT present: slice concatenated stream by offsets.
	if len(offsetTable) > 0 {
		if frameCount != len(offsetTable) {
			return nil, fmt.Errorf("offset table frames mismatch: expected %d, got %d entries", frameCount, len(offsetTable))
		}
		// Concatenate all fragment data.
		var concat []byte
		for _, frag := range fragments {
			concat = append(concat, frag.Data()...)
		}

		var frames [][]byte
		for i := 0; i < frameCount; i++ {
			start := int(offsetTable[i])
			end := len(concat)
			if i+1 < len(offsetTable) {
				end = int(offsetTable[i+1])
			}
			if start < 0 || end < 0 || start > end || end > len(concat) {
				return nil, fmt.Errorf("invalid BOT slice for frame %d: start %d end %d total %d", i, start, end, len(concat))
			}
			if start == end {
				return nil, fmt.Errorf("frame %d derived from BOT is empty", i)
			}
			frames = append(frames, concat[start:end])
		}
		return frames, nil
	}

	// Fallback: require one fragment per frame and consistent sizes.
	if frameCount > len(fragments) {
		return nil, fmt.Errorf("frame count %d exceeds available fragments %d without BOT", frameCount, len(fragments))
	}
	framesToUse := frameCount
	var frames [][]byte
	var expectedSize int
	for i := 0; i < framesToUse; i++ {
		data := fragments[i].Data()
		if i == 0 {
			expectedSize = len(data)
		} else if len(data) != expectedSize {
			return nil, fmt.Errorf("fragment size mismatch at index %d: got %d, expected %d", i, len(data), expectedSize)
		}
		frames = append(frames, data)
	}
	return frames, nil
}

// buildFragmentSequence creates an OB fragment sequence from per-frame compressed data,
// populating the Basic Offset Table for multi-frame images.
func buildFragmentSequence(frames [][]byte) (*element.OtherByteFragment, error) {
	obf := element.NewOtherByteFragment(tag.PixelData)

	// Single-frame images may omit BOT; still add the fragment.
	if len(frames) == 0 {
		return nil, fmt.Errorf("no frame data provided for fragment sequence")
	}

	var offsets []uint32
	var runningOffset uint32
	for i, frame := range frames {
		offsets = append(offsets, runningOffset)

		if len(frame) > int(math.MaxUint32-runningOffset) {
			return nil, fmt.Errorf("fragment too large to represent in BOT at frame %d", i)
		}
		runningOffset += uint32(len(frame))

		obf.AddFragment(buffer.NewMemory(frame))
	}

	if len(frames) > 1 {
		obf.SetOffsetTable(offsets)
	}

	return obf, nil
}

// buildFrameInfoFromDataset extracts frame metadata from a dataset.
func (t *Transcoder) buildFrameInfoFromDataset(ds *dataset.Dataset) (*types.FrameInfo, error) {
	// Extract pixel data element to verify it exists
	_, exists := ds.Get(tag.PixelData)
	if !exists {
		return nil, fmt.Errorf("dataset does not contain pixel data")
	}

	// Extract image attributes using dataset accessors
	width := ds.TryGetUInt16(tag.Columns, 0)
	height := ds.TryGetUInt16(tag.Rows, 0)

	bitsAlloc := ds.TryGetUInt16(tag.BitsAllocated, 0)
	bitsStored := ds.TryGetUInt16(tag.BitsStored, 0)
	if bitsStored == 0 {
		bitsStored = bitsAlloc
	}

	highBit := ds.TryGetUInt16(tag.HighBit, 0)
	if highBit == 0 && bitsStored > 0 {
		highBit = bitsStored - 1
	}

	samplesPerPixel := ds.TryGetUInt16(tag.SamplesPerPixel, 0)
	if samplesPerPixel == 0 {
		samplesPerPixel = 1
	}

	pixelRep := ds.TryGetUInt16(tag.PixelRepresentation, 0)
	planarConfig := ds.TryGetUInt16(tag.PlanarConfiguration, 0)

	photometric := ds.TryGetString(tag.PhotometricInterpretation)
	if photometric == "" {
		photometric = "MONOCHROME2"
	}

	return &types.FrameInfo{
		Width:                     width,
		Height:                    height,
		BitsAllocated:             bitsAlloc,
		BitsStored:                bitsStored,
		HighBit:                   highBit,
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       pixelRep,
		PlanarConfiguration:       planarConfig,
		PhotometricInterpretation: photometric,
	}, nil
}

// extractUncompressedFrame extracts a single frame from uncompressed pixel data.
func (t *Transcoder) extractUncompressedFrame(ds *dataset.Dataset, frameIndex int) ([]byte, error) {
	frameInfo, err := t.buildFrameInfoFromDataset(ds)
	if err != nil {
		return nil, err
	}

	// Get frame count
	frameCount := 1
	if nf, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil {
		frameCount = int(nf)
	}

	if frameIndex >= frameCount {
		return nil, fmt.Errorf("frame index %d out of range (0-%d)",
			frameIndex, frameCount-1)
	}

	// Extract pixel data
	pixelDataElem, _ := ds.Get(tag.PixelData)
	var pixelData []byte
	switch elem := pixelDataElem.(type) {
	case *element.OtherByte:
		pixelData = elem.GetData()
	case *element.OtherWord:
		pixelData = elem.GetData()
	default:
		return nil, fmt.Errorf("unexpected pixel data element type")
	}

	// Calculate frame size
	bytesAllocated := int((frameInfo.BitsAllocated-1)/8 + 1)
	frameSize := bytesAllocated * int(frameInfo.SamplesPerPixel) * int(frameInfo.Width) * int(frameInfo.Height)
	offset := frameIndex * frameSize

	if offset+frameSize > len(pixelData) {
		return nil, fmt.Errorf("frame data out of bounds")
	}

	return pixelData[offset : offset+frameSize], nil
}
