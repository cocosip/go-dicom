// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"context"
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

func TestVolumeSampleInterpolatesPatientSpaceAndIncludesBoundaryCenters(t *testing.T) {
	volume := testVolumeWithPixels(t,
		[]float64{0, 10},
		[][]uint16{{0, 10, 20, 30}, {100, 110, 120, 130}},
	)

	value, valid, err := volume.Sample(math3d.Point3{X: 0.5, Y: 0.5, Z: 5})
	if err != nil {
		t.Fatalf("Sample(center) error = %v", err)
	}
	if !valid || value != 65 {
		t.Fatalf("Sample(center) = %v/%v, want 65/true", value, valid)
	}
	value, valid, err = volume.Sample(math3d.Point3{X: 1, Y: 1, Z: 10})
	if err != nil {
		t.Fatalf("Sample(last center) error = %v", err)
	}
	if !valid || value != 130 {
		t.Fatalf("Sample(last center) = %v/%v, want 130/true", value, valid)
	}
	if _, valid, err = volume.Sample(math3d.Point3{X: 1.01, Y: 1, Z: 10}); err != nil || valid {
		t.Fatalf("Sample(outside x) = valid %v, error %v; want false/nil", valid, err)
	}
	if _, valid, err = volume.Sample(math3d.Point3{X: 0, Y: 0, Z: -0.01}); err != nil || valid {
		t.Fatalf("Sample(outside z) = valid %v, error %v; want false/nil", valid, err)
	}
}

func TestVolumeSampleUsesActualIrregularSliceDistances(t *testing.T) {
	images := testImagesWithPixels(t,
		[]float64{0, 2, 5},
		[][]uint16{{0, 0, 0, 0}, {20, 20, 20, 20}, {50, 50, 50, 50}},
	)
	volume, err := NewVolumeData(images, WithIrregularSpacingAllowed())
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}
	value, valid, err := volume.Sample(math3d.Point3{Z: 3.5})
	if err != nil || !valid || value != 35 {
		t.Fatalf("Sample(z=3.5) = %v/%v, error %v; want 35/true", value, valid, err)
	}
}

func TestVolumeSamplePropagatesPaddingThroughValidity(t *testing.T) {
	images := testImagesWithPixels(t,
		[]float64{0, 10},
		[][]uint16{{0, 10, 20, 30}, {100, 110, 120, 130}},
	)
	padding := int32(0)
	images[0].pixelData.Info.PixelPaddingValue = &padding
	volume, err := NewVolumeData(images)
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}
	if _, valid, err := volume.Sample(math3d.Point3{X: 0.5, Y: 0.5, Z: 5}); err != nil || valid {
		t.Fatalf("Sample(with padding neighbor) = valid %v, error %v; want false/nil", valid, err)
	}
}

func TestVolumeCutIsDeterministicAcrossWorkerCountsAndCancelable(t *testing.T) {
	volume := testVolumeWithPixels(t,
		[]float64{0, 10},
		[][]uint16{{0, 10, 20, 30}, {100, 110, 120, 130}},
	)
	spec := CutSpec{
		TopLeft:         math3d.Point3{Z: 5},
		RowDirection:    math3d.Vector3{X: 1},
		ColumnDirection: math3d.Vector3{Y: 1},
		Rows:            2, Columns: 2,
		PixelSpacingRows: 1, PixelSpacingColumns: 1,
	}
	one, err := volume.Cut(context.Background(), spec, CutOptions{Workers: 1})
	if err != nil {
		t.Fatalf("Cut(workers=1) error = %v", err)
	}
	four, err := volume.Cut(context.Background(), spec, CutOptions{Workers: 4})
	if err != nil {
		t.Fatalf("Cut(workers=4) error = %v", err)
	}
	if !reflect.DeepEqual(one.Values, []float64{50, 60, 70, 80}) || !reflect.DeepEqual(one.Valid, []bool{true, true, true, true}) {
		t.Fatalf("Cut() = %v/%v", one.Values, one.Valid)
	}
	if !reflect.DeepEqual(one.Values, four.Values) || !reflect.DeepEqual(one.Valid, four.Valid) {
		t.Fatal("Cut() changed with worker count")
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := volume.Cut(ctx, spec, CutOptions{Workers: 2}); !errors.Is(err, context.Canceled) {
		t.Fatalf("Cut(canceled) error = %v, want context.Canceled", err)
	}
}

func TestVolumeCutRejectsInvalidSpecification(t *testing.T) {
	volume := testVolumeWithPixels(t, []float64{0, 1}, [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}})
	_, err := volume.Cut(context.Background(), CutSpec{
		RowDirection: math3d.Vector3{X: 1}, ColumnDirection: math3d.Vector3{X: 1},
		Rows: 2, Columns: 2, PixelSpacingRows: 1, PixelSpacingColumns: 1,
	}, CutOptions{})
	if err == nil {
		t.Fatal("Cut() accepted parallel row and column directions")
	}
}

func TestVolumeCutRejectsDimensionsOutsideDicomRange(t *testing.T) {
	volume := testVolumeWithPixels(t, []float64{0, 1}, [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}})
	_, err := volume.Cut(context.Background(), CutSpec{
		RowDirection: math3d.Vector3{X: 1}, ColumnDirection: math3d.Vector3{Y: 1},
		Rows: 65536, Columns: 1, PixelSpacingRows: 1, PixelSpacingColumns: 1,
	}, CutOptions{})
	if err == nil {
		t.Fatal("Cut() accepted dimensions outside the DICOM US range")
	}
}

func TestVolumeCutSupportsObliquePatientSpacePlane(t *testing.T) {
	volume := testVolumeWithPixels(t,
		[]float64{0, 10},
		[][]uint16{{0, 10, 20, 30}, {100, 110, 120, 130}},
	)
	sqrtHalf := math.Sqrt(0.5)
	slice, err := volume.Cut(context.Background(), CutSpec{
		TopLeft:             math3d.Point3{},
		RowDirection:        math3d.Vector3{X: sqrtHalf, Y: sqrtHalf},
		ColumnDirection:     math3d.Vector3{Z: 1},
		Rows:                2,
		Columns:             2,
		PixelSpacingRows:    5,
		PixelSpacingColumns: math.Sqrt(2),
	}, CutOptions{Workers: 2})
	if err != nil {
		t.Fatalf("Cut(oblique) error = %v", err)
	}
	if !reflect.DeepEqual(slice.Values, []float64{0, 30, 50, 80}) || !reflect.DeepEqual(slice.Valid, []bool{true, true, true, true}) {
		t.Fatalf("Cut(oblique) = %v/%v", slice.Values, slice.Valid)
	}
}

func testVolumeWithPixels(t testing.TB, positions []float64, pixels [][]uint16) *VolumeData {
	t.Helper()
	volume, err := NewVolumeData(testImagesWithPixels(t, positions, pixels))
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}
	return volume
}

func testImagesWithPixels(t testing.TB, positions []float64, pixels [][]uint16) []*ImageData {
	t.Helper()
	images := make([]*ImageData, len(positions))
	for index, position := range positions {
		ds := testClassicImageDataset(t, uid.CTImageStorage.UID(), "1.2.840."+string(rune('1'+index)), position, pixels[index])
		mustAddElement(t, ds, element.NewString(tag.Modality, vr.CS, []string{"CT"}))
		image, err := NewImageData(ds, 0)
		if err != nil {
			t.Fatalf("NewImageData(z=%v) error = %v", position, err)
		}
		images[index] = image
	}
	return images
}
