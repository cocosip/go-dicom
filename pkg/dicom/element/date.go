// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"fmt"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
)

// Date represents a DICOM element with VR = DA (Date).
// Format: YYYYMMDD
type Date struct {
	*String
}

// NewDate creates a new DA element with the given date strings.
// Date format should be YYYYMMDD.
func NewDate(t *tag.Tag, dates []string) *Date {
	str := NewString(t, vr.DA, dates)
	return &Date{String: str}
}

// NewDateFromTime creates a new DA element from time.Time values.
func NewDateFromTime(t *tag.Tag, times []time.Time) *Date {
	dates := make([]string, len(times))
	for i, tm := range times {
		dates[i] = tm.Format("20060102") // YYYYMMDD
	}
	return NewDate(t, dates)
}

// NewDateFromBuffer creates a DA element from an existing buffer.
func NewDateFromBuffer(t *tag.Tag, buf buffer.ByteBuffer, enc encoding.Encoding) *Date {
	str := NewStringFromBuffer(t, vr.DA, buf, enc)
	return &Date{String: str}
}

// GetDate parses and returns the date at the specified index.
func (d *Date) GetDate(index int) (time.Time, error) {
	dateStr := d.GetValue(index)
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("date at index %d is empty", index)
	}

	// DICOM DA format: YYYYMMDD
	return time.Parse("20060102", dateStr)
}

// GetDates parses and returns all dates.
func (d *Date) GetDates() ([]time.Time, error) {
	count := d.Count()
	if count == 0 {
		return nil, nil
	}

	dates := make([]time.Time, count)
	for i := 0; i < count; i++ {
		date, err := d.GetDate(i)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date at index %d: %w", i, err)
		}
		dates[i] = date
	}
	return dates, nil
}

// Time represents a DICOM element with VR = TM (Time).
// Format: HHMMSS.FFFFFF (fractional seconds are optional)
type Time struct {
	*String
}

// NewTime creates a new TM element with the given time strings.
// Time format should be HHMMSS or HHMMSS.FFFFFF.
func NewTime(t *tag.Tag, times []string) *Time {
	str := NewString(t, vr.TM, times)
	return &Time{String: str}
}

// NewTimeFromTime creates a new TM element from time.Time values.
func NewTimeFromTime(t *tag.Tag, times []time.Time) *Time {
	timeStrs := make([]string, len(times))
	for i, tm := range times {
		timeStrs[i] = tm.Format("150405.000000") // HHMMSS.FFFFFF
	}
	return NewTime(t, timeStrs)
}

// NewTimeFromBuffer creates a TM element from an existing buffer.
func NewTimeFromBuffer(t *tag.Tag, buf buffer.ByteBuffer, enc encoding.Encoding) *Time {
	str := NewStringFromBuffer(t, vr.TM, buf, enc)
	return &Time{String: str}
}

// GetTime parses and returns the time at the specified index.
func (tm *Time) GetTime(index int) (time.Time, error) {
	timeStr := tm.GetValue(index)
	if timeStr == "" {
		return time.Time{}, fmt.Errorf("time at index %d is empty", index)
	}

	// Try different DICOM TM formats
	formats := []string{
		"150405.000000", // HHMMSS.FFFFFF
		"150405.000",    // HHMMSS.FFF
		"150405",        // HHMMSS
		"1504",          // HHMM
		"15",            // HH
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("failed to parse time %q", timeStr)
}

// GetTimes parses and returns all times.
func (tm *Time) GetTimes() ([]time.Time, error) {
	count := tm.Count()
	if count == 0 {
		return nil, nil
	}

	times := make([]time.Time, count)
	for i := 0; i < count; i++ {
		t, err := tm.GetTime(i)
		if err != nil {
			return nil, fmt.Errorf("failed to parse time at index %d: %w", i, err)
		}
		times[i] = t
	}
	return times, nil
}

// DateTime represents a DICOM element with VR = DT (Date Time).
// Format: YYYYMMDDHHMMSS.FFFFFF&ZZXX
type DateTime struct {
	*String
}

// NewDateTime creates a new DT element with the given datetime strings.
// DateTime format should be YYYYMMDDHHMMSS or YYYYMMDDHHMMSS.FFFFFF&ZZXX.
func NewDateTime(t *tag.Tag, datetimes []string) *DateTime {
	str := NewString(t, vr.DT, datetimes)
	return &DateTime{String: str}
}

// NewDateTimeFromTime creates a new DT element from time.Time values.
func NewDateTimeFromTime(t *tag.Tag, times []time.Time) *DateTime {
	datetimes := make([]string, len(times))
	for i, tm := range times {
		// Format with timezone: YYYYMMDDHHMMSS.FFFFFF&ZZXX
		datetimes[i] = tm.Format("20060102150405.000000-0700")
	}
	return NewDateTime(t, datetimes)
}

// NewDateTimeFromBuffer creates a DT element from an existing buffer.
func NewDateTimeFromBuffer(t *tag.Tag, buf buffer.ByteBuffer, enc encoding.Encoding) *DateTime {
	str := NewStringFromBuffer(t, vr.DT, buf, enc)
	return &DateTime{String: str}
}

// GetDateTime parses and returns the datetime at the specified index.
func (dt *DateTime) GetDateTime(index int) (time.Time, error) {
	dtStr := dt.GetValue(index)
	if dtStr == "" {
		return time.Time{}, fmt.Errorf("datetime at index %d is empty", index)
	}

	// Try different DICOM DT formats
	formats := []string{
		"20060102150405.000000-0700", // YYYYMMDDHHMMSS.FFFFFF&ZZXX
		"20060102150405.000000",      // YYYYMMDDHHMMSS.FFFFFF
		"20060102150405",             // YYYYMMDDHHMMSS
		"200601021504",               // YYYYMMDDHHMM
		"2006010215",                 // YYYYMMDDHH
		"20060102",                   // YYYYMMDD
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dtStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("failed to parse datetime %q", dtStr)
}

// GetDateTimes parses and returns all datetimes.
func (dt *DateTime) GetDateTimes() ([]time.Time, error) {
	count := dt.Count()
	if count == 0 {
		return nil, nil
	}

	times := make([]time.Time, count)
	for i := 0; i < count; i++ {
		t, err := dt.GetDateTime(i)
		if err != nil {
			return nil, fmt.Errorf("failed to parse datetime at index %d: %w", i, err)
		}
		times[i] = t
	}
	return times, nil
}
