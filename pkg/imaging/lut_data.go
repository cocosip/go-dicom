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
	appendWords := func(raw []byte) error {
		if len(raw) < descriptor.entryCount*2 {
			return fmt.Errorf("element %s has %d bytes, want at least %d", dataTag, len(raw), descriptor.entryCount*2)
		}
		for index := 0; index < descriptor.entryCount; index++ {
			entry := byteOrder.Uint16(raw[index*2:])
			if descriptor.bitsPerEntry <= 8 {
				entry &= 0x00ff
			}
			values = append(values, entry)
		}
		return nil
	}

	switch value := elem.(type) {
	case *element.OtherByte:
		raw := value.GetData()
		if descriptor.bitsPerEntry <= 8 {
			if len(raw) < descriptor.entryCount {
				return nil, fmt.Errorf("element %s has %d bytes, want at least %d", dataTag, len(raw), descriptor.entryCount)
			}
			for _, entry := range raw[:descriptor.entryCount] {
				values = append(values, uint16(entry))
			}
		} else if err := appendWords(raw); err != nil {
			return nil, err
		}
	case *element.OtherWord:
		raw := value.GetData()
		if descriptor.bitsPerEntry <= 8 && len(raw) < descriptor.entryCount*2 {
			if len(raw) < descriptor.entryCount {
				return nil, fmt.Errorf("element %s has %d bytes, want at least %d", dataTag, len(raw), descriptor.entryCount)
			}
			for _, entry := range raw[:descriptor.entryCount] {
				values = append(values, uint16(entry))
			}
		} else if err := appendWords(raw); err != nil {
			return nil, err
		}
	case *element.UnsignedShort:
		entries, err := value.GetValues()
		if err != nil {
			return nil, err
		}
		if len(entries) < descriptor.entryCount {
			return nil, fmt.Errorf("element %s has %d entries, want at least %d", dataTag, len(entries), descriptor.entryCount)
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
		if len(entries) < descriptor.entryCount {
			return nil, fmt.Errorf("element %s has %d entries, want at least %d", dataTag, len(entries), descriptor.entryCount)
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

	return values, nil
}
