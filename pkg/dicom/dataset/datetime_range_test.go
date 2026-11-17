// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package dataset_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// TestAddRangeMatchingDateTimeValue verifies PS3.5 Table 6.2-1 and PS3.4
// C.2.2.2.5: a DA/TM/DT query key may carry a range ("from-to"), so the value
// can be longer than the single-value length.
func TestAddRangeMatchingDateTimeValue(t *testing.T) {
	for _, test := range []struct {
		name  string
		tag   *tag.Tag
		vr    *vr.VR
		value string
	}{
		{name: "DA range", tag: tag.StudyDate, vr: vr.DA, value: "20240101-20241231"},
		{name: "TM range", tag: tag.StudyTime, vr: vr.TM, value: "080000.000000-180000.000000"},
		{name: "DT range", tag: tag.HangingProtocolCreationDateTime, vr: vr.DT, value: "20240101000000-20240201000000"},
	} {
		t.Run(test.name, func(t *testing.T) {
			ds := dataset.New()
			if err := ds.Add(element.NewString(test.tag, test.vr, []string{test.value})); err != nil {
				t.Fatalf("Add(%q) error = %v", test.value, err)
			}
		})
	}
}

// TestAddRejectsOverlongSingleDateTimeValue keeps the length check meaningful
// for the values a plain dataset (not a query key) may carry.
func TestAddRejectsOverlongSingleDateTimeValue(t *testing.T) {
	ds := dataset.New()
	value := "20240101000000.123456+0800X"
	if err := ds.Add(element.NewString(tag.HangingProtocolCreationDateTime, vr.DT, []string{value})); err == nil {
		t.Fatal("Add() accepted an overlong DT value")
	}
}
