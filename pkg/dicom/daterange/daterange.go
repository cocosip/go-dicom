// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package daterange provides DICOM date range types for query/retrieve operations.
// DICOM date ranges are used in C-FIND, C-GET, and C-MOVE queries to specify
// date/time constraints.
//
// Format: "YYYYMMDD-YYYYMMDD" where either side can be empty:
//   - "20240101-20241231" - from Jan 1 to Dec 31, 2024
//   - "20240101-" - from Jan 1, 2024 onwards (open-ended)
//   - "-20241231" - up to Dec 31, 2024 (open-started)
//   - "20240101" - exactly Jan 1, 2024 (single date)
package daterange

import (
	"errors"
	"fmt"
	"time"
)

var (
	// ErrInvalidDateFormat indicates an invalid DICOM date format
	ErrInvalidDateFormat = errors.New("invalid DICOM date format")
	// ErrInvalidTimeFormat indicates an invalid DICOM time format
	ErrInvalidTimeFormat = errors.New("invalid DICOM time format")
	// ErrInvalidDateTimeFormat indicates an invalid DICOM datetime format
	ErrInvalidDateTimeFormat = errors.New("invalid DICOM datetime format")
	// ErrInvalidRangeFormat indicates an invalid DICOM range format
	ErrInvalidRangeFormat = errors.New("invalid DICOM range format")
)

// Range represents a generic range with minimum and maximum bounds.
type Range[T comparable] struct {
	Minimum T
	Maximum T
}

// NewRange creates a new Range with the given minimum and maximum values.
func NewRange[T comparable](minVal, maxVal T) *Range[T] {
	return &Range[T]{Minimum: minVal, Maximum: maxVal}
}

// DateRange represents a DICOM date range used in query/retrieve operations.
type DateRange struct {
	*Range[time.Time]
}

// NewDateRange creates a new DateRange with the given minimum and maximum dates.
func NewDateRange(minVal, maxVal time.Time) *DateRange {
	return &DateRange{Range: NewRange(minVal, maxVal)}
}

// NewDateRangeAll creates a DateRange that encompasses all dates.
func NewDateRangeAll() *DateRange {
	return &DateRange{Range: NewRange(time.Time{}, time.Time{})}
}

// ParseDateRange parses a DICOM date range string.
// Supported formats:
//   - "YYYYMMDD-YYYYMMDD" - range from start to end
//   - "YYYYMMDD-" - from start date onwards
//   - "-YYYYMMDD" - up to end date
//   - "YYYYMMDD" - exact date (single value)
func ParseDateRange(s string) (*DateRange, error) {
	if s == "" || s == "-" {
		return NewDateRangeAll(), nil
	}

	minVal := time.Time{}
	maxVal := time.Time{}

	dashIdx := findRangeDash(s)

	if dashIdx == -1 {
		date, err := parseDicomDate(s)
		if err != nil {
			return nil, err
		}
		return NewDateRange(date, date), nil
	}

	startPart := s[:dashIdx]
	endPart := s[dashIdx+1:]

	if startPart != "" {
		var err error
		minVal, err = parseDicomDate(startPart)
		if err != nil {
			return nil, err
		}
	}

	if endPart != "" {
		var err error
		maxVal, err = parseDicomDate(endPart)
		if err != nil {
			return nil, err
		}
	}

	return NewDateRange(minVal, maxVal), nil
}

// MustParseDateRange parses a DICOM date range string, panicking on error.
func MustParseDateRange(s string) *DateRange {
	r, err := ParseDateRange(s)
	if err != nil {
		panic(err)
	}
	return r
}

// parseDicomDate parses a DICOM date string (DA format: YYYYMMDD).
func parseDicomDate(s string) (time.Time, error) {
	if len(s) < 8 {
		return time.Time{}, fmt.Errorf("%w: %q", ErrInvalidDateFormat, s)
	}

	formats := []string{
		"20060102",
		"2006.01.02",
		"2006/01/02",
		"200601",
		"2006.01",
		"2006",
	}

	for _, format := range formats {
		if len(s) >= len(format) {
			if t, err := time.Parse(format, s[:len(format)]); err == nil {
				return t, nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("%w: %q", ErrInvalidDateFormat, s)
}

// String returns the DICOM date range string representation (DA format).
func (dr *DateRange) String() string {
	return dr.Format("20060102")
}

// Format returns the date range formatted according to the given layout.
func (dr *DateRange) Format(layout string) string {
	minStr := ""
	maxStr := ""

	if !dr.Minimum.IsZero() {
		minStr = dr.Minimum.Format(layout)
	}
	if !dr.Maximum.IsZero() {
		maxStr = dr.Maximum.Format(layout)
	}

	if minStr == "" && maxStr == "" {
		return ""
	}

	return minStr + "-" + maxStr
}

// Contains checks if the given date is within the range.
func (dr *DateRange) Contains(date time.Time) bool {
	minOK := dr.Minimum.IsZero() || !date.Before(dr.Minimum)
	maxOK := dr.Maximum.IsZero() || !date.After(dr.Maximum)
	return minOK && maxOK
}

// IsEmpty returns true if both minimum and maximum are zero values (all dates).
func (dr *DateRange) IsEmpty() bool {
	return dr.Minimum.IsZero() && dr.Maximum.IsZero()
}

// IsExact returns true if the range represents a single exact date.
func (dr *DateRange) IsExact() bool {
	return !dr.Minimum.IsZero() && dr.Minimum.Equal(dr.Maximum)
}

// IsOpenEnded returns true if the range has no upper bound.
func (dr *DateRange) IsOpenEnded() bool {
	return !dr.Minimum.IsZero() && dr.Maximum.IsZero()
}

// IsOpenStarted returns true if the range has no lower bound.
func (dr *DateRange) IsOpenStarted() bool {
	return dr.Minimum.IsZero() && !dr.Maximum.IsZero()
}

// IsRange returns true if the range spans multiple dates.
func (dr *DateRange) IsRange() bool {
	return !dr.Minimum.IsZero() && !dr.Maximum.IsZero() && !dr.Minimum.Equal(dr.Maximum)
}

// Join extends the range to include the given date if necessary.
func (dr *DateRange) Join(date time.Time) {
	if dr.Minimum.IsZero() || date.Before(dr.Minimum) {
		dr.Minimum = date
	}
	if dr.Maximum.IsZero() || date.After(dr.Maximum) {
		dr.Maximum = date
	}
}

// TimeRange represents a DICOM time range used in query/retrieve operations.
type TimeRange struct {
	*Range[time.Time]
}

// NewTimeRange creates a new TimeRange with the given minimum and maximum times.
func NewTimeRange(minVal, maxVal time.Time) *TimeRange {
	return &TimeRange{Range: NewRange(minVal, maxVal)}
}

// ParseTimeRange parses a DICOM time range string (TM format).
func ParseTimeRange(s string) (*TimeRange, error) {
	if s == "" || s == "-" {
		return &TimeRange{Range: NewRange(time.Time{}, time.Time{})}, nil
	}

	minVal := time.Time{}
	maxVal := time.Time{}

	dashIdx := findRangeDash(s)

	if dashIdx == -1 {
		t, err := parseDicomTime(s)
		if err != nil {
			return nil, err
		}
		return NewTimeRange(t, t), nil
	}

	startPart := s[:dashIdx]
	endPart := s[dashIdx+1:]

	if startPart != "" {
		var err error
		minVal, err = parseDicomTime(startPart)
		if err != nil {
			return nil, err
		}
	}

	if endPart != "" {
		var err error
		maxVal, err = parseDicomTime(endPart)
		if err != nil {
			return nil, err
		}
	}

	return NewTimeRange(minVal, maxVal), nil
}

// parseDicomTime parses a DICOM time string (TM format).
func parseDicomTime(s string) (time.Time, error) {
	formats := []string{
		"150405.000000",
		"150405.000",
		"150405",
		"1504",
		"15",
	}

	for _, format := range formats {
		if len(s) >= len(format) {
			if t, err := time.Parse(format, s[:len(format)]); err == nil {
				return t, nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("%w: %q", ErrInvalidTimeFormat, s)
}

// String returns the DICOM time range string representation (TM format).
func (tr *TimeRange) String() string {
	minStr := ""
	maxStr := ""

	if !tr.Minimum.IsZero() {
		minStr = tr.Minimum.Format("150405")
	}
	if !tr.Maximum.IsZero() {
		maxStr = tr.Maximum.Format("150405")
	}

	if minStr == "" && maxStr == "" {
		return ""
	}

	return minStr + "-" + maxStr
}

// Contains checks if the given time is within the range.
func (tr *TimeRange) Contains(t time.Time) bool {
	minOK := tr.Minimum.IsZero() || !t.Before(tr.Minimum)
	maxOK := tr.Maximum.IsZero() || !t.After(tr.Maximum)
	return minOK && maxOK
}

// DateTimeRange represents a DICOM datetime range used in query/retrieve operations.
type DateTimeRange struct {
	*Range[time.Time]
}

// NewDateTimeRange creates a new DateTimeRange with the given minimum and maximum datetimes.
func NewDateTimeRange(minVal, maxVal time.Time) *DateTimeRange {
	return &DateTimeRange{Range: NewRange(minVal, maxVal)}
}

// ParseDateTimeRange parses a DICOM datetime range string (DT format).
// Format: "YYYYMMDDHHMMSS.FFFFFF&ZZXX-YYYYMMDDHHMMSS.FFFFFF&ZZXX"
func ParseDateTimeRange(s string) (*DateTimeRange, error) {
	if s == "" || s == "-" {
		return &DateTimeRange{Range: NewRange(time.Time{}, time.Time{})}, nil
	}

	minVal := time.Time{}
	maxVal := time.Time{}

	dashIdx := findRangeDash(s)

	if dashIdx == -1 {
		dt, err := parseDicomDateTime(s)
		if err != nil {
			return nil, err
		}
		return NewDateTimeRange(dt, dt), nil
	}

	startPart := s[:dashIdx]
	endPart := s[dashIdx+1:]

	if startPart != "" {
		var err error
		minVal, err = parseDicomDateTime(startPart)
		if err != nil {
			return nil, err
		}
	}

	if endPart != "" {
		var err error
		maxVal, err = parseDicomDateTime(endPart)
		if err != nil {
			return nil, err
		}
	}

	return NewDateTimeRange(minVal, maxVal), nil
}

// parseDicomDateTime parses a DICOM datetime string (DT format).
func parseDicomDateTime(s string) (time.Time, error) {
	formats := []string{
		"20060102150405.000000-0700",
		"20060102150405.000000",
		"20060102150405",
		"200601021504",
		"2006010215",
		"20060102",
	}

	for _, format := range formats {
		if len(s) >= len(format) {
			if t, err := time.Parse(format, s[:len(format)]); err == nil {
				return t, nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("%w: %q", ErrInvalidDateTimeFormat, s)
}

// String returns the DICOM datetime range string representation (DT format).
func (dtr *DateTimeRange) String() string {
	minStr := ""
	maxStr := ""

	if !dtr.Minimum.IsZero() {
		minStr = dtr.Minimum.Format("20060102150405")
	}
	if !dtr.Maximum.IsZero() {
		maxStr = dtr.Maximum.Format("20060102150405")
	}

	if minStr == "" && maxStr == "" {
		return ""
	}

	return minStr + "-" + maxStr
}

// Contains checks if the given datetime is within the range.
func (dtr *DateTimeRange) Contains(dt time.Time) bool {
	minOK := dtr.Minimum.IsZero() || !dt.Before(dtr.Minimum)
	maxOK := dtr.Maximum.IsZero() || !dt.After(dtr.Maximum)
	return minOK && maxOK
}

// findRangeDash finds the dash character that separates range bounds.
// It returns -1 if no dash is found or if the dash is part of a timezone offset.
// Timezone offsets have format like -0700 (dash followed by 4 digits after another digit).
func findRangeDash(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			// Check if this is a timezone offset (dash followed by 4+ digits after a digit)
			// Timezone pattern: digit + '-' + 4 digits
			if i > 0 && i+4 < len(s) {
				isDigitBefore := s[i-1] >= '0' && s[i-1] <= '9'
				isDigitsAfter := (s[i+1] >= '0' && s[i+1] <= '9') &&
					(s[i+2] >= '0' && s[i+2] <= '9') &&
					(s[i+3] >= '0' && s[i+3] <= '9') &&
					(s[i+4] >= '0' && s[i+4] <= '9')
				// Also check if there are more digits after (could be longer timezone)
				hasMoreDigits := i+5 < len(s) && (s[i+5] >= '0' && s[i+5] <= '9')

				if isDigitBefore && isDigitsAfter && !hasMoreDigits {
					// This looks like a timezone offset, skip it
					continue
				}
			}
			return i
		}
	}
	return -1
}
