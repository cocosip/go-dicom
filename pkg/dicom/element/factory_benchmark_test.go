// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/charset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/endian"
	"golang.org/x/text/encoding"
)

var (
	benchmarkFactoryElement Element
	benchmarkFactoryError   error
)

func TestNewElementFromValueCommonSlicesAvoidBoxing(t *testing.T) {
	context := CanonicalValueContext{
		TextEncodings: []encoding.Encoding{charset.Default},
		Endian:        endian.Little,
	}
	values := []float64{0, 100}
	allocations := testing.AllocsPerRun(100, func() {
		benchmarkFactoryElement, benchmarkFactoryError = NewElementFromValueWithContext(
			tag.WindowCenter,
			vr.DS,
			values,
			context,
		)
	})
	if benchmarkFactoryError != nil {
		t.Fatalf("NewElementFromValueWithContext() error = %v", benchmarkFactoryError)
	}
	if allocations > 21 {
		t.Fatalf("NewElementFromValueWithContext([]float64) allocations = %.0f, want at most 21", allocations)
	}
}

func BenchmarkNewElementFromValueWithContext(b *testing.B) {
	context := CanonicalValueContext{
		TextEncodings: []encoding.Encoding{charset.Default},
		Endian:        endian.Little,
	}
	tests := []struct {
		name  string
		tag   *tag.Tag
		vr    *vr.VR
		value any
	}{
		{name: "string", tag: tag.PatientName, vr: vr.PN, value: "Zhang^San"},
		{name: "uint16", tag: tag.Rows, vr: vr.US, value: uint16(512)},
		{name: "float64s", tag: tag.WindowCenter, vr: vr.DS, value: []float64{0, 100}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				benchmarkFactoryElement, benchmarkFactoryError = NewElementFromValueWithContext(
					tt.tag,
					tt.vr,
					tt.value,
					context,
				)
			}
		})
	}
}
