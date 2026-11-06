// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestDumpDatasetDefaultLayoutIncludesLengthAndKeyword(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^Jane"})); err != nil {
		t.Fatal(err)
	}

	var full bytes.Buffer
	if err := dumpDataset(&full, ds, -1, true, false); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"(0010,0010) PN [Doe^Jane]", "#", ", PatientName"} {
		if !strings.Contains(full.String(), want) {
			t.Fatalf("dump output = %q, want %q", full.String(), want)
		}
	}

	var hidden bytes.Buffer
	if err := dumpDataset(&hidden, ds, -1, false, false); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(hidden.String(), "Doe^Jane") {
		t.Fatalf("hide-values output leaked a value: %q", hidden.String())
	}
}

func TestDumpDatasetUsesWalkerOrderAndDepth(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^Jane"})); err != nil {
		t.Fatal(err)
	}
	root := dataset.New()
	if err := root.Add(dataset.NewSequenceWithItems(tag.ReferencedImageSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatal(err)
	}

	var full bytes.Buffer
	if err := dumpDataset(&full, root, -1, true, true); err != nil {
		t.Fatal(err)
	}
	wantLines := []string{
		"(0008,1140) SQ (1 item)",
		"  Item #1:",
		"  (0010,0010) PN [Doe^Jane]",
	}
	for _, line := range wantLines {
		if !strings.Contains(full.String(), line+"\n") {
			t.Fatalf("dump output = %q, want line %q", full.String(), line)
		}
	}

	var shallow bytes.Buffer
	if err := dumpDataset(&shallow, root, 0, true, true); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(shallow.String(), "Item #") || strings.Contains(shallow.String(), "Doe^Jane") {
		t.Fatalf("depth-limited output contains Sequence children: %q", shallow.String())
	}
}
