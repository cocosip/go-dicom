// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
)

// TestParseUndefinedLengthUNUsesImplicitVR verifies PS3.5 Section 6.2.2: the
// value of an undefined-length UN element is encoded as Implicit VR Little
// Endian even when the enclosing data set is Explicit VR.
func TestParseUndefinedLengthUNUsesImplicitVR(t *testing.T) {
	var encoded bytes.Buffer

	// Private creator (0011,0010) LO "ACME".
	encoded.Write([]byte{0x11, 0x00, 0x10, 0x00, 'L', 'O', 0x04, 0x00})
	encoded.WriteString("ACME")

	// (0011,1001) UN with undefined length.
	encoded.Write([]byte{0x11, 0x00, 0x01, 0x10, 'U', 'N', 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF})

	// Undefined-length item holding an Implicit VR (0010,0010) PN "DOE ".
	encoded.Write([]byte{0xFE, 0xFF, 0x00, 0xE0, 0xFF, 0xFF, 0xFF, 0xFF})
	encoded.Write([]byte{0x10, 0x00, 0x10, 0x00, 0x04, 0x00, 0x00, 0x00})
	encoded.WriteString("DOE ")
	encoded.Write([]byte{0xFE, 0xFF, 0x0D, 0xE0, 0x00, 0x00, 0x00, 0x00})

	// Sequence delimitation.
	encoded.Write([]byte{0xFE, 0xFF, 0xDD, 0xE0, 0x00, 0x00, 0x00, 0x00})

	result, err := Parse(bytes.NewReader(encoded.Bytes()),
		WithAssumedTransferSyntax(transfer.ExplicitVRLittleEndian))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	sequence, err := result.Dataset.GetSequence(tag.New(0x0011, 0x1001))
	if err != nil {
		t.Fatalf("GetSequence() error = %v", err)
	}
	if sequence.Count() != 1 {
		t.Fatalf("item count = %d, want 1", sequence.Count())
	}
	name, ok := sequence.GetItem(0).GetString(tag.PatientName)
	if !ok {
		t.Fatal("GetString(PatientName) reported missing value")
	}
	if name != "DOE" {
		t.Fatalf("PatientName = %q, want %q", name, "DOE")
	}
}

// TestParseImplicitVRForcedOWTags verifies PS3.5 Annex A.1: Pixel Data, Overlay
// Data and Waveform Data are OW in Implicit VR Little Endian, even though the
// dictionary lists OB first.
func TestParseImplicitVRForcedOWTags(t *testing.T) {
	tags := []*tag.Tag{
		tag.PixelData,
		tag.New(0x6000, 0x3000),
		tag.New(0x60FE, 0x3000),
		tag.WaveformData,
	}

	var encoded bytes.Buffer
	for range tags {
		encoded.Write([]byte{0, 0, 0, 0, 0x04, 0x00, 0x00, 0x00}) // placeholder tag/length
		encoded.Write([]byte{0x01, 0x00, 0x02, 0x00})
	}
	// Patch the real tags in place.
	raw := encoded.Bytes()
	for i, value := range tags {
		offset := i * 12
		raw[offset] = byte(value.Group())
		raw[offset+1] = byte(value.Group() >> 8)
		raw[offset+2] = byte(value.Element())
		raw[offset+3] = byte(value.Element() >> 8)
	}

	result, err := Parse(bytes.NewReader(raw),
		WithAssumedTransferSyntax(transfer.ImplicitVRLittleEndian))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	for _, value := range tags {
		elem, ok := result.Dataset.Get(value)
		if !ok {
			t.Fatalf("element %s missing", value)
		}
		if got := elem.ValueRepresentation(); got.Code() != vr.CodeOW {
			t.Errorf("%s VR = %s, want OW", value, got)
		}
	}
}

// TestObserverOffsetsAfterLazyElement verifies that sequence item offsets stay
// absolute when a lazily loaded element precedes them in an *os.File source.
func TestObserverOffsetsAfterLazyElement(t *testing.T) {
	sequence := dataset.NewSequence(tag.ReferencedSeriesSequence)
	largeItem := dataset.New()
	if err := largeItem.Add(element.NewOtherByte(tag.New(0x0011, 0x1010), make([]byte, 70000))); err != nil {
		t.Fatalf("Add(large element) error = %v", err)
	}
	if err := largeItem.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3"})); err != nil {
		t.Fatalf("Add(SeriesInstanceUID) error = %v", err)
	}
	smallItem := dataset.New()
	if err := smallItem.Add(element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.4"})); err != nil {
		t.Fatalf("Add(SeriesInstanceUID) error = %v", err)
	}
	sequence.AddItem(largeItem)
	sequence.AddItem(smallItem)

	ds := dataset.New()
	if err := ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID})); err != nil {
		t.Fatalf("Add(SOPClassUID) error = %v", err)
	}
	if err := ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5"})); err != nil {
		t.Fatalf("Add(SOPInstanceUID) error = %v", err)
	}
	if err := ds.Add(sequence); err != nil {
		t.Fatalf("Add(sequence) error = %v", err)
	}

	path := filepath.Join(t.TempDir(), "observer_lazy.dcm")
	var written []uint64
	if err := writer.WriteFile(path, ds, writer.WithSequenceItemObserver(func(position writer.SequenceItemPosition) error {
		written = append(written, position.Offset)
		return nil
	})); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer func() { _ = file.Close() }()

	var parsed []uint64
	if _, err := Parse(file, WithSequenceItemObserver(func(position SequenceItemPosition) error {
		parsed = append(parsed, position.Offset)
		return nil
	})); err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(parsed) != len(written) {
		t.Fatalf("observer calls = %d, want %d", len(parsed), len(written))
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	for i := range written {
		if parsed[i] != written[i] {
			t.Errorf("offset[%d] = %d, want %d", i, parsed[i], written[i])
		}
		if parsed[i]+4 > uint64(len(raw)) {
			t.Fatalf("offset[%d] = %d outside file length %d", i, parsed[i], len(raw))
		}
		if raw[parsed[i]] != 0xFE || raw[parsed[i]+1] != 0xFF {
			t.Errorf("offset[%d] = %d does not point at an item tag", i, parsed[i])
		}
	}
}

// TestParseImplicitVRPixelRepresentationDependentTags verifies PS3.5 Table
// 6.2-1: attributes registered as US or SS follow Pixel Representation.
func TestParseImplicitVRPixelRepresentationDependentTags(t *testing.T) {
	tags := []*tag.Tag{
		tag.ZeroVelocityPixelValue,
		tag.MappedPixelValue,
		tag.PerimeterValueRETIRED,
		tag.GrayLookupTableDataRETIRED,
	}

	var encoded bytes.Buffer
	writeImplicitUS := func(value *tag.Tag, v uint16) {
		encoded.Write([]byte{byte(value.Group()), byte(value.Group() >> 8), byte(value.Element()), byte(value.Element() >> 8), 0x02, 0x00, 0x00, 0x00})
		encoded.Write([]byte{byte(v), byte(v >> 8)})
	}
	// Pixel Representation (0028,0103) = 1 (signed).
	writeImplicitUS(tag.PixelRepresentation, 1)
	for _, value := range tags {
		writeImplicitUS(value, 0xFFFF)
	}

	result, err := Parse(bytes.NewReader(encoded.Bytes()),
		WithAssumedTransferSyntax(transfer.ImplicitVRLittleEndian))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	for _, value := range tags {
		elem, ok := result.Dataset.Get(value)
		if !ok {
			t.Fatalf("element %s missing", value)
		}
		if got := elem.ValueRepresentation(); got.Code() != vr.CodeSS {
			t.Errorf("%s VR = %s, want SS", value, got)
		}
	}
}
