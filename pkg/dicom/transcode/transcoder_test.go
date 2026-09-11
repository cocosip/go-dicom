// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package transcode

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/pixel"
	"github.com/cocosip/go-dicom/pkg/imaging/pixeldata"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func newTestTranscoder(
	t *testing.T,
	input, output *transfer.Syntax,
	dependenciesAndOptions ...any,
) *Transcoder {
	t.Helper()
	registry := codec.NewRegistry()
	var options []Option
	for _, value := range dependenciesAndOptions {
		switch value := value.(type) {
		case codec.Codec:
			if _, err := registry.Replace(value); err != nil {
				t.Fatalf("Registry.Replace() error = %v", err)
			}
		case Option:
			options = append(options, value)
		default:
			t.Fatalf("unsupported test dependency %T", value)
		}
	}
	manager, err := NewManager(registry)
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}
	transcoder, err := manager.NewTranscoder(input, output, options...)
	if err != nil {
		t.Fatalf("Manager.NewTranscoder() error = %v", err)
	}
	return transcoder
}

func TestManagerNewTranscoder(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		transcoder := newTestTranscoder(t,
			transfer.ExplicitVRLittleEndian,
			transfer.ImplicitVRLittleEndian,
		)

		if transcoder == nil {
			t.Fatal("Manager.NewTranscoder returned nil")
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
	}, 16, pixeldata.PixelDataStandard)
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

	transcoder := newTestTranscoder(t,
		transfer.ExplicitVRLittleEndian,
		transfer.ImplicitVRLittleEndian,
	)

	result, err := transcoder.Transcode(context.Background(), ds)
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

	transcoder := newTestTranscoder(t,
		transfer.ExplicitVRLittleEndian,
		transfer.ExplicitVRBigEndian,
	)
	result, err := transcoder.Transcode(context.Background(), ds)
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
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherWord(tag.PixelData, []byte{0x12, 0x34}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", elem.Tag(), err)
		}
	}

	transcoder := newTestTranscoder(t,
		transfer.ExplicitVRBigEndian,
		transfer.JPEG2000Lossless,
		echoDecodeCodec{},
	)
	result, err := transcoder.Transcode(context.Background(), ds)
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

		transcoder := newTestTranscoder(t,
			transfer.ExplicitVRLittleEndian,
			transfer.JPEG2000Lossless,
			mutatingCodec{},
		)
		if _, err := transcoder.Transcode(context.Background(), ds); err != nil {
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

		transcoder := newTestTranscoder(t,
			transfer.JPEG2000Lossless,
			transfer.ExplicitVRLittleEndian,
			mutatingCodec{},
		)
		if _, err := transcoder.Transcode(context.Background(), ds); err != nil {
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
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
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

func (mutatingCodec) DefaultParameters() codec.Parameters { return codec.NoParameters{} }

func (mutatingCodec) Encode(ctx context.Context, oldPixelData codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	return mutateAndCopyFrame(ctx, oldPixelData, newPixelData)
}

func (mutatingCodec) Decode(ctx context.Context, oldPixelData codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	return mutateAndCopyFrame(ctx, oldPixelData, newPixelData)
}

func mutateAndCopyFrame(ctx context.Context, oldPixelData codec.FrameSource, newPixelData codec.FrameSink) error {
	frame, err := oldPixelData.Frame(ctx, 0)
	if err != nil {
		return err
	}
	frame[0] = 0
	return newPixelData.AddFrame(ctx, frame)
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
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))

	// Create simple pixel data (512x512 = 262144 bytes)
	pixelData := make([]byte, 512*512)
	for i := range pixelData {
		pixelData[i] = byte(i % 256)
	}
	_ = ds.Add(element.NewOtherByte(tag.PixelData, pixelData))

	transcoder := newTestTranscoder(t,
		transfer.ExplicitVRLittleEndian,
		transfer.ImplicitVRLittleEndian,
	)

	result, err := transcoder.Transcode(context.Background(), ds)
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

func TestTranscoderNativePixelDataByteOrderUsesVR(t *testing.T) {
	tests := []struct {
		name       string
		input      *transfer.Syntax
		output     *transfer.Syntax
		pixelData  func() element.Element
		want       []byte
		wantOrder  binary.ByteOrder
		wantWordVR bool
	}{
		{
			name:   "8-bit OW little to big swaps words",
			input:  transfer.ExplicitVRLittleEndian,
			output: transfer.ExplicitVRBigEndian,
			pixelData: func() element.Element {
				return element.NewOtherWord(tag.PixelData, []byte{0x01, 0x02, 0x03, 0x04})
			},
			want:       []byte{0x02, 0x01, 0x04, 0x03},
			wantOrder:  binary.BigEndian,
			wantWordVR: true,
		},
		{
			name:   "8-bit OW big to little swaps words",
			input:  transfer.ExplicitVRBigEndian,
			output: transfer.ExplicitVRLittleEndian,
			pixelData: func() element.Element {
				e := element.NewOtherWord(tag.PixelData, []byte{0x01, 0x02, 0x03, 0x04})
				element.SetByteOrder(e, binary.BigEndian)
				return e
			},
			want:       []byte{0x02, 0x01, 0x04, 0x03},
			wantOrder:  binary.LittleEndian,
			wantWordVR: true,
		},
		{
			name:   "8-bit OB little to big remains byte-identical",
			input:  transfer.ExplicitVRLittleEndian,
			output: transfer.ExplicitVRBigEndian,
			pixelData: func() element.Element {
				return element.NewOtherByte(tag.PixelData, []byte{0x01, 0x02, 0x03, 0x04})
			},
			want: []byte{0x01, 0x02, 0x03, 0x04},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.NewWithTransferSyntax(tt.input)
			for _, elem := range []element.Element{
				element.NewUnsignedShort(tag.Rows, []uint16{1}),
				element.NewUnsignedShort(tag.Columns, []uint16{4}),
				element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
				element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
				element.NewUnsignedShort(tag.HighBit, []uint16{7}),
				element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
				element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
				element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
				tt.pixelData(),
			} {
				if err := ds.Add(elem); err != nil {
					t.Fatalf("Dataset.Add(%s) error = %v", elem.Tag(), err)
				}
			}

			gotDS, err := newTestTranscoder(t, tt.input, tt.output).Transcode(context.Background(), ds)
			if err != nil {
				t.Fatalf("Transcode() error = %v", err)
			}
			gotElem, ok := gotDS.Get(tag.PixelData)
			if !ok {
				t.Fatal("transcoded dataset has no Pixel Data")
			}
			var got []byte
			switch value := gotElem.(type) {
			case *element.OtherByte:
				got = value.GetData()
				if tt.wantWordVR {
					t.Fatalf("transcoded Pixel Data = %T, want OW", gotElem)
				}
			case *element.OtherWord:
				got = value.GetData()
				if !tt.wantWordVR {
					t.Fatalf("transcoded Pixel Data = %T, want OB", gotElem)
				}
			default:
				t.Fatalf("transcoded Pixel Data = %T, want native OB or OW", gotElem)
			}
			if !bytes.Equal(got, tt.want) {
				t.Fatalf("Pixel Data = % x, want % x", got, tt.want)
			}
			gotOrder, known := element.NumericByteOrder(gotElem)
			if tt.wantOrder == nil {
				if known {
					t.Fatalf("endian-neutral Pixel Data unexpectedly has byte order %T", gotOrder)
				}
			} else if !known || gotOrder.Uint16([]byte{0x01, 0x02}) != tt.wantOrder.Uint16([]byte{0x01, 0x02}) {
				t.Fatalf("Pixel Data byte order = %T, want %T", gotOrder, tt.wantOrder)
			}
		})
	}
}

func TestTranscoderNativeThirtyTwoBitOWSwapsSixteenBitWords(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{32}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{32}),
		element.NewUnsignedShort(tag.HighBit, []uint16{31}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherWord(tag.PixelData, []byte{0x11, 0x22, 0x33, 0x44}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", elem.Tag(), err)
		}
	}

	gotDS, err := newTestTranscoder(
		t,
		transfer.ExplicitVRLittleEndian,
		transfer.ExplicitVRBigEndian,
	).Transcode(context.Background(), ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}
	gotElem, ok := gotDS.Get(tag.PixelData)
	if !ok {
		t.Fatal("transcoded dataset has no Pixel Data")
	}
	word, ok := gotElem.(*element.OtherWord)
	if !ok {
		t.Fatalf("transcoded Pixel Data = %T, want OW", gotElem)
	}
	want := []byte{0x22, 0x11, 0x44, 0x33}
	if got := word.GetData(); !bytes.Equal(got, want) {
		t.Fatalf("Pixel Data = % x, want % x", got, want)
	}
}

func TestTranscoderDefaultNormalizesNativeOBAboveEightBits(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x34, 0x12}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	result, err := newTestTranscoder(
		t,
		transfer.ExplicitVRLittleEndian,
		transfer.ExplicitVRLittleEndian,
	).Transcode(context.Background(), ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}
	if got, _ := result.Get(tag.PixelData); got == nil {
		t.Fatal("transcoded Dataset has no Pixel Data")
	} else if _, ok := got.(*element.OtherWord); !ok {
		t.Fatalf("transcoded Pixel Data = %T, want native OW", got)
	}
}

func TestTranscoderStandardReadRejectsNativeOBAboveEightBits(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x34, 0x12}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	_, err := newTestTranscoder(
		t,
		transfer.ExplicitVRLittleEndian,
		transfer.ExplicitVRLittleEndian,
		WithPixelDataReadMode(pixeldata.PixelDataStandard),
	).Transcode(context.Background(), ds)
	if err == nil || !strings.Contains(err.Error(), "native Pixel Data uses OB") {
		t.Fatalf("Transcode() error = %v, want native OB rejection", err)
	}
}

func TestTranscoderStandardReadRejectsImplicitNativeOB(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ImplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x7f, 0x00}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	transcoder := newTestTranscoder(
		t,
		transfer.ImplicitVRLittleEndian,
		transfer.ExplicitVRLittleEndian,
		WithPixelDataReadMode(pixeldata.PixelDataStandard),
	)
	if _, err := transcoder.Transcode(context.Background(), ds); err == nil {
		t.Fatal("Transcode() error = nil, want implicit native OB rejection")
	}
}

func TestTranscoderCompatibleNativeOBCrossEndianRecoversAsOW(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x34, 0x12}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	gotDS, err := newTestTranscoder(
		t,
		transfer.ExplicitVRLittleEndian,
		transfer.ExplicitVRBigEndian,
	).Transcode(context.Background(), ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}
	gotElem, ok := gotDS.Get(tag.PixelData)
	if !ok {
		t.Fatal("transcoded dataset has no Pixel Data")
	}
	word, ok := gotElem.(*element.OtherWord)
	if !ok {
		t.Fatalf("transcoded Pixel Data = %T, want native OW", gotElem)
	}
	if got := word.GetData(); !bytes.Equal(got, []byte{0x12, 0x34}) {
		t.Fatalf("Pixel Data = % x, want big-endian OW bytes 12 34", got)
	}
}

func TestTranscoderDecodeFrameStandardReadRejectsNativeOBAboveEightBits(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
		element.NewOtherByte(tag.PixelData, []byte{0x34, 0x12}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	transcoder := newTestTranscoder(
		t,
		transfer.ExplicitVRLittleEndian,
		transfer.ExplicitVRLittleEndian,
		WithPixelDataReadMode(pixeldata.PixelDataStandard),
	)
	if _, err := transcoder.DecodeFrame(context.Background(), ds, 0); err == nil {
		t.Fatal("DecodeFrame() error = nil, want native OB rejection")
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
		_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))

		pixelData := make([]byte, 10*10)
		for i := range pixelData {
			pixelData[i] = byte(i)
		}
		_ = ds.Add(element.NewOtherByte(tag.PixelData, pixelData))

		transcoder := newTestTranscoder(t,
			transfer.ExplicitVRLittleEndian,
			transfer.ExplicitVRLittleEndian,
		)

		frame, err := transcoder.DecodeFrame(context.Background(), ds, 0)
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
		_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))

		pixelData := make([]byte, 10*10)
		_ = ds.Add(element.NewOtherByte(tag.PixelData, pixelData))

		transcoder := newTestTranscoder(t,
			transfer.ExplicitVRLittleEndian,
			transfer.ExplicitVRLittleEndian,
		)

		_, err := transcoder.DecodeFrame(context.Background(), ds, 5) // Frame 5 doesn't exist
		if err == nil {
			t.Error("DecodeFrame() should return error for invalid frame index")
		}
	})
}

func TestTranscoder_DecodeFrameUsesBOTFrameBoundaries(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{2}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))
	_ = ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"}))

	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.SetOffsetTable([]uint32{0, 20})
	obf.AddFragment(buffer.NewMemory([]byte("AA")))
	obf.AddFragment(buffer.NewMemory([]byte("BB")))
	obf.AddFragment(buffer.NewMemory([]byte("CC")))
	_ = ds.Add(obf)

	transcoder := newTestTranscoder(t,
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		echoDecodeCodec{},
	)

	frame, err := transcoder.DecodeFrame(context.Background(), ds, 1)
	if err != nil {
		t.Fatalf("DecodeFrame() error = %v", err)
	}
	if !bytes.Equal(frame, []byte("CC")) {
		t.Fatalf("DecodeFrame() = %q, want %q", frame, []byte("CC"))
	}
}

func TestTranscoderDecodeFrameUsesExtendedOffsetTable(t *testing.T) {
	ds := newCodecTestDataset(t, transfer.JPEG2000Lossless)
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Columns, []uint16{4}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{8}),
		element.NewUnsignedShort(tag.HighBit, []uint16{7}),
	} {
		if err := ds.AddOrUpdate(elem); err != nil {
			t.Fatal(err)
		}
	}
	if err := ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"})); err != nil {
		t.Fatal(err)
	}
	offsets := make([]byte, 16)
	lengths := make([]byte, 16)
	binary.LittleEndian.PutUint64(offsets, 0)
	binary.LittleEndian.PutUint64(offsets[8:], 20)
	binary.LittleEndian.PutUint64(lengths, 20)
	binary.LittleEndian.PutUint64(lengths[8:], 20)
	if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTable, offsets)); err != nil {
		t.Fatal(err)
	}
	if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTableLengths, lengths)); err != nil {
		t.Fatal(err)
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	for _, data := range [][]byte{[]byte("AA"), []byte("BB"), []byte("CC"), []byte("DD")} {
		fragments.AddFragment(buffer.NewMemory(data))
	}
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	transcoder := newTestTranscoder(
		t,
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		echoDecodeCodec{},
	)
	got, err := transcoder.DecodeFrame(context.Background(), ds, 1)
	if err != nil {
		t.Fatalf("DecodeFrame() error = %v", err)
	}
	if !bytes.Equal(got, []byte("CCDD")) {
		t.Fatalf("DecodeFrame() = %q, want %q", got, []byte("CCDD"))
	}
}

func TestTranscoderRemovesStaleExtendedOffsetTables(t *testing.T) {
	addTables := func(t *testing.T, ds *dataset.Dataset, offsets, lengths []uint64) {
		t.Helper()
		encode := func(values []uint64) []byte {
			data := make([]byte, len(values)*8)
			for i, value := range values {
				binary.LittleEndian.PutUint64(data[i*8:], value)
			}
			return data
		}
		if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTable, encode(offsets))); err != nil {
			t.Fatal(err)
		}
		if err := ds.Add(element.NewOtherVeryLong(tag.ExtendedOffsetTableLengths, encode(lengths))); err != nil {
			t.Fatal(err)
		}
	}
	assertTablesRemoved := func(t *testing.T, ds *dataset.Dataset) {
		t.Helper()
		for _, tableTag := range []*tag.Tag{tag.ExtendedOffsetTable, tag.ExtendedOffsetTableLengths} {
			if ds.Contains(tableTag) {
				t.Fatalf("transcoded dataset retained %s", tableTag)
			}
		}
	}

	t.Run("native to native", func(t *testing.T) {
		ds := newCodecTestDataset(t, transfer.ExplicitVRLittleEndian)
		if err := ds.Add(element.NewOtherWord(tag.PixelData, []byte{0x34, 0x12})); err != nil {
			t.Fatal(err)
		}
		addTables(t, ds, []uint64{0}, []uint64{10})

		got, err := newTestTranscoder(
			t,
			transfer.ExplicitVRLittleEndian,
			transfer.ExplicitVRLittleEndian,
		).Transcode(context.Background(), ds)
		if err != nil {
			t.Fatalf("Transcode() error = %v", err)
		}
		assertTablesRemoved(t, got)
	})

	t.Run("native to encapsulated", func(t *testing.T) {
		ds := newCodecTestDataset(t, transfer.ExplicitVRLittleEndian)
		if err := ds.Add(element.NewOtherWord(tag.PixelData, []byte{0x34, 0x12})); err != nil {
			t.Fatal(err)
		}
		addTables(t, ds, []uint64{0}, []uint64{10})

		got, err := newTestTranscoder(
			t,
			transfer.ExplicitVRLittleEndian,
			transfer.JPEG2000Lossless,
			echoDecodeCodec{},
		).Transcode(context.Background(), ds)
		if err != nil {
			t.Fatalf("Transcode() error = %v", err)
		}
		assertTablesRemoved(t, got)
	})

	t.Run("encapsulated to native", func(t *testing.T) {
		ds := newCodecTestDataset(t, transfer.JPEG2000Lossless)
		addTables(t, ds, []uint64{0}, []uint64{10})
		fragments := element.NewOtherByteFragment(tag.PixelData)
		fragments.AddFragment(buffer.NewMemory([]byte{0x34, 0x12}))
		if err := ds.Add(fragments); err != nil {
			t.Fatal(err)
		}

		got, err := newTestTranscoder(
			t,
			transfer.JPEG2000Lossless,
			transfer.ExplicitVRLittleEndian,
			echoDecodeCodec{},
		).Transcode(context.Background(), ds)
		if err != nil {
			t.Fatalf("Transcode() error = %v", err)
		}
		assertTablesRemoved(t, got)
	})
}

func TestTranscoderDecodeRejectsAmbiguousEmptyOffsetTable(t *testing.T) {
	ds := newCodecTestDataset(t, transfer.JPEG2000Lossless)
	if err := ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"})); err != nil {
		t.Fatal(err)
	}
	fragments := element.NewOtherByteFragment(tag.PixelData)
	for _, data := range [][]byte{[]byte("AA"), []byte("BB"), []byte("CC"), []byte("DD")} {
		fragments.AddFragment(buffer.NewMemory(data))
	}
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	transcoder := newTestTranscoder(
		t,
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		echoDecodeCodec{},
	)
	if _, err := transcoder.decode(context.Background(), ds, transfer.ExplicitVRLittleEndian); err == nil {
		t.Fatal("decode() error = nil, want indeterminate frame-boundaries error")
	}
}

func TestTranscoderDecodeParsesStringNumberOfFrames(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{2}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))
	_ = ds.Add(element.NewString(tag.NumberOfFrames, vr.IS, []string{"2"}))

	obf := element.NewOtherByteFragment(tag.PixelData)
	obf.SetOffsetTable([]uint32{0, 10})
	obf.AddFragment(buffer.NewMemory([]byte("AA")))
	obf.AddFragment(buffer.NewMemory([]byte("BB")))
	_ = ds.Add(obf)

	transcoder := newTestTranscoder(t,
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		echoDecodeCodec{},
	)

	result, err := transcoder.decode(context.Background(), ds, transfer.ExplicitVRLittleEndian)
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

func TestTranscoderDecodeAcceptsEncapsulatedOWPixelData(t *testing.T) {
	ds := dataset.New()
	for _, elem := range []element.Element{
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}
	owf := element.NewOtherWordFragment(tag.PixelData)
	owf.AddFragment(buffer.NewMemory([]byte{0x34, 0x12}))
	if err := ds.Add(owf); err != nil {
		t.Fatalf("add PixelData: %v", err)
	}

	transcoder := newTestTranscoder(t, transfer.JPEG2000Lossless, transfer.ExplicitVRLittleEndian, echoDecodeCodec{})
	decoded, err := transcoder.decode(context.Background(), ds, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("decode() error = %v", err)
	}
	elem, ok := decoded.Get(tag.PixelData)
	if !ok {
		t.Fatal("decoded PixelData not found")
	}
	word, ok := elem.(*element.OtherWord)
	if !ok {
		t.Fatalf("decoded PixelData = %T, want *element.OtherWord", elem)
	}
	if got := word.GetData(); !bytes.Equal(got, []byte{0x34, 0x12}) {
		t.Fatalf("decoded PixelData bytes = %x, want 3412", got)
	}
}

func TestTranscoderStandardReadModeRejectsEncapsulatedOWPixelData(t *testing.T) {
	ds := newCodecTestDataset(t, transfer.JPEG2000Lossless)
	fragments := element.NewOtherWordFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0x34, 0x12}))
	if err := ds.Add(fragments); err != nil {
		t.Fatal(err)
	}

	transcoder := newTestTranscoder(
		t,
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		echoDecodeCodec{},
		WithPixelDataReadMode(pixeldata.PixelDataStandard),
	)
	if _, err := transcoder.decode(context.Background(), ds, transfer.ExplicitVRLittleEndian); err == nil {
		t.Fatal("decode() error = nil, want encapsulated OW rejection")
	}
}

func TestTranscoderDecodeFrameLoadsOnlyRequestedBOTFragments(t *testing.T) {
	ds := dataset.New()
	_ = ds.Add(element.NewUnsignedShort(tag.Rows, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.Columns, []uint16{2}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	_ = ds.Add(element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	_ = ds.Add(element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	_ = ds.Add(element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))
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

	transcoder := newTestTranscoder(t,
		transfer.JPEG2000Lossless,
		transfer.ExplicitVRLittleEndian,
		echoDecodeCodec{},
	)

	frame, err := transcoder.DecodeFrame(context.Background(), ds, 1)
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

func (echoDecodeCodec) DefaultParameters() codec.Parameters {
	return codec.NoParameters{}
}

func (echoDecodeCodec) Encode(ctx context.Context, oldPixelData codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	for i := 0; i < oldPixelData.FrameCount(); i++ {
		frame, err := oldPixelData.Frame(ctx, i)
		if err != nil {
			return err
		}
		if err := newPixelData.AddFrame(ctx, frame); err != nil {
			return err
		}
	}
	return nil
}

func (echoDecodeCodec) Decode(ctx context.Context, oldPixelData codec.FrameSource, newPixelData codec.FrameSink, _ codec.Parameters) error {
	for i := 0; i < oldPixelData.FrameCount(); i++ {
		frame, err := oldPixelData.Frame(ctx, i)
		if err != nil {
			return err
		}
		if err := newPixelData.AddFrame(ctx, frame); err != nil {
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

func TestTranscoderManager(t *testing.T) {
	t.Run("NewTranscoder", func(t *testing.T) {
		manager, err := NewManager(codec.NewRegistry())
		if err != nil {
			t.Fatalf("NewManager() error = %v", err)
		}

		transcoder, err := manager.NewTranscoder(
			transfer.ExplicitVRLittleEndian,
			transfer.ImplicitVRLittleEndian,
		)

		if err != nil {
			t.Fatalf("NewTranscoder() error = %v", err)
		}

		if transcoder == nil {
			t.Fatal("NewTranscoder() returned nil")
		}
	})

	t.Run("CanTranscode", func(t *testing.T) {
		manager, err := NewManager(codec.NewRegistry())
		if err != nil {
			t.Fatalf("NewManager() error = %v", err)
		}

		// Uncompressed to uncompressed should always be supported
		if !manager.CanTranscode(transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian) {
			t.Error("CanTranscode() returned false for uncompressed to uncompressed")
		}

	})
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
			_ = ds.Add(element.NewString(tag.PhotometricInterpretation, vr.CS, []string{pixel.Monochrome2.Value}))

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
			transcoder := newTestTranscoder(t,
				transfer.ExplicitVRLittleEndian,
				transfer.JPEG2000Lossless,
				echoDecodeCodec{},
			)

			// Encode
			encodedDS, err := transcoder.Transcode(context.Background(), ds)
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

func TestTranscoderPixelDataVRRoundTripsThroughWriter(t *testing.T) {
	tests := []struct {
		name   string
		strict bool
		wantVR *vr.VR
	}{
		{name: "strict encapsulated Pixel Data uses OB", strict: true, wantVR: vr.OB},
		{name: "compatibility mode preserves 16-bit OW", strict: false, wantVR: vr.OW},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := newCodecTestDataset(t, transfer.ExplicitVRLittleEndian)
			for _, elem := range []element.Element{
				element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}),
				element.NewString(tag.SOPInstanceUID, vr.UI, []string{"2.25.146"}),
				element.NewOtherWord(tag.PixelData, []byte{0x34, 0x12}),
			} {
				if err := ds.Add(elem); err != nil {
					t.Fatalf("add %s: %v", elem.Tag(), err)
				}
			}

			encodedDS, err := newTestTranscoder(t,
				transfer.ExplicitVRLittleEndian,
				transfer.JPEG2000Lossless,
				echoDecodeCodec{},
				WithStrictDICOMVR(tt.strict),
			).Transcode(context.Background(), ds)
			if err != nil {
				t.Fatalf("Transcode() error = %v", err)
			}

			var encoded bytes.Buffer
			if err := writer.Write(&encoded, encodedDS); err != nil {
				t.Fatalf("writer.Write() error = %v", err)
			}
			parsed, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
			if err != nil {
				t.Fatalf("parser.Parse() error = %v", err)
			}
			pixelData, ok := parsed.Dataset.Get(tag.PixelData)
			if !ok {
				t.Fatal("round-tripped dataset has no Pixel Data")
			}
			if got := pixelData.ValueRepresentation(); got != tt.wantVR {
				t.Fatalf("round-tripped Pixel Data VR = %s, want %s", got, tt.wantVR)
			}
		})
	}
}

func TestTranscoderUsesWordVRForImplicitNativePixelData(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	for _, elem := range []element.Element{
		element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{"2.25.147"}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}),
		element.NewOtherByte(tag.PixelData, []byte{0x7f, 0x00}),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("add %s: %v", elem.Tag(), err)
		}
	}

	got, err := newTestTranscoder(t, transfer.ExplicitVRLittleEndian, transfer.ImplicitVRLittleEndian).Transcode(context.Background(), ds)
	if err != nil {
		t.Fatalf("Transcode() error = %v", err)
	}
	pixelData, ok := got.Get(tag.PixelData)
	if !ok {
		t.Fatal("transcoded dataset has no Pixel Data")
	}
	if _, ok := pixelData.(*element.OtherWord); !ok {
		t.Fatalf("transcoded Pixel Data = %T, want *element.OtherWord", pixelData)
	}

	var encoded bytes.Buffer
	if err := writer.Write(&encoded, got); err != nil {
		t.Fatalf("writer.Write() error = %v", err)
	}
	parsed, err := parser.Parse(bytes.NewReader(encoded.Bytes()))
	if err != nil {
		t.Fatalf("parser.Parse() error = %v", err)
	}
	roundTripped, ok := parsed.Dataset.Get(tag.PixelData)
	if !ok {
		t.Fatal("round-tripped dataset has no Pixel Data")
	}
	if _, ok := roundTripped.(*element.OtherWord); !ok {
		t.Fatalf("round-tripped Pixel Data = %T, want *element.OtherWord", roundTripped)
	}
}
