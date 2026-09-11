// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
)

type contractSource struct {
	frameRead bool
	info      FrameInfo
}

func (s *contractSource) FrameCount() int { return 1 }

func (s *contractSource) Frame(ctx context.Context, _ int) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	s.frameRead = true
	return []byte{0x34, 0x12}, nil
}

func (s *contractSource) FrameInfo() FrameInfo { return s.info }

func (s *contractSource) Encapsulated() bool { return false }

type contractSink struct {
	frames [][]byte
	info   FrameInfo
}

func (s *contractSink) AddFrame(ctx context.Context, frame []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.frames = append(s.frames, append([]byte(nil), frame...))
	return nil
}

func (s *contractSink) SetFrameInfo(info FrameInfo) error {
	s.info = info
	return nil
}

type wrongParameters struct{}

func (wrongParameters) Clone() Parameters { return wrongParameters{} }
func (wrongParameters) Validate() error   { return nil }

type thirdPartyParameters struct {
	values  []int
	invalid bool
}

func (p *thirdPartyParameters) Clone() Parameters {
	if p == nil {
		return (*thirdPartyParameters)(nil)
	}
	return &thirdPartyParameters{
		values:  append([]int(nil), p.values...),
		invalid: p.invalid,
	}
}

func (p *thirdPartyParameters) Validate() error {
	if p.invalid {
		return errors.New("unsupported quality")
	}
	return nil
}

type parameterContractCodec struct {
	defaults Parameters
}

func (parameterContractCodec) Name() string { return "parameter-contract" }

func (parameterContractCodec) TransferSyntax() *transfer.Syntax { return transfer.JPEGBaseline8Bit }

func (c parameterContractCodec) DefaultParameters() Parameters { return c.defaults }

func (parameterContractCodec) Encode(context.Context, FrameSource, FrameSink, Parameters) error {
	return nil
}

func (parameterContractCodec) Decode(context.Context, FrameSource, FrameSink, Parameters) error {
	return nil
}

type typedNilCodec struct{}

func (*typedNilCodec) Name() string { return "typed-nil" }

func (*typedNilCodec) TransferSyntax() *transfer.Syntax { return transfer.JPEGBaseline8Bit }

func (*typedNilCodec) DefaultParameters() Parameters { return NoParameters{} }

func (*typedNilCodec) Encode(context.Context, FrameSource, FrameSink, Parameters) error { return nil }

func (*typedNilCodec) Decode(context.Context, FrameSource, FrameSink, Parameters) error { return nil }

type panicOnNilParameters struct{}

func (p *panicOnNilParameters) Clone() Parameters {
	if p == nil {
		panic("Clone called on typed-nil parameters")
	}
	return &panicOnNilParameters{}
}

func (*panicOnNilParameters) Validate() error { return nil }

type typedNilCloneParameters struct{}

func (typedNilCloneParameters) Clone() Parameters { return (*panicOnNilParameters)(nil) }

func (typedNilCloneParameters) Validate() error { return nil }

func nativeContractFrameInfo() FrameInfo {
	return FrameInfo{
		Width:                     1,
		Height:                    1,
		BitDepth:                  *pixel.NewBitDepth(16, 12, 11, false),
		SamplesPerPixel:           1,
		PixelRepresentation:       pixel.UnsignedPixels,
		PlanarConfiguration:       pixel.InterleavedPlanar,
		PhotometricInterpretation: *pixel.Monochrome2,
	}
}

func TestNativeCodecCanceledBeforeFrameRead(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	source := &contractSource{info: nativeContractFrameInfo()}
	sink := &contractSink{}

	err := NewExplicitVRLittleEndianCodec().Encode(ctx, source, sink, NativeParameters{})

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Encode() error = %v, want context.Canceled", err)
	}
	if source.frameRead {
		t.Fatal("Encode() read a frame after context cancellation")
	}
	if len(sink.frames) != 0 {
		t.Fatalf("Encode() wrote %d frames after cancellation", len(sink.frames))
	}
}

func TestNativeCodecRejectsWrongParameterType(t *testing.T) {
	source := &contractSource{info: nativeContractFrameInfo()}
	sink := &contractSink{}

	err := NewExplicitVRLittleEndianCodec().Encode(
		context.Background(), source, sink, wrongParameters{},
	)

	if !errors.Is(err, ErrInvalidParameters) {
		t.Fatalf("Encode() error = %v, want ErrInvalidParameters", err)
	}
	if source.frameRead {
		t.Fatal("Encode() started frame work with invalid parameters")
	}
}

func TestNativeCodecRejectsInvalidNativeParameters(t *testing.T) {
	source := &contractSource{info: nativeContractFrameInfo()}
	sink := &contractSink{}

	err := NewExplicitVRLittleEndianCodec().Encode(
		context.Background(),
		source,
		sink,
		NativeParameters{ByteSwap: ByteSwapMode(255)},
	)

	if !errors.Is(err, ErrInvalidParameters) {
		t.Fatalf("Encode() error = %v, want ErrInvalidParameters", err)
	}
	if source.frameRead {
		t.Fatal("Encode() started frame work with invalid native parameters")
	}
}

func TestNativeParametersCloneIsIndependentValue(t *testing.T) {
	original := NativeParameters{ByteSwap: ByteSwapEnabled}
	clone, ok := original.Clone().(NativeParameters)
	if !ok {
		t.Fatalf("Clone() type = %T, want NativeParameters", original.Clone())
	}
	clone.ByteSwap = ByteSwapDisabled
	if original.ByteSwap != ByteSwapEnabled {
		t.Fatal("modifying cloned parameters changed the original")
	}
	if err := original.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestPrepareParametersClonesThirdPartyParameters(t *testing.T) {
	original := &thirdPartyParameters{values: []int{90}}
	prepared, err := PrepareParameters(parameterContractCodec{}, original)
	if err != nil {
		t.Fatalf("PrepareParameters() error = %v", err)
	}
	clone, ok := prepared.(*thirdPartyParameters)
	if !ok {
		t.Fatalf("PrepareParameters() type = %T, want *thirdPartyParameters", prepared)
	}
	clone.values[0] = 10
	if original.values[0] != 90 {
		t.Fatalf("codec parameter mutation changed caller value to %d", original.values[0])
	}
}

func TestPrepareParametersClonesDefaultsOnEveryCall(t *testing.T) {
	defaults := &thirdPartyParameters{values: []int{90}}
	c := parameterContractCodec{defaults: defaults}

	first, err := PrepareParameters(c, nil)
	if err != nil {
		t.Fatalf("first PrepareParameters() error = %v", err)
	}
	second, err := PrepareParameters(c, nil)
	if err != nil {
		t.Fatalf("second PrepareParameters() error = %v", err)
	}
	first.(*thirdPartyParameters).values[0] = 10
	if got := second.(*thirdPartyParameters).values[0]; got != 90 {
		t.Fatalf("second prepared defaults changed to %d", got)
	}
	if got := defaults.values[0]; got != 90 {
		t.Fatalf("codec defaults changed to %d", got)
	}
}

func TestPrepareParametersWrapsValidationError(t *testing.T) {
	cause := &thirdPartyParameters{invalid: true}
	_, err := PrepareParameters(parameterContractCodec{}, cause)
	if !errors.Is(err, ErrInvalidParameters) {
		t.Fatalf("PrepareParameters() error = %v, want ErrInvalidParameters", err)
	}
	if err == nil || !strings.Contains(err.Error(), "unsupported quality") {
		t.Fatalf("PrepareParameters() error = %v, want validation context", err)
	}
}

func TestPrepareParametersRejectsTypedNilValuesWithoutPanic(t *testing.T) {
	tests := []struct {
		name       string
		codec      Codec
		parameters Parameters
	}{
		{
			name:       "codec",
			codec:      (*typedNilCodec)(nil),
			parameters: NoParameters{},
		},
		{
			name:       "supplied parameters",
			codec:      parameterContractCodec{},
			parameters: (*panicOnNilParameters)(nil),
		},
		{
			name:  "default parameters",
			codec: parameterContractCodec{defaults: (*panicOnNilParameters)(nil)},
		},
		{
			name:       "parameter clone",
			codec:      parameterContractCodec{},
			parameters: typedNilCloneParameters{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Fatalf("PrepareParameters() panic = %v", recovered)
				}
			}()
			_, err := PrepareParameters(tt.codec, tt.parameters)
			if !errors.Is(err, ErrInvalidParameters) {
				t.Fatalf("PrepareParameters() error = %v, want ErrInvalidParameters", err)
			}
		})
	}
}
