// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package vr_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

func TestRegexValidatorsReuseCompiledPatterns(t *testing.T) {
	tests := []struct {
		name     string
		validate func(string) error
		value    string
	}{
		{name: "AS", validate: vr.ValidateAS, value: "018Y"},
		{name: "CS", validate: vr.ValidateCS, value: "ORIGINAL"},
		{name: "DS", validate: vr.ValidateDS, value: "1.25"},
		{name: "IS", validate: vr.ValidateIS, value: "42"},
		{name: "TM", validate: vr.ValidateTM, value: "123456.123"},
		{name: "UI", validate: vr.ValidateUI, value: "1.2.840.10008.1.2.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocations := testing.AllocsPerRun(100, func() {
				if err := tt.validate(tt.value); err != nil {
					t.Fatalf("validate(%q) error = %v", tt.value, err)
				}
			})
			if allocations > 1 {
				t.Fatalf("validate(%q) allocations = %.0f, want at most 1", tt.value, allocations)
			}
		})
	}
}
