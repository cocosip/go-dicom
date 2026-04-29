// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
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
	strictDICOMVR bool // Controls VR selection according to DICOM standard
}

// TranscoderOption is a functional option for creating a Transcoder.
type TranscoderOption func(*Transcoder)

// NewTranscoder creates a new Transcoder.
func NewTranscoder(inputSyntax, outputSyntax *transfer.Syntax, opts ...TranscoderOption) *Transcoder {
	t := &Transcoder{
		inputSyntax:   inputSyntax,
		outputSyntax:  outputSyntax,
		codecRegistry: GetGlobalRegistry(),
		strictDICOMVR: true, // Default: follow DICOM standard for VR selection (OB for encapsulated data)
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

// WithStrictDICOMVR controls VR selection for pixel data according to DICOM standard.
// When true (default, recommended):
//   - Uncompressed data: VR selected based on BitsAllocated (OB for <=8 bits, OW for >8 bits)
//   - Compressed data: VR is always OB (DICOM Part 5 Section 8.2)
//
// When false (compatibility mode):
//   - Both compressed and uncompressed data: VR selected based on BitsAllocated
//
// Note: Setting this to false may violate DICOM standard but might be needed for compatibility.
func WithStrictDICOMVR(strict bool) TranscoderOption {
	return func(t *Transcoder) {
		t.strictDICOMVR = strict
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
// Note: This method only transcodes the dataset and does not preserve File Meta Information.
// If you need to preserve File Meta Information, use TranscodeWithMetadata instead.
func (t *Transcoder) Transcode(ds *dataset.Dataset) (*dataset.Dataset, error) {
	// Check if dataset contains pixel data
	if !ds.Contains(tag.PixelData) {
		// No pixel data - just clone and update transfer syntax
		newDS := ds.Clone()
		newDS.SetInternalTransferSyntax(t.outputSyntax)
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

// TranscodeWithMetadata converts a dataset from input transfer syntax to output transfer syntax
// while preserving File Meta Information from the source.
//
// This method transcodes the dataset and preserves the File Meta Information from the source,
// updating only the Transfer Syntax UID to match the output format. This is useful when you
// want to preserve all File Meta Information tags (like SourceApplicationEntityTitle) after
// transcoding.
//
// Parameters:
//   - ds: The main DICOM dataset to transcode
//   - sourceMeta: The File Meta Information from the source file (optional)
//     If nil, a new minimal File Meta Information will be created
//
// Returns:
//   - Transcoded dataset
//   - Updated File Meta Information (preserves all fields except Transfer Syntax UID)
//   - Error if transcoding fails
//
// Example usage:
//
//	result, _ := parser.ParseFile("input.dcm")
//	transcoder := NewTranscoder(inputTS, outputTS)
//	newDS, newMeta, _ := transcoder.TranscodeWithMetadata(
//	    result.Dataset,
//	    result.FileMetaInformation)
//	writer.WriteFile("output.dcm", newDS, writer.WithFileMetaInfo(newMeta.Dataset()))
func (t *Transcoder) TranscodeWithMetadata(ds *dataset.Dataset, sourceMeta *dataset.FileMetaInformation) (*dataset.Dataset, *dataset.FileMetaInformation, error) {
	// Transcode the dataset
	newDS, err := t.Transcode(ds)
	if err != nil {
		return nil, nil, err
	}

	// If no source metadata provided, create a new minimal one
	if sourceMeta == nil {
		newMeta, err := dataset.NewFileMetaInformationFromMainDataset(newDS, t.outputSyntax)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create file meta information: %w", err)
		}
		return newDS, newMeta, nil
	}

	// Clone the source File Meta Information by creating a new one from the cloned dataset
	newMeta := dataset.NewFileMetaInformationFromDataset(sourceMeta.Dataset().Clone())

	// Update Transfer Syntax UID to match the output transfer syntax
	if err := newMeta.SetTransferSyntax(t.outputSyntax); err != nil {
		return nil, nil, fmt.Errorf("failed to update transfer syntax: %w", err)
	}

	return newDS, newMeta, nil
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
	var offsetTable []uint32
	switch pd := pixelDataElem.(type) {
	case *element.OtherByteFragment:
		fragments = pd.Fragments()
		offsetTable = pd.OffsetTable()
	case *element.OtherWordFragment:
		fragments = pd.Fragments()
		offsetTable = pd.OffsetTable()
	default:
		return nil, fmt.Errorf("expected fragment sequence for encapsulated transfer syntax")
	}

	frameCount := frameCountFromDataset(ds)
	compressedFrame, err := frameFromFragments(fragments, offsetTable, frameCount, frameIndex)
	if err != nil {
		return nil, err
	}

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
// This includes byte order conversion, planar configuration changes, etc.
func (t *Transcoder) transcodeUncompressedToUncompressed(ds *dataset.Dataset) (*dataset.Dataset, error) {
	// Check if we need any actual conversion
	inputEndian := t.inputSyntax.Endian()
	outputEndian := t.outputSyntax.Endian()

	// If input and output are the same, just clone and update transfer syntax
	if inputEndian == outputEndian {
		newDS := ds.Clone()
		newDS.SetInternalTransferSyntax(t.outputSyntax)
		return newDS, nil
	}

	// Need byte order conversion for pixel data
	// Get pixel data element
	pixelDataElem, exists := ds.Get(tag.PixelData)
	if !exists {
		// No pixel data, just clone
		newDS := ds.Clone()
		newDS.SetInternalTransferSyntax(t.outputSyntax)
		return newDS, nil
	}

	// Extract pixel data
	var pixelData []byte
	switch elem := pixelDataElem.(type) {
	case *element.OtherByte:
		// 8-bit data doesn't need byte order conversion
		newDS := ds.Clone()
		newDS.SetInternalTransferSyntax(t.outputSyntax)
		return newDS, nil
	case *element.OtherWord:
		pixelData = elem.GetData()
	default:
		return nil, fmt.Errorf("unexpected pixel data element type for uncompressed data")
	}

	// Get frame info
	frameInfo, err := t.buildFrameInfoFromDataset(ds)
	if err != nil {
		return nil, err
	}

	// Only need to swap bytes for multi-byte data (BitsAllocated > 8)
	if frameInfo.BitsAllocated <= 8 {
		newDS := ds.Clone()
		newDS.SetInternalTransferSyntax(t.outputSyntax)
		return newDS, nil
	}

	// Swap byte order for 16-bit data
	if len(pixelData)%2 != 0 {
		return nil, fmt.Errorf("pixel data length is not even for 16-bit data")
	}

	convertedData := make([]byte, len(pixelData))
	for i := 0; i < len(pixelData); i += 2 {
		// Swap bytes
		convertedData[i] = pixelData[i+1]
		convertedData[i+1] = pixelData[i]
	}

	// Create new dataset with converted pixel data
	newDS := dataset.NewWithTransferSyntax(t.outputSyntax)

	// Copy all elements except PixelData
	for _, elem := range ds.Elements() {
		if elem.Tag().ToUint32() != tag.PixelData.ToUint32() {
			_ = newDS.Add(elem)
		}
	}

	// Add converted pixel data
	_ = newDS.Add(element.NewOtherWord(tag.PixelData, convertedData))

	return newDS, nil
}

// decode decompresses pixel data from a dataset using the high-level codec.Decode method.
// This follows the fo-dicom pattern of creating complete PixelData objects and using
// the codec's Decode method instead of frame-by-frame processing.
//
// Note: This method requires imagetypes.PixelData implementations that also support conversion
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
	frameCount := frameCountFromDataset(ds)
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

	// Reconstruct uncompressed pixel data from decoded frames
	var uncompressedData []byte
	for i := 0; i < newPixelData.FrameCount(); i++ {
		frameData, err := newPixelData.GetFrame(i)
		if err != nil {
			return nil, fmt.Errorf("failed to get decoded frame %d: %w", i, err)
		}
		uncompressedData = append(uncompressedData, frameData...)
	}

	// Create new dataset with the output transfer syntax
	newDS := dataset.NewWithTransferSyntax(t.outputSyntax)

	// Copy all elements except PixelData
	for _, elem := range ds.Elements() {
		if elem.Tag().ToUint32() != tag.PixelData.ToUint32() {
			_ = newDS.Add(elem)
		}
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
	// Note: NumberOfFrames has VR of IS (Integer String) per DICOM standard
	frameCount := frameCountFromDataset(ds)

	// Calculate frame size
	bytesAllocated := int((frameInfo.BitsAllocated-1)/8 + 1)
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

	fragSeq, err := buildFragmentSequence(frameFragments, frameInfo.BitsAllocated, t.strictDICOMVR)
	if err != nil {
		return nil, err
	}

	// Create new dataset with the output transfer syntax
	// Clone preserves the original transfer syntax, so we need to create a new one
	newDS := dataset.NewWithTransferSyntax(outputTS)

	// Copy all elements except PixelData
	for _, elem := range ds.Elements() {
		if elem.Tag().ToUint32() != tag.PixelData.ToUint32() {
			_ = newDS.Add(elem)
		}
	}

	// Add encoded pixel data
	_ = newDS.Add(fragSeq)

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

		fragmentStartByOffset := make(map[uint32]int, len(fragments))
		var runningOffset uint32
		for i, frag := range fragments {
			fragmentStartByOffset[runningOffset] = i
			size := frag.Size()
			if size%2 != 0 {
				size++
			}
			if size > math.MaxUint32-8 || runningOffset > math.MaxUint32-8-size {
				return nil, fmt.Errorf("fragment offset overflow at index %d", i)
			}
			runningOffset += 8 + size
		}

		frameStartIndexes := make([]int, frameCount)
		for i := 0; i < frameCount; i++ {
			fragmentIndex, ok := fragmentStartByOffset[offsetTable[i]]
			if !ok {
				return nil, fmt.Errorf("BOT offset %d for frame %d does not align with a fragment item", offsetTable[i], i)
			}
			frameStartIndexes[i] = fragmentIndex
		}

		var frames [][]byte
		for i, start := range frameStartIndexes {
			end := len(fragments)
			if i+1 < len(frameStartIndexes) {
				end = frameStartIndexes[i+1]
			}
			if start >= end {
				return nil, fmt.Errorf("frame %d derived from BOT is empty", i)
			}

			var frame []byte
			for _, frag := range fragments[start:end] {
				frame = append(frame, frag.Data()...)
			}
			frames = append(frames, StripTrailingPadding(frame))
		}
		return frames, nil
	}

	if frameCount == 1 {
		var frame []byte
		for _, frag := range fragments {
			frame = append(frame, frag.Data()...)
		}
		return [][]byte{StripTrailingPadding(frame)}, nil
	}

	// Fallback: one fragment per frame. Compressed frames are inherently
	// variable-length, so no size-equality check is applied here.
	if frameCount > len(fragments) {
		return nil, fmt.Errorf("frame count %d exceeds available fragments %d without BOT", frameCount, len(fragments))
	}
	var frames [][]byte
	for i := 0; i < frameCount; i++ {
		frames = append(frames, StripTrailingPadding(fragments[i].Data()))
	}
	return frames, nil
}

func frameFromFragments(fragments []buffer.ByteBuffer, offsetTable []uint32, frameCount, frameIndex int) ([]byte, error) {
	if len(fragments) == 0 {
		return nil, fmt.Errorf("no fragments available to decode")
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
	if frameIndex < 0 || frameIndex >= frameCount {
		return nil, fmt.Errorf("frame index %d out of range [0, %d)", frameIndex, frameCount)
	}

	if len(offsetTable) > 0 {
		if frameCount != len(offsetTable) {
			return nil, fmt.Errorf("offset table frames mismatch: expected %d, got %d entries", frameCount, len(offsetTable))
		}

		fragmentStartByOffset := make(map[uint32]int, len(fragments))
		var runningOffset uint32
		for i, frag := range fragments {
			fragmentStartByOffset[runningOffset] = i
			size := frag.Size()
			if size%2 != 0 {
				size++
			}
			if size > math.MaxUint32-8 || runningOffset > math.MaxUint32-8-size {
				return nil, fmt.Errorf("fragment offset overflow at index %d", i)
			}
			runningOffset += 8 + size
		}

		start, ok := fragmentStartByOffset[offsetTable[frameIndex]]
		if !ok {
			return nil, fmt.Errorf("BOT offset %d for frame %d does not align with a fragment item", offsetTable[frameIndex], frameIndex)
		}
		end := len(fragments)
		if frameIndex+1 < len(offsetTable) {
			nextEnd, ok := fragmentStartByOffset[offsetTable[frameIndex+1]]
			if !ok {
				return nil, fmt.Errorf("BOT offset %d for frame %d does not align with a fragment item", offsetTable[frameIndex+1], frameIndex+1)
			}
			end = nextEnd
		}
		if start >= end {
			return nil, fmt.Errorf("frame %d derived from BOT is empty", frameIndex)
		}

		var frame []byte
		for _, frag := range fragments[start:end] {
			frame = append(frame, frag.Data()...)
		}
		if len(frame) == 0 {
			return nil, fmt.Errorf("frame %d is empty", frameIndex)
		}
		return StripTrailingPadding(frame), nil
	}

	if frameIndex >= len(fragments) {
		return nil, fmt.Errorf("frame index %d out of range [0, %d)", frameIndex, len(fragments))
	}
	data := fragments[frameIndex].Data()
	if len(data) == 0 {
		return nil, fmt.Errorf("fragment %d is empty", frameIndex)
	}
	return StripTrailingPadding(data), nil
}

func frameCountFromDataset(ds *dataset.Dataset) int {
	if nf, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil && nf > 0 {
		return int(nf)
	}
	if nfStr, ok := ds.GetString(tag.NumberOfFrames); ok {
		if parsed, err := strconv.Atoi(strings.TrimSpace(nfStr)); err == nil && parsed > 0 {
			return parsed
		}
	}
	return 1
}

// buildFragmentSequence creates a fragment sequence from per-frame compressed data,
// populating the Basic Offset Table.
//
// VR Selection according to DICOM standard:
//   - For encapsulated (compressed) pixel data, DICOM Part 5 Section 8.2 specifies
//     that VR should always be OB (Other Byte), regardless of BitsAllocated.
//   - For native (uncompressed) pixel data, VR is selected based on BitsAllocated:
//     OB for <=8 bits, OW for >8 bits.
//
// Parameters:
//   - frames: Per-frame compressed data
//   - bitsAllocated: Bits allocated per pixel
//   - strictDICOM: Controls VR selection mode:
//     true:  Use OB for encapsulated data (DICOM standard compliant, recommended)
//     false: Select VR based on BitsAllocated even for encapsulated data (compatibility mode)
func buildFragmentSequence(frames [][]byte, bitsAllocated uint16, strictDICOM bool) (element.Element, error) {
	if len(frames) == 0 {
		return nil, fmt.Errorf("no frame data provided for fragment sequence")
	}

	// Build Basic Offset Table
	// According to DICOM standard, Basic Offset Table should contain at least one offset (0x00000000) for single-frame,
	// and all frame offsets for multi-frame images.
	var offsets []uint32
	var runningOffset uint32
	for i, frame := range frames {
		offsets = append(offsets, runningOffset)

		paddedSize := len(frame)
		if paddedSize%2 != 0 {
			paddedSize++
		}
		if paddedSize > int(math.MaxUint32-8) {
			return nil, fmt.Errorf("fragment too large to represent in BOT at frame %d", i)
		}
		padded := uint32(paddedSize)
		if runningOffset > math.MaxUint32-8-padded {
			return nil, fmt.Errorf("fragment offset overflow at frame %d", i)
		}
		runningOffset += 8 + padded
	}

	// Determine VR based on strictDICOM mode
	if strictDICOM {
		// According to DICOM Part 5 Section 8.2:
		// "If sent in an Encapsulated Format (i.e., other than the Native Format)
		//  the Value Representation OB is used."
		// This is the recommended and standard-compliant mode.
		obf := element.NewOtherByteFragment(tag.PixelData)
		for _, frame := range frames {
			obf.AddFragment(buffer.NewMemory(frame))
		}
		obf.SetOffsetTable(offsets)
		return obf, nil
	}

	// strictDICOM is false: select VR based on BitsAllocated (compatibility mode)
	// This may be needed for compatibility with certain non-standard implementations
	// that expect OW for 16-bit data even when encapsulated.
	if bitsAllocated <= 8 {
		obf := element.NewOtherByteFragment(tag.PixelData)
		for _, frame := range frames {
			obf.AddFragment(buffer.NewMemory(frame))
		}
		obf.SetOffsetTable(offsets)
		return obf, nil
	}

	// BitsAllocated > 8: use OW
	owf := element.NewOtherWordFragment(tag.PixelData)
	for _, frame := range frames {
		owf.AddFragment(buffer.NewMemory(frame))
	}
	owf.SetOffsetTable(offsets)
	return owf, nil
}

// buildFrameInfoFromDataset extracts frame metadata from a dataset.
func (t *Transcoder) buildFrameInfoFromDataset(ds *dataset.Dataset) (*imagetypes.FrameInfo, error) {
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

	return &imagetypes.FrameInfo{
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
	frameCount := frameCountFromDataset(ds)

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
