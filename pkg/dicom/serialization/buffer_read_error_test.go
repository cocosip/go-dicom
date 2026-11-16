// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package serialization

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func failingLazy(t *testing.T, size uint32) buffer.ByteBuffer {
	t.Helper()
	buf, err := buffer.NewLazySizedWithError(size, func() ([]byte, error) {
		return nil, errors.New("boom")
	})
	if err != nil {
		t.Fatalf("NewLazySizedWithError() error = %v", err)
	}
	return buf
}

func TestToJSONAndXMLSurfaceDeferredBufferReadFailures(t *testing.T) {
	tests := []struct {
		name string
		buf  func(t *testing.T) buffer.ByteBuffer
	}{
		{name: "sized lazy", buf: func(t *testing.T) buffer.ByteBuffer { return failingLazy(t, 4) }},
		{
			name: "range over failing lazy",
			buf: func(t *testing.T) buffer.ByteBuffer {
				ranged, err := buffer.NewRange(failingLazy(t, 4), 0, 4)
				if err != nil {
					t.Fatalf("NewRange() error = %v", err)
				}
				return ranged
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
			if err := ds.Add(element.NewOtherByteFromBuffer(tag.PixelData, tc.buf(t))); err != nil {
				t.Fatalf("Add() error = %v", err)
			}
			if _, err := ToJSON(ds); err == nil {
				t.Fatal("ToJSON() error = nil, want deferred read failure")
			}
			if _, err := ToXML(ds); err == nil {
				t.Fatal("ToXML() error = nil, want deferred read failure")
			}
		})
	}
}

func TestToJSONAndXMLSurfaceFileReadFailures(t *testing.T) {
	path := filepath.Join(t.TempDir(), "value.bin")
	if err := os.WriteFile(path, []byte{1, 2, 3, 4}, 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	fileBuf, err := buffer.NewFileAt(path, 0, 4)
	if err != nil {
		t.Fatalf("NewFileAt() error = %v", err)
	}
	if err := os.Remove(path); err != nil {
		t.Fatalf("Remove() error = %v", err)
	}

	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	if err := ds.Add(element.NewOtherByteFromBuffer(tag.PixelData, fileBuf)); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if _, err := ToJSON(ds); err == nil {
		t.Fatal("ToJSON() error = nil, want file read failure")
	}
	if _, err := ToXML(ds); err == nil {
		t.Fatal("ToXML() error = nil, want file read failure")
	}
}

func TestToJSONAndXMLSurfaceLazyStringReadFailures(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	ds.SetAutoValidate(false) // the failing value is the point of this test
	elem, err := element.NewElementFromBuffer(tag.TextValue, vr.UT, failingLazy(t, 4), nil)
	if err != nil {
		t.Fatalf("NewElementFromBuffer() error = %v", err)
	}
	if err := ds.Add(elem); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if _, err := ToJSON(ds); err == nil {
		t.Fatal("ToJSON() error = nil, want deferred string read failure")
	}
	if _, err := ToXML(ds); err == nil {
		t.Fatal("ToXML() error = nil, want deferred string read failure")
	}
}

func TestToJSONAndXMLSurfaceLazyNumericReadFailures(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	elem, err := element.NewElementFromBuffer(tag.Rows, vr.US, failingLazy(t, 2), nil)
	if err != nil {
		t.Fatalf("NewElementFromBuffer() error = %v", err)
	}
	if err := ds.Add(elem); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if _, err := ToJSON(ds); err == nil {
		t.Fatal("ToJSON() error = nil, want deferred numeric read failure")
	}
	if _, err := ToXML(ds); err == nil {
		t.Fatal("ToXML() error = nil, want deferred numeric read failure")
	}
}

func TestToJSONAndXMLRejectFragmentSequencesWithoutPanicking(t *testing.T) {
	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	if err := ds.Add(element.NewOtherByteFragment(tag.PixelData)); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if _, err := ToJSON(ds); err == nil {
		t.Fatal("ToJSON() error = nil, want unsupported fragment sequence error")
	}
	if _, err := ToXML(ds); err == nil {
		t.Fatal("ToXML() error = nil, want unsupported fragment sequence error")
	}
}

func TestToJSONRejectsParsedFragmentSequenceWithoutPanicking(t *testing.T) {
	result, err := parser.ParseFile(filepath.Join("..", "..", "..", "test-data", "HasFragmentSequence.dcm"))
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}
	if _, err := ToJSON(result.Dataset); err == nil || !strings.Contains(err.Error(), "no value buffer") {
		t.Fatalf("ToJSON() error = %v, want unsupported fragment sequence error", err)
	}
}
