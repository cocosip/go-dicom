// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer_test

import (
	"bytes"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestNewLazy(t *testing.T) {
	loader := func() []byte {
		return []byte{1, 2, 3, 4}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	if lb == nil {
		t.Fatal("NewLazy() returned nil buffer")
	}
}

func TestNewLazy_NilLoader(t *testing.T) {
	_, err := buffer.NewLazy(nil)
	if err == nil {
		t.Error("NewLazy(nil) expected error, got nil")
	}
}

func TestLazyByteBuffer_IsMemory(t *testing.T) {
	loader := func() []byte {
		return []byte{1, 2, 3, 4}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	if !lb.IsMemory() {
		t.Error("IsMemory() = false, want true")
	}
}

func TestLazyByteBuffer_LazyLoading(t *testing.T) {
	loadCount := 0
	loader := func() []byte {
		loadCount++
		return []byte{1, 2, 3, 4, 5}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	// Initially, data should not be loaded
	if lb.IsLoaded() {
		t.Error("IsLoaded() = true, want false before first access")
	}

	if loadCount != 0 {
		t.Errorf("Loader called %d times before access, want 0", loadCount)
	}

	// First access should trigger loading
	data1 := lb.Data()
	if loadCount != 1 {
		t.Errorf("Loader called %d times after first access, want 1", loadCount)
	}

	if !lb.IsLoaded() {
		t.Error("IsLoaded() = false, want true after first access")
	}

	// Second access should not trigger loading again
	data2 := lb.Data()
	if loadCount != 1 {
		t.Errorf("Loader called %d times after second access, want 1 (should be cached)", loadCount)
	}

	// Data should be the same
	if !bytes.Equal(data1, data2) {
		t.Error("Data() returned different results on subsequent calls")
	}
}

func TestLazyByteBuffer_Size(t *testing.T) {
	loader := func() []byte {
		return []byte{1, 2, 3, 4, 5, 6, 7, 8}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	if lb.Size() != 8 {
		t.Errorf("Size() = %d, want 8", lb.Size())
	}

	// Should be loaded now
	if !lb.IsLoaded() {
		t.Error("IsLoaded() = false, want true after Size() call")
	}
}

func TestLazyByteBuffer_Size_NilData(t *testing.T) {
	loader := func() []byte {
		return nil
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	if lb.Size() != 0 {
		t.Errorf("Size() = %d, want 0 for nil data", lb.Size())
	}
}

func TestLazyByteBuffer_Data(t *testing.T) {
	want := []byte{1, 2, 3, 4, 5}
	loader := func() []byte {
		return want
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	data := lb.Data()
	if !bytes.Equal(data, want) {
		t.Errorf("Data() = %v, want %v", data, want)
	}
}

func TestLazyByteBuffer_GetByteRange(t *testing.T) {
	loader := func() []byte {
		return []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	tests := []struct {
		name      string
		offset    uint32
		count     uint32
		wantData  []byte
		wantError bool
	}{
		{
			name:     "full range",
			offset:   0,
			count:    10,
			wantData: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "first 5 bytes",
			offset:   0,
			count:    5,
			wantData: []byte{0, 1, 2, 3, 4},
		},
		{
			name:     "middle bytes",
			offset:   3,
			count:    4,
			wantData: []byte{3, 4, 5, 6},
		},
		{
			name:     "last bytes",
			offset:   8,
			count:    2,
			wantData: []byte{8, 9},
		},
		{
			name:      "exceeds buffer",
			offset:    8,
			count:     10,
			wantError: true,
		},
		{
			name:      "offset too large",
			offset:    20,
			count:     2,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := make([]byte, tt.count)
			err := lb.GetByteRange(tt.offset, tt.count, output)

			if tt.wantError {
				if err == nil {
					t.Error("GetByteRange() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("GetByteRange() unexpected error: %v", err)
			}

			if !bytes.Equal(output, tt.wantData) {
				t.Errorf("GetByteRange() = %v, want %v", output, tt.wantData)
			}
		})
	}
}

func TestLazyByteBuffer_GetByteRange_Errors(t *testing.T) {
	loader := func() []byte {
		return []byte{1, 2, 3, 4, 5}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	t.Run("nil output", func(t *testing.T) {
		err := lb.GetByteRange(0, 3, nil)
		if err == nil {
			t.Error("GetByteRange(nil) expected error, got nil")
		}
	})

	t.Run("output too small", func(t *testing.T) {
		output := make([]byte, 2)
		err := lb.GetByteRange(0, 5, output)
		if err == nil {
			t.Error("GetByteRange() with small output expected error, got nil")
		}
	})
}

func TestLazyByteBuffer_GetByteRange_NilData(t *testing.T) {
	loader := func() []byte {
		return nil
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	output := make([]byte, 5)
	err = lb.GetByteRange(0, 5, output)
	if err == nil {
		t.Error("GetByteRange() with nil data expected error, got nil")
	}
}

func TestLazyByteBuffer_WriteTo(t *testing.T) {
	loader := func() []byte {
		return []byte{1, 2, 3, 4, 5, 6, 7, 8}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := lb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != 8 {
		t.Errorf("WriteTo() wrote %d bytes, want 8", n)
	}

	want := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Errorf("WriteTo() wrote %v, want %v", buf.Bytes(), want)
	}
}

func TestLazyByteBuffer_WriteTo_Empty(t *testing.T) {
	loader := func() []byte {
		return nil
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	var buf bytes.Buffer
	n, err := lb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != 0 {
		t.Errorf("WriteTo() wrote %d bytes, want 0", n)
	}
}

func TestLazyByteBuffer_WriteTo_NilWriter(t *testing.T) {
	loader := func() []byte {
		return []byte{1, 2, 3}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	_, err = lb.WriteTo(nil)
	if err == nil {
		t.Error("WriteTo(nil) expected error, got nil")
	}
}

func TestLazyByteBuffer_ConcurrentAccess(t *testing.T) {
	loadCount := 0
	var mu sync.Mutex

	loader := func() []byte {
		mu.Lock()
		loadCount++
		mu.Unlock()
		return []byte{1, 2, 3, 4, 5}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	// Access the buffer concurrently
	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			_ = lb.Data()
		}()
	}

	wg.Wait()

	// Loader should only be called once despite concurrent access
	mu.Lock()
	count := loadCount
	mu.Unlock()

	if count != 1 {
		t.Errorf("Loader called %d times with concurrent access, want 1", count)
	}
}

func TestLazyByteBuffer_ExpensiveLoader(t *testing.T) {
	// Simulate an expensive operation
	loadCalled := false
	loader := func() []byte {
		loadCalled = true
		// Simulate expensive computation
		result := make([]byte, 1000)
		for i := range result {
			result[i] = byte(i % 256)
		}
		return result
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	// Creation should not trigger loading
	if loadCalled {
		t.Error("Loader was called during buffer creation")
	}

	// Access should trigger loading
	size := lb.Size()
	if !loadCalled {
		t.Error("Loader was not called on first access")
	}

	if size != 1000 {
		t.Errorf("Size() = %d, want 1000", size)
	}
}

func TestLazyByteBuffer_MultipleOperations(t *testing.T) {
	loader := func() []byte {
		return []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	lb, err := buffer.NewLazy(loader)
	if err != nil {
		t.Fatalf("NewLazy() unexpected error: %v", err)
	}

	// Perform various operations
	size := lb.Size()
	if size != 10 {
		t.Errorf("Size() = %d, want 10", size)
	}

	data := lb.Data()
	if len(data) != 10 {
		t.Errorf("len(Data()) = %d, want 10", len(data))
	}

	output := make([]byte, 5)
	err = lb.GetByteRange(2, 5, output)
	if err != nil {
		t.Fatalf("GetByteRange() unexpected error: %v", err)
	}

	want := []byte{2, 3, 4, 5, 6}
	if !bytes.Equal(output, want) {
		t.Errorf("GetByteRange() = %v, want %v", output, want)
	}

	var buf bytes.Buffer
	n, err := lb.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo() unexpected error: %v", err)
	}

	if n != 10 {
		t.Errorf("WriteTo() wrote %d bytes, want 10", n)
	}
}
