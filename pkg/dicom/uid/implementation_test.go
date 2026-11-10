// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package uid

import "testing"

func TestFormatImplementationVersionName(t *testing.T) {
	tests := []struct {
		version string
		want    string
	}{
		{"v0.9.0", "GO_DICOM_0_9_0"},
		{"v1.12.3", "GO_DICOM_1_12_3"},
		{"v1.4.2+dirty", "GO_DICOM_1_4_2"},
		{"v0.10.0-0.20260101000000-abcdef123456", "GO_DICOM_0_10_0"},
		{"v0.9.0-rc.1", "GO_DICOM_0_9_0"},
		{"v0.9.1-0.20260101120000-abcdef123456", "GO_DICOM_0_9_1"},
		{"v10.123.456", "GO_DICOM_10_123"},
	}
	for _, tt := range tests {
		got := formatImplementationVersionName(tt.version)
		if got != tt.want {
			t.Errorf("formatImplementationVersionName(%q) = %q, want %q", tt.version, got, tt.want)
		}
		if len(got) > implementationVersionNameMaxLength {
			t.Errorf("formatImplementationVersionName(%q) = %q, length %d exceeds %d",
				tt.version, got, len(got), implementationVersionNameMaxLength)
		}
		for i := 0; i < len(got); i++ {
			c := got[i]
			if (c < 'A' || c > 'Z') && (c < '0' || c > '9') && c != '_' {
				t.Errorf("formatImplementationVersionName(%q) = %q, contains invalid character %q",
					tt.version, got, c)
			}
		}
	}
}
