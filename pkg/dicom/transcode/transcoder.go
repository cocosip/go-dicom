// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transcode

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"github.com/cocosip/go-dicom/pkg/logging"
)

// Transcoder handles transcoding of DICOM datasets between different transfer syntaxes.
// It can compress, decompress, and convert pixel data formats.
type Transcoder struct {
	inputSyntax   *transfer.Syntax
	outputSyntax  *transfer.Syntax
	inputCodec    codec.Codec
	outputCodec   codec.Codec
	inputParams   codec.Parameters
	outputParams  codec.Parameters
	manager       *Manager
	strictDICOMVR bool // Controls VR selection according to DICOM standard
}

// Option configures a Transcoder created by Manager.
type Option func(*Transcoder)

// WithInputParameters sets the input codec parameters.
func WithInputParameters(params codec.Parameters) Option {
	return func(t *Transcoder) {
		t.inputParams = params
	}
}

// WithOutputParameters sets the output codec parameters.
func WithOutputParameters(params codec.Parameters) Option {
	return func(t *Transcoder) {
		t.outputParams = params
	}
}

// WithStrictDICOMVR controls VR selection for encapsulated pixel data.
// When true (default), encapsulated data always uses OB as required by PS3.5
// Section 8.2. When false, 16-bit encapsulated data uses the non-standard OW
// VR for compatibility with systems that emit or expect that representation.
//
// Native pixel data is independent of this option: Implicit VR Little Endian
// uses OW (PS3.5 A.1), while explicit VR uses OW above 8 Bits Allocated and
// may use OB or OW at 8 Bits Allocated or below.
func WithStrictDICOMVR(strict bool) Option {
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

// Transcode converts a Dataset and associates logging and cancellation with ctx.
func (t *Transcoder) Transcode(ctx context.Context, ds *dataset.Dataset) (result *dataset.Dataset, err error) {
	logDebug := logging.Enabled(ctx, slog.LevelDebug)
	logError := logging.Enabled(ctx, slog.LevelError)
	if logDebug || logError {
		started := time.Now()
		defer func() {
			level := slog.LevelDebug
			event := "transcode_completed"
			message := "DICOM transcode completed"
			if err != nil {
				level = slog.LevelError
				event = "transcode_failed"
				message = "DICOM transcode failed"
			}
			if !logging.Enabled(ctx, level) {
				return
			}
			frameCount := 0
			if ds != nil && ds.Contains(tag.PixelData) {
				frameCount = frameCountFromDataset(ds)
			}
			attrs := []slog.Attr{
				slog.String("input_transfer_syntax", t.inputSyntax.UID().UID()),
				slog.String("output_transfer_syntax", t.outputSyntax.UID().UID()),
				slog.Int("frame_count", frameCount),
				slog.Duration("duration", time.Since(started)),
			}
			if err != nil {
				attrs = append(attrs,
					slog.String("failure_stage", "transcode"),
					slog.String("error_type", fmt.Sprintf("%T", err)),
				)
			}
			logging.Emit(ctx, logging.Record{
				Level: level, Component: "imaging.codec", Event: event, Message: message, Attrs: attrs,
			})
		}()
	}

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
		tempDS, err := t.decode(ctx, ds, transfer.ExplicitVRLittleEndian)
		if err != nil {
			return nil, fmt.Errorf("failed to decode: %w", err)
		}
		return t.encode(ctx, tempDS, t.outputSyntax)
	}

	if inputEncapsulated {
		// Compressed to uncompressed
		return t.decode(ctx, ds, t.outputSyntax)
	}

	if outputEncapsulated {
		// Uncompressed to compressed
		return t.encode(ctx, ds, t.outputSyntax)
	}

	return nil, fmt.Errorf("unable to determine transcode path from %s to %s",
		t.inputSyntax.UID().UID(), t.outputSyntax.UID().UID())
}

// TranscodeWithMetadata transcodes a Dataset while preserving File Meta Information.
func (t *Transcoder) TranscodeWithMetadata(
	ctx context.Context,
	ds *dataset.Dataset,
	sourceMeta *dataset.FileMetaInformation,
) (*dataset.Dataset, *dataset.FileMetaInformation, error) {
	// Transcode the dataset
	newDS, err := t.Transcode(ctx, ds)
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

// DecodeFrame decodes one frame and associates logging and cancellation with ctx.
func (t *Transcoder) DecodeFrame(ctx context.Context, ds *dataset.Dataset, frameIndex int) (decoded []byte, err error) {
	logDebug := logging.Enabled(ctx, slog.LevelDebug)
	logError := logging.Enabled(ctx, slog.LevelError)
	if logDebug || logError {
		started := time.Now()
		defer func() {
			level := slog.LevelDebug
			event := "frame_decode_completed"
			message := "DICOM frame decode completed"
			if err != nil {
				level = slog.LevelError
				event = "frame_decode_failed"
				message = "DICOM frame decode failed"
			}
			if !logging.Enabled(ctx, level) {
				return
			}
			attrs := []slog.Attr{
				slog.String("input_transfer_syntax", t.inputSyntax.UID().UID()),
				slog.Int("frame", frameIndex),
				slog.Duration("duration", time.Since(started)),
			}
			if err != nil {
				attrs = append(attrs,
					slog.String("failure_stage", "frame_decode"),
					slog.String("error_type", fmt.Sprintf("%T", err)),
				)
			} else {
				attrs = append(attrs, slog.Int("output_bytes", len(decoded)))
			}
			logging.Emit(ctx, logging.Record{
				Level: level, Component: "imaging.codec", Event: event, Message: message, Attrs: attrs,
			})
		}()
	}

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

	oldPixelData, err := newFrameData(frameInfo, true)
	if err != nil {
		return nil, fmt.Errorf("create compressed frame source: %w", err)
	}
	if err := oldPixelData.AddFrame(ctx, compressedFrame); err != nil {
		return nil, fmt.Errorf("failed to add frame: %w", err)
	}

	newPixelData, err := newFrameData(frameInfo, false)
	if err != nil {
		return nil, fmt.Errorf("create decoded frame sink: %w", err)
	}

	// Decode using high-level codec method
	parameters, err := codec.PrepareParameters(t.inputCodec, t.inputParams)
	if err != nil {
		return nil, err
	}
	if err := t.inputCodec.Decode(ctx, oldPixelData, newPixelData, parameters); err != nil {
		return nil, fmt.Errorf("failed to decode frame: %w", err)
	}

	// Return the decoded frame
	return newPixelData.Frame(ctx, 0)
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
		if err := normalizeImplicitPixelDataVR(newDS, t.outputSyntax); err != nil {
			return nil, err
		}
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
		if err := normalizeImplicitPixelDataVR(newDS, t.outputSyntax); err != nil {
			return nil, err
		}
		return newDS, nil
	case *element.OtherWord:
		pixelData = elem.GetData()
	default:
		return nil, fmt.Errorf("unexpected pixel data element type for uncompressed data")
	}

	// OW byte order follows the native sample width. An 8-bit sample carried
	// in OW still uses 16-bit words; 32-bit samples are swapped as one value.
	bitsAllocated, _ := ds.GetUInt16(tag.BitsAllocated, 0)
	bytesPerSample := 2
	if bitsAllocated > 8 {
		bytesPerSample = int((bitsAllocated-1)/8 + 1)
	}
	if bytesPerSample != 2 && bytesPerSample != 4 {
		return nil, fmt.Errorf("unsupported OW Bits Allocated=%d", bitsAllocated)
	}
	if len(pixelData)%bytesPerSample != 0 {
		return nil, fmt.Errorf("pixel data length is not aligned to %d-byte OW samples", bytesPerSample)
	}

	convertedData := make([]byte, len(pixelData))
	for i := 0; i < len(pixelData); i += bytesPerSample {
		// Swap bytes
		for left, right := i, i+bytesPerSample-1; left < right; left, right = left+1, right-1 {
			convertedData[left] = pixelData[right]
			convertedData[right] = pixelData[left]
		}
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
	convertedPixelData := element.NewOtherWord(tag.PixelData, convertedData)
	element.SetByteOrder(convertedPixelData, outputEndian.ByteOrder())
	_ = newDS.Add(convertedPixelData)

	return newDS, nil
}

func normalizeImplicitPixelDataVR(ds *dataset.Dataset, outputSyntax *transfer.Syntax) error {
	if outputSyntax == nil || outputSyntax.UID().UID() != transfer.ImplicitVRLittleEndian.UID().UID() {
		return nil
	}
	pixelData, ok := ds.Get(tag.PixelData)
	if !ok {
		return nil
	}
	otherByte, ok := pixelData.(*element.OtherByte)
	if !ok {
		return nil
	}
	if err := ds.AddOrUpdate(element.NewOtherWord(tag.PixelData, otherByte.GetData())); err != nil {
		return fmt.Errorf("normalize implicit VR Pixel Data: %w", err)
	}
	return nil
}

// decode decompresses pixel data from a dataset using the high-level codec.Decode method.
// This follows the fo-dicom pattern of creating complete PixelData objects and using
// the codec's Decode method instead of frame-by-frame processing.
//
// Note: This method uses internal frame ports and converts their output back to DICOM elements.
func (t *Transcoder) decode(ctx context.Context, ds *dataset.Dataset, outputSyntax *transfer.Syntax) (*dataset.Dataset, error) {
	if t.inputCodec == nil {
		return nil, fmt.Errorf("no codec available for decoding %s", t.inputSyntax.UID().UID())
	}

	// Build frame info from dataset
	frameInfo, err := t.buildFrameInfoFromDataset(ds)
	if err != nil {
		return nil, err
	}

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

	oldPixelData, err := newFrameData(frameInfo, true)
	if err != nil {
		return nil, fmt.Errorf("create compressed frame source: %w", err)
	}
	for _, frame := range compressedFrames {
		if err := oldPixelData.AddFrame(ctx, frame); err != nil {
			return nil, fmt.Errorf("failed to add frame to oldPixelData: %w", err)
		}
	}

	newPixelData, err := newFrameData(frameInfo, false)
	if err != nil {
		return nil, fmt.Errorf("create decoded frame sink: %w", err)
	}

	// Use the high-level codec.Decode method
	parameters, err := codec.PrepareParameters(t.inputCodec, t.inputParams)
	if err != nil {
		return nil, err
	}
	if err := t.inputCodec.Decode(ctx, oldPixelData, newPixelData, parameters); err != nil {
		return nil, fmt.Errorf("failed to decode pixel data: %w", err)
	}

	// Reconstruct uncompressed pixel data from decoded frames
	var uncompressedData []byte
	for i := 0; i < newPixelData.FrameCount(); i++ {
		frameData, err := newPixelData.Frame(ctx, i)
		if err != nil {
			return nil, fmt.Errorf("failed to get decoded frame %d: %w", i, err)
		}
		uncompressedData = append(uncompressedData, frameData...)
	}

	// Codec output frames use the native little-endian pixel representation.
	// Build that intermediate Dataset first, then convert it to the requested
	// native syntax so 16-bit pixels are swapped when Big Endian was negotiated.
	newDS := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)

	// Copy all elements except PixelData
	for _, elem := range ds.Elements() {
		if elem.Tag().ToUint32() != tag.PixelData.ToUint32() {
			_ = newDS.Add(elem)
		}
	}

	// Add uncompressed pixel data
	if frameInfo.BitDepth.BitsAllocated <= 8 {
		_ = newDS.Add(element.NewOtherByte(tag.PixelData, uncompressedData))
	} else {
		_ = newDS.Add(element.NewOtherWord(tag.PixelData, uncompressedData))
	}
	if err := applyOutputFrameInfo(newDS, frameInfo, newPixelData.FrameInfo()); err != nil {
		return nil, err
	}

	if outputSyntax.UID().UID() == transfer.ExplicitVRLittleEndian.UID().UID() {
		return newDS, nil
	}

	nativeTranscoder, err := t.manager.NewTranscoder(transfer.ExplicitVRLittleEndian, outputSyntax)
	if err != nil {
		return nil, err
	}
	return nativeTranscoder.Transcode(ctx, newDS)
}

// encode compresses pixel data from a dataset using the high-level codec.Encode method.
// This follows the fo-dicom pattern of creating complete PixelData objects and using
// the codec's Encode method instead of frame-by-frame processing.
func (t *Transcoder) encode(ctx context.Context, ds *dataset.Dataset, outputTS *transfer.Syntax) (*dataset.Dataset, error) {
	if t.outputCodec == nil {
		return nil, fmt.Errorf("no codec available for encoding to %s", outputTS.UID().UID())
	}

	sourceDS := ds
	sourceSyntax := ds.InternalTransferSyntax()
	if sourceSyntax == nil {
		sourceSyntax = t.inputSyntax
	}
	if sourceSyntax != nil && sourceSyntax.UID().UID() != transfer.ExplicitVRLittleEndian.UID().UID() {
		nativeTranscoder, err := t.manager.NewTranscoder(sourceSyntax, transfer.ExplicitVRLittleEndian)
		if err != nil {
			return nil, err
		}
		normalizedDS, err := nativeTranscoder.Transcode(ctx, ds)
		if err != nil {
			return nil, fmt.Errorf("failed to normalize pixel data for encoding: %w", err)
		}
		sourceDS = normalizedDS
	}

	// Build frame info from dataset
	frameInfo, err := t.buildFrameInfoFromDataset(sourceDS)
	if err != nil {
		return nil, err
	}

	// Extract uncompressed pixel data
	pixelDataElem, _ := sourceDS.Get(tag.PixelData)
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
	frameCount := frameCountFromDataset(sourceDS)

	oldPixelData, err := newFrameData(frameInfo, false)
	if err != nil {
		return nil, fmt.Errorf("create native frame source: %w", err)
	}
	frameSize := oldPixelData.Info.UncompressedFrameSize()
	if frameSize <= 0 {
		return nil, fmt.Errorf("invalid frame size calculated")
	}
	for i := 0; i < frameCount; i++ {
		start := i * frameSize
		end := start + frameSize
		if end > len(pixelData) {
			end = len(pixelData)
		}
		if start < len(pixelData) {
			frameData := append([]byte(nil), pixelData[start:end]...)
			if err := oldPixelData.AddFrame(ctx, frameData); err != nil {
				return nil, fmt.Errorf("failed to add frame to oldPixelData: %w", err)
			}
		}
	}

	newPixelData, err := newFrameData(frameInfo, true)
	if err != nil {
		return nil, fmt.Errorf("create encoded frame sink: %w", err)
	}

	// Use the high-level codec.Encode method
	parameters, err := codec.PrepareParameters(t.outputCodec, t.outputParams)
	if err != nil {
		return nil, err
	}
	if err := t.outputCodec.Encode(ctx, oldPixelData, newPixelData, parameters); err != nil {
		return nil, fmt.Errorf("failed to encode pixel data: %w", err)
	}

	// Collect encoded frames and build fragment sequence
	var frameFragments [][]byte
	for i := 0; i < newPixelData.FrameCount(); i++ {
		frameData, err := newPixelData.Frame(ctx, i)
		if err != nil {
			return nil, fmt.Errorf("failed to get encoded frame %d: %w", i, err)
		}
		frameFragments = append(frameFragments, frameData)
	}

	fragSeq, err := buildFragmentSequence(frameFragments, frameInfo.BitDepth.BitsAllocated, t.strictDICOMVR)
	if err != nil {
		return nil, err
	}

	// Create new dataset with the output transfer syntax
	// Clone preserves the original transfer syntax, so we need to create a new one
	newDS := dataset.NewWithTransferSyntax(outputTS)

	// Copy all elements except PixelData
	for _, elem := range sourceDS.Elements() {
		if elem.Tag().ToUint32() != tag.PixelData.ToUint32() {
			_ = newDS.Add(elem)
		}
	}

	// Add encoded pixel data
	_ = newDS.Add(fragSeq)
	if err := applyOutputFrameInfo(newDS, frameInfo, newPixelData.FrameInfo()); err != nil {
		return nil, err
	}
	if err := applyLossyImageCompressionMetadata(ctx, newDS, outputTS, oldPixelData, newPixelData); err != nil {
		return nil, err
	}

	return newDS, nil
}

func applyOutputFrameInfo(ds *dataset.Dataset, input, output codec.FrameInfo) error {
	if output.SamplesPerPixel != input.SamplesPerPixel {
		if err := ds.AddOrUpdate(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{output.SamplesPerPixel})); err != nil {
			return fmt.Errorf("update samples per pixel: %w", err)
		}
	}
	if output.PhotometricInterpretation.Value != "" && output.PhotometricInterpretation.Value != input.PhotometricInterpretation.Value {
		if err := ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{output.PhotometricInterpretation.Value})); err != nil {
			return fmt.Errorf("update photometric interpretation: %w", err)
		}
	}
	if output.PlanarConfiguration != input.PlanarConfiguration {
		if err := ds.AddOrUpdate(element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{uint16(output.PlanarConfiguration)})); err != nil {
			return fmt.Errorf("update planar configuration: %w", err)
		}
	}
	return nil
}

func applyLossyImageCompressionMetadata(
	ctx context.Context,
	ds *dataset.Dataset,
	outputTS *transfer.Syntax,
	inputPixelData, outputPixelData codec.FrameSource,
) error {
	if !outputTS.IsLossy() || outputPixelData.FrameCount() == 0 {
		return nil
	}

	inputBytes, err := pixelDataSize(ctx, inputPixelData)
	if err != nil {
		return fmt.Errorf("measure uncompressed pixel data: %w", err)
	}
	outputBytes, err := pixelDataSize(ctx, outputPixelData)
	if err != nil {
		return fmt.Errorf("measure compressed pixel data: %w", err)
	}
	if inputBytes == 0 || outputBytes == 0 {
		return fmt.Errorf("invalid compression sizes: uncompressed=%d compressed=%d", inputBytes, outputBytes)
	}

	if err := ds.AddOrUpdate(element.NewString(tag.LossyImageCompression, vr.CS, []string{"01"})); err != nil {
		return fmt.Errorf("set lossy image compression: %w", err)
	}
	methods, _ := ds.GetStrings(tag.LossyImageCompressionMethod)
	methods = append(methods, outputTS.LossyCompressionMethod())
	if err := ds.AddOrUpdate(element.NewString(tag.LossyImageCompressionMethod, vr.CS, methods)); err != nil {
		return fmt.Errorf("set lossy image compression method: %w", err)
	}
	ratios, _ := ds.GetStrings(tag.LossyImageCompressionRatio)
	ratios = append(ratios, fmt.Sprintf("%.3f", float64(inputBytes)/float64(outputBytes)))
	if err := ds.AddOrUpdate(element.NewString(tag.LossyImageCompressionRatio, vr.DS, ratios)); err != nil {
		return fmt.Errorf("set lossy image compression ratio: %w", err)
	}
	return nil
}

func pixelDataSize(ctx context.Context, pixelData codec.FrameSource) (int, error) {
	total := 0
	for frame := 0; frame < pixelData.FrameCount(); frame++ {
		data, err := pixelData.Frame(ctx, frame)
		if err != nil {
			return 0, fmt.Errorf("read frame %d: %w", frame, err)
		}
		total += len(data)
	}
	return total, nil
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
			frames = append(frames, codec.StripTrailingPadding(frame))
		}
		return frames, nil
	}

	if frameCount == 1 {
		var frame []byte
		for _, frag := range fragments {
			frame = append(frame, frag.Data()...)
		}
		return [][]byte{codec.StripTrailingPadding(frame)}, nil
	}

	// Fallback: one fragment per frame. Compressed frames are inherently
	// variable-length, so no size-equality check is applied here.
	if frameCount > len(fragments) {
		return nil, fmt.Errorf("frame count %d exceeds available fragments %d without BOT", frameCount, len(fragments))
	}
	var frames [][]byte
	for i := 0; i < frameCount; i++ {
		frames = append(frames, codec.StripTrailingPadding(fragments[i].Data()))
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
		return codec.StripTrailingPadding(frame), nil
	}

	if frameIndex >= len(fragments) {
		return nil, fmt.Errorf("frame index %d out of range [0, %d)", frameIndex, len(fragments))
	}
	data := fragments[frameIndex].Data()
	if len(data) == 0 {
		return nil, fmt.Errorf("fragment %d is empty", frameIndex)
	}
	return codec.StripTrailingPadding(data), nil
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
func (t *Transcoder) buildFrameInfoFromDataset(ds *dataset.Dataset) (codec.FrameInfo, error) {
	// Extract pixel data element to verify it exists
	_, exists := ds.Get(tag.PixelData)
	if !exists {
		return codec.FrameInfo{}, fmt.Errorf("dataset does not contain pixel data")
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
		photometric = pixel.Monochrome2.Value
	}

	photometricValue, err := pixel.ParsePhotometricInterpretation(photometric)
	if err != nil {
		return codec.FrameInfo{}, err
	}
	representation := pixel.Representation(pixelRep)
	return codec.FrameInfo{
		Width:                     width,
		Height:                    height,
		BitDepth:                  *pixel.NewBitDepth(bitsAlloc, bitsStored, highBit, representation.IsSigned()),
		SamplesPerPixel:           samplesPerPixel,
		PixelRepresentation:       representation,
		PlanarConfiguration:       pixel.PlanarConfiguration(planarConfig),
		PhotometricInterpretation: *photometricValue,
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
	bytesAllocated := frameInfo.BitDepth.BytesAllocated()
	frameSize := bytesAllocated * int(frameInfo.SamplesPerPixel) * int(frameInfo.Width) * int(frameInfo.Height)
	offset := frameIndex * frameSize

	if offset+frameSize > len(pixelData) {
		return nil, fmt.Errorf("frame data out of bounds")
	}

	return pixelData[offset : offset+frameSize], nil
}
