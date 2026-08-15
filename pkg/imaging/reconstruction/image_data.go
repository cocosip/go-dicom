// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/geometry"
)

const (
	maxReconstructionFrames     = 65535
	maxReconstructionDimension  = 16384
	maxReconstructionPixelBytes = int64(1 << 30)
)

// ImageData is one immutable source frame used by a reconstructed volume.
type ImageData struct {
	dataset              *dataset.Dataset
	frameIndex           int
	geometry             *geometry.FrameGeometry
	pixelData            *imaging.DicomPixelData
	sourceSOPClassUID    string
	sourceSOPInstanceUID string
	sortingPosition      float64
	rescaleSlope         float64
	rescaleIntercept     float64
	pixelFrame           int
}

// NewImageData builds one source-frame view. Native frame bytes are retained;
// encapsulated inputs are decoded one frame at a time through the codec registry.
func NewImageData(ds *dataset.Dataset, frame int) (*ImageData, error) {
	if ds == nil {
		return nil, fmt.Errorf("reconstruction source dataset is nil")
	}
	if err := validateEnhancedDimensions(ds); err != nil {
		return nil, err
	}
	count, err := sourceFrameCount(ds)
	if err != nil {
		return nil, err
	}
	if err := validateSourceDimensions(ds, count); err != nil {
		return nil, err
	}
	pixels, err := imaging.CreatePixelData(ds)
	if err != nil {
		return nil, fmt.Errorf("read frame %d pixel data: %w", frame, err)
	}
	return newImageData(ds, frame, pixels)
}

func newImageData(ds *dataset.Dataset, frame int, sourcePixels *imaging.DicomPixelData) (*ImageData, error) {
	sopClass := ds.TryGetString(tag.SOPClassUID)
	if !supportedSourceSOPClass(sopClass) {
		return nil, fmt.Errorf("unsupported reconstruction SOP Class UID %q", sopClass)
	}
	frameGeometry, err := geometry.NewFrameGeometry(ds, frame)
	if err != nil {
		return nil, fmt.Errorf("read frame %d geometry: %w", frame, err)
	}
	if frameGeometry.Type != geometry.GeometryVolume {
		return nil, fmt.Errorf("frame %d does not contain complete patient-space geometry", frame)
	}
	if frameGeometry.FrameOfReferenceUID == "" {
		return nil, fmt.Errorf("frame %d is missing Frame of Reference UID", frame)
	}

	pixels := sourcePixels
	if pixels == nil || pixels.Info == nil {
		return nil, fmt.Errorf("read frame %d pixel data: pixel data is nil", frame)
	}
	if pixels.Info.SamplesPerPixel != 1 || pixels.Info.PhotometricInterpretation == nil ||
		(pixels.Info.PhotometricInterpretation.Value != "MONOCHROME1" && pixels.Info.PhotometricInterpretation.Value != "MONOCHROME2") {
		return nil, fmt.Errorf("reconstruction requires monochrome single-sample pixel data")
	}
	if sourceModality(sopClass) == "CT" && pixels.Info.PhotometricInterpretation.Value == "MONOCHROME1" {
		return nil, fmt.Errorf("CT reconstruction requires MONOCHROME2 pixel data, got MONOCHROME1")
	}
	if pixels.Info.BitsAllocated != 8 && pixels.Info.BitsAllocated != 16 && pixels.Info.BitsAllocated != 32 {
		return nil, fmt.Errorf("unsupported Bits Allocated %d", pixels.Info.BitsAllocated)
	}
	if frame < 0 || frame >= pixels.FrameCount() {
		return nil, fmt.Errorf("frame index %d out of pixel data range [0, %d)", frame, pixels.FrameCount())
	}

	pixelFrame := frame
	if pixels.IsEncapsulated() {
		pixels, err = decodeFrame(ds, pixels, frame)
		if err != nil {
			return nil, err
		}
		pixelFrame = 0
	}

	slope, intercept, err := frameRescale(ds, frame)
	if err != nil {
		return nil, err
	}
	normal := frameGeometry.DirectionNormal
	position := frameGeometry.TopLeft
	return &ImageData{
		dataset:              ds.Clone(),
		frameIndex:           frame,
		geometry:             frameGeometry,
		pixelData:            pixels,
		sourceSOPClassUID:    sopClass,
		sourceSOPInstanceUID: ds.TryGetString(tag.SOPInstanceUID),
		sortingPosition:      position.X*normal.X + position.Y*normal.Y + position.Z*normal.Z,
		rescaleSlope:         slope,
		rescaleIntercept:     intercept,
		pixelFrame:           pixelFrame,
	}, nil
}

// NewImageDataFromDataset expands all frames in a classic or Enhanced CT/MR dataset.
func NewImageDataFromDataset(ds *dataset.Dataset) ([]*ImageData, error) {
	if ds == nil {
		return nil, fmt.Errorf("reconstruction source dataset is nil")
	}
	if err := validateEnhancedDimensions(ds); err != nil {
		return nil, err
	}
	count, err := sourceFrameCount(ds)
	if err != nil {
		return nil, err
	}
	if err := validateSourceDimensions(ds, count); err != nil {
		return nil, err
	}
	pixels, err := imaging.CreatePixelData(ds)
	if err != nil {
		return nil, fmt.Errorf("read pixel data: %w", err)
	}
	images := make([]*ImageData, count)
	for frame := range count {
		image, err := newImageData(ds, frame, pixels)
		if err != nil {
			return nil, err
		}
		images[frame] = image
	}
	return images, nil
}

func sourceFrameCount(ds *dataset.Dataset) (int, error) {
	count := 1
	if value, ok := ds.Get(tag.NumberOfFrames); ok {
		parsed, err := integerStringValue(value)
		if err != nil || parsed < 1 {
			return 0, fmt.Errorf("invalid Number of Frames")
		}
		count = parsed
	}
	if count > maxReconstructionFrames {
		return 0, fmt.Errorf("source frame count %d exceeds reconstruction limit %d", count, maxReconstructionFrames)
	}
	return count, nil
}

func validateSourceDimensions(ds *dataset.Dataset, frames int) error {
	rows, err := ds.GetUInt16(tag.Rows, 0)
	if err != nil {
		return fmt.Errorf("read source Rows: %w", err)
	}
	columns, err := ds.GetUInt16(tag.Columns, 0)
	if err != nil {
		return fmt.Errorf("read source Columns: %w", err)
	}
	if rows == 0 || columns == 0 || rows > maxReconstructionDimension || columns > maxReconstructionDimension {
		return fmt.Errorf("source dimensions %dx%d exceed reconstruction limit %dx%d",
			rows, columns, maxReconstructionDimension, maxReconstructionDimension)
	}
	bitsAllocated, err := ds.GetUInt16(tag.BitsAllocated, 0)
	if err != nil {
		return fmt.Errorf("read source Bits Allocated: %w", err)
	}
	samplesPerPixel, err := ds.GetUInt16(tag.SamplesPerPixel, 0)
	if err != nil {
		return fmt.Errorf("read source Samples per Pixel: %w", err)
	}
	bytesPerSample := int64((bitsAllocated + 7) / 8)
	total, ok := checkedProduct64(int64(rows), int64(columns), int64(frames), int64(samplesPerPixel), bytesPerSample)
	if !ok || total > maxReconstructionPixelBytes {
		return fmt.Errorf("source pixel data size exceeds reconstruction limit %d bytes", maxReconstructionPixelBytes)
	}
	return nil
}

func checkedProduct64(values ...int64) (int64, bool) {
	product := int64(1)
	for _, value := range values {
		if value <= 0 || product > (1<<63-1)/value {
			return 0, false
		}
		product *= value
	}
	return product, true
}

func validateEnhancedDimensions(ds *dataset.Dataset) error {
	if !ds.Contains(tag.DimensionIndexSequence) {
		return nil
	}
	sequence, err := ds.GetSequence(tag.DimensionIndexSequence)
	if err != nil {
		return fmt.Errorf("read Dimension Index Sequence: %w", err)
	}
	if sequence.Count() == 0 {
		return fmt.Errorf("dimension index sequence is empty")
	}
	pointers := make([]*tag.Tag, sequence.Count())
	variableSpatial := make([]bool, sequence.Count())
	stackIndex := -1
	for index := 0; index < sequence.Count(); index++ {
		item := sequence.GetItem(index)
		if item == nil {
			return fmt.Errorf("dimension index sequence item %d is nil", index)
		}
		value, ok := item.Get(tag.DimensionIndexPointer)
		if !ok {
			return fmt.Errorf("dimension index sequence item %d is missing Dimension Index Pointer", index)
		}
		pointer, ok := value.(*element.AttributeTag)
		if !ok || pointer.Count() != 1 {
			return fmt.Errorf("dimension index pointer in item %d is not a single Attribute Tag", index)
		}
		target, err := pointer.GetValue(0)
		if err != nil {
			return fmt.Errorf("read Dimension Index Pointer in item %d: %w", index, err)
		}
		pointers[index] = target
		switch target.ToUint32() {
		case tag.StackID.ToUint32():
			stackIndex = index
		case tag.InStackPositionNumber.ToUint32(), tag.ImagePositionPatient.ToUint32():
			variableSpatial[index] = true
		}
	}

	perFrame, err := ds.GetSequence(tag.PerFrameFunctionalGroupsSequence)
	if err != nil {
		return fmt.Errorf("read Per-frame Functional Groups Sequence: %w", err)
	}
	var firstIndexes []uint32
	stackID := ""
	for frame := 0; frame < perFrame.Count(); frame++ {
		frameGroup := perFrame.GetItem(frame)
		if frameGroup == nil {
			return fmt.Errorf("per-frame functional group item %d is nil", frame)
		}
		frameContent, err := frameGroup.GetSequence(tag.FrameContentSequence)
		if err != nil || frameContent.Count() != 1 || frameContent.GetItem(0) == nil {
			return fmt.Errorf("frame %d requires one Frame Content Sequence item", frame)
		}
		content := frameContent.GetItem(0)
		indexes, err := content.GetUInt32s(tag.DimensionIndexValues)
		if err != nil {
			return fmt.Errorf("read frame %d Dimension Index Values: %w", frame, err)
		}
		if len(indexes) != len(pointers) {
			return fmt.Errorf("frame %d has %d Dimension Index Values, want %d", frame, len(indexes), len(pointers))
		}

		frameStackID, hasStackID := content.GetString(tag.StackID)
		frameStackID = strings.TrimSpace(frameStackID)
		if stackIndex >= 0 && (!hasStackID || frameStackID == "") {
			return fmt.Errorf("frame %d is missing Stack ID", frame)
		}
		if frameStackID != "" {
			if stackID == "" {
				stackID = frameStackID
			} else if frameStackID != stackID {
				return fmt.Errorf("enhanced source contains multiple Stack IDs %q and %q", stackID, frameStackID)
			}
		}

		if frame == 0 {
			firstIndexes = append([]uint32(nil), indexes...)
			continue
		}
		for index, pointer := range pointers {
			if !variableSpatial[index] && indexes[index] != firstIndexes[index] {
				return fmt.Errorf("non-spatial dimension %s varies between frames", pointer)
			}
		}
	}
	return nil
}

// ValueAt returns a modality-space sample. Padding samples are reported invalid.
func (image *ImageData) ValueAt(x, y int) (float64, bool, error) {
	if image == nil || image.pixelData == nil {
		return 0, false, fmt.Errorf("image pixel data is nil")
	}
	stored, err := image.pixelData.GetSample(image.pixelFrame, x, y, 0)
	if err != nil {
		return 0, false, err
	}
	if image.pixelData.IsPaddingSample(stored) {
		return 0, false, nil
	}
	return float64(stored)*image.rescaleSlope + image.rescaleIntercept, true, nil
}

// FrameIndex returns the zero-based frame number in the source instance.
func (image *ImageData) FrameIndex() int {
	if image == nil {
		return 0
	}
	return image.frameIndex
}

// Geometry returns a copy of the source frame geometry.
func (image *ImageData) Geometry() geometry.FrameGeometry {
	if image == nil || image.geometry == nil {
		return geometry.FrameGeometry{}
	}
	return *image.geometry
}

// SourceSOPClassUID returns the source instance's SOP Class UID.
func (image *ImageData) SourceSOPClassUID() string {
	if image == nil {
		return ""
	}
	return image.sourceSOPClassUID
}

// SourceSOPInstanceUID returns the source instance's SOP Instance UID.
func (image *ImageData) SourceSOPInstanceUID() string {
	if image == nil {
		return ""
	}
	return image.sourceSOPInstanceUID
}

// SourceDataset returns an independent Dataset container for source metadata.
func (image *ImageData) SourceDataset() *dataset.Dataset {
	if image == nil || image.dataset == nil {
		return dataset.New()
	}
	return image.dataset.DeepClone()
}

// SortingPosition returns the frame position projected onto the volume normal.
func (image *ImageData) SortingPosition() float64 {
	if image == nil {
		return 0
	}
	return image.sortingPosition
}

// RescaleSlope returns the frame's modality rescale slope.
func (image *ImageData) RescaleSlope() float64 {
	if image == nil {
		return 0
	}
	return image.rescaleSlope
}

// RescaleIntercept returns the frame's modality rescale intercept.
func (image *ImageData) RescaleIntercept() float64 {
	if image == nil {
		return 0
	}
	return image.rescaleIntercept
}

func supportedSourceSOPClass(value string) bool {
	return value == uid.CTImageStorage.UID() || value == uid.EnhancedCTImageStorage.UID() ||
		value == uid.MRImageStorage.UID() || value == uid.EnhancedMRImageStorage.UID()
}

func decodeFrame(ds *dataset.Dataset, source *imaging.DicomPixelData, frame int) (*imaging.DicomPixelData, error) {
	inputSyntax := ds.InternalTransferSyntax()
	if inputSyntax == nil {
		parsed, err := transfer.Parse(source.Info.TransferSyntaxUID)
		if err != nil {
			return nil, fmt.Errorf("parse source transfer syntax: %w", err)
		}
		inputSyntax = parsed
	}
	transcoder, err := codec.GetDefaultManager().CreateTranscoder(inputSyntax, transfer.ExplicitVRLittleEndian)
	if err != nil {
		return nil, fmt.Errorf("prepare decoder for frame %d: %w", frame, err)
	}
	data, err := transcoder.DecodeFrame(ds, frame)
	if err != nil {
		return nil, fmt.Errorf("decode frame %d: %w", frame, err)
	}
	info := *source.Info
	info.NumberOfFrames = 1
	info.Encapsulated = false
	info.TransferSyntaxUID = transfer.ExplicitVRLittleEndian.UID().UID()
	return imaging.NewDicomPixelDataFromBytes(&info, data)
}

func frameRescale(ds *dataset.Dataset, frame int) (float64, float64, error) {
	values := dataset.New()
	mergeFunctionalValues(values, ds, tag.SharedFunctionalGroupsSequence, 0)
	mergeFunctionalValues(values, ds, tag.PerFrameFunctionalGroupsSequence, frame)
	slope, err := decimalValue(values, ds, tag.RescaleSlope, 1)
	if err != nil {
		return 0, 0, err
	}
	intercept, err := decimalValue(values, ds, tag.RescaleIntercept, 0)
	if err != nil {
		return 0, 0, err
	}
	if slope == 0 {
		return 0, 0, fmt.Errorf("rescale slope must not be zero")
	}
	return slope, intercept, nil
}

func decimalValue(functional, primary *dataset.Dataset, target *tag.Tag, fallback float64) (float64, error) {
	for _, candidate := range []*dataset.Dataset{functional, primary} {
		if candidate == nil {
			continue
		}
		value, ok := candidate.Get(target)
		if !ok {
			continue
		}
		parsed, err := decimalStringValue(value)
		if err != nil {
			return 0, fmt.Errorf("read %s: %w", target, err)
		}
		return parsed, nil
	}
	return fallback, nil
}

func integerStringValue(value element.Element) (int, error) {
	switch typed := value.(type) {
	case *element.IntegerString:
		return typed.GetInt(0)
	case *element.String:
		if typed.ValueRepresentation().Code() != vr.CodeIS {
			return 0, fmt.Errorf("value representation is %s, want IS", typed.ValueRepresentation().Code())
		}
		return strconv.Atoi(strings.TrimSpace(typed.GetValue(0)))
	default:
		return 0, fmt.Errorf("element type is %T, want Integer String", value)
	}
}

func decimalStringValue(value element.Element) (float64, error) {
	switch typed := value.(type) {
	case *element.DecimalString:
		values, err := typed.GetFloats()
		if err != nil {
			return 0, err
		}
		if len(values) == 0 {
			return 0, fmt.Errorf("value is empty")
		}
		return values[0], nil
	case *element.String:
		if typed.ValueRepresentation().Code() != vr.CodeDS {
			return 0, fmt.Errorf("value representation is %s, want DS", typed.ValueRepresentation().Code())
		}
		return strconv.ParseFloat(strings.TrimSpace(typed.GetValue(0)), 64)
	default:
		return 0, fmt.Errorf("element type is %T, want Decimal String", value)
	}
}

func mergeFunctionalValues(target, source *dataset.Dataset, sequenceTag *tag.Tag, item int) {
	sequence, err := source.GetSequence(sequenceTag)
	if err != nil || item < 0 || item >= sequence.Count() || sequence.GetItem(item) == nil {
		return
	}
	for _, value := range sequence.GetItem(item).Elements() {
		nested, ok := value.(*dataset.Sequence)
		if !ok || nested.Count() == 0 || nested.GetItem(0) == nil {
			continue
		}
		for _, nestedValue := range nested.GetItem(0).Elements() {
			_ = target.AddOrUpdate(nestedValue)
		}
	}
}
