// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package interpolation

import (
	"runtime"
	"testing"
)

func BenchmarkGridResizeBilinear(b *testing.B) {
	const inputSize = 256
	data := make([]float64, inputSize*inputSize)
	for index := range data {
		data[index] = float64(index % 4096)
	}
	grid, err := NewGrid(data, inputSize, inputSize, inputSize)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.SetBytes(512 * 512 * 8)
	for b.Loop() {
		output, resizeErr := grid.Resize(512, 512, ModeBilinear)
		if resizeErr != nil {
			b.Fatal(resizeErr)
		}
		runtime.KeepAlive(output)
	}
}
