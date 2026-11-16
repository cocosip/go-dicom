// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package imaging

import (
	"context"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
)

type blockingDecodeCodec struct {
	entered chan struct{}
	release chan struct{}
}

func (blockingDecodeCodec) Name() string { return "blocking decode race test codec" }
func (blockingDecodeCodec) TransferSyntax() *transfer.Syntax {
	return transfer.JPEG2000Lossless
}
func (blockingDecodeCodec) DefaultParameters() codec.Parameters { return codec.NoParameters{} }
func (blockingDecodeCodec) Encode(context.Context, codec.FrameSource, codec.FrameSink, codec.Parameters) error {
	return nil
}
func (c *blockingDecodeCodec) Decode(context.Context, codec.FrameSource, codec.FrameSink, codec.Parameters) error {
	close(c.entered)
	<-c.release
	return nil
}

// TestDicomImageDecodeIsSafeForConcurrentReaders guards the pixelData swap in
// DecodeIfNeeded: a reader may inspect the image while another goroutine
// decodes compressed pixel data. Run with -race to detect a regression.
func TestDicomImageDecodeIsSafeForConcurrentReaders(t *testing.T) {
	pd, err := pixeldata.New(&pixeldata.Info{
		Width:                     2,
		Height:                    2,
		NumberOfFrames:            1,
		BitsAllocated:             8,
		BitsStored:                8,
		HighBit:                   7,
		SamplesPerPixel:           1,
		PhotometricInterpretation: pixel.Monochrome2,
		VRCode:                    "OB",
		Encapsulated:              true,
		TransferSyntaxUID:         "1.2.840.10008.1.2.4.90",
	})
	if err != nil {
		t.Fatalf("pixeldata.New() error = %v", err)
	}
	img := NewDicomImage(pd)
	c := &blockingDecodeCodec{entered: make(chan struct{}), release: make(chan struct{})}

	decodeDone := make(chan error, 1)
	go func() {
		decodeDone <- img.DecodeIfNeeded(c, codec.NoParameters{})
	}()
	<-c.entered

	readerStarted := make(chan struct{})
	readerStop := make(chan struct{})
	readerDone := make(chan struct{})
	go func() {
		defer close(readerDone)
		close(readerStarted)
		for {
			select {
			case <-readerStop:
				return
			default:
				_ = img.Width()
				_ = img.Height()
				_ = img.NumberOfFrames()
				// Rendering and cloning read the pixel data snapshot; they must
				// tolerate the swap happening underneath them.
				_, _ = img.RenderFrameImage(0)
				_ = img.Clone()
			}
		}
	}()
	<-readerStarted

	close(c.release)
	if err := <-decodeDone; err != nil {
		t.Fatalf("DecodeIfNeeded() error = %v", err)
	}
	close(readerStop)
	<-readerDone

	if img.pixelData.Load() == pd {
		t.Fatal("DecodeIfNeeded did not publish new pixel data")
	}
}
