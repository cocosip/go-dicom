// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"bytes"
	"context"
	"math"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

func TestEnhancedMRFixtureMatchesFoDicomReferenceAndRoundTrips(t *testing.T) {
	path := filepath.Join("..", "..", "..", "test-data", "TestMultiFrame.dcm")
	parsed, err := parser.ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile(%q) error = %v", path, err)
	}
	images, err := NewImageDataFromDataset(parsed.Dataset)
	if err != nil {
		t.Fatalf("NewImageDataFromDataset() error = %v", err)
	}
	if len(images) != 7 {
		t.Fatalf("frame count = %d, want fo-dicom reference 7", len(images))
	}

	for _, reference := range []struct {
		frame        int
		z            float64
		centerStored int64
	}{
		{frame: 0, z: -51, centerStored: 1},
		{frame: 3, z: 0, centerStored: 2},
		{frame: 6, z: 51, centerStored: 2},
	} {
		frameGeometry := images[reference.frame].Geometry()
		if !closeReference(frameGeometry.TopLeft.Z, reference.z, 1e-9) {
			t.Fatalf("frame %d top-left Z = %.12g, want %.12g", reference.frame, frameGeometry.TopLeft.Z, reference.z)
		}
		stored, err := images[reference.frame].pixelData.GetSample(images[reference.frame].pixelFrame, frameGeometry.Columns/2, frameGeometry.Rows/2, 0)
		if err != nil || stored != reference.centerStored {
			t.Fatalf("frame %d stored center sample = %v, error %v; want fo-dicom reference %v", reference.frame, stored, err, reference.centerStored)
		}
		value, valid, err := images[reference.frame].ValueAt(frameGeometry.Columns/2, frameGeometry.Rows/2)
		wantModality := float64(reference.centerStored) * 1.75848595848595
		if err != nil || !valid || !closeReference(value, wantModality, 1e-12) {
			t.Fatalf("frame %d modality center sample = %v/%v, error %v; want %v/true", reference.frame, value, valid, err, wantModality)
		}
	}

	volume, err := NewVolumeData(images)
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}
	if volume.MinSliceSpacing() != 17 || volume.MaxSliceSpacing() != 17 {
		t.Fatalf("slice spacing = %v/%v, want fo-dicom reference 17/17", volume.MinSliceSpacing(), volume.MaxSliceSpacing())
	}
	bounds := volume.Bounds()
	foMin := math3d.Point3{X: -250.00000762939, Y: -310.00000762939, Z: -51}
	if !closePointReference(bounds.Min, foMin, 1e-9) {
		t.Fatalf("minimum bounds = %+v, want fo-dicom reference %+v", bounds.Min, foMin)
	}

	cut, err := volume.Cut(context.Background(), CutSpec{
		TopLeft:             math3d.Point3{X: foMin.X, Y: foMin.Y, Z: -51},
		RowDirection:        math3d.Vector3{X: 1},
		ColumnDirection:     math3d.Vector3{Y: 1},
		Rows:                289,
		Columns:             289,
		PixelSpacingRows:    1.73,
		PixelSpacingColumns: 1.73,
	}, CutOptions{Workers: 2})
	if err != nil {
		t.Fatalf("Cut(fo-dicom reference plane) error = %v", err)
	}
	minimum, maximum, err := cut.MinMax()
	if err != nil {
		t.Fatalf("cut.MinMax() error = %v", err)
	}
	if minimum != 0 || !closeReference(maximum, 879.1460137753795, 1e-9) {
		t.Fatalf("cut range = %.12g..%.12g, want rescaled fo-dicom reference 0..879.1460137753795", minimum, maximum)
	}

	stack, err := NewStack(volume, StackTypeAxial, 1.73, 17)
	if err != nil {
		t.Fatalf("NewStack() error = %v", err)
	}
	generator, err := NewDicomGenerator(volume, WithGeneratorUIDFactory(sequenceUIDFactory()))
	if err != nil {
		t.Fatalf("NewDicomGenerator() error = %v", err)
	}
	outputs, err := generator.Generate(context.Background(), stack, "Fixture MPR", CutOptions{Workers: 2})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if len(outputs) != 7 || outputs[0].TryGetString(tag.Modality) != "MR" {
		t.Fatalf("generated count/modality = %d/%q, want 7/MR", len(outputs), outputs[0].TryGetString(tag.Modality))
	}
	encoded := &bytes.Buffer{}
	if err := writer.Write(encoded, outputs[0], writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		t.Fatalf("writer.Write() error = %v", err)
	}
	roundTrip, err := parser.Parse(encoded)
	if err != nil {
		t.Fatalf("parser.Parse(generated output) error = %v", err)
	}
	if _, err := NewImageData(roundTrip.Dataset, 0); err != nil {
		t.Fatalf("NewImageData(round-trip output) error = %v", err)
	}
}

func closePointReference(got, want math3d.Point3, tolerance float64) bool {
	return closeReference(got.X, want.X, tolerance) && closeReference(got.Y, want.Y, tolerance) && closeReference(got.Z, want.Z, tolerance)
}

func closeReference(got, want, tolerance float64) bool {
	return math.Abs(got-want) <= tolerance
}
