// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestTranscoderEncodeAppliesCodecMetadataAndAppendsLossyHistory(t *testing.T) {
	ds := metadataTestDataset(t, transfer.ExplicitVRLittleEndian, "YBR_FULL_422", 1, 2)
	for _, elem := range []element.Element{
		element.NewString(tag.LossyImageCompression, vr.CS, []string{"01"}),
		element.NewString(tag.LossyImageCompressionMethod, vr.CS, []string{"ISO_15444_1"}),
		element.NewString(tag.LossyImageCompressionRatio, vr.DS, []string{"1.500"}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", elem.Tag(), err)
		}
	}
	if err := ds.Add(element.NewOtherByte(tag.PixelData, []byte{1, 2, 3, 4, 5, 6})); err != nil {
		t.Fatal(err)
	}

	transcoder := NewTranscoder(
		transfer.ExplicitVRLittleEndian,
		transfer.JPEGBaseline8Bit,
		WithOutputCodec(metadataCodec{frames: [][]byte{{1}, {2, 3}}}),
	)
	result, err := transcoder.Transcode(ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	if got := result.TryGetString(tag.PhotometricInterpretation); got != photometricRGB {
		t.Fatalf("PhotometricInterpretation = %q, want RGB", got)
	}
	if got := result.TryGetUInt16(tag.PlanarConfiguration, 99); got != 0 {
		t.Fatalf("PlanarConfiguration = %d, want 0", got)
	}
	if got := result.TryGetString(tag.LossyImageCompression); got != "01" {
		t.Fatalf("LossyImageCompression = %q, want 01", got)
	}
	if got, _ := result.GetStrings(tag.LossyImageCompressionMethod); !reflect.DeepEqual(got, []string{"ISO_15444_1", "ISO_10918_1"}) {
		t.Fatalf("LossyImageCompressionMethod = %v, want prior and current methods", got)
	}
	if got, _ := result.GetStrings(tag.LossyImageCompressionRatio); !reflect.DeepEqual(got, []string{"1.500", "2.000"}) {
		t.Fatalf("LossyImageCompressionRatio = %v, want prior ratio and all-frame ratio", got)
	}
}

func TestTranscoderDecodeAppliesCodecOutputMetadata(t *testing.T) {
	ds := metadataTestDataset(t, transfer.JPEGBaseline8Bit, "YBR_FULL_422", 1, 1)
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{1, 2, 3}))
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	transcoder := NewTranscoder(
		transfer.JPEGBaseline8Bit,
		transfer.ExplicitVRLittleEndian,
		WithInputCodec(metadataCodec{frames: [][]byte{{4, 5, 6}}}),
	)
	result, err := transcoder.Transcode(ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	if got := result.TryGetString(tag.PhotometricInterpretation); got != photometricRGB {
		t.Fatalf("PhotometricInterpretation = %q, want RGB", got)
	}
	if got := result.TryGetUInt16(tag.PlanarConfiguration, 99); got != 0 {
		t.Fatalf("PlanarConfiguration = %d, want 0", got)
	}
}

func metadataTestDataset(t *testing.T, syntax *transfer.Syntax, photometric string, planar uint16, frames int) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(syntax)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{3}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewUnsignedShort(tag.PlanarConfiguration, []uint16{planar}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometric}),
		element.NewString(tag.NumberOfFrames, vr.IS, []string{strconv.Itoa(frames)}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", elem.Tag(), err)
		}
	}
	return ds
}

type metadataCodec struct {
	frames [][]byte
}

func (metadataCodec) Name() string { return "metadata" }

func (metadataCodec) TransferSyntax() *transfer.Syntax { return transfer.JPEGBaseline8Bit }

func (metadataCodec) GetDefaultParameters() Parameters { return NewBaseParameters() }

func (c metadataCodec) Encode(_ imagetypes.PixelData, newPixelData imagetypes.PixelData, _ Parameters) error {
	return c.writeFramesAndMetadata(newPixelData)
}

func (c metadataCodec) Decode(_ imagetypes.PixelData, newPixelData imagetypes.PixelData, _ Parameters) error {
	return c.writeFramesAndMetadata(newPixelData)
}

func (c metadataCodec) writeFramesAndMetadata(newPixelData imagetypes.PixelData) error {
	info := newPixelData.GetFrameInfo()
	info.PhotometricInterpretation = photometricRGB
	info.PlanarConfiguration = 0
	if !imagetypes.SetFrameInfo(newPixelData, info) {
		return errors.New("output pixel data does not support frame metadata updates")
	}
	for _, frame := range c.frames {
		if err := newPixelData.AddFrame(frame); err != nil {
			return err
		}
	}
	return nil
}
