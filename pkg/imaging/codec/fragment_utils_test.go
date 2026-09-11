// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"bytes"
	"testing"
)

func TestStripTrailingPaddingOnlyRemovesByteAfterEOI(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want []byte
	}{
		{name: "jpeg padding", data: []byte{0x11, 0xff, 0xd9, 0x00}, want: []byte{0x11, 0xff, 0xd9}},
		{name: "zero before eoi is data", data: []byte{0x11, 0x00, 0xff, 0xd9}, want: []byte{0x11, 0x00, 0xff, 0xd9}},
		{name: "post eoi data is preserved", data: []byte{0xff, 0xd9, 0x01, 0x00}, want: []byte{0xff, 0xd9, 0x01, 0x00}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripTrailingPadding(tt.data); !bytes.Equal(got, tt.want) {
				t.Fatalf("StripTrailingPadding() = %x, want %x", got, tt.want)
			}
		})
	}
}
