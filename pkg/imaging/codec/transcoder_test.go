// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"bytes"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestNewTranscoder(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		transcoder := NewTranscoder(
			transfer.ExplicitVRLittleEndian,
			transfer.ImplicitVRLittleEndian,
		)

		if transcoder == nil {
			t.Fatal("NewTranscoder returned nil")
		}

		if transcoder.InputSyntax() != transfer.ExplicitVRLittleEndian {
			t.Errorf("InputSyntax = %v, want ExplicitVRLittleEndian", transcoder.InputSyntax())
		}

		if transcoder.OutputSyntax() != transfer.ImplicitVRLittleEndian {
			t.Errorf("OutputSyntax = %v, want ImplicitVRLittleEndian", transcoder.OutputSyntax())
		}
	})
}

func TestTranscoder_TranscodeNoPixelData(t *testing.T) {
	// Create dataset without pixel data
	ds := dataset.New()
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"Test^Patient"}))
	_ = ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	transcoder := NewTranscoder(
		transfer.ExplicitVRLittleEndian,
		transfer.ImplicitVRLittleEndian,
	)

	result, err := transcoder.Transcode(ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	if result == nil {
		t.Fatal("Transcode() returned nil dataset")
	}

	// Verify data is preserved
	if !result.Contains(tag.PatientName) {
		t.Error("PatientName not found in transcoded dataset")
	}
}

func TestTranscoder_TranscodeUncompressedToUncompressed(t *testing.T) {
	// Create dataset with uncompressed pixel data
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{512}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{512}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

	// Create simple pixel data (512x512 = 262144 bytes)
	pixelData := make([]byte, 512*512)
	for i := range pixelData {
		pixelData[i] = byte(i % 256)
	}
	_ = ds.Add(element.NewOtherByte(tag.PixelData, pixelData))

	transcoder := NewTranscoder(
		transfer.ExplicitVRLittleEndian,
		transfer.ImplicitVRLittleEndian,
	)

	result, err := transcoder.Transcode(ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}

	if result == nil {
		t.Fatal("Transcode() returned nil dataset")
	}

	// Verify pixel data is present
	if !result.Contains(tag.PixelData) {
		t.Error("PixelData not found in transcoded dataset")
	}
}

func TestTranscoder_DecodeFrame(t *testing.T) {
	t.Run("UncompressedSingleFrame", func(t *testing.T) {
		// Create dataset with uncompressed pixel data
		ds := dataset.New()
		_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{10}))
		_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{10}))
		_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
		_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
		_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
		_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
		_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
		_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

		pixelData := make([]byte, 10*10)
		for i := range pixelData {
			pixelData[i] = byte(i)
		}
		_ = ds.Add(element.NewOtherByte(tag.PixelData, pixelData))

		transcoder := NewTranscoder(
			transfer.ExplicitVRLittleEndian,
			transfer.ExplicitVRLittleEndian,
		)

		frame, err := transcoder.DecodeFrame(ds, 0)
		if err != nil {
			t.Fatalf("DecodeFrame() error = %v", err)
		}

		if len(frame) != 100 {
			t.Errorf("DecodeFrame() frame size = %d, want 100", len(frame))
		}
	})

	t.Run("InvalidFrameIndex", func(t *testing.T) {
		ds := dataset.New()
		_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{10}))
		_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{10}))
		_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
		_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
		_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
		_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
		_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
		_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

		pixelData := make([]byte, 10*10)
		_ = ds.Add(element.NewOtherByte(tag.PixelData, pixelData))

		transcoder := NewTranscoder(
			transfer.ExplicitVRLittleEndian,
			transfer.ExplicitVRLittleEndian,
		)

		_, err := transcoder.DecodeFrame(ds, 5) // Frame 5 doesn't exist
		if err == nil {
			t.Error("DecodeFrame() should return error for invalid frame index")
		}
	})
}

func TestFramesFromFragments_UsesBOTItemOffsets(t *testing.T) {
	fragments := []buffer.ByteBuffer{
		buffer.NewMemory([]byte("AA")),
		buffer.NewMemory([]byte("BBB")),
		buffer.NewMemory([]byte("CC")),
	}

	frames, err := framesFromFragments(fragments, []uint32{0, 22}, 2)
	if err != nil {
		t.Fatalf("framesFromFragments() error = %v", err)
	}
	if len(frames) != 2 {
		t.Fatalf("len(frames) = %d, want 2", len(frames))
	}
	if !bytes.Equal(frames[0], []byte("AABBB")) {
		t.Errorf("frames[0] = %q, want %q", frames[0], []byte("AABBB"))
	}
	if !bytes.Equal(frames[1], []byte("CC")) {
		t.Errorf("frames[1] = %q, want %q", frames[1], []byte("CC"))
	}
}

func TestCodecRegistry(t *testing.T) {
	t.Run("RegisterAndGetCodec", func(t *testing.T) {
		registry := NewCodecRegistry()

		codec := NewNativeCodec(transfer.ExplicitVRLittleEndian, false)
		registry.RegisterCodec(transfer.ExplicitVRLittleEndian, codec)

		retrieved, ok := registry.GetCodec(transfer.ExplicitVRLittleEndian)
		if !ok {
			t.Fatal("GetCodec() returned false for registered codec")
		}

		if retrieved != codec {
			t.Error("GetCodec() returned different codec instance")
		}
	})

	t.Run("HasCodec", func(t *testing.T) {
		registry := NewCodecRegistry()

		if registry.HasCodec(transfer.ExplicitVRLittleEndian) {
			t.Error("HasCodec() returned true for unregistered codec")
		}

		codec := NewNativeCodec(transfer.ExplicitVRLittleEndian, false)
		registry.RegisterCodec(transfer.ExplicitVRLittleEndian, codec)

		if !registry.HasCodec(transfer.ExplicitVRLittleEndian) {
			t.Error("HasCodec() returned false for registered codec")
		}
	})

	t.Run("UnregisterCodec", func(t *testing.T) {
		registry := NewCodecRegistry()

		codec := NewNativeCodec(transfer.ExplicitVRLittleEndian, false)
		registry.RegisterCodec(transfer.ExplicitVRLittleEndian, codec)

		registry.UnregisterCodec(transfer.ExplicitVRLittleEndian)

		if registry.HasCodec(transfer.ExplicitVRLittleEndian) {
			t.Error("HasCodec() returned true after unregistering codec")
		}
	})

	t.Run("ListCodecs", func(t *testing.T) {
		registry := NewCodecRegistry()

		explicitCodec := NewNativeCodec(transfer.ExplicitVRLittleEndian, false)
		implicitCodec := NewNativeCodec(transfer.ImplicitVRLittleEndian, false)
		registry.RegisterCodec(transfer.ExplicitVRLittleEndian, explicitCodec)
		registry.RegisterCodec(transfer.ImplicitVRLittleEndian, implicitCodec)

		codecs := registry.ListCodecs()
		if len(codecs) != 2 {
			t.Errorf("ListCodecs() returned %d codecs, want 2", len(codecs))
		}
	})
}

func TestTranscoderManager(t *testing.T) {
	t.Run("CreateTranscoder", func(t *testing.T) {
		manager := GetDefaultManager()

		transcoder, err := manager.CreateTranscoder(
			transfer.ExplicitVRLittleEndian,
			transfer.ImplicitVRLittleEndian,
		)

		if err != nil {
			t.Fatalf("CreateTranscoder() error = %v", err)
		}

		if transcoder == nil {
			t.Fatal("CreateTranscoder() returned nil")
		}
	})

	t.Run("CanTranscode", func(t *testing.T) {
		manager := GetDefaultManager()

		// Uncompressed to uncompressed should always be supported
		if !manager.CanTranscode(transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian) {
			t.Error("CanTranscode() returned false for uncompressed to uncompressed")
		}

		// RLE should be supported (registered in global registry)
		if !manager.CanTranscode(transfer.RLELossless, transfer.ExplicitVRLittleEndian) {
			t.Error("CanTranscode() returned false for RLE to uncompressed")
		}
	})
}

func TestGetGlobalRegistry(t *testing.T) {
	registry1 := GetGlobalRegistry()
	registry2 := GetGlobalRegistry()

	if registry1 != registry2 {
		t.Error("GetGlobalRegistry() returned different instances")
	}

	// Check that built-in codecs are registered
	if !registry1.HasCodec(transfer.ExplicitVRLittleEndian) {
		t.Error("Global registry does not have ExplicitVRLittleEndian codec")
	}

	if !registry1.HasCodec(transfer.RLELossless) {
		t.Error("Global registry does not have RLE codec")
	}
}

// TestTranscoder_VRSelection tests that the transcoder correctly uses OB
// for all encapsulated/compressed formats according to DICOM Part 5 Section 8.2.
//
// DICOM Standard: "If sent in an Encapsulated Format (i.e., other than the Native Format)
// the Value Representation OB is used."
//
// This means ALL compressed/encapsulated pixel data must use OB, regardless of BitsAllocated.
func TestTranscoder_VRSelection(t *testing.T) {
	tests := []struct {
		name           string
		bitsAllocated  uint16
		expectedVRType string // Should always be "OB" for encapsulated formats
	}{
		{
			name:           "8-bit should use OB",
			bitsAllocated:  8,
			expectedVRType: "OB",
		},
		{
			name:           "16-bit should use OB (not OW for encapsulated)",
			bitsAllocated:  16,
			expectedVRType: "OB", // Changed from "OW" to comply with DICOM standard
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create dataset with uncompressed pixel data
			ds := dataset.New()

			// Add image attributes
			_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{10}))
			_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{10}))
			_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{tt.bitsAllocated}))
			_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{tt.bitsAllocated}))
			_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{tt.bitsAllocated - 1}))
			_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
			_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
			_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))

			// Create pixel data
			bytesPerPixel := (tt.bitsAllocated + 7) / 8
			pixelDataSize := int(10 * 10 * bytesPerPixel)
			pixelData := make([]byte, pixelDataSize)
			for i := range pixelData {
				pixelData[i] = byte(i % 256)
			}

			if tt.bitsAllocated <= 8 {
				_ = ds.Add(element.NewOtherByte(tag.PixelData, pixelData))
			} else {
				_ = ds.Add(element.NewOtherWord(tag.PixelData, pixelData))
			}

			// Create transcoder to encode to RLE
			transcoder := NewTranscoder(transfer.ExplicitVRLittleEndian, transfer.RLELossless)

			// Encode
			encodedDS, err := transcoder.Transcode(ds)
			if err != nil {
				t.Fatalf("Transcode error = %v", err)
			}

			// Check pixel data VR type
			pixelDataElem, exists := encodedDS.Get(tag.PixelData)
			if !exists {
				t.Fatal("PixelData not found in encoded dataset")
			}

			var actualVR string
			switch pixelDataElem.(type) {
			case *element.OtherByteFragment:
				actualVR = "OB"
			case *element.OtherWordFragment:
				actualVR = "OW"
			default:
				t.Fatalf("Unexpected pixel data element type: %T", pixelDataElem)
			}

			if actualVR != tt.expectedVRType {
				t.Errorf("VR type = %s, want %s (BitsAllocated=%d)",
					actualVR, tt.expectedVRType, tt.bitsAllocated)
			}
		})
	}
}
