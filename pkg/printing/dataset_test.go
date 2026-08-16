// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"reflect"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

func TestEmptyUIDsAreUnique(t *testing.T) {
	constructors := map[string]func() string{
		"film session": func() string {
			return NewFilmSession("", "", false).SOPInstanceUID
		},
		"film box": func() string {
			return NewFilmBox("", `STANDARD\1,1`).SOPInstanceUID
		},
		"image box": func() string {
			return NewImageBox("", false).SOPInstanceUID
		},
		"presentation LUT": func() string {
			return NewPresentationLUT("").SOPInstanceUID
		},
	}

	for name, construct := range constructors {
		t.Run(name, func(t *testing.T) {
			first := construct()
			second := construct()
			if first == second {
				t.Fatalf("empty UID constructor reused %q", first)
			}
			for _, value := range []string{first, second} {
				if !strings.HasPrefix(value, "2.25.") {
					t.Errorf("UID %q does not use the UUID-derived 2.25 root", value)
				}
				if !uid.IsValid(value) {
					t.Errorf("UID %q is not valid", value)
				}
			}
		})
	}
}

func TestFilmSessionDatasetRoundTrip(t *testing.T) {
	want := NewFilmSession("1.2.840.10008.5.1.1.1", "2.25.101", true)
	want.FilmDestination = FilmDestinationMagazine
	want.FilmSessionLabel = "urgent chest"
	want.MemoryAllocation = 4096
	want.MediumType = MediumTypeBlueFilm
	want.PrintPriority = PrintPriorityHigh
	want.NumberOfCopies = 3

	ds, err := want.ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	got, err := NewFilmSessionFromDataset("", "", ds, true)
	if err != nil {
		t.Fatalf("NewFilmSessionFromDataset() error = %v", err)
	}

	if got.SOPClassUID != want.SOPClassUID || got.SOPInstanceUID != want.SOPInstanceUID ||
		got.FilmDestination != want.FilmDestination || got.FilmSessionLabel != want.FilmSessionLabel ||
		got.MemoryAllocation != want.MemoryAllocation || got.MediumType != want.MediumType ||
		got.PrintPriority != want.PrintPriority || got.NumberOfCopies != want.NumberOfCopies || !got.IsColor {
		t.Fatalf("round trip mismatch: got %#v, want %#v", got, want)
	}
}

func TestFilmBoxDatasetRoundTrip(t *testing.T) {
	want := NewFilmBox("2.25.102", `ROW\1,2`)
	want.FilmOrientation = FilmOrientationLandscape
	want.FilmSizeID = FilmSizeA3
	want.MagnificationType = MagnificationCubic
	want.MaxDensity = 320
	want.MinDensity = 20
	want.Trim = TrimYes
	want.BorderDensity = BorderDensityWhite
	want.EmptyImageDensity = EmptyImageDensityWhite
	want.ConfigurationInformation = "CS123"

	ds, err := want.ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	got, err := NewFilmBoxFromDataset("", ds)
	if err != nil {
		t.Fatalf("NewFilmBoxFromDataset() error = %v", err)
	}

	if got.SOPInstanceUID != want.SOPInstanceUID || got.ImageDisplayFormat != want.ImageDisplayFormat ||
		got.FilmOrientation != want.FilmOrientation || got.FilmSizeID != want.FilmSizeID ||
		got.MagnificationType != want.MagnificationType || got.MaxDensity != want.MaxDensity ||
		got.MinDensity != want.MinDensity || got.Trim != want.Trim ||
		got.BorderDensity != want.BorderDensity || got.EmptyImageDensity != want.EmptyImageDensity ||
		got.ConfigurationInformation != want.ConfigurationInformation {
		t.Fatalf("round trip mismatch: got %#v, want %#v", got, want)
	}
}

func TestImageBoxDatasetRoundTrip(t *testing.T) {
	want := NewImageBox("2.25.103", true)
	want.ImageBoxPosition = 4
	want.Polarity = PolarityReverse
	want.MagnificationType = MagnificationBilinear
	want.SmoothingType = "SMOOTH_2"
	want.RequestedImageSize = "120.5"
	want.SetImageData([]byte{1, 2, 3, 4})

	ds, err := want.ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	got, err := NewImageBoxFromDataset("", ds, true)
	if err != nil {
		t.Fatalf("NewImageBoxFromDataset() error = %v", err)
	}

	if got.SOPClassUID != want.SOPClassUID || got.SOPInstanceUID != want.SOPInstanceUID ||
		got.ImageBoxPosition != want.ImageBoxPosition || got.Polarity != want.Polarity ||
		got.MagnificationType != want.MagnificationType || got.SmoothingType != want.SmoothingType ||
		got.RequestedImageSize != want.RequestedImageSize || !reflect.DeepEqual(got.GetImageData(), []byte{1, 2, 3, 4}) || !got.IsColor {
		t.Fatalf("round trip mismatch: got %#v, want %#v", got, want)
	}

	want.PreformattedColorImageSequence[0] = 99
	if got.GetImageData()[0] != 1 {
		t.Fatal("loaded ImageBox aliases source image bytes")
	}
}

func TestPresentationLUTDatasetRoundTrip(t *testing.T) {
	want := NewPresentationLUT("2.25.104")
	want.LUTExplanation = "calibrated"
	want.PresentationLUTShape = PresentationLUTShapeLinOD
	if err := want.SetLUT(3, 0, 12, []uint16{10, 20, 30}); err != nil {
		t.Fatalf("SetLUT() error = %v", err)
	}

	ds, err := want.ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	got, err := NewPresentationLUTFromDataset("", ds)
	if err != nil {
		t.Fatalf("NewPresentationLUTFromDataset() error = %v", err)
	}

	if got.SOPInstanceUID != want.SOPInstanceUID || got.LUTExplanation != want.LUTExplanation ||
		got.PresentationLUTShape != want.PresentationLUTShape ||
		!reflect.DeepEqual(got.LUTDescriptor, []uint16{3, 0, 12}) ||
		!reflect.DeepEqual(got.LUTData, []uint16{10, 20, 30}) {
		t.Fatalf("round trip mismatch: got %#v, want %#v", got, want)
	}

	want.LUTDescriptor[0] = 99
	want.LUTData[0] = 99
	if got.LUTDescriptor[0] != 3 || got.LUTData[0] != 10 {
		t.Fatal("loaded PresentationLUT aliases source slices")
	}
}
