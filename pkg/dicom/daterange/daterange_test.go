// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package daterange

import (
	"testing"
	"time"
)

const (
	testFullRangeName     = "full range"
	testFullDateRange     = "20240101-20241231"
	testOpenEndedDate     = "20240101-"
	testOpenEndedName     = "open-ended"
	testOpenStartedName   = "open-started"
	testEmptyRangeName    = "empty"
	testDateRangeLayout   = "20060102"
	testRangeStartDate    = "20240101"
	testRangeEndDate      = "20241231"
	testInvalidRangeInput = "invalid"
	testOpenEndedDateTime = "20240101080000-"

	testDateTimeWithFraction      = "20240115120000.5"
	testDateTimePartialWithOffset = "202401-0500"
	testTimeFraction              = "120000.5"
)

func TestDateRange_Parse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantMin   string
		wantMax   string
		wantErr   bool
		wantEmpty bool
		wantExact bool
		wantOpen  bool
	}{
		{
			name:      "empty string",
			input:     "",
			wantEmpty: true,
		},
		{
			name:      "single dash",
			input:     "-",
			wantEmpty: true,
		},
		{
			name:      "single date",
			input:     "20240115",
			wantMin:   "20240115",
			wantMax:   "20240115",
			wantExact: true,
		},
		{
			name:    testFullRangeName,
			input:   testFullDateRange,
			wantMin: testRangeStartDate,
			wantMax: testRangeEndDate,
		},
		{
			name:     "open-ended range",
			input:    testOpenEndedDate,
			wantMin:  testRangeStartDate,
			wantOpen: true,
		},
		{
			name:    "open-started range",
			input:   "-20241231",
			wantMax: testRangeEndDate,
		},
		{
			name:    "invalid date",
			input:   testInvalidRangeInput,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := ParseDateRange(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("ParseDateRange() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("ParseDateRange() error = %v", err)
				return
			}

			if tt.wantEmpty {
				if !r.IsEmpty() {
					t.Error("ParseDateRange() expected IsEmpty() = true")
				}
				return
			}

			if tt.wantMin != "" {
				wantMin, _ := time.Parse(testDateRangeLayout, tt.wantMin)
				if !r.Minimum.Equal(wantMin) {
					t.Errorf("Minimum = %v, want %v", r.Minimum, wantMin)
				}
			}

			if tt.wantMax != "" {
				wantMax, _ := time.Parse(testDateRangeLayout, tt.wantMax)
				if !r.Maximum.Equal(wantMax) {
					t.Errorf("Maximum = %v, want %v", r.Maximum, wantMax)
				}
			}

			if tt.wantExact && !r.IsExact() {
				t.Error("IsExact() = false, want true")
			}

			if tt.wantOpen && !r.IsOpenEnded() {
				t.Error("IsOpenEnded() = false, want true")
			}
		})
	}
}

func TestDateRange_String(t *testing.T) {
	tests := []struct {
		name string
		min  string
		max  string
		want string
	}{
		{
			name: testFullRangeName,
			min:  testRangeStartDate,
			max:  testRangeEndDate,
			want: testFullDateRange,
		},
		{
			name: testOpenEndedName,
			min:  testRangeStartDate,
			max:  "",
			want: testOpenEndedDate,
		},
		{
			name: testOpenStartedName,
			min:  "",
			max:  testRangeEndDate,
			want: "-20241231",
		},
		{
			name: "exact date",
			min:  "20240115",
			max:  "20240115",
			want: "20240115-20240115",
		},
		{
			name: testEmptyRangeName,
			min:  "",
			max:  "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDateRangeAll()
			if tt.min != "" {
				r.Minimum, _ = time.Parse(testDateRangeLayout, tt.min)
			}
			if tt.max != "" {
				r.Maximum, _ = time.Parse(testDateRangeLayout, tt.max)
			}

			got := r.String()
			if got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDateRange_Contains(t *testing.T) {
	r, _ := ParseDateRange(testFullDateRange)

	tests := []struct {
		date string
		want bool
	}{
		{"20240101", true},  // start boundary
		{"20241231", true},  // end boundary
		{"20240615", true},  // middle
		{"20231231", false}, // before
		{"20250101", false}, // after
	}

	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			date, _ := time.Parse(testDateRangeLayout, tt.date)
			if got := r.Contains(date); got != tt.want {
				t.Errorf("Contains(%s) = %v, want %v", tt.date, got, tt.want)
			}
		})
	}
}

func TestDateRange_Join(t *testing.T) {
	r := NewDateRangeAll()

	date1, _ := time.Parse(testDateRangeLayout, testRangeStartDate)
	date2, _ := time.Parse(testDateRangeLayout, testRangeEndDate)
	date3, _ := time.Parse(testDateRangeLayout, "20240615")

	r.Join(date2)
	if !r.Minimum.Equal(date2) {
		t.Errorf("after first join, Minimum = %v, want %v", r.Minimum, date2)
	}

	r.Join(date1)
	if !r.Minimum.Equal(date1) {
		t.Errorf("after second join, Minimum = %v, want %v", r.Minimum, date1)
	}
	if !r.Maximum.Equal(date2) {
		t.Errorf("Maximum = %v, want %v", r.Maximum, date2)
	}

	r.Join(date3)
	if !r.Minimum.Equal(date1) {
		t.Errorf("Minimum changed unexpectedly")
	}
	if !r.Maximum.Equal(date2) {
		t.Errorf("Maximum changed unexpectedly")
	}
}

func TestTimeRange_Parse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:  "single time",
			input: "120000",
		},
		{
			name:  testFullRangeName,
			input: "080000-170000",
		},
		{
			name:  testOpenEndedName,
			input: "080000-",
		},
		{
			name:  testOpenStartedName,
			input: "-170000",
		},
		{
			name:  testEmptyRangeName,
			input: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseTimeRange(tt.input)
			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestDateTimeRange_Parse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:  "single datetime",
			input: "20240115120000",
		},
		{
			name:  testFullRangeName,
			input: "20240101080000-20241231170000",
		},
		{
			name:  testOpenEndedName,
			input: testOpenEndedDateTime,
		},
		{
			name:  testOpenStartedName,
			input: "-20241231170000",
		},
		{
			name:  testEmptyRangeName,
			input: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseDateTimeRange(tt.input)
			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestFindRangeDash(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"20240101080000-20241231170000", 14},
		{testOpenEndedDateTime, 14},
		{"-20241231170000", 0},
		{"20240115120000", -1},
		{"", -1},
		{"-", 0},
		// A trailing -ZZXX is a UTC offset, not a range separator.
		{testDateTimePartialWithOffset, -1},
		{"2024-0500", -1},
		{"2007-0500", -1},
		{"20240101-0500", -1},
		{"20240101080000-0500", -1},
		// "1800" is not a legal UTC offset, so this stays a range.
		{"1000-1800", 4},
		{"2024-2025", 4},
		// An offset may precede the range separator.
		{"20240101-0500-20241231-0500", 13},
		{"20240101-0500-", 13},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := findRangeDash(tt.input)
			if got != tt.want {
				t.Errorf("findRangeDash(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseDateRangeRejectsNonConformantDA(t *testing.T) {
	for _, input := range []string{
		"20241301",   // impossible month
		"20240230",   // impossible day
		"2024",       // partial year
		"202401",     // partial year/month
		"2024.01.01", // retired ACR-NEMA form
		"2024/01/01", // non-conformant form
		"20240101x",  // trailing garbage
	} {
		t.Run(input, func(t *testing.T) {
			if _, err := ParseDateRange(input); err == nil {
				t.Errorf("ParseDateRange(%q) = nil error, want ErrInvalidDateFormat", input)
			}
		})
	}
}

func TestParseDateTimeRangeTreatsTrailingUTCOffsetAsOffset(t *testing.T) {
	tests := []struct {
		input    string
		wantZone int // seconds east of UTC
		wantYear int
		wantOpen bool
	}{
		{testDateTimePartialWithOffset, -5 * 3600, 2024, false},
		{"2024-0500", -5 * 3600, 2024, false},
		{"20240101+0800", 8 * 3600, 2024, false},
		{"20240101-0500-", -5 * 3600, 2024, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r, err := ParseDateTimeRange(tt.input)
			if err != nil {
				t.Fatalf("ParseDateTimeRange(%q) error = %v", tt.input, err)
			}
			if _, offset := r.Minimum.Zone(); offset != tt.wantZone {
				t.Errorf("Minimum zone offset = %d, want %d", offset, tt.wantZone)
			}
			if r.Minimum.Year() != tt.wantYear {
				t.Errorf("Minimum year = %d, want %d", r.Minimum.Year(), tt.wantYear)
			}
			if got := r.Maximum.IsZero(); got != tt.wantOpen {
				t.Errorf("Maximum.IsZero() = %v, want %v", got, tt.wantOpen)
			}
		})
	}
}

func TestParseDateTimeRangeWithOffsetsOnBothBounds(t *testing.T) {
	r, err := ParseDateTimeRange("20240101-0500-20241231-0500")
	if err != nil {
		t.Fatalf("ParseDateTimeRange() error = %v", err)
	}
	if _, offset := r.Minimum.Zone(); offset != -5*3600 {
		t.Errorf("Minimum zone offset = %d, want %d", offset, -5*3600)
	}
	if _, offset := r.Maximum.Zone(); offset != -5*3600 {
		t.Errorf("Maximum zone offset = %d, want %d", offset, -5*3600)
	}
	if r.Minimum.Year() != 2024 || r.Maximum.Month() != time.December {
		t.Fatalf("range = %v..%v, want 2024-01-01..2024-12-31", r.Minimum, r.Maximum)
	}
}

func TestMustParseDateRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParseDateRange did not panic on invalid input")
		}
	}()

	MustParseDateRange(testInvalidRangeInput)
}

func TestParseTimeRangeFractionalSeconds(t *testing.T) {
	tests := []struct {
		input string
		want  time.Duration
	}{
		{testTimeFraction, 500 * time.Millisecond},
		{"120000.05", 50 * time.Millisecond},
		{"120000.123", 123 * time.Millisecond},
		{"120000.1234", 123400 * time.Microsecond},
		{"120000.123456", 123456 * time.Microsecond},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r, err := ParseTimeRange(tt.input)
			if err != nil {
				t.Fatalf("ParseTimeRange(%q) error = %v", tt.input, err)
			}
			if r.Minimum.Hour() != 12 || r.Minimum.Minute() != 0 || r.Minimum.Second() != 0 {
				t.Fatalf("Minimum = %v, want 12:00:00", r.Minimum)
			}
			if got := time.Duration(r.Minimum.Nanosecond()); got != tt.want {
				t.Errorf("fractional seconds = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTimeRangeDoesNotTreatBoundsAsTimezone(t *testing.T) {
	r, err := ParseTimeRange("1000-1800")
	if err != nil {
		t.Fatalf("ParseTimeRange() error = %v", err)
	}
	if r.Minimum.Hour() != 10 || r.Maximum.Hour() != 18 {
		t.Fatalf("range = %v..%v, want 10:00..18:00", r.Minimum, r.Maximum)
	}
}

func TestTimeRangeStringRoundTrip(t *testing.T) {
	values := []string{
		"120000",
		testTimeFraction,
		"120000.123456",
		"080000.5-170000.25",
		"080000-170000",
	}
	for _, input := range values {
		t.Run(input, func(t *testing.T) {
			r, err := ParseTimeRange(input)
			if err != nil {
				t.Fatalf("ParseTimeRange(%q) error = %v", input, err)
			}
			formatted := r.String()
			again, err := ParseTimeRange(formatted)
			if err != nil {
				t.Fatalf("ParseTimeRange(%q) error = %v", formatted, err)
			}
			if !again.Minimum.Equal(r.Minimum) {
				t.Errorf("minimum after round trip = %v, want %v (String() = %q)", again.Minimum, r.Minimum, formatted)
			}
			if !again.Maximum.Equal(r.Maximum) {
				t.Errorf("maximum after round trip = %v, want %v (String() = %q)", again.Maximum, r.Maximum, formatted)
			}
		})
	}
}

func TestDateTimeRangeStringRoundTrip(t *testing.T) {
	values := []string{
		"20240115120000",
		"20240115120000.5",
		"20240115120000.123456",
		"20240115120000+0800",
		"20240115120000.123456+0800",
		"20240115120000.5-0500",
		"20240101080000.5+0800-20241231170000-0500",
		testOpenEndedDateTime,
	}
	for _, input := range values {
		t.Run(input, func(t *testing.T) {
			r, err := ParseDateTimeRange(input)
			if err != nil {
				t.Fatalf("ParseDateTimeRange(%q) error = %v", input, err)
			}
			formatted := r.String()
			again, err := ParseDateTimeRange(formatted)
			if err != nil {
				t.Fatalf("ParseDateTimeRange(%q) error = %v", formatted, err)
			}
			if !again.Minimum.Equal(r.Minimum) {
				t.Errorf("minimum after round trip = %v, want %v (String() = %q)", again.Minimum, r.Minimum, formatted)
			}
			if !again.Maximum.Equal(r.Maximum) {
				t.Errorf("maximum after round trip = %v, want %v (String() = %q)", again.Maximum, r.Maximum, formatted)
			}
		})
	}
}

func TestDateTimeRangeStringOmitsAbsentComponents(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"20240115120000", "20240115120000-20240115120000"},
		{testDateTimeWithFraction, "20240115120000.5-20240115120000.5"},
		{"20240115120000.123456+0800", "20240115120000.123456+0800-20240115120000.123456+0800"},
		{"20240115120000-0500", "20240115120000-0500-20240115120000-0500"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r, err := ParseDateTimeRange(tt.input)
			if err != nil {
				t.Fatalf("ParseDateTimeRange(%q) error = %v", tt.input, err)
			}
			if got := r.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTimeRangeStringEmitsFraction(t *testing.T) {
	for input, want := range map[string]string{
		"120000":         "120000-120000",
		testTimeFraction: testTimeFraction + "-" + testTimeFraction,
		"120000.12345":   "120000.12345-120000.12345",
	} {
		r, err := ParseTimeRange(input)
		if err != nil {
			t.Fatalf("ParseTimeRange(%q) error = %v", input, err)
		}
		if got := r.String(); got != want {
			t.Errorf("ParseTimeRange(%q).String() = %q, want %q", input, got, want)
		}
	}
}

func TestTimeRangeStringPreservesReducedPrecision(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"12", "12-12"},
		{"1200", "1200-1200"},
		{"1200-1800", "1200-1800"},
		{"120000.5", "120000.5-120000.5"},
		{"120000-180000", "120000-180000"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r, err := ParseTimeRange(tt.input)
			if err != nil {
				t.Fatalf("ParseTimeRange(%q) error = %v", tt.input, err)
			}
			if got := r.String(); got != tt.want {
				t.Fatalf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDateTimeRangeStringPreservesReducedPrecision(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"2024", "2024-2024"},
		{"202401", "202401-202401"},
		{"2024011512", "2024011512-2024011512"},
		{testDateTimeWithFraction, "20240115120000.5-20240115120000.5"},
		{testDateTimePartialWithOffset, "202401-0500-202401-0500"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r, err := ParseDateTimeRange(tt.input)
			if err != nil {
				t.Fatalf("ParseDateTimeRange(%q) error = %v", tt.input, err)
			}
			if got := r.String(); got != tt.want {
				t.Fatalf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
