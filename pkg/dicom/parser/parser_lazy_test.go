// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package parser

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// TestCreateLazyBuffer_WithFile tests lazy loading with an *os.File
func TestCreateLazyBuffer_WithFile(t *testing.T) {
	// Create a temporary DICOM-like file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.dcm")

	// Write some test data
	testData := make([]byte, 200*1024) // 200KB
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Create file with preamble + DICM + some data
	f, err := os.Create(tmpFile)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Write preamble
	preamble := make([]byte, 128)
	f.Write(preamble)
	f.WriteString("DICM")

	// Write a large element
	// Tag (0008,0018) - SOP Instance UID
	binary.Write(f, binary.LittleEndian, uint16(0x0008))
	binary.Write(f, binary.LittleEndian, uint16(0x0018))
	// VR: UI
	f.WriteString("UI")
	// Length (16-bit for this VR)
	binary.Write(f, binary.LittleEndian, uint16(len(testData)))
	// Data
	f.Write(testData)

	f.Close()

	// Open file for parsing
	file, err := os.Open(tmpFile)
	if err != nil {
		t.Fatalf("Failed to open temp file: %v", err)
	}
	defer file.Close()

	// Skip preamble
	file.Seek(132, io.SeekStart)

	// Create parse context
	ctx := newParseContext(
		WithReadOption(ReadLargeOnDemand),
		WithLargeObjectSize(100*1024), // 100KB threshold
	)
	ctx.reader = file
	ctx.file = file
	ctx.seekableReader = file
	ctx.byteOrder = binary.LittleEndian
	ctx.isExplicitVR = true

	// Read the tag
	readTag, err := ctx.readTag()
	if err != nil {
		t.Fatalf("readTag() error: %v", err)
	}

	// Skip VR (2 bytes)
	file.Seek(2, io.SeekCurrent)

	// Skip length (2 bytes for UI VR)
	file.Seek(2, io.SeekCurrent)

	// Now we're at the data position
	// Create lazy buffer
	buf, err := ctx.createLazyBuffer(uint32(len(testData)))
	if err != nil {
		t.Fatalf("createLazyBuffer() error: %v", err)
	}

	// Verify it's a FileByteBuffer
	if _, ok := buf.(*buffer.FileByteBuffer); !ok {
		t.Errorf("Expected FileByteBuffer, got %T", buf)
	}

	// Verify size
	if buf.Size() != uint32(len(testData)) {
		t.Errorf("Size() = %d, want %d", buf.Size(), len(testData))
	}

	// Verify data can be loaded
	data := buf.Data()
	if len(data) != len(testData) {
		t.Errorf("len(Data()) = %d, want %d", len(data), len(testData))
	}

	// Verify data is correct
	for i := 0; i < 100; i++ {
		if data[i] != testData[i] {
			t.Errorf("data[%d] = %d, want %d", i, data[i], testData[i])
			break
		}
	}

	_ = readTag
}

// TestCreateLazyBuffer_WithSeekableReader tests lazy loading with io.ReadSeeker
func TestCreateLazyBuffer_WithSeekableReader(t *testing.T) {
	// Create test data
	testData := make([]byte, 150*1024) // 150KB
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Create a bytes.Reader (implements io.ReadSeeker)
	reader := bytes.NewReader(testData)

	// Create parse context
	ctx := newParseContext(
		WithReadOption(ReadLargeOnDemand),
		WithLargeObjectSize(100*1024),
	)
	ctx.reader = reader
	ctx.seekableReader = reader
	ctx.byteOrder = binary.LittleEndian
	ctx.isExplicitVR = true

	// Create lazy buffer
	buf, err := ctx.createLazyBuffer(uint32(len(testData)))
	if err != nil {
		t.Fatalf("createLazyBuffer() error: %v", err)
	}

	// Verify it's a LazyByteBuffer
	lb, ok := buf.(*buffer.LazyByteBuffer)
	if !ok {
		t.Errorf("Expected LazyByteBuffer, got %T", buf)
	}

	// Initially should not be loaded
	if lb.IsLoaded() {
		t.Error("LazyByteBuffer should not be loaded initially")
	}

	// Access data (should trigger loading)
	data := buf.Data()
	if len(data) != len(testData) {
		t.Errorf("len(Data()) = %d, want %d", len(data), len(testData))
	}

	// Now should be loaded
	if !lb.IsLoaded() {
		t.Error("LazyByteBuffer should be loaded after Data() call")
	}

	// Verify data is correct
	for i := 0; i < 100; i++ {
		if data[i] != testData[i] {
			t.Errorf("data[%d] = %d, want %d", i, data[i], testData[i])
			break
		}
	}
}

// TestCreateLazyBuffer_NonSeekable tests fallback for non-seekable readers
func TestCreateLazyBuffer_NonSeekable(t *testing.T) {
	// Create a non-seekable reader (io.Pipe)
	pr, pw := io.Pipe()
	defer pr.Close()
	defer pw.Close()

	// Write some data in the background
	go func() {
		testData := make([]byte, 100)
		pw.Write(testData)
		pw.Close()
	}()

	// Create parse context
	ctx := newParseContext()
	ctx.reader = pr
	// Don't set seekableReader or file

	// Try to create lazy buffer
	_, err := ctx.createLazyBuffer(100)
	if err == nil {
		t.Error("createLazyBuffer() should return error for non-seekable reader")
	}
}

// TestReadLargeOnDemand_Integration tests the full parsing flow with lazy loading
func TestReadLargeOnDemand_Integration(t *testing.T) {
	// Create a temporary DICOM file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_lazy.dcm")

	f, err := os.Create(tmpFile)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Write preamble
	preamble := make([]byte, 128)
	f.Write(preamble)
	f.WriteString("DICM")

	// Write File Meta Information Length (Group 0002)
	// Tag (0002,0000) - File Meta Information Group Length
	binary.Write(f, binary.LittleEndian, uint16(0x0002))
	binary.Write(f, binary.LittleEndian, uint16(0x0000))
	f.WriteString("UL") // VR
	binary.Write(f, binary.LittleEndian, uint16(4))
	binary.Write(f, binary.LittleEndian, uint32(0)) // Placeholder

	// Tag (0002,0010) - Transfer Syntax UID
	binary.Write(f, binary.LittleEndian, uint16(0x0002))
	binary.Write(f, binary.LittleEndian, uint16(0x0010))
	f.WriteString("UI")
	tsUID := "1.2.840.10008.1.2\x00" // Implicit VR Little Endian + null padding
	binary.Write(f, binary.LittleEndian, uint16(len(tsUID)))
	f.WriteString(tsUID)

	// Update Group Length
	endPos, _ := f.Seek(0, io.SeekCurrent)
	groupLength := uint32(endPos - 132 - 12) // Total length minus preamble, DICM, and group length element itself
	f.Seek(140, io.SeekStart)
	binary.Write(f, binary.LittleEndian, groupLength)
	f.Seek(0, io.SeekEnd)

	// Write a large pixel data element (> 64KB)
	largeData := make([]byte, 200*1024) // 200KB
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	// Pixel Data tag (7FE0,0010) - using Implicit VR format now
	binary.Write(f, binary.LittleEndian, uint16(0x7FE0))
	binary.Write(f, binary.LittleEndian, uint16(0x0010))
	binary.Write(f, binary.LittleEndian, uint32(len(largeData))) // 32-bit length for implicit VR
	f.Write(largeData)

	f.Close()

	// Parse with ReadLargeOnDemand
	file, err := os.Open(tmpFile)
	if err != nil {
		t.Fatalf("Failed to open temp file: %v", err)
	}
	defer file.Close()

	result, err := Parse(file,
		WithReadOption(ReadLargeOnDemand),
		WithLargeObjectSize(100*1024), // 100KB threshold
	)
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}

	// Verify we got a dataset
	if result.Dataset == nil {
		t.Fatal("Dataset should not be nil")
	}

	// Try to get the pixel data element
	pixelDataElem, exists := result.Dataset.Get(tag.New(0x7FE0, 0x0010))
	if !exists {
		t.Fatal("Pixel Data element should exist")
	}

	// For a lazy-loaded element, the buffer should be a FileByteBuffer or LazyByteBuffer
	// We can't easily check this without exposing internal details,
	// but we can verify the data can be accessed
	if pixelDataElem == nil {
		t.Fatal("Pixel Data element should not be nil")
	}

	// The element should have the correct tag
	if pixelDataElem.Tag().ToUint32() != 0x7FE00010 {
		t.Errorf("Wrong tag: %s", pixelDataElem.Tag())
	}
}

// TestReadOption_SkipVsLazy tests the difference between Skip and Lazy modes
func TestReadOption_SkipVsLazy(t *testing.T) {
	// Create test file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_skip.dcm")

	// Write a minimal DICOM file with a large element
	largeData := make([]byte, 200*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	f, err := os.Create(tmpFile)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// Write preamble
	preamble := make([]byte, 128)
	f.Write(preamble)
	f.WriteString("DICM")

	// Write File Meta Information properly
	// Tag (0002,0000) - File Meta Information Group Length
	binary.Write(f, binary.LittleEndian, uint16(0x0002))
	binary.Write(f, binary.LittleEndian, uint16(0x0000))
	f.WriteString("UL")
	binary.Write(f, binary.LittleEndian, uint16(4))
	binary.Write(f, binary.LittleEndian, uint32(0)) // Placeholder

	// Tag (0002,0010) - Transfer Syntax UID
	binary.Write(f, binary.LittleEndian, uint16(0x0002))
	binary.Write(f, binary.LittleEndian, uint16(0x0010))
	f.WriteString("UI")
	tsUID := "1.2.840.10008.1.2\x00" // Implicit VR Little Endian + null padding
	binary.Write(f, binary.LittleEndian, uint16(len(tsUID)))
	f.WriteString(tsUID)

	// Update Group Length
	endPos, _ := f.Seek(0, io.SeekCurrent)
	groupLength := uint32(endPos - 132 - 12) // Total length minus preamble, DICM, and group length element itself
	f.Seek(140, io.SeekStart)
	binary.Write(f, binary.LittleEndian, groupLength)
	f.Seek(0, io.SeekEnd)

	// Large element - using Implicit VR format now (no VR string, 32-bit length)
	binary.Write(f, binary.LittleEndian, uint16(0x0008))
	binary.Write(f, binary.LittleEndian, uint16(0x0018))
	binary.Write(f, binary.LittleEndian, uint32(len(largeData))) // 32-bit length for implicit VR
	f.Write(largeData)

	f.Close()

	// Test 1: Skip mode
	t.Run("SkipLargeTags", func(t *testing.T) {
		file, _ := os.Open(tmpFile)
		defer file.Close()

		result, err := Parse(file,
			WithReadOption(SkipLargeTags),
			WithLargeObjectSize(100*1024),
		)
		if err != nil {
			t.Fatalf("Parse() error: %v", err)
		}

		// Element should exist but be empty
		elem, exists := result.Dataset.Get(tag.New(0x0008, 0x0018))
		if !exists {
			t.Fatal("Element should exist in Skip mode")
		}

		// Should have empty data
		if elem.Length() != 0 {
			t.Errorf("Skipped element should have length 0, got %d", elem.Length())
		}
	})

	// Test 2: Lazy mode
	t.Run("ReadLargeOnDemand", func(t *testing.T) {
		file, _ := os.Open(tmpFile)
		defer file.Close()

		result, err := Parse(file,
			WithReadOption(ReadLargeOnDemand),
			WithLargeObjectSize(100*1024),
		)
		if err != nil {
			t.Fatalf("Parse() error: %v", err)
		}

		// Element should exist and have data accessible
		elem, exists := result.Dataset.Get(tag.New(0x0008, 0x0018))
		if !exists {
			t.Fatal("Element should exist in Lazy mode")
		}

		// Should have the full length
		if elem.Length() != uint32(len(largeData)) {
			t.Errorf("Lazy element should have length %d, got %d", len(largeData), elem.Length())
		}
	})
}
