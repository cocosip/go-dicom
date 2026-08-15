// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging_test

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/imaging"
)

func ExampleHistogram_ApplyPercentWindow() {
	histogram, err := imaging.NewHistogram(0, 3)
	if err != nil {
		panic(err)
	}
	for _, value := range []int{0, 1, 1, 2, 2, 2, 3, 3, 3, 3} {
		histogram.Add(value)
	}

	if err := histogram.ApplyPercentWindow(60); err != nil {
		panic(err)
	}
	fmt.Printf("percent: %d..%d, count=%d\n", histogram.WindowStart(), histogram.WindowEnd(), histogram.WindowTotal())

	if err := histogram.ApplyWindow(1, 2); err != nil {
		panic(err)
	}
	fmt.Printf("explicit: %d..%d, count=%d\n", histogram.WindowStart(), histogram.WindowEnd(), histogram.WindowTotal())
	// Output:
	// percent: 0..2, count=6
	// explicit: 1..2, count=5
}
