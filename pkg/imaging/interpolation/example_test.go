// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package interpolation_test

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging/interpolation"
)

func ExampleGrid_Resize() {
	// Each row has two pixels followed by one padding value.
	grid, err := interpolation.NewGrid(
		[]float64{0, 10, 99, 20, 30},
		2,
		2,
		3,
	)
	if err != nil {
		panic(err)
	}

	sample, inside := grid.Bilinear(0.5, 0.5)
	resized, err := grid.Resize(3, 3, interpolation.ModeBilinear)
	if err != nil {
		panic(err)
	}
	_, outside := grid.Nearest(-0.1, 0)

	fmt.Printf("sample: %.0f, inside=%t\n", sample, inside)
	fmt.Printf("resized: %v\n", resized)
	fmt.Printf("outside: %t\n", outside)
	// Output:
	// sample: 15, inside=true
	// resized: [0 5 10 10 15 20 20 25 30]
	// outside: false
}
