// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestNewImageDataReadsClassicGeometryAndModalityValues(t *testing.T) {
	ds := testClassicImageDataset(t, uid.CTImageStorage.UID(), "1.2.3.1", 5, []uint16{1, 2, 3, 4})
	mustAddElement(t, ds,
		element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}),
		element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{-100}),
	)

	image, err := NewImageData(ds, 0)
	if err != nil {
		t.Fatalf("NewImageData() error = %v", err)
	}
	if image.FrameIndex() != 0 || image.SourceSOPInstanceUID() != "1.2.3.1" {
		t.Fatalf("frame/source identity = %d/%q", image.FrameIndex(), image.SourceSOPInstanceUID())
	}
	frameGeometry := image.Geometry()
	if frameGeometry.TopLeft.Z != 5 || image.SortingPosition() != 5 {
		t.Fatalf("geometry/sorting position = %v/%v, want z=5", frameGeometry.TopLeft, image.SortingPosition())
	}
	value, valid, err := image.ValueAt(1, 1)
	if err != nil {
		t.Fatalf("ValueAt() error = %v", err)
	}
	if !valid || value != -92 {
		t.Fatalf("ValueAt(1,1) = %v/%v, want -92/true", value, valid)
	}
}

func TestNewImageDataReadsParserStyleRescaleValues(t *testing.T) {
	ds := testClassicImageDataset(t, uid.CTImageStorage.UID(), "1.2.3.11", 5, []uint16{1, 2, 3, 4})
	mustAddElement(t, ds,
		element.NewString(tag.RescaleSlope, vr.DS, []string{"2"}),
		element.NewString(tag.RescaleIntercept, vr.DS, []string{"-100"}),
	)

	image, err := NewImageData(ds, 0)
	if err != nil {
		t.Fatalf("NewImageData() error = %v", err)
	}
	value, valid, err := image.ValueAt(1, 1)
	if err != nil {
		t.Fatalf("ValueAt() error = %v", err)
	}
	if !valid || value != -92 {
		t.Fatalf("ValueAt(1,1) = %v/%v, want -92/true", value, valid)
	}
}

func TestNewImageDataFromParsedEnhancedFixture(t *testing.T) {
	path := filepath.Join("..", "..", "..", "test-data", "TestMultiFrame.dcm")
	result, err := parser.ParseFile(path)
	if err != nil {
		t.Fatalf("ParseFile(%q) error = %v", path, err)
	}

	images, err := NewImageDataFromDataset(result.Dataset)
	if err != nil {
		t.Fatalf("NewImageDataFromDataset(parsed fixture) error = %v", err)
	}
	if len(images) < 2 {
		t.Fatalf("parsed fixture frame count = %d, want at least 2", len(images))
	}
	if _, valid, err := images[len(images)-1].ValueAt(0, 0); err != nil || !valid {
		t.Fatalf("last parsed frame sample valid/error = %v/%v, want true/nil", valid, err)
	}
}

func TestNewImageDataFromDatasetExpandsEnhancedFrames(t *testing.T) {
	ds := testEnhancedImageDataset(t, uid.EnhancedCTImageStorage.UID(), []float64{3, 8}, []uint16{
		10, 20, 30, 40,
		50, 60, 70, 80,
	})

	images, err := NewImageDataFromDataset(ds)
	if err != nil {
		t.Fatalf("NewImageDataFromDataset() error = %v", err)
	}
	if len(images) != 2 {
		t.Fatalf("image count = %d, want 2", len(images))
	}
	firstGeometry, secondGeometry := images[0].Geometry(), images[1].Geometry()
	if firstGeometry.TopLeft.Z != 3 || secondGeometry.TopLeft.Z != 8 {
		t.Fatalf("frame positions = %v/%v, want 3/8", firstGeometry.TopLeft.Z, secondGeometry.TopLeft.Z)
	}
	value, valid, err := images[1].ValueAt(0, 0)
	if err != nil {
		t.Fatalf("ValueAt() error = %v", err)
	}
	if !valid || value != 50 {
		t.Fatalf("enhanced frame 1 ValueAt(0,0) = %v/%v, want 50/true", value, valid)
	}
	if images[0].pixelData != images[1].pixelData {
		t.Fatal("native Enhanced frames retained duplicate full pixel-data containers")
	}
}

func TestImageDataMetadataAccessorsReturnDefensiveCopies(t *testing.T) {
	ds := testClassicImageDataset(t, uid.CTImageStorage.UID(), "1.2.3.12", 5, []uint16{1, 2, 3, 4})
	reference := dataset.New()
	mustAddElement(t, reference, element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{"1.2.3.99"}))
	mustAddElement(t, ds, dataset.NewSequenceWithItems(tag.SourceImageSequence, []*dataset.Dataset{reference}))
	image, err := NewImageData(ds, 0)
	if err != nil {
		t.Fatalf("NewImageData() error = %v", err)
	}

	geometryCopy := image.Geometry()
	geometryCopy.TopLeft.Z = 99
	if image.Geometry().TopLeft.Z != 5 {
		t.Fatal("Geometry() exposed mutable image geometry")
	}

	datasetCopy := image.SourceDataset()
	datasetCopy.Remove(tag.SOPInstanceUID)
	if image.SourceDataset().TryGetString(tag.SOPInstanceUID) != "1.2.3.12" {
		t.Fatal("SourceDataset() exposed mutable image metadata")
	}
	sequence, err := datasetCopy.GetSequence(tag.SourceImageSequence)
	if err != nil {
		t.Fatalf("SourceDataset().GetSequence() error = %v", err)
	}
	sequence.GetItem(0).Remove(tag.ReferencedSOPInstanceUID)
	pixelData, ok := datasetCopy.GetOrNil(tag.PixelData).(*element.OtherWord)
	if !ok {
		t.Fatalf("SourceDataset() Pixel Data type = %T, want OtherWord", datasetCopy.GetOrNil(tag.PixelData))
	}
	pixelData.GetData()[0] = 0xff

	freshCopy := image.SourceDataset()
	freshSequence, err := freshCopy.GetSequence(tag.SourceImageSequence)
	if err != nil || freshSequence.GetItem(0).TryGetString(tag.ReferencedSOPInstanceUID) != "1.2.3.99" {
		t.Fatalf("SourceDataset() exposed nested Sequence mutation: sequence=%v error=%v", freshSequence, err)
	}
	freshPixelData := freshCopy.GetOrNil(tag.PixelData).(*element.OtherWord)
	if freshPixelData.GetData()[0] == 0xff {
		t.Fatal("SourceDataset() exposed mutable OtherWord data")
	}
}

func TestNewImageDataRejectsUnsupportedOrIncompleteInputs(t *testing.T) {
	t.Run("unsupported SOP class", func(t *testing.T) {
		ds := testClassicImageDataset(t, uid.SecondaryCaptureImageStorage.UID(), "1.2.3.2", 0, []uint16{1, 2, 3, 4})
		_, err := NewImageData(ds, 0)
		if err == nil || !strings.Contains(err.Error(), "SOP Class") {
			t.Fatalf("NewImageData() error = %v, want unsupported SOP Class", err)
		}
	})

	t.Run("missing frame of reference", func(t *testing.T) {
		ds := testClassicImageDataset(t, uid.MRImageStorage.UID(), "1.2.3.3", 0, []uint16{1, 2, 3, 4})
		ds.Remove(tag.FrameOfReferenceUID)
		_, err := NewImageData(ds, 0)
		if err == nil || !strings.Contains(err.Error(), "Frame of Reference") {
			t.Fatalf("NewImageData() error = %v, want Frame of Reference rejection", err)
		}
	})

	t.Run("CT MONOCHROME1", func(t *testing.T) {
		ds := testClassicImageDataset(t, uid.CTImageStorage.UID(), "1.2.3.4", 0, []uint16{1, 2, 3, 4})
		mustAddElement(t, ds, element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME1"}))
		_, err := NewImageData(ds, 0)
		if err == nil || !strings.Contains(err.Error(), "CT") || !strings.Contains(err.Error(), "MONOCHROME1") {
			t.Fatalf("NewImageData() error = %v, want CT MONOCHROME1 rejection", err)
		}
	})
}

func TestNewImageDataFromDatasetAcceptsConstantNonSpatialEnhancedDimension(t *testing.T) {
	ds := testEnhancedImageDataset(t, uid.EnhancedMRImageStorage.UID(), []float64{0, 1}, []uint16{
		1, 2, 3, 4,
		5, 6, 7, 8,
	})
	addDimensionIndexes(t, ds, tag.StackID, tag.InStackPositionNumber, tag.TemporalPositionIndex)
	addFrameContent(t, ds, []frameContentValues{
		{stackID: "STACK-1", indexes: []uint32{1, 1, 1}},
		{stackID: "STACK-1", indexes: []uint32{1, 2, 1}},
	})

	if _, err := NewImageDataFromDataset(ds); err != nil {
		t.Fatalf("NewImageDataFromDataset() error = %v, want constant non-spatial dimension accepted", err)
	}
}

func TestNewImageDataFromDatasetRejectsAmbiguousEnhancedDimensions(t *testing.T) {
	tests := []struct {
		name   string
		frames []frameContentValues
		want   string
	}{
		{
			name: "multiple stacks",
			frames: []frameContentValues{
				{stackID: "STACK-1", indexes: []uint32{1, 1, 1}},
				{stackID: "STACK-2", indexes: []uint32{2, 1, 1}},
			},
			want: "multiple Stack IDs",
		},
		{
			name: "varying temporal position",
			frames: []frameContentValues{
				{stackID: "STACK-1", indexes: []uint32{1, 1, 1}},
				{stackID: "STACK-1", indexes: []uint32{1, 2, 2}},
			},
			want: "non-spatial dimension",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ds := testEnhancedImageDataset(t, uid.EnhancedMRImageStorage.UID(), []float64{0, 1}, []uint16{
				1, 2, 3, 4,
				5, 6, 7, 8,
			})
			addDimensionIndexes(t, ds, tag.StackID, tag.InStackPositionNumber, tag.TemporalPositionIndex)
			addFrameContent(t, ds, test.frames)

			_, err := NewImageDataFromDataset(ds)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("NewImageDataFromDataset() error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestNewImageDataRejectsAmbiguousEnhancedDimensions(t *testing.T) {
	ds := testEnhancedImageDataset(t, uid.EnhancedMRImageStorage.UID(), []float64{0, 1}, []uint16{
		1, 2, 3, 4,
		5, 6, 7, 8,
	})
	addDimensionIndexes(t, ds, tag.StackID, tag.InStackPositionNumber, tag.TemporalPositionIndex)
	addFrameContent(t, ds, []frameContentValues{
		{stackID: "STACK-1", indexes: []uint32{1, 1, 1}},
		{stackID: "STACK-2", indexes: []uint32{2, 1, 1}},
	})

	_, err := NewImageData(ds, 0)
	if err == nil || !strings.Contains(err.Error(), "multiple Stack IDs") {
		t.Fatalf("NewImageData() error = %v, want multiple Stack IDs rejection", err)
	}
}

func TestNewImageDataFromDatasetRejectsExcessiveSourceDimensionsBeforePixelAllocation(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*dataset.Dataset)
		want   string
	}{
		{
			name: "frame count",
			mutate: func(ds *dataset.Dataset) {
				mustAddElement(t, ds, element.NewString(tag.NumberOfFrames, vr.IS, []string{"65536"}))
			},
			want: "frame count",
		},
		{
			name: "rows",
			mutate: func(ds *dataset.Dataset) {
				mustAddElement(t, ds, element.NewUnsignedShort(tag.Rows, []uint16{16385}))
			},
			want: "source dimensions",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ds := testClassicImageDataset(t, uid.MRImageStorage.UID(), "1.2.3.50", 0, []uint16{1, 2, 3, 4})
			test.mutate(ds)
			_, err := NewImageDataFromDataset(ds)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("NewImageDataFromDataset() error = %v, want containing %q", err, test.want)
			}
		})
	}
}

type frameContentValues struct {
	stackID string
	indexes []uint32
}

func addDimensionIndexes(t testing.TB, ds *dataset.Dataset, pointers ...*tag.Tag) {
	t.Helper()
	items := make([]*dataset.Dataset, len(pointers))
	for index, pointer := range pointers {
		items[index] = dataset.New()
		mustAddElement(t, items[index], element.NewAttributeTag(tag.DimensionIndexPointer, []*tag.Tag{pointer}))
	}
	mustAddElement(t, ds, dataset.NewSequenceWithItems(tag.DimensionIndexSequence, items))
}

func addFrameContent(t testing.TB, ds *dataset.Dataset, frames []frameContentValues) {
	t.Helper()
	sequence, err := ds.GetSequence(tag.PerFrameFunctionalGroupsSequence)
	if err != nil {
		t.Fatalf("GetSequence(PerFrameFunctionalGroupsSequence) error = %v", err)
	}
	for index, values := range frames {
		addFunctionalMacro(t, sequence.GetItem(index), tag.FrameContentSequence,
			element.NewString(tag.StackID, vr.SH, []string{values.stackID}),
			element.NewUnsignedLong(tag.DimensionIndexValues, values.indexes),
		)
	}
}

func testClassicImageDataset(t testing.TB, sopClass, sopInstance string, z float64, pixels []uint16) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	mustAddElement(t, ds,
		element.NewString(tag.SOPClassUID, vr.UI, []string{sopClass}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstance}),
		element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.100"}),
		element.NewString(tag.SeriesInstanceUID, vr.UI, []string{"1.2.3.200"}),
		element.NewString(tag.FrameOfReferenceUID, vr.UI, []string{"1.2.3.300"}),
		element.NewUnsignedShort(tag.Rows, []uint16{2}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewDecimalStringFromFloat(tag.PixelSpacing, []float64{1, 1}),
		element.NewDecimalStringFromFloat(tag.ImagePositionPatient, []float64{0, 0, z}),
		element.NewDecimalStringFromFloat(tag.ImageOrientationPatient, []float64{1, 0, 0, 0, 1, 0}),
		element.NewOtherWord(tag.PixelData, uint16Bytes(pixels)),
	)
	return ds
}

func testEnhancedImageDataset(t testing.TB, sopClass string, positions []float64, pixels []uint16) *dataset.Dataset {
	t.Helper()
	ds := testClassicImageDataset(t, sopClass, "1.2.3.9", 0, pixels)
	ds.Remove(tag.ImagePositionPatient)
	ds.Remove(tag.ImageOrientationPatient)
	ds.Remove(tag.PixelSpacing)
	mustAddElement(t, ds, element.NewIntegerStringFromInt(tag.NumberOfFrames, []int{len(positions)}))

	shared := dataset.New()
	addFunctionalMacro(t, shared, tag.PixelMeasuresSequence,
		element.NewDecimalStringFromFloat(tag.PixelSpacing, []float64{1, 1}))
	addFunctionalMacro(t, shared, tag.PlaneOrientationSequence,
		element.NewDecimalStringFromFloat(tag.ImageOrientationPatient, []float64{1, 0, 0, 0, 1, 0}))
	mustAddElement(t, ds, dataset.NewSequenceWithItems(tag.SharedFunctionalGroupsSequence, []*dataset.Dataset{shared}))

	frames := make([]*dataset.Dataset, len(positions))
	for index, z := range positions {
		frames[index] = dataset.New()
		addFunctionalMacro(t, frames[index], tag.PlanePositionSequence,
			element.NewDecimalStringFromFloat(tag.ImagePositionPatient, []float64{0, 0, z}))
	}
	mustAddElement(t, ds, dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, frames))
	return ds
}

func addFunctionalMacro(t testing.TB, target *dataset.Dataset, sequenceTag *tag.Tag, values ...element.Element) {
	t.Helper()
	item := dataset.New()
	mustAddElement(t, item, values...)
	mustAddElement(t, target, dataset.NewSequenceWithItems(sequenceTag, []*dataset.Dataset{item}))
}

func mustAddElement(t testing.TB, ds *dataset.Dataset, values ...element.Element) {
	t.Helper()
	for _, value := range values {
		if err := ds.AddOrUpdate(value); err != nil {
			t.Fatalf("add %s: %v", value.Tag(), err)
		}
	}
}

func uint16Bytes(values []uint16) []byte {
	data := make([]byte, len(values)*2)
	for index, value := range values {
		data[index*2] = byte(value)
		data[index*2+1] = byte(value >> 8)
	}
	return data
}
