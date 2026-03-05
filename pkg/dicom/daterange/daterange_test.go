// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package daterange

import (
	"testing"
	"time"
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
			name:    "full range",
			input:   "20240101-20241231",
			wantMin: "20240101",
			wantMax: "20241231",
		},
		{
			name:     "open-ended range",
			input:    "20240101-",
			wantMin:  "20240101",
			wantOpen: true,
		},
		{
			name:    "open-started range",
			input:   "-20241231",
			wantMax: "20241231",
		},
		{
			name:    "invalid date",
			input:   "invalid",
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
				wantMin, _ := time.Parse("20060102", tt.wantMin)
				if !r.Minimum.Equal(wantMin) {
					t.Errorf("Minimum = %v, want %v", r.Minimum, wantMin)
				}
			}

			if tt.wantMax != "" {
				wantMax, _ := time.Parse("20060102", tt.wantMax)
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
			name: "full range",
			min:  "20240101",
			max:  "20241231",
			want: "20240101-20241231",
		},
		{
			name: "open-ended",
			min:  "20240101",
			max:  "",
			want: "20240101-",
		},
		{
			name: "open-started",
			min:  "",
			max:  "20241231",
			want: "-20241231",
		},
		{
			name: "exact date",
			min:  "20240115",
			max:  "20240115",
			want: "20240115-20240115",
		},
		{
			name: "empty",
			min:  "",
			max:  "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDateRangeAll()
			if tt.min != "" {
				r.Minimum, _ = time.Parse("20060102", tt.min)
			}
			if tt.max != "" {
				r.Maximum, _ = time.Parse("20060102", tt.max)
			}

			got := r.String()
			if got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDateRange_Contains(t *testing.T) {
	r, _ := ParseDateRange("20240101-20241231")

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
			date, _ := time.Parse("20060102", tt.date)
			if got := r.Contains(date); got != tt.want {
				t.Errorf("Contains(%s) = %v, want %v", tt.date, got, tt.want)
			}
		})
	}
}

func TestDateRange_Join(t *testing.T) {
	r := NewDateRangeAll()

	date1, _ := time.Parse("20060102", "20240101")
	date2, _ := time.Parse("20060102", "20241231")
	date3, _ := time.Parse("20060102", "20240615")

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
			name:  "full range",
			input: "080000-170000",
		},
		{
			name:  "open-ended",
			input: "080000-",
		},
		{
			name:  "open-started",
			input: "-170000",
		},
		{
			name:  "empty",
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
			name:  "full range",
			input: "20240101080000-20241231170000",
		},
		{
			name:  "open-ended",
			input: "20240101080000-",
		},
		{
			name:  "open-started",
			input: "-20241231170000",
		},
		{
			name:  "empty",
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
		{"20240101-20241231", 8},
		{"20240101-", 8},
		{"-20241231", 0},
		{"20240101", -1},
		{"", -1},
		{"-", 0},
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

func TestMustParseDateRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParseDateRange did not panic on invalid input")
		}
	}()

	MustParseDateRange("invalid")
}
