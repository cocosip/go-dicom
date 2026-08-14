// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dicomxml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadReadsDictionaryTagsAndUIDs(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dictionary.xml")
	const source = `<dictionary version="2026b">
  <tag group="0010" element="0010" keyword="PatientName" vr="PN" vm="1">Patient Name</tag>
  <uid uid="1.2.840.10008.1.1" keyword="Verification" type="SOP Class">Verification SOP Class</uid>
</dictionary>`
	if err := os.WriteFile(path, []byte(source), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	dictionary, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if dictionary.Version != "2026b" {
		t.Errorf("Version = %q, want 2026b", dictionary.Version)
	}
	if len(dictionary.Tags) != 1 || dictionary.Tags[0].Keyword != "PatientName" {
		t.Fatalf("Tags = %#v, want PatientName", dictionary.Tags)
	}
	if len(dictionary.UIDs) != 1 || dictionary.UIDs[0].Keyword != "Verification" {
		t.Fatalf("UIDs = %#v, want Verification", dictionary.UIDs)
	}
}
