// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse

import "testing"

func TestPriority_String(t *testing.T) {
	tests := []struct {
		name     string
		priority Priority
		expected string
	}{
		{"Low", PriorityLow, "LOW"},
		{"Medium", PriorityMedium, "MEDIUM"},
		{"High", PriorityHigh, "HIGH"},
		{"Unknown", Priority(0x9999), "UNKNOWN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.priority.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestPriority_Values(t *testing.T) {
	tests := []struct {
		name     string
		priority Priority
		expected uint16
	}{
		{"Medium", PriorityMedium, 0x0000},
		{"High", PriorityHigh, 0x0001},
		{"Low", PriorityLow, 0x0002},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint16(tt.priority) != tt.expected {
				t.Errorf("Expected value 0x%04X for %s, got 0x%04X", tt.expected, tt.name, uint16(tt.priority))
			}
		})
	}
}
