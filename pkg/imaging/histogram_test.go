// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import "testing"

func TestHistogramCountsAndClearsBins(t *testing.T) {
	histogram, err := NewHistogram(-1, 3)
	if err != nil {
		t.Fatalf("NewHistogram() error = %v", err)
	}
	for _, value := range []int{-2, -1, 0, 0, 2, 2, 2, 4} {
		histogram.Add(value)
	}
	if histogram.Total() != 6 || histogram.Count(-1) != 1 || histogram.Count(2) != 3 || histogram.Count(4) != 0 {
		t.Fatalf("histogram counts: total=%d, -1=%d, 2=%d, 4=%d", histogram.Total(), histogram.Count(-1), histogram.Count(2), histogram.Count(4))
	}
	histogram.Clear(2)
	if histogram.Total() != 3 || histogram.Count(2) != 0 {
		t.Fatalf("Clear(2): total=%d count=%d", histogram.Total(), histogram.Count(2))
	}
}

func TestHistogramWindows(t *testing.T) {
	histogram, err := NewHistogram(0, 3)
	if err != nil {
		t.Fatalf("NewHistogram() error = %v", err)
	}
	for value, count := range []int{1, 2, 3, 4} {
		for range count {
			histogram.Add(value)
		}
	}
	if err := histogram.ApplyPercentWindow(60); err != nil {
		t.Fatalf("ApplyPercentWindow() error = %v", err)
	}
	if histogram.WindowStart() != 0 || histogram.WindowEnd() != 2 || histogram.WindowTotal() != 6 {
		t.Fatalf("60%% window = %d..%d total %d", histogram.WindowStart(), histogram.WindowEnd(), histogram.WindowTotal())
	}
	if err := histogram.ApplyWindow(1, 2); err != nil {
		t.Fatalf("ApplyWindow() error = %v", err)
	}
	if histogram.WindowTotal() != 5 {
		t.Fatalf("explicit window total = %d, want 5", histogram.WindowTotal())
	}
}

func TestHistogramRejectsInvalidRanges(t *testing.T) {
	if _, err := NewHistogram(2, 1); err == nil {
		t.Fatal("NewHistogram() accepted max below min")
	}
	histogram, _ := NewHistogram(0, 3)
	if err := histogram.ApplyPercentWindow(0); err == nil {
		t.Fatal("ApplyPercentWindow() accepted 0 percent")
	}
	if err := histogram.ApplyWindow(-1, 2); err == nil {
		t.Fatal("ApplyWindow() accepted a window outside the histogram")
	}
}
