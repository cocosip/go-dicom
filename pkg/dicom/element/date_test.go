// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element_test

import (
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

// TestDate tests DA (Date) elements
func TestDate(t *testing.T) {
	t.Run("FromString", func(t *testing.T) {
		elem := element.NewDate(tag.StudyDate, []string{"20250106"})

		date, err := elem.GetDate(0)
		if err != nil {
			t.Fatalf("GetDate() error = %v", err)
		}

		if date.Year() != 2025 || date.Month() != 1 || date.Day() != 6 {
			t.Errorf("GetDate() = %v, want 2025-01-06", date)
		}
	})

	t.Run("FromTime", func(t *testing.T) {
		testTime := time.Date(2025, 1, 6, 0, 0, 0, 0, time.UTC)
		elem := element.NewDateFromTime(tag.StudyDate, []time.Time{testTime})

		date, err := elem.GetDate(0)
		if err != nil {
			t.Fatalf("GetDate() error = %v", err)
		}

		if date.Year() != 2025 || date.Month() != 1 || date.Day() != 6 {
			t.Errorf("GetDate() = %v, want 2025-01-06", date)
		}
	})

	t.Run("MultipleValues", func(t *testing.T) {
		elem := element.NewDate(tag.StudyDate, []string{"20250106", "20250107"})

		dates, err := elem.GetDates()
		if err != nil {
			t.Fatalf("GetDates() error = %v", err)
		}

		if len(dates) != 2 {
			t.Errorf("len(GetDates()) = %d, want 2", len(dates))
		}
	})

	t.Run("InvalidDate", func(t *testing.T) {
		elem := element.NewDate(tag.StudyDate, []string{"invalid"})

		_, err := elem.GetDate(0)
		if err == nil {
			t.Error("GetDate() should return error for invalid date")
		}
	})
}

// TestTime tests TM (Time) elements
func TestTime(t *testing.T) {
	t.Run("FullFormat", func(t *testing.T) {
		elem := element.NewTime(tag.StudyTime, []string{"123456.789012"})

		tm, err := elem.GetTime(0)
		if err != nil {
			t.Fatalf("GetTime() error = %v", err)
		}

		if tm.Hour() != 12 || tm.Minute() != 34 || tm.Second() != 56 {
			t.Errorf("GetTime() = %v, want 12:34:56", tm)
		}
	})

	t.Run("ShortFormat", func(t *testing.T) {
		elem := element.NewTime(tag.StudyTime, []string{"1234"})

		tm, err := elem.GetTime(0)
		if err != nil {
			t.Fatalf("GetTime() error = %v", err)
		}

		if tm.Hour() != 12 || tm.Minute() != 34 {
			t.Errorf("GetTime() = %v, want 12:34", tm)
		}
	})

	t.Run("FromTime", func(t *testing.T) {
		testTime := time.Date(2025, 1, 6, 12, 34, 56, 0, time.UTC)
		elem := element.NewTimeFromTime(tag.StudyTime, []time.Time{testTime})

		tm, err := elem.GetTime(0)
		if err != nil {
			t.Fatalf("GetTime() error = %v", err)
		}

		if tm.Hour() != 12 || tm.Minute() != 34 || tm.Second() != 56 {
			t.Errorf("GetTime() = %v, want 12:34:56", tm)
		}
	})
}

// TestDateTime tests DT (DateTime) elements
func TestDateTime(t *testing.T) {
	t.Run("FullFormat", func(t *testing.T) {
		elem := element.NewDateTime(tag.AcquisitionDateTime, []string{"20250106123456"})

		dt, err := elem.GetDateTime(0)
		if err != nil {
			t.Fatalf("GetDateTime() error = %v", err)
		}

		if dt.Year() != 2025 || dt.Month() != 1 || dt.Day() != 6 {
			t.Errorf("GetDateTime() date = %v, want 2025-01-06", dt)
		}
		if dt.Hour() != 12 || dt.Minute() != 34 || dt.Second() != 56 {
			t.Errorf("GetDateTime() time = %v, want 12:34:56", dt)
		}
	})

	t.Run("DateOnly", func(t *testing.T) {
		elem := element.NewDateTime(tag.AcquisitionDateTime, []string{"20250106"})

		dt, err := elem.GetDateTime(0)
		if err != nil {
			t.Fatalf("GetDateTime() error = %v", err)
		}

		if dt.Year() != 2025 || dt.Month() != 1 || dt.Day() != 6 {
			t.Errorf("GetDateTime() = %v, want 2025-01-06", dt)
		}
	})

	t.Run("FromTime", func(t *testing.T) {
		testTime := time.Date(2025, 1, 6, 12, 34, 56, 0, time.UTC)
		elem := element.NewDateTimeFromTime(tag.AcquisitionDateTime, []time.Time{testTime})

		dt, err := elem.GetDateTime(0)
		if err != nil {
			t.Fatalf("GetDateTime() error = %v", err)
		}

		if dt.Year() != 2025 || dt.Month() != 1 || dt.Day() != 6 {
			t.Errorf("GetDateTime() date = %v, want 2025-01-06", dt)
		}
	})
}
