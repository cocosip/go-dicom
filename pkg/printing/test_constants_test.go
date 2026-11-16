// Copyright (c) 2025 go-dicom authors.
// Licensed under the MIT License.

package printing

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const (
	testSOPInstanceUID         = "1.2.3.4.5"
	testImageDisplayFormat2x2  = "STANDARD\\2,2"
	testRemoteImageBoxUID      = "2.25.900"
	testStandardOneByOne       = "STANDARD\\1,1"
	testStandardTwoByOne       = "STANDARD\\2,1"
	testPhotometricMonochrome2 = "MONOCHROME2"
	testPhotometricRGB         = "RGB"
	testSmoothingType          = "SMOOTH_1"
)

// newTestImageItem builds a Preformatted Grayscale Image Sequence item that
// carries the Image Pixel Macro attributes required by PS3.3 Table C.13-4.
func newTestImageItem(t *testing.T, rows, columns uint16, pixels []byte) *dataset.Dataset {
	t.Helper()
	item := dataset.New()
	for _, imageElement := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{rows}),
		element.NewUnsignedShort(tag.Columns, []uint16{columns}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{testPhotometricMonochrome2}),
		element.NewOtherByte(tag.PixelData, pixels),
	} {
		if err := item.Add(imageElement); err != nil {
			t.Fatalf("image item Add() error = %v", err)
		}
	}
	return item
}

// attachTestImage gives an Image Box a conformant one-pixel image item.
func attachTestImage(t *testing.T, imageBox *ImageBox) {
	t.Helper()
	if err := imageBox.SetImageSequence(newTestImageItem(t, 1, 2, []byte{0, 1})); err != nil {
		t.Fatalf("SetImageSequence() error = %v", err)
	}
}
