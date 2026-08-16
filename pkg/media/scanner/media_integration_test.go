// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package scanner_test

import (
	"bytes"
	"context"
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/media"
	"github.com/cocosip/go-dicom/pkg/media/scanner"
)

func TestScannerResultCanPopulateMediaDirectory(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "IMAGE001")
	if err := writeExternalScannerDICOM(path); err != nil {
		t.Fatalf("write DICOM fixture: %v", err)
	}
	value, err := scanner.New()
	if err != nil {
		t.Fatalf("scanner.New() error = %v", err)
	}
	directory, err := media.NewDirectory()
	if err != nil {
		t.Fatalf("media.NewDirectory() error = %v", err)
	}

	summary, err := value.Scan(context.Background(), []string{root}, func(result scanner.Result) error {
		fileID, parseErr := media.ParseFileID(result.RelativePath)
		if parseErr != nil {
			return parseErr
		}
		_, addErr := directory.AddFile(result.File, fileID)
		return addErr
	})
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if summary != (scanner.Summary{Results: 1, DICOMFiles: 1}) {
		t.Fatalf("Summary = %#v", summary)
	}
	imageRecords := 0
	if err := directory.Walk(func(record *media.Record) error {
		if record.Type() == media.RecordImage {
			imageRecords++
		}
		return nil
	}); err != nil {
		t.Fatalf("Directory.Walk() error = %v", err)
	}
	if imageRecords != 1 {
		t.Fatalf("IMAGE records = %d, want 1", imageRecords)
	}
}

func writeExternalScannerDICOM(path string) error {
	var file bytes.Buffer
	file.Write(make([]byte, 128))
	file.WriteString("DICM")
	appendExternalShortElement(&file, 0x0002, 0x0002, "UI", "1.2.840.10008.5.1.4.1.1.2")
	appendExternalShortElement(&file, 0x0002, 0x0003, "UI", "1.2.826.0.1.3680043.10.543.1")
	appendExternalShortElement(&file, 0x0002, 0x0010, "UI", "1.2.840.10008.1.2.1")
	appendExternalShortElement(&file, 0x0010, 0x0010, "PN", "Scanner^Patient")
	appendExternalShortElement(&file, 0x0010, 0x0020, "LO", "SCANNER1")
	appendExternalShortElement(&file, 0x0020, 0x000D, "UI", "1.2.826.0.1.3680043.10.543.2")
	appendExternalShortElement(&file, 0x0020, 0x000E, "UI", "1.2.826.0.1.3680043.10.543.3")
	appendExternalShortElement(&file, 0x0008, 0x0018, "UI", "1.2.826.0.1.3680043.10.543.1")
	appendExternalLongElement(&file, 0x7FE0, 0x0010, "OB", bytes.Repeat([]byte{0x5A}, 1024))
	return os.WriteFile(path, file.Bytes(), 0o600)
}

func appendExternalShortElement(dst *bytes.Buffer, group, element uint16, valueRepresentation, value string) {
	padding := byte(' ')
	if valueRepresentation == "UI" {
		padding = 0
	}
	data := []byte(value)
	if len(data)%2 != 0 {
		data = append(data, padding)
	}
	var header [8]byte
	binary.LittleEndian.PutUint16(header[0:2], group)
	binary.LittleEndian.PutUint16(header[2:4], element)
	copy(header[4:6], valueRepresentation)
	binary.LittleEndian.PutUint16(header[6:8], uint16(len(data)))
	dst.Write(header[:])
	dst.Write(data)
}

func appendExternalLongElement(dst *bytes.Buffer, group, element uint16, valueRepresentation string, value []byte) {
	var header [12]byte
	binary.LittleEndian.PutUint16(header[0:2], group)
	binary.LittleEndian.PutUint16(header[2:4], element)
	copy(header[4:6], valueRepresentation)
	binary.LittleEndian.PutUint32(header[8:12], uint32(len(value)))
	dst.Write(header[:])
	dst.Write(value)
}
