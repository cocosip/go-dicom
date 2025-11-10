// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package element

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

func TestFragmentSequence(t *testing.T) {
	t.Run("NewFragmentSequence", func(t *testing.T) {
		fs := NewFragmentSequence(tag.PixelData, vr.OB)
		if fs == nil {
			t.Fatal("NewFragmentSequence returned nil")
		}
		if fs.Tag() != tag.PixelData {
			t.Errorf("Tag() = %v, want %v", fs.Tag(), tag.PixelData)
		}
		if fs.ValueRepresentation() != vr.OB {
			t.Errorf("ValueRepresentation() = %v, want %v", fs.ValueRepresentation(), vr.OB)
		}
		if fs.FragmentCount() != 0 {
			t.Errorf("FragmentCount() = %d, want 0", fs.FragmentCount())
		}
	})

	t.Run("OffsetTable", func(t *testing.T) {
		fs := NewFragmentSequence(tag.PixelData, vr.OB)

		// Initially empty
		if len(fs.OffsetTable()) != 0 {
			t.Errorf("OffsetTable length = %d, want 0", len(fs.OffsetTable()))
		}

		// Set offset table
		offsets := []uint32{0, 1024, 2048}
		fs.SetOffsetTable(offsets)

		if len(fs.OffsetTable()) != 3 {
			t.Errorf("OffsetTable length = %d, want 3", len(fs.OffsetTable()))
		}

		for i, offset := range offsets {
			if fs.OffsetTable()[i] != offset {
				t.Errorf("OffsetTable[%d] = %d, want %d", i, fs.OffsetTable()[i], offset)
			}
		}
	})

	t.Run("AddFragment", func(t *testing.T) {
		fs := NewFragmentSequence(tag.PixelData, vr.OB)

		// Add fragments
		frag1 := buffer.NewMemory([]byte{1, 2, 3, 4})
		frag2 := buffer.NewMemory([]byte{5, 6, 7, 8})

		fs.AddFragment(frag1)
		if fs.FragmentCount() != 1 {
			t.Errorf("FragmentCount() = %d, want 1", fs.FragmentCount())
		}

		fs.AddFragment(frag2)
		if fs.FragmentCount() != 2 {
			t.Errorf("FragmentCount() = %d, want 2", fs.FragmentCount())
		}
	})

	t.Run("GetFragment", func(t *testing.T) {
		fs := NewFragmentSequence(tag.PixelData, vr.OB)

		frag1 := buffer.NewMemory([]byte{1, 2, 3, 4})
		frag2 := buffer.NewMemory([]byte{5, 6, 7, 8})

		fs.AddFragment(frag1)
		fs.AddFragment(frag2)

		// Get valid fragments
		got1, err := fs.GetFragment(0)
		if err != nil {
			t.Errorf("GetFragment(0) error = %v", err)
		}
		if got1.Size() != 4 {
			t.Errorf("Fragment 0 size = %d, want 4", got1.Size())
		}

		got2, err := fs.GetFragment(1)
		if err != nil {
			t.Errorf("GetFragment(1) error = %v", err)
		}
		if got2.Size() != 4 {
			t.Errorf("Fragment 1 size = %d, want 4", got2.Size())
		}

		// Get invalid index
		_, err = fs.GetFragment(2)
		if err == nil {
			t.Error("GetFragment(2) should return error for out of range index")
		}

		_, err = fs.GetFragment(-1)
		if err == nil {
			t.Error("GetFragment(-1) should return error for negative index")
		}
	})

	t.Run("Fragments", func(t *testing.T) {
		fs := NewFragmentSequence(tag.PixelData, vr.OB)

		frag1 := buffer.NewMemory([]byte{1, 2, 3})
		frag2 := buffer.NewMemory([]byte{4, 5, 6})

		fs.AddFragment(frag1)
		fs.AddFragment(frag2)

		fragments := fs.Fragments()
		if len(fragments) != 2 {
			t.Errorf("Fragments() length = %d, want 2", len(fragments))
		}
	})
}

func TestOtherByteFragment(t *testing.T) {
	t.Run("NewOtherByteFragment", func(t *testing.T) {
		obf := NewOtherByteFragment(tag.PixelData)
		if obf == nil {
			t.Fatal("NewOtherByteFragment returned nil")
		}
		if obf.ValueRepresentation() != vr.OB {
			t.Errorf("ValueRepresentation() = %v, want OB", obf.ValueRepresentation())
		}
	})

	t.Run("AddFragment", func(t *testing.T) {
		obf := NewOtherByteFragment(tag.PixelData)
		frag := buffer.NewMemory([]byte{1, 2, 3, 4})
		obf.AddFragment(frag)

		if obf.FragmentCount() != 1 {
			t.Errorf("FragmentCount() = %d, want 1", obf.FragmentCount())
		}
	})
}

func TestOtherWordFragment(t *testing.T) {
	t.Run("NewOtherWordFragment", func(t *testing.T) {
		owf := NewOtherWordFragment(tag.PixelData)
		if owf == nil {
			t.Fatal("NewOtherWordFragment returned nil")
		}
		if owf.ValueRepresentation() != vr.OW {
			t.Errorf("ValueRepresentation() = %v, want OW", owf.ValueRepresentation())
		}
	})

	t.Run("AddFragment", func(t *testing.T) {
		owf := NewOtherWordFragment(tag.PixelData)
		frag := buffer.NewMemory([]byte{1, 2, 3, 4})
		owf.AddFragment(frag)

		if owf.FragmentCount() != 1 {
			t.Errorf("FragmentCount() = %d, want 1", owf.FragmentCount())
		}
	})
}
