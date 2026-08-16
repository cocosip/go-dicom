// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import "testing"

func TestParseFileIDNormalizesDICOMPathComponents(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantString string
		wantParts  []string
	}{
		{name: "forward slash", path: testDirectoryID + "/" + testImageID, wantString: testDirectoryID + `\` + testImageID, wantParts: []string{testDirectoryID, testImageID}},
		{name: "backslash", path: testDirectoryID + `\` + testImageID, wantString: testDirectoryID + `\` + testImageID, wantParts: []string{testDirectoryID, testImageID}},
		{name: "single component", path: testDICOMDIRFileID, wantString: testDICOMDIRFileID, wantParts: []string{testDICOMDIRFileID}},
		{name: "eight components", path: "A/B/C/D/E/F/G/H", wantString: `A\B\C\D\E\F\G\H`, wantParts: []string{"A", "B", "C", "D", "E", "F", "G", "H"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ParseFileID(tt.path)
			if err != nil {
				t.Fatalf("ParseFileID() error = %v", err)
			}
			if got := id.String(); got != tt.wantString {
				t.Fatalf("String() = %q, want %q", got, tt.wantString)
			}
			parts := id.Components()
			if len(parts) != len(tt.wantParts) {
				t.Fatalf("Components() length = %d, want %d", len(parts), len(tt.wantParts))
			}
			for i := range parts {
				if parts[i] != tt.wantParts[i] {
					t.Fatalf("Components()[%d] = %q, want %q", i, parts[i], tt.wantParts[i])
				}
			}
		})
	}
}

func TestParseFileIDRejectsInvalidDICOMPaths(t *testing.T) {
	tests := []string{
		"",
		"/DIR1",
		`\DIR1`,
		"DIR1/",
		"DIR1//IMAGE001",
		"DIR1/../IMAGE001",
		"DIR1/./IMAGE001",
		"dir1/IMAGE001",
		"DIR-1/IMAGE001",
		"DIRECTORY1/IMAGE001",
		"A/B/C/D/E/F/G/H/I",
		`C:\DIR1\IMAGE001`,
	}

	for _, path := range tests {
		t.Run(path, func(t *testing.T) {
			if _, err := ParseFileID(path); err == nil {
				t.Fatalf("ParseFileID(%q) succeeded, want error", path)
			}
		})
	}
}

func TestFileIDComponentsReturnsDefensiveCopy(t *testing.T) {
	id, err := NewFileID(testDirectoryID, testImageID)
	if err != nil {
		t.Fatalf("NewFileID() error = %v", err)
	}

	parts := id.Components()
	parts[0] = "CHANGED"

	if got := id.String(); got != `DIR1\IMAGE001` {
		t.Fatalf("String() after caller mutation = %q", got)
	}
}
