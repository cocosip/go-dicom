// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestGetByteRangeOverflowDoesNotPanic(t *testing.T) {
	tests := []struct {
		name string
		run  func() error
	}{
		{
			name: "MemoryByteBuffer",
			run: func() error {
				out := make([]byte, 2)
				return buffer.NewMemory([]byte{1, 2, 3, 4}).GetByteRange(^uint32(0)-1, 2, out)
			},
		},
		{
			name: "LazyByteBuffer",
			run: func() error {
				out := make([]byte, 2)
				lb, err := buffer.NewLazy(func() []byte { return []byte{1, 2, 3, 4} })
				if err != nil {
					return err
				}
				return lb.GetByteRange(^uint32(0)-1, 2, out)
			},
		},
		{
			name: "BulkDataURIByteBuffer",
			run: func() error {
				out := make([]byte, 2)
				b := buffer.NewBulkDataURIWithData("memory://test", []byte{1, 2, 3, 4})
				return b.GetByteRange(^uint32(0)-1, 2, out)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("GetByteRange panicked: %v", r)
				}
			}()
			if err := tc.run(); err == nil {
				t.Fatalf("expected bounds error, got nil")
			}
		})
	}
}

func TestNewFileLargeSparseFile(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "sparse.dcm")

	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("Create() error: %v", err)
	}
	defer func() { _ = f.Close() }()

	const size5GiB = int64(5) << 30
	if err := f.Truncate(size5GiB); err != nil {
		t.Skipf("sparse truncate unsupported on this filesystem: %v", err)
	}

	if _, err := buffer.NewFile(path, 2<<30, 1); err != nil {
		t.Fatalf("NewFile() unexpected error for valid range in sparse file: %v", err)
	}
}

func TestNewFileAtKeepsOffsetsAbove4GiB(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "sparse-large-offset.dcm")

	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("Create() error: %v", err)
	}
	defer func() { _ = f.Close() }()

	const offset = int64(5)<<30 + 123
	want := []byte{0xAA, 0xBB, 0xCC}
	if err := f.Truncate(offset + int64(len(want))); err != nil {
		t.Skipf("sparse truncate unsupported on this filesystem: %v", err)
	}
	if _, err := f.WriteAt(want, offset); err != nil {
		t.Fatalf("WriteAt() error: %v", err)
	}

	fb, err := buffer.NewFileAt(path, offset, uint32(len(want)))
	if err != nil {
		t.Fatalf("NewFileAt() error: %v", err)
	}
	if fb.Position64() != offset {
		t.Fatalf("Position64() = %d, want %d", fb.Position64(), offset)
	}

	got := make([]byte, len(want))
	if err := fb.GetByteRange(0, uint32(len(got)), got); err != nil {
		t.Fatalf("GetByteRange() error: %v", err)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("GetByteRange()[%d] = 0x%02X, want 0x%02X", i, got[i], want[i])
		}
	}
}
