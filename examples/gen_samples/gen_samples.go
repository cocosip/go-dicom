// Package main generates sample DICOM files for manual verification.
package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/cocosip/go-dicom/examples/internal/examplepath"
	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	dicomwriter "github.com/cocosip/go-dicom/pkg/dicom/writer"
)

func mustAddValue(ds *dataset.Dataset, t *tag.Tag, value any) {
	if err := ds.AddValue(t, value); err != nil {
		log.Fatalf("add value failed: %v", err)
	}
}

func buildDataset(instanceUID, seriesUID, studyUID string, rows, cols uint16, frames int, pixels []byte) *dataset.Dataset {
	ds := dataset.New()
	mustAddValue(ds, tag.SOPClassUID, "1.2.840.10008.5.1.4.1.1.7") // Secondary Capture Image Storage
	mustAddValue(ds, tag.SOPInstanceUID, instanceUID)
	mustAddValue(ds, tag.StudyInstanceUID, studyUID)
	mustAddValue(ds, tag.SeriesInstanceUID, seriesUID)
	mustAddValue(ds, tag.Modality, "OT")
	mustAddValue(ds, tag.Rows, rows)
	mustAddValue(ds, tag.Columns, cols)
	mustAddValue(ds, tag.SamplesPerPixel, uint16(1))
	mustAddValue(ds, tag.PhotometricInterpretation, "MONOCHROME2")
	mustAddValue(ds, tag.BitsAllocated, uint16(8))
	mustAddValue(ds, tag.BitsStored, uint16(8))
	mustAddValue(ds, tag.HighBit, uint16(7))
	mustAddValue(ds, tag.PixelRepresentation, uint16(0))
	if frames > 1 {
		mustAddValue(ds, tag.NumberOfFrames, fmt.Sprintf("%d", frames))
	}
	if err := ds.AddValueWithVR(tag.PixelData, vr.OB, pixels); err != nil {
		log.Fatalf("add pixel data: %v", err)
	}
	return ds
}

func writeFile(path string, ds *dataset.Dataset) {
	if err := dicomwriter.WriteFile(path, ds, dicomwriter.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		log.Fatalf("write %s failed: %v", path, err)
	}
}

func main() {
	outputDir := flag.String("output-dir", ".", "Directory to write generated DICOM samples")
	flag.Parse()

	if err := examplepath.PrepareOutputDir(*outputDir); err != nil {
		log.Fatal(err)
	}

	// Single-frame 2x2 image: pixel values 0,64,128,255
	singlePixels := []byte{0x00, 0x40, 0x80, 0xFF}
	single := buildDataset("1.2.826.0.1.3680043.10.1142.1001.1", "1.2.826.0.1.3680043.10.1142.1001", "1.2.826.0.1.3680043.10.1142.1000", 2, 2, 1, singlePixels)
	writeFile(filepath.Join(*outputDir, "sample-single-frame.dcm"), single)

	// Multi-frame 2 frames, each 2x2: frame0 values 0..3, frame1 values 128..131
	multiPixels := []byte{
		0x00, 0x20, 0x40, 0x60, // frame 1
		0x80, 0x90, 0xA0, 0xB0, // frame 2
	}
	multi := buildDataset("1.2.826.0.1.3680043.10.1142.2001.1", "1.2.826.0.1.3680043.10.1142.2001", "1.2.826.0.1.3680043.10.1142.2000", 2, 2, 2, multiPixels)
	writeFile(filepath.Join(*outputDir, "sample-multiframe.dcm"), multi)
}
