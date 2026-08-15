// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewVolumeDataSortsSlicesAndCalculatesSpacingAndBounds(t *testing.T) {
	images := testImagesAt(t, []float64{10, 0, 5})
	for index, image := range images {
		mustAddElement(t, image.dataset,
			element.NewString(tag.PatientID, vr.LO, []string{"P-1"}),
			element.NewString(tag.SeriesDescription, vr.LO, []string{"source-" + string(rune('A'+index))}),
		)
	}

	volume, err := NewVolumeData(images)
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}
	positions := volume.SlicePositions()
	if volume.Len() != 3 || positions[0] != 0 || positions[1] != 5 || positions[2] != 10 {
		t.Fatalf("sorted positions = %v, want [0 5 10]", positions)
	}
	if volume.MinSliceSpacing() != 5 || volume.MaxSliceSpacing() != 5 {
		t.Fatalf("slice spacing = %v/%v, want 5/5", volume.MinSliceSpacing(), volume.MaxSliceSpacing())
	}
	bounds := volume.Bounds()
	if bounds.Min.Z != 0 || bounds.Max.Z != 10 || bounds.Max.X != 1 || bounds.Max.Y != 1 {
		t.Fatalf("bounds = %+v, want min z=0 max=(1,1,10)", bounds)
	}
	common := volume.CommonDataset()
	if common.TryGetString(tag.PatientID) != "P-1" {
		t.Fatal("CommonData omitted equal patient metadata")
	}
	if common.Contains(tag.SeriesDescription) {
		t.Fatal("CommonData retained frame-varying metadata")
	}
}

func TestVolumeDataMetadataAccessorsReturnDefensiveCopies(t *testing.T) {
	images := testImagesWithPixels(t, []float64{0, 5}, [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}})
	privateBinaryTag := tag.New(0x0011, 0x1010)
	for _, image := range images {
		mustAddElement(t, image.dataset,
			element.NewString(tag.PatientID, vr.LO, []string{"P-1"}),
			element.NewOtherByte(privateBinaryTag, []byte{1, 2, 3}),
		)
	}
	volume, err := NewVolumeData(images)
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}

	positions := volume.SlicePositions()
	positions[0] = 99
	if volume.SlicePositions()[0] != 0 {
		t.Fatal("SlicePositions() exposed mutable volume positions")
	}

	common := volume.CommonDataset()
	common.Remove(tag.PatientID)
	if volume.CommonDataset().TryGetString(tag.PatientID) != "P-1" {
		t.Fatal("CommonDataset() exposed mutable common metadata")
	}
	binaryData, ok := common.GetOrNil(privateBinaryTag).(*element.OtherByte)
	if !ok {
		t.Fatalf("CommonDataset() private binary type = %T, want OtherByte", common.GetOrNil(privateBinaryTag))
	}
	binaryData.GetData()[0] = 0xff
	freshBinaryData := volume.CommonDataset().GetOrNil(privateBinaryTag).(*element.OtherByte)
	if freshBinaryData.GetData()[0] == 0xff {
		t.Fatal("CommonDataset() exposed mutable OtherByte data")
	}
}

func TestNewVolumeDataRejectsIrregularSpacingUnlessEnabled(t *testing.T) {
	images := testImagesAt(t, []float64{0, 2, 5})
	if _, err := NewVolumeData(images); err == nil || !strings.Contains(err.Error(), "irregular") {
		t.Fatalf("NewVolumeData() error = %v, want irregular spacing rejection", err)
	}
	volume, err := NewVolumeData(images, WithIrregularSpacingAllowed())
	if err != nil {
		t.Fatalf("NewVolumeData(allow irregular) error = %v", err)
	}
	if volume.MinSliceSpacing() != 2 || volume.MaxSliceSpacing() != 3 {
		t.Fatalf("slice spacing = %v/%v, want 2/3", volume.MinSliceSpacing(), volume.MaxSliceSpacing())
	}
}

func TestNewVolumeDataRejectsIncompatibleSources(t *testing.T) {
	tests := []struct {
		name   string
		mutate func([]*ImageData)
		want   string
	}{
		{name: "too few", mutate: func(images []*ImageData) { images[1] = nil }, want: "at least two"},
		{name: "duplicate position", mutate: func(images []*ImageData) {
			images[1].sortingPosition = images[0].sortingPosition
			images[1].geometry.TopLeft.Z = images[0].geometry.TopLeft.Z
		}, want: "unique"},
		{name: "frame of reference", mutate: func(images []*ImageData) { images[1].geometry.FrameOfReferenceUID = "9.9.9" }, want: "Frame of Reference"},
		{name: "dimensions", mutate: func(images []*ImageData) { images[1].geometry.Columns = 3 }, want: "dimensions"},
		{name: "orientation", mutate: func(images []*ImageData) { images[1].geometry.DirectionRow.X = -1 }, want: "orientation"},
		{name: "pixel semantics", mutate: func(images []*ImageData) { images[1].pixelData.Info.BitsStored = 12 }, want: "pixel semantics"},
		{name: "modality", mutate: func(images []*ImageData) { images[1].sourceSOPClassUID = uid.MRImageStorage.UID() }, want: "modality"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			images := testImagesAt(t, []float64{0, 5})
			test.mutate(images)
			_, err := NewVolumeData(images)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("NewVolumeData() error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func testImagesAt(t testing.TB, positions []float64) []*ImageData {
	t.Helper()
	images := make([]*ImageData, len(positions))
	for index, position := range positions {
		ds := testClassicImageDataset(t, uid.CTImageStorage.UID(), "1.2.3."+string(rune('1'+index)), position, []uint16{1, 2, 3, 4})
		image, err := NewImageData(ds, 0)
		if err != nil {
			t.Fatalf("NewImageData(z=%v) error = %v", position, err)
		}
		images[index] = image
	}
	return images
}
