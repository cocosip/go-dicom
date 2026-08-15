// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package reconstruction

import (
	"context"
	"math"
	"reflect"
	"testing"
)

func TestNewStackBuildsInclusiveLazyStandardPlanes(t *testing.T) {
	volume := testVolumeWithPixels(t,
		[]float64{0, 5, 10},
		[][]uint16{{0, 10, 20, 30}, {50, 60, 70, 80}, {100, 110, 120, 130}},
	)
	tests := []struct {
		name      string
		stackType StackType
		wantLen   int
		wantRows  int
		wantCols  int
	}{
		{name: "axial", stackType: StackTypeAxial, wantLen: 3, wantRows: 2, wantCols: 2},
		{name: "coronal", stackType: StackTypeCoronal, wantLen: 2, wantRows: 11, wantCols: 2},
		{name: "sagittal", stackType: StackTypeSagittal, wantLen: 2, wantRows: 11, wantCols: 2},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stack, err := NewStack(volume, test.stackType, 1, map[StackType]float64{
				StackTypeAxial: 5, StackTypeCoronal: 1, StackTypeSagittal: 1,
			}[test.stackType])
			if err != nil {
				t.Fatalf("NewStack() error = %v", err)
			}
			if stack.Len() != test.wantLen {
				t.Fatalf("Len() = %d, want %d", stack.Len(), test.wantLen)
			}
			spec, err := stack.Spec(stack.Len() - 1)
			if err != nil {
				t.Fatalf("Spec(last) error = %v", err)
			}
			if spec.Rows != test.wantRows || spec.Columns != test.wantCols {
				t.Fatalf("last dimensions = %dx%d, want %dx%d", spec.Rows, spec.Columns, test.wantRows, test.wantCols)
			}
		})
	}

	axial, err := NewStack(volume, StackTypeAxial, 1, 5)
	if err != nil {
		t.Fatalf("NewStack(axial) error = %v", err)
	}
	last, err := axial.Materialize(context.Background(), 2, CutOptions{Workers: 2})
	if err != nil {
		t.Fatalf("Materialize(last) error = %v", err)
	}
	if !reflect.DeepEqual(last.Values, []float64{100, 110, 120, 130}) {
		t.Fatalf("last axial values = %v", last.Values)
	}
}

func TestSliceMinMaxRejectsNonFiniteValues(t *testing.T) {
	slice := &Slice{Values: []float64{1, math.NaN()}, Valid: []bool{true, true}}
	if _, _, err := slice.MinMax(); err == nil {
		t.Fatal("MinMax() accepted a non-finite valid sample")
	}
}

func TestNewStackRejectsInvalidConfigurationAndIndex(t *testing.T) {
	volume := testVolumeWithPixels(t, []float64{0, 1}, [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}})
	for _, call := range []func() error{
		func() error { _, err := NewStack(nil, StackTypeAxial, 1, 1); return err },
		func() error { _, err := NewStack(volume, StackType(99), 1, 1); return err },
		func() error { _, err := NewStack(volume, StackTypeAxial, 0, 1); return err },
		func() error { _, err := NewStack(volume, StackTypeAxial, 1, 0); return err },
	} {
		if err := call(); err == nil {
			t.Fatal("NewStack() accepted invalid configuration")
		}
	}
	stack, err := NewStack(volume, StackTypeAxial, 1, 1)
	if err != nil {
		t.Fatalf("NewStack() error = %v", err)
	}
	if _, err := stack.Materialize(context.Background(), stack.Len(), CutOptions{}); err == nil {
		t.Fatal("Materialize() accepted out-of-range index")
	}
}

func TestNewStackRejectsExcessiveSliceCountBeforeAllocation(t *testing.T) {
	volume := testVolumeWithPixels(t, []float64{0, 1}, [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}})
	if _, err := NewStack(volume, StackTypeAxial, 1, 0.00001); err == nil {
		t.Fatal("NewStack() accepted more than 65535 output slices")
	}
}

func TestStandardCoronalAndSagittalStacksProduceExpectedPixels(t *testing.T) {
	volume := testVolumeWithPixels(t,
		[]float64{0, 5, 10},
		[][]uint16{{0, 10, 20, 30}, {50, 60, 70, 80}, {100, 110, 120, 130}},
	)
	for _, test := range []struct {
		name      string
		stackType StackType
		first     []float64
		middle    []float64
		last      []float64
	}{
		{name: "coronal", stackType: StackTypeCoronal, first: []float64{0, 10}, middle: []float64{50, 60}, last: []float64{100, 110}},
		{name: "sagittal", stackType: StackTypeSagittal, first: []float64{0, 20}, middle: []float64{50, 70}, last: []float64{100, 120}},
	} {
		t.Run(test.name, func(t *testing.T) {
			stack, err := NewStack(volume, test.stackType, 1, 1)
			if err != nil {
				t.Fatalf("NewStack() error = %v", err)
			}
			slice, err := stack.Materialize(context.Background(), 0, CutOptions{Workers: 2})
			if err != nil {
				t.Fatalf("Materialize() error = %v", err)
			}
			if !reflect.DeepEqual(slice.Values[0:2], test.first) ||
				!reflect.DeepEqual(slice.Values[10:12], test.middle) ||
				!reflect.DeepEqual(slice.Values[20:22], test.last) {
				t.Fatalf("first/middle/last rows = %v/%v/%v", slice.Values[0:2], slice.Values[10:12], slice.Values[20:22])
			}
			for index, valid := range slice.Valid {
				if !valid {
					t.Fatalf("sample %d is invalid", index)
				}
			}
		})
	}
}

func TestSliceMinMaxAndRender8BitHonorValidityMask(t *testing.T) {
	slice := &Slice{
		Spec:   CutSpec{Rows: 2, Columns: 2},
		Values: []float64{0, 50, 100, -1000},
		Valid:  []bool{true, true, true, false},
	}
	minimum, maximum, err := slice.MinMax()
	if err != nil || minimum != 0 || maximum != 100 {
		t.Fatalf("MinMax() = %v/%v, error %v; want 0/100", minimum, maximum, err)
	}
	rendered, err := slice.Render8Bit(50.5, 101, 7)
	if err != nil {
		t.Fatalf("Render8Bit() error = %v", err)
	}
	if !reflect.DeepEqual(rendered, []byte{0, 128, 255, 7}) {
		t.Fatalf("Render8Bit() = %v, want [0 128 255 7]", rendered)
	}
}
