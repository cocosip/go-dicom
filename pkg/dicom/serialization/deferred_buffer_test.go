// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package serialization

import (
	"errors"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func datasetWithDeferredOB(t *testing.T, loader func() ([]byte, error)) *dataset.Dataset {
	t.Helper()
	lazy, err := buffer.NewLazySizedWithError(4, loader)
	if err != nil {
		t.Fatalf("NewLazySizedWithError() error = %v", err)
	}
	ds := dataset.New()
	if err := ds.Add(element.NewOtherByteFromBuffer(tag.PixelData, lazy)); err != nil {
		t.Fatalf("Add(PixelData) error = %v", err)
	}
	return ds
}

func TestSerializersReportDeferredBufferLoadFailure(t *testing.T) {
	ds := datasetWithDeferredOB(t, func() ([]byte, error) {
		return nil, errors.New("deferred read failed")
	})

	if _, err := ToJSON(ds); err == nil || !strings.Contains(err.Error(), "deferred read failed") {
		t.Fatalf("ToJSON() error = %v, want the deferred load failure", err)
	}
	if _, err := ToXML(ds); err == nil || !strings.Contains(err.Error(), "deferred read failed") {
		t.Fatalf("ToXML() error = %v, want the deferred load failure", err)
	}
}

func TestSerializersWriteDeferredBufferData(t *testing.T) {
	ds := datasetWithDeferredOB(t, func() ([]byte, error) {
		return []byte{1, 2, 3, 4}, nil
	})

	jsonData, err := ToJSON(ds)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}
	if !strings.Contains(string(jsonData), `"InlineBinary":"AQIDBA=="`) {
		t.Fatalf("ToJSON() = %s, want the deferred data as InlineBinary", jsonData)
	}

	xmlData, err := ToXML(ds)
	if err != nil {
		t.Fatalf("ToXML() error = %v", err)
	}
	if !strings.Contains(string(xmlData), "<InlineBinary>AQIDBA==</InlineBinary>") {
		t.Fatalf("ToXML() = %s, want the deferred data as InlineBinary", xmlData)
	}
}
