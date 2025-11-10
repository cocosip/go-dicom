// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// FragmentSequence represents a DICOM fragment sequence used for encapsulated pixel data.
// This is typically used for compressed image formats like JPEG, JPEG 2000, JPEG-LS, and RLE.
//
// Structure:
// - First item: Offset Table (optional, can be empty)
// - Subsequent items: Compressed frame fragments
// - Terminated by Sequence Delimitation Item
type FragmentSequence struct {
	tag         *tag.Tag
	vr          *vr.VR
	offsetTable []uint32            // Offset table for multi-frame images
	fragments   []buffer.ByteBuffer // Fragment data buffers
}

// NewFragmentSequence creates a new fragment sequence element.
func NewFragmentSequence(t *tag.Tag, vrValue *vr.VR) *FragmentSequence {
	return &FragmentSequence{
		tag:         t,
		vr:          vrValue,
		offsetTable: []uint32{},
		fragments:   []buffer.ByteBuffer{},
	}
}

// Tag returns the DICOM tag.
func (f *FragmentSequence) Tag() *tag.Tag {
	return f.tag
}

// ValueRepresentation returns the VR.
func (f *FragmentSequence) ValueRepresentation() *vr.VR {
	return f.vr
}

// Buffer returns nil as fragment sequences don't have a single buffer.
func (f *FragmentSequence) Buffer() buffer.ByteBuffer {
	return nil
}

// Length returns the total length of all fragments.
func (f *FragmentSequence) Length() uint32 {
	var total uint32
	for _, frag := range f.fragments {
		total += frag.Size()
	}
	return total
}

// Validate checks if the fragment sequence is valid.
func (f *FragmentSequence) Validate() error {
	if f.tag == nil {
		return fmt.Errorf("fragment sequence has nil tag")
	}
	if f.vr != vr.OB && f.vr != vr.OW {
		return fmt.Errorf("fragment sequence must have VR of OB or OW, got %v", f.vr)
	}
	return nil
}

// OffsetTable returns the offset table.
// The offset table contains byte offsets to the first fragment of each frame.
// For single-frame images, this may be empty.
func (f *FragmentSequence) OffsetTable() []uint32 {
	return f.offsetTable
}

// SetOffsetTable sets the offset table.
func (f *FragmentSequence) SetOffsetTable(offsets []uint32) {
	f.offsetTable = offsets
}

// Fragments returns all fragment buffers.
func (f *FragmentSequence) Fragments() []buffer.ByteBuffer {
	return f.fragments
}

// AddFragment adds a fragment buffer to the sequence.
func (f *FragmentSequence) AddFragment(fragment buffer.ByteBuffer) {
	f.fragments = append(f.fragments, fragment)
}

// FragmentCount returns the number of fragments.
func (f *FragmentSequence) FragmentCount() int {
	return len(f.fragments)
}

// GetFragment returns the fragment at the specified index.
func (f *FragmentSequence) GetFragment(index int) (buffer.ByteBuffer, error) {
	if index < 0 || index >= len(f.fragments) {
		return nil, fmt.Errorf("fragment index %d out of range [0, %d)", index, len(f.fragments))
	}
	return f.fragments[index], nil
}

// Count returns the number of fragments (for VM compatibility).
func (f *FragmentSequence) Count() int {
	return len(f.fragments)
}

// String returns a string representation of the fragment sequence.
func (f *FragmentSequence) String() string {
	return fmt.Sprintf("FragmentSequence[%s, %d fragments, offset table: %d entries]",
		f.tag, len(f.fragments), len(f.offsetTable))
}

// OtherByteFragment represents a fragment sequence with VR = OB.
type OtherByteFragment struct {
	*FragmentSequence
}

// NewOtherByteFragment creates a new OB fragment sequence.
func NewOtherByteFragment(t *tag.Tag) *OtherByteFragment {
	return &OtherByteFragment{
		FragmentSequence: NewFragmentSequence(t, vr.OB),
	}
}

// OtherWordFragment represents a fragment sequence with VR = OW.
type OtherWordFragment struct {
	*FragmentSequence
}

// NewOtherWordFragment creates a new OW fragment sequence.
func NewOtherWordFragment(t *tag.Tag) *OtherWordFragment {
	return &OtherWordFragment{
		FragmentSequence: NewFragmentSequence(t, vr.OW),
	}
}
