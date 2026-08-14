// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"bytes"
	"encoding/binary"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestParseImplicitVRUsesLoadedPrivateDictionary(t *testing.T) {
	const dictionaryXML = `<dictionary creator="ACME">
  <tag group="0019" element="1000" keyword="FieldOfView" vr="DS" vm="1">Field Of View</tag>
</dictionary>`
	dictionary, err := dict.NewFromXML(strings.NewReader(dictionaryXML))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}

	var raw bytes.Buffer
	writeImplicitElement(&raw, 0x0019, 0x0011, []byte("ACME"))
	writeImplicitElement(&raw, 0x0019, 0x1100, []byte("42"))

	result, err := Parse(
		bytes.NewReader(raw.Bytes()),
		WithAssumedTransferSyntax(transfer.ImplicitVRLittleEndian),
		WithDictionary(dictionary),
	)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	creatorTag := tag.New(0x0019, 0x0011)
	creatorElement, ok := result.Dataset.Get(creatorTag)
	if !ok {
		t.Fatal("private creator element not found")
	}
	if creatorElement.ValueRepresentation() != vr.LO {
		t.Errorf("private creator VR = %s, want LO", creatorElement.ValueRepresentation())
	}

	privateTag := tag.NewWithPrivateCreator(
		0x0019,
		0x1100,
		dictionary.GetPrivateCreator("ACME"),
	)
	privateElement, ok := result.Dataset.Get(privateTag)
	if !ok {
		t.Fatal("private element not found through creator-aware tag")
	}
	if privateElement.ValueRepresentation() != vr.DS {
		t.Errorf("private element VR = %s, want DS", privateElement.ValueRepresentation())
	}
	if got, ok := result.Dataset.GetString(privateTag); !ok || got != "42" {
		t.Errorf("private value = %q, %v, want 42, true", got, ok)
	}
}

func TestParseImplicitVRKeepsPrivateCreatorsScopedToSequenceItems(t *testing.T) {
	const dictionaryXML = `<dictionaries>
  <dictionary>
    <tag group="0008" element="1110" keyword="ReferencedStudySequence" vr="SQ" vm="1">Referenced Study Sequence</tag>
  </dictionary>
  <dictionary creator="ACME">
    <tag group="0019" element="1000" keyword="AcmeValue" vr="DS" vm="1">ACME Value</tag>
  </dictionary>
  <dictionary creator="BETA">
    <tag group="0019" element="1000" keyword="BetaValue" vr="LO" vm="1">BETA Value</tag>
  </dictionary>
</dictionaries>`
	dictionary, err := dict.NewFromXML(strings.NewReader(dictionaryXML))
	if err != nil {
		t.Fatalf("NewFromXML() error = %v", err)
	}

	var firstItem bytes.Buffer
	writeImplicitElement(&firstItem, 0x0019, 0x0011, []byte("ACME"))
	writeImplicitElement(&firstItem, 0x0019, 0x1100, []byte("42"))

	var secondItem bytes.Buffer
	writeImplicitElement(&secondItem, 0x0019, 0x0011, []byte("BETA"))
	writeImplicitElement(&secondItem, 0x0019, 0x1100, []byte("TEXT"))
	var unreservedItem bytes.Buffer
	writeImplicitElement(&unreservedItem, 0x0019, 0x1100, []byte("RAW "))

	var sequenceValue bytes.Buffer
	writeImplicitItem(&sequenceValue, firstItem.Bytes())
	writeImplicitItem(&sequenceValue, secondItem.Bytes())
	writeImplicitItem(&sequenceValue, unreservedItem.Bytes())

	var raw bytes.Buffer
	writeImplicitElement(
		&raw,
		tag.ReferencedStudySequence.Group(),
		tag.ReferencedStudySequence.Element(),
		sequenceValue.Bytes(),
	)

	result, err := Parse(
		bytes.NewReader(raw.Bytes()),
		WithAssumedTransferSyntax(transfer.ImplicitVRLittleEndian),
		WithDictionary(dictionary),
	)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	sequence, err := result.Dataset.GetSequence(tag.ReferencedStudySequence)
	if err != nil {
		t.Fatalf("GetSequence() error = %v", err)
	}
	if sequence.Count() != 3 {
		t.Fatalf("sequence item count = %d, want 3", sequence.Count())
	}

	tests := []struct {
		name    string
		item    int
		creator string
		wantVR  *vr.VR
		want    string
	}{
		{name: "first item", item: 0, creator: "ACME", wantVR: vr.DS, want: "42"},
		{name: "second item", item: 1, creator: "BETA", wantVR: vr.LO, want: "TEXT"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			privateTag := tag.NewWithPrivateCreator(
				0x0019,
				0x1100,
				dictionary.GetPrivateCreator(test.creator),
			)
			privateElement, ok := sequence.GetItem(test.item).Get(privateTag)
			if !ok {
				t.Fatalf("private element for creator %q not found", test.creator)
			}
			if privateElement.ValueRepresentation() != test.wantVR {
				t.Errorf("private element VR = %s, want %s", privateElement.ValueRepresentation(), test.wantVR)
			}
			if got, ok := sequence.GetItem(test.item).GetString(privateTag); !ok || got != test.want {
				t.Errorf("private value = %q, %v, want %q, true", got, ok, test.want)
			}
		})
	}

	unreservedElement, ok := sequence.GetItem(2).Get(tag.New(0x0019, 0x1100))
	if !ok {
		t.Fatal("unreserved private element not found")
	}
	if got := unreservedElement.Tag().PrivateCreator(); got != nil {
		t.Errorf("unreserved private element creator = %v, want nil", got)
	}
	if got := unreservedElement.ValueRepresentation(); got != vr.UN {
		t.Errorf("unreserved private element VR = %s, want UN", got)
	}
}

func writeImplicitElement(buffer *bytes.Buffer, group, element uint16, value []byte) {
	_ = binary.Write(buffer, binary.LittleEndian, group)
	_ = binary.Write(buffer, binary.LittleEndian, element)
	_ = binary.Write(buffer, binary.LittleEndian, uint32(len(value)))
	_, _ = buffer.Write(value)
}

func writeImplicitItem(buffer *bytes.Buffer, value []byte) {
	writeImplicitElement(buffer, 0xFFFE, 0xE000, value)
}
