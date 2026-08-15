// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package main

import (
	"context"
	"strings"
	"testing"
)

func TestRunRejectsInvalidPlaneBeforeReadingInputs(t *testing.T) {
	err := run(context.Background(), []string{"does-not-exist.dcm"}, t.TempDir(), "diagonal", 1, 1, 1, false)
	if err == nil || !strings.Contains(err.Error(), "unknown plane") {
		t.Fatalf("run() error = %v, want unknown plane", err)
	}
}
