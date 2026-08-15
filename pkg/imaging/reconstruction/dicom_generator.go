// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// DicomGenerator creates classic single-frame CT/MR derived instances.
type DicomGenerator struct {
	volume     *VolumeData
	clock      func() time.Time
	uidFactory func() string
}

type generatorConfig struct {
	clock      func() time.Time
	uidFactory func() string
}

// GeneratorOption configures deterministic derived-instance generation.
type GeneratorOption func(*generatorConfig)

// WithGeneratorClock sets the clock used for derived instance dates and times.
func WithGeneratorClock(clock func() time.Time) GeneratorOption {
	return func(config *generatorConfig) { config.clock = clock }
}

// WithGeneratorUIDFactory sets the factory used for Series and SOP Instance UIDs.
func WithGeneratorUIDFactory(factory func() string) GeneratorOption {
	return func(config *generatorConfig) { config.uidFactory = factory }
}

// NewDicomGenerator validates the source modality and configures generation.
func NewDicomGenerator(volume *VolumeData, options ...GeneratorOption) (*DicomGenerator, error) {
	if volume == nil || len(volume.images) < 2 || volume.commonData == nil {
		return nil, fmt.Errorf("DICOM generator requires a complete volume")
	}
	if sourceModality(volume.images[0].sourceSOPClassUID) == "" {
		return nil, fmt.Errorf("unsupported source SOP Class UID %q", volume.images[0].sourceSOPClassUID)
	}
	config := generatorConfig{
		clock:      time.Now,
		uidFactory: func() string { return uid.GenerateDerivedFromUUID().UID() },
	}
	for _, option := range options {
		if option != nil {
			option(&config)
		}
	}
	if config.clock == nil || config.uidFactory == nil {
		return nil, fmt.Errorf("DICOM generator clock and UID factory must not be nil")
	}
	return &DicomGenerator{volume: volume, clock: config.clock, uidFactory: config.uidFactory}, nil
}

// Stream generates and emits one Dataset at a time.
func (generator *DicomGenerator) Stream(
	ctx context.Context,
	stack *Stack,
	seriesDescription string,
	cutOptions CutOptions,
	consume func(index int, output *dataset.Dataset) error,
) error {
	if generator == nil || generator.volume == nil {
		return fmt.Errorf("DICOM generator is nil")
	}
	if stack == nil || stack.volume != generator.volume {
		return fmt.Errorf("stack does not belong to the generator volume")
	}
	if consume == nil {
		return fmt.Errorf("DICOM stream consumer is nil")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	seriesUID := generator.uidFactory()
	if seriesUID == "" {
		return fmt.Errorf("UID factory returned an empty Series Instance UID")
	}
	now := generator.clock()
	sourceSequence, err := generator.sourceImageSequence()
	if err != nil {
		return err
	}
	for index := 0; index < stack.Len(); index++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		slice, err := stack.Materialize(ctx, index, cutOptions)
		if err != nil {
			return fmt.Errorf("materialize stack slice %d: %w", index, err)
		}
		sopUID := generator.uidFactory()
		if sopUID == "" {
			return fmt.Errorf("UID factory returned an empty SOP Instance UID for slice %d", index)
		}
		output, err := generator.datasetForSlice(slice, stack, index, seriesUID, sopUID, seriesDescription, now, sourceSequence)
		if err != nil {
			return fmt.Errorf("generate stack slice %d: %w", index, err)
		}
		if err := consume(index, output); err != nil {
			return fmt.Errorf("consume stack slice %d: %w", index, err)
		}
	}
	return nil
}

// Generate is the collect-all convenience API over Stream.
func (generator *DicomGenerator) Generate(ctx context.Context, stack *Stack, seriesDescription string, options CutOptions) ([]*dataset.Dataset, error) {
	outputs := make([]*dataset.Dataset, 0, stack.Len())
	err := generator.Stream(ctx, stack, seriesDescription, options, func(_ int, output *dataset.Dataset) error {
		outputs = append(outputs, output)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return outputs, nil
}

func (generator *DicomGenerator) datasetForSlice(
	slice *Slice,
	stack *Stack,
	index int,
	seriesUID, sopUID, description string,
	now time.Time,
	sourceSequence *dataset.Sequence,
) (*dataset.Dataset, error) {
	pixelBytes, representation, padding, minimum, maximum, err := encodeSlicePixels(slice)
	if err != nil {
		return nil, err
	}
	output := generator.volume.commonData.DeepClone()
	output.SetInternalTransferSyntax(transfer.ExplicitVRLittleEndian)
	removeEnhancedAndStalePixelTags(output)

	sopClass := uid.CTImageStorage.UID()
	modality := sourceModality(generator.volume.images[0].sourceSOPClassUID)
	if modality == "MR" {
		sopClass = uid.MRImageStorage.UID()
	}
	photometric := generator.volume.images[0].pixelData.Info.PhotometricInterpretation.Value
	date := now.Format("20060102")
	timeValue := now.Format("150405.000000")
	decimalDefinitions := []struct {
		target *tag.Tag
		values []float64
	}{
		{tag.RescaleSlope, []float64{1}},
		{tag.RescaleIntercept, []float64{0}},
		{tag.WindowCenter, []float64{(minimum + maximum) / 2}},
		{tag.WindowWidth, []float64{math.Max(1, maximum-minimum+1)}},
		{tag.ImagePositionPatient, []float64{slice.Spec.TopLeft.X, slice.Spec.TopLeft.Y, slice.Spec.TopLeft.Z}},
		{tag.ImageOrientationPatient, []float64{
			slice.Spec.RowDirection.X, slice.Spec.RowDirection.Y, slice.Spec.RowDirection.Z,
			slice.Spec.ColumnDirection.X, slice.Spec.ColumnDirection.Y, slice.Spec.ColumnDirection.Z,
		}},
		{tag.PixelSpacing, []float64{slice.Spec.PixelSpacingRows, slice.Spec.PixelSpacingColumns}},
		{tag.SliceThickness, []float64{stack.SliceDistance}},
		{tag.SpacingBetweenSlices, []float64{stack.SliceDistance}},
		{tag.SliceLocation, []float64{sliceLocation(slice.Spec)}},
	}
	decimalElements := make([]element.Element, len(decimalDefinitions))
	for index, definition := range decimalDefinitions {
		decimal, err := newGeneratorDecimalString(definition.target, definition.values)
		if err != nil {
			return nil, err
		}
		decimalElements[index] = decimal
	}
	values := []element.Element{
		element.NewString(tag.SOPClassUID, vr.UI, []string{sopClass}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopUID}),
		element.NewString(tag.SeriesInstanceUID, vr.UI, []string{seriesUID}),
		element.NewString(tag.Modality, vr.CS, []string{modality}),
		element.NewString(tag.ImageType, vr.CS, []string{"DERIVED", "SECONDARY", "MPR"}),
		element.NewString(tag.SeriesDescription, vr.LO, []string{description}),
		element.NewString(tag.DerivationDescription, vr.ST, []string{"Multi-planar reformation from source image volume"}),
		element.NewString(tag.InstanceCreationDate, vr.DA, []string{date}),
		element.NewString(tag.InstanceCreationTime, vr.TM, []string{timeValue}),
		element.NewString(tag.ContentDate, vr.DA, []string{date}),
		element.NewString(tag.ContentTime, vr.TM, []string{timeValue}),
		element.NewIntegerStringFromInt(tag.InstanceNumber, []int{index + 1}),
		element.NewUnsignedShort(tag.Rows, []uint16{uint16(slice.Spec.Rows)}),       // #nosec G115 -- validated below
		element.NewUnsignedShort(tag.Columns, []uint16{uint16(slice.Spec.Columns)}), // #nosec G115 -- validated below
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometric}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{representation}),
		sourceSequence.DeepClone(),
		element.NewOtherWord(tag.PixelData, pixelBytes),
	}
	values = append(values, decimalElements...)
	if slice.Spec.Rows > math.MaxUint16 || slice.Spec.Columns > math.MaxUint16 {
		return nil, fmt.Errorf("output dimensions %dx%d exceed DICOM US range", slice.Spec.Rows, slice.Spec.Columns)
	}
	if padding != nil {
		values = append(values, padding)
	}
	for _, value := range values {
		if err := output.AddOrUpdate(value); err != nil {
			return nil, fmt.Errorf("set %s: %w", value.Tag(), err)
		}
	}
	return output, nil
}

func newGeneratorDecimalString(target *tag.Tag, values []float64) (*element.DecimalString, error) {
	formatted := make([]string, len(values))
	for index, value := range values {
		if math.IsNaN(value) || math.IsInf(value, 0) {
			return nil, fmt.Errorf("format %s value %d: value is not finite", target, index)
		}
		for precision := 15; precision >= 1; precision-- {
			candidate := strconv.FormatFloat(value, 'g', precision, 64)
			if len(candidate) <= 16 {
				formatted[index] = candidate
				break
			}
		}
		if formatted[index] == "" {
			return nil, fmt.Errorf("format %s value %d: cannot represent %.17g as DICOM DS", target, index, value)
		}
	}
	return element.NewDecimalString(target, formatted), nil
}

func encodeSlicePixels(slice *Slice) ([]byte, uint16, element.Element, float64, float64, error) {
	minimum, maximum, err := slice.MinMax()
	if err != nil {
		return nil, 0, nil, 0, 0, err
	}
	hasInvalid := false
	for _, valid := range slice.Valid {
		hasInvalid = hasInvalid || !valid
	}
	representation := uint16(0)
	if minimum < 0 {
		representation = 1
	}
	data := make([]byte, len(slice.Values)*2)
	var padding element.Element
	if representation == 1 {
		lower := -32768.0
		if hasInvalid {
			lower = -32767
			padding = element.NewSignedShort(tag.PixelPaddingValue, []int16{-32768})
		}
		if minimum < lower || maximum > 32767 {
			return nil, 0, nil, 0, 0, fmt.Errorf("signed output range %.6g..%.6g exceeds 16-bit representation", minimum, maximum)
		}
		for index, value := range slice.Values {
			encoded := int16(-32768)
			if slice.Valid[index] {
				encoded = int16(math.Round(value))
			}
			binary.LittleEndian.PutUint16(data[index*2:], uint16(encoded)) // #nosec G115 -- bit-preserving conversion
		}
	} else {
		upper := 65535.0
		if hasInvalid {
			upper = 65534
			padding = element.NewUnsignedShort(tag.PixelPaddingValue, []uint16{65535})
		}
		if minimum < 0 || maximum > upper {
			return nil, 0, nil, 0, 0, fmt.Errorf("unsigned output range %.6g..%.6g exceeds 16-bit representation", minimum, maximum)
		}
		for index, value := range slice.Values {
			encoded := uint16(65535)
			if slice.Valid[index] {
				encoded = uint16(math.Round(value))
			}
			binary.LittleEndian.PutUint16(data[index*2:], encoded)
		}
	}
	return data, representation, padding, minimum, maximum, nil
}

func (generator *DicomGenerator) sourceImageSequence() (*dataset.Sequence, error) {
	type sourceReference struct {
		class, instance string
		frames          []int
	}
	references := make([]sourceReference, 0, len(generator.volume.images))
	indexes := make(map[string]int)
	for _, image := range generator.volume.images {
		if image.sourceSOPClassUID == "" || image.sourceSOPInstanceUID == "" {
			return nil, fmt.Errorf("source frame is missing SOP identity")
		}
		key := image.sourceSOPClassUID + "\x00" + image.sourceSOPInstanceUID
		index, exists := indexes[key]
		if !exists {
			index = len(references)
			indexes[key] = index
			references = append(references, sourceReference{class: image.sourceSOPClassUID, instance: image.sourceSOPInstanceUID})
		}
		if image.sourceSOPClassUID == uid.EnhancedCTImageStorage.UID() || image.sourceSOPClassUID == uid.EnhancedMRImageStorage.UID() {
			references[index].frames = append(references[index].frames, image.frameIndex+1)
		}
	}
	items := make([]*dataset.Dataset, len(references))
	for index, reference := range references {
		item := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
		_ = item.Add(element.NewString(tag.ReferencedSOPClassUID, vr.UI, []string{reference.class}))
		_ = item.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{reference.instance}))
		if len(reference.frames) > 0 {
			_ = item.Add(element.NewIntegerStringFromInt(tag.ReferencedFrameNumber, reference.frames))
		}
		items[index] = item
	}
	return dataset.NewSequenceWithItems(tag.SourceImageSequence, items), nil
}

func removeEnhancedAndStalePixelTags(output *dataset.Dataset) {
	output.RemoveAll(
		tag.NumberOfFrames,
		tag.SharedFunctionalGroupsSequence,
		tag.PerFrameFunctionalGroupsSequence,
		tag.FrameIncrementPointer,
		tag.FrameType,
		tag.CTImageFrameTypeSequence,
		tag.MRImageFrameTypeSequence,
		tag.DimensionOrganizationSequence,
		tag.DimensionIndexSequence,
		tag.DimensionOrganizationUID,
		tag.DimensionIndexValues,
		tag.DimensionIndexPointer,
		tag.DimensionIndexPrivateCreator,
		tag.ConcatenationUID,
		tag.InConcatenationNumber,
		tag.InConcatenationTotalNumber,
		tag.ConcatenationFrameOffsetNumber,
		tag.SOPInstanceUIDOfConcatenationSource,
		tag.StackID,
		tag.InStackPositionNumber,
		tag.TemporalPositionIndex,
		tag.RepresentativeFrameNumber,
		tag.PlanarConfiguration,
		tag.SmallestImagePixelValue,
		tag.LargestImagePixelValue,
		tag.ExtendedOffsetTable,
		tag.ExtendedOffsetTableLengths,
		tag.FloatPixelData,
		tag.DoubleFloatPixelData,
		tag.PixelDataProviderURL,
		tag.PixelPaddingValue,
		tag.PixelPaddingRangeLimit,
		tag.PixelData,
	)
}

func sliceLocation(spec CutSpec) float64 {
	normal := spec.RowDirection.Cross(spec.ColumnDirection)
	return dot3(spec.TopLeft, normal)
}
