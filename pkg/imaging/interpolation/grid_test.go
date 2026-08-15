// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package interpolation

import (
	"math"
	"testing"
)

const interpolationTolerance = 1e-9

func TestGridNearestAndBilinearSampling(t *testing.T) {
	grid, err := NewGrid([]float64{0, 10, 20, 30}, 2, 2, 2)
	if err != nil {
		t.Fatalf("NewGrid() error = %v", err)
	}
	if got, ok := grid.Nearest(0.5, 0.5); !ok || got != 30 {
		t.Fatalf("Nearest(0.5,0.5) = %v, %v, want 30", got, ok)
	}
	if got, ok := grid.Bilinear(0.5, 0.5); !ok || math.Abs(got-15) > interpolationTolerance {
		t.Fatalf("Bilinear(0.5,0.5) = %v, %v, want 15", got, ok)
	}
	if got, ok := grid.Bilinear(1, 1); !ok || got != 30 {
		t.Fatalf("Bilinear(1,1) = %v, %v, want 30", got, ok)
	}
	if _, ok := grid.Nearest(-0.01, 0); ok {
		t.Fatal("Nearest() accepted an out-of-bounds coordinate")
	}
}

func TestGridUsesStride(t *testing.T) {
	grid, err := NewGrid([]float64{1, 2, 99, 3, 4}, 2, 2, 3)
	if err != nil {
		t.Fatalf("NewGrid() error = %v", err)
	}
	if got, ok := grid.Nearest(1, 1); !ok || got != 4 {
		t.Fatalf("Nearest(1,1) = %v, %v, want 4", got, ok)
	}
}

func TestGridResizeBilinearAlignsEndpoints(t *testing.T) {
	grid, _ := NewGrid([]float64{0, 10, 20, 30}, 2, 2, 2)
	got, err := grid.Resize(3, 3, ModeBilinear)
	if err != nil {
		t.Fatalf("Resize() error = %v", err)
	}
	want := []float64{0, 5, 10, 10, 15, 20, 20, 25, 30}
	for index := range want {
		if math.Abs(got[index]-want[index]) > interpolationTolerance {
			t.Fatalf("Resize()[%d] = %v, want %v", index, got[index], want[index])
		}
	}
}

func TestGridResizeSinglePixel(t *testing.T) {
	grid, _ := NewGrid([]float64{7}, 1, 1, 1)
	for _, mode := range []Mode{ModeNearestNeighbor, ModeBilinear} {
		got, err := grid.Resize(3, 2, mode)
		if err != nil {
			t.Fatalf("Resize(%v) error = %v", mode, err)
		}
		for index, value := range got {
			if value != 7 {
				t.Fatalf("Resize(%v)[%d] = %v, want 7", mode, index, value)
			}
		}
	}
}

func TestGridRejectsInvalidLayoutAndResize(t *testing.T) {
	if _, err := NewGrid(nil, 1, 1, 1); err == nil {
		t.Fatal("NewGrid() accepted missing data")
	}
	if _, err := NewGrid([]float64{1, 2, 3, 4}, 2, 2, 1); err == nil {
		t.Fatal("NewGrid() accepted stride below width")
	}
	grid, _ := NewGrid([]float64{1}, 1, 1, 1)
	if _, err := grid.Resize(0, 1, ModeNearestNeighbor); err == nil {
		t.Fatal("Resize() accepted zero output width")
	}
	if _, err := grid.Resize(1, 1, Mode(99)); err == nil {
		t.Fatal("Resize() accepted an unknown mode")
	}
}
