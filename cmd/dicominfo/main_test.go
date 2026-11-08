// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	original := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() error = %v", err)
	}
	os.Stdout = writer
	defer func() { os.Stdout = original }()

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("close stdout pipe error = %v", err)
	}
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("read stdout pipe error = %v", err)
	}
	if err := reader.Close(); err != nil {
		t.Fatalf("close stdout reader error = %v", err)
	}
	return string(output)
}

func TestPrintSeriesAndInstanceNumberReadISValues(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewIntegerString(tag.SeriesNumber, []string{"3"})); err != nil {
		t.Fatalf("Add(SeriesNumber) error = %v", err)
	}
	if err := ds.Add(element.NewIntegerString(tag.InstanceNumber, []string{"7"})); err != nil {
		t.Fatalf("Add(InstanceNumber) error = %v", err)
	}

	output := captureStdout(t, func() {
		printSeriesInfo(ds)
		printImageInfo(ds)
	})

	if !strings.Contains(output, "Series Number: 3") {
		t.Errorf("output %q does not contain Series Number: 3", output)
	}
	if !strings.Contains(output, "Instance Number: 7") {
		t.Errorf("output %q does not contain Instance Number: 7", output)
	}
}
