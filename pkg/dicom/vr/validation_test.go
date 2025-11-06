// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

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
		{"valid AE", "MYAETITLE", false},
		{"valid with spaces", "MY_AE", false},
		{"too long", "VERYLONGAETITLE1234", true},
		{"only spaces", "    ", true},
		{"contains backslash", "AE\\TITLE", true},
		{"empty", "", false}, // Empty is typically allowed
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
		{"invalid format", "25Y", true},
		{"invalid unit", "025X", true},
		{"not digits", "ABCY", true},
		{"empty", "", false},
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
		{"too long", "VERYLONGCODESTRIN", true},
		{"lowercase", "lowercase", true},
		{"special char", "CODE#1", true},
		{"empty", "", false},
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
		{"empty", "", false},
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
		{"with spaces", " 123.45 ", false},
		{"too long", "12345678901234567", true},
		{"invalid format", "abc", true},
		{"multiple dots", "12.34.56", true},
		{"empty", "", false},
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
		{"with spaces", " 123 ", false},
		{"too long", "1234567890123", true},
		{"decimal", "123.45", true},
		{"not a number", "abc", true},
		{"too large", "999999999999", true},
		{"empty", "", false},
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
		{"valid short", "Patient Name", false},
		{"valid max length", "A234567890123456789012345678901234567890123456789012345678901234", false},
		{"too long", "A2345678901234567890123456789012345678901234567890123456789012345", true},
		{"empty", "", false},
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
		{"with spaces", "143015 ", false},
		{"invalid hour", "2530", true},
		{"invalid minute", "1460", true},
		{"invalid second", "145061", true},
		{"invalid format", "14:30", true},
		{"empty", "", false},
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
		{"too long", "1.234567890123456789012345678901234567890123456789012345678901234567890", true},
		{"leading zero", "1.2.01.4", true},
		{"invalid char", "1.2.3a.4", true},
		{"trailing dot", "1.2.3.", true},
		{"leading dot", ".1.2.3", true},
		{"empty", "", false},
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
