// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package pixeldata

import (
	"context"
	"errors"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

func contractInfo() *Info {
	return &Info{
		Width:                     2,
		Height:                    1,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: pixel.Monochrome2,
	}
}

func TestFrameOwnershipIsIsolatedInBothDirections(t *testing.T) {
	pixels, err := New(contractInfo())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	input := []byte{1, 2}
	if err := pixels.AddFrame(context.Background(), input); err != nil {
		t.Fatalf("AddFrame() error = %v", err)
	}
	input[0] = 9

	first, err := pixels.Frame(context.Background(), 0)
	if err != nil {
		t.Fatalf("Frame() error = %v", err)
	}
	if first[0] != 1 {
		t.Fatalf("Frame()[0] = %d, want owned value 1", first[0])
	}
	first[0] = 8

	second, err := pixels.Frame(context.Background(), 0)
	if err != nil {
		t.Fatalf("Frame() second read error = %v", err)
	}
	if second[0] != 1 {
		t.Fatalf("second Frame()[0] = %d, want internal value 1", second[0])
	}
}

func TestFrameAndAddFramePreserveCancellation(t *testing.T) {
	pixels, err := New(contractInfo())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	canceled, cancel := context.WithCancel(context.Background())
	cancel()

	if err := pixels.AddFrame(canceled, []byte{1, 2}); !errors.Is(err, context.Canceled) {
		t.Fatalf("AddFrame() error = %v, want context.Canceled", err)
	}
	if pixels.FrameCount() != 0 {
		t.Fatalf("FrameCount() = %d after canceled write, want 0", pixels.FrameCount())
	}
	if _, err := pixels.Frame(canceled, 0); !errors.Is(err, context.Canceled) {
		t.Fatalf("Frame() error = %v, want context.Canceled", err)
	}
}

func TestSetFrameInfoRejectsInvalidMetadataWithoutMutation(t *testing.T) {
	pixels, err := New(contractInfo())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	original := pixels.FrameInfo()
	invalid := codec.FrameInfo{
		Width:                     2,
		Height:                    1,
		BitDepth:                  *pixel.NewBitDepth(0, 0, 0, false),
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: *pixel.Monochrome2,
	}

	if err := pixels.SetFrameInfo(invalid); err == nil {
		t.Fatal("SetFrameInfo() accepted invalid bit depth")
	}
	if got := pixels.FrameInfo(); got != original {
		t.Fatalf("FrameInfo() changed after rejected update: got %+v, want %+v", got, original)
	}
}

type mutableParameters struct {
	quality []int
}

func (p *mutableParameters) Clone() codec.Parameters {
	return &mutableParameters{quality: append([]int(nil), p.quality...)}
}

func (*mutableParameters) Validate() error { return nil }

type parameterMutatingCodec struct{}

func (parameterMutatingCodec) Name() string { return "parameter-mutating" }

func (parameterMutatingCodec) TransferSyntax() *transfer.Syntax { return transfer.JPEGBaseline8Bit }

func (parameterMutatingCodec) DefaultParameters() codec.Parameters {
	return &mutableParameters{quality: []int{90}}
}

func (parameterMutatingCodec) Encode(
	ctx context.Context,
	source codec.FrameSource,
	sink codec.FrameSink,
	parameters codec.Parameters,
) error {
	parameters.(*mutableParameters).quality[0] = 10
	frame, err := source.Frame(ctx, 0)
	if err != nil {
		return err
	}
	return sink.AddFrame(ctx, frame)
}

func (c parameterMutatingCodec) Decode(
	ctx context.Context,
	source codec.FrameSource,
	sink codec.FrameSink,
	parameters codec.Parameters,
) error {
	return c.Encode(ctx, source, sink, parameters)
}

func TestEncodeProtectsCallerOwnedThirdPartyParameters(t *testing.T) {
	pixels, err := NewFromBytes(contractInfo(), []byte{1, 2})
	if err != nil {
		t.Fatalf("NewFromBytes() error = %v", err)
	}
	parameters := &mutableParameters{quality: []int{90}}

	if _, err := pixels.Encode(context.Background(), parameterMutatingCodec{}, parameters); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	if got := parameters.quality[0]; got != 90 {
		t.Fatalf("caller parameter quality = %d after codec mutation, want 90", got)
	}
}
