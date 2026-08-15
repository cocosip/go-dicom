// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"fmt"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/daterange"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"golang.org/x/text/encoding"
)

// Compile-time checks to ensure date/time types implement Element interface
var _ Element = (*Date)(nil)
var _ Element = (*Time)(nil)
var _ Element = (*DateTime)(nil)

// Date represents a DICOM element with VR = DA (Date).
// Format: YYYYMMDD
type Date struct {
	str *String
}

// NewDate creates a new DA element with the given date strings.
// Date format should be YYYYMMDD.
func NewDate(t *tag.Tag, dates []string) *Date {
	return &Date{str: NewString(t, vr.DA, dates)}
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
	return &Date{str: NewStringFromBuffer(t, vr.DA, buf, enc)}
}

// NewDateFromRange creates a DA element from a DateRange.
// This is used in DICOM query operations (C-FIND, C-GET, C-MOVE).
func NewDateFromRange(t *tag.Tag, r *daterange.DateRange) *Date {
	return &Date{str: NewString(t, vr.DA, []string{r.String()})}
}

// Tag returns the DICOM tag.
func (d *Date) Tag() *tag.Tag {
	return d.str.Tag()
}

// ValueRepresentation returns the VR.
func (d *Date) ValueRepresentation() *vr.VR {
	return d.str.ValueRepresentation()
}

// Buffer returns the binary data buffer.
func (d *Date) Buffer() buffer.ByteBuffer {
	return d.str.Buffer()
}

// Length returns the length in bytes.
func (d *Date) Length() uint32 {
	return d.str.Length()
}

// Count returns the number of values.
func (d *Date) Count() int {
	return d.str.Count()
}

// String returns a string representation.
func (d *Date) String() string {
	return d.str.String()
}

// Validate performs validation.
func (d *Date) Validate() error {
	return d.str.Validate()
}

func (d *Date) validateValue() error {
	return d.str.validateValue()
}

// GetValue returns the string value at the specified index.
func (d *Date) GetValue(index int) string {
	return d.str.GetValue(index)
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

// GetDateRange parses and returns the date as a DateRange.
// This supports DICOM query format: "YYYYMMDD-YYYYMMDD", "YYYYMMDD-", "-YYYYMMDD", or single date.
func (d *Date) GetDateRange() (*daterange.DateRange, error) {
	dateStr := d.str.GetString()
	if dateStr == "" {
		return daterange.NewDateRangeAll(), nil
	}
	return daterange.ParseDateRange(dateStr)
}

// IsRange returns true if the date value represents a range (contains "-").
func (d *Date) IsRange() bool {
	dateStr := d.str.GetString()
	for i := 0; i < len(dateStr); i++ {
		if dateStr[i] == '-' {
			return true
		}
	}
	return false
}

// Time represents a DICOM element with VR = TM (Time).
// Format: HHMMSS.FFFFFF (fractional seconds are optional)
type Time struct {
	str *String
}

// NewTime creates a new TM element with the given time strings.
// Time format should be HHMMSS or HHMMSS.FFFFFF.
func NewTime(t *tag.Tag, times []string) *Time {
	return &Time{str: NewString(t, vr.TM, times)}
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
	return &Time{str: NewStringFromBuffer(t, vr.TM, buf, enc)}
}

// Tag returns the DICOM tag.
func (tm *Time) Tag() *tag.Tag {
	return tm.str.Tag()
}

// ValueRepresentation returns the VR.
func (tm *Time) ValueRepresentation() *vr.VR {
	return tm.str.ValueRepresentation()
}

// Buffer returns the binary data buffer.
func (tm *Time) Buffer() buffer.ByteBuffer {
	return tm.str.Buffer()
}

// Length returns the length in bytes.
func (tm *Time) Length() uint32 {
	return tm.str.Length()
}

// Count returns the number of values.
func (tm *Time) Count() int {
	return tm.str.Count()
}

// String returns a string representation.
func (tm *Time) String() string {
	return tm.str.String()
}

// Validate performs validation.
func (tm *Time) Validate() error {
	return tm.str.Validate()
}

func (tm *Time) validateValue() error {
	return tm.str.validateValue()
}

// GetValue returns the string value at the specified index.
func (tm *Time) GetValue(index int) string {
	return tm.str.GetValue(index)
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

// NewTimeFromRange creates a TM element from a TimeRange.
// This is used in DICOM query operations (C-FIND, C-GET, C-MOVE).
func NewTimeFromRange(t *tag.Tag, r *daterange.TimeRange) *Time {
	return &Time{str: NewString(t, vr.TM, []string{r.String()})}
}

// GetTimeRange parses and returns the time as a TimeRange.
// This supports DICOM query format: "HHMMSS-HHMMSS", "HHMMSS-", "-HHMMSS", or single time.
func (tm *Time) GetTimeRange() (*daterange.TimeRange, error) {
	timeStr := tm.str.GetString()
	if timeStr == "" {
		return &daterange.TimeRange{Range: daterange.NewRange(time.Time{}, time.Time{})}, nil
	}
	return daterange.ParseTimeRange(timeStr)
}

// IsRange returns true if the time value represents a range (contains "-").
func (tm *Time) IsRange() bool {
	timeStr := tm.str.GetString()
	for i := 0; i < len(timeStr); i++ {
		if timeStr[i] == '-' {
			return true
		}
	}
	return false
}

// DateTime represents a DICOM element with VR = DT (Date Time).
// Format: YYYYMMDDHHMMSS.FFFFFF&ZZXX
type DateTime struct {
	str *String
}

// NewDateTime creates a new DT element with the given datetime strings.
// DateTime format should be YYYYMMDDHHMMSS or YYYYMMDDHHMMSS.FFFFFF&ZZXX.
func NewDateTime(t *tag.Tag, datetimes []string) *DateTime {
	return &DateTime{str: NewString(t, vr.DT, datetimes)}
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
	return &DateTime{str: NewStringFromBuffer(t, vr.DT, buf, enc)}
}

// Tag returns the DICOM tag.
func (dt *DateTime) Tag() *tag.Tag {
	return dt.str.Tag()
}

// ValueRepresentation returns the VR.
func (dt *DateTime) ValueRepresentation() *vr.VR {
	return dt.str.ValueRepresentation()
}

// Buffer returns the binary data buffer.
func (dt *DateTime) Buffer() buffer.ByteBuffer {
	return dt.str.Buffer()
}

// Length returns the length in bytes.
func (dt *DateTime) Length() uint32 {
	return dt.str.Length()
}

// Count returns the number of values.
func (dt *DateTime) Count() int {
	return dt.str.Count()
}

// String returns a string representation.
func (dt *DateTime) String() string {
	return dt.str.String()
}

// Validate performs validation.
func (dt *DateTime) Validate() error {
	return dt.str.Validate()
}

func (dt *DateTime) validateValue() error {
	return dt.str.validateValue()
}

// GetValue returns the string value at the specified index.
func (dt *DateTime) GetValue(index int) string {
	return dt.str.GetValue(index)
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

// NewDateTimeFromRange creates a DT element from a DateTimeRange.
// This is used in DICOM query operations (C-FIND, C-GET, C-MOVE).
func NewDateTimeFromRange(t *tag.Tag, r *daterange.DateTimeRange) *DateTime {
	return &DateTime{str: NewString(t, vr.DT, []string{r.String()})}
}

// GetDateTimeRange parses and returns the datetime as a DateTimeRange.
// This supports DICOM query format: "YYYYMMDDHHMMSS-YYYYMMDDHHMMSS", etc.
func (dt *DateTime) GetDateTimeRange() (*daterange.DateTimeRange, error) {
	dtStr := dt.str.GetString()
	if dtStr == "" {
		return &daterange.DateTimeRange{Range: daterange.NewRange(time.Time{}, time.Time{})}, nil
	}
	return daterange.ParseDateTimeRange(dtStr)
}

// IsRange returns true if the datetime value represents a range (contains "-").
// Distinguished from timezone offsets (&ZZXX) which also contain a dash:
// a timezone dash has exactly 4 digits after it; a range dash continues beyond.
func (dt *DateTime) IsRange() bool {
	dtStr := dt.str.GetString()
	return findRangeDash(dtStr) != -1
}

// findRangeDash returns the index of the range-separating dash, or -1 if none.
// Skips dashes that are part of a timezone offset (digit + '-' + exactly 4 digits,
// no more digits after). This matches the logic in daterange.findRangeDash.
func findRangeDash(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			if i > 0 && i+4 < len(s) {
				isDigitBefore := s[i-1] >= '0' && s[i-1] <= '9'
				isDigitsAfter := (s[i+1] >= '0' && s[i+1] <= '9') &&
					(s[i+2] >= '0' && s[i+2] <= '9') &&
					(s[i+3] >= '0' && s[i+3] <= '9') &&
					(s[i+4] >= '0' && s[i+4] <= '9')
				hasMoreDigits := i+5 < len(s) && (s[i+5] >= '0' && s[i+5] <= '9')
				if isDigitBefore && isDigitsAfter && !hasMoreDigits {
					continue
				}
			}
			return i
		}
	}
	return -1
}
