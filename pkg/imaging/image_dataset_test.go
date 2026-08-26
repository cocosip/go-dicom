// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/color"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
	"github.com/cocosip/go-dicom/pkg/imaging/transform"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestOpenDicomImageRendersFile(t *testing.T) {
	dicomImage, err := OpenDicomImage(filepath.Join("..", "..", "test-data", "TestPattern_RGB.dcm"))
	if err != nil {
		t.Fatalf("OpenDicomImage() error = %v", err)
	}
	if dicomImage.Width() == 0 || dicomImage.Height() == 0 {
		t.Fatalf("image dimensions = %dx%d", dicomImage.Width(), dicomImage.Height())
	}
	if _, err := dicomImage.RenderFrameImage(0); err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
}

func TestNewDicomImageFromParseResultUsesParsedDataset(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})

	dicomImage, err := NewDicomImageFromParseResult(&parser.ParseResult{Dataset: ds})
	if err != nil {
		t.Fatalf("NewDicomImageFromParseResult() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got, want := color.GrayModel.Convert(rendered.At(1, 0)).(color.Gray).Y, uint8(255); got != want {
		t.Fatalf("pixel 1 = %d, want %d", got, want)
	}
}

func TestNewDicomImageFromParseResultRejectsNil(t *testing.T) {
	if _, err := NewDicomImageFromParseResult(nil); err == nil {
		t.Fatal("NewDicomImageFromParseResult(nil) succeeded")
	}
}

func TestNewDicomImageFromDatasetRendersNativePixels(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if got, want := dicomImage.Width(), uint16(2); got != want {
		t.Fatalf("Width() = %d, want %d", got, want)
	}
	if got, want := dicomImage.Height(), uint16(1); got != want {
		t.Fatalf("Height() = %d, want %d", got, want)
	}

	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got, want := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y, uint8(0); got != want {
		t.Fatalf("pixel 0 = %d, want %d", got, want)
	}
	if got, want := color.GrayModel.Convert(rendered.At(1, 0)).(color.Gray).Y, uint8(255); got != want {
		t.Fatalf("pixel 1 = %d, want %d", got, want)
	}
}

func TestDatasetRescaleAndWindowControlRendering(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 100, 200})
	for _, elem := range []element.Element{
		element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}),
		element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{-100}),
		element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{0}),
		element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{200}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}

	want := []uint8{0, 255, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestFunctionalGroupsBuildFrameSpecificPipelines(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{50, 50})
	if err := ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"})); err != nil {
		t.Fatalf("add NumberOfFrames: %v", err)
	}

	sharedTransform := dataset.New()
	for _, elem := range []element.Element{
		element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}),
		element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}),
	} {
		if err := sharedTransform.Add(elem); err != nil {
			t.Fatalf("add shared transform value: %v", err)
		}
	}
	sharedTransformSequence := dataset.NewSequenceWithItems(tag.PixelValueTransformationSequence, []*dataset.Dataset{sharedTransform})
	sharedItem := dataset.New()
	if err := sharedItem.Add(sharedTransformSequence); err != nil {
		t.Fatalf("add PixelValueTransformationSequence: %v", err)
	}
	if err := ds.Add(dataset.NewSequenceWithItems(tag.SharedFunctionalGroupsSequence, []*dataset.Dataset{sharedItem})); err != nil {
		t.Fatalf("add SharedFunctionalGroupsSequence: %v", err)
	}

	perFrameItems := make([]*dataset.Dataset, 0, 2)
	for _, center := range []float64{50, 200} {
		frameVOI := dataset.New()
		if err := frameVOI.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{center})); err != nil {
			t.Fatalf("add frame WindowCenter: %v", err)
		}
		if err := frameVOI.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100})); err != nil {
			t.Fatalf("add frame WindowWidth: %v", err)
		}
		frameItem := dataset.New()
		if err := frameItem.Add(dataset.NewSequenceWithItems(tag.FrameVOILUTSequence, []*dataset.Dataset{frameVOI})); err != nil {
			t.Fatalf("add FrameVOILUTSequence: %v", err)
		}
		perFrameItems = append(perFrameItems, frameItem)
	}
	if err := ds.Add(dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, perFrameItems)); err != nil {
		t.Fatalf("add PerFrameFunctionalGroupsSequence: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	want := []uint8{255, 0}
	for frame, expected := range want {
		rendered, err := dicomImage.RenderFrameImage(frame)
		if err != nil {
			t.Fatalf("RenderFrameImage(%d) error = %v", frame, err)
		}
		if got := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("frame %d pixel = %d, want %d", frame, got, expected)
		}
	}
}

func TestIncompleteTopLevelWindowFallsBackToFunctionalGroupPair(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{100})
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{10})); err != nil {
		t.Fatalf("add top-level WindowCenter: %v", err)
	}
	frameVOI := dataset.New()
	_ = frameVOI.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{100}))
	_ = frameVOI.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{200}))
	frameItem := dataset.New()
	_ = frameItem.Add(dataset.NewSequenceWithItems(tag.FrameVOILUTSequence, []*dataset.Dataset{frameVOI}))
	if err := ds.Add(dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, []*dataset.Dataset{frameItem})); err != nil {
		t.Fatalf("add PerFrameFunctionalGroupsSequence: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if got, want := dicomImage.WindowCenter(), 100.0; got != want {
		t.Fatalf("WindowCenter() = %v, want functional-group value %v", got, want)
	}
	if got, want := dicomImage.WindowWidth(), 200.0; got != want {
		t.Fatalf("WindowWidth() = %v, want functional-group value %v", got, want)
	}
}

func TestPerFrameFunctionalGroupVOILUTSequenceIsUsed(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{1}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{1}))
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8}))
	_ = lutItem.Add(element.NewOtherByte(tag.LUTData, []byte{0, 100, 200}))
	frameVOI := dataset.New()
	_ = frameVOI.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem}))
	frameItem := dataset.New()
	_ = frameItem.Add(dataset.NewSequenceWithItems(tag.FrameVOILUTSequence, []*dataset.Dataset{frameVOI}))
	if err := ds.Add(dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, []*dataset.Dataset{frameItem})); err != nil {
		t.Fatalf("add PerFrameFunctionalGroupsSequence: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{0, 128, 255} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d", x, got, want)
		}
	}
}

func TestDatasetVOILUTFunctionControlsWindowing(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{125})
	for _, elem := range []element.Element{
		element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{100}),
		element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100}),
		element.NewString(tag.VOILUTFunction, vr.CS, []string{"SIGMOID"}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got, want := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y, uint8(186); got != want {
		t.Fatalf("SIGMOID pixel = %d, want %d", got, want)
	}
}

func TestModalityLUTSequencePrecedesRescale(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	lutItem := dataset.New()
	if err := lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 16})); err != nil {
		t.Fatalf("add LUTDescriptor: %v", err)
	}
	lutBytes := make([]byte, 6)
	for i, value := range []uint16{0, 10, 20} {
		binary.LittleEndian.PutUint16(lutBytes[i*2:], value)
	}
	if err := lutItem.Add(element.NewOtherWord(tag.LUTData, lutBytes)); err != nil {
		t.Fatalf("add LUTData: %v", err)
	}
	if err := ds.Add(dataset.NewSequenceWithItems(tag.ModalityLUTSequence, []*dataset.Dataset{lutItem})); err != nil {
		t.Fatalf("add ModalityLUTSequence: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{100})); err != nil {
		t.Fatalf("add RescaleSlope: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{10})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{20})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	want := []uint8{0, 134, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestDatasetImageReadsBigEndianModalityLUTData(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	ds.SetInternalTransferSyntax(transfer.ExplicitVRBigEndian)
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 16}))
	lutBytes := make([]byte, 6)
	for index, value := range []uint16{0, 100, 200} {
		binary.BigEndian.PutUint16(lutBytes[index*2:], value)
	}
	_ = lutItem.Add(element.NewOtherWord(tag.LUTData, lutBytes))
	_ = ds.Add(dataset.NewSequenceWithItems(tag.ModalityLUTSequence, []*dataset.Dataset{lutItem}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{100}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{200}))

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	want := []uint8{0, 128, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestVOILUTSequencePrecedesWindow(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	lutItem := dataset.New()
	if err := lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8})); err != nil {
		t.Fatalf("add LUTDescriptor: %v", err)
	}
	if err := lutItem.Add(element.NewOtherByte(tag.LUTData, []byte{0, 100, 200})); err != nil {
		t.Fatalf("add LUTData: %v", err)
	}
	if err := ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem})); err != nil {
		t.Fatalf("add VOILUTSequence: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{1})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{1})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	want := []uint8{0, 128, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}

	dicomImage.SetUseVOILUT(false)
	rendered, err = dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) without VOI LUT error = %v", err)
	}
	want = []uint8{0, 255, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("windowed pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestVOILUTSequenceNormalizes16BitEntries(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	lutItem := dataset.New()
	if err := lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 16})); err != nil {
		t.Fatalf("add LUTDescriptor: %v", err)
	}
	lutBytes := make([]byte, 6)
	for index, value := range []uint16{0, 32768, 65535} {
		binary.LittleEndian.PutUint16(lutBytes[index*2:], value)
	}
	if err := lutItem.Add(element.NewOtherWord(tag.LUTData, lutBytes)); err != nil {
		t.Fatalf("add LUTData: %v", err)
	}
	if err := ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem})); err != nil {
		t.Fatalf("add VOILUTSequence: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	want := []uint8{0, 128, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestDatasetImageReadsBigEndianVOILUTData(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	ds.SetInternalTransferSyntax(transfer.ExplicitVRBigEndian)
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 16}))
	lutBytes := make([]byte, 6)
	for index, value := range []uint16{0, 32768, 65535} {
		binary.BigEndian.PutUint16(lutBytes[index*2:], value)
	}
	_ = lutItem.Add(element.NewOtherWord(tag.LUTData, lutBytes))
	_ = ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem}))

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	want := []uint8{0, 128, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestDefaultWindowUsesModalityValueRange(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 50, 100})
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2})); err != nil {
		t.Fatalf("add RescaleSlope: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	want := []uint8{0, 128, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestDefaultWindowUsesImagePixelValueTagsBeforePixelRange(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{50, 75, 100})
	if err := ds.Add(element.NewUnsignedShort(tag.SmallestImagePixelValue, []uint16{0})); err != nil {
		t.Fatalf("add SmallestImagePixelValue: %v", err)
	}
	if err := ds.Add(element.NewUnsignedShort(tag.LargestImagePixelValue, []uint16{200})); err != nil {
		t.Fatalf("add LargestImagePixelValue: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2})); err != nil {
		t.Fatalf("add RescaleSlope: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if got, want := dicomImage.WindowCenter(), 200.0; got != want {
		t.Fatalf("WindowCenter() = %v, want %v", got, want)
	}
	if got, want := dicomImage.WindowWidth(), 400.0; got != want {
		t.Fatalf("WindowWidth() = %v, want %v", got, want)
	}
}

func TestDefaultWindowExcludesDatasetPixelPadding(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 50, 100})
	if err := ds.Add(element.NewUnsignedShort(tag.PixelPaddingValue, []uint16{0})); err != nil {
		t.Fatalf("add PixelPaddingValue: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if got, want := dicomImage.WindowCenter(), 75.0; got != want {
		t.Fatalf("WindowCenter() = %v, want %v", got, want)
	}
	if got, want := dicomImage.WindowWidth(), 50.0; got != want {
		t.Fatalf("WindowWidth() = %v, want %v", got, want)
	}
}

func TestRenderFrameImageAppliesScale(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	dicomImage.SetScale(2)

	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got, want := rendered.Bounds().Dx(), 4; got != want {
		t.Fatalf("rendered width = %d, want %d", got, want)
	}
	if got, want := rendered.Bounds().Dy(), 2; got != want {
		t.Fatalf("rendered height = %d, want %d", got, want)
	}
	got := color.GrayModel.Convert(rendered.At(1, 0)).(color.Gray).Y
	if got == 0 || got == 255 {
		t.Fatalf("interpolated pixel = %d, want a value between the source pixels", got)
	}
}

func TestRenderFrameImageWithOptionsAppliesSpatialTransformInsteadOfLegacyScale(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	dicomImage.SetScale(2)
	spatial := transform.NewSpatialTransform()
	spatial.Rotate(90)

	rendered, err := dicomImage.RenderFrameImageWithOptions(0, FrameRenderOptions{SpatialTransform: spatial})
	if err != nil {
		t.Fatalf("RenderFrameImageWithOptions() error = %v", err)
	}
	if got, want := rendered.Bounds().Size().X, 1; got != want {
		t.Fatalf("rendered width = %d, want %d", got, want)
	}
	if got, want := rendered.Bounds().Size().Y, 2; got != want {
		t.Fatalf("rendered height = %d, want %d", got, want)
	}
}

func TestDatasetImageRendersPackedOneBitPixels(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	elements := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{8}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{1}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{1}),
		element.NewUnsignedShort(tag.HighBit, []uint16{0}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{monochrome2}),
		element.NewOtherByte(tag.PixelData, []byte{0b10101010}),
	}
	for _, elem := range elements {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	want := []uint8{0, 255, 0, 255, 0, 255, 0, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestDatasetImageRenders32BitGrayscalePixels(t *testing.T) {
	tests := []struct {
		name           string
		representation uint16
		values         []uint32
	}{
		{name: "unsigned", representation: 0, values: []uint32{0, 1 << 31, ^uint32(0)}},
		{name: "signed", representation: 1, values: []uint32{1 << 31, 0, uint32(1<<31) - 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pixels := make([]byte, 4*len(tt.values))
			for index, value := range tt.values {
				binary.LittleEndian.PutUint32(pixels[index*4:], value)
			}
			ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
			elements := []element.Element{
				element.NewUnsignedShort(tag.Rows, []uint16{1}),
				element.NewUnsignedShort(tag.Columns, []uint16{uint16(len(tt.values))}),
				element.NewUnsignedShort(tag.BitsAllocated, []uint16{32}),
				element.NewUnsignedShort(tag.BitsStored, []uint16{32}),
				element.NewUnsignedShort(tag.HighBit, []uint16{31}),
				element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
				element.NewUnsignedShort(tag.PixelRepresentation, []uint16{tt.representation}),
				element.NewString(tag.PhotometricInterpretation, vr.CS, []string{monochrome2}),
				element.NewOtherWord(tag.PixelData, pixels),
			}
			for _, elem := range elements {
				if err := ds.Add(elem); err != nil {
					t.Fatalf("add %s: %v", elem.Tag(), err)
				}
			}

			dicomImage, err := NewDicomImageFromDataset(ds)
			if err != nil {
				t.Fatalf("NewDicomImageFromDataset() error = %v", err)
			}
			rendered, err := dicomImage.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("RenderFrameImage(0) error = %v", err)
			}
			want := []uint8{0, 128, 255}
			for x, expected := range want {
				if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
					t.Fatalf("pixel %d = %d, want %d", x, got, expected)
				}
			}
		})
	}
}

func TestDatasetImageRendersExplicitVRBigEndian32BitPixels(t *testing.T) {
	pixels := make([]byte, 12)
	for index, value := range []uint32{0, 1 << 31, ^uint32(0)} {
		binary.BigEndian.PutUint32(pixels[index*4:], value)
	}
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	elements := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{3}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{32}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{32}),
		element.NewUnsignedShort(tag.HighBit, []uint16{31}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{monochrome2}),
		element.NewOtherWord(tag.PixelData, pixels),
	}
	for _, elem := range elements {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	want := []uint8{0, 128, 255}
	for x, expected := range want {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != expected {
			t.Fatalf("pixel %d = %d, want %d", x, got, expected)
		}
	}
}

func TestDatasetImageAutomaticallyDecodesEncapsulatedPixels(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 0})
	ds.SetInternalTransferSyntax(transfer.JPEG2000Lossless)
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0, 255}))
	if err := ds.AddOrUpdate(fragments); err != nil {
		t.Fatalf("replace PixelData: %v", err)
	}
	registry := codec.NewCodecRegistry()
	registry.RegisterCodec(transfer.JPEG2000Lossless, imagePassthroughCodec{})

	dicomImage, err := NewDicomImageFromDataset(ds, WithImageCodecRegistry(registry))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got := color.GrayModel.Convert(rendered.At(1, 0)).(color.Gray).Y; got != 255 {
		t.Fatalf("pixel 1 = %d, want 255", got)
	}
	if ds.InternalTransferSyntax() != transfer.JPEG2000Lossless {
		t.Fatalf("source transfer syntax = %v, want JPEG 2000 Lossless", ds.InternalTransferSyntax())
	}
}

func TestDatasetImageDecodeDoesNotMutateSourceFragmentsWithoutBOT(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 0, 0, 0})
	ds.SetInternalTransferSyntax(transfer.JPEG2000Lossless)
	if err := ds.Add(element.NewIntegerStringFromInt(tag.NumberOfFrames, []int{2})); err != nil {
		t.Fatalf("add NumberOfFrames: %v", err)
	}
	firstFragment := []byte{1, 2}
	secondFragment := []byte{3, 4}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory(firstFragment))
	fragments.AddFragment(buffer.NewMemory(secondFragment))
	if err := ds.AddOrUpdate(fragments); err != nil {
		t.Fatalf("replace PixelData: %v", err)
	}
	registry := codec.NewCodecRegistry()
	registry.RegisterCodec(transfer.JPEG2000Lossless, mutatingImageCodec{})

	if _, err := NewDicomImageFromDataset(ds, WithImageCodecRegistry(registry)); err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if got, want := firstFragment, []byte{1, 2}; !bytes.Equal(got, want) {
		t.Fatalf("source fragment 0 = %v, want %v", got, want)
	}
	if got, want := secondFragment, []byte{3, 4}; !bytes.Equal(got, want) {
		t.Fatalf("source fragment 1 = %v, want %v", got, want)
	}
}

func TestDatasetImageDecodesEncapsulatedPaletteBeforeRGBConversion(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.JPEG2000Lossless)
	elements := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricPaletteColor}),
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewOtherByte(tag.RedPaletteColorLookupTableData, []byte{0, 255}),
		element.NewOtherByte(tag.GreenPaletteColorLookupTableData, []byte{0, 0}),
		element.NewOtherByte(tag.BluePaletteColorLookupTableData, []byte{0, 0}),
	}
	for _, elem := range elements {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0, 1}))
	if err := ds.Add(fragments); err != nil {
		t.Fatalf("add PixelData: %v", err)
	}
	registry := codec.NewCodecRegistry()
	registry.RegisterCodec(transfer.JPEG2000Lossless, paletteIndexCodec{})

	dicomImage, err := NewDicomImageFromDataset(ds, WithImageCodecRegistry(registry))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got := color.RGBAModel.Convert(rendered.At(0, 0)).(color.RGBA); got != (color.RGBA{A: 255}) {
		t.Fatalf("palette pixel 0 = %#v, want opaque black", got)
	}
	if got := color.RGBAModel.Convert(rendered.At(1, 0)).(color.RGBA); got != (color.RGBA{R: 255, A: 255}) {
		t.Fatalf("palette pixel 1 = %#v, want opaque red", got)
	}
}

func TestDatasetImageRendersExplicitOverlayWithOriginAndVisibility(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 4, 4, make([]byte, 16))
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{127.5})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{255})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}
	groupTag := func(elementNumber uint16) *tag.Tag { return tag.New(0x6000, elementNumber) }
	overlayElements := []element.Element{
		element.NewUnsignedShort(groupTag(0x0010), []uint16{2}),
		element.NewUnsignedShort(groupTag(0x0011), []uint16{2}),
		element.NewSignedShort(groupTag(0x0050), []int16{2, 2}),
		element.NewOtherWord(groupTag(0x3000), []byte{0b00000001, 0}),
	}
	for _, elem := range overlayElements {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add overlay %s: %v", elem.Tag(), err)
		}
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	dicomImage.SetOverlayColor(imagetypes.Color32{A: 255, R: 255})
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got := color.RGBAModel.Convert(rendered.At(1, 1)).(color.RGBA); got != (color.RGBA{R: 255, A: 255}) {
		t.Fatalf("overlay pixel = %#v, want opaque red", got)
	}
	if got := color.RGBAModel.Convert(rendered.At(0, 0)).(color.RGBA); got.R != 0 {
		t.Fatalf("pixel outside overlay red = %d, want 0", got.R)
	}

	dicomImage.SetShowOverlays(false)
	rendered, err = dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) without overlays error = %v", err)
	}
	if got := color.RGBAModel.Convert(rendered.At(1, 1)).(color.RGBA); got.R != 0 {
		t.Fatalf("hidden overlay pixel red = %d, want 0", got.R)
	}
}

func TestDatasetImageExtractsEmbeddedOverlayBit(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	groupTag := func(elementNumber uint16) *tag.Tag { return tag.New(0x6000, elementNumber) }
	elements := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{12}),
		element.NewUnsignedShort(tag.HighBit, []uint16{11}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{monochrome2}),
		element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{2047.5}),
		element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{4095}),
		element.NewOtherWord(tag.PixelData, []byte{0, 0, 0, 0x80}),
		element.NewUnsignedShort(groupTag(0x0010), []uint16{1}),
		element.NewUnsignedShort(groupTag(0x0011), []uint16{2}),
		element.NewSignedShort(groupTag(0x0050), []int16{1, 1}),
		element.NewUnsignedShort(groupTag(0x0100), []uint16{16}),
		element.NewUnsignedShort(groupTag(0x0102), []uint16{15}),
	}
	for _, elem := range elements {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	dicomImage.SetOverlayColor(imagetypes.Color32{A: 255, R: 255})
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got := color.RGBAModel.Convert(rendered.At(0, 0)).(color.RGBA).R; got != 0 {
		t.Fatalf("pixel 0 red = %d, want 0", got)
	}
	if got := color.RGBAModel.Convert(rendered.At(1, 0)).(color.RGBA); got != (color.RGBA{R: 255, A: 255}) {
		t.Fatalf("embedded overlay pixel = %#v, want opaque red", got)
	}
}

func TestWindowOverridesApplyToAllFramesUnlessDisabled(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255, 0, 255})
	if err := ds.Add(element.NewIntegerStringFromInt(tag.NumberOfFrames, []int{2})); err != nil {
		t.Fatalf("add NumberOfFrames: %v", err)
	}
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}

	dicomImage.SetWindow(10, 20)
	if err := dicomImage.SetCurrentFrame(1); err != nil {
		t.Fatalf("SetCurrentFrame(1) error = %v", err)
	}
	if got, want := dicomImage.WindowCenter(), 10.0; got != want {
		t.Fatalf("frame 1 WindowCenter() = %v, want %v", got, want)
	}
	if got, want := dicomImage.WindowWidth(), 20.0; got != want {
		t.Fatalf("frame 1 WindowWidth() = %v, want %v", got, want)
	}

	dicomImage.SetAutoApplyLUTToAllFrames(false)
	dicomImage.SetWindow(30, 40)
	if err := dicomImage.SetCurrentFrame(0); err != nil {
		t.Fatalf("SetCurrentFrame(0) error = %v", err)
	}
	if got, want := dicomImage.WindowCenter(), 10.0; got != want {
		t.Fatalf("frame 0 WindowCenter() = %v, want %v", got, want)
	}
}

func TestLegacyPixelDataConstructorKeepsWindowBehavior(t *testing.T) {
	info := &PixelDataInfo{
		Width: 2, Height: 1, NumberOfFrames: 2,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PixelRepresentation: UnsignedPixels, PhotometricInterpretation: Monochrome2,
	}
	pixelData, err := NewDicomPixelDataFromBytes(info, []byte{0, 100, 200, 255})
	if err != nil {
		t.Fatalf("NewDicomPixelDataFromBytes() error = %v", err)
	}
	dicomImage := NewDicomImage(pixelData)
	if got, want := dicomImage.WindowCenter(), 50.0; got != want {
		t.Fatalf("WindowCenter() = %v, want first-frame value %v", got, want)
	}
	dicomImage.SetWindow(10, 20)
	if err := dicomImage.SetCurrentFrame(1); err != nil {
		t.Fatalf("SetCurrentFrame(1) error = %v", err)
	}
	if got, want := dicomImage.WindowCenter(), 50.0; got != want {
		t.Fatalf("frame 1 WindowCenter() = %v, want untouched value %v", got, want)
	}
}

func TestCallerGrayscaleColorMapChangesRenderedColors(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	var colorMap [256]imagetypes.Color32
	for index := range colorMap {
		colorMap[index] = imagetypes.Color32{A: 255, G: uint8(index)}
	}
	dicomImage.SetGrayscaleColorMap(colorMap)

	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got := color.RGBAModel.Convert(rendered.At(1, 0)).(color.RGBA); got != (color.RGBA{G: 255, A: 255}) {
		t.Fatalf("mapped pixel = %#v, want opaque green", got)
	}
}

func TestClonePreservesDatasetDrivenRendering(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 50, 100})
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2})); err != nil {
		t.Fatalf("add RescaleSlope: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{50})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}
	if err := ds.Add(element.NewString(tag.VOILUTFunction, vr.CS, []string{"SIGMOID"})); err != nil {
		t.Fatalf("add VOILUTFunction: %v", err)
	}
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}

	original, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("original RenderFrameImage(0) error = %v", err)
	}
	cloned, err := dicomImage.Clone().RenderFrameImage(0)
	if err != nil {
		t.Fatalf("clone RenderFrameImage(0) error = %v", err)
	}
	for x := 0; x < 3; x++ {
		got := color.GrayModel.Convert(cloned.At(x, 0)).(color.Gray).Y
		want := color.GrayModel.Convert(original.At(x, 0)).(color.Gray).Y
		if got != want {
			t.Fatalf("clone pixel %d = %d, want %d", x, got, want)
		}
	}
}

func TestClonePreservesCallerPipelineOverrides(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8}))
	_ = lutItem.Add(element.NewOtherByte(tag.LUTData, []byte{0, 100, 200}))
	_ = ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{1}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{1}))

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	dicomImage.SetWindow(1, 3)
	dicomImage.SetInvert(true)
	dicomImage.SetUseVOILUT(false)
	original, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("original RenderFrameImage(0) error = %v", err)
	}

	clone := dicomImage.Clone()
	if got, want := clone.WindowCenter(), 1.0; got != want {
		t.Fatalf("clone WindowCenter() = %v, want %v", got, want)
	}
	if got, want := clone.WindowWidth(), 3.0; got != want {
		t.Fatalf("clone WindowWidth() = %v, want %v", got, want)
	}
	if !clone.Invert() {
		t.Fatal("clone Invert() = false, want true")
	}
	if clone.UseVOILUT() {
		t.Fatal("clone UseVOILUT() = true, want false")
	}
	cloned, err := clone.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("clone RenderFrameImage(0) error = %v", err)
	}
	for x := 0; x < 3; x++ {
		got := color.GrayModel.Convert(cloned.At(x, 0)).(color.Gray).Y
		want := color.GrayModel.Convert(original.At(x, 0)).(color.Gray).Y
		if got != want {
			t.Fatalf("clone pixel %d = %d, want %d", x, got, want)
		}
	}
}

func TestCloneCopiesExplicitGrayscalePipeline(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	override := render.NewGrayscalePipeline(1, 0, 1, 3, 0, 2, true)
	if err := dicomImage.SetPipeline(0, override); err != nil {
		t.Fatalf("SetPipeline(0) error = %v", err)
	}

	clone := dicomImage.Clone()
	override.SetWindow(100, 1)
	override.SetInvert(false)
	if got, want := clone.WindowCenter(), 1.0; got != want {
		t.Fatalf("clone WindowCenter() = %v, want %v", got, want)
	}
	if got, want := clone.WindowWidth(), 3.0; got != want {
		t.Fatalf("clone WindowWidth() = %v, want %v", got, want)
	}
	if !clone.Invert() {
		t.Fatal("clone Invert() = false after original mutation, want true")
	}
}

func TestCloneCallsCustomPipelineClonerOutsideImageLock(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	calledWithoutImageLock := false
	pipeline := &imageLockCheckingPipeline{
		image:                  dicomImage,
		delegate:               render.NewGrayscalePipeline(1, 0, 0, 1, 0, 255, false),
		calledWithoutImageLock: &calledWithoutImageLock,
	}
	if err := dicomImage.SetPipeline(0, pipeline); err != nil {
		t.Fatalf("SetPipeline(0) error = %v", err)
	}

	_ = dicomImage.Clone()
	if !calledWithoutImageLock {
		t.Fatal("ClonePipeline() was called while DicomImage.mu was held")
	}
}

type imagePassthroughCodec struct{}

func (imagePassthroughCodec) Name() string { return "image test passthrough" }
func (imagePassthroughCodec) TransferSyntax() *transfer.Syntax {
	return transfer.JPEG2000Lossless
}
func (imagePassthroughCodec) GetDefaultParameters() codec.Parameters {
	return codec.NewBaseParameters()
}
func (imagePassthroughCodec) Encode(imagetypes.PixelData, imagetypes.PixelData, codec.Parameters) error {
	return nil
}
func (imagePassthroughCodec) Decode(oldPixelData, newPixelData imagetypes.PixelData, _ codec.Parameters) error {
	for frame := 0; frame < oldPixelData.FrameCount(); frame++ {
		data, err := oldPixelData.GetFrame(frame)
		if err != nil {
			return err
		}
		if err := newPixelData.AddFrame(data); err != nil {
			return err
		}
	}
	return nil
}

type paletteIndexCodec struct{}

func (paletteIndexCodec) Name() string { return "palette index test codec" }
func (paletteIndexCodec) TransferSyntax() *transfer.Syntax {
	return transfer.JPEG2000Lossless
}
func (paletteIndexCodec) GetDefaultParameters() codec.Parameters {
	return codec.NewBaseParameters()
}
func (paletteIndexCodec) Encode(imagetypes.PixelData, imagetypes.PixelData, codec.Parameters) error {
	return nil
}
func (paletteIndexCodec) Decode(oldPixelData, newPixelData imagetypes.PixelData, _ codec.Parameters) error {
	info := oldPixelData.GetFrameInfo()
	if info.PhotometricInterpretation != photometricPaletteColor || info.SamplesPerPixel != 1 {
		return fmt.Errorf("codec input metadata = %s/%d, want PALETTE COLOR/1", info.PhotometricInterpretation, info.SamplesPerPixel)
	}
	frame, err := oldPixelData.GetFrame(0)
	if err != nil {
		return err
	}
	if !bytes.Equal(frame, []byte{0, 1}) {
		return fmt.Errorf("codec input frame = %v, want [0 1]", frame)
	}
	return newPixelData.AddFrame(frame)
}

type mutatingImageCodec struct{}

func (mutatingImageCodec) Name() string { return "mutating image test codec" }
func (mutatingImageCodec) TransferSyntax() *transfer.Syntax {
	return transfer.JPEG2000Lossless
}
func (mutatingImageCodec) GetDefaultParameters() codec.Parameters {
	return codec.NewBaseParameters()
}
func (mutatingImageCodec) Encode(imagetypes.PixelData, imagetypes.PixelData, codec.Parameters) error {
	return nil
}
func (mutatingImageCodec) Decode(oldPixelData, newPixelData imagetypes.PixelData, _ codec.Parameters) error {
	for frameIndex := 0; frameIndex < oldPixelData.FrameCount(); frameIndex++ {
		frame, err := oldPixelData.GetFrame(frameIndex)
		if err != nil {
			return err
		}
		frame[0] ^= 0xff
		if err := newPixelData.AddFrame(frame); err != nil {
			return err
		}
	}
	return nil
}

type imageLockCheckingPipeline struct {
	image                  *DicomImage
	delegate               render.Pipeline
	calledWithoutImageLock *bool
}

func (p *imageLockCheckingPipeline) LUT() render.LUT { return p.delegate.LUT() }
func (p *imageLockCheckingPipeline) ClearCache()     { p.delegate.ClearCache() }
func (p *imageLockCheckingPipeline) ClonePipeline() render.Pipeline {
	if p.image.mu.TryLock() {
		*p.calledWithoutImageLock = true
		p.image.mu.Unlock()
	}
	return p.delegate
}

func newNativeMonochromeDataset(t *testing.T, width, height uint16, pixels []byte) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	elements := []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{height}),
		element.NewUnsignedShort(tag.Columns, []uint16{width}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{monochrome2}),
		element.NewOtherByte(tag.PixelData, pixels),
	}
	for _, elem := range elements {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}
	return ds
}
