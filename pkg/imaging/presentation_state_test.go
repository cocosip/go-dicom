// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package imaging

import (
	"encoding/binary"
	"image/color"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

// addPresentationLUTSequence appends an identity Presentation LUT Sequence with
// the given number of entries and bits per entry to ds.
func addPresentationLUTSequence(t *testing.T, ds *dataset.Dataset, entries int, bits uint16) {
	t.Helper()
	count := uint16(entries)
	if entries == 65536 {
		count = 0
	}
	item := dataset.New()
	if err := item.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{count, 0, bits})); err != nil {
		t.Fatalf("add Presentation LUT Descriptor: %v", err)
	}
	var data []byte
	var dataElement element.Element
	if bits == 8 {
		data = make([]byte, entries)
		for index := range data {
			data[index] = byte(index)
		}
		dataElement = element.NewOtherByte(tag.LUTData, data)
	} else {
		data = make([]byte, entries*2)
		for index := 0; index < entries; index++ {
			// Identity ramp: entry i maps to i.
			binary.LittleEndian.PutUint16(data[index*2:], uint16(index))
		}
		dataElement = element.NewOtherWord(tag.LUTData, data)
	}
	if err := addLegacyTestElement(item, dataElement); err != nil {
		t.Fatalf("add Presentation LUT Data: %v", err)
	}
	if err := addLegacyTestElement(ds, dataset.NewSequenceWithItems(tag.PresentationLUTSequence, []*dataset.Dataset{item})); err != nil {
		t.Fatalf("add Presentation LUT Sequence: %v", err)
	}
}

// TestPresentationLUTSequenceIdentityRamps verifies PS3.3 C.11.6.1: the
// preceding 0..255 intermediate is linearly scaled onto the Presentation LUT
// input range, so an identity ramp of any legal size maps 0->0 and 255->255.
func TestPresentationLUTSequenceIdentityRamps(t *testing.T) {
	for _, test := range []struct {
		name    string
		entries int
		bits    uint16
	}{
		{name: "8-bit", entries: 256, bits: 8},
		{name: "10-bit", entries: 1024, bits: 10},
		{name: "12-bit", entries: 4096, bits: 12},
		{name: "16-bit", entries: 65536, bits: 16},
	} {
		t.Run(test.name, func(t *testing.T) {
			ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
			addPresentationLUTSequence(t, ds, test.entries, test.bits)

			dicomImage, err := NewDicomImageFromDataset(ds)
			if err != nil {
				t.Fatalf("NewDicomImageFromDataset() error = %v", err)
			}
			rendered, err := dicomImage.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("RenderFrameImage(0) error = %v", err)
			}
			for x, want := range []uint8{0, 255} {
				if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
					t.Fatalf("pixel %d = %d, want %d for the identity ramp", x, got, want)
				}
			}
		})
	}
}

// presentationStateWithSoftcopyWindow builds a Presentation State carrying a
// Softcopy VOI LUT item with the given window.
func presentationStateWithSoftcopyWindow(center, width float64) *dataset.Dataset {
	ps := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	item := dataset.New()
	_ = item.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{center}))
	_ = item.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{width}))
	_ = ps.Add(dataset.NewSequenceWithItems(tag.SoftcopyVOILUTSequence, []*dataset.Dataset{item}))
	return ps
}

// TestPresentationStateWindowOnlyReplacesImageTransform verifies PS3.4 N.2:
// the PS window applies to the identity PS modality transform even when the
// image declares its own Rescale Slope.
func TestPresentationStateWindowOnlyReplacesImageTransform(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{100})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))
	// Image window would map the (ignored) image rescale of 200 to 255.
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{200}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{1}))

	dicomImage, err := NewDicomImageFromDataset(ds, WithPresentationState(
		presentationStateWithSoftcopyWindow(100, 4),
	))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	// Identity modality (stored 100) through center 100 / width 4 => 170.
	if got, want := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y, uint8(170); got != want {
		t.Fatalf("pixel = %d, want %d from the PS window applied to the identity PS modality", got, want)
	}
}

// TestPresentationStateRescaleOnlyUsesIdentityVOI verifies that a PS which
// supplies only a rescale does not fall back to the image window or to the
// image data-derived window (PS3.4 N.2).
func TestPresentationStateRescaleOnlyUsesIdentityVOI(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{10, 20})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))
	// The image window maps 10..20 to 0..255; N.2 requires identity instead.
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{15}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{10}))

	ps := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	_ = ps.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{1}))
	_ = ps.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))

	dicomImage, err := NewDicomImageFromDataset(ds, WithPresentationState(ps))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{10, 20} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want identity %d", x, got, want)
		}
	}
}

// TestPresentationStateRescaleFeedsPresentationWindow verifies that the PS
// rescale is actually applied before the PS window.
func TestPresentationStateRescaleFeedsPresentationWindow(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 1, 1, []byte{100})
	// The image rescale would leave the stored value at 100, below the PS
	// window threshold; the PS rescale of 2 moves it above.
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{1}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))

	ps := presentationStateWithSoftcopyWindow(200, 1)
	_ = ps.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}))
	_ = ps.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))
	_ = ps.Add(element.NewString(tag.RescaleType, vr.LO, []string{"HU"}))

	dicomImage, err := NewDicomImageFromDataset(ds, WithPresentationState(ps))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got, want := color.GrayModel.Convert(rendered.At(0, 0)).(color.Gray).Y, uint8(255); got != want {
		t.Fatalf("pixel = %d, want %d from the PS rescale feeding the PS window", got, want)
	}
}

// TestPresentationStateModalityLUTReplacesImageTransform verifies that a
// non-linear PS Modality LUT is used instead of the image Rescale Slope.
func TestPresentationStateModalityLUTReplacesImageTransform(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleSlope, []float64{2}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.RescaleIntercept, []float64{0}))

	values := make([]byte, 256)
	for index := range values {
		values[index] = byte(255 - index)
	}
	lutItem := dataset.New()
	_ = lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{256, 0, 8}))
	_ = lutItem.Add(element.NewString(tag.ModalityLUTType, vr.LO, []string{"US"}))
	if err := addLegacyTestElement(lutItem, element.NewOtherByte(tag.LUTData, values)); err != nil {
		t.Fatalf("add Modality LUT Data: %v", err)
	}
	ps := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	if err := addLegacyTestElement(ps, dataset.NewSequenceWithItems(tag.ModalityLUTSequence, []*dataset.Dataset{lutItem})); err != nil {
		t.Fatalf("add Modality LUT Sequence: %v", err)
	}

	dicomImage, err := NewDicomImageFromDataset(ds, WithPresentationState(ps))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{255, 0} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want %d from the inverted PS Modality LUT", x, got, want)
		}
	}
}

// TestPresentationStateEmptyUsesIdentityTransforms verifies that an empty
// Presentation State replaces the image transforms with identity (PS3.4 N.2)
// and does not reuse the image window, data-derived window, or Presentation
// LUT.
func TestPresentationStateEmptyUsesIdentityTransforms(t *testing.T) {
	ds := newNativeMonochromeDataset(t, 2, 1, []byte{10, 20})
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowCenter, []float64{15}))
	_ = ds.Add(element.NewDecimalStringFromFloat(tag.WindowWidth, []float64{10}))
	// An image Presentation LUT must not be used when a PS is applied.
	_ = ds.Add(element.NewString(tag.PresentationLUTShape, vr.CS, []string{lutShapeInverse}))

	ps := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)

	dicomImage, err := NewDicomImageFromDataset(ds, WithPresentationState(ps))
	if err != nil {
		t.Fatalf("NewDicomImageFromDataset() error = %v", err)
	}
	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	for x, want := range []uint8{10, 20} {
		if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
			t.Fatalf("pixel %d = %d, want identity %d", x, got, want)
		}
	}
}

// Presentation LUT Shape values used by the presentation state tests.
const (
	lutShapeIdentity = "IDENTITY"
	lutShapeInverse  = "INVERSE"
)

// TestPresentationStateIgnoresMonochrome1Polarity verifies PS3.4 N.2: with a
// Presentation State, the referenced image's MONOCHROME1 polarity is ignored
// and the presentation LUT Shape owns polarity.
func TestPresentationStateIgnoresMonochrome1Polarity(t *testing.T) {
	for _, test := range []struct {
		name  string
		shape string
		want  []uint8
	}{
		{name: "IDENTITY", shape: lutShapeIdentity, want: []uint8{0, 255}},
		{name: "INVERSE", shape: lutShapeInverse, want: []uint8{255, 0}},
	} {
		t.Run(test.name, func(t *testing.T) {
			ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
			_ = ds.AddOrUpdate(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome1.Value}))
			ps := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
			_ = ps.Add(element.NewString(tag.PresentationLUTShape, vr.CS, []string{test.shape}))

			dicomImage, err := NewDicomImageFromDataset(ds, WithPresentationState(ps))
			if err != nil {
				t.Fatalf("NewDicomImageFromDataset() error = %v", err)
			}
			rendered, err := dicomImage.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("RenderFrameImage(0) error = %v", err)
			}
			for x, want := range test.want {
				if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
					t.Fatalf("pixel %d = %d, want %d (image MONOCHROME1 ignored)", x, got, want)
				}
			}
		})
	}
}

// TestSoftcopyVOILUTAcceptsWideBitsPerEntry verifies PS3.3 C.11.2.1.1: a
// Presentation State Softcopy VOI LUT may use 8 through 16 bits per entry.
func TestSoftcopyVOILUTAcceptsWideBitsPerEntry(t *testing.T) {
	for _, test := range []struct {
		name string
		bits uint16
		want []uint8
	}{
		{name: "9-bit", bits: 9, want: []uint8{0, 127}},
		{name: "12-bit", bits: 12, want: []uint8{0, 16}},
		{name: "15-bit", bits: 15, want: []uint8{0, 2}},
	} {
		t.Run(test.name, func(t *testing.T) {
			const entries = 512
			lutItem := dataset.New()
			if err := lutItem.Add(element.NewUnsignedShort(tag.LUTDescriptor, []uint16{entries, 0, test.bits})); err != nil {
				t.Fatalf("add VOI LUT Descriptor: %v", err)
			}
			data := make([]byte, entries*2)
			for index := 0; index < entries; index++ {
				binary.LittleEndian.PutUint16(data[index*2:], uint16(index))
			}
			if err := addLegacyTestElement(lutItem, element.NewOtherWord(tag.LUTData, data)); err != nil {
				t.Fatalf("add VOI LUT Data: %v", err)
			}
			softcopyItem := dataset.New()
			if err := addLegacyTestElement(softcopyItem, dataset.NewSequenceWithItems(tag.VOILUTSequence, []*dataset.Dataset{lutItem})); err != nil {
				t.Fatalf("add nested VOI LUT Sequence: %v", err)
			}
			ps := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
			if err := addLegacyTestElement(ps, dataset.NewSequenceWithItems(tag.SoftcopyVOILUTSequence, []*dataset.Dataset{softcopyItem})); err != nil {
				t.Fatalf("add Softcopy VOI LUT Sequence: %v", err)
			}

			ds := newNativeMonochromeDataset(t, 2, 1, []byte{0, 255})
			dicomImage, err := NewDicomImageFromDataset(ds, WithPresentationState(ps))
			if err != nil {
				t.Fatalf("NewDicomImageFromDataset() error = %v", err)
			}
			rendered, err := dicomImage.RenderFrameImage(0)
			if err != nil {
				t.Fatalf("RenderFrameImage(0) error = %v, want a legal Presentation State VOI LUT", err)
			}
			for x, want := range test.want {
				if got := color.GrayModel.Convert(rendered.At(x, 0)).(color.Gray).Y; got != want {
					t.Fatalf("pixel %d = %d, want %d", x, got, want)
				}
			}
		})
	}
}



