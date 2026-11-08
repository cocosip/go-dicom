// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package vr_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestValidateAE(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{testCaseValidAE, "MYAETITLE", false},
		{"valid with spaces", "MY_AE", false},
		{testCaseTooLong, "VERYLONGAETITLE1234", true},
		{"only spaces", "    ", true},
		{"contains backslash", "AE\\TITLE", true},
		{"outside default repertoire", "医院", true},
		{testCaseEmpty, "", false}, // Empty is typically allowed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateAE(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAE(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidateAS(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid years", "025Y", false},
		{"valid months", "006M", false},
		{"valid weeks", "052W", false},
		{"valid days", "365D", false},
		{testCaseInvalidFormat, "25Y", true},
		{"invalid unit", "025X", true},
		{"not digits", "ABCY", true},
		{testCaseEmpty, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateAS(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAS(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidateCS(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid code", "CT", false},
		{"valid with numbers", "T1", false},
		{"valid with underscore", "MY_CODE", false},
		{testCaseTooLong, "VERYLONGCODESTRIN", true},
		{"lowercase", "lowercase", true},
		{"special char", "CODE#1", true},
		{testCaseEmpty, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateCS(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCS(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidateDA(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid date", "20250106", false},
		{"valid with spaces", "20250106 ", false},
		{"date range", "20250101-20250131", false},
		{"invalid month", "20251301", true},
		{"invalid day", "20250132", true},
		{"wrong format", "2025-01-06", true},
		{"too short", "202501", true},
		{testCaseEmpty, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateDA(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDA(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidateDS(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"integer", "123", false},
		{"decimal", "123.456", false},
		{"scientific", "1.23e10", false},
		{"negative", "-123.45", false},
		{testCaseWithSpaces, " 123.45 ", false},
		{testCaseTooLong, "12345678901234567", true},
		{testCaseInvalidFormat, "abc", true},
		{"multiple dots", "12.34.56", true},
		{testCaseEmpty, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateDS(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDS(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidateIS(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"positive integer", "123", false},
		{"negative integer", "-456", false},
		{"with plus", "+789", false},
		{testCaseWithSpaces, " 123 ", false},
		{testCaseTooLong, "1234567890123", true},
		{"decimal", "123.45", true},
		{"not a number", "abc", true},
		{"too large", "999999999999", true},
		{testCaseEmpty, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateIS(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateIS(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidateLO(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid short", testPatientNameText, false},
		{"valid max length", "A234567890123456789012345678901234567890123456789012345678901234", false},
		{testCaseTooLong, "A2345678901234567890123456789012345678901234567890123456789012345", true},
		{testCaseEmpty, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateLO(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateLO(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidatePN(t *testing.T) {
	// Each component group may be up to 64 characters; the whole value may be
	// longer when it contains multiple component groups.
	group64 := "A234567890123456789012345678901234567890123456789012345678901234"
	group65 := group64 + "5"
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"single group at limit", group64, false},
		{"single group too long", group65, true},
		{"multiple groups each at limit", group64 + "=" + group64 + "=" + group64, false},
		{"one group too long among several", group64 + "=" + group65, true},
		{testCaseEmpty, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := vr.ValidatePN(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTM(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"hours only", "14", false},
		{"hours minutes", "1430", false},
		{"full time", "143015", false},
		{"with fraction", "143015.123", false},
		{testCaseWithSpaces, "143015 ", false},
		{"invalid hour", "2530", true},
		{"invalid minute", "1460", true},
		{"invalid second", "145061", true},
		{testCaseInvalidFormat, "14:30", true},
		{testCaseEmpty, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateTM(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTM(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidateUI(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"valid UID", "1.2.840.10008.1.1", false},
		{"valid with trailing space", "1.2.3.4 ", false},
		{testCaseTooLong, "1.234567890123456789012345678901234567890123456789012345678901234567890", true},
		{"leading zero", "1.2.01.4", true},
		{"invalid char", "1.2.3a.4", true},
		{"trailing dot", "1.2.3.", true},
		{"leading dot", ".1.2.3", true},
		{testCaseEmpty, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateUI(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUI(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestVRValidateString(t *testing.T) {
	// Test validation through VR object
	err := vr.AE.ValidateString("VALIDAETITLE")
	if err != nil {
		t.Errorf("AE.ValidateString() unexpected error: %v", err)
	}

	err = vr.AE.ValidateString("TOOLONGAETITLE12345")
	if err == nil {
		t.Error("AE.ValidateString() expected error for too long value")
	}

	// Test with validation disabled
	vr.PerformValidation = false
	err = vr.AE.ValidateString("TOOLONGAETITLE12345")
	if err != nil {
		t.Errorf("AE.ValidateString() should not error when validation disabled: %v", err)
	}
	vr.PerformValidation = true // Reset
}

func TestValidateDT(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"date only", "20251108", false},
		{"full datetime", "20251108172000", false},
		{"fractional seconds", "20251108172000.123456", false},
		{"conformant offset", "20251108172000.000000+0800", false},
		{"negative offset", "20251108172000-0500", false},
		{"trailing space", "20251108172000 ", false},
		{"ampersand offset", "20251108172000.000000&+0800", true},
		{"negative zero offset", "20251108172000+0000", false},
		{"minus zero offset", "20251108172000-0000", true},
		{"odd digit count", "202511081", true},
		{"bad month", "20251308", true},
		{"bad hour", "2025110825", true},
		{"fraction without seconds", "2025110817.5", true},
		{"too many fraction digits", "20251108172000.1234567", true},
		{"nondigit", "2025110X", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vr.ValidateDT(tt.value)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateDT(%q) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestValidateDARejectsInvalidCalendar(t *testing.T) {
	for _, value := range []string{"20250230", "20250001", "20200132"} {
		if err := vr.ValidateDA(value); err == nil {
			t.Fatalf("ValidateDA(%q) = nil, want error", value)
		}
	}
	if err := vr.ValidateDA("20240229"); err != nil {
		t.Fatalf("ValidateDA(20240229) error = %v", err)
	}
}

func TestValidateDTOffsetRange(t *testing.T) {
	if err := vr.ValidateDT("20251108172000+1400"); err != nil {
		t.Fatalf("ValidateDT(+1400) error = %v", err)
	}
	for _, value := range []string{"20251108172000+2300", "20251108172000-1300", "20251108172000+1401"} {
		if err := vr.ValidateDT(value); err == nil {
			t.Fatalf("ValidateDT(%q) = nil, want error", value)
		}
	}
}
