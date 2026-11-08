// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package uid

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

// TestConvertUUIDToIntegerUsesNetworkOrder pins the 2.25 conversion to the
// numeric value of the RFC 4122 UUID (big-endian octets), per PS3.5 B.2.
func TestConvertUUIDToIntegerUsesNetworkOrder(t *testing.T) {
	u := uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	got := convertUUIDToInteger(u)
	want := "2.25.143098242404177361603877621312831893704"
	if got != want {
		t.Fatalf("convertUUIDToInteger() = %q, want %q", got, want)
	}
}

func TestAppendRejectsOverlongUID(t *testing.T) {
	// Build a 62-character base: appending ".100" would exceed the 64-char limit.
	base := New("1.2."+strings.Repeat("1", 58), "Base", TypeSOPInstance, false)
	if len(base.UID()) != 62 {
		t.Fatalf("base length = %d, want 62", len(base.UID()))
	}
	if got := Append(base, 100); got != nil {
		t.Fatalf("Append() = %q, want nil for overlong UID", got.UID())
	}
	if got := Append(base, 0); got == nil || got.UID() != base.UID()+".0" {
		t.Fatalf("Append(base, 0) = %v, want %q", got, base.UID()+".0")
	}
}
