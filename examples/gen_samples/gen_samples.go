// Package main generates sample DICOM files for manual verification.
package main

import (
	"fmt"
	"log"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	dicomwriter "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func mustAdd(ds *dataset.Dataset, elem element.Element) {
	if err := ds.Add(elem); err != nil {
		log.Fatalf("add element failed: %v", err)
	}
}

func buildDataset(instanceUID, seriesUID, studyUID string, rows, cols uint16, frames int, pixels []byte) *dataset.Dataset {
	ds := dataset.New()
	mustAdd(ds, element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.7"})) // Secondary Capture Image Storage
	mustAdd(ds, element.NewString(tag.SOPInstanceUID, vr.UI, []string{instanceUID}))
	mustAdd(ds, element.NewString(tag.StudyInstanceUID, vr.UI, []string{studyUID}))
	mustAdd(ds, element.NewString(tag.SeriesInstanceUID, vr.UI, []string{seriesUID}))
	mustAdd(ds, element.NewString(tag.Modality, vr.CS, []string{"OT"}))
	mustAdd(ds, element.NewUnsignedShort(tag.Rows, []uint16{rows}))
	mustAdd(ds, element.NewUnsignedShort(tag.Columns, []uint16{cols}))
	mustAdd(ds, element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}))
	mustAdd(ds, element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}))
	mustAdd(ds, element.NewUnsignedShort(tag.BitsAllocated, []uint16{8}))
	mustAdd(ds, element.NewUnsignedShort(tag.BitsStored, []uint16{8}))
	mustAdd(ds, element.NewUnsignedShort(tag.HighBit, []uint16{7}))
	mustAdd(ds, element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}))
	if frames > 1 {
		mustAdd(ds, element.NewString(tag.NumberOfFrames, vr.IS, []string{fmt.Sprintf("%d", frames)}))
	}
	mustAdd(ds, element.NewOtherByte(tag.PixelData, pixels))
	return ds
}

func writeFile(path string, ds *dataset.Dataset) {
	if err := dicomwriter.WriteFile(path, ds, dicomwriter.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		log.Fatalf("write %s failed: %v", path, err)
	}
}

func main() {
	// Single-frame 2x2 image: pixel values 0,64,128,255
	singlePixels := []byte{0x00, 0x40, 0x80, 0xFF}
	single := buildDataset("1.2.826.0.1.3680043.10.1142.1001.1", "1.2.826.0.1.3680043.10.1142.1001", "1.2.826.0.1.3680043.10.1142.1000", 2, 2, 1, singlePixels)
	writeFile("sample-single-frame.dcm", single)

	// Multi-frame 2 frames, each 2x2: frame0 values 0..3, frame1 values 128..131
	multiPixels := []byte{
		0x00, 0x20, 0x40, 0x60, // frame 1
		0x80, 0x90, 0xA0, 0xB0, // frame 2
	}
	multi := buildDataset("1.2.826.0.1.3680043.10.1142.2001.1", "1.2.826.0.1.3680043.10.1142.2001", "1.2.826.0.1.3680043.10.1142.2000", 2, 2, 2, multiPixels)
	writeFile("sample-multiframe.dcm", multi)
}
