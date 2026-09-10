// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestReadLUTDataReadsSingleEightBitBigEndianCompactWord(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewOtherWord(tag.LUTData, []byte{0xab, 0x00})); err != nil {
		t.Fatalf("add LUT Data: %v", err)
	}

	got, err := readLUTData(ds, tag.LUTData, lutDescriptor{entryCount: 1, bitsPerEntry: 8}, binary.BigEndian)
	if err != nil {
		t.Fatalf("readLUTData() error = %v", err)
	}
	if len(got) != 1 || got[0] != 0xab {
		t.Fatalf("readLUTData() = %v, want [171]", got)
	}
}

func TestReadLUTDataRejectsAmbiguousEightBitWordLength(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewOtherWord(tag.LUTData, []byte{1, 2, 3, 4, 5, 6})); err != nil {
		t.Fatalf("add LUT Data: %v", err)
	}

	if _, err := readLUTData(ds, tag.LUTData, lutDescriptor{entryCount: 4, bitsPerEntry: 8}, binary.LittleEndian); err == nil {
		t.Fatal("readLUTData() accepted a length that is neither compact nor one-word-per-entry")
	}
}

func TestReadLUTDataUsesElementNumericByteOrder(t *testing.T) {
	ds := dataset.New()
	data := element.NewOtherWord(tag.LUTData, []byte{0x12, 0x34})
	element.SetByteOrder(data, binary.BigEndian)
	if err := ds.Add(data); err != nil {
		t.Fatalf("add LUT Data: %v", err)
	}

	got, err := readLUTData(ds, tag.LUTData, lutDescriptor{entryCount: 1, bitsPerEntry: 16}, binary.LittleEndian)
	if err != nil {
		t.Fatalf("readLUTData() error = %v", err)
	}
	if len(got) != 1 || got[0] != 0x1234 {
		t.Fatalf("readLUTData() = %v, want [4660]", got)
	}
}

func TestReadLUTDataFallsBackToDatasetByteOrderForUnspecifiedElementOrder(t *testing.T) {
	ds := dataset.New()
	data := element.NewOtherWord(tag.LUTData, []byte{0x12, 0x34})
	if err := ds.Add(data); err != nil {
		t.Fatalf("add LUT Data: %v", err)
	}

	got, err := readLUTData(ds, tag.LUTData, lutDescriptor{entryCount: 1, bitsPerEntry: 16}, binary.BigEndian)
	if err != nil {
		t.Fatalf("readLUTData() error = %v", err)
	}
	if len(got) != 1 || got[0] != 0x1234 {
		t.Fatalf("readLUTData() = %v, want [4660]", got)
	}
}

func TestReadLUTDataRejectsValuesOutsideDeclaredRange(t *testing.T) {
	tests := []struct {
		name string
		data element.Element
	}{
		{
			name: "US",
			data: element.NewUnsignedShort(tag.LUTData, []uint16{0, 1024}),
		},
		{
			name: "OW",
			data: element.NewOtherWord(tag.LUTData, []byte{0, 0, 0, 4}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			if err := ds.Add(tt.data); err != nil {
				t.Fatalf("add LUT Data: %v", err)
			}

			if _, err := readLUTData(ds, tag.LUTData, lutDescriptor{entryCount: 2, bitsPerEntry: 10}, binary.LittleEndian); err == nil {
				t.Fatal("readLUTData() accepted value 1024 for a 10-bit LUT")
			}
		})
	}
}
