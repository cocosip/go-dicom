// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	dicomendian "github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
)

func datasetByteOrder(ds *dataset.Dataset) binary.ByteOrder {
	if ds != nil {
		if syntax := ds.InternalTransferSyntax(); syntax != nil && syntax.Endian() == dicomendian.Big {
			return binary.BigEndian
		}
	}
	return binary.LittleEndian
}

func datasetShortValue(ds *dataset.Dataset, t *tag.Tag) (int, error) {
	elem, ok := ds.Get(t)
	if !ok {
		return 0, fmt.Errorf("element %s not found", t)
	}
	switch value := elem.(type) {
	case *element.SignedShort:
		entry, err := value.GetValue(0)
		return int(entry), err
	case *element.UnsignedShort:
		entry, err := value.GetValue(0)
		return int(entry), err
	default:
		return 0, fmt.Errorf("element %s is not SignedShort or UnsignedShort", t)
	}
}

func imagePixelValueRange(ds *dataset.Dataset) (float64, float64, error) {
	minimum, err := datasetShortValue(ds, tag.SmallestImagePixelValue)
	if err != nil {
		return 0, 0, err
	}
	maximum, err := datasetShortValue(ds, tag.LargestImagePixelValue)
	if err != nil {
		return 0, 0, err
	}
	if minimum >= maximum {
		return 0, 0, fmt.Errorf("invalid image pixel value range %d..%d", minimum, maximum)
	}
	return float64(minimum), float64(maximum), nil
}

func imageVOILUT(ds *dataset.Dataset, signed bool) (render.LUT, error) {
	return imageVOILUTAt(ds, signed, 0)
}

func imageVOILUTAt(ds *dataset.Dataset, signed bool, index int) (render.LUT, error) {
	sequence, err := ds.GetSequence(tag.VOILUTSequence)
	if err != nil || sequence.Count() == 0 {
		return nil, fmt.Errorf("VOI LUT Sequence is missing or empty")
	}
	if index < 0 || index >= sequence.Count() {
		return nil, fmt.Errorf("VOI LUT Sequence item index %d is out of range [0, %d)", index, sequence.Count())
	}
	item := sequence.GetItem(index)
	if item == nil {
		return nil, fmt.Errorf("VOI LUT Sequence item %d is nil", index)
	}
	descriptor, err := readLUTDescriptor(item, tag.LUTDescriptor, signed)
	if err != nil {
		return nil, fmt.Errorf("read VOI LUT descriptor: %w", err)
	}
	if err := validateImageVOILUTDescriptor(descriptor); err != nil {
		return nil, fmt.Errorf("read VOI LUT descriptor: %w", err)
	}
	values, err := readLUTData(item, tag.LUTData, descriptor, datasetByteOrder(ds))
	if err != nil {
		return nil, fmt.Errorf("read VOI LUT Data: %w", err)
	}
	table := &voiTableLUT{values: values, first: descriptor.firstMappedValue}
	maximumOutput := float64(uint32(1)<<descriptor.bitsPerEntry - 1)
	return scaledVOITable(table, 0, maximumOutput), nil
}

func imageVOILUTFromAt(primary, fallback *dataset.Dataset, signed bool, index int) (render.LUT, error) {
	if primary != nil && primary.Contains(tag.VOILUTSequence) {
		return imageVOILUTAt(primary, signed, index)
	}
	if fallback != nil && fallback.Contains(tag.VOILUTSequence) {
		return imageVOILUTAt(fallback, signed, index)
	}
	return nil, fmt.Errorf("VOI LUT Sequence is missing")
}

func imageModalityLUT(ds *dataset.Dataset, signed bool) (render.ModalityLUT, error) {
	sequence, err := ds.GetSequence(tag.ModalityLUTSequence)
	if err != nil {
		return nil, fmt.Errorf("modality LUT Sequence is missing: %w", err)
	}
	if sequence.Count() != 1 || sequence.GetItem(0) == nil {
		return nil, fmt.Errorf("modality LUT Sequence must contain exactly one item, got %d", sequence.Count())
	}
	item := sequence.GetItem(0)
	descriptor, err := readLUTDescriptor(item, tag.LUTDescriptor, signed)
	if err != nil {
		return nil, fmt.Errorf("read Modality LUT descriptor: %w", err)
	}
	if descriptor.bitsPerEntry != 8 && descriptor.bitsPerEntry != 16 {
		return nil, fmt.Errorf("read Modality LUT descriptor: bits per entry must be 8 or 16, got %d", descriptor.bitsPerEntry)
	}
	entries, err := readLUTData(item, tag.LUTData, descriptor, datasetByteOrder(ds))
	if err != nil {
		return nil, fmt.Errorf("read Modality LUT Data: %w", err)
	}
	values := make([]float64, len(entries))
	for index, entry := range entries {
		values[index] = float64(entry)
	}
	return render.NewModalitySequenceLUT(values, descriptor.firstMappedValue, signed), nil
}

func imageWindowPairAt(primary, fallback *dataset.Dataset, index int) (float64, float64, error) {
	source := datasetWithCompletePair(primary, fallback, tag.WindowCenter, tag.WindowWidth)
	if source == nil {
		source = datasetWithAnyTag(primary, fallback, tag.WindowCenter, tag.WindowWidth)
	}
	if source == nil {
		return 0, 0, fmt.Errorf("window center and width are missing")
	}
	centerElement, centerOK := source.Get(tag.WindowCenter)
	widthElement, widthOK := source.Get(tag.WindowWidth)
	if !centerOK || !widthOK {
		return 0, 0, fmt.Errorf("window center and width must both be present")
	}
	centers, centerOK := centerElement.(*element.DecimalString)
	widths, widthOK := widthElement.(*element.DecimalString)
	if !centerOK || !widthOK {
		return 0, 0, fmt.Errorf("window center and width must use Decimal String VR")
	}
	if centers.Count() != widths.Count() {
		return 0, 0, fmt.Errorf("window center and width value counts differ: %d and %d", centers.Count(), widths.Count())
	}
	if index < 0 || index >= centers.Count() {
		return 0, 0, fmt.Errorf("window index %d is out of range [0, %d)", index, centers.Count())
	}
	center, err := centers.GetFloat(index)
	if err != nil {
		return 0, 0, err
	}
	width, err := widths.GetFloat(index)
	if err != nil {
		return 0, 0, err
	}
	return center, width, nil
}

func datasetWithCompletePair(primary, fallback *dataset.Dataset, first, second *tag.Tag) *dataset.Dataset {
	for _, ds := range []*dataset.Dataset{primary, fallback} {
		if ds != nil && ds.Contains(first) && ds.Contains(second) {
			return ds
		}
	}
	return nil
}

func datasetWithAnyTag(primary, fallback *dataset.Dataset, tags ...*tag.Tag) *dataset.Dataset {
	for _, ds := range []*dataset.Dataset{primary, fallback} {
		if ds == nil {
			continue
		}
		for _, t := range tags {
			if ds.Contains(t) {
				return ds
			}
		}
	}
	return nil
}

func imageModalityTransform(primary, fallback *dataset.Dataset, pixelSigned bool, minInput, maxInput float64) (
	modalityLUT render.ModalityLUT,
	slope float64,
	intercept float64,
	voiDescriptorSigned bool,
	err error,
) {
	slope, intercept = 1, 0
	modalitySource := datasetWithAnyTag(primary, fallback, tag.ModalityLUTSequence)
	rescaleSource := datasetWithAnyTag(primary, fallback, tag.RescaleSlope, tag.RescaleIntercept)
	if modalitySource != nil && rescaleSource != nil {
		return nil, 0, 0, false, fmt.Errorf("modality LUT sequence cannot coexist with rescale slope/intercept")
	}
	if modalitySource != nil {
		modalityLUT, err = imageModalityLUT(modalitySource, pixelSigned)
		if err != nil {
			return nil, 0, 0, false, err
		}
		return modalityLUT, slope, intercept, false, nil
	}
	if rescaleSource == nil {
		return nil, slope, intercept, pixelSigned, nil
	}
	if !rescaleSource.Contains(tag.RescaleSlope) || !rescaleSource.Contains(tag.RescaleIntercept) {
		return nil, 0, 0, false, fmt.Errorf("rescale slope and intercept must both be present")
	}
	slope, err = imageSingleDecimal(rescaleSource, tag.RescaleSlope)
	if err != nil {
		return nil, 0, 0, false, err
	}
	if slope == 0 || math.IsNaN(slope) || math.IsInf(slope, 0) {
		return nil, 0, 0, false, fmt.Errorf("rescale slope must be finite and non-zero")
	}
	intercept, err = imageSingleDecimal(rescaleSource, tag.RescaleIntercept)
	if err != nil {
		return nil, 0, 0, false, err
	}
	if math.IsNaN(intercept) || math.IsInf(intercept, 0) {
		return nil, 0, 0, false, fmt.Errorf("rescale intercept must be finite")
	}
	minimum := math.Min(minInput*slope+intercept, maxInput*slope+intercept)
	return nil, slope, intercept, minimum < 0, nil
}

func imageSingleDecimal(ds *dataset.Dataset, t *tag.Tag) (float64, error) {
	elem, ok := ds.Get(t)
	if !ok {
		return 0, fmt.Errorf("element %s not found", t)
	}
	value, ok := elem.(*element.DecimalString)
	if !ok || value.Count() != 1 {
		return 0, fmt.Errorf("element %s must contain exactly one Decimal String value", t)
	}
	return value.GetFloat(0)
}

func imageStringFrom(primary, fallback *dataset.Dataset, t *tag.Tag) (string, bool) {
	for _, source := range []*dataset.Dataset{primary, fallback} {
		if value, ok := source.GetString(t); ok {
			value = strings.TrimSpace(value)
			if value != "" {
				return value, true
			}
		}
	}
	return "", false
}

// imageFunctionalGroupValues flattens the first item of each shared and
// per-frame functional-group macro. Per-frame values replace shared values.
func imageFunctionalGroupValues(ds *dataset.Dataset, frame int) *dataset.Dataset {
	values := dataset.New()
	values.SetInternalTransferSyntax(ds.InternalTransferSyntax())
	mergeFunctionalGroupValues(values, ds, tag.SharedFunctionalGroupsSequence, 0)
	mergeFunctionalGroupValues(values, ds, tag.PerFrameFunctionalGroupsSequence, frame)
	return values
}

func mergeFunctionalGroupValues(values, ds *dataset.Dataset, sequenceTag *tag.Tag, itemIndex int) {
	sequence, err := ds.GetSequence(sequenceTag)
	if err != nil || itemIndex < 0 || itemIndex >= sequence.Count() {
		return
	}
	item := sequence.GetItem(itemIndex)
	if item == nil {
		return
	}
	for _, elem := range item.Elements() {
		nested, ok := elem.(*dataset.Sequence)
		if !ok || nested.Count() == 0 || nested.GetItem(0) == nil {
			continue
		}
		for _, value := range nested.GetItem(0).Elements() {
			_ = values.AddOrUpdate(value)
		}
	}
}
