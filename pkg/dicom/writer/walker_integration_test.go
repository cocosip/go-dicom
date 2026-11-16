// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package writer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
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

func TestWriteReturnsWalkerCycleError(t *testing.T) {
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
	// The writer no longer pre-walks the dataset solely to compute explicit
	// lengths (audit #13a), so a cycle is detected while the sequence is being
	// encoded and bytes may already have reached the destination. Callers
	// writing to a stream must discard the output when Write returns an error.
	if encoded.Len() > 12 {
		t.Fatalf("wrote %d Dataset bytes before cycle detection, want at most the sequence header", encoded.Len())
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

// TestCalculateWalkLengthsAllowsOver4GiBUndefinedLengthContainers verifies that
// the 4 GiB explicit-length limit is only enforced when the container is
// actually encoded with an explicit length. A sequence/item larger than 4 GiB is
// legal when it uses undefined length plus delimiters.
func TestCalculateWalkLengthsAllowsOver4GiBUndefinedLengthContainers(t *testing.T) {
	path := dataset.FormatPath(dataset.Path{{Tag: tag.ReferencedImageSequence}})
	eventPath := dataset.Path{{Tag: tag.ReferencedImageSequence}}
	newCalculator := func(explicitSequences, explicitItems bool) *walkLengthCalculator {
		return &walkLengthCalculator{
			writer: &Writer{
				isExplicitVR:                true,
				explicitLengthSequences:     explicitSequences,
				explicitLengthSequenceItems: explicitItems,
			},
			lengths: &walkLengths{sequences: map[string]uint32{}, items: map[string]uint32{}},
			frames:  []lengthFrame{{kind: dataset.WalkInvalid}},
			skipped: make(map[string]dataset.WalkEventKind),
		}
	}

	tests := []struct {
		name              string
		kind              dataset.WalkEventKind
		explicitSequences bool
		explicitItems     bool
		wantErr           bool
	}{
		{"undefined sequence length", dataset.WalkSequenceEnd, false, false, false},
		{"explicit sequence length", dataset.WalkSequenceEnd, true, false, true},
		{"undefined item length", dataset.WalkSequenceItemEnd, false, false, false},
		{"explicit item length", dataset.WalkSequenceItemEnd, false, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newCalculator(tt.explicitSequences, tt.explicitItems)
			frameKind := dataset.WalkSequenceBegin
			if tt.kind == dataset.WalkSequenceItemEnd {
				frameKind = dataset.WalkSequenceItemBegin
			}
			c.frames = append(c.frames, lengthFrame{kind: frameKind, path: path, content: math.MaxUint32})
			_, err := c.visit(dataset.WalkEvent{Kind: tt.kind, Path: eventPath})
			if (err != nil) != tt.wantErr {
				t.Fatalf("visit(%s) error = %v, wantErr %v", tt.kind, err, tt.wantErr)
			}
		})
	}
}
