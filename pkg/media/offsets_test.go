// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package media

import "testing"

func TestRecoverOffsetDeltaRejectsAmbiguousCandidates(t *testing.T) {
	_, err := recoverOffsetDelta([]uint32{102}, []uint32{100, 104})
	if err == nil {
		t.Fatal("recoverOffsetDelta() succeeded, want ambiguous-candidate error")
	}
}

func TestRecoverOffsetDeltaFindsUniqueCommonDelta(t *testing.T) {
	delta, err := recoverOffsetDelta([]uint32{102, 202}, []uint32{100, 200})
	if err != nil {
		t.Fatalf("recoverOffsetDelta() error = %v", err)
	}
	if delta != 2 {
		t.Fatalf("delta = %d, want 2", delta)
	}
}
