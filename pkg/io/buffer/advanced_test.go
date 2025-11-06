// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package buffer_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// TestCompositeByteBuffer tests the composite buffer
func TestCompositeByteBuffer(t *testing.T) {
	buf1 := buffer.NewMemory([]byte{1, 2, 3})
	buf2 := buffer.NewMemory([]byte{4, 5, 6})
	buf3 := buffer.NewMemory([]byte{7, 8, 9, 10})

	composite := buffer.NewComposite(buf1, buf2, buf3)

	t.Run("Size", func(t *testing.T) {
		if composite.Size() != 10 {
			t.Errorf("Size() = %d, want 10", composite.Size())
		}
	})

	t.Run("IsMemory", func(t *testing.T) {
		if !composite.IsMemory() {
			t.Error("IsMemory() should be true when all buffers are in memory")
		}
	})

	t.Run("Data", func(t *testing.T) {
		data := composite.Data()
		want := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		if !bytes.Equal(data, want) {
			t.Errorf("Data() = %v, want %v", data, want)
		}
	})

	t.Run("GetByteRange spanning buffers", func(t *testing.T) {
		// Get bytes that span across multiple buffers
		output := make([]byte, 5)
		err := composite.GetByteRange(2, 5, output)
		if err != nil {
			t.Fatalf("GetByteRange() error = %v", err)
		}
		want := []byte{3, 4, 5, 6, 7}
		if !bytes.Equal(output, want) {
			t.Errorf("GetByteRange() = %v, want %v", output, want)
		}
	})

	t.Run("WriteTo", func(t *testing.T) {
		var w bytes.Buffer
		n, err := composite.WriteTo(&w)
		if err != nil {
			t.Fatalf("WriteTo() error = %v", err)
		}
		if n != 10 {
			t.Errorf("WriteTo() wrote %d bytes, want 10", n)
		}
		want := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		if !bytes.Equal(w.Bytes(), want) {
			t.Errorf("WriteTo() wrote %v, want %v", w.Bytes(), want)
		}
	})
}

// TestStreamByteBuffer tests the stream buffer
func TestStreamByteBuffer(t *testing.T) {
	data := []byte("Hello, DICOM world!")
	reader := bytes.NewReader(data)

	stream, err := buffer.NewStream(reader, 7, 5) // "DICOM"
	if err != nil {
		t.Fatalf("NewStream() error = %v", err)
	}

	t.Run("IsMemory", func(t *testing.T) {
		if stream.IsMemory() {
			t.Error("IsMemory() should be false for stream buffer")
		}
	})

	t.Run("Size", func(t *testing.T) {
		if stream.Size() != 5 {
			t.Errorf("Size() = %d, want 5", stream.Size())
		}
	})

	t.Run("Data", func(t *testing.T) {
		got := stream.Data()
		want := []byte("DICOM")
		if !bytes.Equal(got, want) {
			t.Errorf("Data() = %q, want %q", got, want)
		}
	})

	t.Run("GetByteRange", func(t *testing.T) {
		output := make([]byte, 3)
		err := stream.GetByteRange(1, 3, output) // "ICO"
		if err != nil {
			t.Fatalf("GetByteRange() error = %v", err)
		}
		want := []byte("ICO")
		if !bytes.Equal(output, want) {
			t.Errorf("GetByteRange() = %q, want %q", output, want)
		}
	})
}

// TestEndianByteBuffer tests the endian buffer
func TestEndianByteBuffer(t *testing.T) {
	// Create buffer with uint16 values in little endian: [0x0201, 0x0403]
	data := []byte{0x01, 0x02, 0x03, 0x04}
	mem := buffer.NewMemory(data)

	// If local machine is little endian, swapping to big endian should reverse bytes
	endianBuf := buffer.NewEndian(mem, endian.Big, 2)

	t.Run("No wrapping for matching endian", func(t *testing.T) {
		sameEndian := buffer.NewEndian(mem, endian.LocalMachine, 2)
		// Should return the original buffer
		if sameEndian != mem {
			t.Error("NewEndian should return original buffer when endianness matches")
		}
	})

	t.Run("No wrapping for unitSize=1", func(t *testing.T) {
		byteBuf := buffer.NewEndian(mem, endian.Big, 1)
		// Should return the original buffer
		if byteBuf != mem {
			t.Error("NewEndian should return original buffer when unitSize=1")
		}
	})

	t.Run("Delegates IsMemory", func(t *testing.T) {
		if !endianBuf.IsMemory() {
			t.Error("EndianByteBuffer should delegate IsMemory to internal buffer")
		}
	})

	t.Run("Delegates Size", func(t *testing.T) {
		if endianBuf.Size() != 4 {
			t.Errorf("Size() = %d, want 4", endianBuf.Size())
		}
	})
}

// TestEvenLengthBuffer tests the even length buffer
func TestEvenLengthBuffer(t *testing.T) {
	t.Run("Odd length buffer", func(t *testing.T) {
		odd := buffer.NewMemory([]byte{1, 2, 3})
		even := buffer.NewEvenLength(odd)

		if even.Size() != 4 {
			t.Errorf("Size() = %d, want 4", even.Size())
		}

		data := even.Data()
		want := []byte{1, 2, 3, 0}
		if !bytes.Equal(data, want) {
			t.Errorf("Data() = %v, want %v", data, want)
		}
	})

	t.Run("Even length buffer unchanged", func(t *testing.T) {
		evenBuf := buffer.NewMemory([]byte{1, 2, 3, 4})
		wrapped := buffer.NewEvenLength(evenBuf)

		// Should return the original buffer without wrapping
		if wrapped != evenBuf {
			t.Error("NewEvenLength should return original buffer when already even")
		}
	})

	t.Run("GetByteRange with padding", func(t *testing.T) {
		odd := buffer.NewMemory([]byte{1, 2, 3})
		even := buffer.NewEvenLength(odd)

		// Read the last 2 bytes (includes padding)
		output := make([]byte, 2)
		err := even.GetByteRange(2, 2, output)
		if err != nil {
			t.Fatalf("GetByteRange() error = %v", err)
		}
		want := []byte{3, 0}
		if !bytes.Equal(output, want) {
			t.Errorf("GetByteRange() = %v, want %v", output, want)
		}
	})

	t.Run("WriteTo with padding", func(t *testing.T) {
		odd := buffer.NewMemory([]byte{1, 2, 3})
		even := buffer.NewEvenLength(odd)

		var w bytes.Buffer
		n, err := even.WriteTo(&w)
		if err != nil {
			t.Fatalf("WriteTo() error = %v", err)
		}
		if n != 4 {
			t.Errorf("WriteTo() wrote %d bytes, want 4", n)
		}
		want := []byte{1, 2, 3, 0}
		if !bytes.Equal(w.Bytes(), want) {
			t.Errorf("WriteTo() wrote %v, want %v", w.Bytes(), want)
		}
	})
}

// TestRangeByteBuffer tests the range buffer
func TestRangeByteBuffer(t *testing.T) {
	// Create a buffer with data: "Hello, DICOM world!"
	data := []byte("Hello, DICOM world!")
	mem := buffer.NewMemory(data)

	// Create a range that extracts "DICOM" (offset 7, length 5)
	rangeBuf, err := buffer.NewRange(mem, 7, 5)
	if err != nil {
		t.Fatalf("NewRange() error = %v", err)
	}

	t.Run("Size", func(t *testing.T) {
		if rangeBuf.Size() != 5 {
			t.Errorf("Size() = %d, want 5", rangeBuf.Size())
		}
	})

	t.Run("Data", func(t *testing.T) {
		got := rangeBuf.Data()
		want := []byte("DICOM")
		if !bytes.Equal(got, want) {
			t.Errorf("Data() = %q, want %q", got, want)
		}
	})

	t.Run("GetByteRange within range", func(t *testing.T) {
		output := make([]byte, 3)
		err := rangeBuf.GetByteRange(1, 3, output) // "ICO"
		if err != nil {
			t.Fatalf("GetByteRange() error = %v", err)
		}
		want := []byte("ICO")
		if !bytes.Equal(output, want) {
			t.Errorf("GetByteRange() = %q, want %q", output, want)
		}
	})

	t.Run("Error on invalid range", func(t *testing.T) {
		_, err := buffer.NewRange(mem, 15, 10) // Exceeds buffer size
		if err == nil {
			t.Error("NewRange should error when range exceeds buffer size")
		}
	})

	t.Run("WriteTo", func(t *testing.T) {
		var w bytes.Buffer
		n, err := rangeBuf.WriteTo(&w)
		if err != nil {
			t.Fatalf("WriteTo() error = %v", err)
		}
		if n != 5 {
			t.Errorf("WriteTo() wrote %d bytes, want 5", n)
		}
		want := []byte("DICOM")
		if !bytes.Equal(w.Bytes(), want) {
			t.Errorf("WriteTo() wrote %q, want %q", w.Bytes(), want)
		}
	})
}

// TestRangeByteBuffer_LargeData tests range buffer with chunked writing
func TestRangeByteBuffer_LargeData(t *testing.T) {
	// Create a large buffer (2MB)
	size := 2 * 1024 * 1024
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(i % 256)
	}
	mem := buffer.NewMemory(data)

	// Create a range of 1.5MB
	rangeSize := uint32(1536 * 1024)
	rangeBuf, err := buffer.NewRange(mem, 100, rangeSize)
	if err != nil {
		t.Fatalf("NewRange() error = %v", err)
	}

	var w bytes.Buffer
	n, err := rangeBuf.WriteTo(&w)
	if err != nil {
		t.Fatalf("WriteTo() error = %v", err)
	}
	if n != int64(rangeSize) {
		t.Errorf("WriteTo() wrote %d bytes, want %d", n, rangeSize)
	}
	if uint32(w.Len()) != rangeSize {
		t.Errorf("Buffer length = %d, want %d", w.Len(), rangeSize)
	}
}

// TestCompositeWithRangeBuffers tests combining range buffers
func TestCompositeWithRangeBuffers(t *testing.T) {
	data := []byte("ABCDEFGHIJKLMNOP")
	mem := buffer.NewMemory(data)

	// Create ranges: "ABC", "GHI", "MNO"
	range1, _ := buffer.NewRange(mem, 0, 3)
	range2, _ := buffer.NewRange(mem, 6, 3)
	range3, _ := buffer.NewRange(mem, 12, 3)

	composite := buffer.NewComposite(range1, range2, range3)

	got := composite.Data()
	want := []byte("ABCGHIMNO")
	if !bytes.Equal(got, want) {
		t.Errorf("Composite of ranges = %q, want %q", got, want)
	}
}

// TestStreamByteBuffer_Concurrent tests concurrent access to stream buffer
func TestStreamByteBuffer_Concurrent(t *testing.T) {
	data := []byte(strings.Repeat("DICOM", 1000))
	reader := bytes.NewReader(data)
	stream, _ := buffer.NewStream(reader, 0, uint32(len(data)))

	// Multiple goroutines reading different parts
	done := make(chan bool, 3)
	for i := 0; i < 3; i++ {
		go func(offset uint32) {
			output := make([]byte, 100)
			err := stream.GetByteRange(offset, 100, output)
			if err != nil {
				t.Errorf("GetByteRange() error = %v", err)
			}
			done <- true
		}(uint32(i * 100))
	}

	// Wait for all goroutines
	for i := 0; i < 3; i++ {
		<-done
	}
}
