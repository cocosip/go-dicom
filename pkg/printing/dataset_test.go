// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestEmptyUIDsAreUnique(t *testing.T) {
	constructors := map[string]func() string{
		"film session": func() string {
			return NewFilmSession("", "", false).SOPInstanceUID
		},
		"film box": func() string {
			return NewFilmBox("", testStandardOneByOne).SOPInstanceUID
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

func TestFilmBoxDatasetRoundTripsStandardAttributesAndReferences(t *testing.T) {
	const presentationLUTInstanceUID = "2.25.122"

	session := NewFilmSession("", "2.25.120", false)
	box := NewFilmBox("2.25.121", testStandardTwoByOne)
	box.AnnotationDisplayFormatID = "ANNOTATION_1"
	box.SmoothingType = "SMOOTH_1"
	box.Illumination = 2000
	box.ReflectedAmbientLight = 10
	box.RequestedResolutionID = "HIGH"
	box.ReferencedPresentationLUT = SOPReference{
		SOPClassUID:    presentationLUTSOPClassUID,
		SOPInstanceUID: presentationLUTInstanceUID,
	}
	first := NewImageBox("2.25.123", false)
	first.ImageBoxPosition = 1
	second := NewImageBox("2.25.124", false)
	second.ImageBoxPosition = 2
	box.AddImageBox(first)
	box.AddImageBox(second)
	session.AddFilmBox(box)

	ds, err := box.ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	for target, want := range map[*tag.Tag]string{
		tag.AnnotationDisplayFormatID: "ANNOTATION_1",
		tag.SmoothingType:             "SMOOTH_1",
		tag.RequestedResolutionID:     "HIGH",
	} {
		if got, ok := ds.GetString(target); !ok || got != want {
			t.Errorf("%s = %q, %v; want %q", target, got, ok, want)
		}
	}
	if got, err := ds.GetUInt16(tag.Illumination, 0); err != nil || got != 2000 {
		t.Errorf("Illumination = %d, %v", got, err)
	}
	if got, err := ds.GetUInt16(tag.ReflectedAmbientLight, 0); err != nil || got != 10 {
		t.Errorf("ReflectedAmbientLight = %d, %v", got, err)
	}
	assertReferenceSequence(t, ds, tag.ReferencedFilmSessionSequence, []SOPReference{{
		SOPClassUID: basicFilmSessionSOPClassUID, SOPInstanceUID: "2.25.120",
	}})
	assertReferenceSequence(t, ds, tag.ReferencedImageBoxSequence, []SOPReference{
		{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.123"},
		{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.124"},
	})
	assertReferenceSequence(t, ds, tag.ReferencedPresentationLUTSequence, []SOPReference{{
		SOPClassUID: presentationLUTSOPClassUID, SOPInstanceUID: presentationLUTInstanceUID,
	}})

	got, err := NewFilmBoxFromDataset("", ds)
	if err != nil {
		t.Fatalf("NewFilmBoxFromDataset() error = %v", err)
	}
	if got.AnnotationDisplayFormatID != box.AnnotationDisplayFormatID ||
		got.SmoothingType != box.SmoothingType ||
		got.Illumination != box.Illumination ||
		got.ReflectedAmbientLight != box.ReflectedAmbientLight ||
		got.RequestedResolutionID != box.RequestedResolutionID ||
		got.ReferencedFilmSession.SOPInstanceUID != "2.25.120" ||
		got.ReferencedPresentationLUT.SOPInstanceUID != presentationLUTInstanceUID ||
		!reflect.DeepEqual(got.ReferencedImageBoxes, []SOPReference{
			{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.123"},
			{SOPClassUID: SOPClassGrayscaleImageBox, SOPInstanceUID: "2.25.124"},
		}) {
		t.Fatalf("Film Box standard round trip = %#v", got)
	}
}

func TestFilmBoxDatasetRejectsMalformedStandardUInt16Attributes(t *testing.T) {
	for _, target := range []*tag.Tag{tag.Illumination, tag.ReflectedAmbientLight} {
		t.Run(target.String(), func(t *testing.T) {
			ds, err := NewFilmBox("2.25.125", testStandardOneByOne).ToDataset()
			if err != nil {
				t.Fatalf("ToDataset() error = %v", err)
			}
			ds.SetAutoValidate(false)
			if err := ds.AddOrUpdate(element.NewString(target, vr.LO, []string{"invalid"})); err != nil {
				t.Fatalf("AddOrUpdate(%s) error = %v", target, err)
			}

			if _, err := NewFilmBoxFromDataset("", ds); err == nil {
				t.Fatalf("NewFilmBoxFromDataset() accepted malformed %s", target)
			}
		})
	}
}

func TestFilmBoxDatasetRejectsMultiValuedStandardUInt16Attributes(t *testing.T) {
	for _, target := range []*tag.Tag{tag.Illumination, tag.ReflectedAmbientLight} {
		t.Run(target.String(), func(t *testing.T) {
			ds, err := NewFilmBox("2.25.126", testStandardOneByOne).ToDataset()
			if err != nil {
				t.Fatalf("ToDataset() error = %v", err)
			}
			ds.SetAutoValidate(false)
			if err := ds.AddOrUpdate(element.NewUnsignedShort(target, []uint16{1, 2})); err != nil {
				t.Fatalf("AddOrUpdate(%s) error = %v", target, err)
			}

			if _, err := NewFilmBoxFromDataset("", ds); err == nil {
				t.Fatalf("NewFilmBoxFromDataset() accepted multi-valued %s", target)
			}
		})
	}
}

func TestFilmBoxDatasetAllowsEmptyOptionalUInt16Attributes(t *testing.T) {
	for _, target := range []*tag.Tag{tag.Illumination, tag.ReflectedAmbientLight} {
		t.Run(target.String(), func(t *testing.T) {
			ds, err := NewFilmBox("2.25.127", testStandardOneByOne).ToDataset()
			if err != nil {
				t.Fatalf("ToDataset() error = %v", err)
			}
			if err := ds.AddOrUpdate(element.NewUnsignedShort(target, nil)); err != nil {
				t.Fatalf("AddOrUpdate(%s) error = %v", target, err)
			}

			if _, err := NewFilmBoxFromDataset("", ds); err != nil {
				t.Fatalf("NewFilmBoxFromDataset() error = %v", err)
			}
		})
	}
}

func TestFilmBoxDatasetRejectsOddLengthOptionalUInt16Attributes(t *testing.T) {
	for _, length := range []int{1, 3} {
		t.Run(fmt.Sprintf("length_%d", length), func(t *testing.T) {
			ds, err := NewFilmBox("2.25.128", testStandardOneByOne).ToDataset()
			if err != nil {
				t.Fatalf("ToDataset() error = %v", err)
			}
			ds.SetAutoValidate(false)
			malformed := element.NewUnsignedShortFromBuffer(tag.Illumination, buffer.NewMemory(make([]byte, length)))
			if err := ds.AddOrUpdate(malformed); err != nil {
				t.Fatalf("AddOrUpdate(Illumination) error = %v", err)
			}

			if _, err := NewFilmBoxFromDataset("", ds); err == nil {
				t.Fatalf("NewFilmBoxFromDataset() accepted a %d-byte US value", length)
			}
		})
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

func TestImageBoxDatasetRoundTripsOptionalOverridesIncludingZero(t *testing.T) {
	maxDensity := uint16(0)
	minDensity := uint16(15)
	configuration := "CS456"
	want := NewImageBox("2.25.130", false)
	want.MaxDensity = &maxDensity
	want.MinDensity = &minDensity
	want.ConfigurationInformation = &configuration
	want.RequestedDecimateCropBehavior = "CROP"

	ds, err := want.ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	got, err := NewImageBoxFromDataset("", ds, false)
	if err != nil {
		t.Fatalf("NewImageBoxFromDataset() error = %v", err)
	}
	if got.MaxDensity == nil || *got.MaxDensity != 0 ||
		got.MinDensity == nil || *got.MinDensity != 15 ||
		got.ConfigurationInformation == nil || *got.ConfigurationInformation != "CS456" ||
		got.RequestedDecimateCropBehavior != "CROP" {
		t.Fatalf("Image Box overrides = %#v", got)
	}
}

func TestImageBoxDatasetAllowsEmptyOptionalUInt16Attributes(t *testing.T) {
	for _, target := range []*tag.Tag{tag.MaxDensity, tag.MinDensity} {
		t.Run(target.String(), func(t *testing.T) {
			ds, err := NewImageBox("2.25.134", false).ToDataset()
			if err != nil {
				t.Fatalf("ToDataset() error = %v", err)
			}
			if err := ds.AddOrUpdate(element.NewUnsignedShort(target, nil)); err != nil {
				t.Fatalf("AddOrUpdate(%s) error = %v", target, err)
			}

			got, err := NewImageBoxFromDataset("", ds, false)
			if err != nil {
				t.Fatalf("NewImageBoxFromDataset() error = %v", err)
			}
			if got.MaxDensity != nil || got.MinDensity != nil {
				t.Fatalf("empty density attributes became values: %#v", got)
			}
		})
	}
}

func TestImageBoxDatasetAllowsEmptyConfigurationInformation(t *testing.T) {
	ds, err := NewImageBox("2.25.135", false).ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	if err := ds.AddOrUpdate(element.NewString(tag.ConfigurationInformation, vr.ST, []string{""})); err != nil {
		t.Fatalf("AddOrUpdate(ConfigurationInformation) error = %v", err)
	}

	got, err := NewImageBoxFromDataset("", ds, false)
	if err != nil {
		t.Fatalf("NewImageBoxFromDataset() error = %v", err)
	}
	if got.ConfigurationInformation == nil || *got.ConfigurationInformation != "" {
		t.Fatalf("ConfigurationInformation = %#v, want present empty value", got.ConfigurationInformation)
	}
}

func TestImageBoxDatasetRejectsMalformedConfigurationInformation(t *testing.T) {
	for _, malformed := range []element.Element{
		element.NewUnsignedShort(tag.ConfigurationInformation, []uint16{1}),
		element.NewString(tag.ConfigurationInformation, vr.LO, []string{"wrong VR"}),
	} {
		t.Run(malformed.ValueRepresentation().String(), func(t *testing.T) {
			ds, err := NewImageBox("2.25.135", false).ToDataset()
			if err != nil {
				t.Fatalf("ToDataset() error = %v", err)
			}
			ds.SetAutoValidate(false)
			if err := ds.AddOrUpdate(malformed); err != nil {
				t.Fatalf("AddOrUpdate(ConfigurationInformation) error = %v", err)
			}

			if _, err := NewImageBoxFromDataset("", ds, false); err == nil {
				t.Fatal("NewImageBoxFromDataset() accepted malformed Configuration Information")
			}
		})
	}
}

func TestImageBoxDatasetRoundTripsCompleteImageSequenceIndependently(t *testing.T) {
	for _, test := range []struct {
		name        string
		color       bool
		sequenceTag *tag.Tag
		classUID    string
		photometric string
		samples     uint16
		pixelData   []byte
	}{
		{name: "grayscale", sequenceTag: tag.BasicGrayscaleImageSequence, classUID: SOPClassGrayscaleImageBox, photometric: "MONOCHROME2", samples: 1, pixelData: []byte{1, 2, 3, 4}},
		{name: "color", color: true, sequenceTag: tag.BasicColorImageSequence, classUID: SOPClassColorImageBox, photometric: "RGB", samples: 3, pixelData: []byte{1, 2, 3, 4, 5, 6}},
	} {
		t.Run(test.name, func(t *testing.T) {
			item := dataset.New()
			values := []element.Element{
				element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{test.samples}),
				element.NewString(tag.PhotometricInterpretation, vr.CS, []string{test.photometric}),
				element.NewUnsignedShort(tag.Rows, []uint16{1}),
				element.NewUnsignedShort(tag.Columns, []uint16{2}),
				element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
				element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
				element.NewUnsignedShort(tag.HighBit, []uint16{7}),
				element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
				element.NewOtherByte(tag.PixelData, append([]byte(nil), test.pixelData...)),
			}
			for _, value := range values {
				if err := item.Add(value); err != nil {
					t.Fatalf("add image item element %s: %v", value.Tag(), err)
				}
			}

			image := NewImageBox("2.25.131", test.color)
			if err := image.SetImageSequence(item); err != nil {
				t.Fatalf("SetImageSequence() error = %v", err)
			}
			item.Remove(tag.Rows)
			stored, err := image.ImageSequence()
			if err != nil {
				t.Fatalf("ImageSequence() error = %v", err)
			}
			if rows, err := stored.GetUInt16(tag.Rows, 0); err != nil || rows != 1 {
				t.Fatalf("stored Rows = %d, %v; source mutation leaked", rows, err)
			}

			ds, err := image.ToDataset()
			if err != nil {
				t.Fatalf("ToDataset() error = %v", err)
			}
			sequence, err := ds.GetSequence(test.sequenceTag)
			if err != nil || sequence.Count() != 1 {
				t.Fatalf("image sequence = %#v, %v", sequence, err)
			}
			if got, _ := sequence.GetItem(0).GetString(tag.PhotometricInterpretation); got != test.photometric {
				t.Errorf("PhotometricInterpretation = %q, want %q", got, test.photometric)
			}

			loaded, err := NewImageBoxFromDataset("", ds, test.color)
			if err != nil {
				t.Fatalf("NewImageBoxFromDataset() error = %v", err)
			}
			if loaded.SOPClassUID != test.classUID || !reflect.DeepEqual(loaded.GetImageData(), test.pixelData) {
				t.Fatalf("loaded Image Box = %#v", loaded)
			}
			firstCopy, err := loaded.ImageSequence()
			if err != nil {
				t.Fatalf("first ImageSequence() error = %v", err)
			}
			firstCopy.Remove(tag.Columns)
			secondCopy, err := loaded.ImageSequence()
			if err != nil {
				t.Fatalf("second ImageSequence() error = %v", err)
			}
			if columns, err := secondCopy.GetUInt16(tag.Columns, 0); err != nil || columns != 2 {
				t.Fatalf("stored Columns = %d, %v; returned Dataset mutation leaked", columns, err)
			}
		})
	}
}

func TestImageBoxImageSequencePreservesPixelDataVR(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewOtherWord(tag.PixelData, []byte{1, 0, 2, 0})); err != nil {
		t.Fatalf("add OW Pixel Data: %v", err)
	}
	image := NewImageBox("2.25.132", false)
	if err := image.SetImageSequence(item); err != nil {
		t.Fatalf("SetImageSequence() error = %v", err)
	}

	stored, err := image.ImageSequence()
	if err != nil {
		t.Fatalf("ImageSequence() error = %v", err)
	}
	pixelData, ok := stored.Get(tag.PixelData)
	if !ok || pixelData.ValueRepresentation() != vr.OW {
		t.Fatalf("Pixel Data VR = %v, want OW", pixelData)
	}
}

func TestImageBoxSetImageDataPreservesExistingPixelDataVR(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewOtherWord(tag.PixelData, []byte{1, 0, 2, 0})); err != nil {
		t.Fatalf("add OW Pixel Data: %v", err)
	}
	image := NewImageBox("2.25.134", false)
	if err := image.SetImageSequence(item); err != nil {
		t.Fatalf("SetImageSequence() error = %v", err)
	}

	image.SetImageData([]byte{3, 0, 4, 0})
	stored, err := image.ImageSequence()
	if err != nil {
		t.Fatalf("ImageSequence() error = %v", err)
	}
	pixelData, ok := stored.Get(tag.PixelData)
	if !ok || pixelData.ValueRepresentation() != vr.OW {
		t.Fatalf("updated Pixel Data VR = %v, want OW", pixelData)
	}
	if got := pixelData.Buffer().Data(); !reflect.DeepEqual(got, []byte{3, 0, 4, 0}) {
		t.Fatalf("updated Pixel Data = %v", got)
	}
}

func TestImageBoxDatasetRejectsMalformedImageSequence(t *testing.T) {
	ds, err := NewImageBox("2.25.133", false).ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	ds.SetAutoValidate(false)
	if err := ds.Add(element.NewString(tag.BasicGrayscaleImageSequence, vr.LO, []string{"not a sequence"})); err != nil {
		t.Fatalf("add malformed image sequence: %v", err)
	}

	if _, err := NewImageBoxFromDataset("", ds, false); err == nil {
		t.Fatal("NewImageBoxFromDataset() accepted a malformed image sequence")
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

func TestPresentationLUTDatasetUsesCanonicalSequence(t *testing.T) {
	lut := NewPresentationLUT("2.25.105")
	lut.LUTExplanation = "canonical"
	if err := lut.SetLUT(2, 0, 12, []uint16{11, 22}); err != nil {
		t.Fatalf("SetLUT() error = %v", err)
	}

	ds, err := lut.ToDataset()
	if err != nil {
		t.Fatalf("ToDataset() error = %v", err)
	}
	sequence, err := ds.GetSequence(tag.PresentationLUTSequence)
	if err != nil {
		t.Fatalf("PresentationLUTSequence missing: %v", err)
	}
	if sequence.Count() != 1 {
		t.Fatalf("PresentationLUTSequence item count = %d, want 1", sequence.Count())
	}
	item := sequence.GetItem(0)
	if got, err := item.GetUInt16s(tag.LUTDescriptor); err != nil || !reflect.DeepEqual(got, []uint16{2, 0, 12}) {
		t.Fatalf("nested LUT Descriptor = %v, %v", got, err)
	}
	if got, ok := item.GetString(tag.LUTExplanation); !ok || got != "canonical" {
		t.Fatalf("nested LUT Explanation = %q, %v", got, ok)
	}
	if got, err := item.GetUInt16s(tag.LUTData); err != nil || !reflect.DeepEqual(got, []uint16{11, 22}) {
		t.Fatalf("nested LUT Data = %v, %v", got, err)
	}
	for _, forbidden := range []*tag.Tag{tag.LUTDescriptor, tag.LUTExplanation, tag.LUTData} {
		if _, ok := ds.Get(forbidden); ok {
			t.Errorf("top-level Dataset unexpectedly contains %s", forbidden)
		}
	}
}

func TestPresentationLUTDatasetReadsLegacyTopLevelValues(t *testing.T) {
	ds := dataset.New()
	legacy := []element.Element{
		element.NewString(tag.SOPClassUID, vr.UI, []string{presentationLUTSOPClassUID}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{"2.25.106"}),
		element.NewUnsignedShort(tag.LUTDescriptor, []uint16{2, 0, 12}),
		element.NewString(tag.LUTExplanation, vr.LO, []string{"legacy"}),
		element.NewUnsignedShort(tag.LUTData, []uint16{33, 44}),
		element.NewString(tag.PresentationLUTShape, vr.CS, []string{string(PresentationLUTShapeIdentity)}),
	}
	for _, value := range legacy {
		if err := ds.Add(value); err != nil {
			t.Fatalf("add legacy element %s: %v", value.Tag(), err)
		}
	}

	lut, err := NewPresentationLUTFromDataset("", ds)
	if err != nil {
		t.Fatalf("NewPresentationLUTFromDataset() error = %v", err)
	}
	if !reflect.DeepEqual(lut.LUTDescriptor, []uint16{2, 0, 12}) ||
		lut.LUTExplanation != "legacy" ||
		!reflect.DeepEqual(lut.LUTData, []uint16{33, 44}) {
		t.Fatalf("legacy LUT values = %#v", lut)
	}
}

func TestPresentationLUTDatasetRejectsMalformedCanonicalSequence(t *testing.T) {
	ds := dataset.New()
	ds.SetAutoValidate(false)
	for _, value := range []element.Element{
		element.NewString(tag.SOPClassUID, vr.UI, []string{presentationLUTSOPClassUID}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{"2.25.107"}),
		element.NewString(tag.PresentationLUTSequence, vr.LO, []string{"not a sequence"}),
		element.NewUnsignedShort(tag.LUTDescriptor, []uint16{2, 0, 12}),
		element.NewUnsignedShort(tag.LUTData, []uint16{33, 44}),
	} {
		if err := ds.Add(value); err != nil {
			t.Fatalf("add malformed element %s: %v", value.Tag(), err)
		}
	}

	if _, err := NewPresentationLUTFromDataset("", ds); err == nil {
		t.Fatal("NewPresentationLUTFromDataset() accepted a malformed canonical sequence")
	}
}

func assertReferenceSequence(t *testing.T, ds *dataset.Dataset, sequenceTag *tag.Tag, want []SOPReference) {
	t.Helper()
	sequence, err := ds.GetSequence(sequenceTag)
	if err != nil {
		t.Fatalf("%s missing: %v", sequenceTag, err)
	}
	if sequence.Count() != len(want) {
		t.Fatalf("%s item count = %d, want %d", sequenceTag, sequence.Count(), len(want))
	}
	for index, reference := range want {
		item := sequence.GetItem(index)
		classUID, classOK := item.GetString(tag.ReferencedSOPClassUID)
		instanceUID, instanceOK := item.GetString(tag.ReferencedSOPInstanceUID)
		if !classOK || !instanceOK || classUID != reference.SOPClassUID || instanceUID != reference.SOPInstanceUID {
			t.Errorf("%s item %d = %q/%q, want %q/%q", sequenceTag, index, classUID, instanceUID, reference.SOPClassUID, reference.SOPInstanceUID)
		}
	}
}
