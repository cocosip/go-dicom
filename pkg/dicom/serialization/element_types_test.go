// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package serialization

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// assertTypedTextElements verifies that deserialization produces the same
// concrete element types as the file parser and element.NewElementFromBuffer.
func assertTypedTextElements(t *testing.T, ds *dataset.Dataset) {
	t.Helper()

	tests := []struct {
		tag      *tag.Tag
		wantType string
		isType   func(element.Element) bool
	}{
		{tag.StudyDate, "*element.Date", func(e element.Element) bool { _, ok := e.(*element.Date); return ok }},
		{tag.StudyTime, "*element.Time", func(e element.Element) bool { _, ok := e.(*element.Time); return ok }},
		{tag.AcquisitionDateTime, "*element.DateTime", func(e element.Element) bool { _, ok := e.(*element.DateTime); return ok }},
		{tag.SliceThickness, "*element.DecimalString", func(e element.Element) bool { _, ok := e.(*element.DecimalString); return ok }},
		{tag.InstanceNumber, "*element.IntegerString", func(e element.Element) bool { _, ok := e.(*element.IntegerString); return ok }},
	}
	for _, tt := range tests {
		elem, ok := ds.Get(tt.tag)
		if !ok {
			t.Errorf("%s: element missing", tt.tag)
			continue
		}
		if !tt.isType(elem) {
			t.Errorf("%s: element type = %T, want %s", tt.tag, elem, tt.wantType)
		}
	}
}

func TestFromJSONUsesTypedDateElements(t *testing.T) {
	input := []byte(`{
		"00080020":{"vr":"DA","Value":["20240102"]},
		"00080030":{"vr":"TM","Value":["120000.5"]},
		"0008002A":{"vr":"DT","Value":["20240102120000"]},
		"00180050":{"vr":"DS","Value":[1.25]},
		"00200013":{"vr":"IS","Value":[42]}
	}`)

	ds, err := FromJSON(input)
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}
	assertTypedTextElements(t, ds)
}

func TestFromXMLUsesTypedDateElements(t *testing.T) {
	input := []byte(`<?xml version="1.0" encoding="utf-8"?>
<NativeDicomModel>
  <DicomAttribute tag="00080020" vr="DA"><Value number="1">20240102</Value></DicomAttribute>
  <DicomAttribute tag="00080030" vr="TM"><Value number="1">120000.5</Value></DicomAttribute>
  <DicomAttribute tag="0008002A" vr="DT"><Value number="1">20240102120000</Value></DicomAttribute>
  <DicomAttribute tag="00180050" vr="DS"><Value number="1">1.25</Value></DicomAttribute>
  <DicomAttribute tag="00200013" vr="IS"><Value number="1">42</Value></DicomAttribute>
</NativeDicomModel>`)

	ds, err := FromXML(input)
	if err != nil {
		t.Fatalf("FromXML() error = %v", err)
	}
	assertTypedTextElements(t, ds)
}
