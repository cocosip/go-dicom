// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestWriteTreatsNilSequenceItemAsEmptyItem(t *testing.T) {
	tests := []struct {
		name    string
		options []WriteOption
	}{
		{name: "undefined lengths"},
		{name: "explicit sequence", options: []WriteOption{WithExplicitLengthSequences(true)}},
		{name: "explicit item", options: []WriteOption{WithExplicitLengthSequenceItems()}},
		{name: "explicit sequence and item", options: []WriteOption{
			WithExplicitLengthSequences(true),
			WithExplicitLengthSequenceItems(),
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			sequence := dataset.NewSequenceWithItems(tag.ReferencedImageSequence, []*dataset.Dataset{nil})
			if err := ds.Add(sequence); err != nil {
				t.Fatal(err)
			}

			var encoded bytes.Buffer
			options := append([]WriteOption{WithoutPreamble()}, tt.options...)
			if err := Write(&encoded, ds, options...); err != nil {
				t.Fatalf("Write() error = %v", err)
			}

			result, err := parser.Parse(&encoded, parser.WithAssumedTransferSyntax(transfer.ExplicitVRLittleEndian))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			got, ok := result.Dataset.Get(tag.ReferencedImageSequence)
			if !ok {
				t.Fatal("ReferencedImageSequence not found")
			}
			parsed := got.(*dataset.Sequence)
			if parsed.Count() != 1 || parsed.GetItem(0) == nil || parsed.GetItem(0).Count() != 0 {
				t.Fatalf("parsed items = %#v, want one empty Dataset", parsed.GetItems())
			}
		})
	}
}

func TestWriteReturnsWalkerCycleErrorBeforeDatasetBytes(t *testing.T) {
	ds := dataset.New()
	sequence := dataset.NewSequence(tag.ReferencedImageSequence)
	if err := ds.Add(sequence); err != nil {
		t.Fatal(err)
	}
	sequence.AddItem(ds)

	var encoded bytes.Buffer
	err := Write(&encoded, ds, WithoutPreamble())
	var walkErr *dataset.WalkError
	if !errors.As(err, &walkErr) {
		t.Fatalf("Write() error = %v, want *dataset.WalkError", err)
	}
	if got := dataset.FormatPath(walkErr.Path); got != "(0008,1140)[0]" {
		t.Fatalf("cycle path = %q, want (0008,1140)[0]", got)
	}
	if encoded.Len() != 0 {
		t.Fatalf("wrote %d Dataset bytes before cycle detection", encoded.Len())
	}
}

func TestCalculateWalkLengthsForExplicitSequenceAndItem(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewString(tag.PatientName, vr.PN, []string{"A"})); err != nil {
		t.Fatal(err)
	}
	ds := dataset.New()
	if err := ds.Add(dataset.NewSequenceWithItems(tag.ReferencedImageSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatal(err)
	}

	w := &Writer{
		byteOrder:                   binary.LittleEndian,
		isExplicitVR:                true,
		explicitLengthSequences:     true,
		explicitLengthSequenceItems: true,
	}
	lengths, err := w.calculateWalkLengths(ds)
	if err != nil {
		t.Fatal(err)
	}
	if got := lengths.items["(0008,1140)[0]"]; got != 10 {
		t.Fatalf("item content length = %d, want 10", got)
	}
	if got := lengths.sequences["(0008,1140)"]; got != 18 {
		t.Fatalf("sequence content length = %d, want 18", got)
	}
}
