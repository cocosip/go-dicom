// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transcode

import (
	"context"
	"errors"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
)

func TestNewManagerRejectsNilRegistry(t *testing.T) {
	manager, err := NewManager(nil)
	if manager != nil || !errors.Is(err, ErrNilRegistry) {
		t.Fatalf("NewManager(nil) = (%#v, %v), want nil and ErrNilRegistry", manager, err)
	}
}

func TestManagerNewTranscoderRejectsMissingTransferSyntax(t *testing.T) {
	manager, err := NewManager(codec.NewRegistry())
	if err != nil {
		t.Fatal(err)
	}
	for _, syntaxes := range [][2]*transfer.Syntax{
		{nil, transfer.ExplicitVRLittleEndian},
		{transfer.ExplicitVRLittleEndian, nil},
	} {
		transcoder, createErr := manager.NewTranscoder(syntaxes[0], syntaxes[1])
		if transcoder != nil || !errors.Is(createErr, ErrTransferSyntaxRequired) {
			t.Fatalf("NewTranscoder(%v, %v) = (%#v, %v), want nil and ErrTransferSyntaxRequired",
				syntaxes[0], syntaxes[1], transcoder, createErr)
		}
	}
}

func TestManagerRequiresCompressedCodecFromItsRegistry(t *testing.T) {
	manager, err := NewManager(codec.NewRegistry())
	if err != nil {
		t.Fatal(err)
	}
	if manager.CanTranscode(transfer.ExplicitVRLittleEndian, transfer.JPEG2000Lossless) {
		t.Fatal("CanTranscode returned true without a JPEG 2000 codec")
	}
	transcoder, createErr := manager.NewTranscoder(
		transfer.ExplicitVRLittleEndian,
		transfer.JPEG2000Lossless,
	)
	if transcoder != nil || !errors.Is(createErr, ErrCodecUnavailable) {
		t.Fatalf("NewTranscoder() = (%#v, %v), want nil and ErrCodecUnavailable", transcoder, createErr)
	}
}

func TestManagerSupportsNativeTranscodingWithoutExternalCodecs(t *testing.T) {
	manager, err := NewManager(codec.NewRegistry())
	if err != nil {
		t.Fatal(err)
	}
	if !manager.CanTranscode(transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian) {
		t.Fatal("CanTranscode returned false for native transfer syntaxes")
	}
	if _, err := manager.NewTranscoder(transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian); err != nil {
		t.Fatalf("NewTranscoder() error = %v", err)
	}
}

func TestManagerDoesNotFallBackToGlobalRegistry(t *testing.T) {
	globalCodec := managerTestCodec{syntax: transfer.JPEG2000Lossless}
	previous, err := codec.GlobalRegistry().Replace(globalCodec)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if previous != nil {
			_, _ = codec.GlobalRegistry().Replace(previous)
			return
		}
		codec.GlobalRegistry().Unregister(transfer.JPEG2000Lossless)
	})

	manager, err := NewManager(codec.NewRegistry())
	if err != nil {
		t.Fatal(err)
	}
	if manager.CanTranscode(transfer.ExplicitVRLittleEndian, transfer.JPEG2000Lossless) {
		t.Fatal("Manager used a codec registered only in GlobalRegistry")
	}
}

type managerTestParameters struct {
	quality int
}

func (p *managerTestParameters) Clone() codec.Parameters {
	clone := *p
	return &clone
}

func (*managerTestParameters) Validate() error { return nil }

type managerTestCodec struct {
	syntax *transfer.Syntax
}

func (managerTestCodec) Name() string { return "manager-test" }

func (c managerTestCodec) TransferSyntax() *transfer.Syntax { return c.syntax }

func (managerTestCodec) DefaultParameters() codec.Parameters { return &managerTestParameters{} }

func (managerTestCodec) Encode(context.Context, codec.FrameSource, codec.FrameSink, codec.Parameters) error {
	return nil
}

func (managerTestCodec) Decode(context.Context, codec.FrameSource, codec.FrameSink, codec.Parameters) error {
	return nil
}

func TestManagerNewTranscoderAppliesParametersAndStrictVR(t *testing.T) {
	registry := codec.NewRegistry()
	if err := registry.Register(managerTestCodec{syntax: transfer.JPEG2000Lossless}); err != nil {
		t.Fatal(err)
	}
	manager, err := NewManager(registry)
	if err != nil {
		t.Fatal(err)
	}
	inputParameters := &managerTestParameters{quality: 80}
	outputParameters := &managerTestParameters{quality: 90}
	transcoder, err := manager.NewTranscoder(
		transfer.JPEG2000Lossless,
		transfer.JPEG2000Lossless,
		WithInputParameters(inputParameters),
		WithOutputParameters(outputParameters),
		WithStrictDICOMVR(false),
	)
	if err != nil {
		t.Fatal(err)
	}
	if transcoder.inputParams != inputParameters || transcoder.outputParams != outputParameters {
		t.Fatal("NewTranscoder did not retain the supplied codec parameters")
	}
	if transcoder.pixelDataWriteMode != pixeldata.PixelDataCompatible {
		t.Fatal("NewTranscoder ignored WithStrictDICOMVR(false)")
	}
}

func TestManagerNewTranscoderAppliesPixelDataModes(t *testing.T) {
	manager, err := NewManager(codec.NewRegistry())
	if err != nil {
		t.Fatal(err)
	}
	transcoder, err := manager.NewTranscoder(
		transfer.ExplicitVRLittleEndian,
		transfer.ImplicitVRLittleEndian,
		WithPixelDataReadMode(pixeldata.PixelDataStandard),
		WithPixelDataWriteMode(pixeldata.PixelDataCompatible),
	)
	if err != nil {
		t.Fatal(err)
	}
	if transcoder.pixelDataReadMode != pixeldata.PixelDataStandard {
		t.Fatalf("read mode = %v, want standard", transcoder.pixelDataReadMode)
	}
	if transcoder.pixelDataWriteMode != pixeldata.PixelDataCompatible {
		t.Fatal("compatible write mode did not disable strict DICOM VR output")
	}
}

func TestManagerNewTranscoderRejectsInvalidPixelDataModes(t *testing.T) {
	manager, err := NewManager(codec.NewRegistry())
	if err != nil {
		t.Fatal(err)
	}
	for _, option := range []Option{
		WithPixelDataReadMode(pixeldata.VRMode(99)),
		WithPixelDataWriteMode(pixeldata.VRMode(99)),
	} {
		if _, err := manager.NewTranscoder(
			transfer.ExplicitVRLittleEndian,
			transfer.ImplicitVRLittleEndian,
			option,
		); err == nil {
			t.Fatal("NewTranscoder() accepted an invalid Pixel Data VR mode")
		}
	}
}
