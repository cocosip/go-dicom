// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

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
	"strconv"
	"strings"
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
//
// PS3.5 Table 6.2-1 defines DA as exactly eight digits, so partial-year and
// dotted/slashed ACR-NEMA forms are rejected.
func ParseDateRange(s string) (*DateRange, error) {
	if s == "" || s == "-" {
		return NewDateRangeAll(), nil
	}

	minVal := time.Time{}
	maxVal := time.Time{}

	dashIdx := strings.IndexByte(s, '-')

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
//
// PS3.5 Table 6.2-1 fixes DA at exactly eight digits; the retired ACR-NEMA
// YYYY.MM.DD form and partial year/month forms are not accepted. time.Parse
// rejects impossible calendar dates such as 20241301 and 20240230.
func parseDicomDate(s string) (time.Time, error) {
	if len(s) != 8 || !isASCIIDigits(s) {
		return time.Time{}, fmt.Errorf("%w: %q", ErrInvalidDateFormat, s)
	}
	t, err := time.Parse("20060102", s)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %q", ErrInvalidDateFormat, s)
	}
	return t, nil
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

	// minDigits and maxDigits record the significant TM digits of the parsed
	// bounds (2, 4, 6, or 6 plus fractional digits); 0 means full precision.
	// Keeping them preserves reduced-precision query values across String.
	minDigits, maxDigits int
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
	minDigits := 0
	maxDigits := 0

	dashIdx := strings.IndexByte(s, '-')

	if dashIdx == -1 {
		t, digits, err := parseDicomTime(s)
		if err != nil {
			return nil, err
		}
		r := NewTimeRange(t, t)
		r.minDigits, r.maxDigits = digits, digits
		return r, nil
	}

	startPart := s[:dashIdx]
	endPart := s[dashIdx+1:]

	if startPart != "" {
		var err error
		minVal, minDigits, err = parseDicomTime(startPart)
		if err != nil {
			return nil, err
		}
	}

	if endPart != "" {
		var err error
		maxVal, maxDigits, err = parseDicomTime(endPart)
		if err != nil {
			return nil, err
		}
	}

	r := NewTimeRange(minVal, maxVal)
	r.minDigits, r.maxDigits = minDigits, maxDigits
	return r, nil
}

// NormalizeTimeFraction expands a fractional-seconds component of 1-6 digits
// to the six-digit form expected by the time layouts. DICOM permits any
// precision from 1 to 6 digits (PS3.5 Table 6.2-1). Values without a valid
// fraction are returned unchanged.
func NormalizeTimeFraction(value string) string {
	dot := strings.IndexByte(value, '.')
	if dot < 0 {
		return value
	}
	end := dot + 1
	for end < len(value) && value[end] >= '0' && value[end] <= '9' {
		end++
	}
	digits := end - dot - 1
	if digits < 1 || digits >= 6 {
		return value
	}
	return value[:dot+1] + value[dot+1:end] + strings.Repeat("0", 6-digits) + value[end:]
}

// parseDicomTime parses a DICOM time string (TM format) and reports how many
// significant digits it carried (2, 4, 6, or 6 plus fractional digits).
func parseDicomTime(s string) (time.Time, int, error) {
	s = strings.TrimRight(s, " \x00")
	digits := len(s)
	// PS3.5 Table 6.2-1 permits 1 to 6 fractional-second digits, but the layouts
	// below use the fixed six-digit form. Normalize the fraction first so e.g.
	// "120000.5" is 12:00:00.5 instead of silently truncating the fraction.
	normalized := NormalizeTimeFraction(s)

	formats := []string{
		"150405.000000",
		"150405.000",
		"150405",
		"1504",
		"15",
	}

	for _, format := range formats {
		if len(normalized) == len(format) {
			if t, err := time.Parse(format, normalized[:len(format)]); err == nil {
				return t, digits, nil
			}
		}
	}

	return time.Time{}, 0, fmt.Errorf("%w: %q", ErrInvalidTimeFormat, s)
}

// String returns the DICOM time range string representation (TM format). A
// reduced-precision bound (e.g. "1200") keeps its precision so it keeps its
// standard matching semantics (PS3.4 C.2.2.2.5.2).
func (tr *TimeRange) String() string {
	minStr := ""
	maxStr := ""

	if !tr.Minimum.IsZero() {
		minStr = formatDicomTimeWithDigits(tr.Minimum, tr.minDigits)
	}
	if !tr.Maximum.IsZero() {
		maxStr = formatDicomTimeWithDigits(tr.Maximum, tr.maxDigits)
	}

	if minStr == "" && maxStr == "" {
		return ""
	}

	return minStr + "-" + maxStr
}

// formatDicomTime formats a time as PS3.5 TM: HHMMSS[.FFFFFF]. The fractional
// seconds component is emitted only when the value has one (up to six digits).
func formatDicomTime(t time.Time) string {
	return t.Format("150405") + formatDicomFraction(t)
}

// formatDicomTimeWithDigits formats a time with the precision it was parsed
// with. Zero means full precision.
func formatDicomTimeWithDigits(t time.Time, digits int) string {
	switch {
	case digits == 2:
		return t.Format("15")
	case digits == 4:
		return t.Format("1504")
	case digits == 6:
		return t.Format("150405")
	case digits > 6:
		return t.Format("150405") + formatDicomFractionDigits(t, digits-7)
	default:
		return formatDicomTime(t)
	}
}

// formatDicomFractionDigits returns the fractional-seconds suffix truncated to
// the given number of digits (PS3.5 permits at most six).
func formatDicomFractionDigits(t time.Time, digits int) string {
	if digits <= 0 {
		return ""
	}
	if digits > 6 {
		digits = 6
	}
	fmtSeconds := fmt.Sprintf("%06d", t.Nanosecond()/1000)[:digits]
	fmtSeconds = strings.TrimRight(fmtSeconds, "0")
	if fmtSeconds == "" {
		return ""
	}
	return "." + fmtSeconds
}

// formatDicomFraction returns the fractional-seconds suffix for a time,
// truncated to the six digits PS3.5 Table 6.2-1 permits. Trailing zeros are
// dropped, so "120000.5" round-trips as "120000.5" and not "120000.500000".
func formatDicomFraction(t time.Time) string {
	nanoseconds := t.Nanosecond()
	if nanoseconds == 0 {
		return ""
	}
	microseconds := nanoseconds / 1000
	if microseconds == 0 {
		// PS3.5 permits at most six fractional digits, so a sub-microsecond
		// value cannot be represented; omit it rather than emit a bare ".".
		return ""
	}
	return "." + strings.TrimRight(fmt.Sprintf("%06d", microseconds), "0")
}

// formatDicomUTCOffset returns the &ZZXX suffix for a time, or an empty string
// when the time is in UTC. PS3.5 Table 6.2-1 forbids the "-0000" offset, so
// sub-minute offsets (which DICOM cannot express) are rounded to the nearest
// minute instead.
func formatDicomUTCOffset(t time.Time) string {
	_, offset := t.Zone()
	sign := byte('+')
	if offset < 0 {
		sign = '-'
		offset = -offset
	}
	minutes := (offset + 30) / 60
	if minutes == 0 {
		return ""
	}
	if minutes > 14*60 {
		minutes = 14 * 60
	}
	return fmt.Sprintf("%c%02d%02d", sign, minutes/60, minutes%60)
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

	// minDigits/maxDigits and minFrac/maxFrac record the precision the bounds
	// were parsed with (date/time digits and fractional digits); zero means
	// full precision. String uses them to keep reduced-precision values.
	minDigits, maxDigits int
	minFrac, maxFrac     int
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
	minDigits, maxDigits := 0, 0
	minFrac, maxFrac := 0, 0

	dashIdx := findRangeDash(s)

	if dashIdx == -1 {
		dt, digits, frac, err := parseDicomDateTime(s)
		if err != nil {
			return nil, err
		}
		r := NewDateTimeRange(dt, dt)
		r.minDigits, r.maxDigits = digits, digits
		r.minFrac, r.maxFrac = frac, frac
		return r, nil
	}

	startPart := s[:dashIdx]
	endPart := s[dashIdx+1:]

	if startPart != "" {
		var err error
		minVal, minDigits, minFrac, err = parseDicomDateTime(startPart)
		if err != nil {
			return nil, err
		}
	}

	if endPart != "" {
		var err error
		maxVal, maxDigits, maxFrac, err = parseDicomDateTime(endPart)
		if err != nil {
			return nil, err
		}
	}

	r := NewDateTimeRange(minVal, maxVal)
	r.minDigits, r.maxDigits = minDigits, maxDigits
	r.minFrac, r.maxFrac = minFrac, maxFrac
	return r, nil
}

// parseDicomDateTime parses a DICOM datetime string (DT format), honouring an
// optional fractional-second part (1-6 digits) and an optional &ZZXX timezone
// offset such as "+0800". A missing offset is treated as UTC. It also reports
// the number of date/time digits and fractional digits for round-tripping.
func parseDicomDateTime(s string) (time.Time, int, int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, 0, 0, fmt.Errorf("%w: empty value", ErrInvalidDateTimeFormat)
	}

	loc := time.UTC
	if i := strings.IndexAny(s, "+-"); i > 0 {
		off := s[i:]
		s = s[:i]
		if len(off) != 5 {
			return time.Time{}, 0, 0, fmt.Errorf("%w: invalid timezone offset %q", ErrInvalidDateTimeFormat, off)
		}
		sign := 1
		switch off[0] {
		case '-':
			sign = -1
		case '+':
		default:
			return time.Time{}, 0, 0, fmt.Errorf("%w: invalid timezone offset %q", ErrInvalidDateTimeFormat, off)
		}
		hh, err1 := strconv.Atoi(off[1:3])
		mm, err2 := strconv.Atoi(off[3:5])
		// PS3.5 Table 6.2-1: the UTC offset range is -1200 to +1400, and -0000
		// is not permitted. Keep this in step with vr.ValidateDT.
		if err1 != nil || err2 != nil || off == "-0000" ||
			hh > 14 || mm > 59 || (hh == 14 && mm != 0) || (sign < 0 && hh > 12) {
			return time.Time{}, 0, 0, fmt.Errorf("%w: invalid timezone offset %q", ErrInvalidDateTimeFormat, off)
		}
		loc = time.FixedZone("", sign*(hh*3600+mm*60))
	}

	frac := ""
	if dot := strings.IndexByte(s, '.'); dot >= 0 {
		frac = s[dot+1:]
		s = s[:dot]
	}
	digits := len(s)

	layouts := map[int]string{
		4:  "2006",
		6:  "200601",
		8:  "20060102",
		10: "2006010215",
		12: "200601021504",
		14: "20060102150405",
	}
	layout, ok := layouts[digits]
	if !ok {
		return time.Time{}, 0, 0, fmt.Errorf("%w: %q", ErrInvalidDateTimeFormat, s)
	}
	t, err := time.ParseInLocation(layout, s, loc)
	if err != nil {
		return time.Time{}, 0, 0, fmt.Errorf("%w: %q", ErrInvalidDateTimeFormat, s)
	}

	fracDigits := 0
	if frac != "" {
		if len(frac) > 9 {
			frac = frac[:9]
		}
		fracDigits = len(frac)
		if fracDigits > 6 {
			fracDigits = 6
		}
		ns, err := strconv.Atoi(frac)
		if err != nil {
			return time.Time{}, 0, 0, fmt.Errorf("%w: invalid fraction %q", ErrInvalidDateTimeFormat, frac)
		}
		for i := len(frac); i < 9; i++ {
			ns *= 10
		}
		t = t.Add(time.Duration(ns))
	}

	return t, digits, fracDigits, nil
}

// String returns the DICOM datetime range string representation (DT format). A
// reduced-precision bound (e.g. "202401") keeps its precision so it keeps its
// standard matching semantics (PS3.4 C.2.2.2.5.3).
func (dtr *DateTimeRange) String() string {
	minStr := ""
	maxStr := ""

	if !dtr.Minimum.IsZero() {
		minStr = formatDicomDateTimeWithDigits(dtr.Minimum, dtr.minDigits, dtr.minFrac)
	}
	if !dtr.Maximum.IsZero() {
		maxStr = formatDicomDateTimeWithDigits(dtr.Maximum, dtr.maxDigits, dtr.maxFrac)
	}

	if minStr == "" && maxStr == "" {
		return ""
	}

	return minStr + "-" + maxStr
}

// formatDicomDateTimeWithDigits formats a datetime as PS3.5 DT:
// YYYYMMDDHHMMSS[.FFFFFF][&ZZXX] using the precision it was parsed with;
// digits == 0 means full precision.
func formatDicomDateTimeWithDigits(t time.Time, digits, fracDigits int) string {
	var layout string
	switch digits {
	case 4:
		layout = "2006"
	case 6:
		layout = "200601"
	case 8:
		layout = "20060102"
	case 10:
		layout = "2006010215"
	case 12:
		layout = "200601021504"
	default:
		layout = "20060102150405"
	}
	return t.Format(layout) + formatDicomFractionDigits(t, fracDigits) + formatDicomUTCOffset(t)
}

// Contains checks if the given datetime is within the range.
func (dtr *DateTimeRange) Contains(dt time.Time) bool {
	minOK := dtr.Minimum.IsZero() || !dt.Before(dtr.Minimum)
	maxOK := dtr.Maximum.IsZero() || !dt.After(dtr.Maximum)
	return minOK && maxOK
}

// isASCIIDigits reports whether every byte is an ASCII digit.
func isASCIIDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// validZZXX reports whether d is a plausible 4-digit UTC offset (HHMM within
// the PS3.5 range -1200..+1400).
func validZZXX(d string) bool {
	if len(d) != 4 {
		return false
	}
	hh := int(d[0]-'0')*10 + int(d[1]-'0')
	mm := int(d[2]-'0')*10 + int(d[3]-'0')
	return mm < 60 && (hh < 14 || (hh == 14 && mm == 0))
}

// findRangeDash finds the dash that separates the bounds of a DT range.
//
// It returns -1 when no separator is present. A dash that introduces a
// trailing "-ZZXX" UTC offset is not a separator: PS3.5 Table 6.2-1 note 5
// allows the offset after a partial date/time (e.g. "202401-0500" is the
// datetime 2024-01 with a -05:00 offset). The offset can only appear at the
// end of a value (or immediately before the range separator), and only when
// the text before it is a legal DT prefix, so the decision is made by
// look-ahead instead of by absolute position. This keeps TM values such as
// "1000-1800" a range because "1800" is not a legal UTC offset.
func findRangeDash(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != '-' {
			continue
		}
		if i == 0 {
			return 0 // open-started range
		}
		if isUTCOffsetAt(s, i) {
			continue
		}
		return i
	}
	return -1
}

// FindRangeDash returns the index of the range-separating dash in a DT value,
// or -1 when the value has no range separator. Dashes that introduce a trailing
// UTC offset (e.g. "202401-0500") are ignored.
func FindRangeDash(s string) int {
	return findRangeDash(s)
}

// isUTCOffsetAt reports whether the dash at s[i] introduces a "-ZZXX" UTC
// offset on a legal DT value prefix.
func isUTCOffsetAt(s string, i int) bool {
	if i+5 > len(s) {
		return false
	}
	offset := s[i+1 : i+5]
	if !isASCIIDigits(offset) || !validZZXX(offset) {
		return false
	}
	if i+5 != len(s) && s[i+5] != '-' {
		return false
	}
	return isDateTimePrefix(s[:i])
}

// isDateTimePrefix reports whether s is a legal DT value prefix: YYYY,
// YYYYMM, YYYYMMDD, YYYYMMDDHH, YYYYMMDDHHMM or YYYYMMDDHHMMSS, optionally
// followed by a fractional-second part.
func isDateTimePrefix(s string) bool {
	if dot := strings.IndexByte(s, '.'); dot >= 0 {
		if !isASCIIDigits(s[dot+1:]) {
			return false
		}
		s = s[:dot]
	}
	if !isASCIIDigits(s) {
		return false
	}
	switch len(s) {
	case 4, 6, 8, 10, 12, 14:
		return true
	default:
		return false
	}
}
