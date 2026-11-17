// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package imaging

import (
	"encoding/binary"
	"image"
	"image/color"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/colorconv"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
)

// TestOverlayDataOtherWordBytesAreSwappedForBigEndian verifies IMG-14: Overlay
// Data is OW, so a big-endian element's words must be swapped into the little-
// endian bit stream the renderer unpacks.
func TestOverlayDataOtherWordBytesAreSwappedForBigEndian(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.New(0x6000, 0x0010), []uint16{16}),
		element.NewUnsignedShort(tag.New(0x6000, 0x0011), []uint16{1}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}
	data := element.NewOtherWord(tag.New(0x6000, 0x3000), []byte{0x00, 0xFF})
	element.SetByteOrder(data, binary.BigEndian)
	if err := ds.Add(data); err != nil {
		t.Fatalf("add Overlay Data: %v", err)
	}

	overlays := imageOverlays(ds, nil)
	if len(overlays) != 1 {
		t.Fatalf("imageOverlays() returned %d overlays, want 1", len(overlays))
	}
	if got := overlays[0].Data; len(got) != 2 || got[0] != 0xFF || got[1] != 0x00 {
		t.Fatalf("overlay data = % x, want ff 00 after OW byte swap", got)
	}
}

func TestOverlayType_String(t *testing.T) {
	testCases := []struct {
		overlayType OverlayType
		expected    string
	}{
		{OverlayGraphics, "Graphics"},
		{OverlayROI, "ROI"},
	}

	for _, tc := range testCases {
		result := tc.overlayType.String()
		if result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}
}

func TestNewDicomOverlayData(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)

	if overlay.Group != 0x6000 {
		t.Errorf("Expected Group=0x6000, got 0x%X", overlay.Group)
	}

	if overlay.BitsAllocated != 1 {
		t.Errorf("Expected BitsAllocated=1, got %d", overlay.BitsAllocated)
	}

	if overlay.NumberOfFrames != 1 {
		t.Errorf("Expected NumberOfFrames=1, got %d", overlay.NumberOfFrames)
	}
}

func TestDicomOverlayData_IsValid(t *testing.T) {
	testCases := []struct {
		name     string
		group    uint16
		expected bool
	}{
		{"valid 0x6000", 0x6000, true},
		{"valid 0x6002", 0x6002, true},
		{"valid 0x60FE", 0x60FE, true},
		{"invalid: too low", 0x5FFE, false},
		{"invalid: too high", 0x6100, false},
		{"invalid: odd", 0x6001, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			overlay := NewDicomOverlayData(tc.group)
			result := overlay.IsValid()
			if result != tc.expected {
				t.Errorf("IsValid for group 0x%X: expected %v, got %v",
					tc.group, tc.expected, result)
			}
		})
	}
}

func TestDicomOverlayData_UnpackedDataSize(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 512
	overlay.Columns = 512
	overlay.NumberOfFrames = 1

	expected := 512 * 512 * 1
	result := overlay.UnpackedDataSize()

	if result != expected {
		t.Errorf("UnpackedDataSize: expected %d, got %d", expected, result)
	}

	// Multi-frame
	overlay.NumberOfFrames = 10
	expected = 512 * 512 * 10
	result = overlay.UnpackedDataSize()

	if result != expected {
		t.Errorf("UnpackedDataSize (multi-frame): expected %d, got %d", expected, result)
	}
}

func TestDicomOverlayData_PackedDataSize(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 100
	overlay.Columns = 100
	overlay.NumberOfFrames = 1

	// 100*100 = 10000 bits = 1250 bytes
	expected := 1250
	result := overlay.PackedDataSize()

	if result != expected {
		t.Errorf("PackedDataSize: expected %d, got %d", expected, result)
	}

	// Test with size not divisible by 8
	overlay.Rows = 101
	overlay.Columns = 101
	// 101*101 = 10201 bits = 1276 bytes (rounded up)
	expected = 1276
	result = overlay.PackedDataSize()

	if result != expected {
		t.Errorf("PackedDataSize (rounded): expected %d, got %d", expected, result)
	}
}

func TestDicomOverlayData_PackUnpack(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 4
	overlay.Columns = 4
	overlay.NumberOfFrames = 1

	// Create test data (4x4 = 16 pixels)
	unpacked := []byte{
		255, 0, 255, 0,
		0, 255, 0, 255,
		255, 0, 255, 0,
		0, 255, 0, 255,
	}

	// Pack
	err := overlay.PackData(unpacked)
	if err != nil {
		t.Fatalf("PackData failed: %v", err)
	}

	// Check packed size (16 bits = 2 bytes)
	if len(overlay.Data) != 2 {
		t.Errorf("Packed size: expected 2 bytes, got %d", len(overlay.Data))
	}

	// Unpack
	result := overlay.UnpackData()

	// Verify
	if len(result) != len(unpacked) {
		t.Fatalf("Unpacked size: expected %d, got %d", len(unpacked), len(result))
	}

	for i := 0; i < len(unpacked); i++ {
		expected := byte(0)
		if unpacked[i] != 0 {
			expected = 255
		}
		if result[i] != expected {
			t.Errorf("Pixel %d: expected %d, got %d", i, expected, result[i])
		}
	}
}

func TestDicomOverlayData_PackData_InvalidSize(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 4
	overlay.Columns = 4
	overlay.NumberOfFrames = 1

	// Wrong size data
	unpacked := []byte{1, 2, 3} // Should be 16

	err := overlay.PackData(unpacked)
	if err == nil {
		t.Error("Expected error for invalid data size, got nil")
	}
}

func TestDicomOverlayData_GetFrame(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 2
	overlay.Columns = 2
	overlay.NumberOfFrames = 2

	// Create data for 2 frames (2x2x2 = 8 pixels)
	unpacked := []byte{
		255, 0, // Frame 0
		0, 255,
		0, 255, // Frame 1
		255, 0,
	}

	_ = overlay.PackData(unpacked)

	// Get frame 0
	frame0, err := overlay.GetFrame(0)
	if err != nil {
		t.Fatalf("GetFrame(0) failed: %v", err)
	}

	if len(frame0) != 4 {
		t.Fatalf("Frame 0 size: expected 4, got %d", len(frame0))
	}

	// Check frame 0 values
	if frame0[0] != 255 || frame0[1] != 0 || frame0[2] != 0 || frame0[3] != 255 {
		t.Errorf("Frame 0 values incorrect: %v", frame0)
	}

	// Get frame 1
	frame1, err := overlay.GetFrame(1)
	if err != nil {
		t.Fatalf("GetFrame(1) failed: %v", err)
	}

	if frame1[0] != 0 || frame1[1] != 255 || frame1[2] != 255 || frame1[3] != 0 {
		t.Errorf("Frame 1 values incorrect: %v", frame1)
	}

	// Invalid frame index
	_, err = overlay.GetFrame(2)
	if err == nil {
		t.Error("Expected error for invalid frame index, got nil")
	}
}

func TestDicomOverlayData_ApplyToImage_Grayscale(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 4
	overlay.Columns = 4
	overlay.NumberOfFrames = 1

	// Create overlay data (checkerboard)
	unpacked := []byte{
		255, 0, 255, 0,
		0, 255, 0, 255,
		255, 0, 255, 0,
		0, 255, 0, 255,
	}
	_ = overlay.PackData(unpacked)

	// Create grayscale image (all 128)
	imageData := make([]byte, 16)
	for i := range imageData {
		imageData[i] = 128
	}

	// Apply overlay
	overlayColor := colorconv.NewColor32(255, 200, 200, 200) // Light gray
	result, err := overlay.ApplyToImage(imageData, 4, 4, 1, overlayColor)
	if err != nil {
		t.Fatalf("ApplyToImage failed: %v", err)
	}

	// Check that overlay pixels were modified
	if result[0] == 128 {
		t.Error("Overlay pixel should have been modified")
	}

	// Check that non-overlay pixels remained the same
	if result[1] != 128 {
		t.Error("Non-overlay pixel should remain unchanged")
	}
}

func TestDicomOverlayData_ApplyToImage_RGB(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 2
	overlay.Columns = 2
	overlay.NumberOfFrames = 1

	// Simple overlay (top-left pixel only)
	unpacked := []byte{255, 0, 0, 0}
	_ = overlay.PackData(unpacked)

	// Create RGB image (all black)
	imageData := make([]byte, 2*2*3)

	// Apply red overlay
	overlayColor := colorconv.NewColor32(255, 255, 0, 0) // Red
	result, err := overlay.ApplyToImage(imageData, 2, 2, 3, overlayColor)
	if err != nil {
		t.Fatalf("ApplyToImage failed: %v", err)
	}

	// First pixel should be blended with red
	if result[0] == 0 {
		t.Error("Overlay pixel R should have been modified")
	}
}

func TestDicomOverlayData_ApplyToImage_DimensionMismatch(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 4
	overlay.Columns = 4

	imageData := make([]byte, 16)

	overlayColor := colorconv.NewColor32(255, 255, 255, 255)

	// Mismatched dimensions
	_, err := overlay.ApplyToImage(imageData, 8, 8, 1, overlayColor)
	if err == nil {
		t.Error("Expected error for dimension mismatch, got nil")
	}
}

func TestOverlayGroupNumbers(t *testing.T) {
	groups := OverlayGroupNumbers()

	if len(groups) != 128 {
		t.Errorf("Expected 128 overlay groups, got %d", len(groups))
	}

	// Check first and last
	if groups[0] != 0x6000 {
		t.Errorf("First group: expected 0x6000, got 0x%X", groups[0])
	}

	if groups[127] != 0x60FE {
		t.Errorf("Last group: expected 0x60FE, got 0x%X", groups[127])
	}

	// Check that all are even
	for i, g := range groups {
		if g%2 != 0 {
			t.Errorf("Group[%d]=0x%X is odd", i, g)
		}
	}
}

func TestDicomOverlayData_ApplyToImageAlpha(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 2
	overlay.Columns = 2
	if err := overlay.PackData([]byte{255, 0, 0, 0}); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}

	gray := []byte{100, 100, 100, 100}
	for _, test := range []struct {
		name string
		rgb  colorconv.Color32
		want uint8
	}{
		{name: "transparent gray", rgb: colorconv.NewColor32(0, 200, 0, 0), want: 100},
		{name: "opaque gray", rgb: colorconv.NewColor32(255, 200, 0, 0), want: 66},
	} {
		t.Run("grayscale/"+test.name, func(t *testing.T) {
			result, err := overlay.ApplyToImage(gray, 2, 2, 1, test.rgb)
			if err != nil {
				t.Fatalf("ApplyToImage() error = %v", err)
			}
			if result[0] != test.want {
				t.Errorf("marked pixel = %d, want %d", result[0], test.want)
			}
			if result[1] != 100 {
				t.Errorf("unmarked pixel = %d, want 100", result[1])
			}
		})
	}

	rgb := []byte{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	for _, test := range []struct {
		name string
		rgb  colorconv.Color32
		want [3]uint8
	}{
		{name: "transparent rgb", rgb: colorconv.NewColor32(0, 255, 0, 0), want: [3]uint8{100, 100, 100}},
		{name: "opaque rgb", rgb: colorconv.NewColor32(255, 255, 0, 0), want: [3]uint8{255, 0, 0}},
	} {
		t.Run("rgb/"+test.name, func(t *testing.T) {
			result, err := overlay.ApplyToImage(rgb, 2, 2, 3, test.rgb)
			if err != nil {
				t.Fatalf("ApplyToImage() error = %v", err)
			}
			if got := [3]uint8{result[0], result[1], result[2]}; got != test.want {
				t.Errorf("marked pixel = %v, want %v", got, test.want)
			}
			if got := [3]uint8{result[3], result[4], result[5]}; got != [3]uint8{100, 100, 100} {
				t.Errorf("unmarked pixel = %v, want [100 100 100]", got)
			}
		})
	}
}

func TestDicomOverlayData_ApplyToImageUsesFirstFrame(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 2
	overlay.Columns = 2
	overlay.NumberOfFrames = 2
	// Only the second frame marks a pixel; the single image frame only sees
	// overlay frame 0.
	if err := overlay.PackData([]byte{0, 0, 0, 0, 255, 0, 0, 0}); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}
	gray := []byte{100, 100, 100, 100}
	result, err := overlay.ApplyToImage(gray, 2, 2, 1, colorconv.NewColor32(255, 255, 255, 255))
	if err != nil {
		t.Fatalf("ApplyToImage() error = %v", err)
	}
	for index, value := range result {
		if value != 100 {
			t.Fatalf("pixel %d = %d, want 100 (frame 1 must not be composited)", index, value)
		}
	}
}

func TestApplyOverlaysBlendsAlpha(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 2
	overlay.Columns = 2
	overlay.OriginRow = 1
	overlay.OriginColumn = 1
	overlay.ImageFrameOrigin = 1
	if err := overlay.PackData([]byte{255, 0, 0, 0}); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}
	source := image.NewRGBA(image.Rect(0, 0, 2, 2))
	for index := 0; index < len(source.Pix); index += 4 {
		source.Pix[index] = 100
		source.Pix[index+1] = 100
		source.Pix[index+2] = 100
		source.Pix[index+3] = 255
	}

	for _, test := range []struct {
		name string
		rgb  colorconv.Color32
		want color.NRGBA
	}{
		{name: "transparent overlay", rgb: colorconv.NewColor32(0, 255, 0, 0), want: color.NRGBA{R: 100, G: 100, B: 100, A: 255}},
		{name: "opaque overlay", rgb: colorconv.NewColor32(255, 255, 0, 0), want: color.NRGBA{R: 255, A: 255}},
	} {
		t.Run(test.name, func(t *testing.T) {
			img := &DicomImage{
				showOverlays: true,
				overlays:     []*DicomOverlayData{overlay},
				overlayColor: test.rgb,
			}
			result := img.applyOverlays(source, 0)
			if got := color.NRGBAModel.Convert(result.At(0, 0)).(color.NRGBA); got != test.want {
				t.Errorf("marked pixel = %#v, want %#v", got, test.want)
			}
			want := color.NRGBA{R: 100, G: 100, B: 100, A: 255}
			if got := color.NRGBAModel.Convert(result.At(1, 0)).(color.NRGBA); got != want {
				t.Errorf("unmarked pixel = %#v, want %#v", got, want)
			}
		})
	}
}

// TestApplyOverlaysSingleFrameOverlayCoversEveryFrame verifies PS3.3 C.9.2.1.4:
// an overlay without Number of Frames in Overlay or Image Frame Origin applies
// to all frames of a multi-frame image.
func TestApplyOverlaysSingleFrameOverlayCoversEveryFrame(t *testing.T) {
	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows = 2
	overlay.Columns = 2
	overlay.OriginRow = 1
	overlay.OriginColumn = 1
	overlay.ImageFrameOrigin = 1
	if err := overlay.PackData([]byte{255, 0, 0, 0}); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}
	source := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img := &DicomImage{
		showOverlays: true,
		overlays:     []*DicomOverlayData{overlay},
		overlayColor: colorconv.NewColor32(255, 255, 0, 0),
	}

	want := color.NRGBA{R: 255, A: 255}
	for _, frame := range []int{0, 1, 7} {
		result := img.applyOverlays(source, frame)
		if got := color.NRGBAModel.Convert(result.At(0, 0)).(color.NRGBA); got != want {
			t.Errorf("frame %d marked pixel = %#v, want %#v", frame, got, want)
		}
	}
}

// TestApplyOverlaysKeepsSourceWhenNothingDraws verifies an overlay that cannot
// contribute pixels (no set bits, frame out of range, transparent color) does
// not force an RGBA copy of the image.
func TestApplyOverlaysKeepsSourceWhenNothingDraws(t *testing.T) {
	blank := NewDicomOverlayData(0x6000)
	blank.Rows, blank.Columns = 2, 2
	blank.OriginRow, blank.OriginColumn = 1, 1
	blank.ImageFrameOrigin = 1
	blank.Data = []byte{0}

	outOfRange := NewDicomOverlayData(0x6002)
	outOfRange.Rows, outOfRange.Columns = 2, 2
	outOfRange.OriginRow, outOfRange.OriginColumn = 1, 1
	outOfRange.ImageFrameOrigin = 5
	outOfRange.MultiFrameOverlay = true
	if err := outOfRange.PackData([]byte{255, 255, 255, 255}); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}

	marked := NewDicomOverlayData(0x6004)
	marked.Rows, marked.Columns = 2, 2
	marked.OriginRow, marked.OriginColumn = 1, 1
	marked.ImageFrameOrigin = 1
	if err := marked.PackData([]byte{255, 0, 0, 0}); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}

	tests := []struct {
		name     string
		overlays []*DicomOverlayData
		color    colorconv.Color32
	}{
		{"blank overlay", []*DicomOverlayData{blank}, colorconv.NewColor32(255, 255, 255, 255)},
		{"overlay frame out of range", []*DicomOverlayData{outOfRange}, colorconv.NewColor32(255, 255, 255, 255)},
		{"transparent overlay color", []*DicomOverlayData{marked}, colorconv.NewColor32(0, 255, 255, 255)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := image.NewGray(image.Rect(0, 0, 2, 2))
			img := &DicomImage{
				showOverlays: true,
				overlays:     test.overlays,
				overlayColor: test.color,
			}
			result := img.applyOverlays(source, 0)
			gray, ok := result.(*image.Gray)
			if !ok || gray != source {
				t.Fatalf("applyOverlays() = %T, want the original *image.Gray", result)
			}
		})
	}
}

// TestOneBitImageWithOverlayScalesNearestNeighbor verifies that an overlay
// composited into RGBA on a 1-bit image does not make the scaler interpolate
// gray levels between the two stored values.
func TestOneBitImageWithOverlayScalesNearestNeighbor(t *testing.T) {
	const width, height = 8, 8
	pixelData, err := pixeldata.NewFromBytes(&pixeldata.Info{
		Width: width, Height: height, NumberOfFrames: 1,
		BitsAllocated: 1, BitsStored: 1, HighBit: 0, SamplesPerPixel: 1,
		PixelRepresentation: pixel.UnsignedPixels, PhotometricInterpretation: pixel.Monochrome2,
	}, []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA})
	if err != nil {
		t.Fatalf("pixeldata.NewFromBytes() error = %v", err)
	}

	overlay := NewDicomOverlayData(0x6000)
	overlay.Rows, overlay.Columns = height, width
	overlay.OriginRow, overlay.OriginColumn = 1, 1
	overlay.ImageFrameOrigin = 1
	marked := make([]byte, width*height)
	marked[0] = 255
	if err := overlay.PackData(marked); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}

	dicomImage := NewDicomImage(pixelData)
	dicomImage.overlays = []*DicomOverlayData{overlay}
	dicomImage.overlayColor = colorconv.NewColor32(255, 255, 0, 0)
	dicomImage.SetScale(2)

	rendered, err := dicomImage.RenderFrameImage(0)
	if err != nil {
		t.Fatalf("RenderFrameImage(0) error = %v", err)
	}
	if got, want := rendered.Bounds(), image.Rect(0, 0, 2*width, 2*height); got != want {
		t.Fatalf("bounds = %v, want %v", got, want)
	}
	overlayColor := color.RGBA{R: 255, A: 255}
	for y := rendered.Bounds().Min.Y; y < rendered.Bounds().Max.Y; y++ {
		for x := rendered.Bounds().Min.X; x < rendered.Bounds().Max.X; x++ {
			value := color.RGBAModel.Convert(rendered.At(x, y)).(color.RGBA)
			if value == overlayColor {
				continue
			}
			if value.R != value.G || value.G != value.B {
				t.Fatalf("pixel (%d,%d) = %#v, want a gray or the overlay color", x, y, value)
			}
			if value.R != 0 && value.R != 255 {
				t.Fatalf("pixel (%d,%d) gray = %d, want 0 or 255 (interpolated scaling)", x, y, value.R)
			}
		}
	}
}

// TestApplyToImageRejectsShortBuffers verifies the composited buffer is checked
// against the whole image instead of panicking on the first marked pixel.
func TestApplyToImageRejectsShortBuffers(t *testing.T) {
	grayOverlay := NewDicomOverlayData(0x6000)
	grayOverlay.Rows, grayOverlay.Columns = 2, 2
	if err := grayOverlay.PackData([]byte{255, 0, 0, 255}); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}
	if _, err := grayOverlay.ApplyToImage(make([]byte, 3), 2, 2, 1, colorconv.NewColor32(0, 255, 255, 255)); err == nil {
		t.Fatal("ApplyToImage() with a short gray buffer = nil error, want a length error")
	}

	rgbOverlay := NewDicomOverlayData(0x6002)
	rgbOverlay.Rows, rgbOverlay.Columns = 1, 1
	if err := rgbOverlay.PackData([]byte{255}); err != nil {
		t.Fatalf("PackData() error = %v", err)
	}
	if _, err := rgbOverlay.ApplyToImage(make([]byte, 2), 1, 1, 3, colorconv.NewColor32(0, 255, 255, 255)); err == nil {
		t.Fatal("ApplyToImage() with a short RGB buffer = nil error, want a length error")
	}
}
