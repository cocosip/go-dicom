// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package scanner

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
)

func BenchmarkScanMetadata(b *testing.B) {
	const fixtureCount = 32
	root := b.TempDir()
	for index := range fixtureCount {
		path := filepath.Join(root, fmt.Sprintf("IMAGE%03d", index))
		writeScannerDICOM(b, path, "Benchmark^Patient")
	}

	for _, workers := range []int{1, 4} {
		b.Run(fmt.Sprintf("Workers%d", workers), func(b *testing.B) {
			value, err := New(WithWorkers(workers))
			if err != nil {
				b.Fatalf("New() error = %v", err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for range b.N {
				summary, scanErr := value.Scan(context.Background(), []string{root}, nil)
				if scanErr != nil {
					b.Fatalf("Scan() error = %v", scanErr)
				}
				if summary.DICOMFiles != fixtureCount {
					b.Fatalf("DICOM files = %d, want %d", summary.DICOMFiles, fixtureCount)
				}
			}
		})
	}
}
