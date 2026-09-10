// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dicomlut

import (
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	dicomendian "github.com/cocosip/go-dicom/pkg/io/endian"
)

// ByteOrder returns the byte order of a Dataset's internal transfer syntax.
func ByteOrder(ds *dataset.Dataset) binary.ByteOrder {
	if ds != nil {
		if syntax := ds.InternalTransferSyntax(); syntax != nil && syntax.Endian() == dicomendian.Big {
			return binary.BigEndian
		}
	}
	return binary.LittleEndian
}

// ShortValue reads a signed or unsigned short Dataset value as an int.
func ShortValue(ds *dataset.Dataset, t *tag.Tag) (int, error) {
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

// Descriptor is the normalized three-value DICOM LUT descriptor.
type Descriptor struct {
	EntryCount       int
	FirstMappedValue int
	BitsPerEntry     uint16
}

// ReadDescriptor reads and validates a LUT descriptor from a Dataset.
func ReadDescriptor(ds *dataset.Dataset, descriptorTag *tag.Tag, signedFirstMapped bool) (Descriptor, error) {
	elem, ok := ds.Get(descriptorTag)
	if !ok {
		return Descriptor{}, fmt.Errorf("element %s not found", descriptorTag)
	}

	var raw [3]uint16
	var first int
	elementIsSigned := false
	switch value := elem.(type) {
	case *element.UnsignedShort:
		values, err := value.GetValues()
		if err != nil {
			return Descriptor{}, err
		}
		if len(values) != 3 {
			return Descriptor{}, fmt.Errorf("element %s has %d values, want 3", descriptorTag, len(values))
		}
		copy(raw[:], values)
		first = int(values[1])
	case *element.SignedShort:
		elementIsSigned = true
		values, err := value.GetValues()
		if err != nil {
			return Descriptor{}, err
		}
		if len(values) != 3 {
			return Descriptor{}, fmt.Errorf("element %s has %d values, want 3", descriptorTag, len(values))
		}
		for index, entry := range values {
			raw[index] = uint16(entry)
		}
		first = int(values[1])
	default:
		return Descriptor{}, fmt.Errorf("element %s is not UnsignedShort or SignedShort", descriptorTag)
	}
	if signedFirstMapped || elementIsSigned {
		first = int(int16(raw[1]))
	}

	entryCount := int(raw[0])
	if entryCount == 0 {
		entryCount = 65536
	}
	bitsPerEntry := raw[2]
	if bitsPerEntry < 8 || bitsPerEntry > 16 {
		return Descriptor{}, fmt.Errorf("element %s has invalid bits per entry %d", descriptorTag, bitsPerEntry)
	}
	return Descriptor{
		EntryCount:       entryCount,
		FirstMappedValue: first,
		BitsPerEntry:     bitsPerEntry,
	}, nil
}

// ReadData reads LUT entries according to a normalized descriptor.
func ReadData(ds *dataset.Dataset, dataTag *tag.Tag, descriptor Descriptor, byteOrder binary.ByteOrder) ([]uint16, error) {
	elem, ok := ds.Get(dataTag)
	if !ok {
		return nil, fmt.Errorf("element %s not found", dataTag)
	}

	values := make([]uint16, 0, descriptor.EntryCount)

	switch value := elem.(type) {
	case *element.OtherByte:
		if descriptor.BitsPerEntry <= 8 {
			entries, err := readCompactLUTBytes(value.GetData(), dataTag, descriptor)
			return validateLUTDataRange(entries, dataTag, descriptor, err)
		}
		entries, err := readLUTWords(value.GetData(), dataTag, descriptor, byteOrder)
		return validateLUTDataRange(entries, dataTag, descriptor, err)
	case *element.OtherWord:
		entries, err := readOtherWordLUTData(value.GetData(), dataTag, descriptor, NumericByteOrderOr(value, byteOrder))
		return validateLUTDataRange(entries, dataTag, descriptor, err)
	case *element.UnsignedShort:
		entries, err := value.GetValues()
		if err != nil {
			return nil, err
		}
		if len(entries) != descriptor.EntryCount {
			return nil, fmt.Errorf("element %s has %d entries, want %d", dataTag, len(entries), descriptor.EntryCount)
		}
		for _, entry := range entries[:descriptor.EntryCount] {
			if descriptor.BitsPerEntry <= 8 {
				entry &= 0x00ff
			}
			values = append(values, entry)
		}
	case *element.SignedShort:
		entries, err := value.GetValues()
		if err != nil {
			return nil, err
		}
		if len(entries) != descriptor.EntryCount {
			return nil, fmt.Errorf("element %s has %d entries, want %d", dataTag, len(entries), descriptor.EntryCount)
		}
		for _, entry := range entries[:descriptor.EntryCount] {
			word := uint16(entry)
			if descriptor.BitsPerEntry <= 8 {
				word &= 0x00ff
			}
			values = append(values, word)
		}
	default:
		return nil, fmt.Errorf("unsupported LUT Data element %T", elem)
	}

	return validateLUTDataRange(values, dataTag, descriptor, nil)
}

// NumericByteOrderOr returns an element's byte order or the supplied fallback.
func NumericByteOrderOr(elem element.Element, fallback binary.ByteOrder) binary.ByteOrder {
	if order, ok := element.NumericByteOrder(elem); ok {
		return order
	}
	if fallback == nil {
		return binary.LittleEndian
	}
	return fallback
}

func validateLUTDataRange(values []uint16, dataTag *tag.Tag, descriptor Descriptor, err error) ([]uint16, error) {
	if err != nil {
		return nil, err
	}
	if descriptor.BitsPerEntry <= 8 || descriptor.BitsPerEntry == 16 {
		return values, nil
	}

	maximum := uint16((uint32(1) << descriptor.BitsPerEntry) - 1)
	for index, value := range values {
		if value > maximum {
			return nil, fmt.Errorf("element %s entry %d has value %d, exceeds %d-bit maximum %d",
				dataTag, index, value, descriptor.BitsPerEntry, maximum)
		}
	}
	return values, nil
}

func readOtherWordLUTData(raw []byte, dataTag *tag.Tag, descriptor Descriptor, byteOrder binary.ByteOrder) ([]uint16, error) {
	if descriptor.BitsPerEntry > 8 {
		return readLUTWords(raw, dataTag, descriptor, byteOrder)
	}
	compactLength := descriptor.EntryCount
	if compactLength%2 != 0 {
		compactLength++
	}
	switch len(raw) {
	case descriptor.EntryCount, compactLength:
		return readCompactLUTBytes(raw, dataTag, descriptor)
	case descriptor.EntryCount * 2:
		return readLUTWords(raw, dataTag, descriptor, byteOrder)
	default:
		return nil, fmt.Errorf("element %s has %d bytes, want compact length %d or word length %d",
			dataTag, len(raw), compactLength, descriptor.EntryCount*2)
	}
}

func readLUTWords(raw []byte, dataTag *tag.Tag, descriptor Descriptor, byteOrder binary.ByteOrder) ([]uint16, error) {
	expectedLength := descriptor.EntryCount * 2
	if len(raw) != expectedLength {
		return nil, fmt.Errorf("element %s has %d bytes, want %d", dataTag, len(raw), expectedLength)
	}
	values := make([]uint16, descriptor.EntryCount)
	for index := range values {
		values[index] = byteOrder.Uint16(raw[index*2:])
		if descriptor.BitsPerEntry <= 8 {
			values[index] &= 0x00ff
		}
	}
	return values, nil
}

func readCompactLUTBytes(raw []byte, dataTag *tag.Tag, descriptor Descriptor) ([]uint16, error) {
	unpaddedLength := descriptor.EntryCount
	paddedLength := unpaddedLength
	if paddedLength%2 != 0 {
		paddedLength++
	}
	if len(raw) != unpaddedLength && len(raw) != paddedLength {
		return nil, fmt.Errorf("element %s has %d bytes, want %d bytes plus optional even-length padding", dataTag, len(raw), unpaddedLength)
	}
	values := make([]uint16, descriptor.EntryCount)
	for index, entry := range raw[:descriptor.EntryCount] {
		values[index] = uint16(entry)
	}
	return values, nil
}
