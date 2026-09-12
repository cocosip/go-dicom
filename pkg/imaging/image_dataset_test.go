// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"math"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transcode"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/colorconv"
	"github.com/cocosip/go-dicom/pkg/imaging/lut"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
	"github.com/cocosip/go-dicom/pkg/imaging/render"
	"github.com/cocosip/go-dicom/pkg/imaging/transform"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

const testVOILUTFunctionSigmoid = "SIGMOID"

func addLegacyTestElement(ds *dataset.Dataset, elem element.Element) error {
	ds.SetAutoValidate(false)
	defer ds.SetAutoValidate(true)
	return ds.Add(elem)
}

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
		element.NewString(tag.RescaleType, vr.LO, []string{"HU"}),
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

func TestDatasetRescaleRendersWithoutRescaleType(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 100})
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

	for x, expected := range []uint8{0, 255} {
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
		element.NewString(tag.RescaleType, vr.LO, []string{"HU"}),
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
	_ = addLegacyTestElement(lutItem, element.NewOtherByte(tag.LUTData, []byte{0, 100, 200}))
	frameVOI := dataset.New()
	_ = addLegacyTestElement(frameVOI, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem}))
	frameItem := dataset.New()
	_ = addLegacyTestElement(frameItem, dataset.NewSequenceWithItems(tag.FrameVOILUTSequence, []*dataset.Dataset{frameVOI}))
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, []*dataset.Dataset{frameItem})); err != nil {
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
	for x, want := range []uint8{0, 100, 200} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d", x, got, want)
		}
	}
}

func TestDatasetImageStandardLUTModeRejectsOtherByteData(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 8}))
	_ = addLegacyTestElement(lutItem, element.NewOtherByte(tag.LUTData, []byte{255}))
	_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem}))

	image, err := NewDicomImageFromDataset(ds, WithLUTVRMode(pixeldata.LUTStandard))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if _, err := image.RenderFrameImage(0); err == nil {
		t.Fatal("RenderFrameImage() accepted VOI LUT Data with non-standard OB VR")
	}
}

func TestPerFrameFunctionalGroupWindowOverridesTopLevelWindow(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{50})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{50}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100}))

	frameVOI := dataset.New()
	_ = frameVOI.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{100}))
	_ = frameVOI.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100}))
	frameItem := dataset.New()
	_ = frameItem.Add(dataset.NewSequenceWithItems(tag.FrameVOILUTSequence, []*dataset.Dataset{frameVOI}))
	if err := ds.Add(dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, []*dataset.Dataset{frameItem})); err != nil {
		t.Fatal(err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatal(err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatal(err)
	}
	if got := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y; got != 0 {
		t.Fatalf("pixel = %d, want 0 from per-frame window", got)
	}
}

func TestPerFrameFunctionalGroupRescaleOverridesTopLevelRescale(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{50})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{1}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))
	_ = ds.Add(element.NewString(tag.RescaleType, vr.LO, []string{"US"}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{100}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100}))

	transform := dataset.New()
	_ = transform.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}))
	_ = transform.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))
	_ = transform.Add(element.NewString(tag.RescaleType, vr.LO, []string{"US"}))
	frameItem := dataset.New()
	_ = frameItem.Add(dataset.NewSequenceWithItems(tag.PixelValueTransformationSequence, []*dataset.Dataset{transform}))
	if err := ds.Add(dataset.NewSequenceWithItems(tag.PerFrameFunctionalGroupsSequence, []*dataset.Dataset{frameItem})); err != nil {
		t.Fatal(err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatal(err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatal(err)
	}
	got := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y
	if got < 127 || got > 129 {
		t.Fatalf("pixel = %d, want approximately 128 from per-frame rescale", got)
	}
}

func TestDatasetVOILUTFunctionControlsWindowing(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{125})
	for _, elem := range []element.Element{
		element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{100}),
		element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100}),
		element.NewString(tag.VOILUTFunction, vr.CS, []string{testVOILUTFunctionSigmoid}),
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

func TestModalityLUTSequenceCannotCoexistWithRescale(t *testing.T) {
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
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0})); err != nil {
		t.Fatalf("add RescaleIntercept: %v", err)
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
	if _, err := dicomImage.RenderFrameImage(0); err == nil {
		t.Fatal("RenderFrameImage(0) accepted Modality LUT Sequence with Rescale Slope/Intercept")
	}
}

func TestRenderRejectsIncompleteRescalePair(t *testing.T) {
	for _, rescaleTag := range []*tag.Tag{tag.RescaleSlope, tag.RescaleIntercept} {
		t.Run(rescaleTag.String(), func(t *testing.T) {
			ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
			_ = ds.Add(element.NewDecimalStringFromFloat(rescaleTag, []float64{1}))
			image, err := NewDicomImageFromDataset(ds)
			if err != nil {
				t.Fatalf("NewDicomImageFromDataset() error = %v", err)
			}
			if _, err := image.RenderFrameImage(0); err == nil {
				t.Fatalf("RenderFrameImage(0) accepted lone %s", rescaleTag)
			}
		})
	}
}

func TestRenderSurfacesMalformedVOILUT(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
	item := dataset.New()
	_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 8}))
	_ = ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item}))
	image, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if _, err := image.RenderFrameImage(0); err == nil {
		t.Fatal("RenderFrameImage(0) silently ignored malformed VOI LUT Sequence")
	}
}

func TestRenderRejectsInvalidDatasetWindow(t *testing.T) {
	for _, tt := range []struct {
		name     string
		function string
		width    float64
	}{
		{name: "LINEAR below one", function: "LINEAR", width: 0.5},
		{name: "LINEAR_EXACT zero", function: "LINEAR_EXACT", width: 0},
		{name: "SIGMOID zero", function: testVOILUTFunctionSigmoid, width: 0},
		{name: "unknown function", function: "UNKNOWN", width: 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
			_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{0}))
			_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{tt.width}))
			_ = ds.Add(element.NewString(tag.VOILUTFunction, vr.CS, []string{tt.function}))
			image, err := NewDicomImageFromDataset(ds)
			if err != nil {
				t.Fatalf("NewDicomImageFromDataset() error = %v", err)
			}
			if _, err := image.RenderFrameImage(0); err == nil {
				t.Fatal("RenderFrameImage(0) accepted invalid window")
			}
		})
	}
}

func TestRenderTreatsEmptyVOILUTFunctionAsLinear(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{100})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{100}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100}))
	_ = ds.Add(element.NewString(tag.VOILUTFunction, vr.CS, []string{""}))
	image, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if _, err := image.RenderFrameImage(0); err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
}

func TestRenderRejectsZeroRescaleSlope(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{0}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))
	_ = ds.Add(element.NewString(tag.RescaleType, vr.LO, []string{"US"}))
	image, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	if _, err := image.RenderFrameImage(0); err == nil {
		t.Fatal("RenderFrameImage(0) accepted Rescale Slope=0")
	}
}

func TestVOILUTDescriptorSignednessUsesRescaledDomain(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 1})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{1}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{-100}))
	_ = ds.Add(element.NewString(tag.RescaleType, vr.LO, []string{"HU"}))
	item := dataset.New()
	_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{2, 0xff9c, 8}))
	_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{0, 255}))
	_ = ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item}))

	image, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := image.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{0, 255} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d", x, got, want)
		}
	}
}

func TestWindowAndVOILUTAlternativeSelection(t *testing.T) {
	t.Run("window pair", func(t *testing.T) {
		ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
		_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{0, 100}))
		_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{1, 1}))

		defaultImage, err := NewDicomImageFromDataset(ds)
		if err != nil {
			t.Fatalf("NewDicomImageFromDataset() error = %v", err)
		}
		selectedImage, err := NewDicomImageFromDataset(ds, WithWindowIndex(1))
		if err != nil {
			t.Fatalf("NewDicomImageFromDataset(WithWindowIndex) error = %v", err)
		}
		for _, tt := range []struct {
			name  string
			image *DicomImage
			want  uint8
		}{
			{name: "default first", image: defaultImage, want: 255},
			{name: "selected second", image: selectedImage, want: 0},
		} {
			rendered, err := tt.image.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("%s RenderFrameImage() error = %v", tt.name, err)
			}
			if got := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y; got != tt.want {
				t.Fatalf("%s pixel = %d, want %d", tt.name, got, tt.want)
			}
		}
	})

	t.Run("VOI LUT item", func(t *testing.T) {
		ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
		items := make([]*dataset.Dataset, 2)
		for index, output := range []byte{0, 255} {
			item := dataset.New()
			_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 8}))
			_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{output}))
			items[index] = item
		}
		_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, items))

		for _, tt := range []struct {
			name    string
			options []DicomImageOption
			want    uint8
		}{
			{name: "default first", want: 0},
			{name: "selected second", options: []DicomImageOption{WithVOILUTIndex(1)}, want: 255},
		} {
			image, err := NewDicomImageFromDataset(ds, tt.options...)
			if err != nil {
				t.Fatalf("%s NewDicomImageFromDataset() error = %v", tt.name, err)
			}
			rendered, err := image.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("%s RenderFrameImage() error = %v", tt.name, err)
			}
			if got := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y; got != tt.want {
				t.Fatalf("%s pixel = %d, want %d", tt.name, got, tt.want)
			}
		}
	})
}

func TestAlternativeSelectionValidation(t *testing.T) {
	t.Run("window count mismatch", func(t *testing.T) {
		ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
		_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{0, 1}))
		_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{1}))
		image, _ := NewDicomImageFromDataset(ds)
		if _, err := image.RenderFrameImage(0); err == nil {
			t.Fatal("RenderFrameImage() accepted mismatched Window Center/Width counts")
		}
	})

	t.Run("window index", func(t *testing.T) {
		ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
		_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{0}))
		_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{1}))
		image, _ := NewDicomImageFromDataset(ds, WithWindowIndex(1))
		if _, err := image.RenderFrameImage(0); err == nil {
			t.Fatal("RenderFrameImage() accepted out-of-range window index")
		}
	})

	t.Run("VOI LUT index", func(t *testing.T) {
		ds := newNativeMonochromeDataset(t, 1, 1, []byte{0})
		item := dataset.New()
		_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 8}))
		_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{0}))
		_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item}))
		image, _ := NewDicomImageFromDataset(ds, WithVOILUTIndex(1))
		if _, err := image.RenderFrameImage(0); err == nil {
			t.Fatal("RenderFrameImage() accepted out-of-range VOI LUT index")
		}
	})
}

func TestFunctionalGroupModalityLUTOverridesTopLevelRescale(t *testing.T) {
	primary := dataset.New()
	lutItem := dataset.New()
	if err := lutItem.Add(element.NewString(tag.ModalityLUTType, vr.LO, []string{"US"})); err != nil {
		t.Fatalf("add Modality LUT Type: %v", err)
	}
	if err := lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{2, 0, 8})); err != nil {
		t.Fatalf("add LUT Descriptor: %v", err)
	}
	if err := addLegacyTestElement(lutItem, element.NewOtherByte(tag.LUTData, []byte{10, 20})); err != nil {
		t.Fatalf("add LUT Data: %v", err)
	}
	if err := addLegacyTestElement(primary, dataset.NewSequenceWithItems(tag.ModalityLUTSequence, []*dataset.Dataset{lutItem})); err != nil {
		t.Fatalf("add Modality LUT Sequence: %v", err)
	}

	fallback := dataset.New()
	for _, elem := range []element.Element{
		element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{100}),
		element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}),
		element.NewString(tag.RescaleType, vr.LO, []string{"US"}),
	} {
		if err := fallback.Add(elem); err != nil {
			t.Fatalf("add top-level rescale element: %v", err)
		}
	}

	modality, slope, intercept, _, err := imageModalityTransform(primary, fallback, false, 0, 1)
	if err != nil {
		t.Fatalf("imageModalityTransform() error = %v", err)
	}
	if modality == nil {
		t.Fatal("imageModalityTransform() returned no Functional Group LUT")
	}
	if slope != 1 || intercept != 0 {
		t.Fatalf("rescale = (%v, %v), want identity when Functional Group LUT is selected", slope, intercept)
	}
	if got := modality.Transform(1); got != 20 {
		t.Fatalf("Functional Group LUT Transform(1) = %v, want 20", got)
	}
}

func TestModalityLUTRejectsNonStandardBitDepth(t *testing.T) {
	ds := dataset.New()
	item := dataset.New()
	_ = item.Add(element.NewString(tag.ModalityLUTType, vr.LO, []string{"US"}))
	if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 12})); err != nil {
		t.Fatalf("add LUT Descriptor: %v", err)
	}
	if err := item.Add(element.NewOtherWord(tag.LUTData, []byte{0xff, 0x0f})); err != nil {
		t.Fatalf("add LUT Data: %v", err)
	}
	if err := ds.Add(dataset.NewSequenceWithItems(tag.ModalityLUTSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("add Modality LUT Sequence: %v", err)
	}

	if _, err := pixeldata.ModalityLUT(ds, false); err == nil {
		t.Fatal("pixeldata.ModalityLUT() accepted a bit depth other than 8 or 16")
	}
}

func TestModalityLUTRequiresExactlyOneItem(t *testing.T) {
	for _, count := range []int{0, 2} {
		t.Run(fmt.Sprintf("items_%d", count), func(t *testing.T) {
			ds := dataset.New()
			items := make([]*dataset.Dataset, count)
			for index := range items {
				item := dataset.New()
				_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 8}))
				_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{byte(index)}))
				items[index] = item
			}
			_ = ds.Add(dataset.NewSequenceWithItems(tag.ModalityLUTSequence, items))

			if _, err := pixeldata.ModalityLUT(ds, false); err == nil {
				t.Fatalf("pixeldata.ModalityLUT() accepted %d items", count)
			}
		})
	}
}

func TestDatasetImageReadsBigEndianModalityLUTData(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	ds.SetInternalTransferSyntax(transfer.ExplicitVRBigEndian)
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewString(tag.ModalityLUTType, vr.LO, []string{"US"}))
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 16}))
	lutBytes := make([]byte, 6)
	for index, value := range []uint16{0, 100, 200} {
		binary.BigEndian.PutUint16(lutBytes[index*2:], value)
	}
	lutData := element.NewOtherWord(tag.LUTData, lutBytes)
	element.SetByteOrder(lutData, binary.BigEndian)
	_ = lutItem.Add(lutData)
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

func TestSignedShortModalityLUTDescriptorMapsSignedInputs(t *testing.T) {
	ds := dataset.New()
	item := dataset.New()
	_ = item.Add(element.NewString(tag.ModalityLUTType, vr.LO, []string{"US"}))
	_ = item.Add(element.NewSignedShort(tag.LUTDescriptor, []int16{3, -1, 8}))
	_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{10, 20, 30}))
	_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.ModalityLUTSequence, []*dataset.Dataset{item}))

	table, err := pixeldata.ModalityLUT(ds, true)
	if err != nil {
		t.Fatalf("pixeldata.ModalityLUT() error = %v", err)
	}
	for input, want := range map[int]float64{-1: 10, 0: 20, 1: 30} {
		if got := table.Transform(float64(input)); got != want {
			t.Fatalf("Transform(%d) = %v, want %v", input, got, want)
		}
	}
}

func TestSignedShortVOILUTDescriptorMapsSignedInputs(t *testing.T) {
	ds := dataset.New()
	item := dataset.New()
	_ = item.Add(element.NewSignedShort(tag.LUTDescriptor, []int16{3, -1, 8}))
	_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{10, 20, 30}))
	_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item}))

	table, err := pixeldata.VOILUT(ds, true)
	if err != nil {
		t.Fatalf("pixeldata.VOILUT() error = %v", err)
	}
	for input, want := range map[int]float64{-1: 10, 0: 20, 1: 30} {
		if got := table.Transform(float64(input)); math.Abs(got-want) > 1e-9 {
			t.Fatalf("Transform(%d) = %v, want %v", input, got, want)
		}
	}
}

func TestUnsignedShortLUTDescriptorUsesSignedPixelRepresentation(t *testing.T) {
	for _, tt := range []struct {
		name string
		tag  *tag.Tag
		load func(*dataset.Dataset) (lut.LUT, error)
	}{
		{
			name: "modality",
			tag:  tag.ModalityLUTSequence,
			load: func(ds *dataset.Dataset) (lut.LUT, error) {
				return pixeldata.ModalityLUT(ds, true)
			},
		},
		{
			name: "VOI",
			tag:  tag.VOILUTSequence,
			load: func(ds *dataset.Dataset) (lut.LUT, error) {
				return pixeldata.VOILUT(ds, true)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			item := dataset.New()
			if tt.tag == tag.ModalityLUTSequence {
				_ = item.Add(element.NewString(tag.ModalityLUTType, vr.LO, []string{"US"}))
			}
			_ = item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0xffff, 8}))
			_ = addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{10, 20, 30}))
			_ = addLegacyTestElement(ds, dataset.NewSequenceWithItems(tt.tag, []*dataset.Dataset{item}))

			table, err := tt.load(ds)
			if err != nil {
				t.Fatalf("load LUT error = %v", err)
			}
			for input, want := range map[int]float64{-1: 10, 0: 20, 1: 30} {
				if got := table.Transform(float64(input)); math.Abs(got-want) > 1e-9 {
					t.Fatalf("Transform(%d) = %v, want %v", input, got, want)
				}
			}
		})
	}
}

func TestVOILUTSequencePrecedesWindow(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 1, 2})
	lutItem := dataset.New()
	if err := lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8})); err != nil {
		t.Fatalf("add LUTDescriptor: %v", err)
	}
	if err := addLegacyTestElement(lutItem, element.NewOtherByte(tag.LUTData, []byte{0, 100, 200})); err != nil {
		t.Fatalf("add LUTData: %v", err)
	}
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem})); err != nil {
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
	want := []uint8{0, 100, 200}
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

func TestPresentationLUTShapeInverseIsAppliedAfterVOI(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{127.5})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{256})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}
	if err := ds.Add(element.NewString(tag.PresentationLUTShape, vr.CS, []string{"INVERSE"})); err != nil {
		t.Fatalf("add PresentationLUTShape: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{255, 0} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d", x, got, want)
		}
	}
}

func TestPresentationLUTSequenceIsAppliedAfterVOI(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{127.5})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{256})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}
	data := make([]byte, 256)
	for index := range data {
		data[index] = byte(255 - index)
	}
	item := dataset.New()
	if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{256, 0, 8})); err != nil {
		t.Fatalf("add LUTDescriptor: %v", err)
	}
	if err := addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, data)); err != nil {
		t.Fatalf("add LUTData: %v", err)
	}
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.PresentationLUTSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("add PresentationLUTSequence: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{255, 0} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d", x, got, want)
		}
	}
}

func TestPresentationLUTRejectsSequenceAndShapeTogether(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 8})); err != nil {
		t.Fatal(err)
	}
	if err := addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{0})); err != nil {
		t.Fatal(err)
	}
	ds := dataset.New()
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.PresentationLUTSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatal(err)
	}
	if err := ds.Add(element.NewString(tag.PresentationLUTShape, vr.CS, []string{"IDENTITY"})); err != nil {
		t.Fatal(err)
	}

	_, _, err := pixeldata.PresentationLUT(ds)
	if err == nil || !strings.Contains(err.Error(), "mutually exclusive") {
		t.Fatalf("PresentationLUT() error = %v, want mutual-exclusion error", err)
	}
}

func TestPresentationLUTAcceptsTenBitDescriptor(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 10})); err != nil {
		t.Fatal(err)
	}
	if err := addLegacyTestElement(item, element.NewOtherWord(tag.LUTData, []byte{0xff, 0x03})); err != nil {
		t.Fatal(err)
	}
	ds := dataset.New()
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.PresentationLUTSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatal(err)
	}
	if _, present, err := pixeldata.PresentationLUTWithVRMode(ds, pixeldata.LUTStandard); err != nil || !present {
		t.Fatalf("PresentationLUTWithVRMode() = present %v, error %v; want a valid 10-bit LUT", present, err)
	}
}

func TestPresentationLUTStandardRejectsEightBitDescriptor(t *testing.T) {
	item := dataset.New()
	if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, 8})); err != nil {
		t.Fatal(err)
	}
	if err := addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{0})); err != nil {
		t.Fatal(err)
	}
	ds := dataset.New()
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.PresentationLUTSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatal(err)
	}
	if _, _, err := pixeldata.PresentationLUTWithVRMode(ds, pixeldata.LUTStandard); err == nil {
		t.Fatal("PresentationLUTWithVRMode() accepted an 8-bit descriptor in standard mode")
	}
}

func TestSoftcopyVOILUTSequenceSelectsReferencedFrame(t *testing.T) {
	const sopInstanceUID = "1.2.826.0.1.3680043.10.543.1"
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{50, 50})
	if err := ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"})); err != nil {
		t.Fatalf("add NumberOfFrames: %v", err)
	}
	if err := ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstanceUID})); err != nil {
		t.Fatalf("add SOPInstanceUID: %v", err)
	}

	defaultVOI := dataset.New()
	_ = defaultVOI.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{50}))
	_ = defaultVOI.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100}))
	specificVOI := dataset.New()
	_ = specificVOI.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{100}))
	_ = specificVOI.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100}))
	reference := dataset.New()
	_ = reference.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{sopInstanceUID}))
	_ = reference.Add(element.NewString(tag.ReferencedFrameNumber, vr.IS, []string{"2"}))
	_ = specificVOI.Add(dataset.NewSequenceWithItems(tag.ReferencedImageSequence, []*dataset.Dataset{reference}))
	presentationState := dataset.New()
	if err := presentationState.Add(dataset.NewSequenceWithItems(
		tag.SoftcopyVOILUTSequence,
		[]*dataset.Dataset{defaultVOI, specificVOI},
	)); err != nil {
		t.Fatalf("add SoftcopyVOILUTSequence: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds, WithPresentationState(presentationState))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	frame0, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got := color.GrayModel.Convert(frame0.At(0, 0)).(color.Gray).Y; got < 127 || got > 129 {
		t.Fatalf("frame 0 = %d, want approximately 128 from default Softcopy VOI", got)
	}
	frame1, err := dicomImage.RenderFrameImage(1)
	if err != nil {
		t.Fatalf("RenderFrameImage(1) error = %v", err)
	}
	if got := color.GrayModel.Convert(frame1.At(0, 0)).(color.Gray).Y; got != 0 {
		t.Fatalf("frame 1 = %d, want 0 from referenced-frame Softcopy VOI", got)
	}
}

func TestSoftcopyVOILUTSequenceRejectsAmbiguousSelection(t *testing.T) {
	const sopInstanceUID = "1.2.826.0.1.3680043.10.543.2"
	imageDataset := dataset.New()
	if err := imageDataset.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstanceUID})); err != nil {
		t.Fatal(err)
	}

	referencedItem := func() *dataset.Dataset {
		reference := dataset.New()
		_ = reference.Add(element.NewString(tag.ReferencedSOPInstanceUID, vr.UI, []string{sopInstanceUID}))
		_ = reference.Add(element.NewString(tag.ReferencedFrameNumber, vr.IS, []string{"1"}))
		item := dataset.New()
		_ = item.Add(dataset.NewSequenceWithItems(tag.ReferencedImageSequence, []*dataset.Dataset{reference}))
		return item
	}

	for _, test := range []struct {
		name  string
		items []*dataset.Dataset
		want  string
	}{
		{
			name:  "duplicate exact frame matches",
			items: []*dataset.Dataset{referencedItem(), referencedItem()},
			want:  "multiple Softcopy VOI LUT Sequence items reference frame 1",
		},
		{
			name:  "multiple defaults",
			items: []*dataset.Dataset{dataset.New(), dataset.New()},
			want:  "multiple default items",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			presentationState := dataset.New()
			if err := presentationState.Add(dataset.NewSequenceWithItems(tag.SoftcopyVOILUTSequence, test.items)); err != nil {
				t.Fatal(err)
			}
			_, err := imageSoftcopyVOI(presentationState, imageDataset, 0)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("imageSoftcopyVOI() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestImageVOILUTReportsValidForStandardSequence(t *testing.T) {
	ds := dataset.New()
	item := dataset.New()
	if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 8})); err != nil {
		t.Fatalf("add LUT Descriptor: %v", err)
	}
	if err := addLegacyTestElement(item, element.NewOtherByte(tag.LUTData, []byte{0, 100, 200})); err != nil {
		t.Fatalf("add LUT Data: %v", err)
	}
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("add VOI LUT Sequence: %v", err)
	}

	table, err := pixeldata.VOILUT(ds, false)
	if err != nil {
		t.Fatalf("pixeldata.VOILUT() error = %v", err)
	}
	if !table.IsValid() {
		t.Fatal("pixeldata.VOILUT() returned a LUT that reports itself invalid")
	}
}

func TestImageVOILUTRejectsPresentationStateBitDepths(t *testing.T) {
	for _, bitsPerEntry := range []uint16{9, 12, 15} {
		t.Run(fmt.Sprintf("bits_%d", bitsPerEntry), func(t *testing.T) {
			ds := dataset.New()
			item := dataset.New()
			if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{1, 0, bitsPerEntry})); err != nil {
				t.Fatalf("add LUT Descriptor: %v", err)
			}
			if err := item.Add(element.NewOtherWord(tag.LUTData, []byte{0, 0})); err != nil {
				t.Fatalf("add LUT Data: %v", err)
			}
			if err := ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{item})); err != nil {
				t.Fatalf("add VOI LUT Sequence: %v", err)
			}

			if _, err := pixeldata.VOILUT(ds, false); err == nil {
				t.Fatalf("pixeldata.VOILUT() accepted image VOI LUT bitsPerEntry=%d", bitsPerEntry)
			}
		})
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
	lutData := element.NewOtherWord(tag.LUTData, lutBytes)
	element.SetByteOrder(lutData, binary.BigEndian)
	_ = lutItem.Add(lutData)
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

func TestDatasetImageReadsSourceEndianVOILUTAfterNativeTranscode(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{3}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	voiItem := dataset.New()
	_ = voiItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{3, 0, 16}))
	voiData := element.NewOtherWord(tag.LUTData, []byte{0, 0, 0x80, 0, 0xff, 0xff})
	element.SetByteOrder(voiData, binary.BigEndian)
	_ = voiItem.Add(voiData)
	_ = ds.Add(dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{voiItem}))

	pixelData := element.NewOtherWord(tag.PixelData, []byte{1, 0, 0, 2})
	element.SetByteOrder(pixelData, binary.BigEndian)
	_ = ds.Add(pixelData)

	manager, err := transcode.NewManager(codec.NewRegistry())
	if err != nil {
		t.Fatalf("transcode.NewManager() error = %v", err)
	}
	transcoder, err := manager.NewTranscoder(transfer.ExplicitVRBigEndian, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("Manager.NewTranscoder() error = %v", err)
	}
	transcoded, err := transcoder.Transcode(context.Background(), ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}
	transcodedPixelData, ok := transcoded.Get(tag.PixelData)
	if !ok {
		t.Fatal("transcoded dataset has no Pixel Data")
	}
	if got, want := transcodedPixelData.(*element.OtherWord).GetData(), []byte{0, 1, 2, 0}; !bytes.Equal(got, want) {
		t.Fatalf("transcoded Pixel Data = %v, want %v", got, want)
	}
	if got, known := element.NumericByteOrder(voiData); !known || got.Uint16([]byte{0, 1}) != binary.BigEndian.Uint16([]byte{0, 1}) {
		t.Fatalf("VOI LUT byte order changed to %T, want Big Endian source order", got)
	}

	dicomImage, err := NewDicomImageFromDataset(transcoded)
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

func TestDefaultWindowUsesModalityValueRange(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 3, 1, []byte{0, 50, 100})
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2})); err != nil {
		t.Fatalf("add RescaleSlope: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0})); err != nil {
		t.Fatalf("add RescaleIntercept: %v", err)
	}
	if err := ds.Add(element.NewString(tag.RescaleType, vr.LO, []string{"US"})); err != nil {
		t.Fatalf("add RescaleType: %v", err)
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
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0})); err != nil {
		t.Fatalf("add RescaleIntercept: %v", err)
	}
	if err := ds.Add(element.NewString(tag.RescaleType, vr.LO, []string{"US"})); err != nil {
		t.Fatalf("add RescaleType: %v", err)
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
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
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
				element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
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
		// OW byte order is applied independently to each 16-bit word. The
		// low word remains before the high word for a 32-bit native sample.
		binary.BigEndian.PutUint16(pixels[index*4:], uint16(value))
		binary.BigEndian.PutUint16(pixels[index*4+2:], uint16(value>>16))
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
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
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
	registry := codec.NewRegistry()
	if err := registry.Register(imagePassthroughCodec{}); err != nil {
		t.Fatal(err)
	}

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
	registry := codec.NewRegistry()
	if err := registry.Register(mutatingImageCodec{}); err != nil {
		t.Fatal(err)
	}

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
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.PaletteColor.Value}),
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{0, 255}),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{0, 0}),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{0, 0}),
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
	registry := codec.NewRegistry()
	if err := registry.Register(paletteIndexCodec{}); err != nil {
		t.Fatal(err)
	}

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

func TestDatasetPaletteAlphaIsPreservedInRendering(t *testing.T) {
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{2}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.PaletteColor.Value}),
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewUnsignedShort(tag.AlphaPaletteColorLookupTableDescriptor, []uint16{2, 0, 8}),
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{10, 20}),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{30, 40}),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{50, 60}),
		element.NewOtherWord(tag.AlphaPaletteColorLookupTableData, []byte{70, 80}),
		element.NewOtherByte(tag.PixelData, []byte{0, 1}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage() error = %v", err)
	}
	for x, want := range []color.NRGBA{
		{R: 10, G: 30, B: 50, A: 70},
		{R: 20, G: 40, B: 60, A: 80},
	} {
		if got := color.NRGBAModel.Convert(rendered.At(x, 0)).(color.NRGBA); got != want {
			t.Fatalf("pixel %d = %#v, want %#v", x, got, want)
		}
	}
}

func TestSupplementalPaletteRendering(t *testing.T) {
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{3}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewString(tag.PixelPresentation, vr.CS, []string{"COLOR"}),
		element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{127.5}),
		element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{256}),
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{2, 10, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{2, 10, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{2, 10, 8}),
		element.NewUnsignedShort(tag.AlphaPaletteColorLookupTableDescriptor, []uint16{2, 10, 8}),
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{100, 200}),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{0, 0}),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{0, 0}),
		element.NewOtherWord(tag.AlphaPaletteColorLookupTableData, []byte{128, 255}),
		element.NewOtherByte(tag.PixelData, []byte{5, 10, 11, 5, 10, 11}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage() error = %v", err)
	}
	below := color.NRGBAModel.Convert(rendered.At(0, 0)).(color.NRGBA)
	if below.R != below.G || below.G != below.B || below.A != 255 {
		t.Fatalf("value below first mapped rendered as %#v, want opaque grayscale", below)
	}
	for x, want := range map[int]color.NRGBA{
		1: {R: 100, A: 128},
		2: {R: 200, A: 255},
	} {
		if got := color.NRGBAModel.Convert(rendered.At(x, 0)).(color.NRGBA); got != want {
			t.Fatalf("pixel %d = %#v, want %#v", x, got, want)
		}
	}
}

func TestSupplementalPaletteRenderingSupportsSingleFrame(t *testing.T) {
	ds, err := dataset.NewWithElements([]element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.NumberOfFrames, vr.IS, []string{"1"}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewString(tag.PixelPresentation, vr.CS, []string{"COLOR"}),
		element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{1, 10, 8}),
		element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{1, 10, 8}),
		element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{1, 10, 8}),
		element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{100, 0}),
		element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{20, 0}),
		element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{30, 0}),
		element.NewOtherByte(tag.PixelData, []byte{10, 0}),
	})
	if err != nil {
		t.Fatalf("NewWithElements() error = %v", err)
	}
	dicomImage, err := NewDicomImageFromDataset(ds)
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage() error = %v", err)
	}
	if got, want := color.NRGBAModel.Convert(rendered.At(0, 0)).(color.NRGBA), (color.NRGBA{R: 100, G: 20, B: 30, A: 255}); got != want {
		t.Fatalf("pixel = %#v, want %#v", got, want)
	}
}

func TestSupplementalPaletteRequiresEligibility(t *testing.T) {
	for _, tt := range []struct {
		name              string
		numberOfFrames    string
		pixelPresentation string
		pixelData         []byte
	}{
		{name: "missing Pixel Presentation", numberOfFrames: "2", pixelData: []byte{10, 10}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			elements := []element.Element{
				element.NewUnsignedShort(tag.Rows, []uint16{1}),
				element.NewUnsignedShort(tag.Columns, []uint16{1}),
				element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
				element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
				element.NewUnsignedShort(tag.HighBit, []uint16{7}),
				element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
				element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
				element.NewString(tag.NumberOfFrames, vr.IS, []string{tt.numberOfFrames}),
				element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
				element.NewUnsignedShort(tag.RedPaletteColorLookupTableDescriptor, []uint16{1, 10, 8}),
				element.NewUnsignedShort(tag.GreenPaletteColorLookupTableDescriptor, []uint16{1, 10, 8}),
				element.NewUnsignedShort(tag.BluePaletteColorLookupTableDescriptor, []uint16{1, 10, 8}),
				element.NewOtherWord(tag.RedPaletteColorLookupTableData, []byte{100, 0}),
				element.NewOtherWord(tag.GreenPaletteColorLookupTableData, []byte{0, 0}),
				element.NewOtherWord(tag.BluePaletteColorLookupTableData, []byte{0, 0}),
				element.NewOtherByte(tag.PixelData, tt.pixelData),
			}
			if tt.pixelPresentation != "" {
				elements = append(elements, element.NewString(tag.PixelPresentation, vr.CS, []string{tt.pixelPresentation}))
			}
			ds, err := dataset.NewWithElements(elements)
			if err != nil {
				t.Fatalf("NewWithElements() error = %v", err)
			}
			dicomImage, err := NewDicomImageFromDataset(ds)
			if err != nil {
				t.Fatalf("NewDicomImageFromDataset() error = %v", err)
			}
			rendered, err := dicomImage.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("RenderFrameImage() error = %v", err)
			}
			if _, ok := rendered.(*image.Gray); !ok {
				t.Fatalf("rendered image = %T, want grayscale when Supplemental Palette is ineligible", rendered)
			}
		})
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
	dicomImage.SetOverlayColor(colorconv.Color32{A: 255, R: 255})
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
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
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
	dicomImage.SetOverlayColor(colorconv.Color32{A: 255, R: 255})
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
	info := &pixeldata.Info{
		Width: 2, Height: 1, NumberOfFrames: 2,
		BitsAllocated: 8, BitsStored: 8, HighBit: 7, SamplesPerPixel: 1,
		PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.Monochrome2,
	}
	pixelData, err := pixeldata.NewFromBytes(info, []byte{0, 100, 200, 255})
	if err != nil {
		t.Fatalf("pixeldata.NewFromBytes() error = %v", err)
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
	var colorMap [256]colorconv.Color32
	for index := range colorMap {
		colorMap[index] = colorconv.Color32{A: 255, G: uint8(index)}
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
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0})); err != nil {
		t.Fatalf("add RescaleIntercept: %v", err)
	}
	if err := ds.Add(element.NewString(tag.RescaleType, vr.LO, []string{"US"})); err != nil {
		t.Fatalf("add RescaleType: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{50})); err != nil {
		t.Fatalf("add WindowCenter: %v", err)
	}
	if err := ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{100})); err != nil {
		t.Fatalf("add WindowWidth: %v", err)
	}
	if err := ds.Add(element.NewString(tag.VOILUTFunction, vr.CS, []string{testVOILUTFunctionSigmoid})); err != nil {
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
	_ = addLegacyTestElement(lutItem, element.NewOtherByte(tag.LUTData, []byte{0, 100, 200}))
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
func (imagePassthroughCodec) DefaultParameters() codec.Parameters {
	return codec.NoParameters{}
}
func (imagePassthroughCodec) Encode(context.Context, codec.FrameSource, codec.FrameSink, codec.Parameters) error {
	return nil
}
func (imagePassthroughCodec) Decode(ctx context.Context, oldPixelData codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	for frame := 0; frame < oldPixelData.FrameCount(); frame++ {
		data, err := oldPixelData.Frame(ctx, frame)
		if err != nil {
			return err
		}
		if err := newPixelData.AddFrame(ctx, data); err != nil {
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
func (paletteIndexCodec) DefaultParameters() codec.Parameters {
	return codec.NoParameters{}
}
func (paletteIndexCodec) Encode(context.Context, codec.FrameSource, codec.FrameSink, codec.Parameters) error {
	return nil
}
func (paletteIndexCodec) Decode(ctx context.Context, oldPixelData codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	info := oldPixelData.FrameInfo()
	if info.PhotometricInterpretation.Value != pixel.PaletteColor.Value || info.SamplesPerPixel != 1 {
		return fmt.Errorf("codec input metadata = %s/%d, want PALETTE COLOR/1", info.PhotometricInterpretation.Value, info.SamplesPerPixel)
	}
	frame, err := oldPixelData.Frame(ctx, 0)
	if err != nil {
		return err
	}
	if !bytes.Equal(frame, []byte{0, 1}) {
		return fmt.Errorf("codec input frame = %v, want [0 1]", frame)
	}
	return newPixelData.AddFrame(ctx, frame)
}

type mutatingImageCodec struct{}

func (mutatingImageCodec) Name() string { return "mutating image test codec" }
func (mutatingImageCodec) TransferSyntax() *transfer.Syntax {
	return transfer.JPEG2000Lossless
}
func (mutatingImageCodec) DefaultParameters() codec.Parameters {
	return codec.NoParameters{}
}
func (mutatingImageCodec) Encode(context.Context, codec.FrameSource, codec.FrameSink, codec.Parameters) error {
	return nil
}
func (mutatingImageCodec) Decode(ctx context.Context, oldPixelData codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	for frameIndex := 0; frameIndex < oldPixelData.FrameCount(); frameIndex++ {
		frame, err := oldPixelData.Frame(ctx, frameIndex)
		if err != nil {
			return err
		}
		frame[0] ^= 0xff
		if err := newPixelData.AddFrame(ctx, frame); err != nil {
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

func (p *imageLockCheckingPipeline) LUT() lut.LUT { return p.delegate.LUT() }
func (p *imageLockCheckingPipeline) ClearCache()  { p.delegate.ClearCache() }
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
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, pixels),
	}
	for _, elem := range elements {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}
	return ds
}
