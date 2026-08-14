// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package writer

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestWriteSequenceItemObserverReportsAbsoluteItemOffset(t *testing.T) {
	ds, item := observedSequenceDataset(t)

	tests := []struct {
		name    string
		options []WriteOption
	}{
		{name: "undefined lengths"},
		{name: "explicit sequence length", options: []WriteOption{WithExplicitLengthSequences(true)}},
		{name: "explicit item length", options: []WriteOption{WithExplicitLengthSequenceItems()}},
		{name: "explicit sequence and item lengths", options: []WriteOption{
			WithExplicitLengthSequences(true),
			WithExplicitLengthSequenceItems(),
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var positions []SequenceItemPosition
			options := append([]WriteOption{}, tt.options...)
			options = append(options,
				WithoutPreamble(),
				WithSequenceItemObserver(func(position SequenceItemPosition) error {
					positions = append(positions, position)
					return nil
				}),
			)

			var encoded bytes.Buffer
			if err := Write(&encoded, ds, options...); err != nil {
				t.Fatalf("Write() error = %v", err)
			}
			if len(positions) != 1 {
				t.Fatalf("observer calls = %d, want 1", len(positions))
			}

			position := positions[0]
			if !position.SequenceTag.Equals(tag.ReferencedSeriesSequence) {
				t.Fatalf("SequenceTag = %s, want %s", position.SequenceTag, tag.ReferencedSeriesSequence)
			}
			if position.Item != item {
				t.Fatal("observer Item does not match the written sequence item")
			}
			if got := binary.LittleEndian.Uint32(encoded.Bytes()[position.Offset : position.Offset+4]); got != 0xE000FFFE {
				t.Fatalf("offset points to %#08x, want Item Tag", got)
			}
		})
	}
}

func TestWriteSequenceItemObserverDoesNotChangeEncodedBytes(t *testing.T) {
	ds, _ := observedSequenceDataset(t)

	var baseline bytes.Buffer
	if err := Write(&baseline, ds, WithoutPreamble()); err != nil {
		t.Fatalf("baseline Write() error = %v", err)
	}

	var observed bytes.Buffer
	if err := Write(&observed, ds,
		WithoutPreamble(),
		WithSequenceItemObserver(func(SequenceItemPosition) error { return nil }),
	); err != nil {
		t.Fatalf("observed Write() error = %v", err)
	}

	if !bytes.Equal(baseline.Bytes(), observed.Bytes()) {
		t.Fatal("sequence item observer changed encoded bytes")
	}
}

func observedSequenceDataset(t *testing.T) (*dataset.Dataset, *dataset.Dataset) {
	t.Helper()

	item := dataset.New()
	if err := item.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3"})); err != nil {
		t.Fatalf("item.Add() error = %v", err)
	}

	ds := dataset.New()
	if err := ds.Add(dataset.NewSequenceWithItems(tag.ReferencedSeriesSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("dataset.Add() error = %v", err)
	}
	return ds, item
}
