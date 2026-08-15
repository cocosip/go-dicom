// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"context"
	"testing"

	"github.com/cocosip/go-dicom/pkg/imaging/math3d"
)

func BenchmarkVolumeSample(b *testing.B) {
	volume := testVolumeWithPixels(b,
		[]float64{0, 5, 10},
		[][]uint16{{0, 10, 20, 30}, {50, 60, 70, 80}, {100, 110, 120, 130}},
	)
	point := math3d.Point3{X: 0.5, Y: 0.5, Z: 2.5}
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if _, _, err := volume.Sample(point); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkVolumeCut64x64(b *testing.B) {
	volume := testVolumeWithPixels(b,
		[]float64{0, 10},
		[][]uint16{{0, 10, 20, 30}, {100, 110, 120, 130}},
	)
	spec := CutSpec{
		TopLeft:             math3d.Point3{Z: 5},
		RowDirection:        math3d.Vector3{X: 1},
		ColumnDirection:     math3d.Vector3{Y: 1},
		Rows:                64,
		Columns:             64,
		PixelSpacingRows:    1.0 / 63,
		PixelSpacingColumns: 1.0 / 63,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if _, err := volume.Cut(context.Background(), spec, CutOptions{Workers: 1}); err != nil {
			b.Fatal(err)
		}
	}
}
