// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
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

	// Build PixelData for decoding
	srcPixelData, err := t.buildPixelDataFromDataset(ds)
	if err != nil {
		return nil, fmt.Errorf("failed to build pixel data: %w", err)
	}

	srcPixelData.Data = compressedFrame
	srcPixelData.NumberOfFrames = 1

	// Create destination pixel data
	dstPixelData := &PixelData{
		Width:                     srcPixelData.Width,
		Height:                    srcPixelData.Height,
		NumberOfFrames:            1,
		BitsAllocated:             srcPixelData.BitsAllocated,
		BitsStored:                srcPixelData.BitsStored,
		HighBit:                   srcPixelData.HighBit,
		SamplesPerPixel:           srcPixelData.SamplesPerPixel,
		PixelRepresentation:       srcPixelData.PixelRepresentation,
		PlanarConfiguration:       srcPixelData.PlanarConfiguration,
		PhotometricInterpretation: srcPixelData.PhotometricInterpretation,
		TransferSyntaxUID:         t.outputSyntax.UID().UID(),
	}

	// Decode
	if err := t.inputCodec.Decode(srcPixelData, dstPixelData, t.inputParams); err != nil {
		return nil, fmt.Errorf("failed to decode frame: %w", err)
	}

	return dstPixelData.Data, nil
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

// decode decompresses pixel data from a dataset.
func (t *Transcoder) decode(ds *dataset.Dataset, outputTS *transfer.Syntax) (*dataset.Dataset, error) {
	if t.inputCodec == nil {
		return nil, fmt.Errorf("no codec available for decoding %s", t.inputSyntax.UID().UID())
	}

	// Build source pixel data
	srcPixelData, err := t.buildPixelDataFromDataset(ds)
	if err != nil {
		return nil, err
	}

	// Extract compressed data from fragment sequence
	pixelDataElem, _ := ds.Get(tag.PixelData)

	var allCompressedData []byte
	switch pd := pixelDataElem.(type) {
	case *element.OtherByteFragment:
		for _, frag := range pd.Fragments() {
			allCompressedData = append(allCompressedData, frag.Data()...)
		}
	case *element.OtherWordFragment:
		for _, frag := range pd.Fragments() {
			allCompressedData = append(allCompressedData, frag.Data()...)
		}
	default:
		return nil, fmt.Errorf("unexpected pixel data element type for encapsulated transfer syntax")
	}

	srcPixelData.Data = allCompressedData

	// Create destination pixel data
	dstPixelData := &PixelData{
		Width:                     srcPixelData.Width,
		Height:                    srcPixelData.Height,
		NumberOfFrames:            srcPixelData.NumberOfFrames,
		BitsAllocated:             srcPixelData.BitsAllocated,
		BitsStored:                srcPixelData.BitsStored,
		HighBit:                   srcPixelData.HighBit,
		SamplesPerPixel:           srcPixelData.SamplesPerPixel,
		PixelRepresentation:       srcPixelData.PixelRepresentation,
		PlanarConfiguration:       srcPixelData.PlanarConfiguration,
		PhotometricInterpretation: srcPixelData.PhotometricInterpretation,
		TransferSyntaxUID:         outputTS.UID().UID(),
	}

	// Decode
	if err := t.inputCodec.Decode(srcPixelData, dstPixelData, t.inputParams); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}

	// Create new dataset with decoded pixel data
	newDS := ds.Clone()

	// Remove old pixel data
	newDS.Remove(tag.PixelData)

	// Add uncompressed pixel data
	// Use appropriate element type based on bits allocated
	if dstPixelData.BitsAllocated <= 8 {
		_ = newDS.Add(element.NewOtherByte(tag.PixelData, dstPixelData.Data))
	} else {
		_ = newDS.Add(element.NewOtherWord(tag.PixelData, dstPixelData.Data))
	}

	return newDS, nil
}

// encode compresses pixel data from a dataset.
func (t *Transcoder) encode(ds *dataset.Dataset, outputTS *transfer.Syntax) (*dataset.Dataset, error) {
	if t.outputCodec == nil {
		return nil, fmt.Errorf("no codec available for encoding to %s", outputTS.UID().UID())
	}

	// Build source pixel data
	srcPixelData, err := t.buildPixelDataFromDataset(ds)
	if err != nil {
		return nil, err
	}

	// Create destination pixel data
	dstPixelData := &PixelData{
		Width:                     srcPixelData.Width,
		Height:                    srcPixelData.Height,
		NumberOfFrames:            srcPixelData.NumberOfFrames,
		BitsAllocated:             srcPixelData.BitsAllocated,
		BitsStored:                srcPixelData.BitsStored,
		HighBit:                   srcPixelData.HighBit,
		SamplesPerPixel:           srcPixelData.SamplesPerPixel,
		PixelRepresentation:       srcPixelData.PixelRepresentation,
		PlanarConfiguration:       srcPixelData.PlanarConfiguration,
		PhotometricInterpretation: srcPixelData.PhotometricInterpretation,
		TransferSyntaxUID:         outputTS.UID().UID(),
	}

	// Encode
	if err := t.outputCodec.Encode(srcPixelData, dstPixelData, t.outputParams); err != nil {
		return nil, fmt.Errorf("encode failed: %w", err)
	}

	// Create new dataset with encoded pixel data
	newDS := ds.Clone()

	// Remove old pixel data
	newDS.Remove(tag.PixelData)

	// Add compressed pixel data as fragment sequence
	// This is simplified - in reality, we need to split into proper frames
	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.AddFragment(buffer.NewMemory(dstPixelData.Data))
	_ = newDS.Add(obf)

	return newDS, nil
}

// buildPixelDataFromDataset extracts pixel data information from a dataset.
func (t *Transcoder) buildPixelDataFromDataset(ds *dataset.Dataset) (*PixelData, error) {
	// Extract pixel data element
	pixelDataElem, exists := ds.Get(tag.PixelData)
	if !exists {
		return nil, fmt.Errorf("dataset does not contain pixel data")
	}

	// Extract image attributes using dataset accessors
	width := ds.TryGetUInt16(tag.Columns, 0)
	height := ds.TryGetUInt16(tag.Rows, 0)

	// NumberOfFrames is often stored as IS (Integer String), try int32 first
	frames := int32(1)
	if val, err := ds.GetInt32(tag.NumberOfFrames, 0); err == nil {
		frames = val
	}

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

	pd := &PixelData{
		Width:                     width,
		Height:                    height,
		NumberOfFrames:            int(frames),
		BitsAllocated:             bitsAlloc,
		BitsStored:                bitsStored,
		HighBit:                   highBit,
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       pixelRep,
		PlanarConfiguration:       planarConfig,
		PhotometricInterpretation: photometric,
		TransferSyntaxUID:         t.inputSyntax.UID().UID(),
	}

	// Extract data
	switch elem := pixelDataElem.(type) {
	case *element.OtherByte:
		pd.Data = elem.GetData()
	case *element.OtherWord:
		pd.Data = elem.GetData()
	case *element.OtherByteFragment:
		// Will be handled by decode function
	case *element.OtherWordFragment:
		// Will be handled by decode function
	default:
		return nil, fmt.Errorf("unexpected pixel data element type: %T", elem)
	}

	return pd, nil
}

// extractUncompressedFrame extracts a single frame from uncompressed pixel data.
func (t *Transcoder) extractUncompressedFrame(ds *dataset.Dataset, frameIndex int) ([]byte, error) {
	pixelData, err := t.buildPixelDataFromDataset(ds)
	if err != nil {
		return nil, err
	}

	if frameIndex >= pixelData.NumberOfFrames {
		return nil, fmt.Errorf("frame index %d out of range (0-%d)",
			frameIndex, pixelData.NumberOfFrames-1)
	}

	frameSize := pixelData.UncompressedFrameSize()
	offset := frameIndex * frameSize

	if offset+frameSize > len(pixelData.Data) {
		return nil, fmt.Errorf("frame data out of bounds")
	}

	return pixelData.Data[offset : offset+frameSize], nil
}
