// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package main

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestCountElementsIgnoresGroupLength(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewUnsignedLong(tag.New(0x0009, 0x0000), []uint32{0})); err != nil {
		t.Fatalf("Add(group length) error = %v", err)
	}
	if err := ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Doe^Jane"})); err != nil {
		t.Fatalf("Add(PatientName) error = %v", err)
	}

	total, private := countElements(ds)
	if total != 1 {
		t.Errorf("countElements() total = %d, want 1 (group length ignored)", total)
	}
	if private != 0 {
		t.Errorf("countElements() private = %d, want 0 (private group length ignored)", private)
	}
}

func TestVerifyPreservation(t *testing.T) {
	base := metadataSnapshot{datasetCount: 10, privateCount: 3, fmiCount: 4}

	// Identical snapshots pass, and a source without (0002,0016) must not be
	// reported as losing it.
	if failures := verifyPreservation(base, base); len(failures) != 0 {
		t.Fatalf("verifyPreservation(identical) = %v, want no failures", failures)
	}

	withAE := base
	withAE.hasSourceAE = true
	withAE.sourceAE = "SCU"
	failures := verifyPreservation(withAE, base)
	if len(failures) != 1 || !strings.Contains(failures[0], "SourceApplicationEntityTitle") {
		t.Fatalf("verifyPreservation(lost AE) = %v, want one SourceApplicationEntityTitle failure", failures)
	}

	lostPrivate := base
	lostPrivate.privateCount = 2
	failures = verifyPreservation(base, lostPrivate)
	if len(failures) != 1 || !strings.Contains(failures[0], "Private tags") {
		t.Fatalf("verifyPreservation(lost private tags) = %v, want one Private tags failure", failures)
	}
}
