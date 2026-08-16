// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestDeepCloneRecursivelyCopiesElementsAndFragments(t *testing.T) {
	privateOB := tag.New(0x0011, 0x1010)
	privateOW := tag.New(0x0011, 0x1011)
	item := New()
	if err := item.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{testStudyInstanceUID})); err != nil {
		t.Fatalf("add nested item: %v", err)
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.SetOffsetTable([]uint32{0, 12})
	fragments.AddFragment(buffer.NewMemory([]byte{4, 5, 6}))

	source := NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	if err := source.Add(NewSequenceWithItems(tag.SourceImageSequence, []*Dataset{item})); err != nil {
		t.Fatalf("add sequence: %v", err)
	}
	for _, value := range []element.Element{
		element.NewOtherByte(privateOB, []byte{1, 2, 3}),
		element.NewOtherWord(privateOW, []byte{7, 8, 9, 10}),
		fragments,
	} {
		if err := source.Add(value); err != nil {
			t.Fatalf("add %s: %v", value.Tag(), err)
		}
	}
	source.SetAutoValidate(false)

	clone := source.DeepClone()
	if clone.InternalTransferSyntax() != transfer.ExplicitVRBigEndian || clone.AutoValidate() != source.AutoValidate() {
		t.Fatal("DeepClone() did not preserve Dataset settings")
	}
	sequence, _ := clone.GetSequence(tag.SourceImageSequence)
	sequence.GetItem(0).Remove(tag.ReferencedSOPInstanceUID)
	clone.GetOrNil(privateOB).(*element.OtherByte).GetData()[0] = 0xff
	clone.GetOrNil(privateOW).(*element.OtherWord).GetData()[0] = 0xff
	clonedFragments := clone.GetOrNil(tag.PixelData).(*element.OtherByteFragment)
	clonedFragments.OffsetTable()[0] = 99
	clonedFragment, _ := clonedFragments.GetFragment(0)
	clonedFragment.Data()[0] = 0xff

	originalSequence, _ := source.GetSequence(tag.SourceImageSequence)
	if originalSequence.GetItem(0).TryGetString(tag.ReferencedSOPInstanceUID) != testStudyInstanceUID {
		t.Fatal("DeepClone() shared nested Sequence items")
	}
	if source.GetOrNil(privateOB).(*element.OtherByte).GetData()[0] != 1 {
		t.Fatal("DeepClone() shared OtherByte data")
	}
	if source.GetOrNil(privateOW).(*element.OtherWord).GetData()[0] != 7 {
		t.Fatal("DeepClone() shared OtherWord data")
	}
	originalFragments := source.GetOrNil(tag.PixelData).(*element.OtherByteFragment)
	originalFragment, _ := originalFragments.GetFragment(0)
	if originalFragments.OffsetTable()[0] != 0 || originalFragment.Data()[0] != 4 {
		t.Fatal("DeepClone() shared fragment offset or data buffers")
	}
}

func TestDeepClonePreservesConcreteElementTypesAndNumericByteOrder(t *testing.T) {
	privateUS := tag.New(0x0011, 0x1020)
	data := []byte{0x12, 0x34}
	numeric := element.NewUnsignedShortFromBuffer(privateUS, buffer.NewMemory(data))
	element.SetByteOrder(numeric, binary.BigEndian)
	source := New()
	if err := source.Add(numeric); err != nil {
		t.Fatalf("add numeric: %v", err)
	}

	clone := source.DeepClone()
	clonedNumeric, ok := clone.GetOrNil(privateUS).(*element.UnsignedShort)
	if !ok {
		t.Fatalf("DeepClone() element type = %T, want UnsignedShort", clone.GetOrNil(privateUS))
	}
	value, err := clonedNumeric.GetValue(0)
	if err != nil || value != 0x1234 {
		t.Fatalf("DeepClone() numeric value = %#x, error %v", value, err)
	}
	clonedNumeric.Buffer().Data()[0] = 0xff
	if numeric.Buffer().Data()[0] != 0x12 {
		t.Fatal("DeepClone() shared numeric buffer")
	}
}
