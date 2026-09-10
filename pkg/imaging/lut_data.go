// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"encoding/binary"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

type lutDescriptor struct {
	entryCount       int
	firstMappedValue int
	bitsPerEntry     uint16
}

func readLUTDescriptor(ds *dataset.Dataset, descriptorTag *tag.Tag, signedFirstMapped bool) (lutDescriptor, error) {
	elem, ok := ds.Get(descriptorTag)
	if !ok {
		return lutDescriptor{}, fmt.Errorf("element %s not found", descriptorTag)
	}

	var raw [3]uint16
	var first int
	elementIsSigned := false
	switch value := elem.(type) {
	case *element.UnsignedShort:
		values, err := value.GetValues()
		if err != nil {
			return lutDescriptor{}, err
		}
		if len(values) != 3 {
			return lutDescriptor{}, fmt.Errorf("element %s has %d values, want 3", descriptorTag, len(values))
		}
		copy(raw[:], values)
		first = int(values[1])
	case *element.SignedShort:
		elementIsSigned = true
		values, err := value.GetValues()
		if err != nil {
			return lutDescriptor{}, err
		}
		if len(values) != 3 {
			return lutDescriptor{}, fmt.Errorf("element %s has %d values, want 3", descriptorTag, len(values))
		}
		for index, entry := range values {
			raw[index] = uint16(entry)
		}
		first = int(values[1])
	default:
		return lutDescriptor{}, fmt.Errorf("element %s is not UnsignedShort or SignedShort", descriptorTag)
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
		return lutDescriptor{}, fmt.Errorf("element %s has invalid bits per entry %d", descriptorTag, bitsPerEntry)
	}
	return lutDescriptor{
		entryCount:       entryCount,
		firstMappedValue: first,
		bitsPerEntry:     bitsPerEntry,
	}, nil
}

func readLUTData(ds *dataset.Dataset, dataTag *tag.Tag, descriptor lutDescriptor, byteOrder binary.ByteOrder) ([]uint16, error) {
	elem, ok := ds.Get(dataTag)
	if !ok {
		return nil, fmt.Errorf("element %s not found", dataTag)
	}

	values := make([]uint16, 0, descriptor.entryCount)

	switch value := elem.(type) {
	case *element.OtherByte:
		if descriptor.bitsPerEntry <= 8 {
			entries, err := readCompactLUTBytes(value.GetData(), dataTag, descriptor)
			return validateLUTDataRange(entries, dataTag, descriptor, err)
		}
		entries, err := readLUTWords(value.GetData(), dataTag, descriptor, byteOrder)
		return validateLUTDataRange(entries, dataTag, descriptor, err)
	case *element.OtherWord:
		entries, err := readOtherWordLUTData(value.GetData(), dataTag, descriptor, byteOrder)
		return validateLUTDataRange(entries, dataTag, descriptor, err)
	case *element.UnsignedShort:
		entries, err := value.GetValues()
		if err != nil {
			return nil, err
		}
		if len(entries) != descriptor.entryCount {
			return nil, fmt.Errorf("element %s has %d entries, want %d", dataTag, len(entries), descriptor.entryCount)
		}
		for _, entry := range entries[:descriptor.entryCount] {
			if descriptor.bitsPerEntry <= 8 {
				entry &= 0x00ff
			}
			values = append(values, entry)
		}
	case *element.SignedShort:
		entries, err := value.GetValues()
		if err != nil {
			return nil, err
		}
		if len(entries) != descriptor.entryCount {
			return nil, fmt.Errorf("element %s has %d entries, want %d", dataTag, len(entries), descriptor.entryCount)
		}
		for _, entry := range entries[:descriptor.entryCount] {
			word := uint16(entry)
			if descriptor.bitsPerEntry <= 8 {
				word &= 0x00ff
			}
			values = append(values, word)
		}
	default:
		return nil, fmt.Errorf("unsupported LUT Data element %T", elem)
	}

	return validateLUTDataRange(values, dataTag, descriptor, nil)
}

func validateLUTDataRange(values []uint16, dataTag *tag.Tag, descriptor lutDescriptor, err error) ([]uint16, error) {
	if err != nil {
		return nil, err
	}
	if descriptor.bitsPerEntry <= 8 || descriptor.bitsPerEntry == 16 {
		return values, nil
	}

	maximum := uint16((uint32(1) << descriptor.bitsPerEntry) - 1)
	for index, value := range values {
		if value > maximum {
			return nil, fmt.Errorf("element %s entry %d has value %d, exceeds %d-bit maximum %d",
				dataTag, index, value, descriptor.bitsPerEntry, maximum)
		}
	}
	return values, nil
}

func readOtherWordLUTData(raw []byte, dataTag *tag.Tag, descriptor lutDescriptor, byteOrder binary.ByteOrder) ([]uint16, error) {
	if descriptor.bitsPerEntry > 8 {
		return readLUTWords(raw, dataTag, descriptor, byteOrder)
	}
	compactLength := descriptor.entryCount
	if compactLength%2 != 0 {
		compactLength++
	}
	switch len(raw) {
	case descriptor.entryCount, compactLength:
		return readCompactLUTBytes(raw, dataTag, descriptor)
	case descriptor.entryCount * 2:
		return readLUTWords(raw, dataTag, descriptor, byteOrder)
	default:
		return nil, fmt.Errorf("element %s has %d bytes, want compact length %d or word length %d",
			dataTag, len(raw), compactLength, descriptor.entryCount*2)
	}
}

func readLUTWords(raw []byte, dataTag *tag.Tag, descriptor lutDescriptor, byteOrder binary.ByteOrder) ([]uint16, error) {
	expectedLength := descriptor.entryCount * 2
	if len(raw) != expectedLength {
		return nil, fmt.Errorf("element %s has %d bytes, want %d", dataTag, len(raw), expectedLength)
	}
	values := make([]uint16, descriptor.entryCount)
	for index := range values {
		values[index] = byteOrder.Uint16(raw[index*2:])
		if descriptor.bitsPerEntry <= 8 {
			values[index] &= 0x00ff
		}
	}
	return values, nil
}

func readCompactLUTBytes(raw []byte, dataTag *tag.Tag, descriptor lutDescriptor) ([]uint16, error) {
	unpaddedLength := descriptor.entryCount
	paddedLength := unpaddedLength
	if paddedLength%2 != 0 {
		paddedLength++
	}
	if len(raw) != unpaddedLength && len(raw) != paddedLength {
		return nil, fmt.Errorf("element %s has %d bytes, want %d bytes plus optional even-length padding", dataTag, len(raw), unpaddedLength)
	}
	values := make([]uint16, descriptor.entryCount)
	for index, entry := range raw[:descriptor.entryCount] {
		values[index] = uint16(entry)
	}
	return values, nil
}
