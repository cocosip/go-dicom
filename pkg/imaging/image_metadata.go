// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"encoding/binary"
	"fmt"

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

func imageDecimal(ds *dataset.Dataset, t *tag.Tag) (float64, error) {
	elem, ok := ds.Get(t)
	if !ok {
		return 0, fmt.Errorf("element %s not found", t)
	}
	value, ok := elem.(*element.DecimalString)
	if !ok {
		return 0, fmt.Errorf("element %s is not DecimalString", t)
	}
	return value.GetFloat(0)
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
	sequence, err := ds.GetSequence(tag.VOILUTSequence)
	if err != nil || sequence.Count() == 0 || sequence.GetItem(0) == nil {
		return nil, fmt.Errorf("VOI LUT Sequence is missing or empty")
	}
	item := sequence.GetItem(0)
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

func imageVOILUTFrom(primary, fallback *dataset.Dataset, signed bool) (render.LUT, error) {
	if value, err := imageVOILUT(primary, signed); err == nil {
		return value, nil
	}
	return imageVOILUT(fallback, signed)
}

func imageModalityLUT(ds *dataset.Dataset, signed bool) (render.ModalityLUT, error) {
	sequence, err := ds.GetSequence(tag.ModalityLUTSequence)
	if err != nil || sequence.Count() == 0 || sequence.GetItem(0) == nil {
		return nil, fmt.Errorf("modality LUT Sequence is missing or empty")
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

func imageDecimalFrom(primary, fallback *dataset.Dataset, t *tag.Tag) (float64, error) {
	if value, err := imageDecimal(primary, t); err == nil {
		return value, nil
	}
	return imageDecimal(fallback, t)
}

func imageWindowPair(primary, fallback *dataset.Dataset) (float64, float64, error) {
	if center, err := imageDecimal(primary, tag.WindowCenter); err == nil {
		if width, err := imageDecimal(primary, tag.WindowWidth); err == nil && width > 0 {
			return center, width, nil
		}
	}
	center, err := imageDecimal(fallback, tag.WindowCenter)
	if err != nil {
		return 0, 0, err
	}
	width, err := imageDecimal(fallback, tag.WindowWidth)
	if err != nil || width <= 0 {
		return 0, 0, fmt.Errorf("functional-group window is missing or invalid")
	}
	return center, width, nil
}

func imageStringFrom(primary, fallback *dataset.Dataset, t *tag.Tag) (string, bool) {
	if value, ok := primary.GetString(t); ok {
		return value, true
	}
	return fallback.GetString(t)
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
