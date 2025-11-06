// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package endian_test

import (
	"encoding/binary"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/endian"
)

func TestEndianString(t *testing.T) {
	tests := []struct {
		name string
		e    endian.Endian
		want string
	}{
		{"Little Endian", endian.Little, "Little Endian"},
		{"Big Endian", endian.Big, "Big Endian"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Endian.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestEndianIsBig(t *testing.T) {
	if endian.Big.IsBig() != true {
		t.Error("Big.IsBig() = false, want true")
	}
	if endian.Little.IsBig() != false {
		t.Error("Little.IsBig() = true, want false")
	}
}

func TestEndianIsLittle(t *testing.T) {
	if endian.Little.IsLittle() != true {
		t.Error("Little.IsLittle() = false, want true")
	}
	if endian.Big.IsLittle() != false {
		t.Error("Big.IsLittle() = true, want false")
	}
}

func TestEndianByteOrder(t *testing.T) {
	if endian.Little.ByteOrder() != binary.LittleEndian {
		t.Error("Little.ByteOrder() != binary.LittleEndian")
	}
	if endian.Big.ByteOrder() != binary.BigEndian {
		t.Error("Big.ByteOrder() != binary.BigEndian")
	}
}

func TestSwapUint16(t *testing.T) {
	tests := []struct {
		name  string
		value uint16
		want  uint16
	}{
		{"0x0102", 0x0102, 0x0201},
		{"0x1234", 0x1234, 0x3412},
		{"0xABCD", 0xABCD, 0xCDAB},
		{"0x0000", 0x0000, 0x0000},
		{"0xFFFF", 0xFFFF, 0xFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := endian.SwapUint16(tt.value); got != tt.want {
				t.Errorf("SwapUint16(0x%04X) = 0x%04X, want 0x%04X", tt.value, got, tt.want)
			}
		})
	}
}

func TestSwapUint32(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
		want  uint32
	}{
		{"0x01020304", 0x01020304, 0x04030201},
		{"0x12345678", 0x12345678, 0x78563412},
		{"0xABCDEF00", 0xABCDEF00, 0x00EFCDAB},
		{"0x00000000", 0x00000000, 0x00000000},
		{"0xFFFFFFFF", 0xFFFFFFFF, 0xFFFFFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := endian.SwapUint32(tt.value); got != tt.want {
				t.Errorf("SwapUint32(0x%08X) = 0x%08X, want 0x%08X", tt.value, got, tt.want)
			}
		})
	}
}

func TestSwapUint64(t *testing.T) {
	tests := []struct {
		name  string
		value uint64
		want  uint64
	}{
		{"0x0102030405060708", 0x0102030405060708, 0x0807060504030201},
		{"0x123456789ABCDEF0", 0x123456789ABCDEF0, 0xF0DEBC9A78563412},
		{"0x0000000000000000", 0x0000000000000000, 0x0000000000000000},
		{"0xFFFFFFFFFFFFFFFF", 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := endian.SwapUint64(tt.value); got != tt.want {
				t.Errorf("SwapUint64(0x%016X) = 0x%016X, want 0x%016X", tt.value, got, tt.want)
			}
		})
	}
}

func TestSwapInt16(t *testing.T) {
	value := int16(0x0102)
	want := int16(0x0201)
	if got := endian.SwapInt16(value); got != want {
		t.Errorf("SwapInt16(0x%04X) = 0x%04X, want 0x%04X", value, got, want)
	}
}

func TestSwapInt32(t *testing.T) {
	value := int32(0x01020304)
	want := int32(0x04030201)
	if got := endian.SwapInt32(value); got != want {
		t.Errorf("SwapInt32(0x%08X) = 0x%08X, want 0x%08X", value, got, want)
	}
}

func TestSwapInt64(t *testing.T) {
	value := int64(0x0102030405060708)
	want := int64(0x0807060504030201)
	if got := endian.SwapInt64(value); got != want {
		t.Errorf("SwapInt64(0x%016X) = 0x%016X, want 0x%016X", value, got, want)
	}
}

func TestSwapBytes2(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	want := []byte{0x02, 0x01, 0x04, 0x03, 0x06, 0x05}

	endian.SwapBytes(2, data)

	for i := range data {
		if data[i] != want[i] {
			t.Errorf("SwapBytes(2) failed: got %v, want %v", data, want)
			break
		}
	}
}

func TestSwapBytes4(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	want := []byte{0x04, 0x03, 0x02, 0x01, 0x08, 0x07, 0x06, 0x05}

	endian.SwapBytes(4, data)

	for i := range data {
		if data[i] != want[i] {
			t.Errorf("SwapBytes(4) failed: got %v, want %v", data, want)
			break
		}
	}
}

func TestSwapBytesN(t *testing.T) {
	// Test swapping only first 4 bytes
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	want := []byte{0x02, 0x01, 0x04, 0x03, 0x05, 0x06} // Only first 4 bytes swapped

	endian.SwapBytesN(2, data, 4)

	for i := range data {
		if data[i] != want[i] {
			t.Errorf("SwapBytesN(2, data, 4) failed: got %v, want %v", data, want)
			break
		}
	}
}

func TestSwapBytes1(t *testing.T) {
	// Swapping 1 byte at a time should not change anything
	data := []byte{0x01, 0x02, 0x03, 0x04}
	want := []byte{0x01, 0x02, 0x03, 0x04}

	endian.SwapBytes(1, data)

	for i := range data {
		if data[i] != want[i] {
			t.Errorf("SwapBytes(1) failed: got %v, want %v", data, want)
			break
		}
	}
}

func TestLocalMachineEndian(t *testing.T) {
	// Just verify it's one of the valid values
	if endian.LocalMachine != endian.Little && endian.LocalMachine != endian.Big {
		t.Errorf("LocalMachine has invalid value: %v", endian.LocalMachine)
	}

	// Verify it matches Go's understanding
	buf := [2]byte{0x01, 0x00}
	isLittle := binary.LittleEndian.Uint16(buf[:]) == 1

	if isLittle && endian.LocalMachine != endian.Little {
		t.Error("LocalMachine should be Little on little-endian machine")
	}
	if !isLittle && endian.LocalMachine != endian.Big {
		t.Error("LocalMachine should be Big on big-endian machine")
	}
}

func TestNetworkEndian(t *testing.T) {
	if endian.Network != endian.Big {
		t.Error("Network endian should always be Big")
	}
}

// Benchmark tests
func BenchmarkSwapUint16(b *testing.B) {
	value := uint16(0x1234)
	for i := 0; i < b.N; i++ {
		endian.SwapUint16(value)
	}
}

func BenchmarkSwapUint32(b *testing.B) {
	value := uint32(0x12345678)
	for i := 0; i < b.N; i++ {
		endian.SwapUint32(value)
	}
}

func BenchmarkSwapUint64(b *testing.B) {
	value := uint64(0x123456789ABCDEF0)
	for i := 0; i < b.N; i++ {
		endian.SwapUint64(value)
	}
}

func BenchmarkSwapBytes2(b *testing.B) {
	data := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		endian.SwapBytes(2, data)
	}
}

func BenchmarkSwapBytes4(b *testing.B) {
	data := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		endian.SwapBytes(4, data)
	}
}
