// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package parser

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// writeNestedSequence appends `depth` nested undefined-length SQ sequences, each
// containing a single undefined-length item. The innermost item holds a small
// PN element so the recursion has a terminating dataset.
func writeNestedSequence(buf *bytes.Buffer, depth int) {
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x0008))
	_ = binary.Write(buf, binary.LittleEndian, uint16(0x1110))
	buf.WriteString("SQ")
	_ = binary.Write(buf, binary.LittleEndian, uint16(0))
	_ = binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF)) // undefined sequence length

	_ = binary.Write(buf, binary.LittleEndian, uint16(0xFFFE)) // Item
	_ = binary.Write(buf, binary.LittleEndian, uint16(0xE000))
	_ = binary.Write(buf, binary.LittleEndian, uint32(0xFFFFFFFF)) // undefined item length

	if depth > 1 {
		writeNestedSequence(buf, depth-1)
	} else {
		writeExplicitStringElement(buf, tag.PatientName, "PN", []byte("Test"))
	}

	_ = binary.Write(buf, binary.LittleEndian, uint16(0xFFFE)) // Item Delimitation
	_ = binary.Write(buf, binary.LittleEndian, uint16(0xE00D))
	_ = binary.Write(buf, binary.LittleEndian, uint32(0))
	_ = binary.Write(buf, binary.LittleEndian, uint16(0xFFFE)) // Sequence Delimitation
	_ = binary.Write(buf, binary.LittleEndian, uint16(0xE0DD))
	_ = binary.Write(buf, binary.LittleEndian, uint32(0))
}

func TestParseRejectsExcessiveSequenceNesting(t *testing.T) {
	t.Run("within limit", func(t *testing.T) {
		file := createMiniDICOM()
		writeNestedSequence(file, defaultMaxSequenceDepth)
		if _, err := Parse(file); err != nil {
			t.Fatalf("Parse() error = %v, want nil at depth %d", err, defaultMaxSequenceDepth)
		}
	})

	t.Run("exceeds limit", func(t *testing.T) {
		file := createMiniDICOM()
		writeNestedSequence(file, defaultMaxSequenceDepth+1)
		if _, err := Parse(file); err == nil {
			t.Fatalf("Parse() error = nil, want depth-limit error at depth %d", defaultMaxSequenceDepth+1)
		}
	})

	t.Run("custom limit", func(t *testing.T) {
		limited := createMiniDICOM()
		writeNestedSequence(limited, 3)
		if _, err := Parse(limited, WithMaxSequenceDepth(2)); err == nil {
			t.Fatal("Parse() error = nil, want depth-limit error with WithMaxSequenceDepth(2)")
		}

		allowed := createMiniDICOM()
		writeNestedSequence(allowed, 3)
		if _, err := Parse(allowed, WithMaxSequenceDepth(3)); err != nil {
			t.Fatalf("Parse() error = %v, want nil at depth 3 with limit 3", err)
		}
	})
}

// writeUnknownVRElement appends an explicit-VR element whose VR code is not in
// the DICOM registry. Because the fallback is UN, the length is encoded as a
// 32-bit length (2 reserved bytes + uint32).
func writeUnknownVRElement(buf *bytes.Buffer, tg *tag.Tag, code string, value []byte) {
	_ = binary.Write(buf, binary.LittleEndian, tg.Group())
	_ = binary.Write(buf, binary.LittleEndian, tg.Element())
	buf.WriteString(code)
	_ = binary.Write(buf, binary.LittleEndian, uint16(0))
	_ = binary.Write(buf, binary.LittleEndian, uint32(len(value)))
	buf.Write(value)
}

func TestParseUnknownExplicitVRToleratedAsUN(t *testing.T) {
	target := tag.New(0x0008, 0x0001)
	file := createMiniDICOM()
	writeUnknownVRElement(file, target, "ZZ", []byte{1, 2, 3, 4})

	result, err := Parse(file)
	if err != nil {
		t.Fatalf("Parse() error = %v, want nil (tolerated unknown VR)", err)
	}
	elem, ok := result.Dataset.Get(target)
	if !ok {
		t.Fatal("unknown-VR element was dropped")
	}
	if elem.ValueRepresentation() != vr.UN {
		t.Fatalf("VR = %s, want UN", elem.ValueRepresentation())
	}
	if elem.Length() != 4 {
		t.Fatalf("length = %d, want 4", elem.Length())
	}

	strictFile := createMiniDICOM()
	writeUnknownVRElement(strictFile, target, "ZZ", []byte{1, 2, 3, 4})
	if _, err := Parse(strictFile, WithStrictVR(true)); err == nil {
		t.Fatal("Parse() error = nil, want error with WithStrictVR(true)")
	}
}

type callCountingReader struct {
	r     io.Reader
	reads int
}

func (c *callCountingReader) Read(p []byte) (int, error) {
	c.reads++
	return c.r.Read(p)
}

func TestParseNonSeekableReaderUsesBufferedReads(t *testing.T) {
	file := createMiniDICOM()
	// Add many small elements so an unbuffered parser would issue many Reads.
	for i := 0; i < 20; i++ {
		writeExplicitStringElement(file, tag.New(0x0008, uint16(0x0100+i)), "LO", []byte("v"))
	}

	cr := &callCountingReader{r: bytes.NewReader(file.Bytes())}
	if _, err := Parse(cr); err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	// With bufio the whole small stream is fetched in a handful of underlying
	// Read calls; without it each field read hits the reader (dozens of calls).
	if cr.reads > 8 {
		t.Fatalf("underlying Read calls = %d, want <= 8 (buffered)", cr.reads)
	}
}
