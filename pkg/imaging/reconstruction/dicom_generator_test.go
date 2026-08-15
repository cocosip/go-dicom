// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/imaging"
)

const testMonochrome1 = "MONOCHROME1"

func TestDicomGeneratorCreatesClassicDerivedSeriesAndRoundTrips(t *testing.T) {
	source := testEnhancedImageDataset(t, uid.EnhancedCTImageStorage.UID(), []float64{0, 5, 10}, []uint16{
		0, 10, 20, 30,
		50, 60, 70, 80,
		100, 110, 120, 130,
	})
	mustAddElement(t, source,
		element.NewString(tag.PatientID, vr.LO, []string{"P-IMG-003"}),
		element.NewString(tag.Modality, vr.CS, []string{"MR"}),
	)
	images, err := NewImageDataFromDataset(source)
	if err != nil {
		t.Fatalf("NewImageDataFromDataset() error = %v", err)
	}
	volume, err := NewVolumeData(images)
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}
	stack, err := NewStack(volume, StackTypeAxial, 1, 5)
	if err != nil {
		t.Fatalf("NewStack() error = %v", err)
	}
	uids := []string{"2.25.100", "2.25.101", "2.25.102", "2.25.103"}
	nextUID := 0
	generator, err := NewDicomGenerator(volume,
		WithGeneratorClock(func() time.Time { return time.Date(2026, 8, 15, 12, 34, 56, 0, time.UTC) }),
		WithGeneratorUIDFactory(func() string {
			value := uids[nextUID]
			nextUID++
			return value
		}),
	)
	if err != nil {
		t.Fatalf("NewDicomGenerator() error = %v", err)
	}

	outputs, err := generator.Generate(context.Background(), stack, "Axial MPR", CutOptions{Workers: 2})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if len(outputs) != 3 {
		t.Fatalf("output count = %d, want 3", len(outputs))
	}
	for index, output := range outputs {
		if got := output.TryGetString(tag.SOPClassUID); got != uid.CTImageStorage.UID() {
			t.Fatalf("output %d SOP Class UID = %q", index, got)
		}
		if got := output.TryGetString(tag.Modality); got != "CT" {
			t.Fatalf("output %d Modality = %q, want CT from output SOP Class", index, got)
		}
		if got := output.TryGetString(tag.SeriesInstanceUID); got != "2.25.100" {
			t.Fatalf("output %d Series Instance UID = %q", index, got)
		}
		if got := output.TryGetString(tag.SOPInstanceUID); got != uids[index+1] {
			t.Fatalf("output %d SOP Instance UID = %q", index, got)
		}
		if values, ok := output.GetStrings(tag.ImageType); !ok || !reflect.DeepEqual(values, []string{"DERIVED", "SECONDARY", "MPR"}) {
			t.Fatalf("output %d Image Type = %v/%v", index, values, ok)
		}
		if output.Contains(tag.NumberOfFrames) || output.Contains(tag.SharedFunctionalGroupsSequence) || output.Contains(tag.PerFrameFunctionalGroupsSequence) {
			t.Fatalf("output %d retained Enhanced multi-frame tags", index)
		}
		if output.InternalTransferSyntax() != transfer.ExplicitVRLittleEndian {
			t.Fatalf("output %d transfer syntax = %v", index, output.InternalTransferSyntax())
		}
		if output.TryGetString(tag.PatientID) != "P-IMG-003" {
			t.Fatalf("output %d lost common patient metadata", index)
		}
		sequence, err := output.GetSequence(tag.SourceImageSequence)
		if err != nil || sequence.Count() != 1 {
			t.Fatalf("output %d Source Image Sequence = %v, error %v", index, sequence, err)
		}
		referencedFramesElement, ok := sequence.GetItem(0).Get(tag.ReferencedFrameNumber)
		if !ok {
			t.Fatalf("output %d source reference omitted Referenced Frame Number", index)
		}
		referencedFrames, err := referencedFramesElement.(*element.IntegerString).GetInts()
		if err != nil || !reflect.DeepEqual(referencedFrames, []int{1, 2, 3}) {
			t.Fatalf("output %d Referenced Frame Number = %v, error %v", index, referencedFrames, err)
		}

		encoded := &bytes.Buffer{}
		if err := writer.Write(encoded, output, writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
			t.Fatalf("writer.Write(output %d) error = %v", index, err)
		}
		parsed, err := parser.Parse(encoded)
		if err != nil {
			t.Fatalf("parser.Parse(output %d) error = %v", index, err)
		}
		if parsed.TransferSyntax != transfer.ExplicitVRLittleEndian {
			t.Fatalf("parsed output %d transfer syntax = %v", index, parsed.TransferSyntax)
		}
		pixels, err := imaging.CreatePixelData(parsed.Dataset)
		if err != nil {
			t.Fatalf("CreatePixelData(parsed output %d) error = %v", index, err)
		}
		first, err := pixels.GetSample(0, 0, 0, 0)
		if err != nil {
			t.Fatalf("GetSample(parsed output %d) error = %v", index, err)
		}
		if first != int64(index*50) {
			t.Fatalf("parsed output %d first sample = %d, want %d", index, first, index*50)
		}
	}
}

func TestEncodeSlicePixelsUsesSignedPaddingAndRejectsOverflow(t *testing.T) {
	data, representation, padding, _, _, err := encodeSlicePixels(&Slice{
		Values: []float64{-10, 0, 20},
		Valid:  []bool{true, false, true},
	})
	if err != nil {
		t.Fatalf("encodeSlicePixels() error = %v", err)
	}
	if representation != 1 {
		t.Fatalf("pixel representation = %d, want signed", representation)
	}
	if _, ok := padding.(*element.SignedShort); !ok {
		t.Fatalf("padding element type = %T, want SignedShort", padding)
	}
	got := []int16{
		int16(binary.LittleEndian.Uint16(data[0:2])),
		int16(binary.LittleEndian.Uint16(data[2:4])),
		int16(binary.LittleEndian.Uint16(data[4:6])),
	}
	if !reflect.DeepEqual(got, []int16{-10, -32768, 20}) {
		t.Fatalf("encoded samples = %v", got)
	}
	if _, _, _, _, _, err := encodeSlicePixels(&Slice{Values: []float64{70000}, Valid: []bool{true}}); err == nil {
		t.Fatal("encodeSlicePixels() accepted an overflowing value")
	}
}

func TestDicomGeneratorMapsEnhancedMRToClassicMR(t *testing.T) {
	images := testImagesWithPixels(t, []float64{0, 1}, [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}})
	for _, image := range images {
		image.sourceSOPClassUID = uid.EnhancedMRImageStorage.UID()
		_ = image.dataset.AddOrUpdate(element.NewString(tag.SOPClassUID, vr.UI, []string{uid.EnhancedMRImageStorage.UID()}))
	}
	volume, err := NewVolumeData(images)
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}
	stack, _ := NewStack(volume, StackTypeAxial, 1, 1)
	generator, _ := NewDicomGenerator(volume, WithGeneratorUIDFactory(sequenceUIDFactory()))
	outputs, err := generator.Generate(context.Background(), stack, "MR MPR", CutOptions{})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if outputs[0].TryGetString(tag.SOPClassUID) != uid.MRImageStorage.UID() {
		t.Fatalf("SOP Class UID = %q, want classic MR", outputs[0].TryGetString(tag.SOPClassUID))
	}
}

func TestDicomGeneratorOutputsHaveIndependentSourceImageSequences(t *testing.T) {
	volume := testVolumeWithPixels(t, []float64{0, 1}, [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}})
	stack, err := NewStack(volume, StackTypeAxial, 1, 1)
	if err != nil {
		t.Fatalf("NewStack() error = %v", err)
	}
	generator, err := NewDicomGenerator(volume, WithGeneratorUIDFactory(sequenceUIDFactory()))
	if err != nil {
		t.Fatalf("NewDicomGenerator() error = %v", err)
	}
	outputs, err := generator.Generate(context.Background(), stack, "MPR", CutOptions{})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if len(outputs) < 2 {
		t.Fatalf("Generate() output count = %d, want at least 2", len(outputs))
	}

	first, err := outputs[0].GetSequence(tag.SourceImageSequence)
	if err != nil {
		t.Fatalf("first output Source Image Sequence error = %v", err)
	}
	first.GetItem(0).Remove(tag.ReferencedSOPInstanceUID)
	second, err := outputs[1].GetSequence(tag.SourceImageSequence)
	if err != nil {
		t.Fatalf("second output Source Image Sequence error = %v", err)
	}
	if second.GetItem(0).TryGetString(tag.ReferencedSOPInstanceUID) == "" {
		t.Fatal("mutating one output changed another output's Source Image Sequence")
	}
}

func TestDicomGeneratorPreservesMRMonochrome1Polarity(t *testing.T) {
	images := make([]*ImageData, 2)
	for index, position := range []float64{0, 1} {
		ds := testClassicImageDataset(t, uid.MRImageStorage.UID(), "1.2.3."+strconv.Itoa(20+index), position, []uint16{1, 2, 3, 4})
		mustAddElement(t, ds, element.NewString(tag.PhotometricInterpretation, vr.CS, []string{testMonochrome1}))
		image, err := NewImageData(ds, 0)
		if err != nil {
			t.Fatalf("NewImageData(MR MONOCHROME1) error = %v", err)
		}
		images[index] = image
	}
	volume, err := NewVolumeData(images)
	if err != nil {
		t.Fatalf("NewVolumeData() error = %v", err)
	}
	stack, err := NewStack(volume, StackTypeAxial, 1, 1)
	if err != nil {
		t.Fatalf("NewStack() error = %v", err)
	}
	generator, err := NewDicomGenerator(volume, WithGeneratorUIDFactory(sequenceUIDFactory()))
	if err != nil {
		t.Fatalf("NewDicomGenerator() error = %v", err)
	}
	outputs, err := generator.Generate(context.Background(), stack, "MR MPR", CutOptions{})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if got := outputs[0].TryGetString(tag.Modality); got != "MR" {
		t.Fatalf("Modality = %q, want MR", got)
	}
	if got := outputs[0].TryGetString(tag.PhotometricInterpretation); got != testMonochrome1 {
		t.Fatalf("Photometric Interpretation = %q, want MONOCHROME1", got)
	}
}

func TestDicomGeneratorStreamsAndStopsOnConsumerError(t *testing.T) {
	volume := testVolumeWithPixels(t, []float64{0, 1}, [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}})
	stack, _ := NewStack(volume, StackTypeAxial, 1, 1)
	generator, _ := NewDicomGenerator(volume, WithGeneratorUIDFactory(sequenceUIDFactory()))
	wantErr := errors.New("stop")
	count := 0
	err := generator.Stream(context.Background(), stack, "MPR", CutOptions{}, func(_ int, _ *dataset.Dataset) error {
		count++
		return wantErr
	})
	if !errors.Is(err, wantErr) || count != 1 {
		t.Fatalf("Stream() error/count = %v/%d, want stop/1", err, count)
	}
}

func sequenceUIDFactory() func() string {
	next := 1
	return func() string {
		value := "2.25." + strconv.Itoa(next)
		next++
		return value
	}
}
