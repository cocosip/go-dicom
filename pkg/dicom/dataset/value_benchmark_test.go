// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dataset

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

const (
	benchmarkFloat64sName = "float64s"
	benchmarkStringName   = "string"
	benchmarkStringValue  = "Zhang^San"
	benchmarkUInt16Name   = "uint16"
)

var (
	benchmarkValueMatch bool
	benchmarkValueVR    *vr.VR
	benchmarkValueError error
)

func TestValueMatchesVRCommonSlicesDoNotAllocate(t *testing.T) {
	tests := []struct {
		name  string
		vr    *vr.VR
		value any
	}{
		{name: "strings", vr: vr.LO, value: []string{"one", "two"}},
		{name: "uint16s", vr: vr.US, value: []uint16{512, 1024}},
		{name: benchmarkFloat64sName, vr: vr.DS, value: []float64{0, 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocations := testing.AllocsPerRun(100, func() {
				benchmarkValueMatch = valueMatchesVR(tt.vr, tt.value)
			})
			if allocations != 0 {
				t.Fatalf("valueMatchesVR(%T) allocations = %.0f, want 0", tt.value, allocations)
			}
		})
	}
}

func BenchmarkValueMatchesVR(b *testing.B) {
	tests := []struct {
		name  string
		vr    *vr.VR
		value any
	}{
		{name: benchmarkStringName, vr: vr.PN, value: benchmarkStringValue},
		{name: "strings", vr: vr.LO, value: []string{"one", "two"}},
		{name: benchmarkUInt16Name, vr: vr.US, value: uint16(512)},
		{name: benchmarkFloat64sName, vr: vr.DS, value: []float64{0, 100}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				benchmarkValueMatch = valueMatchesVR(tt.vr, tt.value)
			}
		})
	}
}

func BenchmarkResolveValueVR(b *testing.B) {
	tests := []struct {
		name  string
		tag   *tag.Tag
		value any
	}{
		{name: benchmarkStringName, tag: tag.PatientName, value: benchmarkStringValue},
		{name: benchmarkUInt16Name, tag: tag.Rows, value: uint16(512)},
		{name: benchmarkFloat64sName, tag: tag.WindowCenter, value: []float64{0, 100}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				benchmarkValueVR, benchmarkValueError = resolveValueVR(tt.tag, tt.value)
			}
		})
	}
}

func BenchmarkDatasetAddOrUpdateValue(b *testing.B) {
	tests := []struct {
		name  string
		tag   *tag.Tag
		value any
	}{
		{name: benchmarkStringName, tag: tag.PatientName, value: benchmarkStringValue},
		{name: benchmarkUInt16Name, tag: tag.Rows, value: uint16(512)},
		{name: benchmarkFloat64sName, tag: tag.WindowCenter, value: []float64{0, 100}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			ds := New()
			b.ReportAllocs()
			for b.Loop() {
				benchmarkValueError = ds.AddOrUpdateValue(tt.tag, tt.value)
			}
		})
	}
}
