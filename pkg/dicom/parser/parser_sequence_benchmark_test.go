// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

//revive:disable:var-naming // package name must match public import path (pkg/dicom/parser)
package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func buildSequenceBenchmarkDICOM(definedLength bool, itemCount int, payloadSize int) []byte {
	buf := bytes.NewBuffer(nil)

	// Preamble + DICM.
	buf.Write(make([]byte, 128))
	buf.WriteString("DICM")

	// File Meta: Transfer Syntax UID (Explicit VR Little Endian).
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0002))
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	buf.WriteString("UI")
	tsUID := testExplicitVRLittleLE
	_ = binary.Write(buf, binary.LittleEndian, uint16(len(tsUID)))
	buf.WriteString(tsUID)

	// Build sequence content.
	seqData := bytes.NewBuffer(nil)
	payload := bytes.Repeat([]byte{0xAB}, payloadSize)

	for i := 0; i < itemCount; i++ {
		itemData := bytes.NewBuffer(nil)

		// (0008,1155) UI - Referenced SOP Instance UID
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x0008))
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x1155))
		itemData.WriteString("UI")
		sopUID := "1.2.840.10008.5.1.4.1.1.2"
		_ = binary.Write(itemData, binary.LittleEndian, uint16(len(sopUID)))
		itemData.WriteString(sopUID)

		// (7FE1,0010) OB - large payload to stress sequence/item parsing.
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x7FE1))
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0x0010))
		itemData.WriteString("OB")
		_ = binary.Write(itemData, binary.LittleEndian, uint16(0))
		_ = binary.Write(itemData, binary.LittleEndian, uint32(len(payload)))
		_, _ = itemData.Write(payload)

		// Item header (FFFE,E000).
		_ = binary.Write(seqData, binary.LittleEndian, uint16(0xFFFE))
		_ = binary.Write(seqData, binary.LittleEndian, uint16(0xE000))
		if definedLength {
			_ = binary.Write(seqData, binary.LittleEndian, uint32(itemData.Len()))
			seqData.Write(itemData.Bytes())
		} else {
			_ = binary.Write(seqData, binary.LittleEndian, uint32(0xFFFFFFFF))
			seqData.Write(itemData.Bytes())
			_ = binary.Write(seqData, binary.LittleEndian, uint16(0xFFFE))
			_ = binary.Write(seqData, binary.LittleEndian, uint16(0xE00D))
			_ = binary.Write(seqData, binary.LittleEndian, uint32(0))
		}
	}

	// Sequence element (0008,1140) SQ.
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0008))
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x1140))
	buf.WriteString("SQ")
	_ = binary.Write(buf, binary.LittleEndian, uint16(0))
	if definedLength {
		_ = binary.Write(buf, binary.LittleEndian, uint32(seqData.Len()))
		buf.Write(seqData.Bytes())
	} else {
		_ = binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF))
		buf.Write(seqData.Bytes())
		_ = binary.Write(buf, binary.LittleEndian, uint16(0xFFFE))
		_ = binary.Write(buf, binary.LittleEndian, uint16(0xE0DD))
		_ = binary.Write(buf, binary.LittleEndian, uint32(0))
	}

	// Trailing element to ensure stream alignment.
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0010))
	buf.WriteString("PN")
	pn := "Benchmark^Tail"
	_ = binary.Write(buf, binary.LittleEndian, uint16(len(pn)))
	buf.WriteString(pn)

	return buf.Bytes()
}

func BenchmarkParseSequenceLayouts(b *testing.B) {
	type benchCase struct {
		name         string
		defined      bool
		itemCount    int
		payloadBytes int
	}

	cases := []benchCase{
		{name: "defined_1item_256KB", defined: true, itemCount: 1, payloadBytes: 256 * 1024},
		{name: "defined_8items_32KB", defined: true, itemCount: 8, payloadBytes: 32 * 1024},
		{name: "undefined_1item_256KB", defined: false, itemCount: 1, payloadBytes: 256 * 1024},
		{name: "undefined_8items_32KB", defined: false, itemCount: 8, payloadBytes: 32 * 1024},
	}

	for _, tc := range cases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			data := buildSequenceBenchmarkDICOM(tc.defined, tc.itemCount, tc.payloadBytes)
			b.SetBytes(int64(len(data)))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				reader := bytes.NewReader(data)
				if _, err := Parse(reader); err != nil {
					b.Fatalf("Parse() error: %v", err)
				}
			}
		})
	}
}

func BenchmarkParseDefinedLengthSequenceReadModes(b *testing.B) {
	data := buildSequenceBenchmarkDICOM(true, 4, 128*1024)

	cases := []struct {
		name string
		opt  ReadOption
	}{
		{name: "read_default", opt: ReadDefault},
		{name: "read_large_on_demand", opt: ReadLargeOnDemand},
		{name: "skip_large_tags", opt: SkipLargeTags},
	}

	for _, tc := range cases {
		tc := tc
		b.Run(fmt.Sprintf("defined_4items_128KB_%s", tc.name), func(b *testing.B) {
			b.SetBytes(int64(len(data)))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				reader := bytes.NewReader(data)
				if _, err := Parse(reader,
					WithReadOption(tc.opt),
					WithLargeObjectSize(64*1024),
				); err != nil {
					b.Fatalf("Parse() error: %v", err)
				}
			}
		})
	}
}
