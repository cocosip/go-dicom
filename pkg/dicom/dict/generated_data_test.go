// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

//revive:disable:var-naming // package name must match public import path (pkg/dicom/dict)
package dict_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dict"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
)

func TestDefaultContainsGenerated2026bStandardAndPrivateDictionaries(t *testing.T) {
	dictionary := dict.Default()
	if got := len(dictionary.Entries()); got != 5347 {
		t.Errorf("default standard dictionary entries = %d, want 5347", got)
	}

	creator := dictionary.GetPrivateCreator("MED NM")
	entry := dictionary.Lookup(tag.NewWithPrivateCreator(0x0011, 0x1010, creator))
	if entry == nil || entry.Name() != "Unknown" {
		t.Fatalf("default MED NM (0011,xx10) entry = %#v, want Unknown", entry)
	}
}
