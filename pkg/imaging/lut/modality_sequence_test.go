// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package lut

import "testing"

func TestModalitySequenceLUTOwnsInputValues(t *testing.T) {
	values := []float64{10, 20, 30}
	table := NewModalitySequenceLUT(values, 0, false)
	if table == nil {
		t.Fatal("NewModalitySequenceLUT() returned nil")
	}

	values[1] = 99

	if got := table.Transform(1); got != 20 {
		t.Fatalf("Transform(1) = %v after caller mutation, want 20", got)
	}
}
