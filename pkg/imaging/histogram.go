// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package imaging

import "fmt"

// Histogram counts integer values in an inclusive range and maintains an
// independently selectable active window.
type Histogram struct {
	minimum int
	maximum int
	counts  []int
	total   int

	windowStart int
	windowEnd   int
	windowTotal int
}

// NewHistogram creates a histogram whose bins cover min through max.
func NewHistogram(minimum, maximum int) (*Histogram, error) {
	if maximum < minimum {
		return nil, fmt.Errorf("histogram maximum %d is below minimum %d", maximum, minimum)
	}
	histogram := &Histogram{
		minimum:     minimum,
		maximum:     maximum,
		counts:      make([]int, maximum-minimum+1),
		windowStart: minimum,
		windowEnd:   maximum,
	}
	return histogram, nil
}

// Add increments value when it lies inside the histogram range.
func (h *Histogram) Add(value int) {
	index, ok := h.index(value)
	if !ok {
		return
	}
	h.counts[index]++
	h.total++
	if value >= h.windowStart && value <= h.windowEnd {
		h.windowTotal++
	}
}

// Clear resets one bin to zero.
func (h *Histogram) Clear(value int) {
	index, ok := h.index(value)
	if !ok {
		return
	}
	count := h.counts[index]
	h.counts[index] = 0
	h.total -= count
	if value >= h.windowStart && value <= h.windowEnd {
		h.windowTotal -= count
	}
}

// Count returns the count for value, or zero outside the histogram range.
func (h *Histogram) Count(value int) int {
	index, ok := h.index(value)
	if !ok {
		return 0
	}
	return h.counts[index]
}

// Total returns the number of values in all bins.
func (h *Histogram) Total() int { return h.total }

// WindowStart returns the first active bin.
func (h *Histogram) WindowStart() int { return h.windowStart }

// WindowEnd returns the last active bin.
func (h *Histogram) WindowEnd() int { return h.windowEnd }

// WindowTotal returns the number of values in the active window.
func (h *Histogram) WindowTotal() int { return h.windowTotal }

// ApplyPercentWindow trims the less-populated edge until removing another bin
// would leave fewer than percent of all values.
func (h *Histogram) ApplyPercentWindow(percent int) error {
	if percent < 1 || percent > 100 {
		return fmt.Errorf("histogram percent must be in 1..100")
	}
	start, end, total := 0, len(h.counts)-1, h.total
	if percent != 100 && total != 0 {
		target := int(float64(total) * float64(percent) / 100)
		for total > target && start < end {
			if h.counts[start] >= h.counts[end] {
				if total-h.counts[start] < target {
					break
				}
				total -= h.counts[start]
				start++
			} else {
				if total-h.counts[end] < target {
					break
				}
				total -= h.counts[end]
				end--
			}
		}
	}
	h.windowStart = start + h.minimum
	h.windowEnd = end + h.minimum
	h.windowTotal = total
	return nil
}

// ApplyWindow selects an inclusive absolute bin range.
func (h *Histogram) ApplyWindow(start, end int) error {
	if start < h.minimum || end > h.maximum || start > end {
		return fmt.Errorf("histogram window %d..%d is outside %d..%d", start, end, h.minimum, h.maximum)
	}
	total := 0
	for value := start; value <= end; value++ {
		total += h.Count(value)
	}
	h.windowStart = start
	h.windowEnd = end
	h.windowTotal = total
	return nil
}

func (h *Histogram) index(value int) (int, bool) {
	if h == nil || value < h.minimum || value > h.maximum {
		return 0, false
	}
	return value - h.minimum, true
}
