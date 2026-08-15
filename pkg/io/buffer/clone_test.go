// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer

import (
	"errors"
	"os"
	"testing"
)

func TestClonePreservesUnresolvedBulkDataURI(t *testing.T) {
	source := NewBulkDataURI("https://example.invalid/pixel-data")

	clonedBuffer, err := Clone(source)
	if err != nil {
		t.Fatal(err)
	}
	clone, ok := clonedBuffer.(*BulkDataURIByteBuffer)
	if !ok || clone == source {
		t.Fatalf("Clone() = %T %p, want independent BulkDataURIByteBuffer", clonedBuffer, clonedBuffer)
	}
	if clone.BulkDataURI() != source.BulkDataURI() || clone.IsMemory() {
		t.Fatalf("cloned BulkDataURI = %q, IsMemory = %t", clone.BulkDataURI(), clone.IsMemory())
	}
	clone.SetData([]byte{1, 2})
	if source.IsMemory() {
		t.Fatal("setting cloned bulk data changed unresolved source state")
	}
}

func TestCloneReportsUnavailableFileBuffer(t *testing.T) {
	path := t.TempDir() + string(os.PathSeparator) + "value.bin"
	if err := os.WriteFile(path, []byte{1, 2, 3, 4}, 0o600); err != nil {
		t.Fatal(err)
	}
	source, err := NewFile(path, 0, 4)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}

	clone, err := Clone(source)
	if clone != nil || err == nil || !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Clone() = %v, %v, want file-not-found error", clone, err)
	}
}
