// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"bytes"
	"io"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
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

func TestBuildFragmentSequenceBOTUsesEncodedItemOffsets(t *testing.T) {
	seqElem, err := buildFragmentSequence([][]byte{
		[]byte("AA"),
		[]byte("BBB"),
		[]byte("C"),
	}, 16, true)
	if err != nil {
		t.Fatalf("buildFragmentSequence() error = %v", err)
	}

	obf, ok := seqElem.(*element.OtherByteFragment)
	if !ok {
		t.Fatalf("buildFragmentSequence() returned %T, want *element.OtherByteFragment", seqElem)
	}

	want := []uint32{0, 10, 22}
	got := obf.OffsetTable()
	if len(got) != len(want) {
		t.Fatalf("OffsetTable length = %d, want %d (%v)", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("OffsetTable[%d] = %d, want %d", i, got[i], want[i])
		}
	}
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

func TestTranscodeNoPixelDataSetsOutputTransferSyntax(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"NoPixel^Data"}))

	transcoder := NewTranscoder(
		transfer.ExplicitVRLittleEndian,
		transfer.ExplicitVRBigEndian,
	)
	result, err := transcoder.Transcode(ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}
	if result.InternalTransferSyntax() != transfer.ExplicitVRBigEndian {
		t.Fatalf("InternalTransferSyntax() = %v, want ExplicitVRBigEndian", result.InternalTransferSyntax())
	}
}

func TestTranscoderEncodeNormalizesBigEndianPixelsForCodec(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRBigEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}),
		element.NewOtherWord(tag.PixelData, []byte{0x12, 0x34}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", elem.Tag(), err)
		}
	}

	transcoder := NewTranscoder(
		transfer.ExplicitVRBigEndian,
		transfer.JPEG2000Lossless,
		WithOutputCodec(echoDecodeCodec{}),
	)
	result, err := transcoder.Transcode(ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}
	pixelData, ok := result.Get(tag.PixelData)
	if !ok {
		t.Fatal("transcoded dataset has no PixelData")
	}
	var fragments []buffer.ByteBuffer
	switch value := pixelData.(type) {
	case *element.OtherByteFragment:
		fragments = value.Fragments()
	case *element.OtherWordFragment:
		fragments = value.Fragments()
	default:
		t.Fatalf("transcoded PixelData = %T, want fragment sequence", pixelData)
	}
	if len(fragments) != 1 {
		t.Fatalf("fragment count = %d, want 1", len(fragments))
	}
	if got := fragments[0].Data(); !bytes.Equal(got, []byte{0x34, 0x12}) {
		t.Fatalf("codec input bytes = %v, want normalized Little Endian bytes [52 18]", got)
	}
}

func TestTranscoderCodecCannotMutateSourceDataset(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		ds := newCodecTestDataset(t, transfer.ExplicitVRLittleEndian)
		original := []byte{0x34, 0x12}
		if err := ds.Add(element.NewOtherWord(tag.PixelData, append([]byte(nil), original...))); err != nil {
			t.Fatalf("Dataset.Add(PixelData) error = %v", err)
		}

		transcoder := NewTranscoder(
			transfer.ExplicitVRLittleEndian,
			transfer.JPEG2000Lossless,
			WithOutputCodec(mutatingCodec{}),
		)
		if _, err := transcoder.Transcode(ds); err != nil {
			t.Fatalf("Transcode() error = %v", err)
		}
		pixelData, _ := ds.Get(tag.PixelData)
		if got := pixelData.(*element.OtherWord).GetData(); !bytes.Equal(got, original) {
			t.Fatalf("source PixelData = %v, want unchanged %v", got, original)
		}
	})

	t.Run("decode", func(t *testing.T) {
		ds := newCodecTestDataset(t, transfer.JPEG2000Lossless)
		if err := ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"})); err != nil {
			t.Fatalf("Dataset.Add(NumberOfFrames) error = %v", err)
		}
		original := []byte{0xff, 0x4f, 0xff, 0x51}
		fragments := element.NewOtherByteFragment(tag.PixelData)
		fragments.AddFragment(buffer.NewMemory(append([]byte(nil), original...)))
		fragments.AddFragment(buffer.NewMemory([]byte{0xff, 0x4f, 0xff, 0x52}))
		if err := ds.Add(fragments); err != nil {
			t.Fatalf("Dataset.Add(PixelData) error = %v", err)
		}

		transcoder := NewTranscoder(
			transfer.JPEG2000Lossless,
			transfer.ExplicitVRLittleEndian,
			WithInputCodec(mutatingCodec{}),
		)
		if _, err := transcoder.Transcode(ds); err != nil {
			t.Fatalf("Transcode() error = %v", err)
		}
		if got := fragments.Fragments()[0].Data(); !bytes.Equal(got, original) {
			t.Fatalf("source fragment = %v, want unchanged %v", got, original)
		}
	})
}

func newCodecTestDataset(t *testing.T, ts *transfer.Syntax) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(ts)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", elem.Tag(), err)
		}
	}
	return ds
}

type mutatingCodec struct{}

func (mutatingCodec) Name() string { return "mutating" }

func (mutatingCodec) TransferSyntax() *transfer.Syntax { return transfer.JPEG2000Lossless }

func (mutatingCodec) GetDefaultParameters() Parameters { return NewBaseParameters() }

func (mutatingCodec) Encode(oldPixelData, newPixelData imagetypes.PixelData, _ Parameters) error {
	return mutateAndCopyFrame(oldPixelData, newPixelData)
}

func (mutatingCodec) Decode(oldPixelData, newPixelData imagetypes.PixelData, _ Parameters) error {
	return mutateAndCopyFrame(oldPixelData, newPixelData)
}

func mutateAndCopyFrame(oldPixelData, newPixelData imagetypes.PixelData) error {
	frame, err := oldPixelData.GetFrame(0)
	if err != nil {
		return err
	}
	frame[0] = 0
	return newPixelData.AddFrame(frame)
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
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}))

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
		_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}))

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
		_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}))

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

func TestTranscoder_DecodeFrameUsesBOTFrameBoundaries(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}))
	_ = ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"}))

	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.SetOffsetTable([]uint32{0, 20})
	obf.AddFragment(buffer.NewMemory([]byte("AA")))
	obf.AddFragment(buffer.NewMemory([]byte("BB")))
	obf.AddFragment(buffer.NewMemory([]byte("CC")))
	_ = ds.Add(obf)

	transcoder := NewTranscoder(
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		WithInputCodec(echoDecodeCodec{}),
	)

	frame, err := transcoder.DecodeFrame(ds, 1)
	if err != nil {
		t.Fatalf("DecodeFrame() error = %v", err)
	}
	if !bytes.Equal(frame, []byte("CC")) {
		t.Fatalf("DecodeFrame() = %q, want %q", frame, []byte("CC"))
	}
}

func TestTranscoderDecodeParsesStringNumberOfFrames(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}))
	_ = ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"}))

	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.SetOffsetTable([]uint32{0, 10})
	obf.AddFragment(buffer.NewMemory([]byte("AA")))
	obf.AddFragment(buffer.NewMemory([]byte("BB")))
	_ = ds.Add(obf)

	transcoder := NewTranscoder(
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		WithInputCodec(echoDecodeCodec{}),
	)

	result, err := transcoder.decode(ds, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("decode() error = %v", err)
	}

	elem, exists := result.Get(tag.PixelData)
	if !exists {
		t.Fatal("decoded PixelData not found")
	}
	ob, ok := elem.(*element.OtherByte)
	if !ok {
		t.Fatalf("decoded PixelData = %T, want *element.OtherByte", elem)
	}
	if !bytes.Equal(ob.GetData(), []byte("AABB")) {
		t.Fatalf("decoded PixelData = %q, want %q", ob.GetData(), []byte("AABB"))
	}
}

func TestTranscoderDecodeFrameLoadsOnlyRequestedBOTFragments(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}))
	_ = ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"}))

	first := &countingBuffer{data: []byte("AA")}
	second := &countingBuffer{data: []byte("BB")}
	third := &countingBuffer{data: []byte("CC")}

	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.SetOffsetTable([]uint32{0, 20})
	obf.AddFragment(first)
	obf.AddFragment(second)
	obf.AddFragment(third)
	_ = ds.Add(obf)

	transcoder := NewTranscoder(
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		WithInputCodec(echoDecodeCodec{}),
	)

	frame, err := transcoder.DecodeFrame(ds, 1)
	if err != nil {
		t.Fatalf("DecodeFrame() error = %v", err)
	}
	if !bytes.Equal(frame, []byte("CC")) {
		t.Fatalf("DecodeFrame() = %q, want %q", frame, []byte("CC"))
	}
	if first.dataCalls != 0 || second.dataCalls != 0 {
		t.Fatalf("DecodeFrame() loaded unrequested fragments: first=%d second=%d", first.dataCalls, second.dataCalls)
	}
	if third.dataCalls == 0 {
		t.Fatal("DecodeFrame() did not load requested fragment")
	}
}

type echoDecodeCodec struct{}

func (echoDecodeCodec) Name() string {
	return "echo"
}

func (echoDecodeCodec) TransferSyntax() *transfer.Syntax {
	return transfer.JPEG2000Lossless
}

func (echoDecodeCodec) GetDefaultParameters() Parameters {
	return NewBaseParameters()
}

func (echoDecodeCodec) Encode(oldPixelData imagetypes.PixelData, newPixelData imagetypes.PixelData, _ Parameters) error {
	for i := 0; i < oldPixelData.FrameCount(); i++ {
		frame, err := oldPixelData.GetFrame(i)
		if err != nil {
			return err
		}
		if err := newPixelData.AddFrame(frame); err != nil {
			return err
		}
	}
	return nil
}

func (echoDecodeCodec) Decode(oldPixelData imagetypes.PixelData, newPixelData imagetypes.PixelData, _ Parameters) error {
	for i := 0; i < oldPixelData.FrameCount(); i++ {
		frame, err := oldPixelData.GetFrame(i)
		if err != nil {
			return err
		}
		if err := newPixelData.AddFrame(frame); err != nil {
			return err
		}
	}
	return nil
}

type countingBuffer struct {
	data      []byte
	dataCalls int
}

func (c *countingBuffer) IsMemory() bool {
	return false
}

func (c *countingBuffer) Size() uint32 {
	return uint32(len(c.data)) //nolint:gosec // test data is tiny
}

func (c *countingBuffer) Data() []byte {
	c.dataCalls++
	return c.data
}

func (c *countingBuffer) GetByteRange(offset, count uint32, output []byte) error {
	if offset > uint32(len(c.data)) || count > uint32(len(c.data))-offset {
		return io.ErrUnexpectedEOF
	}
	copy(output, c.data[offset:offset+count])
	return nil
}

func (c *countingBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(c.data)
	return int64(n), err
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

func TestFramesFromFragments_StripsTrailingPaddingWithBOT(t *testing.T) {
	fragments := []buffer.ByteBuffer{
		buffer.NewMemory([]byte{0xFF, 0xD9, 0x00}),
		buffer.NewMemory([]byte("B")),
	}

	frames, err := framesFromFragments(fragments, []uint32{0, 12}, 2)
	if err != nil {
		t.Fatalf("framesFromFragments() error = %v", err)
	}
	if !bytes.Equal(frames[0], []byte{0xFF, 0xD9}) {
		t.Fatalf("frames[0] = %v, want JPEG EOI without padding", frames[0])
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
			_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{photometricMonochrome2}))

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

			// Use a test codec to encode to an encapsulated transfer syntax.
			transcoder := NewTranscoder(
				transfer.ExplicitVRLittleEndian,
				transfer.JPEG2000Lossless,
				WithOutputCodec(echoDecodeCodec{}),
			)

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
