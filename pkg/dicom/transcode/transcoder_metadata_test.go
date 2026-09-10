// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transcode

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
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
	if err := ds.Add(element.NewOtherByte(tag.PixelData, []byte{1, 2, 3, 4, 5, 6, 7, 8})); err != nil {
		t.Fatal(err)
	}

	transcoder := newTestTranscoder(t,
		transfer.ExplicitVRLittleEndian,
		transfer.JPEGBaseline8Bit,
		metadataCodec{frames: [][]byte{{1, 2}, {3, 4}}},
	)
	result, err := transcoder.Transcode(context.Background(), ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	if got := result.TryGetString(tag.PhotometricInterpretation); got != pixel.RGBPhotometric.Value {
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

	transcoder := newTestTranscoder(t,
		transfer.JPEGBaseline8Bit,
		transfer.ExplicitVRLittleEndian,
		metadataCodec{frames: [][]byte{{4, 5, 6}}},
	)
	result, err := transcoder.Transcode(context.Background(), ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	if got := result.TryGetString(tag.PhotometricInterpretation); got != pixel.RGBPhotometric.Value {
		t.Fatalf("PhotometricInterpretation = %q, want RGB", got)
	}
	if got := result.TryGetUInt16(tag.PlanarConfiguration, 99); got != 0 {
		t.Fatalf("PlanarConfiguration = %d, want 0", got)
	}
}

func TestTranscoderRejectsInvalidCodecOutputMetadata(t *testing.T) {
	ds := metadataTestDataset(t, transfer.ExplicitVRLittleEndian, "MONOCHROME2", 0, 1)
	if err := ds.Add(element.NewOtherByte(tag.PixelData, []byte{1})); err != nil {
		t.Fatal(err)
	}

	_, err := newTestTranscoder(t,
		transfer.ExplicitVRLittleEndian,
		transfer.JPEGBaseline8Bit,
		invalidMetadataCodec{},
	).Transcode(context.Background(), ds)
	if err == nil {
		t.Fatal("Transcode() error = nil, want invalid codec metadata error")
	}
	if !strings.Contains(err.Error(), "invalid frame info") {
		t.Fatalf("Transcode() error = %q, want invalid frame info context", err)
	}
}

func metadataTestDataset(t *testing.T, syntax *transfer.Syntax, photometric string, planar uint16, frames int) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(syntax)
	samplesPerPixel := uint16(3)
	if strings.HasPrefix(photometric, "MONOCHROME") {
		samplesPerPixel = 1
	}
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{samplesPerPixel}),
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

type invalidMetadataCodec struct{}

func (invalidMetadataCodec) Name() string { return "invalid-metadata" }

func (invalidMetadataCodec) TransferSyntax() *transfer.Syntax { return transfer.JPEGBaseline8Bit }

func (invalidMetadataCodec) DefaultParameters() codec.Parameters { return codec.NoParameters{} }

func (invalidMetadataCodec) Encode(
	ctx context.Context,
	source codec.FrameSource,
	sink codec.FrameSink,
	_ codec.Parameters,
) error {
	info := source.FrameInfo()
	info.BitDepth.BitsStored = info.BitDepth.BitsAllocated + 1
	if err := sink.SetFrameInfo(info); err != nil {
		return err
	}
	return sink.AddFrame(ctx, []byte{1})
}

func (c invalidMetadataCodec) Decode(
	ctx context.Context,
	source codec.FrameSource,
	sink codec.FrameSink,
	parameters codec.Parameters,
) error {
	return c.Encode(ctx, source, sink, parameters)
}

func (metadataCodec) Name() string { return "metadata" }

func (metadataCodec) TransferSyntax() *transfer.Syntax { return transfer.JPEGBaseline8Bit }

func (metadataCodec) DefaultParameters() codec.Parameters { return codec.NoParameters{} }

func (c metadataCodec) Encode(ctx context.Context, _ codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	return c.writeFramesAndMetadata(ctx, newPixelData)
}

func (c metadataCodec) Decode(ctx context.Context, _ codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	return c.writeFramesAndMetadata(ctx, newPixelData)
}

func (c metadataCodec) writeFramesAndMetadata(ctx context.Context, newPixelData codec.FrameSink) error {
	info := codec.FrameInfo{
		Width:                     1,
		Height:                    1,
		BitDepth:                  *pixel.NewBitDepth(8, 8, 7, false),
		SamplesPerPixel:           3,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: *pixel.RGBPhotometric,
	}
	if err := newPixelData.SetFrameInfo(info); err != nil {
		return errors.New("output pixel data does not support frame metadata updates")
	}
	for _, frame := range c.frames {
		if err := newPixelData.AddFrame(ctx, frame); err != nil {
			return err
		}
	}
	return nil
}
